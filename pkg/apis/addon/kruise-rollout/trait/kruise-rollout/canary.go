/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kruise_rollout

import (
	"encoding/json"
	"errors"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Canary type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Canary{}

// Canary struct for Canary
type Canary struct {
	// Defines the entire rollout process in steps
	Steps []CanaryStep `json:"steps"`
	// Define traffic routing related service, ingress information
	TrafficRoutings []TrafficRouting `json:"trafficRoutings,omitempty"`
}

// NewCanaryWith instantiates a new Canary object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewCanaryWith(steps []CanaryStep) *Canary {
	this := Canary{}
	this.Steps = steps
	return &this
}

// NewCanaryWithDefault instantiates a new Canary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCanaryWithDefault() *Canary {
	this := Canary{}
	return &this
}

// NewCanary is short for NewCanaryWithDefault which instantiates a new Canary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCanary() *Canary {
	return NewCanaryWithDefault()
}

// NewCanaryEmpty instantiates a new Canary object with no properties set.
// This constructor will not assign any default values to properties.
func NewCanaryEmpty() *Canary {
	this := Canary{}
	return &this
}

// NewCanarys converts a list Canary pointers to objects.
// This is helpful when the SetCanary requires a list of objects
func NewCanaryList(ps ...*Canary) []Canary {
	objs := []Canary{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Canary
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Canary) Validate() error {
	if o.Steps == nil {
		return errors.New("Steps in Canary must be set")
	}
	// validate all nested properties
	return nil
}

// GetSteps returns the Steps field value
func (o *Canary) GetSteps() []CanaryStep {
	if o == nil {
		var ret []CanaryStep
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *Canary) GetStepsOk() ([]CanaryStep, bool) {
	if o == nil {
		return nil, false
	}
	return o.Steps, true
}

// SetSteps sets field value
func (o *Canary) SetSteps(v []CanaryStep) *Canary {
	o.Steps = v
	return o
}

// GetTrafficRoutings returns the TrafficRoutings field value if set, zero value otherwise.
func (o *Canary) GetTrafficRoutings() []TrafficRouting {
	if o == nil || utils.IsNil(o.TrafficRoutings) {
		var ret []TrafficRouting
		return ret
	}
	return o.TrafficRoutings
}

// GetTrafficRoutingsOk returns a tuple with the TrafficRoutings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Canary) GetTrafficRoutingsOk() ([]TrafficRouting, bool) {
	if o == nil || utils.IsNil(o.TrafficRoutings) {
		return nil, false
	}
	return o.TrafficRoutings, true
}

// HasTrafficRoutings returns a boolean if a field has been set.
func (o *Canary) HasTrafficRoutings() bool {
	if o != nil && !utils.IsNil(o.TrafficRoutings) {
		return true
	}

	return false
}

// SetTrafficRoutings gets a reference to the given []TrafficRouting and assigns it to the trafficRoutings field.
// TrafficRoutings:  Define traffic routing related service, ingress information
func (o *Canary) SetTrafficRoutings(v []TrafficRouting) *Canary {
	o.TrafficRoutings = v
	return o
}

func (o Canary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Canary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["steps"] = o.Steps
	if !utils.IsNil(o.TrafficRoutings) {
		toSerialize["trafficRoutings"] = o.TrafficRoutings
	}
	return toSerialize, nil
}

type NullableCanary struct {
	value *Canary
	isSet bool
}

func (v *NullableCanary) Get() *Canary {
	return v.value
}

func (v *NullableCanary) Set(val *Canary) {
	v.value = val
	v.isSet = true
}

func (v *NullableCanary) IsSet() bool {
	return v.isSet
}

func (v *NullableCanary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCanary(val *Canary) *NullableCanary {
	return &NullableCanary{value: val, isSet: true}
}

func (v NullableCanary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCanary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
