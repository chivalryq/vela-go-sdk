/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_deployment

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the ApplyDeploymentSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApplyDeploymentSpec{}

// ApplyDeploymentSpec struct for ApplyDeploymentSpec
type ApplyDeploymentSpec struct {
	cmd   []string `json:"cmd,omitempty"`
	image string   `json:"image"`
}

// NewApplyDeploymentSpecWith instantiates a new ApplyDeploymentSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyDeploymentSpecWith(image string) *ApplyDeploymentSpec {
	this := ApplyDeploymentSpec{}
	this.image = image
	return &this
}

// NewApplyDeploymentSpec instantiates a new ApplyDeploymentSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyDeploymentSpec() *ApplyDeploymentSpec {
	this := ApplyDeploymentSpec{}
	return &this
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *ApplyDeploymentWorkflowStep) GetCmd() []string {
	if o == nil || utils.IsNil(o.Properties.cmd) {
		var ret []string
		return ret
	}
	return o.Properties.cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyDeploymentWorkflowStep) GetCmdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.cmd) {
		return nil, false
	}
	return o.Properties.cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *ApplyDeploymentWorkflowStep) HasCmd() bool {
	if o != nil && !utils.IsNil(o.Properties.cmd) {
		return true
	}

	return false
}

// Cmd gets a reference to the given []string and assigns it to the cmd field.
// cmd:
func (o *ApplyDeploymentWorkflowStep) Cmd(v []string) *ApplyDeploymentWorkflowStep {
	o.Properties.cmd = v
	return o
}

// GetImage returns the Image field value
func (o *ApplyDeploymentWorkflowStep) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *ApplyDeploymentWorkflowStep) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.image, true
}

// Image sets field value
func (o *ApplyDeploymentWorkflowStep) Image(v string) *ApplyDeploymentWorkflowStep {
	o.Properties.image = v
	return o
}

func (o ApplyDeploymentSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyDeploymentSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.cmd) {
		toSerialize["cmd"] = o.cmd
	}
	toSerialize["image"] = o.image
	return toSerialize, nil
}

type NullableApplyDeploymentSpec struct {
	value *ApplyDeploymentSpec
	isSet bool
}

func (v NullableApplyDeploymentSpec) Get() *ApplyDeploymentSpec {
	return v.value
}

func (v *NullableApplyDeploymentSpec) Set(val *ApplyDeploymentSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyDeploymentSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyDeploymentSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyDeploymentSpec(val *ApplyDeploymentSpec) *NullableApplyDeploymentSpec {
	return &NullableApplyDeploymentSpec{value: val, isSet: true}
}

func (v NullableApplyDeploymentSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyDeploymentSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ApplyDeploymentType = "apply-deployment"

func init() {
	sdkcommon.RegisterWorkflowStep(ApplyDeploymentType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(ApplyDeploymentType, FromWorkflowSubStep)
}

type ApplyDeploymentWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties ApplyDeploymentSpec
}

func ApplyDeployment(name string) *ApplyDeploymentWorkflowStep {
	a := &ApplyDeploymentWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
	}}
	return a
}

func (a *ApplyDeploymentWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range a.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties})
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
		Type:       ApplyDeploymentType,
	}
	return res
}

func (a *ApplyDeploymentWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*ApplyDeploymentWorkflowStep, error) {
	var properties ApplyDeploymentSpec
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
	a.Properties = properties
	a.Base.SubSteps = subSteps
	return a, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	a := &ApplyDeploymentWorkflowStep{}
	return a.FromWorkflowStep(from)
}

func (a *ApplyDeploymentWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*ApplyDeploymentWorkflowStep, error) {
	var properties ApplyDeploymentSpec
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
	a.Properties = properties
	return a, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	a := &ApplyDeploymentWorkflowStep{}
	return a.FromWorkflowSubStep(from)
}

func (a *ApplyDeploymentWorkflowStep) If(_if string) *ApplyDeploymentWorkflowStep {
	a.Base.If = _if
	return a
}

func (a *ApplyDeploymentWorkflowStep) Alias(alias string) *ApplyDeploymentWorkflowStep {
	a.Base.Meta.Alias = alias
	return a
}

func (a *ApplyDeploymentWorkflowStep) Timeout(timeout string) *ApplyDeploymentWorkflowStep {
	a.Base.Timeout = timeout
	return a
}

func (a *ApplyDeploymentWorkflowStep) DependsOn(dependsOn []string) *ApplyDeploymentWorkflowStep {
	a.Base.DependsOn = dependsOn
	return a
}

func (a *ApplyDeploymentWorkflowStep) Inputs(input common.StepInputs) *ApplyDeploymentWorkflowStep {
	a.Base.Inputs = input
	return a
}

func (a *ApplyDeploymentWorkflowStep) Outputs(output common.StepOutputs) *ApplyDeploymentWorkflowStep {
	a.Base.Outputs = output
	return a
}

func (a *ApplyDeploymentWorkflowStep) DefName() string {
	return a.Base.Name
}

func (a *ApplyDeploymentWorkflowStep) DefType() string {
	return ApplyDeploymentType
}
