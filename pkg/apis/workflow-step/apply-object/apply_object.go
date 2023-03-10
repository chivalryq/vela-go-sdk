/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_object

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the ApplyObjectSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApplyObjectSpec{}

// ApplyObjectSpec struct for ApplyObjectSpec
type ApplyObjectSpec struct {
	// The cluster you want to apply the resource to, default is the current control plane cluster
	Cluster *string `json:"cluster,omitempty"`
	// Specify Kubernetes native resource object to be applied
	Value map[string]interface{} `json:"value,omitempty"`
}

// NewApplyObjectSpecWith instantiates a new ApplyObjectSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyObjectSpecWith() *ApplyObjectSpec {
	this := ApplyObjectSpec{}
	var cluster string = ""
	this.Cluster = &cluster
	return &this
}

// NewApplyObjectSpec instantiates a new ApplyObjectSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyObjectSpec() *ApplyObjectSpec {
	this := ApplyObjectSpec{}
	var cluster string = ""
	this.Cluster = &cluster
	return &this
}

// NewApplyObjectSpecs converts a list ApplyObjectSpec pointers to objects.
// This is helpful when the SetApplyObjectSpec requires a list of objects
func NewApplyObjectSpecs(ps ...*ApplyObjectSpec) []ApplyObjectSpec {
	objs := []ApplyObjectSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *ApplyObjectWorkflowStep) GetCluster() string {
	if o == nil || utils.IsNil(o.Properties.Cluster) {
		var ret string
		return ret
	}
	return *o.Properties.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyObjectWorkflowStep) GetClusterOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Cluster) {
		return nil, false
	}
	return o.Properties.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ApplyObjectWorkflowStep) HasCluster() bool {
	if o != nil && !utils.IsNil(o.Properties.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the cluster field.
// Cluster:  The cluster you want to apply the resource to, default is the current control plane cluster
func (o *ApplyObjectWorkflowStep) SetCluster(v string) *ApplyObjectWorkflowStep {
	o.Properties.Cluster = &v
	return o
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApplyObjectWorkflowStep) GetValue() map[string]interface{} {
	if o == nil || utils.IsNil(o.Properties.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyObjectWorkflowStep) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Properties.Value) {
		return map[string]interface{}{}, false
	}
	return o.Properties.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApplyObjectWorkflowStep) HasValue() bool {
	if o != nil && !utils.IsNil(o.Properties.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the value field.
// Value:  Specify Kubernetes native resource object to be applied
func (o *ApplyObjectWorkflowStep) SetValue(v map[string]interface{}) *ApplyObjectWorkflowStep {
	o.Properties.Value = v
	return o
}

func (o ApplyObjectSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyObjectSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !utils.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableApplyObjectSpec struct {
	value *ApplyObjectSpec
	isSet bool
}

func (v NullableApplyObjectSpec) Get() *ApplyObjectSpec {
	return v.value
}

func (v *NullableApplyObjectSpec) Set(val *ApplyObjectSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyObjectSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyObjectSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyObjectSpec(val *ApplyObjectSpec) *NullableApplyObjectSpec {
	return &NullableApplyObjectSpec{value: val, isSet: true}
}

func (v NullableApplyObjectSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyObjectSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ApplyObjectType = "apply-object"

func init() {
	sdkcommon.RegisterWorkflowStep(ApplyObjectType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(ApplyObjectType, FromWorkflowSubStep)
}

type ApplyObjectWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties ApplyObjectSpec
}

func ApplyObject(name string) *ApplyObjectWorkflowStep {
	a := &ApplyObjectWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: ApplyObjectType,
	}}
	return a
}

func (a *ApplyObjectWorkflowStep) Build() v1beta1.WorkflowStep {
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
		Type:       ApplyObjectType,
	}
	return res
}

func (a *ApplyObjectWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*ApplyObjectWorkflowStep, error) {
	var properties ApplyObjectSpec
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
	a.Base.Type = ApplyObjectType
	a.Properties = properties
	a.Base.SubSteps = subSteps
	return a, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	a := &ApplyObjectWorkflowStep{}
	return a.FromWorkflowStep(from)
}

func (a *ApplyObjectWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*ApplyObjectWorkflowStep, error) {
	var properties ApplyObjectSpec
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
	a.Base.Type = ApplyObjectType
	a.Properties = properties
	return a, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	a := &ApplyObjectWorkflowStep{}
	return a.FromWorkflowSubStep(from)
}

func (a *ApplyObjectWorkflowStep) WorkflowStepName() string {
	return a.Base.Name
}

func (a *ApplyObjectWorkflowStep) DefType() string {
	return ApplyObjectType
}

func (a *ApplyObjectWorkflowStep) If(_if string) *ApplyObjectWorkflowStep {
	a.Base.If = _if
	return a
}

func (a *ApplyObjectWorkflowStep) Alias(alias string) *ApplyObjectWorkflowStep {
	a.Base.Meta.Alias = alias
	return a
}

func (a *ApplyObjectWorkflowStep) Timeout(timeout string) *ApplyObjectWorkflowStep {
	a.Base.Timeout = timeout
	return a
}

func (a *ApplyObjectWorkflowStep) DependsOn(dependsOn []string) *ApplyObjectWorkflowStep {
	a.Base.DependsOn = dependsOn
	return a
}

func (a *ApplyObjectWorkflowStep) Inputs(input common.StepInputs) *ApplyObjectWorkflowStep {
	a.Base.Inputs = input
	return a
}

func (a *ApplyObjectWorkflowStep) Outputs(output common.StepOutputs) *ApplyObjectWorkflowStep {
	a.Base.Outputs = output
	return a
}
