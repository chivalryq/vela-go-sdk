/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cron_task

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the SecretKeyRef type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SecretKeyRef{}

// SecretKeyRef Selects a key of a secret in the pod's namespace
type SecretKeyRef struct {
	// The key of the secret to select from. Must be a valid secret key
	Key *string `json:"key,omitempty"`
	// The name of the secret in the pod's namespace to select from
	Name *string `json:"name,omitempty"`
}

// NewSecretKeyRefWith instantiates a new SecretKeyRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretKeyRefWith() *SecretKeyRef {
	this := SecretKeyRef{}
	return &this
}

// NewSecretKeyRef instantiates a new SecretKeyRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretKeyRef() *SecretKeyRef {
	this := SecretKeyRef{}
	return &this
}

// NewSecretKeyRefs converts a list SecretKeyRef pointers to objects.
// This is helpful when the SetSecretKeyRef requires a list of objects
func NewSecretKeyRefs(ps ...*SecretKeyRef) []SecretKeyRef {
	objs := []SecretKeyRef{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SecretKeyRef) GetKey() string {
	if o == nil || utils.IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretKeyRef) GetKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SecretKeyRef) HasKey() bool {
	if o != nil && !utils.IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the key field.
// Key:  The key of the secret to select from. Must be a valid secret key
func (o *SecretKeyRef) SetKey(v string) *SecretKeyRef {
	o.Key = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecretKeyRef) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretKeyRef) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecretKeyRef) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:  The name of the secret in the pod's namespace to select from
func (o *SecretKeyRef) SetName(v string) *SecretKeyRef {
	o.Name = &v
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
	if !utils.IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
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
