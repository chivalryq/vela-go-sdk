/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_terraform_config

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the ProviderRef type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ProviderRef{}

// ProviderRef providerRef specifies the reference to Provider
type ProviderRef struct {
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// NewProviderRefWith instantiates a new ProviderRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderRefWith() *ProviderRef {
	this := ProviderRef{}
	return &this
}

// NewProviderRef instantiates a new ProviderRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderRef() *ProviderRef {
	this := ProviderRef{}
	return &this
}

// NewProviderRefs converts a list ProviderRef pointers to objects.
// This is helpful when the SetProviderRef requires a list of objects
func NewProviderRefs(ps ...*ProviderRef) []ProviderRef {
	objs := []ProviderRef{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProviderRef) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderRef) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProviderRef) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:
func (o *ProviderRef) SetName(v string) *ProviderRef {
	o.Name = &v
	return o
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ProviderRef) GetNamespace() string {
	if o == nil || utils.IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderRef) GetNamespaceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ProviderRef) HasNamespace() bool {
	if o != nil && !utils.IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the namespace field.
// Namespace:
func (o *ProviderRef) SetNamespace(v string) *ProviderRef {
	o.Namespace = &v
	return o
}

func (o ProviderRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	return toSerialize, nil
}

type NullableProviderRef struct {
	value *ProviderRef
	isSet bool
}

func (v NullableProviderRef) Get() *ProviderRef {
	return v.value
}

func (v *NullableProviderRef) Set(val *ProviderRef) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderRef) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderRef(val *ProviderRef) *NullableProviderRef {
	return &NullableProviderRef{value: val, isSet: true}
}

func (v NullableProviderRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
