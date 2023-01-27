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

const CpuscalerType = "cpuscaler"

// CpuscalerTrait is the root struct of cpuscaler
type CpuscalerTrait struct {
	Base  TraitBase
	Props CpuscalerSpec
}

// CpuscalerSpec -
type CpuscalerSpec struct {
	Min              int    `json:"min"`
	Max              int    `json:"max"`
	CPUUtil          int    `json:"cpuUtil"`
	TargetAPIVersion string `json:"targetAPIVersion"`
	TargetKind       string `json:"targetKind"`
}

func Cpuscaler() *CpuscalerTrait {
	trait := CpuscalerTrait{
		Base: TraitBase{},
	}
	return &trait
}

func (c *CpuscalerTrait) Build() common.ApplicationTrait {
	trait := common.ApplicationTrait{
		Type:       CpuscalerType,
		Properties: util.Object2RawExtension(c.Props),
	}
	return trait
}

// Min Specify the minimal number of replicas to which the autoscaler can scale down
func (c *CpuscalerTrait) Min(value int) *CpuscalerTrait {
	c.Props.Min = value
	return c
}

// Max Specify the maximum number of of replicas to which the autoscaler can scale up
func (c *CpuscalerTrait) Max(value int) *CpuscalerTrait {
	c.Props.Max = value
	return c
}

// CPUUtil Specify the average CPU utilization, for example, 50 means the CPU usage is 50%
func (c *CpuscalerTrait) CPUUtil(value int) *CpuscalerTrait {
	c.Props.CPUUtil = value
	return c
}

// TargetAPIVersion Specify the apiVersion of scale target
func (c *CpuscalerTrait) TargetAPIVersion(value string) *CpuscalerTrait {
	c.Props.TargetAPIVersion = value
	return c
}

// TargetKind Specify the kind of scale target
func (c *CpuscalerTrait) TargetKind(value string) *CpuscalerTrait {
	c.Props.TargetKind = value
	return c
}
