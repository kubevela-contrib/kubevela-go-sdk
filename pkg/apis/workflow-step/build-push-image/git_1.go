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

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Git1 type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Git1{}

// Git1 Specify the credentials to access git
type Git1 struct {
	// Specify the secret key
	Key *string `json:"key"`
	// Specify the secret name
	Name *string `json:"name"`
}

// NewGit1With instantiates a new Git1 object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewGit1With(key string, name string) *Git1 {
	this := Git1{}
	this.Key = &key
	this.Name = &name
	return &this
}

// NewGit1WithDefault instantiates a new Git1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGit1WithDefault() *Git1 {
	this := Git1{}
	return &this
}

// NewGit1 is short for NewGit1WithDefault which instantiates a new Git1 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGit1() *Git1 {
	return NewGit1WithDefault()
}

// NewGit1Empty instantiates a new Git1 object with no properties set.
// This constructor will not assign any default values to properties.
func NewGit1Empty() *Git1 {
	this := Git1{}
	return &this
}

// NewGit1s converts a list Git1 pointers to objects.
// This is helpful when the SetGit1 requires a list of objects
func NewGit1List(ps ...*Git1) []Git1 {
	objs := []Git1{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Git1
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Git1) Validate() error {
	if o.Key == nil {
		return errors.New("Key in Git1 must be set")
	}
	if o.Name == nil {
		return errors.New("Name in Git1 must be set")
	}
	// validate all nested properties
	return nil
}

// GetKey returns the Key field value
func (o *Git1) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Git1) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key, true
}

// SetKey sets field value
func (o *Git1) SetKey(v string) *Git1 {
	o.Key = &v
	return o
}

// GetName returns the Name field value
func (o *Git1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Git1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *Git1) SetName(v string) *Git1 {
	o.Name = &v
	return o
}

func (o Git1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Git1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableGit1 struct {
	value *Git1
	isSet bool
}

func (v *NullableGit1) Get() *Git1 {
	return v.value
}

func (v *NullableGit1) Set(val *Git1) {
	v.value = val
	v.isSet = true
}

func (v *NullableGit1) IsSet() bool {
	return v.isSet
}

func (v *NullableGit1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGit1(val *Git1) *NullableGit1 {
	return &NullableGit1{value: val, isSet: true}
}

func (v NullableGit1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGit1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
