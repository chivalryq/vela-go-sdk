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

const WorkerType = "worker"

// WorkerComponent is the root struct of worker
type WorkerComponent struct {
	Base  ComponentBase
	Props WorkerSpec
}

// Secret -
type Secret struct {
	Name        string `json:"name"`
	MountPath   string `json:"mountPath"`
	DefaultMode int    `json:"defaultMode"`
	SecretName  string `json:"secretName"`
	Items       Items  `json:"items,omitempty"`
}

// Volumes -
type Volumes struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
	Medium    string `json:"medium"`
	Type      string `json:"type"`
}

// HTTPGet Instructions for assessing container health by executing an HTTP GET request. Either this attribute or the exec attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with both the exec attribute and the tcpSocket attribute.
type HTTPGet struct {
	Path        string      `json:"path"`
	Port        int         `json:"port"`
	HTTPHeaders HTTPHeaders `json:"httpHeaders,omitempty"`
}

// ConfigMap -
type ConfigMap struct {
	Name        string `json:"name"`
	MountPath   string `json:"mountPath"`
	DefaultMode int    `json:"defaultMode"`
	CmName      string `json:"cmName"`
	Items       Items  `json:"items,omitempty"`
}

// WorkerSpec -
type WorkerSpec struct {
	Image            string       `json:"image"`
	ImagePullPolicy  string       `json:"imagePullPolicy,omitempty"`
	ImagePullSecrets []string     `json:"imagePullSecrets,omitempty"`
	Cmd              []string     `json:"cmd,omitempty"`
	Env              Env          `json:"env,omitempty"`
	CPU              string       `json:"cpu,omitempty"`
	Memory           string       `json:"memory,omitempty"`
	VolumeMounts     VolumeMounts `json:"volumeMounts,omitempty"`
	Volumes          Volumes      `json:"volumes,omitempty"`
	LivenessProbe    HealthProbe  `json:"livenessProbe,omitempty"`
	ReadinessProbe   HealthProbe  `json:"readinessProbe,omitempty"`
}

// Exec Instructions for assessing container health by executing a command. Either this attribute or the httpGet attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with both the httpGet attribute and the tcpSocket attribute.
type Exec struct {
	Command []string `json:"command"`
}

// SecretKeyRef Selects a key of a secret in the pod's namespace
type SecretKeyRef struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// ConfigMapKeyRef Selects a key of a config map in the pod's namespace
type ConfigMapKeyRef struct {
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

// Items -
type Items struct {
	Key  string `json:"key"`
	Path string `json:"path"`
	Mode int    `json:"mode"`
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

// VolumeMounts -
type VolumeMounts struct {
	PVC       PVC       `json:"pvc,omitempty"`
	ConfigMap ConfigMap `json:"configMap,omitempty"`
	Secret    Secret    `json:"secret,omitempty"`
	EmptyDir  EmptyDir  `json:"emptyDir,omitempty"`
	HostPath  HostPath  `json:"hostPath,omitempty"`
}

// HTTPHeaders -
type HTTPHeaders struct {
	Name  string `json:"name"`
	Value string `json:"value"`
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

// EmptyDir -
type EmptyDir struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
	Medium    string `json:"medium"`
}

// HostPath -
type HostPath struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
	Path      string `json:"path"`
}

func Worker(name string) *WorkerComponent {
	comp := WorkerComponent{
		Base: ComponentBase{
			Name: name,
		},
	}
	return &comp
}

func (w *WorkerComponent) Trait(traits ...Trait) *WorkerComponent {
	w.Base.Traits = append(w.Base.Traits, traits...)
	return w
}

func (w *WorkerComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range w.Base.Traits {
		traits = append(traits, trait.Build())
	}
	comp := common.ApplicationComponent{
		Name:       w.Base.Name,
		Type:       WorkerType,
		Properties: util.Object2RawExtension(w.Props),
		DependsOn:  w.Base.DependsOn,
		Inputs:     w.Base.Inputs,
		Outputs:    w.Base.Outputs,
		Traits:     traits,
	}
	return comp
}

// Image Which image would you like to use for your service
func (w *WorkerComponent) Image(value string) *WorkerComponent {
	w.Props.Image = value
	return w
}

// ImagePullPolicy Specify image pull policy for your service
func (w *WorkerComponent) ImagePullPolicy(value string) *WorkerComponent {
	w.Props.ImagePullPolicy = value
	return w
}

// ImagePullSecrets Specify image pull secrets for your service
func (w *WorkerComponent) ImagePullSecrets(value []string) *WorkerComponent {
	w.Props.ImagePullSecrets = value
	return w
}

// Cmd Commands to run in the container
func (w *WorkerComponent) Cmd(value []string) *WorkerComponent {
	w.Props.Cmd = value
	return w
}

// Env Define arguments by using environment variables
func (w *WorkerComponent) Env(value Env) *WorkerComponent {
	w.Props.Env = value
	return w
}

// CPU Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
func (w *WorkerComponent) CPU(value string) *WorkerComponent {
	w.Props.CPU = value
	return w
}

// Memory Specifies the attributes of the memory resource required for the container.
func (w *WorkerComponent) Memory(value string) *WorkerComponent {
	w.Props.Memory = value
	return w
}

// VolumeMounts -
func (w *WorkerComponent) VolumeMounts(value VolumeMounts) *WorkerComponent {
	w.Props.VolumeMounts = value
	return w
}

// Volumes Deprecated field, use volumeMounts instead.
func (w *WorkerComponent) Volumes(value Volumes) *WorkerComponent {
	w.Props.Volumes = value
	return w
}

// LivenessProbe Instructions for assessing whether the container is alive.
func (w *WorkerComponent) LivenessProbe(value HealthProbe) *WorkerComponent {
	w.Props.LivenessProbe = value
	return w
}

// ReadinessProbe Instructions for assessing whether the container is in a suitable state to serve traffic.
func (w *WorkerComponent) ReadinessProbe(value HealthProbe) *WorkerComponent {
	w.Props.ReadinessProbe = value
	return w
}
