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

// checks if the Content type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Content{}

// Content Specify the content of the email
type Content struct {
	// Specify the context body of the email
	Body *string `json:"body,omitempty"`
	// Specify the subject of the email
	Subject *string `json:"subject,omitempty"`
}

// NewContentWith instantiates a new Content object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentWith() *Content {
	this := Content{}
	return &this
}

// NewContent instantiates a new Content object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContent() *Content {
	this := Content{}
	return &this
}

// NewContents converts a list Content pointers to objects.
// This is helpful when the SetContent requires a list of objects
func NewContentList(ps ...*Content) []Content {
	objs := []Content{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *Content) GetBody() string {
	if o == nil || utils.IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetBodyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *Content) HasBody() bool {
	if o != nil && !utils.IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the body field.
// Body:  Specify the context body of the email
func (o *Content) SetBody(v string) *Content {
	o.Body = &v
	return o
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *Content) GetSubject() string {
	if o == nil || utils.IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetSubjectOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *Content) HasSubject() bool {
	if o != nil && !utils.IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the subject field.
// Subject:  Specify the subject of the email
func (o *Content) SetSubject(v string) *Content {
	o.Subject = &v
	return o
}

func (o Content) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Content) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !utils.IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	return toSerialize, nil
}

type NullableContent struct {
	value *Content
	isSet bool
}

func (v NullableContent) Get() *Content {
	return v.value
}

func (v *NullableContent) Set(val *Content) {
	v.value = val
	v.isSet = true
}

func (v NullableContent) IsSet() bool {
	return v.isSet
}

func (v *NullableContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContent(val *Content) *NullableContent {
	return &NullableContent{value: val, isSet: true}
}

func (v NullableContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
