/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package daemon

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the DaemonSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DaemonSpec{}

// DaemonSpec struct for DaemonSpec
type DaemonSpec struct {
	AddRevisionLabel *bool `json:"addRevisionLabel,omitempty"`
	// Specify the annotations in the workload
	Annotations *map[string]string `json:"annotations,omitempty"`
	// Commands to run in the container
	Cmd []string `json:"cmd,omitempty"`
	// Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
	Cpu *string `json:"cpu,omitempty"`
	// Define arguments by using environment variables
	Env        []Env   `json:"env,omitempty"`
	ExposeType *string `json:"exposeType,omitempty"`
	// Specify the hostAliases to add
	HostAliases []HostAliases `json:"hostAliases,omitempty"`
	// Which image would you like to use for your service +short=i
	Image *string `json:"image,omitempty"`
	// Specify image pull policy for your service
	ImagePullPolicy *string `json:"imagePullPolicy,omitempty"`
	// Specify image pull secrets for your service
	ImagePullSecrets []string `json:"imagePullSecrets,omitempty"`
	// Specify the labels in the workload
	Labels        *map[string]string `json:"labels,omitempty"`
	LivenessProbe *HealthProbe       `json:"livenessProbe,omitempty"`
	// Specifies the attributes of the memory resource required for the container.
	Memory *string `json:"memory,omitempty"`
	Port   *int32  `json:"port,omitempty"`
	// Which ports do you want customer traffic sent to, defaults to 80
	Ports          []Ports       `json:"ports,omitempty"`
	ReadinessProbe *HealthProbe  `json:"readinessProbe,omitempty"`
	VolumeMounts   *VolumeMounts `json:"volumeMounts,omitempty"`
	// Deprecated field, use volumeMounts instead.
	Volumes []Volumes `json:"volumes,omitempty"`
}

// NewDaemonSpecWith instantiates a new DaemonSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDaemonSpecWith() *DaemonSpec {
	this := DaemonSpec{}
	var addRevisionLabel bool = false
	this.AddRevisionLabel = &addRevisionLabel
	var exposeType string = "ClusterIP"
	this.ExposeType = &exposeType
	return &this
}

// NewDaemonSpec instantiates a new DaemonSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDaemonSpec() *DaemonSpec {
	this := DaemonSpec{}
	var addRevisionLabel bool = false
	this.AddRevisionLabel = &addRevisionLabel
	var exposeType string = "ClusterIP"
	this.ExposeType = &exposeType
	return &this
}

// NewDaemonSpecs converts a list DaemonSpec pointers to objects.
// This is helpful when the SetDaemonSpec requires a list of objects
func NewDaemonSpecs(ps ...*DaemonSpec) []DaemonSpec {
	objs := []DaemonSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetAddRevisionLabel returns the AddRevisionLabel field value if set, zero value otherwise.
func (o *DaemonComponent) GetAddRevisionLabel() bool {
	if o == nil || utils.IsNil(o.Properties.AddRevisionLabel) {
		var ret bool
		return ret
	}
	return *o.Properties.AddRevisionLabel
}

// GetAddRevisionLabelOk returns a tuple with the AddRevisionLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetAddRevisionLabelOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Properties.AddRevisionLabel) {
		return nil, false
	}
	return o.Properties.AddRevisionLabel, true
}

// HasAddRevisionLabel returns a boolean if a field has been set.
func (o *DaemonComponent) HasAddRevisionLabel() bool {
	if o != nil && !utils.IsNil(o.Properties.AddRevisionLabel) {
		return true
	}

	return false
}

// SetAddRevisionLabel gets a reference to the given bool and assigns it to the addRevisionLabel field.
// AddRevisionLabel:
func (o *DaemonComponent) SetAddRevisionLabel(v bool) *DaemonComponent {
	o.Properties.AddRevisionLabel = &v
	return o
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *DaemonComponent) GetAnnotations() map[string]string {
	if o == nil || utils.IsNil(o.Properties.Annotations) {
		var ret map[string]string
		return ret
	}
	return *o.Properties.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Annotations) {
		return nil, false
	}
	return o.Properties.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *DaemonComponent) HasAnnotations() bool {
	if o != nil && !utils.IsNil(o.Properties.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the annotations field.
// Annotations:  Specify the annotations in the workload
func (o *DaemonComponent) SetAnnotations(v map[string]string) *DaemonComponent {
	o.Properties.Annotations = &v
	return o
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *DaemonComponent) GetCmd() []string {
	if o == nil || utils.IsNil(o.Properties.Cmd) {
		var ret []string
		return ret
	}
	return o.Properties.Cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetCmdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Cmd) {
		return nil, false
	}
	return o.Properties.Cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *DaemonComponent) HasCmd() bool {
	if o != nil && !utils.IsNil(o.Properties.Cmd) {
		return true
	}

	return false
}

// SetCmd gets a reference to the given []string and assigns it to the cmd field.
// Cmd:  Commands to run in the container
func (o *DaemonComponent) SetCmd(v []string) *DaemonComponent {
	o.Properties.Cmd = v
	return o
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *DaemonComponent) GetCpu() string {
	if o == nil || utils.IsNil(o.Properties.Cpu) {
		var ret string
		return ret
	}
	return *o.Properties.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetCpuOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Cpu) {
		return nil, false
	}
	return o.Properties.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *DaemonComponent) HasCpu() bool {
	if o != nil && !utils.IsNil(o.Properties.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given string and assigns it to the cpu field.
// Cpu:  Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
func (o *DaemonComponent) SetCpu(v string) *DaemonComponent {
	o.Properties.Cpu = &v
	return o
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *DaemonComponent) GetEnv() []Env {
	if o == nil || utils.IsNil(o.Properties.Env) {
		var ret []Env
		return ret
	}
	return o.Properties.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetEnvOk() ([]Env, bool) {
	if o == nil || utils.IsNil(o.Properties.Env) {
		return nil, false
	}
	return o.Properties.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *DaemonComponent) HasEnv() bool {
	if o != nil && !utils.IsNil(o.Properties.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []Env and assigns it to the env field.
// Env:  Define arguments by using environment variables
func (o *DaemonComponent) SetEnv(v []Env) *DaemonComponent {
	o.Properties.Env = v
	return o
}

// GetExposeType returns the ExposeType field value if set, zero value otherwise.
func (o *DaemonComponent) GetExposeType() string {
	if o == nil || utils.IsNil(o.Properties.ExposeType) {
		var ret string
		return ret
	}
	return *o.Properties.ExposeType
}

// GetExposeTypeOk returns a tuple with the ExposeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetExposeTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.ExposeType) {
		return nil, false
	}
	return o.Properties.ExposeType, true
}

// HasExposeType returns a boolean if a field has been set.
func (o *DaemonComponent) HasExposeType() bool {
	if o != nil && !utils.IsNil(o.Properties.ExposeType) {
		return true
	}

	return false
}

// SetExposeType gets a reference to the given string and assigns it to the exposeType field.
// ExposeType:
func (o *DaemonComponent) SetExposeType(v string) *DaemonComponent {
	o.Properties.ExposeType = &v
	return o
}

// GetHostAliases returns the HostAliases field value if set, zero value otherwise.
func (o *DaemonComponent) GetHostAliases() []HostAliases {
	if o == nil || utils.IsNil(o.Properties.HostAliases) {
		var ret []HostAliases
		return ret
	}
	return o.Properties.HostAliases
}

// GetHostAliasesOk returns a tuple with the HostAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetHostAliasesOk() ([]HostAliases, bool) {
	if o == nil || utils.IsNil(o.Properties.HostAliases) {
		return nil, false
	}
	return o.Properties.HostAliases, true
}

// HasHostAliases returns a boolean if a field has been set.
func (o *DaemonComponent) HasHostAliases() bool {
	if o != nil && !utils.IsNil(o.Properties.HostAliases) {
		return true
	}

	return false
}

// SetHostAliases gets a reference to the given []HostAliases and assigns it to the hostAliases field.
// HostAliases:  Specify the hostAliases to add
func (o *DaemonComponent) SetHostAliases(v []HostAliases) *DaemonComponent {
	o.Properties.HostAliases = v
	return o
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *DaemonComponent) GetImage() string {
	if o == nil || utils.IsNil(o.Properties.Image) {
		var ret string
		return ret
	}
	return *o.Properties.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetImageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Image) {
		return nil, false
	}
	return o.Properties.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *DaemonComponent) HasImage() bool {
	if o != nil && !utils.IsNil(o.Properties.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the image field.
// Image:  Which image would you like to use for your service +short=i
func (o *DaemonComponent) SetImage(v string) *DaemonComponent {
	o.Properties.Image = &v
	return o
}

// GetImagePullPolicy returns the ImagePullPolicy field value if set, zero value otherwise.
func (o *DaemonComponent) GetImagePullPolicy() string {
	if o == nil || utils.IsNil(o.Properties.ImagePullPolicy) {
		var ret string
		return ret
	}
	return *o.Properties.ImagePullPolicy
}

// GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetImagePullPolicyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.ImagePullPolicy) {
		return nil, false
	}
	return o.Properties.ImagePullPolicy, true
}

// HasImagePullPolicy returns a boolean if a field has been set.
func (o *DaemonComponent) HasImagePullPolicy() bool {
	if o != nil && !utils.IsNil(o.Properties.ImagePullPolicy) {
		return true
	}

	return false
}

// SetImagePullPolicy gets a reference to the given string and assigns it to the imagePullPolicy field.
// ImagePullPolicy:  Specify image pull policy for your service
func (o *DaemonComponent) SetImagePullPolicy(v string) *DaemonComponent {
	o.Properties.ImagePullPolicy = &v
	return o
}

// GetImagePullSecrets returns the ImagePullSecrets field value if set, zero value otherwise.
func (o *DaemonComponent) GetImagePullSecrets() []string {
	if o == nil || utils.IsNil(o.Properties.ImagePullSecrets) {
		var ret []string
		return ret
	}
	return o.Properties.ImagePullSecrets
}

// GetImagePullSecretsOk returns a tuple with the ImagePullSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetImagePullSecretsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.ImagePullSecrets) {
		return nil, false
	}
	return o.Properties.ImagePullSecrets, true
}

// HasImagePullSecrets returns a boolean if a field has been set.
func (o *DaemonComponent) HasImagePullSecrets() bool {
	if o != nil && !utils.IsNil(o.Properties.ImagePullSecrets) {
		return true
	}

	return false
}

// SetImagePullSecrets gets a reference to the given []string and assigns it to the imagePullSecrets field.
// ImagePullSecrets:  Specify image pull secrets for your service
func (o *DaemonComponent) SetImagePullSecrets(v []string) *DaemonComponent {
	o.Properties.ImagePullSecrets = v
	return o
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *DaemonComponent) GetLabels() map[string]string {
	if o == nil || utils.IsNil(o.Properties.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Properties.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Labels) {
		return nil, false
	}
	return o.Properties.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *DaemonComponent) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Properties.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the labels field.
// Labels:  Specify the labels in the workload
func (o *DaemonComponent) SetLabels(v map[string]string) *DaemonComponent {
	o.Properties.Labels = &v
	return o
}

// GetLivenessProbe returns the LivenessProbe field value if set, zero value otherwise.
func (o *DaemonComponent) GetLivenessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.Properties.LivenessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.Properties.LivenessProbe
}

// GetLivenessProbeOk returns a tuple with the LivenessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetLivenessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.Properties.LivenessProbe) {
		return nil, false
	}
	return o.Properties.LivenessProbe, true
}

// HasLivenessProbe returns a boolean if a field has been set.
func (o *DaemonComponent) HasLivenessProbe() bool {
	if o != nil && !utils.IsNil(o.Properties.LivenessProbe) {
		return true
	}

	return false
}

// SetLivenessProbe gets a reference to the given HealthProbe and assigns it to the livenessProbe field.
// LivenessProbe:
func (o *DaemonComponent) SetLivenessProbe(v HealthProbe) *DaemonComponent {
	o.Properties.LivenessProbe = &v
	return o
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *DaemonComponent) GetMemory() string {
	if o == nil || utils.IsNil(o.Properties.Memory) {
		var ret string
		return ret
	}
	return *o.Properties.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetMemoryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Memory) {
		return nil, false
	}
	return o.Properties.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *DaemonComponent) HasMemory() bool {
	if o != nil && !utils.IsNil(o.Properties.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given string and assigns it to the memory field.
// Memory:  Specifies the attributes of the memory resource required for the container.
func (o *DaemonComponent) SetMemory(v string) *DaemonComponent {
	o.Properties.Memory = &v
	return o
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *DaemonComponent) GetPort() int32 {
	if o == nil || utils.IsNil(o.Properties.Port) {
		var ret int32
		return ret
	}
	return *o.Properties.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.Port) {
		return nil, false
	}
	return o.Properties.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *DaemonComponent) HasPort() bool {
	if o != nil && !utils.IsNil(o.Properties.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the port field.
// Port:
func (o *DaemonComponent) SetPort(v int32) *DaemonComponent {
	o.Properties.Port = &v
	return o
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *DaemonComponent) GetPorts() []Ports {
	if o == nil || utils.IsNil(o.Properties.Ports) {
		var ret []Ports
		return ret
	}
	return o.Properties.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetPortsOk() ([]Ports, bool) {
	if o == nil || utils.IsNil(o.Properties.Ports) {
		return nil, false
	}
	return o.Properties.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *DaemonComponent) HasPorts() bool {
	if o != nil && !utils.IsNil(o.Properties.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []Ports and assigns it to the ports field.
// Ports:  Which ports do you want customer traffic sent to, defaults to 80
func (o *DaemonComponent) SetPorts(v []Ports) *DaemonComponent {
	o.Properties.Ports = v
	return o
}

// GetReadinessProbe returns the ReadinessProbe field value if set, zero value otherwise.
func (o *DaemonComponent) GetReadinessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.Properties.ReadinessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.Properties.ReadinessProbe
}

// GetReadinessProbeOk returns a tuple with the ReadinessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetReadinessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.Properties.ReadinessProbe) {
		return nil, false
	}
	return o.Properties.ReadinessProbe, true
}

// HasReadinessProbe returns a boolean if a field has been set.
func (o *DaemonComponent) HasReadinessProbe() bool {
	if o != nil && !utils.IsNil(o.Properties.ReadinessProbe) {
		return true
	}

	return false
}

// SetReadinessProbe gets a reference to the given HealthProbe and assigns it to the readinessProbe field.
// ReadinessProbe:
func (o *DaemonComponent) SetReadinessProbe(v HealthProbe) *DaemonComponent {
	o.Properties.ReadinessProbe = &v
	return o
}

// GetVolumeMounts returns the VolumeMounts field value if set, zero value otherwise.
func (o *DaemonComponent) GetVolumeMounts() VolumeMounts {
	if o == nil || utils.IsNil(o.Properties.VolumeMounts) {
		var ret VolumeMounts
		return ret
	}
	return *o.Properties.VolumeMounts
}

// GetVolumeMountsOk returns a tuple with the VolumeMounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetVolumeMountsOk() (*VolumeMounts, bool) {
	if o == nil || utils.IsNil(o.Properties.VolumeMounts) {
		return nil, false
	}
	return o.Properties.VolumeMounts, true
}

// HasVolumeMounts returns a boolean if a field has been set.
func (o *DaemonComponent) HasVolumeMounts() bool {
	if o != nil && !utils.IsNil(o.Properties.VolumeMounts) {
		return true
	}

	return false
}

// SetVolumeMounts gets a reference to the given VolumeMounts and assigns it to the volumeMounts field.
// VolumeMounts:
func (o *DaemonComponent) SetVolumeMounts(v VolumeMounts) *DaemonComponent {
	o.Properties.VolumeMounts = &v
	return o
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *DaemonComponent) GetVolumes() []Volumes {
	if o == nil || utils.IsNil(o.Properties.Volumes) {
		var ret []Volumes
		return ret
	}
	return o.Properties.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonComponent) GetVolumesOk() ([]Volumes, bool) {
	if o == nil || utils.IsNil(o.Properties.Volumes) {
		return nil, false
	}
	return o.Properties.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *DaemonComponent) HasVolumes() bool {
	if o != nil && !utils.IsNil(o.Properties.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []Volumes and assigns it to the volumes field.
// Volumes:  Deprecated field, use volumeMounts instead.
func (o *DaemonComponent) SetVolumes(v []Volumes) *DaemonComponent {
	o.Properties.Volumes = v
	return o
}

func (o DaemonSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DaemonSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AddRevisionLabel) {
		toSerialize["addRevisionLabel"] = o.AddRevisionLabel
	}
	if !utils.IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !utils.IsNil(o.Cmd) {
		toSerialize["cmd"] = o.Cmd
	}
	if !utils.IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !utils.IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !utils.IsNil(o.ExposeType) {
		toSerialize["exposeType"] = o.ExposeType
	}
	if !utils.IsNil(o.HostAliases) {
		toSerialize["hostAliases"] = o.HostAliases
	}
	if !utils.IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !utils.IsNil(o.ImagePullPolicy) {
		toSerialize["imagePullPolicy"] = o.ImagePullPolicy
	}
	if !utils.IsNil(o.ImagePullSecrets) {
		toSerialize["imagePullSecrets"] = o.ImagePullSecrets
	}
	if !utils.IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !utils.IsNil(o.LivenessProbe) {
		toSerialize["livenessProbe"] = o.LivenessProbe
	}
	if !utils.IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !utils.IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !utils.IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !utils.IsNil(o.ReadinessProbe) {
		toSerialize["readinessProbe"] = o.ReadinessProbe
	}
	if !utils.IsNil(o.VolumeMounts) {
		toSerialize["volumeMounts"] = o.VolumeMounts
	}
	if !utils.IsNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	return toSerialize, nil
}

type NullableDaemonSpec struct {
	value *DaemonSpec
	isSet bool
}

func (v NullableDaemonSpec) Get() *DaemonSpec {
	return v.value
}

func (v *NullableDaemonSpec) Set(val *DaemonSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDaemonSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDaemonSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaemonSpec(val *DaemonSpec) *NullableDaemonSpec {
	return &NullableDaemonSpec{value: val, isSet: true}
}

func (v NullableDaemonSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaemonSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const DaemonType = "daemon"

func init() {
	sdkcommon.RegisterComponent(DaemonType, FromComponent)
}

type DaemonComponent struct {
	Base       apis.ComponentBase
	Properties DaemonSpec
}

func Daemon(name string) *DaemonComponent {
	d := &DaemonComponent{Base: apis.ComponentBase{
		Name: name,
		Type: DaemonType,
	}}
	return d
}

func (d *DaemonComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range d.Base.Traits {
		traits = append(traits, trait.Build())
	}
	res := common.ApplicationComponent{
		DependsOn:  d.Base.DependsOn,
		Inputs:     d.Base.Inputs,
		Name:       d.Base.Name,
		Outputs:    d.Base.Outputs,
		Properties: util.Object2RawExtension(d.Properties),
		Traits:     traits,
		Type:       DaemonType,
	}
	return res
}

func (d *DaemonComponent) FromComponent(from common.ApplicationComponent) (*DaemonComponent, error) {
	for _, trait := range from.Traits {
		_t, err := sdkcommon.FromTrait(&trait)
		if err != nil {
			return nil, err
		}
		d.Base.Traits = append(d.Base.Traits, _t)
	}
	var properties DaemonSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	d.Base.Name = from.Name
	d.Base.DependsOn = from.DependsOn
	d.Base.Inputs = from.Inputs
	d.Base.Outputs = from.Outputs
	d.Base.Type = DaemonType
	d.Properties = properties
	return d, nil
}

func FromComponent(from common.ApplicationComponent) (apis.Component, error) {
	d := &DaemonComponent{}
	return d.FromComponent(from)
}

func (d *DaemonComponent) AddTrait(traits ...apis.Trait) *DaemonComponent {
	d.Base.Traits = append(d.Base.Traits, traits...)
	return d
}

func (d *DaemonComponent) GetTrait(typ string) apis.Trait {
	for _, _t := range d.Base.Traits {
		if _t.DefType() == typ {
			return _t
		}
	}
	return nil
}

func (d *DaemonComponent) ComponentName() string {
	return d.Base.Name
}

func (d *DaemonComponent) DefType() string {
	return DaemonType
}

func (d *DaemonComponent) DependsOn(dependsOn []string) *DaemonComponent {
	d.Base.DependsOn = dependsOn
	return d
}

func (d *DaemonComponent) Inputs(input common.StepInputs) *DaemonComponent {
	d.Base.Inputs = input
	return d
}

func (d *DaemonComponent) Outputs(output common.StepOutputs) *DaemonComponent {
	d.Base.Outputs = output
	return d
}

func (d *DaemonComponent) AddDependsOn(dependsOn string) *DaemonComponent {
	d.Base.DependsOn = append(d.Base.DependsOn, dependsOn)
	return d
}
