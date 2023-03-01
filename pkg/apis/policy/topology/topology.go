/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package topology

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/vela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/vela-go-sdk/pkg/apis/utils"
)

// checks if the TopologySpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TopologySpec{}

// TopologySpec struct for TopologySpec
type TopologySpec struct {
	// Ignore empty cluster error
	AllowEmpty *bool `json:"allowEmpty,omitempty"`
	// Specify the label selector for clusters
	ClusterLabelSelector *map[string]string `json:"clusterLabelSelector,omitempty"`
	// Deprecated: Use clusterLabelSelector instead.
	ClusterSelector *map[string]string `json:"clusterSelector,omitempty"`
	// Specify the names of the clusters to select.
	Clusters []string `json:"clusters,omitempty"`
	// Specify the target namespace to deploy in the selected clusters, default inherit the original namespace.
	Namespace *string `json:"namespace,omitempty"`
}

// NewTopologySpecWith instantiates a new TopologySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopologySpecWith() *TopologySpec {
	this := TopologySpec{}
	return &this
}

// NewTopologySpec instantiates a new TopologySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopologySpec() *TopologySpec {
	this := TopologySpec{}
	return &this
}

// NewTopologySpecs converts a list TopologySpec pointers to objects.
// This is helpful when the SetTopologySpec requires a list of objects
func NewTopologySpecList(ps ...*TopologySpec) []TopologySpec {
	objs := []TopologySpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetAllowEmpty returns the AllowEmpty field value if set, zero value otherwise.
func (o *TopologyPolicy) GetAllowEmpty() bool {
	if o == nil || utils.IsNil(o.Properties.AllowEmpty) {
		var ret bool
		return ret
	}
	return *o.Properties.AllowEmpty
}

// GetAllowEmptyOk returns a tuple with the AllowEmpty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyPolicy) GetAllowEmptyOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Properties.AllowEmpty) {
		return nil, false
	}
	return o.Properties.AllowEmpty, true
}

// HasAllowEmpty returns a boolean if a field has been set.
func (o *TopologyPolicy) HasAllowEmpty() bool {
	if o != nil && !utils.IsNil(o.Properties.AllowEmpty) {
		return true
	}

	return false
}

// SetAllowEmpty gets a reference to the given bool and assigns it to the allowEmpty field.
// AllowEmpty:  Ignore empty cluster error
func (o *TopologyPolicy) SetAllowEmpty(v bool) *TopologyPolicy {
	o.Properties.AllowEmpty = &v
	return o
}

// GetClusterLabelSelector returns the ClusterLabelSelector field value if set, zero value otherwise.
func (o *TopologyPolicy) GetClusterLabelSelector() map[string]string {
	if o == nil || utils.IsNil(o.Properties.ClusterLabelSelector) {
		var ret map[string]string
		return ret
	}
	return *o.Properties.ClusterLabelSelector
}

// GetClusterLabelSelectorOk returns a tuple with the ClusterLabelSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyPolicy) GetClusterLabelSelectorOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.Properties.ClusterLabelSelector) {
		return nil, false
	}
	return o.Properties.ClusterLabelSelector, true
}

// HasClusterLabelSelector returns a boolean if a field has been set.
func (o *TopologyPolicy) HasClusterLabelSelector() bool {
	if o != nil && !utils.IsNil(o.Properties.ClusterLabelSelector) {
		return true
	}

	return false
}

// SetClusterLabelSelector gets a reference to the given map[string]string and assigns it to the clusterLabelSelector field.
// ClusterLabelSelector:  Specify the label selector for clusters
func (o *TopologyPolicy) SetClusterLabelSelector(v map[string]string) *TopologyPolicy {
	o.Properties.ClusterLabelSelector = &v
	return o
}

// GetClusterSelector returns the ClusterSelector field value if set, zero value otherwise.
func (o *TopologyPolicy) GetClusterSelector() map[string]string {
	if o == nil || utils.IsNil(o.Properties.ClusterSelector) {
		var ret map[string]string
		return ret
	}
	return *o.Properties.ClusterSelector
}

// GetClusterSelectorOk returns a tuple with the ClusterSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyPolicy) GetClusterSelectorOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.Properties.ClusterSelector) {
		return nil, false
	}
	return o.Properties.ClusterSelector, true
}

// HasClusterSelector returns a boolean if a field has been set.
func (o *TopologyPolicy) HasClusterSelector() bool {
	if o != nil && !utils.IsNil(o.Properties.ClusterSelector) {
		return true
	}

	return false
}

// SetClusterSelector gets a reference to the given map[string]string and assigns it to the clusterSelector field.
// ClusterSelector:  Deprecated: Use clusterLabelSelector instead.
func (o *TopologyPolicy) SetClusterSelector(v map[string]string) *TopologyPolicy {
	o.Properties.ClusterSelector = &v
	return o
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *TopologyPolicy) GetClusters() []string {
	if o == nil || utils.IsNil(o.Properties.Clusters) {
		var ret []string
		return ret
	}
	return o.Properties.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyPolicy) GetClustersOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Clusters) {
		return nil, false
	}
	return o.Properties.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *TopologyPolicy) HasClusters() bool {
	if o != nil && !utils.IsNil(o.Properties.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []string and assigns it to the clusters field.
// Clusters:  Specify the names of the clusters to select.
func (o *TopologyPolicy) SetClusters(v []string) *TopologyPolicy {
	o.Properties.Clusters = v
	return o
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *TopologyPolicy) GetNamespace() string {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		var ret string
		return ret
	}
	return *o.Properties.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyPolicy) GetNamespaceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		return nil, false
	}
	return o.Properties.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *TopologyPolicy) HasNamespace() bool {
	if o != nil && !utils.IsNil(o.Properties.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the namespace field.
// Namespace:  Specify the target namespace to deploy in the selected clusters, default inherit the original namespace.
func (o *TopologyPolicy) SetNamespace(v string) *TopologyPolicy {
	o.Properties.Namespace = &v
	return o
}

func (o TopologySpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopologySpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AllowEmpty) {
		toSerialize["allowEmpty"] = o.AllowEmpty
	}
	if !utils.IsNil(o.ClusterLabelSelector) {
		toSerialize["clusterLabelSelector"] = o.ClusterLabelSelector
	}
	if !utils.IsNil(o.ClusterSelector) {
		toSerialize["clusterSelector"] = o.ClusterSelector
	}
	if !utils.IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	if !utils.IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	return toSerialize, nil
}

type NullableTopologySpec struct {
	value *TopologySpec
	isSet bool
}

func (v NullableTopologySpec) Get() *TopologySpec {
	return v.value
}

func (v *NullableTopologySpec) Set(val *TopologySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTopologySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTopologySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopologySpec(val *TopologySpec) *NullableTopologySpec {
	return &NullableTopologySpec{value: val, isSet: true}
}

func (v NullableTopologySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopologySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const TopologyType = "topology"

func init() {
	sdkcommon.RegisterPolicy(TopologyType, FromPolicy)
}

type TopologyPolicy struct {
	Base       apis.PolicyBase
	Properties TopologySpec
}

func Topology(name string) *TopologyPolicy {
	t := &TopologyPolicy{Base: apis.PolicyBase{
		Name: name,
		Type: TopologyType,
	}}
	return t
}

func (t *TopologyPolicy) Build() v1beta1.AppPolicy {
	res := v1beta1.AppPolicy{
		Name:       t.Base.Name,
		Properties: util.Object2RawExtension(t.Properties),
		Type:       TopologyType,
	}
	return res
}

func (t *TopologyPolicy) FromPolicy(from v1beta1.AppPolicy) (*TopologyPolicy, error) {
	var properties TopologySpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	t.Base.Name = from.Name
	t.Base.Type = TopologyType
	t.Properties = properties
	return t, nil
}

func FromPolicy(from v1beta1.AppPolicy) (apis.Policy, error) {
	t := &TopologyPolicy{}
	return t.FromPolicy(from)
}

func (t *TopologyPolicy) PolicyName() string {
	return t.Base.Name
}

func (t *TopologyPolicy) DefType() string {
	return TopologyType
}