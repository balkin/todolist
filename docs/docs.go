// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-03-06 11:38:35.1061606 +0300 MSK m=+0.078011501

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "localhost:8000",
    "basePath": "/api/v1",
    "paths": {
        "/todo/count": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Count todo items without a parent",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.CountStruct"
                        }
                    }
                }
            }
        },
        "/todo/countall": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Count todo items without a parent",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.CountStruct"
                        }
                    }
                }
            }
        },
        "/todo/item/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "List todo items",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/todo.TodoItem"
                            }
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Add todo item (root)",
                "parameters": [
                    {
                        "description": "Add TodoItem",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.SimpleTodoItem"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.TodoItem"
                        }
                    }
                }
            }
        },
        "/todo/item/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Show todo item including subitems",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "TodoItem ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/todo.TodoItem"
                            }
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Add todo sub item",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "TodoItem ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Add sub TodoItem",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.SimpleTodoItem"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.TodoItem"
                        }
                    }
                }
            },
            "delete": {
                "summary": "Delete todo item including subitems",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "TodoItem ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        }
    },
    "definitions": {
        "controllers.CountStruct": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                }
            }
        },
        "todo.SimpleTodoItem": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "todo.TodoItem": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "parentId": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
