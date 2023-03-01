/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_once

import (
	"encoding/json"

	"github.com/kubevela-contrib/vela-go-sdk/pkg/apis/utils"
)

// checks if the ApplyOncePolicyRule type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApplyOncePolicyRule{}

// ApplyOncePolicyRule struct for ApplyOncePolicyRule
type ApplyOncePolicyRule struct {
	Selector *ResourcePolicyRuleSelector `json:"selector,omitempty"`
	Strategy *ApplyOnceStrategy          `json:"strategy,omitempty"`
}

// NewApplyOncePolicyRuleWith instantiates a new ApplyOncePolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyOncePolicyRuleWith() *ApplyOncePolicyRule {
	this := ApplyOncePolicyRule{}
	return &this
}

// NewApplyOncePolicyRule instantiates a new ApplyOncePolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyOncePolicyRule() *ApplyOncePolicyRule {
	this := ApplyOncePolicyRule{}
	return &this
}

// NewApplyOncePolicyRules converts a list ApplyOncePolicyRule pointers to objects.
// This is helpful when the SetApplyOncePolicyRule requires a list of objects
func NewApplyOncePolicyRuleList(ps ...*ApplyOncePolicyRule) []ApplyOncePolicyRule {
	objs := []ApplyOncePolicyRule{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *ApplyOncePolicyRule) GetSelector() ResourcePolicyRuleSelector {
	if o == nil || utils.IsNil(o.Selector) {
		var ret ResourcePolicyRuleSelector
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOncePolicyRule) GetSelectorOk() (*ResourcePolicyRuleSelector, bool) {
	if o == nil || utils.IsNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *ApplyOncePolicyRule) HasSelector() bool {
	if o != nil && !utils.IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given ResourcePolicyRuleSelector and assigns it to the selector field.
// Selector:
func (o *ApplyOncePolicyRule) SetSelector(v ResourcePolicyRuleSelector) *ApplyOncePolicyRule {
	o.Selector = &v
	return o
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *ApplyOncePolicyRule) GetStrategy() ApplyOnceStrategy {
	if o == nil || utils.IsNil(o.Strategy) {
		var ret ApplyOnceStrategy
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOncePolicyRule) GetStrategyOk() (*ApplyOnceStrategy, bool) {
	if o == nil || utils.IsNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *ApplyOncePolicyRule) HasStrategy() bool {
	if o != nil && !utils.IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given ApplyOnceStrategy and assigns it to the strategy field.
// Strategy:
func (o *ApplyOncePolicyRule) SetStrategy(v ApplyOnceStrategy) *ApplyOncePolicyRule {
	o.Strategy = &v
	return o
}

func (o ApplyOncePolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyOncePolicyRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	if !utils.IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	return toSerialize, nil
}

type NullableApplyOncePolicyRule struct {
	value *ApplyOncePolicyRule
	isSet bool
}

func (v NullableApplyOncePolicyRule) Get() *ApplyOncePolicyRule {
	return v.value
}

func (v *NullableApplyOncePolicyRule) Set(val *ApplyOncePolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyOncePolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyOncePolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyOncePolicyRule(val *ApplyOncePolicyRule) *NullableApplyOncePolicyRule {
	return &NullableApplyOncePolicyRule{value: val, isSet: true}
}

func (v NullableApplyOncePolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyOncePolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
