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

const PureIngressType = "pure-ingress"

// PureIngressTrait is the root struct of pure-ingress
type PureIngressTrait struct {
	Base  TraitBase
	Props PureIngressSpec
}

// PureIngressSpec -
type PureIngressSpec struct {
	Domain string         `json:"domain"`
	HTTP   map[string]int `json:"http"`
}

func PureIngress() *PureIngressTrait {
	trait := PureIngressTrait{
		Base: TraitBase{},
	}
	return &trait
}

func (p *PureIngressTrait) Build() common.ApplicationTrait {
	trait := common.ApplicationTrait{
		Type:       PureIngressType,
		Properties: util.Object2RawExtension(p.Props),
	}
	return trait
}

// Domain Specify the domain you want to expose
func (p *PureIngressTrait) Domain(value string) *PureIngressTrait {
	p.Props.Domain = value
	return p
}

// HTTP Specify the mapping relationship between the http path and the workload port
func (p *PureIngressTrait) HTTP(value map[string]int) *PureIngressTrait {
	p.Props.HTTP = value
	return p
}
