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

const LifecycleType = "lifecycle"

// LifecycleTrait is the root struct of lifecycle
type LifecycleTrait struct {
	Base  TraitBase
	Props LifecycleSpec
}

// HTTPHeaders -
type HTTPHeaders struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// HTTPGet -
type HTTPGet struct {
	Path        string      `json:"path,omitempty"`
	Port        int         `json:"port"`
	Host        string      `json:"host,omitempty"`
	Scheme      string      `json:"scheme"`
	HTTPHeaders HTTPHeaders `json:"httpHeaders,omitempty"`
}

// TcpSocket -
type TcpSocket struct {
	Port int    `json:"port"`
	Host string `json:"host,omitempty"`
}

// LifeCycleHandler -
type LifeCycleHandler struct {
	Exec      Exec      `json:"exec,omitempty"`
	HTTPGet   HTTPGet   `json:"httpGet,omitempty"`
	TcpSocket TcpSocket `json:"tcpSocket,omitempty"`
}

// LifecycleSpec -
type LifecycleSpec struct {
	PostStart LifeCycleHandler `json:"postStart,omitempty"`
	PreStop   LifeCycleHandler `json:"preStop,omitempty"`
}

// Port -
type Port Port

// Exec -
type Exec struct {
	Command []string `json:"command"`
}

func Lifecycle() *LifecycleTrait {
	trait := LifecycleTrait{
		Base: TraitBase{},
	}
	return &trait
}

func (l *LifecycleTrait) Build() common.ApplicationTrait {
	trait := common.ApplicationTrait{
		Type:       LifecycleType,
		Properties: util.Object2RawExtension(l.Props),
	}
	return trait
}

// PostStart -
func (l *LifecycleTrait) PostStart(value LifeCycleHandler) *LifecycleTrait {
	l.Props.PostStart = value
	return l
}

// PreStop -
func (l *LifecycleTrait) PreStop(value LifeCycleHandler) *LifecycleTrait {
	l.Props.PreStop = value
	return l
}
