package api

import (
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
)

type Application interface {
	WithComponent(component Component) Application
	WithWorkflowStep(step WorkflowStep) Application
}

type Component interface {
	Name() string
	Build() common.ApplicationComponent
}

type Trait interface {
	Build() common.ApplicationTrait
}

type WorkflowStep interface {
}

type ComponentBase struct {
	Name      string
	Type      string
	DependsOn []string
	Inputs    common.StepInputs
	Outputs   common.StepOutputs
	Traits    []Trait
}
