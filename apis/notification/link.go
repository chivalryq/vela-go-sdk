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



    
    // checks if the Link type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Link{}

// Link struct for Link
type Link struct {
	messageUrl *string `json:"messageUrl,omitempty"`
	picUrl *string `json:"picUrl,omitempty"`
	text *string `json:"text,omitempty"`
	title *string `json:"title,omitempty"`
}

// NewLink instantiates a new Link object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLink() *Link {
	this := Link{}
	return &this
}

// NewLinkWithDefaults instantiates a new Link object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkWithDefaults() *Link {
	this := Link{}
	return &this
}

// GetMessageUrl returns the MessageUrl field value if set, zero value otherwise.
func (o *Link) GetMessageUrl() string {
	if o == nil || utils.IsNil(o.messageUrl) {
		var ret string
		return ret
	}
	return *o.messageUrl
}

// GetMessageUrlOk returns a tuple with the MessageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetMessageUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.messageUrl) {
		return nil, false
	}
	return o.messageUrl, true
}

// HasMessageUrl returns a boolean if a field has been set.
func (o *Link) HasMessageUrl() bool {
	if o != nil && !utils.IsNil(o.messageUrl) {
		return true
	}

	return false
}

// SetMessageUrl gets a reference to the given string and assigns it to the messageUrl field.
// messageUrl: 

func (o *Link) MessageUrl(v string) (*Link){
	o.messageUrl = &v
return o
}

// GetPicUrl returns the PicUrl field value if set, zero value otherwise.
func (o *Link) GetPicUrl() string {
	if o == nil || utils.IsNil(o.picUrl) {
		var ret string
		return ret
	}
	return *o.picUrl
}

// GetPicUrlOk returns a tuple with the PicUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetPicUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.picUrl) {
		return nil, false
	}
	return o.picUrl, true
}

// HasPicUrl returns a boolean if a field has been set.
func (o *Link) HasPicUrl() bool {
	if o != nil && !utils.IsNil(o.picUrl) {
		return true
	}

	return false
}

// SetPicUrl gets a reference to the given string and assigns it to the picUrl field.
// picUrl: 

func (o *Link) PicUrl(v string) (*Link){
	o.picUrl = &v
return o
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Link) GetText() string {
	if o == nil || utils.IsNil(o.text) {
		var ret string
		return ret
	}
	return *o.text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetTextOk() (*string, bool) {
	if o == nil || utils.IsNil(o.text) {
		return nil, false
	}
	return o.text, true
}

// HasText returns a boolean if a field has been set.
func (o *Link) HasText() bool {
	if o != nil && !utils.IsNil(o.text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the text field.
// text: 

func (o *Link) Text(v string) (*Link){
	o.text = &v
return o
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Link) GetTitle() string {
	if o == nil || utils.IsNil(o.title) {
		var ret string
		return ret
	}
	return *o.title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.title) {
		return nil, false
	}
	return o.title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Link) HasTitle() bool {
	if o != nil && !utils.IsNil(o.title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the title field.
// title: 

func (o *Link) Title(v string) (*Link){
	o.title = &v
return o
}

func (o Link) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Link) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.messageUrl) {
		toSerialize["messageUrl"] = o.messageUrl
	}
	if !utils.IsNil(o.picUrl) {
		toSerialize["picUrl"] = o.picUrl
	}
	if !utils.IsNil(o.text) {
		toSerialize["text"] = o.text
	}
	if !utils.IsNil(o.title) {
		toSerialize["title"] = o.title
	}
	return toSerialize, nil
}

type NullableLink struct {
	value *Link
	isSet bool
}

func (v NullableLink) Get() *Link {
	return v.value
}

func (v *NullableLink) Set(val *Link) {
	v.value = val
	v.isSet = true
}

func (v NullableLink) IsSet() bool {
	return v.isSet
}

func (v *NullableLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLink(val *Link) *NullableLink {
	return &NullableLink{value: val, isSet: true}
}

func (v NullableLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 