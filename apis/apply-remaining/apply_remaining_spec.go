/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_remaining

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the ApplyRemainingSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApplyRemainingSpec{}

// ApplyRemainingSpec struct for ApplyRemainingSpec
type ApplyRemainingSpec struct {
	// Declare the name of the component
	exceptions []string `json:"exceptions,omitempty"`
}

// NewApplyRemainingSpec instantiates a new ApplyRemainingSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyRemainingSpec() *ApplyRemainingSpec {
	this := ApplyRemainingSpec{}
	return &this
}

// NewApplyRemainingSpecWithDefaults instantiates a new ApplyRemainingSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyRemainingSpecWithDefaults() *ApplyRemainingSpec {
	this := ApplyRemainingSpec{}
	return &this
}

// GetExceptions returns the Exceptions field value if set, zero value otherwise.
func (o *ApplyRemainingSpec) GetExceptions() []string {
	if o == nil || utils.IsNil(o.exceptions) {
		var ret []string
		return ret
	}
	return o.exceptions
}

// GetExceptionsOk returns a tuple with the Exceptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyRemainingSpec) GetExceptionsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.exceptions) {
		return nil, false
	}
	return o.exceptions, true
}

// HasExceptions returns a boolean if a field has been set.
func (o *ApplyRemainingSpec) HasExceptions() bool {
	if o != nil && !utils.IsNil(o.exceptions) {
		return true
	}

	return false
}

// SetExceptions gets a reference to the given []string and assigns it to the exceptions field.
// exceptions:  Declare the name of the component 

func (o *ApplyRemainingSpec) Exceptions(v []string) (*ApplyRemainingSpec){
	o.exceptions = v
return o
}

func (o ApplyRemainingSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyRemainingSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.exceptions) {
		toSerialize["exceptions"] = o.exceptions
	}
	return toSerialize, nil
}

type NullableApplyRemainingSpec struct {
	value *ApplyRemainingSpec
	isSet bool
}

func (v NullableApplyRemainingSpec) Get() *ApplyRemainingSpec {
	return v.value
}

func (v *NullableApplyRemainingSpec) Set(val *ApplyRemainingSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyRemainingSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyRemainingSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyRemainingSpec(val *ApplyRemainingSpec) *NullableApplyRemainingSpec {
	return &NullableApplyRemainingSpec{value: val, isSet: true}
}

func (v NullableApplyRemainingSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyRemainingSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 
