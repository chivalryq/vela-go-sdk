/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package task

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the TaskSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TaskSpec{}

// TaskSpec struct for TaskSpec
type TaskSpec struct {
	// Specify the annotations in the workload
	annotations *map[string]string `json:"annotations,omitempty"`
	// Commands to run in the container
	cmd []string `json:"cmd,omitempty"`
	// Specify number of tasks to run in parallel +short=c
	count int32 `json:"count"`
	// Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
	cpu *string `json:"cpu,omitempty"`
	// Define arguments by using environment variables
	env []TaskSpecEnv `json:"env,omitempty"`
	// Which image would you like to use for your service +short=i
	image string `json:"image"`
	// Specify image pull policy for your service
	imagePullPolicy *string `json:"imagePullPolicy,omitempty"`
	// Specify image pull secrets for your service
	imagePullSecrets []string `json:"imagePullSecrets,omitempty"`
	// Specify the labels in the workload
	labels *map[string]string `json:"labels,omitempty"`
	livenessProbe *HealthProbe `json:"livenessProbe,omitempty"`
	// Specifies the attributes of the memory resource required for the container.
	memory *string `json:"memory,omitempty"`
	readinessProbe *HealthProbe `json:"readinessProbe,omitempty"`
	// Define the job restart policy, the value can only be Never or OnFailure. By default, it's Never.
	restart string `json:"restart"`
	// Declare volumes and volumeMounts
	volumes []TaskSpecVolumes `json:"volumes,omitempty"`
}

// NewTaskSpec instantiates a new TaskSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskSpec(count int32, image string, restart string) *TaskSpec {
	this := TaskSpec{}
	this.count = count
	this.image = image
	this.restart = restart
	return &this
}

// NewTaskSpecWithDefaults instantiates a new TaskSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskSpecWithDefaults() *TaskSpec {
	this := TaskSpec{}
	var count int32 = 1
	this.count = count
	var restart string = "Never"
	this.restart = restart
	return &this
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *TaskSpec) GetAnnotations() map[string]string {
	if o == nil || utils.IsNil(o.annotations) {
		var ret map[string]string
		return ret
	}
	return *o.annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSpec) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.annotations) {
		return nil, false
	}
	return o.annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *TaskSpec) HasAnnotations() bool {
	if o != nil && !utils.IsNil(o.annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the annotations field.
// annotations:  Specify the annotations in the workload 

func (o *TaskSpec) Annotations(v map[string]string) (*TaskSpec){
	o.annotations = &v
return o
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *TaskSpec) GetCmd() []string {
	if o == nil || utils.IsNil(o.cmd) {
		var ret []string
		return ret
	}
	return o.cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSpec) GetCmdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.cmd) {
		return nil, false
	}
	return o.cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *TaskSpec) HasCmd() bool {
	if o != nil && !utils.IsNil(o.cmd) {
		return true
	}

	return false
}

// SetCmd gets a reference to the given []string and assigns it to the cmd field.
// cmd:  Commands to run in the container 

func (o *TaskSpec) Cmd(v []string) (*TaskSpec){
	o.cmd = v
return o
}

// GetCount returns the Count field value
func (o *TaskSpec) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *TaskSpec) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.count, true
}

// Count sets field value
func (o *TaskSpec) Count(v int32) *TaskSpec {
	o.count = v
    return o
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *TaskSpec) GetCpu() string {
	if o == nil || utils.IsNil(o.cpu) {
		var ret string
		return ret
	}
	return *o.cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSpec) GetCpuOk() (*string, bool) {
	if o == nil || utils.IsNil(o.cpu) {
		return nil, false
	}
	return o.cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *TaskSpec) HasCpu() bool {
	if o != nil && !utils.IsNil(o.cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given string and assigns it to the cpu field.
// cpu:  Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core) 

func (o *TaskSpec) Cpu(v string) (*TaskSpec){
	o.cpu = &v
return o
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *TaskSpec) GetEnv() []TaskSpecEnv {
	if o == nil || utils.IsNil(o.env) {
		var ret []TaskSpecEnv
		return ret
	}
	return o.env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSpec) GetEnvOk() ([]TaskSpecEnv, bool) {
	if o == nil || utils.IsNil(o.env) {
		return nil, false
	}
	return o.env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *TaskSpec) HasEnv() bool {
	if o != nil && !utils.IsNil(o.env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []TaskSpecEnv and assigns it to the env field.
// env:  Define arguments by using environment variables 

func (o *TaskSpec) Env(v []TaskSpecEnv) (*TaskSpec){
	o.env = v
return o
}

// GetImage returns the Image field value
func (o *TaskSpec) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *TaskSpec) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.image, true
}

// Image sets field value
func (o *TaskSpec) Image(v string) *TaskSpec {
	o.image = v
    return o
}

// GetImagePullPolicy returns the ImagePullPolicy field value if set, zero value otherwise.
func (o *TaskSpec) GetImagePullPolicy() string {
	if o == nil || utils.IsNil(o.imagePullPolicy) {
		var ret string
		return ret
	}
	return *o.imagePullPolicy
}

// GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSpec) GetImagePullPolicyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.imagePullPolicy) {
		return nil, false
	}
	return o.imagePullPolicy, true
}

// HasImagePullPolicy returns a boolean if a field has been set.
func (o *TaskSpec) HasImagePullPolicy() bool {
	if o != nil && !utils.IsNil(o.imagePullPolicy) {
		return true
	}

	return false
}

// SetImagePullPolicy gets a reference to the given string and assigns it to the imagePullPolicy field.
// imagePullPolicy:  Specify image pull policy for your service 

func (o *TaskSpec) ImagePullPolicy(v string) (*TaskSpec){
	o.imagePullPolicy = &v
return o
}

// GetImagePullSecrets returns the ImagePullSecrets field value if set, zero value otherwise.
func (o *TaskSpec) GetImagePullSecrets() []string {
	if o == nil || utils.IsNil(o.imagePullSecrets) {
		var ret []string
		return ret
	}
	return o.imagePullSecrets
}

// GetImagePullSecretsOk returns a tuple with the ImagePullSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSpec) GetImagePullSecretsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.imagePullSecrets) {
		return nil, false
	}
	return o.imagePullSecrets, true
}

// HasImagePullSecrets returns a boolean if a field has been set.
func (o *TaskSpec) HasImagePullSecrets() bool {
	if o != nil && !utils.IsNil(o.imagePullSecrets) {
		return true
	}

	return false
}

// SetImagePullSecrets gets a reference to the given []string and assigns it to the imagePullSecrets field.
// imagePullSecrets:  Specify image pull secrets for your service 

func (o *TaskSpec) ImagePullSecrets(v []string) (*TaskSpec){
	o.imagePullSecrets = v
return o
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *TaskSpec) GetLabels() map[string]string {
	if o == nil || utils.IsNil(o.labels) {
		var ret map[string]string
		return ret
	}
	return *o.labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSpec) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.labels) {
		return nil, false
	}
	return o.labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *TaskSpec) HasLabels() bool {
	if o != nil && !utils.IsNil(o.labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the labels field.
// labels:  Specify the labels in the workload 

func (o *TaskSpec) Labels(v map[string]string) (*TaskSpec){
	o.labels = &v
return o
}

// GetLivenessProbe returns the LivenessProbe field value if set, zero value otherwise.
func (o *TaskSpec) GetLivenessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.livenessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.livenessProbe
}

// GetLivenessProbeOk returns a tuple with the LivenessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSpec) GetLivenessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.livenessProbe) {
		return nil, false
	}
	return o.livenessProbe, true
}

// HasLivenessProbe returns a boolean if a field has been set.
func (o *TaskSpec) HasLivenessProbe() bool {
	if o != nil && !utils.IsNil(o.livenessProbe) {
		return true
	}

	return false
}

// SetLivenessProbe gets a reference to the given HealthProbe and assigns it to the livenessProbe field.
// livenessProbe: 

func (o *TaskSpec) LivenessProbe(v HealthProbe) (*TaskSpec){
	o.livenessProbe = &v
return o
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *TaskSpec) GetMemory() string {
	if o == nil || utils.IsNil(o.memory) {
		var ret string
		return ret
	}
	return *o.memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSpec) GetMemoryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.memory) {
		return nil, false
	}
	return o.memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *TaskSpec) HasMemory() bool {
	if o != nil && !utils.IsNil(o.memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given string and assigns it to the memory field.
// memory:  Specifies the attributes of the memory resource required for the container. 

func (o *TaskSpec) Memory(v string) (*TaskSpec){
	o.memory = &v
return o
}

// GetReadinessProbe returns the ReadinessProbe field value if set, zero value otherwise.
func (o *TaskSpec) GetReadinessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.readinessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.readinessProbe
}

// GetReadinessProbeOk returns a tuple with the ReadinessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSpec) GetReadinessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.readinessProbe) {
		return nil, false
	}
	return o.readinessProbe, true
}

// HasReadinessProbe returns a boolean if a field has been set.
func (o *TaskSpec) HasReadinessProbe() bool {
	if o != nil && !utils.IsNil(o.readinessProbe) {
		return true
	}

	return false
}

// SetReadinessProbe gets a reference to the given HealthProbe and assigns it to the readinessProbe field.
// readinessProbe: 

func (o *TaskSpec) ReadinessProbe(v HealthProbe) (*TaskSpec){
	o.readinessProbe = &v
return o
}

// GetRestart returns the Restart field value
func (o *TaskSpec) GetRestart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.restart
}

// GetRestartOk returns a tuple with the Restart field value
// and a boolean to check if the value has been set.
func (o *TaskSpec) GetRestartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.restart, true
}

// Restart sets field value
func (o *TaskSpec) Restart(v string) *TaskSpec {
	o.restart = v
    return o
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *TaskSpec) GetVolumes() []TaskSpecVolumes {
	if o == nil || utils.IsNil(o.volumes) {
		var ret []TaskSpecVolumes
		return ret
	}
	return o.volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSpec) GetVolumesOk() ([]TaskSpecVolumes, bool) {
	if o == nil || utils.IsNil(o.volumes) {
		return nil, false
	}
	return o.volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *TaskSpec) HasVolumes() bool {
	if o != nil && !utils.IsNil(o.volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []TaskSpecVolumes and assigns it to the volumes field.
// volumes:  Declare volumes and volumeMounts 

func (o *TaskSpec) Volumes(v []TaskSpecVolumes) (*TaskSpec){
	o.volumes = v
return o
}

func (o TaskSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.annotations) {
		toSerialize["annotations"] = o.annotations
	}
	if !utils.IsNil(o.cmd) {
		toSerialize["cmd"] = o.cmd
	}
	toSerialize["count"] = o.count
	if !utils.IsNil(o.cpu) {
		toSerialize["cpu"] = o.cpu
	}
	if !utils.IsNil(o.env) {
		toSerialize["env"] = o.env
	}
	toSerialize["image"] = o.image
	if !utils.IsNil(o.imagePullPolicy) {
		toSerialize["imagePullPolicy"] = o.imagePullPolicy
	}
	if !utils.IsNil(o.imagePullSecrets) {
		toSerialize["imagePullSecrets"] = o.imagePullSecrets
	}
	if !utils.IsNil(o.labels) {
		toSerialize["labels"] = o.labels
	}
	if !utils.IsNil(o.livenessProbe) {
		toSerialize["livenessProbe"] = o.livenessProbe
	}
	if !utils.IsNil(o.memory) {
		toSerialize["memory"] = o.memory
	}
	if !utils.IsNil(o.readinessProbe) {
		toSerialize["readinessProbe"] = o.readinessProbe
	}
	toSerialize["restart"] = o.restart
	if !utils.IsNil(o.volumes) {
		toSerialize["volumes"] = o.volumes
	}
	return toSerialize, nil
}

type NullableTaskSpec struct {
	value *TaskSpec
	isSet bool
}

func (v NullableTaskSpec) Get() *TaskSpec {
	return v.value
}

func (v *NullableTaskSpec) Set(val *TaskSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskSpec(val *TaskSpec) *NullableTaskSpec {
	return &NullableTaskSpec{value: val, isSet: true}
}

func (v NullableTaskSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 
