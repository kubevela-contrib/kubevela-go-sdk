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

// checks if the Source type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Source{}

// Source struct for Source
type Source struct {
	Properties map[string]interface{} `json:"properties"`
	Type       *string                `json:"type"`
}

// NewSourceWith instantiates a new Source object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewSourceWith(properties map[string]interface{}, type_ string) *Source {
	this := Source{}
	this.Properties = properties
	this.Type = &type_
	return &this
}

// NewSourceWithDefault instantiates a new Source object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceWithDefault() *Source {
	this := Source{}
	return &this
}

// NewSource is short for NewSourceWithDefault which instantiates a new Source object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSource() *Source {
	return NewSourceWithDefault()
}

// NewSourceEmpty instantiates a new Source object with no properties set.
// This constructor will not assign any default values to properties.
func NewSourceEmpty() *Source {
	this := Source{}
	return &this
}

// NewSources converts a list Source pointers to objects.
// This is helpful when the SetSource requires a list of objects
func NewSourceList(ps ...*Source) []Source {
	objs := []Source{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Source
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Source) Validate() error {
	if o.Properties == nil {
		return errors.New("Properties in Source must be set")
	}
	if o.Type == nil {
		return errors.New("Type in Source must be set")
	}
	// validate all nested properties
	return nil
}

// GetProperties returns the Properties field value
func (o *Source) GetProperties() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *Source) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *Source) SetProperties(v map[string]interface{}) *Source {
	o.Properties = v
	return o
}

// GetType returns the Type field value
func (o *Source) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Source) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *Source) SetType(v string) *Source {
	o.Type = &v
	return o
}

func (o Source) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Source) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["properties"] = o.Properties
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableSource struct {
	value *Source
	isSet bool
}

func (v *NullableSource) Get() *Source {
	return v.value
}

func (v *NullableSource) Set(val *Source) {
	v.value = val
	v.isSet = true
}

func (v *NullableSource) IsSet() bool {
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
