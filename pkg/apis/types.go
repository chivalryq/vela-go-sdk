package apis

import (
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
)

type Application interface {
	Name(name string) Application
	Namespace(namespace string) Application
	Labels(labels map[string]string) Application
	Annotations(annotations map[string]string) Application

	WithComponents(component ...Component) Application
	WithWorkflowSteps(step ...WorkflowStep) Application
	WithPolicies(policy ...Policy) Application
	WithWorkflowMode(steps, subSteps common.WorkflowMode) Application
}

type Component interface {
	Name() string
	Type() string
	Build() common.ApplicationComponent
}

type Trait interface {
	Type() string
	Build() common.ApplicationTrait
}

type WorkflowStep interface {
	Name() string
	Type() string
	Build() v1beta1.WorkflowStep
}

type Policy interface {
	Type() string
	Build() v1beta1.AppPolicy
}

type WorkflowSubStep interface {
	Build() common.WorkflowSubStep
}

type ComponentBase struct {
	Name      string
	Type      string
	DependsOn []string
	Inputs    common.StepInputs
	Outputs   common.StepOutputs
	Traits    []Trait
}

type TraitBase struct{}

type WorkflowSubStepBase struct {
	Name      string
	Type      string
	Meta      *common.WorkflowStepMeta
	If        string
	Timeout   string
	DependsOn []string
	Inputs    common.StepInputs
	Outputs   common.StepOutputs
}

type WorkflowStepBase struct {
	Name      string
	Type      string
	Meta      *common.WorkflowStepMeta
	SubSteps  []WorkflowStep
	If        string
	Timeout   string
	DependsOn []string
	Inputs    common.StepInputs
	Outputs   common.StepOutputs
}

type PolicyBase struct {
	Name string
}
