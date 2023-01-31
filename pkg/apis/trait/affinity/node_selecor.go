/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affinity

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the NodeSelecor type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NodeSelecor{}

// NodeSelecor struct for NodeSelecor
type NodeSelecor struct {
	key      string   `json:"key"`
	operator string   `json:"operator"`
	values   []string `json:"values,omitempty"`
}

// NewNodeSelecorWith instantiates a new NodeSelecor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeSelecorWith(key string, operator string) *NodeSelecor {
	this := NodeSelecor{}
	this.key = key
	this.operator = operator
	return &this
}

// NewNodeSelecor instantiates a new NodeSelecor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeSelecor() *NodeSelecor {
	this := NodeSelecor{}
	var operator string = "In"
	this.operator = operator
	return &this
}

// GetKey returns the Key field value
func (o *NodeSelecor) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *NodeSelecor) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.key, true
}

// Key sets field value
func (o *NodeSelecor) Key(v string) *NodeSelecor {
	o.key = v
	return o
}

// GetOperator returns the Operator field value
func (o *NodeSelecor) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *NodeSelecor) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.operator, true
}

// Operator sets field value
func (o *NodeSelecor) Operator(v string) *NodeSelecor {
	o.operator = v
	return o
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *NodeSelecor) GetValues() []string {
	if o == nil || utils.IsNil(o.values) {
		var ret []string
		return ret
	}
	return o.values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeSelecor) GetValuesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.values) {
		return nil, false
	}
	return o.values, true
}

// HasValues returns a boolean if a field has been set.
func (o *NodeSelecor) HasValues() bool {
	if o != nil && !utils.IsNil(o.values) {
		return true
	}

	return false
}

// Values gets a reference to the given []string and assigns it to the values field.
// values:
func (o *NodeSelecor) Values(v []string) *NodeSelecor {
	o.values = v
	return o
}

func (o NodeSelecor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeSelecor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.key
	toSerialize["operator"] = o.operator
	if !utils.IsNil(o.values) {
		toSerialize["values"] = o.values
	}
	return toSerialize, nil
}

type NullableNodeSelecor struct {
	value *NodeSelecor
	isSet bool
}

func (v NullableNodeSelecor) Get() *NodeSelecor {
	return v.value
}

func (v *NullableNodeSelecor) Set(val *NodeSelecor) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeSelecor) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeSelecor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeSelecor(val *NodeSelecor) *NullableNodeSelecor {
	return &NullableNodeSelecor{value: val, isSet: true}
}

func (v NullableNodeSelecor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeSelecor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
