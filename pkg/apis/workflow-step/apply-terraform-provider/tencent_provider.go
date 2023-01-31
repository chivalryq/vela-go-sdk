/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_terraform_provider

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the TencentProvider type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TencentProvider{}

// TencentProvider struct for TencentProvider
type TencentProvider struct {
	name      string `json:"name"`
	region    string `json:"region"`
	secretID  string `json:"secretID"`
	secretKey string `json:"secretKey"`
	type_     string `json:"type"`
}

// NewTencentProviderWith instantiates a new TencentProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTencentProviderWith(name string, region string, secretID string, secretKey string, type_ string) *TencentProvider {
	this := TencentProvider{}
	this.name = name
	this.region = region
	this.secretID = secretID
	this.secretKey = secretKey
	this.type_ = type_
	return &this
}

// NewTencentProvider instantiates a new TencentProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTencentProvider() *TencentProvider {
	this := TencentProvider{}
	var name string = "tencent-provider"
	this.name = name
	return &this
}

// GetName returns the Name field value
func (o *TencentProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TencentProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *TencentProvider) Name(v string) *TencentProvider {
	o.name = v
	return o
}

// GetRegion returns the Region field value
func (o *TencentProvider) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *TencentProvider) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.region, true
}

// Region sets field value
func (o *TencentProvider) Region(v string) *TencentProvider {
	o.region = v
	return o
}

// GetSecretID returns the SecretID field value
func (o *TencentProvider) GetSecretID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.secretID
}

// GetSecretIDOk returns a tuple with the SecretID field value
// and a boolean to check if the value has been set.
func (o *TencentProvider) GetSecretIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.secretID, true
}

// SecretID sets field value
func (o *TencentProvider) SecretID(v string) *TencentProvider {
	o.secretID = v
	return o
}

// GetSecretKey returns the SecretKey field value
func (o *TencentProvider) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.secretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *TencentProvider) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.secretKey, true
}

// SecretKey sets field value
func (o *TencentProvider) SecretKey(v string) *TencentProvider {
	o.secretKey = v
	return o
}

// GetType returns the Type field value
func (o *TencentProvider) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.type_
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TencentProvider) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.type_, true
}

// Type sets field value
func (o *TencentProvider) Type(v string) *TencentProvider {
	o.type_ = v
	return o
}

func (o TencentProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TencentProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.name
	toSerialize["region"] = o.region
	toSerialize["secretID"] = o.secretID
	toSerialize["secretKey"] = o.secretKey
	toSerialize["type"] = o.type_
	return toSerialize, nil
}

type NullableTencentProvider struct {
	value *TencentProvider
	isSet bool
}

func (v NullableTencentProvider) Get() *TencentProvider {
	return v.value
}

func (v *NullableTencentProvider) Set(val *TencentProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableTencentProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableTencentProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTencentProvider(val *TencentProvider) *NullableTencentProvider {
	return &NullableTencentProvider{value: val, isSet: true}
}

func (v NullableTencentProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTencentProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
