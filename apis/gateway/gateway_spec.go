/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gateway

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the GatewaySpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GatewaySpec{}

// GatewaySpec struct for GatewaySpec
type GatewaySpec struct {
	// Specify the class of ingress to use
	class string `json:"class"`
	// Set ingress class in '.spec.ingressClassName' instead of 'kubernetes.io/ingress.class' annotation.
	classInSpec bool `json:"classInSpec"`
	// Specify the domain you want to expose
	domain *string `json:"domain,omitempty"`
	// Specify the host of the ingress gateway, which is used to generate the endpoints when the host is empty.
	gatewayHost *string `json:"gatewayHost,omitempty"`
	// Specify the mapping relationship between the http path and the workload port
	http map[string]int32 `json:"http"`
	// Specify the secret name you want to quote to use tls.
	secretName *string `json:"secretName,omitempty"`
}

// NewGatewaySpec instantiates a new GatewaySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewaySpec(class string, classInSpec bool, http map[string]int32) *GatewaySpec {
	this := GatewaySpec{}
	this.class = class
	this.classInSpec = classInSpec
	this.http = http
	return &this
}

// NewGatewaySpecWithDefaults instantiates a new GatewaySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewaySpecWithDefaults() *GatewaySpec {
	this := GatewaySpec{}
	var class string = "nginx"
	this.class = class
	var classInSpec bool = false
	this.classInSpec = classInSpec
	return &this
}

// GetClass returns the Class field value
func (o *GatewaySpec) GetClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.class
}

// GetClassOk returns a tuple with the Class field value
// and a boolean to check if the value has been set.
func (o *GatewaySpec) GetClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.class, true
}

// Class sets field value
func (o *GatewaySpec) Class(v string) *GatewaySpec {
	o.class = v
    return o
}

// GetClassInSpec returns the ClassInSpec field value
func (o *GatewaySpec) GetClassInSpec() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.classInSpec
}

// GetClassInSpecOk returns a tuple with the ClassInSpec field value
// and a boolean to check if the value has been set.
func (o *GatewaySpec) GetClassInSpecOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.classInSpec, true
}

// ClassInSpec sets field value
func (o *GatewaySpec) ClassInSpec(v bool) *GatewaySpec {
	o.classInSpec = v
    return o
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *GatewaySpec) GetDomain() string {
	if o == nil || utils.IsNil(o.domain) {
		var ret string
		return ret
	}
	return *o.domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySpec) GetDomainOk() (*string, bool) {
	if o == nil || utils.IsNil(o.domain) {
		return nil, false
	}
	return o.domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *GatewaySpec) HasDomain() bool {
	if o != nil && !utils.IsNil(o.domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the domain field.
// domain:  Specify the domain you want to expose 

func (o *GatewaySpec) Domain(v string) (*GatewaySpec){
	o.domain = &v
return o
}

// GetGatewayHost returns the GatewayHost field value if set, zero value otherwise.
func (o *GatewaySpec) GetGatewayHost() string {
	if o == nil || utils.IsNil(o.gatewayHost) {
		var ret string
		return ret
	}
	return *o.gatewayHost
}

// GetGatewayHostOk returns a tuple with the GatewayHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySpec) GetGatewayHostOk() (*string, bool) {
	if o == nil || utils.IsNil(o.gatewayHost) {
		return nil, false
	}
	return o.gatewayHost, true
}

// HasGatewayHost returns a boolean if a field has been set.
func (o *GatewaySpec) HasGatewayHost() bool {
	if o != nil && !utils.IsNil(o.gatewayHost) {
		return true
	}

	return false
}

// SetGatewayHost gets a reference to the given string and assigns it to the gatewayHost field.
// gatewayHost:  Specify the host of the ingress gateway, which is used to generate the endpoints when the host is empty. 

func (o *GatewaySpec) GatewayHost(v string) (*GatewaySpec){
	o.gatewayHost = &v
return o
}

// GetHttp returns the Http field value
func (o *GatewaySpec) GetHttp() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.http
}

// GetHttpOk returns a tuple with the Http field value
// and a boolean to check if the value has been set.
func (o *GatewaySpec) GetHttpOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.http, true
}

// Http sets field value
func (o *GatewaySpec) Http(v map[string]int32) *GatewaySpec {
	o.http = v
    return o
}

// GetSecretName returns the SecretName field value if set, zero value otherwise.
func (o *GatewaySpec) GetSecretName() string {
	if o == nil || utils.IsNil(o.secretName) {
		var ret string
		return ret
	}
	return *o.secretName
}

// GetSecretNameOk returns a tuple with the SecretName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySpec) GetSecretNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.secretName) {
		return nil, false
	}
	return o.secretName, true
}

// HasSecretName returns a boolean if a field has been set.
func (o *GatewaySpec) HasSecretName() bool {
	if o != nil && !utils.IsNil(o.secretName) {
		return true
	}

	return false
}

// SetSecretName gets a reference to the given string and assigns it to the secretName field.
// secretName:  Specify the secret name you want to quote to use tls. 

func (o *GatewaySpec) SecretName(v string) (*GatewaySpec){
	o.secretName = &v
return o
}

func (o GatewaySpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewaySpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["class"] = o.class
	toSerialize["classInSpec"] = o.classInSpec
	if !utils.IsNil(o.domain) {
		toSerialize["domain"] = o.domain
	}
	if !utils.IsNil(o.gatewayHost) {
		toSerialize["gatewayHost"] = o.gatewayHost
	}
	toSerialize["http"] = o.http
	if !utils.IsNil(o.secretName) {
		toSerialize["secretName"] = o.secretName
	}
	return toSerialize, nil
}

type NullableGatewaySpec struct {
	value *GatewaySpec
	isSet bool
}

func (v NullableGatewaySpec) Get() *GatewaySpec {
	return v.value
}

func (v *NullableGatewaySpec) Set(val *GatewaySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewaySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewaySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewaySpec(val *GatewaySpec) *NullableGatewaySpec {
	return &NullableGatewaySpec{value: val, isSet: true}
}

func (v NullableGatewaySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewaySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 