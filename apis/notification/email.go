/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the Email type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Email{}

// Email Please fulfill its from, to and content if you want to send email
type Email struct {
	content Content `json:"content"`
	from From `json:"from"`
	// Specify the email address that you want to send to
	to []string `json:"to"`
}

// NewEmail instantiates a new Email object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmail(content Content, from From, to []string) *Email {
	this := Email{}
	this.content = content
	this.from = from
	this.to = to
	return &this
}

// NewEmailWithDefaults instantiates a new Email object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailWithDefaults() *Email {
	this := Email{}
	return &this
}

// GetContent returns the Content field value
func (o *Email) GetContent() Content {
	if o == nil {
		var ret Content
		return ret
	}

	return o.content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *Email) GetContentOk() (*Content, bool) {
	if o == nil {
		return nil, false
	}
	return &o.content, true
}

// Content sets field value
func (o *Email) Content(v Content) *Email {
	o.content = v
    return o
}

// GetFrom returns the From field value
func (o *Email) GetFrom() From {
	if o == nil {
		var ret From
		return ret
	}

	return o.from
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *Email) GetFromOk() (*From, bool) {
	if o == nil {
		return nil, false
	}
	return &o.from, true
}

// From sets field value
func (o *Email) From(v From) *Email {
	o.from = v
    return o
}

// GetTo returns the To field value
func (o *Email) GetTo() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.to
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *Email) GetToOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.to, true
}

// To sets field value
func (o *Email) To(v []string) *Email {
	o.to = v
    return o
}

func (o Email) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Email) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.content
	toSerialize["from"] = o.from
	toSerialize["to"] = o.to
	return toSerialize, nil
}

type NullableEmail struct {
	value *Email
	isSet bool
}

func (v NullableEmail) Get() *Email {
	return v.value
}

func (v *NullableEmail) Set(val *Email) {
	v.value = val
	v.isSet = true
}

func (v NullableEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmail(val *Email) *NullableEmail {
	return &NullableEmail{value: val, isSet: true}
}

func (v NullableEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 
