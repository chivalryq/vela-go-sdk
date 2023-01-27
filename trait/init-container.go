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

const InitContainerType = "init-container"

// InitContainerTrait is the root struct of init-container
type InitContainerTrait struct {
	Base  TraitBase
	Props InitContainerSpec
}

// Env -
type Env struct {
	Name      string    `json:"name"`
	Value     string    `json:"value,omitempty"`
	ValueFrom ValueFrom `json:"valueFrom,omitempty"`
}

// ExtraVolumeMounts -
type ExtraVolumeMounts struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
}

// InitContainerSpec -
type InitContainerSpec struct {
	Name              string            `json:"name"`
	Image             string            `json:"image"`
	ImagePullPolicy   string            `json:"imagePullPolicy"`
	Cmd               []string          `json:"cmd,omitempty"`
	Args              []string          `json:"args,omitempty"`
	Env               Env               `json:"env,omitempty"`
	MountName         string            `json:"mountName"`
	AppMountPath      string            `json:"appMountPath"`
	InitMountPath     string            `json:"initMountPath"`
	ExtraVolumeMounts ExtraVolumeMounts `json:"extraVolumeMounts"`
}

// SecretKeyRef Selects a key of a secret in the pod's namespace
type SecretKeyRef struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// ConfigMapKeyRef Selects a key of a config map in the pod's namespace
type ConfigMapKeyRef struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// ValueFrom Specifies a source the value of this var should come from
type ValueFrom struct {
	SecretKeyRef    SecretKeyRef    `json:"secretKeyRef,omitempty"`
	ConfigMapKeyRef ConfigMapKeyRef `json:"configMapKeyRef,omitempty"`
}

func InitContainer() *InitContainerTrait {
	trait := InitContainerTrait{
		Base: TraitBase{},
	}
	return &trait
}

func (i *InitContainerTrait) Build() common.ApplicationTrait {
	trait := common.ApplicationTrait{
		Type:       InitContainerType,
		Properties: util.Object2RawExtension(i.Props),
	}
	return trait
}

// Name Specify the name of init container
func (i *InitContainerTrait) Name(value string) *InitContainerTrait {
	i.Props.Name = value
	return i
}

// Image Specify the image of init container
func (i *InitContainerTrait) Image(value string) *InitContainerTrait {
	i.Props.Image = value
	return i
}

// ImagePullPolicy Specify image pull policy for your service
func (i *InitContainerTrait) ImagePullPolicy(value string) *InitContainerTrait {
	i.Props.ImagePullPolicy = value
	return i
}

// Cmd Specify the commands run in the init container
func (i *InitContainerTrait) Cmd(value []string) *InitContainerTrait {
	i.Props.Cmd = value
	return i
}

// Args Specify the args run in the init container
func (i *InitContainerTrait) Args(value []string) *InitContainerTrait {
	i.Props.Args = value
	return i
}

// Env Specify the env run in the init container
func (i *InitContainerTrait) Env(value Env) *InitContainerTrait {
	i.Props.Env = value
	return i
}

// MountName Specify the mount name of shared volume
func (i *InitContainerTrait) MountName(value string) *InitContainerTrait {
	i.Props.MountName = value
	return i
}

// AppMountPath Specify the mount path of app container
func (i *InitContainerTrait) AppMountPath(value string) *InitContainerTrait {
	i.Props.AppMountPath = value
	return i
}

// InitMountPath Specify the mount path of init container
func (i *InitContainerTrait) InitMountPath(value string) *InitContainerTrait {
	i.Props.InitMountPath = value
	return i
}

// ExtraVolumeMounts Specify the extra volume mounts for the init container
func (i *InitContainerTrait) ExtraVolumeMounts(value ExtraVolumeMounts) *InitContainerTrait {
	i.Props.ExtraVolumeMounts = value
	return i
}
