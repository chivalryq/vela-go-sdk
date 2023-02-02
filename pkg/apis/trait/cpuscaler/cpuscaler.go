/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cpuscaler

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the CpuscalerSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CpuscalerSpec{}

// CpuscalerSpec struct for CpuscalerSpec
type CpuscalerSpec struct {
	// Specify the average CPU utilization, for example, 50 means the CPU usage is 50%
	cpuUtil int32 `json:"cpuUtil"`
	// Specify the maximum number of of replicas to which the autoscaler can scale up
	max int32 `json:"max"`
	// Specify the minimal number of replicas to which the autoscaler can scale down
	min int32 `json:"min"`
	// Specify the apiVersion of scale target
	targetAPIVersion string `json:"targetAPIVersion"`
	// Specify the kind of scale target
	targetKind string `json:"targetKind"`
}

// NewCpuscalerSpecWith instantiates a new CpuscalerSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpuscalerSpecWith(cpuUtil int32, max int32, min int32, targetAPIVersion string, targetKind string) *CpuscalerSpec {
	this := CpuscalerSpec{}
	this.cpuUtil = cpuUtil
	this.max = max
	this.min = min
	this.targetAPIVersion = targetAPIVersion
	this.targetKind = targetKind
	return &this
}

// NewCpuscalerSpec instantiates a new CpuscalerSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpuscalerSpec() *CpuscalerSpec {
	this := CpuscalerSpec{}
	var cpuUtil int32 = 50
	this.cpuUtil = cpuUtil
	var max int32 = 10
	this.max = max
	var min int32 = 1
	this.min = min
	var targetAPIVersion string = "apps/v1"
	this.targetAPIVersion = targetAPIVersion
	var targetKind string = "Deployment"
	this.targetKind = targetKind
	return &this
}

// GetCpuUtil returns the CpuUtil field value
func (o *CpuscalerTrait) GetCpuUtil() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Properties.cpuUtil
}

// GetCpuUtilOk returns a tuple with the CpuUtil field value
// and a boolean to check if the value has been set.
func (o *CpuscalerTrait) GetCpuUtilOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.cpuUtil, true
}

// CpuUtil sets field value
func (o *CpuscalerTrait) CpuUtil(v int32) *CpuscalerTrait {
	o.Properties.cpuUtil = v
	return o
}

// GetMax returns the Max field value
func (o *CpuscalerTrait) GetMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Properties.max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *CpuscalerTrait) GetMaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.max, true
}

// Max sets field value
func (o *CpuscalerTrait) Max(v int32) *CpuscalerTrait {
	o.Properties.max = v
	return o
}

// GetMin returns the Min field value
func (o *CpuscalerTrait) GetMin() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Properties.min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *CpuscalerTrait) GetMinOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.min, true
}

// Min sets field value
func (o *CpuscalerTrait) Min(v int32) *CpuscalerTrait {
	o.Properties.min = v
	return o
}

// GetTargetAPIVersion returns the TargetAPIVersion field value
func (o *CpuscalerTrait) GetTargetAPIVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.targetAPIVersion
}

// GetTargetAPIVersionOk returns a tuple with the TargetAPIVersion field value
// and a boolean to check if the value has been set.
func (o *CpuscalerTrait) GetTargetAPIVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.targetAPIVersion, true
}

// TargetAPIVersion sets field value
func (o *CpuscalerTrait) TargetAPIVersion(v string) *CpuscalerTrait {
	o.Properties.targetAPIVersion = v
	return o
}

// GetTargetKind returns the TargetKind field value
func (o *CpuscalerTrait) GetTargetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.targetKind
}

// GetTargetKindOk returns a tuple with the TargetKind field value
// and a boolean to check if the value has been set.
func (o *CpuscalerTrait) GetTargetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.targetKind, true
}

// TargetKind sets field value
func (o *CpuscalerTrait) TargetKind(v string) *CpuscalerTrait {
	o.Properties.targetKind = v
	return o
}

func (o CpuscalerSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CpuscalerSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cpuUtil"] = o.cpuUtil
	toSerialize["max"] = o.max
	toSerialize["min"] = o.min
	toSerialize["targetAPIVersion"] = o.targetAPIVersion
	toSerialize["targetKind"] = o.targetKind
	return toSerialize, nil
}

type NullableCpuscalerSpec struct {
	value *CpuscalerSpec
	isSet bool
}

func (v NullableCpuscalerSpec) Get() *CpuscalerSpec {
	return v.value
}

func (v *NullableCpuscalerSpec) Set(val *CpuscalerSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCpuscalerSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCpuscalerSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpuscalerSpec(val *CpuscalerSpec) *NullableCpuscalerSpec {
	return &NullableCpuscalerSpec{value: val, isSet: true}
}

func (v NullableCpuscalerSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpuscalerSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const CpuscalerType = "cpuscaler"

func init() {
	sdkcommon.RegisterTrait(CpuscalerType, FromTrait)
}

type CpuscalerTrait struct {
	Base       apis.TraitBase
	Properties CpuscalerSpec
}

func Cpuscaler() *CpuscalerTrait {
	c := &CpuscalerTrait{Base: apis.TraitBase{}}
	return c
}

func (c *CpuscalerTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(c.Properties),
		Type:       CpuscalerType,
	}
	return res
}

func (c *CpuscalerTrait) FromTrait(from common.ApplicationTrait) (*CpuscalerTrait, error) {
	var properties CpuscalerSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	c.Properties = properties
	return c, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	c := &CpuscalerTrait{}
	return c.FromTrait(from)
}

func (c *CpuscalerTrait) DefType() string {
	return CpuscalerType
}
