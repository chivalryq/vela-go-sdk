/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the MountToEnvs1 type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MountToEnvs1{}

// MountToEnvs1 struct for MountToEnvs1
type MountToEnvs1 struct {
	EnvName   *string `json:"envName,omitempty"`
	SecretKey *string `json:"secretKey,omitempty"`
}

// NewMountToEnvs1With instantiates a new MountToEnvs1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMountToEnvs1With() *MountToEnvs1 {
	this := MountToEnvs1{}
	return &this
}

// NewMountToEnvs1 instantiates a new MountToEnvs1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMountToEnvs1() *MountToEnvs1 {
	this := MountToEnvs1{}
	return &this
}

// NewMountToEnvs1s converts a list MountToEnvs1 pointers to objects.
// This is helpful when the SetMountToEnvs1 requires a list of objects
func NewMountToEnvs1s(ps ...*MountToEnvs1) []MountToEnvs1 {
	objs := []MountToEnvs1{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetEnvName returns the EnvName field value if set, zero value otherwise.
func (o *MountToEnvs1) GetEnvName() string {
	if o == nil || utils.IsNil(o.EnvName) {
		var ret string
		return ret
	}
	return *o.EnvName
}

// GetEnvNameOk returns a tuple with the EnvName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MountToEnvs1) GetEnvNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.EnvName) {
		return nil, false
	}
	return o.EnvName, true
}

// HasEnvName returns a boolean if a field has been set.
func (o *MountToEnvs1) HasEnvName() bool {
	if o != nil && !utils.IsNil(o.EnvName) {
		return true
	}

	return false
}

// SetEnvName gets a reference to the given string and assigns it to the envName field.
// EnvName:
func (o *MountToEnvs1) SetEnvName(v string) *MountToEnvs1 {
	o.EnvName = &v
	return o
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *MountToEnvs1) GetSecretKey() string {
	if o == nil || utils.IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MountToEnvs1) GetSecretKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *MountToEnvs1) HasSecretKey() bool {
	if o != nil && !utils.IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the secretKey field.
// SecretKey:
func (o *MountToEnvs1) SetSecretKey(v string) *MountToEnvs1 {
	o.SecretKey = &v
	return o
}

func (o MountToEnvs1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MountToEnvs1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.EnvName) {
		toSerialize["envName"] = o.EnvName
	}
	if !utils.IsNil(o.SecretKey) {
		toSerialize["secretKey"] = o.SecretKey
	}
	return toSerialize, nil
}

type NullableMountToEnvs1 struct {
	value *MountToEnvs1
	isSet bool
}

func (v NullableMountToEnvs1) Get() *MountToEnvs1 {
	return v.value
}

func (v *NullableMountToEnvs1) Set(val *MountToEnvs1) {
	v.value = val
	v.isSet = true
}

func (v NullableMountToEnvs1) IsSet() bool {
	return v.isSet
}

func (v *NullableMountToEnvs1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMountToEnvs1(val *MountToEnvs1) *NullableMountToEnvs1 {
	return &NullableMountToEnvs1{value: val, isSet: true}
}

func (v NullableMountToEnvs1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMountToEnvs1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
