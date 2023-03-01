/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s_update_strategy

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/vela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/vela-go-sdk/pkg/apis/utils"
)

// checks if the K8sUpdateStrategySpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &K8sUpdateStrategySpec{}

// K8sUpdateStrategySpec struct for K8sUpdateStrategySpec
type K8sUpdateStrategySpec struct {
	Strategy *Strategy `json:"strategy,omitempty"`
	// Specify the apiVersion of target
	TargetAPIVersion *string `json:"targetAPIVersion,omitempty"`
	// Specify the kind of target
	TargetKind *string `json:"targetKind,omitempty"`
}

// NewK8sUpdateStrategySpecWith instantiates a new K8sUpdateStrategySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewK8sUpdateStrategySpecWith() *K8sUpdateStrategySpec {
	this := K8sUpdateStrategySpec{}
	var targetAPIVersion string = "apps/v1"
	this.TargetAPIVersion = &targetAPIVersion
	var targetKind string = "Deployment"
	this.TargetKind = &targetKind
	return &this
}

// NewK8sUpdateStrategySpec instantiates a new K8sUpdateStrategySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewK8sUpdateStrategySpec() *K8sUpdateStrategySpec {
	this := K8sUpdateStrategySpec{}
	var targetAPIVersion string = "apps/v1"
	this.TargetAPIVersion = &targetAPIVersion
	var targetKind string = "Deployment"
	this.TargetKind = &targetKind
	return &this
}

// NewK8sUpdateStrategySpecs converts a list K8sUpdateStrategySpec pointers to objects.
// This is helpful when the SetK8sUpdateStrategySpec requires a list of objects
func NewK8sUpdateStrategySpecList(ps ...*K8sUpdateStrategySpec) []K8sUpdateStrategySpec {
	objs := []K8sUpdateStrategySpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *K8sUpdateStrategyTrait) GetStrategy() Strategy {
	if o == nil || utils.IsNil(o.Properties.Strategy) {
		var ret Strategy
		return ret
	}
	return *o.Properties.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sUpdateStrategyTrait) GetStrategyOk() (*Strategy, bool) {
	if o == nil || utils.IsNil(o.Properties.Strategy) {
		return nil, false
	}
	return o.Properties.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *K8sUpdateStrategyTrait) HasStrategy() bool {
	if o != nil && !utils.IsNil(o.Properties.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given Strategy and assigns it to the strategy field.
// Strategy:
func (o *K8sUpdateStrategyTrait) SetStrategy(v Strategy) *K8sUpdateStrategyTrait {
	o.Properties.Strategy = &v
	return o
}

// GetTargetAPIVersion returns the TargetAPIVersion field value if set, zero value otherwise.
func (o *K8sUpdateStrategyTrait) GetTargetAPIVersion() string {
	if o == nil || utils.IsNil(o.Properties.TargetAPIVersion) {
		var ret string
		return ret
	}
	return *o.Properties.TargetAPIVersion
}

// GetTargetAPIVersionOk returns a tuple with the TargetAPIVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sUpdateStrategyTrait) GetTargetAPIVersionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.TargetAPIVersion) {
		return nil, false
	}
	return o.Properties.TargetAPIVersion, true
}

// HasTargetAPIVersion returns a boolean if a field has been set.
func (o *K8sUpdateStrategyTrait) HasTargetAPIVersion() bool {
	if o != nil && !utils.IsNil(o.Properties.TargetAPIVersion) {
		return true
	}

	return false
}

// SetTargetAPIVersion gets a reference to the given string and assigns it to the targetAPIVersion field.
// TargetAPIVersion:  Specify the apiVersion of target
func (o *K8sUpdateStrategyTrait) SetTargetAPIVersion(v string) *K8sUpdateStrategyTrait {
	o.Properties.TargetAPIVersion = &v
	return o
}

// GetTargetKind returns the TargetKind field value if set, zero value otherwise.
func (o *K8sUpdateStrategyTrait) GetTargetKind() string {
	if o == nil || utils.IsNil(o.Properties.TargetKind) {
		var ret string
		return ret
	}
	return *o.Properties.TargetKind
}

// GetTargetKindOk returns a tuple with the TargetKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sUpdateStrategyTrait) GetTargetKindOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.TargetKind) {
		return nil, false
	}
	return o.Properties.TargetKind, true
}

// HasTargetKind returns a boolean if a field has been set.
func (o *K8sUpdateStrategyTrait) HasTargetKind() bool {
	if o != nil && !utils.IsNil(o.Properties.TargetKind) {
		return true
	}

	return false
}

// SetTargetKind gets a reference to the given string and assigns it to the targetKind field.
// TargetKind:  Specify the kind of target
func (o *K8sUpdateStrategyTrait) SetTargetKind(v string) *K8sUpdateStrategyTrait {
	o.Properties.TargetKind = &v
	return o
}

func (o K8sUpdateStrategySpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o K8sUpdateStrategySpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	if !utils.IsNil(o.TargetAPIVersion) {
		toSerialize["targetAPIVersion"] = o.TargetAPIVersion
	}
	if !utils.IsNil(o.TargetKind) {
		toSerialize["targetKind"] = o.TargetKind
	}
	return toSerialize, nil
}

type NullableK8sUpdateStrategySpec struct {
	value *K8sUpdateStrategySpec
	isSet bool
}

func (v NullableK8sUpdateStrategySpec) Get() *K8sUpdateStrategySpec {
	return v.value
}

func (v *NullableK8sUpdateStrategySpec) Set(val *K8sUpdateStrategySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableK8sUpdateStrategySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableK8sUpdateStrategySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableK8sUpdateStrategySpec(val *K8sUpdateStrategySpec) *NullableK8sUpdateStrategySpec {
	return &NullableK8sUpdateStrategySpec{value: val, isSet: true}
}

func (v NullableK8sUpdateStrategySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableK8sUpdateStrategySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const K8sUpdateStrategyType = "k8s-update-strategy"

func init() {
	sdkcommon.RegisterTrait(K8sUpdateStrategyType, FromTrait)
}

type K8sUpdateStrategyTrait struct {
	Base       apis.TraitBase
	Properties K8sUpdateStrategySpec
}

func K8sUpdateStrategy() *K8sUpdateStrategyTrait {
	k := &K8sUpdateStrategyTrait{Base: apis.TraitBase{}}
	return k
}

func (k *K8sUpdateStrategyTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(k.Properties),
		Type:       K8sUpdateStrategyType,
	}
	return res
}

func (k *K8sUpdateStrategyTrait) FromTrait(from common.ApplicationTrait) (*K8sUpdateStrategyTrait, error) {
	var properties K8sUpdateStrategySpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	k.Base.Type = K8sUpdateStrategyType
	k.Properties = properties
	return k, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	k := &K8sUpdateStrategyTrait{}
	return k.FromTrait(from)
}

func (k *K8sUpdateStrategyTrait) DefType() string {
	return K8sUpdateStrategyType
}
