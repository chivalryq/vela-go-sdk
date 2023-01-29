/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package topologyspreadconstraints

import (
	"encoding/json"
	"vela-go-sdk/pkg/apis"
	"vela-go-sdk/pkg/apis/utils"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)

// checks if the TopologyspreadconstraintsSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TopologyspreadconstraintsSpec{}

// TopologyspreadconstraintsSpec struct for TopologyspreadconstraintsSpec
type TopologyspreadconstraintsSpec struct {
	constraints []Constraints `json:"constraints"`
}

// NewTopologyspreadconstraintsSpecWith instantiates a new TopologyspreadconstraintsSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopologyspreadconstraintsSpecWith(constraints []Constraints) *TopologyspreadconstraintsSpec {
	this := TopologyspreadconstraintsSpec{}
	this.constraints = constraints
	return &this
}

// NewTopologyspreadconstraintsSpec instantiates a new TopologyspreadconstraintsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopologyspreadconstraintsSpec() *TopologyspreadconstraintsSpec {
	this := TopologyspreadconstraintsSpec{}
	return &this
}

// GetConstraints returns the Constraints field value
func (o *TopologyspreadconstraintsTrait) GetConstraints() []Constraints {
	if o == nil {
		var ret []Constraints
		return ret
	}

	return o.Properties.constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value
// and a boolean to check if the value has been set.
func (o *TopologyspreadconstraintsTrait) GetConstraintsOk() ([]Constraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.constraints, true
}

// Constraints sets field value
func (o *TopologyspreadconstraintsTrait) Constraints(v []Constraints) *TopologyspreadconstraintsTrait {
	o.Properties.constraints = v
	return o
}

func (o TopologyspreadconstraintsSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopologyspreadconstraintsSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["constraints"] = o.constraints
	return toSerialize, nil
}

type NullableTopologyspreadconstraintsSpec struct {
	value *TopologyspreadconstraintsSpec
	isSet bool
}

func (v NullableTopologyspreadconstraintsSpec) Get() *TopologyspreadconstraintsSpec {
	return v.value
}

func (v *NullableTopologyspreadconstraintsSpec) Set(val *TopologyspreadconstraintsSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTopologyspreadconstraintsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTopologyspreadconstraintsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopologyspreadconstraintsSpec(val *TopologyspreadconstraintsSpec) *NullableTopologyspreadconstraintsSpec {
	return &NullableTopologyspreadconstraintsSpec{value: val, isSet: true}
}

func (v NullableTopologyspreadconstraintsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopologyspreadconstraintsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const TopologyspreadconstraintsType = "topologyspreadconstraints"

type TopologyspreadconstraintsTrait struct {
	Base       apis.TraitBase
	Properties TopologyspreadconstraintsSpec
}

func Topologyspreadconstraints() *TopologyspreadconstraintsTrait {
	t := &TopologyspreadconstraintsTrait{Base: apis.TraitBase{}}
	return t
}

func (t *TopologyspreadconstraintsTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(t.Properties),
		Type:       TopologyspreadconstraintsType,
	}
	return res
}