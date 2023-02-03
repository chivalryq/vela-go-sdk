/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kustomize_patch

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the KustomizePatchSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &KustomizePatchSpec{}

// KustomizePatchSpec struct for KustomizePatchSpec
type KustomizePatchSpec struct {
	// a list of StrategicMerge or JSON6902 patch to selected target
	Patches []PatchItem `json:"patches,omitempty"`
}

// NewKustomizePatchSpecWith instantiates a new KustomizePatchSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKustomizePatchSpecWith() *KustomizePatchSpec {
	this := KustomizePatchSpec{}
	return &this
}

// NewKustomizePatchSpec instantiates a new KustomizePatchSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKustomizePatchSpec() *KustomizePatchSpec {
	this := KustomizePatchSpec{}
	return &this
}

// NewKustomizePatchSpecs converts a list KustomizePatchSpec pointers to objects.
// This is helpful when the SetKustomizePatchSpec requires a list of objects
func NewKustomizePatchSpecs(ps ...*KustomizePatchSpec) []KustomizePatchSpec {
	objs := []KustomizePatchSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetPatches returns the Patches field value if set, zero value otherwise.
func (o *KustomizePatchTrait) GetPatches() []PatchItem {
	if o == nil || utils.IsNil(o.Properties.Patches) {
		var ret []PatchItem
		return ret
	}
	return o.Properties.Patches
}

// GetPatchesOk returns a tuple with the Patches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KustomizePatchTrait) GetPatchesOk() ([]PatchItem, bool) {
	if o == nil || utils.IsNil(o.Properties.Patches) {
		return nil, false
	}
	return o.Properties.Patches, true
}

// HasPatches returns a boolean if a field has been set.
func (o *KustomizePatchTrait) HasPatches() bool {
	if o != nil && !utils.IsNil(o.Properties.Patches) {
		return true
	}

	return false
}

// SetPatches gets a reference to the given []PatchItem and assigns it to the patches field.
// Patches:  a list of StrategicMerge or JSON6902 patch to selected target
func (o *KustomizePatchTrait) SetPatches(v []PatchItem) *KustomizePatchTrait {
	o.Properties.Patches = v
	return o
}

func (o KustomizePatchSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KustomizePatchSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Patches) {
		toSerialize["patches"] = o.Patches
	}
	return toSerialize, nil
}

type NullableKustomizePatchSpec struct {
	value *KustomizePatchSpec
	isSet bool
}

func (v NullableKustomizePatchSpec) Get() *KustomizePatchSpec {
	return v.value
}

func (v *NullableKustomizePatchSpec) Set(val *KustomizePatchSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableKustomizePatchSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableKustomizePatchSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKustomizePatchSpec(val *KustomizePatchSpec) *NullableKustomizePatchSpec {
	return &NullableKustomizePatchSpec{value: val, isSet: true}
}

func (v NullableKustomizePatchSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKustomizePatchSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const KustomizePatchType = "kustomize-patch"

func init() {
	sdkcommon.RegisterTrait(KustomizePatchType, FromTrait)
}

type KustomizePatchTrait struct {
	Base       apis.TraitBase
	Properties KustomizePatchSpec
}

func KustomizePatch() *KustomizePatchTrait {
	k := &KustomizePatchTrait{Base: apis.TraitBase{}}
	return k
}

func (k *KustomizePatchTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(k.Properties),
		Type:       KustomizePatchType,
	}
	return res
}

func (k *KustomizePatchTrait) FromTrait(from common.ApplicationTrait) (*KustomizePatchTrait, error) {
	var properties KustomizePatchSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	k.Base.Type = KustomizePatchType
	k.Properties = properties
	return k, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	k := &KustomizePatchTrait{}
	return k.FromTrait(from)
}

func (k *KustomizePatchTrait) DefType() string {
	return KustomizePatchType
}
