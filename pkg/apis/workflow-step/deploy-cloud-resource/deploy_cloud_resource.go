/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deploy_cloud_resource

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"vela-go-sdk/pkg/apis"
	"vela-go-sdk/pkg/apis/utils"
)

// checks if the DeployCloudResourceSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DeployCloudResourceSpec{}

// DeployCloudResourceSpec struct for DeployCloudResourceSpec
type DeployCloudResourceSpec struct {
	// Declare the name of the env in policy
	env string `json:"env"`
	// Declare the name of the env-binding policy, if empty, the first env-binding policy will be used
	policy string `json:"policy"`
}

// NewDeployCloudResourceSpecWith instantiates a new DeployCloudResourceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployCloudResourceSpecWith(env string, policy string) *DeployCloudResourceSpec {
	this := DeployCloudResourceSpec{}
	this.env = env
	this.policy = policy
	return &this
}

// NewDeployCloudResourceSpec instantiates a new DeployCloudResourceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeployCloudResourceSpec() *DeployCloudResourceSpec {
	this := DeployCloudResourceSpec{}
	var policy string = ""
	this.policy = policy
	return &this
}

// GetEnv returns the Env field value
func (o *DeployCloudResourceWorkflowStep) GetEnv() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *DeployCloudResourceWorkflowStep) GetEnvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.env, true
}

// Env sets field value
func (o *DeployCloudResourceWorkflowStep) Env(v string) *DeployCloudResourceWorkflowStep {
	o.Properties.env = v
	return o
}

// GetPolicy returns the Policy field value
func (o *DeployCloudResourceWorkflowStep) GetPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *DeployCloudResourceWorkflowStep) GetPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.policy, true
}

// Policy sets field value
func (o *DeployCloudResourceWorkflowStep) Policy(v string) *DeployCloudResourceWorkflowStep {
	o.Properties.policy = v
	return o
}

func (o DeployCloudResourceSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeployCloudResourceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["env"] = o.env
	toSerialize["policy"] = o.policy
	return toSerialize, nil
}

type NullableDeployCloudResourceSpec struct {
	value *DeployCloudResourceSpec
	isSet bool
}

func (v NullableDeployCloudResourceSpec) Get() *DeployCloudResourceSpec {
	return v.value
}

func (v *NullableDeployCloudResourceSpec) Set(val *DeployCloudResourceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployCloudResourceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployCloudResourceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployCloudResourceSpec(val *DeployCloudResourceSpec) *NullableDeployCloudResourceSpec {
	return &NullableDeployCloudResourceSpec{value: val, isSet: true}
}

func (v NullableDeployCloudResourceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployCloudResourceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const DeployCloudResourceType = "deploy-cloud-resource"

type DeployCloudResourceWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties DeployCloudResourceSpec
}

func DeployCloudResource(name string) *DeployCloudResourceWorkflowStep {
	d := &DeployCloudResourceWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
	}}
	return d
}

func (d *DeployCloudResourceWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range d.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  d.Base.DependsOn,
		If:         d.Base.If,
		Inputs:     d.Base.Inputs,
		Meta:       d.Base.Meta,
		Name:       d.Base.Name,
		Outputs:    d.Base.Outputs,
		Properties: util.Object2RawExtension(d.Properties),
		SubSteps:   subSteps,
		Timeout:    d.Base.Timeout,
		Type:       DeployCloudResourceType,
	}
	return res
}

func (d *DeployCloudResourceWorkflowStep) If(_if string) *DeployCloudResourceWorkflowStep {
	d.Base.If = _if
	return d
}

func (d *DeployCloudResourceWorkflowStep) Alias(alias string) *DeployCloudResourceWorkflowStep {
	d.Base.Meta.Alias = alias
	return d
}

func (d *DeployCloudResourceWorkflowStep) Timeout(timeout string) *DeployCloudResourceWorkflowStep {
	d.Base.Timeout = timeout
	return d
}

func (d *DeployCloudResourceWorkflowStep) DependsOn(dependsOn []string) *DeployCloudResourceWorkflowStep {
	d.Base.DependsOn = dependsOn
	return d
}

func (d *DeployCloudResourceWorkflowStep) Inputs(input common.StepInputs) *DeployCloudResourceWorkflowStep {
	d.Base.Inputs = input
	return d
}

func (d *DeployCloudResourceWorkflowStep) Outputs(output common.StepOutputs) *DeployCloudResourceWorkflowStep {
	d.Base.Outputs = output
	return d
}

func (d *DeployCloudResourceWorkflowStep) Name() string {
	return d.Base.Name
}

func (d *DeployCloudResourceWorkflowStep) Type() string {
	return DeployCloudResourceType
}
