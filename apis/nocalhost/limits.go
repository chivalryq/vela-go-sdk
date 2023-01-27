/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nocalhost

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the Limits type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Limits{}

// Limits struct for Limits
type Limits struct {
	cpu string `json:"cpu"`
	memory string `json:"memory"`
}

// NewLimits instantiates a new Limits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLimits(cpu string, memory string) *Limits {
	this := Limits{}
	this.cpu = cpu
	this.memory = memory
	return &this
}

// NewLimitsWithDefaults instantiates a new Limits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimitsWithDefaults() *Limits {
	this := Limits{}
	var cpu string = "2"
	this.cpu = cpu
	var memory string = "2Gi"
	this.memory = memory
	return &this
}

// GetCpu returns the Cpu field value
func (o *Limits) GetCpu() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *Limits) GetCpuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.cpu, true
}

// Cpu sets field value
func (o *Limits) Cpu(v string) *Limits {
	o.cpu = v
    return o
}

// GetMemory returns the Memory field value
func (o *Limits) GetMemory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *Limits) GetMemoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.memory, true
}

// Memory sets field value
func (o *Limits) Memory(v string) *Limits {
	o.memory = v
    return o
}

func (o Limits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Limits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cpu"] = o.cpu
	toSerialize["memory"] = o.memory
	return toSerialize, nil
}

type NullableLimits struct {
	value *Limits
	isSet bool
}

func (v NullableLimits) Get() *Limits {
	return v.value
}

func (v *NullableLimits) Set(val *Limits) {
	v.value = val
	v.isSet = true
}

func (v NullableLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLimits(val *Limits) *NullableLimits {
	return &NullableLimits{value: val, isSet: true}
}

func (v NullableLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 
