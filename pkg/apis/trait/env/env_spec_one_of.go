/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package env

import (
	"encoding/json"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the EnvSpecOneOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EnvSpecOneOf{}

// EnvSpecOneOf struct for EnvSpecOneOf
type EnvSpecOneOf struct {
	// Specify the environment variables for multiple containers
	Containers []PatchParams `json:"containers"`
}

// NewEnvSpecOneOfWith instantiates a new EnvSpecOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvSpecOneOfWith(containers []PatchParams) *EnvSpecOneOf {
	this := EnvSpecOneOf{}
	this.Containers = containers
	return &this
}

// NewEnvSpecOneOf instantiates a new EnvSpecOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvSpecOneOf() *EnvSpecOneOf {
	this := EnvSpecOneOf{}
	return &this
}

// NewEnvSpecOneOfs converts a list EnvSpecOneOf pointers to objects.
// This is helpful when the SetEnvSpecOneOf requires a list of objects
func NewEnvSpecOneOfList(ps ...*EnvSpecOneOf) []EnvSpecOneOf {
	objs := []EnvSpecOneOf{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetContainers returns the Containers field value
func (o *EnvSpecOneOf) GetContainers() []PatchParams {
	if o == nil {
		var ret []PatchParams
		return ret
	}

	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value
// and a boolean to check if the value has been set.
func (o *EnvSpecOneOf) GetContainersOk() ([]PatchParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.Containers, true
}

// SetContainers sets field value
func (o *EnvSpecOneOf) SetContainers(v []PatchParams) *EnvSpecOneOf {
	o.Containers = v
	return o
}

func (o EnvSpecOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvSpecOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["containers"] = o.Containers
	return toSerialize, nil
}

type NullableEnvSpecOneOf struct {
	value *EnvSpecOneOf
	isSet bool
}

func (v NullableEnvSpecOneOf) Get() *EnvSpecOneOf {
	return v.value
}

func (v *NullableEnvSpecOneOf) Set(val *EnvSpecOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvSpecOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvSpecOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvSpecOneOf(val *EnvSpecOneOf) *NullableEnvSpecOneOf {
	return &NullableEnvSpecOneOf{value: val, isSet: true}
}

func (v NullableEnvSpecOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvSpecOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
