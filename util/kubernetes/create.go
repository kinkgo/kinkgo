package kubernetes

import (
	"context"
	"fmt"
	"os"

	"sigs.k8s.io/cluster-api/util/yaml"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CreateFromManifest creates resources from given manifest string.
func CreateFromManifest(ctx context.Context, c client.Client, manifest string) error {
	if manifest == "" {
		return ErrEmptyManifest
	}

	objects, err := yaml.ToUnstructured([]byte(manifest))
	if err != nil {
		return err
	}

	for _, object := range objects {
		if err = c.Create(ctx, &object); err != nil {
			return err
		}
	}

	return nil
}

// CreateFromFile creates resources from given file.
func CreateFromFile(ctx context.Context, c client.Client, filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	return CreateFromManifest(ctx, c, string(data))
}

// CreateFromManifestWithModifications updates the given manifest with the given modifications and creates resources from it.
//
// For manifest modifications:
// - The modifications are applied in the order they are given.
// - If any of the modifications fails, the whole operation fails.
// - Manifest modifications are not supported for lists.
func CreateFromManifestWithModifications(ctx context.Context, c client.Client, manifest string, modifiers ...ManifestModifierFn) error {
	if manifest == "" {
		return ErrEmptyManifest
	}

	objects, err := yaml.ToUnstructured([]byte(manifest))
	if err != nil {
		return err
	}

	if len(objects) > 1 {
		return fmt.Errorf("%w: manifest contains more than one object", ErrNotSupported)
	}

	object := objects[0]

	// Apply the modifications for each object.
	for _, modifier := range modifiers {
		if err = modifier(object); err != nil {
			return err
		}
	}

	err = c.Create(ctx, &object)
	if err != nil {
		return err
	}

	return nil
}

// CreateFromFileWithModifications creates resources from given file, and updates the given manifest with the given modifications.
//
// For manifest modifications:
// - The modifications are applied in the order they are given.
// - If any of the modifications fails, the whole operation fails.
// - Manifest modifications are not supported for lists.
func CreateFromFileWithModifications(ctx context.Context, c client.Client, filePath string, modifiers ...ManifestModifierFn) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	return CreateFromManifestWithModifications(ctx, c, string(data), modifiers...)
}
