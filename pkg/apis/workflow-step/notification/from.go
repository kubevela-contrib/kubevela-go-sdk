/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the From type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &From{}

// From Specify the email info that you want to send from
type From struct {
	// Specify the email address that you want to send from
	Address *string `json:"address,omitempty"`
	// The alias is the email alias to show after sending the email
	Alias *string `json:"alias,omitempty"`
	// Specify the host of your email
	Host     *string   `json:"host,omitempty"`
	Password *Password `json:"password,omitempty"`
	// Specify the port of the email host, default to 587
	Port *int32 `json:"port,omitempty"`
}

// NewFromWith instantiates a new From object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFromWith() *From {
	this := From{}
	var port int32 = 587
	this.Port = &port
	return &this
}

// NewFrom instantiates a new From object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrom() *From {
	this := From{}
	var port int32 = 587
	this.Port = &port
	return &this
}

// NewFroms converts a list From pointers to objects.
// This is helpful when the SetFrom requires a list of objects
func NewFromList(ps ...*From) []From {
	objs := []From{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *From) GetAddress() string {
	if o == nil || utils.IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *From) GetAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *From) HasAddress() bool {
	if o != nil && !utils.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the address field.
// Address:  Specify the email address that you want to send from
func (o *From) SetAddress(v string) *From {
	o.Address = &v
	return o
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *From) GetAlias() string {
	if o == nil || utils.IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *From) GetAliasOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *From) HasAlias() bool {
	if o != nil && !utils.IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the alias field.
// Alias:  The alias is the email alias to show after sending the email
func (o *From) SetAlias(v string) *From {
	o.Alias = &v
	return o
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *From) GetHost() string {
	if o == nil || utils.IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *From) GetHostOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *From) HasHost() bool {
	if o != nil && !utils.IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the host field.
// Host:  Specify the host of your email
func (o *From) SetHost(v string) *From {
	o.Host = &v
	return o
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *From) GetPassword() Password {
	if o == nil || utils.IsNil(o.Password) {
		var ret Password
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *From) GetPasswordOk() (*Password, bool) {
	if o == nil || utils.IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *From) HasPassword() bool {
	if o != nil && !utils.IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given Password and assigns it to the password field.
// Password:
func (o *From) SetPassword(v Password) *From {
	o.Password = &v
	return o
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *From) GetPort() int32 {
	if o == nil || utils.IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *From) GetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *From) HasPort() bool {
	if o != nil && !utils.IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the port field.
// Port:  Specify the port of the email host, default to 587
func (o *From) SetPort(v int32) *From {
	o.Port = &v
	return o
}

func (o From) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o From) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !utils.IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !utils.IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !utils.IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !utils.IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

type NullableFrom struct {
	value *From
	isSet bool
}

func (v NullableFrom) Get() *From {
	return v.value
}

func (v *NullableFrom) Set(val *From) {
	v.value = val
	v.isSet = true
}

func (v NullableFrom) IsSet() bool {
	return v.isSet
}

func (v *NullableFrom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrom(val *From) *NullableFrom {
	return &NullableFrom{value: val, isSet: true}
}

func (v NullableFrom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
