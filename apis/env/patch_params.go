/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package env

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the PatchParams type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PatchParams{}

// PatchParams struct for PatchParams
type PatchParams struct {
	// Specify the name of the target container, if not set, use the component name
	containerName string `json:"containerName"`
	// Specify the  environment variables to merge, if key already existing, override its value
	env map[string]string `json:"env"`
	// Specify if replacing the whole environment settings for the container
	replace bool `json:"replace"`
	// Specify which existing environment variables to unset
	unset []string `json:"unset"`
}

// NewPatchParams instantiates a new PatchParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchParams(containerName string, env map[string]string, replace bool, unset []string) *PatchParams {
	this := PatchParams{}
	this.containerName = containerName
	this.env = env
	this.replace = replace
	this.unset = unset
	return &this
}

// NewPatchParamsWithDefaults instantiates a new PatchParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchParamsWithDefaults() *PatchParams {
	this := PatchParams{}
	var containerName string = ""
	this.containerName = containerName
	var replace bool = false
	this.replace = replace
	return &this
}

// GetContainerName returns the ContainerName field value
func (o *PatchParams) GetContainerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.containerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value
// and a boolean to check if the value has been set.
func (o *PatchParams) GetContainerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.containerName, true
}

// ContainerName sets field value
func (o *PatchParams) ContainerName(v string) *PatchParams {
	o.containerName = v
    return o
}

// GetEnv returns the Env field value
func (o *PatchParams) GetEnv() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *PatchParams) GetEnvOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.env, true
}

// Env sets field value
func (o *PatchParams) Env(v map[string]string) *PatchParams {
	o.env = v
    return o
}

// GetReplace returns the Replace field value
func (o *PatchParams) GetReplace() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.replace
}

// GetReplaceOk returns a tuple with the Replace field value
// and a boolean to check if the value has been set.
func (o *PatchParams) GetReplaceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.replace, true
}

// Replace sets field value
func (o *PatchParams) Replace(v bool) *PatchParams {
	o.replace = v
    return o
}

// GetUnset returns the Unset field value
func (o *PatchParams) GetUnset() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.unset
}

// GetUnsetOk returns a tuple with the Unset field value
// and a boolean to check if the value has been set.
func (o *PatchParams) GetUnsetOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.unset, true
}

// Unset sets field value
func (o *PatchParams) Unset(v []string) *PatchParams {
	o.unset = v
    return o
}

func (o PatchParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["containerName"] = o.containerName
	toSerialize["env"] = o.env
	toSerialize["replace"] = o.replace
	toSerialize["unset"] = o.unset
	return toSerialize, nil
}

type NullablePatchParams struct {
	value *PatchParams
	isSet bool
}

func (v NullablePatchParams) Get() *PatchParams {
	return v.value
}

func (v *NullablePatchParams) Set(val *PatchParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchParams(val *PatchParams) *NullablePatchParams {
	return &NullablePatchParams{value: val, isSet: true}
}

func (v NullablePatchParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 