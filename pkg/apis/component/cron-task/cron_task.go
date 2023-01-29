/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cron_task

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"vela-go-sdk/pkg/apis"
	"vela-go-sdk/pkg/apis/utils"
)

// checks if the CronTaskSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CronTaskSpec{}

// CronTaskSpec struct for CronTaskSpec
type CronTaskSpec struct {
	// The duration in seconds relative to the startTime that the job may be continuously active before the system tries to terminate it
	activeDeadlineSeconds *int32 `json:"activeDeadlineSeconds,omitempty"`
	// Specify the annotations in the workload
	annotations *map[string]string `json:"annotations,omitempty"`
	// The number of retries before marking this job failed
	backoffLimit int32 `json:"backoffLimit"`
	// Commands to run in the container
	cmd []string `json:"cmd,omitempty"`
	// Specifies how to treat concurrent executions of a Job
	concurrencyPolicy string `json:"concurrencyPolicy"`
	// Specify number of tasks to run in parallel +short=c
	count int32 `json:"count"`
	// Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
	cpu *string `json:"cpu,omitempty"`
	// Define arguments by using environment variables
	env []Env `json:"env,omitempty"`
	// The number of failed finished jobs to retain
	failedJobsHistoryLimit int32 `json:"failedJobsHistoryLimit"`
	// An optional list of hosts and IPs that will be injected into the pod's hosts file
	hostAliases []HostAliases `json:"hostAliases,omitempty"`
	// Which image would you like to use for your service +short=i
	image string `json:"image"`
	// Specify image pull policy for your service
	imagePullPolicy *string `json:"imagePullPolicy,omitempty"`
	// Specify image pull secrets for your service
	imagePullSecrets []string `json:"imagePullSecrets,omitempty"`
	// Specify the labels in the workload
	labels        *map[string]string `json:"labels,omitempty"`
	livenessProbe *HealthProbe       `json:"livenessProbe,omitempty"`
	// Specifies the attributes of the memory resource required for the container.
	memory         *string      `json:"memory,omitempty"`
	readinessProbe *HealthProbe `json:"readinessProbe,omitempty"`
	// Define the job restart policy, the value can only be Never or OnFailure. By default, it's Never.
	restart string `json:"restart"`
	// Specify the schedule in Cron format, see https://en.wikipedia.org/wiki/Cron
	schedule string `json:"schedule"`
	// Specify deadline in seconds for starting the job if it misses scheduled
	startingDeadlineSeconds *int32 `json:"startingDeadlineSeconds,omitempty"`
	// The number of successful finished jobs to retain
	successfulJobsHistoryLimit int32 `json:"successfulJobsHistoryLimit"`
	// suspend subsequent executions
	suspend bool `json:"suspend"`
	// Limits the lifetime of a Job that has finished
	ttlSecondsAfterFinished *int32 `json:"ttlSecondsAfterFinished,omitempty"`
	// Declare volumes and volumeMounts
	volumes []Volumes `json:"volumes,omitempty"`
}

// NewCronTaskSpecWith instantiates a new CronTaskSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCronTaskSpecWith(backoffLimit int32, concurrencyPolicy string, count int32, failedJobsHistoryLimit int32, image string, restart string, schedule string, successfulJobsHistoryLimit int32, suspend bool) *CronTaskSpec {
	this := CronTaskSpec{}
	this.backoffLimit = backoffLimit
	this.concurrencyPolicy = concurrencyPolicy
	this.count = count
	this.failedJobsHistoryLimit = failedJobsHistoryLimit
	this.image = image
	this.restart = restart
	this.schedule = schedule
	this.successfulJobsHistoryLimit = successfulJobsHistoryLimit
	this.suspend = suspend
	return &this
}

// NewCronTaskSpec instantiates a new CronTaskSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCronTaskSpec() *CronTaskSpec {
	this := CronTaskSpec{}
	var backoffLimit int32 = 6
	this.backoffLimit = backoffLimit
	var concurrencyPolicy string = "Allow"
	this.concurrencyPolicy = concurrencyPolicy
	var count int32 = 1
	this.count = count
	var failedJobsHistoryLimit int32 = 1
	this.failedJobsHistoryLimit = failedJobsHistoryLimit
	var restart string = "Never"
	this.restart = restart
	var successfulJobsHistoryLimit int32 = 3
	this.successfulJobsHistoryLimit = successfulJobsHistoryLimit
	var suspend bool = false
	this.suspend = suspend
	return &this
}

// GetActiveDeadlineSeconds returns the ActiveDeadlineSeconds field value if set, zero value otherwise.
func (o *CronTaskComponent) GetActiveDeadlineSeconds() int32 {
	if o == nil || utils.IsNil(o.Properties.activeDeadlineSeconds) {
		var ret int32
		return ret
	}
	return *o.Properties.activeDeadlineSeconds
}

// GetActiveDeadlineSecondsOk returns a tuple with the ActiveDeadlineSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetActiveDeadlineSecondsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.activeDeadlineSeconds) {
		return nil, false
	}
	return o.Properties.activeDeadlineSeconds, true
}

// HasActiveDeadlineSeconds returns a boolean if a field has been set.
func (o *CronTaskComponent) HasActiveDeadlineSeconds() bool {
	if o != nil && !utils.IsNil(o.Properties.activeDeadlineSeconds) {
		return true
	}

	return false
}

// ActiveDeadlineSeconds gets a reference to the given int32 and assigns it to the activeDeadlineSeconds field.
// activeDeadlineSeconds:  The duration in seconds relative to the startTime that the job may be continuously active before the system tries to terminate it
func (o *CronTaskComponent) ActiveDeadlineSeconds(v int32) *CronTaskComponent {
	o.Properties.activeDeadlineSeconds = &v
	return o
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *CronTaskComponent) GetAnnotations() map[string]string {
	if o == nil || utils.IsNil(o.Properties.annotations) {
		var ret map[string]string
		return ret
	}
	return *o.Properties.annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.Properties.annotations) {
		return nil, false
	}
	return o.Properties.annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *CronTaskComponent) HasAnnotations() bool {
	if o != nil && !utils.IsNil(o.Properties.annotations) {
		return true
	}

	return false
}

// Annotations gets a reference to the given map[string]string and assigns it to the annotations field.
// annotations:  Specify the annotations in the workload
func (o *CronTaskComponent) Annotations(v map[string]string) *CronTaskComponent {
	o.Properties.annotations = &v
	return o
}

// GetBackoffLimit returns the BackoffLimit field value
func (o *CronTaskComponent) GetBackoffLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Properties.backoffLimit
}

// GetBackoffLimitOk returns a tuple with the BackoffLimit field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetBackoffLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.backoffLimit, true
}

// BackoffLimit sets field value
func (o *CronTaskComponent) BackoffLimit(v int32) *CronTaskComponent {
	o.Properties.backoffLimit = v
	return o
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *CronTaskComponent) GetCmd() []string {
	if o == nil || utils.IsNil(o.Properties.cmd) {
		var ret []string
		return ret
	}
	return o.Properties.cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetCmdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.cmd) {
		return nil, false
	}
	return o.Properties.cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *CronTaskComponent) HasCmd() bool {
	if o != nil && !utils.IsNil(o.Properties.cmd) {
		return true
	}

	return false
}

// Cmd gets a reference to the given []string and assigns it to the cmd field.
// cmd:  Commands to run in the container
func (o *CronTaskComponent) Cmd(v []string) *CronTaskComponent {
	o.Properties.cmd = v
	return o
}

// GetConcurrencyPolicy returns the ConcurrencyPolicy field value
func (o *CronTaskComponent) GetConcurrencyPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.concurrencyPolicy
}

// GetConcurrencyPolicyOk returns a tuple with the ConcurrencyPolicy field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetConcurrencyPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.concurrencyPolicy, true
}

// ConcurrencyPolicy sets field value
func (o *CronTaskComponent) ConcurrencyPolicy(v string) *CronTaskComponent {
	o.Properties.concurrencyPolicy = v
	return o
}

// GetCount returns the Count field value
func (o *CronTaskComponent) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Properties.count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.count, true
}

// Count sets field value
func (o *CronTaskComponent) Count(v int32) *CronTaskComponent {
	o.Properties.count = v
	return o
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *CronTaskComponent) GetCpu() string {
	if o == nil || utils.IsNil(o.Properties.cpu) {
		var ret string
		return ret
	}
	return *o.Properties.cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetCpuOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.cpu) {
		return nil, false
	}
	return o.Properties.cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *CronTaskComponent) HasCpu() bool {
	if o != nil && !utils.IsNil(o.Properties.cpu) {
		return true
	}

	return false
}

// Cpu gets a reference to the given string and assigns it to the cpu field.
// cpu:  Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
func (o *CronTaskComponent) Cpu(v string) *CronTaskComponent {
	o.Properties.cpu = &v
	return o
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *CronTaskComponent) GetEnv() []Env {
	if o == nil || utils.IsNil(o.Properties.env) {
		var ret []Env
		return ret
	}
	return o.Properties.env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetEnvOk() ([]Env, bool) {
	if o == nil || utils.IsNil(o.Properties.env) {
		return nil, false
	}
	return o.Properties.env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *CronTaskComponent) HasEnv() bool {
	if o != nil && !utils.IsNil(o.Properties.env) {
		return true
	}

	return false
}

// Env gets a reference to the given []Env and assigns it to the env field.
// env:  Define arguments by using environment variables
func (o *CronTaskComponent) Env(v []Env) *CronTaskComponent {
	o.Properties.env = v
	return o
}

// GetFailedJobsHistoryLimit returns the FailedJobsHistoryLimit field value
func (o *CronTaskComponent) GetFailedJobsHistoryLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Properties.failedJobsHistoryLimit
}

// GetFailedJobsHistoryLimitOk returns a tuple with the FailedJobsHistoryLimit field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetFailedJobsHistoryLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.failedJobsHistoryLimit, true
}

// FailedJobsHistoryLimit sets field value
func (o *CronTaskComponent) FailedJobsHistoryLimit(v int32) *CronTaskComponent {
	o.Properties.failedJobsHistoryLimit = v
	return o
}

// GetHostAliases returns the HostAliases field value if set, zero value otherwise.
func (o *CronTaskComponent) GetHostAliases() []HostAliases {
	if o == nil || utils.IsNil(o.Properties.hostAliases) {
		var ret []HostAliases
		return ret
	}
	return o.Properties.hostAliases
}

// GetHostAliasesOk returns a tuple with the HostAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetHostAliasesOk() ([]HostAliases, bool) {
	if o == nil || utils.IsNil(o.Properties.hostAliases) {
		return nil, false
	}
	return o.Properties.hostAliases, true
}

// HasHostAliases returns a boolean if a field has been set.
func (o *CronTaskComponent) HasHostAliases() bool {
	if o != nil && !utils.IsNil(o.Properties.hostAliases) {
		return true
	}

	return false
}

// HostAliases gets a reference to the given []HostAliases and assigns it to the hostAliases field.
// hostAliases:  An optional list of hosts and IPs that will be injected into the pod's hosts file
func (o *CronTaskComponent) HostAliases(v []HostAliases) *CronTaskComponent {
	o.Properties.hostAliases = v
	return o
}

// GetImage returns the Image field value
func (o *CronTaskComponent) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.image, true
}

// Image sets field value
func (o *CronTaskComponent) Image(v string) *CronTaskComponent {
	o.Properties.image = v
	return o
}

// GetImagePullPolicy returns the ImagePullPolicy field value if set, zero value otherwise.
func (o *CronTaskComponent) GetImagePullPolicy() string {
	if o == nil || utils.IsNil(o.Properties.imagePullPolicy) {
		var ret string
		return ret
	}
	return *o.Properties.imagePullPolicy
}

// GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetImagePullPolicyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.imagePullPolicy) {
		return nil, false
	}
	return o.Properties.imagePullPolicy, true
}

// HasImagePullPolicy returns a boolean if a field has been set.
func (o *CronTaskComponent) HasImagePullPolicy() bool {
	if o != nil && !utils.IsNil(o.Properties.imagePullPolicy) {
		return true
	}

	return false
}

// ImagePullPolicy gets a reference to the given string and assigns it to the imagePullPolicy field.
// imagePullPolicy:  Specify image pull policy for your service
func (o *CronTaskComponent) ImagePullPolicy(v string) *CronTaskComponent {
	o.Properties.imagePullPolicy = &v
	return o
}

// GetImagePullSecrets returns the ImagePullSecrets field value if set, zero value otherwise.
func (o *CronTaskComponent) GetImagePullSecrets() []string {
	if o == nil || utils.IsNil(o.Properties.imagePullSecrets) {
		var ret []string
		return ret
	}
	return o.Properties.imagePullSecrets
}

// GetImagePullSecretsOk returns a tuple with the ImagePullSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetImagePullSecretsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.imagePullSecrets) {
		return nil, false
	}
	return o.Properties.imagePullSecrets, true
}

// HasImagePullSecrets returns a boolean if a field has been set.
func (o *CronTaskComponent) HasImagePullSecrets() bool {
	if o != nil && !utils.IsNil(o.Properties.imagePullSecrets) {
		return true
	}

	return false
}

// ImagePullSecrets gets a reference to the given []string and assigns it to the imagePullSecrets field.
// imagePullSecrets:  Specify image pull secrets for your service
func (o *CronTaskComponent) ImagePullSecrets(v []string) *CronTaskComponent {
	o.Properties.imagePullSecrets = v
	return o
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CronTaskComponent) GetLabels() map[string]string {
	if o == nil || utils.IsNil(o.Properties.labels) {
		var ret map[string]string
		return ret
	}
	return *o.Properties.labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.Properties.labels) {
		return nil, false
	}
	return o.Properties.labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CronTaskComponent) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Properties.labels) {
		return true
	}

	return false
}

// Labels gets a reference to the given map[string]string and assigns it to the labels field.
// labels:  Specify the labels in the workload
func (o *CronTaskComponent) Labels(v map[string]string) *CronTaskComponent {
	o.Properties.labels = &v
	return o
}

// GetLivenessProbe returns the LivenessProbe field value if set, zero value otherwise.
func (o *CronTaskComponent) GetLivenessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.Properties.livenessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.Properties.livenessProbe
}

// GetLivenessProbeOk returns a tuple with the LivenessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetLivenessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.Properties.livenessProbe) {
		return nil, false
	}
	return o.Properties.livenessProbe, true
}

// HasLivenessProbe returns a boolean if a field has been set.
func (o *CronTaskComponent) HasLivenessProbe() bool {
	if o != nil && !utils.IsNil(o.Properties.livenessProbe) {
		return true
	}

	return false
}

// LivenessProbe gets a reference to the given HealthProbe and assigns it to the livenessProbe field.
// livenessProbe:
func (o *CronTaskComponent) LivenessProbe(v HealthProbe) *CronTaskComponent {
	o.Properties.livenessProbe = &v
	return o
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *CronTaskComponent) GetMemory() string {
	if o == nil || utils.IsNil(o.Properties.memory) {
		var ret string
		return ret
	}
	return *o.Properties.memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetMemoryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.memory) {
		return nil, false
	}
	return o.Properties.memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *CronTaskComponent) HasMemory() bool {
	if o != nil && !utils.IsNil(o.Properties.memory) {
		return true
	}

	return false
}

// Memory gets a reference to the given string and assigns it to the memory field.
// memory:  Specifies the attributes of the memory resource required for the container.
func (o *CronTaskComponent) Memory(v string) *CronTaskComponent {
	o.Properties.memory = &v
	return o
}

// GetReadinessProbe returns the ReadinessProbe field value if set, zero value otherwise.
func (o *CronTaskComponent) GetReadinessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.Properties.readinessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.Properties.readinessProbe
}

// GetReadinessProbeOk returns a tuple with the ReadinessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetReadinessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.Properties.readinessProbe) {
		return nil, false
	}
	return o.Properties.readinessProbe, true
}

// HasReadinessProbe returns a boolean if a field has been set.
func (o *CronTaskComponent) HasReadinessProbe() bool {
	if o != nil && !utils.IsNil(o.Properties.readinessProbe) {
		return true
	}

	return false
}

// ReadinessProbe gets a reference to the given HealthProbe and assigns it to the readinessProbe field.
// readinessProbe:
func (o *CronTaskComponent) ReadinessProbe(v HealthProbe) *CronTaskComponent {
	o.Properties.readinessProbe = &v
	return o
}

// GetRestart returns the Restart field value
func (o *CronTaskComponent) GetRestart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.restart
}

// GetRestartOk returns a tuple with the Restart field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetRestartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.restart, true
}

// Restart sets field value
func (o *CronTaskComponent) Restart(v string) *CronTaskComponent {
	o.Properties.restart = v
	return o
}

// GetSchedule returns the Schedule field value
func (o *CronTaskComponent) GetSchedule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetScheduleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.schedule, true
}

// Schedule sets field value
func (o *CronTaskComponent) Schedule(v string) *CronTaskComponent {
	o.Properties.schedule = v
	return o
}

// GetStartingDeadlineSeconds returns the StartingDeadlineSeconds field value if set, zero value otherwise.
func (o *CronTaskComponent) GetStartingDeadlineSeconds() int32 {
	if o == nil || utils.IsNil(o.Properties.startingDeadlineSeconds) {
		var ret int32
		return ret
	}
	return *o.Properties.startingDeadlineSeconds
}

// GetStartingDeadlineSecondsOk returns a tuple with the StartingDeadlineSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetStartingDeadlineSecondsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.startingDeadlineSeconds) {
		return nil, false
	}
	return o.Properties.startingDeadlineSeconds, true
}

// HasStartingDeadlineSeconds returns a boolean if a field has been set.
func (o *CronTaskComponent) HasStartingDeadlineSeconds() bool {
	if o != nil && !utils.IsNil(o.Properties.startingDeadlineSeconds) {
		return true
	}

	return false
}

// StartingDeadlineSeconds gets a reference to the given int32 and assigns it to the startingDeadlineSeconds field.
// startingDeadlineSeconds:  Specify deadline in seconds for starting the job if it misses scheduled
func (o *CronTaskComponent) StartingDeadlineSeconds(v int32) *CronTaskComponent {
	o.Properties.startingDeadlineSeconds = &v
	return o
}

// GetSuccessfulJobsHistoryLimit returns the SuccessfulJobsHistoryLimit field value
func (o *CronTaskComponent) GetSuccessfulJobsHistoryLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Properties.successfulJobsHistoryLimit
}

// GetSuccessfulJobsHistoryLimitOk returns a tuple with the SuccessfulJobsHistoryLimit field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetSuccessfulJobsHistoryLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.successfulJobsHistoryLimit, true
}

// SuccessfulJobsHistoryLimit sets field value
func (o *CronTaskComponent) SuccessfulJobsHistoryLimit(v int32) *CronTaskComponent {
	o.Properties.successfulJobsHistoryLimit = v
	return o
}

// GetSuspend returns the Suspend field value
func (o *CronTaskComponent) GetSuspend() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Properties.suspend
}

// GetSuspendOk returns a tuple with the Suspend field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetSuspendOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.suspend, true
}

// Suspend sets field value
func (o *CronTaskComponent) Suspend(v bool) *CronTaskComponent {
	o.Properties.suspend = v
	return o
}

// GetTtlSecondsAfterFinished returns the TtlSecondsAfterFinished field value if set, zero value otherwise.
func (o *CronTaskComponent) GetTtlSecondsAfterFinished() int32 {
	if o == nil || utils.IsNil(o.Properties.ttlSecondsAfterFinished) {
		var ret int32
		return ret
	}
	return *o.Properties.ttlSecondsAfterFinished
}

// GetTtlSecondsAfterFinishedOk returns a tuple with the TtlSecondsAfterFinished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetTtlSecondsAfterFinishedOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.ttlSecondsAfterFinished) {
		return nil, false
	}
	return o.Properties.ttlSecondsAfterFinished, true
}

// HasTtlSecondsAfterFinished returns a boolean if a field has been set.
func (o *CronTaskComponent) HasTtlSecondsAfterFinished() bool {
	if o != nil && !utils.IsNil(o.Properties.ttlSecondsAfterFinished) {
		return true
	}

	return false
}

// TtlSecondsAfterFinished gets a reference to the given int32 and assigns it to the ttlSecondsAfterFinished field.
// ttlSecondsAfterFinished:  Limits the lifetime of a Job that has finished
func (o *CronTaskComponent) TtlSecondsAfterFinished(v int32) *CronTaskComponent {
	o.Properties.ttlSecondsAfterFinished = &v
	return o
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *CronTaskComponent) GetVolumes() []Volumes {
	if o == nil || utils.IsNil(o.Properties.volumes) {
		var ret []Volumes
		return ret
	}
	return o.Properties.volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetVolumesOk() ([]Volumes, bool) {
	if o == nil || utils.IsNil(o.Properties.volumes) {
		return nil, false
	}
	return o.Properties.volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *CronTaskComponent) HasVolumes() bool {
	if o != nil && !utils.IsNil(o.Properties.volumes) {
		return true
	}

	return false
}

// Volumes gets a reference to the given []Volumes and assigns it to the volumes field.
// volumes:  Declare volumes and volumeMounts
func (o *CronTaskComponent) Volumes(v []Volumes) *CronTaskComponent {
	o.Properties.volumes = v
	return o
}

func (o CronTaskSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CronTaskSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.activeDeadlineSeconds) {
		toSerialize["activeDeadlineSeconds"] = o.activeDeadlineSeconds
	}
	if !utils.IsNil(o.annotations) {
		toSerialize["annotations"] = o.annotations
	}
	toSerialize["backoffLimit"] = o.backoffLimit
	if !utils.IsNil(o.cmd) {
		toSerialize["cmd"] = o.cmd
	}
	toSerialize["concurrencyPolicy"] = o.concurrencyPolicy
	toSerialize["count"] = o.count
	if !utils.IsNil(o.cpu) {
		toSerialize["cpu"] = o.cpu
	}
	if !utils.IsNil(o.env) {
		toSerialize["env"] = o.env
	}
	toSerialize["failedJobsHistoryLimit"] = o.failedJobsHistoryLimit
	if !utils.IsNil(o.hostAliases) {
		toSerialize["hostAliases"] = o.hostAliases
	}
	toSerialize["image"] = o.image
	if !utils.IsNil(o.imagePullPolicy) {
		toSerialize["imagePullPolicy"] = o.imagePullPolicy
	}
	if !utils.IsNil(o.imagePullSecrets) {
		toSerialize["imagePullSecrets"] = o.imagePullSecrets
	}
	if !utils.IsNil(o.labels) {
		toSerialize["labels"] = o.labels
	}
	if !utils.IsNil(o.livenessProbe) {
		toSerialize["livenessProbe"] = o.livenessProbe
	}
	if !utils.IsNil(o.memory) {
		toSerialize["memory"] = o.memory
	}
	if !utils.IsNil(o.readinessProbe) {
		toSerialize["readinessProbe"] = o.readinessProbe
	}
	toSerialize["restart"] = o.restart
	toSerialize["schedule"] = o.schedule
	if !utils.IsNil(o.startingDeadlineSeconds) {
		toSerialize["startingDeadlineSeconds"] = o.startingDeadlineSeconds
	}
	toSerialize["successfulJobsHistoryLimit"] = o.successfulJobsHistoryLimit
	toSerialize["suspend"] = o.suspend
	if !utils.IsNil(o.ttlSecondsAfterFinished) {
		toSerialize["ttlSecondsAfterFinished"] = o.ttlSecondsAfterFinished
	}
	if !utils.IsNil(o.volumes) {
		toSerialize["volumes"] = o.volumes
	}
	return toSerialize, nil
}

type NullableCronTaskSpec struct {
	value *CronTaskSpec
	isSet bool
}

func (v NullableCronTaskSpec) Get() *CronTaskSpec {
	return v.value
}

func (v *NullableCronTaskSpec) Set(val *CronTaskSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCronTaskSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCronTaskSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCronTaskSpec(val *CronTaskSpec) *NullableCronTaskSpec {
	return &NullableCronTaskSpec{value: val, isSet: true}
}

func (v NullableCronTaskSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCronTaskSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const CronTaskType = "cron-task"

type CronTaskComponent struct {
	Base       apis.ComponentBase
	Properties CronTaskSpec
}

func CronTask(name string) *CronTaskComponent {
	c := &CronTaskComponent{Base: apis.ComponentBase{
		Name: name,
	}}
	return c
}

func (c *CronTaskComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range c.Base.Traits {
		traits = append(traits, trait.Build())
	}
	res := common.ApplicationComponent{
		DependsOn:  c.Base.DependsOn,
		Inputs:     c.Base.Inputs,
		Name:       c.Base.Name,
		Outputs:    c.Base.Outputs,
		Properties: util.Object2RawExtension(c.Properties),
		Traits:     traits,
		Type:       CronTaskType,
	}
	return res
}

func (c *CronTaskComponent) AddTrait(traits ...apis.Trait) *CronTaskComponent {
	c.Base.Traits = append(c.Base.Traits, traits...)
	return c
}

func (c *CronTaskComponent) Name() string {
	return c.Base.Name
}

func (c *CronTaskComponent) Type() string {
	return CronTaskType
}
