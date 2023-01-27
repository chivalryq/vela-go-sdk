/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package create_config

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the CreateConfigSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateConfigSpec{}

// CreateConfigSpec struct for CreateConfigSpec
type CreateConfigSpec struct {
	// Specify the content of the config.
	config map[string]interface{} `json:"config"`
	// Specify the name of the config.
	name string `json:"name"`
	// Specify the namespace of the config.
	namespace *string `json:"namespace,omitempty"`
	// Specify the template of the config.
	template *string `json:"template,omitempty"`
}

// NewCreateConfigSpec instantiates a new CreateConfigSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConfigSpec(config map[string]interface{}, name string) *CreateConfigSpec {
	this := CreateConfigSpec{}
	this.config = config
	this.name = name
	return &this
}

// NewCreateConfigSpecWithDefaults instantiates a new CreateConfigSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConfigSpecWithDefaults() *CreateConfigSpec {
	this := CreateConfigSpec{}
	return &this
}

// GetConfig returns the Config field value
func (o *CreateConfigSpec) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *CreateConfigSpec) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.config, true
}

// Config sets field value
func (o *CreateConfigSpec) Config(v map[string]interface{}) *CreateConfigSpec {
	o.config = v
    return o
}

// GetName returns the Name field value
func (o *CreateConfigSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateConfigSpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *CreateConfigSpec) Name(v string) *CreateConfigSpec {
	o.name = v
    return o
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *CreateConfigSpec) GetNamespace() string {
	if o == nil || utils.IsNil(o.namespace) {
		var ret string
		return ret
	}
	return *o.namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConfigSpec) GetNamespaceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.namespace) {
		return nil, false
	}
	return o.namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *CreateConfigSpec) HasNamespace() bool {
	if o != nil && !utils.IsNil(o.namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the namespace field.
// namespace:  Specify the namespace of the config. 

func (o *CreateConfigSpec) Namespace(v string) (*CreateConfigSpec){
	o.namespace = &v
return o
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *CreateConfigSpec) GetTemplate() string {
	if o == nil || utils.IsNil(o.template) {
		var ret string
		return ret
	}
	return *o.template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConfigSpec) GetTemplateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.template) {
		return nil, false
	}
	return o.template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *CreateConfigSpec) HasTemplate() bool {
	if o != nil && !utils.IsNil(o.template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the template field.
// template:  Specify the template of the config. 

func (o *CreateConfigSpec) Template(v string) (*CreateConfigSpec){
	o.template = &v
return o
}

func (o CreateConfigSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateConfigSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.config
	toSerialize["name"] = o.name
	if !utils.IsNil(o.namespace) {
		toSerialize["namespace"] = o.namespace
	}
	if !utils.IsNil(o.template) {
		toSerialize["template"] = o.template
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

 
