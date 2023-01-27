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

const ContainerImageType = "container-image"

// ContainerImageTrait is the root struct of container-image
type ContainerImageTrait struct {
	Base  TraitBase
	Props ContainerImageSpec
}

// PatchParams -
type PatchParams struct {
	ContainerName   string `json:"containerName"`
	Image           string `json:"image"`
	ImagePullPolicy string `json:"imagePullPolicy"`
}

// ContainerImageSpec -
type ContainerImageSpec struct {
	ContainerName   string `json:"containerName"`
	Image           string `json:"image"`
	ImagePullPolicy string `json:"imagePullPolicy"`
}

func ContainerImage() *ContainerImageTrait {
	trait := ContainerImageTrait{
		Base: TraitBase{},
	}
	return &trait
}

func (c *ContainerImageTrait) Build() common.ApplicationTrait {
	trait := common.ApplicationTrait{
		Type:       ContainerImageType,
		Properties: util.Object2RawExtension(c.Props),
	}
	return trait
}

// ContainerName Specify the name of the target container, if not set, use the component name
func (c *ContainerImageTrait) ContainerName(value string) *ContainerImageTrait {
	c.Props.ContainerName = value
	return c
}

// Image Specify the image of the container
func (c *ContainerImageTrait) Image(value string) *ContainerImageTrait {
	c.Props.Image = value
	return c
}

// ImagePullPolicy Specify the image pull policy of the container
func (c *ContainerImageTrait) ImagePullPolicy(value string) *ContainerImageTrait {
	c.Props.ImagePullPolicy = value
	return c
}
