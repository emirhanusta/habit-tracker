// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "emirhan usta",
            "email": "emirhan1usta@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/habit": {
            "put": {
                "description": "update habit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Habit"
                ],
                "summary": "This method update habit",
                "parameters": [
                    {
                        "description": "Handle Request Body",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.HabitUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "save habit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Habit"
                ],
                "summary": "This method save habit",
                "parameters": [
                    {
                        "description": "habit",
                        "name": "habit",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.HabitCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/v1/habit-tracker/user/email": {
            "get": {
                "description": "get user by email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "This method get user by given email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/v1/habit/user/{userId}": {
            "get": {
                "description": "get all habits by given user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Habit"
                ],
                "summary": "This method get all habits by given user id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userId",
                        "name": "userId",
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
                                "$ref": "#/definitions/response.HabitResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/v1/habit/{id}": {
            "get": {
                "description": "get habit by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Habit"
                ],
                "summary": "This method get habit by given id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.HabitResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "delete habit by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Habit"
                ],
                "summary": "This method delete habit by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/v1/reminder": {
            "put": {
                "description": "update reminder",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reminder"
                ],
                "summary": "This method update reminder",
                "parameters": [
                    {
                        "description": "Handle Request Body",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ReminderUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "save reminder",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reminder"
                ],
                "summary": "This method save reminder",
                "parameters": [
                    {
                        "description": "reminder",
                        "name": "reminder",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ReminderCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/v1/reminder/habit/{habitId}": {
            "get": {
                "description": "get all reminders by given habit id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reminder"
                ],
                "summary": "This method get all reminders by given habit id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "habitId",
                        "name": "habitId",
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
                                "$ref": "#/definitions/response.ReminderResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/v1/reminder/{id}": {
            "get": {
                "description": "get reminder by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reminder"
                ],
                "summary": "This method get reminder by given id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ReminderResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "delete reminder by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reminder"
                ],
                "summary": "This method delete reminder by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/v1/user": {
            "get": {
                "description": "get all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "This method get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.UserResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "put": {
                "description": "Update user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "description": "Handle Request Body",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "saving new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "This method used for saving new user",
                "parameters": [
                    {
                        "description": "Handle Request Body",
                        "name": "requestBody",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.UserCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/v1/user/{id}": {
            "get": {
                "description": "get user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "This method get user by given id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "Delete user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "request.HabitCreateRequest": {
            "type": "object",
            "required": [
                "name",
                "userId"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "request.HabitUpdateRequest": {
            "type": "object",
            "required": [
                "description",
                "id",
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "request.ReminderCreateRequest": {
            "type": "object",
            "properties": {
                "habitId": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "remindAt": {
                    "type": "string"
                }
            }
        },
        "request.ReminderUpdateRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "remindAt": {
                    "type": "string"
                }
            }
        },
        "request.UserCreateRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.UserUpdateRequest": {
            "type": "object",
            "required": [
                "email",
                "id",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.HabitResponse": {
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
                "userId": {
                    "type": "string"
                }
            }
        },
        "response.ReminderResponse": {
            "type": "object",
            "properties": {
                "habitId": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "remindAt": {
                    "type": "string"
                }
            }
        },
        "response.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Habit Tracker Fiber REST API",
	Description:      "This is a sample swagger for Habit Tracker Fiber REST API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
