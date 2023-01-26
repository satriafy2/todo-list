// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/todo": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Get all todo list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Todo"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Create to-do list",
                "parameters": [
                    {
                        "description": "Create Todo Body, include parent_id field to create sub todo list",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateTodoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Todo"
                        }
                    }
                }
            }
        },
        "/todo/{id}": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Get single todo",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Todo"
                            }
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Delete to-do",
                "responses": {
                    "202": {
                        "description": "Accepted"
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
                    "Todo"
                ],
                "summary": "Update to-do",
                "parameters": [
                    {
                        "description": "Update todo request body, id field is required",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UpdateTodoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Todo"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "handlers.CreateTodoRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "parent_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "handlers.UpdateTodoRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "parent_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.Todo": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "description": {
                    "type": "string"
                },
                "file_url": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "parent_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:1323",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "To-Do List API",
	Description:      "This is a to-do list mini project.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
