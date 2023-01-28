/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affinity

import (
	"encoding/json"

	"vela-go-sdk/apis/utils"
)

// checks if the AffinitySpecNodeAffinityRequired type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffinitySpecNodeAffinityRequired{}

// AffinitySpecNodeAffinityRequired Specify the required during scheduling ignored during execution
type AffinitySpecNodeAffinityRequired struct {
	// Specify a list of node selector
	nodeSelectorTerms []NodeSelectorTerm `json:"nodeSelectorTerms"`
}

// NewAffinitySpecNodeAffinityRequiredWith instantiates a new AffinitySpecNodeAffinityRequired object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffinitySpecNodeAffinityRequiredWith(nodeSelectorTerms []NodeSelectorTerm) *AffinitySpecNodeAffinityRequired {
	this := AffinitySpecNodeAffinityRequired{}
	this.nodeSelectorTerms = nodeSelectorTerms
	return &this
}

// NewAffinitySpecNodeAffinityRequired instantiates a new AffinitySpecNodeAffinityRequired object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffinitySpecNodeAffinityRequired() *AffinitySpecNodeAffinityRequired {
	this := AffinitySpecNodeAffinityRequired{}
	return &this
}

// GetNodeSelectorTerms returns the NodeSelectorTerms field value
func (o *AffinitySpecNodeAffinityRequired) GetNodeSelectorTerms() []NodeSelectorTerm {
	if o == nil {
		var ret []NodeSelectorTerm
		return ret
	}

	return o.nodeSelectorTerms
}

// GetNodeSelectorTermsOk returns a tuple with the NodeSelectorTerms field value
// and a boolean to check if the value has been set.
func (o *AffinitySpecNodeAffinityRequired) GetNodeSelectorTermsOk() ([]NodeSelectorTerm, bool) {
	if o == nil {
		return nil, false
	}
	return o.nodeSelectorTerms, true
}

// NodeSelectorTerms sets field value
func (o *AffinitySpecNodeAffinityRequired) NodeSelectorTerms(v []NodeSelectorTerm) *AffinitySpecNodeAffinityRequired {
	o.nodeSelectorTerms = v
	return o
}

func (o AffinitySpecNodeAffinityRequired) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AffinitySpecNodeAffinityRequired) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nodeSelectorTerms"] = o.nodeSelectorTerms
	return toSerialize, nil
}

type NullableAffinitySpecNodeAffinityRequired struct {
	value *AffinitySpecNodeAffinityRequired
	isSet bool
}

func (v NullableAffinitySpecNodeAffinityRequired) Get() *AffinitySpecNodeAffinityRequired {
	return v.value
}

func (v *NullableAffinitySpecNodeAffinityRequired) Set(val *AffinitySpecNodeAffinityRequired) {
	v.value = val
	v.isSet = true
}

func (v NullableAffinitySpecNodeAffinityRequired) IsSet() bool {
	return v.isSet
}

func (v *NullableAffinitySpecNodeAffinityRequired) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffinitySpecNodeAffinityRequired(val *AffinitySpecNodeAffinityRequired) *NullableAffinitySpecNodeAffinityRequired {
	return &NullableAffinitySpecNodeAffinityRequired{value: val, isSet: true}
}

func (v NullableAffinitySpecNodeAffinityRequired) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffinitySpecNodeAffinityRequired) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
