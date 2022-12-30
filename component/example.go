package component

import (
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	. "vela-go-sdk/api"
)

const webserviceType = "webservice"

var _ Component = &WebService{}

type WebService struct {
	Base  ComponentBase
	Props WebserviceSpec
}

type WebserviceSpec struct {
	Image string
}

func Webservice(name string) *WebService {
	comp := WebService{
		Base: ComponentBase{
			Name: name,
		},
	}
	return &comp
}

func (e *WebService) Traits(traits ...Trait) *WebService {
	e.Base.Traits = append(e.Base.Traits, traits...)
	return e
}

func (e *WebService) Name() string {
	return e.Base.Name
}

// Build convert a component to the kubevela-core-api type.
func (e *WebService) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range e.Base.Traits {
		traits = append(traits, trait.Build())
	}
	comp := common.ApplicationComponent{
		Name:       e.Base.Name,
		Type:       webserviceType,
		Properties: util.Object2RawExtension(e.Props),
		DependsOn:  e.Base.DependsOn,
		Inputs:     e.Base.Inputs,
		Outputs:    e.Base.Outputs,
		Traits:     traits,
	}
	return comp
}

// Properties modify the whole properties directly
// todo maybe add patchProperties?
func (e *WebService) Properties(props WebserviceSpec) *WebService {
	e.Props = props
	return e
}

// Modify single property //

// Image will set the image of the component.
func (e *WebService) Image(image string) *WebService {
	e.Props.Image = image
	return e
}

// Template functions to make it easier to use, deal with:
// DependsOn
// Inputs/Outputs

// DependsOn will mark the component depends on the given components.
func (e *WebService) DependsOn(dependsOn []string) *WebService {
	if e.Base.DependsOn == nil {
		e.Base.DependsOn = make([]string, 0)
	}
	e.Base.DependsOn = dependsOn
	return e
}

// Inputs will add the given inputs to the component.
func (e *WebService) Inputs(inputs common.StepInputs) *WebService {
	e.Base.Inputs = inputs
	return e
}

// Outputs will add the given outputs to the component.
func (e *WebService) Outputs(outputs common.StepOutputs) *WebService {
	e.Base.Outputs = outputs
	return e
}

// Extension approach //

type WebserviceOps func(builder *WebService) *WebService

func (e *WebService) Do(ops ...WebserviceOps) *WebService {
	for _, op := range ops {
		op(e)
	}
	return e
}
