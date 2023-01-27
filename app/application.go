package app

import (
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	. "vela-go-sdk/api"
)

type BuildOps func(app Application)

type ApplicationBuilder struct {
	components []Component
	steps      []WorkflowStep
	policies   []Policy
}

func (a *ApplicationBuilder) WithWorkflowStep(step WorkflowStep) Application {
	a.steps = append(a.steps, step)
	return a
}

// WithComponent append a component by adding the given traits.
func (a *ApplicationBuilder) WithComponent(component Component) Application {
	a.components = append(a.components, component)
	return a
}

// New creates a new application with the given components.
func New(ops ...BuildOps) Application {
	app := &ApplicationBuilder{
		components: make([]Component, 0),
		steps:      make([]WorkflowStep, 0),
		policies:   make([]Policy, 0),
	}
	for _, op := range ops {
		op(app)
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
