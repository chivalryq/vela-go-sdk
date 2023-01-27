/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package service_binding

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the ServiceBindingSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ServiceBindingSpec{}

// ServiceBindingSpec struct for ServiceBindingSpec
type ServiceBindingSpec struct {
	// The mapping of environment variables to secret
	envMappings map[string]KeySecret `json:"envMappings"`
}

// NewServiceBindingSpec instantiates a new ServiceBindingSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceBindingSpec(envMappings map[string]KeySecret) *ServiceBindingSpec {
	this := ServiceBindingSpec{}
	this.envMappings = envMappings
	return &this
}

// NewServiceBindingSpecWithDefaults instantiates a new ServiceBindingSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceBindingSpecWithDefaults() *ServiceBindingSpec {
	this := ServiceBindingSpec{}
	return &this
}

// GetEnvMappings returns the EnvMappings field value
func (o *ServiceBindingSpec) GetEnvMappings() map[string]KeySecret {
	if o == nil {
		var ret map[string]KeySecret
		return ret
	}

	return o.envMappings
}

// GetEnvMappingsOk returns a tuple with the EnvMappings field value
// and a boolean to check if the value has been set.
func (o *ServiceBindingSpec) GetEnvMappingsOk() (*map[string]KeySecret, bool) {
	if o == nil {
		return nil, false
	}
	return &o.envMappings, true
}

// EnvMappings sets field value
func (o *ServiceBindingSpec) EnvMappings(v map[string]KeySecret) *ServiceBindingSpec {
	o.envMappings = v
    return o
}

func (o ServiceBindingSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceBindingSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["envMappings"] = o.envMappings
	return toSerialize, nil
}

type NullableServiceBindingSpec struct {
	value *ServiceBindingSpec
	isSet bool
}

func (v NullableServiceBindingSpec) Get() *ServiceBindingSpec {
	return v.value
}

func (v *NullableServiceBindingSpec) Set(val *ServiceBindingSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceBindingSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceBindingSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceBindingSpec(val *ServiceBindingSpec) *NullableServiceBindingSpec {
	return &NullableServiceBindingSpec{value: val, isSet: true}
}

func (v NullableServiceBindingSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceBindingSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 