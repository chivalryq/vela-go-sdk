/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s_update_strategy

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the RollingStrategy type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RollingStrategy{}

// RollingStrategy Specify the parameters of rollong update strategy
type RollingStrategy struct {
	maxSurge string `json:"maxSurge"`
	maxUnavailable string `json:"maxUnavailable"`
	partition int32 `json:"partition"`
}

// NewRollingStrategy instantiates a new RollingStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRollingStrategy(maxSurge string, maxUnavailable string, partition int32) *RollingStrategy {
	this := RollingStrategy{}
	this.maxSurge = maxSurge
	this.maxUnavailable = maxUnavailable
	this.partition = partition
	return &this
}

// NewRollingStrategyWithDefaults instantiates a new RollingStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollingStrategyWithDefaults() *RollingStrategy {
	this := RollingStrategy{}
	var maxSurge string = "25%"
	this.maxSurge = maxSurge
	var maxUnavailable string = "25%"
	this.maxUnavailable = maxUnavailable
	var partition int32 = 0
	this.partition = partition
	return &this
}

// GetMaxSurge returns the MaxSurge field value
func (o *RollingStrategy) GetMaxSurge() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.maxSurge
}

// GetMaxSurgeOk returns a tuple with the MaxSurge field value
// and a boolean to check if the value has been set.
func (o *RollingStrategy) GetMaxSurgeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.maxSurge, true
}

// MaxSurge sets field value
func (o *RollingStrategy) MaxSurge(v string) *RollingStrategy {
	o.maxSurge = v
    return o
}

// GetMaxUnavailable returns the MaxUnavailable field value
func (o *RollingStrategy) GetMaxUnavailable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.maxUnavailable
}

// GetMaxUnavailableOk returns a tuple with the MaxUnavailable field value
// and a boolean to check if the value has been set.
func (o *RollingStrategy) GetMaxUnavailableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.maxUnavailable, true
}

// MaxUnavailable sets field value
func (o *RollingStrategy) MaxUnavailable(v string) *RollingStrategy {
	o.maxUnavailable = v
    return o
}

// GetPartition returns the Partition field value
func (o *RollingStrategy) GetPartition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.partition
}

// GetPartitionOk returns a tuple with the Partition field value
// and a boolean to check if the value has been set.
func (o *RollingStrategy) GetPartitionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.partition, true
}

// Partition sets field value
func (o *RollingStrategy) Partition(v int32) *RollingStrategy {
	o.partition = v
    return o
}

func (o RollingStrategy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RollingStrategy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maxSurge"] = o.maxSurge
	toSerialize["maxUnavailable"] = o.maxUnavailable
	toSerialize["partition"] = o.partition
	return toSerialize, nil
}

type NullableRollingStrategy struct {
	value *RollingStrategy
	isSet bool
}

func (v NullableRollingStrategy) Get() *RollingStrategy {
	return v.value
}

func (v *NullableRollingStrategy) Set(val *RollingStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableRollingStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableRollingStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRollingStrategy(val *RollingStrategy) *NullableRollingStrategy {
	return &NullableRollingStrategy{value: val, isSet: true}
}

func (v NullableRollingStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRollingStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 