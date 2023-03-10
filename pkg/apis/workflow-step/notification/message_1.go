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

// checks if the Message1 type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Message1{}

// Message1 Specify the message that you want to sent, refer to [Lark messaging](https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN#8b0f2a1b).
type Message1 struct {
	// content should be json encode string
	Content *string `json:"content,omitempty"`
	// msg_type can be text, post, image, interactive, share_chat, share_user, audio, media, file, sticker
	MsgType *string `json:"msg_type,omitempty"`
}

// NewMessage1With instantiates a new Message1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessage1With() *Message1 {
	this := Message1{}
	return &this
}

// NewMessage1 instantiates a new Message1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessage1() *Message1 {
	this := Message1{}
	return &this
}

// NewMessage1s converts a list Message1 pointers to objects.
// This is helpful when the SetMessage1 requires a list of objects
func NewMessage1s(ps ...*Message1) []Message1 {
	objs := []Message1{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Message1) GetContent() string {
	if o == nil || utils.IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message1) GetContentOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Message1) HasContent() bool {
	if o != nil && !utils.IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the content field.
// Content:  content should be json encode string
func (o *Message1) SetContent(v string) *Message1 {
	o.Content = &v
	return o
}

// GetMsgType returns the MsgType field value if set, zero value otherwise.
func (o *Message1) GetMsgType() string {
	if o == nil || utils.IsNil(o.MsgType) {
		var ret string
		return ret
	}
	return *o.MsgType
}

// GetMsgTypeOk returns a tuple with the MsgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message1) GetMsgTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MsgType) {
		return nil, false
	}
	return o.MsgType, true
}

// HasMsgType returns a boolean if a field has been set.
func (o *Message1) HasMsgType() bool {
	if o != nil && !utils.IsNil(o.MsgType) {
		return true
	}

	return false
}

// SetMsgType gets a reference to the given string and assigns it to the msgType field.
// MsgType:  msg_type can be text, post, image, interactive, share_chat, share_user, audio, media, file, sticker
func (o *Message1) SetMsgType(v string) *Message1 {
	o.MsgType = &v
	return o
}

func (o Message1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Message1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !utils.IsNil(o.MsgType) {
		toSerialize["msg_type"] = o.MsgType
	}
	return toSerialize, nil
}

type NullableMessage1 struct {
	value *Message1
	isSet bool
}

func (v NullableMessage1) Get() *Message1 {
	return v.value
}

func (v *NullableMessage1) Set(val *Message1) {
	v.value = val
	v.isSet = true
}

func (v NullableMessage1) IsSet() bool {
	return v.isSet
}

func (v *NullableMessage1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessage1(val *Message1) *NullableMessage1 {
	return &NullableMessage1{value: val, isSet: true}
}

func (v NullableMessage1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessage1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
