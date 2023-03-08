/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package build_push_image

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

// checks if the Credentials type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Credentials{}

// Credentials Specify the credentials to access git and image registry
type Credentials struct {
	Git   *Git1  `json:"git,omitempty"`
	Image *Image `json:"image,omitempty"`
}

// NewCredentialsWith instantiates a new Credentials object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewCredentialsWith() *Credentials {
	this := Credentials{}
	return &this
}

// NewCredentialsWithDefault instantiates a new Credentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsWithDefault() *Credentials {
	this := Credentials{}
	return &this
}

// NewCredentials is short for NewCredentialsWithDefault which instantiates a new Credentials object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentials() *Credentials {
	return NewCredentialsWithDefault()
}

// NewCredentialsEmpty instantiates a new Credentials object with no properties set.
// This constructor will not assign any default values to properties.
func NewCredentialsEmpty() *Credentials {
	this := Credentials{}
	return &this
}

// NewCredentialss converts a list Credentials pointers to objects.
// This is helpful when the SetCredentials requires a list of objects
func NewCredentialsList(ps ...*Credentials) []Credentials {
	objs := []Credentials{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Credentials
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Credentials) Validate() error {
	// validate all nested properties
	if o.Git != nil {
		if err := o.Git.Validate(); err != nil {
			return err
		}
	}
	if o.Image != nil {
		if err := o.Image.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *Credentials) GetGit() Git1 {
	if o == nil || utils.IsNil(o.Git) {
		var ret Git1
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credentials) GetGitOk() (*Git1, bool) {
	if o == nil || utils.IsNil(o.Git) {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *Credentials) HasGit() bool {
	if o != nil && !utils.IsNil(o.Git) {
		return true
	}

	return false
}

// SetGit gets a reference to the given Git1 and assigns it to the git field.
// Git:
func (o *Credentials) SetGit(v Git1) *Credentials {
	o.Git = &v
	return o
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *Credentials) GetImage() Image {
	if o == nil || utils.IsNil(o.Image) {
		var ret Image
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credentials) GetImageOk() (*Image, bool) {
	if o == nil || utils.IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *Credentials) HasImage() bool {
	if o != nil && !utils.IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given Image and assigns it to the image field.
// Image:
func (o *Credentials) SetImage(v Image) *Credentials {
	o.Image = &v
	return o
}

func (o Credentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Credentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Git) {
		toSerialize["git"] = o.Git
	}
	if !utils.IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	return toSerialize, nil
}

type NullableCredentials struct {
	value *Credentials
	isSet bool
}

func (v *NullableCredentials) Get() *Credentials {
	return v.value
}

func (v *NullableCredentials) Set(val *Credentials) {
	v.value = val
	v.isSet = true
}

func (v *NullableCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentials(val *Credentials) *NullableCredentials {
	return &NullableCredentials{value: val, isSet: true}
}

func (v NullableCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
