// Code generated by swaggo/swag. DO NOT EDIT.
// Edited: Made multi-lines for Failures responses

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
        "/api/cars/available": {
            "get": {
                "description": "Get a list of available cars by options(param)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Available Cars",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Date in the format of yyyy-mm-dd",
                        "name": "receiver_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Date in the format of yyyy-mm-dd",
                        "name": "delivery_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Time in the format of hh.mm",
                        "name": "time_start(hour)",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Time in the format of hh.mm",
                        "name": "time_end(hour)",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Location by id",
                        "name": "location",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Cars"
                            }
                        }
                    },
                    "400": {
                        "description": "missing required parameters, error occurred while converting data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "failed to find data, an error occurred while retrieving data from the database, no active offices available, no data found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/cars/reserve": {
            "post": {
                "description": "Reserve a car by ID and user information",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Reserve a car",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Car ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "User information",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "successfully reserved car {id}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "missing required parameters, error occurred while converting data, empty request body, error occurred while decoding JSON",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "failed to update data, the car is already reserved",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/cars/showreservedcars": {
            "get": {
                "description": "Retrieve the list of reserved cars from the database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Show reserved cars",
                "operationId": "ShowReservedCars",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Cars"
                            }
                        }
                    },
                    "500": {
                        "description": "failed to find data, an error occurred while retrieving data from the database, no data found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/locations/show": {
            "get": {
                "description": "Retrieve a list of active locations",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "locations"
                ],
                "summary": "Show active locations",
                "operationId": "ShowLocations",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Locations"
                            }
                        }
                    },
                    "500": {
                        "description": "failed to find data, an error occurred while retrieving data from the database, no data found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Cars": {
            "type": "object",
            "properties": {
                "fuel": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "office_id": {
                    "type": "integer"
                },
                "reserved": {
                    "type": "boolean"
                },
                "reserved_by": {
                    "$ref": "#/definitions/models.User"
                },
                "transmission": {
                    "type": "string"
                },
                "vendor": {
                    "type": "string"
                }
            }
        },
        "models.Locations": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "integer"
                },
                "active": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "nation_id": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "surname": {
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
