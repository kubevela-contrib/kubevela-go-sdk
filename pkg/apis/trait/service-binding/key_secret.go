/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package service_binding

import (
	"encoding/json"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the KeySecret type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &KeySecret{}

// KeySecret struct for KeySecret
type KeySecret struct {
	Key    *string `json:"key,omitempty"`
	Secret *string `json:"secret,omitempty"`
}

// NewKeySecretWith instantiates a new KeySecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeySecretWith() *KeySecret {
	this := KeySecret{}
	return &this
}

// NewKeySecret instantiates a new KeySecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeySecret() *KeySecret {
	this := KeySecret{}
	return &this
}

// NewKeySecrets converts a list KeySecret pointers to objects.
// This is helpful when the SetKeySecret requires a list of objects
func NewKeySecretList(ps ...*KeySecret) []KeySecret {
	objs := []KeySecret{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *KeySecret) GetKey() string {
	if o == nil || utils.IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeySecret) GetKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *KeySecret) HasKey() bool {
	if o != nil && !utils.IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the key field.
// Key:
func (o *KeySecret) SetKey(v string) *KeySecret {
	o.Key = &v
	return o
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *KeySecret) GetSecret() string {
	if o == nil || utils.IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeySecret) GetSecretOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *KeySecret) HasSecret() bool {
	if o != nil && !utils.IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the secret field.
// Secret:
func (o *KeySecret) SetSecret(v string) *KeySecret {
	o.Secret = &v
	return o
}

func (o KeySecret) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeySecret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !utils.IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return toSerialize, nil
}

type NullableKeySecret struct {
	value *KeySecret
	isSet bool
}

func (v NullableKeySecret) Get() *KeySecret {
	return v.value
}

func (v *NullableKeySecret) Set(val *KeySecret) {
	v.value = val
	v.isSet = true
}

func (v NullableKeySecret) IsSet() bool {
	return v.isSet
}

func (v *NullableKeySecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeySecret(val *KeySecret) *NullableKeySecret {
	return &NullableKeySecret{value: val, isSet: true}
}

func (v NullableKeySecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeySecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
