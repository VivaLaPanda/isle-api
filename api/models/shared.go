package models

// Edge handles the fact that DGraph expects edges to be sub objects, not fields
type Edge struct {
	UID string `json:"uid"`
}
