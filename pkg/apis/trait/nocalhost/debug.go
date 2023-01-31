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

// checks if the Debug type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Debug{}

// Debug struct for Debug
type Debug struct {
	remoteDebugPort *int32 `json:"remoteDebugPort,omitempty"`
}

// NewDebugWith instantiates a new Debug object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebugWith() *Debug {
	this := Debug{}
	return &this
}

// NewDebug instantiates a new Debug object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebug() *Debug {
	this := Debug{}
	return &this
}

// GetRemoteDebugPort returns the RemoteDebugPort field value if set, zero value otherwise.
func (o *Debug) GetRemoteDebugPort() int32 {
	if o == nil || utils.IsNil(o.remoteDebugPort) {
		var ret int32
		return ret
	}
	return *o.remoteDebugPort
}

// GetRemoteDebugPortOk returns a tuple with the RemoteDebugPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Debug) GetRemoteDebugPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.remoteDebugPort) {
		return nil, false
	}
	return o.remoteDebugPort, true
}

// HasRemoteDebugPort returns a boolean if a field has been set.
func (o *Debug) HasRemoteDebugPort() bool {
	if o != nil && !utils.IsNil(o.remoteDebugPort) {
		return true
	}

	return false
}

// RemoteDebugPort gets a reference to the given int32 and assigns it to the remoteDebugPort field.
// remoteDebugPort:
func (o *Debug) RemoteDebugPort(v int32) *Debug {
	o.remoteDebugPort = &v
	return o
}

func (o Debug) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Debug) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.remoteDebugPort) {
		toSerialize["remoteDebugPort"] = o.remoteDebugPort
	}
	return toSerialize, nil
}

type NullableDebug struct {
	value *Debug
	isSet bool
}

func (v NullableDebug) Get() *Debug {
	return v.value
}

func (v *NullableDebug) Set(val *Debug) {
	v.value = val
	v.isSet = true
}

func (v NullableDebug) IsSet() bool {
	return v.isSet
}

func (v *NullableDebug) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebug(val *Debug) *NullableDebug {
	return &NullableDebug{value: val, isSet: true}
}

func (v NullableDebug) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebug) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
