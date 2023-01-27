/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webservice

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the WebserviceSpecEnvValueFromConfigMapKeyRef type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WebserviceSpecEnvValueFromConfigMapKeyRef{}

// WebserviceSpecEnvValueFromConfigMapKeyRef Selects a key of a config map in the pod's namespace
type WebserviceSpecEnvValueFromConfigMapKeyRef struct {
	// The key of the config map to select from. Must be a valid secret key
	key string `json:"key"`
	// The name of the config map in the pod's namespace to select from
	name string `json:"name"`
}

// NewWebserviceSpecEnvValueFromConfigMapKeyRef instantiates a new WebserviceSpecEnvValueFromConfigMapKeyRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebserviceSpecEnvValueFromConfigMapKeyRef(key string, name string) *WebserviceSpecEnvValueFromConfigMapKeyRef {
	this := WebserviceSpecEnvValueFromConfigMapKeyRef{}
	this.key = key
	this.name = name
	return &this
}

// NewWebserviceSpecEnvValueFromConfigMapKeyRefWithDefaults instantiates a new WebserviceSpecEnvValueFromConfigMapKeyRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebserviceSpecEnvValueFromConfigMapKeyRefWithDefaults() *WebserviceSpecEnvValueFromConfigMapKeyRef {
	this := WebserviceSpecEnvValueFromConfigMapKeyRef{}
	return &this
}

// GetKey returns the Key field value
func (o *WebserviceSpecEnvValueFromConfigMapKeyRef) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *WebserviceSpecEnvValueFromConfigMapKeyRef) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.key, true
}

// Key sets field value
func (o *WebserviceSpecEnvValueFromConfigMapKeyRef) Key(v string) *WebserviceSpecEnvValueFromConfigMapKeyRef {
	o.key = v
    return o
}

// GetName returns the Name field value
func (o *WebserviceSpecEnvValueFromConfigMapKeyRef) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebserviceSpecEnvValueFromConfigMapKeyRef) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *WebserviceSpecEnvValueFromConfigMapKeyRef) Name(v string) *WebserviceSpecEnvValueFromConfigMapKeyRef {
	o.name = v
    return o
}

func (o WebserviceSpecEnvValueFromConfigMapKeyRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebserviceSpecEnvValueFromConfigMapKeyRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.key
	toSerialize["name"] = o.name
	return toSerialize, nil
}

type NullableWebserviceSpecEnvValueFromConfigMapKeyRef struct {
	value *WebserviceSpecEnvValueFromConfigMapKeyRef
	isSet bool
}

func (v NullableWebserviceSpecEnvValueFromConfigMapKeyRef) Get() *WebserviceSpecEnvValueFromConfigMapKeyRef {
	return v.value
}

func (v *NullableWebserviceSpecEnvValueFromConfigMapKeyRef) Set(val *WebserviceSpecEnvValueFromConfigMapKeyRef) {
	v.value = val
	v.isSet = true
}

func (v NullableWebserviceSpecEnvValueFromConfigMapKeyRef) IsSet() bool {
	return v.isSet
}

func (v *NullableWebserviceSpecEnvValueFromConfigMapKeyRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebserviceSpecEnvValueFromConfigMapKeyRef(val *WebserviceSpecEnvValueFromConfigMapKeyRef) *NullableWebserviceSpecEnvValueFromConfigMapKeyRef {
	return &NullableWebserviceSpecEnvValueFromConfigMapKeyRef{value: val, isSet: true}
}

func (v NullableWebserviceSpecEnvValueFromConfigMapKeyRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebserviceSpecEnvValueFromConfigMapKeyRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 
