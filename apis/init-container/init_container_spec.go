/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package init_container

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the InitContainerSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InitContainerSpec{}

// InitContainerSpec struct for InitContainerSpec
type InitContainerSpec struct {
	// Specify the mount path of app container
	appMountPath string `json:"appMountPath"`
	// Specify the args run in the init container
	args []string `json:"args,omitempty"`
	// Specify the commands run in the init container
	cmd []string `json:"cmd,omitempty"`
	// Specify the env run in the init container
	env []Env `json:"env,omitempty"`
	// Specify the extra volume mounts for the init container
	extraVolumeMounts []ExtraVolumeMounts `json:"extraVolumeMounts"`
	// Specify the image of init container
	image string `json:"image"`
	// Specify image pull policy for your service
	imagePullPolicy string `json:"imagePullPolicy"`
	// Specify the mount path of init container
	initMountPath string `json:"initMountPath"`
	// Specify the mount name of shared volume
	mountName string `json:"mountName"`
	// Specify the name of init container
	name string `json:"name"`
}

// NewInitContainerSpec instantiates a new InitContainerSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitContainerSpec(appMountPath string, extraVolumeMounts []ExtraVolumeMounts, image string, imagePullPolicy string, initMountPath string, mountName string, name string) *InitContainerSpec {
	this := InitContainerSpec{}
	this.appMountPath = appMountPath
	this.extraVolumeMounts = extraVolumeMounts
	this.image = image
	this.imagePullPolicy = imagePullPolicy
	this.initMountPath = initMountPath
	this.mountName = mountName
	this.name = name
	return &this
}

// NewInitContainerSpecWithDefaults instantiates a new InitContainerSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitContainerSpecWithDefaults() *InitContainerSpec {
	this := InitContainerSpec{}
	var imagePullPolicy string = "IfNotPresent"
	this.imagePullPolicy = imagePullPolicy
	var mountName string = "workdir"
	this.mountName = mountName
	return &this
}

// GetAppMountPath returns the AppMountPath field value
func (o *InitContainerSpec) GetAppMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.appMountPath
}

// GetAppMountPathOk returns a tuple with the AppMountPath field value
// and a boolean to check if the value has been set.
func (o *InitContainerSpec) GetAppMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.appMountPath, true
}

// AppMountPath sets field value
func (o *InitContainerSpec) AppMountPath(v string) *InitContainerSpec {
	o.appMountPath = v
    return o
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *InitContainerSpec) GetArgs() []string {
	if o == nil || utils.IsNil(o.args) {
		var ret []string
		return ret
	}
	return o.args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitContainerSpec) GetArgsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.args) {
		return nil, false
	}
	return o.args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *InitContainerSpec) HasArgs() bool {
	if o != nil && !utils.IsNil(o.args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []string and assigns it to the args field.
// args:  Specify the args run in the init container 

func (o *InitContainerSpec) Args(v []string) (*InitContainerSpec){
	o.args = v
return o
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *InitContainerSpec) GetCmd() []string {
	if o == nil || utils.IsNil(o.cmd) {
		var ret []string
		return ret
	}
	return o.cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitContainerSpec) GetCmdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.cmd) {
		return nil, false
	}
	return o.cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *InitContainerSpec) HasCmd() bool {
	if o != nil && !utils.IsNil(o.cmd) {
		return true
	}

	return false
}

// SetCmd gets a reference to the given []string and assigns it to the cmd field.
// cmd:  Specify the commands run in the init container 

func (o *InitContainerSpec) Cmd(v []string) (*InitContainerSpec){
	o.cmd = v
return o
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *InitContainerSpec) GetEnv() []Env {
	if o == nil || utils.IsNil(o.env) {
		var ret []Env
		return ret
	}
	return o.env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitContainerSpec) GetEnvOk() ([]Env, bool) {
	if o == nil || utils.IsNil(o.env) {
		return nil, false
	}
	return o.env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *InitContainerSpec) HasEnv() bool {
	if o != nil && !utils.IsNil(o.env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []Env and assigns it to the env field.
// env:  Specify the env run in the init container 

func (o *InitContainerSpec) Env(v []Env) (*InitContainerSpec){
	o.env = v
return o
}

// GetExtraVolumeMounts returns the ExtraVolumeMounts field value
func (o *InitContainerSpec) GetExtraVolumeMounts() []ExtraVolumeMounts {
	if o == nil {
		var ret []ExtraVolumeMounts
		return ret
	}

	return o.extraVolumeMounts
}

// GetExtraVolumeMountsOk returns a tuple with the ExtraVolumeMounts field value
// and a boolean to check if the value has been set.
func (o *InitContainerSpec) GetExtraVolumeMountsOk() ([]ExtraVolumeMounts, bool) {
	if o == nil {
		return nil, false
	}
	return o.extraVolumeMounts, true
}

// ExtraVolumeMounts sets field value
func (o *InitContainerSpec) ExtraVolumeMounts(v []ExtraVolumeMounts) *InitContainerSpec {
	o.extraVolumeMounts = v
    return o
}

// GetImage returns the Image field value
func (o *InitContainerSpec) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *InitContainerSpec) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.image, true
}

// Image sets field value
func (o *InitContainerSpec) Image(v string) *InitContainerSpec {
	o.image = v
    return o
}

// GetImagePullPolicy returns the ImagePullPolicy field value
func (o *InitContainerSpec) GetImagePullPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.imagePullPolicy
}

// GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field value
// and a boolean to check if the value has been set.
func (o *InitContainerSpec) GetImagePullPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.imagePullPolicy, true
}

// ImagePullPolicy sets field value
func (o *InitContainerSpec) ImagePullPolicy(v string) *InitContainerSpec {
	o.imagePullPolicy = v
    return o
}

// GetInitMountPath returns the InitMountPath field value
func (o *InitContainerSpec) GetInitMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.initMountPath
}

// GetInitMountPathOk returns a tuple with the InitMountPath field value
// and a boolean to check if the value has been set.
func (o *InitContainerSpec) GetInitMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.initMountPath, true
}

// InitMountPath sets field value
func (o *InitContainerSpec) InitMountPath(v string) *InitContainerSpec {
	o.initMountPath = v
    return o
}

// GetMountName returns the MountName field value
func (o *InitContainerSpec) GetMountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.mountName
}

// GetMountNameOk returns a tuple with the MountName field value
// and a boolean to check if the value has been set.
func (o *InitContainerSpec) GetMountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.mountName, true
}

// MountName sets field value
func (o *InitContainerSpec) MountName(v string) *InitContainerSpec {
	o.mountName = v
    return o
}

// GetName returns the Name field value
func (o *InitContainerSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InitContainerSpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *InitContainerSpec) Name(v string) *InitContainerSpec {
	o.name = v
    return o
}

func (o InitContainerSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InitContainerSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appMountPath"] = o.appMountPath
	if !utils.IsNil(o.args) {
		toSerialize["args"] = o.args
	}
	if !utils.IsNil(o.cmd) {
		toSerialize["cmd"] = o.cmd
	}
	if !utils.IsNil(o.env) {
		toSerialize["env"] = o.env
	}
	toSerialize["extraVolumeMounts"] = o.extraVolumeMounts
	toSerialize["image"] = o.image
	toSerialize["imagePullPolicy"] = o.imagePullPolicy
	toSerialize["initMountPath"] = o.initMountPath
	toSerialize["mountName"] = o.mountName
	toSerialize["name"] = o.name
	return toSerialize, nil
}

type NullableInitContainerSpec struct {
	value *InitContainerSpec
	isSet bool
}

func (v NullableInitContainerSpec) Get() *InitContainerSpec {
	return v.value
}

func (v *NullableInitContainerSpec) Set(val *InitContainerSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableInitContainerSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableInitContainerSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitContainerSpec(val *InitContainerSpec) *NullableInitContainerSpec {
	return &NullableInitContainerSpec{value: val, isSet: true}
}

func (v NullableInitContainerSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitContainerSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 
