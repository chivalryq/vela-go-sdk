/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kustomize_strategy_merge

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the KustomizeStrategyMergeSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &KustomizeStrategyMergeSpec{}

// KustomizeStrategyMergeSpec struct for KustomizeStrategyMergeSpec
type KustomizeStrategyMergeSpec struct {
	// a list of strategicmerge, defined as inline yaml objects.
	PatchesStrategicMerge []map[string]interface{} `json:"patchesStrategicMerge,omitempty"`
}

// NewKustomizeStrategyMergeSpecWith instantiates a new KustomizeStrategyMergeSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKustomizeStrategyMergeSpecWith() *KustomizeStrategyMergeSpec {
	this := KustomizeStrategyMergeSpec{}
	return &this
}

// NewKustomizeStrategyMergeSpec instantiates a new KustomizeStrategyMergeSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKustomizeStrategyMergeSpec() *KustomizeStrategyMergeSpec {
	this := KustomizeStrategyMergeSpec{}
	return &this
}

// NewKustomizeStrategyMergeSpecs converts a list KustomizeStrategyMergeSpec pointers to objects.
// This is helpful when the SetKustomizeStrategyMergeSpec requires a list of objects
func NewKustomizeStrategyMergeSpecs(ps ...*KustomizeStrategyMergeSpec) []KustomizeStrategyMergeSpec {
	objs := []KustomizeStrategyMergeSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetPatchesStrategicMerge returns the PatchesStrategicMerge field value if set, zero value otherwise.
func (o *KustomizeStrategyMergeTrait) GetPatchesStrategicMerge() []map[string]interface{} {
	if o == nil || utils.IsNil(o.Properties.PatchesStrategicMerge) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Properties.PatchesStrategicMerge
}

// GetPatchesStrategicMergeOk returns a tuple with the PatchesStrategicMerge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KustomizeStrategyMergeTrait) GetPatchesStrategicMergeOk() ([]map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Properties.PatchesStrategicMerge) {
		return nil, false
	}
	return o.Properties.PatchesStrategicMerge, true
}

// HasPatchesStrategicMerge returns a boolean if a field has been set.
func (o *KustomizeStrategyMergeTrait) HasPatchesStrategicMerge() bool {
	if o != nil && !utils.IsNil(o.Properties.PatchesStrategicMerge) {
		return true
	}

	return false
}

// SetPatchesStrategicMerge gets a reference to the given []map[string]interface{} and assigns it to the patchesStrategicMerge field.
// PatchesStrategicMerge:  a list of strategicmerge, defined as inline yaml objects.
func (o *KustomizeStrategyMergeTrait) SetPatchesStrategicMerge(v []map[string]interface{}) *KustomizeStrategyMergeTrait {
	o.Properties.PatchesStrategicMerge = v
	return o
}

func (o KustomizeStrategyMergeSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KustomizeStrategyMergeSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.PatchesStrategicMerge) {
		toSerialize["patchesStrategicMerge"] = o.PatchesStrategicMerge
	}
	return toSerialize, nil
}

type NullableKustomizeStrategyMergeSpec struct {
	value *KustomizeStrategyMergeSpec
	isSet bool
}

func (v NullableKustomizeStrategyMergeSpec) Get() *KustomizeStrategyMergeSpec {
	return v.value
}

func (v *NullableKustomizeStrategyMergeSpec) Set(val *KustomizeStrategyMergeSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableKustomizeStrategyMergeSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableKustomizeStrategyMergeSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKustomizeStrategyMergeSpec(val *KustomizeStrategyMergeSpec) *NullableKustomizeStrategyMergeSpec {
	return &NullableKustomizeStrategyMergeSpec{value: val, isSet: true}
}

func (v NullableKustomizeStrategyMergeSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKustomizeStrategyMergeSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const KustomizeStrategyMergeType = "kustomize-strategy-merge"

func init() {
	sdkcommon.RegisterTrait(KustomizeStrategyMergeType, FromTrait)
}

type KustomizeStrategyMergeTrait struct {
	Base       apis.TraitBase
	Properties KustomizeStrategyMergeSpec
}

func KustomizeStrategyMerge() *KustomizeStrategyMergeTrait {
	k := &KustomizeStrategyMergeTrait{Base: apis.TraitBase{}}
	return k
}

func (k *KustomizeStrategyMergeTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(k.Properties),
		Type:       KustomizeStrategyMergeType,
	}
	return res
}

func (k *KustomizeStrategyMergeTrait) FromTrait(from common.ApplicationTrait) (*KustomizeStrategyMergeTrait, error) {
	var properties KustomizeStrategyMergeSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	k.Base.Type = KustomizeStrategyMergeType
	k.Properties = properties
	return k, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	k := &KustomizeStrategyMergeTrait{}
	return k.FromTrait(from)
}

func (k *KustomizeStrategyMergeTrait) DefType() string {
	return KustomizeStrategyMergeType
}
