/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package shared_resource

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the SharedResourcePolicyRule type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SharedResourcePolicyRule{}

// SharedResourcePolicyRule struct for SharedResourcePolicyRule
type SharedResourcePolicyRule struct {
	// Specify how to select the targets of the rule
	Selector []ResourcePolicyRuleSelector `json:"selector,omitempty"`
}

// NewSharedResourcePolicyRuleWith instantiates a new SharedResourcePolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedResourcePolicyRuleWith() *SharedResourcePolicyRule {
	this := SharedResourcePolicyRule{}
	return &this
}

// NewSharedResourcePolicyRule instantiates a new SharedResourcePolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedResourcePolicyRule() *SharedResourcePolicyRule {
	this := SharedResourcePolicyRule{}
	return &this
}

// NewSharedResourcePolicyRules converts a list SharedResourcePolicyRule pointers to objects.
// This is helpful when the SetSharedResourcePolicyRule requires a list of objects
func NewSharedResourcePolicyRules(ps ...*SharedResourcePolicyRule) []SharedResourcePolicyRule {
	objs := []SharedResourcePolicyRule{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *SharedResourcePolicyRule) GetSelector() []ResourcePolicyRuleSelector {
	if o == nil || utils.IsNil(o.Selector) {
		var ret []ResourcePolicyRuleSelector
		return ret
	}
	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedResourcePolicyRule) GetSelectorOk() ([]ResourcePolicyRuleSelector, bool) {
	if o == nil || utils.IsNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *SharedResourcePolicyRule) HasSelector() bool {
	if o != nil && !utils.IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given []ResourcePolicyRuleSelector and assigns it to the selector field.
// Selector:  Specify how to select the targets of the rule
func (o *SharedResourcePolicyRule) SetSelector(v []ResourcePolicyRuleSelector) *SharedResourcePolicyRule {
	o.Selector = v
	return o
}

func (o SharedResourcePolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedResourcePolicyRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	return toSerialize, nil
}

type NullableSharedResourcePolicyRule struct {
	value *SharedResourcePolicyRule
	isSet bool
}

func (v NullableSharedResourcePolicyRule) Get() *SharedResourcePolicyRule {
	return v.value
}

func (v *NullableSharedResourcePolicyRule) Set(val *SharedResourcePolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedResourcePolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedResourcePolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedResourcePolicyRule(val *SharedResourcePolicyRule) *NullableSharedResourcePolicyRule {
	return &NullableSharedResourcePolicyRule{value: val, isSet: true}
}

func (v NullableSharedResourcePolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedResourcePolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
