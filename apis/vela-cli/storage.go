/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vela_cli

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the Storage type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Storage{}

// Storage struct for Storage
type Storage struct {
	// Declare host path type storage
	hostPath []HostPath `json:"hostPath,omitempty"`
	// Mount Secret type storage
	secret []Secret `json:"secret,omitempty"`
}

// NewStorage instantiates a new Storage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorage() *Storage {
	this := Storage{}
	return &this
}

// NewStorageWithDefaults instantiates a new Storage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageWithDefaults() *Storage {
	this := Storage{}
	return &this
}

// GetHostPath returns the HostPath field value if set, zero value otherwise.
func (o *Storage) GetHostPath() []HostPath {
	if o == nil || utils.IsNil(o.hostPath) {
		var ret []HostPath
		return ret
	}
	return o.hostPath
}

// GetHostPathOk returns a tuple with the HostPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetHostPathOk() ([]HostPath, bool) {
	if o == nil || utils.IsNil(o.hostPath) {
		return nil, false
	}
	return o.hostPath, true
}

// HasHostPath returns a boolean if a field has been set.
func (o *Storage) HasHostPath() bool {
	if o != nil && !utils.IsNil(o.hostPath) {
		return true
	}

	return false
}

// SetHostPath gets a reference to the given []HostPath and assigns it to the hostPath field.
// hostPath:  Declare host path type storage 

func (o *Storage) HostPath(v []HostPath) (*Storage){
	o.hostPath = v
return o
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *Storage) GetSecret() []Secret {
	if o == nil || utils.IsNil(o.secret) {
		var ret []Secret
		return ret
	}
	return o.secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetSecretOk() ([]Secret, bool) {
	if o == nil || utils.IsNil(o.secret) {
		return nil, false
	}
	return o.secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *Storage) HasSecret() bool {
	if o != nil && !utils.IsNil(o.secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given []Secret and assigns it to the secret field.
// secret:  Mount Secret type storage 

func (o *Storage) Secret(v []Secret) (*Storage){
	o.secret = v
return o
}

func (o Storage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Storage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.hostPath) {
		toSerialize["hostPath"] = o.hostPath
	}
	if !utils.IsNil(o.secret) {
		toSerialize["secret"] = o.secret
	}
	return toSerialize, nil
}

type NullableStorage struct {
	value *Storage
	isSet bool
}

func (v NullableStorage) Get() *Storage {
	return v.value
}

func (v *NullableStorage) Set(val *Storage) {
	v.value = val
	v.isSet = true
}

func (v NullableStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorage(val *Storage) *NullableStorage {
	return &NullableStorage{value: val, isSet: true}
}

func (v NullableStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 
