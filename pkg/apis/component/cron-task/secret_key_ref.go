/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cron_task

import (
	"encoding/json"
	"errors"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the SecretKeyRef type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SecretKeyRef{}

// SecretKeyRef Selects a key of a secret in the pod's namespace
type SecretKeyRef struct {
	// The key of the secret to select from. Must be a valid secret key
	Key *string `json:"key"`
	// The name of the secret in the pod's namespace to select from
	Name *string `json:"name"`
}

// NewSecretKeyRefWith instantiates a new SecretKeyRef object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewSecretKeyRefWith(key string, name string) *SecretKeyRef {
	this := SecretKeyRef{}
	this.Key = &key
	this.Name = &name
	return &this
}

// NewSecretKeyRefWithDefault instantiates a new SecretKeyRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretKeyRefWithDefault() *SecretKeyRef {
	this := SecretKeyRef{}
	return &this
}

// NewSecretKeyRef is short for NewSecretKeyRefWithDefault which instantiates a new SecretKeyRef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretKeyRef() *SecretKeyRef {
	return NewSecretKeyRefWithDefault()
}

// NewSecretKeyRefEmpty instantiates a new SecretKeyRef object with no properties set.
// This constructor will not assign any default values to properties.
func NewSecretKeyRefEmpty() *SecretKeyRef {
	this := SecretKeyRef{}
	return &this
}

// NewSecretKeyRefs converts a list SecretKeyRef pointers to objects.
// This is helpful when the SetSecretKeyRef requires a list of objects
func NewSecretKeyRefList(ps ...*SecretKeyRef) []SecretKeyRef {
	objs := []SecretKeyRef{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this SecretKeyRef
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *SecretKeyRef) Validate() error {
	if o.Key == nil {
		return errors.New("Key in SecretKeyRef must be set")
	}
	if o.Name == nil {
		return errors.New("Name in SecretKeyRef must be set")
	}
	// validate all nested properties
	return nil
}

// GetKey returns the Key field value
func (o *SecretKeyRef) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SecretKeyRef) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key, true
}

// SetKey sets field value
func (o *SecretKeyRef) SetKey(v string) *SecretKeyRef {
	o.Key = &v
	return o
}

// GetName returns the Name field value
func (o *SecretKeyRef) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecretKeyRef) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *SecretKeyRef) SetName(v string) *SecretKeyRef {
	o.Name = &v
	return o
}

func (o SecretKeyRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretKeyRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableSecretKeyRef struct {
	value *SecretKeyRef
	isSet bool
}

func (v *NullableSecretKeyRef) Get() *SecretKeyRef {
	return v.value
}

func (v *NullableSecretKeyRef) Set(val *SecretKeyRef) {
	v.value = val
	v.isSet = true
}

func (v *NullableSecretKeyRef) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretKeyRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretKeyRef(val *SecretKeyRef) *NullableSecretKeyRef {
	return &NullableSecretKeyRef{value: val, isSet: true}
}

func (v NullableSecretKeyRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretKeyRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
