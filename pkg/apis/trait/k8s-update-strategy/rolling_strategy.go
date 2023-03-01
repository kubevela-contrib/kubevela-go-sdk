/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s_update_strategy

import (
	"encoding/json"

	"github.com/kubevela-contrib/vela-go-sdk/pkg/apis/utils"
)

// checks if the RollingStrategy type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RollingStrategy{}

// RollingStrategy Specify the parameters of rollong update strategy
type RollingStrategy struct {
	MaxSurge       *string `json:"maxSurge,omitempty"`
	MaxUnavailable *string `json:"maxUnavailable,omitempty"`
	Partition      *int32  `json:"partition,omitempty"`
}

// NewRollingStrategyWith instantiates a new RollingStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRollingStrategyWith() *RollingStrategy {
	this := RollingStrategy{}
	var maxSurge string = "25%"
	this.MaxSurge = &maxSurge
	var maxUnavailable string = "25%"
	this.MaxUnavailable = &maxUnavailable
	var partition int32 = 0
	this.Partition = &partition
	return &this
}

// NewRollingStrategy instantiates a new RollingStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollingStrategy() *RollingStrategy {
	this := RollingStrategy{}
	var maxSurge string = "25%"
	this.MaxSurge = &maxSurge
	var maxUnavailable string = "25%"
	this.MaxUnavailable = &maxUnavailable
	var partition int32 = 0
	this.Partition = &partition
	return &this
}

// NewRollingStrategys converts a list RollingStrategy pointers to objects.
// This is helpful when the SetRollingStrategy requires a list of objects
func NewRollingStrategyList(ps ...*RollingStrategy) []RollingStrategy {
	objs := []RollingStrategy{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetMaxSurge returns the MaxSurge field value if set, zero value otherwise.
func (o *RollingStrategy) GetMaxSurge() string {
	if o == nil || utils.IsNil(o.MaxSurge) {
		var ret string
		return ret
	}
	return *o.MaxSurge
}

// GetMaxSurgeOk returns a tuple with the MaxSurge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollingStrategy) GetMaxSurgeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MaxSurge) {
		return nil, false
	}
	return o.MaxSurge, true
}

// HasMaxSurge returns a boolean if a field has been set.
func (o *RollingStrategy) HasMaxSurge() bool {
	if o != nil && !utils.IsNil(o.MaxSurge) {
		return true
	}

	return false
}

// SetMaxSurge gets a reference to the given string and assigns it to the maxSurge field.
// MaxSurge:
func (o *RollingStrategy) SetMaxSurge(v string) *RollingStrategy {
	o.MaxSurge = &v
	return o
}

// GetMaxUnavailable returns the MaxUnavailable field value if set, zero value otherwise.
func (o *RollingStrategy) GetMaxUnavailable() string {
	if o == nil || utils.IsNil(o.MaxUnavailable) {
		var ret string
		return ret
	}
	return *o.MaxUnavailable
}

// GetMaxUnavailableOk returns a tuple with the MaxUnavailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollingStrategy) GetMaxUnavailableOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MaxUnavailable) {
		return nil, false
	}
	return o.MaxUnavailable, true
}

// HasMaxUnavailable returns a boolean if a field has been set.
func (o *RollingStrategy) HasMaxUnavailable() bool {
	if o != nil && !utils.IsNil(o.MaxUnavailable) {
		return true
	}

	return false
}

// SetMaxUnavailable gets a reference to the given string and assigns it to the maxUnavailable field.
// MaxUnavailable:
func (o *RollingStrategy) SetMaxUnavailable(v string) *RollingStrategy {
	o.MaxUnavailable = &v
	return o
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *RollingStrategy) GetPartition() int32 {
	if o == nil || utils.IsNil(o.Partition) {
		var ret int32
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollingStrategy) GetPartitionOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Partition) {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *RollingStrategy) HasPartition() bool {
	if o != nil && !utils.IsNil(o.Partition) {
		return true
	}

	return false
}

// SetPartition gets a reference to the given int32 and assigns it to the partition field.
// Partition:
func (o *RollingStrategy) SetPartition(v int32) *RollingStrategy {
	o.Partition = &v
	return o
}

func (o RollingStrategy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RollingStrategy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.MaxSurge) {
		toSerialize["maxSurge"] = o.MaxSurge
	}
	if !utils.IsNil(o.MaxUnavailable) {
		toSerialize["maxUnavailable"] = o.MaxUnavailable
	}
	if !utils.IsNil(o.Partition) {
		toSerialize["partition"] = o.Partition
	}
	return toSerialize, nil
}

type NullableRollingStrategy struct {
	value *RollingStrategy
	isSet bool
}

func (v NullableRollingStrategy) Get() *RollingStrategy {
	return v.value
}

func (v *NullableRollingStrategy) Set(val *RollingStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableRollingStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableRollingStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRollingStrategy(val *RollingStrategy) *NullableRollingStrategy {
	return &NullableRollingStrategy{value: val, isSet: true}
}

func (v NullableRollingStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRollingStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
