/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package daemon

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the EmptyDir type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EmptyDir{}

// EmptyDir struct for EmptyDir
type EmptyDir struct {
	medium string `json:"medium"`
	mountPath string `json:"mountPath"`
	name string `json:"name"`
}

// NewEmptyDir instantiates a new EmptyDir object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmptyDir(medium string, mountPath string, name string) *EmptyDir {
	this := EmptyDir{}
	this.medium = medium
	this.mountPath = mountPath
	this.name = name
	return &this
}

// NewEmptyDirWithDefaults instantiates a new EmptyDir object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmptyDirWithDefaults() *EmptyDir {
	this := EmptyDir{}
	var medium string = ""
	this.medium = medium
	return &this
}

// GetMedium returns the Medium field value
func (o *EmptyDir) GetMedium() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.medium
}

// GetMediumOk returns a tuple with the Medium field value
// and a boolean to check if the value has been set.
func (o *EmptyDir) GetMediumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.medium, true
}

// Medium sets field value
func (o *EmptyDir) Medium(v string) *EmptyDir {
	o.medium = v
    return o
}

// GetMountPath returns the MountPath field value
func (o *EmptyDir) GetMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.mountPath
}

// GetMountPathOk returns a tuple with the MountPath field value
// and a boolean to check if the value has been set.
func (o *EmptyDir) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.mountPath, true
}

// MountPath sets field value
func (o *EmptyDir) MountPath(v string) *EmptyDir {
	o.mountPath = v
    return o
}

// GetName returns the Name field value
func (o *EmptyDir) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EmptyDir) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *EmptyDir) Name(v string) *EmptyDir {
	o.name = v
    return o
}

func (o EmptyDir) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmptyDir) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["medium"] = o.medium
	toSerialize["mountPath"] = o.mountPath
	toSerialize["name"] = o.name
	return toSerialize, nil
}

type NullableEmptyDir struct {
	value *EmptyDir
	isSet bool
}

func (v NullableEmptyDir) Get() *EmptyDir {
	return v.value
}

func (v *NullableEmptyDir) Set(val *EmptyDir) {
	v.value = val
	v.isSet = true
}

func (v NullableEmptyDir) IsSet() bool {
	return v.isSet
}

func (v *NullableEmptyDir) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmptyDir(val *EmptyDir) *NullableEmptyDir {
	return &NullableEmptyDir{value: val, isSet: true}
}

func (v NullableEmptyDir) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmptyDir) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 
