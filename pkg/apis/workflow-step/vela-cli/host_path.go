/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vela_cli

import (
	"encoding/json"
	"errors"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the HostPath type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HostPath{}

// HostPath struct for HostPath
type HostPath struct {
	MountPath *string `json:"mountPath"`
	Name      *string `json:"name"`
	Path      *string `json:"path"`
	Type      *string `json:"type"`
}

// NewHostPathWith instantiates a new HostPath object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewHostPathWith(mountPath string, name string, path string, type_ string) *HostPath {
	this := HostPath{}
	this.MountPath = &mountPath
	this.Name = &name
	this.Path = &path
	this.Type = &type_
	return &this
}

// NewHostPathWithDefault instantiates a new HostPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostPathWithDefault() *HostPath {
	this := HostPath{}
	var type_ string = "Directory"
	this.Type = &type_
	return &this
}

// NewHostPath is short for NewHostPathWithDefault which instantiates a new HostPath object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostPath() *HostPath {
	return NewHostPathWithDefault()
}

// NewHostPathEmpty instantiates a new HostPath object with no properties set.
// This constructor will not assign any default values to properties.
func NewHostPathEmpty() *HostPath {
	this := HostPath{}
	return &this
}

// NewHostPaths converts a list HostPath pointers to objects.
// This is helpful when the SetHostPath requires a list of objects
func NewHostPathList(ps ...*HostPath) []HostPath {
	objs := []HostPath{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this HostPath
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *HostPath) Validate() error {
	if o.MountPath == nil {
		return errors.New("MountPath in HostPath must be set")
	}
	if o.Name == nil {
		return errors.New("Name in HostPath must be set")
	}
	if o.Path == nil {
		return errors.New("Path in HostPath must be set")
	}
	if o.Type == nil {
		return errors.New("Type in HostPath must be set")
	}
	// validate all nested properties
	return nil
}

// GetMountPath returns the MountPath field value
func (o *HostPath) GetMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value
// and a boolean to check if the value has been set.
func (o *HostPath) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountPath, true
}

// SetMountPath sets field value
func (o *HostPath) SetMountPath(v string) *HostPath {
	o.MountPath = &v
	return o
}

// GetName returns the Name field value
func (o *HostPath) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HostPath) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *HostPath) SetName(v string) *HostPath {
	o.Name = &v
	return o
}

// GetPath returns the Path field value
func (o *HostPath) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *HostPath) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path, true
}

// SetPath sets field value
func (o *HostPath) SetPath(v string) *HostPath {
	o.Path = &v
	return o
}

// GetType returns the Type field value
func (o *HostPath) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HostPath) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *HostPath) SetType(v string) *HostPath {
	o.Type = &v
	return o
}

func (o HostPath) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostPath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mountPath"] = o.MountPath
	toSerialize["name"] = o.Name
	toSerialize["path"] = o.Path
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableHostPath struct {
	value *HostPath
	isSet bool
}

func (v *NullableHostPath) Get() *HostPath {
	return v.value
}

func (v *NullableHostPath) Set(val *HostPath) {
	v.value = val
	v.isSet = true
}

func (v *NullableHostPath) IsSet() bool {
	return v.isSet
}

func (v *NullableHostPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostPath(val *HostPath) *NullableHostPath {
	return &NullableHostPath{value: val, isSet: true}
}

func (v NullableHostPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
