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

const AffinityType = "affinity"

// AffinityTrait is the root struct of affinity
type AffinityTrait struct {
	Base  TraitBase
	Props AffinitySpec
}

// MatchFields -
type MatchFields struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Values   []string `json:"values,omitempty"`
}

// NodeSelectorTerm -
type NodeSelectorTerm struct {
	MatchExpressions MatchExpressions `json:"matchExpressions,omitempty"`
	MatchFields      MatchFields      `json:"matchFields,omitempty"`
}

// PodAffinity Specify the pod affinity scheduling rules
type PodAffinity struct {
	Required  PodAffinityRequired  `json:"required,omitempty"`
	Preferred PodAffinityPreferred `json:"preferred,omitempty"`
}

// NodeAffinityRequired Specify the required during scheduling ignored during execution
type NodeAffinityRequired struct {
	NodeSelectorTerms NodeSelectorTerms `json:"nodeSelectorTerms"`
}

// NodeAffinity Specify the node affinity scheduling rules for the pod
type NodeAffinity struct {
	Required  NodeAffinityRequired  `json:"required,omitempty"`
	Preferred NodeAffinityPreferred `json:"preferred,omitempty"`
}

// MatchExpressions -
type MatchExpressions struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Values   []string `json:"values,omitempty"`
}

// NodeSelecor -
type NodeSelecor struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Values   []string `json:"values,omitempty"`
}

// PodAffinityPreferred -
type PodAffinityPreferred struct {
	Weight          int             `json:"weight"`
	PodAffinityTerm podAffinityTerm `json:"podAffinityTerm"`
}

// Tolerations -
type Tolerations struct {
	Key               string `json:"key,omitempty"`
	Operator          string `json:"operator"`
	Value             string `json:"value,omitempty"`
	Effect            string `json:"effect,omitempty"`
	TolerationSeconds int    `json:"tolerationSeconds,omitempty"`
}

// NodeAffinityPreferred -
type NodeAffinityPreferred struct {
	Weight     int              `json:"weight"`
	Preference nodeSelectorTerm `json:"preference"`
}

// AffinitySpec -
type AffinitySpec struct {
	PodAffinity     PodAffinity     `json:"podAffinity,omitempty"`
	PodAntiAffinity PodAntiAffinity `json:"podAntiAffinity,omitempty"`
	NodeAffinity    NodeAffinity    `json:"nodeAffinity,omitempty"`
	Tolerations     Tolerations     `json:"tolerations,omitempty"`
}

// LabelSelector -
type LabelSelector struct {
	MatchLabels      map[string]string `json:"matchLabels,omitempty"`
	MatchExpressions MatchExpressions  `json:"matchExpressions,omitempty"`
}

// PodAffinityTerm -
type PodAffinityTerm struct {
	LabelSelector     labelSelector `json:"labelSelector,omitempty"`
	Namespaces        []string      `json:"namespaces,omitempty"`
	TopologyKey       string        `json:"topologyKey"`
	NamespaceSelector labelSelector `json:"namespaceSelector,omitempty"`
}

// NodeSelectorTerms -
type NodeSelectorTerms struct {
	MatchExpressions MatchExpressions `json:"matchExpressions,omitempty"`
	MatchFields      MatchFields      `json:"matchFields,omitempty"`
}

// PodAntiAffinity Specify the pod anti-affinity scheduling rules
type PodAntiAffinity struct {
	Required  Required  `json:"required,omitempty"`
	Preferred Preferred `json:"preferred,omitempty"`
}

// PodAffinityRequired -
type PodAffinityRequired struct {
	LabelSelector     labelSelector `json:"labelSelector,omitempty"`
	Namespaces        []string      `json:"namespaces,omitempty"`
	TopologyKey       string        `json:"topologyKey"`
	NamespaceSelector labelSelector `json:"namespaceSelector,omitempty"`
}

func Affinity() *AffinityTrait {
	trait := AffinityTrait{
		Base: TraitBase{},
	}
	return &trait
}

func (a *AffinityTrait) Build() common.ApplicationTrait {
	trait := common.ApplicationTrait{
		Type:       AffinityType,
		Properties: util.Object2RawExtension(a.Props),
	}
	return trait
}

// PodAffinity Specify the pod affinity scheduling rules
func (a *AffinityTrait) PodAffinity(value PodAffinity) *AffinityTrait {
	a.Props.PodAffinity = value
	return a
}

// PodAntiAffinity Specify the pod anti-affinity scheduling rules
func (a *AffinityTrait) PodAntiAffinity(value PodAntiAffinity) *AffinityTrait {
	a.Props.PodAntiAffinity = value
	return a
}

// NodeAffinity Specify the node affinity scheduling rules for the pod
func (a *AffinityTrait) NodeAffinity(value NodeAffinity) *AffinityTrait {
	a.Props.NodeAffinity = value
	return a
}

// Tolerations Specify tolerant taint
func (a *AffinityTrait) Tolerations(value Tolerations) *AffinityTrait {
	a.Props.Tolerations = value
	return a
}
