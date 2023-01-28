/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package garbage_collect

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1alpha1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"vela-go-sdk/api"
	"vela-go-sdk/apis/utils"
)

// checks if the GarbageCollectSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GarbageCollectSpec{}

// GarbageCollectSpec struct for GarbageCollectSpec
type GarbageCollectSpec struct {
	// If is set, outdated versioned resourcetracker will not be recycled automatically, outdated resources will be kept until resourcetracker be deleted manually
	keepLegacyResource bool `json:"keepLegacyResource"`
	// Specify the list of rules to control gc strategy at resource level, if one resource is controlled by multiple rules, first rule will be used
	rules []GarbageCollectPolicyRule `json:"rules,omitempty"`
}

// NewGarbageCollectSpecWith instantiates a new GarbageCollectSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGarbageCollectSpecWith(keepLegacyResource bool) *GarbageCollectSpec {
	this := GarbageCollectSpec{}
	this.keepLegacyResource = keepLegacyResource
	return &this
}

// NewGarbageCollectSpec instantiates a new GarbageCollectSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGarbageCollectSpec() *GarbageCollectSpec {
	this := GarbageCollectSpec{}
	var keepLegacyResource bool = false
	this.keepLegacyResource = keepLegacyResource
	return &this
}

// GetKeepLegacyResource returns the KeepLegacyResource field value
func (o *GarbageCollectPolicy) GetKeepLegacyResource() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Properties.keepLegacyResource
}

// GetKeepLegacyResourceOk returns a tuple with the KeepLegacyResource field value
// and a boolean to check if the value has been set.
func (o *GarbageCollectPolicy) GetKeepLegacyResourceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.keepLegacyResource, true
}

// KeepLegacyResource sets field value
func (o *GarbageCollectPolicy) KeepLegacyResource(v bool) *GarbageCollectPolicy {
	o.Properties.keepLegacyResource = v
	return o
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *GarbageCollectPolicy) GetRules() []GarbageCollectPolicyRule {
	if o == nil || utils.IsNil(o.Properties.rules) {
		var ret []GarbageCollectPolicyRule
		return ret
	}
	return o.Properties.rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GarbageCollectPolicy) GetRulesOk() ([]GarbageCollectPolicyRule, bool) {
	if o == nil || utils.IsNil(o.Properties.rules) {
		return nil, false
	}
	return o.Properties.rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *GarbageCollectPolicy) HasRules() bool {
	if o != nil && !utils.IsNil(o.Properties.rules) {
		return true
	}

	return false
}

// Rules gets a reference to the given []GarbageCollectPolicyRule and assigns it to the rules field.
// rules:  Specify the list of rules to control gc strategy at resource level, if one resource is controlled by multiple rules, first rule will be used
func (o *GarbageCollectPolicy) Rules(v []GarbageCollectPolicyRule) *GarbageCollectPolicy {
	o.Properties.rules = v
	return o
}

func (o GarbageCollectSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GarbageCollectSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keepLegacyResource"] = o.keepLegacyResource
	if !utils.IsNil(o.rules) {
		toSerialize["rules"] = o.rules
	}
	return toSerialize, nil
}

type NullableGarbageCollectSpec struct {
	value *GarbageCollectSpec
	isSet bool
}

func (v NullableGarbageCollectSpec) Get() *GarbageCollectSpec {
	return v.value
}

func (v *NullableGarbageCollectSpec) Set(val *GarbageCollectSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableGarbageCollectSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGarbageCollectSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGarbageCollectSpec(val *GarbageCollectSpec) *NullableGarbageCollectSpec {
	return &NullableGarbageCollectSpec{value: val, isSet: true}
}

func (v NullableGarbageCollectSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGarbageCollectSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const GarbageCollectType = "garbage-collect"

type GarbageCollectPolicy struct {
	Base       api.PolicyBase
	Properties GarbageCollectSpec
}

func GarbageCollect() *GarbageCollectPolicy {
	g := &GarbageCollectPolicy{Base: api.PolicyBase{}}
	return g
}

func (g *GarbageCollectPolicy) Build() v1alpha1.Policy {
	res := v1alpha1.Policy{
		Properties: util.Object2RawExtension(g.Properties),
		Type:       GarbageCollectType,
	}
	return res
}

func (g *GarbageCollectPolicy) Props() *GarbageCollectSpec {
	return &g.Properties
}