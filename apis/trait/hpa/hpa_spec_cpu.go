/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hpa

import (
	"encoding/json"

	"vela-go-sdk/apis/utils"
)

// checks if the HpaSpecCpu type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HpaSpecCpu{}

// HpaSpecCpu struct for HpaSpecCpu
type HpaSpecCpu struct {
	// Specify resource metrics in terms of percentage(\"Utilization\") or direct value(\"AverageValue\")
	type_ string `json:"type"`
	// Specify the value of CPU utilization or averageValue
	value int32 `json:"value"`
}

// HpaSpecCpuWith instantiates a new HpaSpecCpu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func HpaSpecCpuWith(type_ string, value int32) *HpaSpecCpu {
	this := HpaSpecCpu{}
	this.type_ = type_
	this.value = value
	return &this
}

// NewHpaSpecCpu instantiates a new HpaSpecCpu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHpaSpecCpu() *HpaSpecCpu {
	this := HpaSpecCpu{}
	var type_ string = "Utilization"
	this.type_ = type_
	var value int32 = 50
	this.value = value
	return &this
}

// GetType returns the Type field value
func (o *HpaSpecCpu) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.type_
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HpaSpecCpu) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.type_, true
}

// Type sets field value
func (o *HpaSpecCpu) Type(v string) *HpaSpecCpu {
	o.type_ = v
	return o
}

// GetValue returns the Value field value
func (o *HpaSpecCpu) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *HpaSpecCpu) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.value, true
}

// Value sets field value
func (o *HpaSpecCpu) Value(v int32) *HpaSpecCpu {
	o.value = v
	return o
}

func (o HpaSpecCpu) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HpaSpecCpu) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.type_
	toSerialize["value"] = o.value
	return toSerialize, nil
}

type NullableHpaSpecCpu struct {
	value *HpaSpecCpu
	isSet bool
}

func (v NullableHpaSpecCpu) Get() *HpaSpecCpu {
	return v.value
}

func (v *NullableHpaSpecCpu) Set(val *HpaSpecCpu) {
	v.value = val
	v.isSet = true
}

func (v NullableHpaSpecCpu) IsSet() bool {
	return v.isSet
}

func (v *NullableHpaSpecCpu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHpaSpecCpu(val *HpaSpecCpu) *NullableHpaSpecCpu {
	return &NullableHpaSpecCpu{value: val, isSet: true}
}

func (v NullableHpaSpecCpu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHpaSpecCpu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
