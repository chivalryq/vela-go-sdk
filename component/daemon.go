/*
Copyright 2023 The KubeVela Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated from  using `vela def gen-api`. DO NOT EDIT.

package component

import (
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	. "vela-go-sdk/api"
)

const DaemonType = "daemon"

// DaemonComponent is the root struct of daemon
type DaemonComponent struct {
	Base  ComponentBase
	Props DaemonSpec
}

// ConfigMap -
type ConfigMap struct {
	Name        string `json:"name"`
	MountPath   string `json:"mountPath"`
	DefaultMode int    `json:"defaultMode"`
	CmName      string `json:"cmName"`
	Items       Items  `json:"items,omitempty"`
}

// Secret -
type Secret struct {
	Name        string `json:"name"`
	MountPath   string `json:"mountPath"`
	DefaultMode int    `json:"defaultMode"`
	SecretName  string `json:"secretName"`
	Items       Items  `json:"items,omitempty"`
}

// HostPath -
type HostPath struct {
	Name             string `json:"name"`
	MountPath        string `json:"mountPath"`
	MountPropagation string `json:"mountPropagation,omitempty"`
	Path             string `json:"path"`
	ReadOnly         bool   `json:"readOnly,omitempty"`
}

// HTTPHeaders -
type HTTPHeaders struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// HTTPGet Instructions for assessing container health by executing an HTTP GET request. Either this attribute or the exec attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with both the exec attribute and the tcpSocket attribute.
type HTTPGet struct {
	Path        string      `json:"path"`
	Port        int         `json:"port"`
	Host        string      `json:"host,omitempty"`
	Scheme      string      `json:"scheme,omitempty"`
	HTTPHeaders HTTPHeaders `json:"httpHeaders,omitempty"`
}

// SecretKeyRef Selects a key of a secret in the pod's namespace
type SecretKeyRef struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// ValueFrom Specifies a source the value of this var should come from
type ValueFrom struct {
	SecretKeyRef    SecretKeyRef    `json:"secretKeyRef,omitempty"`
	ConfigMapKeyRef ConfigMapKeyRef `json:"configMapKeyRef,omitempty"`
}

// Env -
type Env struct {
	Name      string    `json:"name"`
	Value     string    `json:"value,omitempty"`
	ValueFrom ValueFrom `json:"valueFrom,omitempty"`
}

// Exec Instructions for assessing container health by executing a command. Either this attribute or the httpGet attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with both the httpGet attribute and the tcpSocket attribute.
type Exec struct {
	Command []string `json:"command"`
}

// HealthProbe -
type HealthProbe struct {
	Exec                Exec      `json:"exec,omitempty"`
	HTTPGet             HTTPGet   `json:"httpGet,omitempty"`
	TcpSocket           TcpSocket `json:"tcpSocket,omitempty"`
	InitialDelaySeconds int       `json:"initialDelaySeconds"`
	PeriodSeconds       int       `json:"periodSeconds"`
	TimeoutSeconds      int       `json:"timeoutSeconds"`
	SuccessThreshold    int       `json:"successThreshold"`
	FailureThreshold    int       `json:"failureThreshold"`
}

// Ports -
type Ports struct {
	Port     int    `json:"port"`
	Name     string `json:"name,omitempty"`
	Protocol string `json:"protocol"`
	Expose   bool   `json:"expose"`
}

// ConfigMapKeyRef Selects a key of a config map in the pod's namespace
type ConfigMapKeyRef struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// EmptyDir -
type EmptyDir struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
	Medium    string `json:"medium"`
}

// DaemonSpec -
type DaemonSpec struct {
	Labels           map[string]string `json:"labels,omitempty"`
	Annotations      map[string]string `json:"annotations,omitempty"`
	Image            string            `json:"image"`
	ImagePullPolicy  string            `json:"imagePullPolicy,omitempty"`
	ImagePullSecrets []string          `json:"imagePullSecrets,omitempty"`
	Port             int               `json:"port,omitempty"`
	Ports            Ports             `json:"ports,omitempty"`
	ExposeType       string            `json:"exposeType"`
	AddRevisionLabel bool              `json:"addRevisionLabel"`
	Cmd              []string          `json:"cmd,omitempty"`
	Env              Env               `json:"env,omitempty"`
	CPU              string            `json:"cpu,omitempty"`
	Memory           string            `json:"memory,omitempty"`
	VolumeMounts     VolumeMounts      `json:"volumeMounts,omitempty"`
	Volumes          Volumes           `json:"volumes,omitempty"`
	LivenessProbe    HealthProbe       `json:"livenessProbe,omitempty"`
	ReadinessProbe   HealthProbe       `json:"readinessProbe,omitempty"`
	HostAliases      HostAliases       `json:"hostAliases,omitempty"`
}

// TcpSocket Instructions for assessing container health by probing a TCP socket. Either this attribute or the exec attribute or the httpGet attribute MUST be specified. This attribute is mutually exclusive with both the exec attribute and the httpGet attribute.
type TcpSocket struct {
	Port int `json:"port"`
}

// PVC -
type PVC struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
	ClaimName string `json:"claimName"`
}

// Items -
type Items struct {
	Key  string `json:"key"`
	Path string `json:"path"`
	Mode int    `json:"mode"`
}

// VolumeMounts -
type VolumeMounts struct {
	PVC       PVC       `json:"pvc,omitempty"`
	ConfigMap ConfigMap `json:"configMap,omitempty"`
	Secret    Secret    `json:"secret,omitempty"`
	EmptyDir  EmptyDir  `json:"emptyDir,omitempty"`
	HostPath  HostPath  `json:"hostPath,omitempty"`
}

// Volumes -
type Volumes struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
	Medium    string `json:"medium"`
	Type      string `json:"type"`
}

// HostAliases -
type HostAliases struct {
	IP        string   `json:"ip"`
	Hostnames []string `json:"hostnames"`
}

func Daemon(name string) *DaemonComponent {
	comp := DaemonComponent{
		Base: ComponentBase{
			Name: name,
		},
	}
	return &comp
}

func (d *DaemonComponent) Trait(traits ...Trait) *DaemonComponent {
	d.Base.Traits = append(d.Base.Traits, traits...)
	return d
}

func (d *DaemonComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range d.Base.Traits {
		traits = append(traits, trait.Build())
	}
	comp := common.ApplicationComponent{
		Name:       d.Base.Name,
		Type:       DaemonType,
		Properties: util.Object2RawExtension(d.Props),
		DependsOn:  d.Base.DependsOn,
		Inputs:     d.Base.Inputs,
		Outputs:    d.Base.Outputs,
		Traits:     traits,
	}
	return comp
}

// Labels Specify the labels in the workload
func (d *DaemonComponent) Labels(value map[string]string) *DaemonComponent {
	d.Props.Labels = value
	return d
}

// Annotations Specify the annotations in the workload
func (d *DaemonComponent) Annotations(value map[string]string) *DaemonComponent {
	d.Props.Annotations = value
	return d
}

// Image Which image would you like to use for your service
func (d *DaemonComponent) Image(value string) *DaemonComponent {
	d.Props.Image = value
	return d
}

// ImagePullPolicy Specify image pull policy for your service
func (d *DaemonComponent) ImagePullPolicy(value string) *DaemonComponent {
	d.Props.ImagePullPolicy = value
	return d
}

// ImagePullSecrets Specify image pull secrets for your service
func (d *DaemonComponent) ImagePullSecrets(value []string) *DaemonComponent {
	d.Props.ImagePullSecrets = value
	return d
}

// Port Deprecated field, please use ports instead
func (d *DaemonComponent) Port(value int) *DaemonComponent {
	d.Props.Port = value
	return d
}

// Ports Which ports do you want customer traffic sent to, defaults to 80
func (d *DaemonComponent) Ports(value Ports) *DaemonComponent {
	d.Props.Ports = value
	return d
}

// ExposeType Specify what kind of Service you want. options: "ClusterIP", "NodePort", "LoadBalancer", "ExternalName"
func (d *DaemonComponent) ExposeType(value string) *DaemonComponent {
	d.Props.ExposeType = value
	return d
}

// AddRevisionLabel If addRevisionLabel is true, the revision label will be added to the underlying pods
func (d *DaemonComponent) AddRevisionLabel(value bool) *DaemonComponent {
	d.Props.AddRevisionLabel = value
	return d
}

// Cmd Commands to run in the container
func (d *DaemonComponent) Cmd(value []string) *DaemonComponent {
	d.Props.Cmd = value
	return d
}

// Env Define arguments by using environment variables
func (d *DaemonComponent) Env(value Env) *DaemonComponent {
	d.Props.Env = value
	return d
}

// CPU Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
func (d *DaemonComponent) CPU(value string) *DaemonComponent {
	d.Props.CPU = value
	return d
}

// Memory Specifies the attributes of the memory resource required for the container.
func (d *DaemonComponent) Memory(value string) *DaemonComponent {
	d.Props.Memory = value
	return d
}

// VolumeMounts -
func (d *DaemonComponent) VolumeMounts(value VolumeMounts) *DaemonComponent {
	d.Props.VolumeMounts = value
	return d
}

// Volumes Deprecated field, use volumeMounts instead.
func (d *DaemonComponent) Volumes(value Volumes) *DaemonComponent {
	d.Props.Volumes = value
	return d
}

// LivenessProbe Instructions for assessing whether the container is alive.
func (d *DaemonComponent) LivenessProbe(value HealthProbe) *DaemonComponent {
	d.Props.LivenessProbe = value
	return d
}

// ReadinessProbe Instructions for assessing whether the container is in a suitable state to serve traffic.
func (d *DaemonComponent) ReadinessProbe(value HealthProbe) *DaemonComponent {
	d.Props.ReadinessProbe = value
	return d
}

// HostAliases Specify the hostAliases to add
func (d *DaemonComponent) HostAliases(value HostAliases) *DaemonComponent {
	d.Props.HostAliases = value
	return d
}
