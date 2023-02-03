/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hpa

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the Cpu type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Cpu{}

// Cpu struct for Cpu
type Cpu struct {
	// Specify resource metrics in terms of percentage(\"Utilization\") or direct value(\"AverageValue\")
	Type *string `json:"type,omitempty"`
	// Specify the value of CPU utilization or averageValue
	Value *int32 `json:"value,omitempty"`
}

// NewCpuWith instantiates a new Cpu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpuWith() *Cpu {
	this := Cpu{}
	var type_ string = "Utilization"
	this.Type = &type_
	var value int32 = 50
	this.Value = &value
	return &this
}

// NewCpu instantiates a new Cpu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpu() *Cpu {
	this := Cpu{}
	var type_ string = "Utilization"
	this.Type = &type_
	var value int32 = 50
	this.Value = &value
	return &this
}

// NewCpus converts a list Cpu pointers to objects.
// This is helpful when the SetCpu requires a list of objects
func NewCpus(ps ...*Cpu) []Cpu {
	objs := []Cpu{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Cpu) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cpu) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Cpu) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the type_ field.
// Type:  Specify resource metrics in terms of percentage(\"Utilization\") or direct value(\"AverageValue\")
func (o *Cpu) SetType(v string) *Cpu {
	o.Type = &v
	return o
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Cpu) GetValue() int32 {
	if o == nil || utils.IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cpu) GetValueOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Cpu) HasValue() bool {
	if o != nil && !utils.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the value field.
// Value:  Specify the value of CPU utilization or averageValue
func (o *Cpu) SetValue(v int32) *Cpu {
	o.Value = &v
	return o
}

func (o Cpu) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cpu) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCpu struct {
	value *Cpu
	isSet bool
}

func (v NullableCpu) Get() *Cpu {
	return v.value
}

func (v *NullableCpu) Set(val *Cpu) {
	v.value = val
	v.isSet = true
}

func (v NullableCpu) IsSet() bool {
	return v.isSet
}

func (v *NullableCpu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpu(val *Cpu) *NullableCpu {
	return &NullableCpu{value: val, isSet: true}
}

func (v NullableCpu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
