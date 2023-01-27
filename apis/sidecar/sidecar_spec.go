/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sidecar

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the SidecarSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SidecarSpec{}

// SidecarSpec struct for SidecarSpec
type SidecarSpec struct {
	// Specify the args in the sidecar
	args []string `json:"args,omitempty"`
	// Specify the commands run in the sidecar
	cmd []string `json:"cmd,omitempty"`
	// Specify the env in the sidecar
	env []Env `json:"env,omitempty"`
	// Specify the image of sidecar container
	image string `json:"image"`
	livenessProbe *HealthProbe `json:"livenessProbe,omitempty"`
	// Specify the name of sidecar container
	name string `json:"name"`
	readinessProbe *HealthProbe `json:"readinessProbe,omitempty"`
	// Specify the shared volume path
	volumes []Volumes `json:"volumes,omitempty"`
}

// NewSidecarSpec instantiates a new SidecarSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSidecarSpec(image string, name string) *SidecarSpec {
	this := SidecarSpec{}
	this.image = image
	this.name = name
	return &this
}

// NewSidecarSpecWithDefaults instantiates a new SidecarSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSidecarSpecWithDefaults() *SidecarSpec {
	this := SidecarSpec{}
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *SidecarSpec) GetArgs() []string {
	if o == nil || utils.IsNil(o.args) {
		var ret []string
		return ret
	}
	return o.args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SidecarSpec) GetArgsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.args) {
		return nil, false
	}
	return o.args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *SidecarSpec) HasArgs() bool {
	if o != nil && !utils.IsNil(o.args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []string and assigns it to the args field.
// args:  Specify the args in the sidecar 

func (o *SidecarSpec) Args(v []string) (*SidecarSpec){
	o.args = v
return o
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *SidecarSpec) GetCmd() []string {
	if o == nil || utils.IsNil(o.cmd) {
		var ret []string
		return ret
	}
	return o.cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SidecarSpec) GetCmdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.cmd) {
		return nil, false
	}
	return o.cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *SidecarSpec) HasCmd() bool {
	if o != nil && !utils.IsNil(o.cmd) {
		return true
	}

	return false
}

// SetCmd gets a reference to the given []string and assigns it to the cmd field.
// cmd:  Specify the commands run in the sidecar 

func (o *SidecarSpec) Cmd(v []string) (*SidecarSpec){
	o.cmd = v
return o
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *SidecarSpec) GetEnv() []Env {
	if o == nil || utils.IsNil(o.env) {
		var ret []Env
		return ret
	}
	return o.env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SidecarSpec) GetEnvOk() ([]Env, bool) {
	if o == nil || utils.IsNil(o.env) {
		return nil, false
	}
	return o.env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *SidecarSpec) HasEnv() bool {
	if o != nil && !utils.IsNil(o.env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []Env and assigns it to the env field.
// env:  Specify the env in the sidecar 

func (o *SidecarSpec) Env(v []Env) (*SidecarSpec){
	o.env = v
return o
}

// GetImage returns the Image field value
func (o *SidecarSpec) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *SidecarSpec) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.image, true
}

// Image sets field value
func (o *SidecarSpec) Image(v string) *SidecarSpec {
	o.image = v
    return o
}

// GetLivenessProbe returns the LivenessProbe field value if set, zero value otherwise.
func (o *SidecarSpec) GetLivenessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.livenessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.livenessProbe
}

// GetLivenessProbeOk returns a tuple with the LivenessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SidecarSpec) GetLivenessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.livenessProbe) {
		return nil, false
	}
	return o.livenessProbe, true
}

// HasLivenessProbe returns a boolean if a field has been set.
func (o *SidecarSpec) HasLivenessProbe() bool {
	if o != nil && !utils.IsNil(o.livenessProbe) {
		return true
	}

	return false
}

// SetLivenessProbe gets a reference to the given HealthProbe and assigns it to the livenessProbe field.
// livenessProbe: 

func (o *SidecarSpec) LivenessProbe(v HealthProbe) (*SidecarSpec){
	o.livenessProbe = &v
return o
}

// GetName returns the Name field value
func (o *SidecarSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SidecarSpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *SidecarSpec) Name(v string) *SidecarSpec {
	o.name = v
    return o
}

// GetReadinessProbe returns the ReadinessProbe field value if set, zero value otherwise.
func (o *SidecarSpec) GetReadinessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.readinessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.readinessProbe
}

// GetReadinessProbeOk returns a tuple with the ReadinessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SidecarSpec) GetReadinessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.readinessProbe) {
		return nil, false
	}
	return o.readinessProbe, true
}

// HasReadinessProbe returns a boolean if a field has been set.
func (o *SidecarSpec) HasReadinessProbe() bool {
	if o != nil && !utils.IsNil(o.readinessProbe) {
		return true
	}

	return false
}

// SetReadinessProbe gets a reference to the given HealthProbe and assigns it to the readinessProbe field.
// readinessProbe: 

func (o *SidecarSpec) ReadinessProbe(v HealthProbe) (*SidecarSpec){
	o.readinessProbe = &v
return o
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *SidecarSpec) GetVolumes() []Volumes {
	if o == nil || utils.IsNil(o.volumes) {
		var ret []Volumes
		return ret
	}
	return o.volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SidecarSpec) GetVolumesOk() ([]Volumes, bool) {
	if o == nil || utils.IsNil(o.volumes) {
		return nil, false
	}
	return o.volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *SidecarSpec) HasVolumes() bool {
	if o != nil && !utils.IsNil(o.volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []Volumes and assigns it to the volumes field.
// volumes:  Specify the shared volume path 

func (o *SidecarSpec) Volumes(v []Volumes) (*SidecarSpec){
	o.volumes = v
return o
}

func (o SidecarSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SidecarSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.args) {
		toSerialize["args"] = o.args
	}
	if !utils.IsNil(o.cmd) {
		toSerialize["cmd"] = o.cmd
	}
	if !utils.IsNil(o.env) {
		toSerialize["env"] = o.env
	}
	toSerialize["image"] = o.image
	if !utils.IsNil(o.livenessProbe) {
		toSerialize["livenessProbe"] = o.livenessProbe
	}
	toSerialize["name"] = o.name
	if !utils.IsNil(o.readinessProbe) {
		toSerialize["readinessProbe"] = o.readinessProbe
	}
	if !utils.IsNil(o.volumes) {
		toSerialize["volumes"] = o.volumes
	}
	return toSerialize, nil
}

type NullableSidecarSpec struct {
	value *SidecarSpec
	isSet bool
}

func (v NullableSidecarSpec) Get() *SidecarSpec {
	return v.value
}

func (v *NullableSidecarSpec) Set(val *SidecarSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSidecarSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSidecarSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSidecarSpec(val *SidecarSpec) *NullableSidecarSpec {
	return &NullableSidecarSpec{value: val, isSet: true}
}

func (v NullableSidecarSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSidecarSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 