/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gateway

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

// checks if the GatewaySpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GatewaySpec{}

// GatewaySpec struct for GatewaySpec
type GatewaySpec struct {
	// Specify the class of ingress to use
	Class *string `json:"class"`
	// Set ingress class in '.spec.ingressClassName' instead of 'kubernetes.io/ingress.class' annotation.
	ClassInSpec *bool `json:"classInSpec"`
	// Specify the domain you want to expose
	Domain *string `json:"domain,omitempty"`
	// Specify the host of the ingress gateway, which is used to generate the endpoints when the host is empty.
	GatewayHost *string `json:"gatewayHost,omitempty"`
	// Specify the mapping relationship between the http path and the workload port
	Http map[string]int32 `json:"http"`
	// Specify the secret name you want to quote to use tls.
	SecretName *string `json:"secretName,omitempty"`
}

// NewGatewaySpecWith instantiates a new GatewaySpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewGatewaySpecWith(class string, classInSpec bool, http map[string]int32) *GatewaySpec {
	this := GatewaySpec{}
	this.Class = &class
	this.ClassInSpec = &classInSpec
	this.Http = http
	return &this
}

// NewGatewaySpecWithDefault instantiates a new GatewaySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewaySpecWithDefault() *GatewaySpec {
	this := GatewaySpec{}
	var class string = "nginx"
	this.Class = &class
	var classInSpec bool = false
	this.ClassInSpec = &classInSpec
	return &this
}

// NewGatewaySpec is short for NewGatewaySpecWithDefault which instantiates a new GatewaySpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewaySpec() *GatewaySpec {
	return NewGatewaySpecWithDefault()
}

// NewGatewaySpecEmpty instantiates a new GatewaySpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewGatewaySpecEmpty() *GatewaySpec {
	this := GatewaySpec{}
	return &this
}

// NewGatewaySpecs converts a list GatewaySpec pointers to objects.
// This is helpful when the SetGatewaySpec requires a list of objects
func NewGatewaySpecList(ps ...*GatewaySpec) []GatewaySpec {
	objs := []GatewaySpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this GatewaySpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *GatewayTrait) Validate() error {
	if o.Properties.Class == nil {
		return errors.New("Class in GatewaySpec must be set")
	}
	if o.Properties.ClassInSpec == nil {
		return errors.New("ClassInSpec in GatewaySpec must be set")
	}
	if o.Properties.Http == nil {
		return errors.New("Http in GatewaySpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetClass returns the Class field value
func (o *GatewayTrait) GetClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Class
}

// GetClassOk returns a tuple with the Class field value
// and a boolean to check if the value has been set.
func (o *GatewayTrait) GetClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Class, true
}

// SetClass sets field value
func (o *GatewayTrait) SetClass(v string) *GatewayTrait {
	o.Properties.Class = &v
	return o
}

// GetClassInSpec returns the ClassInSpec field value
func (o *GatewayTrait) GetClassInSpec() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return *o.Properties.ClassInSpec
}

// GetClassInSpecOk returns a tuple with the ClassInSpec field value
// and a boolean to check if the value has been set.
func (o *GatewayTrait) GetClassInSpecOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.ClassInSpec, true
}

// SetClassInSpec sets field value
func (o *GatewayTrait) SetClassInSpec(v bool) *GatewayTrait {
	o.Properties.ClassInSpec = &v
	return o
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *GatewayTrait) GetDomain() string {
	if o == nil || utils.IsNil(o.Properties.Domain) {
		var ret string
		return ret
	}
	return *o.Properties.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTrait) GetDomainOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Domain) {
		return nil, false
	}
	return o.Properties.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *GatewayTrait) HasDomain() bool {
	if o != nil && !utils.IsNil(o.Properties.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the domain field.
// Domain:  Specify the domain you want to expose
func (o *GatewayTrait) SetDomain(v string) *GatewayTrait {
	o.Properties.Domain = &v
	return o
}

// GetGatewayHost returns the GatewayHost field value if set, zero value otherwise.
func (o *GatewayTrait) GetGatewayHost() string {
	if o == nil || utils.IsNil(o.Properties.GatewayHost) {
		var ret string
		return ret
	}
	return *o.Properties.GatewayHost
}

// GetGatewayHostOk returns a tuple with the GatewayHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTrait) GetGatewayHostOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.GatewayHost) {
		return nil, false
	}
	return o.Properties.GatewayHost, true
}

// HasGatewayHost returns a boolean if a field has been set.
func (o *GatewayTrait) HasGatewayHost() bool {
	if o != nil && !utils.IsNil(o.Properties.GatewayHost) {
		return true
	}

	return false
}

// SetGatewayHost gets a reference to the given string and assigns it to the gatewayHost field.
// GatewayHost:  Specify the host of the ingress gateway, which is used to generate the endpoints when the host is empty.
func (o *GatewayTrait) SetGatewayHost(v string) *GatewayTrait {
	o.Properties.GatewayHost = &v
	return o
}

// GetHttp returns the Http field value
func (o *GatewayTrait) GetHttp() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.Properties.Http
}

// GetHttpOk returns a tuple with the Http field value
// and a boolean to check if the value has been set.
func (o *GatewayTrait) GetHttpOk() (map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Http, true
}

// SetHttp sets field value
func (o *GatewayTrait) SetHttp(v map[string]int32) *GatewayTrait {
	o.Properties.Http = v
	return o
}

// GetSecretName returns the SecretName field value if set, zero value otherwise.
func (o *GatewayTrait) GetSecretName() string {
	if o == nil || utils.IsNil(o.Properties.SecretName) {
		var ret string
		return ret
	}
	return *o.Properties.SecretName
}

// GetSecretNameOk returns a tuple with the SecretName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTrait) GetSecretNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.SecretName) {
		return nil, false
	}
	return o.Properties.SecretName, true
}

// HasSecretName returns a boolean if a field has been set.
func (o *GatewayTrait) HasSecretName() bool {
	if o != nil && !utils.IsNil(o.Properties.SecretName) {
		return true
	}

	return false
}

// SetSecretName gets a reference to the given string and assigns it to the secretName field.
// SecretName:  Specify the secret name you want to quote to use tls.
func (o *GatewayTrait) SetSecretName(v string) *GatewayTrait {
	o.Properties.SecretName = &v
	return o
}

func (o GatewaySpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewaySpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["class"] = o.Class
	toSerialize["classInSpec"] = o.ClassInSpec
	if !utils.IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !utils.IsNil(o.GatewayHost) {
		toSerialize["gatewayHost"] = o.GatewayHost
	}
	toSerialize["http"] = o.Http
	if !utils.IsNil(o.SecretName) {
		toSerialize["secretName"] = o.SecretName
	}
	return toSerialize, nil
}

type NullableGatewaySpec struct {
	value *GatewaySpec
	isSet bool
}

func (v *NullableGatewaySpec) Get() *GatewaySpec {
	return v.value
}

func (v *NullableGatewaySpec) Set(val *GatewaySpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableGatewaySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewaySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewaySpec(val *GatewaySpec) *NullableGatewaySpec {
	return &NullableGatewaySpec{value: val, isSet: true}
}

func (v NullableGatewaySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewaySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const GatewayType = "gateway"

func init() {
	sdkcommon.RegisterTrait(GatewayType, FromTrait)
}

type GatewayTrait struct {
	Base       apis.TraitBase
	Properties GatewaySpec
}

func Gateway() *GatewayTrait {
	g := &GatewayTrait{Base: apis.TraitBase{}}
	return g
}

func (g *GatewayTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(g.Properties),
		Type:       GatewayType,
	}
	return res
}

func (g *GatewayTrait) FromTrait(from common.ApplicationTrait) (*GatewayTrait, error) {
	var properties GatewaySpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	g.Base.Type = GatewayType
	g.Properties = properties
	return g, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	g := &GatewayTrait{}
	return g.FromTrait(from)
}

func (g *GatewayTrait) DefType() string {
	return GatewayType
}
