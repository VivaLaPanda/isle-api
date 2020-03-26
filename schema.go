package main

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)

// LoadSchema makes sure the DGraph schema is loaded into the database on starting
func LoadSchema(db *dgo.Dgraph) {
	op := &api.Operation{}
	op.Schema = `
		# Type predicate
		type: [uid] .
		
		# ContentNode Predicates
		title: string @index(exact, fulltext) .
		imageUri: string .
		body: string @index(fulltext) .
		created: dateTime @index(hour) .
		edited: dateTime @index(hour) .
		author: [uid] @reverse . 
		tags: [uid] @reverse .
		score: float .
		sentiment: float .
		parent: [uid] @reverse .
		
		type ContentNode {
			title
			imageUri
			body
			created
			edited
			author
			tags
			score
			sentiment
			parent
		}
		
		# Role Predicates
		text: string @index(exact).
		type Role {
			text
		}
		
		type Tag {
			text
		}
		
		# Invite Predicates
		code: string @index(hash) .
		invitedBy: [uid] .
		type Invite {
			code
			invitedBy
		}
		
		
		# User Predicates 
		name: string @index(exact) .
		password: password .
		email: string @index(hash) .
		aviImgUri: string .
		reputation: float .
		spent: float .
		role: [uid] @reverse .
		invitedBy: [uid] @reverse .
		joined: dateTime .
		type User {
			name
			password
			email
			aviImgUri
			reputation
			spent
			role
			invitedBy
			joined
		}

	`

	err := db.Alter(context.Background(), op)
	if err != nil {
		log.Fatal(err)
	}
}
