package kubernetes_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/cluster-api/util/yaml"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	kubeutils "github.com/kinkgo/kinkgo/util/kubernetes"
)

var _ = Describe("Kubernetes modifiers", func() {
	It("should modify manifests", func() {
		modifiers := []kubeutils.ManifestModifierFn{
			kubeutils.WithNamespace("default"),
			kubeutils.WithName("test-config"),
			kubeutils.WithAnnotations(
				map[string]string{
					"annotation-1": "value1",
					"annotation-2": "value2",
				},
			),
			kubeutils.WithLabels(
				map[string]string{
					"label-1": "value1",
					"label-2": "value2",
				},
			),
		}

		objects, err := yaml.ToUnstructured([]byte("kind: ConfigMap\napiVersion: v1"))
		Expect(err).NotTo(HaveOccurred())
		Expect(objects).To(HaveLen(1))

		object := objects[0]

		for _, modifier := range modifiers {
			Expect(modifier(object)).ShouldNot(HaveOccurred())
		}

		validateField(object.Object, "ConfigMap", "kind")
		validateField(object.Object, "v1", "apiVersion")
		validateField(object.Object, "test-config", "metadata", "name")
		validateField(object.Object, "default", "metadata", "namespace")
		validateField(object.Object, "value1", "metadata", "annotations", "annotation-1")
		validateField(object.Object, "value2", "metadata", "annotations", "annotation-2")
		validateField(object.Object, "value1", "metadata", "labels", "label-1")
		validateField(object.Object, "value2", "metadata", "labels", "label-2")

	})
})

func validateField(object map[string]interface{}, expected string, nestedKeys ...string) {
	value, ok, err := unstructured.NestedString(object, nestedKeys...)
	Expect(err).ToNot(HaveOccurred())
	Expect(ok).To(BeTrue())
	Expect(value).To(Equal(expected))
}
