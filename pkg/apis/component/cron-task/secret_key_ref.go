/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cron_task

import (
	"encoding/json"

	"vela-go-sdk/pkg/apis/utils"
)

// checks if the SecretKeyRef type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SecretKeyRef{}

// SecretKeyRef Selects a key of a secret in the pod's namespace
type SecretKeyRef struct {
	// The key of the secret to select from. Must be a valid secret key
	key string `json:"key"`
	// The name of the secret in the pod's namespace to select from
	name string `json:"name"`
}

// NewSecretKeyRefWith instantiates a new SecretKeyRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretKeyRefWith(key string, name string) *SecretKeyRef {
	this := SecretKeyRef{}
	this.key = key
	this.name = name
	return &this
}

// NewSecretKeyRef instantiates a new SecretKeyRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretKeyRef() *SecretKeyRef {
	this := SecretKeyRef{}
	return &this
}

// GetKey returns the Key field value
func (o *SecretKeyRef) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SecretKeyRef) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.key, true
}

// Key sets field value
func (o *SecretKeyRef) Key(v string) *SecretKeyRef {
	o.key = v
	return o
}

// GetName returns the Name field value
func (o *SecretKeyRef) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecretKeyRef) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *SecretKeyRef) Name(v string) *SecretKeyRef {
	o.name = v
	return o
}

func (o SecretKeyRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretKeyRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.key
	toSerialize["name"] = o.name
	return toSerialize, nil
}

type NullableSecretKeyRef struct {
	value *SecretKeyRef
	isSet bool
}

func (v NullableSecretKeyRef) Get() *SecretKeyRef {
	return v.value
}

func (v *NullableSecretKeyRef) Set(val *SecretKeyRef) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretKeyRef) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretKeyRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretKeyRef(val *SecretKeyRef) *NullableSecretKeyRef {
	return &NullableSecretKeyRef{value: val, isSet: true}
}

func (v NullableSecretKeyRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretKeyRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
