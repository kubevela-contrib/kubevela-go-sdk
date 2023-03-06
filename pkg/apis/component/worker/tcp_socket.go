/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package worker

import (
	"encoding/json"
	"errors"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the TcpSocket type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TcpSocket{}

// TcpSocket Instructions for assessing container health by probing a TCP socket. Either this attribute or the exec attribute or the httpGet attribute MUST be specified. This attribute is mutually exclusive with both the exec attribute and the httpGet attribute.
type TcpSocket struct {
	// The TCP socket within the container that should be probed to assess container health.
	Port *int32 `json:"port"`
}

// NewTcpSocketWith instantiates a new TcpSocket object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewTcpSocketWith(port int32) *TcpSocket {
	this := TcpSocket{}
	this.Port = &port
	return &this
}

// NewTcpSocketWithDefault instantiates a new TcpSocket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTcpSocketWithDefault() *TcpSocket {
	this := TcpSocket{}
	return &this
}

// NewTcpSocket is short for NewTcpSocketWithDefault which instantiates a new TcpSocket object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTcpSocket() *TcpSocket {
	return NewTcpSocketWithDefault()
}

// NewTcpSocketEmpty instantiates a new TcpSocket object with no properties set.
// This constructor will not assign any default values to properties.
func NewTcpSocketEmpty() *TcpSocket {
	this := TcpSocket{}
	return &this
}

// NewTcpSockets converts a list TcpSocket pointers to objects.
// This is helpful when the SetTcpSocket requires a list of objects
func NewTcpSocketList(ps ...*TcpSocket) []TcpSocket {
	objs := []TcpSocket{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this TcpSocket
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *TcpSocket) Validate() error {
	if o.Port == nil {
		return errors.New("Port in TcpSocket must be set")
	}
	// validate all nested properties
	return nil
}

// GetPort returns the Port field value
func (o *TcpSocket) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *TcpSocket) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port, true
}

// SetPort sets field value
func (o *TcpSocket) SetPort(v int32) *TcpSocket {
	o.Port = &v
	return o
}

func (o TcpSocket) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TcpSocket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["port"] = o.Port
	return toSerialize, nil
}

type NullableTcpSocket struct {
	value *TcpSocket
	isSet bool
}

func (v *NullableTcpSocket) Get() *TcpSocket {
	return v.value
}

func (v *NullableTcpSocket) Set(val *TcpSocket) {
	v.value = val
	v.isSet = true
}

func (v *NullableTcpSocket) IsSet() bool {
	return v.isSet
}

func (v *NullableTcpSocket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTcpSocket(val *TcpSocket) *NullableTcpSocket {
	return &NullableTcpSocket{value: val, isSet: true}
}

func (v NullableTcpSocket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTcpSocket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
