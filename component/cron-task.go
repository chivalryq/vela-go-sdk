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

package component

import (
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	. "vela-go-sdk/api"
)

const CronTaskType = "cron-task"

// CronTaskComponent is the root struct of cron-task
type CronTaskComponent struct {
	Base  ComponentBase
	Props CronTaskSpec
}

// HostAliases -
type HostAliases struct {
	IP        string   `json:"ip"`
	Hostnames []string `json:"hostnames"`
}

// CronTaskSpec -
type CronTaskSpec struct {
	Labels                     map[string]string `json:"labels,omitempty"`
	Annotations                map[string]string `json:"annotations,omitempty"`
	Schedule                   string            `json:"schedule"`
	StartingDeadlineSeconds    int               `json:"startingDeadlineSeconds,omitempty"`
	Suspend                    bool              `json:"suspend"`
	ConcurrencyPolicy          string            `json:"concurrencyPolicy"`
	SuccessfulJobsHistoryLimit int               `json:"successfulJobsHistoryLimit"`
	FailedJobsHistoryLimit     int               `json:"failedJobsHistoryLimit"`
	Count                      int               `json:"count"`
	Image                      string            `json:"image"`
	ImagePullPolicy            string            `json:"imagePullPolicy,omitempty"`
	ImagePullSecrets           []string          `json:"imagePullSecrets,omitempty"`
	Restart                    string            `json:"restart"`
	Cmd                        []string          `json:"cmd,omitempty"`
	Env                        Env               `json:"env,omitempty"`
	CPU                        string            `json:"cpu,omitempty"`
	Memory                     string            `json:"memory,omitempty"`
	Volumes                    Volumes           `json:"volumes,omitempty"`
	HostAliases                HostAliases       `json:"hostAliases,omitempty"`
	TtlSecondsAfterFinished    int               `json:"ttlSecondsAfterFinished,omitempty"`
	ActiveDeadlineSeconds      int               `json:"activeDeadlineSeconds,omitempty"`
	BackoffLimit               int               `json:"backoffLimit"`
	LivenessProbe              HealthProbe       `json:"livenessProbe,omitempty"`
	ReadinessProbe             HealthProbe       `json:"readinessProbe,omitempty"`
}

// HTTPHeaders -
type HTTPHeaders struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// HTTPGet Instructions for assessing container health by executing an HTTP GET request. Either this attribute or the exec attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with both the exec attribute and the tcpSocket attribute.
type HTTPGet struct {
	Path        string      `json:"path"`
	Port        int         `json:"port"`
	HTTPHeaders HTTPHeaders `json:"httpHeaders,omitempty"`
}

// SecretKeyRef Selects a key of a secret in the pod's namespace
type SecretKeyRef struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// ValueFrom Specifies a source the value of this var should come from
type ValueFrom struct {
	SecretKeyRef    SecretKeyRef    `json:"secretKeyRef,omitempty"`
	ConfigMapKeyRef ConfigMapKeyRef `json:"configMapKeyRef,omitempty"`
}

// Env -
type Env struct {
	Name      string    `json:"name"`
	Value     string    `json:"value,omitempty"`
	ValueFrom ValueFrom `json:"valueFrom,omitempty"`
}

// Volumes -
type Volumes struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
	Medium    string `json:"medium"`
	Type      string `json:"type"`
}

// TcpSocket Instructions for assessing container health by probing a TCP socket. Either this attribute or the exec attribute or the httpGet attribute MUST be specified. This attribute is mutually exclusive with both the exec attribute and the httpGet attribute.
type TcpSocket struct {
	Port int `json:"port"`
}

// HealthProbe -
type HealthProbe struct {
	Exec                Exec      `json:"exec,omitempty"`
	HTTPGet             HTTPGet   `json:"httpGet,omitempty"`
	TcpSocket           TcpSocket `json:"tcpSocket,omitempty"`
	InitialDelaySeconds int       `json:"initialDelaySeconds"`
	PeriodSeconds       int       `json:"periodSeconds"`
	TimeoutSeconds      int       `json:"timeoutSeconds"`
	SuccessThreshold    int       `json:"successThreshold"`
	FailureThreshold    int       `json:"failureThreshold"`
}

// ConfigMapKeyRef Selects a key of a config map in the pod's namespace
type ConfigMapKeyRef struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// Exec Instructions for assessing container health by executing a command. Either this attribute or the httpGet attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with both the httpGet attribute and the tcpSocket attribute.
type Exec struct {
	Command []string `json:"command"`
}

func CronTask(name string) *CronTaskComponent {
	comp := CronTaskComponent{
		Base: ComponentBase{
			Name: name,
		},
	}
	return &comp
}

func (c *CronTaskComponent) Trait(traits ...Trait) *CronTaskComponent {
	c.Base.Traits = append(c.Base.Traits, traits...)
	return c
}

func (c *CronTaskComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range c.Base.Traits {
		traits = append(traits, trait.Build())
	}
	comp := common.ApplicationComponent{
		Name:       c.Base.Name,
		Type:       CronTaskType,
		Properties: util.Object2RawExtension(c.Props),
		DependsOn:  c.Base.DependsOn,
		Inputs:     c.Base.Inputs,
		Outputs:    c.Base.Outputs,
		Traits:     traits,
	}
	return comp
}

// Labels Specify the labels in the workload
func (c *CronTaskComponent) Labels(value map[string]string) *CronTaskComponent {
	c.Props.Labels = value
	return c
}

// Annotations Specify the annotations in the workload
func (c *CronTaskComponent) Annotations(value map[string]string) *CronTaskComponent {
	c.Props.Annotations = value
	return c
}

// Schedule Specify the schedule in Cron format, see https://en.wikipedia.org/wiki/Cron
func (c *CronTaskComponent) Schedule(value string) *CronTaskComponent {
	c.Props.Schedule = value
	return c
}

// StartingDeadlineSeconds Specify deadline in seconds for starting the job if it misses scheduled
func (c *CronTaskComponent) StartingDeadlineSeconds(value int) *CronTaskComponent {
	c.Props.StartingDeadlineSeconds = value
	return c
}

// Suspend suspend subsequent executions
func (c *CronTaskComponent) Suspend(value bool) *CronTaskComponent {
	c.Props.Suspend = value
	return c
}

// ConcurrencyPolicy Specifies how to treat concurrent executions of a Job
func (c *CronTaskComponent) ConcurrencyPolicy(value string) *CronTaskComponent {
	c.Props.ConcurrencyPolicy = value
	return c
}

// SuccessfulJobsHistoryLimit The number of successful finished jobs to retain
func (c *CronTaskComponent) SuccessfulJobsHistoryLimit(value int) *CronTaskComponent {
	c.Props.SuccessfulJobsHistoryLimit = value
	return c
}

// FailedJobsHistoryLimit The number of failed finished jobs to retain
func (c *CronTaskComponent) FailedJobsHistoryLimit(value int) *CronTaskComponent {
	c.Props.FailedJobsHistoryLimit = value
	return c
}

// Count Specify number of tasks to run in parallel
func (c *CronTaskComponent) Count(value int) *CronTaskComponent {
	c.Props.Count = value
	return c
}

// Image Which image would you like to use for your service
func (c *CronTaskComponent) Image(value string) *CronTaskComponent {
	c.Props.Image = value
	return c
}

// ImagePullPolicy Specify image pull policy for your service
func (c *CronTaskComponent) ImagePullPolicy(value string) *CronTaskComponent {
	c.Props.ImagePullPolicy = value
	return c
}

// ImagePullSecrets Specify image pull secrets for your service
func (c *CronTaskComponent) ImagePullSecrets(value []string) *CronTaskComponent {
	c.Props.ImagePullSecrets = value
	return c
}

// Restart Define the job restart policy, the value can only be Never or OnFailure. By default, it's Never.
func (c *CronTaskComponent) Restart(value string) *CronTaskComponent {
	c.Props.Restart = value
	return c
}

// Cmd Commands to run in the container
func (c *CronTaskComponent) Cmd(value []string) *CronTaskComponent {
	c.Props.Cmd = value
	return c
}

// Env Define arguments by using environment variables
func (c *CronTaskComponent) Env(value Env) *CronTaskComponent {
	c.Props.Env = value
	return c
}

// CPU Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
func (c *CronTaskComponent) CPU(value string) *CronTaskComponent {
	c.Props.CPU = value
	return c
}

// Memory Specifies the attributes of the memory resource required for the container.
func (c *CronTaskComponent) Memory(value string) *CronTaskComponent {
	c.Props.Memory = value
	return c
}

// Volumes Declare volumes and volumeMounts
func (c *CronTaskComponent) Volumes(value Volumes) *CronTaskComponent {
	c.Props.Volumes = value
	return c
}

// HostAliases An optional list of hosts and IPs that will be injected into the pod's hosts file
func (c *CronTaskComponent) HostAliases(value HostAliases) *CronTaskComponent {
	c.Props.HostAliases = value
	return c
}

// TtlSecondsAfterFinished Limits the lifetime of a Job that has finished
func (c *CronTaskComponent) TtlSecondsAfterFinished(value int) *CronTaskComponent {
	c.Props.TtlSecondsAfterFinished = value
	return c
}

// ActiveDeadlineSeconds The duration in seconds relative to the startTime that the job may be continuously active before the system tries to terminate it
func (c *CronTaskComponent) ActiveDeadlineSeconds(value int) *CronTaskComponent {
	c.Props.ActiveDeadlineSeconds = value
	return c
}

// BackoffLimit The number of retries before marking this job failed
func (c *CronTaskComponent) BackoffLimit(value int) *CronTaskComponent {
	c.Props.BackoffLimit = value
	return c
}

// LivenessProbe Instructions for assessing whether the container is alive.
func (c *CronTaskComponent) LivenessProbe(value HealthProbe) *CronTaskComponent {
	c.Props.LivenessProbe = value
	return c
}

// ReadinessProbe Instructions for assessing whether the container is in a suitable state to serve traffic.
func (c *CronTaskComponent) ReadinessProbe(value HealthProbe) *CronTaskComponent {
	c.Props.ReadinessProbe = value
	return c
}
