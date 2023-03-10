/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package startup_probe

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the StartupProbeParams type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &StartupProbeParams{}

// StartupProbeParams struct for StartupProbeParams
type StartupProbeParams struct {
	// Specify the name of the target container, if not set, use the component name
	ContainerName *string `json:"containerName,omitempty"`
	Exec          *Exec   `json:"exec,omitempty"`
	// Minimum consecutive failures for the probe to be considered failed after having succeeded. Minimum value is 1.
	FailureThreshold *int32   `json:"failureThreshold,omitempty"`
	Grpc             *Grpc    `json:"grpc,omitempty"`
	HttpGet          *HttpGet `json:"httpGet,omitempty"`
	// Number of seconds after the container has started before liveness probes are initiated. Minimum value is 0.
	InitialDelaySeconds *int32 `json:"initialDelaySeconds,omitempty"`
	// How often, in seconds, to execute the probe. Minimum value is 1.
	PeriodSeconds *int32 `json:"periodSeconds,omitempty"`
	// Minimum consecutive successes for the probe to be considered successful after having failed.  Minimum value is 1.
	SuccessThreshold *int32     `json:"successThreshold,omitempty"`
	TcpSocket        *TcpSocket `json:"tcpSocket,omitempty"`
	// Optional duration in seconds the pod needs to terminate gracefully upon probe failure. Set this value longer than the expected cleanup time for your process.
	TerminationGracePeriodSeconds *int32 `json:"terminationGracePeriodSeconds,omitempty"`
	// Number of seconds after which the probe times out. Minimum value is 1.
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`
}

// NewStartupProbeParamsWith instantiates a new StartupProbeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartupProbeParamsWith() *StartupProbeParams {
	this := StartupProbeParams{}
	var containerName string = ""
	this.ContainerName = &containerName
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

// NewStartupProbeParams instantiates a new StartupProbeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartupProbeParams() *StartupProbeParams {
	this := StartupProbeParams{}
	var containerName string = ""
	this.ContainerName = &containerName
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

// NewStartupProbeParamss converts a list StartupProbeParams pointers to objects.
// This is helpful when the SetStartupProbeParams requires a list of objects
func NewStartupProbeParamss(ps ...*StartupProbeParams) []StartupProbeParams {
	objs := []StartupProbeParams{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetContainerName returns the ContainerName field value if set, zero value otherwise.
func (o *StartupProbeParams) GetContainerName() string {
	if o == nil || utils.IsNil(o.ContainerName) {
		var ret string
		return ret
	}
	return *o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetContainerNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ContainerName) {
		return nil, false
	}
	return o.ContainerName, true
}

// HasContainerName returns a boolean if a field has been set.
func (o *StartupProbeParams) HasContainerName() bool {
	if o != nil && !utils.IsNil(o.ContainerName) {
		return true
	}

	return false
}

// SetContainerName gets a reference to the given string and assigns it to the containerName field.
// ContainerName:  Specify the name of the target container, if not set, use the component name
func (o *StartupProbeParams) SetContainerName(v string) *StartupProbeParams {
	o.ContainerName = &v
	return o
}

// GetExec returns the Exec field value if set, zero value otherwise.
func (o *StartupProbeParams) GetExec() Exec {
	if o == nil || utils.IsNil(o.Exec) {
		var ret Exec
		return ret
	}
	return *o.Exec
}

// GetExecOk returns a tuple with the Exec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetExecOk() (*Exec, bool) {
	if o == nil || utils.IsNil(o.Exec) {
		return nil, false
	}
	return o.Exec, true
}

// HasExec returns a boolean if a field has been set.
func (o *StartupProbeParams) HasExec() bool {
	if o != nil && !utils.IsNil(o.Exec) {
		return true
	}

	return false
}

// SetExec gets a reference to the given Exec and assigns it to the exec field.
// Exec:
func (o *StartupProbeParams) SetExec(v Exec) *StartupProbeParams {
	o.Exec = &v
	return o
}

// GetFailureThreshold returns the FailureThreshold field value if set, zero value otherwise.
func (o *StartupProbeParams) GetFailureThreshold() int32 {
	if o == nil || utils.IsNil(o.FailureThreshold) {
		var ret int32
		return ret
	}
	return *o.FailureThreshold
}

// GetFailureThresholdOk returns a tuple with the FailureThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetFailureThresholdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.FailureThreshold) {
		return nil, false
	}
	return o.FailureThreshold, true
}

// HasFailureThreshold returns a boolean if a field has been set.
func (o *StartupProbeParams) HasFailureThreshold() bool {
	if o != nil && !utils.IsNil(o.FailureThreshold) {
		return true
	}

	return false
}

// SetFailureThreshold gets a reference to the given int32 and assigns it to the failureThreshold field.
// FailureThreshold:  Minimum consecutive failures for the probe to be considered failed after having succeeded. Minimum value is 1.
func (o *StartupProbeParams) SetFailureThreshold(v int32) *StartupProbeParams {
	o.FailureThreshold = &v
	return o
}

// GetGrpc returns the Grpc field value if set, zero value otherwise.
func (o *StartupProbeParams) GetGrpc() Grpc {
	if o == nil || utils.IsNil(o.Grpc) {
		var ret Grpc
		return ret
	}
	return *o.Grpc
}

// GetGrpcOk returns a tuple with the Grpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetGrpcOk() (*Grpc, bool) {
	if o == nil || utils.IsNil(o.Grpc) {
		return nil, false
	}
	return o.Grpc, true
}

// HasGrpc returns a boolean if a field has been set.
func (o *StartupProbeParams) HasGrpc() bool {
	if o != nil && !utils.IsNil(o.Grpc) {
		return true
	}

	return false
}

// SetGrpc gets a reference to the given Grpc and assigns it to the grpc field.
// Grpc:
func (o *StartupProbeParams) SetGrpc(v Grpc) *StartupProbeParams {
	o.Grpc = &v
	return o
}

// GetHttpGet returns the HttpGet field value if set, zero value otherwise.
func (o *StartupProbeParams) GetHttpGet() HttpGet {
	if o == nil || utils.IsNil(o.HttpGet) {
		var ret HttpGet
		return ret
	}
	return *o.HttpGet
}

// GetHttpGetOk returns a tuple with the HttpGet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetHttpGetOk() (*HttpGet, bool) {
	if o == nil || utils.IsNil(o.HttpGet) {
		return nil, false
	}
	return o.HttpGet, true
}

// HasHttpGet returns a boolean if a field has been set.
func (o *StartupProbeParams) HasHttpGet() bool {
	if o != nil && !utils.IsNil(o.HttpGet) {
		return true
	}

	return false
}

// SetHttpGet gets a reference to the given HttpGet and assigns it to the httpGet field.
// HttpGet:
func (o *StartupProbeParams) SetHttpGet(v HttpGet) *StartupProbeParams {
	o.HttpGet = &v
	return o
}

// GetInitialDelaySeconds returns the InitialDelaySeconds field value if set, zero value otherwise.
func (o *StartupProbeParams) GetInitialDelaySeconds() int32 {
	if o == nil || utils.IsNil(o.InitialDelaySeconds) {
		var ret int32
		return ret
	}
	return *o.InitialDelaySeconds
}

// GetInitialDelaySecondsOk returns a tuple with the InitialDelaySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetInitialDelaySecondsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.InitialDelaySeconds) {
		return nil, false
	}
	return o.InitialDelaySeconds, true
}

// HasInitialDelaySeconds returns a boolean if a field has been set.
func (o *StartupProbeParams) HasInitialDelaySeconds() bool {
	if o != nil && !utils.IsNil(o.InitialDelaySeconds) {
		return true
	}

	return false
}

// SetInitialDelaySeconds gets a reference to the given int32 and assigns it to the initialDelaySeconds field.
// InitialDelaySeconds:  Number of seconds after the container has started before liveness probes are initiated. Minimum value is 0.
func (o *StartupProbeParams) SetInitialDelaySeconds(v int32) *StartupProbeParams {
	o.InitialDelaySeconds = &v
	return o
}

// GetPeriodSeconds returns the PeriodSeconds field value if set, zero value otherwise.
func (o *StartupProbeParams) GetPeriodSeconds() int32 {
	if o == nil || utils.IsNil(o.PeriodSeconds) {
		var ret int32
		return ret
	}
	return *o.PeriodSeconds
}

// GetPeriodSecondsOk returns a tuple with the PeriodSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetPeriodSecondsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PeriodSeconds) {
		return nil, false
	}
	return o.PeriodSeconds, true
}

// HasPeriodSeconds returns a boolean if a field has been set.
func (o *StartupProbeParams) HasPeriodSeconds() bool {
	if o != nil && !utils.IsNil(o.PeriodSeconds) {
		return true
	}

	return false
}

// SetPeriodSeconds gets a reference to the given int32 and assigns it to the periodSeconds field.
// PeriodSeconds:  How often, in seconds, to execute the probe. Minimum value is 1.
func (o *StartupProbeParams) SetPeriodSeconds(v int32) *StartupProbeParams {
	o.PeriodSeconds = &v
	return o
}

// GetSuccessThreshold returns the SuccessThreshold field value if set, zero value otherwise.
func (o *StartupProbeParams) GetSuccessThreshold() int32 {
	if o == nil || utils.IsNil(o.SuccessThreshold) {
		var ret int32
		return ret
	}
	return *o.SuccessThreshold
}

// GetSuccessThresholdOk returns a tuple with the SuccessThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetSuccessThresholdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.SuccessThreshold) {
		return nil, false
	}
	return o.SuccessThreshold, true
}

// HasSuccessThreshold returns a boolean if a field has been set.
func (o *StartupProbeParams) HasSuccessThreshold() bool {
	if o != nil && !utils.IsNil(o.SuccessThreshold) {
		return true
	}

	return false
}

// SetSuccessThreshold gets a reference to the given int32 and assigns it to the successThreshold field.
// SuccessThreshold:  Minimum consecutive successes for the probe to be considered successful after having failed.  Minimum value is 1.
func (o *StartupProbeParams) SetSuccessThreshold(v int32) *StartupProbeParams {
	o.SuccessThreshold = &v
	return o
}

// GetTcpSocket returns the TcpSocket field value if set, zero value otherwise.
func (o *StartupProbeParams) GetTcpSocket() TcpSocket {
	if o == nil || utils.IsNil(o.TcpSocket) {
		var ret TcpSocket
		return ret
	}
	return *o.TcpSocket
}

// GetTcpSocketOk returns a tuple with the TcpSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetTcpSocketOk() (*TcpSocket, bool) {
	if o == nil || utils.IsNil(o.TcpSocket) {
		return nil, false
	}
	return o.TcpSocket, true
}

// HasTcpSocket returns a boolean if a field has been set.
func (o *StartupProbeParams) HasTcpSocket() bool {
	if o != nil && !utils.IsNil(o.TcpSocket) {
		return true
	}

	return false
}

// SetTcpSocket gets a reference to the given TcpSocket and assigns it to the tcpSocket field.
// TcpSocket:
func (o *StartupProbeParams) SetTcpSocket(v TcpSocket) *StartupProbeParams {
	o.TcpSocket = &v
	return o
}

// GetTerminationGracePeriodSeconds returns the TerminationGracePeriodSeconds field value if set, zero value otherwise.
func (o *StartupProbeParams) GetTerminationGracePeriodSeconds() int32 {
	if o == nil || utils.IsNil(o.TerminationGracePeriodSeconds) {
		var ret int32
		return ret
	}
	return *o.TerminationGracePeriodSeconds
}

// GetTerminationGracePeriodSecondsOk returns a tuple with the TerminationGracePeriodSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetTerminationGracePeriodSecondsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TerminationGracePeriodSeconds) {
		return nil, false
	}
	return o.TerminationGracePeriodSeconds, true
}

// HasTerminationGracePeriodSeconds returns a boolean if a field has been set.
func (o *StartupProbeParams) HasTerminationGracePeriodSeconds() bool {
	if o != nil && !utils.IsNil(o.TerminationGracePeriodSeconds) {
		return true
	}

	return false
}

// SetTerminationGracePeriodSeconds gets a reference to the given int32 and assigns it to the terminationGracePeriodSeconds field.
// TerminationGracePeriodSeconds:  Optional duration in seconds the pod needs to terminate gracefully upon probe failure. Set this value longer than the expected cleanup time for your process.
func (o *StartupProbeParams) SetTerminationGracePeriodSeconds(v int32) *StartupProbeParams {
	o.TerminationGracePeriodSeconds = &v
	return o
}

// GetTimeoutSeconds returns the TimeoutSeconds field value if set, zero value otherwise.
func (o *StartupProbeParams) GetTimeoutSeconds() int32 {
	if o == nil || utils.IsNil(o.TimeoutSeconds) {
		var ret int32
		return ret
	}
	return *o.TimeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetTimeoutSecondsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TimeoutSeconds) {
		return nil, false
	}
	return o.TimeoutSeconds, true
}

// HasTimeoutSeconds returns a boolean if a field has been set.
func (o *StartupProbeParams) HasTimeoutSeconds() bool {
	if o != nil && !utils.IsNil(o.TimeoutSeconds) {
		return true
	}

	return false
}

// SetTimeoutSeconds gets a reference to the given int32 and assigns it to the timeoutSeconds field.
// TimeoutSeconds:  Number of seconds after which the probe times out. Minimum value is 1.
func (o *StartupProbeParams) SetTimeoutSeconds(v int32) *StartupProbeParams {
	o.TimeoutSeconds = &v
	return o
}

func (o StartupProbeParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartupProbeParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ContainerName) {
		toSerialize["containerName"] = o.ContainerName
	}
	if !utils.IsNil(o.Exec) {
		toSerialize["exec"] = o.Exec
	}
	if !utils.IsNil(o.FailureThreshold) {
		toSerialize["failureThreshold"] = o.FailureThreshold
	}
	if !utils.IsNil(o.Grpc) {
		toSerialize["grpc"] = o.Grpc
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
	if !utils.IsNil(o.TerminationGracePeriodSeconds) {
		toSerialize["terminationGracePeriodSeconds"] = o.TerminationGracePeriodSeconds
	}
	if !utils.IsNil(o.TimeoutSeconds) {
		toSerialize["timeoutSeconds"] = o.TimeoutSeconds
	}
	return toSerialize, nil
}

type NullableStartupProbeParams struct {
	value *StartupProbeParams
	isSet bool
}

func (v NullableStartupProbeParams) Get() *StartupProbeParams {
	return v.value
}

func (v *NullableStartupProbeParams) Set(val *StartupProbeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableStartupProbeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableStartupProbeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartupProbeParams(val *StartupProbeParams) *NullableStartupProbeParams {
	return &NullableStartupProbeParams{value: val, isSet: true}
}

func (v NullableStartupProbeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartupProbeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
