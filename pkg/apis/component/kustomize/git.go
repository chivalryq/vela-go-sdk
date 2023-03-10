/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kustomize

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the Git type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Git{}

// Git struct for Git
type Git struct {
	// The Git reference to checkout and monitor for changes, defaults to master branch
	Branch *string `json:"branch,omitempty"`
	// Determines which git client library to use. Defaults to GitHub, it will pick go-git. AzureDevOps will pick libgit2.
	Provider *string `json:"provider,omitempty"`
}

// NewGitWith instantiates a new Git object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitWith() *Git {
	this := Git{}
	var provider string = "GitHub"
	this.Provider = &provider
	return &this
}

// NewGit instantiates a new Git object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGit() *Git {
	this := Git{}
	var provider string = "GitHub"
	this.Provider = &provider
	return &this
}

// NewGits converts a list Git pointers to objects.
// This is helpful when the SetGit requires a list of objects
func NewGits(ps ...*Git) []Git {
	objs := []Git{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *Git) GetBranch() string {
	if o == nil || utils.IsNil(o.Branch) {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Git) GetBranchOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Branch) {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *Git) HasBranch() bool {
	if o != nil && !utils.IsNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the branch field.
// Branch:  The Git reference to checkout and monitor for changes, defaults to master branch
func (o *Git) SetBranch(v string) *Git {
	o.Branch = &v
	return o
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *Git) GetProvider() string {
	if o == nil || utils.IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Git) GetProviderOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *Git) HasProvider() bool {
	if o != nil && !utils.IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the provider field.
// Provider:  Determines which git client library to use. Defaults to GitHub, it will pick go-git. AzureDevOps will pick libgit2.
func (o *Git) SetProvider(v string) *Git {
	o.Provider = &v
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
	if !utils.IsNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	if !utils.IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
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
