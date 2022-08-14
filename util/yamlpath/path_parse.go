package yamlpath

import (
	"fmt"
	"strconv"
	"strings"

	yamlv3 "gopkg.in/yaml.v3"
)

// ParsePath parses the given path and returns a path object.
func ParsePath(path string, node *yamlv3.Node) (Path, bool, error) {
	if node == nil {
		return Path{}, false, ErrEmptyYamlNode
	}

	if node.Kind != yamlv3.DocumentNode {
		return Path{}, false, fmt.Errorf("%w: %s", ErrInvalidType, "path is not a document")
	}

	// path segments
	segments := make([]PathSegment, 0)

	// pointer of the current node in the path
	pointer := node.Content[0]

	ok := true

	for _, segment := range strings.Split(path, ".") {
		// if the pointer is nil, previous segment of the path is not found. This check assumes that rest of the path
		// segments are assumed to type of map.  d
		if pointer == nil {
			ok = false
			segments = append(segments, PathSegment{Index: -1, Name: segment})
			continue
		}

		switch pointer.Kind {
		case yamlv3.MappingNode:
			if value, err := getValueByKey(pointer, segment); err == nil {
				pointer = value
			} else {
				pointer = nil
				ok = false
			}
			segments = append(segments, PathSegment{Index: -1, Name: segment})
		case yamlv3.SequenceNode:
			list := pointer.Content

			if id, err := strconv.Atoi(segment); err == nil {
				if id < 0 || id >= len(list) {
					return Path{}, false, fmt.Errorf("%w: index %d out of range", ErrInvalidPathString, id)
				}

				pointer = list[id]
				segments = append(segments, PathSegment{Index: id})
			} else {
				pointer = nil
				ok = false
				segments = append(segments, PathSegment{Name: segment})
			}
		default:
			return Path{}, false, fmt.Errorf("%w: %s", ErrInvalidType, "path is not a mapping or sequence")
		}
	}

	return Path{Root: node, Index: 0, Segments: segments}, ok, nil
}

// getValueByKey returns the value for a given key in a provided mapping node,
// or nil with an error if there is no such entry. This is comparable to getting
// a value from a map with `foobar[key]`.
func getValueByKey(mappingNode *yamlv3.Node, key string) (*yamlv3.Node, error) {
	for i := 0; i < len(mappingNode.Content); i += 2 {
		k, v := mappingNode.Content[i], mappingNode.Content[i+1]
		if k.Value == key {
			return v, nil
		}
	}

	return nil, fmt.Errorf("key %s not found", key)
}
