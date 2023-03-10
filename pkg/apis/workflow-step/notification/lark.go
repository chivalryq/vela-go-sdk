/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the Lark type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Lark{}

// Lark Please fulfill its url and message if you want to send Lark messages
type Lark struct {
	Message *Message1 `json:"message,omitempty"`
	Url     *Url1     `json:"url,omitempty"`
}

// NewLarkWith instantiates a new Lark object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLarkWith() *Lark {
	this := Lark{}
	return &this
}

// NewLark instantiates a new Lark object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLark() *Lark {
	this := Lark{}
	return &this
}

// NewLarks converts a list Lark pointers to objects.
// This is helpful when the SetLark requires a list of objects
func NewLarks(ps ...*Lark) []Lark {
	objs := []Lark{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Lark) GetMessage() Message1 {
	if o == nil || utils.IsNil(o.Message) {
		var ret Message1
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lark) GetMessageOk() (*Message1, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Lark) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given Message1 and assigns it to the message field.
// Message:
func (o *Lark) SetMessage(v Message1) *Lark {
	o.Message = &v
	return o
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Lark) GetUrl() Url1 {
	if o == nil || utils.IsNil(o.Url) {
		var ret Url1
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lark) GetUrlOk() (*Url1, bool) {
	if o == nil || utils.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Lark) HasUrl() bool {
	if o != nil && !utils.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given Url1 and assigns it to the url field.
// Url:
func (o *Lark) SetUrl(v Url1) *Lark {
	o.Url = &v
	return o
}

func (o Lark) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Lark) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !utils.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableLark struct {
	value *Lark
	isSet bool
}

func (v NullableLark) Get() *Lark {
	return v.value
}

func (v *NullableLark) Set(val *Lark) {
	v.value = val
	v.isSet = true
}

func (v NullableLark) IsSet() bool {
	return v.isSet
}

func (v *NullableLark) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLark(val *Lark) *NullableLark {
	return &NullableLark{value: val, isSet: true}
}

func (v NullableLark) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLark) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
