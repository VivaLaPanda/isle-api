package models

import (
	"net/url"
	"time"
)

// ContentNode content node
type ContentNode struct {

	// body
	Body string `json:"body,omitempty"`

	// children
	// Required: true
	Children *ContentNode `json:"children,omitempty"`

	// created
	// Required: true
	Created time.Time `json:"created"`
	// edited
	// Required: true
	Edited time.Time `json:"edited"`

	// image Uri
	ImageURI url.URL `json:"imageUri,omitempty"`

	// score
	// Required: true
	Score float64 `json:"score"`

	// sentiment
	// Required: true
	Sentiment float64 `json:"sentiment"`

	// tags
	Tags []Tag `json:"tags"`

	// title
	Title string `json:"title,omitempty"`

	// type
	DgraphType string `json:"dgraph.type,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// NewContentNode is a struct designed to handle special parameters for creating a node in DGraph
type NewContentNode struct {
	ContentNode
	ParentUID string `json:"parent"`
}
