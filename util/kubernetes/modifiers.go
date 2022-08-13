package kubernetes

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type ManifestModifierFn func(unstructured.Unstructured) error

// WithNamespace sets the namespace of the object.
func WithNamespace(namespace string) ManifestModifierFn {
	return func(m unstructured.Unstructured) error {
		return unstructured.SetNestedField(m.Object, namespace, "metadata", "namespace")
	}
}

// WithName sets the name of the object.
func WithName(name string) ManifestModifierFn {
	return func(m unstructured.Unstructured) error {
		return unstructured.SetNestedField(m.Object, name, "metadata", "name")
	}
}

// WithAnnotations sets the annotations of the object.
func WithAnnotations(annotations map[string]string) ManifestModifierFn {
	return func(m unstructured.Unstructured) error {
		for key, value := range annotations {
			if err := unstructured.SetNestedField(m.Object, value, "metadata", "annotations", key); err != nil {
				return err
			}
		}

		return nil
	}
}

// WithLabels sets the labels of the object.
func WithLabels(labels map[string]string) ManifestModifierFn {
	return func(m unstructured.Unstructured) error {
		for key, value := range labels {
			if err := unstructured.SetNestedField(m.Object, value, "metadata", "labels", key); err != nil {
				return err
			}
		}

		return nil
	}
}
