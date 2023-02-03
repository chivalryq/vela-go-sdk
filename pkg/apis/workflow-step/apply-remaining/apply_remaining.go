/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_remaining

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the ApplyRemainingSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApplyRemainingSpec{}

// ApplyRemainingSpec struct for ApplyRemainingSpec
type ApplyRemainingSpec struct {
	// Declare the name of the component
	Exceptions []string `json:"exceptions,omitempty"`
}

// NewApplyRemainingSpecWith instantiates a new ApplyRemainingSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyRemainingSpecWith() *ApplyRemainingSpec {
	this := ApplyRemainingSpec{}
	return &this
}

// NewApplyRemainingSpec instantiates a new ApplyRemainingSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyRemainingSpec() *ApplyRemainingSpec {
	this := ApplyRemainingSpec{}
	return &this
}

// NewApplyRemainingSpecs converts a list ApplyRemainingSpec pointers to objects.
// This is helpful when the SetApplyRemainingSpec requires a list of objects
func NewApplyRemainingSpecs(ps ...*ApplyRemainingSpec) []ApplyRemainingSpec {
	objs := []ApplyRemainingSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetExceptions returns the Exceptions field value if set, zero value otherwise.
func (o *ApplyRemainingWorkflowStep) GetExceptions() []string {
	if o == nil || utils.IsNil(o.Properties.Exceptions) {
		var ret []string
		return ret
	}
	return o.Properties.Exceptions
}

// GetExceptionsOk returns a tuple with the Exceptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyRemainingWorkflowStep) GetExceptionsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Exceptions) {
		return nil, false
	}
	return o.Properties.Exceptions, true
}

// HasExceptions returns a boolean if a field has been set.
func (o *ApplyRemainingWorkflowStep) HasExceptions() bool {
	if o != nil && !utils.IsNil(o.Properties.Exceptions) {
		return true
	}

	return false
}

// SetExceptions gets a reference to the given []string and assigns it to the exceptions field.
// Exceptions:  Declare the name of the component
func (o *ApplyRemainingWorkflowStep) SetExceptions(v []string) *ApplyRemainingWorkflowStep {
	o.Properties.Exceptions = v
	return o
}

func (o ApplyRemainingSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyRemainingSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Exceptions) {
		toSerialize["exceptions"] = o.Exceptions
	}
	return toSerialize, nil
}

type NullableApplyRemainingSpec struct {
	value *ApplyRemainingSpec
	isSet bool
}

func (v NullableApplyRemainingSpec) Get() *ApplyRemainingSpec {
	return v.value
}

func (v *NullableApplyRemainingSpec) Set(val *ApplyRemainingSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyRemainingSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyRemainingSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyRemainingSpec(val *ApplyRemainingSpec) *NullableApplyRemainingSpec {
	return &NullableApplyRemainingSpec{value: val, isSet: true}
}

func (v NullableApplyRemainingSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyRemainingSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ApplyRemainingType = "apply-remaining"

func init() {
	sdkcommon.RegisterWorkflowStep(ApplyRemainingType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(ApplyRemainingType, FromWorkflowSubStep)
}

type ApplyRemainingWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties ApplyRemainingSpec
}

func ApplyRemaining(name string) *ApplyRemainingWorkflowStep {
	a := &ApplyRemainingWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: ApplyRemainingType,
	}}
	return a
}

func (a *ApplyRemainingWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range a.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  a.Base.DependsOn,
		If:         a.Base.If,
		Inputs:     a.Base.Inputs,
		Meta:       a.Base.Meta,
		Name:       a.Base.Name,
		Outputs:    a.Base.Outputs,
		Properties: util.Object2RawExtension(a.Properties),
		SubSteps:   subSteps,
		Timeout:    a.Base.Timeout,
		Type:       ApplyRemainingType,
	}
	return res
}

func (a *ApplyRemainingWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*ApplyRemainingWorkflowStep, error) {
	var properties ApplyRemainingSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := a.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	a.Base.Name = from.Name
	a.Base.DependsOn = from.DependsOn
	a.Base.Inputs = from.Inputs
	a.Base.Outputs = from.Outputs
	a.Base.If = from.If
	a.Base.Timeout = from.Timeout
	a.Base.Meta = from.Meta
	a.Base.Type = ApplyRemainingType
	a.Properties = properties
	a.Base.SubSteps = subSteps
	return a, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	a := &ApplyRemainingWorkflowStep{}
	return a.FromWorkflowStep(from)
}

func (a *ApplyRemainingWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*ApplyRemainingWorkflowStep, error) {
	var properties ApplyRemainingSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	a.Base.Name = from.Name
	a.Base.DependsOn = from.DependsOn
	a.Base.Inputs = from.Inputs
	a.Base.Outputs = from.Outputs
	a.Base.If = from.If
	a.Base.Timeout = from.Timeout
	a.Base.Meta = from.Meta
	a.Base.Type = ApplyRemainingType
	a.Properties = properties
	return a, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	a := &ApplyRemainingWorkflowStep{}
	return a.FromWorkflowSubStep(from)
}

func (a *ApplyRemainingWorkflowStep) WorkflowStepName() string {
	return a.Base.Name
}

func (a *ApplyRemainingWorkflowStep) DefType() string {
	return ApplyRemainingType
}

func (a *ApplyRemainingWorkflowStep) If(_if string) *ApplyRemainingWorkflowStep {
	a.Base.If = _if
	return a
}

func (a *ApplyRemainingWorkflowStep) Alias(alias string) *ApplyRemainingWorkflowStep {
	a.Base.Meta.Alias = alias
	return a
}

func (a *ApplyRemainingWorkflowStep) Timeout(timeout string) *ApplyRemainingWorkflowStep {
	a.Base.Timeout = timeout
	return a
}

func (a *ApplyRemainingWorkflowStep) DependsOn(dependsOn []string) *ApplyRemainingWorkflowStep {
	a.Base.DependsOn = dependsOn
	return a
}

func (a *ApplyRemainingWorkflowStep) Inputs(input common.StepInputs) *ApplyRemainingWorkflowStep {
	a.Base.Inputs = input
	return a
}

func (a *ApplyRemainingWorkflowStep) Outputs(output common.StepOutputs) *ApplyRemainingWorkflowStep {
	a.Base.Outputs = output
	return a
}
