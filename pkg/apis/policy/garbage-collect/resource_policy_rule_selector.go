/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package garbage_collect

import (
	"encoding/json"

	"vela-go-sdk/pkg/apis/utils"
)

// checks if the ResourcePolicyRuleSelector type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ResourcePolicyRuleSelector{}

// ResourcePolicyRuleSelector struct for ResourcePolicyRuleSelector
type ResourcePolicyRuleSelector struct {
	// Select resources by component names
	componentNames []string `json:"componentNames,omitempty"`
	// Select resources by component types
	componentTypes []string `json:"componentTypes,omitempty"`
	// Select resources by oamTypes (COMPONENT or TRAIT)
	oamTypes []string `json:"oamTypes,omitempty"`
	// Select resources by their names
	resourceNames []string `json:"resourceNames,omitempty"`
	// Select resources by resource types (like Deployment)
	resourceTypes []string `json:"resourceTypes,omitempty"`
	// Select resources by trait types
	traitTypes []string `json:"traitTypes,omitempty"`
}

// NewResourcePolicyRuleSelectorWith instantiates a new ResourcePolicyRuleSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcePolicyRuleSelectorWith() *ResourcePolicyRuleSelector {
	this := ResourcePolicyRuleSelector{}
	return &this
}

// NewResourcePolicyRuleSelector instantiates a new ResourcePolicyRuleSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcePolicyRuleSelector() *ResourcePolicyRuleSelector {
	this := ResourcePolicyRuleSelector{}
	return &this
}

// GetComponentNames returns the ComponentNames field value if set, zero value otherwise.
func (o *ResourcePolicyRuleSelector) GetComponentNames() []string {
	if o == nil || utils.IsNil(o.componentNames) {
		var ret []string
		return ret
	}
	return o.componentNames
}

// GetComponentNamesOk returns a tuple with the ComponentNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePolicyRuleSelector) GetComponentNamesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.componentNames) {
		return nil, false
	}
	return o.componentNames, true
}

// HasComponentNames returns a boolean if a field has been set.
func (o *ResourcePolicyRuleSelector) HasComponentNames() bool {
	if o != nil && !utils.IsNil(o.componentNames) {
		return true
	}

	return false
}

// ComponentNames gets a reference to the given []string and assigns it to the componentNames field.
// componentNames:  Select resources by component names
func (o *ResourcePolicyRuleSelector) ComponentNames(v []string) *ResourcePolicyRuleSelector {
	o.componentNames = v
	return o
}

// GetComponentTypes returns the ComponentTypes field value if set, zero value otherwise.
func (o *ResourcePolicyRuleSelector) GetComponentTypes() []string {
	if o == nil || utils.IsNil(o.componentTypes) {
		var ret []string
		return ret
	}
	return o.componentTypes
}

// GetComponentTypesOk returns a tuple with the ComponentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePolicyRuleSelector) GetComponentTypesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.componentTypes) {
		return nil, false
	}
	return o.componentTypes, true
}

// HasComponentTypes returns a boolean if a field has been set.
func (o *ResourcePolicyRuleSelector) HasComponentTypes() bool {
	if o != nil && !utils.IsNil(o.componentTypes) {
		return true
	}

	return false
}

// ComponentTypes gets a reference to the given []string and assigns it to the componentTypes field.
// componentTypes:  Select resources by component types
func (o *ResourcePolicyRuleSelector) ComponentTypes(v []string) *ResourcePolicyRuleSelector {
	o.componentTypes = v
	return o
}

// GetOamTypes returns the OamTypes field value if set, zero value otherwise.
func (o *ResourcePolicyRuleSelector) GetOamTypes() []string {
	if o == nil || utils.IsNil(o.oamTypes) {
		var ret []string
		return ret
	}
	return o.oamTypes
}

// GetOamTypesOk returns a tuple with the OamTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePolicyRuleSelector) GetOamTypesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.oamTypes) {
		return nil, false
	}
	return o.oamTypes, true
}

// HasOamTypes returns a boolean if a field has been set.
func (o *ResourcePolicyRuleSelector) HasOamTypes() bool {
	if o != nil && !utils.IsNil(o.oamTypes) {
		return true
	}

	return false
}

// OamTypes gets a reference to the given []string and assigns it to the oamTypes field.
// oamTypes:  Select resources by oamTypes (COMPONENT or TRAIT)
func (o *ResourcePolicyRuleSelector) OamTypes(v []string) *ResourcePolicyRuleSelector {
	o.oamTypes = v
	return o
}

// GetResourceNames returns the ResourceNames field value if set, zero value otherwise.
func (o *ResourcePolicyRuleSelector) GetResourceNames() []string {
	if o == nil || utils.IsNil(o.resourceNames) {
		var ret []string
		return ret
	}
	return o.resourceNames
}

// GetResourceNamesOk returns a tuple with the ResourceNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePolicyRuleSelector) GetResourceNamesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.resourceNames) {
		return nil, false
	}
	return o.resourceNames, true
}

// HasResourceNames returns a boolean if a field has been set.
func (o *ResourcePolicyRuleSelector) HasResourceNames() bool {
	if o != nil && !utils.IsNil(o.resourceNames) {
		return true
	}

	return false
}

// ResourceNames gets a reference to the given []string and assigns it to the resourceNames field.
// resourceNames:  Select resources by their names
func (o *ResourcePolicyRuleSelector) ResourceNames(v []string) *ResourcePolicyRuleSelector {
	o.resourceNames = v
	return o
}

// GetResourceTypes returns the ResourceTypes field value if set, zero value otherwise.
func (o *ResourcePolicyRuleSelector) GetResourceTypes() []string {
	if o == nil || utils.IsNil(o.resourceTypes) {
		var ret []string
		return ret
	}
	return o.resourceTypes
}

// GetResourceTypesOk returns a tuple with the ResourceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePolicyRuleSelector) GetResourceTypesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.resourceTypes) {
		return nil, false
	}
	return o.resourceTypes, true
}

// HasResourceTypes returns a boolean if a field has been set.
func (o *ResourcePolicyRuleSelector) HasResourceTypes() bool {
	if o != nil && !utils.IsNil(o.resourceTypes) {
		return true
	}

	return false
}

// ResourceTypes gets a reference to the given []string and assigns it to the resourceTypes field.
// resourceTypes:  Select resources by resource types (like Deployment)
func (o *ResourcePolicyRuleSelector) ResourceTypes(v []string) *ResourcePolicyRuleSelector {
	o.resourceTypes = v
	return o
}

// GetTraitTypes returns the TraitTypes field value if set, zero value otherwise.
func (o *ResourcePolicyRuleSelector) GetTraitTypes() []string {
	if o == nil || utils.IsNil(o.traitTypes) {
		var ret []string
		return ret
	}
	return o.traitTypes
}

// GetTraitTypesOk returns a tuple with the TraitTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePolicyRuleSelector) GetTraitTypesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.traitTypes) {
		return nil, false
	}
	return o.traitTypes, true
}

// HasTraitTypes returns a boolean if a field has been set.
func (o *ResourcePolicyRuleSelector) HasTraitTypes() bool {
	if o != nil && !utils.IsNil(o.traitTypes) {
		return true
	}

	return false
}

// TraitTypes gets a reference to the given []string and assigns it to the traitTypes field.
// traitTypes:  Select resources by trait types
func (o *ResourcePolicyRuleSelector) TraitTypes(v []string) *ResourcePolicyRuleSelector {
	o.traitTypes = v
	return o
}

func (o ResourcePolicyRuleSelector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourcePolicyRuleSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.componentNames) {
		toSerialize["componentNames"] = o.componentNames
	}
	if !utils.IsNil(o.componentTypes) {
		toSerialize["componentTypes"] = o.componentTypes
	}
	if !utils.IsNil(o.oamTypes) {
		toSerialize["oamTypes"] = o.oamTypes
	}
	if !utils.IsNil(o.resourceNames) {
		toSerialize["resourceNames"] = o.resourceNames
	}
	if !utils.IsNil(o.resourceTypes) {
		toSerialize["resourceTypes"] = o.resourceTypes
	}
	if !utils.IsNil(o.traitTypes) {
		toSerialize["traitTypes"] = o.traitTypes
	}
	return toSerialize, nil
}

type NullableResourcePolicyRuleSelector struct {
	value *ResourcePolicyRuleSelector
	isSet bool
}

func (v NullableResourcePolicyRuleSelector) Get() *ResourcePolicyRuleSelector {
	return v.value
}

func (v *NullableResourcePolicyRuleSelector) Set(val *ResourcePolicyRuleSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcePolicyRuleSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcePolicyRuleSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcePolicyRuleSelector(val *ResourcePolicyRuleSelector) *NullableResourcePolicyRuleSelector {
	return &NullableResourcePolicyRuleSelector{value: val, isSet: true}
}

func (v NullableResourcePolicyRuleSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcePolicyRuleSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
