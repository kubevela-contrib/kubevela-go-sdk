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

// checks if the Pvc type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Pvc{}

// Pvc struct for Pvc
type Pvc struct {
	// The name of the PVC
	ClaimName *string `json:"claimName"`
	MountPath *string `json:"mountPath"`
	Name      *string `json:"name"`
	SubPath   *string `json:"subPath,omitempty"`
}

// NewPvcWith instantiates a new Pvc object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewPvcWith(claimName string, mountPath string, name string) *Pvc {
	this := Pvc{}
	this.ClaimName = &claimName
	this.MountPath = &mountPath
	this.Name = &name
	return &this
}

// NewPvcWithDefault instantiates a new Pvc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPvcWithDefault() *Pvc {
	this := Pvc{}
	return &this
}

// NewPvc is short for NewPvcWithDefault which instantiates a new Pvc object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPvc() *Pvc {
	return NewPvcWithDefault()
}

// NewPvcEmpty instantiates a new Pvc object with no properties set.
// This constructor will not assign any default values to properties.
func NewPvcEmpty() *Pvc {
	this := Pvc{}
	return &this
}

// NewPvcs converts a list Pvc pointers to objects.
// This is helpful when the SetPvc requires a list of objects
func NewPvcList(ps ...*Pvc) []Pvc {
	objs := []Pvc{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Pvc
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Pvc) Validate() error {
	if o.ClaimName == nil {
		return errors.New("ClaimName in Pvc must be set")
	}
	if o.MountPath == nil {
		return errors.New("MountPath in Pvc must be set")
	}
	if o.Name == nil {
		return errors.New("Name in Pvc must be set")
	}
	// validate all nested properties
	return nil
}

// GetClaimName returns the ClaimName field value
func (o *Pvc) GetClaimName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.ClaimName
}

// GetClaimNameOk returns a tuple with the ClaimName field value
// and a boolean to check if the value has been set.
func (o *Pvc) GetClaimNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClaimName, true
}

// SetClaimName sets field value
func (o *Pvc) SetClaimName(v string) *Pvc {
	o.ClaimName = &v
	return o
}

// GetMountPath returns the MountPath field value
func (o *Pvc) GetMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value
// and a boolean to check if the value has been set.
func (o *Pvc) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountPath, true
}

// SetMountPath sets field value
func (o *Pvc) SetMountPath(v string) *Pvc {
	o.MountPath = &v
	return o
}

// GetName returns the Name field value
func (o *Pvc) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Pvc) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *Pvc) SetName(v string) *Pvc {
	o.Name = &v
	return o
}

// GetSubPath returns the SubPath field value if set, zero value otherwise.
func (o *Pvc) GetSubPath() string {
	if o == nil || utils.IsNil(o.SubPath) {
		var ret string
		return ret
	}
	return *o.SubPath
}

// GetSubPathOk returns a tuple with the SubPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pvc) GetSubPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SubPath) {
		return nil, false
	}
	return o.SubPath, true
}

// HasSubPath returns a boolean if a field has been set.
func (o *Pvc) HasSubPath() bool {
	if o != nil && !utils.IsNil(o.SubPath) {
		return true
	}

	return false
}

// SetSubPath gets a reference to the given string and assigns it to the subPath field.
// SubPath:
func (o *Pvc) SetSubPath(v string) *Pvc {
	o.SubPath = &v
	return o
}

func (o Pvc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pvc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["claimName"] = o.ClaimName
	toSerialize["mountPath"] = o.MountPath
	toSerialize["name"] = o.Name
	if !utils.IsNil(o.SubPath) {
		toSerialize["subPath"] = o.SubPath
	}
	return toSerialize, nil
}

type NullablePvc struct {
	value *Pvc
	isSet bool
}

func (v *NullablePvc) Get() *Pvc {
	return v.value
}

func (v *NullablePvc) Set(val *Pvc) {
	v.value = val
	v.isSet = true
}

func (v *NullablePvc) IsSet() bool {
	return v.isSet
}

func (v *NullablePvc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePvc(val *Pvc) *NullablePvc {
	return &NullablePvc{value: val, isSet: true}
}

func (v NullablePvc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePvc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
