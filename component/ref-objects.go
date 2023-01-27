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

const RefObjectsType = "ref-objects"

// RefObjectsComponent is the root struct of ref-objects
type RefObjectsComponent struct {
	Base  ComponentBase
	Props RefObjectsSpec
}

// K8SObject -
type K8SObject struct {
	Resource      string            `json:"resource,omitempty"`
	Group         string            `json:"group,omitempty"`
	Name          string            `json:"name,omitempty"`
	Namespace     string            `json:"namespace,omitempty"`
	Cluster       string            `json:"cluster,omitempty"`
	LabelSelector map[string]string `json:"labelSelector,omitempty"`
}

// Objects -
type Objects struct {
	Resource      string            `json:"resource,omitempty"`
	Group         string            `json:"group,omitempty"`
	Name          string            `json:"name,omitempty"`
	Namespace     string            `json:"namespace,omitempty"`
	Cluster       string            `json:"cluster,omitempty"`
	LabelSelector map[string]string `json:"labelSelector,omitempty"`
}

// RefObjectsSpec -
type RefObjectsSpec struct {
	Objects Objects  `json:"objects,omitempty"`
	Urls    []string `json:"urls,omitempty"`
}

func RefObjects(name string) *RefObjectsComponent {
	comp := RefObjectsComponent{
		Base: ComponentBase{
			Name: name,
		},
	}
	return &comp
}

func (r *RefObjectsComponent) Trait(traits ...Trait) *RefObjectsComponent {
	r.Base.Traits = append(r.Base.Traits, traits...)
	return r
}

func (r *RefObjectsComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range r.Base.Traits {
		traits = append(traits, trait.Build())
	}
	comp := common.ApplicationComponent{
		Name:       r.Base.Name,
		Type:       RefObjectsType,
		Properties: util.Object2RawExtension(r.Props),
		DependsOn:  r.Base.DependsOn,
		Inputs:     r.Base.Inputs,
		Outputs:    r.Base.Outputs,
		Traits:     traits,
	}
	return comp
}

// Objects If specified, application will fetch native Kubernetes objects according to the object description
func (r *RefObjectsComponent) Objects(value Objects) *RefObjectsComponent {
	r.Props.Objects = value
	return r
}

// Urls If specified, the objects in the urls will be loaded.
func (r *RefObjectsComponent) Urls(value []string) *RefObjectsComponent {
	r.Props.Urls = value
	return r
}
