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

// checks if the ConfigMap type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ConfigMap{}

// ConfigMap struct for ConfigMap
type ConfigMap struct {
	CmName      *string `json:"cmName,omitempty"`
	DefaultMode *int32  `json:"defaultMode,omitempty"`
	Items       []Items `json:"items,omitempty"`
	MountPath   *string `json:"mountPath,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// NewConfigMapWith instantiates a new ConfigMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigMapWith() *ConfigMap {
	this := ConfigMap{}
	var defaultMode int32 = 420
	this.DefaultMode = &defaultMode
	return &this
}

// NewConfigMap instantiates a new ConfigMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigMap() *ConfigMap {
	this := ConfigMap{}
	var defaultMode int32 = 420
	this.DefaultMode = &defaultMode
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

// GetCmName returns the CmName field value if set, zero value otherwise.
func (o *ConfigMap) GetCmName() string {
	if o == nil || utils.IsNil(o.CmName) {
		var ret string
		return ret
	}
	return *o.CmName
}

// GetCmNameOk returns a tuple with the CmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetCmNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CmName) {
		return nil, false
	}
	return o.CmName, true
}

// HasCmName returns a boolean if a field has been set.
func (o *ConfigMap) HasCmName() bool {
	if o != nil && !utils.IsNil(o.CmName) {
		return true
	}

	return false
}

// SetCmName gets a reference to the given string and assigns it to the cmName field.
// CmName:
func (o *ConfigMap) SetCmName(v string) *ConfigMap {
	o.CmName = &v
	return o
}

// GetDefaultMode returns the DefaultMode field value if set, zero value otherwise.
func (o *ConfigMap) GetDefaultMode() int32 {
	if o == nil || utils.IsNil(o.DefaultMode) {
		var ret int32
		return ret
	}
	return *o.DefaultMode
}

// GetDefaultModeOk returns a tuple with the DefaultMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetDefaultModeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.DefaultMode) {
		return nil, false
	}
	return o.DefaultMode, true
}

// HasDefaultMode returns a boolean if a field has been set.
func (o *ConfigMap) HasDefaultMode() bool {
	if o != nil && !utils.IsNil(o.DefaultMode) {
		return true
	}

	return false
}

// SetDefaultMode gets a reference to the given int32 and assigns it to the defaultMode field.
// DefaultMode:
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

// GetMountPath returns the MountPath field value if set, zero value otherwise.
func (o *ConfigMap) GetMountPath() string {
	if o == nil || utils.IsNil(o.MountPath) {
		var ret string
		return ret
	}
	return *o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetMountPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MountPath) {
		return nil, false
	}
	return o.MountPath, true
}

// HasMountPath returns a boolean if a field has been set.
func (o *ConfigMap) HasMountPath() bool {
	if o != nil && !utils.IsNil(o.MountPath) {
		return true
	}

	return false
}

// SetMountPath gets a reference to the given string and assigns it to the mountPath field.
// MountPath:
func (o *ConfigMap) SetMountPath(v string) *ConfigMap {
	o.MountPath = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConfigMap) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConfigMap) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:
func (o *ConfigMap) SetName(v string) *ConfigMap {
	o.Name = &v
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
	if !utils.IsNil(o.CmName) {
		toSerialize["cmName"] = o.CmName
	}
	if !utils.IsNil(o.DefaultMode) {
		toSerialize["defaultMode"] = o.DefaultMode
	}
	if !utils.IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !utils.IsNil(o.MountPath) {
		toSerialize["mountPath"] = o.MountPath
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableConfigMap struct {
	value *ConfigMap
	isSet bool
}

func (v NullableConfigMap) Get() *ConfigMap {
	return v.value
}

func (v *NullableConfigMap) Set(val *ConfigMap) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigMap) IsSet() bool {
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
