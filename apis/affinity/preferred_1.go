/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affinity

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the Preferred1 type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Preferred1{}

// Preferred1 struct for Preferred1
type Preferred1 struct {
	podAffinityTerm PodAffinityTerm `json:"podAffinityTerm"`
	// Specify weight associated with matching the corresponding podAffinityTerm
	weight int32 `json:"weight"`
}

// NewPreferred1 instantiates a new Preferred1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreferred1(podAffinityTerm PodAffinityTerm, weight int32) *Preferred1 {
	this := Preferred1{}
	this.podAffinityTerm = podAffinityTerm
	this.weight = weight
	return &this
}

// NewPreferred1WithDefaults instantiates a new Preferred1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferred1WithDefaults() *Preferred1 {
	this := Preferred1{}
	return &this
}

// GetPodAffinityTerm returns the PodAffinityTerm field value
func (o *Preferred1) GetPodAffinityTerm() PodAffinityTerm {
	if o == nil {
		var ret PodAffinityTerm
		return ret
	}

	return o.podAffinityTerm
}

// GetPodAffinityTermOk returns a tuple with the PodAffinityTerm field value
// and a boolean to check if the value has been set.
func (o *Preferred1) GetPodAffinityTermOk() (*PodAffinityTerm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.podAffinityTerm, true
}

// PodAffinityTerm sets field value
func (o *Preferred1) PodAffinityTerm(v PodAffinityTerm) *Preferred1 {
	o.podAffinityTerm = v
    return o
}

// GetWeight returns the Weight field value
func (o *Preferred1) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *Preferred1) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.weight, true
}

// Weight sets field value
func (o *Preferred1) Weight(v int32) *Preferred1 {
	o.weight = v
    return o
}

func (o Preferred1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Preferred1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["podAffinityTerm"] = o.podAffinityTerm
	toSerialize["weight"] = o.weight
	return toSerialize, nil
}

type NullablePreferred1 struct {
	value *Preferred1
	isSet bool
}

func (v NullablePreferred1) Get() *Preferred1 {
	return v.value
}

func (v *NullablePreferred1) Set(val *Preferred1) {
	v.value = val
	v.isSet = true
}

func (v NullablePreferred1) IsSet() bool {
	return v.isSet
}

func (v *NullablePreferred1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreferred1(val *Preferred1) *NullablePreferred1 {
	return &NullablePreferred1{value: val, isSet: true}
}

func (v NullablePreferred1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreferred1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 
