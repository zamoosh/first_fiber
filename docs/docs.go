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
        "/": {
            "get": {
                "description": "Root endpoint with ali ali ali",
                "tags": [
                    "default"
                ],
                "summary": "Get Root docs",
                "responses": {
                    "200": {
                        "description": "hello root!",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/value/{value}": {
            "get": {
                "description": "Get value by parameter",
                "tags": [
                    "default"
                ],
                "summary": "Get Value",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Value",
                        "name": "value",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "value is: \u003cvalue\u003e",
                        "schema": {
                            "type": "string"
                        }
                    }
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
