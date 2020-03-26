package dal

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/VivaLaPanda/isle-api/api/models"
	"github.com/dgraph-io/dgo"
)

// NewUser creates a new User based on the provided struct
func NewUser(db *dgo.Dgraph, user models.User) (uid string, err error) {
	// Handle validation, etc
	// Making sure creation time is current
	user.Joined = time.Now()

	return Mutator(db, user)
}

// NewRole creates a new User based on the provided struct
func NewRole(db *dgo.Dgraph, role models.Role) (uid string, err error) {
	// Handle validation, etc

	return Mutator(db, role)
}

// GetUser returns a struct with a summary of the information for a user
func GetUser(db *dgo.Dgraph, uid string) (resp models.User, err error) {
	// Set up transaction
	txn := db.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	// Construct the query
	const q = `
	query GetUser($id: string) {
		user(func: uid($id)) {
			uid
			name
			email
			joined
			reputation
			spent
			aviImgUri
			role {
				text
				uid
			}
			invitedBy {
				uid
				name
				aviImgUri
			}
			~author(orderasc: edited, first: 10) {
				uid
				title
				body
				created
				imageUri
				sentiment
				score
				tags
			}
		}
	}
	`

	// Make variables map
	variables := map[string]string{"$id": uid}

	// Run the query
	out, err := txn.QueryWithVars(context.Background(), q, variables)
	if err != nil {
		return resp, err
	}

	txn.Commit(context.Background())

	// Decode the response
	var decode struct {
		User []models.User
	}
	if err := json.Unmarshal(out.GetJson(), &decode); err != nil {
		log.Fatal(err)
	}

	// There's only ever one node
	return decode.User[0], nil
}
