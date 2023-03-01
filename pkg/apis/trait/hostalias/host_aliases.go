/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hostalias

import (
	"encoding/json"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the HostAliases type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HostAliases{}

// HostAliases struct for HostAliases
type HostAliases struct {
	Hostnames []string `json:"hostnames,omitempty"`
	Ip        *string  `json:"ip,omitempty"`
}

// NewHostAliasesWith instantiates a new HostAliases object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostAliasesWith() *HostAliases {
	this := HostAliases{}
	return &this
}

// NewHostAliases instantiates a new HostAliases object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostAliases() *HostAliases {
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

// GetHostnames returns the Hostnames field value if set, zero value otherwise.
func (o *HostAliases) GetHostnames() []string {
	if o == nil || utils.IsNil(o.Hostnames) {
		var ret []string
		return ret
	}
	return o.Hostnames
}

// GetHostnamesOk returns a tuple with the Hostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostAliases) GetHostnamesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Hostnames) {
		return nil, false
	}
	return o.Hostnames, true
}

// HasHostnames returns a boolean if a field has been set.
func (o *HostAliases) HasHostnames() bool {
	if o != nil && !utils.IsNil(o.Hostnames) {
		return true
	}

	return false
}

// SetHostnames gets a reference to the given []string and assigns it to the hostnames field.
// Hostnames:
func (o *HostAliases) SetHostnames(v []string) *HostAliases {
	o.Hostnames = v
	return o
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *HostAliases) GetIp() string {
	if o == nil || utils.IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostAliases) GetIpOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *HostAliases) HasIp() bool {
	if o != nil && !utils.IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the ip field.
// Ip:
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
	if !utils.IsNil(o.Hostnames) {
		toSerialize["hostnames"] = o.Hostnames
	}
	if !utils.IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	return toSerialize, nil
}

type NullableHostAliases struct {
	value *HostAliases
	isSet bool
}

func (v NullableHostAliases) Get() *HostAliases {
	return v.value
}

func (v *NullableHostAliases) Set(val *HostAliases) {
	v.value = val
	v.isSet = true
}

func (v NullableHostAliases) IsSet() bool {
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
