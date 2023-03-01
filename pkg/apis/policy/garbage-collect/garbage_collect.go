/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package garbage_collect

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the GarbageCollectSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GarbageCollectSpec{}

// GarbageCollectSpec struct for GarbageCollectSpec
type GarbageCollectSpec struct {
	// If is set, outdated versioned resourcetracker will not be recycled automatically, outdated resources will be kept until resourcetracker be deleted manually
	KeepLegacyResource *bool `json:"keepLegacyResource,omitempty"`
	// Specify the list of rules to control gc strategy at resource level, if one resource is controlled by multiple rules, first rule will be used
	Rules []GarbageCollectPolicyRule `json:"rules,omitempty"`
}

// NewGarbageCollectSpecWith instantiates a new GarbageCollectSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGarbageCollectSpecWith() *GarbageCollectSpec {
	this := GarbageCollectSpec{}
	var keepLegacyResource bool = false
	this.KeepLegacyResource = &keepLegacyResource
	return &this
}

// NewGarbageCollectSpec instantiates a new GarbageCollectSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGarbageCollectSpec() *GarbageCollectSpec {
	this := GarbageCollectSpec{}
	var keepLegacyResource bool = false
	this.KeepLegacyResource = &keepLegacyResource
	return &this
}

// NewGarbageCollectSpecs converts a list GarbageCollectSpec pointers to objects.
// This is helpful when the SetGarbageCollectSpec requires a list of objects
func NewGarbageCollectSpecList(ps ...*GarbageCollectSpec) []GarbageCollectSpec {
	objs := []GarbageCollectSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetKeepLegacyResource returns the KeepLegacyResource field value if set, zero value otherwise.
func (o *GarbageCollectPolicy) GetKeepLegacyResource() bool {
	if o == nil || utils.IsNil(o.Properties.KeepLegacyResource) {
		var ret bool
		return ret
	}
	return *o.Properties.KeepLegacyResource
}

// GetKeepLegacyResourceOk returns a tuple with the KeepLegacyResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GarbageCollectPolicy) GetKeepLegacyResourceOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Properties.KeepLegacyResource) {
		return nil, false
	}
	return o.Properties.KeepLegacyResource, true
}

// HasKeepLegacyResource returns a boolean if a field has been set.
func (o *GarbageCollectPolicy) HasKeepLegacyResource() bool {
	if o != nil && !utils.IsNil(o.Properties.KeepLegacyResource) {
		return true
	}

	return false
}

// SetKeepLegacyResource gets a reference to the given bool and assigns it to the keepLegacyResource field.
// KeepLegacyResource:  If is set, outdated versioned resourcetracker will not be recycled automatically, outdated resources will be kept until resourcetracker be deleted manually
func (o *GarbageCollectPolicy) SetKeepLegacyResource(v bool) *GarbageCollectPolicy {
	o.Properties.KeepLegacyResource = &v
	return o
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *GarbageCollectPolicy) GetRules() []GarbageCollectPolicyRule {
	if o == nil || utils.IsNil(o.Properties.Rules) {
		var ret []GarbageCollectPolicyRule
		return ret
	}
	return o.Properties.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GarbageCollectPolicy) GetRulesOk() ([]GarbageCollectPolicyRule, bool) {
	if o == nil || utils.IsNil(o.Properties.Rules) {
		return nil, false
	}
	return o.Properties.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *GarbageCollectPolicy) HasRules() bool {
	if o != nil && !utils.IsNil(o.Properties.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []GarbageCollectPolicyRule and assigns it to the rules field.
// Rules:  Specify the list of rules to control gc strategy at resource level, if one resource is controlled by multiple rules, first rule will be used
func (o *GarbageCollectPolicy) SetRules(v []GarbageCollectPolicyRule) *GarbageCollectPolicy {
	o.Properties.Rules = v
	return o
}

func (o GarbageCollectSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GarbageCollectSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.KeepLegacyResource) {
		toSerialize["keepLegacyResource"] = o.KeepLegacyResource
	}
	if !utils.IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableGarbageCollectSpec struct {
	value *GarbageCollectSpec
	isSet bool
}

func (v NullableGarbageCollectSpec) Get() *GarbageCollectSpec {
	return v.value
}

func (v *NullableGarbageCollectSpec) Set(val *GarbageCollectSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableGarbageCollectSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGarbageCollectSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGarbageCollectSpec(val *GarbageCollectSpec) *NullableGarbageCollectSpec {
	return &NullableGarbageCollectSpec{value: val, isSet: true}
}

func (v NullableGarbageCollectSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGarbageCollectSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const GarbageCollectType = "garbage-collect"

func init() {
	sdkcommon.RegisterPolicy(GarbageCollectType, FromPolicy)
}

type GarbageCollectPolicy struct {
	Base       apis.PolicyBase
	Properties GarbageCollectSpec
}

func GarbageCollect(name string) *GarbageCollectPolicy {
	g := &GarbageCollectPolicy{Base: apis.PolicyBase{
		Name: name,
		Type: GarbageCollectType,
	}}
	return g
}

func (g *GarbageCollectPolicy) Build() v1beta1.AppPolicy {
	res := v1beta1.AppPolicy{
		Name:       g.Base.Name,
		Properties: util.Object2RawExtension(g.Properties),
		Type:       GarbageCollectType,
	}
	return res
}

func (g *GarbageCollectPolicy) FromPolicy(from v1beta1.AppPolicy) (*GarbageCollectPolicy, error) {
	var properties GarbageCollectSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	g.Base.Name = from.Name
	g.Base.Type = GarbageCollectType
	g.Properties = properties
	return g, nil
}

func FromPolicy(from v1beta1.AppPolicy) (apis.Policy, error) {
	g := &GarbageCollectPolicy{}
	return g.FromPolicy(from)
}

func (g *GarbageCollectPolicy) PolicyName() string {
	return g.Base.Name
}

func (g *GarbageCollectPolicy) DefType() string {
	return GarbageCollectType
}
