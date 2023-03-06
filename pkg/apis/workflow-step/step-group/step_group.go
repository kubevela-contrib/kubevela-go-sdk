/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package step_group

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the StepGroupSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &StepGroupSpec{}

// StepGroupSpec struct for StepGroupSpec
type StepGroupSpec struct {
}

// NewStepGroupSpecWith instantiates a new StepGroupSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewStepGroupSpecWith() *StepGroupSpec {
	this := StepGroupSpec{}
	return &this
}

// NewStepGroupSpecWithDefault instantiates a new StepGroupSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStepGroupSpecWithDefault() *StepGroupSpec {
	this := StepGroupSpec{}
	return &this
}

// NewStepGroupSpec is short for NewStepGroupSpecWithDefault which instantiates a new StepGroupSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStepGroupSpec() *StepGroupSpec {
	return NewStepGroupSpecWithDefault()
}

// NewStepGroupSpecEmpty instantiates a new StepGroupSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewStepGroupSpecEmpty() *StepGroupSpec {
	this := StepGroupSpec{}
	return &this
}

// NewStepGroupSpecs converts a list StepGroupSpec pointers to objects.
// This is helpful when the SetStepGroupSpec requires a list of objects
func NewStepGroupSpecList(ps ...*StepGroupSpec) []StepGroupSpec {
	objs := []StepGroupSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this StepGroupSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *StepGroupWorkflowStep) Validate() error {
	// validate all nested properties
	return nil
}

func (o StepGroupSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StepGroupSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableStepGroupSpec struct {
	value *StepGroupSpec
	isSet bool
}

func (v *NullableStepGroupSpec) Get() *StepGroupSpec {
	return v.value
}

func (v *NullableStepGroupSpec) Set(val *StepGroupSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableStepGroupSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableStepGroupSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStepGroupSpec(val *StepGroupSpec) *NullableStepGroupSpec {
	return &NullableStepGroupSpec{value: val, isSet: true}
}

func (v NullableStepGroupSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStepGroupSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const StepGroupType = "step-group"

func init() {
	sdkcommon.RegisterWorkflowStep(StepGroupType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(StepGroupType, FromWorkflowSubStep)
}

type StepGroupWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties StepGroupSpec
}

func StepGroup(name string) *StepGroupWorkflowStep {
	s := &StepGroupWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: StepGroupType,
	}}
	return s
}

func (s *StepGroupWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range s.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  s.Base.DependsOn,
		If:         s.Base.If,
		Inputs:     s.Base.Inputs,
		Meta:       s.Base.Meta,
		Name:       s.Base.Name,
		Outputs:    s.Base.Outputs,
		Properties: util.Object2RawExtension(s.Properties),
		SubSteps:   subSteps,
		Timeout:    s.Base.Timeout,
		Type:       StepGroupType,
	}
	return res
}

func (s *StepGroupWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*StepGroupWorkflowStep, error) {
	var properties StepGroupSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := s.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	s.Base.Name = from.Name
	s.Base.DependsOn = from.DependsOn
	s.Base.Inputs = from.Inputs
	s.Base.Outputs = from.Outputs
	s.Base.If = from.If
	s.Base.Timeout = from.Timeout
	s.Base.Meta = from.Meta
	s.Base.Type = StepGroupType
	s.Properties = properties
	s.Base.SubSteps = subSteps
	return s, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	s := &StepGroupWorkflowStep{}
	return s.FromWorkflowStep(from)
}

func (s *StepGroupWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*StepGroupWorkflowStep, error) {
	var properties StepGroupSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	s.Base.Name = from.Name
	s.Base.DependsOn = from.DependsOn
	s.Base.Inputs = from.Inputs
	s.Base.Outputs = from.Outputs
	s.Base.If = from.If
	s.Base.Timeout = from.Timeout
	s.Base.Meta = from.Meta
	s.Base.Type = StepGroupType
	s.Properties = properties
	return s, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	s := &StepGroupWorkflowStep{}
	return s.FromWorkflowSubStep(from)
}

func (s *StepGroupWorkflowStep) WorkflowStepName() string {
	return s.Base.Name
}

func (s *StepGroupWorkflowStep) DefType() string {
	return StepGroupType
}

func (s *StepGroupWorkflowStep) If(_if string) *StepGroupWorkflowStep {
	s.Base.If = _if
	return s
}

func (s *StepGroupWorkflowStep) Alias(alias string) *StepGroupWorkflowStep {
	s.Base.Meta.Alias = alias
	return s
}

func (s *StepGroupWorkflowStep) Timeout(timeout string) *StepGroupWorkflowStep {
	s.Base.Timeout = timeout
	return s
}

func (s *StepGroupWorkflowStep) DependsOn(dependsOn []string) *StepGroupWorkflowStep {
	s.Base.DependsOn = dependsOn
	return s
}

func (s *StepGroupWorkflowStep) Inputs(input common.StepInputs) *StepGroupWorkflowStep {
	s.Base.Inputs = input
	return s
}

func (s *StepGroupWorkflowStep) Outputs(output common.StepOutputs) *StepGroupWorkflowStep {
	s.Base.Outputs = output
	return s
}

func (s *StepGroupWorkflowStep) AddSubStep(subStep apis.WorkflowStep) *StepGroupWorkflowStep {
	s.Base.SubSteps = append(s.Base.SubSteps, subStep)
	return s
}
