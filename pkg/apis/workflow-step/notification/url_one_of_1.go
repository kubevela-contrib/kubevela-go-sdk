/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"
	"errors"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the UrlOneOf1 type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UrlOneOf1{}

// UrlOneOf1 struct for UrlOneOf1
type UrlOneOf1 struct {
	SecretRef *SecretRef `json:"secretRef"`
}

// NewUrlOneOf1With instantiates a new UrlOneOf1 object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewUrlOneOf1With(secretRef SecretRef) *UrlOneOf1 {
	this := UrlOneOf1{}
	this.SecretRef = &secretRef
	return &this
}

// NewUrlOneOf1WithDefault instantiates a new UrlOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrlOneOf1WithDefault() *UrlOneOf1 {
	this := UrlOneOf1{}
	return &this
}

// NewUrlOneOf1 is short for NewUrlOneOf1WithDefault which instantiates a new UrlOneOf1 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrlOneOf1() *UrlOneOf1 {
	return NewUrlOneOf1WithDefault()
}

// NewUrlOneOf1Empty instantiates a new UrlOneOf1 object with no properties set.
// This constructor will not assign any default values to properties.
func NewUrlOneOf1Empty() *UrlOneOf1 {
	this := UrlOneOf1{}
	return &this
}

// NewUrlOneOf1s converts a list UrlOneOf1 pointers to objects.
// This is helpful when the SetUrlOneOf1 requires a list of objects
func NewUrlOneOf1List(ps ...*UrlOneOf1) []UrlOneOf1 {
	objs := []UrlOneOf1{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this UrlOneOf1
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *UrlOneOf1) Validate() error {
	if o.SecretRef == nil {
		return errors.New("SecretRef in UrlOneOf1 must be set")
	}
	// validate all nested properties
	if o.SecretRef != nil {
		if err := o.SecretRef.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// GetSecretRef returns the SecretRef field value
func (o *UrlOneOf1) GetSecretRef() SecretRef {
	if o == nil {
		var ret SecretRef
		return ret
	}

	return *o.SecretRef
}

// GetSecretRefOk returns a tuple with the SecretRef field value
// and a boolean to check if the value has been set.
func (o *UrlOneOf1) GetSecretRefOk() (*SecretRef, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretRef, true
}

// SetSecretRef sets field value
func (o *UrlOneOf1) SetSecretRef(v SecretRef) *UrlOneOf1 {
	o.SecretRef = &v
	return o
}

func (o UrlOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UrlOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["secretRef"] = o.SecretRef
	return toSerialize, nil
}

type NullableUrlOneOf1 struct {
	value *UrlOneOf1
	isSet bool
}

func (v *NullableUrlOneOf1) Get() *UrlOneOf1 {
	return v.value
}

func (v *NullableUrlOneOf1) Set(val *UrlOneOf1) {
	v.value = val
	v.isSet = true
}

func (v *NullableUrlOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableUrlOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrlOneOf1(val *UrlOneOf1) *NullableUrlOneOf1 {
	return &NullableUrlOneOf1{value: val, isSet: true}
}

func (v NullableUrlOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrlOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
