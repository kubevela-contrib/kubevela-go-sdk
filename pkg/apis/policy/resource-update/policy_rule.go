/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resource_update

import (
	"encoding/json"
	"errors"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the PolicyRule type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PolicyRule{}

// PolicyRule struct for PolicyRule
type PolicyRule struct {
	Selector *RuleSelector `json:"selector"`
	Strategy *Strategy     `json:"strategy"`
}

// NewPolicyRuleWith instantiates a new PolicyRule object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewPolicyRuleWith(selector RuleSelector, strategy Strategy) *PolicyRule {
	this := PolicyRule{}
	this.Selector = &selector
	this.Strategy = &strategy
	return &this
}

// NewPolicyRuleWithDefault instantiates a new PolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyRuleWithDefault() *PolicyRule {
	this := PolicyRule{}
	return &this
}

// NewPolicyRule is short for NewPolicyRuleWithDefault which instantiates a new PolicyRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyRule() *PolicyRule {
	return NewPolicyRuleWithDefault()
}

// NewPolicyRuleEmpty instantiates a new PolicyRule object with no properties set.
// This constructor will not assign any default values to properties.
func NewPolicyRuleEmpty() *PolicyRule {
	this := PolicyRule{}
	return &this
}

// NewPolicyRules converts a list PolicyRule pointers to objects.
// This is helpful when the SetPolicyRule requires a list of objects
func NewPolicyRuleList(ps ...*PolicyRule) []PolicyRule {
	objs := []PolicyRule{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this PolicyRule
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *PolicyRule) Validate() error {
	if o.Selector == nil {
		return errors.New("Selector in PolicyRule must be set")
	}
	if o.Strategy == nil {
		return errors.New("Strategy in PolicyRule must be set")
	}
	// validate all nested properties
	if o.Selector != nil {
		if err := o.Selector.Validate(); err != nil {
			return err
		}
	}
	if o.Strategy != nil {
		if err := o.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// GetSelector returns the Selector field value
func (o *PolicyRule) GetSelector() RuleSelector {
	if o == nil {
		var ret RuleSelector
		return ret
	}

	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value
// and a boolean to check if the value has been set.
func (o *PolicyRule) GetSelectorOk() (*RuleSelector, bool) {
	if o == nil {
		return nil, false
	}
	return o.Selector, true
}

// SetSelector sets field value
func (o *PolicyRule) SetSelector(v RuleSelector) *PolicyRule {
	o.Selector = &v
	return o
}

// GetStrategy returns the Strategy field value
func (o *PolicyRule) GetStrategy() Strategy {
	if o == nil {
		var ret Strategy
		return ret
	}

	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *PolicyRule) GetStrategyOk() (*Strategy, bool) {
	if o == nil {
		return nil, false
	}
	return o.Strategy, true
}

// SetStrategy sets field value
func (o *PolicyRule) SetStrategy(v Strategy) *PolicyRule {
	o.Strategy = &v
	return o
}

func (o PolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["selector"] = o.Selector
	toSerialize["strategy"] = o.Strategy
	return toSerialize, nil
}

type NullablePolicyRule struct {
	value *PolicyRule
	isSet bool
}

func (v *NullablePolicyRule) Get() *PolicyRule {
	return v.value
}

func (v *NullablePolicyRule) Set(val *PolicyRule) {
	v.value = val
	v.isSet = true
}

func (v *NullablePolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyRule(val *PolicyRule) *NullablePolicyRule {
	return &NullablePolicyRule{value: val, isSet: true}
}

func (v NullablePolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
