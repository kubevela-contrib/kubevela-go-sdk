/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package read_object

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/vela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/vela-go-sdk/pkg/apis/utils"
)

// checks if the ReadObjectSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReadObjectSpec{}

// ReadObjectSpec struct for ReadObjectSpec
type ReadObjectSpec struct {
	// Specify the apiVersion of the object, defaults to 'core.oam.dev/v1beta1'
	ApiVersion *string `json:"apiVersion,omitempty"`
	// The cluster you want to apply the resource to, default is the current control plane cluster
	Cluster *string `json:"cluster,omitempty"`
	// Specify the kind of the object, defaults to Application
	Kind *string `json:"kind,omitempty"`
	// Specify the name of the object
	Name *string `json:"name,omitempty"`
	// The namespace of the resource you want to read
	Namespace *string `json:"namespace,omitempty"`
}

// NewReadObjectSpecWith instantiates a new ReadObjectSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadObjectSpecWith() *ReadObjectSpec {
	this := ReadObjectSpec{}
	var cluster string = ""
	this.Cluster = &cluster
	var namespace string = "default"
	this.Namespace = &namespace
	return &this
}

// NewReadObjectSpec instantiates a new ReadObjectSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadObjectSpec() *ReadObjectSpec {
	this := ReadObjectSpec{}
	var cluster string = ""
	this.Cluster = &cluster
	var namespace string = "default"
	this.Namespace = &namespace
	return &this
}

// NewReadObjectSpecs converts a list ReadObjectSpec pointers to objects.
// This is helpful when the SetReadObjectSpec requires a list of objects
func NewReadObjectSpecList(ps ...*ReadObjectSpec) []ReadObjectSpec {
	objs := []ReadObjectSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ReadObjectWorkflowStep) GetApiVersion() string {
	if o == nil || utils.IsNil(o.Properties.ApiVersion) {
		var ret string
		return ret
	}
	return *o.Properties.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadObjectWorkflowStep) GetApiVersionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.ApiVersion) {
		return nil, false
	}
	return o.Properties.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ReadObjectWorkflowStep) HasApiVersion() bool {
	if o != nil && !utils.IsNil(o.Properties.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the apiVersion field.
// ApiVersion:  Specify the apiVersion of the object, defaults to 'core.oam.dev/v1beta1'
func (o *ReadObjectWorkflowStep) SetApiVersion(v string) *ReadObjectWorkflowStep {
	o.Properties.ApiVersion = &v
	return o
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *ReadObjectWorkflowStep) GetCluster() string {
	if o == nil || utils.IsNil(o.Properties.Cluster) {
		var ret string
		return ret
	}
	return *o.Properties.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadObjectWorkflowStep) GetClusterOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Cluster) {
		return nil, false
	}
	return o.Properties.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ReadObjectWorkflowStep) HasCluster() bool {
	if o != nil && !utils.IsNil(o.Properties.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the cluster field.
// Cluster:  The cluster you want to apply the resource to, default is the current control plane cluster
func (o *ReadObjectWorkflowStep) SetCluster(v string) *ReadObjectWorkflowStep {
	o.Properties.Cluster = &v
	return o
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ReadObjectWorkflowStep) GetKind() string {
	if o == nil || utils.IsNil(o.Properties.Kind) {
		var ret string
		return ret
	}
	return *o.Properties.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadObjectWorkflowStep) GetKindOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Kind) {
		return nil, false
	}
	return o.Properties.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ReadObjectWorkflowStep) HasKind() bool {
	if o != nil && !utils.IsNil(o.Properties.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the kind field.
// Kind:  Specify the kind of the object, defaults to Application
func (o *ReadObjectWorkflowStep) SetKind(v string) *ReadObjectWorkflowStep {
	o.Properties.Kind = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReadObjectWorkflowStep) GetName() string {
	if o == nil || utils.IsNil(o.Properties.Name) {
		var ret string
		return ret
	}
	return *o.Properties.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadObjectWorkflowStep) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Name) {
		return nil, false
	}
	return o.Properties.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReadObjectWorkflowStep) HasName() bool {
	if o != nil && !utils.IsNil(o.Properties.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:  Specify the name of the object
func (o *ReadObjectWorkflowStep) SetName(v string) *ReadObjectWorkflowStep {
	o.Properties.Name = &v
	return o
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ReadObjectWorkflowStep) GetNamespace() string {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		var ret string
		return ret
	}
	return *o.Properties.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadObjectWorkflowStep) GetNamespaceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		return nil, false
	}
	return o.Properties.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ReadObjectWorkflowStep) HasNamespace() bool {
	if o != nil && !utils.IsNil(o.Properties.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the namespace field.
// Namespace:  The namespace of the resource you want to read
func (o *ReadObjectWorkflowStep) SetNamespace(v string) *ReadObjectWorkflowStep {
	o.Properties.Namespace = &v
	return o
}

func (o ReadObjectSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadObjectSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !utils.IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !utils.IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	return toSerialize, nil
}

type NullableReadObjectSpec struct {
	value *ReadObjectSpec
	isSet bool
}

func (v NullableReadObjectSpec) Get() *ReadObjectSpec {
	return v.value
}

func (v *NullableReadObjectSpec) Set(val *ReadObjectSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableReadObjectSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableReadObjectSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadObjectSpec(val *ReadObjectSpec) *NullableReadObjectSpec {
	return &NullableReadObjectSpec{value: val, isSet: true}
}

func (v NullableReadObjectSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadObjectSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ReadObjectType = "read-object"

func init() {
	sdkcommon.RegisterWorkflowStep(ReadObjectType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(ReadObjectType, FromWorkflowSubStep)
}

type ReadObjectWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties ReadObjectSpec
}

func ReadObject(name string) *ReadObjectWorkflowStep {
	r := &ReadObjectWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: ReadObjectType,
	}}
	return r
}

func (r *ReadObjectWorkflowStep) Build() v1beta1.WorkflowStep {
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
		Type:       ReadObjectType,
	}
	return res
}

func (r *ReadObjectWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*ReadObjectWorkflowStep, error) {
	var properties ReadObjectSpec
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
	r.Base.Type = ReadObjectType
	r.Properties = properties
	r.Base.SubSteps = subSteps
	return r, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	r := &ReadObjectWorkflowStep{}
	return r.FromWorkflowStep(from)
}

func (r *ReadObjectWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*ReadObjectWorkflowStep, error) {
	var properties ReadObjectSpec
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
	r.Base.Type = ReadObjectType
	r.Properties = properties
	return r, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	r := &ReadObjectWorkflowStep{}
	return r.FromWorkflowSubStep(from)
}

func (r *ReadObjectWorkflowStep) WorkflowStepName() string {
	return r.Base.Name
}

func (r *ReadObjectWorkflowStep) DefType() string {
	return ReadObjectType
}

func (r *ReadObjectWorkflowStep) If(_if string) *ReadObjectWorkflowStep {
	r.Base.If = _if
	return r
}

func (r *ReadObjectWorkflowStep) Alias(alias string) *ReadObjectWorkflowStep {
	r.Base.Meta.Alias = alias
	return r
}

func (r *ReadObjectWorkflowStep) Timeout(timeout string) *ReadObjectWorkflowStep {
	r.Base.Timeout = timeout
	return r
}

func (r *ReadObjectWorkflowStep) DependsOn(dependsOn []string) *ReadObjectWorkflowStep {
	r.Base.DependsOn = dependsOn
	return r
}

func (r *ReadObjectWorkflowStep) Inputs(input common.StepInputs) *ReadObjectWorkflowStep {
	r.Base.Inputs = input
	return r
}

func (r *ReadObjectWorkflowStep) Outputs(output common.StepOutputs) *ReadObjectWorkflowStep {
	r.Base.Outputs = output
	return r
}