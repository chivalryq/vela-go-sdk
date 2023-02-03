/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ref_objects

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the RefObjectsSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RefObjectsSpec{}

// RefObjectsSpec struct for RefObjectsSpec
type RefObjectsSpec struct {
	// If specified, application will fetch native Kubernetes objects according to the object description
	Objects []K8sObject `json:"objects,omitempty"`
	// If specified, the objects in the urls will be loaded.
	Urls []string `json:"urls,omitempty"`
}

// NewRefObjectsSpecWith instantiates a new RefObjectsSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefObjectsSpecWith() *RefObjectsSpec {
	this := RefObjectsSpec{}
	return &this
}

// NewRefObjectsSpec instantiates a new RefObjectsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefObjectsSpec() *RefObjectsSpec {
	this := RefObjectsSpec{}
	return &this
}

// NewRefObjectsSpecs converts a list RefObjectsSpec pointers to objects.
// This is helpful when the SetRefObjectsSpec requires a list of objects
func NewRefObjectsSpecs(ps ...*RefObjectsSpec) []RefObjectsSpec {
	objs := []RefObjectsSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *RefObjectsComponent) GetObjects() []K8sObject {
	if o == nil || utils.IsNil(o.Properties.Objects) {
		var ret []K8sObject
		return ret
	}
	return o.Properties.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefObjectsComponent) GetObjectsOk() ([]K8sObject, bool) {
	if o == nil || utils.IsNil(o.Properties.Objects) {
		return nil, false
	}
	return o.Properties.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *RefObjectsComponent) HasObjects() bool {
	if o != nil && !utils.IsNil(o.Properties.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []K8sObject and assigns it to the objects field.
// Objects:  If specified, application will fetch native Kubernetes objects according to the object description
func (o *RefObjectsComponent) SetObjects(v []K8sObject) *RefObjectsComponent {
	o.Properties.Objects = v
	return o
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *RefObjectsComponent) GetUrls() []string {
	if o == nil || utils.IsNil(o.Properties.Urls) {
		var ret []string
		return ret
	}
	return o.Properties.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefObjectsComponent) GetUrlsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Urls) {
		return nil, false
	}
	return o.Properties.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *RefObjectsComponent) HasUrls() bool {
	if o != nil && !utils.IsNil(o.Properties.Urls) {
		return true
	}

	return false
}

// SetUrls gets a reference to the given []string and assigns it to the urls field.
// Urls:  If specified, the objects in the urls will be loaded.
func (o *RefObjectsComponent) SetUrls(v []string) *RefObjectsComponent {
	o.Properties.Urls = v
	return o
}

func (o RefObjectsSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefObjectsSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Objects) {
		toSerialize["objects"] = o.Objects
	}
	if !utils.IsNil(o.Urls) {
		toSerialize["urls"] = o.Urls
	}
	return toSerialize, nil
}

type NullableRefObjectsSpec struct {
	value *RefObjectsSpec
	isSet bool
}

func (v NullableRefObjectsSpec) Get() *RefObjectsSpec {
	return v.value
}

func (v *NullableRefObjectsSpec) Set(val *RefObjectsSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableRefObjectsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableRefObjectsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefObjectsSpec(val *RefObjectsSpec) *NullableRefObjectsSpec {
	return &NullableRefObjectsSpec{value: val, isSet: true}
}

func (v NullableRefObjectsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefObjectsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const RefObjectsType = "ref-objects"

func init() {
	sdkcommon.RegisterComponent(RefObjectsType, FromComponent)
}

type RefObjectsComponent struct {
	Base       apis.ComponentBase
	Properties RefObjectsSpec
}

func RefObjects(name string) *RefObjectsComponent {
	r := &RefObjectsComponent{Base: apis.ComponentBase{
		Name: name,
		Type: RefObjectsType,
	}}
	return r
}

func (r *RefObjectsComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range r.Base.Traits {
		traits = append(traits, trait.Build())
	}
	res := common.ApplicationComponent{
		DependsOn:  r.Base.DependsOn,
		Inputs:     r.Base.Inputs,
		Name:       r.Base.Name,
		Outputs:    r.Base.Outputs,
		Properties: util.Object2RawExtension(r.Properties),
		Traits:     traits,
		Type:       RefObjectsType,
	}
	return res
}

func (r *RefObjectsComponent) FromComponent(from common.ApplicationComponent) (*RefObjectsComponent, error) {
	for _, trait := range from.Traits {
		_t, err := sdkcommon.FromTrait(&trait)
		if err != nil {
			return nil, err
		}
		r.Base.Traits = append(r.Base.Traits, _t)
	}
	var properties RefObjectsSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	r.Base.Name = from.Name
	r.Base.DependsOn = from.DependsOn
	r.Base.Inputs = from.Inputs
	r.Base.Outputs = from.Outputs
	r.Base.Type = RefObjectsType
	r.Properties = properties
	return r, nil
}

func FromComponent(from common.ApplicationComponent) (apis.Component, error) {
	r := &RefObjectsComponent{}
	return r.FromComponent(from)
}

func (r *RefObjectsComponent) AddTrait(traits ...apis.Trait) *RefObjectsComponent {
	r.Base.Traits = append(r.Base.Traits, traits...)
	return r
}

func (r *RefObjectsComponent) GetTrait(_type string) apis.Trait {
	for _, _t := range r.Base.Traits {
		if _t.DefType() == _type {
			return _t
		}
	}
	return nil
}

func (r *RefObjectsComponent) ComponentName() string {
	return r.Base.Name
}

func (r *RefObjectsComponent) DefType() string {
	return RefObjectsType
}

func (r *RefObjectsComponent) DependsOn(dependsOn []string) *RefObjectsComponent {
	r.Base.DependsOn = dependsOn
	return r
}

func (r *RefObjectsComponent) Inputs(input common.StepInputs) *RefObjectsComponent {
	r.Base.Inputs = input
	return r
}

func (r *RefObjectsComponent) Outputs(output common.StepOutputs) *RefObjectsComponent {
	r.Base.Outputs = output
	return r
}

func (r *RefObjectsComponent) AddDependsOn(dependsOn string) *RefObjectsComponent {
	r.Base.DependsOn = append(r.Base.DependsOn, dependsOn)
	return r
}
