/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package health

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

// checks if the HealthSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HealthSpec{}

// HealthSpec struct for HealthSpec
type HealthSpec struct {
	// Specify health checking interval(seconds), default 30s
	ProbeInterval *int32 `json:"probeInterval"`
	// Specify health checking timeout(seconds), default 10s
	ProbeTimeout *int32 `json:"probeTimeout"`
}

// NewHealthSpecWith instantiates a new HealthSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewHealthSpecWith(probeInterval int32, probeTimeout int32) *HealthSpec {
	this := HealthSpec{}
	this.ProbeInterval = &probeInterval
	this.ProbeTimeout = &probeTimeout
	return &this
}

// NewHealthSpecWithDefault instantiates a new HealthSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthSpecWithDefault() *HealthSpec {
	this := HealthSpec{}
	var probeInterval int32 = 30
	this.ProbeInterval = &probeInterval
	var probeTimeout int32 = 10
	this.ProbeTimeout = &probeTimeout
	return &this
}

// NewHealthSpec is short for NewHealthSpecWithDefault which instantiates a new HealthSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthSpec() *HealthSpec {
	return NewHealthSpecWithDefault()
}

// NewHealthSpecEmpty instantiates a new HealthSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewHealthSpecEmpty() *HealthSpec {
	this := HealthSpec{}
	return &this
}

// NewHealthSpecs converts a list HealthSpec pointers to objects.
// This is helpful when the SetHealthSpec requires a list of objects
func NewHealthSpecList(ps ...*HealthSpec) []HealthSpec {
	objs := []HealthSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this HealthSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *HealthPolicy) Validate() error {
	if o.Properties.ProbeInterval == nil {
		return errors.New("ProbeInterval in HealthSpec must be set")
	}
	if o.Properties.ProbeTimeout == nil {
		return errors.New("ProbeTimeout in HealthSpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetProbeInterval returns the ProbeInterval field value
func (o *HealthPolicy) GetProbeInterval() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Properties.ProbeInterval
}

// GetProbeIntervalOk returns a tuple with the ProbeInterval field value
// and a boolean to check if the value has been set.
func (o *HealthPolicy) GetProbeIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.ProbeInterval, true
}

// SetProbeInterval sets field value
func (o *HealthPolicy) SetProbeInterval(v int32) *HealthPolicy {
	o.Properties.ProbeInterval = &v
	return o
}

// GetProbeTimeout returns the ProbeTimeout field value
func (o *HealthPolicy) GetProbeTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Properties.ProbeTimeout
}

// GetProbeTimeoutOk returns a tuple with the ProbeTimeout field value
// and a boolean to check if the value has been set.
func (o *HealthPolicy) GetProbeTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.ProbeTimeout, true
}

// SetProbeTimeout sets field value
func (o *HealthPolicy) SetProbeTimeout(v int32) *HealthPolicy {
	o.Properties.ProbeTimeout = &v
	return o
}

func (o HealthSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["probeInterval"] = o.ProbeInterval
	toSerialize["probeTimeout"] = o.ProbeTimeout
	return toSerialize, nil
}

type NullableHealthSpec struct {
	value *HealthSpec
	isSet bool
}

func (v *NullableHealthSpec) Get() *HealthSpec {
	return v.value
}

func (v *NullableHealthSpec) Set(val *HealthSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableHealthSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthSpec(val *HealthSpec) *NullableHealthSpec {
	return &NullableHealthSpec{value: val, isSet: true}
}

func (v NullableHealthSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const HealthType = "health"

func init() {
	sdkcommon.RegisterPolicy(HealthType, FromPolicy)
}

type HealthPolicy struct {
	Base       apis.PolicyBase
	Properties HealthSpec
}

func Health(name string) *HealthPolicy {
	h := &HealthPolicy{Base: apis.PolicyBase{
		Name: name,
		Type: HealthType,
	}}
	return h
}

func (h *HealthPolicy) Build() v1beta1.AppPolicy {
	res := v1beta1.AppPolicy{
		Name:       h.Base.Name,
		Properties: util.Object2RawExtension(h.Properties),
		Type:       HealthType,
	}
	return res
}

func (h *HealthPolicy) FromPolicy(from v1beta1.AppPolicy) (*HealthPolicy, error) {
	var properties HealthSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	h.Base.Name = from.Name
	h.Base.Type = HealthType
	h.Properties = properties
	return h, nil
}

func FromPolicy(from v1beta1.AppPolicy) (apis.Policy, error) {
	h := &HealthPolicy{}
	return h.FromPolicy(from)
}

func (h *HealthPolicy) PolicyName() string {
	return h.Base.Name
}

func (h *HealthPolicy) DefType() string {
	return HealthType
}
