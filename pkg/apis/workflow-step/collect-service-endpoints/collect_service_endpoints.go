/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package collect_service_endpoints

import (
	"encoding/json"
	"vela-go-sdk/pkg/apis"
	"vela-go-sdk/pkg/apis/utils"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)

// checks if the CollectServiceEndpointsSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CollectServiceEndpointsSpec{}

// CollectServiceEndpointsSpec struct for CollectServiceEndpointsSpec
type CollectServiceEndpointsSpec struct {
	// Filter the component of the endpoints
	components []string `json:"components,omitempty"`
	// Specify the name of the application
	name *string `json:"name,omitempty"`
	// Specify the namespace of the application
	namespace *string `json:"namespace,omitempty"`
	// Filter the endpoint that are only outer
	outer *bool `json:"outer,omitempty"`
	// Filter the port of the endpoints
	port *int32 `json:"port,omitempty"`
	// Filter the port name of the endpoints
	portName *string `json:"portName,omitempty"`
	// The protocal of endpoint url
	protocal string `json:"protocal"`
}

// NewCollectServiceEndpointsSpecWith instantiates a new CollectServiceEndpointsSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectServiceEndpointsSpecWith(protocal string) *CollectServiceEndpointsSpec {
	this := CollectServiceEndpointsSpec{}
	this.protocal = protocal
	return &this
}

// NewCollectServiceEndpointsSpec instantiates a new CollectServiceEndpointsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectServiceEndpointsSpec() *CollectServiceEndpointsSpec {
	this := CollectServiceEndpointsSpec{}
	var protocal string = "http"
	this.protocal = protocal
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *CollectServiceEndpointsWorkflowStep) GetComponents() []string {
	if o == nil || utils.IsNil(o.Properties.components) {
		var ret []string
		return ret
	}
	return o.Properties.components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectServiceEndpointsWorkflowStep) GetComponentsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.components) {
		return nil, false
	}
	return o.Properties.components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *CollectServiceEndpointsWorkflowStep) HasComponents() bool {
	if o != nil && !utils.IsNil(o.Properties.components) {
		return true
	}

	return false
}

// Components gets a reference to the given []string and assigns it to the components field.
// components:  Filter the component of the endpoints
func (o *CollectServiceEndpointsWorkflowStep) Components(v []string) *CollectServiceEndpointsWorkflowStep {
	o.Properties.components = v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CollectServiceEndpointsWorkflowStep) GetName() string {
	if o == nil || utils.IsNil(o.Properties.name) {
		var ret string
		return ret
	}
	return *o.Properties.name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectServiceEndpointsWorkflowStep) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.name) {
		return nil, false
	}
	return o.Properties.name, true
}

// HasName returns a boolean if a field has been set.
func (o *CollectServiceEndpointsWorkflowStep) HasName() bool {
	if o != nil && !utils.IsNil(o.Properties.name) {
		return true
	}

	return false
}

// Name gets a reference to the given string and assigns it to the name field.
// name:  Specify the name of the application
func (o *CollectServiceEndpointsWorkflowStep) Name(v string) *CollectServiceEndpointsWorkflowStep {
	o.Properties.name = &v
	return o
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *CollectServiceEndpointsWorkflowStep) GetNamespace() string {
	if o == nil || utils.IsNil(o.Properties.namespace) {
		var ret string
		return ret
	}
	return *o.Properties.namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectServiceEndpointsWorkflowStep) GetNamespaceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.namespace) {
		return nil, false
	}
	return o.Properties.namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *CollectServiceEndpointsWorkflowStep) HasNamespace() bool {
	if o != nil && !utils.IsNil(o.Properties.namespace) {
		return true
	}

	return false
}

// Namespace gets a reference to the given string and assigns it to the namespace field.
// namespace:  Specify the namespace of the application
func (o *CollectServiceEndpointsWorkflowStep) Namespace(v string) *CollectServiceEndpointsWorkflowStep {
	o.Properties.namespace = &v
	return o
}

// GetOuter returns the Outer field value if set, zero value otherwise.
func (o *CollectServiceEndpointsWorkflowStep) GetOuter() bool {
	if o == nil || utils.IsNil(o.Properties.outer) {
		var ret bool
		return ret
	}
	return *o.Properties.outer
}

// GetOuterOk returns a tuple with the Outer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectServiceEndpointsWorkflowStep) GetOuterOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Properties.outer) {
		return nil, false
	}
	return o.Properties.outer, true
}

// HasOuter returns a boolean if a field has been set.
func (o *CollectServiceEndpointsWorkflowStep) HasOuter() bool {
	if o != nil && !utils.IsNil(o.Properties.outer) {
		return true
	}

	return false
}

// Outer gets a reference to the given bool and assigns it to the outer field.
// outer:  Filter the endpoint that are only outer
func (o *CollectServiceEndpointsWorkflowStep) Outer(v bool) *CollectServiceEndpointsWorkflowStep {
	o.Properties.outer = &v
	return o
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *CollectServiceEndpointsWorkflowStep) GetPort() int32 {
	if o == nil || utils.IsNil(o.Properties.port) {
		var ret int32
		return ret
	}
	return *o.Properties.port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectServiceEndpointsWorkflowStep) GetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.port) {
		return nil, false
	}
	return o.Properties.port, true
}

// HasPort returns a boolean if a field has been set.
func (o *CollectServiceEndpointsWorkflowStep) HasPort() bool {
	if o != nil && !utils.IsNil(o.Properties.port) {
		return true
	}

	return false
}

// Port gets a reference to the given int32 and assigns it to the port field.
// port:  Filter the port of the endpoints
func (o *CollectServiceEndpointsWorkflowStep) Port(v int32) *CollectServiceEndpointsWorkflowStep {
	o.Properties.port = &v
	return o
}

// GetPortName returns the PortName field value if set, zero value otherwise.
func (o *CollectServiceEndpointsWorkflowStep) GetPortName() string {
	if o == nil || utils.IsNil(o.Properties.portName) {
		var ret string
		return ret
	}
	return *o.Properties.portName
}

// GetPortNameOk returns a tuple with the PortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectServiceEndpointsWorkflowStep) GetPortNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.portName) {
		return nil, false
	}
	return o.Properties.portName, true
}

// HasPortName returns a boolean if a field has been set.
func (o *CollectServiceEndpointsWorkflowStep) HasPortName() bool {
	if o != nil && !utils.IsNil(o.Properties.portName) {
		return true
	}

	return false
}

// PortName gets a reference to the given string and assigns it to the portName field.
// portName:  Filter the port name of the endpoints
func (o *CollectServiceEndpointsWorkflowStep) PortName(v string) *CollectServiceEndpointsWorkflowStep {
	o.Properties.portName = &v
	return o
}

// GetProtocal returns the Protocal field value
func (o *CollectServiceEndpointsWorkflowStep) GetProtocal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.protocal
}

// GetProtocalOk returns a tuple with the Protocal field value
// and a boolean to check if the value has been set.
func (o *CollectServiceEndpointsWorkflowStep) GetProtocalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.protocal, true
}

// Protocal sets field value
func (o *CollectServiceEndpointsWorkflowStep) Protocal(v string) *CollectServiceEndpointsWorkflowStep {
	o.Properties.protocal = v
	return o
}

func (o CollectServiceEndpointsSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectServiceEndpointsSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.components) {
		toSerialize["components"] = o.components
	}
	if !utils.IsNil(o.name) {
		toSerialize["name"] = o.name
	}
	if !utils.IsNil(o.namespace) {
		toSerialize["namespace"] = o.namespace
	}
	if !utils.IsNil(o.outer) {
		toSerialize["outer"] = o.outer
	}
	if !utils.IsNil(o.port) {
		toSerialize["port"] = o.port
	}
	if !utils.IsNil(o.portName) {
		toSerialize["portName"] = o.portName
	}
	toSerialize["protocal"] = o.protocal
	return toSerialize, nil
}

type NullableCollectServiceEndpointsSpec struct {
	value *CollectServiceEndpointsSpec
	isSet bool
}

func (v NullableCollectServiceEndpointsSpec) Get() *CollectServiceEndpointsSpec {
	return v.value
}

func (v *NullableCollectServiceEndpointsSpec) Set(val *CollectServiceEndpointsSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectServiceEndpointsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectServiceEndpointsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectServiceEndpointsSpec(val *CollectServiceEndpointsSpec) *NullableCollectServiceEndpointsSpec {
	return &NullableCollectServiceEndpointsSpec{value: val, isSet: true}
}

func (v NullableCollectServiceEndpointsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectServiceEndpointsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const CollectServiceEndpointsType = "collect-service-endpoints"

type CollectServiceEndpointsWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties CollectServiceEndpointsSpec
}

func CollectServiceEndpoints(name string) *CollectServiceEndpointsWorkflowStep {
	c := &CollectServiceEndpointsWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
	}}
	return c
}

func (c *CollectServiceEndpointsWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range c.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties})
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
		Type:       CollectServiceEndpointsType,
	}
	return res
}

func (c *CollectServiceEndpointsWorkflowStep) If(_if string) *CollectServiceEndpointsWorkflowStep {
	c.Base.If = _if
	return c
}

func (c *CollectServiceEndpointsWorkflowStep) Alias(alias string) *CollectServiceEndpointsWorkflowStep {
	c.Base.Meta.Alias = alias
	return c
}

func (c *CollectServiceEndpointsWorkflowStep) Timeout(timeout string) *CollectServiceEndpointsWorkflowStep {
	c.Base.Timeout = timeout
	return c
}

func (c *CollectServiceEndpointsWorkflowStep) DependsOn(dependsOn []string) *CollectServiceEndpointsWorkflowStep {
	c.Base.DependsOn = dependsOn
	return c
}

func (c *CollectServiceEndpointsWorkflowStep) Inputs(input common.StepInputs) *CollectServiceEndpointsWorkflowStep {
	c.Base.Inputs = input
	return c
}

func (c *CollectServiceEndpointsWorkflowStep) Outputs(output common.StepOutputs) *CollectServiceEndpointsWorkflowStep {
	c.Base.Outputs = output
	return c
}