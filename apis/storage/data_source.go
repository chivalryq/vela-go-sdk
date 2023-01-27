/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the DataSource type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DataSource{}

// DataSource struct for DataSource
type DataSource struct {
	apiGroup string `json:"apiGroup"`
	kind string `json:"kind"`
	name string `json:"name"`
}

// NewDataSource instantiates a new DataSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSource(apiGroup string, kind string, name string) *DataSource {
	this := DataSource{}
	this.apiGroup = apiGroup
	this.kind = kind
	this.name = name
	return &this
}

// NewDataSourceWithDefaults instantiates a new DataSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourceWithDefaults() *DataSource {
	this := DataSource{}
	return &this
}

// GetApiGroup returns the ApiGroup field value
func (o *DataSource) GetApiGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.apiGroup
}

// GetApiGroupOk returns a tuple with the ApiGroup field value
// and a boolean to check if the value has been set.
func (o *DataSource) GetApiGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.apiGroup, true
}

// ApiGroup sets field value
func (o *DataSource) ApiGroup(v string) *DataSource {
	o.apiGroup = v
    return o
}

// GetKind returns the Kind field value
func (o *DataSource) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *DataSource) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.kind, true
}

// Kind sets field value
func (o *DataSource) Kind(v string) *DataSource {
	o.kind = v
    return o
}

// GetName returns the Name field value
func (o *DataSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DataSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *DataSource) Name(v string) *DataSource {
	o.name = v
    return o
}

func (o DataSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiGroup"] = o.apiGroup
	toSerialize["kind"] = o.kind
	toSerialize["name"] = o.name
	return toSerialize, nil
}

type NullableDataSource struct {
	value *DataSource
	isSet bool
}

func (v NullableDataSource) Get() *DataSource {
	return v.value
}

func (v *NullableDataSource) Set(val *DataSource) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSource) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSource(val *DataSource) *NullableDataSource {
	return &NullableDataSource{value: val, isSet: true}
}

func (v NullableDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 
