package yamlpath_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	yamlv3 "gopkg.in/yaml.v3"

	"github.com/kinkgo/kinkgo/util/yamlpath"
)

var _ = Describe("Traverse", func() {
	var node yamlv3.Node

	BeforeEach(func() {
		data, err := os.ReadFile("./testdata/traverse_test.yaml")
		Expect(err).ShouldNot(HaveOccurred())

		Expect(yamlv3.Unmarshal(data, &node)).ShouldNot(HaveOccurred())
	})

	It("should traverse tree", func() {
		leafs := map[string]string{
			"data.key1":                 "value1-from-config-cm",
			"data.key2":                 "value2-from-config-cm",
			"data.list.0":               "item1",
			"data.list-of-map.0.key1":   "value1",
			"data.list-of-map.0.key2.0": "value2",
			"data.list-of-map.0.key2.1": "value3",
			"data.list-of-map.1.key1":   "value3",
		}
		yamlpath.TraverseTree(yamlpath.NewRootPath(), nil, &node, func(path yamlpath.Path, parent *yamlv3.Node, leaf *yamlv3.Node) {
			value := leafs[path.String()]
			Expect(parent).ShouldNot(BeNil())
			Expect(leaf.Value).Should(Equal(value), "path: %s expected: %s actual: %s", path.String(), value, leaf.Value)
		})
	})
})
