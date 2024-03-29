/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package container_ports

import (
	"encoding/json"
	"errors"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Ports type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Ports{}

// Ports struct for Ports
type Ports struct {
	// Number of port to expose on the pod's IP address
	ContainerPort *int32 `json:"containerPort"`
	// What host IP to bind the external port to.
	HostIP *string `json:"hostIP,omitempty"`
	// Number of port to expose on the host
	HostPort *int32 `json:"hostPort,omitempty"`
	// Protocol for port. Must be UDP, TCP, or SCTP
	Protocol *string `json:"protocol"`
}

// NewPortsWith instantiates a new Ports object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewPortsWith(containerPort int32, protocol string) *Ports {
	this := Ports{}
	this.ContainerPort = &containerPort
	this.Protocol = &protocol
	return &this
}

// NewPortsWithDefault instantiates a new Ports object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortsWithDefault() *Ports {
	this := Ports{}
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
	if o.ContainerPort == nil {
		return errors.New("ContainerPort in Ports must be set")
	}
	if o.Protocol == nil {
		return errors.New("Protocol in Ports must be set")
	}
	// validate all nested properties
	return nil
}

// GetContainerPort returns the ContainerPort field value
func (o *Ports) GetContainerPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.ContainerPort
}

// GetContainerPortOk returns a tuple with the ContainerPort field value
// and a boolean to check if the value has been set.
func (o *Ports) GetContainerPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerPort, true
}

// SetContainerPort sets field value
func (o *Ports) SetContainerPort(v int32) *Ports {
	o.ContainerPort = &v
	return o
}

// GetHostIP returns the HostIP field value if set, zero value otherwise.
func (o *Ports) GetHostIP() string {
	if o == nil || utils.IsNil(o.HostIP) {
		var ret string
		return ret
	}
	return *o.HostIP
}

// GetHostIPOk returns a tuple with the HostIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ports) GetHostIPOk() (*string, bool) {
	if o == nil || utils.IsNil(o.HostIP) {
		return nil, false
	}
	return o.HostIP, true
}

// HasHostIP returns a boolean if a field has been set.
func (o *Ports) HasHostIP() bool {
	if o != nil && !utils.IsNil(o.HostIP) {
		return true
	}

	return false
}

// SetHostIP gets a reference to the given string and assigns it to the hostIP field.
// HostIP:  What host IP to bind the external port to.
func (o *Ports) SetHostIP(v string) *Ports {
	o.HostIP = &v
	return o
}

// GetHostPort returns the HostPort field value if set, zero value otherwise.
func (o *Ports) GetHostPort() int32 {
	if o == nil || utils.IsNil(o.HostPort) {
		var ret int32
		return ret
	}
	return *o.HostPort
}

// GetHostPortOk returns a tuple with the HostPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ports) GetHostPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.HostPort) {
		return nil, false
	}
	return o.HostPort, true
}

// HasHostPort returns a boolean if a field has been set.
func (o *Ports) HasHostPort() bool {
	if o != nil && !utils.IsNil(o.HostPort) {
		return true
	}

	return false
}

// SetHostPort gets a reference to the given int32 and assigns it to the hostPort field.
// HostPort:  Number of port to expose on the host
func (o *Ports) SetHostPort(v int32) *Ports {
	o.HostPort = &v
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
	toSerialize["containerPort"] = o.ContainerPort
	if !utils.IsNil(o.HostIP) {
		toSerialize["hostIP"] = o.HostIP
	}
	if !utils.IsNil(o.HostPort) {
		toSerialize["hostPort"] = o.HostPort
	}
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
