/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package startup_probe

import (
	"encoding/json"
	"errors"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the HttpHeaders type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HttpHeaders{}

// HttpHeaders struct for HttpHeaders
type HttpHeaders struct {
	// The header field name
	Name *string `json:"name"`
	// The header field value
	Value *string `json:"value"`
}

// NewHttpHeadersWith instantiates a new HttpHeaders object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewHttpHeadersWith(name string, value string) *HttpHeaders {
	this := HttpHeaders{}
	this.Name = &name
	this.Value = &value
	return &this
}

// NewHttpHeadersWithDefault instantiates a new HttpHeaders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpHeadersWithDefault() *HttpHeaders {
	this := HttpHeaders{}
	return &this
}

// NewHttpHeaders is short for NewHttpHeadersWithDefault which instantiates a new HttpHeaders object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpHeaders() *HttpHeaders {
	return NewHttpHeadersWithDefault()
}

// NewHttpHeadersEmpty instantiates a new HttpHeaders object with no properties set.
// This constructor will not assign any default values to properties.
func NewHttpHeadersEmpty() *HttpHeaders {
	this := HttpHeaders{}
	return &this
}

// NewHttpHeaderss converts a list HttpHeaders pointers to objects.
// This is helpful when the SetHttpHeaders requires a list of objects
func NewHttpHeadersList(ps ...*HttpHeaders) []HttpHeaders {
	objs := []HttpHeaders{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this HttpHeaders
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *HttpHeaders) Validate() error {
	if o.Name == nil {
		return errors.New("Name in HttpHeaders must be set")
	}
	if o.Value == nil {
		return errors.New("Value in HttpHeaders must be set")
	}
	// validate all nested properties
	return nil
}

// GetName returns the Name field value
func (o *HttpHeaders) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HttpHeaders) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *HttpHeaders) SetName(v string) *HttpHeaders {
	o.Name = &v
	return o
}

// GetValue returns the Value field value
func (o *HttpHeaders) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *HttpHeaders) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *HttpHeaders) SetValue(v string) *HttpHeaders {
	o.Value = &v
	return o
}

func (o HttpHeaders) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpHeaders) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableHttpHeaders struct {
	value *HttpHeaders
	isSet bool
}

func (v *NullableHttpHeaders) Get() *HttpHeaders {
	return v.value
}

func (v *NullableHttpHeaders) Set(val *HttpHeaders) {
	v.value = val
	v.isSet = true
}

func (v *NullableHttpHeaders) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpHeaders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpHeaders(val *HttpHeaders) *NullableHttpHeaders {
	return &NullableHttpHeaders{value: val, isSet: true}
}

func (v NullableHttpHeaders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpHeaders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
