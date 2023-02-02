/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the StorageSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &StorageSpec{}

// StorageSpec struct for StorageSpec
type StorageSpec struct {
	// Declare config map type storage
	configMap []ConfigMap `json:"configMap,omitempty"`
	// Declare empty dir type storage
	emptyDir []EmptyDir `json:"emptyDir,omitempty"`
	// Declare host path type storage
	hostPath []HostPath `json:"hostPath,omitempty"`
	// Declare pvc type storage
	pvc []Pvc `json:"pvc,omitempty"`
	// Declare secret type storage
	secret []Secret `json:"secret,omitempty"`
}

// NewStorageSpecWith instantiates a new StorageSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSpecWith() *StorageSpec {
	this := StorageSpec{}
	return &this
}

// NewStorageSpec instantiates a new StorageSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSpec() *StorageSpec {
	this := StorageSpec{}
	return &this
}

// GetConfigMap returns the ConfigMap field value if set, zero value otherwise.
func (o *StorageTrait) GetConfigMap() []ConfigMap {
	if o == nil || utils.IsNil(o.Properties.configMap) {
		var ret []ConfigMap
		return ret
	}
	return o.Properties.configMap
}

// GetConfigMapOk returns a tuple with the ConfigMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageTrait) GetConfigMapOk() ([]ConfigMap, bool) {
	if o == nil || utils.IsNil(o.Properties.configMap) {
		return nil, false
	}
	return o.Properties.configMap, true
}

// HasConfigMap returns a boolean if a field has been set.
func (o *StorageTrait) HasConfigMap() bool {
	if o != nil && !utils.IsNil(o.Properties.configMap) {
		return true
	}

	return false
}

// ConfigMap gets a reference to the given []ConfigMap and assigns it to the configMap field.
// configMap:  Declare config map type storage
func (o *StorageTrait) ConfigMap(v []ConfigMap) *StorageTrait {
	o.Properties.configMap = v
	return o
}

// GetEmptyDir returns the EmptyDir field value if set, zero value otherwise.
func (o *StorageTrait) GetEmptyDir() []EmptyDir {
	if o == nil || utils.IsNil(o.Properties.emptyDir) {
		var ret []EmptyDir
		return ret
	}
	return o.Properties.emptyDir
}

// GetEmptyDirOk returns a tuple with the EmptyDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageTrait) GetEmptyDirOk() ([]EmptyDir, bool) {
	if o == nil || utils.IsNil(o.Properties.emptyDir) {
		return nil, false
	}
	return o.Properties.emptyDir, true
}

// HasEmptyDir returns a boolean if a field has been set.
func (o *StorageTrait) HasEmptyDir() bool {
	if o != nil && !utils.IsNil(o.Properties.emptyDir) {
		return true
	}

	return false
}

// EmptyDir gets a reference to the given []EmptyDir and assigns it to the emptyDir field.
// emptyDir:  Declare empty dir type storage
func (o *StorageTrait) EmptyDir(v []EmptyDir) *StorageTrait {
	o.Properties.emptyDir = v
	return o
}

// GetHostPath returns the HostPath field value if set, zero value otherwise.
func (o *StorageTrait) GetHostPath() []HostPath {
	if o == nil || utils.IsNil(o.Properties.hostPath) {
		var ret []HostPath
		return ret
	}
	return o.Properties.hostPath
}

// GetHostPathOk returns a tuple with the HostPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageTrait) GetHostPathOk() ([]HostPath, bool) {
	if o == nil || utils.IsNil(o.Properties.hostPath) {
		return nil, false
	}
	return o.Properties.hostPath, true
}

// HasHostPath returns a boolean if a field has been set.
func (o *StorageTrait) HasHostPath() bool {
	if o != nil && !utils.IsNil(o.Properties.hostPath) {
		return true
	}

	return false
}

// HostPath gets a reference to the given []HostPath and assigns it to the hostPath field.
// hostPath:  Declare host path type storage
func (o *StorageTrait) HostPath(v []HostPath) *StorageTrait {
	o.Properties.hostPath = v
	return o
}

// GetPvc returns the Pvc field value if set, zero value otherwise.
func (o *StorageTrait) GetPvc() []Pvc {
	if o == nil || utils.IsNil(o.Properties.pvc) {
		var ret []Pvc
		return ret
	}
	return o.Properties.pvc
}

// GetPvcOk returns a tuple with the Pvc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageTrait) GetPvcOk() ([]Pvc, bool) {
	if o == nil || utils.IsNil(o.Properties.pvc) {
		return nil, false
	}
	return o.Properties.pvc, true
}

// HasPvc returns a boolean if a field has been set.
func (o *StorageTrait) HasPvc() bool {
	if o != nil && !utils.IsNil(o.Properties.pvc) {
		return true
	}

	return false
}

// Pvc gets a reference to the given []Pvc and assigns it to the pvc field.
// pvc:  Declare pvc type storage
func (o *StorageTrait) Pvc(v []Pvc) *StorageTrait {
	o.Properties.pvc = v
	return o
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *StorageTrait) GetSecret() []Secret {
	if o == nil || utils.IsNil(o.Properties.secret) {
		var ret []Secret
		return ret
	}
	return o.Properties.secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageTrait) GetSecretOk() ([]Secret, bool) {
	if o == nil || utils.IsNil(o.Properties.secret) {
		return nil, false
	}
	return o.Properties.secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *StorageTrait) HasSecret() bool {
	if o != nil && !utils.IsNil(o.Properties.secret) {
		return true
	}

	return false
}

// Secret gets a reference to the given []Secret and assigns it to the secret field.
// secret:  Declare secret type storage
func (o *StorageTrait) Secret(v []Secret) *StorageTrait {
	o.Properties.secret = v
	return o
}

func (o StorageSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.configMap) {
		toSerialize["configMap"] = o.configMap
	}
	if !utils.IsNil(o.emptyDir) {
		toSerialize["emptyDir"] = o.emptyDir
	}
	if !utils.IsNil(o.hostPath) {
		toSerialize["hostPath"] = o.hostPath
	}
	if !utils.IsNil(o.pvc) {
		toSerialize["pvc"] = o.pvc
	}
	if !utils.IsNil(o.secret) {
		toSerialize["secret"] = o.secret
	}
	return toSerialize, nil
}

type NullableStorageSpec struct {
	value *StorageSpec
	isSet bool
}

func (v NullableStorageSpec) Get() *StorageSpec {
	return v.value
}

func (v *NullableStorageSpec) Set(val *StorageSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSpec(val *StorageSpec) *NullableStorageSpec {
	return &NullableStorageSpec{value: val, isSet: true}
}

func (v NullableStorageSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const StorageType = "storage"

func init() {
	sdkcommon.RegisterTrait(StorageType, FromTrait)
}

type StorageTrait struct {
	Base       apis.TraitBase
	Properties StorageSpec
}

func Storage() *StorageTrait {
	s := &StorageTrait{Base: apis.TraitBase{}}
	return s
}

func (s *StorageTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(s.Properties),
		Type:       StorageType,
	}
	return res
}

func (s *StorageTrait) FromTrait(from common.ApplicationTrait) (*StorageTrait, error) {
	var properties StorageSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	s.Properties = properties
	return s, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	s := &StorageTrait{}
	return s.FromTrait(from)
}

func (s *StorageTrait) DefType() string {
	return StorageType
}
