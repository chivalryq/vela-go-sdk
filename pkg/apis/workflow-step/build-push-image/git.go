/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package build_push_image

import (
	"encoding/json"

	"vela-go-sdk/pkg/apis/utils"
)

// checks if the Git type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Git{}

// Git struct for Git
type Git struct {
	branch string `json:"branch"`
	git    string `json:"git"`
}

// NewGitWith instantiates a new Git object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitWith(branch string, git string) *Git {
	this := Git{}
	this.branch = branch
	this.git = git
	return &this
}

// NewGit instantiates a new Git object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGit() *Git {
	this := Git{}
	var branch string = "master"
	this.branch = branch
	return &this
}

// GetBranch returns the Branch field value
func (o *Git) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *Git) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.branch, true
}

// Branch sets field value
func (o *Git) Branch(v string) *Git {
	o.branch = v
	return o
}

// GetGit returns the Git field value
func (o *Git) GetGit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.git
}

// GetGitOk returns a tuple with the Git field value
// and a boolean to check if the value has been set.
func (o *Git) GetGitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.git, true
}

// Git sets field value
func (o *Git) Git(v string) *Git {
	o.git = v
	return o
}

func (o Git) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Git) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["branch"] = o.branch
	toSerialize["git"] = o.git
	return toSerialize, nil
}

type NullableGit struct {
	value *Git
	isSet bool
}

func (v NullableGit) Get() *Git {
	return v.value
}

func (v *NullableGit) Set(val *Git) {
	v.value = val
	v.isSet = true
}

func (v NullableGit) IsSet() bool {
	return v.isSet
}

func (v *NullableGit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGit(val *Git) *NullableGit {
	return &NullableGit{value: val, isSet: true}
}

func (v NullableGit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
