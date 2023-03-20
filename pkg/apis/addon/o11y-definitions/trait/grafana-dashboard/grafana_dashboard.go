/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grafana_dashboard

import (
	"encoding/json"
	"errors"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the GrafanaDashboardSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GrafanaDashboardSpec{}

// GrafanaDashboardSpec struct for GrafanaDashboardSpec
type GrafanaDashboardSpec struct {
	// The json model of the grafana dashboard
	Data *string `json:"data"`
	// The name of the grafana instance.
	Grafana *string `json:"grafana"`
	// The title of the grafana dashboard, if not specified, the default title will be used.
	Title *string `json:"title,omitempty"`
	// The uid of the grafana dashboard, if not specified, the default uid will be used.
	Uid *string `json:"uid,omitempty"`
}

// NewGrafanaDashboardSpecWith instantiates a new GrafanaDashboardSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewGrafanaDashboardSpecWith(data string, grafana string) *GrafanaDashboardSpec {
	this := GrafanaDashboardSpec{}
	this.Data = &data
	this.Grafana = &grafana
	return &this
}

// NewGrafanaDashboardSpecWithDefault instantiates a new GrafanaDashboardSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrafanaDashboardSpecWithDefault() *GrafanaDashboardSpec {
	this := GrafanaDashboardSpec{}
	var grafana string = "default"
	this.Grafana = &grafana
	return &this
}

// NewGrafanaDashboardSpec is short for NewGrafanaDashboardSpecWithDefault which instantiates a new GrafanaDashboardSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrafanaDashboardSpec() *GrafanaDashboardSpec {
	return NewGrafanaDashboardSpecWithDefault()
}

// NewGrafanaDashboardSpecEmpty instantiates a new GrafanaDashboardSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewGrafanaDashboardSpecEmpty() *GrafanaDashboardSpec {
	this := GrafanaDashboardSpec{}
	return &this
}

// NewGrafanaDashboardSpecs converts a list GrafanaDashboardSpec pointers to objects.
// This is helpful when the SetGrafanaDashboardSpec requires a list of objects
func NewGrafanaDashboardSpecList(ps ...*GrafanaDashboardSpec) []GrafanaDashboardSpec {
	objs := []GrafanaDashboardSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this GrafanaDashboardSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *GrafanaDashboardTrait) Validate() error {
	if o.Properties.Data == nil {
		return errors.New("Data in GrafanaDashboardSpec must be set")
	}
	if o.Properties.Grafana == nil {
		return errors.New("Grafana in GrafanaDashboardSpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetData returns the Data field value
func (o *GrafanaDashboardTrait) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GrafanaDashboardTrait) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Data, true
}

// SetData sets field value
func (o *GrafanaDashboardTrait) SetData(v string) *GrafanaDashboardTrait {
	o.Properties.Data = &v
	return o
}

// GetGrafana returns the Grafana field value
func (o *GrafanaDashboardTrait) GetGrafana() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Grafana
}

// GetGrafanaOk returns a tuple with the Grafana field value
// and a boolean to check if the value has been set.
func (o *GrafanaDashboardTrait) GetGrafanaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Grafana, true
}

// SetGrafana sets field value
func (o *GrafanaDashboardTrait) SetGrafana(v string) *GrafanaDashboardTrait {
	o.Properties.Grafana = &v
	return o
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GrafanaDashboardTrait) GetTitle() string {
	if o == nil || utils.IsNil(o.Properties.Title) {
		var ret string
		return ret
	}
	return *o.Properties.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaDashboardTrait) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Title) {
		return nil, false
	}
	return o.Properties.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GrafanaDashboardTrait) HasTitle() bool {
	if o != nil && !utils.IsNil(o.Properties.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the title field.
// Title:  The title of the grafana dashboard, if not specified, the default title will be used.
func (o *GrafanaDashboardTrait) SetTitle(v string) *GrafanaDashboardTrait {
	o.Properties.Title = &v
	return o
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *GrafanaDashboardTrait) GetUid() string {
	if o == nil || utils.IsNil(o.Properties.Uid) {
		var ret string
		return ret
	}
	return *o.Properties.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaDashboardTrait) GetUidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Uid) {
		return nil, false
	}
	return o.Properties.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *GrafanaDashboardTrait) HasUid() bool {
	if o != nil && !utils.IsNil(o.Properties.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the uid field.
// Uid:  The uid of the grafana dashboard, if not specified, the default uid will be used.
func (o *GrafanaDashboardTrait) SetUid(v string) *GrafanaDashboardTrait {
	o.Properties.Uid = &v
	return o
}

func (o GrafanaDashboardSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrafanaDashboardSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["grafana"] = o.Grafana
	if !utils.IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !utils.IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	return toSerialize, nil
}

type NullableGrafanaDashboardSpec struct {
	value *GrafanaDashboardSpec
	isSet bool
}

func (v *NullableGrafanaDashboardSpec) Get() *GrafanaDashboardSpec {
	return v.value
}

func (v *NullableGrafanaDashboardSpec) Set(val *GrafanaDashboardSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableGrafanaDashboardSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGrafanaDashboardSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrafanaDashboardSpec(val *GrafanaDashboardSpec) *NullableGrafanaDashboardSpec {
	return &NullableGrafanaDashboardSpec{value: val, isSet: true}
}

func (v NullableGrafanaDashboardSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrafanaDashboardSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const GrafanaDashboardType = "grafana-dashboard"

func init() {
	sdkcommon.RegisterTrait(GrafanaDashboardType, FromTrait)
}

type GrafanaDashboardTrait struct {
	Base       apis.TraitBase
	Properties GrafanaDashboardSpec
}

func GrafanaDashboard() *GrafanaDashboardTrait {
	g := &GrafanaDashboardTrait{Base: apis.TraitBase{}}
	return g
}

func (g *GrafanaDashboardTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(g.Properties),
		Type:       GrafanaDashboardType,
	}
	return res
}

func (g *GrafanaDashboardTrait) FromTrait(from common.ApplicationTrait) (*GrafanaDashboardTrait, error) {
	var properties GrafanaDashboardSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	g.Base.Type = GrafanaDashboardType
	g.Properties = properties
	return g, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	g := &GrafanaDashboardTrait{}
	return g.FromTrait(from)
}

func (g *GrafanaDashboardTrait) DefType() string {
	return GrafanaDashboardType
}
