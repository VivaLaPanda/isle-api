package models

import (
	"net/url"
	"time"
)

// ContentNode content node
type ContentNode struct {

	// title
	Title string `json:"title,omitempty"`

	// body
	Body string `json:"body,omitempty"`

	// children
	// Required: true
	Children []*ContentNode `json:"~parent,omitempty"`

	// created
	// Required: true
	Created time.Time `json:"created,omitempty"`
	// edited
	// Required: true
	Edited time.Time `json:"edited,omitempty"`

	// image Uri
	ImageURI *url.URL `json:"imageUri,omitempty"`

	// score
	// Required: true
	Score float64 `json:"score,omitempty"`

	// sentiment
	// Required: true
	Sentiment float64 `json:"sentiment,omitempty"`

	// tags
	Tags []Tag `json:"tags,omitempty"`

	// type
	DgraphType []string `json:"dgraph.type,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// NewContentNodeNode is a struct designed to handle special parameters for creating a node in DGraph
type NewContentNodeNode struct {
	ContentNode
	ParentUID *Edge `json:"parent,omitempty"`
}
