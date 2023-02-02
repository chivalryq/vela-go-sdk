/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package json_patch

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the JsonPatchSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &JsonPatchSpec{}

// JsonPatchSpec struct for JsonPatchSpec
type JsonPatchSpec struct {
	operations []map[string]interface{} `json:"operations"`
}

// NewJsonPatchSpecWith instantiates a new JsonPatchSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonPatchSpecWith(operations []map[string]interface{}) *JsonPatchSpec {
	this := JsonPatchSpec{}
	this.operations = operations
	return &this
}

// NewJsonPatchSpec instantiates a new JsonPatchSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonPatchSpec() *JsonPatchSpec {
	this := JsonPatchSpec{}
	return &this
}

// GetOperations returns the Operations field value
func (o *JSONPatchTrait) GetOperations() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Properties.operations
}

// GetOperationsOk returns a tuple with the Operations field value
// and a boolean to check if the value has been set.
func (o *JSONPatchTrait) GetOperationsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.operations, true
}

// Operations sets field value
func (o *JSONPatchTrait) Operations(v []map[string]interface{}) *JSONPatchTrait {
	o.Properties.operations = v
	return o
}

func (o JsonPatchSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JsonPatchSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operations"] = o.operations
	return toSerialize, nil
}

type NullableJsonPatchSpec struct {
	value *JsonPatchSpec
	isSet bool
}

func (v NullableJsonPatchSpec) Get() *JsonPatchSpec {
	return v.value
}

func (v *NullableJsonPatchSpec) Set(val *JsonPatchSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonPatchSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonPatchSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonPatchSpec(val *JsonPatchSpec) *NullableJsonPatchSpec {
	return &NullableJsonPatchSpec{value: val, isSet: true}
}

func (v NullableJsonPatchSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonPatchSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const JsonPatchType = "json-patch"

func init() {
	sdkcommon.RegisterTrait(JsonPatchType, FromTrait)
}

type JSONPatchTrait struct {
	Base       apis.TraitBase
	Properties JsonPatchSpec
}

func JsonPatch() *JSONPatchTrait {
	j := &JSONPatchTrait{Base: apis.TraitBase{}}
	return j
}

func (j *JSONPatchTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(j.Properties),
		Type:       JsonPatchType,
	}
	return res
}

func (j *JSONPatchTrait) FromTrait(from common.ApplicationTrait) (*JSONPatchTrait, error) {
	var properties JsonPatchSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	j.Properties = properties
	return j, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	j := &JSONPatchTrait{}
	return j.FromTrait(from)
}

func (j *JSONPatchTrait) DefType() string {
	return JsonPatchType
}
