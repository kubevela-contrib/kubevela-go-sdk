/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nocalhost

import (
	"encoding/json"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Sync type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Sync{}

// Sync struct for Sync
type Sync struct {
	FilePattern       []string `json:"filePattern,omitempty"`
	IgnoreFilePattern []string `json:"ignoreFilePattern,omitempty"`
	Type              *string  `json:"type,omitempty"`
}

// NewSyncWith instantiates a new Sync object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncWith() *Sync {
	this := Sync{}
	var type_ string = "send"
	this.Type = &type_
	return &this
}

// NewSync instantiates a new Sync object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSync() *Sync {
	this := Sync{}
	var type_ string = "send"
	this.Type = &type_
	return &this
}

// NewSyncs converts a list Sync pointers to objects.
// This is helpful when the SetSync requires a list of objects
func NewSyncList(ps ...*Sync) []Sync {
	objs := []Sync{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetFilePattern returns the FilePattern field value if set, zero value otherwise.
func (o *Sync) GetFilePattern() []string {
	if o == nil || utils.IsNil(o.FilePattern) {
		var ret []string
		return ret
	}
	return o.FilePattern
}

// GetFilePatternOk returns a tuple with the FilePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sync) GetFilePatternOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.FilePattern) {
		return nil, false
	}
	return o.FilePattern, true
}

// HasFilePattern returns a boolean if a field has been set.
func (o *Sync) HasFilePattern() bool {
	if o != nil && !utils.IsNil(o.FilePattern) {
		return true
	}

	return false
}

// SetFilePattern gets a reference to the given []string and assigns it to the filePattern field.
// FilePattern:
func (o *Sync) SetFilePattern(v []string) *Sync {
	o.FilePattern = v
	return o
}

// GetIgnoreFilePattern returns the IgnoreFilePattern field value if set, zero value otherwise.
func (o *Sync) GetIgnoreFilePattern() []string {
	if o == nil || utils.IsNil(o.IgnoreFilePattern) {
		var ret []string
		return ret
	}
	return o.IgnoreFilePattern
}

// GetIgnoreFilePatternOk returns a tuple with the IgnoreFilePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sync) GetIgnoreFilePatternOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.IgnoreFilePattern) {
		return nil, false
	}
	return o.IgnoreFilePattern, true
}

// HasIgnoreFilePattern returns a boolean if a field has been set.
func (o *Sync) HasIgnoreFilePattern() bool {
	if o != nil && !utils.IsNil(o.IgnoreFilePattern) {
		return true
	}

	return false
}

// SetIgnoreFilePattern gets a reference to the given []string and assigns it to the ignoreFilePattern field.
// IgnoreFilePattern:
func (o *Sync) SetIgnoreFilePattern(v []string) *Sync {
	o.IgnoreFilePattern = v
	return o
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Sync) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sync) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Sync) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the type_ field.
// Type:
func (o *Sync) SetType(v string) *Sync {
	o.Type = &v
	return o
}

func (o Sync) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sync) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.FilePattern) {
		toSerialize["filePattern"] = o.FilePattern
	}
	if !utils.IsNil(o.IgnoreFilePattern) {
		toSerialize["ignoreFilePattern"] = o.IgnoreFilePattern
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableSync struct {
	value *Sync
	isSet bool
}

func (v NullableSync) Get() *Sync {
	return v.value
}

func (v *NullableSync) Set(val *Sync) {
	v.value = val
	v.isSet = true
}

func (v NullableSync) IsSet() bool {
	return v.isSet
}

func (v *NullableSync) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSync(val *Sync) *NullableSync {
	return &NullableSync{value: val, isSet: true}
}

func (v NullableSync) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSync) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
