/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_component

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"vela-go-sdk/pkg/apis"
	"vela-go-sdk/pkg/apis/utils"
)

// checks if the ApplyComponentSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApplyComponentSpec{}

// ApplyComponentSpec struct for ApplyComponentSpec
type ApplyComponentSpec struct {
	// Specify the cluster
	cluster string `json:"cluster"`
	// Specify the component name to apply
	component string `json:"component"`
}

// NewApplyComponentSpecWith instantiates a new ApplyComponentSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyComponentSpecWith(cluster string, component string) *ApplyComponentSpec {
	this := ApplyComponentSpec{}
	this.cluster = cluster
	this.component = component
	return &this
}

// NewApplyComponentSpec instantiates a new ApplyComponentSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyComponentSpec() *ApplyComponentSpec {
	this := ApplyComponentSpec{}
	var cluster string = ""
	this.cluster = cluster
	return &this
}

// GetCluster returns the Cluster field value
func (o *ApplyComponentWorkflowStep) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *ApplyComponentWorkflowStep) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.cluster, true
}

// Cluster sets field value
func (o *ApplyComponentWorkflowStep) Cluster(v string) *ApplyComponentWorkflowStep {
	o.Properties.cluster = v
	return o
}

// GetComponent returns the Component field value
func (o *ApplyComponentWorkflowStep) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ApplyComponentWorkflowStep) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.component, true
}

// Component sets field value
func (o *ApplyComponentWorkflowStep) Component(v string) *ApplyComponentWorkflowStep {
	o.Properties.component = v
	return o
}

func (o ApplyComponentSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyComponentSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster"] = o.cluster
	toSerialize["component"] = o.component
	return toSerialize, nil
}

type NullableApplyComponentSpec struct {
	value *ApplyComponentSpec
	isSet bool
}

func (v NullableApplyComponentSpec) Get() *ApplyComponentSpec {
	return v.value
}

func (v *NullableApplyComponentSpec) Set(val *ApplyComponentSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyComponentSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyComponentSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyComponentSpec(val *ApplyComponentSpec) *NullableApplyComponentSpec {
	return &NullableApplyComponentSpec{value: val, isSet: true}
}

func (v NullableApplyComponentSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyComponentSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ApplyComponentType = "apply-component"

type ApplyComponentWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties ApplyComponentSpec
}

func ApplyComponent(name string) *ApplyComponentWorkflowStep {
	a := &ApplyComponentWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
	}}
	return a
}

func (a *ApplyComponentWorkflowStep) Build() v1beta1.WorkflowStep {
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
		Type:       ApplyComponentType,
	}
	return res
}

func (a *ApplyComponentWorkflowStep) If(_if string) *ApplyComponentWorkflowStep {
	a.Base.If = _if
	return a
}

func (a *ApplyComponentWorkflowStep) Alias(alias string) *ApplyComponentWorkflowStep {
	a.Base.Meta.Alias = alias
	return a
}

func (a *ApplyComponentWorkflowStep) Timeout(timeout string) *ApplyComponentWorkflowStep {
	a.Base.Timeout = timeout
	return a
}

func (a *ApplyComponentWorkflowStep) DependsOn(dependsOn []string) *ApplyComponentWorkflowStep {
	a.Base.DependsOn = dependsOn
	return a
}

func (a *ApplyComponentWorkflowStep) Inputs(input common.StepInputs) *ApplyComponentWorkflowStep {
	a.Base.Inputs = input
	return a
}

func (a *ApplyComponentWorkflowStep) Outputs(output common.StepOutputs) *ApplyComponentWorkflowStep {
	a.Base.Outputs = output
	return a
}

func (a *ApplyComponentWorkflowStep) Name() string {
	return a.Base.Name
}

func (a *ApplyComponentWorkflowStep) Type() string {
	return ApplyComponentType
}
