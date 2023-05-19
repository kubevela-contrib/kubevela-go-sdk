/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resource_update

import (
	"encoding/json"
	"errors"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Strategy type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Strategy{}

// Strategy struct for Strategy
type Strategy struct {
	// Specify the op for updating target resources
	Op *string `json:"op"`
	// Specify which fields would trigger recreation when updated
	RecreateFields []string `json:"recreateFields,omitempty"`
}

// NewStrategyWith instantiates a new Strategy object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewStrategyWith(op string) *Strategy {
	this := Strategy{}
	this.Op = &op
	return &this
}

// NewStrategyWithDefault instantiates a new Strategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStrategyWithDefault() *Strategy {
	this := Strategy{}
	var op string = "patch"
	this.Op = &op
	return &this
}

// NewStrategy is short for NewStrategyWithDefault which instantiates a new Strategy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStrategy() *Strategy {
	return NewStrategyWithDefault()
}

// NewStrategyEmpty instantiates a new Strategy object with no properties set.
// This constructor will not assign any default values to properties.
func NewStrategyEmpty() *Strategy {
	this := Strategy{}
	return &this
}

// NewStrategys converts a list Strategy pointers to objects.
// This is helpful when the SetStrategy requires a list of objects
func NewStrategyList(ps ...*Strategy) []Strategy {
	objs := []Strategy{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Strategy
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Strategy) Validate() error {
	if o.Op == nil {
		return errors.New("Op in Strategy must be set")
	}
	// validate all nested properties
	return nil
}

// GetOp returns the Op field value
func (o *Strategy) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *Strategy) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Op, true
}

// SetOp sets field value
func (o *Strategy) SetOp(v string) *Strategy {
	o.Op = &v
	return o
}

// GetRecreateFields returns the RecreateFields field value if set, zero value otherwise.
func (o *Strategy) GetRecreateFields() []string {
	if o == nil || utils.IsNil(o.RecreateFields) {
		var ret []string
		return ret
	}
	return o.RecreateFields
}

// GetRecreateFieldsOk returns a tuple with the RecreateFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Strategy) GetRecreateFieldsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.RecreateFields) {
		return nil, false
	}
	return o.RecreateFields, true
}

// HasRecreateFields returns a boolean if a field has been set.
func (o *Strategy) HasRecreateFields() bool {
	if o != nil && !utils.IsNil(o.RecreateFields) {
		return true
	}

	return false
}

// SetRecreateFields gets a reference to the given []string and assigns it to the recreateFields field.
// RecreateFields:  Specify which fields would trigger recreation when updated
func (o *Strategy) SetRecreateFields(v []string) *Strategy {
	o.RecreateFields = v
	return o
}

func (o Strategy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Strategy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["op"] = o.Op
	if !utils.IsNil(o.RecreateFields) {
		toSerialize["recreateFields"] = o.RecreateFields
	}
	return toSerialize, nil
}

type NullableStrategy struct {
	value *Strategy
	isSet bool
}

func (v *NullableStrategy) Get() *Strategy {
	return v.value
}

func (v *NullableStrategy) Set(val *Strategy) {
	v.value = val
	v.isSet = true
}

func (v *NullableStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStrategy(val *Strategy) *NullableStrategy {
	return &NullableStrategy{value: val, isSet: true}
}

func (v NullableStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
