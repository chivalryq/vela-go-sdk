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

const TopologyspreadconstraintsType = "topologyspreadconstraints"

// TopologyspreadconstraintsTrait is the root struct of topologyspreadconstraints
type TopologyspreadconstraintsTrait struct {
	Base  TraitBase
	Props TopologyspreadconstraintsSpec
}

// LabSelector -
type LabSelector struct {
	MatchLabels      map[string]string `json:"matchLabels,omitempty"`
	MatchExpressions MatchExpressions  `json:"matchExpressions,omitempty"`
}

// Constraints -
type Constraints struct {
	MaxSkew            int         `json:"maxSkew"`
	TopologyKey        string      `json:"topologyKey"`
	WhenUnsatisfiable  string      `json:"whenUnsatisfiable"`
	LabelSelector      labSelector `json:"labelSelector"`
	MinDomains         int         `json:"minDomains,omitempty"`
	MatchLabelKeys     []string    `json:"matchLabelKeys,omitempty"`
	NodeAffinityPolicy string      `json:"nodeAffinityPolicy,omitempty"`
	NodeTaintsPolicy   string      `json:"nodeTaintsPolicy,omitempty"`
}

// TopologyspreadconstraintsSpec -
type TopologyspreadconstraintsSpec struct {
	Constraints Constraints `json:"constraints"`
}

// MatchExpressions -
type MatchExpressions struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Values   []string `json:"values,omitempty"`
}

func Topologyspreadconstraints() *TopologyspreadconstraintsTrait {
	trait := TopologyspreadconstraintsTrait{
		Base: TraitBase{},
	}
	return &trait
}

func (t *TopologyspreadconstraintsTrait) Build() common.ApplicationTrait {
	trait := common.ApplicationTrait{
		Type:       TopologyspreadconstraintsType,
		Properties: util.Object2RawExtension(t.Props),
	}
	return trait
}

// Constraints -
func (t *TopologyspreadconstraintsTrait) Constraints(value Constraints) *TopologyspreadconstraintsTrait {
	t.Props.Constraints = value
	return t
}
