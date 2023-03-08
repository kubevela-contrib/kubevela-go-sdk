/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affinity

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

// checks if the PodAffinity type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PodAffinity{}

// PodAffinity Specify the pod affinity scheduling rules
type PodAffinity struct {
	// Specify the preferred during scheduling ignored during execution
	Preferred []Preferred1 `json:"preferred,omitempty"`
	// Specify the required during scheduling ignored during execution
	Required []PodAffinityTerm `json:"required,omitempty"`
}

// NewPodAffinityWith instantiates a new PodAffinity object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewPodAffinityWith() *PodAffinity {
	this := PodAffinity{}
	return &this
}

// NewPodAffinityWithDefault instantiates a new PodAffinity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPodAffinityWithDefault() *PodAffinity {
	this := PodAffinity{}
	return &this
}

// NewPodAffinity is short for NewPodAffinityWithDefault which instantiates a new PodAffinity object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPodAffinity() *PodAffinity {
	return NewPodAffinityWithDefault()
}

// NewPodAffinityEmpty instantiates a new PodAffinity object with no properties set.
// This constructor will not assign any default values to properties.
func NewPodAffinityEmpty() *PodAffinity {
	this := PodAffinity{}
	return &this
}

// NewPodAffinitys converts a list PodAffinity pointers to objects.
// This is helpful when the SetPodAffinity requires a list of objects
func NewPodAffinityList(ps ...*PodAffinity) []PodAffinity {
	objs := []PodAffinity{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this PodAffinity
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *PodAffinity) Validate() error {
	// validate all nested properties
	return nil
}

// GetPreferred returns the Preferred field value if set, zero value otherwise.
func (o *PodAffinity) GetPreferred() []Preferred1 {
	if o == nil || utils.IsNil(o.Preferred) {
		var ret []Preferred1
		return ret
	}
	return o.Preferred
}

// GetPreferredOk returns a tuple with the Preferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodAffinity) GetPreferredOk() ([]Preferred1, bool) {
	if o == nil || utils.IsNil(o.Preferred) {
		return nil, false
	}
	return o.Preferred, true
}

// HasPreferred returns a boolean if a field has been set.
func (o *PodAffinity) HasPreferred() bool {
	if o != nil && !utils.IsNil(o.Preferred) {
		return true
	}

	return false
}

// SetPreferred gets a reference to the given []Preferred1 and assigns it to the preferred field.
// Preferred:  Specify the preferred during scheduling ignored during execution
func (o *PodAffinity) SetPreferred(v []Preferred1) *PodAffinity {
	o.Preferred = v
	return o
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *PodAffinity) GetRequired() []PodAffinityTerm {
	if o == nil || utils.IsNil(o.Required) {
		var ret []PodAffinityTerm
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodAffinity) GetRequiredOk() ([]PodAffinityTerm, bool) {
	if o == nil || utils.IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *PodAffinity) HasRequired() bool {
	if o != nil && !utils.IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given []PodAffinityTerm and assigns it to the required field.
// Required:  Specify the required during scheduling ignored during execution
func (o *PodAffinity) SetRequired(v []PodAffinityTerm) *PodAffinity {
	o.Required = v
	return o
}

func (o PodAffinity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PodAffinity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Preferred) {
		toSerialize["preferred"] = o.Preferred
	}
	if !utils.IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	return toSerialize, nil
}

type NullablePodAffinity struct {
	value *PodAffinity
	isSet bool
}

func (v *NullablePodAffinity) Get() *PodAffinity {
	return v.value
}

func (v *NullablePodAffinity) Set(val *PodAffinity) {
	v.value = val
	v.isSet = true
}

func (v *NullablePodAffinity) IsSet() bool {
	return v.isSet
}

func (v *NullablePodAffinity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePodAffinity(val *PodAffinity) *NullablePodAffinity {
	return &NullablePodAffinity{value: val, isSet: true}
}

func (v NullablePodAffinity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePodAffinity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
