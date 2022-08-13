package kubernetes_test

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	kubeutlis "github.com/kinkgo/kinkgo/util/kubernetes"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

const manifest = `apiVersion: v1
kind: ConfigMap
metadata:
  name: test-config
  namespace: default
data: 
  key1: value1
  key2: value2
`

const manifestList = `---
apiVersion: v1
kind: ConfigMap
metadata:
  name: test-config
  namespace: default
data: 
  key1: value1
  key2: value2
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: test-config-2
  namespace: default
data: 
  key1: value1
  key2: value2
`

var _ = Describe("Create", func() {
	var (
		ctx        context.Context
		fakeClient client.Client
	)

	BeforeEach(func() {
		ctx = context.Background()
		fakeClient = fake.NewClientBuilder().Build()
	})

	Context("from manifest", func() {
		It("should create resources", func() {
			Expect(kubeutlis.CreateFromManifest(ctx, fakeClient, manifest)).ShouldNot(HaveOccurred())

			cm := &corev1.ConfigMap{}
			cmName := types.NamespacedName{Name: "test-config", Namespace: "default"}

			Expect(fakeClient.Get(context.Background(), cmName, cm)).ShouldNot(HaveOccurred())
		})

		It("when manifest is empty should return error", func() {
			Expect(kubeutlis.CreateFromManifest(ctx, fakeClient, "")).Should(MatchError(kubeutlis.ErrEmptyManifest))
		})
	})

	Context("from manifest file", func() {
		It("should create resources", func() {
			Expect(kubeutlis.CreateFromFile(ctx, fakeClient, "./testdata/test-config-cm.yaml")).ShouldNot(HaveOccurred())

			cm := &corev1.ConfigMap{}
			cmName := types.NamespacedName{Name: "test-config-cm", Namespace: "default"}

			Expect(fakeClient.Get(context.Background(), cmName, cm)).ShouldNot(HaveOccurred())
		})
	})

	Context("from manifest with modifier functions", func() {
		It("should create resources", func() {
			err := kubeutlis.CreateFromManifestWithModifications(
				ctx,
				fakeClient,
				manifest,
				kubeutlis.WithName("just-another-config"),
				kubeutlis.WithNamespace("kube-system"),
			)
			Expect(err).ShouldNot(HaveOccurred())

			cm := &corev1.ConfigMap{}
			cmName := types.NamespacedName{Name: "just-another-config", Namespace: "kube-system"}

			Expect(fakeClient.Get(context.Background(), cmName, cm)).ShouldNot(HaveOccurred())
		})

		It("when manifest is empty should return error", func() {
			err := kubeutlis.CreateFromManifestWithModifications(
				ctx,
				fakeClient,
				"",
				kubeutlis.WithName("just-another-config"),
				kubeutlis.WithNamespace("kube-system"),
			)
			Expect(err).Should(MatchError(kubeutlis.ErrEmptyManifest))
		})

		It("when manifest is a list should return error", func() {
			err := kubeutlis.CreateFromManifestWithModifications(
				ctx,
				fakeClient,
				manifestList,
				kubeutlis.WithName("just-another-config"),
				kubeutlis.WithNamespace("kube-system"),
			)
			Expect(err).Should(MatchError(kubeutlis.ErrNotSupported))
		})
	})

	Context("from manifest file with modifier functions", func() {
		It("should create resources", func() {
			err := kubeutlis.CreateFromFileWithModifications(
				ctx,
				fakeClient,
				"./testdata/test-config-cm.yaml",
				kubeutlis.WithName("just-another-config"),
				kubeutlis.WithNamespace("kube-system"),
			)
			Expect(err).ShouldNot(HaveOccurred())

			cm := &corev1.ConfigMap{}
			cmName := types.NamespacedName{Name: "just-another-config", Namespace: "kube-system"}

			Expect(fakeClient.Get(context.Background(), cmName, cm)).ShouldNot(HaveOccurred())
		})
	})
})
