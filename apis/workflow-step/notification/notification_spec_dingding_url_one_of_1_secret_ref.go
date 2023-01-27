/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"

	"vela-go-sdk/apis/utils"
)

// checks if the NotificationSpecDingdingUrlOneOf1SecretRef type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NotificationSpecDingdingUrlOneOf1SecretRef{}

// NotificationSpecDingdingUrlOneOf1SecretRef struct for NotificationSpecDingdingUrlOneOf1SecretRef
type NotificationSpecDingdingUrlOneOf1SecretRef struct {
	// key is the key in the secret
	key string `json:"key"`
	// name is the name of the secret
	name string `json:"name"`
}

// NotificationSpecDingdingUrlOneOf1SecretRefWith instantiates a new NotificationSpecDingdingUrlOneOf1SecretRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NotificationSpecDingdingUrlOneOf1SecretRefWith(key string, name string) *NotificationSpecDingdingUrlOneOf1SecretRef {
	this := NotificationSpecDingdingUrlOneOf1SecretRef{}
	this.key = key
	this.name = name
	return &this
}

// NewNotificationSpecDingdingUrlOneOf1SecretRef instantiates a new NotificationSpecDingdingUrlOneOf1SecretRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSpecDingdingUrlOneOf1SecretRef() *NotificationSpecDingdingUrlOneOf1SecretRef {
	this := NotificationSpecDingdingUrlOneOf1SecretRef{}
	return &this
}

// GetKey returns the Key field value
func (o *NotificationSpecDingdingUrlOneOf1SecretRef) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *NotificationSpecDingdingUrlOneOf1SecretRef) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.key, true
}

// Key sets field value
func (o *NotificationSpecDingdingUrlOneOf1SecretRef) Key(v string) *NotificationSpecDingdingUrlOneOf1SecretRef {
	o.key = v
	return o
}

// GetName returns the Name field value
func (o *NotificationSpecDingdingUrlOneOf1SecretRef) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotificationSpecDingdingUrlOneOf1SecretRef) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *NotificationSpecDingdingUrlOneOf1SecretRef) Name(v string) *NotificationSpecDingdingUrlOneOf1SecretRef {
	o.name = v
	return o
}

func (o NotificationSpecDingdingUrlOneOf1SecretRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationSpecDingdingUrlOneOf1SecretRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.key
	toSerialize["name"] = o.name
	return toSerialize, nil
}

type NullableNotificationSpecDingdingUrlOneOf1SecretRef struct {
	value *NotificationSpecDingdingUrlOneOf1SecretRef
	isSet bool
}

func (v NullableNotificationSpecDingdingUrlOneOf1SecretRef) Get() *NotificationSpecDingdingUrlOneOf1SecretRef {
	return v.value
}

func (v *NullableNotificationSpecDingdingUrlOneOf1SecretRef) Set(val *NotificationSpecDingdingUrlOneOf1SecretRef) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSpecDingdingUrlOneOf1SecretRef) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSpecDingdingUrlOneOf1SecretRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSpecDingdingUrlOneOf1SecretRef(val *NotificationSpecDingdingUrlOneOf1SecretRef) *NullableNotificationSpecDingdingUrlOneOf1SecretRef {
	return &NullableNotificationSpecDingdingUrlOneOf1SecretRef{value: val, isSet: true}
}

func (v NullableNotificationSpecDingdingUrlOneOf1SecretRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSpecDingdingUrlOneOf1SecretRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
