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

// checks if the AzureProvider type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AzureProvider{}

// AzureProvider struct for AzureProvider
type AzureProvider struct {
	clientID       string `json:"clientID"`
	clientSecret   string `json:"clientSecret"`
	name           string `json:"name"`
	subscriptionID string `json:"subscriptionID"`
	tenantID       string `json:"tenantID"`
}

// NewAzureProviderWith instantiates a new AzureProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureProviderWith(clientID string, clientSecret string, name string, subscriptionID string, tenantID string) *AzureProvider {
	this := AzureProvider{}
	this.clientID = clientID
	this.clientSecret = clientSecret
	this.name = name
	this.subscriptionID = subscriptionID
	this.tenantID = tenantID
	return &this
}

// NewAzureProvider instantiates a new AzureProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureProvider() *AzureProvider {
	this := AzureProvider{}
	var name string = "azure-provider"
	this.name = name
	return &this
}

// GetClientID returns the ClientID field value
func (o *AzureProvider) GetClientID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.clientID
}

// GetClientIDOk returns a tuple with the ClientID field value
// and a boolean to check if the value has been set.
func (o *AzureProvider) GetClientIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.clientID, true
}

// ClientID sets field value
func (o *AzureProvider) ClientID(v string) *AzureProvider {
	o.clientID = v
	return o
}

// GetClientSecret returns the ClientSecret field value
func (o *AzureProvider) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.clientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *AzureProvider) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.clientSecret, true
}

// ClientSecret sets field value
func (o *AzureProvider) ClientSecret(v string) *AzureProvider {
	o.clientSecret = v
	return o
}

// GetName returns the Name field value
func (o *AzureProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AzureProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *AzureProvider) Name(v string) *AzureProvider {
	o.name = v
	return o
}

// GetSubscriptionID returns the SubscriptionID field value
func (o *AzureProvider) GetSubscriptionID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.subscriptionID
}

// GetSubscriptionIDOk returns a tuple with the SubscriptionID field value
// and a boolean to check if the value has been set.
func (o *AzureProvider) GetSubscriptionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.subscriptionID, true
}

// SubscriptionID sets field value
func (o *AzureProvider) SubscriptionID(v string) *AzureProvider {
	o.subscriptionID = v
	return o
}

// GetTenantID returns the TenantID field value
func (o *AzureProvider) GetTenantID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.tenantID
}

// GetTenantIDOk returns a tuple with the TenantID field value
// and a boolean to check if the value has been set.
func (o *AzureProvider) GetTenantIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.tenantID, true
}

// TenantID sets field value
func (o *AzureProvider) TenantID(v string) *AzureProvider {
	o.tenantID = v
	return o
}

func (o AzureProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientID"] = o.clientID
	toSerialize["clientSecret"] = o.clientSecret
	toSerialize["name"] = o.name
	toSerialize["subscriptionID"] = o.subscriptionID
	toSerialize["tenantID"] = o.tenantID
	return toSerialize, nil
}

type NullableAzureProvider struct {
	value *AzureProvider
	isSet bool
}

func (v NullableAzureProvider) Get() *AzureProvider {
	return v.value
}

func (v *NullableAzureProvider) Set(val *AzureProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureProvider(val *AzureProvider) *NullableAzureProvider {
	return &NullableAzureProvider{value: val, isSet: true}
}

func (v NullableAzureProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
