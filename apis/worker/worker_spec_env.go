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



    
    // checks if the WorkerSpecEnv type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WorkerSpecEnv{}

// WorkerSpecEnv struct for WorkerSpecEnv
type WorkerSpecEnv struct {
	// Environment variable name
	name string `json:"name"`
	// The value of the environment variable
	value *string `json:"value,omitempty"`
	valueFrom *WorkerSpecEnvValueFrom `json:"valueFrom,omitempty"`
}

// NewWorkerSpecEnv instantiates a new WorkerSpecEnv object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkerSpecEnv(name string) *WorkerSpecEnv {
	this := WorkerSpecEnv{}
	this.name = name
	return &this
}

// NewWorkerSpecEnvWithDefaults instantiates a new WorkerSpecEnv object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkerSpecEnvWithDefaults() *WorkerSpecEnv {
	this := WorkerSpecEnv{}
	return &this
}

// GetName returns the Name field value
func (o *WorkerSpecEnv) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkerSpecEnv) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *WorkerSpecEnv) Name(v string) *WorkerSpecEnv {
	o.name = v
    return o
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WorkerSpecEnv) GetValue() string {
	if o == nil || utils.IsNil(o.value) {
		var ret string
		return ret
	}
	return *o.value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerSpecEnv) GetValueOk() (*string, bool) {
	if o == nil || utils.IsNil(o.value) {
		return nil, false
	}
	return o.value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WorkerSpecEnv) HasValue() bool {
	if o != nil && !utils.IsNil(o.value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the value field.
// value:  The value of the environment variable 

func (o *WorkerSpecEnv) Value(v string) (*WorkerSpecEnv){
	o.value = &v
return o
}

// GetValueFrom returns the ValueFrom field value if set, zero value otherwise.
func (o *WorkerSpecEnv) GetValueFrom() WorkerSpecEnvValueFrom {
	if o == nil || utils.IsNil(o.valueFrom) {
		var ret WorkerSpecEnvValueFrom
		return ret
	}
	return *o.valueFrom
}

// GetValueFromOk returns a tuple with the ValueFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerSpecEnv) GetValueFromOk() (*WorkerSpecEnvValueFrom, bool) {
	if o == nil || utils.IsNil(o.valueFrom) {
		return nil, false
	}
	return o.valueFrom, true
}

// HasValueFrom returns a boolean if a field has been set.
func (o *WorkerSpecEnv) HasValueFrom() bool {
	if o != nil && !utils.IsNil(o.valueFrom) {
		return true
	}

	return false
}

// SetValueFrom gets a reference to the given WorkerSpecEnvValueFrom and assigns it to the valueFrom field.
// valueFrom: 

func (o *WorkerSpecEnv) ValueFrom(v WorkerSpecEnvValueFrom) (*WorkerSpecEnv){
	o.valueFrom = &v
return o
}

func (o WorkerSpecEnv) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkerSpecEnv) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.name
	if !utils.IsNil(o.value) {
		toSerialize["value"] = o.value
	}
	if !utils.IsNil(o.valueFrom) {
		toSerialize["valueFrom"] = o.valueFrom
	}
	return toSerialize, nil
}

type NullableWorkerSpecEnv struct {
	value *WorkerSpecEnv
	isSet bool
}

func (v NullableWorkerSpecEnv) Get() *WorkerSpecEnv {
	return v.value
}

func (v *NullableWorkerSpecEnv) Set(val *WorkerSpecEnv) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkerSpecEnv) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkerSpecEnv) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkerSpecEnv(val *WorkerSpecEnv) *NullableWorkerSpecEnv {
	return &NullableWorkerSpecEnv{value: val, isSet: true}
}

func (v NullableWorkerSpecEnv) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkerSpecEnv) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 
