/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package daemon

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the Secret type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Secret{}

// Secret struct for Secret
type Secret struct {
	defaultMode int32   `json:"defaultMode"`
	items       []Items `json:"items,omitempty"`
	mountPath   string  `json:"mountPath"`
	name        string  `json:"name"`
	secretName  string  `json:"secretName"`
}

// NewSecretWith instantiates a new Secret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretWith(defaultMode int32, mountPath string, name string, secretName string) *Secret {
	this := Secret{}
	this.defaultMode = defaultMode
	this.mountPath = mountPath
	this.name = name
	this.secretName = secretName
	return &this
}

// NewSecret instantiates a new Secret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecret() *Secret {
	this := Secret{}
	var defaultMode int32 = 420
	this.defaultMode = defaultMode
	return &this
}

// GetDefaultMode returns the DefaultMode field value
func (o *Secret) GetDefaultMode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.defaultMode
}

// GetDefaultModeOk returns a tuple with the DefaultMode field value
// and a boolean to check if the value has been set.
func (o *Secret) GetDefaultModeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.defaultMode, true
}

// DefaultMode sets field value
func (o *Secret) DefaultMode(v int32) *Secret {
	o.defaultMode = v
	return o
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Secret) GetItems() []Items {
	if o == nil || utils.IsNil(o.items) {
		var ret []Items
		return ret
	}
	return o.items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetItemsOk() ([]Items, bool) {
	if o == nil || utils.IsNil(o.items) {
		return nil, false
	}
	return o.items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Secret) HasItems() bool {
	if o != nil && !utils.IsNil(o.items) {
		return true
	}

	return false
}

// Items gets a reference to the given []Items and assigns it to the items field.
// items:
func (o *Secret) Items(v []Items) *Secret {
	o.items = v
	return o
}

// GetMountPath returns the MountPath field value
func (o *Secret) GetMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.mountPath
}

// GetMountPathOk returns a tuple with the MountPath field value
// and a boolean to check if the value has been set.
func (o *Secret) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.mountPath, true
}

// MountPath sets field value
func (o *Secret) MountPath(v string) *Secret {
	o.mountPath = v
	return o
}

// GetName returns the Name field value
func (o *Secret) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Secret) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *Secret) Name(v string) *Secret {
	o.name = v
	return o
}

// GetSecretName returns the SecretName field value
func (o *Secret) GetSecretName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.secretName
}

// GetSecretNameOk returns a tuple with the SecretName field value
// and a boolean to check if the value has been set.
func (o *Secret) GetSecretNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.secretName, true
}

// SecretName sets field value
func (o *Secret) SecretName(v string) *Secret {
	o.secretName = v
	return o
}

func (o Secret) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Secret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["defaultMode"] = o.defaultMode
	if !utils.IsNil(o.items) {
		toSerialize["items"] = o.items
	}
	toSerialize["mountPath"] = o.mountPath
	toSerialize["name"] = o.name
	toSerialize["secretName"] = o.secretName
	return toSerialize, nil
}

type NullableSecret struct {
	value *Secret
	isSet bool
}

func (v NullableSecret) Get() *Secret {
	return v.value
}

func (v *NullableSecret) Set(val *Secret) {
	v.value = val
	v.isSet = true
}

func (v NullableSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecret(val *Secret) *NullableSecret {
	return &NullableSecret{value: val, isSet: true}
}

func (v NullableSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
