package repository

import (
	"api/config"
	"api/types"
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	conn *sql.DB
}

func NewSQLite() (*SQLite, error) {
	conf := config.Get()
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?_foreign_keys=on", conf.Database.Address))
	if err != nil {
		return nil, err
	}

	// Set the maximum number of open connections
	db.SetMaxOpenConns(conf.Database.MaxConn)

	// Ping to check if the database connection is established
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	repo := &SQLite{
		conn: db,
	}

	err = repo.migrate("database/schema.sql")
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (s *SQLite) migrate(filepath string) error {
	// Read the schema file
	schema, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	// Execute the SQL statements from the schema file
	_, err = s.conn.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLite) CreateRetrospective(ctx context.Context, retro *types.Retrospective) error {
	sql := `INSERT INTO retrospectives (id, name, description) VALUES ($1, $2, $3)`
	_, err := s.conn.Exec(sql,
		retro.ID,
		retro.Name,
		retro.Description,
	)
	return err
}

func (s *SQLite) UpdateRetrospective(ctx context.Context, retro *types.Retrospective) error {
	foundRetro := &types.Retrospective{
		ID: retro.ID,
	}

	sqlQuery := `SELECT name, description FROM retrospectives WHERE id = $1`
	err := s.conn.QueryRow(sqlQuery, foundRetro.ID).Scan(
		&foundRetro.Name,
		&foundRetro.Description,
	)
	if err != nil {
		return err
	}

	if len(retro.Name) == 0 {
		retro.Name = foundRetro.Name
	}

	if len(retro.Description) == 0 {
		retro.Description = foundRetro.Description
	}

	sqlQuery = `UPDATE retrospectives SET name = $1, description = $2 WHERE id = $3`
	_, err = s.conn.Exec(sqlQuery,
		retro.Name,
		retro.Description,
		retro.ID,
	)

	return err
}

func (s *SQLite) DeleteRetrospective(ctx context.Context, id uuid.UUID) (*types.Retrospective, error) {
	retro := &types.Retrospective{
		ID: id,
	}

	sqlQuery := `SELECT name, description FROM retrospectives WHERE id = $1`
	err := s.conn.QueryRow(sqlQuery, id).Scan(
		&retro.Name,
		&retro.Description,
	)
	if err != nil {
		return nil, err
	}

	tx, err := s.conn.Begin()
	if err != nil {
		return retro, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Delete answers associated with questions of the retrospective
	sqlQuery = `DELETE FROM answers WHERE question_id IN (SELECT id FROM questions WHERE retrospective_id = $1)`
	_, err = tx.Exec(sqlQuery, id)
	if err != nil {
		return retro, err
	}

	// Delete questions associated with the retrospective
	sqlQuery = `DELETE FROM questions WHERE retrospective_id = $1`
	_, err = tx.Exec(sqlQuery, id)
	if err != nil {
		return retro, err
	}

	// Delete the retrospective
	sqlQuery = `DELETE FROM retrospectives WHERE id = $1`
	_, err = tx.Exec(sqlQuery, id)
	if err != nil {
		return retro, err
	}

	return retro, nil
}

func (s *SQLite) GetAllRetrospectives(ctx context.Context) ([]uuid.UUID, error) {
	sqlQuery := `SELECT id FROM retrospectives`
	rows, err := s.conn.Query(sqlQuery)
	if err != nil {
		return nil, err
	}

	IDs := make([]uuid.UUID, 0)

	for rows.Next() {
		var ID uuid.UUID
		err := rows.Scan(&ID)
		if err != nil {
			return nil, err
		}
		IDs = append(IDs, ID)
	}
	return IDs, nil
}

func (s *SQLite) GetRetrospective(ctx context.Context, id uuid.UUID) (*types.Retrospective, error) {
	retro := &types.Retrospective{
		ID: id,
	}

	sqlQuery := `SELECT name, description FROM retrospectives WHERE id = $1`
	err := s.conn.QueryRow(sqlQuery, id).Scan(
		&retro.Name,
		&retro.Description,
	)
	if err != nil {
		return nil, err
	}

	// Query questions for the retrospective
	sqlQuery = `SELECT id, text FROM questions WHERE retrospective_id = $1`
	rows, err := s.conn.Query(sqlQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var question types.Question
		err := rows.Scan(
			&question.ID,
			&question.Text,
		)
		if err != nil {
			return nil, err
		}

		// Query answers for the question
		sqlQuery = `SELECT id, text, position FROM answers WHERE question_id = $1`
		answerRows, err := s.conn.Query(sqlQuery, question.ID)
		if err != nil {
			return nil, err
		}
		defer answerRows.Close()

		// Loop through answers and append to the question
		for answerRows.Next() {
			var answer types.Answer
			err := answerRows.Scan(
				&answer.ID,
				&answer.Text,
				&answer.Position,
			)
			if err != nil {
				return nil, err
			}
			question.Answers = append(question.Answers, answer)
		}

		// Append the question to the retrospective
		retro.Questions = append(retro.Questions, question)
	}

	return retro, nil
}

func (s *SQLite) CreateQuestion(ctx context.Context, question *types.Question) error {
	retrospectiveID, ok := ctx.Value("retrospective_id").(uuid.UUID)
	if !ok {
		return fmt.Errorf("retrospective id not found")
	}
	sql := `INSERT INTO questions (id, text, retrospective_id) VALUES ($1, $2, $3)`
	_, err := s.conn.Exec(sql,
		question.ID,
		question.Text,
		retrospectiveID,
	)
	return err
}

func (s *SQLite) UpdateQuestion(ctx context.Context, question *types.Question) error {
	retrospectiveID, ok := ctx.Value("retrospective_id").(uuid.UUID)
	if !ok {
		return fmt.Errorf("retrospective id not found")
	}

	foundQuestion := &types.Question{
		ID: question.ID,
	}

	sqlQuery := `SELECT text FROM questions WHERE id = $1 and retrospective_id = $2`
	err := s.conn.QueryRow(sqlQuery, foundQuestion.ID, retrospectiveID).Scan(
		&foundQuestion.Text,
	)
	if err != nil {
		return err
	}

	if len(question.Text) == 0 {
		question.Text = foundQuestion.Text
	}

	sqlQuery = `UPDATE questions SET text = $1 WHERE id = $2 and retrospective_id = $3`
	_, err = s.conn.Exec(sqlQuery,
		question.Text,
		question.ID,
		retrospectiveID,
	)

	return err
}

func (s *SQLite) DeleteQuestion(ctx context.Context, id uuid.UUID) (*types.Question, error) {
	retrospectiveID, ok := ctx.Value("retrospective_id").(uuid.UUID)
	if !ok {
		return nil, fmt.Errorf("retrospective id not found")
	}
	question := &types.Question{
		ID: id,
	}

	sqlQuery := `SELECT text FROM questions WHERE id = $1 and retrospective_id = $2`
	err := s.conn.QueryRow(sqlQuery, id, retrospectiveID).Scan(
		&question.Text,
	)
	if err != nil {
		return nil, err
	}

	tx, err := s.conn.Begin()
	if err != nil {
		return question, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Delete answers associated with questions of the retrospective
	sqlQuery = `DELETE FROM answers WHERE question_id = $1`
	_, err = tx.Exec(sqlQuery, id)
	if err != nil {
		return question, err
	}

	// Delete questions associated with the retrospective
	sqlQuery = `DELETE FROM questions WHERE id = $1`
	_, err = tx.Exec(sqlQuery, id)

	return question, nil
}

func (s *SQLite) CreateAnswer(ctx context.Context, answer *types.Answer) error {
	sqlQuery := `INSERT INTO answers 
								(id, text, question_id, position) 
								VALUES ($1, $2, $3, (SELECT IFNULL(MAX(position),0) + 1 FROM answers WHERE question_id = $3)) returning position`
	err := s.conn.QueryRow(sqlQuery,
		answer.ID,
		answer.Text,
		answer.QuestionID,
	).Scan(
		&answer.Position,
	)
	return err
}

func (s *SQLite) UpdateAnswer(ctx context.Context, answer *types.Answer) error {
	foundAnswer := &types.Answer{
		ID:         answer.ID,
		QuestionID: answer.QuestionID,
	}

	sqlQuery := `SELECT text, position FROM answers WHERE id = $1 and question_id = $2`
	err := s.conn.QueryRow(sqlQuery,
		foundAnswer.ID,
		foundAnswer.QuestionID,
	).Scan(
		&foundAnswer.Text,
		&foundAnswer.Position,
	)
	if err != nil {
		return err
	}

	if len(answer.Text) == 0 {
		answer.Text = foundAnswer.Text
	}

	sqlQuery = `UPDATE answers SET text = $1 WHERE id = $2 and question_id = $3`
	_, err = s.conn.Exec(sqlQuery,
		answer.Text,
		answer.ID,
		answer.QuestionID,
	)

	return err
}

func (s *SQLite) DeleteAnswer(ctx context.Context, id uuid.UUID) (*types.Answer, error) {
	answer := &types.Answer{
		ID: id,
	}

	sqlQuery := `SELECT text, position, question_id FROM answers WHERE id = $1`
	err := s.conn.QueryRow(sqlQuery, id).Scan(
		&answer.Text,
		&answer.Position,
		&answer.QuestionID,
	)
	if err != nil {
		return nil, err
	}

	tx, err := s.conn.Begin()
	if err != nil {
		return answer, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	sqlQuery = `DELETE FROM answers WHERE id = $1`
	_, err = tx.Exec(sqlQuery, id)
	if err != nil {
		return answer, err
	}

	return answer, nil
}
