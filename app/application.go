package app

import (
	. "vela-go-sdk/api"
)

type BuildOps func(app Application)

type ApplicationBuilder struct {
	components map[string]Component
	steps      []WorkflowStep
}

func (a *ApplicationBuilder) WithWorkflowStep(step WorkflowStep) Application {
	a.steps = append(a.steps, step)
	return a
}

// WithComponent append a component by adding the given traits.
func (a *ApplicationBuilder) WithComponent(component Component) Application {
	a.components[component.Name()] = component
	return a
}

// New creates a new application with the given components.
func New(ops ...BuildOps) Application {
	app := &ApplicationBuilder{
		components: make(map[string]Component),
	}
	for _, op := range ops {
		op(app)
	}
	return app
}
