// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/adverts": {
            "get": {
                "description": "广告列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "广告管理"
                ],
                "summary": "广告列表",
                "parameters": [
                    {
                        "type": "string",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/res.ListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "创建广告",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "广告管理"
                ],
                "summary": "创建广告",
                "parameters": [
                    {
                        "description": "表示多个参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/advert_api.AdvertRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "批量删除广告",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "广告管理"
                ],
                "summary": "批量删除广告",
                "parameters": [
                    {
                        "description": "广告id列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RemoveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/adverts/:id": {
            "put": {
                "description": "更新广告",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "广告管理"
                ],
                "summary": "更新广告",
                "parameters": [
                    {
                        "description": "广告的一些参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/advert_api.AdvertRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "advert_api.AdvertRequest": {
            "type": "object",
            "required": [
                "href",
                "images",
                "title"
            ],
            "properties": {
                "href": {
                    "description": "跳转链接",
                    "type": "string"
                },
                "images": {
                    "description": "图片",
                    "type": "string"
                },
                "is_show": {
                    "description": "是否展示",
                    "type": "boolean"
                },
                "title": {
                    "description": "显示的标题",
                    "type": "string"
                }
            }
        },
        "models.RemoveRequest": {
            "type": "object",
            "properties": {
                "id_list": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "res.ListResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {}
            }
        },
        "res.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.01:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "gvb_server API文档",
	Description:      "gvb_server API文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
