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

// Password - Specify the password of the email, you can either sepcify it in value or use secretRef
type Password struct {
	PasswordOneOf *PasswordOneOf
	UrlOneOf1     *UrlOneOf1
}

// PasswordOneOfAsPassword is is a convenience function that returns PasswordOneOf wrapped in Password
func PasswordOneOfAsPassword(v *PasswordOneOf) Password {
	return Password{
		PasswordOneOf: v,
	}
}

// UrlOneOf1AsPassword is is a convenience function that returns UrlOneOf1 wrapped in Password
func UrlOneOf1AsPassword(v *UrlOneOf1) Password {
	return Password{
		UrlOneOf1: v,
	}
}

// Validate validates this Password
func (o *Password) Validate() error {
	if o.PasswordOneOf != nil {
		return nil
	}

	if o.UrlOneOf1 != nil {
		return nil
	}

	return fmt.Errorf("No oneOf schemas were matched in Password")
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Password) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PasswordOneOf
	err = utils.NewStrictDecoder(data).Decode(&dst.PasswordOneOf)
	if err == nil {
		jsonPasswordOneOf, _ := json.Marshal(dst.PasswordOneOf)
		if string(jsonPasswordOneOf) == "{}" { // empty struct
			dst.PasswordOneOf = nil
		} else {
			match++
		}
	} else {
		dst.PasswordOneOf = nil
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
		dst.PasswordOneOf = nil
		dst.UrlOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Password)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Password)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Password) MarshalJSON() ([]byte, error) {
	if src.PasswordOneOf != nil {
		return json.Marshal(&src.PasswordOneOf)
	}

	if src.UrlOneOf1 != nil {
		return json.Marshal(&src.UrlOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Password) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.PasswordOneOf != nil {
		return obj.PasswordOneOf
	}

	if obj.UrlOneOf1 != nil {
		return obj.UrlOneOf1
	}

	// all schemas are nil
	return nil
}

type NullablePassword struct {
	value *Password
	isSet bool
}

func (v *NullablePassword) Get() *Password {
	return v.value
}

func (v *NullablePassword) Set(val *Password) {
	v.value = val
	v.isSet = true
}

func (v *NullablePassword) IsSet() bool {
	return v.isSet
}

func (v *NullablePassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePassword(val *Password) *NullablePassword {
	return &NullablePassword{value: val, isSet: true}
}

func (v NullablePassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
