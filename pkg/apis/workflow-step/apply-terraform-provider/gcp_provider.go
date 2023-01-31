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

// checks if the GCPProvider type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GCPProvider{}

// GCPProvider struct for GCPProvider
type GCPProvider struct {
	credentials string `json:"credentials"`
	name        string `json:"name"`
	project     string `json:"project"`
	region      string `json:"region"`
	type_       string `json:"type"`
}

// NewGCPProviderWith instantiates a new GCPProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCPProviderWith(credentials string, name string, project string, region string, type_ string) *GCPProvider {
	this := GCPProvider{}
	this.credentials = credentials
	this.name = name
	this.project = project
	this.region = region
	this.type_ = type_
	return &this
}

// NewGCPProvider instantiates a new GCPProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCPProvider() *GCPProvider {
	this := GCPProvider{}
	var name string = "gcp-provider"
	this.name = name
	return &this
}

// GetCredentials returns the Credentials field value
func (o *GCPProvider) GetCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *GCPProvider) GetCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.credentials, true
}

// Credentials sets field value
func (o *GCPProvider) Credentials(v string) *GCPProvider {
	o.credentials = v
	return o
}

// GetName returns the Name field value
func (o *GCPProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GCPProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *GCPProvider) Name(v string) *GCPProvider {
	o.name = v
	return o
}

// GetProject returns the Project field value
func (o *GCPProvider) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *GCPProvider) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.project, true
}

// Project sets field value
func (o *GCPProvider) Project(v string) *GCPProvider {
	o.project = v
	return o
}

// GetRegion returns the Region field value
func (o *GCPProvider) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *GCPProvider) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.region, true
}

// Region sets field value
func (o *GCPProvider) Region(v string) *GCPProvider {
	o.region = v
	return o
}

// GetType returns the Type field value
func (o *GCPProvider) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.type_
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GCPProvider) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.type_, true
}

// Type sets field value
func (o *GCPProvider) Type(v string) *GCPProvider {
	o.type_ = v
	return o
}

func (o GCPProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GCPProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["credentials"] = o.credentials
	toSerialize["name"] = o.name
	toSerialize["project"] = o.project
	toSerialize["region"] = o.region
	toSerialize["type"] = o.type_
	return toSerialize, nil
}

type NullableGCPProvider struct {
	value *GCPProvider
	isSet bool
}

func (v NullableGCPProvider) Get() *GCPProvider {
	return v.value
}

func (v *NullableGCPProvider) Set(val *GCPProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableGCPProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableGCPProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCPProvider(val *GCPProvider) *NullableGCPProvider {
	return &NullableGCPProvider{value: val, isSet: true}
}

func (v NullableGCPProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCPProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
