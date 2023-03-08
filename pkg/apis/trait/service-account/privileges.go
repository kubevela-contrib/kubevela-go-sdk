/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package service_account

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

// checks if the Privileges type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Privileges{}

// Privileges struct for Privileges
type Privileges struct {
	// Specify the apiGroups of the resource
	ApiGroups []string `json:"apiGroups,omitempty"`
	// Specify the resource url to be allowed
	NonResourceURLs []string `json:"nonResourceURLs,omitempty"`
	// Specify the resourceNames to be allowed
	ResourceNames []string `json:"resourceNames,omitempty"`
	// Specify the resources to be allowed
	Resources []string `json:"resources,omitempty"`
	// Specify the scope of the privileges, default to be namespace scope
	Scope *string `json:"scope"`
	// Specify the verbs to be allowed for the resource
	Verbs []string `json:"verbs"`
}

// NewPrivilegesWith instantiates a new Privileges object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewPrivilegesWith(scope string, verbs []string) *Privileges {
	this := Privileges{}
	this.Scope = &scope
	this.Verbs = verbs
	return &this
}

// NewPrivilegesWithDefault instantiates a new Privileges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegesWithDefault() *Privileges {
	this := Privileges{}
	var scope string = "namespace"
	this.Scope = &scope
	return &this
}

// NewPrivileges is short for NewPrivilegesWithDefault which instantiates a new Privileges object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivileges() *Privileges {
	return NewPrivilegesWithDefault()
}

// NewPrivilegesEmpty instantiates a new Privileges object with no properties set.
// This constructor will not assign any default values to properties.
func NewPrivilegesEmpty() *Privileges {
	this := Privileges{}
	return &this
}

// NewPrivilegess converts a list Privileges pointers to objects.
// This is helpful when the SetPrivileges requires a list of objects
func NewPrivilegesList(ps ...*Privileges) []Privileges {
	objs := []Privileges{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Privileges
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Privileges) Validate() error {
	if o.Scope == nil {
		return errors.New("Scope in Privileges must be set")
	}
	if o.Verbs == nil {
		return errors.New("Verbs in Privileges must be set")
	}
	// validate all nested properties
	return nil
}

// GetApiGroups returns the ApiGroups field value if set, zero value otherwise.
func (o *Privileges) GetApiGroups() []string {
	if o == nil || utils.IsNil(o.ApiGroups) {
		var ret []string
		return ret
	}
	return o.ApiGroups
}

// GetApiGroupsOk returns a tuple with the ApiGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privileges) GetApiGroupsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ApiGroups) {
		return nil, false
	}
	return o.ApiGroups, true
}

// HasApiGroups returns a boolean if a field has been set.
func (o *Privileges) HasApiGroups() bool {
	if o != nil && !utils.IsNil(o.ApiGroups) {
		return true
	}

	return false
}

// SetApiGroups gets a reference to the given []string and assigns it to the apiGroups field.
// ApiGroups:  Specify the apiGroups of the resource
func (o *Privileges) SetApiGroups(v []string) *Privileges {
	o.ApiGroups = v
	return o
}

// GetNonResourceURLs returns the NonResourceURLs field value if set, zero value otherwise.
func (o *Privileges) GetNonResourceURLs() []string {
	if o == nil || utils.IsNil(o.NonResourceURLs) {
		var ret []string
		return ret
	}
	return o.NonResourceURLs
}

// GetNonResourceURLsOk returns a tuple with the NonResourceURLs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privileges) GetNonResourceURLsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.NonResourceURLs) {
		return nil, false
	}
	return o.NonResourceURLs, true
}

// HasNonResourceURLs returns a boolean if a field has been set.
func (o *Privileges) HasNonResourceURLs() bool {
	if o != nil && !utils.IsNil(o.NonResourceURLs) {
		return true
	}

	return false
}

// SetNonResourceURLs gets a reference to the given []string and assigns it to the nonResourceURLs field.
// NonResourceURLs:  Specify the resource url to be allowed
func (o *Privileges) SetNonResourceURLs(v []string) *Privileges {
	o.NonResourceURLs = v
	return o
}

// GetResourceNames returns the ResourceNames field value if set, zero value otherwise.
func (o *Privileges) GetResourceNames() []string {
	if o == nil || utils.IsNil(o.ResourceNames) {
		var ret []string
		return ret
	}
	return o.ResourceNames
}

// GetResourceNamesOk returns a tuple with the ResourceNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privileges) GetResourceNamesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ResourceNames) {
		return nil, false
	}
	return o.ResourceNames, true
}

// HasResourceNames returns a boolean if a field has been set.
func (o *Privileges) HasResourceNames() bool {
	if o != nil && !utils.IsNil(o.ResourceNames) {
		return true
	}

	return false
}

// SetResourceNames gets a reference to the given []string and assigns it to the resourceNames field.
// ResourceNames:  Specify the resourceNames to be allowed
func (o *Privileges) SetResourceNames(v []string) *Privileges {
	o.ResourceNames = v
	return o
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *Privileges) GetResources() []string {
	if o == nil || utils.IsNil(o.Resources) {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privileges) GetResourcesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *Privileges) HasResources() bool {
	if o != nil && !utils.IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the resources field.
// Resources:  Specify the resources to be allowed
func (o *Privileges) SetResources(v []string) *Privileges {
	o.Resources = v
	return o
}

// GetScope returns the Scope field value
func (o *Privileges) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *Privileges) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope, true
}

// SetScope sets field value
func (o *Privileges) SetScope(v string) *Privileges {
	o.Scope = &v
	return o
}

// GetVerbs returns the Verbs field value
func (o *Privileges) GetVerbs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Verbs
}

// GetVerbsOk returns a tuple with the Verbs field value
// and a boolean to check if the value has been set.
func (o *Privileges) GetVerbsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verbs, true
}

// SetVerbs sets field value
func (o *Privileges) SetVerbs(v []string) *Privileges {
	o.Verbs = v
	return o
}

func (o Privileges) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Privileges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ApiGroups) {
		toSerialize["apiGroups"] = o.ApiGroups
	}
	if !utils.IsNil(o.NonResourceURLs) {
		toSerialize["nonResourceURLs"] = o.NonResourceURLs
	}
	if !utils.IsNil(o.ResourceNames) {
		toSerialize["resourceNames"] = o.ResourceNames
	}
	if !utils.IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	toSerialize["scope"] = o.Scope
	toSerialize["verbs"] = o.Verbs
	return toSerialize, nil
}

type NullablePrivileges struct {
	value *Privileges
	isSet bool
}

func (v *NullablePrivileges) Get() *Privileges {
	return v.value
}

func (v *NullablePrivileges) Set(val *Privileges) {
	v.value = val
	v.isSet = true
}

func (v *NullablePrivileges) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivileges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivileges(val *Privileges) *NullablePrivileges {
	return &NullablePrivileges{value: val, isSet: true}
}

func (v NullablePrivileges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivileges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
