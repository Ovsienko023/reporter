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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/profile": {
            "get": {
                "description": "Getting user data",
                "tags": [
                    "Profile"
                ],
                "summary": "Get Profile",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.GetProfileRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.GetProfileResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/reports": {
            "get": {
                "description": "get all reports",
                "tags": [
                    "Reports"
                ],
                "summary": "Get all reports",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.GetReportsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.GetReportsResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create report",
                "tags": [
                    "Reports"
                ],
                "summary": "Create report",
                "parameters": [
                    {
                        "description": "body params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CreateReportRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.CreateReportResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/reports/{report_id}": {
            "get": {
                "description": "get report",
                "tags": [
                    "Reports"
                ],
                "summary": "Get report",
                "parameters": [
                    {
                        "type": "string",
                        "description": "report_id",
                        "name": "report_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.GetReportResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httperror.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update report",
                "tags": [
                    "Reports"
                ],
                "summary": "Update report",
                "parameters": [
                    {
                        "type": "string",
                        "description": "report_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UpdateReportRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "delete": {
                "description": "Delete report",
                "tags": [
                    "Reports"
                ],
                "summary": "Delete report",
                "parameters": [
                    {
                        "type": "string",
                        "description": "report_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/api/v1/sign_in": {
            "post": {
                "description": "Getting an authorization token",
                "tags": [
                    "Auth"
                ],
                "summary": "Sign In",
                "parameters": [
                    {
                        "description": "body params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SignInResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
        "/api/v1/sign_up": {
            "post": {
                "description": "User registration",
                "tags": [
                    "Auth"
                ],
                "summary": "Sign Up",
                "parameters": [
                    {
                        "description": "body params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.CreateReportRequest": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "break_time": {
                    "type": "integer"
                },
                "date": {
                    "type": "integer"
                },
                "end_time": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "work_time": {
                    "type": "integer"
                }
            }
        },
        "domain.CreateReportResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "domain.GetProfileRequest": {
            "type": "object"
        },
        "domain.GetProfileResponse": {
            "type": "object",
            "properties": {
                "display_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                }
            }
        },
        "domain.GetReportResponse": {
            "type": "object",
            "properties": {
                "report": {
                    "$ref": "#/definitions/domain.Report"
                }
            }
        },
        "domain.GetReportsRequest": {
            "type": "object"
        },
        "domain.GetReportsResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "reports": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.ReportItem"
                    }
                }
            }
        },
        "domain.Report": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "break_time": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "integer"
                },
                "creator_id": {
                    "type": "string"
                },
                "date": {
                    "type": "integer"
                },
                "end_time": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "start_time": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "integer"
                },
                "work_time": {
                    "type": "integer"
                }
            }
        },
        "domain.ReportItem": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "break_time": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "integer"
                },
                "creator_id": {
                    "type": "string"
                },
                "date": {
                    "type": "integer"
                },
                "end_time": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "start_time": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "integer"
                },
                "work_time": {
                    "type": "integer"
                }
            }
        },
        "domain.SignInRequest": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.SignInResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "domain.SignUpRequest": {
            "type": "object",
            "properties": {
                "display_name": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.UpdateReportRequest": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "break_time": {
                    "type": "integer"
                },
                "date": {
                    "type": "integer"
                },
                "end_time": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "work_time": {
                    "type": "integer"
                }
            }
        },
        "httperror.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/httperror.ErrorResponseError"
                }
            }
        },
        "httperror.ErrorResponseDetails": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                },
                "reason": {
                    "type": "string"
                }
            }
        },
        "httperror.ErrorResponseError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "details": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/httperror.ErrorResponseDetails"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Reporter API",
	Description:      "This is a report server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
