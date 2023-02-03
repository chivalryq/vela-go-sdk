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

// checks if the Markdown type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Markdown{}

// Markdown struct for Markdown
type Markdown struct {
	Text  *string `json:"text,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewMarkdownWith instantiates a new Markdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkdownWith() *Markdown {
	this := Markdown{}
	return &this
}

// NewMarkdown instantiates a new Markdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkdown() *Markdown {
	this := Markdown{}
	return &this
}

// NewMarkdowns converts a list Markdown pointers to objects.
// This is helpful when the SetMarkdown requires a list of objects
func NewMarkdowns(ps ...*Markdown) []Markdown {
	objs := []Markdown{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Markdown) GetText() string {
	if o == nil || utils.IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Markdown) GetTextOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Markdown) HasText() bool {
	if o != nil && !utils.IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the text field.
// Text:
func (o *Markdown) SetText(v string) *Markdown {
	o.Text = &v
	return o
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Markdown) GetTitle() string {
	if o == nil || utils.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Markdown) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Markdown) HasTitle() bool {
	if o != nil && !utils.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the title field.
// Title:
func (o *Markdown) SetTitle(v string) *Markdown {
	o.Title = &v
	return o
}

func (o Markdown) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Markdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !utils.IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableMarkdown struct {
	value *Markdown
	isSet bool
}

func (v NullableMarkdown) Get() *Markdown {
	return v.value
}

func (v *NullableMarkdown) Set(val *Markdown) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkdown) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkdown(val *Markdown) *NullableMarkdown {
	return &NullableMarkdown{value: val, isSet: true}
}

func (v NullableMarkdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
