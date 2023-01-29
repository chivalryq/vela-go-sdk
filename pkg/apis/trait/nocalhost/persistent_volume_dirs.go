/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nocalhost

import (
	"encoding/json"

	"vela-go-sdk/pkg/apis/utils"
)

// checks if the PersistentVolumeDirs type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PersistentVolumeDirs{}

// PersistentVolumeDirs struct for PersistentVolumeDirs
type PersistentVolumeDirs struct {
	capacity string `json:"capacity"`
	path     string `json:"path"`
}

// NewPersistentVolumeDirsWith instantiates a new PersistentVolumeDirs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersistentVolumeDirsWith(capacity string, path string) *PersistentVolumeDirs {
	this := PersistentVolumeDirs{}
	this.capacity = capacity
	this.path = path
	return &this
}

// NewPersistentVolumeDirs instantiates a new PersistentVolumeDirs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersistentVolumeDirs() *PersistentVolumeDirs {
	this := PersistentVolumeDirs{}
	return &this
}

// GetCapacity returns the Capacity field value
func (o *PersistentVolumeDirs) GetCapacity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.capacity
}

// GetCapacityOk returns a tuple with the Capacity field value
// and a boolean to check if the value has been set.
func (o *PersistentVolumeDirs) GetCapacityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.capacity, true
}

// Capacity sets field value
func (o *PersistentVolumeDirs) Capacity(v string) *PersistentVolumeDirs {
	o.capacity = v
	return o
}

// GetPath returns the Path field value
func (o *PersistentVolumeDirs) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *PersistentVolumeDirs) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.path, true
}

// Path sets field value
func (o *PersistentVolumeDirs) Path(v string) *PersistentVolumeDirs {
	o.path = v
	return o
}

func (o PersistentVolumeDirs) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersistentVolumeDirs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["capacity"] = o.capacity
	toSerialize["path"] = o.path
	return toSerialize, nil
}

type NullablePersistentVolumeDirs struct {
	value *PersistentVolumeDirs
	isSet bool
}

func (v NullablePersistentVolumeDirs) Get() *PersistentVolumeDirs {
	return v.value
}

func (v *NullablePersistentVolumeDirs) Set(val *PersistentVolumeDirs) {
	v.value = val
	v.isSet = true
}

func (v NullablePersistentVolumeDirs) IsSet() bool {
	return v.isSet
}

func (v *NullablePersistentVolumeDirs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersistentVolumeDirs(val *PersistentVolumeDirs) *NullablePersistentVolumeDirs {
	return &NullablePersistentVolumeDirs{value: val, isSet: true}
}

func (v NullablePersistentVolumeDirs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersistentVolumeDirs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
