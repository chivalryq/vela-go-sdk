/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pure_ingress

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the PureIngressSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PureIngressSpec{}

// PureIngressSpec struct for PureIngressSpec
type PureIngressSpec struct {
	// Specify the domain you want to expose
	domain string `json:"domain"`
	// Specify the mapping relationship between the http path and the workload port
	http map[string]int32 `json:"http"`
}

// NewPureIngressSpec instantiates a new PureIngressSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPureIngressSpec(domain string, http map[string]int32) *PureIngressSpec {
	this := PureIngressSpec{}
	this.domain = domain
	this.http = http
	return &this
}

// NewPureIngressSpecWithDefaults instantiates a new PureIngressSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPureIngressSpecWithDefaults() *PureIngressSpec {
	this := PureIngressSpec{}
	return &this
}

// GetDomain returns the Domain field value
func (o *PureIngressSpec) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *PureIngressSpec) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.domain, true
}

// Domain sets field value
func (o *PureIngressSpec) Domain(v string) *PureIngressSpec {
	o.domain = v
    return o
}

// GetHttp returns the Http field value
func (o *PureIngressSpec) GetHttp() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.http
}

// GetHttpOk returns a tuple with the Http field value
// and a boolean to check if the value has been set.
func (o *PureIngressSpec) GetHttpOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.http, true
}

// Http sets field value
func (o *PureIngressSpec) Http(v map[string]int32) *PureIngressSpec {
	o.http = v
    return o
}

func (o PureIngressSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PureIngressSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.domain
	toSerialize["http"] = o.http
	return toSerialize, nil
}

type NullablePureIngressSpec struct {
	value *PureIngressSpec
	isSet bool
}

func (v NullablePureIngressSpec) Get() *PureIngressSpec {
	return v.value
}

func (v *NullablePureIngressSpec) Set(val *PureIngressSpec) {
	v.value = val
	v.isSet = true
}

func (v NullablePureIngressSpec) IsSet() bool {
	return v.isSet
}

func (v *NullablePureIngressSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePureIngressSpec(val *PureIngressSpec) *NullablePureIngressSpec {
	return &NullablePureIngressSpec{value: val, isSet: true}
}

func (v NullablePureIngressSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePureIngressSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 