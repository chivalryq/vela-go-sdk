/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cron_task

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the HostAliases type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HostAliases{}

// HostAliases struct for HostAliases
type HostAliases struct {
	hostnames []string `json:"hostnames"`
	ip string `json:"ip"`
}

// NewHostAliases instantiates a new HostAliases object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostAliases(hostnames []string, ip string) *HostAliases {
	this := HostAliases{}
	this.hostnames = hostnames
	this.ip = ip
	return &this
}

// NewHostAliasesWithDefaults instantiates a new HostAliases object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostAliasesWithDefaults() *HostAliases {
	this := HostAliases{}
	return &this
}

// GetHostnames returns the Hostnames field value
func (o *HostAliases) GetHostnames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.hostnames
}

// GetHostnamesOk returns a tuple with the Hostnames field value
// and a boolean to check if the value has been set.
func (o *HostAliases) GetHostnamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.hostnames, true
}

// Hostnames sets field value
func (o *HostAliases) Hostnames(v []string) *HostAliases {
	o.hostnames = v
    return o
}

// GetIp returns the Ip field value
func (o *HostAliases) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *HostAliases) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ip, true
}

// Ip sets field value
func (o *HostAliases) Ip(v string) *HostAliases {
	o.ip = v
    return o
}

func (o HostAliases) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostAliases) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hostnames"] = o.hostnames
	toSerialize["ip"] = o.ip
	return toSerialize, nil
}

type NullableHostAliases struct {
	value *HostAliases
	isSet bool
}

func (v NullableHostAliases) Get() *HostAliases {
	return v.value
}

func (v *NullableHostAliases) Set(val *HostAliases) {
	v.value = val
	v.isSet = true
}

func (v NullableHostAliases) IsSet() bool {
	return v.isSet
}

func (v *NullableHostAliases) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostAliases(val *HostAliases) *NullableHostAliases {
	return &NullableHostAliases{value: val, isSet: true}
}

func (v NullableHostAliases) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostAliases) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 
