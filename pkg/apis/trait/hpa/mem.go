/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hpa

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

// checks if the Mem type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Mem{}

// Mem struct for Mem
type Mem struct {
	// Specify resource metrics in terms of percentage(\"Utilization\") or direct value(\"AverageValue\")
	Type *string `json:"type"`
	// Specify  the value of MEM utilization or averageValue
	Value *int32 `json:"value"`
}

// NewMemWith instantiates a new Mem object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewMemWith(type_ string, value int32) *Mem {
	this := Mem{}
	this.Type = &type_
	this.Value = &value
	return &this
}

// NewMemWithDefault instantiates a new Mem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemWithDefault() *Mem {
	this := Mem{}
	var type_ string = "Utilization"
	this.Type = &type_
	var value int32 = 50
	this.Value = &value
	return &this
}

// NewMem is short for NewMemWithDefault which instantiates a new Mem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMem() *Mem {
	return NewMemWithDefault()
}

// NewMemEmpty instantiates a new Mem object with no properties set.
// This constructor will not assign any default values to properties.
func NewMemEmpty() *Mem {
	this := Mem{}
	return &this
}

// NewMems converts a list Mem pointers to objects.
// This is helpful when the SetMem requires a list of objects
func NewMemList(ps ...*Mem) []Mem {
	objs := []Mem{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Mem
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Mem) Validate() error {
	if o.Type == nil {
		return errors.New("Type in Mem must be set")
	}
	if o.Value == nil {
		return errors.New("Value in Mem must be set")
	}
	// validate all nested properties
	return nil
}

// GetType returns the Type field value
func (o *Mem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Mem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *Mem) SetType(v string) *Mem {
	o.Type = &v
	return o
}

// GetValue returns the Value field value
func (o *Mem) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Mem) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *Mem) SetValue(v int32) *Mem {
	o.Value = &v
	return o
}

func (o Mem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Mem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableMem struct {
	value *Mem
	isSet bool
}

func (v *NullableMem) Get() *Mem {
	return v.value
}

func (v *NullableMem) Set(val *Mem) {
	v.value = val
	v.isSet = true
}

func (v *NullableMem) IsSet() bool {
	return v.isSet
}

func (v *NullableMem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMem(val *Mem) *NullableMem {
	return &NullableMem{value: val, isSet: true}
}

func (v NullableMem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
