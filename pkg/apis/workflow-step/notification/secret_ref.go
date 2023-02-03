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

// checks if the SecretRef type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SecretRef{}

// SecretRef struct for SecretRef
type SecretRef struct {
	// key is the key in the secret
	Key *string `json:"key,omitempty"`
	// name is the name of the secret
	Name *string `json:"name,omitempty"`
}

// NewSecretRefWith instantiates a new SecretRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretRefWith() *SecretRef {
	this := SecretRef{}
	return &this
}

// NewSecretRef instantiates a new SecretRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretRef() *SecretRef {
	this := SecretRef{}
	return &this
}

// NewSecretRefs converts a list SecretRef pointers to objects.
// This is helpful when the SetSecretRef requires a list of objects
func NewSecretRefs(ps ...*SecretRef) []SecretRef {
	objs := []SecretRef{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SecretRef) GetKey() string {
	if o == nil || utils.IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRef) GetKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SecretRef) HasKey() bool {
	if o != nil && !utils.IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the key field.
// Key:  key is the key in the secret
func (o *SecretRef) SetKey(v string) *SecretRef {
	o.Key = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecretRef) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRef) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecretRef) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:  name is the name of the secret
func (o *SecretRef) SetName(v string) *SecretRef {
	o.Name = &v
	return o
}

func (o SecretRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableSecretRef struct {
	value *SecretRef
	isSet bool
}

func (v NullableSecretRef) Get() *SecretRef {
	return v.value
}

func (v *NullableSecretRef) Set(val *SecretRef) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretRef) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretRef(val *SecretRef) *NullableSecretRef {
	return &NullableSecretRef{value: val, isSet: true}
}

func (v NullableSecretRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
