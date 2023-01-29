/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package worker

import (
	"encoding/json"

	"vela-go-sdk/pkg/apis/utils"
)

// checks if the ConfigMap type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ConfigMap{}

// ConfigMap struct for ConfigMap
type ConfigMap struct {
	cmName      string  `json:"cmName"`
	defaultMode int32   `json:"defaultMode"`
	items       []Items `json:"items,omitempty"`
	mountPath   string  `json:"mountPath"`
	name        string  `json:"name"`
}

// NewConfigMapWith instantiates a new ConfigMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigMapWith(cmName string, defaultMode int32, mountPath string, name string) *ConfigMap {
	this := ConfigMap{}
	this.cmName = cmName
	this.defaultMode = defaultMode
	this.mountPath = mountPath
	this.name = name
	return &this
}

// NewConfigMap instantiates a new ConfigMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigMap() *ConfigMap {
	this := ConfigMap{}
	var defaultMode int32 = 420
	this.defaultMode = defaultMode
	return &this
}

// GetCmName returns the CmName field value
func (o *ConfigMap) GetCmName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.cmName
}

// GetCmNameOk returns a tuple with the CmName field value
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetCmNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.cmName, true
}

// CmName sets field value
func (o *ConfigMap) CmName(v string) *ConfigMap {
	o.cmName = v
	return o
}

// GetDefaultMode returns the DefaultMode field value
func (o *ConfigMap) GetDefaultMode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.defaultMode
}

// GetDefaultModeOk returns a tuple with the DefaultMode field value
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetDefaultModeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.defaultMode, true
}

// DefaultMode sets field value
func (o *ConfigMap) DefaultMode(v int32) *ConfigMap {
	o.defaultMode = v
	return o
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ConfigMap) GetItems() []Items {
	if o == nil || utils.IsNil(o.items) {
		var ret []Items
		return ret
	}
	return o.items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetItemsOk() ([]Items, bool) {
	if o == nil || utils.IsNil(o.items) {
		return nil, false
	}
	return o.items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ConfigMap) HasItems() bool {
	if o != nil && !utils.IsNil(o.items) {
		return true
	}

	return false
}

// Items gets a reference to the given []Items and assigns it to the items field.
// items:
func (o *ConfigMap) Items(v []Items) *ConfigMap {
	o.items = v
	return o
}

// GetMountPath returns the MountPath field value
func (o *ConfigMap) GetMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.mountPath
}

// GetMountPathOk returns a tuple with the MountPath field value
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.mountPath, true
}

// MountPath sets field value
func (o *ConfigMap) MountPath(v string) *ConfigMap {
	o.mountPath = v
	return o
}

// GetName returns the Name field value
func (o *ConfigMap) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *ConfigMap) Name(v string) *ConfigMap {
	o.name = v
	return o
}

func (o ConfigMap) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigMap) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cmName"] = o.cmName
	toSerialize["defaultMode"] = o.defaultMode
	if !utils.IsNil(o.items) {
		toSerialize["items"] = o.items
	}
	toSerialize["mountPath"] = o.mountPath
	toSerialize["name"] = o.name
	return toSerialize, nil
}

type NullableConfigMap struct {
	value *ConfigMap
	isSet bool
}

func (v NullableConfigMap) Get() *ConfigMap {
	return v.value
}

func (v *NullableConfigMap) Set(val *ConfigMap) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigMap) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigMap(val *ConfigMap) *NullableConfigMap {
	return &NullableConfigMap{value: val, isSet: true}
}

func (v NullableConfigMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
