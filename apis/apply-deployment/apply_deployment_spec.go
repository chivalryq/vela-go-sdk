/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_deployment

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the ApplyDeploymentSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApplyDeploymentSpec{}

// ApplyDeploymentSpec struct for ApplyDeploymentSpec
type ApplyDeploymentSpec struct {
	cmd []string `json:"cmd,omitempty"`
	image string `json:"image"`
}

// NewApplyDeploymentSpec instantiates a new ApplyDeploymentSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyDeploymentSpec(image string) *ApplyDeploymentSpec {
	this := ApplyDeploymentSpec{}
	this.image = image
	return &this
}

// NewApplyDeploymentSpecWithDefaults instantiates a new ApplyDeploymentSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyDeploymentSpecWithDefaults() *ApplyDeploymentSpec {
	this := ApplyDeploymentSpec{}
	return &this
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *ApplyDeploymentSpec) GetCmd() []string {
	if o == nil || utils.IsNil(o.cmd) {
		var ret []string
		return ret
	}
	return o.cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyDeploymentSpec) GetCmdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.cmd) {
		return nil, false
	}
	return o.cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *ApplyDeploymentSpec) HasCmd() bool {
	if o != nil && !utils.IsNil(o.cmd) {
		return true
	}

	return false
}

// SetCmd gets a reference to the given []string and assigns it to the cmd field.
// cmd: 

func (o *ApplyDeploymentSpec) Cmd(v []string) (*ApplyDeploymentSpec){
	o.cmd = v
return o
}

// GetImage returns the Image field value
func (o *ApplyDeploymentSpec) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *ApplyDeploymentSpec) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.image, true
}

// Image sets field value
func (o *ApplyDeploymentSpec) Image(v string) *ApplyDeploymentSpec {
	o.image = v
    return o
}

func (o ApplyDeploymentSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyDeploymentSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.cmd) {
		toSerialize["cmd"] = o.cmd
	}
	toSerialize["image"] = o.image
	return toSerialize, nil
}

type NullableApplyDeploymentSpec struct {
	value *ApplyDeploymentSpec
	isSet bool
}

func (v NullableApplyDeploymentSpec) Get() *ApplyDeploymentSpec {
	return v.value
}

func (v *NullableApplyDeploymentSpec) Set(val *ApplyDeploymentSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyDeploymentSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyDeploymentSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyDeploymentSpec(val *ApplyDeploymentSpec) *NullableApplyDeploymentSpec {
	return &NullableApplyDeploymentSpec{value: val, isSet: true}
}

func (v NullableApplyDeploymentSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyDeploymentSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 
