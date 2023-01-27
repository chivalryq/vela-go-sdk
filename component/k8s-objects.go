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

package component

import (
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	. "vela-go-sdk/api"
)

const K8sObjectsType = "k8s-objects"

// K8sObjectsComponent is the root struct of k8s-objects
type K8sObjectsComponent struct {
	Base  ComponentBase
	Props K8sObjectsSpec
}

// Objects -
type Objects struct {
}

// K8SObjectsSpec -
type K8SObjectsSpec struct {
	Objects Objects `json:"objects"`
}

func K8sObjects(name string) *K8sObjectsComponent {
	comp := K8sObjectsComponent{
		Base: ComponentBase{
			Name: name,
		},
	}
	return &comp
}

func (k *K8sObjectsComponent) Trait(traits ...Trait) *K8sObjectsComponent {
	k.Base.Traits = append(k.Base.Traits, traits...)
	return k
}

func (k *K8sObjectsComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range k.Base.Traits {
		traits = append(traits, trait.Build())
	}
	comp := common.ApplicationComponent{
		Name:       k.Base.Name,
		Type:       K8sObjectsType,
		Properties: util.Object2RawExtension(k.Props),
		DependsOn:  k.Base.DependsOn,
		Inputs:     k.Base.Inputs,
		Outputs:    k.Base.Outputs,
		Traits:     traits,
	}
	return comp
}

// Objects -
func (k *K8sObjectsComponent) Objects(value Objects) *K8sObjectsComponent {
	k.Props.Objects = value
	return k
}
