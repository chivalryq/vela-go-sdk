/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package take_over

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the TakeOverSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TakeOverSpec{}

// TakeOverSpec struct for TakeOverSpec
type TakeOverSpec struct {
	// Specify the list of rules to control take over strategy at resource level. The selected resource will be able to be taken over by the current application when the resource belongs to no one.
	Rules []PolicyRule `json:"rules,omitempty"`
}

// NewTakeOverSpecWith instantiates a new TakeOverSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTakeOverSpecWith() *TakeOverSpec {
	this := TakeOverSpec{}
	return &this
}

// NewTakeOverSpec instantiates a new TakeOverSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTakeOverSpec() *TakeOverSpec {
	this := TakeOverSpec{}
	return &this
}

// NewTakeOverSpecs converts a list TakeOverSpec pointers to objects.
// This is helpful when the SetTakeOverSpec requires a list of objects
func NewTakeOverSpecs(ps ...*TakeOverSpec) []TakeOverSpec {
	objs := []TakeOverSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *TakeOverPolicy) GetRules() []PolicyRule {
	if o == nil || utils.IsNil(o.Properties.Rules) {
		var ret []PolicyRule
		return ret
	}
	return o.Properties.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TakeOverPolicy) GetRulesOk() ([]PolicyRule, bool) {
	if o == nil || utils.IsNil(o.Properties.Rules) {
		return nil, false
	}
	return o.Properties.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *TakeOverPolicy) HasRules() bool {
	if o != nil && !utils.IsNil(o.Properties.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []PolicyRule and assigns it to the rules field.
// Rules:  Specify the list of rules to control take over strategy at resource level. The selected resource will be able to be taken over by the current application when the resource belongs to no one.
func (o *TakeOverPolicy) SetRules(v []PolicyRule) *TakeOverPolicy {
	o.Properties.Rules = v
	return o
}

func (o TakeOverSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TakeOverSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableTakeOverSpec struct {
	value *TakeOverSpec
	isSet bool
}

func (v NullableTakeOverSpec) Get() *TakeOverSpec {
	return v.value
}

func (v *NullableTakeOverSpec) Set(val *TakeOverSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTakeOverSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTakeOverSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTakeOverSpec(val *TakeOverSpec) *NullableTakeOverSpec {
	return &NullableTakeOverSpec{value: val, isSet: true}
}

func (v NullableTakeOverSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTakeOverSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const TakeOverType = "take-over"

func init() {
	sdkcommon.RegisterPolicy(TakeOverType, FromPolicy)
}

type TakeOverPolicy struct {
	Base       apis.PolicyBase
	Properties TakeOverSpec
}

func TakeOver(name string) *TakeOverPolicy {
	t := &TakeOverPolicy{Base: apis.PolicyBase{
		Name: name,
		Type: TakeOverType,
	}}
	return t
}

func (t *TakeOverPolicy) Build() v1beta1.AppPolicy {
	res := v1beta1.AppPolicy{
		Name:       t.Base.Name,
		Properties: util.Object2RawExtension(t.Properties),
		Type:       TakeOverType,
	}
	return res
}

func (t *TakeOverPolicy) FromPolicy(from v1beta1.AppPolicy) (*TakeOverPolicy, error) {
	var properties TakeOverSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	t.Base.Name = from.Name
	t.Base.Type = TakeOverType
	t.Properties = properties
	return t, nil
}

func FromPolicy(from v1beta1.AppPolicy) (apis.Policy, error) {
	t := &TakeOverPolicy{}
	return t.FromPolicy(from)
}

func (t *TakeOverPolicy) PolicyName() string {
	return t.Base.Name
}

func (t *TakeOverPolicy) DefType() string {
	return TakeOverType
}
