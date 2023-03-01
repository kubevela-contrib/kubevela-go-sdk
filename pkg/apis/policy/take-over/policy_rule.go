/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package take_over

import (
	"encoding/json"

	"github.com/kubevela-contrib/vela-go-sdk/pkg/apis/utils"
)

// checks if the PolicyRule type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PolicyRule{}

// PolicyRule struct for PolicyRule
type PolicyRule struct {
	// Specify how to select the targets of the rule
	Selector []RuleSelector `json:"selector,omitempty"`
}

// NewPolicyRuleWith instantiates a new PolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyRuleWith() *PolicyRule {
	this := PolicyRule{}
	return &this
}

// NewPolicyRule instantiates a new PolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyRule() *PolicyRule {
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

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *PolicyRule) GetSelector() []RuleSelector {
	if o == nil || utils.IsNil(o.Selector) {
		var ret []RuleSelector
		return ret
	}
	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRule) GetSelectorOk() ([]RuleSelector, bool) {
	if o == nil || utils.IsNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *PolicyRule) HasSelector() bool {
	if o != nil && !utils.IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given []RuleSelector and assigns it to the selector field.
// Selector:  Specify how to select the targets of the rule
func (o *PolicyRule) SetSelector(v []RuleSelector) *PolicyRule {
	o.Selector = v
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
	if !utils.IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	return toSerialize, nil
}

type NullablePolicyRule struct {
	value *PolicyRule
	isSet bool
}

func (v NullablePolicyRule) Get() *PolicyRule {
	return v.value
}

func (v *NullablePolicyRule) Set(val *PolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyRule) IsSet() bool {
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