package models

// Tag tag
type Tag struct {

	// text
	// Required: true
	Text string `json:"text,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}
