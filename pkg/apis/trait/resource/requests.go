/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resource

import (
	"encoding/json"
	"vela-go-sdk/pkg/apis/utils"
)

// checks if the Requests type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Requests{}

// Requests Specify the resources in requests
type Requests struct {
	// Specify the amount of cpu for requests
	cpu float32 `json:"cpu"`
	// Specify the amount of memory for requests
	memory string `json:"memory"`
}

// NewRequestsWith instantiates a new Requests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestsWith(cpu float32, memory string) *Requests {
	this := Requests{}
	this.cpu = cpu
	this.memory = memory
	return &this
}

// NewRequests instantiates a new Requests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequests() *Requests {
	this := Requests{}
	var cpu float32 = 1
	this.cpu = cpu
	var memory string = "2048Mi"
	this.memory = memory
	return &this
}

// GetCpu returns the Cpu field value
func (o *Requests) GetCpu() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *Requests) GetCpuOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.cpu, true
}

// Cpu sets field value
func (o *Requests) Cpu(v float32) *Requests {
	o.cpu = v
	return o
}

// GetMemory returns the Memory field value
func (o *Requests) GetMemory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *Requests) GetMemoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.memory, true
}

// Memory sets field value
func (o *Requests) Memory(v string) *Requests {
	o.memory = v
	return o
}

func (o Requests) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Requests) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cpu"] = o.cpu
	toSerialize["memory"] = o.memory
	return toSerialize, nil
}

type NullableRequests struct {
	value *Requests
	isSet bool
}

func (v NullableRequests) Get() *Requests {
	return v.value
}

func (v *NullableRequests) Set(val *Requests) {
	v.value = val
	v.isSet = true
}

func (v NullableRequests) IsSet() bool {
	return v.isSet
}

func (v *NullableRequests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequests(val *Requests) *NullableRequests {
	return &NullableRequests{value: val, isSet: true}
}

func (v NullableRequests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}