/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package trigger_service

import (
	"encoding/json"
	"errors"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Action type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Action{}

// Action struct for Action
type Action struct {
	Properties map[string]interface{} `json:"properties"`
	Type       *string                `json:"type"`
}

// NewActionWith instantiates a new Action object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewActionWith(properties map[string]interface{}, type_ string) *Action {
	this := Action{}
	this.Properties = properties
	this.Type = &type_
	return &this
}

// NewActionWithDefault instantiates a new Action object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionWithDefault() *Action {
	this := Action{}
	return &this
}

// NewAction is short for NewActionWithDefault which instantiates a new Action object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAction() *Action {
	return NewActionWithDefault()
}

// NewActionEmpty instantiates a new Action object with no properties set.
// This constructor will not assign any default values to properties.
func NewActionEmpty() *Action {
	this := Action{}
	return &this
}

// NewActions converts a list Action pointers to objects.
// This is helpful when the SetAction requires a list of objects
func NewActionList(ps ...*Action) []Action {
	objs := []Action{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Action
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Action) Validate() error {
	if o.Properties == nil {
		return errors.New("Properties in Action must be set")
	}
	if o.Type == nil {
		return errors.New("Type in Action must be set")
	}
	// validate all nested properties
	return nil
}

// GetProperties returns the Properties field value
func (o *Action) GetProperties() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *Action) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *Action) SetProperties(v map[string]interface{}) *Action {
	o.Properties = v
	return o
}

// GetType returns the Type field value
func (o *Action) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Action) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *Action) SetType(v string) *Action {
	o.Type = &v
	return o
}

func (o Action) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Action) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["properties"] = o.Properties
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAction struct {
	value *Action
	isSet bool
}

func (v *NullableAction) Get() *Action {
	return v.value
}

func (v *NullableAction) Set(val *Action) {
	v.value = val
	v.isSet = true
}

func (v *NullableAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAction(val *Action) *NullableAction {
	return &NullableAction{value: val, isSet: true}
}

func (v NullableAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
