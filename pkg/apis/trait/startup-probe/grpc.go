/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package startup_probe

import (
	"encoding/json"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Grpc type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Grpc{}

// Grpc Instructions for assessing container startup status by probing a gRPC service. Either this attribute or the exec attribute or the grpc attribute or the httpGet attribute MUST be specified. This attribute is mutually exclusive with the exec attribute and the httpGet attribute and the tcpSocket attribute.
type Grpc struct {
	// The port number of the gRPC service.
	Port *int32 `json:"port,omitempty"`
	// The name of the service to place in the gRPC HealthCheckRequest
	Service *string `json:"service,omitempty"`
}

// NewGrpcWith instantiates a new Grpc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrpcWith() *Grpc {
	this := Grpc{}
	return &this
}

// NewGrpc instantiates a new Grpc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrpc() *Grpc {
	this := Grpc{}
	return &this
}

// NewGrpcs converts a list Grpc pointers to objects.
// This is helpful when the SetGrpc requires a list of objects
func NewGrpcList(ps ...*Grpc) []Grpc {
	objs := []Grpc{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *Grpc) GetPort() int32 {
	if o == nil || utils.IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grpc) GetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *Grpc) HasPort() bool {
	if o != nil && !utils.IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the port field.
// Port:  The port number of the gRPC service.
func (o *Grpc) SetPort(v int32) *Grpc {
	o.Port = &v
	return o
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *Grpc) GetService() string {
	if o == nil || utils.IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grpc) GetServiceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *Grpc) HasService() bool {
	if o != nil && !utils.IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the service field.
// Service:  The name of the service to place in the gRPC HealthCheckRequest
func (o *Grpc) SetService(v string) *Grpc {
	o.Service = &v
	return o
}

func (o Grpc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Grpc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !utils.IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	return toSerialize, nil
}

type NullableGrpc struct {
	value *Grpc
	isSet bool
}

func (v NullableGrpc) Get() *Grpc {
	return v.value
}

func (v *NullableGrpc) Set(val *Grpc) {
	v.value = val
	v.isSet = true
}

func (v NullableGrpc) IsSet() bool {
	return v.isSet
}

func (v *NullableGrpc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrpc(val *Grpc) *NullableGrpc {
	return &NullableGrpc{value: val, isSet: true}
}

func (v NullableGrpc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrpc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
