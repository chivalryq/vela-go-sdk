/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
"encoding/json"
. "vela-go-sdk/api"
"vela-go-sdk/apis/utils"

"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
    "fmt"
)


// Url1 - Specify the the lark url, you can either sepcify it in value or use secretRef
type Url1 struct {
	UrlOneOf *UrlOneOf
	UrlOneOf1 *UrlOneOf1
}

// UrlOneOfAsUrl1 is a convenience function that returns UrlOneOf wrapped in Url1
func UrlOneOfAsUrl1(v *UrlOneOf) Url1 {
	return Url1{
		UrlOneOf: v,
	}
}

// UrlOneOf1AsUrl1 is a convenience function that returns UrlOneOf1 wrapped in Url1
func UrlOneOf1AsUrl1(v *UrlOneOf1) Url1 {
	return Url1{
		UrlOneOf1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Url1) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UrlOneOf
	err = newStrictDecoder(data).Decode(&dst.UrlOneOf)
	if err == nil {
		jsonUrlOneOf, _ := json.Marshal(dst.UrlOneOf)
		if string(jsonUrlOneOf) == "{}" { // empty struct
			dst.UrlOneOf = nil
		} else {
			match++
		}
	} else {
		dst.UrlOneOf = nil
	}

	// try to unmarshal data into UrlOneOf1
	err = newStrictDecoder(data).Decode(&dst.UrlOneOf1)
	if err == nil {
		jsonUrlOneOf1, _ := json.Marshal(dst.UrlOneOf1)
		if string(jsonUrlOneOf1) == "{}" { // empty struct
			dst.UrlOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.UrlOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.UrlOneOf = nil
		dst.UrlOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Url1)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Url1)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Url1) MarshalJSON() ([]byte, error) {
	if src.UrlOneOf != nil {
		return json.Marshal(&src.UrlOneOf)
	}

	if src.UrlOneOf1 != nil {
		return json.Marshal(&src.UrlOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Url1) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.UrlOneOf != nil {
		return obj.UrlOneOf
	}

	if obj.UrlOneOf1 != nil {
		return obj.UrlOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableUrl1 struct {
	value *Url1
	isSet bool
}

func (v NullableUrl1) Get() *Url1 {
	return v.value
}

func (v *NullableUrl1) Set(val *Url1) {
	v.value = val
	v.isSet = true
}

func (v NullableUrl1) IsSet() bool {
	return v.isSet
}

func (v *NullableUrl1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrl1(val *Url1) *NullableUrl1 {
	return &NullableUrl1{value: val, isSet: true}
}

func (v NullableUrl1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrl1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


