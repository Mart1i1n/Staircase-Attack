// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
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
        "/duties/{epoch}": {
            "get": {
                "description": "get duties by epoch",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get duties by epoch",
                "operationId": "get-duties-by-epoch",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Epoch",
                        "name": "epoch",
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
                                "$ref": "#/definitions/types.ProposerDuty"
                            }
                        }
                    }
                }
            }
        },
        "/reorgs": {
            "get": {
                "description": "get reorgs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get reorgs",
                "operationId": "get-reorgs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.ReorgEvent"
                            }
                        }
                    }
                }
            }
        },
        "/reward/{epoch}": {
            "get": {
                "description": "get reward by epoch",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get reward by epoch",
                "operationId": "get-reward-by-epoch",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Epoch",
                        "name": "epoch",
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
                                "$ref": "#/definitions/dbmodel.BlockReward"
                            }
                        }
                    }
                }
            }
        },
        "/strategy": {
            "get": {
                "description": "get strategy config",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get strategy config",
                "operationId": "get-strategy",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Strategy"
                        }
                    }
                }
            }
        },
        "/update-strategy": {
            "post": {
                "description": "update strategy",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update strategy",
                "operationId": "update-strategy",
                "parameters": [
                    {
                        "description": "Strategy",
                        "name": "strategy",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Strategy"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dbmodel.BlockReward": {
            "type": "object",
            "properties": {
                "epoch": {
                    "description": "epoch",
                    "type": "integer"
                },
                "head_amount": {
                    "description": "Head 奖励数量",
                    "type": "integer"
                },
                "id": {
                    "description": "任务类型id",
                    "type": "integer"
                },
                "target_amount": {
                    "description": "Target 奖励数量",
                    "type": "integer"
                },
                "validator_index": {
                    "description": "验证者索引",
                    "type": "integer"
                }
            }
        },
        "types.AttestStrategy": {
            "type": "object",
            "properties": {
                "broad_cast_delay": {
                    "description": "unit millisecond",
                    "type": "integer"
                },
                "delay_enable": {
                    "type": "boolean"
                },
                "modify_enable": {
                    "type": "boolean"
                }
            }
        },
        "types.BlockStrategy": {
            "type": "object",
            "properties": {
                "broad_cast_delay": {
                    "description": "unit millisecond",
                    "type": "integer"
                },
                "delay_enable": {
                    "type": "boolean"
                },
                "modify_enable": {
                    "type": "boolean"
                }
            }
        },
        "types.ProposerDuty": {
            "type": "object",
            "properties": {
                "pubkey": {
                    "type": "string"
                },
                "slot": {
                    "type": "string"
                },
                "validator_index": {
                    "type": "string"
                }
            }
        },
        "types.ReorgEvent": {
            "type": "object",
            "properties": {
                "depth": {
                    "type": "integer"
                },
                "epoch": {
                    "type": "integer"
                },
                "new_block_proposer_index": {
                    "type": "integer"
                },
                "new_block_slot": {
                    "type": "integer"
                },
                "new_head_state": {
                    "type": "string"
                },
                "old_block_proposer_index": {
                    "type": "integer"
                },
                "old_block_slot": {
                    "type": "integer"
                },
                "old_head_state": {
                    "type": "string"
                },
                "slot": {
                    "type": "integer"
                }
            }
        },
        "types.SlotStrategy": {
            "type": "object",
            "properties": {
                "actions": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "level": {
                    "type": "integer"
                },
                "slot": {
                    "type": "string"
                }
            }
        },
        "types.Strategy": {
            "type": "object",
            "properties": {
                "attest": {
                    "$ref": "#/definitions/types.AttestStrategy"
                },
                "block": {
                    "$ref": "#/definitions/types.BlockStrategy"
                },
                "slots": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.SlotStrategy"
                    }
                },
                "validator": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.ValidatorStrategy"
                    }
                }
            }
        },
        "types.ValidatorStrategy": {
            "type": "object",
            "properties": {
                "attacker_end_slot": {
                    "type": "integer"
                },
                "attacker_start_slot": {
                    "type": "integer"
                },
                "validator_index": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "localhost:20001",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Attacker Service API",
	Description:      "This is the attacker service API server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
