/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package read_only

import (
	"encoding/json"

	"vela-go-sdk/pkg/apis/utils"
)

// checks if the PolicyRule type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PolicyRule{}

// PolicyRule struct for PolicyRule
type PolicyRule struct {
	// Specify how to select the targets of the rule
	selector []RuleSelector `json:"selector"`
}

// NewPolicyRuleWith instantiates a new PolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyRuleWith(selector []RuleSelector) *PolicyRule {
	this := PolicyRule{}
	this.selector = selector
	return &this
}

// NewPolicyRule instantiates a new PolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyRule() *PolicyRule {
	this := PolicyRule{}
	return &this
}

// GetSelector returns the Selector field value
func (o *PolicyRule) GetSelector() []RuleSelector {
	if o == nil {
		var ret []RuleSelector
		return ret
	}

	return o.selector
}

// GetSelectorOk returns a tuple with the Selector field value
// and a boolean to check if the value has been set.
func (o *PolicyRule) GetSelectorOk() ([]RuleSelector, bool) {
	if o == nil {
		return nil, false
	}
	return o.selector, true
}

// Selector sets field value
func (o *PolicyRule) Selector(v []RuleSelector) *PolicyRule {
	o.selector = v
	return o
}

func (o PolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["selector"] = o.selector
	return toSerialize, nil
}

type NullablePolicyRule struct {
	value *PolicyRule
	isSet bool
}

func (v NullablePolicyRule) Get() *PolicyRule {
	return v.value
}

func (v *NullablePolicyRule) Set(val *PolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyRule(val *PolicyRule) *NullablePolicyRule {
	return &NullablePolicyRule{value: val, isSet: true}
}

func (v NullablePolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
