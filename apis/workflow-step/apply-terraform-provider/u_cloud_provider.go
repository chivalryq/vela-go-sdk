/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_terraform_provider

import (
	"encoding/json"

	"vela-go-sdk/apis/utils"
)

// checks if the UCloudProvider type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UCloudProvider{}

// UCloudProvider struct for UCloudProvider
type UCloudProvider struct {
	name       string `json:"name"`
	privateKey string `json:"privateKey"`
	projectID  string `json:"projectID"`
	publicKey  string `json:"publicKey"`
	region     string `json:"region"`
	type_      string `json:"type"`
}

// NewUCloudProviderWith instantiates a new UCloudProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUCloudProviderWith(name string, privateKey string, projectID string, publicKey string, region string, type_ string) *UCloudProvider {
	this := UCloudProvider{}
	this.name = name
	this.privateKey = privateKey
	this.projectID = projectID
	this.publicKey = publicKey
	this.region = region
	this.type_ = type_
	return &this
}

// NewUCloudProvider instantiates a new UCloudProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUCloudProvider() *UCloudProvider {
	this := UCloudProvider{}
	var name string = "ucloud-provider"
	this.name = name
	return &this
}

// GetName returns the Name field value
func (o *UCloudProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UCloudProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *UCloudProvider) Name(v string) *UCloudProvider {
	o.name = v
	return o
}

// GetPrivateKey returns the PrivateKey field value
func (o *UCloudProvider) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.privateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *UCloudProvider) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.privateKey, true
}

// PrivateKey sets field value
func (o *UCloudProvider) PrivateKey(v string) *UCloudProvider {
	o.privateKey = v
	return o
}

// GetProjectID returns the ProjectID field value
func (o *UCloudProvider) GetProjectID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.projectID
}

// GetProjectIDOk returns a tuple with the ProjectID field value
// and a boolean to check if the value has been set.
func (o *UCloudProvider) GetProjectIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.projectID, true
}

// ProjectID sets field value
func (o *UCloudProvider) ProjectID(v string) *UCloudProvider {
	o.projectID = v
	return o
}

// GetPublicKey returns the PublicKey field value
func (o *UCloudProvider) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.publicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *UCloudProvider) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.publicKey, true
}

// PublicKey sets field value
func (o *UCloudProvider) PublicKey(v string) *UCloudProvider {
	o.publicKey = v
	return o
}

// GetRegion returns the Region field value
func (o *UCloudProvider) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *UCloudProvider) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.region, true
}

// Region sets field value
func (o *UCloudProvider) Region(v string) *UCloudProvider {
	o.region = v
	return o
}

// GetType returns the Type field value
func (o *UCloudProvider) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.type_
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UCloudProvider) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.type_, true
}

// Type sets field value
func (o *UCloudProvider) Type(v string) *UCloudProvider {
	o.type_ = v
	return o
}

func (o UCloudProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UCloudProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.name
	toSerialize["privateKey"] = o.privateKey
	toSerialize["projectID"] = o.projectID
	toSerialize["publicKey"] = o.publicKey
	toSerialize["region"] = o.region
	toSerialize["type"] = o.type_
	return toSerialize, nil
}

type NullableUCloudProvider struct {
	value *UCloudProvider
	isSet bool
}

func (v NullableUCloudProvider) Get() *UCloudProvider {
	return v.value
}

func (v *NullableUCloudProvider) Set(val *UCloudProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableUCloudProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableUCloudProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUCloudProvider(val *UCloudProvider) *NullableUCloudProvider {
	return &NullableUCloudProvider{value: val, isSet: true}
}

func (v NullableUCloudProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUCloudProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}