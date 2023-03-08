/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package replication

import (
	"encoding/json"
	"errors"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the ReplicationSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReplicationSpec{}

// ReplicationSpec struct for ReplicationSpec
type ReplicationSpec struct {
	// Spicify the keys of replication. Every key coresponds to a replication components
	Keys []string `json:"keys"`
	// Specify the components which will be replicated.
	Selector []string `json:"selector,omitempty"`
}

// NewReplicationSpecWith instantiates a new ReplicationSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewReplicationSpecWith(keys []string) *ReplicationSpec {
	this := ReplicationSpec{}
	this.Keys = keys
	return &this
}

// NewReplicationSpecWithDefault instantiates a new ReplicationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicationSpecWithDefault() *ReplicationSpec {
	this := ReplicationSpec{}
	return &this
}

// NewReplicationSpec is short for NewReplicationSpecWithDefault which instantiates a new ReplicationSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicationSpec() *ReplicationSpec {
	return NewReplicationSpecWithDefault()
}

// NewReplicationSpecEmpty instantiates a new ReplicationSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewReplicationSpecEmpty() *ReplicationSpec {
	this := ReplicationSpec{}
	return &this
}

// NewReplicationSpecs converts a list ReplicationSpec pointers to objects.
// This is helpful when the SetReplicationSpec requires a list of objects
func NewReplicationSpecList(ps ...*ReplicationSpec) []ReplicationSpec {
	objs := []ReplicationSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this ReplicationSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *ReplicationPolicy) Validate() error {
	if o.Properties.Keys == nil {
		return errors.New("Keys in ReplicationSpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetKeys returns the Keys field value
func (o *ReplicationPolicy) GetKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Properties.Keys
}

// GetKeysOk returns a tuple with the Keys field value
// and a boolean to check if the value has been set.
func (o *ReplicationPolicy) GetKeysOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Keys, true
}

// SetKeys sets field value
func (o *ReplicationPolicy) SetKeys(v []string) *ReplicationPolicy {
	o.Properties.Keys = v
	return o
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *ReplicationPolicy) GetSelector() []string {
	if o == nil || utils.IsNil(o.Properties.Selector) {
		var ret []string
		return ret
	}
	return o.Properties.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationPolicy) GetSelectorOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Selector) {
		return nil, false
	}
	return o.Properties.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *ReplicationPolicy) HasSelector() bool {
	if o != nil && !utils.IsNil(o.Properties.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given []string and assigns it to the selector field.
// Selector:  Specify the components which will be replicated.
func (o *ReplicationPolicy) SetSelector(v []string) *ReplicationPolicy {
	o.Properties.Selector = v
	return o
}

func (o ReplicationSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplicationSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keys"] = o.Keys
	if !utils.IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	return toSerialize, nil
}

type NullableReplicationSpec struct {
	value *ReplicationSpec
	isSet bool
}

func (v *NullableReplicationSpec) Get() *ReplicationSpec {
	return v.value
}

func (v *NullableReplicationSpec) Set(val *ReplicationSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableReplicationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicationSpec(val *ReplicationSpec) *NullableReplicationSpec {
	return &NullableReplicationSpec{value: val, isSet: true}
}

func (v NullableReplicationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ReplicationType = "replication"

func init() {
	sdkcommon.RegisterPolicy(ReplicationType, FromPolicy)
}

type ReplicationPolicy struct {
	Base       apis.PolicyBase
	Properties ReplicationSpec
}

func Replication(name string) *ReplicationPolicy {
	r := &ReplicationPolicy{Base: apis.PolicyBase{
		Name: name,
		Type: ReplicationType,
	}}
	return r
}

func (r *ReplicationPolicy) Build() v1beta1.AppPolicy {
	res := v1beta1.AppPolicy{
		Name:       r.Base.Name,
		Properties: util.Object2RawExtension(r.Properties),
		Type:       ReplicationType,
	}
	return res
}

func (r *ReplicationPolicy) FromPolicy(from v1beta1.AppPolicy) (*ReplicationPolicy, error) {
	var properties ReplicationSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	r.Base.Name = from.Name
	r.Base.Type = ReplicationType
	r.Properties = properties
	return r, nil
}

func FromPolicy(from v1beta1.AppPolicy) (apis.Policy, error) {
	r := &ReplicationPolicy{}
	return r.FromPolicy(from)
}

func (r *ReplicationPolicy) PolicyName() string {
	return r.Base.Name
}

func (r *ReplicationPolicy) DefType() string {
	return ReplicationType
}
