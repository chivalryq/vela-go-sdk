/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package worker

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the HostPath type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HostPath{}

// HostPath struct for HostPath
type HostPath struct {
	MountPath *string `json:"mountPath,omitempty"`
	Name      *string `json:"name,omitempty"`
	Path      *string `json:"path,omitempty"`
}

// NewHostPathWith instantiates a new HostPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostPathWith() *HostPath {
	this := HostPath{}
	return &this
}

// NewHostPath instantiates a new HostPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostPath() *HostPath {
	this := HostPath{}
	return &this
}

// NewHostPaths converts a list HostPath pointers to objects.
// This is helpful when the SetHostPath requires a list of objects
func NewHostPaths(ps ...*HostPath) []HostPath {
	objs := []HostPath{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetMountPath returns the MountPath field value if set, zero value otherwise.
func (o *HostPath) GetMountPath() string {
	if o == nil || utils.IsNil(o.MountPath) {
		var ret string
		return ret
	}
	return *o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostPath) GetMountPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MountPath) {
		return nil, false
	}
	return o.MountPath, true
}

// HasMountPath returns a boolean if a field has been set.
func (o *HostPath) HasMountPath() bool {
	if o != nil && !utils.IsNil(o.MountPath) {
		return true
	}

	return false
}

// SetMountPath gets a reference to the given string and assigns it to the mountPath field.
// MountPath:
func (o *HostPath) SetMountPath(v string) *HostPath {
	o.MountPath = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HostPath) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostPath) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HostPath) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:
func (o *HostPath) SetName(v string) *HostPath {
	o.Name = &v
	return o
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HostPath) GetPath() string {
	if o == nil || utils.IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostPath) GetPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *HostPath) HasPath() bool {
	if o != nil && !utils.IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the path field.
// Path:
func (o *HostPath) SetPath(v string) *HostPath {
	o.Path = &v
	return o
}

func (o HostPath) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostPath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.MountPath) {
		toSerialize["mountPath"] = o.MountPath
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableHostPath struct {
	value *HostPath
	isSet bool
}

func (v NullableHostPath) Get() *HostPath {
	return v.value
}

func (v *NullableHostPath) Set(val *HostPath) {
	v.value = val
	v.isSet = true
}

func (v NullableHostPath) IsSet() bool {
	return v.isSet
}

func (v *NullableHostPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostPath(val *HostPath) *NullableHostPath {
	return &NullableHostPath{value: val, isSet: true}
}

func (v NullableHostPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
