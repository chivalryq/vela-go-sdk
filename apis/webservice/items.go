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



    
    // checks if the Items type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Items{}

// Items struct for Items
type Items struct {
	key string `json:"key"`
	mode int32 `json:"mode"`
	path string `json:"path"`
}

// NewItems instantiates a new Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItems(key string, mode int32, path string) *Items {
	this := Items{}
	this.key = key
	this.mode = mode
	this.path = path
	return &this
}

// NewItemsWithDefaults instantiates a new Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemsWithDefaults() *Items {
	this := Items{}
	var mode int32 = 511
	this.mode = mode
	return &this
}

// GetKey returns the Key field value
func (o *Items) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Items) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.key, true
}

// Key sets field value
func (o *Items) Key(v string) *Items {
	o.key = v
    return o
}

// GetMode returns the Mode field value
func (o *Items) GetMode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *Items) GetModeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.mode, true
}

// Mode sets field value
func (o *Items) Mode(v int32) *Items {
	o.mode = v
    return o
}

// GetPath returns the Path field value
func (o *Items) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *Items) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.path, true
}

// Path sets field value
func (o *Items) Path(v string) *Items {
	o.path = v
    return o
}

func (o Items) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Items) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.key
	toSerialize["mode"] = o.mode
	toSerialize["path"] = o.path
	return toSerialize, nil
}

type NullableItems struct {
	value *Items
	isSet bool
}

func (v NullableItems) Get() *Items {
	return v.value
}

func (v *NullableItems) Set(val *Items) {
	v.value = val
	v.isSet = true
}

func (v NullableItems) IsSet() bool {
	return v.isSet
}

func (v *NullableItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItems(val *Items) *NullableItems {
	return &NullableItems{value: val, isSet: true}
}

func (v NullableItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 
