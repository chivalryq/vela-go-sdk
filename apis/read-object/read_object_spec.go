/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package read_object

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the ReadObjectSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReadObjectSpec{}

// ReadObjectSpec struct for ReadObjectSpec
type ReadObjectSpec struct {
	// Specify the apiVersion of the object, defaults to 'core.oam.dev/v1beta1'
	apiVersion *string `json:"apiVersion,omitempty"`
	// The cluster you want to apply the resource to, default is the current control plane cluster
	cluster string `json:"cluster"`
	// Specify the kind of the object, defaults to Application
	kind *string `json:"kind,omitempty"`
	// Specify the name of the object
	name string `json:"name"`
	// The namespace of the resource you want to read
	namespace *string `json:"namespace,omitempty"`
}

// NewReadObjectSpec instantiates a new ReadObjectSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadObjectSpec(cluster string, name string) *ReadObjectSpec {
	this := ReadObjectSpec{}
	this.cluster = cluster
	this.name = name
	var namespace string = "default"
	this.namespace = &namespace
	return &this
}

// NewReadObjectSpecWithDefaults instantiates a new ReadObjectSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadObjectSpecWithDefaults() *ReadObjectSpec {
	this := ReadObjectSpec{}
	var cluster string = ""
	this.cluster = cluster
	var namespace string = "default"
	this.namespace = &namespace
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ReadObjectSpec) GetApiVersion() string {
	if o == nil || utils.IsNil(o.apiVersion) {
		var ret string
		return ret
	}
	return *o.apiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadObjectSpec) GetApiVersionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.apiVersion) {
		return nil, false
	}
	return o.apiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ReadObjectSpec) HasApiVersion() bool {
	if o != nil && !utils.IsNil(o.apiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the apiVersion field.
// apiVersion:  Specify the apiVersion of the object, defaults to 'core.oam.dev/v1beta1' 

func (o *ReadObjectSpec) ApiVersion(v string) (*ReadObjectSpec){
	o.apiVersion = &v
return o
}

// GetCluster returns the Cluster field value
func (o *ReadObjectSpec) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *ReadObjectSpec) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.cluster, true
}

// Cluster sets field value
func (o *ReadObjectSpec) Cluster(v string) *ReadObjectSpec {
	o.cluster = v
    return o
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ReadObjectSpec) GetKind() string {
	if o == nil || utils.IsNil(o.kind) {
		var ret string
		return ret
	}
	return *o.kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadObjectSpec) GetKindOk() (*string, bool) {
	if o == nil || utils.IsNil(o.kind) {
		return nil, false
	}
	return o.kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ReadObjectSpec) HasKind() bool {
	if o != nil && !utils.IsNil(o.kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the kind field.
// kind:  Specify the kind of the object, defaults to Application 

func (o *ReadObjectSpec) Kind(v string) (*ReadObjectSpec){
	o.kind = &v
return o
}

// GetName returns the Name field value
func (o *ReadObjectSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReadObjectSpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *ReadObjectSpec) Name(v string) *ReadObjectSpec {
	o.name = v
    return o
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ReadObjectSpec) GetNamespace() string {
	if o == nil || utils.IsNil(o.namespace) {
		var ret string
		return ret
	}
	return *o.namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadObjectSpec) GetNamespaceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.namespace) {
		return nil, false
	}
	return o.namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ReadObjectSpec) HasNamespace() bool {
	if o != nil && !utils.IsNil(o.namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the namespace field.
// namespace:  The namespace of the resource you want to read 

func (o *ReadObjectSpec) Namespace(v string) (*ReadObjectSpec){
	o.namespace = &v
return o
}

func (o ReadObjectSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadObjectSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.apiVersion) {
		toSerialize["apiVersion"] = o.apiVersion
	}
	toSerialize["cluster"] = o.cluster
	if !utils.IsNil(o.kind) {
		toSerialize["kind"] = o.kind
	}
	toSerialize["name"] = o.name
	if !utils.IsNil(o.namespace) {
		toSerialize["namespace"] = o.namespace
	}
	return toSerialize, nil
}

type NullableReadObjectSpec struct {
	value *ReadObjectSpec
	isSet bool
}

func (v NullableReadObjectSpec) Get() *ReadObjectSpec {
	return v.value
}

func (v *NullableReadObjectSpec) Set(val *ReadObjectSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableReadObjectSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableReadObjectSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadObjectSpec(val *ReadObjectSpec) *NullableReadObjectSpec {
	return &NullableReadObjectSpec{value: val, isSet: true}
}

func (v NullableReadObjectSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadObjectSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 
