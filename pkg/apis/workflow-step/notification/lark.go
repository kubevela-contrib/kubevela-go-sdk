/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

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

// checks if the Lark type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Lark{}

// Lark Please fulfill its url and message if you want to send Lark messages
type Lark struct {
	Message *Message1 `json:"message"`
	Url     *Url1     `json:"url"`
}

// NewLarkWith instantiates a new Lark object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewLarkWith(message Message1, url Url1) *Lark {
	this := Lark{}
	this.Message = &message
	this.Url = &url
	return &this
}

// NewLarkWithDefault instantiates a new Lark object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLarkWithDefault() *Lark {
	this := Lark{}
	return &this
}

// NewLark is short for NewLarkWithDefault which instantiates a new Lark object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLark() *Lark {
	return NewLarkWithDefault()
}

// NewLarkEmpty instantiates a new Lark object with no properties set.
// This constructor will not assign any default values to properties.
func NewLarkEmpty() *Lark {
	this := Lark{}
	return &this
}

// NewLarks converts a list Lark pointers to objects.
// This is helpful when the SetLark requires a list of objects
func NewLarkList(ps ...*Lark) []Lark {
	objs := []Lark{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Lark
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Lark) Validate() error {
	if o.Message == nil {
		return errors.New("Message in Lark must be set")
	}
	if o.Url == nil {
		return errors.New("Url in Lark must be set")
	}
	// validate all nested properties
	if o.Message != nil {
		if err := o.Message.Validate(); err != nil {
			return err
		}
	}
	if o.Url != nil {
		if err := o.Url.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// GetMessage returns the Message field value
func (o *Lark) GetMessage() Message1 {
	if o == nil {
		var ret Message1
		return ret
	}

	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Lark) GetMessageOk() (*Message1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message, true
}

// SetMessage sets field value
func (o *Lark) SetMessage(v Message1) *Lark {
	o.Message = &v
	return o
}

// GetUrl returns the Url field value
func (o *Lark) GetUrl() Url1 {
	if o == nil {
		var ret Url1
		return ret
	}

	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Lark) GetUrlOk() (*Url1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url, true
}

// SetUrl sets field value
func (o *Lark) SetUrl(v Url1) *Lark {
	o.Url = &v
	return o
}

func (o Lark) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Lark) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableLark struct {
	value *Lark
	isSet bool
}

func (v *NullableLark) Get() *Lark {
	return v.value
}

func (v *NullableLark) Set(val *Lark) {
	v.value = val
	v.isSet = true
}

func (v *NullableLark) IsSet() bool {
	return v.isSet
}

func (v *NullableLark) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLark(val *Lark) *NullableLark {
	return &NullableLark{value: val, isSet: true}
}

func (v NullableLark) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLark) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
