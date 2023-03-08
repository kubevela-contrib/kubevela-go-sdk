/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package build_push_image

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// Context - Specify the context to build image, you can use context with git and branch or directly specify the context, please refer to https://github.com/GoogleContainerTools/kaniko#kaniko-build-contexts
type Context struct {
	Git    *Git
	String *string
}

// GitAsContext is a convenience function that returns Git wrapped in Context
func GitAsContext(v *Git) Context {
	return Context{
		Git: v,
	}
}

// stringAsContext is a convenience function that returns string wrapped in Context
func StringAsContext(v *string) Context {
	return Context{
		String: v,
	}
}

// Validate validates this Context
func (o *Context) Validate() error {
	if o.Git != nil {
		return nil
	}

	if o.String != nil {
		return nil
	}

	return fmt.Errorf("No oneOf schemas were matched in Context")
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Context) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Git
	err = utils.NewStrictDecoder(data).Decode(&dst.Git)
	if err == nil {
		jsonGit, _ := json.Marshal(dst.Git)
		if string(jsonGit) == "{}" { // empty struct
			dst.Git = nil
		} else {
			match++
		}
	} else {
		dst.Git = nil
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
		dst.Git = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Context)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Context)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Context) MarshalJSON() ([]byte, error) {
	if src.Git != nil {
		return json.Marshal(&src.Git)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Context) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Git != nil {
		return obj.Git
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableContext struct {
	value *Context
	isSet bool
}

func (v *NullableContext) Get() *Context {
	return v.value
}

func (v *NullableContext) Set(val *Context) {
	v.value = val
	v.isSet = true
}

func (v *NullableContext) IsSet() bool {
	return v.isSet
}

func (v *NullableContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContext(val *Context) *NullableContext {
	return &NullableContext{value: val, isSet: true}
}

func (v NullableContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
