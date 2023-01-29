/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package share_cloud_resource

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"vela-go-sdk/pkg/apis"
	"vela-go-sdk/pkg/apis/utils"
)

// checks if the ShareCloudResourceSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ShareCloudResourceSpec{}

// ShareCloudResourceSpec struct for ShareCloudResourceSpec
type ShareCloudResourceSpec struct {
	// Declare the name of the env in policy
	env string `json:"env"`
	// Declare the location to bind
	placements []Placements `json:"placements"`
	// Declare the name of the env-binding policy, if empty, the first env-binding policy will be used
	policy string `json:"policy"`
}

// NewShareCloudResourceSpecWith instantiates a new ShareCloudResourceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShareCloudResourceSpecWith(env string, placements []Placements, policy string) *ShareCloudResourceSpec {
	this := ShareCloudResourceSpec{}
	this.env = env
	this.placements = placements
	this.policy = policy
	return &this
}

// NewShareCloudResourceSpec instantiates a new ShareCloudResourceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShareCloudResourceSpec() *ShareCloudResourceSpec {
	this := ShareCloudResourceSpec{}
	var policy string = ""
	this.policy = policy
	return &this
}

// GetEnv returns the Env field value
func (o *ShareCloudResourceWorkflowStep) GetEnv() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *ShareCloudResourceWorkflowStep) GetEnvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.env, true
}

// Env sets field value
func (o *ShareCloudResourceWorkflowStep) Env(v string) *ShareCloudResourceWorkflowStep {
	o.Properties.env = v
	return o
}

// GetPlacements returns the Placements field value
func (o *ShareCloudResourceWorkflowStep) GetPlacements() []Placements {
	if o == nil {
		var ret []Placements
		return ret
	}

	return o.Properties.placements
}

// GetPlacementsOk returns a tuple with the Placements field value
// and a boolean to check if the value has been set.
func (o *ShareCloudResourceWorkflowStep) GetPlacementsOk() ([]Placements, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.placements, true
}

// Placements sets field value
func (o *ShareCloudResourceWorkflowStep) Placements(v []Placements) *ShareCloudResourceWorkflowStep {
	o.Properties.placements = v
	return o
}

// GetPolicy returns the Policy field value
func (o *ShareCloudResourceWorkflowStep) GetPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *ShareCloudResourceWorkflowStep) GetPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.policy, true
}

// Policy sets field value
func (o *ShareCloudResourceWorkflowStep) Policy(v string) *ShareCloudResourceWorkflowStep {
	o.Properties.policy = v
	return o
}

func (o ShareCloudResourceSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShareCloudResourceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["env"] = o.env
	toSerialize["placements"] = o.placements
	toSerialize["policy"] = o.policy
	return toSerialize, nil
}

type NullableShareCloudResourceSpec struct {
	value *ShareCloudResourceSpec
	isSet bool
}

func (v NullableShareCloudResourceSpec) Get() *ShareCloudResourceSpec {
	return v.value
}

func (v *NullableShareCloudResourceSpec) Set(val *ShareCloudResourceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableShareCloudResourceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableShareCloudResourceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShareCloudResourceSpec(val *ShareCloudResourceSpec) *NullableShareCloudResourceSpec {
	return &NullableShareCloudResourceSpec{value: val, isSet: true}
}

func (v NullableShareCloudResourceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShareCloudResourceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ShareCloudResourceType = "share-cloud-resource"

type ShareCloudResourceWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties ShareCloudResourceSpec
}

func ShareCloudResource(name string) *ShareCloudResourceWorkflowStep {
	s := &ShareCloudResourceWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
	}}
	return s
}

func (s *ShareCloudResourceWorkflowStep) Build() v1beta1.WorkflowStep {
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
		Type:       ShareCloudResourceType,
	}
	return res
}

func (s *ShareCloudResourceWorkflowStep) If(_if string) *ShareCloudResourceWorkflowStep {
	s.Base.If = _if
	return s
}

func (s *ShareCloudResourceWorkflowStep) Alias(alias string) *ShareCloudResourceWorkflowStep {
	s.Base.Meta.Alias = alias
	return s
}

func (s *ShareCloudResourceWorkflowStep) Timeout(timeout string) *ShareCloudResourceWorkflowStep {
	s.Base.Timeout = timeout
	return s
}

func (s *ShareCloudResourceWorkflowStep) DependsOn(dependsOn []string) *ShareCloudResourceWorkflowStep {
	s.Base.DependsOn = dependsOn
	return s
}

func (s *ShareCloudResourceWorkflowStep) Inputs(input common.StepInputs) *ShareCloudResourceWorkflowStep {
	s.Base.Inputs = input
	return s
}

func (s *ShareCloudResourceWorkflowStep) Outputs(output common.StepOutputs) *ShareCloudResourceWorkflowStep {
	s.Base.Outputs = output
	return s
}

func (s *ShareCloudResourceWorkflowStep) Name() string {
	return s.Base.Name
}

func (s *ShareCloudResourceWorkflowStep) Type() string {
	return ShareCloudResourceType
}
