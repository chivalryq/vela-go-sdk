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

const ResourceType = "resource"

// ResourceTrait is the root struct of resource
type ResourceTrait struct {
	Base  TraitBase
	Props ResourceSpec
}

// Limits Specify the resources in limits
type Limits struct {
	CPU    float64 `json:"cpu"`
	Memory string  `json:"memory"`
}

// ResourceSpec -
type ResourceSpec struct {
	CPU      float64  `json:"cpu,omitempty"`
	Memory   string   `json:"memory,omitempty"`
	Requests Requests `json:"requests,omitempty"`
	Limits   Limits   `json:"limits,omitempty"`
}

// Requests Specify the resources in requests
type Requests struct {
	CPU    float64 `json:"cpu"`
	Memory string  `json:"memory"`
}

func Resource() *ResourceTrait {
	trait := ResourceTrait{
		Base: TraitBase{},
	}
	return &trait
}

func (r *ResourceTrait) Build() common.ApplicationTrait {
	trait := common.ApplicationTrait{
		Type:       ResourceType,
		Properties: util.Object2RawExtension(r.Props),
	}
	return trait
}

// CPU Specify the amount of cpu for requests and limits
func (r *ResourceTrait) CPU(value float64) *ResourceTrait {
	r.Props.CPU = value
	return r
}

// Memory Specify the amount of memory for requests and limits
func (r *ResourceTrait) Memory(value string) *ResourceTrait {
	r.Props.Memory = value
	return r
}

// Requests Specify the resources in requests
func (r *ResourceTrait) Requests(value Requests) *ResourceTrait {
	r.Props.Requests = value
	return r
}

// Limits Specify the resources in limits
func (r *ResourceTrait) Limits(value Limits) *ResourceTrait {
	r.Props.Limits = value
	return r
}
