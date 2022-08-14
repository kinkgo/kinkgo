package yamlpath

import (
	"strconv"
	"strings"

	yamlv3 "gopkg.in/yaml.v3"
)

// Path represents a pointer to a location in a YAML document.
type Path struct {
	// root of the path
	Root *yamlv3.Node

	// index of the node in the root
	Index int

	// Segments of the path
	Segments []PathSegment
}

// PathSegment represents single part of the path. A key can be entry in a map or a named entry in a sequence or
// index in a sequence.
type PathSegment struct {
	// Index of the element. If negative, it's a key in a map. If positive, it's an index in a sequence.
	Index int

	// Name of the element key. If empty, it's an index in a sequence.
	Name string
}

// NewRootPath returns a path to the root of the YAML document.
func NewRootPath() Path {
	return Path{Index: -1}
}

// NewPathWithPathSegment returns a new path based on the given path and segment.
func NewPathWithPathSegment(path Path, segment PathSegment) Path {
	// Init new segments slice and copy existing ones.
	segments := make([]PathSegment, len(path.Segments))
	copy(segments, path.Segments)

	return Path{
		Root:     path.Root,
		Index:    path.Index,
		Segments: append(segments, segment),
	}
}

// NewPathWithNamedSegment returns a new path based on the given path and name.
func NewPathWithNamedSegment(path Path, name string) Path {
	return NewPathWithPathSegment(
		path,
		PathSegment{Index: -1, Name: name},
	)
}

// NewPathWithIndexSegment returns a new path based on the given path and index.
func NewPathWithIndexSegment(path Path, index int) Path {
	return NewPathWithPathSegment(
		path,
		PathSegment{Index: index},
	)
}

func (p Path) String() string {
	var sections []string

	for _, segment := range p.Segments {
		switch {
		case segment.Name != "":
			sections = append(sections, segment.Name)

		case segment.Index >= 0:
			sections = append(sections, strconv.Itoa(segment.Index))
		}
	}

	return strings.Join(sections, ".")
}
