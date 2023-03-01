/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sidecar

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the SidecarSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SidecarSpec{}

// SidecarSpec struct for SidecarSpec
type SidecarSpec struct {
	// Specify the args in the sidecar
	Args []string `json:"args,omitempty"`
	// Specify the commands run in the sidecar
	Cmd []string `json:"cmd,omitempty"`
	// Specify the env in the sidecar
	Env []Env `json:"env,omitempty"`
	// Specify the image of sidecar container
	Image         *string      `json:"image,omitempty"`
	LivenessProbe *HealthProbe `json:"livenessProbe,omitempty"`
	// Specify the name of sidecar container
	Name           *string      `json:"name,omitempty"`
	ReadinessProbe *HealthProbe `json:"readinessProbe,omitempty"`
	// Specify the shared volume path
	Volumes []Volumes `json:"volumes,omitempty"`
}

// NewSidecarSpecWith instantiates a new SidecarSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSidecarSpecWith() *SidecarSpec {
	this := SidecarSpec{}
	return &this
}

// NewSidecarSpec instantiates a new SidecarSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSidecarSpec() *SidecarSpec {
	this := SidecarSpec{}
	return &this
}

// NewSidecarSpecs converts a list SidecarSpec pointers to objects.
// This is helpful when the SetSidecarSpec requires a list of objects
func NewSidecarSpecList(ps ...*SidecarSpec) []SidecarSpec {
	objs := []SidecarSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *SidecarTrait) GetArgs() []string {
	if o == nil || utils.IsNil(o.Properties.Args) {
		var ret []string
		return ret
	}
	return o.Properties.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SidecarTrait) GetArgsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Args) {
		return nil, false
	}
	return o.Properties.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *SidecarTrait) HasArgs() bool {
	if o != nil && !utils.IsNil(o.Properties.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []string and assigns it to the args field.
// Args:  Specify the args in the sidecar
func (o *SidecarTrait) SetArgs(v []string) *SidecarTrait {
	o.Properties.Args = v
	return o
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *SidecarTrait) GetCmd() []string {
	if o == nil || utils.IsNil(o.Properties.Cmd) {
		var ret []string
		return ret
	}
	return o.Properties.Cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SidecarTrait) GetCmdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Cmd) {
		return nil, false
	}
	return o.Properties.Cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *SidecarTrait) HasCmd() bool {
	if o != nil && !utils.IsNil(o.Properties.Cmd) {
		return true
	}

	return false
}

// SetCmd gets a reference to the given []string and assigns it to the cmd field.
// Cmd:  Specify the commands run in the sidecar
func (o *SidecarTrait) SetCmd(v []string) *SidecarTrait {
	o.Properties.Cmd = v
	return o
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *SidecarTrait) GetEnv() []Env {
	if o == nil || utils.IsNil(o.Properties.Env) {
		var ret []Env
		return ret
	}
	return o.Properties.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SidecarTrait) GetEnvOk() ([]Env, bool) {
	if o == nil || utils.IsNil(o.Properties.Env) {
		return nil, false
	}
	return o.Properties.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *SidecarTrait) HasEnv() bool {
	if o != nil && !utils.IsNil(o.Properties.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []Env and assigns it to the env field.
// Env:  Specify the env in the sidecar
func (o *SidecarTrait) SetEnv(v []Env) *SidecarTrait {
	o.Properties.Env = v
	return o
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *SidecarTrait) GetImage() string {
	if o == nil || utils.IsNil(o.Properties.Image) {
		var ret string
		return ret
	}
	return *o.Properties.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SidecarTrait) GetImageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Image) {
		return nil, false
	}
	return o.Properties.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *SidecarTrait) HasImage() bool {
	if o != nil && !utils.IsNil(o.Properties.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the image field.
// Image:  Specify the image of sidecar container
func (o *SidecarTrait) SetImage(v string) *SidecarTrait {
	o.Properties.Image = &v
	return o
}

// GetLivenessProbe returns the LivenessProbe field value if set, zero value otherwise.
func (o *SidecarTrait) GetLivenessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.Properties.LivenessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.Properties.LivenessProbe
}

// GetLivenessProbeOk returns a tuple with the LivenessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SidecarTrait) GetLivenessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.Properties.LivenessProbe) {
		return nil, false
	}
	return o.Properties.LivenessProbe, true
}

// HasLivenessProbe returns a boolean if a field has been set.
func (o *SidecarTrait) HasLivenessProbe() bool {
	if o != nil && !utils.IsNil(o.Properties.LivenessProbe) {
		return true
	}

	return false
}

// SetLivenessProbe gets a reference to the given HealthProbe and assigns it to the livenessProbe field.
// LivenessProbe:
func (o *SidecarTrait) SetLivenessProbe(v HealthProbe) *SidecarTrait {
	o.Properties.LivenessProbe = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SidecarTrait) GetName() string {
	if o == nil || utils.IsNil(o.Properties.Name) {
		var ret string
		return ret
	}
	return *o.Properties.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SidecarTrait) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Name) {
		return nil, false
	}
	return o.Properties.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SidecarTrait) HasName() bool {
	if o != nil && !utils.IsNil(o.Properties.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:  Specify the name of sidecar container
func (o *SidecarTrait) SetName(v string) *SidecarTrait {
	o.Properties.Name = &v
	return o
}

// GetReadinessProbe returns the ReadinessProbe field value if set, zero value otherwise.
func (o *SidecarTrait) GetReadinessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.Properties.ReadinessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.Properties.ReadinessProbe
}

// GetReadinessProbeOk returns a tuple with the ReadinessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SidecarTrait) GetReadinessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.Properties.ReadinessProbe) {
		return nil, false
	}
	return o.Properties.ReadinessProbe, true
}

// HasReadinessProbe returns a boolean if a field has been set.
func (o *SidecarTrait) HasReadinessProbe() bool {
	if o != nil && !utils.IsNil(o.Properties.ReadinessProbe) {
		return true
	}

	return false
}

// SetReadinessProbe gets a reference to the given HealthProbe and assigns it to the readinessProbe field.
// ReadinessProbe:
func (o *SidecarTrait) SetReadinessProbe(v HealthProbe) *SidecarTrait {
	o.Properties.ReadinessProbe = &v
	return o
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *SidecarTrait) GetVolumes() []Volumes {
	if o == nil || utils.IsNil(o.Properties.Volumes) {
		var ret []Volumes
		return ret
	}
	return o.Properties.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SidecarTrait) GetVolumesOk() ([]Volumes, bool) {
	if o == nil || utils.IsNil(o.Properties.Volumes) {
		return nil, false
	}
	return o.Properties.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *SidecarTrait) HasVolumes() bool {
	if o != nil && !utils.IsNil(o.Properties.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []Volumes and assigns it to the volumes field.
// Volumes:  Specify the shared volume path
func (o *SidecarTrait) SetVolumes(v []Volumes) *SidecarTrait {
	o.Properties.Volumes = v
	return o
}

func (o SidecarSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SidecarSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Args) {
		toSerialize["args"] = o.Args
	}
	if !utils.IsNil(o.Cmd) {
		toSerialize["cmd"] = o.Cmd
	}
	if !utils.IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !utils.IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !utils.IsNil(o.LivenessProbe) {
		toSerialize["livenessProbe"] = o.LivenessProbe
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.ReadinessProbe) {
		toSerialize["readinessProbe"] = o.ReadinessProbe
	}
	if !utils.IsNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	return toSerialize, nil
}

type NullableSidecarSpec struct {
	value *SidecarSpec
	isSet bool
}

func (v NullableSidecarSpec) Get() *SidecarSpec {
	return v.value
}

func (v *NullableSidecarSpec) Set(val *SidecarSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSidecarSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSidecarSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSidecarSpec(val *SidecarSpec) *NullableSidecarSpec {
	return &NullableSidecarSpec{value: val, isSet: true}
}

func (v NullableSidecarSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSidecarSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const SidecarType = "sidecar"

func init() {
	sdkcommon.RegisterTrait(SidecarType, FromTrait)
}

type SidecarTrait struct {
	Base       apis.TraitBase
	Properties SidecarSpec
}

func Sidecar() *SidecarTrait {
	s := &SidecarTrait{Base: apis.TraitBase{}}
	return s
}

func (s *SidecarTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(s.Properties),
		Type:       SidecarType,
	}
	return res
}

func (s *SidecarTrait) FromTrait(from common.ApplicationTrait) (*SidecarTrait, error) {
	var properties SidecarSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	s.Base.Type = SidecarType
	s.Properties = properties
	return s, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	s := &SidecarTrait{}
	return s.FromTrait(from)
}

func (s *SidecarTrait) DefType() string {
	return SidecarType
}
