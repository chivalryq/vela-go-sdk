/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package service_account

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the Privileges type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Privileges{}

// Privileges struct for Privileges
type Privileges struct {
	// Specify the apiGroups of the resource
	apiGroups []string `json:"apiGroups,omitempty"`
	// Specify the resource url to be allowed
	nonResourceURLs []string `json:"nonResourceURLs,omitempty"`
	// Specify the resourceNames to be allowed
	resourceNames []string `json:"resourceNames,omitempty"`
	// Specify the resources to be allowed
	resources []string `json:"resources,omitempty"`
	// Specify the scope of the privileges, default to be namespace scope
	scope string `json:"scope"`
	// Specify the verbs to be allowed for the resource
	verbs []string `json:"verbs"`
}

// NewPrivilegesWith instantiates a new Privileges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegesWith(scope string, verbs []string) *Privileges {
	this := Privileges{}
	this.scope = scope
	this.verbs = verbs
	return &this
}

// NewPrivileges instantiates a new Privileges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivileges() *Privileges {
	this := Privileges{}
	var scope string = "namespace"
	this.scope = scope
	return &this
}

// GetApiGroups returns the ApiGroups field value if set, zero value otherwise.
func (o *Privileges) GetApiGroups() []string {
	if o == nil || utils.IsNil(o.apiGroups) {
		var ret []string
		return ret
	}
	return o.apiGroups
}

// GetApiGroupsOk returns a tuple with the ApiGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privileges) GetApiGroupsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.apiGroups) {
		return nil, false
	}
	return o.apiGroups, true
}

// HasApiGroups returns a boolean if a field has been set.
func (o *Privileges) HasApiGroups() bool {
	if o != nil && !utils.IsNil(o.apiGroups) {
		return true
	}

	return false
}

// ApiGroups gets a reference to the given []string and assigns it to the apiGroups field.
// apiGroups:  Specify the apiGroups of the resource
func (o *Privileges) ApiGroups(v []string) *Privileges {
	o.apiGroups = v
	return o
}

// GetNonResourceURLs returns the NonResourceURLs field value if set, zero value otherwise.
func (o *Privileges) GetNonResourceURLs() []string {
	if o == nil || utils.IsNil(o.nonResourceURLs) {
		var ret []string
		return ret
	}
	return o.nonResourceURLs
}

// GetNonResourceURLsOk returns a tuple with the NonResourceURLs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privileges) GetNonResourceURLsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.nonResourceURLs) {
		return nil, false
	}
	return o.nonResourceURLs, true
}

// HasNonResourceURLs returns a boolean if a field has been set.
func (o *Privileges) HasNonResourceURLs() bool {
	if o != nil && !utils.IsNil(o.nonResourceURLs) {
		return true
	}

	return false
}

// NonResourceURLs gets a reference to the given []string and assigns it to the nonResourceURLs field.
// nonResourceURLs:  Specify the resource url to be allowed
func (o *Privileges) NonResourceURLs(v []string) *Privileges {
	o.nonResourceURLs = v
	return o
}

// GetResourceNames returns the ResourceNames field value if set, zero value otherwise.
func (o *Privileges) GetResourceNames() []string {
	if o == nil || utils.IsNil(o.resourceNames) {
		var ret []string
		return ret
	}
	return o.resourceNames
}

// GetResourceNamesOk returns a tuple with the ResourceNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privileges) GetResourceNamesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.resourceNames) {
		return nil, false
	}
	return o.resourceNames, true
}

// HasResourceNames returns a boolean if a field has been set.
func (o *Privileges) HasResourceNames() bool {
	if o != nil && !utils.IsNil(o.resourceNames) {
		return true
	}

	return false
}

// ResourceNames gets a reference to the given []string and assigns it to the resourceNames field.
// resourceNames:  Specify the resourceNames to be allowed
func (o *Privileges) ResourceNames(v []string) *Privileges {
	o.resourceNames = v
	return o
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *Privileges) GetResources() []string {
	if o == nil || utils.IsNil(o.resources) {
		var ret []string
		return ret
	}
	return o.resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privileges) GetResourcesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.resources) {
		return nil, false
	}
	return o.resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *Privileges) HasResources() bool {
	if o != nil && !utils.IsNil(o.resources) {
		return true
	}

	return false
}

// Resources gets a reference to the given []string and assigns it to the resources field.
// resources:  Specify the resources to be allowed
func (o *Privileges) Resources(v []string) *Privileges {
	o.resources = v
	return o
}

// GetScope returns the Scope field value
func (o *Privileges) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *Privileges) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.scope, true
}

// Scope sets field value
func (o *Privileges) Scope(v string) *Privileges {
	o.scope = v
	return o
}

// GetVerbs returns the Verbs field value
func (o *Privileges) GetVerbs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.verbs
}

// GetVerbsOk returns a tuple with the Verbs field value
// and a boolean to check if the value has been set.
func (o *Privileges) GetVerbsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.verbs, true
}

// Verbs sets field value
func (o *Privileges) Verbs(v []string) *Privileges {
	o.verbs = v
	return o
}

func (o Privileges) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Privileges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.apiGroups) {
		toSerialize["apiGroups"] = o.apiGroups
	}
	if !utils.IsNil(o.nonResourceURLs) {
		toSerialize["nonResourceURLs"] = o.nonResourceURLs
	}
	if !utils.IsNil(o.resourceNames) {
		toSerialize["resourceNames"] = o.resourceNames
	}
	if !utils.IsNil(o.resources) {
		toSerialize["resources"] = o.resources
	}
	toSerialize["scope"] = o.scope
	toSerialize["verbs"] = o.verbs
	return toSerialize, nil
}

type NullablePrivileges struct {
	value *Privileges
	isSet bool
}

func (v NullablePrivileges) Get() *Privileges {
	return v.value
}

func (v *NullablePrivileges) Set(val *Privileges) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivileges) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivileges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivileges(val *Privileges) *NullablePrivileges {
	return &NullablePrivileges{value: val, isSet: true}
}

func (v NullablePrivileges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivileges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
