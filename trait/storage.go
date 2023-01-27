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

const StorageType = "storage"

// StorageTrait is the root struct of storage
type StorageTrait struct {
	Base  TraitBase
	Props StorageSpec
}

// MatchExpressions -
type MatchExpressions struct {
	Key      string   `json:"key"`
	Values   []string `json:"values"`
	Operator string   `json:"operator"`
}

// EmptyDir -
type EmptyDir struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
	SubPath   string `json:"subPath,omitempty"`
	Medium    string `json:"medium"`
}

// Resources -
type Resources struct {
	Requests Requests `json:"requests"`
	Limits   Limits   `json:"limits,omitempty"`
}

// Secret -
type Secret struct {
	Name        string            `json:"name"`
	MountOnly   bool              `json:"mountOnly"`
	MountToEnv  SecretMountToEnv  `json:"mountToEnv,omitempty"`
	MountToEnvs SecretMountToEnvs `json:"mountToEnvs,omitempty"`
	MountPath   string            `json:"mountPath,omitempty"`
	SubPath     string            `json:"subPath,omitempty"`
	DefaultMode int               `json:"defaultMode"`
	ReadOnly    bool              `json:"readOnly"`
	StringData  map[string]_      `json:"stringData,omitempty"`
	Data        map[string]_      `json:"data,omitempty"`
	Items       Items             `json:"items,omitempty"`
}

// Items -
type Items struct {
	Key  string `json:"key"`
	Path string `json:"path"`
	Mode int    `json:"mode"`
}

// ConfigMap -
type ConfigMap struct {
	Name        string               `json:"name"`
	MountOnly   bool                 `json:"mountOnly"`
	MountToEnv  ConfigMapMountToEnv  `json:"mountToEnv,omitempty"`
	MountToEnvs ConfigMapMountToEnvs `json:"mountToEnvs,omitempty"`
	MountPath   string               `json:"mountPath,omitempty"`
	SubPath     string               `json:"subPath,omitempty"`
	DefaultMode int                  `json:"defaultMode"`
	ReadOnly    bool                 `json:"readOnly"`
	Data        map[string]_         `json:"data,omitempty"`
	Items       Items                `json:"items,omitempty"`
}

// SecretMountToEnv -
type SecretMountToEnv struct {
	EnvName   string `json:"envName"`
	SecretKey string `json:"secretKey"`
}

// StorageSpec -
type StorageSpec struct {
	PVC       PVC       `json:"pvc,omitempty"`
	ConfigMap ConfigMap `json:"configMap,omitempty"`
	Secret    Secret    `json:"secret,omitempty"`
	EmptyDir  EmptyDir  `json:"emptyDir,omitempty"`
}

// DataSourceRef -
type DataSourceRef struct {
	Name     string `json:"name"`
	Kind     string `json:"kind"`
	APIGroup string `json:"apiGroup"`
}

// Selector -
type Selector struct {
	MatchLabels      map[string]string `json:"matchLabels,omitempty"`
	MatchExpressions MatchExpressions  `json:"matchExpressions,omitempty"`
}

// ConfigMapMountToEnv -
type ConfigMapMountToEnv struct {
	EnvName      string `json:"envName"`
	ConfigMapKey string `json:"configMapKey"`
}

// ConfigMapMountToEnvs -
type ConfigMapMountToEnvs struct {
	EnvName      string `json:"envName"`
	ConfigMapKey string `json:"configMapKey"`
}

// SecretMountToEnvs -
type SecretMountToEnvs struct {
	EnvName   string `json:"envName"`
	SecretKey string `json:"secretKey"`
}

// Requests -
type Requests struct {
	Storage string `json:"storage"`
}

// Limits -
type Limits struct {
	Storage string `json:"storage"`
}

// DataSource -
type DataSource struct {
	Name     string `json:"name"`
	Kind     string `json:"kind"`
	APIGroup string `json:"apiGroup"`
}

// PVC -
type PVC struct {
	Name             string        `json:"name"`
	MountOnly        bool          `json:"mountOnly"`
	MountPath        string        `json:"mountPath"`
	SubPath          string        `json:"subPath,omitempty"`
	VolumeMode       string        `json:"volumeMode"`
	VolumeName       string        `json:"volumeName,omitempty"`
	AccessModes      list          `json:"accessModes"`
	StorageClassName string        `json:"storageClassName,omitempty"`
	Resources        Resources     `json:"resources,omitempty"`
	DataSourceRef    DataSourceRef `json:"dataSourceRef,omitempty"`
	DataSource       DataSource    `json:"dataSource,omitempty"`
	Selector         Selector      `json:"selector,omitempty"`
}

func Storage() *StorageTrait {
	trait := StorageTrait{
		Base: TraitBase{},
	}
	return &trait
}

func (s *StorageTrait) Build() common.ApplicationTrait {
	trait := common.ApplicationTrait{
		Type:       StorageType,
		Properties: util.Object2RawExtension(s.Props),
	}
	return trait
}

// PVC Declare pvc type storage
func (s *StorageTrait) PVC(value PVC) *StorageTrait {
	s.Props.PVC = value
	return s
}

// ConfigMap Declare config map type storage
func (s *StorageTrait) ConfigMap(value ConfigMap) *StorageTrait {
	s.Props.ConfigMap = value
	return s
}

// Secret Declare secret type storage
func (s *StorageTrait) Secret(value Secret) *StorageTrait {
	s.Props.Secret = value
	return s
}

// EmptyDir Declare empty dir type storage
func (s *StorageTrait) EmptyDir(value EmptyDir) *StorageTrait {
	s.Props.EmptyDir = value
	return s
}
