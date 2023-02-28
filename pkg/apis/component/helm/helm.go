/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package helm

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the HelmSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HelmSpec{}

// HelmSpec struct for HelmSpec
type HelmSpec struct {
	// 1.The relative path to helm chart for git/oss source. 2. chart name for helm resource 3. relative path for chart package(e.g. ./charts/podinfo-1.2.3.tgz)
	Chart          *string         `json:"chart,omitempty"`
	Git            *Git            `json:"git,omitempty"`
	Helmrepository *Helmrepository `json:"helmrepository,omitempty"`
	// The timeout for operation `helm install`, optional
	InstallTimeout *string `json:"installTimeout,omitempty"`
	// The  Interval at which to reconcile the Helm release, default to 30s
	Interval *string `json:"interval,omitempty"`
	Oss      *Oss    `json:"oss,omitempty"`
	// The interval at which to check for repository/bucket and release updates, default to 5m
	PullInterval *string `json:"pullInterval,omitempty"`
	// The release name
	ReleaseName *string `json:"releaseName,omitempty"`
	RepoType    *string `json:"repoType,omitempty"`
	// Retry times when install/upgrade fail.
	Retries *int32 `json:"retries,omitempty"`
	// The name of the secret containing authentication credentials
	SecretRef *string `json:"secretRef,omitempty"`
	// The name of the source already existed
	SourceName *string `json:"sourceName,omitempty"`
	// The namespace for helm chart, optional
	TargetNamespace *string `json:"targetNamespace,omitempty"`
	// The timeout for operations like download index/clone repository, optional
	Timeout *string `json:"timeout,omitempty"`
	// The Git or Helm repository URL, OSS endpoint, accept HTTP/S or SSH address as git url,
	Url    *string                `json:"url,omitempty"`
	Values map[string]interface{} `json:"values,omitempty"`
	// Alternative list of values files to use as the chart values (values.yaml is not included by default), expected to be a relative path in the SourceRef.Values files are merged in the order of this list with the last file overriding the first.
	ValuesFiles []string `json:"valuesFiles,omitempty"`
	// Chart version
	Version *string `json:"version,omitempty"`
}

// NewHelmSpecWith instantiates a new HelmSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmSpecWith() *HelmSpec {
	this := HelmSpec{}
	var installTimeout string = "10m"
	this.InstallTimeout = &installTimeout
	var interval string = "30s"
	this.Interval = &interval
	var pullInterval string = "5m"
	this.PullInterval = &pullInterval
	var repoType string = "helm"
	this.RepoType = &repoType
	var retries int32 = 3
	this.Retries = &retries
	var version string = "*"
	this.Version = &version
	return &this
}

// NewHelmSpec instantiates a new HelmSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmSpec() *HelmSpec {
	this := HelmSpec{}
	var installTimeout string = "10m"
	this.InstallTimeout = &installTimeout
	var interval string = "30s"
	this.Interval = &interval
	var pullInterval string = "5m"
	this.PullInterval = &pullInterval
	var repoType string = "helm"
	this.RepoType = &repoType
	var retries int32 = 3
	this.Retries = &retries
	var version string = "*"
	this.Version = &version
	return &this
}

// NewHelmSpecs converts a list HelmSpec pointers to objects.
// This is helpful when the SetHelmSpec requires a list of objects
func NewHelmSpecs(ps ...*HelmSpec) []HelmSpec {
	objs := []HelmSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetChart returns the Chart field value if set, zero value otherwise.
func (o *HelmComponent) GetChart() string {
	if o == nil || utils.IsNil(o.Properties.Chart) {
		var ret string
		return ret
	}
	return *o.Properties.Chart
}

// GetChartOk returns a tuple with the Chart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetChartOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Chart) {
		return nil, false
	}
	return o.Properties.Chart, true
}

// HasChart returns a boolean if a field has been set.
func (o *HelmComponent) HasChart() bool {
	if o != nil && !utils.IsNil(o.Properties.Chart) {
		return true
	}

	return false
}

// SetChart gets a reference to the given string and assigns it to the chart field.
// Chart:  1.The relative path to helm chart for git/oss source. 2. chart name for helm resource 3. relative path for chart package(e.g. ./charts/podinfo-1.2.3.tgz)
func (o *HelmComponent) SetChart(v string) *HelmComponent {
	o.Properties.Chart = &v
	return o
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *HelmComponent) GetGit() Git {
	if o == nil || utils.IsNil(o.Properties.Git) {
		var ret Git
		return ret
	}
	return *o.Properties.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetGitOk() (*Git, bool) {
	if o == nil || utils.IsNil(o.Properties.Git) {
		return nil, false
	}
	return o.Properties.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *HelmComponent) HasGit() bool {
	if o != nil && !utils.IsNil(o.Properties.Git) {
		return true
	}

	return false
}

// SetGit gets a reference to the given Git and assigns it to the git field.
// Git:
func (o *HelmComponent) SetGit(v Git) *HelmComponent {
	o.Properties.Git = &v
	return o
}

// GetHelmrepository returns the Helmrepository field value if set, zero value otherwise.
func (o *HelmComponent) GetHelmrepository() Helmrepository {
	if o == nil || utils.IsNil(o.Properties.Helmrepository) {
		var ret Helmrepository
		return ret
	}
	return *o.Properties.Helmrepository
}

// GetHelmrepositoryOk returns a tuple with the Helmrepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetHelmrepositoryOk() (*Helmrepository, bool) {
	if o == nil || utils.IsNil(o.Properties.Helmrepository) {
		return nil, false
	}
	return o.Properties.Helmrepository, true
}

// HasHelmrepository returns a boolean if a field has been set.
func (o *HelmComponent) HasHelmrepository() bool {
	if o != nil && !utils.IsNil(o.Properties.Helmrepository) {
		return true
	}

	return false
}

// SetHelmrepository gets a reference to the given Helmrepository and assigns it to the helmrepository field.
// Helmrepository:
func (o *HelmComponent) SetHelmrepository(v Helmrepository) *HelmComponent {
	o.Properties.Helmrepository = &v
	return o
}

// GetInstallTimeout returns the InstallTimeout field value if set, zero value otherwise.
func (o *HelmComponent) GetInstallTimeout() string {
	if o == nil || utils.IsNil(o.Properties.InstallTimeout) {
		var ret string
		return ret
	}
	return *o.Properties.InstallTimeout
}

// GetInstallTimeoutOk returns a tuple with the InstallTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetInstallTimeoutOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.InstallTimeout) {
		return nil, false
	}
	return o.Properties.InstallTimeout, true
}

// HasInstallTimeout returns a boolean if a field has been set.
func (o *HelmComponent) HasInstallTimeout() bool {
	if o != nil && !utils.IsNil(o.Properties.InstallTimeout) {
		return true
	}

	return false
}

// SetInstallTimeout gets a reference to the given string and assigns it to the installTimeout field.
// InstallTimeout:  The timeout for operation `helm install`, optional
func (o *HelmComponent) SetInstallTimeout(v string) *HelmComponent {
	o.Properties.InstallTimeout = &v
	return o
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *HelmComponent) GetInterval() string {
	if o == nil || utils.IsNil(o.Properties.Interval) {
		var ret string
		return ret
	}
	return *o.Properties.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetIntervalOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Interval) {
		return nil, false
	}
	return o.Properties.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *HelmComponent) HasInterval() bool {
	if o != nil && !utils.IsNil(o.Properties.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the interval field.
// Interval:  The  Interval at which to reconcile the Helm release, default to 30s
func (o *HelmComponent) SetInterval(v string) *HelmComponent {
	o.Properties.Interval = &v
	return o
}

// GetOss returns the Oss field value if set, zero value otherwise.
func (o *HelmComponent) GetOss() Oss {
	if o == nil || utils.IsNil(o.Properties.Oss) {
		var ret Oss
		return ret
	}
	return *o.Properties.Oss
}

// GetOssOk returns a tuple with the Oss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetOssOk() (*Oss, bool) {
	if o == nil || utils.IsNil(o.Properties.Oss) {
		return nil, false
	}
	return o.Properties.Oss, true
}

// HasOss returns a boolean if a field has been set.
func (o *HelmComponent) HasOss() bool {
	if o != nil && !utils.IsNil(o.Properties.Oss) {
		return true
	}

	return false
}

// SetOss gets a reference to the given Oss and assigns it to the oss field.
// Oss:
func (o *HelmComponent) SetOss(v Oss) *HelmComponent {
	o.Properties.Oss = &v
	return o
}

// GetPullInterval returns the PullInterval field value if set, zero value otherwise.
func (o *HelmComponent) GetPullInterval() string {
	if o == nil || utils.IsNil(o.Properties.PullInterval) {
		var ret string
		return ret
	}
	return *o.Properties.PullInterval
}

// GetPullIntervalOk returns a tuple with the PullInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetPullIntervalOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.PullInterval) {
		return nil, false
	}
	return o.Properties.PullInterval, true
}

// HasPullInterval returns a boolean if a field has been set.
func (o *HelmComponent) HasPullInterval() bool {
	if o != nil && !utils.IsNil(o.Properties.PullInterval) {
		return true
	}

	return false
}

// SetPullInterval gets a reference to the given string and assigns it to the pullInterval field.
// PullInterval:  The interval at which to check for repository/bucket and release updates, default to 5m
func (o *HelmComponent) SetPullInterval(v string) *HelmComponent {
	o.Properties.PullInterval = &v
	return o
}

// GetReleaseName returns the ReleaseName field value if set, zero value otherwise.
func (o *HelmComponent) GetReleaseName() string {
	if o == nil || utils.IsNil(o.Properties.ReleaseName) {
		var ret string
		return ret
	}
	return *o.Properties.ReleaseName
}

// GetReleaseNameOk returns a tuple with the ReleaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetReleaseNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.ReleaseName) {
		return nil, false
	}
	return o.Properties.ReleaseName, true
}

// HasReleaseName returns a boolean if a field has been set.
func (o *HelmComponent) HasReleaseName() bool {
	if o != nil && !utils.IsNil(o.Properties.ReleaseName) {
		return true
	}

	return false
}

// SetReleaseName gets a reference to the given string and assigns it to the releaseName field.
// ReleaseName:  The release name
func (o *HelmComponent) SetReleaseName(v string) *HelmComponent {
	o.Properties.ReleaseName = &v
	return o
}

// GetRepoType returns the RepoType field value if set, zero value otherwise.
func (o *HelmComponent) GetRepoType() string {
	if o == nil || utils.IsNil(o.Properties.RepoType) {
		var ret string
		return ret
	}
	return *o.Properties.RepoType
}

// GetRepoTypeOk returns a tuple with the RepoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetRepoTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.RepoType) {
		return nil, false
	}
	return o.Properties.RepoType, true
}

// HasRepoType returns a boolean if a field has been set.
func (o *HelmComponent) HasRepoType() bool {
	if o != nil && !utils.IsNil(o.Properties.RepoType) {
		return true
	}

	return false
}

// SetRepoType gets a reference to the given string and assigns it to the repoType field.
// RepoType:
func (o *HelmComponent) SetRepoType(v string) *HelmComponent {
	o.Properties.RepoType = &v
	return o
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *HelmComponent) GetRetries() int32 {
	if o == nil || utils.IsNil(o.Properties.Retries) {
		var ret int32
		return ret
	}
	return *o.Properties.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetRetriesOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.Retries) {
		return nil, false
	}
	return o.Properties.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *HelmComponent) HasRetries() bool {
	if o != nil && !utils.IsNil(o.Properties.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the retries field.
// Retries:  Retry times when install/upgrade fail.
func (o *HelmComponent) SetRetries(v int32) *HelmComponent {
	o.Properties.Retries = &v
	return o
}

// GetSecretRef returns the SecretRef field value if set, zero value otherwise.
func (o *HelmComponent) GetSecretRef() string {
	if o == nil || utils.IsNil(o.Properties.SecretRef) {
		var ret string
		return ret
	}
	return *o.Properties.SecretRef
}

// GetSecretRefOk returns a tuple with the SecretRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetSecretRefOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.SecretRef) {
		return nil, false
	}
	return o.Properties.SecretRef, true
}

// HasSecretRef returns a boolean if a field has been set.
func (o *HelmComponent) HasSecretRef() bool {
	if o != nil && !utils.IsNil(o.Properties.SecretRef) {
		return true
	}

	return false
}

// SetSecretRef gets a reference to the given string and assigns it to the secretRef field.
// SecretRef:  The name of the secret containing authentication credentials
func (o *HelmComponent) SetSecretRef(v string) *HelmComponent {
	o.Properties.SecretRef = &v
	return o
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *HelmComponent) GetSourceName() string {
	if o == nil || utils.IsNil(o.Properties.SourceName) {
		var ret string
		return ret
	}
	return *o.Properties.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetSourceNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.SourceName) {
		return nil, false
	}
	return o.Properties.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *HelmComponent) HasSourceName() bool {
	if o != nil && !utils.IsNil(o.Properties.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the sourceName field.
// SourceName:  The name of the source already existed
func (o *HelmComponent) SetSourceName(v string) *HelmComponent {
	o.Properties.SourceName = &v
	return o
}

// GetTargetNamespace returns the TargetNamespace field value if set, zero value otherwise.
func (o *HelmComponent) GetTargetNamespace() string {
	if o == nil || utils.IsNil(o.Properties.TargetNamespace) {
		var ret string
		return ret
	}
	return *o.Properties.TargetNamespace
}

// GetTargetNamespaceOk returns a tuple with the TargetNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetTargetNamespaceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.TargetNamespace) {
		return nil, false
	}
	return o.Properties.TargetNamespace, true
}

// HasTargetNamespace returns a boolean if a field has been set.
func (o *HelmComponent) HasTargetNamespace() bool {
	if o != nil && !utils.IsNil(o.Properties.TargetNamespace) {
		return true
	}

	return false
}

// SetTargetNamespace gets a reference to the given string and assigns it to the targetNamespace field.
// TargetNamespace:  The namespace for helm chart, optional
func (o *HelmComponent) SetTargetNamespace(v string) *HelmComponent {
	o.Properties.TargetNamespace = &v
	return o
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *HelmComponent) GetTimeout() string {
	if o == nil || utils.IsNil(o.Properties.Timeout) {
		var ret string
		return ret
	}
	return *o.Properties.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetTimeoutOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Timeout) {
		return nil, false
	}
	return o.Properties.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *HelmComponent) HasTimeout() bool {
	if o != nil && !utils.IsNil(o.Properties.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given string and assigns it to the timeout field.
// Timeout:  The timeout for operations like download index/clone repository, optional
func (o *HelmComponent) SetTimeout(v string) *HelmComponent {
	o.Properties.Timeout = &v
	return o
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *HelmComponent) GetUrl() string {
	if o == nil || utils.IsNil(o.Properties.Url) {
		var ret string
		return ret
	}
	return *o.Properties.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Url) {
		return nil, false
	}
	return o.Properties.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *HelmComponent) HasUrl() bool {
	if o != nil && !utils.IsNil(o.Properties.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the url field.
// Url:  The Git or Helm repository URL, OSS endpoint, accept HTTP/S or SSH address as git url,
func (o *HelmComponent) SetUrl(v string) *HelmComponent {
	o.Properties.Url = &v
	return o
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *HelmComponent) GetValues() map[string]interface{} {
	if o == nil || utils.IsNil(o.Properties.Values) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetValuesOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Properties.Values) {
		return map[string]interface{}{}, false
	}
	return o.Properties.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *HelmComponent) HasValues() bool {
	if o != nil && !utils.IsNil(o.Properties.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given map[string]interface{} and assigns it to the values field.
// Values:
func (o *HelmComponent) SetValues(v map[string]interface{}) *HelmComponent {
	o.Properties.Values = v
	return o
}

// GetValuesFiles returns the ValuesFiles field value if set, zero value otherwise.
func (o *HelmComponent) GetValuesFiles() []string {
	if o == nil || utils.IsNil(o.Properties.ValuesFiles) {
		var ret []string
		return ret
	}
	return o.Properties.ValuesFiles
}

// GetValuesFilesOk returns a tuple with the ValuesFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetValuesFilesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.ValuesFiles) {
		return nil, false
	}
	return o.Properties.ValuesFiles, true
}

// HasValuesFiles returns a boolean if a field has been set.
func (o *HelmComponent) HasValuesFiles() bool {
	if o != nil && !utils.IsNil(o.Properties.ValuesFiles) {
		return true
	}

	return false
}

// SetValuesFiles gets a reference to the given []string and assigns it to the valuesFiles field.
// ValuesFiles:  Alternative list of values files to use as the chart values (values.yaml is not included by default), expected to be a relative path in the SourceRef.Values files are merged in the order of this list with the last file overriding the first.
func (o *HelmComponent) SetValuesFiles(v []string) *HelmComponent {
	o.Properties.ValuesFiles = v
	return o
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HelmComponent) GetVersion() string {
	if o == nil || utils.IsNil(o.Properties.Version) {
		var ret string
		return ret
	}
	return *o.Properties.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmComponent) GetVersionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Version) {
		return nil, false
	}
	return o.Properties.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HelmComponent) HasVersion() bool {
	if o != nil && !utils.IsNil(o.Properties.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the version field.
// Version:  Chart version
func (o *HelmComponent) SetVersion(v string) *HelmComponent {
	o.Properties.Version = &v
	return o
}

func (o HelmSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Chart) {
		toSerialize["chart"] = o.Chart
	}
	if !utils.IsNil(o.Git) {
		toSerialize["git"] = o.Git
	}
	if !utils.IsNil(o.Helmrepository) {
		toSerialize["helmrepository"] = o.Helmrepository
	}
	if !utils.IsNil(o.InstallTimeout) {
		toSerialize["installTimeout"] = o.InstallTimeout
	}
	if !utils.IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !utils.IsNil(o.Oss) {
		toSerialize["oss"] = o.Oss
	}
	if !utils.IsNil(o.PullInterval) {
		toSerialize["pullInterval"] = o.PullInterval
	}
	if !utils.IsNil(o.ReleaseName) {
		toSerialize["releaseName"] = o.ReleaseName
	}
	if !utils.IsNil(o.RepoType) {
		toSerialize["repoType"] = o.RepoType
	}
	if !utils.IsNil(o.Retries) {
		toSerialize["retries"] = o.Retries
	}
	if !utils.IsNil(o.SecretRef) {
		toSerialize["secretRef"] = o.SecretRef
	}
	if !utils.IsNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}
	if !utils.IsNil(o.TargetNamespace) {
		toSerialize["targetNamespace"] = o.TargetNamespace
	}
	if !utils.IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !utils.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !utils.IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !utils.IsNil(o.ValuesFiles) {
		toSerialize["valuesFiles"] = o.ValuesFiles
	}
	if !utils.IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableHelmSpec struct {
	value *HelmSpec
	isSet bool
}

func (v NullableHelmSpec) Get() *HelmSpec {
	return v.value
}

func (v *NullableHelmSpec) Set(val *HelmSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmSpec(val *HelmSpec) *NullableHelmSpec {
	return &NullableHelmSpec{value: val, isSet: true}
}

func (v NullableHelmSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const HelmType = "helm"

func init() {
	sdkcommon.RegisterComponent(HelmType, FromComponent)
}

type HelmComponent struct {
	Base       apis.ComponentBase
	Properties HelmSpec
}

func Helm(name string) *HelmComponent {
	h := &HelmComponent{Base: apis.ComponentBase{
		Name: name,
		Type: HelmType,
	}}
	return h
}

func (h *HelmComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range h.Base.Traits {
		traits = append(traits, trait.Build())
	}
	res := common.ApplicationComponent{
		DependsOn:  h.Base.DependsOn,
		Inputs:     h.Base.Inputs,
		Name:       h.Base.Name,
		Outputs:    h.Base.Outputs,
		Properties: util.Object2RawExtension(h.Properties),
		Traits:     traits,
		Type:       HelmType,
	}
	return res
}

func (h *HelmComponent) FromComponent(from common.ApplicationComponent) (*HelmComponent, error) {
	for _, trait := range from.Traits {
		_t, err := sdkcommon.FromTrait(&trait)
		if err != nil {
			return nil, err
		}
		h.Base.Traits = append(h.Base.Traits, _t)
	}
	var properties HelmSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	h.Base.Name = from.Name
	h.Base.DependsOn = from.DependsOn
	h.Base.Inputs = from.Inputs
	h.Base.Outputs = from.Outputs
	h.Base.Type = HelmType
	h.Properties = properties
	return h, nil
}

func FromComponent(from common.ApplicationComponent) (apis.Component, error) {
	h := &HelmComponent{}
	return h.FromComponent(from)
}

func (h *HelmComponent) AddTrait(traits ...apis.Trait) *HelmComponent {
	h.Base.Traits = append(h.Base.Traits, traits...)
	return h
}

func (h *HelmComponent) GetTrait(typ string) apis.Trait {
	for _, _t := range h.Base.Traits {
		if _t.DefType() == typ {
			return _t
		}
	}
	return nil
}

func (h *HelmComponent) ComponentName() string {
	return h.Base.Name
}

func (h *HelmComponent) DefType() string {
	return HelmType
}

func (h *HelmComponent) DependsOn(dependsOn []string) *HelmComponent {
	h.Base.DependsOn = dependsOn
	return h
}

func (h *HelmComponent) Inputs(input common.StepInputs) *HelmComponent {
	h.Base.Inputs = input
	return h
}

func (h *HelmComponent) Outputs(output common.StepOutputs) *HelmComponent {
	h.Base.Outputs = output
	return h
}

func (h *HelmComponent) AddDependsOn(dependsOn string) *HelmComponent {
	h.Base.DependsOn = append(h.Base.DependsOn, dependsOn)
	return h
}
