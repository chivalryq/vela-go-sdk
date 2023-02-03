/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webservice

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the HttpGet type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HttpGet{}

// HttpGet Instructions for assessing container health by executing an HTTP GET request. Either this attribute or the exec attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with both the exec attribute and the tcpSocket attribute.
type HttpGet struct {
	Host        *string       `json:"host,omitempty"`
	HttpHeaders []HttpHeaders `json:"httpHeaders,omitempty"`
	// The endpoint, relative to the port, to which the HTTP GET request should be directed.
	Path *string `json:"path,omitempty"`
	// The TCP socket within the container to which the HTTP GET request should be directed.
	Port   *int32  `json:"port,omitempty"`
	Scheme *string `json:"scheme,omitempty"`
}

// NewHttpGetWith instantiates a new HttpGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpGetWith() *HttpGet {
	this := HttpGet{}
	var scheme string = "HTTP"
	this.Scheme = &scheme
	return &this
}

// NewHttpGet instantiates a new HttpGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpGet() *HttpGet {
	this := HttpGet{}
	var scheme string = "HTTP"
	this.Scheme = &scheme
	return &this
}

// NewHttpGets converts a list HttpGet pointers to objects.
// This is helpful when the SetHttpGet requires a list of objects
func NewHttpGets(ps ...*HttpGet) []HttpGet {
	objs := []HttpGet{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *HttpGet) GetHost() string {
	if o == nil || utils.IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpGet) GetHostOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *HttpGet) HasHost() bool {
	if o != nil && !utils.IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the host field.
// Host:
func (o *HttpGet) SetHost(v string) *HttpGet {
	o.Host = &v
	return o
}

// GetHttpHeaders returns the HttpHeaders field value if set, zero value otherwise.
func (o *HttpGet) GetHttpHeaders() []HttpHeaders {
	if o == nil || utils.IsNil(o.HttpHeaders) {
		var ret []HttpHeaders
		return ret
	}
	return o.HttpHeaders
}

// GetHttpHeadersOk returns a tuple with the HttpHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpGet) GetHttpHeadersOk() ([]HttpHeaders, bool) {
	if o == nil || utils.IsNil(o.HttpHeaders) {
		return nil, false
	}
	return o.HttpHeaders, true
}

// HasHttpHeaders returns a boolean if a field has been set.
func (o *HttpGet) HasHttpHeaders() bool {
	if o != nil && !utils.IsNil(o.HttpHeaders) {
		return true
	}

	return false
}

// SetHttpHeaders gets a reference to the given []HttpHeaders and assigns it to the httpHeaders field.
// HttpHeaders:
func (o *HttpGet) SetHttpHeaders(v []HttpHeaders) *HttpGet {
	o.HttpHeaders = v
	return o
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HttpGet) GetPath() string {
	if o == nil || utils.IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpGet) GetPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *HttpGet) HasPath() bool {
	if o != nil && !utils.IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the path field.
// Path:  The endpoint, relative to the port, to which the HTTP GET request should be directed.
func (o *HttpGet) SetPath(v string) *HttpGet {
	o.Path = &v
	return o
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *HttpGet) GetPort() int32 {
	if o == nil || utils.IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpGet) GetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *HttpGet) HasPort() bool {
	if o != nil && !utils.IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the port field.
// Port:  The TCP socket within the container to which the HTTP GET request should be directed.
func (o *HttpGet) SetPort(v int32) *HttpGet {
	o.Port = &v
	return o
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *HttpGet) GetScheme() string {
	if o == nil || utils.IsNil(o.Scheme) {
		var ret string
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpGet) GetSchemeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Scheme) {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *HttpGet) HasScheme() bool {
	if o != nil && !utils.IsNil(o.Scheme) {
		return true
	}

	return false
}

// SetScheme gets a reference to the given string and assigns it to the scheme field.
// Scheme:
func (o *HttpGet) SetScheme(v string) *HttpGet {
	o.Scheme = &v
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
	if !utils.IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !utils.IsNil(o.HttpHeaders) {
		toSerialize["httpHeaders"] = o.HttpHeaders
	}
	if !utils.IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !utils.IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !utils.IsNil(o.Scheme) {
		toSerialize["scheme"] = o.Scheme
	}
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
