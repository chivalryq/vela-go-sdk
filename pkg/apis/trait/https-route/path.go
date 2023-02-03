/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package https_route

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the Path type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Path{}

// Path An HTTP request path matcher. If this field is not specified, a default prefix match on the \"/\" path is provided.
type Path struct {
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewPathWith instantiates a new Path object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathWith() *Path {
	this := Path{}
	var type_ string = "PathPrefix"
	this.Type = &type_
	var value string = "/"
	this.Value = &value
	return &this
}

// NewPath instantiates a new Path object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPath() *Path {
	this := Path{}
	var type_ string = "PathPrefix"
	this.Type = &type_
	var value string = "/"
	this.Value = &value
	return &this
}

// NewPaths converts a list Path pointers to objects.
// This is helpful when the SetPath requires a list of objects
func NewPaths(ps ...*Path) []Path {
	objs := []Path{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Path) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Path) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Path) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the type_ field.
// Type:
func (o *Path) SetType(v string) *Path {
	o.Type = &v
	return o
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Path) GetValue() string {
	if o == nil || utils.IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Path) GetValueOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Path) HasValue() bool {
	if o != nil && !utils.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the value field.
// Value:
func (o *Path) SetValue(v string) *Path {
	o.Value = &v
	return o
}

func (o Path) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Path) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullablePath struct {
	value *Path
	isSet bool
}

func (v NullablePath) Get() *Path {
	return v.value
}

func (v *NullablePath) Set(val *Path) {
	v.value = val
	v.isSet = true
}

func (v NullablePath) IsSet() bool {
	return v.isSet
}

func (v *NullablePath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePath(val *Path) *NullablePath {
	return &NullablePath{value: val, isSet: true}
}

func (v NullablePath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
