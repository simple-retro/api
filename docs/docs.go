// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://github.com/simple-retro/api/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/answer": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Answer"
                ],
                "summary": "Create Answer",
                "parameters": [
                    {
                        "description": "Create Answer",
                        "name": "question",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.AnswerCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retrospective Object",
                        "schema": {
                            "$ref": "#/definitions/types.Answer"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/answer/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Answer"
                ],
                "summary": "Delete Answer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Answer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Answer Object",
                        "schema": {
                            "$ref": "#/definitions/types.Answer"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Answer"
                ],
                "summary": "Update Answer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Answer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Answer",
                        "name": "answer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.AnswerCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Answer Object",
                        "schema": {
                            "$ref": "#/definitions/types.Answer"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Healthcheck"
                ],
                "summary": "Show API health",
                "responses": {
                    "200": {
                        "description": "API metrics",
                        "schema": {
                            "$ref": "#/definitions/server.health"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/hello": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Websocket"
                ],
                "summary": "Subscribe to changes via web socket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Repository ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/question": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Create Question",
                "parameters": [
                    {
                        "description": "Create Question",
                        "name": "question",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.QuestionCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retrospective Object",
                        "schema": {
                            "$ref": "#/definitions/types.Question"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/question/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Delete Question by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Question ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Question Object",
                        "schema": {
                            "$ref": "#/definitions/types.Question"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Update Question by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Question ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Question",
                        "name": "retrospective",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.QuestionCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Question Object",
                        "schema": {
                            "$ref": "#/definitions/types.Retrospective"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/retrospective": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Retrospective"
                ],
                "summary": "Create Retrospective",
                "parameters": [
                    {
                        "description": "Create Retrospective",
                        "name": "retrospective",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.RetrospectiveCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retrospective Object",
                        "schema": {
                            "$ref": "#/definitions/types.Retrospective"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/retrospective/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Retrospective"
                ],
                "summary": "Get Retrospective by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Retrospective ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retrospective Object",
                        "schema": {
                            "$ref": "#/definitions/types.Retrospective"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Retrospective"
                ],
                "summary": "Delete Retrospective by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Retrospective ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retrospective Object",
                        "schema": {
                            "$ref": "#/definitions/types.Retrospective"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Retrospective"
                ],
                "summary": "Update Retrospective by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Retrospective ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Retrospective",
                        "name": "retrospective",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.RetrospectiveCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retrospective Object",
                        "schema": {
                            "$ref": "#/definitions/types.Retrospective"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "server.health": {
            "type": "object",
            "properties": {
                "cpu": {
                    "type": "number"
                },
                "memory": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "types.Answer": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "position": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "types.AnswerCreateRequest": {
            "type": "object",
            "properties": {
                "question_id": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "types.Question": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Answer"
                    }
                },
                "id": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "types.QuestionCreateRequest": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string"
                }
            }
        },
        "types.Retrospective": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "questions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Question"
                    }
                }
            }
        },
        "types.RetrospectiveCreateRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
