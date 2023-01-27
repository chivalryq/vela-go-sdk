/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package daemon

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the VolumeMounts type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &VolumeMounts{}

// VolumeMounts struct for VolumeMounts
type VolumeMounts struct {
	// Mount ConfigMap type volume
	configMap []ConfigMap `json:"configMap,omitempty"`
	// Mount EmptyDir type volume
	emptyDir []EmptyDir `json:"emptyDir,omitempty"`
	// Mount HostPath type volume
	hostPath []HostPath `json:"hostPath,omitempty"`
	// Mount PVC type volume
	pvc []Pvc `json:"pvc,omitempty"`
	// Mount Secret type volume
	secret []Secret `json:"secret,omitempty"`
}

// NewVolumeMounts instantiates a new VolumeMounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeMounts() *VolumeMounts {
	this := VolumeMounts{}
	return &this
}

// NewVolumeMountsWithDefaults instantiates a new VolumeMounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeMountsWithDefaults() *VolumeMounts {
	this := VolumeMounts{}
	return &this
}

// GetConfigMap returns the ConfigMap field value if set, zero value otherwise.
func (o *VolumeMounts) GetConfigMap() []ConfigMap {
	if o == nil || utils.IsNil(o.configMap) {
		var ret []ConfigMap
		return ret
	}
	return o.configMap
}

// GetConfigMapOk returns a tuple with the ConfigMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeMounts) GetConfigMapOk() ([]ConfigMap, bool) {
	if o == nil || utils.IsNil(o.configMap) {
		return nil, false
	}
	return o.configMap, true
}

// HasConfigMap returns a boolean if a field has been set.
func (o *VolumeMounts) HasConfigMap() bool {
	if o != nil && !utils.IsNil(o.configMap) {
		return true
	}

	return false
}

// SetConfigMap gets a reference to the given []ConfigMap and assigns it to the configMap field.
// configMap:  Mount ConfigMap type volume 

func (o *VolumeMounts) ConfigMap(v []ConfigMap) (*VolumeMounts){
	o.configMap = v
return o
}

// GetEmptyDir returns the EmptyDir field value if set, zero value otherwise.
func (o *VolumeMounts) GetEmptyDir() []EmptyDir {
	if o == nil || utils.IsNil(o.emptyDir) {
		var ret []EmptyDir
		return ret
	}
	return o.emptyDir
}

// GetEmptyDirOk returns a tuple with the EmptyDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeMounts) GetEmptyDirOk() ([]EmptyDir, bool) {
	if o == nil || utils.IsNil(o.emptyDir) {
		return nil, false
	}
	return o.emptyDir, true
}

// HasEmptyDir returns a boolean if a field has been set.
func (o *VolumeMounts) HasEmptyDir() bool {
	if o != nil && !utils.IsNil(o.emptyDir) {
		return true
	}

	return false
}

// SetEmptyDir gets a reference to the given []EmptyDir and assigns it to the emptyDir field.
// emptyDir:  Mount EmptyDir type volume 

func (o *VolumeMounts) EmptyDir(v []EmptyDir) (*VolumeMounts){
	o.emptyDir = v
return o
}

// GetHostPath returns the HostPath field value if set, zero value otherwise.
func (o *VolumeMounts) GetHostPath() []HostPath {
	if o == nil || utils.IsNil(o.hostPath) {
		var ret []HostPath
		return ret
	}
	return o.hostPath
}

// GetHostPathOk returns a tuple with the HostPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeMounts) GetHostPathOk() ([]HostPath, bool) {
	if o == nil || utils.IsNil(o.hostPath) {
		return nil, false
	}
	return o.hostPath, true
}

// HasHostPath returns a boolean if a field has been set.
func (o *VolumeMounts) HasHostPath() bool {
	if o != nil && !utils.IsNil(o.hostPath) {
		return true
	}

	return false
}

// SetHostPath gets a reference to the given []HostPath and assigns it to the hostPath field.
// hostPath:  Mount HostPath type volume 

func (o *VolumeMounts) HostPath(v []HostPath) (*VolumeMounts){
	o.hostPath = v
return o
}

// GetPvc returns the Pvc field value if set, zero value otherwise.
func (o *VolumeMounts) GetPvc() []Pvc {
	if o == nil || utils.IsNil(o.pvc) {
		var ret []Pvc
		return ret
	}
	return o.pvc
}

// GetPvcOk returns a tuple with the Pvc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeMounts) GetPvcOk() ([]Pvc, bool) {
	if o == nil || utils.IsNil(o.pvc) {
		return nil, false
	}
	return o.pvc, true
}

// HasPvc returns a boolean if a field has been set.
func (o *VolumeMounts) HasPvc() bool {
	if o != nil && !utils.IsNil(o.pvc) {
		return true
	}

	return false
}

// SetPvc gets a reference to the given []Pvc and assigns it to the pvc field.
// pvc:  Mount PVC type volume 

func (o *VolumeMounts) Pvc(v []Pvc) (*VolumeMounts){
	o.pvc = v
return o
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *VolumeMounts) GetSecret() []Secret {
	if o == nil || utils.IsNil(o.secret) {
		var ret []Secret
		return ret
	}
	return o.secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeMounts) GetSecretOk() ([]Secret, bool) {
	if o == nil || utils.IsNil(o.secret) {
		return nil, false
	}
	return o.secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *VolumeMounts) HasSecret() bool {
	if o != nil && !utils.IsNil(o.secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given []Secret and assigns it to the secret field.
// secret:  Mount Secret type volume 

func (o *VolumeMounts) Secret(v []Secret) (*VolumeMounts){
	o.secret = v
return o
}

func (o VolumeMounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeMounts) ToMap() (map[string]interface{}, error) {
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

type NullableVolumeMounts struct {
	value *VolumeMounts
	isSet bool
}

func (v NullableVolumeMounts) Get() *VolumeMounts {
	return v.value
}

func (v *NullableVolumeMounts) Set(val *VolumeMounts) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeMounts) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeMounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeMounts(val *VolumeMounts) *NullableVolumeMounts {
	return &NullableVolumeMounts{value: val, isSet: true}
}

func (v NullableVolumeMounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeMounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

 
