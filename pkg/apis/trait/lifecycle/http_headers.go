/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lifecycle

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the HttpHeaders type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HttpHeaders{}

// HttpHeaders struct for HttpHeaders
type HttpHeaders struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewHttpHeadersWith instantiates a new HttpHeaders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpHeadersWith() *HttpHeaders {
	this := HttpHeaders{}
	return &this
}

// NewHttpHeaders instantiates a new HttpHeaders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpHeaders() *HttpHeaders {
	this := HttpHeaders{}
	return &this
}

// NewHttpHeaderss converts a list HttpHeaders pointers to objects.
// This is helpful when the SetHttpHeaders requires a list of objects
func NewHttpHeaderss(ps ...*HttpHeaders) []HttpHeaders {
	objs := []HttpHeaders{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HttpHeaders) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpHeaders) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HttpHeaders) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:
func (o *HttpHeaders) SetName(v string) *HttpHeaders {
	o.Name = &v
	return o
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *HttpHeaders) GetValue() string {
	if o == nil || utils.IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpHeaders) GetValueOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *HttpHeaders) HasValue() bool {
	if o != nil && !utils.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the value field.
// Value:
func (o *HttpHeaders) SetValue(v string) *HttpHeaders {
	o.Value = &v
	return o
}

func (o HttpHeaders) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpHeaders) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableHttpHeaders struct {
	value *HttpHeaders
	isSet bool
}

func (v NullableHttpHeaders) Get() *HttpHeaders {
	return v.value
}

func (v *NullableHttpHeaders) Set(val *HttpHeaders) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpHeaders) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpHeaders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpHeaders(val *HttpHeaders) *NullableHttpHeaders {
	return &NullableHttpHeaders{value: val, isSet: true}
}

func (v NullableHttpHeaders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpHeaders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
