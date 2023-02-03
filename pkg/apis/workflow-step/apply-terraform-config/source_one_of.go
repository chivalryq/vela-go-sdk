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

// checks if the SourceOneOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SourceOneOf{}

// SourceOneOf struct for SourceOneOf
type SourceOneOf struct {
	// directly specify the hcl of the terraform configuration
	Hcl string `json:"hcl"`
}

// NewSourceOneOfWith instantiates a new SourceOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceOneOfWith(hcl string) *SourceOneOf {
	this := SourceOneOf{}
	this.Hcl = hcl
	return &this
}

// NewSourceOneOf instantiates a new SourceOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceOneOf() *SourceOneOf {
	this := SourceOneOf{}
	return &this
}

// NewSourceOneOfs converts a list SourceOneOf pointers to objects.
// This is helpful when the SetSourceOneOf requires a list of objects
func NewSourceOneOfs(ps ...*SourceOneOf) []SourceOneOf {
	objs := []SourceOneOf{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetHcl returns the Hcl field value
func (o *SourceOneOf) GetHcl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hcl
}

// GetHclOk returns a tuple with the Hcl field value
// and a boolean to check if the value has been set.
func (o *SourceOneOf) GetHclOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hcl, true
}

// SetHcl sets field value
func (o *SourceOneOf) SetHcl(v string) *SourceOneOf {
	o.Hcl = v
	return o
}

func (o SourceOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hcl"] = o.Hcl
	return toSerialize, nil
}

type NullableSourceOneOf struct {
	value *SourceOneOf
	isSet bool
}

func (v NullableSourceOneOf) Get() *SourceOneOf {
	return v.value
}

func (v *NullableSourceOneOf) Set(val *SourceOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceOneOf(val *SourceOneOf) *NullableSourceOneOf {
	return &NullableSourceOneOf{value: val, isSet: true}
}

func (v NullableSourceOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
