/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grafana_datasource

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the GrafanaDatasourceSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GrafanaDatasourceSpec{}

// GrafanaDatasourceSpec struct for GrafanaDatasourceSpec
type GrafanaDatasourceSpec struct {
	// The json model of the grafana datasource
	Data *string `json:"data"`
	// The name of the grafana instance.
	Grafana *string `json:"grafana"`
	// The uid of the grafana datasource.
	Uid *string `json:"uid"`
}

// NewGrafanaDatasourceSpecWith instantiates a new GrafanaDatasourceSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewGrafanaDatasourceSpecWith(data string, grafana string, uid string) *GrafanaDatasourceSpec {
	this := GrafanaDatasourceSpec{}
	this.Data = &data
	this.Grafana = &grafana
	this.Uid = &uid
	return &this
}

// NewGrafanaDatasourceSpecWithDefault instantiates a new GrafanaDatasourceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrafanaDatasourceSpecWithDefault() *GrafanaDatasourceSpec {
	this := GrafanaDatasourceSpec{}
	var grafana string = "default"
	this.Grafana = &grafana
	return &this
}

// NewGrafanaDatasourceSpec is short for NewGrafanaDatasourceSpecWithDefault which instantiates a new GrafanaDatasourceSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrafanaDatasourceSpec() *GrafanaDatasourceSpec {
	return NewGrafanaDatasourceSpecWithDefault()
}

// NewGrafanaDatasourceSpecEmpty instantiates a new GrafanaDatasourceSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewGrafanaDatasourceSpecEmpty() *GrafanaDatasourceSpec {
	this := GrafanaDatasourceSpec{}
	return &this
}

// NewGrafanaDatasourceSpecs converts a list GrafanaDatasourceSpec pointers to objects.
// This is helpful when the SetGrafanaDatasourceSpec requires a list of objects
func NewGrafanaDatasourceSpecList(ps ...*GrafanaDatasourceSpec) []GrafanaDatasourceSpec {
	objs := []GrafanaDatasourceSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this GrafanaDatasourceSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *GrafanaDatasourceComponent) Validate() error {
	if o.Properties.Data == nil {
		return errors.New("Data in GrafanaDatasourceSpec must be set")
	}
	if o.Properties.Grafana == nil {
		return errors.New("Grafana in GrafanaDatasourceSpec must be set")
	}
	if o.Properties.Uid == nil {
		return errors.New("Uid in GrafanaDatasourceSpec must be set")
	}
	// validate all nested properties

	for i, v := range o.Base.Traits {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("traits[%d] %s in %s component is invalid: %w", i, v.DefType(), GrafanaDatasourceType, err)
		}
	}
	return nil
}

// GetData returns the Data field value
func (o *GrafanaDatasourceComponent) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GrafanaDatasourceComponent) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Data, true
}

// SetData sets field value
func (o *GrafanaDatasourceComponent) SetData(v string) *GrafanaDatasourceComponent {
	o.Properties.Data = &v
	return o
}

// GetGrafana returns the Grafana field value
func (o *GrafanaDatasourceComponent) GetGrafana() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Grafana
}

// GetGrafanaOk returns a tuple with the Grafana field value
// and a boolean to check if the value has been set.
func (o *GrafanaDatasourceComponent) GetGrafanaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Grafana, true
}

// SetGrafana sets field value
func (o *GrafanaDatasourceComponent) SetGrafana(v string) *GrafanaDatasourceComponent {
	o.Properties.Grafana = &v
	return o
}

// GetUid returns the Uid field value
func (o *GrafanaDatasourceComponent) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *GrafanaDatasourceComponent) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Uid, true
}

// SetUid sets field value
func (o *GrafanaDatasourceComponent) SetUid(v string) *GrafanaDatasourceComponent {
	o.Properties.Uid = &v
	return o
}

func (o GrafanaDatasourceSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrafanaDatasourceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["grafana"] = o.Grafana
	toSerialize["uid"] = o.Uid
	return toSerialize, nil
}

type NullableGrafanaDatasourceSpec struct {
	value *GrafanaDatasourceSpec
	isSet bool
}

func (v *NullableGrafanaDatasourceSpec) Get() *GrafanaDatasourceSpec {
	return v.value
}

func (v *NullableGrafanaDatasourceSpec) Set(val *GrafanaDatasourceSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableGrafanaDatasourceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGrafanaDatasourceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrafanaDatasourceSpec(val *GrafanaDatasourceSpec) *NullableGrafanaDatasourceSpec {
	return &NullableGrafanaDatasourceSpec{value: val, isSet: true}
}

func (v NullableGrafanaDatasourceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrafanaDatasourceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const GrafanaDatasourceType = "grafana-datasource"

func init() {
	sdkcommon.RegisterComponent(GrafanaDatasourceType, FromComponent)
}

type GrafanaDatasourceComponent struct {
	Base       apis.ComponentBase
	Properties GrafanaDatasourceSpec
}

func GrafanaDatasource(name string) *GrafanaDatasourceComponent {
	g := &GrafanaDatasourceComponent{Base: apis.ComponentBase{
		Name: name,
		Type: GrafanaDatasourceType,
	}}
	return g
}

func (g *GrafanaDatasourceComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range g.Base.Traits {
		traits = append(traits, trait.Build())
	}
	res := common.ApplicationComponent{
		DependsOn:  g.Base.DependsOn,
		Inputs:     g.Base.Inputs,
		Name:       g.Base.Name,
		Outputs:    g.Base.Outputs,
		Properties: util.Object2RawExtension(g.Properties),
		Traits:     traits,
		Type:       GrafanaDatasourceType,
	}
	return res
}

func (g *GrafanaDatasourceComponent) FromComponent(from common.ApplicationComponent) (*GrafanaDatasourceComponent, error) {
	for _, trait := range from.Traits {
		_t, err := sdkcommon.FromTrait(&trait)
		if err != nil {
			return nil, err
		}
		g.Base.Traits = append(g.Base.Traits, _t)
	}
	var properties GrafanaDatasourceSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	g.Base.Name = from.Name
	g.Base.DependsOn = from.DependsOn
	g.Base.Inputs = from.Inputs
	g.Base.Outputs = from.Outputs
	g.Base.Type = GrafanaDatasourceType
	g.Properties = properties
	return g, nil
}

func FromComponent(from common.ApplicationComponent) (apis.Component, error) {
	g := &GrafanaDatasourceComponent{}
	return g.FromComponent(from)
}

func (g *GrafanaDatasourceComponent) SetTraits(traits ...apis.Trait) *GrafanaDatasourceComponent {
	for _, addTrait := range traits {
		found := false
		for i, _t := range g.Base.Traits {
			if _t.DefType() == addTrait.DefType() {
				g.Base.Traits[i] = addTrait
				found = true
				break
			}
			if !found {
				g.Base.Traits = append(g.Base.Traits, addTrait)
			}
		}
	}
	return g
}

func (g *GrafanaDatasourceComponent) GetTrait(typ string) apis.Trait {
	for _, _t := range g.Base.Traits {
		if _t.DefType() == typ {
			return _t
		}
	}
	return nil
}

func (g *GrafanaDatasourceComponent) GetAllTraits() []apis.Trait {
	return g.Base.Traits
}

func (g *GrafanaDatasourceComponent) ComponentName() string {
	return g.Base.Name
}

func (g *GrafanaDatasourceComponent) DefType() string {
	return GrafanaDatasourceType
}

func (g *GrafanaDatasourceComponent) DependsOn(dependsOn []string) *GrafanaDatasourceComponent {
	g.Base.DependsOn = dependsOn
	return g
}

func (g *GrafanaDatasourceComponent) Inputs(input common.StepInputs) *GrafanaDatasourceComponent {
	g.Base.Inputs = input
	return g
}

func (g *GrafanaDatasourceComponent) Outputs(output common.StepOutputs) *GrafanaDatasourceComponent {
	g.Base.Outputs = output
	return g
}

func (g *GrafanaDatasourceComponent) AddDependsOn(dependsOn string) *GrafanaDatasourceComponent {
	g.Base.DependsOn = append(g.Base.DependsOn, dependsOn)
	return g
}
