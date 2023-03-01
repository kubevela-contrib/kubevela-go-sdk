/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package command

import (
	"encoding/json"

	"github.com/kubevela-contrib/vela-go-sdk/pkg/apis/utils"
)

// checks if the CommandSpecOneOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CommandSpecOneOf{}

// CommandSpecOneOf struct for CommandSpecOneOf
type CommandSpecOneOf struct {
	// Specify the commands for multiple containers
	Containers []PatchParams `json:"containers"`
}

// NewCommandSpecOneOfWith instantiates a new CommandSpecOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandSpecOneOfWith(containers []PatchParams) *CommandSpecOneOf {
	this := CommandSpecOneOf{}
	this.Containers = containers
	return &this
}

// NewCommandSpecOneOf instantiates a new CommandSpecOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandSpecOneOf() *CommandSpecOneOf {
	this := CommandSpecOneOf{}
	return &this
}

// NewCommandSpecOneOfs converts a list CommandSpecOneOf pointers to objects.
// This is helpful when the SetCommandSpecOneOf requires a list of objects
func NewCommandSpecOneOfList(ps ...*CommandSpecOneOf) []CommandSpecOneOf {
	objs := []CommandSpecOneOf{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetContainers returns the Containers field value
func (o *CommandSpecOneOf) GetContainers() []PatchParams {
	if o == nil {
		var ret []PatchParams
		return ret
	}

	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value
// and a boolean to check if the value has been set.
func (o *CommandSpecOneOf) GetContainersOk() ([]PatchParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.Containers, true
}

// SetContainers sets field value
func (o *CommandSpecOneOf) SetContainers(v []PatchParams) *CommandSpecOneOf {
	o.Containers = v
	return o
}

func (o CommandSpecOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommandSpecOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["containers"] = o.Containers
	return toSerialize, nil
}

type NullableCommandSpecOneOf struct {
	value *CommandSpecOneOf
	isSet bool
}

func (v NullableCommandSpecOneOf) Get() *CommandSpecOneOf {
	return v.value
}

func (v *NullableCommandSpecOneOf) Set(val *CommandSpecOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandSpecOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandSpecOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandSpecOneOf(val *CommandSpecOneOf) *NullableCommandSpecOneOf {
	return &NullableCommandSpecOneOf{value: val, isSet: true}
}

func (v NullableCommandSpecOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandSpecOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
