package models

import "time"

type User struct {

	// avi img Uri
	AviImgURI *string `json:"aviImgUri,omitempty"`

	// commented
	Posted []ContentNode `json:"~author,omitempty"`

	// email
	// Required: true
	Email string `json:"email,omitempty"`

	// invited by
	InvitedBy []*User `json:"invitedBy,omitempty"`

	// joined
	// Required: true
	Joined *time.Time `json:"joined,omitempty"`

	// name
	// Required: true
	Name string `json:"name,omitempty"`

	// reputation
	// Required: true
	Reputation float64 `json:"reputation,omitempty"`

	// role
	// Required: true
	Role []*Role `json:"role,omitempty"`

	// spent
	Spent float64 `json:"spent,omitempty"`

	// uid
	UID *string `json:"uid,omitempty"`

	// type
	DgraphType []string `json:"dgraph.type,omitempty"`
}

// // NewContentNodeNode is a struct designed to handle special parameters for creating a node in DGraph
// type NewUserNode struct {
// 	User
// 	InvitedByUID *Edge `json:"invitedBy,omitempty"`
// 	Role *Edge `json:"role,omitempty"`
// }

// Role role
type Role struct {

	// text
	// Required: true
	Text string `json:"text,omitempty"`

	// uid
	UID *string `json:"uid,omitempty"`

	// type
	DgraphType []string `json:"dgraph.type,omitempty"`
}

// Invite invite
// swagger:model Invite
type Invite struct {

	// code
	Code string `json:"code,omitempty"`

	// created by
	CreatedBy *User `json:"createdBy,omitempty"`

	// uid
	UID *string `json:"uid,omitempty"`

	// type
	DgraphType []string `json:"dgraph.type,omitempty"`
}
