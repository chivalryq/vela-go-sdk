/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"

	"vela-go-sdk/pkg/apis/utils"
)

// checks if the MountToEnvs1 type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MountToEnvs1{}

// MountToEnvs1 struct for MountToEnvs1
type MountToEnvs1 struct {
	envName   string `json:"envName"`
	secretKey string `json:"secretKey"`
}

// NewMountToEnvs1With instantiates a new MountToEnvs1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMountToEnvs1With(envName string, secretKey string) *MountToEnvs1 {
	this := MountToEnvs1{}
	this.envName = envName
	this.secretKey = secretKey
	return &this
}

// NewMountToEnvs1 instantiates a new MountToEnvs1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMountToEnvs1() *MountToEnvs1 {
	this := MountToEnvs1{}
	return &this
}

// GetEnvName returns the EnvName field value
func (o *MountToEnvs1) GetEnvName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.envName
}

// GetEnvNameOk returns a tuple with the EnvName field value
// and a boolean to check if the value has been set.
func (o *MountToEnvs1) GetEnvNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.envName, true
}

// EnvName sets field value
func (o *MountToEnvs1) EnvName(v string) *MountToEnvs1 {
	o.envName = v
	return o
}

// GetSecretKey returns the SecretKey field value
func (o *MountToEnvs1) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.secretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *MountToEnvs1) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.secretKey, true
}

// SecretKey sets field value
func (o *MountToEnvs1) SecretKey(v string) *MountToEnvs1 {
	o.secretKey = v
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
	toSerialize["envName"] = o.envName
	toSerialize["secretKey"] = o.secretKey
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
