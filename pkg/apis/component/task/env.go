/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package task

import (
	"encoding/json"
	"errors"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Env type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Env{}

// Env struct for Env
type Env struct {
	// Environment variable name
	Name *string `json:"name"`
	// The value of the environment variable
	Value     *string    `json:"value,omitempty"`
	ValueFrom *ValueFrom `json:"valueFrom,omitempty"`
}

// NewEnvWith instantiates a new Env object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewEnvWith(name string) *Env {
	this := Env{}
	this.Name = &name
	return &this
}

// NewEnvWithDefault instantiates a new Env object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvWithDefault() *Env {
	this := Env{}
	return &this
}

// NewEnv is short for NewEnvWithDefault which instantiates a new Env object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnv() *Env {
	return NewEnvWithDefault()
}

// NewEnvEmpty instantiates a new Env object with no properties set.
// This constructor will not assign any default values to properties.
func NewEnvEmpty() *Env {
	this := Env{}
	return &this
}

// NewEnvs converts a list Env pointers to objects.
// This is helpful when the SetEnv requires a list of objects
func NewEnvList(ps ...*Env) []Env {
	objs := []Env{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Env
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Env) Validate() error {
	if o.Name == nil {
		return errors.New("Name in Env must be set")
	}
	// validate all nested properties
	if o.ValueFrom != nil {
		if err := o.ValueFrom.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// GetName returns the Name field value
func (o *Env) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Env) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *Env) SetName(v string) *Env {
	o.Name = &v
	return o
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Env) GetValue() string {
	if o == nil || utils.IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Env) GetValueOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Env) HasValue() bool {
	if o != nil && !utils.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the value field.
// Value:  The value of the environment variable
func (o *Env) SetValue(v string) *Env {
	o.Value = &v
	return o
}

// GetValueFrom returns the ValueFrom field value if set, zero value otherwise.
func (o *Env) GetValueFrom() ValueFrom {
	if o == nil || utils.IsNil(o.ValueFrom) {
		var ret ValueFrom
		return ret
	}
	return *o.ValueFrom
}

// GetValueFromOk returns a tuple with the ValueFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Env) GetValueFromOk() (*ValueFrom, bool) {
	if o == nil || utils.IsNil(o.ValueFrom) {
		return nil, false
	}
	return o.ValueFrom, true
}

// HasValueFrom returns a boolean if a field has been set.
func (o *Env) HasValueFrom() bool {
	if o != nil && !utils.IsNil(o.ValueFrom) {
		return true
	}

	return false
}

// SetValueFrom gets a reference to the given ValueFrom and assigns it to the valueFrom field.
// ValueFrom:
func (o *Env) SetValueFrom(v ValueFrom) *Env {
	o.ValueFrom = &v
	return o
}

func (o Env) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Env) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !utils.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !utils.IsNil(o.ValueFrom) {
		toSerialize["valueFrom"] = o.ValueFrom
	}
	return toSerialize, nil
}

type NullableEnv struct {
	value *Env
	isSet bool
}

func (v *NullableEnv) Get() *Env {
	return v.value
}

func (v *NullableEnv) Set(val *Env) {
	v.value = val
	v.isSet = true
}

func (v *NullableEnv) IsSet() bool {
	return v.isSet
}

func (v *NullableEnv) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnv(val *Env) *NullableEnv {
	return &NullableEnv{value: val, isSet: true}
}

func (v NullableEnv) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnv) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
