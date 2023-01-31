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

// checks if the Volumes type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Volumes{}

// Volumes struct for Volumes
type Volumes struct {
	medium    string `json:"medium"`
	mountPath string `json:"mountPath"`
	name      string `json:"name"`
	// Specify volume type, options: \"pvc\",\"configMap\",\"secret\",\"emptyDir\", default to emptyDir
	type_ string `json:"type"`
}

// NewVolumesWith instantiates a new Volumes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumesWith(medium string, mountPath string, name string, type_ string) *Volumes {
	this := Volumes{}
	this.medium = medium
	this.mountPath = mountPath
	this.name = name
	this.type_ = type_
	return &this
}

// NewVolumes instantiates a new Volumes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumes() *Volumes {
	this := Volumes{}
	var medium string = ""
	this.medium = medium
	var type_ string = "emptyDir"
	this.type_ = type_
	return &this
}

// GetMedium returns the Medium field value
func (o *Volumes) GetMedium() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.medium
}

// GetMediumOk returns a tuple with the Medium field value
// and a boolean to check if the value has been set.
func (o *Volumes) GetMediumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.medium, true
}

// Medium sets field value
func (o *Volumes) Medium(v string) *Volumes {
	o.medium = v
	return o
}

// GetMountPath returns the MountPath field value
func (o *Volumes) GetMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.mountPath
}

// GetMountPathOk returns a tuple with the MountPath field value
// and a boolean to check if the value has been set.
func (o *Volumes) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.mountPath, true
}

// MountPath sets field value
func (o *Volumes) MountPath(v string) *Volumes {
	o.mountPath = v
	return o
}

// GetName returns the Name field value
func (o *Volumes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Volumes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *Volumes) Name(v string) *Volumes {
	o.name = v
	return o
}

// GetType returns the Type field value
func (o *Volumes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.type_
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Volumes) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.type_, true
}

// Type sets field value
func (o *Volumes) Type(v string) *Volumes {
	o.type_ = v
	return o
}

func (o Volumes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Volumes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["medium"] = o.medium
	toSerialize["mountPath"] = o.mountPath
	toSerialize["name"] = o.name
	toSerialize["type"] = o.type_
	return toSerialize, nil
}

type NullableVolumes struct {
	value *Volumes
	isSet bool
}

func (v NullableVolumes) Get() *Volumes {
	return v.value
}

func (v *NullableVolumes) Set(val *Volumes) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumes) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumes(val *Volumes) *NullableVolumes {
	return &NullableVolumes{value: val, isSet: true}
}

func (v NullableVolumes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
