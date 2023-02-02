package convert

import (
	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	common2 "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/pkg/errors"
)

func FromComponent(component *common.ApplicationComponent) (apis.Component, error) {
	build, ok := common2.ComponentsBuilders[component.Type]
	if !ok {
		return nil, errors.Errorf("no component type %s registered", component.Type)
	}
	return build(*component)
}

func FromWorkflowStep(step *v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	build, ok := common2.WorkflowStepsBuilders[step.Type]
	if !ok {
		return nil, errors.Errorf("no workflow step type %s registered", step.Type)
	}
	return build(*step)
}

func FromPolicy(policy *v1beta1.AppPolicy) (apis.Policy, error) {
	build, ok := common2.PoliciesBuilders[policy.Type]
	if !ok {
		return nil, errors.Errorf("no policy type %s registered", policy.Type)
	}
	return build(*policy)
}

func FromTrait(trait *common.ApplicationTrait) (apis.Trait, error) {
	build, ok := common2.TraitBuilders[trait.Type]
	if !ok {
		return nil, errors.Errorf("no trait type %s registered", trait.Type)
	}
	return build(*trait)
}
