/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cron_task

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the Exec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Exec{}

// Exec Instructions for assessing container health by executing a command. Either this attribute or the httpGet attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with both the httpGet attribute and the tcpSocket attribute.
type Exec struct {
	// A command to be executed inside the container to assess its health. Each space delimited token of the command is a separate array element. Commands exiting 0 are considered to be successful probes, whilst all other exit codes are considered failures.
	Command []string `json:"command,omitempty"`
}

// NewExecWith instantiates a new Exec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecWith() *Exec {
	this := Exec{}
	return &this
}

// NewExec instantiates a new Exec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExec() *Exec {
	this := Exec{}
	return &this
}

// NewExecs converts a list Exec pointers to objects.
// This is helpful when the SetExec requires a list of objects
func NewExecs(ps ...*Exec) []Exec {
	objs := []Exec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *Exec) GetCommand() []string {
	if o == nil || utils.IsNil(o.Command) {
		var ret []string
		return ret
	}
	return o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Exec) GetCommandOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *Exec) HasCommand() bool {
	if o != nil && !utils.IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given []string and assigns it to the command field.
// Command:  A command to be executed inside the container to assess its health. Each space delimited token of the command is a separate array element. Commands exiting 0 are considered to be successful probes, whilst all other exit codes are considered failures.
func (o *Exec) SetCommand(v []string) *Exec {
	o.Command = v
	return o
}

func (o Exec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Exec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	return toSerialize, nil
}

type NullableExec struct {
	value *Exec
	isSet bool
}

func (v NullableExec) Get() *Exec {
	return v.value
}

func (v *NullableExec) Set(val *Exec) {
	v.value = val
	v.isSet = true
}

func (v NullableExec) IsSet() bool {
	return v.isSet
}

func (v *NullableExec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExec(val *Exec) *NullableExec {
	return &NullableExec{value: val, isSet: true}
}

func (v NullableExec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
