/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package container_image

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the ContainerImageSpecOneOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ContainerImageSpecOneOf{}

// ContainerImageSpecOneOf struct for ContainerImageSpecOneOf
type ContainerImageSpecOneOf struct {
	// Specify the container image for multiple containers
	containers []PatchParams `json:"containers"`
}

// NewContainerImageSpecOneOf instantiates a new ContainerImageSpecOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerImageSpecOneOf(containers []PatchParams) *ContainerImageSpecOneOf {
	this := ContainerImageSpecOneOf{}
	this.containers = containers
	return &this
}

// NewContainerImageSpecOneOfWithDefaults instantiates a new ContainerImageSpecOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerImageSpecOneOfWithDefaults() *ContainerImageSpecOneOf {
	this := ContainerImageSpecOneOf{}
	return &this
}

// GetContainers returns the Containers field value
func (o *ContainerImageSpecOneOf) GetContainers() []PatchParams {
	if o == nil {
		var ret []PatchParams
		return ret
	}

	return o.containers
}

// GetContainersOk returns a tuple with the Containers field value
// and a boolean to check if the value has been set.
func (o *ContainerImageSpecOneOf) GetContainersOk() ([]PatchParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.containers, true
}

// Containers sets field value
func (o *ContainerImageSpecOneOf) Containers(v []PatchParams) *ContainerImageSpecOneOf {
	o.containers = v
    return o
}

func (o ContainerImageSpecOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerImageSpecOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["containers"] = o.containers
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

 
