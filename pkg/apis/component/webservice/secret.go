/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webservice

import (
	"encoding/json"
	"errors"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Secret type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Secret{}

// Secret struct for Secret
type Secret struct {
	DefaultMode *int32  `json:"defaultMode"`
	Items       []Items `json:"items,omitempty"`
	MountPath   *string `json:"mountPath"`
	Name        *string `json:"name"`
	SecretName  *string `json:"secretName"`
	SubPath     *string `json:"subPath,omitempty"`
}

// NewSecretWith instantiates a new Secret object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewSecretWith(defaultMode int32, mountPath string, name string, secretName string) *Secret {
	this := Secret{}
	this.DefaultMode = &defaultMode
	this.MountPath = &mountPath
	this.Name = &name
	this.SecretName = &secretName
	return &this
}

// NewSecretWithDefault instantiates a new Secret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretWithDefault() *Secret {
	this := Secret{}
	var defaultMode int32 = 420
	this.DefaultMode = &defaultMode
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
	if o.MountPath == nil {
		return errors.New("MountPath in Secret must be set")
	}
	if o.Name == nil {
		return errors.New("Name in Secret must be set")
	}
	if o.SecretName == nil {
		return errors.New("SecretName in Secret must be set")
	}
	// validate all nested properties
	return nil
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

// GetSecretName returns the SecretName field value
func (o *Secret) GetSecretName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.SecretName
}

// GetSecretNameOk returns a tuple with the SecretName field value
// and a boolean to check if the value has been set.
func (o *Secret) GetSecretNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretName, true
}

// SetSecretName sets field value
func (o *Secret) SetSecretName(v string) *Secret {
	o.SecretName = &v
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
	toSerialize["defaultMode"] = o.DefaultMode
	if !utils.IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	toSerialize["mountPath"] = o.MountPath
	toSerialize["name"] = o.Name
	toSerialize["secretName"] = o.SecretName
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
