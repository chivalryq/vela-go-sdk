/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package annotations

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the AnnotationsSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AnnotationsSpec{}

// AnnotationsSpec struct for AnnotationsSpec
type AnnotationsSpec struct {
}

// NewAnnotationsSpecWith instantiates a new AnnotationsSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnnotationsSpecWith() *AnnotationsSpec {
	this := AnnotationsSpec{}
	return &this
}

// NewAnnotationsSpec instantiates a new AnnotationsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnnotationsSpec() *AnnotationsSpec {
	this := AnnotationsSpec{}
	return &this
}

func (o AnnotationsSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnnotationsSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableAnnotationsSpec struct {
	value *AnnotationsSpec
	isSet bool
}

func (v NullableAnnotationsSpec) Get() *AnnotationsSpec {
	return v.value
}

func (v *NullableAnnotationsSpec) Set(val *AnnotationsSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnotationsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnotationsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnotationsSpec(val *AnnotationsSpec) *NullableAnnotationsSpec {
	return &NullableAnnotationsSpec{value: val, isSet: true}
}

func (v NullableAnnotationsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnotationsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const AnnotationsType = "annotations"

func init() {
	sdkcommon.RegisterTrait(AnnotationsType, FromTrait)
}

type AnnotationsTrait struct {
	Base       apis.TraitBase
	Properties AnnotationsSpec
}

func Annotations() *AnnotationsTrait {
	a := &AnnotationsTrait{Base: apis.TraitBase{}}
	return a
}

func (a *AnnotationsTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(a.Properties),
		Type:       AnnotationsType,
	}
	return res
}

func (a *AnnotationsTrait) FromTrait(from common.ApplicationTrait) (*AnnotationsTrait, error) {
	var properties AnnotationsSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	a.Properties = properties
	return a, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	a := &AnnotationsTrait{}
	return a.FromTrait(from)
}

func (a *AnnotationsTrait) DefType() string {
	return AnnotationsType
}
