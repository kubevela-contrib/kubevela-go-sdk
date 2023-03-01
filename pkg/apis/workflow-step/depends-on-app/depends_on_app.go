/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package depends_on_app

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/vela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/vela-go-sdk/pkg/apis/utils"
)

// checks if the DependsOnAppSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DependsOnAppSpec{}

// DependsOnAppSpec struct for DependsOnAppSpec
type DependsOnAppSpec struct {
	// Specify the name of the dependent Application
	Name *string `json:"name,omitempty"`
	// Specify the namespace of the dependent Application
	Namespace *string `json:"namespace,omitempty"`
}

// NewDependsOnAppSpecWith instantiates a new DependsOnAppSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDependsOnAppSpecWith() *DependsOnAppSpec {
	this := DependsOnAppSpec{}
	return &this
}

// NewDependsOnAppSpec instantiates a new DependsOnAppSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDependsOnAppSpec() *DependsOnAppSpec {
	this := DependsOnAppSpec{}
	return &this
}

// NewDependsOnAppSpecs converts a list DependsOnAppSpec pointers to objects.
// This is helpful when the SetDependsOnAppSpec requires a list of objects
func NewDependsOnAppSpecList(ps ...*DependsOnAppSpec) []DependsOnAppSpec {
	objs := []DependsOnAppSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DependsOnAppWorkflowStep) GetName() string {
	if o == nil || utils.IsNil(o.Properties.Name) {
		var ret string
		return ret
	}
	return *o.Properties.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DependsOnAppWorkflowStep) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Name) {
		return nil, false
	}
	return o.Properties.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DependsOnAppWorkflowStep) HasName() bool {
	if o != nil && !utils.IsNil(o.Properties.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:  Specify the name of the dependent Application
func (o *DependsOnAppWorkflowStep) SetName(v string) *DependsOnAppWorkflowStep {
	o.Properties.Name = &v
	return o
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *DependsOnAppWorkflowStep) GetNamespace() string {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		var ret string
		return ret
	}
	return *o.Properties.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DependsOnAppWorkflowStep) GetNamespaceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		return nil, false
	}
	return o.Properties.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *DependsOnAppWorkflowStep) HasNamespace() bool {
	if o != nil && !utils.IsNil(o.Properties.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the namespace field.
// Namespace:  Specify the namespace of the dependent Application
func (o *DependsOnAppWorkflowStep) SetNamespace(v string) *DependsOnAppWorkflowStep {
	o.Properties.Namespace = &v
	return o
}

func (o DependsOnAppSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DependsOnAppSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	return toSerialize, nil
}

type NullableDependsOnAppSpec struct {
	value *DependsOnAppSpec
	isSet bool
}

func (v NullableDependsOnAppSpec) Get() *DependsOnAppSpec {
	return v.value
}

func (v *NullableDependsOnAppSpec) Set(val *DependsOnAppSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDependsOnAppSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDependsOnAppSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDependsOnAppSpec(val *DependsOnAppSpec) *NullableDependsOnAppSpec {
	return &NullableDependsOnAppSpec{value: val, isSet: true}
}

func (v NullableDependsOnAppSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDependsOnAppSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const DependsOnAppType = "depends-on-app"

func init() {
	sdkcommon.RegisterWorkflowStep(DependsOnAppType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(DependsOnAppType, FromWorkflowSubStep)
}

type DependsOnAppWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties DependsOnAppSpec
}

func DependsOnApp(name string) *DependsOnAppWorkflowStep {
	d := &DependsOnAppWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: DependsOnAppType,
	}}
	return d
}

func (d *DependsOnAppWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range d.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  d.Base.DependsOn,
		If:         d.Base.If,
		Inputs:     d.Base.Inputs,
		Meta:       d.Base.Meta,
		Name:       d.Base.Name,
		Outputs:    d.Base.Outputs,
		Properties: util.Object2RawExtension(d.Properties),
		SubSteps:   subSteps,
		Timeout:    d.Base.Timeout,
		Type:       DependsOnAppType,
	}
	return res
}

func (d *DependsOnAppWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*DependsOnAppWorkflowStep, error) {
	var properties DependsOnAppSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := d.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	d.Base.Name = from.Name
	d.Base.DependsOn = from.DependsOn
	d.Base.Inputs = from.Inputs
	d.Base.Outputs = from.Outputs
	d.Base.If = from.If
	d.Base.Timeout = from.Timeout
	d.Base.Meta = from.Meta
	d.Base.Type = DependsOnAppType
	d.Properties = properties
	d.Base.SubSteps = subSteps
	return d, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	d := &DependsOnAppWorkflowStep{}
	return d.FromWorkflowStep(from)
}

func (d *DependsOnAppWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*DependsOnAppWorkflowStep, error) {
	var properties DependsOnAppSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	d.Base.Name = from.Name
	d.Base.DependsOn = from.DependsOn
	d.Base.Inputs = from.Inputs
	d.Base.Outputs = from.Outputs
	d.Base.If = from.If
	d.Base.Timeout = from.Timeout
	d.Base.Meta = from.Meta
	d.Base.Type = DependsOnAppType
	d.Properties = properties
	return d, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	d := &DependsOnAppWorkflowStep{}
	return d.FromWorkflowSubStep(from)
}

func (d *DependsOnAppWorkflowStep) WorkflowStepName() string {
	return d.Base.Name
}

func (d *DependsOnAppWorkflowStep) DefType() string {
	return DependsOnAppType
}

func (d *DependsOnAppWorkflowStep) If(_if string) *DependsOnAppWorkflowStep {
	d.Base.If = _if
	return d
}

func (d *DependsOnAppWorkflowStep) Alias(alias string) *DependsOnAppWorkflowStep {
	d.Base.Meta.Alias = alias
	return d
}

func (d *DependsOnAppWorkflowStep) Timeout(timeout string) *DependsOnAppWorkflowStep {
	d.Base.Timeout = timeout
	return d
}

func (d *DependsOnAppWorkflowStep) DependsOn(dependsOn []string) *DependsOnAppWorkflowStep {
	d.Base.DependsOn = dependsOn
	return d
}

func (d *DependsOnAppWorkflowStep) Inputs(input common.StepInputs) *DependsOnAppWorkflowStep {
	d.Base.Inputs = input
	return d
}

func (d *DependsOnAppWorkflowStep) Outputs(output common.StepOutputs) *DependsOnAppWorkflowStep {
	d.Base.Outputs = output
	return d
}
