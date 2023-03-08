/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package init_container

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

// checks if the InitContainerSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InitContainerSpec{}

// InitContainerSpec struct for InitContainerSpec
type InitContainerSpec struct {
	// Specify the mount path of app container
	AppMountPath *string `json:"appMountPath"`
	// Specify the args run in the init container
	Args []string `json:"args,omitempty"`
	// Specify the commands run in the init container
	Cmd []string `json:"cmd,omitempty"`
	// Specify the env run in the init container
	Env []Env `json:"env,omitempty"`
	// Specify the extra volume mounts for the init container
	ExtraVolumeMounts []ExtraVolumeMounts `json:"extraVolumeMounts"`
	// Specify the image of init container
	Image *string `json:"image"`
	// Specify image pull policy for your service
	ImagePullPolicy *string `json:"imagePullPolicy"`
	// Specify the mount path of init container
	InitMountPath *string `json:"initMountPath"`
	// Specify the mount name of shared volume
	MountName *string `json:"mountName"`
	// Specify the name of init container
	Name *string `json:"name"`
}

// NewInitContainerSpecWith instantiates a new InitContainerSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewInitContainerSpecWith(appMountPath string, extraVolumeMounts []ExtraVolumeMounts, image string, imagePullPolicy string, initMountPath string, mountName string, name string) *InitContainerSpec {
	this := InitContainerSpec{}
	this.AppMountPath = &appMountPath
	this.ExtraVolumeMounts = extraVolumeMounts
	this.Image = &image
	this.ImagePullPolicy = &imagePullPolicy
	this.InitMountPath = &initMountPath
	this.MountName = &mountName
	this.Name = &name
	return &this
}

// NewInitContainerSpecWithDefault instantiates a new InitContainerSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitContainerSpecWithDefault() *InitContainerSpec {
	this := InitContainerSpec{}
	var imagePullPolicy string = "IfNotPresent"
	this.ImagePullPolicy = &imagePullPolicy
	var mountName string = "workdir"
	this.MountName = &mountName
	return &this
}

// NewInitContainerSpec is short for NewInitContainerSpecWithDefault which instantiates a new InitContainerSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitContainerSpec() *InitContainerSpec {
	return NewInitContainerSpecWithDefault()
}

// NewInitContainerSpecEmpty instantiates a new InitContainerSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewInitContainerSpecEmpty() *InitContainerSpec {
	this := InitContainerSpec{}
	return &this
}

// NewInitContainerSpecs converts a list InitContainerSpec pointers to objects.
// This is helpful when the SetInitContainerSpec requires a list of objects
func NewInitContainerSpecList(ps ...*InitContainerSpec) []InitContainerSpec {
	objs := []InitContainerSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this InitContainerSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *InitContainerTrait) Validate() error {
	if o.Properties.AppMountPath == nil {
		return errors.New("AppMountPath in InitContainerSpec must be set")
	}
	if o.Properties.ExtraVolumeMounts == nil {
		return errors.New("ExtraVolumeMounts in InitContainerSpec must be set")
	}
	if o.Properties.Image == nil {
		return errors.New("Image in InitContainerSpec must be set")
	}
	if o.Properties.ImagePullPolicy == nil {
		return errors.New("ImagePullPolicy in InitContainerSpec must be set")
	}
	if o.Properties.InitMountPath == nil {
		return errors.New("InitMountPath in InitContainerSpec must be set")
	}
	if o.Properties.MountName == nil {
		return errors.New("MountName in InitContainerSpec must be set")
	}
	if o.Properties.Name == nil {
		return errors.New("Name in InitContainerSpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetAppMountPath returns the AppMountPath field value
func (o *InitContainerTrait) GetAppMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.AppMountPath
}

// GetAppMountPathOk returns a tuple with the AppMountPath field value
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetAppMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.AppMountPath, true
}

// SetAppMountPath sets field value
func (o *InitContainerTrait) SetAppMountPath(v string) *InitContainerTrait {
	o.Properties.AppMountPath = &v
	return o
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *InitContainerTrait) GetArgs() []string {
	if o == nil || utils.IsNil(o.Properties.Args) {
		var ret []string
		return ret
	}
	return o.Properties.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetArgsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Args) {
		return nil, false
	}
	return o.Properties.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *InitContainerTrait) HasArgs() bool {
	if o != nil && !utils.IsNil(o.Properties.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []string and assigns it to the args field.
// Args:  Specify the args run in the init container
func (o *InitContainerTrait) SetArgs(v []string) *InitContainerTrait {
	o.Properties.Args = v
	return o
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *InitContainerTrait) GetCmd() []string {
	if o == nil || utils.IsNil(o.Properties.Cmd) {
		var ret []string
		return ret
	}
	return o.Properties.Cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetCmdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Cmd) {
		return nil, false
	}
	return o.Properties.Cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *InitContainerTrait) HasCmd() bool {
	if o != nil && !utils.IsNil(o.Properties.Cmd) {
		return true
	}

	return false
}

// SetCmd gets a reference to the given []string and assigns it to the cmd field.
// Cmd:  Specify the commands run in the init container
func (o *InitContainerTrait) SetCmd(v []string) *InitContainerTrait {
	o.Properties.Cmd = v
	return o
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *InitContainerTrait) GetEnv() []Env {
	if o == nil || utils.IsNil(o.Properties.Env) {
		var ret []Env
		return ret
	}
	return o.Properties.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetEnvOk() ([]Env, bool) {
	if o == nil || utils.IsNil(o.Properties.Env) {
		return nil, false
	}
	return o.Properties.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *InitContainerTrait) HasEnv() bool {
	if o != nil && !utils.IsNil(o.Properties.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []Env and assigns it to the env field.
// Env:  Specify the env run in the init container
func (o *InitContainerTrait) SetEnv(v []Env) *InitContainerTrait {
	o.Properties.Env = v
	return o
}

// GetExtraVolumeMounts returns the ExtraVolumeMounts field value
func (o *InitContainerTrait) GetExtraVolumeMounts() []ExtraVolumeMounts {
	if o == nil {
		var ret []ExtraVolumeMounts
		return ret
	}

	return o.Properties.ExtraVolumeMounts
}

// GetExtraVolumeMountsOk returns a tuple with the ExtraVolumeMounts field value
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetExtraVolumeMountsOk() ([]ExtraVolumeMounts, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.ExtraVolumeMounts, true
}

// SetExtraVolumeMounts sets field value
func (o *InitContainerTrait) SetExtraVolumeMounts(v []ExtraVolumeMounts) *InitContainerTrait {
	o.Properties.ExtraVolumeMounts = v
	return o
}

// GetImage returns the Image field value
func (o *InitContainerTrait) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Image, true
}

// SetImage sets field value
func (o *InitContainerTrait) SetImage(v string) *InitContainerTrait {
	o.Properties.Image = &v
	return o
}

// GetImagePullPolicy returns the ImagePullPolicy field value
func (o *InitContainerTrait) GetImagePullPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.ImagePullPolicy
}

// GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field value
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetImagePullPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.ImagePullPolicy, true
}

// SetImagePullPolicy sets field value
func (o *InitContainerTrait) SetImagePullPolicy(v string) *InitContainerTrait {
	o.Properties.ImagePullPolicy = &v
	return o
}

// GetInitMountPath returns the InitMountPath field value
func (o *InitContainerTrait) GetInitMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.InitMountPath
}

// GetInitMountPathOk returns a tuple with the InitMountPath field value
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetInitMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.InitMountPath, true
}

// SetInitMountPath sets field value
func (o *InitContainerTrait) SetInitMountPath(v string) *InitContainerTrait {
	o.Properties.InitMountPath = &v
	return o
}

// GetMountName returns the MountName field value
func (o *InitContainerTrait) GetMountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.MountName
}

// GetMountNameOk returns a tuple with the MountName field value
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetMountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.MountName, true
}

// SetMountName sets field value
func (o *InitContainerTrait) SetMountName(v string) *InitContainerTrait {
	o.Properties.MountName = &v
	return o
}

// GetName returns the Name field value
func (o *InitContainerTrait) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Name, true
}

// SetName sets field value
func (o *InitContainerTrait) SetName(v string) *InitContainerTrait {
	o.Properties.Name = &v
	return o
}

func (o InitContainerSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InitContainerSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appMountPath"] = o.AppMountPath
	if !utils.IsNil(o.Args) {
		toSerialize["args"] = o.Args
	}
	if !utils.IsNil(o.Cmd) {
		toSerialize["cmd"] = o.Cmd
	}
	if !utils.IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	toSerialize["extraVolumeMounts"] = o.ExtraVolumeMounts
	toSerialize["image"] = o.Image
	toSerialize["imagePullPolicy"] = o.ImagePullPolicy
	toSerialize["initMountPath"] = o.InitMountPath
	toSerialize["mountName"] = o.MountName
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableInitContainerSpec struct {
	value *InitContainerSpec
	isSet bool
}

func (v *NullableInitContainerSpec) Get() *InitContainerSpec {
	return v.value
}

func (v *NullableInitContainerSpec) Set(val *InitContainerSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableInitContainerSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableInitContainerSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitContainerSpec(val *InitContainerSpec) *NullableInitContainerSpec {
	return &NullableInitContainerSpec{value: val, isSet: true}
}

func (v NullableInitContainerSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitContainerSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const InitContainerType = "init-container"

func init() {
	sdkcommon.RegisterTrait(InitContainerType, FromTrait)
}

type InitContainerTrait struct {
	Base       apis.TraitBase
	Properties InitContainerSpec
}

func InitContainer() *InitContainerTrait {
	i := &InitContainerTrait{Base: apis.TraitBase{}}
	return i
}

func (i *InitContainerTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(i.Properties),
		Type:       InitContainerType,
	}
	return res
}

func (i *InitContainerTrait) FromTrait(from common.ApplicationTrait) (*InitContainerTrait, error) {
	var properties InitContainerSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	i.Base.Type = InitContainerType
	i.Properties = properties
	return i, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	i := &InitContainerTrait{}
	return i.FromTrait(from)
}

func (i *InitContainerTrait) DefType() string {
	return InitContainerType
}
