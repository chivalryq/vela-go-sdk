/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hpa

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the PodCustomMetrics type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PodCustomMetrics{}

// PodCustomMetrics struct for PodCustomMetrics
type PodCustomMetrics struct {
	// Specify name of custom metrics
	Name *string `json:"name,omitempty"`
	// Specify target value of custom metrics
	Value *string `json:"value,omitempty"`
}

// NewPodCustomMetricsWith instantiates a new PodCustomMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPodCustomMetricsWith() *PodCustomMetrics {
	this := PodCustomMetrics{}
	return &this
}

// NewPodCustomMetrics instantiates a new PodCustomMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPodCustomMetrics() *PodCustomMetrics {
	this := PodCustomMetrics{}
	return &this
}

// NewPodCustomMetricss converts a list PodCustomMetrics pointers to objects.
// This is helpful when the SetPodCustomMetrics requires a list of objects
func NewPodCustomMetricss(ps ...*PodCustomMetrics) []PodCustomMetrics {
	objs := []PodCustomMetrics{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PodCustomMetrics) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodCustomMetrics) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PodCustomMetrics) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:  Specify name of custom metrics
func (o *PodCustomMetrics) SetName(v string) *PodCustomMetrics {
	o.Name = &v
	return o
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PodCustomMetrics) GetValue() string {
	if o == nil || utils.IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodCustomMetrics) GetValueOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PodCustomMetrics) HasValue() bool {
	if o != nil && !utils.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the value field.
// Value:  Specify target value of custom metrics
func (o *PodCustomMetrics) SetValue(v string) *PodCustomMetrics {
	o.Value = &v
	return o
}

func (o PodCustomMetrics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PodCustomMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullablePodCustomMetrics struct {
	value *PodCustomMetrics
	isSet bool
}

func (v NullablePodCustomMetrics) Get() *PodCustomMetrics {
	return v.value
}

func (v *NullablePodCustomMetrics) Set(val *PodCustomMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullablePodCustomMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullablePodCustomMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePodCustomMetrics(val *PodCustomMetrics) *NullablePodCustomMetrics {
	return &NullablePodCustomMetrics{value: val, isSet: true}
}

func (v NullablePodCustomMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePodCustomMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
