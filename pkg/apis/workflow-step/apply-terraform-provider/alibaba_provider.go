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

// checks if the AlibabaProvider type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AlibabaProvider{}

// AlibabaProvider struct for AlibabaProvider
type AlibabaProvider struct {
	accessKey string `json:"accessKey"`
	region    string `json:"region"`
	secretKey string `json:"secretKey"`
	name      string `json:"name"`
	type_     string `json:"type"`
}

// NewAlibabaProviderWith instantiates a new AlibabaProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlibabaProviderWith(accessKey string, region string, secretKey string, name string, type_ string) *AlibabaProvider {
	this := AlibabaProvider{}
	this.accessKey = accessKey
	this.region = region
	this.secretKey = secretKey
	this.name = name
	this.type_ = type_
	return &this
}

// NewAlibabaProvider instantiates a new AlibabaProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlibabaProvider() *AlibabaProvider {
	this := AlibabaProvider{}
	var name string = "alibaba-provider"
	this.name = name
	return &this
}

// GetAccessKey returns the AccessKey field value
func (o *AlibabaProvider) GetAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.accessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
func (o *AlibabaProvider) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.accessKey, true
}

// AccessKey sets field value
func (o *AlibabaProvider) AccessKey(v string) *AlibabaProvider {
	o.accessKey = v
	return o
}

// GetRegion returns the Region field value
func (o *AlibabaProvider) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AlibabaProvider) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.region, true
}

// Region sets field value
func (o *AlibabaProvider) Region(v string) *AlibabaProvider {
	o.region = v
	return o
}

// GetSecretKey returns the SecretKey field value
func (o *AlibabaProvider) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.secretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *AlibabaProvider) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.secretKey, true
}

// SecretKey sets field value
func (o *AlibabaProvider) SecretKey(v string) *AlibabaProvider {
	o.secretKey = v
	return o
}

// GetName returns the Name field value
func (o *AlibabaProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AlibabaProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *AlibabaProvider) Name(v string) *AlibabaProvider {
	o.name = v
	return o
}

// GetType returns the Type field value
func (o *AlibabaProvider) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.type_
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AlibabaProvider) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.type_, true
}

// Type sets field value
func (o *AlibabaProvider) Type(v string) *AlibabaProvider {
	o.type_ = v
	return o
}

func (o AlibabaProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlibabaProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accessKey"] = o.accessKey
	toSerialize["region"] = o.region
	toSerialize["secretKey"] = o.secretKey
	toSerialize["name"] = o.name
	toSerialize["type"] = o.type_
	return toSerialize, nil
}

type NullableAlibabaProvider struct {
	value *AlibabaProvider
	isSet bool
}

func (v NullableAlibabaProvider) Get() *AlibabaProvider {
	return v.value
}

func (v *NullableAlibabaProvider) Set(val *AlibabaProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableAlibabaProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableAlibabaProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlibabaProvider(val *AlibabaProvider) *NullableAlibabaProvider {
	return &NullableAlibabaProvider{value: val, isSet: true}
}

func (v NullableAlibabaProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlibabaProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
