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
        "termsOfService": "https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones",
        "contact": {
            "name": "API Support Digital House",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/products": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Show list products",
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
                        "description": "List products",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "Not found products",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error ",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
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
                    "Products"
                ],
                "summary": "Store product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token requerido",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Product to Store",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.ProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Update product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "description": "Product to update",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.ProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "Products"
                ],
                "summary": "Delete product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "This endpoint update field name from an product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Update name product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token header",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Product Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Product name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.ProductRequestPatch"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.ProductRequest": {
            "type": "object",
            "required": [
                "cantidad",
                "nombre",
                "precio",
                "tipo"
            ],
            "properties": {
                "cantidad": {
                    "type": "integer"
                },
                "nombre": {
                    "type": "string"
                },
                "precio": {
                    "type": "number"
                },
                "tipo": {
                    "type": "string"
                }
            }
        },
        "handler.ProductRequestPatch": {
            "type": "object",
            "required": [
                "nombre"
            ],
            "properties": {
                "nombre": {
                    "type": "string"
                }
            }
        },
        "web.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Bootcamp Go Wave 6 - API",
	Description:      "This is a simple API development by Digital House's team.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
