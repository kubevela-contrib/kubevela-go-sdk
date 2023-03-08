/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hostalias

import (
	"encoding/json"
	"errors"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the HostAliases type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HostAliases{}

// HostAliases struct for HostAliases
type HostAliases struct {
	Hostnames []string `json:"hostnames"`
	Ip        *string  `json:"ip"`
}

// NewHostAliasesWith instantiates a new HostAliases object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewHostAliasesWith(hostnames []string, ip string) *HostAliases {
	this := HostAliases{}
	this.Hostnames = hostnames
	this.Ip = &ip
	return &this
}

// NewHostAliasesWithDefault instantiates a new HostAliases object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostAliasesWithDefault() *HostAliases {
	this := HostAliases{}
	return &this
}

// NewHostAliases is short for NewHostAliasesWithDefault which instantiates a new HostAliases object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostAliases() *HostAliases {
	return NewHostAliasesWithDefault()
}

// NewHostAliasesEmpty instantiates a new HostAliases object with no properties set.
// This constructor will not assign any default values to properties.
func NewHostAliasesEmpty() *HostAliases {
	this := HostAliases{}
	return &this
}

// NewHostAliasess converts a list HostAliases pointers to objects.
// This is helpful when the SetHostAliases requires a list of objects
func NewHostAliasesList(ps ...*HostAliases) []HostAliases {
	objs := []HostAliases{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this HostAliases
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *HostAliases) Validate() error {
	if o.Hostnames == nil {
		return errors.New("Hostnames in HostAliases must be set")
	}
	if o.Ip == nil {
		return errors.New("Ip in HostAliases must be set")
	}
	// validate all nested properties
	return nil
}

// GetHostnames returns the Hostnames field value
func (o *HostAliases) GetHostnames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Hostnames
}

// GetHostnamesOk returns a tuple with the Hostnames field value
// and a boolean to check if the value has been set.
func (o *HostAliases) GetHostnamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hostnames, true
}

// SetHostnames sets field value
func (o *HostAliases) SetHostnames(v []string) *HostAliases {
	o.Hostnames = v
	return o
}

// GetIp returns the Ip field value
func (o *HostAliases) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *HostAliases) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip, true
}

// SetIp sets field value
func (o *HostAliases) SetIp(v string) *HostAliases {
	o.Ip = &v
	return o
}

func (o HostAliases) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostAliases) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hostnames"] = o.Hostnames
	toSerialize["ip"] = o.Ip
	return toSerialize, nil
}

type NullableHostAliases struct {
	value *HostAliases
	isSet bool
}

func (v *NullableHostAliases) Get() *HostAliases {
	return v.value
}

func (v *NullableHostAliases) Set(val *HostAliases) {
	v.value = val
	v.isSet = true
}

func (v *NullableHostAliases) IsSet() bool {
	return v.isSet
}

func (v *NullableHostAliases) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostAliases(val *HostAliases) *NullableHostAliases {
	return &NullableHostAliases{value: val, isSet: true}
}

func (v NullableHostAliases) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostAliases) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
