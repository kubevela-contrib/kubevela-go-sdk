/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resource_update

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the ResourceUpdateSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ResourceUpdateSpec{}

// ResourceUpdateSpec struct for ResourceUpdateSpec
type ResourceUpdateSpec struct {
	// Specify the list of rules to control resource update strategy at resource level.
	Rules []PolicyRule `json:"rules,omitempty"`
}

// NewResourceUpdateSpecWith instantiates a new ResourceUpdateSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewResourceUpdateSpecWith() *ResourceUpdateSpec {
	this := ResourceUpdateSpec{}
	return &this
}

// NewResourceUpdateSpecWithDefault instantiates a new ResourceUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceUpdateSpecWithDefault() *ResourceUpdateSpec {
	this := ResourceUpdateSpec{}
	return &this
}

// NewResourceUpdateSpec is short for NewResourceUpdateSpecWithDefault which instantiates a new ResourceUpdateSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceUpdateSpec() *ResourceUpdateSpec {
	return NewResourceUpdateSpecWithDefault()
}

// NewResourceUpdateSpecEmpty instantiates a new ResourceUpdateSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewResourceUpdateSpecEmpty() *ResourceUpdateSpec {
	this := ResourceUpdateSpec{}
	return &this
}

// NewResourceUpdateSpecs converts a list ResourceUpdateSpec pointers to objects.
// This is helpful when the SetResourceUpdateSpec requires a list of objects
func NewResourceUpdateSpecList(ps ...*ResourceUpdateSpec) []ResourceUpdateSpec {
	objs := []ResourceUpdateSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this ResourceUpdateSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *ResourceUpdatePolicy) Validate() error {
	// validate all nested properties
	return nil
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *ResourceUpdatePolicy) GetRules() []PolicyRule {
	if o == nil || utils.IsNil(o.Properties.Rules) {
		var ret []PolicyRule
		return ret
	}
	return o.Properties.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceUpdatePolicy) GetRulesOk() ([]PolicyRule, bool) {
	if o == nil || utils.IsNil(o.Properties.Rules) {
		return nil, false
	}
	return o.Properties.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *ResourceUpdatePolicy) HasRules() bool {
	if o != nil && !utils.IsNil(o.Properties.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []PolicyRule and assigns it to the rules field.
// Rules:  Specify the list of rules to control resource update strategy at resource level.
func (o *ResourceUpdatePolicy) SetRules(v []PolicyRule) *ResourceUpdatePolicy {
	o.Properties.Rules = v
	return o
}

func (o ResourceUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceUpdateSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableResourceUpdateSpec struct {
	value *ResourceUpdateSpec
	isSet bool
}

func (v *NullableResourceUpdateSpec) Get() *ResourceUpdateSpec {
	return v.value
}

func (v *NullableResourceUpdateSpec) Set(val *ResourceUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableResourceUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceUpdateSpec(val *ResourceUpdateSpec) *NullableResourceUpdateSpec {
	return &NullableResourceUpdateSpec{value: val, isSet: true}
}

func (v NullableResourceUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ResourceUpdateType = "resource-update"

func init() {
	sdkcommon.RegisterPolicy(ResourceUpdateType, FromPolicy)
}

type ResourceUpdatePolicy struct {
	Base       apis.PolicyBase
	Properties ResourceUpdateSpec
}

func ResourceUpdate(name string) *ResourceUpdatePolicy {
	r := &ResourceUpdatePolicy{Base: apis.PolicyBase{
		Name: name,
		Type: ResourceUpdateType,
	}}
	return r
}

func (r *ResourceUpdatePolicy) Build() v1beta1.AppPolicy {
	res := v1beta1.AppPolicy{
		Name:       r.Base.Name,
		Properties: util.Object2RawExtension(r.Properties),
		Type:       ResourceUpdateType,
	}
	return res
}

func (r *ResourceUpdatePolicy) FromPolicy(from v1beta1.AppPolicy) (*ResourceUpdatePolicy, error) {
	var properties ResourceUpdateSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	r.Base.Name = from.Name
	r.Base.Type = ResourceUpdateType
	r.Properties = properties
	return r, nil
}

func FromPolicy(from v1beta1.AppPolicy) (apis.Policy, error) {
	r := &ResourceUpdatePolicy{}
	return r.FromPolicy(from)
}

func (r *ResourceUpdatePolicy) PolicyName() string {
	return r.Base.Name
}

func (r *ResourceUpdatePolicy) DefType() string {
	return ResourceUpdateType
}
