/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sidecar

import (
	"encoding/json"

	"github.com/kubevela-contrib/vela-go-sdk/pkg/apis/utils"
)

// checks if the FieldRef type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FieldRef{}

// FieldRef Specify the field reference for env
type FieldRef struct {
	// Specify the field path for env
	FieldPath *string `json:"fieldPath,omitempty"`
}

// NewFieldRefWith instantiates a new FieldRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldRefWith() *FieldRef {
	this := FieldRef{}
	return &this
}

// NewFieldRef instantiates a new FieldRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldRef() *FieldRef {
	this := FieldRef{}
	return &this
}

// NewFieldRefs converts a list FieldRef pointers to objects.
// This is helpful when the SetFieldRef requires a list of objects
func NewFieldRefList(ps ...*FieldRef) []FieldRef {
	objs := []FieldRef{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetFieldPath returns the FieldPath field value if set, zero value otherwise.
func (o *FieldRef) GetFieldPath() string {
	if o == nil || utils.IsNil(o.FieldPath) {
		var ret string
		return ret
	}
	return *o.FieldPath
}

// GetFieldPathOk returns a tuple with the FieldPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldRef) GetFieldPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FieldPath) {
		return nil, false
	}
	return o.FieldPath, true
}

// HasFieldPath returns a boolean if a field has been set.
func (o *FieldRef) HasFieldPath() bool {
	if o != nil && !utils.IsNil(o.FieldPath) {
		return true
	}

	return false
}

// SetFieldPath gets a reference to the given string and assigns it to the fieldPath field.
// FieldPath:  Specify the field path for env
func (o *FieldRef) SetFieldPath(v string) *FieldRef {
	o.FieldPath = &v
	return o
}

func (o FieldRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.FieldPath) {
		toSerialize["fieldPath"] = o.FieldPath
	}
	return toSerialize, nil
}

type NullableFieldRef struct {
	value *FieldRef
	isSet bool
}

func (v NullableFieldRef) Get() *FieldRef {
	return v.value
}

func (v *NullableFieldRef) Set(val *FieldRef) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldRef) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldRef(val *FieldRef) *NullableFieldRef {
	return &NullableFieldRef{value: val, isSet: true}
}

func (v NullableFieldRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
