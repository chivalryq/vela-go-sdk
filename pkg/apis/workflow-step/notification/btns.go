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

// checks if the Btns type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Btns{}

// Btns struct for Btns
type Btns struct {
	ActionURL *string `json:"actionURL,omitempty"`
	Title     *string `json:"title,omitempty"`
}

// NewBtnsWith instantiates a new Btns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBtnsWith() *Btns {
	this := Btns{}
	return &this
}

// NewBtns instantiates a new Btns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBtns() *Btns {
	this := Btns{}
	return &this
}

// NewBtnss converts a list Btns pointers to objects.
// This is helpful when the SetBtns requires a list of objects
func NewBtnss(ps ...*Btns) []Btns {
	objs := []Btns{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetActionURL returns the ActionURL field value if set, zero value otherwise.
func (o *Btns) GetActionURL() string {
	if o == nil || utils.IsNil(o.ActionURL) {
		var ret string
		return ret
	}
	return *o.ActionURL
}

// GetActionURLOk returns a tuple with the ActionURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Btns) GetActionURLOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ActionURL) {
		return nil, false
	}
	return o.ActionURL, true
}

// HasActionURL returns a boolean if a field has been set.
func (o *Btns) HasActionURL() bool {
	if o != nil && !utils.IsNil(o.ActionURL) {
		return true
	}

	return false
}

// SetActionURL gets a reference to the given string and assigns it to the actionURL field.
// ActionURL:
func (o *Btns) SetActionURL(v string) *Btns {
	o.ActionURL = &v
	return o
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Btns) GetTitle() string {
	if o == nil || utils.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Btns) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Btns) HasTitle() bool {
	if o != nil && !utils.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the title field.
// Title:
func (o *Btns) SetTitle(v string) *Btns {
	o.Title = &v
	return o
}

func (o Btns) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Btns) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ActionURL) {
		toSerialize["actionURL"] = o.ActionURL
	}
	if !utils.IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableBtns struct {
	value *Btns
	isSet bool
}

func (v NullableBtns) Get() *Btns {
	return v.value
}

func (v *NullableBtns) Set(val *Btns) {
	v.value = val
	v.isSet = true
}

func (v NullableBtns) IsSet() bool {
	return v.isSet
}

func (v *NullableBtns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBtns(val *Btns) *NullableBtns {
	return &NullableBtns{value: val, isSet: true}
}

func (v NullableBtns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBtns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
