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

const TaskType = "task"

// TaskComponent is the root struct of task
type TaskComponent struct {
	Base  ComponentBase
	Props TaskSpec
}

// TaskSpec -
type TaskSpec struct {
	Labels           map[string]string `json:"labels,omitempty"`
	Annotations      map[string]string `json:"annotations,omitempty"`
	Count            int               `json:"count"`
	Image            string            `json:"image"`
	ImagePullPolicy  string            `json:"imagePullPolicy,omitempty"`
	ImagePullSecrets []string          `json:"imagePullSecrets,omitempty"`
	Restart          string            `json:"restart"`
	Cmd              []string          `json:"cmd,omitempty"`
	Env              Env               `json:"env,omitempty"`
	CPU              string            `json:"cpu,omitempty"`
	Memory           string            `json:"memory,omitempty"`
	Volumes          Volumes           `json:"volumes,omitempty"`
	LivenessProbe    HealthProbe       `json:"livenessProbe,omitempty"`
	ReadinessProbe   HealthProbe       `json:"readinessProbe,omitempty"`
}

// Exec Instructions for assessing container health by executing a command. Either this attribute or the httpGet attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with both the httpGet attribute and the tcpSocket attribute.
type Exec struct {
	Command []string `json:"command"`
}

// HTTPHeaders -
type HTTPHeaders struct {
	Name  string `json:"name"`
	Value string `json:"value"`
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

// Volumes -
type Volumes struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
	Medium    string `json:"medium"`
	Type      string `json:"type"`
}

// TcpSocket Instructions for assessing container health by probing a TCP socket. Either this attribute or the exec attribute or the httpGet attribute MUST be specified. This attribute is mutually exclusive with both the exec attribute and the httpGet attribute.
type TcpSocket struct {
	Port int `json:"port"`
}

// SecretKeyRef Selects a key of a secret in the pod's namespace
type SecretKeyRef struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// Env -
type Env struct {
	Name      string    `json:"name"`
	Value     string    `json:"value,omitempty"`
	ValueFrom ValueFrom `json:"valueFrom,omitempty"`
}

// HTTPGet Instructions for assessing container health by executing an HTTP GET request. Either this attribute or the exec attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with both the exec attribute and the tcpSocket attribute.
type HTTPGet struct {
	Path        string      `json:"path"`
	Port        int         `json:"port"`
	HTTPHeaders HTTPHeaders `json:"httpHeaders,omitempty"`
}

func Task(name string) *TaskComponent {
	comp := TaskComponent{
		Base: ComponentBase{
			Name: name,
		},
	}
	return &comp
}

func (t *TaskComponent) Trait(traits ...Trait) *TaskComponent {
	t.Base.Traits = append(t.Base.Traits, traits...)
	return t
}

func (t *TaskComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range t.Base.Traits {
		traits = append(traits, trait.Build())
	}
	comp := common.ApplicationComponent{
		Name:       t.Base.Name,
		Type:       TaskType,
		Properties: util.Object2RawExtension(t.Props),
		DependsOn:  t.Base.DependsOn,
		Inputs:     t.Base.Inputs,
		Outputs:    t.Base.Outputs,
		Traits:     traits,
	}
	return comp
}

// Labels Specify the labels in the workload
func (t *TaskComponent) Labels(value map[string]string) *TaskComponent {
	t.Props.Labels = value
	return t
}

// Annotations Specify the annotations in the workload
func (t *TaskComponent) Annotations(value map[string]string) *TaskComponent {
	t.Props.Annotations = value
	return t
}

// Count Specify number of tasks to run in parallel
func (t *TaskComponent) Count(value int) *TaskComponent {
	t.Props.Count = value
	return t
}

// Image Which image would you like to use for your service
func (t *TaskComponent) Image(value string) *TaskComponent {
	t.Props.Image = value
	return t
}

// ImagePullPolicy Specify image pull policy for your service
func (t *TaskComponent) ImagePullPolicy(value string) *TaskComponent {
	t.Props.ImagePullPolicy = value
	return t
}

// ImagePullSecrets Specify image pull secrets for your service
func (t *TaskComponent) ImagePullSecrets(value []string) *TaskComponent {
	t.Props.ImagePullSecrets = value
	return t
}

// Restart Define the job restart policy, the value can only be Never or OnFailure. By default, it's Never.
func (t *TaskComponent) Restart(value string) *TaskComponent {
	t.Props.Restart = value
	return t
}

// Cmd Commands to run in the container
func (t *TaskComponent) Cmd(value []string) *TaskComponent {
	t.Props.Cmd = value
	return t
}

// Env Define arguments by using environment variables
func (t *TaskComponent) Env(value Env) *TaskComponent {
	t.Props.Env = value
	return t
}

// CPU Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
func (t *TaskComponent) CPU(value string) *TaskComponent {
	t.Props.CPU = value
	return t
}

// Memory Specifies the attributes of the memory resource required for the container.
func (t *TaskComponent) Memory(value string) *TaskComponent {
	t.Props.Memory = value
	return t
}

// Volumes Declare volumes and volumeMounts
func (t *TaskComponent) Volumes(value Volumes) *TaskComponent {
	t.Props.Volumes = value
	return t
}

// LivenessProbe Instructions for assessing whether the container is alive.
func (t *TaskComponent) LivenessProbe(value HealthProbe) *TaskComponent {
	t.Props.LivenessProbe = value
	return t
}

// ReadinessProbe Instructions for assessing whether the container is in a suitable state to serve traffic.
func (t *TaskComponent) ReadinessProbe(value HealthProbe) *TaskComponent {
	t.Props.ReadinessProbe = value
	return t
}
