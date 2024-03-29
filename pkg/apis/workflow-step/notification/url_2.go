/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"
	"fmt"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// Url2 - Specify the the slack url, you can either sepcify it in value or use secretRef
type Url2 struct {
	UrlOneOf  *UrlOneOf
	UrlOneOf1 *UrlOneOf1
}

// UrlOneOfAsUrl2 is is a convenience function that returns UrlOneOf wrapped in Url2
func UrlOneOfAsUrl2(v *UrlOneOf) Url2 {
	return Url2{
		UrlOneOf: v,
	}
}

// UrlOneOf1AsUrl2 is is a convenience function that returns UrlOneOf1 wrapped in Url2
func UrlOneOf1AsUrl2(v *UrlOneOf1) Url2 {
	return Url2{
		UrlOneOf1: v,
	}
}

// Validate validates this Url2
func (o *Url2) Validate() error {
	if o.UrlOneOf != nil {
		return nil
	}

	if o.UrlOneOf1 != nil {
		return nil
	}

	return fmt.Errorf("No oneOf schemas were matched in Url2")
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Url2) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(Url2)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Url2)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Url2) MarshalJSON() ([]byte, error) {
	if src.UrlOneOf != nil {
		return json.Marshal(&src.UrlOneOf)
	}

	if src.UrlOneOf1 != nil {
		return json.Marshal(&src.UrlOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Url2) GetActualInstance() interface{} {
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

type NullableUrl2 struct {
	value *Url2
	isSet bool
}

func (v *NullableUrl2) Get() *Url2 {
	return v.value
}

func (v *NullableUrl2) Set(val *Url2) {
	v.value = val
	v.isSet = true
}

func (v *NullableUrl2) IsSet() bool {
	return v.isSet
}

func (v *NullableUrl2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrl2(val *Url2) *NullableUrl2 {
	return &NullableUrl2{value: val, isSet: true}
}

func (v NullableUrl2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrl2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
