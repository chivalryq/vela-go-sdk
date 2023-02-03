/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package container_image

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the ContainerImageSpecOneOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ContainerImageSpecOneOf{}

// ContainerImageSpecOneOf struct for ContainerImageSpecOneOf
type ContainerImageSpecOneOf struct {
	// Specify the container image for multiple containers
	Containers []PatchParams `json:"containers"`
}

// NewContainerImageSpecOneOfWith instantiates a new ContainerImageSpecOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerImageSpecOneOfWith(containers []PatchParams) *ContainerImageSpecOneOf {
	this := ContainerImageSpecOneOf{}
	this.Containers = containers
	return &this
}

// NewContainerImageSpecOneOf instantiates a new ContainerImageSpecOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerImageSpecOneOf() *ContainerImageSpecOneOf {
	this := ContainerImageSpecOneOf{}
	return &this
}

// NewContainerImageSpecOneOfs converts a list ContainerImageSpecOneOf pointers to objects.
// This is helpful when the SetContainerImageSpecOneOf requires a list of objects
func NewContainerImageSpecOneOfs(ps ...*ContainerImageSpecOneOf) []ContainerImageSpecOneOf {
	objs := []ContainerImageSpecOneOf{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetContainers returns the Containers field value
func (o *ContainerImageSpecOneOf) GetContainers() []PatchParams {
	if o == nil {
		var ret []PatchParams
		return ret
	}

	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value
// and a boolean to check if the value has been set.
func (o *ContainerImageSpecOneOf) GetContainersOk() ([]PatchParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.Containers, true
}

// SetContainers sets field value
func (o *ContainerImageSpecOneOf) SetContainers(v []PatchParams) *ContainerImageSpecOneOf {
	o.Containers = v
	return o
}

func (o ContainerImageSpecOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerImageSpecOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["containers"] = o.Containers
	return toSerialize, nil
}

type NullableContainerImageSpecOneOf struct {
	value *ContainerImageSpecOneOf
	isSet bool
}

func (v NullableContainerImageSpecOneOf) Get() *ContainerImageSpecOneOf {
	return v.value
}

func (v *NullableContainerImageSpecOneOf) Set(val *ContainerImageSpecOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerImageSpecOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerImageSpecOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerImageSpecOneOf(val *ContainerImageSpecOneOf) *NullableContainerImageSpecOneOf {
	return &NullableContainerImageSpecOneOf{value: val, isSet: true}
}

func (v NullableContainerImageSpecOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerImageSpecOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
