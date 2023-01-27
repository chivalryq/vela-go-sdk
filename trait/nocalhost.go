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

const NocalhostType = "nocalhost"

// NocalhostTrait is the root struct of nocalhost
type NocalhostTrait struct {
	Base  TraitBase
	Props NocalhostSpec
}

// Sync -
type Sync struct {
	Type              string `json:"type"`
	FilePattern       list   `json:"filePattern"`
	IgnoreFilePattern list   `json:"ignoreFilePattern"`
}

// Env -
type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// PersistentVolumeDirs -
type PersistentVolumeDirs struct {
	Path     string `json:"path"`
	Capacity string `json:"capacity"`
}

// Limits -
type Limits struct {
	Memory string `json:"memory"`
	CPU    string `json:"cpu"`
}

// Resources -
type Resources struct {
	Limits   Limits   `json:"limits"`
	Requests Requests `json:"requests"`
}

// Debug -
type Debug struct {
	RemoteDebugPort int `json:"remoteDebugPort,omitempty"`
}

// Requests -
type Requests struct {
	Memory string `json:"memory"`
	CPU    string `json:"cpu"`
}

// NocalhostSpec -
type NocalhostSpec struct {
	Port                 int                  `json:"port"`
	ServiceType          string               `json:"serviceType"`
	GitURL               string               `json:"gitUrl,omitempty"`
	Image                string               `json:"image"`
	Shell                string               `json:"shell"`
	WorkDir              string               `json:"workDir"`
	StorageClass         string               `json:"storageClass,omitempty"`
	Command              Command              `json:"command"`
	Debug                Debug                `json:"debug,omitempty"`
	HotReload            bool                 `json:"hotReload"`
	Sync                 Sync                 `json:"sync"`
	Env                  Env                  `json:"env,omitempty"`
	PortForward          []string             `json:"portForward,omitempty"`
	PersistentVolumeDirs PersistentVolumeDirs `json:"persistentVolumeDirs,omitempty"`
	Resources            Resources            `json:"resources"`
}

// Command -
type Command struct {
	Run   list `json:"run"`
	Debug list `json:"debug"`
}

func Nocalhost() *NocalhostTrait {
	trait := NocalhostTrait{
		Base: TraitBase{},
	}
	return &trait
}

func (n *NocalhostTrait) Build() common.ApplicationTrait {
	trait := common.ApplicationTrait{
		Type:       NocalhostType,
		Properties: util.Object2RawExtension(n.Props),
	}
	return trait
}

// Port -
func (n *NocalhostTrait) Port(value int) *NocalhostTrait {
	n.Props.Port = value
	return n
}

// ServiceType -
func (n *NocalhostTrait) ServiceType(value string) *NocalhostTrait {
	n.Props.ServiceType = value
	return n
}

// GitURL -
func (n *NocalhostTrait) GitURL(value string) *NocalhostTrait {
	n.Props.GitURL = value
	return n
}

// Image -
func (n *NocalhostTrait) Image(value string) *NocalhostTrait {
	n.Props.Image = value
	return n
}

// Shell -
func (n *NocalhostTrait) Shell(value string) *NocalhostTrait {
	n.Props.Shell = value
	return n
}

// WorkDir -
func (n *NocalhostTrait) WorkDir(value string) *NocalhostTrait {
	n.Props.WorkDir = value
	return n
}

// StorageClass -
func (n *NocalhostTrait) StorageClass(value string) *NocalhostTrait {
	n.Props.StorageClass = value
	return n
}

// Command -
func (n *NocalhostTrait) Command(value Command) *NocalhostTrait {
	n.Props.Command = value
	return n
}

// Debug -
func (n *NocalhostTrait) Debug(value Debug) *NocalhostTrait {
	n.Props.Debug = value
	return n
}

// HotReload -
func (n *NocalhostTrait) HotReload(value bool) *NocalhostTrait {
	n.Props.HotReload = value
	return n
}

// Sync -
func (n *NocalhostTrait) Sync(value Sync) *NocalhostTrait {
	n.Props.Sync = value
	return n
}

// Env -
func (n *NocalhostTrait) Env(value Env) *NocalhostTrait {
	n.Props.Env = value
	return n
}

// PortForward -
func (n *NocalhostTrait) PortForward(value []string) *NocalhostTrait {
	n.Props.PortForward = value
	return n
}

// PersistentVolumeDirs -
func (n *NocalhostTrait) PersistentVolumeDirs(value PersistentVolumeDirs) *NocalhostTrait {
	n.Props.PersistentVolumeDirs = value
	return n
}

// Resources -
func (n *NocalhostTrait) Resources(value Resources) *NocalhostTrait {
	n.Props.Resources = value
	return n
}
