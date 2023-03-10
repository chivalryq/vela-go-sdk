/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the NotificationSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NotificationSpec{}

// NotificationSpec struct for NotificationSpec
type NotificationSpec struct {
	Dingding *Dingding `json:"dingding,omitempty"`
	Email    *Email    `json:"email,omitempty"`
	Lark     *Lark     `json:"lark,omitempty"`
	Slack    *Slack    `json:"slack,omitempty"`
}

// NewNotificationSpecWith instantiates a new NotificationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSpecWith() *NotificationSpec {
	this := NotificationSpec{}
	return &this
}

// NewNotificationSpec instantiates a new NotificationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSpec() *NotificationSpec {
	this := NotificationSpec{}
	return &this
}

// NewNotificationSpecs converts a list NotificationSpec pointers to objects.
// This is helpful when the SetNotificationSpec requires a list of objects
func NewNotificationSpecs(ps ...*NotificationSpec) []NotificationSpec {
	objs := []NotificationSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetDingding returns the Dingding field value if set, zero value otherwise.
func (o *NotificationWorkflowStep) GetDingding() Dingding {
	if o == nil || utils.IsNil(o.Properties.Dingding) {
		var ret Dingding
		return ret
	}
	return *o.Properties.Dingding
}

// GetDingdingOk returns a tuple with the Dingding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWorkflowStep) GetDingdingOk() (*Dingding, bool) {
	if o == nil || utils.IsNil(o.Properties.Dingding) {
		return nil, false
	}
	return o.Properties.Dingding, true
}

// HasDingding returns a boolean if a field has been set.
func (o *NotificationWorkflowStep) HasDingding() bool {
	if o != nil && !utils.IsNil(o.Properties.Dingding) {
		return true
	}

	return false
}

// SetDingding gets a reference to the given Dingding and assigns it to the dingding field.
// Dingding:
func (o *NotificationWorkflowStep) SetDingding(v Dingding) *NotificationWorkflowStep {
	o.Properties.Dingding = &v
	return o
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *NotificationWorkflowStep) GetEmail() Email {
	if o == nil || utils.IsNil(o.Properties.Email) {
		var ret Email
		return ret
	}
	return *o.Properties.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWorkflowStep) GetEmailOk() (*Email, bool) {
	if o == nil || utils.IsNil(o.Properties.Email) {
		return nil, false
	}
	return o.Properties.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *NotificationWorkflowStep) HasEmail() bool {
	if o != nil && !utils.IsNil(o.Properties.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given Email and assigns it to the email field.
// Email:
func (o *NotificationWorkflowStep) SetEmail(v Email) *NotificationWorkflowStep {
	o.Properties.Email = &v
	return o
}

// GetLark returns the Lark field value if set, zero value otherwise.
func (o *NotificationWorkflowStep) GetLark() Lark {
	if o == nil || utils.IsNil(o.Properties.Lark) {
		var ret Lark
		return ret
	}
	return *o.Properties.Lark
}

// GetLarkOk returns a tuple with the Lark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWorkflowStep) GetLarkOk() (*Lark, bool) {
	if o == nil || utils.IsNil(o.Properties.Lark) {
		return nil, false
	}
	return o.Properties.Lark, true
}

// HasLark returns a boolean if a field has been set.
func (o *NotificationWorkflowStep) HasLark() bool {
	if o != nil && !utils.IsNil(o.Properties.Lark) {
		return true
	}

	return false
}

// SetLark gets a reference to the given Lark and assigns it to the lark field.
// Lark:
func (o *NotificationWorkflowStep) SetLark(v Lark) *NotificationWorkflowStep {
	o.Properties.Lark = &v
	return o
}

// GetSlack returns the Slack field value if set, zero value otherwise.
func (o *NotificationWorkflowStep) GetSlack() Slack {
	if o == nil || utils.IsNil(o.Properties.Slack) {
		var ret Slack
		return ret
	}
	return *o.Properties.Slack
}

// GetSlackOk returns a tuple with the Slack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWorkflowStep) GetSlackOk() (*Slack, bool) {
	if o == nil || utils.IsNil(o.Properties.Slack) {
		return nil, false
	}
	return o.Properties.Slack, true
}

// HasSlack returns a boolean if a field has been set.
func (o *NotificationWorkflowStep) HasSlack() bool {
	if o != nil && !utils.IsNil(o.Properties.Slack) {
		return true
	}

	return false
}

// SetSlack gets a reference to the given Slack and assigns it to the slack field.
// Slack:
func (o *NotificationWorkflowStep) SetSlack(v Slack) *NotificationWorkflowStep {
	o.Properties.Slack = &v
	return o
}

func (o NotificationSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Dingding) {
		toSerialize["dingding"] = o.Dingding
	}
	if !utils.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !utils.IsNil(o.Lark) {
		toSerialize["lark"] = o.Lark
	}
	if !utils.IsNil(o.Slack) {
		toSerialize["slack"] = o.Slack
	}
	return toSerialize, nil
}

type NullableNotificationSpec struct {
	value *NotificationSpec
	isSet bool
}

func (v NullableNotificationSpec) Get() *NotificationSpec {
	return v.value
}

func (v *NullableNotificationSpec) Set(val *NotificationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSpec(val *NotificationSpec) *NullableNotificationSpec {
	return &NullableNotificationSpec{value: val, isSet: true}
}

func (v NullableNotificationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const NotificationType = "notification"

func init() {
	sdkcommon.RegisterWorkflowStep(NotificationType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(NotificationType, FromWorkflowSubStep)
}

type NotificationWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties NotificationSpec
}

func Notification(name string) *NotificationWorkflowStep {
	n := &NotificationWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: NotificationType,
	}}
	return n
}

func (n *NotificationWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range n.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  n.Base.DependsOn,
		If:         n.Base.If,
		Inputs:     n.Base.Inputs,
		Meta:       n.Base.Meta,
		Name:       n.Base.Name,
		Outputs:    n.Base.Outputs,
		Properties: util.Object2RawExtension(n.Properties),
		SubSteps:   subSteps,
		Timeout:    n.Base.Timeout,
		Type:       NotificationType,
	}
	return res
}

func (n *NotificationWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*NotificationWorkflowStep, error) {
	var properties NotificationSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := n.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	n.Base.Name = from.Name
	n.Base.DependsOn = from.DependsOn
	n.Base.Inputs = from.Inputs
	n.Base.Outputs = from.Outputs
	n.Base.If = from.If
	n.Base.Timeout = from.Timeout
	n.Base.Meta = from.Meta
	n.Base.Type = NotificationType
	n.Properties = properties
	n.Base.SubSteps = subSteps
	return n, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	n := &NotificationWorkflowStep{}
	return n.FromWorkflowStep(from)
}

func (n *NotificationWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*NotificationWorkflowStep, error) {
	var properties NotificationSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	n.Base.Name = from.Name
	n.Base.DependsOn = from.DependsOn
	n.Base.Inputs = from.Inputs
	n.Base.Outputs = from.Outputs
	n.Base.If = from.If
	n.Base.Timeout = from.Timeout
	n.Base.Meta = from.Meta
	n.Base.Type = NotificationType
	n.Properties = properties
	return n, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	n := &NotificationWorkflowStep{}
	return n.FromWorkflowSubStep(from)
}

func (n *NotificationWorkflowStep) WorkflowStepName() string {
	return n.Base.Name
}

func (n *NotificationWorkflowStep) DefType() string {
	return NotificationType
}

func (n *NotificationWorkflowStep) If(_if string) *NotificationWorkflowStep {
	n.Base.If = _if
	return n
}

func (n *NotificationWorkflowStep) Alias(alias string) *NotificationWorkflowStep {
	n.Base.Meta.Alias = alias
	return n
}

func (n *NotificationWorkflowStep) Timeout(timeout string) *NotificationWorkflowStep {
	n.Base.Timeout = timeout
	return n
}

func (n *NotificationWorkflowStep) DependsOn(dependsOn []string) *NotificationWorkflowStep {
	n.Base.DependsOn = dependsOn
	return n
}

func (n *NotificationWorkflowStep) Inputs(input common.StepInputs) *NotificationWorkflowStep {
	n.Base.Inputs = input
	return n
}

func (n *NotificationWorkflowStep) Outputs(output common.StepOutputs) *NotificationWorkflowStep {
	n.Base.Outputs = output
	return n
}
