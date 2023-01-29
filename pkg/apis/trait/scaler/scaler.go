/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package scaler

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"vela-go-sdk/pkg/apis"
	"vela-go-sdk/pkg/apis/utils"
)

// checks if the ScalerSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ScalerSpec{}

// ScalerSpec struct for ScalerSpec
type ScalerSpec struct {
	// Specify the number of workload
	replicas int32 `json:"replicas"`
}

// NewScalerSpecWith instantiates a new ScalerSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScalerSpecWith(replicas int32) *ScalerSpec {
	this := ScalerSpec{}
	this.replicas = replicas
	return &this
}

// NewScalerSpec instantiates a new ScalerSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScalerSpec() *ScalerSpec {
	this := ScalerSpec{}
	var replicas int32 = 1
	this.replicas = replicas
	return &this
}

// GetReplicas returns the Replicas field value
func (o *ScalerTrait) GetReplicas() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Properties.replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *ScalerTrait) GetReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.replicas, true
}

// Replicas sets field value
func (o *ScalerTrait) Replicas(v int32) *ScalerTrait {
	o.Properties.replicas = v
	return o
}

func (o ScalerSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScalerSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["replicas"] = o.replicas
	return toSerialize, nil
}

type NullableScalerSpec struct {
	value *ScalerSpec
	isSet bool
}

func (v NullableScalerSpec) Get() *ScalerSpec {
	return v.value
}

func (v *NullableScalerSpec) Set(val *ScalerSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableScalerSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableScalerSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScalerSpec(val *ScalerSpec) *NullableScalerSpec {
	return &NullableScalerSpec{value: val, isSet: true}
}

func (v NullableScalerSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScalerSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ScalerType = "scaler"

type ScalerTrait struct {
	Base       apis.TraitBase
	Properties ScalerSpec
}

func Scaler() *ScalerTrait {
	s := &ScalerTrait{Base: apis.TraitBase{}}
	return s
}

func (s *ScalerTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(s.Properties),
		Type:       ScalerType,
	}
	return res
}

func (s *ScalerTrait) Type() string {
	return ScalerType
}
