/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"
	"errors"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Secret type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Secret{}

// Secret struct for Secret
type Secret struct {
	Data        map[string]interface{} `json:"data,omitempty"`
	DefaultMode *int32                 `json:"defaultMode"`
	Items       []Items                `json:"items,omitempty"`
	MountOnly   *bool                  `json:"mountOnly"`
	MountPath   *string                `json:"mountPath"`
	MountToEnv  *MountToEnv1           `json:"mountToEnv,omitempty"`
	MountToEnvs []MountToEnvs1         `json:"mountToEnvs,omitempty"`
	Name        *string                `json:"name"`
	ReadOnly    *bool                  `json:"readOnly"`
	StringData  map[string]interface{} `json:"stringData,omitempty"`
	SubPath     *string                `json:"subPath,omitempty"`
}

// NewSecretWith instantiates a new Secret object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewSecretWith(defaultMode int32, mountOnly bool, mountPath string, name string, readOnly bool) *Secret {
	this := Secret{}
	this.DefaultMode = &defaultMode
	this.MountOnly = &mountOnly
	this.MountPath = &mountPath
	this.Name = &name
	this.ReadOnly = &readOnly
	return &this
}

// NewSecretWithDefault instantiates a new Secret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretWithDefault() *Secret {
	this := Secret{}
	var defaultMode int32 = 420
	this.DefaultMode = &defaultMode
	var mountOnly bool = false
	this.MountOnly = &mountOnly
	var readOnly bool = false
	this.ReadOnly = &readOnly
	return &this
}

// NewSecret is short for NewSecretWithDefault which instantiates a new Secret object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecret() *Secret {
	return NewSecretWithDefault()
}

// NewSecretEmpty instantiates a new Secret object with no properties set.
// This constructor will not assign any default values to properties.
func NewSecretEmpty() *Secret {
	this := Secret{}
	return &this
}

// NewSecrets converts a list Secret pointers to objects.
// This is helpful when the SetSecret requires a list of objects
func NewSecretList(ps ...*Secret) []Secret {
	objs := []Secret{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Secret
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Secret) Validate() error {
	if o.DefaultMode == nil {
		return errors.New("DefaultMode in Secret must be set")
	}
	if o.MountOnly == nil {
		return errors.New("MountOnly in Secret must be set")
	}
	if o.MountPath == nil {
		return errors.New("MountPath in Secret must be set")
	}
	if o.Name == nil {
		return errors.New("Name in Secret must be set")
	}
	if o.ReadOnly == nil {
		return errors.New("ReadOnly in Secret must be set")
	}
	// validate all nested properties
	if o.MountToEnv != nil {
		if err := o.MountToEnv.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Secret) GetData() map[string]interface{} {
	if o == nil || utils.IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Secret) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the data field.
// Data:
func (o *Secret) SetData(v map[string]interface{}) *Secret {
	o.Data = v
	return o
}

// GetDefaultMode returns the DefaultMode field value
func (o *Secret) GetDefaultMode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.DefaultMode
}

// GetDefaultModeOk returns a tuple with the DefaultMode field value
// and a boolean to check if the value has been set.
func (o *Secret) GetDefaultModeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultMode, true
}

// SetDefaultMode sets field value
func (o *Secret) SetDefaultMode(v int32) *Secret {
	o.DefaultMode = &v
	return o
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Secret) GetItems() []Items {
	if o == nil || utils.IsNil(o.Items) {
		var ret []Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetItemsOk() ([]Items, bool) {
	if o == nil || utils.IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Secret) HasItems() bool {
	if o != nil && !utils.IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Items and assigns it to the items field.
// Items:
func (o *Secret) SetItems(v []Items) *Secret {
	o.Items = v
	return o
}

// GetMountOnly returns the MountOnly field value
func (o *Secret) GetMountOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return *o.MountOnly
}

// GetMountOnlyOk returns a tuple with the MountOnly field value
// and a boolean to check if the value has been set.
func (o *Secret) GetMountOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountOnly, true
}

// SetMountOnly sets field value
func (o *Secret) SetMountOnly(v bool) *Secret {
	o.MountOnly = &v
	return o
}

// GetMountPath returns the MountPath field value
func (o *Secret) GetMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value
// and a boolean to check if the value has been set.
func (o *Secret) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountPath, true
}

// SetMountPath sets field value
func (o *Secret) SetMountPath(v string) *Secret {
	o.MountPath = &v
	return o
}

// GetMountToEnv returns the MountToEnv field value if set, zero value otherwise.
func (o *Secret) GetMountToEnv() MountToEnv1 {
	if o == nil || utils.IsNil(o.MountToEnv) {
		var ret MountToEnv1
		return ret
	}
	return *o.MountToEnv
}

// GetMountToEnvOk returns a tuple with the MountToEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetMountToEnvOk() (*MountToEnv1, bool) {
	if o == nil || utils.IsNil(o.MountToEnv) {
		return nil, false
	}
	return o.MountToEnv, true
}

// HasMountToEnv returns a boolean if a field has been set.
func (o *Secret) HasMountToEnv() bool {
	if o != nil && !utils.IsNil(o.MountToEnv) {
		return true
	}

	return false
}

// SetMountToEnv gets a reference to the given MountToEnv1 and assigns it to the mountToEnv field.
// MountToEnv:
func (o *Secret) SetMountToEnv(v MountToEnv1) *Secret {
	o.MountToEnv = &v
	return o
}

// GetMountToEnvs returns the MountToEnvs field value if set, zero value otherwise.
func (o *Secret) GetMountToEnvs() []MountToEnvs1 {
	if o == nil || utils.IsNil(o.MountToEnvs) {
		var ret []MountToEnvs1
		return ret
	}
	return o.MountToEnvs
}

// GetMountToEnvsOk returns a tuple with the MountToEnvs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetMountToEnvsOk() ([]MountToEnvs1, bool) {
	if o == nil || utils.IsNil(o.MountToEnvs) {
		return nil, false
	}
	return o.MountToEnvs, true
}

// HasMountToEnvs returns a boolean if a field has been set.
func (o *Secret) HasMountToEnvs() bool {
	if o != nil && !utils.IsNil(o.MountToEnvs) {
		return true
	}

	return false
}

// SetMountToEnvs gets a reference to the given []MountToEnvs1 and assigns it to the mountToEnvs field.
// MountToEnvs:
func (o *Secret) SetMountToEnvs(v []MountToEnvs1) *Secret {
	o.MountToEnvs = v
	return o
}

// GetName returns the Name field value
func (o *Secret) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Secret) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *Secret) SetName(v string) *Secret {
	o.Name = &v
	return o
}

// GetReadOnly returns the ReadOnly field value
func (o *Secret) GetReadOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value
// and a boolean to check if the value has been set.
func (o *Secret) GetReadOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// SetReadOnly sets field value
func (o *Secret) SetReadOnly(v bool) *Secret {
	o.ReadOnly = &v
	return o
}

// GetStringData returns the StringData field value if set, zero value otherwise.
func (o *Secret) GetStringData() map[string]interface{} {
	if o == nil || utils.IsNil(o.StringData) {
		var ret map[string]interface{}
		return ret
	}
	return o.StringData
}

// GetStringDataOk returns a tuple with the StringData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetStringDataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.StringData) {
		return map[string]interface{}{}, false
	}
	return o.StringData, true
}

// HasStringData returns a boolean if a field has been set.
func (o *Secret) HasStringData() bool {
	if o != nil && !utils.IsNil(o.StringData) {
		return true
	}

	return false
}

// SetStringData gets a reference to the given map[string]interface{} and assigns it to the stringData field.
// StringData:
func (o *Secret) SetStringData(v map[string]interface{}) *Secret {
	o.StringData = v
	return o
}

// GetSubPath returns the SubPath field value if set, zero value otherwise.
func (o *Secret) GetSubPath() string {
	if o == nil || utils.IsNil(o.SubPath) {
		var ret string
		return ret
	}
	return *o.SubPath
}

// GetSubPathOk returns a tuple with the SubPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetSubPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SubPath) {
		return nil, false
	}
	return o.SubPath, true
}

// HasSubPath returns a boolean if a field has been set.
func (o *Secret) HasSubPath() bool {
	if o != nil && !utils.IsNil(o.SubPath) {
		return true
	}

	return false
}

// SetSubPath gets a reference to the given string and assigns it to the subPath field.
// SubPath:
func (o *Secret) SetSubPath(v string) *Secret {
	o.SubPath = &v
	return o
}

func (o Secret) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Secret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["defaultMode"] = o.DefaultMode
	if !utils.IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	toSerialize["mountOnly"] = o.MountOnly
	toSerialize["mountPath"] = o.MountPath
	if !utils.IsNil(o.MountToEnv) {
		toSerialize["mountToEnv"] = o.MountToEnv
	}
	if !utils.IsNil(o.MountToEnvs) {
		toSerialize["mountToEnvs"] = o.MountToEnvs
	}
	toSerialize["name"] = o.Name
	toSerialize["readOnly"] = o.ReadOnly
	if !utils.IsNil(o.StringData) {
		toSerialize["stringData"] = o.StringData
	}
	if !utils.IsNil(o.SubPath) {
		toSerialize["subPath"] = o.SubPath
	}
	return toSerialize, nil
}

type NullableSecret struct {
	value *Secret
	isSet bool
}

func (v *NullableSecret) Get() *Secret {
	return v.value
}

func (v *NullableSecret) Set(val *Secret) {
	v.value = val
	v.isSet = true
}

func (v *NullableSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecret(val *Secret) *NullableSecret {
	return &NullableSecret{value: val, isSet: true}
}

func (v NullableSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
