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

// checks if the ConfigMap type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ConfigMap{}

// ConfigMap struct for ConfigMap
type ConfigMap struct {
	CmName      *string `json:"cmName"`
	DefaultMode *int32  `json:"defaultMode"`
	Items       []Items `json:"items,omitempty"`
	MountPath   *string `json:"mountPath"`
	Name        *string `json:"name"`
	SubPath     *string `json:"subPath,omitempty"`
}

// NewConfigMapWith instantiates a new ConfigMap object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewConfigMapWith(cmName string, defaultMode int32, mountPath string, name string) *ConfigMap {
	this := ConfigMap{}
	this.CmName = &cmName
	this.DefaultMode = &defaultMode
	this.MountPath = &mountPath
	this.Name = &name
	return &this
}

// NewConfigMapWithDefault instantiates a new ConfigMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigMapWithDefault() *ConfigMap {
	this := ConfigMap{}
	var defaultMode int32 = 420
	this.DefaultMode = &defaultMode
	return &this
}

// NewConfigMap is short for NewConfigMapWithDefault which instantiates a new ConfigMap object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigMap() *ConfigMap {
	return NewConfigMapWithDefault()
}

// NewConfigMapEmpty instantiates a new ConfigMap object with no properties set.
// This constructor will not assign any default values to properties.
func NewConfigMapEmpty() *ConfigMap {
	this := ConfigMap{}
	return &this
}

// NewConfigMaps converts a list ConfigMap pointers to objects.
// This is helpful when the SetConfigMap requires a list of objects
func NewConfigMapList(ps ...*ConfigMap) []ConfigMap {
	objs := []ConfigMap{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this ConfigMap
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *ConfigMap) Validate() error {
	if o.CmName == nil {
		return errors.New("CmName in ConfigMap must be set")
	}
	if o.DefaultMode == nil {
		return errors.New("DefaultMode in ConfigMap must be set")
	}
	if o.MountPath == nil {
		return errors.New("MountPath in ConfigMap must be set")
	}
	if o.Name == nil {
		return errors.New("Name in ConfigMap must be set")
	}
	// validate all nested properties
	return nil
}

// GetCmName returns the CmName field value
func (o *ConfigMap) GetCmName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.CmName
}

// GetCmNameOk returns a tuple with the CmName field value
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetCmNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CmName, true
}

// SetCmName sets field value
func (o *ConfigMap) SetCmName(v string) *ConfigMap {
	o.CmName = &v
	return o
}

// GetDefaultMode returns the DefaultMode field value
func (o *ConfigMap) GetDefaultMode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.DefaultMode
}

// GetDefaultModeOk returns a tuple with the DefaultMode field value
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetDefaultModeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultMode, true
}

// SetDefaultMode sets field value
func (o *ConfigMap) SetDefaultMode(v int32) *ConfigMap {
	o.DefaultMode = &v
	return o
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ConfigMap) GetItems() []Items {
	if o == nil || utils.IsNil(o.Items) {
		var ret []Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetItemsOk() ([]Items, bool) {
	if o == nil || utils.IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ConfigMap) HasItems() bool {
	if o != nil && !utils.IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Items and assigns it to the items field.
// Items:
func (o *ConfigMap) SetItems(v []Items) *ConfigMap {
	o.Items = v
	return o
}

// GetMountPath returns the MountPath field value
func (o *ConfigMap) GetMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountPath, true
}

// SetMountPath sets field value
func (o *ConfigMap) SetMountPath(v string) *ConfigMap {
	o.MountPath = &v
	return o
}

// GetName returns the Name field value
func (o *ConfigMap) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *ConfigMap) SetName(v string) *ConfigMap {
	o.Name = &v
	return o
}

// GetSubPath returns the SubPath field value if set, zero value otherwise.
func (o *ConfigMap) GetSubPath() string {
	if o == nil || utils.IsNil(o.SubPath) {
		var ret string
		return ret
	}
	return *o.SubPath
}

// GetSubPathOk returns a tuple with the SubPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetSubPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SubPath) {
		return nil, false
	}
	return o.SubPath, true
}

// HasSubPath returns a boolean if a field has been set.
func (o *ConfigMap) HasSubPath() bool {
	if o != nil && !utils.IsNil(o.SubPath) {
		return true
	}

	return false
}

// SetSubPath gets a reference to the given string and assigns it to the subPath field.
// SubPath:
func (o *ConfigMap) SetSubPath(v string) *ConfigMap {
	o.SubPath = &v
	return o
}

func (o ConfigMap) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigMap) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cmName"] = o.CmName
	toSerialize["defaultMode"] = o.DefaultMode
	if !utils.IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	toSerialize["mountPath"] = o.MountPath
	toSerialize["name"] = o.Name
	if !utils.IsNil(o.SubPath) {
		toSerialize["subPath"] = o.SubPath
	}
	return toSerialize, nil
}

type NullableConfigMap struct {
	value *ConfigMap
	isSet bool
}

func (v *NullableConfigMap) Get() *ConfigMap {
	return v.value
}

func (v *NullableConfigMap) Set(val *ConfigMap) {
	v.value = val
	v.isSet = true
}

func (v *NullableConfigMap) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigMap(val *ConfigMap) *NullableConfigMap {
	return &NullableConfigMap{value: val, isSet: true}
}

func (v NullableConfigMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
