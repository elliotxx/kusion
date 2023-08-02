package generators

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kusionstack.io/kusion/pkg/models"
)

type namespaceGenerator struct {
	projectName string
}

func NewNamespaceGenerator(projectName string) (Generator, error) {
	if len(projectName) == 0 {
		return nil, fmt.Errorf("project name must not be empty")
	}

	return &namespaceGenerator{
		projectName: projectName,
	}, nil
}

func NewNamespaceGeneratorFunc(projectName string) NewGeneratorFunc {
	return func() (Generator, error) {
		return NewNamespaceGenerator(projectName)
	}
}

func (g *namespaceGenerator) Generate(spec *models.Spec) error {
	if spec.Resources == nil {
		spec.Resources = make(models.Resources, 0)
	}

	ns := &v1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: v1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{Name: g.projectName},
	}

	return appendToSpec(
		kubernetesResourceID(ns.TypeMeta, ns.ObjectMeta),
		ns,
		spec,
	)
}
