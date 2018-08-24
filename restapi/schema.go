package restapi

import (
	"context"

	"github.com/dgraph-io/dgo/protos/api"
)

// This file will write the schema to the databse every time the server is launched
// to ensure the server is in sync. Writing the same schema twice
// isn't a problem, so it's fine for this to run every time
func writeSchema() {
	err := db.Alter(context.Background(), &api.Operation{
		Schema: `
      contentNode: default .
      contentNode.title: string @index(exact, fulltext) .
      contentNode.imageUri: string .
      contentNode.body: string @index(fulltext) .
      contentNode.created: dateTime @index(hour) .
      contentNode.author: uid @reverse .
      contentNode.tags: uid @reverse .
      contentNode.score: float .
      contentNode.sentiment: float .
      contentNode.parent: uid @reverse .

      role: default .
      role.text: string @index(exact).

      tag: default .
      tag.text: string @index(exact) .

      invite: default .
      invite.code: string @index(hash) .
      invite.created: uid .

      user: default .
      user.name: string @index(exact) .
      user.email: string @index(hash) .
      user.aviImgUri: string .
      user.reputation: float .
      user.spent: float .
      user.role: uid @reverse .
      user.invitedBy: uid @reverse .
      user.joined: dateTime .
		`,
	})

	if err != nil {
		panic("Failed to write schema to database on startup.")
	}
}
