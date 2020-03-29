package dal

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/VivaLaPanda/isle-api/api/models"
	"github.com/dgraph-io/dgo"
)

// NewUser creates a new User based on the provided struct
func NewUser(db *dgo.Dgraph, user models.User, inviteCode string) (uid string, err error) {
	// Handle validation, etc
	// Making sure creation time is current
	now := time.Now()
	user.Joined = &now

	if inviteCode == "" {
		// TODO: Validate that user is an admin
	} else {
		invite, err := GetInvite(db, inviteCode)
		if err != nil {
			// Something is wrong with the invite, don't create the user
			return "", err
		}

		user.InvitedBy = []*models.User{invite.CreatedBy}
	}

	// TODO: Delete the invite when it's used
	return Mutator(db, user)
}

// NewRole creates a new User based on the provided struct
func NewRole(db *dgo.Dgraph, role models.Role) (uid string, err error) {
	// Handle validation, etc

	return Mutator(db, role)
}

// NewInvite creates a new invite owned by the current user
func NewInvite(db *dgo.Dgraph, invite models.Invite, currentUser models.User) (uid string, err error) {
	// Handle validation, etc
	invite.CreatedBy = &models.User{UID: currentUser.UID}

	return Mutator(db, invite)
}

// GetInvite will return the struct representing a particular invite matching the given code
func GetInvite(db *dgo.Dgraph, inviteCode string) (resp models.Invite, err error) {
	// Construct the query
	const q = `
	query GetInvite($match: string) {
		invite(func: eq(code, $match)) @filter(type(Invite)) {
			uid
			code
			createdBy {
				uid
			}
			dgraph.type
		}
	}
	`

	jsonResp, err := SimpleQuery(db, q, inviteCode)

	// Decode the response
	var decode struct {
		Invite []models.Invite
	}
	if err := json.Unmarshal(jsonResp, &decode); err != nil {
		return resp, err
	}

	// if len(decode.Invite) != 1 {
	// 	return resp, fmt.Errorf("Found %d invites, should only be 1", len(decode.Invite))
	// }

	// We got the expected 1 result, just return that
	return decode.Invite[0], nil
}

// GetRole returns a struct with a summary of the information for a role
func GetRole(db *dgo.Dgraph, uid string) (resp models.Role, err error) {
	// Construct the query
	const q = `
	query GetRole($match: string) {
		role(func: uid($match)) @filter(type(User)) {
			text
			uid
			dgraph.type
		}
	}
	`

	jsonResp, err := SimpleQuery(db, q, uid)

	// Decode the response
	var decode struct {
		Role []models.Role
	}
	if err := json.Unmarshal(jsonResp, &decode); err != nil {
		return resp, err
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
	query GetUser($match: string) {
		user(func: uid($match)) @filter(type(User)) {
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

	jsonResp, err := SimpleQuery(db, q, uid)

	// Decode the response
	var decode struct {
		User []models.User
	}
	if err := json.Unmarshal(jsonResp, &decode); err != nil {
		return resp, err
	}

	if len(decode.User) != 1 {
		return resp, fmt.Errorf("Query returned %d results, not 1 as expected. Did you query the wrong endpoint?", len(decode.User))
	}

	// There's only ever one node
	return decode.User[0], nil
}
