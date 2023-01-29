/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lifecycle

import (
	"encoding/json"

	"vela-go-sdk/pkg/apis/utils"
)

// checks if the HttpGet type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HttpGet{}

// HttpGet struct for HttpGet
type HttpGet struct {
	host        *string       `json:"host,omitempty"`
	httpHeaders []HttpHeaders `json:"httpHeaders,omitempty"`
	path        *string       `json:"path,omitempty"`
	port        int32         `json:"port"`
	scheme      string        `json:"scheme"`
}

// NewHttpGetWith instantiates a new HttpGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpGetWith(port int32, scheme string) *HttpGet {
	this := HttpGet{}
	this.port = port
	this.scheme = scheme
	return &this
}

// NewHttpGet instantiates a new HttpGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpGet() *HttpGet {
	this := HttpGet{}
	var scheme string = "HTTP"
	this.scheme = scheme
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *HttpGet) GetHost() string {
	if o == nil || utils.IsNil(o.host) {
		var ret string
		return ret
	}
	return *o.host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpGet) GetHostOk() (*string, bool) {
	if o == nil || utils.IsNil(o.host) {
		return nil, false
	}
	return o.host, true
}

// HasHost returns a boolean if a field has been set.
func (o *HttpGet) HasHost() bool {
	if o != nil && !utils.IsNil(o.host) {
		return true
	}

	return false
}

// Host gets a reference to the given string and assigns it to the host field.
// host:
func (o *HttpGet) Host(v string) *HttpGet {
	o.host = &v
	return o
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

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HttpGet) GetPath() string {
	if o == nil || utils.IsNil(o.path) {
		var ret string
		return ret
	}
	return *o.path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpGet) GetPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.path) {
		return nil, false
	}
	return o.path, true
}

// HasPath returns a boolean if a field has been set.
func (o *HttpGet) HasPath() bool {
	if o != nil && !utils.IsNil(o.path) {
		return true
	}

	return false
}

// Path gets a reference to the given string and assigns it to the path field.
// path:
func (o *HttpGet) Path(v string) *HttpGet {
	o.path = &v
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

// GetScheme returns the Scheme field value
func (o *HttpGet) GetScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.scheme
}

// GetSchemeOk returns a tuple with the Scheme field value
// and a boolean to check if the value has been set.
func (o *HttpGet) GetSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.scheme, true
}

// Scheme sets field value
func (o *HttpGet) Scheme(v string) *HttpGet {
	o.scheme = v
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
	if !utils.IsNil(o.host) {
		toSerialize["host"] = o.host
	}
	if !utils.IsNil(o.httpHeaders) {
		toSerialize["httpHeaders"] = o.httpHeaders
	}
	if !utils.IsNil(o.path) {
		toSerialize["path"] = o.path
	}
	toSerialize["port"] = o.port
	toSerialize["scheme"] = o.scheme
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
