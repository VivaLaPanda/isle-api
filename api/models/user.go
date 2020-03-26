package models

import "time"

type User struct {

	// avi img Uri
	AviImgURI *string `json:"aviImgUri,omitempty"`

	// commented
	Posted []ContentNode `json:"~author"`

	// email
	// Required: true
	Email string `json:"email"`

	// invited by
	InvitedBy *User `json:"invitedBy,omitempty"`

	// joined
	// Required: true
	Joined time.Time `json:"joined"`

	// name
	// Required: true
	Name string `json:"name"`

	// reputation
	// Required: true
	Reputation float64 `json:"reputation"`

	// role
	// Required: true
	Role []*Role `json:"role"`

	// spent
	Spent float64 `json:"spent"`

	// uid
	UID *string `json:"uid,omitempty"`

	// type
	DgraphType string `json:"dgraph.type,omitempty"`
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
	Text string `json:"text"`

	// uid
	UID *string `json:"uid,omitempty"`

	// type
	DgraphType string `json:"dgraph.type,omitempty"`
}
