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



    
    // checks if the AffinitySpecPodAntiAffinity type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffinitySpecPodAntiAffinity{}

// AffinitySpecPodAntiAffinity Specify the pod anti-affinity scheduling rules
type AffinitySpecPodAntiAffinity struct {
	// Specify the preferred during scheduling ignored during execution
	preferred []AffinitySpecPodAffinityPreferred `json:"preferred,omitempty"`
	// Specify the required during scheduling ignored during execution
	required []PodAffinityTerm `json:"required,omitempty"`
}

// NewAffinitySpecPodAntiAffinity instantiates a new AffinitySpecPodAntiAffinity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffinitySpecPodAntiAffinity() *AffinitySpecPodAntiAffinity {
	this := AffinitySpecPodAntiAffinity{}
	return &this
}

// NewAffinitySpecPodAntiAffinityWithDefaults instantiates a new AffinitySpecPodAntiAffinity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffinitySpecPodAntiAffinityWithDefaults() *AffinitySpecPodAntiAffinity {
	this := AffinitySpecPodAntiAffinity{}
	return &this
}

// GetPreferred returns the Preferred field value if set, zero value otherwise.
func (o *AffinitySpecPodAntiAffinity) GetPreferred() []AffinitySpecPodAffinityPreferred {
	if o == nil || utils.IsNil(o.preferred) {
		var ret []AffinitySpecPodAffinityPreferred
		return ret
	}
	return o.preferred
}

// GetPreferredOk returns a tuple with the Preferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffinitySpecPodAntiAffinity) GetPreferredOk() ([]AffinitySpecPodAffinityPreferred, bool) {
	if o == nil || utils.IsNil(o.preferred) {
		return nil, false
	}
	return o.preferred, true
}

// HasPreferred returns a boolean if a field has been set.
func (o *AffinitySpecPodAntiAffinity) HasPreferred() bool {
	if o != nil && !utils.IsNil(o.preferred) {
		return true
	}

	return false
}

// SetPreferred gets a reference to the given []AffinitySpecPodAffinityPreferred and assigns it to the preferred field.
// preferred:  Specify the preferred during scheduling ignored during execution 

func (o *AffinitySpecPodAntiAffinity) Preferred(v []AffinitySpecPodAffinityPreferred) (*AffinitySpecPodAntiAffinity){
	o.preferred = v
return o
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *AffinitySpecPodAntiAffinity) GetRequired() []PodAffinityTerm {
	if o == nil || utils.IsNil(o.required) {
		var ret []PodAffinityTerm
		return ret
	}
	return o.required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffinitySpecPodAntiAffinity) GetRequiredOk() ([]PodAffinityTerm, bool) {
	if o == nil || utils.IsNil(o.required) {
		return nil, false
	}
	return o.required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *AffinitySpecPodAntiAffinity) HasRequired() bool {
	if o != nil && !utils.IsNil(o.required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given []PodAffinityTerm and assigns it to the required field.
// required:  Specify the required during scheduling ignored during execution 

func (o *AffinitySpecPodAntiAffinity) Required(v []PodAffinityTerm) (*AffinitySpecPodAntiAffinity){
	o.required = v
return o
}

func (o AffinitySpecPodAntiAffinity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AffinitySpecPodAntiAffinity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.preferred) {
		toSerialize["preferred"] = o.preferred
	}
	if !utils.IsNil(o.required) {
		toSerialize["required"] = o.required
	}
	return toSerialize, nil
}

type NullableAffinitySpecPodAntiAffinity struct {
	value *AffinitySpecPodAntiAffinity
	isSet bool
}

func (v NullableAffinitySpecPodAntiAffinity) Get() *AffinitySpecPodAntiAffinity {
	return v.value
}

func (v *NullableAffinitySpecPodAntiAffinity) Set(val *AffinitySpecPodAntiAffinity) {
	v.value = val
	v.isSet = true
}

func (v NullableAffinitySpecPodAntiAffinity) IsSet() bool {
	return v.isSet
}

func (v *NullableAffinitySpecPodAntiAffinity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffinitySpecPodAntiAffinity(val *AffinitySpecPodAntiAffinity) *NullableAffinitySpecPodAntiAffinity {
	return &NullableAffinitySpecPodAntiAffinity{value: val, isSet: true}
}

func (v NullableAffinitySpecPodAntiAffinity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffinitySpecPodAntiAffinity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 