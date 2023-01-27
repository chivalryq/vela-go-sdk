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

const GatewayType = "gateway"

// GatewayTrait is the root struct of gateway
type GatewayTrait struct {
	Base  TraitBase
	Props GatewaySpec
}

// GatewaySpec -
type GatewaySpec struct {
	Domain      string         `json:"domain,omitempty"`
	HTTP        map[string]int `json:"http"`
	Class       string         `json:"class"`
	ClassInSpec bool           `json:"classInSpec"`
	SecretName  string         `json:"secretName,omitempty"`
	GatewayHost string         `json:"gatewayHost,omitempty"`
}

func Gateway() *GatewayTrait {
	trait := GatewayTrait{
		Base: TraitBase{},
	}
	return &trait
}

func (g *GatewayTrait) Build() common.ApplicationTrait {
	trait := common.ApplicationTrait{
		Type:       GatewayType,
		Properties: util.Object2RawExtension(g.Props),
	}
	return trait
}

// Domain Specify the domain you want to expose
func (g *GatewayTrait) Domain(value string) *GatewayTrait {
	g.Props.Domain = value
	return g
}

// HTTP Specify the mapping relationship between the http path and the workload port
func (g *GatewayTrait) HTTP(value map[string]int) *GatewayTrait {
	g.Props.HTTP = value
	return g
}

// Class Specify the class of ingress to use
func (g *GatewayTrait) Class(value string) *GatewayTrait {
	g.Props.Class = value
	return g
}

// ClassInSpec Set ingress class in '.spec.ingressClassName' instead of 'kubernetes.io/ingress.class' annotation.
func (g *GatewayTrait) ClassInSpec(value bool) *GatewayTrait {
	g.Props.ClassInSpec = value
	return g
}

// SecretName Specify the secret name you want to quote to use tls.
func (g *GatewayTrait) SecretName(value string) *GatewayTrait {
	g.Props.SecretName = value
	return g
}

// GatewayHost Specify the host of the ingress gateway, which is used to generate the endpoints when the host is empty.
func (g *GatewayTrait) GatewayHost(value string) *GatewayTrait {
	g.Props.GatewayHost = value
	return g
}
