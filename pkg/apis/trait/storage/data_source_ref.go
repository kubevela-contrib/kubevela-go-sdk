/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"

	"github.com/kubevela-contrib/vela-go-sdk/pkg/apis/utils"
)

// checks if the DataSourceRef type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DataSourceRef{}

// DataSourceRef struct for DataSourceRef
type DataSourceRef struct {
	ApiGroup *string `json:"apiGroup,omitempty"`
	Kind     *string `json:"kind,omitempty"`
	Name     *string `json:"name,omitempty"`
}

// NewDataSourceRefWith instantiates a new DataSourceRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSourceRefWith() *DataSourceRef {
	this := DataSourceRef{}
	return &this
}

// NewDataSourceRef instantiates a new DataSourceRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourceRef() *DataSourceRef {
	this := DataSourceRef{}
	return &this
}

// NewDataSourceRefs converts a list DataSourceRef pointers to objects.
// This is helpful when the SetDataSourceRef requires a list of objects
func NewDataSourceRefList(ps ...*DataSourceRef) []DataSourceRef {
	objs := []DataSourceRef{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetApiGroup returns the ApiGroup field value if set, zero value otherwise.
func (o *DataSourceRef) GetApiGroup() string {
	if o == nil || utils.IsNil(o.ApiGroup) {
		var ret string
		return ret
	}
	return *o.ApiGroup
}

// GetApiGroupOk returns a tuple with the ApiGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceRef) GetApiGroupOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ApiGroup) {
		return nil, false
	}
	return o.ApiGroup, true
}

// HasApiGroup returns a boolean if a field has been set.
func (o *DataSourceRef) HasApiGroup() bool {
	if o != nil && !utils.IsNil(o.ApiGroup) {
		return true
	}

	return false
}

// SetApiGroup gets a reference to the given string and assigns it to the apiGroup field.
// ApiGroup:
func (o *DataSourceRef) SetApiGroup(v string) *DataSourceRef {
	o.ApiGroup = &v
	return o
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *DataSourceRef) GetKind() string {
	if o == nil || utils.IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceRef) GetKindOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *DataSourceRef) HasKind() bool {
	if o != nil && !utils.IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the kind field.
// Kind:
func (o *DataSourceRef) SetKind(v string) *DataSourceRef {
	o.Kind = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataSourceRef) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceRef) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataSourceRef) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:
func (o *DataSourceRef) SetName(v string) *DataSourceRef {
	o.Name = &v
	return o
}

func (o DataSourceRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSourceRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ApiGroup) {
		toSerialize["apiGroup"] = o.ApiGroup
	}
	if !utils.IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableDataSourceRef struct {
	value *DataSourceRef
	isSet bool
}

func (v NullableDataSourceRef) Get() *DataSourceRef {
	return v.value
}

func (v *NullableDataSourceRef) Set(val *DataSourceRef) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceRef) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceRef(val *DataSourceRef) *NullableDataSourceRef {
	return &NullableDataSourceRef{value: val, isSet: true}
}

func (v NullableDataSourceRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
