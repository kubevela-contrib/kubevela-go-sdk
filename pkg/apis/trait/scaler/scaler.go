/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package scaler

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

// checks if the ScalerSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ScalerSpec{}

// ScalerSpec struct for ScalerSpec
type ScalerSpec struct {
	// Specify the number of workload
	Replicas *int32 `json:"replicas"`
}

// NewScalerSpecWith instantiates a new ScalerSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewScalerSpecWith(replicas int32) *ScalerSpec {
	this := ScalerSpec{}
	this.Replicas = &replicas
	return &this
}

// NewScalerSpecWithDefault instantiates a new ScalerSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScalerSpecWithDefault() *ScalerSpec {
	this := ScalerSpec{}
	var replicas int32 = 1
	this.Replicas = &replicas
	return &this
}

// NewScalerSpec is short for NewScalerSpecWithDefault which instantiates a new ScalerSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScalerSpec() *ScalerSpec {
	return NewScalerSpecWithDefault()
}

// NewScalerSpecEmpty instantiates a new ScalerSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewScalerSpecEmpty() *ScalerSpec {
	this := ScalerSpec{}
	return &this
}

// NewScalerSpecs converts a list ScalerSpec pointers to objects.
// This is helpful when the SetScalerSpec requires a list of objects
func NewScalerSpecList(ps ...*ScalerSpec) []ScalerSpec {
	objs := []ScalerSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this ScalerSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *ScalerTrait) Validate() error {
	if o.Properties.Replicas == nil {
		return errors.New("Replicas in ScalerSpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetReplicas returns the Replicas field value
func (o *ScalerTrait) GetReplicas() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Properties.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *ScalerTrait) GetReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Replicas, true
}

// SetReplicas sets field value
func (o *ScalerTrait) SetReplicas(v int32) *ScalerTrait {
	o.Properties.Replicas = &v
	return o
}

func (o ScalerSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScalerSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["replicas"] = o.Replicas
	return toSerialize, nil
}

type NullableScalerSpec struct {
	value *ScalerSpec
	isSet bool
}

func (v *NullableScalerSpec) Get() *ScalerSpec {
	return v.value
}

func (v *NullableScalerSpec) Set(val *ScalerSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableScalerSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableScalerSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScalerSpec(val *ScalerSpec) *NullableScalerSpec {
	return &NullableScalerSpec{value: val, isSet: true}
}

func (v NullableScalerSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScalerSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ScalerType = "scaler"

func init() {
	sdkcommon.RegisterTrait(ScalerType, FromTrait)
}

type ScalerTrait struct {
	Base       apis.TraitBase
	Properties ScalerSpec
}

func Scaler() *ScalerTrait {
	s := &ScalerTrait{Base: apis.TraitBase{}}
	return s
}

func (s *ScalerTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(s.Properties),
		Type:       ScalerType,
	}
	return res
}

func (s *ScalerTrait) FromTrait(from common.ApplicationTrait) (*ScalerTrait, error) {
	var properties ScalerSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	s.Base.Type = ScalerType
	s.Properties = properties
	return s, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	s := &ScalerTrait{}
	return s.FromTrait(from)
}

func (s *ScalerTrait) DefType() string {
	return ScalerType
}
