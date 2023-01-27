/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package worker

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the WorkerSpecEnvValueFrom type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WorkerSpecEnvValueFrom{}

// WorkerSpecEnvValueFrom Specifies a source the value of this var should come from
type WorkerSpecEnvValueFrom struct {
	configMapKeyRef *WorkerSpecEnvValueFromConfigMapKeyRef `json:"configMapKeyRef,omitempty"`
	secretKeyRef *WorkerSpecEnvValueFromSecretKeyRef `json:"secretKeyRef,omitempty"`
}

// NewWorkerSpecEnvValueFrom instantiates a new WorkerSpecEnvValueFrom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkerSpecEnvValueFrom() *WorkerSpecEnvValueFrom {
	this := WorkerSpecEnvValueFrom{}
	return &this
}

// NewWorkerSpecEnvValueFromWithDefaults instantiates a new WorkerSpecEnvValueFrom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkerSpecEnvValueFromWithDefaults() *WorkerSpecEnvValueFrom {
	this := WorkerSpecEnvValueFrom{}
	return &this
}

// GetConfigMapKeyRef returns the ConfigMapKeyRef field value if set, zero value otherwise.
func (o *WorkerSpecEnvValueFrom) GetConfigMapKeyRef() WorkerSpecEnvValueFromConfigMapKeyRef {
	if o == nil || utils.IsNil(o.configMapKeyRef) {
		var ret WorkerSpecEnvValueFromConfigMapKeyRef
		return ret
	}
	return *o.configMapKeyRef
}

// GetConfigMapKeyRefOk returns a tuple with the ConfigMapKeyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerSpecEnvValueFrom) GetConfigMapKeyRefOk() (*WorkerSpecEnvValueFromConfigMapKeyRef, bool) {
	if o == nil || utils.IsNil(o.configMapKeyRef) {
		return nil, false
	}
	return o.configMapKeyRef, true
}

// HasConfigMapKeyRef returns a boolean if a field has been set.
func (o *WorkerSpecEnvValueFrom) HasConfigMapKeyRef() bool {
	if o != nil && !utils.IsNil(o.configMapKeyRef) {
		return true
	}

	return false
}

// SetConfigMapKeyRef gets a reference to the given WorkerSpecEnvValueFromConfigMapKeyRef and assigns it to the configMapKeyRef field.
// configMapKeyRef: 

func (o *WorkerSpecEnvValueFrom) ConfigMapKeyRef(v WorkerSpecEnvValueFromConfigMapKeyRef) (*WorkerSpecEnvValueFrom){
	o.configMapKeyRef = &v
return o
}

// GetSecretKeyRef returns the SecretKeyRef field value if set, zero value otherwise.
func (o *WorkerSpecEnvValueFrom) GetSecretKeyRef() WorkerSpecEnvValueFromSecretKeyRef {
	if o == nil || utils.IsNil(o.secretKeyRef) {
		var ret WorkerSpecEnvValueFromSecretKeyRef
		return ret
	}
	return *o.secretKeyRef
}

// GetSecretKeyRefOk returns a tuple with the SecretKeyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerSpecEnvValueFrom) GetSecretKeyRefOk() (*WorkerSpecEnvValueFromSecretKeyRef, bool) {
	if o == nil || utils.IsNil(o.secretKeyRef) {
		return nil, false
	}
	return o.secretKeyRef, true
}

// HasSecretKeyRef returns a boolean if a field has been set.
func (o *WorkerSpecEnvValueFrom) HasSecretKeyRef() bool {
	if o != nil && !utils.IsNil(o.secretKeyRef) {
		return true
	}

	return false
}

// SetSecretKeyRef gets a reference to the given WorkerSpecEnvValueFromSecretKeyRef and assigns it to the secretKeyRef field.
// secretKeyRef: 

func (o *WorkerSpecEnvValueFrom) SecretKeyRef(v WorkerSpecEnvValueFromSecretKeyRef) (*WorkerSpecEnvValueFrom){
	o.secretKeyRef = &v
return o
}

func (o WorkerSpecEnvValueFrom) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkerSpecEnvValueFrom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.configMapKeyRef) {
		toSerialize["configMapKeyRef"] = o.configMapKeyRef
	}
	if !utils.IsNil(o.secretKeyRef) {
		toSerialize["secretKeyRef"] = o.secretKeyRef
	}
	return toSerialize, nil
}

type NullableWorkerSpecEnvValueFrom struct {
	value *WorkerSpecEnvValueFrom
	isSet bool
}

func (v NullableWorkerSpecEnvValueFrom) Get() *WorkerSpecEnvValueFrom {
	return v.value
}

func (v *NullableWorkerSpecEnvValueFrom) Set(val *WorkerSpecEnvValueFrom) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkerSpecEnvValueFrom) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkerSpecEnvValueFrom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkerSpecEnvValueFrom(val *WorkerSpecEnvValueFrom) *NullableWorkerSpecEnvValueFrom {
	return &NullableWorkerSpecEnvValueFrom{value: val, isSet: true}
}

func (v NullableWorkerSpecEnvValueFrom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkerSpecEnvValueFrom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 