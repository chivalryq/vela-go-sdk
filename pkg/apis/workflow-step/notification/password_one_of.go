/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the PasswordOneOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PasswordOneOf{}

// PasswordOneOf struct for PasswordOneOf
type PasswordOneOf struct {
	// the password content in string
	value string `json:"value"`
}

// NewPasswordOneOfWith instantiates a new PasswordOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordOneOfWith(value string) *PasswordOneOf {
	this := PasswordOneOf{}
	this.value = value
	return &this
}

// NewPasswordOneOf instantiates a new PasswordOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordOneOf() *PasswordOneOf {
	this := PasswordOneOf{}
	return &this
}

// GetValue returns the Value field value
func (o *PasswordOneOf) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PasswordOneOf) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.value, true
}

// Value sets field value
func (o *PasswordOneOf) Value(v string) *PasswordOneOf {
	o.value = v
	return o
}

func (o PasswordOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.value
	return toSerialize, nil
}

type NullablePasswordOneOf struct {
	value *PasswordOneOf
	isSet bool
}

func (v NullablePasswordOneOf) Get() *PasswordOneOf {
	return v.value
}

func (v *NullablePasswordOneOf) Set(val *PasswordOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordOneOf(val *PasswordOneOf) *NullablePasswordOneOf {
	return &NullablePasswordOneOf{value: val, isSet: true}
}

func (v NullablePasswordOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
