/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affinity

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the AffinitySpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffinitySpec{}

// AffinitySpec struct for AffinitySpec
type AffinitySpec struct {
	nodeAffinity *NodeAffinity `json:"nodeAffinity,omitempty"`
	podAffinity *PodAffinity `json:"podAffinity,omitempty"`
	podAntiAffinity *PodAntiAffinity `json:"podAntiAffinity,omitempty"`
	// Specify tolerant taint
	tolerations []Tolerations `json:"tolerations,omitempty"`
}

// NewAffinitySpec instantiates a new AffinitySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffinitySpec() *AffinitySpec {
	this := AffinitySpec{}
	return &this
}

// NewAffinitySpecWithDefaults instantiates a new AffinitySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffinitySpecWithDefaults() *AffinitySpec {
	this := AffinitySpec{}
	return &this
}

// GetNodeAffinity returns the NodeAffinity field value if set, zero value otherwise.
func (o *AffinitySpec) GetNodeAffinity() NodeAffinity {
	if o == nil || utils.IsNil(o.nodeAffinity) {
		var ret NodeAffinity
		return ret
	}
	return *o.nodeAffinity
}

// GetNodeAffinityOk returns a tuple with the NodeAffinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffinitySpec) GetNodeAffinityOk() (*NodeAffinity, bool) {
	if o == nil || utils.IsNil(o.nodeAffinity) {
		return nil, false
	}
	return o.nodeAffinity, true
}

// HasNodeAffinity returns a boolean if a field has been set.
func (o *AffinitySpec) HasNodeAffinity() bool {
	if o != nil && !utils.IsNil(o.nodeAffinity) {
		return true
	}

	return false
}

// SetNodeAffinity gets a reference to the given NodeAffinity and assigns it to the nodeAffinity field.
// nodeAffinity: 

func (o *AffinitySpec) NodeAffinity(v NodeAffinity) (*AffinitySpec){
	o.nodeAffinity = &v
return o
}

// GetPodAffinity returns the PodAffinity field value if set, zero value otherwise.
func (o *AffinitySpec) GetPodAffinity() PodAffinity {
	if o == nil || utils.IsNil(o.podAffinity) {
		var ret PodAffinity
		return ret
	}
	return *o.podAffinity
}

// GetPodAffinityOk returns a tuple with the PodAffinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffinitySpec) GetPodAffinityOk() (*PodAffinity, bool) {
	if o == nil || utils.IsNil(o.podAffinity) {
		return nil, false
	}
	return o.podAffinity, true
}

// HasPodAffinity returns a boolean if a field has been set.
func (o *AffinitySpec) HasPodAffinity() bool {
	if o != nil && !utils.IsNil(o.podAffinity) {
		return true
	}

	return false
}

// SetPodAffinity gets a reference to the given PodAffinity and assigns it to the podAffinity field.
// podAffinity: 

func (o *AffinitySpec) PodAffinity(v PodAffinity) (*AffinitySpec){
	o.podAffinity = &v
return o
}

// GetPodAntiAffinity returns the PodAntiAffinity field value if set, zero value otherwise.
func (o *AffinitySpec) GetPodAntiAffinity() PodAntiAffinity {
	if o == nil || utils.IsNil(o.podAntiAffinity) {
		var ret PodAntiAffinity
		return ret
	}
	return *o.podAntiAffinity
}

// GetPodAntiAffinityOk returns a tuple with the PodAntiAffinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffinitySpec) GetPodAntiAffinityOk() (*PodAntiAffinity, bool) {
	if o == nil || utils.IsNil(o.podAntiAffinity) {
		return nil, false
	}
	return o.podAntiAffinity, true
}

// HasPodAntiAffinity returns a boolean if a field has been set.
func (o *AffinitySpec) HasPodAntiAffinity() bool {
	if o != nil && !utils.IsNil(o.podAntiAffinity) {
		return true
	}

	return false
}

// SetPodAntiAffinity gets a reference to the given PodAntiAffinity and assigns it to the podAntiAffinity field.
// podAntiAffinity: 

func (o *AffinitySpec) PodAntiAffinity(v PodAntiAffinity) (*AffinitySpec){
	o.podAntiAffinity = &v
return o
}

// GetTolerations returns the Tolerations field value if set, zero value otherwise.
func (o *AffinitySpec) GetTolerations() []Tolerations {
	if o == nil || utils.IsNil(o.tolerations) {
		var ret []Tolerations
		return ret
	}
	return o.tolerations
}

// GetTolerationsOk returns a tuple with the Tolerations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffinitySpec) GetTolerationsOk() ([]Tolerations, bool) {
	if o == nil || utils.IsNil(o.tolerations) {
		return nil, false
	}
	return o.tolerations, true
}

// HasTolerations returns a boolean if a field has been set.
func (o *AffinitySpec) HasTolerations() bool {
	if o != nil && !utils.IsNil(o.tolerations) {
		return true
	}

	return false
}

// SetTolerations gets a reference to the given []Tolerations and assigns it to the tolerations field.
// tolerations:  Specify tolerant taint 

func (o *AffinitySpec) Tolerations(v []Tolerations) (*AffinitySpec){
	o.tolerations = v
return o
}

func (o AffinitySpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
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
	Base  TraitBase
	Props AffinitySpec
}

func Affinity() *AffinityTrait {
	a := &AffinityTrait{Base: TraitBase{}}
	return a
}

func (a *AffinityTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(a.Props),
		Type:       AffinityType,
	}
	return res
}
