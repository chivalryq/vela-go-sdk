/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package read_config

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the ReadConfigSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReadConfigSpec{}

// ReadConfigSpec struct for ReadConfigSpec
type ReadConfigSpec struct {
	// Specify the name of the config.
	Name *string `json:"name,omitempty"`
	// Specify the namespace of the config.
	Namespace *string `json:"namespace,omitempty"`
}

// NewReadConfigSpecWith instantiates a new ReadConfigSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadConfigSpecWith() *ReadConfigSpec {
	this := ReadConfigSpec{}
	return &this
}

// NewReadConfigSpec instantiates a new ReadConfigSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadConfigSpec() *ReadConfigSpec {
	this := ReadConfigSpec{}
	return &this
}

// NewReadConfigSpecs converts a list ReadConfigSpec pointers to objects.
// This is helpful when the SetReadConfigSpec requires a list of objects
func NewReadConfigSpecs(ps ...*ReadConfigSpec) []ReadConfigSpec {
	objs := []ReadConfigSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReadConfigWorkflowStep) GetName() string {
	if o == nil || utils.IsNil(o.Properties.Name) {
		var ret string
		return ret
	}
	return *o.Properties.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadConfigWorkflowStep) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Name) {
		return nil, false
	}
	return o.Properties.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReadConfigWorkflowStep) HasName() bool {
	if o != nil && !utils.IsNil(o.Properties.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:  Specify the name of the config.
func (o *ReadConfigWorkflowStep) SetName(v string) *ReadConfigWorkflowStep {
	o.Properties.Name = &v
	return o
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ReadConfigWorkflowStep) GetNamespace() string {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		var ret string
		return ret
	}
	return *o.Properties.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadConfigWorkflowStep) GetNamespaceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		return nil, false
	}
	return o.Properties.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ReadConfigWorkflowStep) HasNamespace() bool {
	if o != nil && !utils.IsNil(o.Properties.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the namespace field.
// Namespace:  Specify the namespace of the config.
func (o *ReadConfigWorkflowStep) SetNamespace(v string) *ReadConfigWorkflowStep {
	o.Properties.Namespace = &v
	return o
}

func (o ReadConfigSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadConfigSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	return toSerialize, nil
}

type NullableReadConfigSpec struct {
	value *ReadConfigSpec
	isSet bool
}

func (v NullableReadConfigSpec) Get() *ReadConfigSpec {
	return v.value
}

func (v *NullableReadConfigSpec) Set(val *ReadConfigSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableReadConfigSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableReadConfigSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadConfigSpec(val *ReadConfigSpec) *NullableReadConfigSpec {
	return &NullableReadConfigSpec{value: val, isSet: true}
}

func (v NullableReadConfigSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadConfigSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ReadConfigType = "read-config"

func init() {
	sdkcommon.RegisterWorkflowStep(ReadConfigType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(ReadConfigType, FromWorkflowSubStep)
}

type ReadConfigWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties ReadConfigSpec
}

func ReadConfig(name string) *ReadConfigWorkflowStep {
	r := &ReadConfigWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: ReadConfigType,
	}}
	return r
}

func (r *ReadConfigWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range r.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  r.Base.DependsOn,
		If:         r.Base.If,
		Inputs:     r.Base.Inputs,
		Meta:       r.Base.Meta,
		Name:       r.Base.Name,
		Outputs:    r.Base.Outputs,
		Properties: util.Object2RawExtension(r.Properties),
		SubSteps:   subSteps,
		Timeout:    r.Base.Timeout,
		Type:       ReadConfigType,
	}
	return res
}

func (r *ReadConfigWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*ReadConfigWorkflowStep, error) {
	var properties ReadConfigSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := r.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	r.Base.Name = from.Name
	r.Base.DependsOn = from.DependsOn
	r.Base.Inputs = from.Inputs
	r.Base.Outputs = from.Outputs
	r.Base.If = from.If
	r.Base.Timeout = from.Timeout
	r.Base.Meta = from.Meta
	r.Base.Type = ReadConfigType
	r.Properties = properties
	r.Base.SubSteps = subSteps
	return r, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	r := &ReadConfigWorkflowStep{}
	return r.FromWorkflowStep(from)
}

func (r *ReadConfigWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*ReadConfigWorkflowStep, error) {
	var properties ReadConfigSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	r.Base.Name = from.Name
	r.Base.DependsOn = from.DependsOn
	r.Base.Inputs = from.Inputs
	r.Base.Outputs = from.Outputs
	r.Base.If = from.If
	r.Base.Timeout = from.Timeout
	r.Base.Meta = from.Meta
	r.Base.Type = ReadConfigType
	r.Properties = properties
	return r, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	r := &ReadConfigWorkflowStep{}
	return r.FromWorkflowSubStep(from)
}

func (r *ReadConfigWorkflowStep) WorkflowStepName() string {
	return r.Base.Name
}

func (r *ReadConfigWorkflowStep) DefType() string {
	return ReadConfigType
}

func (r *ReadConfigWorkflowStep) If(_if string) *ReadConfigWorkflowStep {
	r.Base.If = _if
	return r
}

func (r *ReadConfigWorkflowStep) Alias(alias string) *ReadConfigWorkflowStep {
	r.Base.Meta.Alias = alias
	return r
}

func (r *ReadConfigWorkflowStep) Timeout(timeout string) *ReadConfigWorkflowStep {
	r.Base.Timeout = timeout
	return r
}

func (r *ReadConfigWorkflowStep) DependsOn(dependsOn []string) *ReadConfigWorkflowStep {
	r.Base.DependsOn = dependsOn
	return r
}

func (r *ReadConfigWorkflowStep) Inputs(input common.StepInputs) *ReadConfigWorkflowStep {
	r.Base.Inputs = input
	return r
}

func (r *ReadConfigWorkflowStep) Outputs(output common.StepOutputs) *ReadConfigWorkflowStep {
	r.Base.Outputs = output
	return r
}
