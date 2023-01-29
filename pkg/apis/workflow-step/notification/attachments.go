/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"

	"vela-go-sdk/pkg/apis/utils"
)

// checks if the Attachments type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Attachments{}

// Attachments struct for Attachments
type Attachments struct {
	blocks utils.NullableString `json:"blocks,omitempty"`
	color  *string              `json:"color,omitempty"`
}

// NewAttachmentsWith instantiates a new Attachments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentsWith() *Attachments {
	this := Attachments{}
	return &this
}

// NewAttachments instantiates a new Attachments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachments() *Attachments {
	this := Attachments{}
	return &this
}

// GetBlocks returns the Blocks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attachments) GetBlocks() string {
	if o == nil || utils.IsNil(o.blocks.Get()) {
		var ret string
		return ret
	}
	return *o.blocks.Get()
}

// GetBlocksOk returns a tuple with the Blocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attachments) GetBlocksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.blocks.Get(), o.blocks.IsSet()
}

// HasBlocks returns a boolean if a field has been set.
func (o *Attachments) HasBlocks() bool {
	if o != nil && o.blocks.IsSet() {
		return true
	}

	return false
}

// Blocks gets a reference to the given utils.NullableString and assigns it to the blocks field.
// blocks:
func (o *Attachments) Blocks(v string) *Attachments {
	o.blocks.Set(&v)
	return o
}

// SetBlocksNil sets the value for Blocks to be an explicit nil
func (o *Attachments) SetBlocksNil() {
	o.blocks.Set(nil)
}

// UnsetBlocks ensures that no value is present for Blocks, not even an explicit nil
func (o *Attachments) UnsetBlocks() {
	o.blocks.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *Attachments) GetColor() string {
	if o == nil || utils.IsNil(o.color) {
		var ret string
		return ret
	}
	return *o.color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attachments) GetColorOk() (*string, bool) {
	if o == nil || utils.IsNil(o.color) {
		return nil, false
	}
	return o.color, true
}

// HasColor returns a boolean if a field has been set.
func (o *Attachments) HasColor() bool {
	if o != nil && !utils.IsNil(o.color) {
		return true
	}

	return false
}

// Color gets a reference to the given string and assigns it to the color field.
// color:
func (o *Attachments) Color(v string) *Attachments {
	o.color = &v
	return o
}

func (o Attachments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Attachments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.blocks.IsSet() {
		toSerialize["blocks"] = o.blocks.Get()
	}
	if !utils.IsNil(o.color) {
		toSerialize["color"] = o.color
	}
	return toSerialize, nil
}

type NullableAttachments struct {
	value *Attachments
	isSet bool
}

func (v NullableAttachments) Get() *Attachments {
	return v.value
}

func (v *NullableAttachments) Set(val *Attachments) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachments) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachments(val *Attachments) *NullableAttachments {
	return &NullableAttachments{value: val, isSet: true}
}

func (v NullableAttachments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
