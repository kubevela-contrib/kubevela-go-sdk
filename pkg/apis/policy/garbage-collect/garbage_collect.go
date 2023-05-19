/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package garbage_collect

import (
	"encoding/json"
	"errors"

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
	// If set, it will override the default revision limit number and customize this number for the current application
	ApplicationRevisionLimit *int32 `json:"applicationRevisionLimit,omitempty"`
	// If is set, continue to execute gc when the workflow fails, by default gc will be executed only after the workflow succeeds
	ContinueOnFailure *bool `json:"continueOnFailure"`
	// If is set, outdated versioned resourcetracker will not be recycled automatically, outdated resources will be kept until resourcetracker be deleted manually
	KeepLegacyResource *bool `json:"keepLegacyResource"`
	// Specify the list of rules to control gc strategy at resource level, if one resource is controlled by multiple rules, first rule will be used
	Rules []GarbageCollectPolicyRule `json:"rules,omitempty"`
}

// NewGarbageCollectSpecWith instantiates a new GarbageCollectSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewGarbageCollectSpecWith(continueOnFailure bool, keepLegacyResource bool) *GarbageCollectSpec {
	this := GarbageCollectSpec{}
	this.ContinueOnFailure = &continueOnFailure
	this.KeepLegacyResource = &keepLegacyResource
	return &this
}

// NewGarbageCollectSpecWithDefault instantiates a new GarbageCollectSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGarbageCollectSpecWithDefault() *GarbageCollectSpec {
	this := GarbageCollectSpec{}
	var continueOnFailure bool = false
	this.ContinueOnFailure = &continueOnFailure
	var keepLegacyResource bool = false
	this.KeepLegacyResource = &keepLegacyResource
	return &this
}

// NewGarbageCollectSpec is short for NewGarbageCollectSpecWithDefault which instantiates a new GarbageCollectSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGarbageCollectSpec() *GarbageCollectSpec {
	return NewGarbageCollectSpecWithDefault()
}

// NewGarbageCollectSpecEmpty instantiates a new GarbageCollectSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewGarbageCollectSpecEmpty() *GarbageCollectSpec {
	this := GarbageCollectSpec{}
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

// Validate validates this GarbageCollectSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *GarbageCollectPolicy) Validate() error {
	if o.Properties.ContinueOnFailure == nil {
		return errors.New("ContinueOnFailure in GarbageCollectSpec must be set")
	}
	if o.Properties.KeepLegacyResource == nil {
		return errors.New("KeepLegacyResource in GarbageCollectSpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetApplicationRevisionLimit returns the ApplicationRevisionLimit field value if set, zero value otherwise.
func (o *GarbageCollectPolicy) GetApplicationRevisionLimit() int32 {
	if o == nil || utils.IsNil(o.Properties.ApplicationRevisionLimit) {
		var ret int32
		return ret
	}
	return *o.Properties.ApplicationRevisionLimit
}

// GetApplicationRevisionLimitOk returns a tuple with the ApplicationRevisionLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GarbageCollectPolicy) GetApplicationRevisionLimitOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.ApplicationRevisionLimit) {
		return nil, false
	}
	return o.Properties.ApplicationRevisionLimit, true
}

// HasApplicationRevisionLimit returns a boolean if a field has been set.
func (o *GarbageCollectPolicy) HasApplicationRevisionLimit() bool {
	if o != nil && !utils.IsNil(o.Properties.ApplicationRevisionLimit) {
		return true
	}

	return false
}

// SetApplicationRevisionLimit gets a reference to the given int32 and assigns it to the applicationRevisionLimit field.
// ApplicationRevisionLimit:  If set, it will override the default revision limit number and customize this number for the current application
func (o *GarbageCollectPolicy) SetApplicationRevisionLimit(v int32) *GarbageCollectPolicy {
	o.Properties.ApplicationRevisionLimit = &v
	return o
}

// GetContinueOnFailure returns the ContinueOnFailure field value
func (o *GarbageCollectPolicy) GetContinueOnFailure() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return *o.Properties.ContinueOnFailure
}

// GetContinueOnFailureOk returns a tuple with the ContinueOnFailure field value
// and a boolean to check if the value has been set.
func (o *GarbageCollectPolicy) GetContinueOnFailureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.ContinueOnFailure, true
}

// SetContinueOnFailure sets field value
func (o *GarbageCollectPolicy) SetContinueOnFailure(v bool) *GarbageCollectPolicy {
	o.Properties.ContinueOnFailure = &v
	return o
}

// GetKeepLegacyResource returns the KeepLegacyResource field value
func (o *GarbageCollectPolicy) GetKeepLegacyResource() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return *o.Properties.KeepLegacyResource
}

// GetKeepLegacyResourceOk returns a tuple with the KeepLegacyResource field value
// and a boolean to check if the value has been set.
func (o *GarbageCollectPolicy) GetKeepLegacyResourceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.KeepLegacyResource, true
}

// SetKeepLegacyResource sets field value
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
	if !utils.IsNil(o.ApplicationRevisionLimit) {
		toSerialize["applicationRevisionLimit"] = o.ApplicationRevisionLimit
	}
	toSerialize["continueOnFailure"] = o.ContinueOnFailure
	toSerialize["keepLegacyResource"] = o.KeepLegacyResource
	if !utils.IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableGarbageCollectSpec struct {
	value *GarbageCollectSpec
	isSet bool
}

func (v *NullableGarbageCollectSpec) Get() *GarbageCollectSpec {
	return v.value
}

func (v *NullableGarbageCollectSpec) Set(val *GarbageCollectSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableGarbageCollectSpec) IsSet() bool {
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
