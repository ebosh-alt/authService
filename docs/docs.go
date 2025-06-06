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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/currency/add": {
            "post": {
                "description": "add currency pair",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add Currency",
                "operationId": "add-currency",
                "parameters": [
                    {
                        "description": "CurrencyAddRequest: currencypair",
                        "name": "CurrencyAddRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_protov1.CurrencyAddRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/auth_protov1.CurrencyAddResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/server.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/server.errorResponse"
                        }
                    }
                }
            }
        },
        "/currency/price": {
            "post": {
                "description": "get currency price at specific time",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Currency Price",
                "operationId": "get-currency-price",
                "parameters": [
                    {
                        "description": "CurrencyPriceRequest: currencyPair, timestamp",
                        "name": "CurrencyPriceRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_protov1.CurrencyPriceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/auth_protov1.CurrencyPriceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/server.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/server.errorResponse"
                        }
                    }
                }
            }
        },
        "/currency/remove": {
            "post": {
                "description": "remove currency pair",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Remove Currency",
                "operationId": "remove-currency",
                "parameters": [
                    {
                        "description": "CurrencyRemoveRequest: currencypair",
                        "name": "CurrencyRemoveRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_protov1.CurrencyRemoveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/auth_protov1.CurrencyRemoveResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/server.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/server.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth_protov1.CurrencyAddRequest": {
            "type": "object",
            "properties": {
                "currencyPair": {
                    "type": "string"
                }
            }
        },
        "auth_protov1.CurrencyAddResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "auth_protov1.CurrencyPriceRequest": {
            "type": "object",
            "properties": {
                "currencyPair": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "integer"
                }
            }
        },
        "auth_protov1.CurrencyPriceResponse": {
            "type": "object",
            "properties": {
                "price": {
                    "type": "string"
                }
            }
        },
        "auth_protov1.CurrencyRemoveRequest": {
            "type": "object",
            "properties": {
                "currencyPair": {
                    "type": "string"
                }
            }
        },
        "auth_protov1.CurrencyRemoveResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "server.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
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
	//LeftDelim:        "{{",
	//RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
