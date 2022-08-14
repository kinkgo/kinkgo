package yamlpath

import "errors"

var (
	// ErrEmptyYamlNode is returned when the YAML node is empty.
	ErrEmptyYamlNode = errors.New("yaml node is empty")

	// ErrInvalidType is returned when the type is invalid.
	ErrInvalidType = errors.New("invalid type")

	// ErrInvalidPathString is returned when the path string is invalid.
	ErrInvalidPathString = errors.New("invalid path string")
)
