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

const AnnotationsType = "annotations"

// AnnotationsTrait is the root struct of annotations
type AnnotationsTrait struct {
	Base  TraitBase
	Props AnnotationsSpec
}

// AnnotationsSpec -
type AnnotationsSpec map[string]string

func Annotations() *AnnotationsTrait {
	trait := AnnotationsTrait{
		Base: TraitBase{},
	}
	return &trait
}

func (a *AnnotationsTrait) Build() common.ApplicationTrait {
	trait := common.ApplicationTrait{
		Type:       AnnotationsType,
		Properties: util.Object2RawExtension(a.Props),
	}
	return trait
}
