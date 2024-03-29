/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resource

import (
	"encoding/json"
	"fmt"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// Cpu2 - Specify the amount of cpu for requests
type Cpu2 struct {
	Float32 *float32
	String  *string
}

// Float32AsCpu2 is is a convenience function that returns float32 wrapped in Cpu2
func Float32AsCpu2(v *float32) Cpu2 {
	return Cpu2{
		Float32: v,
	}
}

// Float32AsCpu2OrDefault returns float32 wrapped in Cpu2 if not nil, or a default value if nil
func Float32AsCpu2OrDefault(v *float32) Cpu2 {
	if v == nil {
		return Float32AsCpu2(utils.PtrFloat32(float32(1)))
	}
	return Float32AsCpu2(v)
}

// StringAsCpu2 is is a convenience function that returns string wrapped in Cpu2
func StringAsCpu2(v *string) Cpu2 {
	return Cpu2{
		String: v,
	}
}

// Validate validates this Cpu2
func (o *Cpu2) Validate() error {
	if o.Float32 != nil {
		return nil
	}

	if o.String != nil {
		return nil
	}

	return fmt.Errorf("No oneOf schemas were matched in Cpu2")
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Cpu2) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Float32
	err = utils.NewStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			match++
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal data into String
	err = utils.NewStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Float32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Cpu2)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Cpu2)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Cpu2) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Cpu2) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Float32 != nil {
		return obj.Float32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCpu2 struct {
	value *Cpu2
	isSet bool
}

func (v *NullableCpu2) Get() *Cpu2 {
	return v.value
}

func (v *NullableCpu2) Set(val *Cpu2) {
	v.value = val
	v.isSet = true
}

func (v *NullableCpu2) IsSet() bool {
	return v.isSet
}

func (v *NullableCpu2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpu2(val *Cpu2) *NullableCpu2 {
	return &NullableCpu2{value: val, isSet: true}
}

func (v NullableCpu2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpu2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
