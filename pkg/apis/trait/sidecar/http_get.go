/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sidecar

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the HttpGet type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HttpGet{}

// HttpGet Instructions for assessing container health by executing an HTTP GET request. Either this attribute or the exec attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with both the exec attribute and the tcpSocket attribute.
type HttpGet struct {
	httpHeaders []HttpHeaders `json:"httpHeaders,omitempty"`
	// The endpoint, relative to the port, to which the HTTP GET request should be directed.
	path string `json:"path"`
	// The TCP socket within the container to which the HTTP GET request should be directed.
	port int32 `json:"port"`
}

// NewHttpGetWith instantiates a new HttpGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpGetWith(path string, port int32) *HttpGet {
	this := HttpGet{}
	this.path = path
	this.port = port
	return &this
}

// NewHttpGet instantiates a new HttpGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpGet() *HttpGet {
	this := HttpGet{}
	return &this
}

// GetHttpHeaders returns the HttpHeaders field value if set, zero value otherwise.
func (o *HttpGet) GetHttpHeaders() []HttpHeaders {
	if o == nil || utils.IsNil(o.httpHeaders) {
		var ret []HttpHeaders
		return ret
	}
	return o.httpHeaders
}

// GetHttpHeadersOk returns a tuple with the HttpHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpGet) GetHttpHeadersOk() ([]HttpHeaders, bool) {
	if o == nil || utils.IsNil(o.httpHeaders) {
		return nil, false
	}
	return o.httpHeaders, true
}

// HasHttpHeaders returns a boolean if a field has been set.
func (o *HttpGet) HasHttpHeaders() bool {
	if o != nil && !utils.IsNil(o.httpHeaders) {
		return true
	}

	return false
}

// HttpHeaders gets a reference to the given []HttpHeaders and assigns it to the httpHeaders field.
// httpHeaders:
func (o *HttpGet) HttpHeaders(v []HttpHeaders) *HttpGet {
	o.httpHeaders = v
	return o
}

// GetPath returns the Path field value
func (o *HttpGet) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *HttpGet) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.path, true
}

// Path sets field value
func (o *HttpGet) Path(v string) *HttpGet {
	o.path = v
	return o
}

// GetPort returns the Port field value
func (o *HttpGet) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *HttpGet) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.port, true
}

// Port sets field value
func (o *HttpGet) Port(v int32) *HttpGet {
	o.port = v
	return o
}

func (o HttpGet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpGet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.httpHeaders) {
		toSerialize["httpHeaders"] = o.httpHeaders
	}
	toSerialize["path"] = o.path
	toSerialize["port"] = o.port
	return toSerialize, nil
}

type NullableHttpGet struct {
	value *HttpGet
	isSet bool
}

func (v NullableHttpGet) Get() *HttpGet {
	return v.value
}

func (v *NullableHttpGet) Set(val *HttpGet) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpGet) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpGet(val *HttpGet) *NullableHttpGet {
	return &NullableHttpGet{value: val, isSet: true}
}

func (v NullableHttpGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
