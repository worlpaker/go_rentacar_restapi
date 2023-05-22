package handlers

import "errors"

var (
	ErrEmptyRequestBody      = errors.New("empty request body")
	ErrMissingRequiredParams = errors.New("missing required parameters")
	ErrConvert               = errors.New("error occurred while converting data")
	ErrDecodeJSON            = errors.New("error occurred while decoding JSON")
)
