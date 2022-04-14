// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/v1/download/game": {
            "get": {
                "description": "get a game, return a zip",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "summary": "get a game",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "the game's id",
                        "name": "game_id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/v1/download/img": {
            "get": {
                "description": "get a game's image",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "image/png"
                ],
                "summary": "get a game's image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the image name",
                        "name": "img_name",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/v1/download/personal": {
            "get": {
                "description": "get a person's image",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "image/png"
                ],
                "summary": "get a person's image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/v1/notoken/login": {
            "post": {
                "description": "using password, email and nickname to create a new account",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "log in a account",
                "parameters": [
                    {
                        "description": "email",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":200, \"msg\":\"register successfully\":, token\":\"string\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/notoken/register": {
            "post": {
                "description": "using password, email and nickname to create a new account",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "register a new account",
                "parameters": [
                    {
                        "description": "email",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "nickname",
                        "name": "nickname",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "person image",
                        "name": "file_body",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":200, \"msg\":\"register successfully\":, token\":\"string\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "cannot save answer",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "can not generate token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "406": {
                        "description": "email, nickname or password is missing",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/rank/download": {
            "get": {
                "description": "get game according to the download number in the same zone",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get game according to the download number in the same zone",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "the number of the return itme",
                        "name": "num",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "the zone",
                        "name": "zone",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":200, \"List\":list}",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.RankItem"
                            }
                        }
                    }
                }
            }
        },
        "/v1/rank/zone": {
            "get": {
                "description": "get game according to the like number in the same zone",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get game according to the like number in the same zone",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "the number of the return itme",
                        "name": "num",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "the zone",
                        "name": "zone",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":200, \"List\":list}",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.RankItem"
                            }
                        }
                    }
                }
            }
        },
        "/v1/search/game": {
            "get": {
                "description": "search a game by keyword",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "search a game by keyword",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "the number of the return item",
                        "name": "num",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the keyword",
                        "name": "keyword",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the order method",
                        "name": "method",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "the offset",
                        "name": "offset",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":200, \"List\":list}",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.RankItem"
                            }
                        }
                    }
                }
            }
        },
        "/v1/user/change/detail": {
            "post": {
                "description": "get a user's detail",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get a user's detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":200, \"detail\":detail}",
                        "schema": {
                            "$ref": "#/definitions/dto.PersonDetail"
                        }
                    }
                }
            }
        },
        "/v1/user/detail": {
            "post": {
                "description": "get a user's detail",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get a user's detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "the attribute to be modified",
                        "name": "FieldKey",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "the attribute value to be modified",
                        "name": "FieldVal",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":200, \"msg\": \"set successfully\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.PersonDetail": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "birth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                }
            }
        },
        "dto.RankItem": {
            "type": "object",
            "properties": {
                "GID": {
                    "type": "integer"
                },
                "authorName": {
                    "type": "string"
                },
                "commentNum": {
                    "type": "integer"
                },
                "downloadNum": {
                    "type": "integer"
                },
                "game_name": {
                    "type": "string"
                },
                "img": {
                    "type": "string"
                },
                "like_num": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8084",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "FLIP backend API",
	Description:      "FLIP backend server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
