/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_terraform_config

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the WriteConnectionSecretToRef type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WriteConnectionSecretToRef{}

// WriteConnectionSecretToRef this specifies the namespace and name of a secret to which any connection details for this managed resource should be written.
type WriteConnectionSecretToRef struct {
	name      string `json:"name"`
	namespace string `json:"namespace"`
}

// NewWriteConnectionSecretToRefWith instantiates a new WriteConnectionSecretToRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteConnectionSecretToRefWith(name string, namespace string) *WriteConnectionSecretToRef {
	this := WriteConnectionSecretToRef{}
	this.name = name
	this.namespace = namespace
	return &this
}

// NewWriteConnectionSecretToRef instantiates a new WriteConnectionSecretToRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteConnectionSecretToRef() *WriteConnectionSecretToRef {
	this := WriteConnectionSecretToRef{}
	return &this
}

// GetName returns the Name field value
func (o *WriteConnectionSecretToRef) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WriteConnectionSecretToRef) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *WriteConnectionSecretToRef) Name(v string) *WriteConnectionSecretToRef {
	o.name = v
	return o
}

// GetNamespace returns the Namespace field value
func (o *WriteConnectionSecretToRef) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *WriteConnectionSecretToRef) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.namespace, true
}

// Namespace sets field value
func (o *WriteConnectionSecretToRef) Namespace(v string) *WriteConnectionSecretToRef {
	o.namespace = v
	return o
}

func (o WriteConnectionSecretToRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WriteConnectionSecretToRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.name
	toSerialize["namespace"] = o.namespace
	return toSerialize, nil
}

type NullableWriteConnectionSecretToRef struct {
	value *WriteConnectionSecretToRef
	isSet bool
}

func (v NullableWriteConnectionSecretToRef) Get() *WriteConnectionSecretToRef {
	return v.value
}

func (v *NullableWriteConnectionSecretToRef) Set(val *WriteConnectionSecretToRef) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteConnectionSecretToRef) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteConnectionSecretToRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteConnectionSecretToRef(val *WriteConnectionSecretToRef) *NullableWriteConnectionSecretToRef {
	return &NullableWriteConnectionSecretToRef{value: val, isSet: true}
}

func (v NullableWriteConnectionSecretToRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteConnectionSecretToRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
