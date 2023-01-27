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

const StartupProbeType = "startup-probe"

// StartupProbeTrait is the root struct of startup-probe
type StartupProbeTrait struct {
	Base  TraitBase
	Props StartupProbeSpec
}

// TcpSocket Instructions for assessing container startup status by probing a TCP socket. Either this attribute or the exec attribute or the tcpSocket attribute or the httpGet attribute MUST be specified. This attribute is mutually exclusive with the exec attribute and the httpGet attribute and the gRPC attribute.
type TcpSocket struct {
	Port string `json:"port"`
	Host string `json:"host,omitempty"`
}

// StartupProbeParams -
type StartupProbeParams struct {
	ContainerName                 string                    `json:"containerName"`
	InitialDelaySeconds           int                       `json:"initialDelaySeconds"`
	PeriodSeconds                 int                       `json:"periodSeconds"`
	TimeoutSeconds                int                       `json:"timeoutSeconds"`
	SuccessThreshold              int                       `json:"successThreshold"`
	FailureThreshold              int                       `json:"failureThreshold"`
	TerminationGracePeriodSeconds int                       `json:"terminationGracePeriodSeconds,omitempty"`
	Exec                          Exec                      `json:"exec,omitempty"`
	HTTPGet                       StartupProbeParamsHTTPGet `json:"httpGet,omitempty"`
	Grpc                          Grpc                      `json:"grpc,omitempty"`
	TcpSocket                     TcpSocket                 `json:"tcpSocket,omitempty"`
}

// StartupProbeParamsHTTPGet Instructions for assessing container startup status by executing an HTTP GET request. Either this attribute or the exec attribute or the grpc attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with the exec attribute and the tcpSocket attribute and the gRPC attribute.
type StartupProbeParamsHTTPGet struct {
	Path        string      `json:"path,omitempty"`
	Port        int         `json:"port"`
	Host        string      `json:"host,omitempty"`
	Scheme      string      `json:"scheme,omitempty"`
	HTTPHeaders HTTPHeaders `json:"httpHeaders,omitempty"`
}

// StartupProbeSpecHTTPGet Instructions for assessing container startup status by executing an HTTP GET request. Either this attribute or the exec attribute or the grpc attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with the exec attribute and the tcpSocket attribute and the gRPC attribute.
type StartupProbeSpecHTTPGet struct {
	Path        string      `json:"path,omitempty"`
	Port        int         `json:"port"`
	Host        string      `json:"host,omitempty"`
	Scheme      string      `json:"scheme,omitempty"`
	HTTPHeaders HTTPHeaders `json:"httpHeaders,omitempty"`
}

// Exec Instructions for assessing container startup status by executing a command. Either this attribute or the httpGet attribute or the grpc attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with the httpGet attribute and the tcpSocket attribute and the gRPC attribute.
type Exec struct {
	Command []string `json:"command"`
}

// HTTPHeaders -
type HTTPHeaders struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// StartupProbeSpec -
type StartupProbeSpec struct {
	ContainerName                 string                  `json:"containerName"`
	InitialDelaySeconds           int                     `json:"initialDelaySeconds"`
	PeriodSeconds                 int                     `json:"periodSeconds"`
	TimeoutSeconds                int                     `json:"timeoutSeconds"`
	SuccessThreshold              int                     `json:"successThreshold"`
	FailureThreshold              int                     `json:"failureThreshold"`
	TerminationGracePeriodSeconds int                     `json:"terminationGracePeriodSeconds,omitempty"`
	Exec                          Exec                    `json:"exec,omitempty"`
	HTTPGet                       StartupProbeSpecHTTPGet `json:"httpGet,omitempty"`
	Grpc                          Grpc                    `json:"grpc,omitempty"`
	TcpSocket                     TcpSocket               `json:"tcpSocket,omitempty"`
}

// Grpc Instructions for assessing container startup status by probing a gRPC service. Either this attribute or the exec attribute or the grpc attribute or the httpGet attribute MUST be specified. This attribute is mutually exclusive with the exec attribute and the httpGet attribute and the tcpSocket attribute.
type Grpc struct {
	Port    int    `json:"port"`
	Service string `json:"service,omitempty"`
}

func StartupProbe() *StartupProbeTrait {
	trait := StartupProbeTrait{
		Base: TraitBase{},
	}
	return &trait
}

func (s *StartupProbeTrait) Build() common.ApplicationTrait {
	trait := common.ApplicationTrait{
		Type:       StartupProbeType,
		Properties: util.Object2RawExtension(s.Props),
	}
	return trait
}

// ContainerName Specify the name of the target container, if not set, use the component name
func (s *StartupProbeTrait) ContainerName(value string) *StartupProbeTrait {
	s.Props.ContainerName = value
	return s
}

// InitialDelaySeconds Number of seconds after the container has started before liveness probes are initiated. Minimum value is 0.
func (s *StartupProbeTrait) InitialDelaySeconds(value int) *StartupProbeTrait {
	s.Props.InitialDelaySeconds = value
	return s
}

// PeriodSeconds How often, in seconds, to execute the probe. Minimum value is 1.
func (s *StartupProbeTrait) PeriodSeconds(value int) *StartupProbeTrait {
	s.Props.PeriodSeconds = value
	return s
}

// TimeoutSeconds Number of seconds after which the probe times out. Minimum value is 1.
func (s *StartupProbeTrait) TimeoutSeconds(value int) *StartupProbeTrait {
	s.Props.TimeoutSeconds = value
	return s
}

// SuccessThreshold Minimum consecutive successes for the probe to be considered successful after having failed.  Minimum value is 1.
func (s *StartupProbeTrait) SuccessThreshold(value int) *StartupProbeTrait {
	s.Props.SuccessThreshold = value
	return s
}

// FailureThreshold Minimum consecutive failures for the probe to be considered failed after having succeeded. Minimum value is 1.
func (s *StartupProbeTrait) FailureThreshold(value int) *StartupProbeTrait {
	s.Props.FailureThreshold = value
	return s
}

// TerminationGracePeriodSeconds Optional duration in seconds the pod needs to terminate gracefully upon probe failure. Set this value longer than the expected cleanup time for your process.
func (s *StartupProbeTrait) TerminationGracePeriodSeconds(value int) *StartupProbeTrait {
	s.Props.TerminationGracePeriodSeconds = value
	return s
}

// Exec Instructions for assessing container startup status by executing a command. Either this attribute or the httpGet attribute or the grpc attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with the httpGet attribute and the tcpSocket attribute and the gRPC attribute.
func (s *StartupProbeTrait) Exec(value Exec) *StartupProbeTrait {
	s.Props.Exec = value
	return s
}

// HTTPGet Instructions for assessing container startup status by executing an HTTP GET request. Either this attribute or the exec attribute or the grpc attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with the exec attribute and the tcpSocket attribute and the gRPC attribute.
func (s *StartupProbeTrait) HTTPGet(value StartupProbeSpecHTTPGet) *StartupProbeTrait {
	s.Props.HTTPGet = value
	return s
}

// Grpc Instructions for assessing container startup status by probing a gRPC service. Either this attribute or the exec attribute or the grpc attribute or the httpGet attribute MUST be specified. This attribute is mutually exclusive with the exec attribute and the httpGet attribute and the tcpSocket attribute.
func (s *StartupProbeTrait) Grpc(value Grpc) *StartupProbeTrait {
	s.Props.Grpc = value
	return s
}

// TcpSocket Instructions for assessing container startup status by probing a TCP socket. Either this attribute or the exec attribute or the tcpSocket attribute or the httpGet attribute MUST be specified. This attribute is mutually exclusive with the exec attribute and the httpGet attribute and the gRPC attribute.
func (s *StartupProbeTrait) TcpSocket(value TcpSocket) *StartupProbeTrait {
	s.Props.TcpSocket = value
	return s
}
