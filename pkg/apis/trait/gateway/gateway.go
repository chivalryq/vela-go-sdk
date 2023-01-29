/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gateway

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"vela-go-sdk/pkg/apis"
	"vela-go-sdk/pkg/apis/utils"
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

// NewGatewaySpecWith instantiates a new GatewaySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewaySpecWith(class string, classInSpec bool, http map[string]int32) *GatewaySpec {
	this := GatewaySpec{}
	this.class = class
	this.classInSpec = classInSpec
	this.http = http
	return &this
}

// NewGatewaySpec instantiates a new GatewaySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewaySpec() *GatewaySpec {
	this := GatewaySpec{}
	var class string = "nginx"
	this.class = class
	var classInSpec bool = false
	this.classInSpec = classInSpec
	return &this
}

// GetClass returns the Class field value
func (o *GatewayTrait) GetClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.class
}

// GetClassOk returns a tuple with the Class field value
// and a boolean to check if the value has been set.
func (o *GatewayTrait) GetClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.class, true
}

// Class sets field value
func (o *GatewayTrait) Class(v string) *GatewayTrait {
	o.Properties.class = v
	return o
}

// GetClassInSpec returns the ClassInSpec field value
func (o *GatewayTrait) GetClassInSpec() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Properties.classInSpec
}

// GetClassInSpecOk returns a tuple with the ClassInSpec field value
// and a boolean to check if the value has been set.
func (o *GatewayTrait) GetClassInSpecOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.classInSpec, true
}

// ClassInSpec sets field value
func (o *GatewayTrait) ClassInSpec(v bool) *GatewayTrait {
	o.Properties.classInSpec = v
	return o
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *GatewayTrait) GetDomain() string {
	if o == nil || utils.IsNil(o.Properties.domain) {
		var ret string
		return ret
	}
	return *o.Properties.domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTrait) GetDomainOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.domain) {
		return nil, false
	}
	return o.Properties.domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *GatewayTrait) HasDomain() bool {
	if o != nil && !utils.IsNil(o.Properties.domain) {
		return true
	}

	return false
}

// Domain gets a reference to the given string and assigns it to the domain field.
// domain:  Specify the domain you want to expose
func (o *GatewayTrait) Domain(v string) *GatewayTrait {
	o.Properties.domain = &v
	return o
}

// GetGatewayHost returns the GatewayHost field value if set, zero value otherwise.
func (o *GatewayTrait) GetGatewayHost() string {
	if o == nil || utils.IsNil(o.Properties.gatewayHost) {
		var ret string
		return ret
	}
	return *o.Properties.gatewayHost
}

// GetGatewayHostOk returns a tuple with the GatewayHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTrait) GetGatewayHostOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.gatewayHost) {
		return nil, false
	}
	return o.Properties.gatewayHost, true
}

// HasGatewayHost returns a boolean if a field has been set.
func (o *GatewayTrait) HasGatewayHost() bool {
	if o != nil && !utils.IsNil(o.Properties.gatewayHost) {
		return true
	}

	return false
}

// GatewayHost gets a reference to the given string and assigns it to the gatewayHost field.
// gatewayHost:  Specify the host of the ingress gateway, which is used to generate the endpoints when the host is empty.
func (o *GatewayTrait) GatewayHost(v string) *GatewayTrait {
	o.Properties.gatewayHost = &v
	return o
}

// GetHttp returns the Http field value
func (o *GatewayTrait) GetHttp() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.Properties.http
}

// GetHttpOk returns a tuple with the Http field value
// and a boolean to check if the value has been set.
func (o *GatewayTrait) GetHttpOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.http, true
}

// Http sets field value
func (o *GatewayTrait) Http(v map[string]int32) *GatewayTrait {
	o.Properties.http = v
	return o
}

// GetSecretName returns the SecretName field value if set, zero value otherwise.
func (o *GatewayTrait) GetSecretName() string {
	if o == nil || utils.IsNil(o.Properties.secretName) {
		var ret string
		return ret
	}
	return *o.Properties.secretName
}

// GetSecretNameOk returns a tuple with the SecretName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTrait) GetSecretNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.secretName) {
		return nil, false
	}
	return o.Properties.secretName, true
}

// HasSecretName returns a boolean if a field has been set.
func (o *GatewayTrait) HasSecretName() bool {
	if o != nil && !utils.IsNil(o.Properties.secretName) {
		return true
	}

	return false
}

// SecretName gets a reference to the given string and assigns it to the secretName field.
// secretName:  Specify the secret name you want to quote to use tls.
func (o *GatewayTrait) SecretName(v string) *GatewayTrait {
	o.Properties.secretName = &v
	return o
}

func (o GatewaySpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
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

const GatewayType = "gateway"

type GatewayTrait struct {
	Base       apis.TraitBase
	Properties GatewaySpec
}

func Gateway() *GatewayTrait {
	g := &GatewayTrait{Base: apis.TraitBase{}}
	return g
}

func (g *GatewayTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(g.Properties),
		Type:       GatewayType,
	}
	return res
}

func (g *GatewayTrait) Type() string {
	return GatewayType
}
