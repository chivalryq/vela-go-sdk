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

const EnvType = "env"

// EnvTrait is the root struct of env
type EnvTrait struct {
	Base  TraitBase
	Props EnvSpec
}

// PatchParams -
type PatchParams struct {
	ContainerName string            `json:"containerName"`
	Replace       bool              `json:"replace"`
	Env           map[string]string `json:"env"`
	Unset         list              `json:"unset"`
}

// EnvSpec -
type EnvSpec struct {
	ContainerName string            `json:"containerName"`
	Replace       bool              `json:"replace"`
	Env           map[string]string `json:"env"`
	Unset         list              `json:"unset"`
}

func Env() *EnvTrait {
	trait := EnvTrait{
		Base: TraitBase{},
	}
	return &trait
}

func (e *EnvTrait) Build() common.ApplicationTrait {
	trait := common.ApplicationTrait{
		Type:       EnvType,
		Properties: util.Object2RawExtension(e.Props),
	}
	return trait
}

// ContainerName Specify the name of the target container, if not set, use the component name
func (e *EnvTrait) ContainerName(value string) *EnvTrait {
	e.Props.ContainerName = value
	return e
}

// Replace Specify if replacing the whole environment settings for the container
func (e *EnvTrait) Replace(value bool) *EnvTrait {
	e.Props.Replace = value
	return e
}

// Env Specify the  environment variables to merge, if key already existing, override its value
func (e *EnvTrait) Env(value map[string]string) *EnvTrait {
	e.Props.Env = value
	return e
}

// Unset Specify which existing environment variables to unset
func (e *EnvTrait) Unset(value list) *EnvTrait {
	e.Props.Unset = value
	return e
}
