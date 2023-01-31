/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package step_group

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the StepGroupSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &StepGroupSpec{}

// StepGroupSpec struct for StepGroupSpec
type StepGroupSpec struct {
}

// NewStepGroupSpecWith instantiates a new StepGroupSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStepGroupSpecWith() *StepGroupSpec {
	this := StepGroupSpec{}
	return &this
}

// NewStepGroupSpec instantiates a new StepGroupSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStepGroupSpec() *StepGroupSpec {
	this := StepGroupSpec{}
	return &this
}

func (o StepGroupSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StepGroupSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableStepGroupSpec struct {
	value *StepGroupSpec
	isSet bool
}

func (v NullableStepGroupSpec) Get() *StepGroupSpec {
	return v.value
}

func (v *NullableStepGroupSpec) Set(val *StepGroupSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableStepGroupSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableStepGroupSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStepGroupSpec(val *StepGroupSpec) *NullableStepGroupSpec {
	return &NullableStepGroupSpec{value: val, isSet: true}
}

func (v NullableStepGroupSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStepGroupSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const StepGroupType = "step-group"

type StepGroupWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties StepGroupSpec
}

func StepGroup(name string) *StepGroupWorkflowStep {
	s := &StepGroupWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
	}}
	return s
}

func (s *StepGroupWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range s.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  s.Base.DependsOn,
		If:         s.Base.If,
		Inputs:     s.Base.Inputs,
		Meta:       s.Base.Meta,
		Name:       s.Base.Name,
		Outputs:    s.Base.Outputs,
		Properties: util.Object2RawExtension(s.Properties),
		SubSteps:   subSteps,
		Timeout:    s.Base.Timeout,
		Type:       StepGroupType,
	}
	return res
}

func (s *StepGroupWorkflowStep) If(_if string) *StepGroupWorkflowStep {
	s.Base.If = _if
	return s
}

func (s *StepGroupWorkflowStep) Alias(alias string) *StepGroupWorkflowStep {
	s.Base.Meta.Alias = alias
	return s
}

func (s *StepGroupWorkflowStep) Timeout(timeout string) *StepGroupWorkflowStep {
	s.Base.Timeout = timeout
	return s
}

func (s *StepGroupWorkflowStep) DependsOn(dependsOn []string) *StepGroupWorkflowStep {
	s.Base.DependsOn = dependsOn
	return s
}

func (s *StepGroupWorkflowStep) Inputs(input common.StepInputs) *StepGroupWorkflowStep {
	s.Base.Inputs = input
	return s
}

func (s *StepGroupWorkflowStep) Outputs(output common.StepOutputs) *StepGroupWorkflowStep {
	s.Base.Outputs = output
	return s
}

func (s *StepGroupWorkflowStep) AddSubStep(subStep apis.WorkflowStep) *StepGroupWorkflowStep {
	s.Base.SubSteps = append(s.Base.SubSteps, subStep)
	return s
}

func (s *StepGroupWorkflowStep) Name() string {
	return s.Base.Name
}

func (s *StepGroupWorkflowStep) Type() string {
	return StepGroupType
}
