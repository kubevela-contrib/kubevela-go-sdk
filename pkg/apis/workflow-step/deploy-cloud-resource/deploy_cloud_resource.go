/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deploy_cloud_resource

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the DeployCloudResourceSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DeployCloudResourceSpec{}

// DeployCloudResourceSpec struct for DeployCloudResourceSpec
type DeployCloudResourceSpec struct {
	// Declare the name of the env in policy
	Env *string `json:"env,omitempty"`
	// Declare the name of the env-binding policy, if empty, the first env-binding policy will be used
	Policy *string `json:"policy,omitempty"`
}

// NewDeployCloudResourceSpecWith instantiates a new DeployCloudResourceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployCloudResourceSpecWith() *DeployCloudResourceSpec {
	this := DeployCloudResourceSpec{}
	var policy string = ""
	this.Policy = &policy
	return &this
}

// NewDeployCloudResourceSpec instantiates a new DeployCloudResourceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeployCloudResourceSpec() *DeployCloudResourceSpec {
	this := DeployCloudResourceSpec{}
	var policy string = ""
	this.Policy = &policy
	return &this
}

// NewDeployCloudResourceSpecs converts a list DeployCloudResourceSpec pointers to objects.
// This is helpful when the SetDeployCloudResourceSpec requires a list of objects
func NewDeployCloudResourceSpecList(ps ...*DeployCloudResourceSpec) []DeployCloudResourceSpec {
	objs := []DeployCloudResourceSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *DeployCloudResourceWorkflowStep) GetEnv() string {
	if o == nil || utils.IsNil(o.Properties.Env) {
		var ret string
		return ret
	}
	return *o.Properties.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployCloudResourceWorkflowStep) GetEnvOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Env) {
		return nil, false
	}
	return o.Properties.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *DeployCloudResourceWorkflowStep) HasEnv() bool {
	if o != nil && !utils.IsNil(o.Properties.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given string and assigns it to the env field.
// Env:  Declare the name of the env in policy
func (o *DeployCloudResourceWorkflowStep) SetEnv(v string) *DeployCloudResourceWorkflowStep {
	o.Properties.Env = &v
	return o
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *DeployCloudResourceWorkflowStep) GetPolicy() string {
	if o == nil || utils.IsNil(o.Properties.Policy) {
		var ret string
		return ret
	}
	return *o.Properties.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployCloudResourceWorkflowStep) GetPolicyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Policy) {
		return nil, false
	}
	return o.Properties.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *DeployCloudResourceWorkflowStep) HasPolicy() bool {
	if o != nil && !utils.IsNil(o.Properties.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the policy field.
// Policy:  Declare the name of the env-binding policy, if empty, the first env-binding policy will be used
func (o *DeployCloudResourceWorkflowStep) SetPolicy(v string) *DeployCloudResourceWorkflowStep {
	o.Properties.Policy = &v
	return o
}

func (o DeployCloudResourceSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeployCloudResourceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !utils.IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	return toSerialize, nil
}

type NullableDeployCloudResourceSpec struct {
	value *DeployCloudResourceSpec
	isSet bool
}

func (v NullableDeployCloudResourceSpec) Get() *DeployCloudResourceSpec {
	return v.value
}

func (v *NullableDeployCloudResourceSpec) Set(val *DeployCloudResourceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployCloudResourceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployCloudResourceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployCloudResourceSpec(val *DeployCloudResourceSpec) *NullableDeployCloudResourceSpec {
	return &NullableDeployCloudResourceSpec{value: val, isSet: true}
}

func (v NullableDeployCloudResourceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployCloudResourceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const DeployCloudResourceType = "deploy-cloud-resource"

func init() {
	sdkcommon.RegisterWorkflowStep(DeployCloudResourceType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(DeployCloudResourceType, FromWorkflowSubStep)
}

type DeployCloudResourceWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties DeployCloudResourceSpec
}

func DeployCloudResource(name string) *DeployCloudResourceWorkflowStep {
	d := &DeployCloudResourceWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: DeployCloudResourceType,
	}}
	return d
}

func (d *DeployCloudResourceWorkflowStep) Build() v1beta1.WorkflowStep {
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
		Type:       DeployCloudResourceType,
	}
	return res
}

func (d *DeployCloudResourceWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*DeployCloudResourceWorkflowStep, error) {
	var properties DeployCloudResourceSpec
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
	d.Base.Type = DeployCloudResourceType
	d.Properties = properties
	d.Base.SubSteps = subSteps
	return d, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	d := &DeployCloudResourceWorkflowStep{}
	return d.FromWorkflowStep(from)
}

func (d *DeployCloudResourceWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*DeployCloudResourceWorkflowStep, error) {
	var properties DeployCloudResourceSpec
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
	d.Base.Type = DeployCloudResourceType
	d.Properties = properties
	return d, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	d := &DeployCloudResourceWorkflowStep{}
	return d.FromWorkflowSubStep(from)
}

func (d *DeployCloudResourceWorkflowStep) WorkflowStepName() string {
	return d.Base.Name
}

func (d *DeployCloudResourceWorkflowStep) DefType() string {
	return DeployCloudResourceType
}

func (d *DeployCloudResourceWorkflowStep) If(_if string) *DeployCloudResourceWorkflowStep {
	d.Base.If = _if
	return d
}

func (d *DeployCloudResourceWorkflowStep) Alias(alias string) *DeployCloudResourceWorkflowStep {
	d.Base.Meta.Alias = alias
	return d
}

func (d *DeployCloudResourceWorkflowStep) Timeout(timeout string) *DeployCloudResourceWorkflowStep {
	d.Base.Timeout = timeout
	return d
}

func (d *DeployCloudResourceWorkflowStep) DependsOn(dependsOn []string) *DeployCloudResourceWorkflowStep {
	d.Base.DependsOn = dependsOn
	return d
}

func (d *DeployCloudResourceWorkflowStep) Inputs(input common.StepInputs) *DeployCloudResourceWorkflowStep {
	d.Base.Inputs = input
	return d
}

func (d *DeployCloudResourceWorkflowStep) Outputs(output common.StepOutputs) *DeployCloudResourceWorkflowStep {
	d.Base.Outputs = output
	return d
}
