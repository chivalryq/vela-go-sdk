/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package read_only

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the ReadOnlySpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReadOnlySpec{}

// ReadOnlySpec struct for ReadOnlySpec
type ReadOnlySpec struct {
	// Specify the list of rules to control read only strategy at resource level. The selected resource will be read-only to the current application. If the target resource does not exist, error will be raised.
	Rules []PolicyRule `json:"rules,omitempty"`
}

// NewReadOnlySpecWith instantiates a new ReadOnlySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadOnlySpecWith() *ReadOnlySpec {
	this := ReadOnlySpec{}
	return &this
}

// NewReadOnlySpec instantiates a new ReadOnlySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadOnlySpec() *ReadOnlySpec {
	this := ReadOnlySpec{}
	return &this
}

// NewReadOnlySpecs converts a list ReadOnlySpec pointers to objects.
// This is helpful when the SetReadOnlySpec requires a list of objects
func NewReadOnlySpecs(ps ...*ReadOnlySpec) []ReadOnlySpec {
	objs := []ReadOnlySpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *ReadOnlyPolicy) GetRules() []PolicyRule {
	if o == nil || utils.IsNil(o.Properties.Rules) {
		var ret []PolicyRule
		return ret
	}
	return o.Properties.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadOnlyPolicy) GetRulesOk() ([]PolicyRule, bool) {
	if o == nil || utils.IsNil(o.Properties.Rules) {
		return nil, false
	}
	return o.Properties.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *ReadOnlyPolicy) HasRules() bool {
	if o != nil && !utils.IsNil(o.Properties.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []PolicyRule and assigns it to the rules field.
// Rules:  Specify the list of rules to control read only strategy at resource level. The selected resource will be read-only to the current application. If the target resource does not exist, error will be raised.
func (o *ReadOnlyPolicy) SetRules(v []PolicyRule) *ReadOnlyPolicy {
	o.Properties.Rules = v
	return o
}

func (o ReadOnlySpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadOnlySpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableReadOnlySpec struct {
	value *ReadOnlySpec
	isSet bool
}

func (v NullableReadOnlySpec) Get() *ReadOnlySpec {
	return v.value
}

func (v *NullableReadOnlySpec) Set(val *ReadOnlySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableReadOnlySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableReadOnlySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadOnlySpec(val *ReadOnlySpec) *NullableReadOnlySpec {
	return &NullableReadOnlySpec{value: val, isSet: true}
}

func (v NullableReadOnlySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadOnlySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ReadOnlyType = "read-only"

func init() {
	sdkcommon.RegisterPolicy(ReadOnlyType, FromPolicy)
}

type ReadOnlyPolicy struct {
	Base       apis.PolicyBase
	Properties ReadOnlySpec
}

func ReadOnly(name string) *ReadOnlyPolicy {
	r := &ReadOnlyPolicy{Base: apis.PolicyBase{
		Name: name,
		Type: ReadOnlyType,
	}}
	return r
}

func (r *ReadOnlyPolicy) Build() v1beta1.AppPolicy {
	res := v1beta1.AppPolicy{
		Name:       r.Base.Name,
		Properties: util.Object2RawExtension(r.Properties),
		Type:       ReadOnlyType,
	}
	return res
}

func (r *ReadOnlyPolicy) FromPolicy(from v1beta1.AppPolicy) (*ReadOnlyPolicy, error) {
	var properties ReadOnlySpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	r.Base.Name = from.Name
	r.Base.Type = ReadOnlyType
	r.Properties = properties
	return r, nil
}

func FromPolicy(from v1beta1.AppPolicy) (apis.Policy, error) {
	r := &ReadOnlyPolicy{}
	return r.FromPolicy(from)
}

func (r *ReadOnlyPolicy) PolicyName() string {
	return r.Base.Name
}

func (r *ReadOnlyPolicy) DefType() string {
	return ReadOnlyType
}
