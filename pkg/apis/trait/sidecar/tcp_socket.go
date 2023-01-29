/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sidecar

import (
	"encoding/json"

	"vela-go-sdk/pkg/apis/utils"
)

// checks if the TcpSocket type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TcpSocket{}

// TcpSocket Instructions for assessing container health by probing a TCP socket. Either this attribute or the exec attribute or the httpGet attribute MUST be specified. This attribute is mutually exclusive with both the exec attribute and the httpGet attribute.
type TcpSocket struct {
	// The TCP socket within the container that should be probed to assess container health.
	port int32 `json:"port"`
}

// NewTcpSocketWith instantiates a new TcpSocket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTcpSocketWith(port int32) *TcpSocket {
	this := TcpSocket{}
	this.port = port
	return &this
}

// NewTcpSocket instantiates a new TcpSocket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTcpSocket() *TcpSocket {
	this := TcpSocket{}
	return &this
}

// GetPort returns the Port field value
func (o *TcpSocket) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *TcpSocket) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.port, true
}

// Port sets field value
func (o *TcpSocket) Port(v int32) *TcpSocket {
	o.port = v
	return o
}

func (o TcpSocket) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TcpSocket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["port"] = o.port
	return toSerialize, nil
}

type NullableTcpSocket struct {
	value *TcpSocket
	isSet bool
}

func (v NullableTcpSocket) Get() *TcpSocket {
	return v.value
}

func (v *NullableTcpSocket) Set(val *TcpSocket) {
	v.value = val
	v.isSet = true
}

func (v NullableTcpSocket) IsSet() bool {
	return v.isSet
}

func (v *NullableTcpSocket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTcpSocket(val *TcpSocket) *NullableTcpSocket {
	return &NullableTcpSocket{value: val, isSet: true}
}

func (v NullableTcpSocket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTcpSocket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
