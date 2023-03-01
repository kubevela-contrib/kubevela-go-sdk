/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webhook

import (
	"encoding/json"

	"fmt"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// Url - Specify the webhook url
type Url struct {
	UrlOneOf  *UrlOneOf
	UrlOneOf1 *UrlOneOf1
}

// UrlOneOfAsUrl is a convenience function that returns UrlOneOf wrapped in Url
func UrlOneOfAsUrl(v *UrlOneOf) Url {
	return Url{
		UrlOneOf: v,
	}
}

// UrlOneOf1AsUrl is a convenience function that returns UrlOneOf1 wrapped in Url
func UrlOneOf1AsUrl(v *UrlOneOf1) Url {
	return Url{
		UrlOneOf1: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Url) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UrlOneOf
	err = utils.NewStrictDecoder(data).Decode(&dst.UrlOneOf)
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
	err = utils.NewStrictDecoder(data).Decode(&dst.UrlOneOf1)
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

		return fmt.Errorf("data matches more than one schema in oneOf(Url)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Url)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Url) MarshalJSON() ([]byte, error) {
	if src.UrlOneOf != nil {
		return json.Marshal(&src.UrlOneOf)
	}

	if src.UrlOneOf1 != nil {
		return json.Marshal(&src.UrlOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Url) GetActualInstance() interface{} {
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

type NullableUrl struct {
	value *Url
	isSet bool
}

func (v NullableUrl) Get() *Url {
	return v.value
}

func (v *NullableUrl) Set(val *Url) {
	v.value = val
	v.isSet = true
}

func (v NullableUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrl(val *Url) *NullableUrl {
	return &NullableUrl{value: val, isSet: true}
}

func (v NullableUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
