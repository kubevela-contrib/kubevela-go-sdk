/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package install_kubernetes_api_datasource

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

// checks if the InstallKubernetesApiDatasourceSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InstallKubernetesApiDatasourceSpec{}

// InstallKubernetesApiDatasourceSpec struct for InstallKubernetesApiDatasourceSpec
type InstallKubernetesApiDatasourceSpec struct {
	Endpoint           *string `json:"endpoint"`
	Grafana            *string `json:"grafana"`
	Namespace          *string `json:"namespace"`
	ServiceAccountName *string `json:"serviceAccountName"`
	Uid                *string `json:"uid"`
}

// NewInstallKubernetesApiDatasourceSpecWith instantiates a new InstallKubernetesApiDatasourceSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewInstallKubernetesApiDatasourceSpecWith(endpoint string, grafana string, namespace string, serviceAccountName string, uid string) *InstallKubernetesApiDatasourceSpec {
	this := InstallKubernetesApiDatasourceSpec{}
	this.Endpoint = &endpoint
	this.Grafana = &grafana
	this.Namespace = &namespace
	this.ServiceAccountName = &serviceAccountName
	this.Uid = &uid
	return &this
}

// NewInstallKubernetesApiDatasourceSpecWithDefault instantiates a new InstallKubernetesApiDatasourceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallKubernetesApiDatasourceSpecWithDefault() *InstallKubernetesApiDatasourceSpec {
	this := InstallKubernetesApiDatasourceSpec{}
	var endpoint string = "https://kubernetes.default"
	this.Endpoint = &endpoint
	var grafana string = "default"
	this.Grafana = &grafana
	var namespace string = "o11y-system"
	this.Namespace = &namespace
	var serviceAccountName string = "grafana"
	this.ServiceAccountName = &serviceAccountName
	var uid string = "kubernetes-api"
	this.Uid = &uid
	return &this
}

// NewInstallKubernetesApiDatasourceSpec is short for NewInstallKubernetesApiDatasourceSpecWithDefault which instantiates a new InstallKubernetesApiDatasourceSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallKubernetesApiDatasourceSpec() *InstallKubernetesApiDatasourceSpec {
	return NewInstallKubernetesApiDatasourceSpecWithDefault()
}

// NewInstallKubernetesApiDatasourceSpecEmpty instantiates a new InstallKubernetesApiDatasourceSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewInstallKubernetesApiDatasourceSpecEmpty() *InstallKubernetesApiDatasourceSpec {
	this := InstallKubernetesApiDatasourceSpec{}
	return &this
}

// NewInstallKubernetesApiDatasourceSpecs converts a list InstallKubernetesApiDatasourceSpec pointers to objects.
// This is helpful when the SetInstallKubernetesApiDatasourceSpec requires a list of objects
func NewInstallKubernetesApiDatasourceSpecList(ps ...*InstallKubernetesApiDatasourceSpec) []InstallKubernetesApiDatasourceSpec {
	objs := []InstallKubernetesApiDatasourceSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this InstallKubernetesApiDatasourceSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *InstallKubernetesAPIDatasourceWorkflowStep) Validate() error {
	if o.Properties.Endpoint == nil {
		return errors.New("Endpoint in InstallKubernetesApiDatasourceSpec must be set")
	}
	if o.Properties.Grafana == nil {
		return errors.New("Grafana in InstallKubernetesApiDatasourceSpec must be set")
	}
	if o.Properties.Namespace == nil {
		return errors.New("Namespace in InstallKubernetesApiDatasourceSpec must be set")
	}
	if o.Properties.ServiceAccountName == nil {
		return errors.New("ServiceAccountName in InstallKubernetesApiDatasourceSpec must be set")
	}
	if o.Properties.Uid == nil {
		return errors.New("Uid in InstallKubernetesApiDatasourceSpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetEndpoint returns the Endpoint field value
func (o *InstallKubernetesAPIDatasourceWorkflowStep) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *InstallKubernetesAPIDatasourceWorkflowStep) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Endpoint, true
}

// SetEndpoint sets field value
func (o *InstallKubernetesAPIDatasourceWorkflowStep) SetEndpoint(v string) *InstallKubernetesAPIDatasourceWorkflowStep {
	o.Properties.Endpoint = &v
	return o
}

// GetGrafana returns the Grafana field value
func (o *InstallKubernetesAPIDatasourceWorkflowStep) GetGrafana() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Grafana
}

// GetGrafanaOk returns a tuple with the Grafana field value
// and a boolean to check if the value has been set.
func (o *InstallKubernetesAPIDatasourceWorkflowStep) GetGrafanaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Grafana, true
}

// SetGrafana sets field value
func (o *InstallKubernetesAPIDatasourceWorkflowStep) SetGrafana(v string) *InstallKubernetesAPIDatasourceWorkflowStep {
	o.Properties.Grafana = &v
	return o
}

// GetNamespace returns the Namespace field value
func (o *InstallKubernetesAPIDatasourceWorkflowStep) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *InstallKubernetesAPIDatasourceWorkflowStep) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Namespace, true
}

// SetNamespace sets field value
func (o *InstallKubernetesAPIDatasourceWorkflowStep) SetNamespace(v string) *InstallKubernetesAPIDatasourceWorkflowStep {
	o.Properties.Namespace = &v
	return o
}

// GetServiceAccountName returns the ServiceAccountName field value
func (o *InstallKubernetesAPIDatasourceWorkflowStep) GetServiceAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.ServiceAccountName
}

// GetServiceAccountNameOk returns a tuple with the ServiceAccountName field value
// and a boolean to check if the value has been set.
func (o *InstallKubernetesAPIDatasourceWorkflowStep) GetServiceAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.ServiceAccountName, true
}

// SetServiceAccountName sets field value
func (o *InstallKubernetesAPIDatasourceWorkflowStep) SetServiceAccountName(v string) *InstallKubernetesAPIDatasourceWorkflowStep {
	o.Properties.ServiceAccountName = &v
	return o
}

// GetUid returns the Uid field value
func (o *InstallKubernetesAPIDatasourceWorkflowStep) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *InstallKubernetesAPIDatasourceWorkflowStep) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Uid, true
}

// SetUid sets field value
func (o *InstallKubernetesAPIDatasourceWorkflowStep) SetUid(v string) *InstallKubernetesAPIDatasourceWorkflowStep {
	o.Properties.Uid = &v
	return o
}

func (o InstallKubernetesApiDatasourceSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstallKubernetesApiDatasourceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["grafana"] = o.Grafana
	toSerialize["namespace"] = o.Namespace
	toSerialize["serviceAccountName"] = o.ServiceAccountName
	toSerialize["uid"] = o.Uid
	return toSerialize, nil
}

type NullableInstallKubernetesApiDatasourceSpec struct {
	value *InstallKubernetesApiDatasourceSpec
	isSet bool
}

func (v *NullableInstallKubernetesApiDatasourceSpec) Get() *InstallKubernetesApiDatasourceSpec {
	return v.value
}

func (v *NullableInstallKubernetesApiDatasourceSpec) Set(val *InstallKubernetesApiDatasourceSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableInstallKubernetesApiDatasourceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallKubernetesApiDatasourceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallKubernetesApiDatasourceSpec(val *InstallKubernetesApiDatasourceSpec) *NullableInstallKubernetesApiDatasourceSpec {
	return &NullableInstallKubernetesApiDatasourceSpec{value: val, isSet: true}
}

func (v NullableInstallKubernetesApiDatasourceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallKubernetesApiDatasourceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const InstallKubernetesApiDatasourceType = "install-kubernetes-api-datasource"

func init() {
	sdkcommon.RegisterWorkflowStep(InstallKubernetesApiDatasourceType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(InstallKubernetesApiDatasourceType, FromWorkflowSubStep)
}

type InstallKubernetesAPIDatasourceWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties InstallKubernetesApiDatasourceSpec
}

func InstallKubernetesApiDatasource(name string) *InstallKubernetesAPIDatasourceWorkflowStep {
	i := &InstallKubernetesAPIDatasourceWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: InstallKubernetesApiDatasourceType,
	}}
	return i
}

func (i *InstallKubernetesAPIDatasourceWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range i.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  i.Base.DependsOn,
		If:         i.Base.If,
		Inputs:     i.Base.Inputs,
		Meta:       i.Base.Meta,
		Name:       i.Base.Name,
		Outputs:    i.Base.Outputs,
		Properties: util.Object2RawExtension(i.Properties),
		SubSteps:   subSteps,
		Timeout:    i.Base.Timeout,
		Type:       InstallKubernetesApiDatasourceType,
	}
	return res
}

func (i *InstallKubernetesAPIDatasourceWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*InstallKubernetesAPIDatasourceWorkflowStep, error) {
	var properties InstallKubernetesApiDatasourceSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := i.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	i.Base.Name = from.Name
	i.Base.DependsOn = from.DependsOn
	i.Base.Inputs = from.Inputs
	i.Base.Outputs = from.Outputs
	i.Base.If = from.If
	i.Base.Timeout = from.Timeout
	i.Base.Meta = from.Meta
	i.Base.Type = InstallKubernetesApiDatasourceType
	i.Properties = properties
	i.Base.SubSteps = subSteps
	return i, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	i := &InstallKubernetesAPIDatasourceWorkflowStep{}
	return i.FromWorkflowStep(from)
}

func (i *InstallKubernetesAPIDatasourceWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*InstallKubernetesAPIDatasourceWorkflowStep, error) {
	var properties InstallKubernetesApiDatasourceSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	i.Base.Name = from.Name
	i.Base.DependsOn = from.DependsOn
	i.Base.Inputs = from.Inputs
	i.Base.Outputs = from.Outputs
	i.Base.If = from.If
	i.Base.Timeout = from.Timeout
	i.Base.Meta = from.Meta
	i.Base.Type = InstallKubernetesApiDatasourceType
	i.Properties = properties
	return i, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	i := &InstallKubernetesAPIDatasourceWorkflowStep{}
	return i.FromWorkflowSubStep(from)
}

func (i *InstallKubernetesAPIDatasourceWorkflowStep) WorkflowStepName() string {
	return i.Base.Name
}

func (i *InstallKubernetesAPIDatasourceWorkflowStep) DefType() string {
	return InstallKubernetesApiDatasourceType
}

func (i *InstallKubernetesAPIDatasourceWorkflowStep) If(_if string) *InstallKubernetesAPIDatasourceWorkflowStep {
	i.Base.If = _if
	return i
}

func (i *InstallKubernetesAPIDatasourceWorkflowStep) Alias(alias string) *InstallKubernetesAPIDatasourceWorkflowStep {
	i.Base.Meta.Alias = alias
	return i
}

func (i *InstallKubernetesAPIDatasourceWorkflowStep) Timeout(timeout string) *InstallKubernetesAPIDatasourceWorkflowStep {
	i.Base.Timeout = timeout
	return i
}

func (i *InstallKubernetesAPIDatasourceWorkflowStep) DependsOn(dependsOn []string) *InstallKubernetesAPIDatasourceWorkflowStep {
	i.Base.DependsOn = dependsOn
	return i
}

func (i *InstallKubernetesAPIDatasourceWorkflowStep) Inputs(input common.StepInputs) *InstallKubernetesAPIDatasourceWorkflowStep {
	i.Base.Inputs = input
	return i
}

func (i *InstallKubernetesAPIDatasourceWorkflowStep) Outputs(output common.StepOutputs) *InstallKubernetesAPIDatasourceWorkflowStep {
	i.Base.Outputs = output
	return i
}
