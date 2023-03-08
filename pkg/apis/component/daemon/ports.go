/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package daemon

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

// checks if the Ports type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Ports{}

// Ports struct for Ports
type Ports struct {
	// Specify if the port should be exposed
	Expose *bool `json:"expose"`
	// Name of the port
	Name *string `json:"name,omitempty"`
	// Number of port to expose on the pod's IP address
	Port *int32 `json:"port"`
	// Protocol for port. Must be UDP, TCP, or SCTP
	Protocol *string `json:"protocol"`
}

// NewPortsWith instantiates a new Ports object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewPortsWith(expose bool, port int32, protocol string) *Ports {
	this := Ports{}
	this.Expose = &expose
	this.Port = &port
	this.Protocol = &protocol
	return &this
}

// NewPortsWithDefault instantiates a new Ports object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortsWithDefault() *Ports {
	this := Ports{}
	var expose bool = false
	this.Expose = &expose
	var protocol string = "TCP"
	this.Protocol = &protocol
	return &this
}

// NewPorts is short for NewPortsWithDefault which instantiates a new Ports object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPorts() *Ports {
	return NewPortsWithDefault()
}

// NewPortsEmpty instantiates a new Ports object with no properties set.
// This constructor will not assign any default values to properties.
func NewPortsEmpty() *Ports {
	this := Ports{}
	return &this
}

// NewPortss converts a list Ports pointers to objects.
// This is helpful when the SetPorts requires a list of objects
func NewPortsList(ps ...*Ports) []Ports {
	objs := []Ports{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Ports
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Ports) Validate() error {
	if o.Expose == nil {
		return errors.New("Expose in Ports must be set")
	}
	if o.Port == nil {
		return errors.New("Port in Ports must be set")
	}
	if o.Protocol == nil {
		return errors.New("Protocol in Ports must be set")
	}
	// validate all nested properties
	return nil
}

// GetExpose returns the Expose field value
func (o *Ports) GetExpose() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return *o.Expose
}

// GetExposeOk returns a tuple with the Expose field value
// and a boolean to check if the value has been set.
func (o *Ports) GetExposeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expose, true
}

// SetExpose sets field value
func (o *Ports) SetExpose(v bool) *Ports {
	o.Expose = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Ports) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ports) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ports) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:  Name of the port
func (o *Ports) SetName(v string) *Ports {
	o.Name = &v
	return o
}

// GetPort returns the Port field value
func (o *Ports) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *Ports) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port, true
}

// SetPort sets field value
func (o *Ports) SetPort(v int32) *Ports {
	o.Port = &v
	return o
}

// GetProtocol returns the Protocol field value
func (o *Ports) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *Ports) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Protocol, true
}

// SetProtocol sets field value
func (o *Ports) SetProtocol(v string) *Ports {
	o.Protocol = &v
	return o
}

func (o Ports) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ports) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expose"] = o.Expose
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["port"] = o.Port
	toSerialize["protocol"] = o.Protocol
	return toSerialize, nil
}

type NullablePorts struct {
	value *Ports
	isSet bool
}

func (v *NullablePorts) Get() *Ports {
	return v.value
}

func (v *NullablePorts) Set(val *Ports) {
	v.value = val
	v.isSet = true
}

func (v *NullablePorts) IsSet() bool {
	return v.isSet
}

func (v *NullablePorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePorts(val *Ports) *NullablePorts {
	return &NullablePorts{value: val, isSet: true}
}

func (v NullablePorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
