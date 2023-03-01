/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s_objects

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/vela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/vela-go-sdk/pkg/apis/utils"
)

// checks if the K8sObjectsSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &K8sObjectsSpec{}

// K8sObjectsSpec struct for K8sObjectsSpec
type K8sObjectsSpec struct {
	Objects []map[string]interface{} `json:"objects,omitempty"`
}

// NewK8sObjectsSpecWith instantiates a new K8sObjectsSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewK8sObjectsSpecWith() *K8sObjectsSpec {
	this := K8sObjectsSpec{}
	return &this
}

// NewK8sObjectsSpec instantiates a new K8sObjectsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewK8sObjectsSpec() *K8sObjectsSpec {
	this := K8sObjectsSpec{}
	return &this
}

// NewK8sObjectsSpecs converts a list K8sObjectsSpec pointers to objects.
// This is helpful when the SetK8sObjectsSpec requires a list of objects
func NewK8sObjectsSpecList(ps ...*K8sObjectsSpec) []K8sObjectsSpec {
	objs := []K8sObjectsSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *K8sObjectsComponent) GetObjects() []map[string]interface{} {
	if o == nil || utils.IsNil(o.Properties.Objects) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Properties.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sObjectsComponent) GetObjectsOk() ([]map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Properties.Objects) {
		return nil, false
	}
	return o.Properties.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *K8sObjectsComponent) HasObjects() bool {
	if o != nil && !utils.IsNil(o.Properties.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []map[string]interface{} and assigns it to the objects field.
// Objects:
func (o *K8sObjectsComponent) SetObjects(v []map[string]interface{}) *K8sObjectsComponent {
	o.Properties.Objects = v
	return o
}

func (o K8sObjectsSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o K8sObjectsSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Objects) {
		toSerialize["objects"] = o.Objects
	}
	return toSerialize, nil
}

type NullableK8sObjectsSpec struct {
	value *K8sObjectsSpec
	isSet bool
}

func (v NullableK8sObjectsSpec) Get() *K8sObjectsSpec {
	return v.value
}

func (v *NullableK8sObjectsSpec) Set(val *K8sObjectsSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableK8sObjectsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableK8sObjectsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableK8sObjectsSpec(val *K8sObjectsSpec) *NullableK8sObjectsSpec {
	return &NullableK8sObjectsSpec{value: val, isSet: true}
}

func (v NullableK8sObjectsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableK8sObjectsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const K8sObjectsType = "k8s-objects"

func init() {
	sdkcommon.RegisterComponent(K8sObjectsType, FromComponent)
}

type K8sObjectsComponent struct {
	Base       apis.ComponentBase
	Properties K8sObjectsSpec
}

func K8sObjects(name string) *K8sObjectsComponent {
	k := &K8sObjectsComponent{Base: apis.ComponentBase{
		Name: name,
		Type: K8sObjectsType,
	}}
	return k
}

func (k *K8sObjectsComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range k.Base.Traits {
		traits = append(traits, trait.Build())
	}
	res := common.ApplicationComponent{
		DependsOn:  k.Base.DependsOn,
		Inputs:     k.Base.Inputs,
		Name:       k.Base.Name,
		Outputs:    k.Base.Outputs,
		Properties: util.Object2RawExtension(k.Properties),
		Traits:     traits,
		Type:       K8sObjectsType,
	}
	return res
}

func (k *K8sObjectsComponent) FromComponent(from common.ApplicationComponent) (*K8sObjectsComponent, error) {
	for _, trait := range from.Traits {
		_t, err := sdkcommon.FromTrait(&trait)
		if err != nil {
			return nil, err
		}
		k.Base.Traits = append(k.Base.Traits, _t)
	}
	var properties K8sObjectsSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	k.Base.Name = from.Name
	k.Base.DependsOn = from.DependsOn
	k.Base.Inputs = from.Inputs
	k.Base.Outputs = from.Outputs
	k.Base.Type = K8sObjectsType
	k.Properties = properties
	return k, nil
}

func FromComponent(from common.ApplicationComponent) (apis.Component, error) {
	k := &K8sObjectsComponent{}
	return k.FromComponent(from)
}

func (k *K8sObjectsComponent) SetTraits(traits ...apis.Trait) *K8sObjectsComponent {
	for _, addTrait := range traits {
		found := false
		for i, _t := range k.Base.Traits {
			if _t.DefType() == addTrait.DefType() {
				k.Base.Traits[i] = addTrait
				found = true
			}
			if !found {
				k.Base.Traits = append(k.Base.Traits, addTrait)
			}
		}
	}
	return k
}

func (k *K8sObjectsComponent) GetTrait(typ string) apis.Trait {
	for _, _t := range k.Base.Traits {
		if _t.DefType() == typ {
			return _t
		}
	}
	return nil
}

func (k *K8sObjectsComponent) ComponentName() string {
	return k.Base.Name
}

func (k *K8sObjectsComponent) DefType() string {
	return K8sObjectsType
}

func (k *K8sObjectsComponent) DependsOn(dependsOn []string) *K8sObjectsComponent {
	k.Base.DependsOn = dependsOn
	return k
}

func (k *K8sObjectsComponent) Inputs(input common.StepInputs) *K8sObjectsComponent {
	k.Base.Inputs = input
	return k
}

func (k *K8sObjectsComponent) Outputs(output common.StepOutputs) *K8sObjectsComponent {
	k.Base.Outputs = output
	return k
}

func (k *K8sObjectsComponent) AddDependsOn(dependsOn string) *K8sObjectsComponent {
	k.Base.DependsOn = append(k.Base.DependsOn, dependsOn)
	return k
}