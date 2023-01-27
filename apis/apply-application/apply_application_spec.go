/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_application

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the ApplyApplicationSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApplyApplicationSpec{}

// ApplyApplicationSpec struct for ApplyApplicationSpec
type ApplyApplicationSpec struct {
}

// NewApplyApplicationSpec instantiates a new ApplyApplicationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyApplicationSpec() *ApplyApplicationSpec {
	this := ApplyApplicationSpec{}
	return &this
}

// NewApplyApplicationSpecWithDefaults instantiates a new ApplyApplicationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyApplicationSpecWithDefaults() *ApplyApplicationSpec {
	this := ApplyApplicationSpec{}
	return &this
}

func (o ApplyApplicationSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyApplicationSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableApplyApplicationSpec struct {
	value *ApplyApplicationSpec
	isSet bool
}

func (v NullableApplyApplicationSpec) Get() *ApplyApplicationSpec {
	return v.value
}

func (v *NullableApplyApplicationSpec) Set(val *ApplyApplicationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyApplicationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyApplicationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyApplicationSpec(val *ApplyApplicationSpec) *NullableApplyApplicationSpec {
	return &NullableApplyApplicationSpec{value: val, isSet: true}
}

func (v NullableApplyApplicationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyApplicationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 