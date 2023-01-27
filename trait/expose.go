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

const ExposeType = "expose"

// ExposeTrait is the root struct of expose
type ExposeTrait struct {
	Base  TraitBase
	Props ExposeSpec
}

// ExposeSpec -
type ExposeSpec struct {
	Port        []int             `json:"port"`
	Annotations map[string]string `json:"annotations"`
	Type        string            `json:"type"`
}

func Expose() *ExposeTrait {
	trait := ExposeTrait{
		Base: TraitBase{},
	}
	return &trait
}

func (e *ExposeTrait) Build() common.ApplicationTrait {
	trait := common.ApplicationTrait{
		Type:       ExposeType,
		Properties: util.Object2RawExtension(e.Props),
	}
	return trait
}

// Port Specify the exposion ports
func (e *ExposeTrait) Port(value []int) *ExposeTrait {
	e.Props.Port = value
	return e
}

// Annotations Specify the annotaions of the exposed service
func (e *ExposeTrait) Annotations(value map[string]string) *ExposeTrait {
	e.Props.Annotations = value
	return e
}

// Type Specify what kind of Service you want. options: "ClusterIP","NodePort","LoadBalancer","ExternalName"
func (e *ExposeTrait) Type(value string) *ExposeTrait {
	e.Props.Type = value
	return e
}
