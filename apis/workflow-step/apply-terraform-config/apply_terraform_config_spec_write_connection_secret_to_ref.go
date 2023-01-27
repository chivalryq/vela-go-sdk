/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_terraform_config

import (
	"encoding/json"

	"vela-go-sdk/apis/utils"
)

// checks if the ApplyTerraformConfigSpecWriteConnectionSecretToRef type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApplyTerraformConfigSpecWriteConnectionSecretToRef{}

// ApplyTerraformConfigSpecWriteConnectionSecretToRef this specifies the namespace and name of a secret to which any connection details for this managed resource should be written.
type ApplyTerraformConfigSpecWriteConnectionSecretToRef struct {
	name      string `json:"name"`
	namespace string `json:"namespace"`
}

// ApplyTerraformConfigSpecWriteConnectionSecretToRefWith instantiates a new ApplyTerraformConfigSpecWriteConnectionSecretToRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func ApplyTerraformConfigSpecWriteConnectionSecretToRefWith(name string, namespace string) *ApplyTerraformConfigSpecWriteConnectionSecretToRef {
	this := ApplyTerraformConfigSpecWriteConnectionSecretToRef{}
	this.name = name
	this.namespace = namespace
	return &this
}

// NewApplyTerraformConfigSpecWriteConnectionSecretToRef instantiates a new ApplyTerraformConfigSpecWriteConnectionSecretToRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyTerraformConfigSpecWriteConnectionSecretToRef() *ApplyTerraformConfigSpecWriteConnectionSecretToRef {
	this := ApplyTerraformConfigSpecWriteConnectionSecretToRef{}
	return &this
}

// GetName returns the Name field value
func (o *ApplyTerraformConfigSpecWriteConnectionSecretToRef) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplyTerraformConfigSpecWriteConnectionSecretToRef) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *ApplyTerraformConfigSpecWriteConnectionSecretToRef) Name(v string) *ApplyTerraformConfigSpecWriteConnectionSecretToRef {
	o.name = v
	return o
}

// GetNamespace returns the Namespace field value
func (o *ApplyTerraformConfigSpecWriteConnectionSecretToRef) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *ApplyTerraformConfigSpecWriteConnectionSecretToRef) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.namespace, true
}

// Namespace sets field value
func (o *ApplyTerraformConfigSpecWriteConnectionSecretToRef) Namespace(v string) *ApplyTerraformConfigSpecWriteConnectionSecretToRef {
	o.namespace = v
	return o
}

func (o ApplyTerraformConfigSpecWriteConnectionSecretToRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyTerraformConfigSpecWriteConnectionSecretToRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.name
	toSerialize["namespace"] = o.namespace
	return toSerialize, nil
}

type NullableApplyTerraformConfigSpecWriteConnectionSecretToRef struct {
	value *ApplyTerraformConfigSpecWriteConnectionSecretToRef
	isSet bool
}

func (v NullableApplyTerraformConfigSpecWriteConnectionSecretToRef) Get() *ApplyTerraformConfigSpecWriteConnectionSecretToRef {
	return v.value
}

func (v *NullableApplyTerraformConfigSpecWriteConnectionSecretToRef) Set(val *ApplyTerraformConfigSpecWriteConnectionSecretToRef) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyTerraformConfigSpecWriteConnectionSecretToRef) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyTerraformConfigSpecWriteConnectionSecretToRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyTerraformConfigSpecWriteConnectionSecretToRef(val *ApplyTerraformConfigSpecWriteConnectionSecretToRef) *NullableApplyTerraformConfigSpecWriteConnectionSecretToRef {
	return &NullableApplyTerraformConfigSpecWriteConnectionSecretToRef{value: val, isSet: true}
}

func (v NullableApplyTerraformConfigSpecWriteConnectionSecretToRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyTerraformConfigSpecWriteConnectionSecretToRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
