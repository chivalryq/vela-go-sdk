package app

import (
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/pkg/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	. "vela-go-sdk/pkg/apis"
)

type ApplicationBuilder struct {
	name        string
	namespace   string
	labels      map[string]string
	annotations map[string]string

	components   []Component
	steps        []WorkflowStep
	policies     []Policy
	workflowMode v1beta1.WorkflowExecuteMode
}

func (a *ApplicationBuilder) WithWorkflowSteps(step ...WorkflowStep) Application {
	a.steps = append(a.steps, step...)
	return a
}

// WithComponents append components to application
func (a *ApplicationBuilder) WithComponents(component ...Component) Application {
	a.components = append(a.components, component...)
	return a
}

// WithPolicies append policies to application
func (a *ApplicationBuilder) WithPolicies(policy ...Policy) Application {
	a.policies = append(a.policies, policy...)
	return a
}

// WithWorkflowMode set the workflow mode of application
func (a *ApplicationBuilder) WithWorkflowMode(steps, subSteps common.WorkflowMode) Application {
	a.workflowMode.Steps = steps
	a.workflowMode.SubSteps = subSteps
	return a
}

func (a *ApplicationBuilder) Name(name string) Application {
	a.name = name
	return a
}

func (a *ApplicationBuilder) Namespace(namespace string) Application {
	a.namespace = namespace
	return a
}

func (a *ApplicationBuilder) Labels(labels map[string]string) Application {
	a.labels = labels
	return a
}

func (a *ApplicationBuilder) Annotations(annotations map[string]string) Application {
	a.annotations = annotations
	return a
}

// New creates a new application with the given components.
func New() Application {
	app := &ApplicationBuilder{
		components: make([]Component, 0),
		steps:      make([]WorkflowStep, 0),
		policies:   make([]Policy, 0),
	}
	return app
}

func (a *ApplicationBuilder) Build() v1beta1.Application {
	components := make([]common.ApplicationComponent, 0, len(a.components))
	for _, component := range a.components {
		components = append(components, component.Build())
	}
	steps := make([]v1beta1.WorkflowStep, 0, len(a.steps))
	for _, step := range a.steps {
		steps = append(steps, step.Build())
	}
	policies := make([]v1beta1.AppPolicy, 0)
	for _, policy := range a.policies {
		policies = append(policies, policy.Build())
	}

	res := v1beta1.Application{
		TypeMeta: v1.TypeMeta{
			Kind:       v1beta1.ApplicationKind,
			APIVersion: v1beta1.SchemeGroupVersion.String(),
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      a.name,
			Namespace: a.namespace,
		},
		Spec: v1beta1.ApplicationSpec{
			Components: components,
			Workflow: &v1beta1.Workflow{
				Steps: steps,
			},
			Policies: policies,
		},
	}
	return res
}

func FromK8sObject(app *v1beta1.Application) (Application, error) {
	a := &ApplicationBuilder{}
	a.Name(app.Name)
	a.Namespace(app.Namespace)
	for _, comp := range app.Spec.Components {
		c, err := CompFromK8sObject(&comp)
		if err != nil {
			return nil, errors.Wrap(err, "convert component from k8s object")
		}
		a.WithComponents(c)
	}
	return a, nil
}

func CompFromK8sObject(component *common.ApplicationComponent) (Component, error) {

}
