// Package swagger GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package swagger

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/api/v1/captures": {
            "get": {
                "description": "list all captures in cdc cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "capture"
                ],
                "summary": "List captures",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Capture"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/changefeeds": {
            "get": {
                "description": "list all changefeeds in cdc cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "changefeed"
                ],
                "summary": "List changefeed",
                "parameters": [
                    {
                        "type": "string",
                        "description": "state",
                        "name": "state",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ChangefeedCommonInfo"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "create a new changefeed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "changefeed"
                ],
                "summary": "Create changefeed",
                "parameters": [
                    {
                        "description": "changefeed config",
                        "name": "changefeed",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ChangefeedConfig"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/changefeeds/{changefeed_id}": {
            "get": {
                "description": "get detail information of a changefeed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "changefeed"
                ],
                "summary": "Get changefeed",
                "parameters": [
                    {
                        "type": "string",
                        "description": "changefeed_id",
                        "name": "changefeed_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ChangefeedDetail"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a changefeed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "changefeed"
                ],
                "summary": "Update a changefeed",
                "parameters": [
                    {
                        "type": "string",
                        "description": "changefeed_id",
                        "name": "changefeed_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "changefeed target ts",
                        "name": "target_ts",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "sink uri",
                        "name": "sink_uri",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "filter rules",
                        "name": "filter_rules",
                        "in": "body",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    {
                        "description": "ignore transaction start ts",
                        "name": "ignore_txn_start_ts",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "mounter worker nums",
                        "name": "mounter_worker_num",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "sink config",
                        "name": "sink_config",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/config.SinkConfig"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove a changefeed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "changefeed"
                ],
                "summary": "Remove a changefeed",
                "parameters": [
                    {
                        "type": "string",
                        "description": "changefeed_id",
                        "name": "changefeed_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/changefeeds/{changefeed_id}/pause": {
            "post": {
                "description": "Pause a changefeed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "changefeed"
                ],
                "summary": "Pause a changefeed",
                "parameters": [
                    {
                        "type": "string",
                        "description": "changefeed_id",
                        "name": "changefeed_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/changefeeds/{changefeed_id}/resume": {
            "post": {
                "description": "Resume a changefeed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "changefeed"
                ],
                "summary": "Resume a changefeed",
                "parameters": [
                    {
                        "type": "string",
                        "description": "changefeed_id",
                        "name": "changefeed-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/changefeeds/{changefeed_id}/tables/move_table": {
            "post": {
                "description": "move one table to the target capture",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "changefeed"
                ],
                "summary": "move table",
                "parameters": [
                    {
                        "type": "string",
                        "description": "changefeed_id",
                        "name": "changefeed_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "table_id",
                        "name": "table_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "capture_id",
                        "name": "capture_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/changefeeds/{changefeed_id}/tables/rebalance_table": {
            "post": {
                "description": "rebalance all tables of a changefeed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "changefeed"
                ],
                "summary": "rebalance tables",
                "parameters": [
                    {
                        "type": "string",
                        "description": "changefeed_id",
                        "name": "changefeed_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/health": {
            "get": {
                "description": "check if CDC cluster is health",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "common"
                ],
                "summary": "Check if CDC cluster is health",
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/log": {
            "post": {
                "description": "change TiCDC log level dynamically",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "common"
                ],
                "summary": "Change TiCDC log level",
                "parameters": [
                    {
                        "description": "log level",
                        "name": "log_level",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/owner/resign": {
            "post": {
                "description": "notify the current owner to resign",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "owner"
                ],
                "summary": "notify the owner to resign",
                "responses": {
                    "202": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/processors": {
            "get": {
                "description": "list all processors in the TiCDC cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "processor"
                ],
                "summary": "List processors",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ProcessorCommonInfo"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/processors/{changefeed_id}/{capture_id}": {
            "get": {
                "description": "get the detail information of a processor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "processor"
                ],
                "summary": "Get processor detail information",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ProcessorDetail"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/status": {
            "get": {
                "description": "get the status of a server(capture)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "common"
                ],
                "summary": "Get server status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ServerStatus"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "config.ColumnSelector": {
            "type": "object",
            "properties": {
                "columns": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "matcher": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "config.DispatchRule": {
            "type": "object",
            "properties": {
                "dispatcher": {
                    "description": "Deprecated, please use PartitionRule.",
                    "type": "string"
                },
                "matcher": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "partition": {
                    "description": "PartitionRule is an alias added for DispatcherRule to mitigate confusions.\nIn the future release, the DispatcherRule is expected to be removed .",
                    "type": "string"
                },
                "topic": {
                    "type": "string"
                }
            }
        },
        "config.SinkConfig": {
            "type": "object",
            "properties": {
                "column-selectors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/config.ColumnSelector"
                    }
                },
                "dispatchers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/config.DispatchRule"
                    }
                },
                "protocol": {
                    "type": "string"
                },
                "schema-registry": {
                    "type": "string"
                },
                "transaction-atomicity": {
                    "type": "string"
                }
            }
        },
        "model.Capture": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_owner": {
                    "type": "boolean"
                }
            }
        },
        "model.CaptureTaskStatus": {
            "type": "object",
            "properties": {
                "capture_id": {
                    "type": "string"
                },
                "table_ids": {
                    "description": "Table list, containing tables that processor should process",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "table_operations": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/model.TableOperation"
                    }
                }
            }
        },
        "model.ChangefeedCommonInfo": {
            "type": "object",
            "properties": {
                "checkpoint_time": {
                    "type": "string"
                },
                "checkpoint_tso": {
                    "type": "integer"
                },
                "error": {
                    "$ref": "#/definitions/model.RunningError"
                },
                "id": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                }
            }
        },
        "model.ChangefeedConfig": {
            "type": "object",
            "properties": {
                "changefeed_id": {
                    "type": "string"
                },
                "filter_rules": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "force_replicate": {
                    "description": "if true, force to replicate some ineligible tables",
                    "type": "boolean",
                    "default": false
                },
                "ignore_ineligible_table": {
                    "type": "boolean",
                    "default": false
                },
                "ignore_txn_start_ts": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "mounter_worker_num": {
                    "type": "integer",
                    "default": 16
                },
                "namespace": {
                    "type": "string"
                },
                "sink_config": {
                    "$ref": "#/definitions/config.SinkConfig"
                },
                "sink_uri": {
                    "type": "string"
                },
                "start_ts": {
                    "type": "integer"
                },
                "target_ts": {
                    "type": "integer"
                },
                "timezone": {
                    "description": "timezone used when checking sink uri",
                    "type": "string",
                    "default": "system"
                }
            }
        },
        "model.ChangefeedDetail": {
            "type": "object",
            "properties": {
                "checkpoint_time": {
                    "type": "string"
                },
                "checkpoint_tso": {
                    "type": "integer"
                },
                "create_time": {
                    "type": "string"
                },
                "creator_version": {
                    "type": "string"
                },
                "error": {
                    "$ref": "#/definitions/model.RunningError"
                },
                "error_history": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "resolved_ts": {
                    "type": "integer"
                },
                "sink_uri": {
                    "type": "string"
                },
                "sort_engine": {
                    "type": "string"
                },
                "start_ts": {
                    "type": "integer"
                },
                "state": {
                    "type": "string"
                },
                "target_ts": {
                    "type": "integer"
                },
                "task_status": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CaptureTaskStatus"
                    }
                }
            }
        },
        "model.HTTPError": {
            "type": "object",
            "properties": {
                "error_code": {
                    "type": "string"
                },
                "error_msg": {
                    "type": "string"
                }
            }
        },
        "model.ProcessorCommonInfo": {
            "type": "object",
            "properties": {
                "capture_id": {
                    "type": "string"
                },
                "changefeed_id": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                }
            }
        },
        "model.ProcessorDetail": {
            "type": "object",
            "properties": {
                "checkpoint_ts": {
                    "description": "The maximum event CommitTs that has been synchronized.",
                    "type": "integer"
                },
                "count": {
                    "description": "The count of events that have been replicated.",
                    "type": "integer"
                },
                "error": {
                    "description": "Error code when error happens",
                    "$ref": "#/definitions/model.RunningError"
                },
                "resolved_ts": {
                    "description": "The event that satisfies CommitTs \u003c= ResolvedTs can be synchronized.",
                    "type": "integer"
                },
                "table_ids": {
                    "description": "all table ids that this processor are replicating",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "model.RunningError": {
            "type": "object",
            "properties": {
                "addr": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.ServerStatus": {
            "type": "object",
            "properties": {
                "git_hash": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_owner": {
                    "type": "boolean"
                },
                "pid": {
                    "type": "integer"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "model.TableOperation": {
            "type": "object",
            "properties": {
                "boundary_ts": {
                    "description": "if the operation is a delete operation, BoundaryTs is checkpoint ts\nif the operation is an add operation, BoundaryTs is start ts",
                    "type": "integer"
                },
                "delete": {
                    "type": "boolean"
                },
                "flag": {
                    "type": "integer"
                },
                "status": {
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
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
