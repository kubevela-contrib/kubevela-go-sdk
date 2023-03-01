/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package daemon

import (
	"encoding/json"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Pvc type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Pvc{}

// Pvc struct for Pvc
type Pvc struct {
	// The name of the PVC
	ClaimName *string `json:"claimName,omitempty"`
	MountPath *string `json:"mountPath,omitempty"`
	Name      *string `json:"name,omitempty"`
}

// NewPvcWith instantiates a new Pvc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPvcWith() *Pvc {
	this := Pvc{}
	return &this
}

// NewPvc instantiates a new Pvc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPvc() *Pvc {
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

// GetClaimName returns the ClaimName field value if set, zero value otherwise.
func (o *Pvc) GetClaimName() string {
	if o == nil || utils.IsNil(o.ClaimName) {
		var ret string
		return ret
	}
	return *o.ClaimName
}

// GetClaimNameOk returns a tuple with the ClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pvc) GetClaimNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ClaimName) {
		return nil, false
	}
	return o.ClaimName, true
}

// HasClaimName returns a boolean if a field has been set.
func (o *Pvc) HasClaimName() bool {
	if o != nil && !utils.IsNil(o.ClaimName) {
		return true
	}

	return false
}

// SetClaimName gets a reference to the given string and assigns it to the claimName field.
// ClaimName:  The name of the PVC
func (o *Pvc) SetClaimName(v string) *Pvc {
	o.ClaimName = &v
	return o
}

// GetMountPath returns the MountPath field value if set, zero value otherwise.
func (o *Pvc) GetMountPath() string {
	if o == nil || utils.IsNil(o.MountPath) {
		var ret string
		return ret
	}
	return *o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pvc) GetMountPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MountPath) {
		return nil, false
	}
	return o.MountPath, true
}

// HasMountPath returns a boolean if a field has been set.
func (o *Pvc) HasMountPath() bool {
	if o != nil && !utils.IsNil(o.MountPath) {
		return true
	}

	return false
}

// SetMountPath gets a reference to the given string and assigns it to the mountPath field.
// MountPath:
func (o *Pvc) SetMountPath(v string) *Pvc {
	o.MountPath = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Pvc) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pvc) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Pvc) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:
func (o *Pvc) SetName(v string) *Pvc {
	o.Name = &v
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
	if !utils.IsNil(o.ClaimName) {
		toSerialize["claimName"] = o.ClaimName
	}
	if !utils.IsNil(o.MountPath) {
		toSerialize["mountPath"] = o.MountPath
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullablePvc struct {
	value *Pvc
	isSet bool
}

func (v NullablePvc) Get() *Pvc {
	return v.value
}

func (v *NullablePvc) Set(val *Pvc) {
	v.value = val
	v.isSet = true
}

func (v NullablePvc) IsSet() bool {
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
