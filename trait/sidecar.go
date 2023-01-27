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

package trait

import (
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	. "vela-go-sdk/api"
)

const SidecarType = "sidecar"

// SidecarTrait is the root struct of sidecar
type SidecarTrait struct {
	Base  TraitBase
	Props SidecarSpec
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

// FieldRef Specify the field reference for env
type FieldRef struct {
	FieldPath string `json:"fieldPath"`
}

// Exec Instructions for assessing container health by executing a command. Either this attribute or the httpGet attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with both the httpGet attribute and the tcpSocket attribute.
type Exec struct {
	Command []string `json:"command"`
}

// HTTPGet Instructions for assessing container health by executing an HTTP GET request. Either this attribute or the exec attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with both the exec attribute and the tcpSocket attribute.
type HTTPGet struct {
	Path        string      `json:"path"`
	Port        int         `json:"port"`
	HTTPHeaders HTTPHeaders `json:"httpHeaders,omitempty"`
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
	FieldRef        FieldRef        `json:"fieldRef,omitempty"`
}

// Env -
type Env struct {
	Name      string    `json:"name"`
	Value     string    `json:"value,omitempty"`
	ValueFrom ValueFrom `json:"valueFrom,omitempty"`
}

// Volumes -
type Volumes struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// SidecarSpec -
type SidecarSpec struct {
	Name           string      `json:"name"`
	Image          string      `json:"image"`
	Cmd            []string    `json:"cmd,omitempty"`
	Args           []string    `json:"args,omitempty"`
	Env            Env         `json:"env,omitempty"`
	Volumes        Volumes     `json:"volumes,omitempty"`
	LivenessProbe  HealthProbe `json:"livenessProbe,omitempty"`
	ReadinessProbe HealthProbe `json:"readinessProbe,omitempty"`
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

func Sidecar() *SidecarTrait {
	trait := SidecarTrait{
		Base: TraitBase{},
	}
	return &trait
}

func (s *SidecarTrait) Build() common.ApplicationTrait {
	trait := common.ApplicationTrait{
		Type:       SidecarType,
		Properties: util.Object2RawExtension(s.Props),
	}
	return trait
}

// Name Specify the name of sidecar container
func (s *SidecarTrait) Name(value string) *SidecarTrait {
	s.Props.Name = value
	return s
}

// Image Specify the image of sidecar container
func (s *SidecarTrait) Image(value string) *SidecarTrait {
	s.Props.Image = value
	return s
}

// Cmd Specify the commands run in the sidecar
func (s *SidecarTrait) Cmd(value []string) *SidecarTrait {
	s.Props.Cmd = value
	return s
}

// Args Specify the args in the sidecar
func (s *SidecarTrait) Args(value []string) *SidecarTrait {
	s.Props.Args = value
	return s
}

// Env Specify the env in the sidecar
func (s *SidecarTrait) Env(value Env) *SidecarTrait {
	s.Props.Env = value
	return s
}

// Volumes Specify the shared volume path
func (s *SidecarTrait) Volumes(value Volumes) *SidecarTrait {
	s.Props.Volumes = value
	return s
}

// LivenessProbe Instructions for assessing whether the container is alive.
func (s *SidecarTrait) LivenessProbe(value HealthProbe) *SidecarTrait {
	s.Props.LivenessProbe = value
	return s
}

// ReadinessProbe Instructions for assessing whether the container is in a suitable state to serve traffic.
func (s *SidecarTrait) ReadinessProbe(value HealthProbe) *SidecarTrait {
	s.Props.ReadinessProbe = value
	return s
}
