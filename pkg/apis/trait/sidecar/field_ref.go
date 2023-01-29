/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sidecar

import (
	"encoding/json"
	"vela-go-sdk/pkg/apis/utils"
)

// checks if the FieldRef type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FieldRef{}

// FieldRef Specify the field reference for env
type FieldRef struct {
	// Specify the field path for env
	fieldPath string `json:"fieldPath"`
}

// NewFieldRefWith instantiates a new FieldRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldRefWith(fieldPath string) *FieldRef {
	this := FieldRef{}
	this.fieldPath = fieldPath
	return &this
}

// NewFieldRef instantiates a new FieldRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldRef() *FieldRef {
	this := FieldRef{}
	return &this
}

// GetFieldPath returns the FieldPath field value
func (o *FieldRef) GetFieldPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.fieldPath
}

// GetFieldPathOk returns a tuple with the FieldPath field value
// and a boolean to check if the value has been set.
func (o *FieldRef) GetFieldPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.fieldPath, true
}

// FieldPath sets field value
func (o *FieldRef) FieldPath(v string) *FieldRef {
	o.fieldPath = v
	return o
}

func (o FieldRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fieldPath"] = o.fieldPath
	return toSerialize, nil
}

type NullableFieldRef struct {
	value *FieldRef
	isSet bool
}

func (v NullableFieldRef) Get() *FieldRef {
	return v.value
}

func (v *NullableFieldRef) Set(val *FieldRef) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldRef) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldRef(val *FieldRef) *NullableFieldRef {
	return &NullableFieldRef{value: val, isSet: true}
}

func (v NullableFieldRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}