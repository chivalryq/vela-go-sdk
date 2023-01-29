/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lifecycle

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"vela-go-sdk/pkg/apis"
	"vela-go-sdk/pkg/apis/utils"
)

// checks if the LifecycleSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &LifecycleSpec{}

// LifecycleSpec struct for LifecycleSpec
type LifecycleSpec struct {
	postStart *LifeCycleHandler `json:"postStart,omitempty"`
	preStop   *LifeCycleHandler `json:"preStop,omitempty"`
}

// NewLifecycleSpecWith instantiates a new LifecycleSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLifecycleSpecWith() *LifecycleSpec {
	this := LifecycleSpec{}
	return &this
}

// NewLifecycleSpec instantiates a new LifecycleSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLifecycleSpec() *LifecycleSpec {
	this := LifecycleSpec{}
	return &this
}

// GetPostStart returns the PostStart field value if set, zero value otherwise.
func (o *LifecycleTrait) GetPostStart() LifeCycleHandler {
	if o == nil || utils.IsNil(o.Properties.postStart) {
		var ret LifeCycleHandler
		return ret
	}
	return *o.Properties.postStart
}

// GetPostStartOk returns a tuple with the PostStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleTrait) GetPostStartOk() (*LifeCycleHandler, bool) {
	if o == nil || utils.IsNil(o.Properties.postStart) {
		return nil, false
	}
	return o.Properties.postStart, true
}

// HasPostStart returns a boolean if a field has been set.
func (o *LifecycleTrait) HasPostStart() bool {
	if o != nil && !utils.IsNil(o.Properties.postStart) {
		return true
	}

	return false
}

// PostStart gets a reference to the given LifeCycleHandler and assigns it to the postStart field.
// postStart:
func (o *LifecycleTrait) PostStart(v LifeCycleHandler) *LifecycleTrait {
	o.Properties.postStart = &v
	return o
}

// GetPreStop returns the PreStop field value if set, zero value otherwise.
func (o *LifecycleTrait) GetPreStop() LifeCycleHandler {
	if o == nil || utils.IsNil(o.Properties.preStop) {
		var ret LifeCycleHandler
		return ret
	}
	return *o.Properties.preStop
}

// GetPreStopOk returns a tuple with the PreStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleTrait) GetPreStopOk() (*LifeCycleHandler, bool) {
	if o == nil || utils.IsNil(o.Properties.preStop) {
		return nil, false
	}
	return o.Properties.preStop, true
}

// HasPreStop returns a boolean if a field has been set.
func (o *LifecycleTrait) HasPreStop() bool {
	if o != nil && !utils.IsNil(o.Properties.preStop) {
		return true
	}

	return false
}

// PreStop gets a reference to the given LifeCycleHandler and assigns it to the preStop field.
// preStop:
func (o *LifecycleTrait) PreStop(v LifeCycleHandler) *LifecycleTrait {
	o.Properties.preStop = &v
	return o
}

func (o LifecycleSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LifecycleSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.postStart) {
		toSerialize["postStart"] = o.postStart
	}
	if !utils.IsNil(o.preStop) {
		toSerialize["preStop"] = o.preStop
	}
	return toSerialize, nil
}

type NullableLifecycleSpec struct {
	value *LifecycleSpec
	isSet bool
}

func (v NullableLifecycleSpec) Get() *LifecycleSpec {
	return v.value
}

func (v *NullableLifecycleSpec) Set(val *LifecycleSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableLifecycleSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableLifecycleSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLifecycleSpec(val *LifecycleSpec) *NullableLifecycleSpec {
	return &NullableLifecycleSpec{value: val, isSet: true}
}

func (v NullableLifecycleSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLifecycleSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const LifecycleType = "lifecycle"

type LifecycleTrait struct {
	Base       apis.TraitBase
	Properties LifecycleSpec
}

func Lifecycle() *LifecycleTrait {
	l := &LifecycleTrait{Base: apis.TraitBase{}}
	return l
}

func (l *LifecycleTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(l.Properties),
		Type:       LifecycleType,
	}
	return res
}

func (l *LifecycleTrait) Type() string {
	return LifecycleType
}
