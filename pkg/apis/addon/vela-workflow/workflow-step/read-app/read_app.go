/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package read_app

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

// checks if the ReadAppSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReadAppSpec{}

// ReadAppSpec struct for ReadAppSpec
type ReadAppSpec struct {
	Name      *string `json:"name"`
	Namespace *string `json:"namespace"`
}

// NewReadAppSpecWith instantiates a new ReadAppSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewReadAppSpecWith(name string, namespace string) *ReadAppSpec {
	this := ReadAppSpec{}
	this.Name = &name
	this.Namespace = &namespace
	return &this
}

// NewReadAppSpecWithDefault instantiates a new ReadAppSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadAppSpecWithDefault() *ReadAppSpec {
	this := ReadAppSpec{}
	return &this
}

// NewReadAppSpec is short for NewReadAppSpecWithDefault which instantiates a new ReadAppSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadAppSpec() *ReadAppSpec {
	return NewReadAppSpecWithDefault()
}

// NewReadAppSpecEmpty instantiates a new ReadAppSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewReadAppSpecEmpty() *ReadAppSpec {
	this := ReadAppSpec{}
	return &this
}

// NewReadAppSpecs converts a list ReadAppSpec pointers to objects.
// This is helpful when the SetReadAppSpec requires a list of objects
func NewReadAppSpecList(ps ...*ReadAppSpec) []ReadAppSpec {
	objs := []ReadAppSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this ReadAppSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *ReadAppWorkflowStep) Validate() error {
	if o.Properties.Name == nil {
		return errors.New("Name in ReadAppSpec must be set")
	}
	if o.Properties.Namespace == nil {
		return errors.New("Namespace in ReadAppSpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetName returns the Name field value
func (o *ReadAppWorkflowStep) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReadAppWorkflowStep) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Name, true
}

// SetName sets field value
func (o *ReadAppWorkflowStep) SetName(v string) *ReadAppWorkflowStep {
	o.Properties.Name = &v
	return o
}

// GetNamespace returns the Namespace field value
func (o *ReadAppWorkflowStep) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *ReadAppWorkflowStep) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Namespace, true
}

// SetNamespace sets field value
func (o *ReadAppWorkflowStep) SetNamespace(v string) *ReadAppWorkflowStep {
	o.Properties.Namespace = &v
	return o
}

func (o ReadAppSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadAppSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullableReadAppSpec struct {
	value *ReadAppSpec
	isSet bool
}

func (v *NullableReadAppSpec) Get() *ReadAppSpec {
	return v.value
}

func (v *NullableReadAppSpec) Set(val *ReadAppSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableReadAppSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableReadAppSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadAppSpec(val *ReadAppSpec) *NullableReadAppSpec {
	return &NullableReadAppSpec{value: val, isSet: true}
}

func (v NullableReadAppSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadAppSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ReadAppType = "read-app"

func init() {
	sdkcommon.RegisterWorkflowStep(ReadAppType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(ReadAppType, FromWorkflowSubStep)
}

type ReadAppWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties ReadAppSpec
}

func ReadApp(name string) *ReadAppWorkflowStep {
	r := &ReadAppWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: ReadAppType,
	}}
	return r
}

func (r *ReadAppWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range r.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  r.Base.DependsOn,
		If:         r.Base.If,
		Inputs:     r.Base.Inputs,
		Meta:       r.Base.Meta,
		Name:       r.Base.Name,
		Outputs:    r.Base.Outputs,
		Properties: util.Object2RawExtension(r.Properties),
		SubSteps:   subSteps,
		Timeout:    r.Base.Timeout,
		Type:       ReadAppType,
	}
	return res
}

func (r *ReadAppWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*ReadAppWorkflowStep, error) {
	var properties ReadAppSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := r.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	r.Base.Name = from.Name
	r.Base.DependsOn = from.DependsOn
	r.Base.Inputs = from.Inputs
	r.Base.Outputs = from.Outputs
	r.Base.If = from.If
	r.Base.Timeout = from.Timeout
	r.Base.Meta = from.Meta
	r.Base.Type = ReadAppType
	r.Properties = properties
	r.Base.SubSteps = subSteps
	return r, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	r := &ReadAppWorkflowStep{}
	return r.FromWorkflowStep(from)
}

func (r *ReadAppWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*ReadAppWorkflowStep, error) {
	var properties ReadAppSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	r.Base.Name = from.Name
	r.Base.DependsOn = from.DependsOn
	r.Base.Inputs = from.Inputs
	r.Base.Outputs = from.Outputs
	r.Base.If = from.If
	r.Base.Timeout = from.Timeout
	r.Base.Meta = from.Meta
	r.Base.Type = ReadAppType
	r.Properties = properties
	return r, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	r := &ReadAppWorkflowStep{}
	return r.FromWorkflowSubStep(from)
}

func (r *ReadAppWorkflowStep) WorkflowStepName() string {
	return r.Base.Name
}

func (r *ReadAppWorkflowStep) DefType() string {
	return ReadAppType
}

func (r *ReadAppWorkflowStep) If(_if string) *ReadAppWorkflowStep {
	r.Base.If = _if
	return r
}

func (r *ReadAppWorkflowStep) Alias(alias string) *ReadAppWorkflowStep {
	r.Base.Meta.Alias = alias
	return r
}

func (r *ReadAppWorkflowStep) Timeout(timeout string) *ReadAppWorkflowStep {
	r.Base.Timeout = timeout
	return r
}

func (r *ReadAppWorkflowStep) DependsOn(dependsOn []string) *ReadAppWorkflowStep {
	r.Base.DependsOn = dependsOn
	return r
}

func (r *ReadAppWorkflowStep) Inputs(input common.StepInputs) *ReadAppWorkflowStep {
	r.Base.Inputs = input
	return r
}

func (r *ReadAppWorkflowStep) Outputs(output common.StepOutputs) *ReadAppWorkflowStep {
	r.Base.Outputs = output
	return r
}
