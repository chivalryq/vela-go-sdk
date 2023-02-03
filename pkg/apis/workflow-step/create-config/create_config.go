/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package create_config

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the CreateConfigSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateConfigSpec{}

// CreateConfigSpec struct for CreateConfigSpec
type CreateConfigSpec struct {
	// Specify the content of the config.
	Config map[string]interface{} `json:"config,omitempty"`
	// Specify the name of the config.
	Name *string `json:"name,omitempty"`
	// Specify the namespace of the config.
	Namespace *string `json:"namespace,omitempty"`
	// Specify the template of the config.
	Template *string `json:"template,omitempty"`
}

// NewCreateConfigSpecWith instantiates a new CreateConfigSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConfigSpecWith() *CreateConfigSpec {
	this := CreateConfigSpec{}
	return &this
}

// NewCreateConfigSpec instantiates a new CreateConfigSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConfigSpec() *CreateConfigSpec {
	this := CreateConfigSpec{}
	return &this
}

// NewCreateConfigSpecs converts a list CreateConfigSpec pointers to objects.
// This is helpful when the SetCreateConfigSpec requires a list of objects
func NewCreateConfigSpecs(ps ...*CreateConfigSpec) []CreateConfigSpec {
	objs := []CreateConfigSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CreateConfigWorkflowStep) GetConfig() map[string]interface{} {
	if o == nil || utils.IsNil(o.Properties.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConfigWorkflowStep) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Properties.Config) {
		return map[string]interface{}{}, false
	}
	return o.Properties.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CreateConfigWorkflowStep) HasConfig() bool {
	if o != nil && !utils.IsNil(o.Properties.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the config field.
// Config:  Specify the content of the config.
func (o *CreateConfigWorkflowStep) SetConfig(v map[string]interface{}) *CreateConfigWorkflowStep {
	o.Properties.Config = v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateConfigWorkflowStep) GetName() string {
	if o == nil || utils.IsNil(o.Properties.Name) {
		var ret string
		return ret
	}
	return *o.Properties.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConfigWorkflowStep) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Name) {
		return nil, false
	}
	return o.Properties.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateConfigWorkflowStep) HasName() bool {
	if o != nil && !utils.IsNil(o.Properties.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:  Specify the name of the config.
func (o *CreateConfigWorkflowStep) SetName(v string) *CreateConfigWorkflowStep {
	o.Properties.Name = &v
	return o
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *CreateConfigWorkflowStep) GetNamespace() string {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		var ret string
		return ret
	}
	return *o.Properties.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConfigWorkflowStep) GetNamespaceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		return nil, false
	}
	return o.Properties.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *CreateConfigWorkflowStep) HasNamespace() bool {
	if o != nil && !utils.IsNil(o.Properties.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the namespace field.
// Namespace:  Specify the namespace of the config.
func (o *CreateConfigWorkflowStep) SetNamespace(v string) *CreateConfigWorkflowStep {
	o.Properties.Namespace = &v
	return o
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *CreateConfigWorkflowStep) GetTemplate() string {
	if o == nil || utils.IsNil(o.Properties.Template) {
		var ret string
		return ret
	}
	return *o.Properties.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConfigWorkflowStep) GetTemplateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Template) {
		return nil, false
	}
	return o.Properties.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *CreateConfigWorkflowStep) HasTemplate() bool {
	if o != nil && !utils.IsNil(o.Properties.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the template field.
// Template:  Specify the template of the config.
func (o *CreateConfigWorkflowStep) SetTemplate(v string) *CreateConfigWorkflowStep {
	o.Properties.Template = &v
	return o
}

func (o CreateConfigSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateConfigSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !utils.IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	return toSerialize, nil
}

type NullableCreateConfigSpec struct {
	value *CreateConfigSpec
	isSet bool
}

func (v NullableCreateConfigSpec) Get() *CreateConfigSpec {
	return v.value
}

func (v *NullableCreateConfigSpec) Set(val *CreateConfigSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConfigSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConfigSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConfigSpec(val *CreateConfigSpec) *NullableCreateConfigSpec {
	return &NullableCreateConfigSpec{value: val, isSet: true}
}

func (v NullableCreateConfigSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateConfigSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const CreateConfigType = "create-config"

func init() {
	sdkcommon.RegisterWorkflowStep(CreateConfigType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(CreateConfigType, FromWorkflowSubStep)
}

type CreateConfigWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties CreateConfigSpec
}

func CreateConfig(name string) *CreateConfigWorkflowStep {
	c := &CreateConfigWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: CreateConfigType,
	}}
	return c
}

func (c *CreateConfigWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range c.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  c.Base.DependsOn,
		If:         c.Base.If,
		Inputs:     c.Base.Inputs,
		Meta:       c.Base.Meta,
		Name:       c.Base.Name,
		Outputs:    c.Base.Outputs,
		Properties: util.Object2RawExtension(c.Properties),
		SubSteps:   subSteps,
		Timeout:    c.Base.Timeout,
		Type:       CreateConfigType,
	}
	return res
}

func (c *CreateConfigWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*CreateConfigWorkflowStep, error) {
	var properties CreateConfigSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := c.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	c.Base.Name = from.Name
	c.Base.DependsOn = from.DependsOn
	c.Base.Inputs = from.Inputs
	c.Base.Outputs = from.Outputs
	c.Base.If = from.If
	c.Base.Timeout = from.Timeout
	c.Base.Meta = from.Meta
	c.Base.Type = CreateConfigType
	c.Properties = properties
	c.Base.SubSteps = subSteps
	return c, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	c := &CreateConfigWorkflowStep{}
	return c.FromWorkflowStep(from)
}

func (c *CreateConfigWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*CreateConfigWorkflowStep, error) {
	var properties CreateConfigSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	c.Base.Name = from.Name
	c.Base.DependsOn = from.DependsOn
	c.Base.Inputs = from.Inputs
	c.Base.Outputs = from.Outputs
	c.Base.If = from.If
	c.Base.Timeout = from.Timeout
	c.Base.Meta = from.Meta
	c.Base.Type = CreateConfigType
	c.Properties = properties
	return c, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	c := &CreateConfigWorkflowStep{}
	return c.FromWorkflowSubStep(from)
}

func (c *CreateConfigWorkflowStep) WorkflowStepName() string {
	return c.Base.Name
}

func (c *CreateConfigWorkflowStep) DefType() string {
	return CreateConfigType
}

func (c *CreateConfigWorkflowStep) If(_if string) *CreateConfigWorkflowStep {
	c.Base.If = _if
	return c
}

func (c *CreateConfigWorkflowStep) Alias(alias string) *CreateConfigWorkflowStep {
	c.Base.Meta.Alias = alias
	return c
}

func (c *CreateConfigWorkflowStep) Timeout(timeout string) *CreateConfigWorkflowStep {
	c.Base.Timeout = timeout
	return c
}

func (c *CreateConfigWorkflowStep) DependsOn(dependsOn []string) *CreateConfigWorkflowStep {
	c.Base.DependsOn = dependsOn
	return c
}

func (c *CreateConfigWorkflowStep) Inputs(input common.StepInputs) *CreateConfigWorkflowStep {
	c.Base.Inputs = input
	return c
}

func (c *CreateConfigWorkflowStep) Outputs(output common.StepOutputs) *CreateConfigWorkflowStep {
	c.Base.Outputs = output
	return c
}
