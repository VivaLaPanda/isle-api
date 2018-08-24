// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is an API meant to interface with the Isle Network database to serve entities to a front-end interface. The reference server is written and Go and the reference client is in Javascript",
    "title": "Isle Network",
    "version": "1.0.0"
  },
  "host": "virtserver.swaggerhub.com",
  "basePath": "/api",
  "paths": {
    "/comments": {
      "get": {
        "tags": [
          "comments"
        ],
        "summary": "Query comments",
        "operationId": "getComments",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "description": "How many items to return at one time (max 100)",
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "description": "What item to start listing at",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Do a full text search of comment bodies",
            "name": "fullText",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Only show comments by a particular user",
            "name": "postedBy",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "An paged array of pets",
            "schema": {
              "$ref": "#/definitions/ContentNodes"
            },
            "headers": {
              "x-next": {
                "type": "string",
                "description": "A link to the next page of responses"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "comments"
        ],
        "summary": "Create a comment",
        "operationId": "newComment",
        "parameters": [
          {
            "description": "The comment to create",
            "name": "comment",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ContentNode"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Null response"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/comments/{commentId}": {
      "get": {
        "tags": [
          "comments"
        ],
        "summary": "Query for a specific comment",
        "operationId": "getCommentsById",
        "parameters": [
          {
            "type": "string",
            "description": "The id of the comment to retrieve",
            "name": "commentId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Expected response to a valid request",
            "schema": {
              "$ref": "#/definitions/ContentNode"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "comments"
        ],
        "summary": "Update a comment",
        "operationId": "updateComment",
        "parameters": [
          {
            "type": "string",
            "description": "The id of the comment to update",
            "name": "commentId",
            "in": "path",
            "required": true
          },
          {
            "description": "The data to replace the old data with",
            "name": "comment",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ContentNode"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Null response"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/ping": {
      "get": {
        "security": [],
        "description": "This operation shows how to override the global security defined above, as we want to open it up for all users.",
        "summary": "Server heartbeat operation",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/posts": {
      "get": {
        "tags": [
          "posts"
        ],
        "summary": "Query posts",
        "operationId": "getPosts",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "description": "How many items to return at one time (max 100)",
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "description": "What item to start listing at",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "string",
            "description": "What tags to include",
            "name": "tags",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Do a full text search of titles/body",
            "name": "fullText",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Only show comments by a particular user",
            "name": "postedBy",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "An paged array of posts",
            "schema": {
              "$ref": "#/definitions/ContentNodes"
            },
            "headers": {
              "x-next": {
                "type": "string",
                "description": "A link to the next page of responses"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "posts"
        ],
        "summary": "Create a post",
        "operationId": "newPost",
        "parameters": [
          {
            "description": "The post to create",
            "name": "post",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ContentNode"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Null response"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/posts/{postId}": {
      "get": {
        "tags": [
          "posts"
        ],
        "summary": "Query for a specific post",
        "operationId": "getPostById",
        "parameters": [
          {
            "type": "string",
            "description": "The id of the post to retrieve",
            "name": "postId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Expected response to a valid request",
            "schema": {
              "$ref": "#/definitions/ContentNode"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "posts"
        ],
        "summary": "Update a post",
        "operationId": "updatePost",
        "parameters": [
          {
            "type": "string",
            "description": "The id of the post to update",
            "name": "postId",
            "in": "path",
            "required": true
          },
          {
            "description": "The data to replace the old data with",
            "name": "post",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ContentNode"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Null response"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/roles": {
      "get": {
        "tags": [
          "roles"
        ],
        "summary": "Query roles",
        "operationId": "getRoles",
        "responses": {
          "200": {
            "description": "An paged array of roles",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Role"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/tags": {
      "get": {
        "tags": [
          "tags"
        ],
        "summary": "Query tags",
        "operationId": "getTags",
        "responses": {
          "200": {
            "description": "An paged array of tags",
            "schema": {
              "$ref": "#/definitions/Tags"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "tags"
        ],
        "summary": "Make new tag",
        "operationId": "newTag",
        "responses": {
          "201": {
            "description": "Null response"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/tags/{tagId}": {
      "put": {
        "tags": [
          "tags"
        ],
        "summary": "Update a tag",
        "operationId": "updateTag",
        "parameters": [
          {
            "type": "string",
            "description": "The id of the tag to update",
            "name": "tagId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Null response"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "tags": [
          "users"
        ],
        "summary": "Query users",
        "operationId": "getUsers",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "description": "How many items to return at one time (max 100)",
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "description": "What item to start listing at",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Search by username",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "An paged array of users",
            "schema": {
              "$ref": "#/definitions/Users"
            },
            "headers": {
              "x-next": {
                "type": "string",
                "description": "A link to the next page of responses"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "users"
        ],
        "summary": "Create a user",
        "operationId": "newUser",
        "parameters": [
          {
            "type": "string",
            "name": "inviteCode",
            "in": "query",
            "required": true
          },
          {
            "description": "The user to create",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ContentNode"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Null response"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/users/{userId}": {
      "get": {
        "tags": [
          "users"
        ],
        "summary": "Query for a specific user",
        "operationId": "getUserById",
        "parameters": [
          {
            "type": "string",
            "description": "The id of the user to retrieve",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Expected response to a valid request",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "users"
        ],
        "summary": "Update a user",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "description": "The id of the user to update",
            "name": "userId",
            "in": "path",
            "required": true
          },
          {
            "description": "The data to replace the old data with",
            "name": "user",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Null response"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ContentNode": {
      "required": [
        "created",
        "score",
        "tags",
        "sentiment",
        "children"
      ],
      "properties": {
        "body": {
          "type": "string"
        },
        "children": {
          "$ref": "#/definitions/ContentNodes"
        },
        "created": {
          "type": "string"
        },
        "imageUri": {
          "type": "string"
        },
        "score": {
          "type": "number"
        },
        "sentiment": {
          "type": "number"
        },
        "tags": {
          "$ref": "#/definitions/Tags"
        },
        "title": {
          "type": "string"
        },
        "uid": {
          "type": "string"
        }
      }
    },
    "ContentNodes": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/ContentNode"
      }
    },
    "Error": {
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Invite": {
      "properties": {
        "code": {
          "type": "string"
        },
        "createdBy": {
          "$ref": "#/definitions/User"
        },
        "uid": {
          "type": "string"
        }
      }
    },
    "Purchasable": {
      "required": [
        "name",
        "cost"
      ],
      "properties": {
        "cost": {
          "type": "number"
        },
        "name": {
          "type": "string"
        },
        "uid": {
          "type": "string"
        }
      }
    },
    "Role": {
      "required": [
        "text"
      ],
      "properties": {
        "text": {
          "type": "string"
        },
        "uid": {
          "type": "string"
        }
      }
    },
    "Tag": {
      "required": [
        "text"
      ],
      "properties": {
        "text": {
          "type": "string"
        },
        "uid": {
          "type": "string"
        }
      }
    },
    "Tags": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Tag"
      }
    },
    "User": {
      "required": [
        "name",
        "email",
        "reputation",
        "role"
      ],
      "properties": {
        "aviImgUri": {
          "type": "string"
        },
        "commented": {
          "$ref": "#/definitions/ContentNodes"
        },
        "email": {
          "type": "string"
        },
        "invitedBy": {
          "$ref": "#/definitions/User"
        },
        "joined": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "posted": {
          "$ref": "#/definitions/ContentNodes"
        },
        "reputation": {
          "type": "number"
        },
        "role": {
          "$ref": "#/definitions/Role"
        },
        "spent": {
          "type": "number"
        },
        "uid": {
          "type": "string"
        }
      }
    },
    "Users": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/User"
      }
    }
  },
  "securityDefinitions": {
    "hasRole": {
      "type": "oauth2",
      "flow": "password",
      "tokenUrl": "https://www.googleapis.com/oauth2/v4/token",
      "scopes": {
        "admin": "can modify almost anything",
        "mod": "user + special operations",
        "user": "can modify own resources"
      }
    }
  },
  "security": [
    {
      "hasRole": [
        "user",
        "mod"
      ]
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is an API meant to interface with the Isle Network database to serve entities to a front-end interface. The reference server is written and Go and the reference client is in Javascript",
    "title": "Isle Network",
    "version": "1.0.0"
  },
  "host": "virtserver.swaggerhub.com",
  "basePath": "/api",
  "paths": {
    "/comments": {
      "get": {
        "tags": [
          "comments"
        ],
        "summary": "Query comments",
        "operationId": "getComments",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "description": "How many items to return at one time (max 100)",
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "description": "What item to start listing at",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Do a full text search of comment bodies",
            "name": "fullText",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Only show comments by a particular user",
            "name": "postedBy",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "An paged array of pets",
            "schema": {
              "$ref": "#/definitions/ContentNodes"
            },
            "headers": {
              "x-next": {
                "type": "string",
                "description": "A link to the next page of responses"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "comments"
        ],
        "summary": "Create a comment",
        "operationId": "newComment",
        "parameters": [
          {
            "description": "The comment to create",
            "name": "comment",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ContentNode"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Null response"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/comments/{commentId}": {
      "get": {
        "tags": [
          "comments"
        ],
        "summary": "Query for a specific comment",
        "operationId": "getCommentsById",
        "parameters": [
          {
            "type": "string",
            "description": "The id of the comment to retrieve",
            "name": "commentId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Expected response to a valid request",
            "schema": {
              "$ref": "#/definitions/ContentNode"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "comments"
        ],
        "summary": "Update a comment",
        "operationId": "updateComment",
        "parameters": [
          {
            "type": "string",
            "description": "The id of the comment to update",
            "name": "commentId",
            "in": "path",
            "required": true
          },
          {
            "description": "The data to replace the old data with",
            "name": "comment",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ContentNode"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Null response"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/ping": {
      "get": {
        "security": [],
        "description": "This operation shows how to override the global security defined above, as we want to open it up for all users.",
        "summary": "Server heartbeat operation",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/posts": {
      "get": {
        "tags": [
          "posts"
        ],
        "summary": "Query posts",
        "operationId": "getPosts",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "description": "How many items to return at one time (max 100)",
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "description": "What item to start listing at",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "string",
            "description": "What tags to include",
            "name": "tags",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Do a full text search of titles/body",
            "name": "fullText",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Only show comments by a particular user",
            "name": "postedBy",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "An paged array of posts",
            "schema": {
              "$ref": "#/definitions/ContentNodes"
            },
            "headers": {
              "x-next": {
                "type": "string",
                "description": "A link to the next page of responses"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "posts"
        ],
        "summary": "Create a post",
        "operationId": "newPost",
        "parameters": [
          {
            "description": "The post to create",
            "name": "post",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ContentNode"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Null response"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/posts/{postId}": {
      "get": {
        "tags": [
          "posts"
        ],
        "summary": "Query for a specific post",
        "operationId": "getPostById",
        "parameters": [
          {
            "type": "string",
            "description": "The id of the post to retrieve",
            "name": "postId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Expected response to a valid request",
            "schema": {
              "$ref": "#/definitions/ContentNode"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "posts"
        ],
        "summary": "Update a post",
        "operationId": "updatePost",
        "parameters": [
          {
            "type": "string",
            "description": "The id of the post to update",
            "name": "postId",
            "in": "path",
            "required": true
          },
          {
            "description": "The data to replace the old data with",
            "name": "post",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ContentNode"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Null response"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/roles": {
      "get": {
        "tags": [
          "roles"
        ],
        "summary": "Query roles",
        "operationId": "getRoles",
        "responses": {
          "200": {
            "description": "An paged array of roles",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Role"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/tags": {
      "get": {
        "tags": [
          "tags"
        ],
        "summary": "Query tags",
        "operationId": "getTags",
        "responses": {
          "200": {
            "description": "An paged array of tags",
            "schema": {
              "$ref": "#/definitions/Tags"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "tags"
        ],
        "summary": "Make new tag",
        "operationId": "newTag",
        "responses": {
          "201": {
            "description": "Null response"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/tags/{tagId}": {
      "put": {
        "tags": [
          "tags"
        ],
        "summary": "Update a tag",
        "operationId": "updateTag",
        "parameters": [
          {
            "type": "string",
            "description": "The id of the tag to update",
            "name": "tagId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Null response"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "tags": [
          "users"
        ],
        "summary": "Query users",
        "operationId": "getUsers",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "description": "How many items to return at one time (max 100)",
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "description": "What item to start listing at",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Search by username",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "An paged array of users",
            "schema": {
              "$ref": "#/definitions/Users"
            },
            "headers": {
              "x-next": {
                "type": "string",
                "description": "A link to the next page of responses"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "users"
        ],
        "summary": "Create a user",
        "operationId": "newUser",
        "parameters": [
          {
            "type": "string",
            "name": "inviteCode",
            "in": "query",
            "required": true
          },
          {
            "description": "The user to create",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ContentNode"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Null response"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/users/{userId}": {
      "get": {
        "tags": [
          "users"
        ],
        "summary": "Query for a specific user",
        "operationId": "getUserById",
        "parameters": [
          {
            "type": "string",
            "description": "The id of the user to retrieve",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Expected response to a valid request",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "users"
        ],
        "summary": "Update a user",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "description": "The id of the user to update",
            "name": "userId",
            "in": "path",
            "required": true
          },
          {
            "description": "The data to replace the old data with",
            "name": "user",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Null response"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ContentNode": {
      "required": [
        "created",
        "score",
        "tags",
        "sentiment",
        "children"
      ],
      "properties": {
        "body": {
          "type": "string"
        },
        "children": {
          "$ref": "#/definitions/ContentNodes"
        },
        "created": {
          "type": "string"
        },
        "imageUri": {
          "type": "string"
        },
        "score": {
          "type": "number"
        },
        "sentiment": {
          "type": "number"
        },
        "tags": {
          "$ref": "#/definitions/Tags"
        },
        "title": {
          "type": "string"
        },
        "uid": {
          "type": "string"
        }
      }
    },
    "ContentNodes": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/ContentNode"
      }
    },
    "Error": {
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Invite": {
      "properties": {
        "code": {
          "type": "string"
        },
        "createdBy": {
          "$ref": "#/definitions/User"
        },
        "uid": {
          "type": "string"
        }
      }
    },
    "Purchasable": {
      "required": [
        "name",
        "cost"
      ],
      "properties": {
        "cost": {
          "type": "number"
        },
        "name": {
          "type": "string"
        },
        "uid": {
          "type": "string"
        }
      }
    },
    "Role": {
      "required": [
        "text"
      ],
      "properties": {
        "text": {
          "type": "string"
        },
        "uid": {
          "type": "string"
        }
      }
    },
    "Tag": {
      "required": [
        "text"
      ],
      "properties": {
        "text": {
          "type": "string"
        },
        "uid": {
          "type": "string"
        }
      }
    },
    "Tags": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Tag"
      }
    },
    "User": {
      "required": [
        "name",
        "email",
        "reputation",
        "role"
      ],
      "properties": {
        "aviImgUri": {
          "type": "string"
        },
        "commented": {
          "$ref": "#/definitions/ContentNodes"
        },
        "email": {
          "type": "string"
        },
        "invitedBy": {
          "$ref": "#/definitions/User"
        },
        "joined": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "posted": {
          "$ref": "#/definitions/ContentNodes"
        },
        "reputation": {
          "type": "number"
        },
        "role": {
          "$ref": "#/definitions/Role"
        },
        "spent": {
          "type": "number"
        },
        "uid": {
          "type": "string"
        }
      }
    },
    "Users": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/User"
      }
    }
  },
  "securityDefinitions": {
    "hasRole": {
      "type": "oauth2",
      "flow": "password",
      "tokenUrl": "https://www.googleapis.com/oauth2/v4/token",
      "scopes": {
        "admin": "can modify almost anything",
        "mod": "user + special operations",
        "user": "can modify own resources"
      }
    }
  },
  "security": [
    {
      "hasRole": [
        "user",
        "mod"
      ]
    }
  ]
}`))
}
