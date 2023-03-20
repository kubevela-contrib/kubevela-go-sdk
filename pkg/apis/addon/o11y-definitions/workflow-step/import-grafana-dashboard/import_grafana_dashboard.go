/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package import_grafana_dashboard

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

// checks if the ImportGrafanaDashboardSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ImportGrafanaDashboardSpec{}

// ImportGrafanaDashboardSpec struct for ImportGrafanaDashboardSpec
type ImportGrafanaDashboardSpec struct {
	// The json model of the grafana dashboard.
	Data *string `json:"data,omitempty"`
	// The name of the grafana instance.
	Grafana *string `json:"grafana"`
	// The title of the grafana dashboard, if not specified, the default title will be used.
	Title *string `json:"title,omitempty"`
	// The uid of the grafana dashboard, if not specified, the default uid will be used.
	Uid *string `json:"uid,omitempty"`
	// The url link of the grafana dashboard.
	Url *string `json:"url,omitempty"`
}

// NewImportGrafanaDashboardSpecWith instantiates a new ImportGrafanaDashboardSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewImportGrafanaDashboardSpecWith(grafana string) *ImportGrafanaDashboardSpec {
	this := ImportGrafanaDashboardSpec{}
	this.Grafana = &grafana
	return &this
}

// NewImportGrafanaDashboardSpecWithDefault instantiates a new ImportGrafanaDashboardSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportGrafanaDashboardSpecWithDefault() *ImportGrafanaDashboardSpec {
	this := ImportGrafanaDashboardSpec{}
	var grafana string = "default"
	this.Grafana = &grafana
	return &this
}

// NewImportGrafanaDashboardSpec is short for NewImportGrafanaDashboardSpecWithDefault which instantiates a new ImportGrafanaDashboardSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportGrafanaDashboardSpec() *ImportGrafanaDashboardSpec {
	return NewImportGrafanaDashboardSpecWithDefault()
}

// NewImportGrafanaDashboardSpecEmpty instantiates a new ImportGrafanaDashboardSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewImportGrafanaDashboardSpecEmpty() *ImportGrafanaDashboardSpec {
	this := ImportGrafanaDashboardSpec{}
	return &this
}

// NewImportGrafanaDashboardSpecs converts a list ImportGrafanaDashboardSpec pointers to objects.
// This is helpful when the SetImportGrafanaDashboardSpec requires a list of objects
func NewImportGrafanaDashboardSpecList(ps ...*ImportGrafanaDashboardSpec) []ImportGrafanaDashboardSpec {
	objs := []ImportGrafanaDashboardSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this ImportGrafanaDashboardSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *ImportGrafanaDashboardWorkflowStep) Validate() error {
	if o.Properties.Grafana == nil {
		return errors.New("Grafana in ImportGrafanaDashboardSpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ImportGrafanaDashboardWorkflowStep) GetData() string {
	if o == nil || utils.IsNil(o.Properties.Data) {
		var ret string
		return ret
	}
	return *o.Properties.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportGrafanaDashboardWorkflowStep) GetDataOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Data) {
		return nil, false
	}
	return o.Properties.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ImportGrafanaDashboardWorkflowStep) HasData() bool {
	if o != nil && !utils.IsNil(o.Properties.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the data field.
// Data:  The json model of the grafana dashboard.
func (o *ImportGrafanaDashboardWorkflowStep) SetData(v string) *ImportGrafanaDashboardWorkflowStep {
	o.Properties.Data = &v
	return o
}

// GetGrafana returns the Grafana field value
func (o *ImportGrafanaDashboardWorkflowStep) GetGrafana() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Grafana
}

// GetGrafanaOk returns a tuple with the Grafana field value
// and a boolean to check if the value has been set.
func (o *ImportGrafanaDashboardWorkflowStep) GetGrafanaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Grafana, true
}

// SetGrafana sets field value
func (o *ImportGrafanaDashboardWorkflowStep) SetGrafana(v string) *ImportGrafanaDashboardWorkflowStep {
	o.Properties.Grafana = &v
	return o
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ImportGrafanaDashboardWorkflowStep) GetTitle() string {
	if o == nil || utils.IsNil(o.Properties.Title) {
		var ret string
		return ret
	}
	return *o.Properties.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportGrafanaDashboardWorkflowStep) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Title) {
		return nil, false
	}
	return o.Properties.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ImportGrafanaDashboardWorkflowStep) HasTitle() bool {
	if o != nil && !utils.IsNil(o.Properties.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the title field.
// Title:  The title of the grafana dashboard, if not specified, the default title will be used.
func (o *ImportGrafanaDashboardWorkflowStep) SetTitle(v string) *ImportGrafanaDashboardWorkflowStep {
	o.Properties.Title = &v
	return o
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *ImportGrafanaDashboardWorkflowStep) GetUid() string {
	if o == nil || utils.IsNil(o.Properties.Uid) {
		var ret string
		return ret
	}
	return *o.Properties.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportGrafanaDashboardWorkflowStep) GetUidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Uid) {
		return nil, false
	}
	return o.Properties.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *ImportGrafanaDashboardWorkflowStep) HasUid() bool {
	if o != nil && !utils.IsNil(o.Properties.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the uid field.
// Uid:  The uid of the grafana dashboard, if not specified, the default uid will be used.
func (o *ImportGrafanaDashboardWorkflowStep) SetUid(v string) *ImportGrafanaDashboardWorkflowStep {
	o.Properties.Uid = &v
	return o
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ImportGrafanaDashboardWorkflowStep) GetUrl() string {
	if o == nil || utils.IsNil(o.Properties.Url) {
		var ret string
		return ret
	}
	return *o.Properties.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportGrafanaDashboardWorkflowStep) GetUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Url) {
		return nil, false
	}
	return o.Properties.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ImportGrafanaDashboardWorkflowStep) HasUrl() bool {
	if o != nil && !utils.IsNil(o.Properties.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the url field.
// Url:  The url link of the grafana dashboard.
func (o *ImportGrafanaDashboardWorkflowStep) SetUrl(v string) *ImportGrafanaDashboardWorkflowStep {
	o.Properties.Url = &v
	return o
}

func (o ImportGrafanaDashboardSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportGrafanaDashboardSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["grafana"] = o.Grafana
	if !utils.IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !utils.IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !utils.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableImportGrafanaDashboardSpec struct {
	value *ImportGrafanaDashboardSpec
	isSet bool
}

func (v *NullableImportGrafanaDashboardSpec) Get() *ImportGrafanaDashboardSpec {
	return v.value
}

func (v *NullableImportGrafanaDashboardSpec) Set(val *ImportGrafanaDashboardSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableImportGrafanaDashboardSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableImportGrafanaDashboardSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportGrafanaDashboardSpec(val *ImportGrafanaDashboardSpec) *NullableImportGrafanaDashboardSpec {
	return &NullableImportGrafanaDashboardSpec{value: val, isSet: true}
}

func (v NullableImportGrafanaDashboardSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportGrafanaDashboardSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ImportGrafanaDashboardType = "import-grafana-dashboard"

func init() {
	sdkcommon.RegisterWorkflowStep(ImportGrafanaDashboardType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(ImportGrafanaDashboardType, FromWorkflowSubStep)
}

type ImportGrafanaDashboardWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties ImportGrafanaDashboardSpec
}

func ImportGrafanaDashboard(name string) *ImportGrafanaDashboardWorkflowStep {
	i := &ImportGrafanaDashboardWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: ImportGrafanaDashboardType,
	}}
	return i
}

func (i *ImportGrafanaDashboardWorkflowStep) Build() v1beta1.WorkflowStep {
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
		Type:       ImportGrafanaDashboardType,
	}
	return res
}

func (i *ImportGrafanaDashboardWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*ImportGrafanaDashboardWorkflowStep, error) {
	var properties ImportGrafanaDashboardSpec
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
	i.Base.Type = ImportGrafanaDashboardType
	i.Properties = properties
	i.Base.SubSteps = subSteps
	return i, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	i := &ImportGrafanaDashboardWorkflowStep{}
	return i.FromWorkflowStep(from)
}

func (i *ImportGrafanaDashboardWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*ImportGrafanaDashboardWorkflowStep, error) {
	var properties ImportGrafanaDashboardSpec
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
	i.Base.Type = ImportGrafanaDashboardType
	i.Properties = properties
	return i, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	i := &ImportGrafanaDashboardWorkflowStep{}
	return i.FromWorkflowSubStep(from)
}

func (i *ImportGrafanaDashboardWorkflowStep) WorkflowStepName() string {
	return i.Base.Name
}

func (i *ImportGrafanaDashboardWorkflowStep) DefType() string {
	return ImportGrafanaDashboardType
}

func (i *ImportGrafanaDashboardWorkflowStep) If(_if string) *ImportGrafanaDashboardWorkflowStep {
	i.Base.If = _if
	return i
}

func (i *ImportGrafanaDashboardWorkflowStep) Alias(alias string) *ImportGrafanaDashboardWorkflowStep {
	i.Base.Meta.Alias = alias
	return i
}

func (i *ImportGrafanaDashboardWorkflowStep) Timeout(timeout string) *ImportGrafanaDashboardWorkflowStep {
	i.Base.Timeout = timeout
	return i
}

func (i *ImportGrafanaDashboardWorkflowStep) DependsOn(dependsOn []string) *ImportGrafanaDashboardWorkflowStep {
	i.Base.DependsOn = dependsOn
	return i
}

func (i *ImportGrafanaDashboardWorkflowStep) Inputs(input common.StepInputs) *ImportGrafanaDashboardWorkflowStep {
	i.Base.Inputs = input
	return i
}

func (i *ImportGrafanaDashboardWorkflowStep) Outputs(output common.StepOutputs) *ImportGrafanaDashboardWorkflowStep {
	i.Base.Outputs = output
	return i
}
