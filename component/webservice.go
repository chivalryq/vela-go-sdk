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

const WebserviceType = "webservice"

// WebserviceComponent is the root struct of webservice
type WebserviceComponent struct {
	Base  ComponentBase
	Props WebserviceSpec
}

// PVC -
type PVC struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
	SubPath   string `json:"subPath,omitempty"`
	ClaimName string `json:"claimName"`
}

// EmptyDir -
type EmptyDir struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
	SubPath   string `json:"subPath,omitempty"`
	Medium    string `json:"medium"`
}

// VolumeMounts -
type VolumeMounts struct {
	PVC       PVC       `json:"pvc,omitempty"`
	ConfigMap ConfigMap `json:"configMap,omitempty"`
	Secret    Secret    `json:"secret,omitempty"`
	EmptyDir  EmptyDir  `json:"emptyDir,omitempty"`
	HostPath  HostPath  `json:"hostPath,omitempty"`
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
	Host        string      `json:"host,omitempty"`
	Scheme      string      `json:"scheme,omitempty"`
	HTTPHeaders HTTPHeaders `json:"httpHeaders,omitempty"`
}

// HostAliases -
type HostAliases struct {
	IP        string   `json:"ip"`
	Hostnames []string `json:"hostnames"`
}

// WebserviceSpec -
type WebserviceSpec struct {
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
	Args             []string          `json:"args,omitempty"`
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

// ConfigMapKeyRef Selects a key of a config map in the pod's namespace
type ConfigMapKeyRef struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// Items -
type Items struct {
	Key  string `json:"key"`
	Path string `json:"path"`
	Mode int    `json:"mode"`
}

// Secret -
type Secret struct {
	Name        string `json:"name"`
	MountPath   string `json:"mountPath"`
	SubPath     string `json:"subPath,omitempty"`
	DefaultMode int    `json:"defaultMode"`
	SecretName  string `json:"secretName"`
	Items       Items  `json:"items,omitempty"`
}

// HostPath -
type HostPath struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
	SubPath   string `json:"subPath,omitempty"`
	Path      string `json:"path"`
}

// HTTPHeaders -
type HTTPHeaders struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Ports -
type Ports struct {
	Port     int    `json:"port"`
	Name     string `json:"name,omitempty"`
	Protocol string `json:"protocol"`
	Expose   bool   `json:"expose"`
	NodePort int    `json:"nodePort,omitempty"`
}

// SecretKeyRef Selects a key of a secret in the pod's namespace
type SecretKeyRef struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// ConfigMap -
type ConfigMap struct {
	Name        string `json:"name"`
	MountPath   string `json:"mountPath"`
	SubPath     string `json:"subPath,omitempty"`
	DefaultMode int    `json:"defaultMode"`
	CmName      string `json:"cmName"`
	Items       Items  `json:"items,omitempty"`
}

// Exec Instructions for assessing container health by executing a command. Either this attribute or the httpGet attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with both the httpGet attribute and the tcpSocket attribute.
type Exec struct {
	Command []string `json:"command"`
}

func Webservice(name string) *WebserviceComponent {
	comp := WebserviceComponent{
		Base: ComponentBase{
			Name: name,
		},
	}
	return &comp
}

func (w *WebserviceComponent) Trait(traits ...Trait) *WebserviceComponent {
	w.Base.Traits = append(w.Base.Traits, traits...)
	return w
}

func (w *WebserviceComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range w.Base.Traits {
		traits = append(traits, trait.Build())
	}
	comp := common.ApplicationComponent{
		Name:       w.Base.Name,
		Type:       WebserviceType,
		Properties: util.Object2RawExtension(w.Props),
		DependsOn:  w.Base.DependsOn,
		Inputs:     w.Base.Inputs,
		Outputs:    w.Base.Outputs,
		Traits:     traits,
	}
	return comp
}

// Labels Specify the labels in the workload
func (w *WebserviceComponent) Labels(value map[string]string) *WebserviceComponent {
	w.Props.Labels = value
	return w
}

// Annotations Specify the annotations in the workload
func (w *WebserviceComponent) Annotations(value map[string]string) *WebserviceComponent {
	w.Props.Annotations = value
	return w
}

// Image Which image would you like to use for your service
func (w *WebserviceComponent) Image(value string) *WebserviceComponent {
	w.Props.Image = value
	return w
}

// ImagePullPolicy Specify image pull policy for your service
func (w *WebserviceComponent) ImagePullPolicy(value string) *WebserviceComponent {
	w.Props.ImagePullPolicy = value
	return w
}

// ImagePullSecrets Specify image pull secrets for your service
func (w *WebserviceComponent) ImagePullSecrets(value []string) *WebserviceComponent {
	w.Props.ImagePullSecrets = value
	return w
}

// Port Deprecated field, please use ports instead
func (w *WebserviceComponent) Port(value int) *WebserviceComponent {
	w.Props.Port = value
	return w
}

// Ports Which ports do you want customer traffic sent to, defaults to 80
func (w *WebserviceComponent) Ports(value Ports) *WebserviceComponent {
	w.Props.Ports = value
	return w
}

// ExposeType Specify what kind of Service you want. options: "ClusterIP", "NodePort", "LoadBalancer"
func (w *WebserviceComponent) ExposeType(value string) *WebserviceComponent {
	w.Props.ExposeType = value
	return w
}

// AddRevisionLabel If addRevisionLabel is true, the revision label will be added to the underlying pods
func (w *WebserviceComponent) AddRevisionLabel(value bool) *WebserviceComponent {
	w.Props.AddRevisionLabel = value
	return w
}

// Cmd Commands to run in the container
func (w *WebserviceComponent) Cmd(value []string) *WebserviceComponent {
	w.Props.Cmd = value
	return w
}

// Args Arguments to the entrypoint
func (w *WebserviceComponent) Args(value []string) *WebserviceComponent {
	w.Props.Args = value
	return w
}

// Env Define arguments by using environment variables
func (w *WebserviceComponent) Env(value Env) *WebserviceComponent {
	w.Props.Env = value
	return w
}

// CPU Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
func (w *WebserviceComponent) CPU(value string) *WebserviceComponent {
	w.Props.CPU = value
	return w
}

// Memory Specifies the attributes of the memory resource required for the container.
func (w *WebserviceComponent) Memory(value string) *WebserviceComponent {
	w.Props.Memory = value
	return w
}

// VolumeMounts -
func (w *WebserviceComponent) VolumeMounts(value VolumeMounts) *WebserviceComponent {
	w.Props.VolumeMounts = value
	return w
}

// Volumes Deprecated field, use volumeMounts instead.
func (w *WebserviceComponent) Volumes(value Volumes) *WebserviceComponent {
	w.Props.Volumes = value
	return w
}

// LivenessProbe Instructions for assessing whether the container is alive.
func (w *WebserviceComponent) LivenessProbe(value HealthProbe) *WebserviceComponent {
	w.Props.LivenessProbe = value
	return w
}

// ReadinessProbe Instructions for assessing whether the container is in a suitable state to serve traffic.
func (w *WebserviceComponent) ReadinessProbe(value HealthProbe) *WebserviceComponent {
	w.Props.ReadinessProbe = value
	return w
}

// HostAliases Specify the hostAliases to add
func (w *WebserviceComponent) HostAliases(value HostAliases) *WebserviceComponent {
	w.Props.HostAliases = value
	return w
}
