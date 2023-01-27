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



    
    // checks if the Required type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Required{}

// Required Specify the required during scheduling ignored during execution
type Required struct {
	// Specify a list of node selector
	nodeSelectorTerms []NodeSelectorTerm `json:"nodeSelectorTerms"`
}

// NewRequired instantiates a new Required object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequired(nodeSelectorTerms []NodeSelectorTerm) *Required {
	this := Required{}
	this.nodeSelectorTerms = nodeSelectorTerms
	return &this
}

// NewRequiredWithDefaults instantiates a new Required object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequiredWithDefaults() *Required {
	this := Required{}
	return &this
}

// GetNodeSelectorTerms returns the NodeSelectorTerms field value
func (o *Required) GetNodeSelectorTerms() []NodeSelectorTerm {
	if o == nil {
		var ret []NodeSelectorTerm
		return ret
	}

	return o.nodeSelectorTerms
}

// GetNodeSelectorTermsOk returns a tuple with the NodeSelectorTerms field value
// and a boolean to check if the value has been set.
func (o *Required) GetNodeSelectorTermsOk() ([]NodeSelectorTerm, bool) {
	if o == nil {
		return nil, false
	}
	return o.nodeSelectorTerms, true
}

// NodeSelectorTerms sets field value
func (o *Required) NodeSelectorTerms(v []NodeSelectorTerm) *Required {
	o.nodeSelectorTerms = v
    return o
}

func (o Required) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Required) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nodeSelectorTerms"] = o.nodeSelectorTerms
	return toSerialize, nil
}

type NullableRequired struct {
	value *Required
	isSet bool
}

func (v NullableRequired) Get() *Required {
	return v.value
}

func (v *NullableRequired) Set(val *Required) {
	v.value = val
	v.isSet = true
}

func (v NullableRequired) IsSet() bool {
	return v.isSet
}

func (v *NullableRequired) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequired(val *Required) *NullableRequired {
	return &NullableRequired{value: val, isSet: true}
}

func (v NullableRequired) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequired) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 
