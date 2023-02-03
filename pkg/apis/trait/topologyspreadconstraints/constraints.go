/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package topologyspreadconstraints

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the Constraints type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Constraints{}

// Constraints struct for Constraints
type Constraints struct {
	LabelSelector *LabSelector `json:"labelSelector,omitempty"`
	// A list of pod label keys to select the pods over which spreading will be calculated
	MatchLabelKeys []string `json:"matchLabelKeys,omitempty"`
	// Describe the degree to which Pods may be unevenly distributed
	MaxSkew *int32 `json:"maxSkew,omitempty"`
	// Indicate a minimum number of eligible domains
	MinDomains *int32 `json:"minDomains,omitempty"`
	// Indicate how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew
	NodeAffinityPolicy *string `json:"nodeAffinityPolicy,omitempty"`
	// Indicate how we will treat node taints when calculating pod topology spread skew
	NodeTaintsPolicy *string `json:"nodeTaintsPolicy,omitempty"`
	// Specify the key of node labels
	TopologyKey *string `json:"topologyKey,omitempty"`
	// Indicate how to deal with a Pod if it doesn't satisfy the spread constraint
	WhenUnsatisfiable *string `json:"whenUnsatisfiable,omitempty"`
}

// NewConstraintsWith instantiates a new Constraints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstraintsWith() *Constraints {
	this := Constraints{}
	var nodeAffinityPolicy string = "Honor"
	this.NodeAffinityPolicy = &nodeAffinityPolicy
	var nodeTaintsPolicy string = "Honor"
	this.NodeTaintsPolicy = &nodeTaintsPolicy
	var whenUnsatisfiable string = "DoNotSchedule"
	this.WhenUnsatisfiable = &whenUnsatisfiable
	return &this
}

// NewConstraints instantiates a new Constraints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstraints() *Constraints {
	this := Constraints{}
	var nodeAffinityPolicy string = "Honor"
	this.NodeAffinityPolicy = &nodeAffinityPolicy
	var nodeTaintsPolicy string = "Honor"
	this.NodeTaintsPolicy = &nodeTaintsPolicy
	var whenUnsatisfiable string = "DoNotSchedule"
	this.WhenUnsatisfiable = &whenUnsatisfiable
	return &this
}

// NewConstraintss converts a list Constraints pointers to objects.
// This is helpful when the SetConstraints requires a list of objects
func NewConstraintss(ps ...*Constraints) []Constraints {
	objs := []Constraints{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetLabelSelector returns the LabelSelector field value if set, zero value otherwise.
func (o *Constraints) GetLabelSelector() LabSelector {
	if o == nil || utils.IsNil(o.LabelSelector) {
		var ret LabSelector
		return ret
	}
	return *o.LabelSelector
}

// GetLabelSelectorOk returns a tuple with the LabelSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraints) GetLabelSelectorOk() (*LabSelector, bool) {
	if o == nil || utils.IsNil(o.LabelSelector) {
		return nil, false
	}
	return o.LabelSelector, true
}

// HasLabelSelector returns a boolean if a field has been set.
func (o *Constraints) HasLabelSelector() bool {
	if o != nil && !utils.IsNil(o.LabelSelector) {
		return true
	}

	return false
}

// SetLabelSelector gets a reference to the given LabSelector and assigns it to the labelSelector field.
// LabelSelector:
func (o *Constraints) SetLabelSelector(v LabSelector) *Constraints {
	o.LabelSelector = &v
	return o
}

// GetMatchLabelKeys returns the MatchLabelKeys field value if set, zero value otherwise.
func (o *Constraints) GetMatchLabelKeys() []string {
	if o == nil || utils.IsNil(o.MatchLabelKeys) {
		var ret []string
		return ret
	}
	return o.MatchLabelKeys
}

// GetMatchLabelKeysOk returns a tuple with the MatchLabelKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraints) GetMatchLabelKeysOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.MatchLabelKeys) {
		return nil, false
	}
	return o.MatchLabelKeys, true
}

// HasMatchLabelKeys returns a boolean if a field has been set.
func (o *Constraints) HasMatchLabelKeys() bool {
	if o != nil && !utils.IsNil(o.MatchLabelKeys) {
		return true
	}

	return false
}

// SetMatchLabelKeys gets a reference to the given []string and assigns it to the matchLabelKeys field.
// MatchLabelKeys:  A list of pod label keys to select the pods over which spreading will be calculated
func (o *Constraints) SetMatchLabelKeys(v []string) *Constraints {
	o.MatchLabelKeys = v
	return o
}

// GetMaxSkew returns the MaxSkew field value if set, zero value otherwise.
func (o *Constraints) GetMaxSkew() int32 {
	if o == nil || utils.IsNil(o.MaxSkew) {
		var ret int32
		return ret
	}
	return *o.MaxSkew
}

// GetMaxSkewOk returns a tuple with the MaxSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraints) GetMaxSkewOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.MaxSkew) {
		return nil, false
	}
	return o.MaxSkew, true
}

// HasMaxSkew returns a boolean if a field has been set.
func (o *Constraints) HasMaxSkew() bool {
	if o != nil && !utils.IsNil(o.MaxSkew) {
		return true
	}

	return false
}

// SetMaxSkew gets a reference to the given int32 and assigns it to the maxSkew field.
// MaxSkew:  Describe the degree to which Pods may be unevenly distributed
func (o *Constraints) SetMaxSkew(v int32) *Constraints {
	o.MaxSkew = &v
	return o
}

// GetMinDomains returns the MinDomains field value if set, zero value otherwise.
func (o *Constraints) GetMinDomains() int32 {
	if o == nil || utils.IsNil(o.MinDomains) {
		var ret int32
		return ret
	}
	return *o.MinDomains
}

// GetMinDomainsOk returns a tuple with the MinDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraints) GetMinDomainsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.MinDomains) {
		return nil, false
	}
	return o.MinDomains, true
}

// HasMinDomains returns a boolean if a field has been set.
func (o *Constraints) HasMinDomains() bool {
	if o != nil && !utils.IsNil(o.MinDomains) {
		return true
	}

	return false
}

// SetMinDomains gets a reference to the given int32 and assigns it to the minDomains field.
// MinDomains:  Indicate a minimum number of eligible domains
func (o *Constraints) SetMinDomains(v int32) *Constraints {
	o.MinDomains = &v
	return o
}

// GetNodeAffinityPolicy returns the NodeAffinityPolicy field value if set, zero value otherwise.
func (o *Constraints) GetNodeAffinityPolicy() string {
	if o == nil || utils.IsNil(o.NodeAffinityPolicy) {
		var ret string
		return ret
	}
	return *o.NodeAffinityPolicy
}

// GetNodeAffinityPolicyOk returns a tuple with the NodeAffinityPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraints) GetNodeAffinityPolicyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.NodeAffinityPolicy) {
		return nil, false
	}
	return o.NodeAffinityPolicy, true
}

// HasNodeAffinityPolicy returns a boolean if a field has been set.
func (o *Constraints) HasNodeAffinityPolicy() bool {
	if o != nil && !utils.IsNil(o.NodeAffinityPolicy) {
		return true
	}

	return false
}

// SetNodeAffinityPolicy gets a reference to the given string and assigns it to the nodeAffinityPolicy field.
// NodeAffinityPolicy:  Indicate how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew
func (o *Constraints) SetNodeAffinityPolicy(v string) *Constraints {
	o.NodeAffinityPolicy = &v
	return o
}

// GetNodeTaintsPolicy returns the NodeTaintsPolicy field value if set, zero value otherwise.
func (o *Constraints) GetNodeTaintsPolicy() string {
	if o == nil || utils.IsNil(o.NodeTaintsPolicy) {
		var ret string
		return ret
	}
	return *o.NodeTaintsPolicy
}

// GetNodeTaintsPolicyOk returns a tuple with the NodeTaintsPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraints) GetNodeTaintsPolicyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.NodeTaintsPolicy) {
		return nil, false
	}
	return o.NodeTaintsPolicy, true
}

// HasNodeTaintsPolicy returns a boolean if a field has been set.
func (o *Constraints) HasNodeTaintsPolicy() bool {
	if o != nil && !utils.IsNil(o.NodeTaintsPolicy) {
		return true
	}

	return false
}

// SetNodeTaintsPolicy gets a reference to the given string and assigns it to the nodeTaintsPolicy field.
// NodeTaintsPolicy:  Indicate how we will treat node taints when calculating pod topology spread skew
func (o *Constraints) SetNodeTaintsPolicy(v string) *Constraints {
	o.NodeTaintsPolicy = &v
	return o
}

// GetTopologyKey returns the TopologyKey field value if set, zero value otherwise.
func (o *Constraints) GetTopologyKey() string {
	if o == nil || utils.IsNil(o.TopologyKey) {
		var ret string
		return ret
	}
	return *o.TopologyKey
}

// GetTopologyKeyOk returns a tuple with the TopologyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraints) GetTopologyKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TopologyKey) {
		return nil, false
	}
	return o.TopologyKey, true
}

// HasTopologyKey returns a boolean if a field has been set.
func (o *Constraints) HasTopologyKey() bool {
	if o != nil && !utils.IsNil(o.TopologyKey) {
		return true
	}

	return false
}

// SetTopologyKey gets a reference to the given string and assigns it to the topologyKey field.
// TopologyKey:  Specify the key of node labels
func (o *Constraints) SetTopologyKey(v string) *Constraints {
	o.TopologyKey = &v
	return o
}

// GetWhenUnsatisfiable returns the WhenUnsatisfiable field value if set, zero value otherwise.
func (o *Constraints) GetWhenUnsatisfiable() string {
	if o == nil || utils.IsNil(o.WhenUnsatisfiable) {
		var ret string
		return ret
	}
	return *o.WhenUnsatisfiable
}

// GetWhenUnsatisfiableOk returns a tuple with the WhenUnsatisfiable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraints) GetWhenUnsatisfiableOk() (*string, bool) {
	if o == nil || utils.IsNil(o.WhenUnsatisfiable) {
		return nil, false
	}
	return o.WhenUnsatisfiable, true
}

// HasWhenUnsatisfiable returns a boolean if a field has been set.
func (o *Constraints) HasWhenUnsatisfiable() bool {
	if o != nil && !utils.IsNil(o.WhenUnsatisfiable) {
		return true
	}

	return false
}

// SetWhenUnsatisfiable gets a reference to the given string and assigns it to the whenUnsatisfiable field.
// WhenUnsatisfiable:  Indicate how to deal with a Pod if it doesn't satisfy the spread constraint
func (o *Constraints) SetWhenUnsatisfiable(v string) *Constraints {
	o.WhenUnsatisfiable = &v
	return o
}

func (o Constraints) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Constraints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.LabelSelector) {
		toSerialize["labelSelector"] = o.LabelSelector
	}
	if !utils.IsNil(o.MatchLabelKeys) {
		toSerialize["matchLabelKeys"] = o.MatchLabelKeys
	}
	if !utils.IsNil(o.MaxSkew) {
		toSerialize["maxSkew"] = o.MaxSkew
	}
	if !utils.IsNil(o.MinDomains) {
		toSerialize["minDomains"] = o.MinDomains
	}
	if !utils.IsNil(o.NodeAffinityPolicy) {
		toSerialize["nodeAffinityPolicy"] = o.NodeAffinityPolicy
	}
	if !utils.IsNil(o.NodeTaintsPolicy) {
		toSerialize["nodeTaintsPolicy"] = o.NodeTaintsPolicy
	}
	if !utils.IsNil(o.TopologyKey) {
		toSerialize["topologyKey"] = o.TopologyKey
	}
	if !utils.IsNil(o.WhenUnsatisfiable) {
		toSerialize["whenUnsatisfiable"] = o.WhenUnsatisfiable
	}
	return toSerialize, nil
}

type NullableConstraints struct {
	value *Constraints
	isSet bool
}

func (v NullableConstraints) Get() *Constraints {
	return v.value
}

func (v *NullableConstraints) Set(val *Constraints) {
	v.value = val
	v.isSet = true
}

func (v NullableConstraints) IsSet() bool {
	return v.isSet
}

func (v *NullableConstraints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstraints(val *Constraints) *NullableConstraints {
	return &NullableConstraints{value: val, isSet: true}
}

func (v NullableConstraints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstraints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
