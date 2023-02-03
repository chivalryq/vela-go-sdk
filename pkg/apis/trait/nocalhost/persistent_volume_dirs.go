/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nocalhost

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the PersistentVolumeDirs type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PersistentVolumeDirs{}

// PersistentVolumeDirs struct for PersistentVolumeDirs
type PersistentVolumeDirs struct {
	Capacity *string `json:"capacity,omitempty"`
	Path     *string `json:"path,omitempty"`
}

// NewPersistentVolumeDirsWith instantiates a new PersistentVolumeDirs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersistentVolumeDirsWith() *PersistentVolumeDirs {
	this := PersistentVolumeDirs{}
	return &this
}

// NewPersistentVolumeDirs instantiates a new PersistentVolumeDirs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersistentVolumeDirs() *PersistentVolumeDirs {
	this := PersistentVolumeDirs{}
	return &this
}

// NewPersistentVolumeDirss converts a list PersistentVolumeDirs pointers to objects.
// This is helpful when the SetPersistentVolumeDirs requires a list of objects
func NewPersistentVolumeDirss(ps ...*PersistentVolumeDirs) []PersistentVolumeDirs {
	objs := []PersistentVolumeDirs{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *PersistentVolumeDirs) GetCapacity() string {
	if o == nil || utils.IsNil(o.Capacity) {
		var ret string
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistentVolumeDirs) GetCapacityOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Capacity) {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *PersistentVolumeDirs) HasCapacity() bool {
	if o != nil && !utils.IsNil(o.Capacity) {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given string and assigns it to the capacity field.
// Capacity:
func (o *PersistentVolumeDirs) SetCapacity(v string) *PersistentVolumeDirs {
	o.Capacity = &v
	return o
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PersistentVolumeDirs) GetPath() string {
	if o == nil || utils.IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistentVolumeDirs) GetPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PersistentVolumeDirs) HasPath() bool {
	if o != nil && !utils.IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the path field.
// Path:
func (o *PersistentVolumeDirs) SetPath(v string) *PersistentVolumeDirs {
	o.Path = &v
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
	if !utils.IsNil(o.Capacity) {
		toSerialize["capacity"] = o.Capacity
	}
	if !utils.IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
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
