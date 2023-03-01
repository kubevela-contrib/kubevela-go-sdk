/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package daemon

import (
	"encoding/json"

	"github.com/kubevela-contrib/vela-go-sdk/pkg/apis/utils"
)

// checks if the EmptyDir type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EmptyDir{}

// EmptyDir struct for EmptyDir
type EmptyDir struct {
	Medium    *string `json:"medium,omitempty"`
	MountPath *string `json:"mountPath,omitempty"`
	Name      *string `json:"name,omitempty"`
}

// NewEmptyDirWith instantiates a new EmptyDir object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmptyDirWith() *EmptyDir {
	this := EmptyDir{}
	var medium string = ""
	this.Medium = &medium
	return &this
}

// NewEmptyDir instantiates a new EmptyDir object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmptyDir() *EmptyDir {
	this := EmptyDir{}
	var medium string = ""
	this.Medium = &medium
	return &this
}

// NewEmptyDirs converts a list EmptyDir pointers to objects.
// This is helpful when the SetEmptyDir requires a list of objects
func NewEmptyDirList(ps ...*EmptyDir) []EmptyDir {
	objs := []EmptyDir{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetMedium returns the Medium field value if set, zero value otherwise.
func (o *EmptyDir) GetMedium() string {
	if o == nil || utils.IsNil(o.Medium) {
		var ret string
		return ret
	}
	return *o.Medium
}

// GetMediumOk returns a tuple with the Medium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmptyDir) GetMediumOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Medium) {
		return nil, false
	}
	return o.Medium, true
}

// HasMedium returns a boolean if a field has been set.
func (o *EmptyDir) HasMedium() bool {
	if o != nil && !utils.IsNil(o.Medium) {
		return true
	}

	return false
}

// SetMedium gets a reference to the given string and assigns it to the medium field.
// Medium:
func (o *EmptyDir) SetMedium(v string) *EmptyDir {
	o.Medium = &v
	return o
}

// GetMountPath returns the MountPath field value if set, zero value otherwise.
func (o *EmptyDir) GetMountPath() string {
	if o == nil || utils.IsNil(o.MountPath) {
		var ret string
		return ret
	}
	return *o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmptyDir) GetMountPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MountPath) {
		return nil, false
	}
	return o.MountPath, true
}

// HasMountPath returns a boolean if a field has been set.
func (o *EmptyDir) HasMountPath() bool {
	if o != nil && !utils.IsNil(o.MountPath) {
		return true
	}

	return false
}

// SetMountPath gets a reference to the given string and assigns it to the mountPath field.
// MountPath:
func (o *EmptyDir) SetMountPath(v string) *EmptyDir {
	o.MountPath = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EmptyDir) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmptyDir) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EmptyDir) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:
func (o *EmptyDir) SetName(v string) *EmptyDir {
	o.Name = &v
	return o
}

func (o EmptyDir) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmptyDir) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Medium) {
		toSerialize["medium"] = o.Medium
	}
	if !utils.IsNil(o.MountPath) {
		toSerialize["mountPath"] = o.MountPath
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableEmptyDir struct {
	value *EmptyDir
	isSet bool
}

func (v NullableEmptyDir) Get() *EmptyDir {
	return v.value
}

func (v *NullableEmptyDir) Set(val *EmptyDir) {
	v.value = val
	v.isSet = true
}

func (v NullableEmptyDir) IsSet() bool {
	return v.isSet
}

func (v *NullableEmptyDir) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmptyDir(val *EmptyDir) *NullableEmptyDir {
	return &NullableEmptyDir{value: val, isSet: true}
}

func (v NullableEmptyDir) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmptyDir) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
