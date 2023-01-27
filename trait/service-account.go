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

const ServiceAccountType = "service-account"

// ServiceAccountTrait is the root struct of service-account
type ServiceAccountTrait struct {
	Base  TraitBase
	Props ServiceAccountSpec
}

// Privileges -
type Privileges struct {
	Verbs           []string `json:"verbs"`
	APIGroups       []string `json:"apiGroups,omitempty"`
	Resources       []string `json:"resources,omitempty"`
	ResourceNames   []string `json:"resourceNames,omitempty"`
	NonResourceUrLs []string `json:"nonResourceURLs,omitempty"`
	Scope           string   `json:"scope"`
}

// Privileges -
type Privileges struct {
	Verbs           []string `json:"verbs"`
	APIGroups       []string `json:"apiGroups,omitempty"`
	Resources       []string `json:"resources,omitempty"`
	ResourceNames   []string `json:"resourceNames,omitempty"`
	NonResourceUrLs []string `json:"nonResourceURLs,omitempty"`
	Scope           string   `json:"scope"`
}

// ServiceAccountSpec -
type ServiceAccountSpec struct {
	Name       string     `json:"name"`
	Create     bool       `json:"create"`
	Privileges Privileges `json:"privileges,omitempty"`
}

func ServiceAccount() *ServiceAccountTrait {
	trait := ServiceAccountTrait{
		Base: TraitBase{},
	}
	return &trait
}

func (s *ServiceAccountTrait) Build() common.ApplicationTrait {
	trait := common.ApplicationTrait{
		Type:       ServiceAccountType,
		Properties: util.Object2RawExtension(s.Props),
	}
	return trait
}

// Name Specify the name of ServiceAccount
func (s *ServiceAccountTrait) Name(value string) *ServiceAccountTrait {
	s.Props.Name = value
	return s
}

// Create Specify whether to create new ServiceAccount or not
func (s *ServiceAccountTrait) Create(value bool) *ServiceAccountTrait {
	s.Props.Create = value
	return s
}

// Privileges Specify the privileges of the ServiceAccount, if not empty, RoleBindings(ClusterRoleBindings) will be created
func (s *ServiceAccountTrait) Privileges(value Privileges) *ServiceAccountTrait {
	s.Props.Privileges = value
	return s
}