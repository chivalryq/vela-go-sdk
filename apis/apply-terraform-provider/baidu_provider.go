/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_terraform_provider

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the BaiduProvider type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BaiduProvider{}

// BaiduProvider struct for BaiduProvider
type BaiduProvider struct {
	name *string `json:"name,omitempty"`
	type_ *string `json:"type,omitempty"`
	accessKey string `json:"accessKey"`
	region string `json:"region"`
	secretKey string `json:"secretKey"`
}

// NewBaiduProvider instantiates a new BaiduProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaiduProvider(accessKey string, region string, secretKey string) *BaiduProvider {
	this := BaiduProvider{}
	this.accessKey = accessKey
	this.region = region
	this.secretKey = secretKey
	return &this
}

// NewBaiduProviderWithDefaults instantiates a new BaiduProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaiduProviderWithDefaults() *BaiduProvider {
	this := BaiduProvider{}
	var name string = "baidu-provider"
	this.name = &name
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BaiduProvider) GetName() string {
	if o == nil || utils.IsNil(o.name) {
		var ret string
		return ret
	}
	return *o.name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaiduProvider) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.name) {
		return nil, false
	}
	return o.name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaiduProvider) HasName() bool {
	if o != nil && !utils.IsNil(o.name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// name: 

func (o *BaiduProvider) Name(v string) (*BaiduProvider){
	o.name = &v
return o
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BaiduProvider) GetType() string {
	if o == nil || utils.IsNil(o.type_) {
		var ret string
		return ret
	}
	return *o.type_
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaiduProvider) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.type_) {
		return nil, false
	}
	return o.type_, true
}

// HasType returns a boolean if a field has been set.
func (o *BaiduProvider) HasType() bool {
	if o != nil && !utils.IsNil(o.type_) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the type_ field.
// type_: 

func (o *BaiduProvider) Type(v string) (*BaiduProvider){
	o.type_ = &v
return o
}

// GetAccessKey returns the AccessKey field value
func (o *BaiduProvider) GetAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.accessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
func (o *BaiduProvider) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.accessKey, true
}

// AccessKey sets field value
func (o *BaiduProvider) AccessKey(v string) *BaiduProvider {
	o.accessKey = v
    return o
}

// GetRegion returns the Region field value
func (o *BaiduProvider) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *BaiduProvider) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.region, true
}

// Region sets field value
func (o *BaiduProvider) Region(v string) *BaiduProvider {
	o.region = v
    return o
}

// GetSecretKey returns the SecretKey field value
func (o *BaiduProvider) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.secretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *BaiduProvider) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.secretKey, true
}

// SecretKey sets field value
func (o *BaiduProvider) SecretKey(v string) *BaiduProvider {
	o.secretKey = v
    return o
}

func (o BaiduProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaiduProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.name) {
		toSerialize["name"] = o.name
	}
	if !utils.IsNil(o.type_) {
		toSerialize["type"] = o.type_
	}
	toSerialize["accessKey"] = o.accessKey
	toSerialize["region"] = o.region
	toSerialize["secretKey"] = o.secretKey
	return toSerialize, nil
}

type NullableBaiduProvider struct {
	value *BaiduProvider
	isSet bool
}

func (v NullableBaiduProvider) Get() *BaiduProvider {
	return v.value
}

func (v *NullableBaiduProvider) Set(val *BaiduProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableBaiduProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableBaiduProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaiduProvider(val *BaiduProvider) *NullableBaiduProvider {
	return &NullableBaiduProvider{value: val, isSet: true}
}

func (v NullableBaiduProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaiduProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 