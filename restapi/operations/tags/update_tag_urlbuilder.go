// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// UpdateTagURL generates an URL for the update tag operation
type UpdateTagURL struct {
	TagID string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *UpdateTagURL) WithBasePath(bp string) *UpdateTagURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *UpdateTagURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *UpdateTagURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/tags/{tagId}"

	tagID := o.TagID
	if tagID != "" {
		_path = strings.Replace(_path, "{tagId}", tagID, -1)
	} else {
		return nil, errors.New("TagID is required on UpdateTagURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/VivaLaPanda/Isle/1.0.0"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *UpdateTagURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *UpdateTagURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *UpdateTagURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on UpdateTagURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on UpdateTagURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *UpdateTagURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}