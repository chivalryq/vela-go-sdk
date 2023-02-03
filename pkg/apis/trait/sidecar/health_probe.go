/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sidecar

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the HealthProbe type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HealthProbe{}

// HealthProbe struct for HealthProbe
type HealthProbe struct {
	Exec *Exec `json:"exec,omitempty"`
	// Number of consecutive failures required to determine the container is not alive (liveness probe) or not ready (readiness probe).
	FailureThreshold *int32   `json:"failureThreshold,omitempty"`
	HttpGet          *HttpGet `json:"httpGet,omitempty"`
	// Number of seconds after the container is started before the first probe is initiated.
	InitialDelaySeconds *int32 `json:"initialDelaySeconds,omitempty"`
	// How often, in seconds, to execute the probe.
	PeriodSeconds *int32 `json:"periodSeconds,omitempty"`
	// Minimum consecutive successes for the probe to be considered successful after having failed.
	SuccessThreshold *int32     `json:"successThreshold,omitempty"`
	TcpSocket        *TcpSocket `json:"tcpSocket,omitempty"`
	// Number of seconds after which the probe times out.
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`
}

// NewHealthProbeWith instantiates a new HealthProbe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthProbeWith() *HealthProbe {
	this := HealthProbe{}
	var failureThreshold int32 = 3
	this.FailureThreshold = &failureThreshold
	var initialDelaySeconds int32 = 0
	this.InitialDelaySeconds = &initialDelaySeconds
	var periodSeconds int32 = 10
	this.PeriodSeconds = &periodSeconds
	var successThreshold int32 = 1
	this.SuccessThreshold = &successThreshold
	var timeoutSeconds int32 = 1
	this.TimeoutSeconds = &timeoutSeconds
	return &this
}

// NewHealthProbe instantiates a new HealthProbe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthProbe() *HealthProbe {
	this := HealthProbe{}
	var failureThreshold int32 = 3
	this.FailureThreshold = &failureThreshold
	var initialDelaySeconds int32 = 0
	this.InitialDelaySeconds = &initialDelaySeconds
	var periodSeconds int32 = 10
	this.PeriodSeconds = &periodSeconds
	var successThreshold int32 = 1
	this.SuccessThreshold = &successThreshold
	var timeoutSeconds int32 = 1
	this.TimeoutSeconds = &timeoutSeconds
	return &this
}

// NewHealthProbes converts a list HealthProbe pointers to objects.
// This is helpful when the SetHealthProbe requires a list of objects
func NewHealthProbes(ps ...*HealthProbe) []HealthProbe {
	objs := []HealthProbe{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetExec returns the Exec field value if set, zero value otherwise.
func (o *HealthProbe) GetExec() Exec {
	if o == nil || utils.IsNil(o.Exec) {
		var ret Exec
		return ret
	}
	return *o.Exec
}

// GetExecOk returns a tuple with the Exec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthProbe) GetExecOk() (*Exec, bool) {
	if o == nil || utils.IsNil(o.Exec) {
		return nil, false
	}
	return o.Exec, true
}

// HasExec returns a boolean if a field has been set.
func (o *HealthProbe) HasExec() bool {
	if o != nil && !utils.IsNil(o.Exec) {
		return true
	}

	return false
}

// SetExec gets a reference to the given Exec and assigns it to the exec field.
// Exec:
func (o *HealthProbe) SetExec(v Exec) *HealthProbe {
	o.Exec = &v
	return o
}

// GetFailureThreshold returns the FailureThreshold field value if set, zero value otherwise.
func (o *HealthProbe) GetFailureThreshold() int32 {
	if o == nil || utils.IsNil(o.FailureThreshold) {
		var ret int32
		return ret
	}
	return *o.FailureThreshold
}

// GetFailureThresholdOk returns a tuple with the FailureThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthProbe) GetFailureThresholdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.FailureThreshold) {
		return nil, false
	}
	return o.FailureThreshold, true
}

// HasFailureThreshold returns a boolean if a field has been set.
func (o *HealthProbe) HasFailureThreshold() bool {
	if o != nil && !utils.IsNil(o.FailureThreshold) {
		return true
	}

	return false
}

// SetFailureThreshold gets a reference to the given int32 and assigns it to the failureThreshold field.
// FailureThreshold:  Number of consecutive failures required to determine the container is not alive (liveness probe) or not ready (readiness probe).
func (o *HealthProbe) SetFailureThreshold(v int32) *HealthProbe {
	o.FailureThreshold = &v
	return o
}

// GetHttpGet returns the HttpGet field value if set, zero value otherwise.
func (o *HealthProbe) GetHttpGet() HttpGet {
	if o == nil || utils.IsNil(o.HttpGet) {
		var ret HttpGet
		return ret
	}
	return *o.HttpGet
}

// GetHttpGetOk returns a tuple with the HttpGet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthProbe) GetHttpGetOk() (*HttpGet, bool) {
	if o == nil || utils.IsNil(o.HttpGet) {
		return nil, false
	}
	return o.HttpGet, true
}

// HasHttpGet returns a boolean if a field has been set.
func (o *HealthProbe) HasHttpGet() bool {
	if o != nil && !utils.IsNil(o.HttpGet) {
		return true
	}

	return false
}

// SetHttpGet gets a reference to the given HttpGet and assigns it to the httpGet field.
// HttpGet:
func (o *HealthProbe) SetHttpGet(v HttpGet) *HealthProbe {
	o.HttpGet = &v
	return o
}

// GetInitialDelaySeconds returns the InitialDelaySeconds field value if set, zero value otherwise.
func (o *HealthProbe) GetInitialDelaySeconds() int32 {
	if o == nil || utils.IsNil(o.InitialDelaySeconds) {
		var ret int32
		return ret
	}
	return *o.InitialDelaySeconds
}

// GetInitialDelaySecondsOk returns a tuple with the InitialDelaySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthProbe) GetInitialDelaySecondsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.InitialDelaySeconds) {
		return nil, false
	}
	return o.InitialDelaySeconds, true
}

// HasInitialDelaySeconds returns a boolean if a field has been set.
func (o *HealthProbe) HasInitialDelaySeconds() bool {
	if o != nil && !utils.IsNil(o.InitialDelaySeconds) {
		return true
	}

	return false
}

// SetInitialDelaySeconds gets a reference to the given int32 and assigns it to the initialDelaySeconds field.
// InitialDelaySeconds:  Number of seconds after the container is started before the first probe is initiated.
func (o *HealthProbe) SetInitialDelaySeconds(v int32) *HealthProbe {
	o.InitialDelaySeconds = &v
	return o
}

// GetPeriodSeconds returns the PeriodSeconds field value if set, zero value otherwise.
func (o *HealthProbe) GetPeriodSeconds() int32 {
	if o == nil || utils.IsNil(o.PeriodSeconds) {
		var ret int32
		return ret
	}
	return *o.PeriodSeconds
}

// GetPeriodSecondsOk returns a tuple with the PeriodSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthProbe) GetPeriodSecondsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PeriodSeconds) {
		return nil, false
	}
	return o.PeriodSeconds, true
}

// HasPeriodSeconds returns a boolean if a field has been set.
func (o *HealthProbe) HasPeriodSeconds() bool {
	if o != nil && !utils.IsNil(o.PeriodSeconds) {
		return true
	}

	return false
}

// SetPeriodSeconds gets a reference to the given int32 and assigns it to the periodSeconds field.
// PeriodSeconds:  How often, in seconds, to execute the probe.
func (o *HealthProbe) SetPeriodSeconds(v int32) *HealthProbe {
	o.PeriodSeconds = &v
	return o
}

// GetSuccessThreshold returns the SuccessThreshold field value if set, zero value otherwise.
func (o *HealthProbe) GetSuccessThreshold() int32 {
	if o == nil || utils.IsNil(o.SuccessThreshold) {
		var ret int32
		return ret
	}
	return *o.SuccessThreshold
}

// GetSuccessThresholdOk returns a tuple with the SuccessThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthProbe) GetSuccessThresholdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.SuccessThreshold) {
		return nil, false
	}
	return o.SuccessThreshold, true
}

// HasSuccessThreshold returns a boolean if a field has been set.
func (o *HealthProbe) HasSuccessThreshold() bool {
	if o != nil && !utils.IsNil(o.SuccessThreshold) {
		return true
	}

	return false
}

// SetSuccessThreshold gets a reference to the given int32 and assigns it to the successThreshold field.
// SuccessThreshold:  Minimum consecutive successes for the probe to be considered successful after having failed.
func (o *HealthProbe) SetSuccessThreshold(v int32) *HealthProbe {
	o.SuccessThreshold = &v
	return o
}

// GetTcpSocket returns the TcpSocket field value if set, zero value otherwise.
func (o *HealthProbe) GetTcpSocket() TcpSocket {
	if o == nil || utils.IsNil(o.TcpSocket) {
		var ret TcpSocket
		return ret
	}
	return *o.TcpSocket
}

// GetTcpSocketOk returns a tuple with the TcpSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthProbe) GetTcpSocketOk() (*TcpSocket, bool) {
	if o == nil || utils.IsNil(o.TcpSocket) {
		return nil, false
	}
	return o.TcpSocket, true
}

// HasTcpSocket returns a boolean if a field has been set.
func (o *HealthProbe) HasTcpSocket() bool {
	if o != nil && !utils.IsNil(o.TcpSocket) {
		return true
	}

	return false
}

// SetTcpSocket gets a reference to the given TcpSocket and assigns it to the tcpSocket field.
// TcpSocket:
func (o *HealthProbe) SetTcpSocket(v TcpSocket) *HealthProbe {
	o.TcpSocket = &v
	return o
}

// GetTimeoutSeconds returns the TimeoutSeconds field value if set, zero value otherwise.
func (o *HealthProbe) GetTimeoutSeconds() int32 {
	if o == nil || utils.IsNil(o.TimeoutSeconds) {
		var ret int32
		return ret
	}
	return *o.TimeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthProbe) GetTimeoutSecondsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TimeoutSeconds) {
		return nil, false
	}
	return o.TimeoutSeconds, true
}

// HasTimeoutSeconds returns a boolean if a field has been set.
func (o *HealthProbe) HasTimeoutSeconds() bool {
	if o != nil && !utils.IsNil(o.TimeoutSeconds) {
		return true
	}

	return false
}

// SetTimeoutSeconds gets a reference to the given int32 and assigns it to the timeoutSeconds field.
// TimeoutSeconds:  Number of seconds after which the probe times out.
func (o *HealthProbe) SetTimeoutSeconds(v int32) *HealthProbe {
	o.TimeoutSeconds = &v
	return o
}

func (o HealthProbe) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthProbe) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Exec) {
		toSerialize["exec"] = o.Exec
	}
	if !utils.IsNil(o.FailureThreshold) {
		toSerialize["failureThreshold"] = o.FailureThreshold
	}
	if !utils.IsNil(o.HttpGet) {
		toSerialize["httpGet"] = o.HttpGet
	}
	if !utils.IsNil(o.InitialDelaySeconds) {
		toSerialize["initialDelaySeconds"] = o.InitialDelaySeconds
	}
	if !utils.IsNil(o.PeriodSeconds) {
		toSerialize["periodSeconds"] = o.PeriodSeconds
	}
	if !utils.IsNil(o.SuccessThreshold) {
		toSerialize["successThreshold"] = o.SuccessThreshold
	}
	if !utils.IsNil(o.TcpSocket) {
		toSerialize["tcpSocket"] = o.TcpSocket
	}
	if !utils.IsNil(o.TimeoutSeconds) {
		toSerialize["timeoutSeconds"] = o.TimeoutSeconds
	}
	return toSerialize, nil
}

type NullableHealthProbe struct {
	value *HealthProbe
	isSet bool
}

func (v NullableHealthProbe) Get() *HealthProbe {
	return v.value
}

func (v *NullableHealthProbe) Set(val *HealthProbe) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthProbe) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthProbe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthProbe(val *HealthProbe) *NullableHealthProbe {
	return &NullableHealthProbe{value: val, isSet: true}
}

func (v NullableHealthProbe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthProbe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
