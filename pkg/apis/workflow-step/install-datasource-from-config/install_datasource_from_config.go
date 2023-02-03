/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package install_datasource_from_config

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the InstallDatasourceFromConfigSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InstallDatasourceFromConfigSpec{}

// InstallDatasourceFromConfigSpec struct for InstallDatasourceFromConfigSpec
type InstallDatasourceFromConfigSpec struct {
	Grafana *string    `json:"grafana,omitempty"`
	Type    *ModelType `json:"type,omitempty"`
}

// NewInstallDatasourceFromConfigSpecWith instantiates a new InstallDatasourceFromConfigSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallDatasourceFromConfigSpecWith() *InstallDatasourceFromConfigSpec {
	this := InstallDatasourceFromConfigSpec{}
	var grafana string = "default"
	this.Grafana = &grafana
	var type_ ModelType = "prometheus-server"
	this.Type = &type_
	return &this
}

// NewInstallDatasourceFromConfigSpec instantiates a new InstallDatasourceFromConfigSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallDatasourceFromConfigSpec() *InstallDatasourceFromConfigSpec {
	this := InstallDatasourceFromConfigSpec{}
	var grafana string = "default"
	this.Grafana = &grafana
	var type_ ModelType = "prometheus-server"
	this.Type = &type_
	return &this
}

// NewInstallDatasourceFromConfigSpecs converts a list InstallDatasourceFromConfigSpec pointers to objects.
// This is helpful when the SetInstallDatasourceFromConfigSpec requires a list of objects
func NewInstallDatasourceFromConfigSpecs(ps ...*InstallDatasourceFromConfigSpec) []InstallDatasourceFromConfigSpec {
	objs := []InstallDatasourceFromConfigSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetGrafana returns the Grafana field value if set, zero value otherwise.
func (o *InstallDatasourceFromConfigWorkflowStep) GetGrafana() string {
	if o == nil || utils.IsNil(o.Properties.Grafana) {
		var ret string
		return ret
	}
	return *o.Properties.Grafana
}

// GetGrafanaOk returns a tuple with the Grafana field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallDatasourceFromConfigWorkflowStep) GetGrafanaOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Grafana) {
		return nil, false
	}
	return o.Properties.Grafana, true
}

// HasGrafana returns a boolean if a field has been set.
func (o *InstallDatasourceFromConfigWorkflowStep) HasGrafana() bool {
	if o != nil && !utils.IsNil(o.Properties.Grafana) {
		return true
	}

	return false
}

// SetGrafana gets a reference to the given string and assigns it to the grafana field.
// Grafana:
func (o *InstallDatasourceFromConfigWorkflowStep) SetGrafana(v string) *InstallDatasourceFromConfigWorkflowStep {
	o.Properties.Grafana = &v
	return o
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InstallDatasourceFromConfigWorkflowStep) GetType() ModelType {
	if o == nil || utils.IsNil(o.Properties.Type) {
		var ret ModelType
		return ret
	}
	return *o.Properties.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallDatasourceFromConfigWorkflowStep) GetTypeOk() (*ModelType, bool) {
	if o == nil || utils.IsNil(o.Properties.Type) {
		return nil, false
	}
	return o.Properties.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InstallDatasourceFromConfigWorkflowStep) HasType() bool {
	if o != nil && !utils.IsNil(o.Properties.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ModelType and assigns it to the type_ field.
// Type:
func (o *InstallDatasourceFromConfigWorkflowStep) SetType(v ModelType) *InstallDatasourceFromConfigWorkflowStep {
	o.Properties.Type = &v
	return o
}

func (o InstallDatasourceFromConfigSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstallDatasourceFromConfigSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Grafana) {
		toSerialize["grafana"] = o.Grafana
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableInstallDatasourceFromConfigSpec struct {
	value *InstallDatasourceFromConfigSpec
	isSet bool
}

func (v NullableInstallDatasourceFromConfigSpec) Get() *InstallDatasourceFromConfigSpec {
	return v.value
}

func (v *NullableInstallDatasourceFromConfigSpec) Set(val *InstallDatasourceFromConfigSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallDatasourceFromConfigSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallDatasourceFromConfigSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallDatasourceFromConfigSpec(val *InstallDatasourceFromConfigSpec) *NullableInstallDatasourceFromConfigSpec {
	return &NullableInstallDatasourceFromConfigSpec{value: val, isSet: true}
}

func (v NullableInstallDatasourceFromConfigSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallDatasourceFromConfigSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const InstallDatasourceFromConfigType = "install-datasource-from-config"

func init() {
	sdkcommon.RegisterWorkflowStep(InstallDatasourceFromConfigType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(InstallDatasourceFromConfigType, FromWorkflowSubStep)
}

type InstallDatasourceFromConfigWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties InstallDatasourceFromConfigSpec
}

func InstallDatasourceFromConfig(name string) *InstallDatasourceFromConfigWorkflowStep {
	i := &InstallDatasourceFromConfigWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: InstallDatasourceFromConfigType,
	}}
	return i
}

func (i *InstallDatasourceFromConfigWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range i.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  i.Base.DependsOn,
		If:         i.Base.If,
		Inputs:     i.Base.Inputs,
		Meta:       i.Base.Meta,
		Name:       i.Base.Name,
		Outputs:    i.Base.Outputs,
		Properties: util.Object2RawExtension(i.Properties),
		SubSteps:   subSteps,
		Timeout:    i.Base.Timeout,
		Type:       InstallDatasourceFromConfigType,
	}
	return res
}

func (i *InstallDatasourceFromConfigWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*InstallDatasourceFromConfigWorkflowStep, error) {
	var properties InstallDatasourceFromConfigSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := i.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	i.Base.Name = from.Name
	i.Base.DependsOn = from.DependsOn
	i.Base.Inputs = from.Inputs
	i.Base.Outputs = from.Outputs
	i.Base.If = from.If
	i.Base.Timeout = from.Timeout
	i.Base.Meta = from.Meta
	i.Base.Type = InstallDatasourceFromConfigType
	i.Properties = properties
	i.Base.SubSteps = subSteps
	return i, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	i := &InstallDatasourceFromConfigWorkflowStep{}
	return i.FromWorkflowStep(from)
}

func (i *InstallDatasourceFromConfigWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*InstallDatasourceFromConfigWorkflowStep, error) {
	var properties InstallDatasourceFromConfigSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	i.Base.Name = from.Name
	i.Base.DependsOn = from.DependsOn
	i.Base.Inputs = from.Inputs
	i.Base.Outputs = from.Outputs
	i.Base.If = from.If
	i.Base.Timeout = from.Timeout
	i.Base.Meta = from.Meta
	i.Base.Type = InstallDatasourceFromConfigType
	i.Properties = properties
	return i, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	i := &InstallDatasourceFromConfigWorkflowStep{}
	return i.FromWorkflowSubStep(from)
}

func (i *InstallDatasourceFromConfigWorkflowStep) WorkflowStepName() string {
	return i.Base.Name
}

func (i *InstallDatasourceFromConfigWorkflowStep) DefType() string {
	return InstallDatasourceFromConfigType
}

func (i *InstallDatasourceFromConfigWorkflowStep) If(_if string) *InstallDatasourceFromConfigWorkflowStep {
	i.Base.If = _if
	return i
}

func (i *InstallDatasourceFromConfigWorkflowStep) Alias(alias string) *InstallDatasourceFromConfigWorkflowStep {
	i.Base.Meta.Alias = alias
	return i
}

func (i *InstallDatasourceFromConfigWorkflowStep) Timeout(timeout string) *InstallDatasourceFromConfigWorkflowStep {
	i.Base.Timeout = timeout
	return i
}

func (i *InstallDatasourceFromConfigWorkflowStep) DependsOn(dependsOn []string) *InstallDatasourceFromConfigWorkflowStep {
	i.Base.DependsOn = dependsOn
	return i
}

func (i *InstallDatasourceFromConfigWorkflowStep) Inputs(input common.StepInputs) *InstallDatasourceFromConfigWorkflowStep {
	i.Base.Inputs = input
	return i
}

func (i *InstallDatasourceFromConfigWorkflowStep) Outputs(output common.StepOutputs) *InstallDatasourceFromConfigWorkflowStep {
	i.Base.Outputs = output
	return i
}
