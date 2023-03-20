/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netlify

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

// checks if the NetlifySpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NetlifySpec{}

// NetlifySpec struct for NetlifySpec
type NetlifySpec struct {
	Backofflimit *int32 `json:"backofflimit"`
	Completions  *int32 `json:"completions"`
	// Specify the repo address your resources of the frontend service  store
	Gitrepo *string `json:"gitrepo"`
	// Specify the direcotry name for saving the resources from git or other repo supporting git-cli
	Reponame      *string `json:"reponame"`
	RestartPolicy *string `json:"restartPolicy"`
	// Specify the website name on app.netlify.com
	Sitname *string `json:"sitname"`
	// Specify the token got from app.netlify.com
	Token *string `json:"token"`
}

// NewNetlifySpecWith instantiates a new NetlifySpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewNetlifySpecWith(backofflimit int32, completions int32, gitrepo string, reponame string, restartPolicy string, sitname string, token string) *NetlifySpec {
	this := NetlifySpec{}
	this.Backofflimit = &backofflimit
	this.Completions = &completions
	this.Gitrepo = &gitrepo
	this.Reponame = &reponame
	this.RestartPolicy = &restartPolicy
	this.Sitname = &sitname
	this.Token = &token
	return &this
}

// NewNetlifySpecWithDefault instantiates a new NetlifySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetlifySpecWithDefault() *NetlifySpec {
	this := NetlifySpec{}
	var backofflimit int32 = 5
	this.Backofflimit = &backofflimit
	var completions int32 = 1
	this.Completions = &completions
	var gitrepo string = "https://github.com/oeular/oeular.github.io.Properties.git"
	this.Gitrepo = &gitrepo
	var reponame string = "oeular.github.io"
	this.Reponame = &reponame
	var restartPolicy string = "OnFailure"
	this.RestartPolicy = &restartPolicy
	var sitname string = "deploy-from-vela"
	this.Sitname = &sitname
	var token string = "your-own-token-after-base64"
	this.Token = &token
	return &this
}

// NewNetlifySpec is short for NewNetlifySpecWithDefault which instantiates a new NetlifySpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetlifySpec() *NetlifySpec {
	return NewNetlifySpecWithDefault()
}

// NewNetlifySpecEmpty instantiates a new NetlifySpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewNetlifySpecEmpty() *NetlifySpec {
	this := NetlifySpec{}
	return &this
}

// NewNetlifySpecs converts a list NetlifySpec pointers to objects.
// This is helpful when the SetNetlifySpec requires a list of objects
func NewNetlifySpecList(ps ...*NetlifySpec) []NetlifySpec {
	objs := []NetlifySpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this NetlifySpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *NetlifyComponent) Validate() error {
	if o.Properties.Backofflimit == nil {
		return errors.New("Backofflimit in NetlifySpec must be set")
	}
	if o.Properties.Completions == nil {
		return errors.New("Completions in NetlifySpec must be set")
	}
	if o.Properties.Gitrepo == nil {
		return errors.New("Gitrepo in NetlifySpec must be set")
	}
	if o.Properties.Reponame == nil {
		return errors.New("Reponame in NetlifySpec must be set")
	}
	if o.Properties.RestartPolicy == nil {
		return errors.New("RestartPolicy in NetlifySpec must be set")
	}
	if o.Properties.Sitname == nil {
		return errors.New("Sitname in NetlifySpec must be set")
	}
	if o.Properties.Token == nil {
		return errors.New("Token in NetlifySpec must be set")
	}
	// validate all nested properties

	for i, v := range o.Base.Traits {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("traits[%d] %s in %s component is invalid: %w", i, v.DefType(), NetlifyType, err)
		}
	}
	return nil
}

// GetBackofflimit returns the Backofflimit field value
func (o *NetlifyComponent) GetBackofflimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Properties.Backofflimit
}

// GetBackofflimitOk returns a tuple with the Backofflimit field value
// and a boolean to check if the value has been set.
func (o *NetlifyComponent) GetBackofflimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Backofflimit, true
}

// SetBackofflimit sets field value
func (o *NetlifyComponent) SetBackofflimit(v int32) *NetlifyComponent {
	o.Properties.Backofflimit = &v
	return o
}

// GetCompletions returns the Completions field value
func (o *NetlifyComponent) GetCompletions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Properties.Completions
}

// GetCompletionsOk returns a tuple with the Completions field value
// and a boolean to check if the value has been set.
func (o *NetlifyComponent) GetCompletionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Completions, true
}

// SetCompletions sets field value
func (o *NetlifyComponent) SetCompletions(v int32) *NetlifyComponent {
	o.Properties.Completions = &v
	return o
}

// GetGitrepo returns the Gitrepo field value
func (o *NetlifyComponent) GetGitrepo() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Gitrepo
}

// GetGitrepoOk returns a tuple with the Gitrepo field value
// and a boolean to check if the value has been set.
func (o *NetlifyComponent) GetGitrepoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Gitrepo, true
}

// SetGitrepo sets field value
func (o *NetlifyComponent) SetGitrepo(v string) *NetlifyComponent {
	o.Properties.Gitrepo = &v
	return o
}

// GetReponame returns the Reponame field value
func (o *NetlifyComponent) GetReponame() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Reponame
}

// GetReponameOk returns a tuple with the Reponame field value
// and a boolean to check if the value has been set.
func (o *NetlifyComponent) GetReponameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Reponame, true
}

// SetReponame sets field value
func (o *NetlifyComponent) SetReponame(v string) *NetlifyComponent {
	o.Properties.Reponame = &v
	return o
}

// GetRestartPolicy returns the RestartPolicy field value
func (o *NetlifyComponent) GetRestartPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.RestartPolicy
}

// GetRestartPolicyOk returns a tuple with the RestartPolicy field value
// and a boolean to check if the value has been set.
func (o *NetlifyComponent) GetRestartPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.RestartPolicy, true
}

// SetRestartPolicy sets field value
func (o *NetlifyComponent) SetRestartPolicy(v string) *NetlifyComponent {
	o.Properties.RestartPolicy = &v
	return o
}

// GetSitname returns the Sitname field value
func (o *NetlifyComponent) GetSitname() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Sitname
}

// GetSitnameOk returns a tuple with the Sitname field value
// and a boolean to check if the value has been set.
func (o *NetlifyComponent) GetSitnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Sitname, true
}

// SetSitname sets field value
func (o *NetlifyComponent) SetSitname(v string) *NetlifyComponent {
	o.Properties.Sitname = &v
	return o
}

// GetToken returns the Token field value
func (o *NetlifyComponent) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *NetlifyComponent) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Token, true
}

// SetToken sets field value
func (o *NetlifyComponent) SetToken(v string) *NetlifyComponent {
	o.Properties.Token = &v
	return o
}

func (o NetlifySpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetlifySpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["backofflimit"] = o.Backofflimit
	toSerialize["completions"] = o.Completions
	toSerialize["gitrepo"] = o.Gitrepo
	toSerialize["reponame"] = o.Reponame
	toSerialize["restartPolicy"] = o.RestartPolicy
	toSerialize["sitname"] = o.Sitname
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableNetlifySpec struct {
	value *NetlifySpec
	isSet bool
}

func (v *NullableNetlifySpec) Get() *NetlifySpec {
	return v.value
}

func (v *NullableNetlifySpec) Set(val *NetlifySpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableNetlifySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableNetlifySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetlifySpec(val *NetlifySpec) *NullableNetlifySpec {
	return &NullableNetlifySpec{value: val, isSet: true}
}

func (v NullableNetlifySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetlifySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const NetlifyType = "netlify"

func init() {
	sdkcommon.RegisterComponent(NetlifyType, FromComponent)
}

type NetlifyComponent struct {
	Base       apis.ComponentBase
	Properties NetlifySpec
}

func Netlify(name string) *NetlifyComponent {
	n := &NetlifyComponent{Base: apis.ComponentBase{
		Name: name,
		Type: NetlifyType,
	}}
	return n
}

func (n *NetlifyComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range n.Base.Traits {
		traits = append(traits, trait.Build())
	}
	res := common.ApplicationComponent{
		DependsOn:  n.Base.DependsOn,
		Inputs:     n.Base.Inputs,
		Name:       n.Base.Name,
		Outputs:    n.Base.Outputs,
		Properties: util.Object2RawExtension(n.Properties),
		Traits:     traits,
		Type:       NetlifyType,
	}
	return res
}

func (n *NetlifyComponent) FromComponent(from common.ApplicationComponent) (*NetlifyComponent, error) {
	for _, trait := range from.Traits {
		_t, err := sdkcommon.FromTrait(&trait)
		if err != nil {
			return nil, err
		}
		n.Base.Traits = append(n.Base.Traits, _t)
	}
	var properties NetlifySpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	n.Base.Name = from.Name
	n.Base.DependsOn = from.DependsOn
	n.Base.Inputs = from.Inputs
	n.Base.Outputs = from.Outputs
	n.Base.Type = NetlifyType
	n.Properties = properties
	return n, nil
}

func FromComponent(from common.ApplicationComponent) (apis.Component, error) {
	n := &NetlifyComponent{}
	return n.FromComponent(from)
}

func (n *NetlifyComponent) SetTraits(traits ...apis.Trait) *NetlifyComponent {
	for _, addTrait := range traits {
		found := false
		for i, _t := range n.Base.Traits {
			if _t.DefType() == addTrait.DefType() {
				n.Base.Traits[i] = addTrait
				found = true
				break
			}
			if !found {
				n.Base.Traits = append(n.Base.Traits, addTrait)
			}
		}
	}
	return n
}

func (n *NetlifyComponent) GetTrait(typ string) apis.Trait {
	for _, _t := range n.Base.Traits {
		if _t.DefType() == typ {
			return _t
		}
	}
	return nil
}

func (n *NetlifyComponent) GetAllTraits() []apis.Trait {
	return n.Base.Traits
}

func (n *NetlifyComponent) ComponentName() string {
	return n.Base.Name
}

func (n *NetlifyComponent) DefType() string {
	return NetlifyType
}

func (n *NetlifyComponent) DependsOn(dependsOn []string) *NetlifyComponent {
	n.Base.DependsOn = dependsOn
	return n
}

func (n *NetlifyComponent) Inputs(input common.StepInputs) *NetlifyComponent {
	n.Base.Inputs = input
	return n
}

func (n *NetlifyComponent) Outputs(output common.StepOutputs) *NetlifyComponent {
	n.Base.Outputs = output
	return n
}

func (n *NetlifyComponent) AddDependsOn(dependsOn string) *NetlifyComponent {
	n.Base.DependsOn = append(n.Base.DependsOn, dependsOn)
	return n
}
