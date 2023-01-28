/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affinity

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"vela-go-sdk/api"
	"vela-go-sdk/apis/utils"
)

// checks if the AffinitySpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffinitySpec{}

// AffinitySpec struct for AffinitySpec
type AffinitySpec struct {
	nodeAffinity    *AffinitySpecNodeAffinity    `json:"nodeAffinity,omitempty"`
	podAffinity     *AffinitySpecPodAffinity     `json:"podAffinity,omitempty"`
	podAntiAffinity *AffinitySpecPodAntiAffinity `json:"podAntiAffinity,omitempty"`
	// Specify tolerant taint
	tolerations []AffinitySpecTolerations `json:"tolerations,omitempty"`
}

// NewAffinitySpecWith instantiates a new AffinitySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffinitySpecWith() *AffinitySpec {
	this := AffinitySpec{}
	return &this
}

// NewAffinitySpec instantiates a new AffinitySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffinitySpec() *AffinitySpec {
	this := AffinitySpec{}
	return &this
}

// GetNodeAffinity returns the NodeAffinity field value if set, zero value otherwise.
func (o *AffinityTrait) GetNodeAffinity() AffinitySpecNodeAffinity {
	if o == nil || utils.IsNil(o.Properties.nodeAffinity) {
		var ret AffinitySpecNodeAffinity
		return ret
	}
	return *o.Properties.nodeAffinity
}

// GetNodeAffinityOk returns a tuple with the NodeAffinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffinityTrait) GetNodeAffinityOk() (*AffinitySpecNodeAffinity, bool) {
	if o == nil || utils.IsNil(o.Properties.nodeAffinity) {
		return nil, false
	}
	return o.Properties.nodeAffinity, true
}

// HasNodeAffinity returns a boolean if a field has been set.
func (o *AffinityTrait) HasNodeAffinity() bool {
	if o != nil && !utils.IsNil(o.Properties.nodeAffinity) {
		return true
	}

	return false
}

// NodeAffinity gets a reference to the given AffinitySpecNodeAffinity and assigns it to the nodeAffinity field.
// nodeAffinity:
func (o *AffinityTrait) NodeAffinity(v AffinitySpecNodeAffinity) *AffinityTrait {
	o.Properties.nodeAffinity = &v
	return o
}

// GetPodAffinity returns the PodAffinity field value if set, zero value otherwise.
func (o *AffinityTrait) GetPodAffinity() AffinitySpecPodAffinity {
	if o == nil || utils.IsNil(o.Properties.podAffinity) {
		var ret AffinitySpecPodAffinity
		return ret
	}
	return *o.Properties.podAffinity
}

// GetPodAffinityOk returns a tuple with the PodAffinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffinityTrait) GetPodAffinityOk() (*AffinitySpecPodAffinity, bool) {
	if o == nil || utils.IsNil(o.Properties.podAffinity) {
		return nil, false
	}
	return o.Properties.podAffinity, true
}

// HasPodAffinity returns a boolean if a field has been set.
func (o *AffinityTrait) HasPodAffinity() bool {
	if o != nil && !utils.IsNil(o.Properties.podAffinity) {
		return true
	}

	return false
}

// PodAffinity gets a reference to the given AffinitySpecPodAffinity and assigns it to the podAffinity field.
// podAffinity:
func (o *AffinityTrait) PodAffinity(v AffinitySpecPodAffinity) *AffinityTrait {
	o.Properties.podAffinity = &v
	return o
}

// GetPodAntiAffinity returns the PodAntiAffinity field value if set, zero value otherwise.
func (o *AffinityTrait) GetPodAntiAffinity() AffinitySpecPodAntiAffinity {
	if o == nil || utils.IsNil(o.Properties.podAntiAffinity) {
		var ret AffinitySpecPodAntiAffinity
		return ret
	}
	return *o.Properties.podAntiAffinity
}

// GetPodAntiAffinityOk returns a tuple with the PodAntiAffinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffinityTrait) GetPodAntiAffinityOk() (*AffinitySpecPodAntiAffinity, bool) {
	if o == nil || utils.IsNil(o.Properties.podAntiAffinity) {
		return nil, false
	}
	return o.Properties.podAntiAffinity, true
}

// HasPodAntiAffinity returns a boolean if a field has been set.
func (o *AffinityTrait) HasPodAntiAffinity() bool {
	if o != nil && !utils.IsNil(o.Properties.podAntiAffinity) {
		return true
	}

	return false
}

// PodAntiAffinity gets a reference to the given AffinitySpecPodAntiAffinity and assigns it to the podAntiAffinity field.
// podAntiAffinity:
func (o *AffinityTrait) PodAntiAffinity(v AffinitySpecPodAntiAffinity) *AffinityTrait {
	o.Properties.podAntiAffinity = &v
	return o
}

// GetTolerations returns the Tolerations field value if set, zero value otherwise.
func (o *AffinityTrait) GetTolerations() []AffinitySpecTolerations {
	if o == nil || utils.IsNil(o.Properties.tolerations) {
		var ret []AffinitySpecTolerations
		return ret
	}
	return o.Properties.tolerations
}

// GetTolerationsOk returns a tuple with the Tolerations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffinityTrait) GetTolerationsOk() ([]AffinitySpecTolerations, bool) {
	if o == nil || utils.IsNil(o.Properties.tolerations) {
		return nil, false
	}
	return o.Properties.tolerations, true
}

// HasTolerations returns a boolean if a field has been set.
func (o *AffinityTrait) HasTolerations() bool {
	if o != nil && !utils.IsNil(o.Properties.tolerations) {
		return true
	}

	return false
}

// Tolerations gets a reference to the given []AffinitySpecTolerations and assigns it to the tolerations field.
// tolerations:  Specify tolerant taint
func (o *AffinityTrait) Tolerations(v []AffinitySpecTolerations) *AffinityTrait {
	o.Properties.tolerations = v
	return o
}

func (o AffinitySpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AffinitySpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.nodeAffinity) {
		toSerialize["nodeAffinity"] = o.nodeAffinity
	}
	if !utils.IsNil(o.podAffinity) {
		toSerialize["podAffinity"] = o.podAffinity
	}
	if !utils.IsNil(o.podAntiAffinity) {
		toSerialize["podAntiAffinity"] = o.podAntiAffinity
	}
	if !utils.IsNil(o.tolerations) {
		toSerialize["tolerations"] = o.tolerations
	}
	return toSerialize, nil
}

type NullableAffinitySpec struct {
	value *AffinitySpec
	isSet bool
}

func (v NullableAffinitySpec) Get() *AffinitySpec {
	return v.value
}

func (v *NullableAffinitySpec) Set(val *AffinitySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAffinitySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAffinitySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffinitySpec(val *AffinitySpec) *NullableAffinitySpec {
	return &NullableAffinitySpec{value: val, isSet: true}
}

func (v NullableAffinitySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffinitySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const AffinityType = "affinity"

type AffinityTrait struct {
	Base       api.TraitBase
	Properties AffinitySpec
}

func Affinity() *AffinityTrait {
	a := &AffinityTrait{Base: api.TraitBase{}}
	return a
}

func (a *AffinityTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(a.Properties),
		Type:       AffinityType,
	}
	return res
}

func (a *AffinityTrait) Props() *AffinitySpec {
	return &a.Properties
}
