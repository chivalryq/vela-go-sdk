/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_terraform_config

import (
	"encoding/json"

	"fmt"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// Source - specify the source of the terraform configuration
type Source struct {
	SourceOneOf  *SourceOneOf
	SourceOneOf1 *SourceOneOf1
}

// SourceOneOfAsSource is a convenience function that returns SourceOneOf wrapped in Source
func SourceOneOfAsSource(v *SourceOneOf) Source {
	return Source{
		SourceOneOf: v,
	}
}

// SourceOneOf1AsSource is a convenience function that returns SourceOneOf1 wrapped in Source
func SourceOneOf1AsSource(v *SourceOneOf1) Source {
	return Source{
		SourceOneOf1: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Source) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SourceOneOf
	err = utils.NewStrictDecoder(data).Decode(&dst.SourceOneOf)
	if err == nil {
		jsonSourceOneOf, _ := json.Marshal(dst.SourceOneOf)
		if string(jsonSourceOneOf) == "{}" { // empty struct
			dst.SourceOneOf = nil
		} else {
			match++
		}
	} else {
		dst.SourceOneOf = nil
	}

	// try to unmarshal data into SourceOneOf1
	err = utils.NewStrictDecoder(data).Decode(&dst.SourceOneOf1)
	if err == nil {
		jsonSourceOneOf1, _ := json.Marshal(dst.SourceOneOf1)
		if string(jsonSourceOneOf1) == "{}" { // empty struct
			dst.SourceOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.SourceOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SourceOneOf = nil
		dst.SourceOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Source)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Source)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Source) MarshalJSON() ([]byte, error) {
	if src.SourceOneOf != nil {
		return json.Marshal(&src.SourceOneOf)
	}

	if src.SourceOneOf1 != nil {
		return json.Marshal(&src.SourceOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Source) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SourceOneOf != nil {
		return obj.SourceOneOf
	}

	if obj.SourceOneOf1 != nil {
		return obj.SourceOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableSource struct {
	value *Source
	isSet bool
}

func (v NullableSource) Get() *Source {
	return v.value
}

func (v *NullableSource) Set(val *Source) {
	v.value = val
	v.isSet = true
}

func (v NullableSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSource(val *Source) *NullableSource {
	return &NullableSource{value: val, isSet: true}
}

func (v NullableSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
