/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webservice

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the Env type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Env{}

// Env struct for Env
type Env struct {
	// Environment variable name
	name string `json:"name"`
	// The value of the environment variable
	value     *string    `json:"value,omitempty"`
	valueFrom *ValueFrom `json:"valueFrom,omitempty"`
}

// NewEnvWith instantiates a new Env object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvWith(name string) *Env {
	this := Env{}
	this.name = name
	return &this
}

// NewEnv instantiates a new Env object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnv() *Env {
	this := Env{}
	return &this
}

// GetName returns the Name field value
func (o *Env) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Env) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *Env) Name(v string) *Env {
	o.name = v
	return o
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Env) GetValue() string {
	if o == nil || utils.IsNil(o.value) {
		var ret string
		return ret
	}
	return *o.value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Env) GetValueOk() (*string, bool) {
	if o == nil || utils.IsNil(o.value) {
		return nil, false
	}
	return o.value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Env) HasValue() bool {
	if o != nil && !utils.IsNil(o.value) {
		return true
	}

	return false
}

// Value gets a reference to the given string and assigns it to the value field.
// value:  The value of the environment variable
func (o *Env) Value(v string) *Env {
	o.value = &v
	return o
}

// GetValueFrom returns the ValueFrom field value if set, zero value otherwise.
func (o *Env) GetValueFrom() ValueFrom {
	if o == nil || utils.IsNil(o.valueFrom) {
		var ret ValueFrom
		return ret
	}
	return *o.valueFrom
}

// GetValueFromOk returns a tuple with the ValueFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Env) GetValueFromOk() (*ValueFrom, bool) {
	if o == nil || utils.IsNil(o.valueFrom) {
		return nil, false
	}
	return o.valueFrom, true
}

// HasValueFrom returns a boolean if a field has been set.
func (o *Env) HasValueFrom() bool {
	if o != nil && !utils.IsNil(o.valueFrom) {
		return true
	}

	return false
}

// ValueFrom gets a reference to the given ValueFrom and assigns it to the valueFrom field.
// valueFrom:
func (o *Env) ValueFrom(v ValueFrom) *Env {
	o.valueFrom = &v
	return o
}

func (o Env) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Env) ToMap() (map[string]interface{}, error) {
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

type NullableEnv struct {
	value *Env
	isSet bool
}

func (v NullableEnv) Get() *Env {
	return v.value
}

func (v *NullableEnv) Set(val *Env) {
	v.value = val
	v.isSet = true
}

func (v NullableEnv) IsSet() bool {
	return v.isSet
}

func (v *NullableEnv) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnv(val *Env) *NullableEnv {
	return &NullableEnv{value: val, isSet: true}
}

func (v NullableEnv) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnv) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
