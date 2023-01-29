/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_terraform_provider

import (
	"encoding/json"
	"vela-go-sdk/pkg/apis/utils"
)

// checks if the AWSProvider type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AWSProvider{}

// AWSProvider struct for AWSProvider
type AWSProvider struct {
	accessKey string `json:"accessKey"`
	region    string `json:"region"`
	secretKey string `json:"secretKey"`
	name      string `json:"name"`
	token     string `json:"token"`
	type_     string `json:"type"`
}

// NewAWSProviderWith instantiates a new AWSProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSProviderWith(accessKey string, region string, secretKey string, name string, token string, type_ string) *AWSProvider {
	this := AWSProvider{}
	this.accessKey = accessKey
	this.region = region
	this.secretKey = secretKey
	this.name = name
	this.token = token
	this.type_ = type_
	return &this
}

// NewAWSProvider instantiates a new AWSProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSProvider() *AWSProvider {
	this := AWSProvider{}
	var name string = "aws-provider"
	this.name = name
	var token string = ""
	this.token = token
	return &this
}

// GetAccessKey returns the AccessKey field value
func (o *AWSProvider) GetAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.accessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
func (o *AWSProvider) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.accessKey, true
}

// AccessKey sets field value
func (o *AWSProvider) AccessKey(v string) *AWSProvider {
	o.accessKey = v
	return o
}

// GetRegion returns the Region field value
func (o *AWSProvider) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AWSProvider) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.region, true
}

// Region sets field value
func (o *AWSProvider) Region(v string) *AWSProvider {
	o.region = v
	return o
}

// GetSecretKey returns the SecretKey field value
func (o *AWSProvider) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.secretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *AWSProvider) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.secretKey, true
}

// SecretKey sets field value
func (o *AWSProvider) SecretKey(v string) *AWSProvider {
	o.secretKey = v
	return o
}

// GetName returns the Name field value
func (o *AWSProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AWSProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *AWSProvider) Name(v string) *AWSProvider {
	o.name = v
	return o
}

// GetToken returns the Token field value
func (o *AWSProvider) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *AWSProvider) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.token, true
}

// Token sets field value
func (o *AWSProvider) Token(v string) *AWSProvider {
	o.token = v
	return o
}

// GetType returns the Type field value
func (o *AWSProvider) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.type_
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AWSProvider) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.type_, true
}

// Type sets field value
func (o *AWSProvider) Type(v string) *AWSProvider {
	o.type_ = v
	return o
}

func (o AWSProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AWSProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accessKey"] = o.accessKey
	toSerialize["region"] = o.region
	toSerialize["secretKey"] = o.secretKey
	toSerialize["name"] = o.name
	toSerialize["token"] = o.token
	toSerialize["type"] = o.type_
	return toSerialize, nil
}

type NullableAWSProvider struct {
	value *AWSProvider
	isSet bool
}

func (v NullableAWSProvider) Get() *AWSProvider {
	return v.value
}

func (v *NullableAWSProvider) Set(val *AWSProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSProvider(val *AWSProvider) *NullableAWSProvider {
	return &NullableAWSProvider{value: val, isSet: true}
}

func (v NullableAWSProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}