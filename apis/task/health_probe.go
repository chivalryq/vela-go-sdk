/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package task

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)



    
    // checks if the HealthProbe type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HealthProbe{}

// HealthProbe struct for HealthProbe
type HealthProbe struct {
	exec *Exec `json:"exec,omitempty"`
	// Number of consecutive failures required to determine the container is not alive (liveness probe) or not ready (readiness probe).
	failureThreshold int32 `json:"failureThreshold"`
	httpGet *HttpGet `json:"httpGet,omitempty"`
	// Number of seconds after the container is started before the first probe is initiated.
	initialDelaySeconds int32 `json:"initialDelaySeconds"`
	// How often, in seconds, to execute the probe.
	periodSeconds int32 `json:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed.
	successThreshold int32 `json:"successThreshold"`
	tcpSocket *TcpSocket `json:"tcpSocket,omitempty"`
	// Number of seconds after which the probe times out.
	timeoutSeconds int32 `json:"timeoutSeconds"`
}

// NewHealthProbe instantiates a new HealthProbe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthProbe(failureThreshold int32, initialDelaySeconds int32, periodSeconds int32, successThreshold int32, timeoutSeconds int32) *HealthProbe {
	this := HealthProbe{}
	this.failureThreshold = failureThreshold
	this.initialDelaySeconds = initialDelaySeconds
	this.periodSeconds = periodSeconds
	this.successThreshold = successThreshold
	this.timeoutSeconds = timeoutSeconds
	return &this
}

// NewHealthProbeWithDefaults instantiates a new HealthProbe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthProbeWithDefaults() *HealthProbe {
	this := HealthProbe{}
	var failureThreshold int32 = 3
	this.failureThreshold = failureThreshold
	var initialDelaySeconds int32 = 0
	this.initialDelaySeconds = initialDelaySeconds
	var periodSeconds int32 = 10
	this.periodSeconds = periodSeconds
	var successThreshold int32 = 1
	this.successThreshold = successThreshold
	var timeoutSeconds int32 = 1
	this.timeoutSeconds = timeoutSeconds
	return &this
}

// GetExec returns the Exec field value if set, zero value otherwise.
func (o *HealthProbe) GetExec() Exec {
	if o == nil || utils.IsNil(o.exec) {
		var ret Exec
		return ret
	}
	return *o.exec
}

// GetExecOk returns a tuple with the Exec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthProbe) GetExecOk() (*Exec, bool) {
	if o == nil || utils.IsNil(o.exec) {
		return nil, false
	}
	return o.exec, true
}

// HasExec returns a boolean if a field has been set.
func (o *HealthProbe) HasExec() bool {
	if o != nil && !utils.IsNil(o.exec) {
		return true
	}

	return false
}

// SetExec gets a reference to the given Exec and assigns it to the exec field.
// exec: 

func (o *HealthProbe) Exec(v Exec) (*HealthProbe){
	o.exec = &v
return o
}

// GetFailureThreshold returns the FailureThreshold field value
func (o *HealthProbe) GetFailureThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.failureThreshold
}

// GetFailureThresholdOk returns a tuple with the FailureThreshold field value
// and a boolean to check if the value has been set.
func (o *HealthProbe) GetFailureThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.failureThreshold, true
}

// FailureThreshold sets field value
func (o *HealthProbe) FailureThreshold(v int32) *HealthProbe {
	o.failureThreshold = v
    return o
}

// GetHttpGet returns the HttpGet field value if set, zero value otherwise.
func (o *HealthProbe) GetHttpGet() HttpGet {
	if o == nil || utils.IsNil(o.httpGet) {
		var ret HttpGet
		return ret
	}
	return *o.httpGet
}

// GetHttpGetOk returns a tuple with the HttpGet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthProbe) GetHttpGetOk() (*HttpGet, bool) {
	if o == nil || utils.IsNil(o.httpGet) {
		return nil, false
	}
	return o.httpGet, true
}

// HasHttpGet returns a boolean if a field has been set.
func (o *HealthProbe) HasHttpGet() bool {
	if o != nil && !utils.IsNil(o.httpGet) {
		return true
	}

	return false
}

// SetHttpGet gets a reference to the given HttpGet and assigns it to the httpGet field.
// httpGet: 

func (o *HealthProbe) HttpGet(v HttpGet) (*HealthProbe){
	o.httpGet = &v
return o
}

// GetInitialDelaySeconds returns the InitialDelaySeconds field value
func (o *HealthProbe) GetInitialDelaySeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.initialDelaySeconds
}

// GetInitialDelaySecondsOk returns a tuple with the InitialDelaySeconds field value
// and a boolean to check if the value has been set.
func (o *HealthProbe) GetInitialDelaySecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.initialDelaySeconds, true
}

// InitialDelaySeconds sets field value
func (o *HealthProbe) InitialDelaySeconds(v int32) *HealthProbe {
	o.initialDelaySeconds = v
    return o
}

// GetPeriodSeconds returns the PeriodSeconds field value
func (o *HealthProbe) GetPeriodSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.periodSeconds
}

// GetPeriodSecondsOk returns a tuple with the PeriodSeconds field value
// and a boolean to check if the value has been set.
func (o *HealthProbe) GetPeriodSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.periodSeconds, true
}

// PeriodSeconds sets field value
func (o *HealthProbe) PeriodSeconds(v int32) *HealthProbe {
	o.periodSeconds = v
    return o
}

// GetSuccessThreshold returns the SuccessThreshold field value
func (o *HealthProbe) GetSuccessThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.successThreshold
}

// GetSuccessThresholdOk returns a tuple with the SuccessThreshold field value
// and a boolean to check if the value has been set.
func (o *HealthProbe) GetSuccessThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.successThreshold, true
}

// SuccessThreshold sets field value
func (o *HealthProbe) SuccessThreshold(v int32) *HealthProbe {
	o.successThreshold = v
    return o
}

// GetTcpSocket returns the TcpSocket field value if set, zero value otherwise.
func (o *HealthProbe) GetTcpSocket() TcpSocket {
	if o == nil || utils.IsNil(o.tcpSocket) {
		var ret TcpSocket
		return ret
	}
	return *o.tcpSocket
}

// GetTcpSocketOk returns a tuple with the TcpSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthProbe) GetTcpSocketOk() (*TcpSocket, bool) {
	if o == nil || utils.IsNil(o.tcpSocket) {
		return nil, false
	}
	return o.tcpSocket, true
}

// HasTcpSocket returns a boolean if a field has been set.
func (o *HealthProbe) HasTcpSocket() bool {
	if o != nil && !utils.IsNil(o.tcpSocket) {
		return true
	}

	return false
}

// SetTcpSocket gets a reference to the given TcpSocket and assigns it to the tcpSocket field.
// tcpSocket: 

func (o *HealthProbe) TcpSocket(v TcpSocket) (*HealthProbe){
	o.tcpSocket = &v
return o
}

// GetTimeoutSeconds returns the TimeoutSeconds field value
func (o *HealthProbe) GetTimeoutSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.timeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value
// and a boolean to check if the value has been set.
func (o *HealthProbe) GetTimeoutSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.timeoutSeconds, true
}

// TimeoutSeconds sets field value
func (o *HealthProbe) TimeoutSeconds(v int32) *HealthProbe {
	o.timeoutSeconds = v
    return o
}

func (o HealthProbe) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthProbe) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.exec) {
		toSerialize["exec"] = o.exec
	}
	toSerialize["failureThreshold"] = o.failureThreshold
	if !utils.IsNil(o.httpGet) {
		toSerialize["httpGet"] = o.httpGet
	}
	toSerialize["initialDelaySeconds"] = o.initialDelaySeconds
	toSerialize["periodSeconds"] = o.periodSeconds
	toSerialize["successThreshold"] = o.successThreshold
	if !utils.IsNil(o.tcpSocket) {
		toSerialize["tcpSocket"] = o.tcpSocket
	}
	toSerialize["timeoutSeconds"] = o.timeoutSeconds
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

 
