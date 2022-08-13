package kubernetes

import "errors"

var (
	// ErrEmptyManifest is returned when the manifest is empty.
	ErrEmptyManifest = errors.New("manifest is empty")

	// ErrNotSupported is returned when the operation is not supported.
	ErrNotSupported = errors.New("not supported")
)
