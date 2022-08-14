package yamlpath

import yamlv3 "gopkg.in/yaml.v3"

// TraverseFn is the function called for each node in the YAML document.
type TraverseFn func(path Path, parent *yamlv3.Node, leaf *yamlv3.Node)

// TraverseTree traverses the YAML document starting at root and calls fn for each node.
func TraverseTree(path Path, parent, leaf *yamlv3.Node, leafFn TraverseFn) {
	switch leaf.Kind {
	case yamlv3.DocumentNode:
		TraverseTree(path, leaf, leaf.Content[0], leafFn)
	case yamlv3.SequenceNode:
		for idx, entry := range leaf.Content {
			TraverseTree(NewPathWithIndexSegment(path, idx), leaf, entry, leafFn)
		}
	case yamlv3.MappingNode:
		for idx := 0; idx < len(leaf.Content); idx += 2 {
			k, v := leaf.Content[idx], leaf.Content[idx+1]
			TraverseTree(NewPathWithNamedSegment(path, k.Value), leaf, v, leafFn)
		}
	default:
		leafFn(path, parent, leaf)
	}
}
