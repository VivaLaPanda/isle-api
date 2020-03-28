package dal

import (
	"encoding/json"
	"fmt"
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

// GetRole returns a struct with a summary of the information for a role
func GetRole(db *dgo.Dgraph, uid string) (resp models.Role, err error) {
	// Construct the query
	const q = `
	query GetRole($id: string) {
		role(func: uid($id)) @filter(type(User)) {
			text
			uid
			dgraph.type
		}
	}
	`

	jsonResp, err := UIDFetcher(db, q, uid)

	// Decode the response
	var decode struct {
		Role []models.Role
	}
	if err := json.Unmarshal(jsonResp, &decode); err != nil {
		log.Fatal(err)
	}

	if len(decode.Role) != 1 {
		return resp, fmt.Errorf("Query returned %d results, not 1 as expected. Did you query the wrong endpoint?", len(decode.Role))
	}

	// There's only ever one node
	return decode.Role[0], nil
}

// GetUser returns a struct with a summary of the information for a user
func GetUser(db *dgo.Dgraph, uid string) (resp models.User, err error) {
	// Construct the query
	const q = `
	query GetUser($id: string) {
		user(func: uid($id)) @filter(type(User)) {
			uid
			name
			email
			joined
			reputation
			spent
			aviImgUri
			dgraph.type
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

	jsonResp, err := UIDFetcher(db, q, uid)

	// Decode the response
	var decode struct {
		User []models.User
	}
	if err := json.Unmarshal(jsonResp, &decode); err != nil {
		log.Fatal(err)
	}

	if len(decode.User) != 1 {
		return resp, fmt.Errorf("Query returned %d results, not 1 as expected. Did you query the wrong endpoint?", len(decode.User))
	}

	// There's only ever one node
	return decode.User[0], nil
}
