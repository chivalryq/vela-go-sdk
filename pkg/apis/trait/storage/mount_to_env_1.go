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

// checks if the MountToEnv1 type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MountToEnv1{}

// MountToEnv1 struct for MountToEnv1
type MountToEnv1 struct {
	EnvName   *string `json:"envName,omitempty"`
	SecretKey *string `json:"secretKey,omitempty"`
}

// NewMountToEnv1With instantiates a new MountToEnv1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMountToEnv1With() *MountToEnv1 {
	this := MountToEnv1{}
	return &this
}

// NewMountToEnv1 instantiates a new MountToEnv1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMountToEnv1() *MountToEnv1 {
	this := MountToEnv1{}
	return &this
}

// NewMountToEnv1s converts a list MountToEnv1 pointers to objects.
// This is helpful when the SetMountToEnv1 requires a list of objects
func NewMountToEnv1s(ps ...*MountToEnv1) []MountToEnv1 {
	objs := []MountToEnv1{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetEnvName returns the EnvName field value if set, zero value otherwise.
func (o *MountToEnv1) GetEnvName() string {
	if o == nil || utils.IsNil(o.EnvName) {
		var ret string
		return ret
	}
	return *o.EnvName
}

// GetEnvNameOk returns a tuple with the EnvName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MountToEnv1) GetEnvNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.EnvName) {
		return nil, false
	}
	return o.EnvName, true
}

// HasEnvName returns a boolean if a field has been set.
func (o *MountToEnv1) HasEnvName() bool {
	if o != nil && !utils.IsNil(o.EnvName) {
		return true
	}

	return false
}

// SetEnvName gets a reference to the given string and assigns it to the envName field.
// EnvName:
func (o *MountToEnv1) SetEnvName(v string) *MountToEnv1 {
	o.EnvName = &v
	return o
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *MountToEnv1) GetSecretKey() string {
	if o == nil || utils.IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MountToEnv1) GetSecretKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *MountToEnv1) HasSecretKey() bool {
	if o != nil && !utils.IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the secretKey field.
// SecretKey:
func (o *MountToEnv1) SetSecretKey(v string) *MountToEnv1 {
	o.SecretKey = &v
	return o
}

func (o MountToEnv1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MountToEnv1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.EnvName) {
		toSerialize["envName"] = o.EnvName
	}
	if !utils.IsNil(o.SecretKey) {
		toSerialize["secretKey"] = o.SecretKey
	}
	return toSerialize, nil
}

type NullableMountToEnv1 struct {
	value *MountToEnv1
	isSet bool
}

func (v NullableMountToEnv1) Get() *MountToEnv1 {
	return v.value
}

func (v *NullableMountToEnv1) Set(val *MountToEnv1) {
	v.value = val
	v.isSet = true
}

func (v NullableMountToEnv1) IsSet() bool {
	return v.isSet
}

func (v *NullableMountToEnv1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMountToEnv1(val *MountToEnv1) *NullableMountToEnv1 {
	return &NullableMountToEnv1{value: val, isSet: true}
}

func (v NullableMountToEnv1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMountToEnv1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
