package yamlpath_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	yamlv3 "gopkg.in/yaml.v3"

	"github.com/kinkgo/kinkgo/util/yamlpath"
)

var node yamlv3.Node

var _ = BeforeSuite(func() {
	data, err := os.ReadFile("./testdata/path_parse_test.yaml")
	Expect(err).ShouldNot(HaveOccurred())

	Expect(yamlv3.Unmarshal(data, &node)).ShouldNot(HaveOccurred())
})

var _ = DescribeTable("Path Parse",
	[]TableEntry{
		Entry("should parse existing leaf value", &node, "data.key1", "data.key1", true, nil),
		Entry("should parse existing slice index", &node, "data.list.0", "data.list.0", true, nil),
		Entry("should parse existing map", &node, "data.list-of-map", "data.list-of-map", true, nil),
		Entry("should parse existing slice of maps", &node, "data.list-of-map.0.key1", "data.list-of-map.0.key1", true, nil),
		Entry("should parse non-existing path and return not found", &node, "data.not-exist", "data.not-exist", false, nil),
		Entry("should return error for nil node", nil, "data.key1", "", false, yamlpath.ErrEmptyYamlNode),
		Entry("should return error for invalid index", &node, "data.list-of-map.11", "", false, yamlpath.ErrInvalidPathString),
	},
	func(node *yamlv3.Node, pathString, expectedPath string, expectedOk bool, expectedErr error) {

		path, ok, err := yamlpath.ParsePath(pathString, node)
		Expect(ok).Should(Equal(expectedOk))
		Expect(path.String()).Should(Equal(expectedPath))

		if expectedErr != nil {
			Expect(err).Should(MatchError(expectedErr))
		} else {
			Expect(err).ShouldNot(HaveOccurred())
		}
	})
