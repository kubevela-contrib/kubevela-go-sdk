/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package task

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

// checks if the TaskSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TaskSpec{}

// TaskSpec struct for TaskSpec
type TaskSpec struct {
	// Specify the annotations in the workload
	Annotations map[string]string `json:"annotations,omitempty"`
	// Commands to run in the container
	Cmd []string `json:"cmd,omitempty"`
	// Specify number of tasks to run in parallel +short=c
	Count *int32 `json:"count"`
	// Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
	Cpu *string `json:"cpu,omitempty"`
	// Define arguments by using environment variables
	Env []Env `json:"env,omitempty"`
	// Which image would you like to use for your service +short=i
	Image *string `json:"image"`
	// Specify image pull policy for your service
	ImagePullPolicy *string `json:"imagePullPolicy,omitempty"`
	// Specify image pull secrets for your service
	ImagePullSecrets []string `json:"imagePullSecrets,omitempty"`
	// Specify the labels in the workload
	Labels        map[string]string `json:"labels,omitempty"`
	LivenessProbe *HealthProbe      `json:"livenessProbe,omitempty"`
	// Specifies the attributes of the memory resource required for the container.
	Memory         *string      `json:"memory,omitempty"`
	ReadinessProbe *HealthProbe `json:"readinessProbe,omitempty"`
	// Define the job restart policy, the value can only be Never or OnFailure. By default, it's Never.
	Restart *string `json:"restart"`
	// Declare volumes and volumeMounts
	Volumes []Volumes `json:"volumes,omitempty"`
}

// NewTaskSpecWith instantiates a new TaskSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewTaskSpecWith(count int32, image string, restart string) *TaskSpec {
	this := TaskSpec{}
	this.Count = &count
	this.Image = &image
	this.Restart = &restart
	return &this
}

// NewTaskSpecWithDefault instantiates a new TaskSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskSpecWithDefault() *TaskSpec {
	this := TaskSpec{}
	var count int32 = 1
	this.Count = &count
	var restart string = "Never"
	this.Restart = &restart
	return &this
}

// NewTaskSpec is short for NewTaskSpecWithDefault which instantiates a new TaskSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskSpec() *TaskSpec {
	return NewTaskSpecWithDefault()
}

// NewTaskSpecEmpty instantiates a new TaskSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewTaskSpecEmpty() *TaskSpec {
	this := TaskSpec{}
	return &this
}

// NewTaskSpecs converts a list TaskSpec pointers to objects.
// This is helpful when the SetTaskSpec requires a list of objects
func NewTaskSpecList(ps ...*TaskSpec) []TaskSpec {
	objs := []TaskSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this TaskSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *TaskComponent) Validate() error {
	if o.Properties.Count == nil {
		return errors.New("Count in TaskSpec must be set")
	}
	if o.Properties.Image == nil {
		return errors.New("Image in TaskSpec must be set")
	}
	if o.Properties.Restart == nil {
		return errors.New("Restart in TaskSpec must be set")
	}
	// validate all nested properties
	if o.Properties.LivenessProbe != nil {
		if err := o.Properties.LivenessProbe.Validate(); err != nil {
			return err
		}
	}
	if o.Properties.ReadinessProbe != nil {
		if err := o.Properties.ReadinessProbe.Validate(); err != nil {
			return err
		}
	}

	for i, v := range o.Base.Traits {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("traits[%d] %s in %s component is invalid: %w", i, v.DefType(), TaskType, err)
		}
	}
	return nil
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *TaskComponent) GetAnnotations() map[string]string {
	if o == nil || utils.IsNil(o.Properties.Annotations) {
		var ret map[string]string
		return ret
	}
	return o.Properties.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskComponent) GetAnnotationsOk() (map[string]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Annotations) {
		return nil, false
	}
	return o.Properties.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *TaskComponent) HasAnnotations() bool {
	if o != nil && !utils.IsNil(o.Properties.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the annotations field.
// Annotations:  Specify the annotations in the workload
func (o *TaskComponent) SetAnnotations(v map[string]string) *TaskComponent {
	o.Properties.Annotations = v
	return o
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *TaskComponent) GetCmd() []string {
	if o == nil || utils.IsNil(o.Properties.Cmd) {
		var ret []string
		return ret
	}
	return o.Properties.Cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskComponent) GetCmdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Cmd) {
		return nil, false
	}
	return o.Properties.Cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *TaskComponent) HasCmd() bool {
	if o != nil && !utils.IsNil(o.Properties.Cmd) {
		return true
	}

	return false
}

// SetCmd gets a reference to the given []string and assigns it to the cmd field.
// Cmd:  Commands to run in the container
func (o *TaskComponent) SetCmd(v []string) *TaskComponent {
	o.Properties.Cmd = v
	return o
}

// GetCount returns the Count field value
func (o *TaskComponent) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Properties.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *TaskComponent) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Count, true
}

// SetCount sets field value
func (o *TaskComponent) SetCount(v int32) *TaskComponent {
	o.Properties.Count = &v
	return o
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *TaskComponent) GetCpu() string {
	if o == nil || utils.IsNil(o.Properties.Cpu) {
		var ret string
		return ret
	}
	return *o.Properties.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskComponent) GetCpuOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Cpu) {
		return nil, false
	}
	return o.Properties.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *TaskComponent) HasCpu() bool {
	if o != nil && !utils.IsNil(o.Properties.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given string and assigns it to the cpu field.
// Cpu:  Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
func (o *TaskComponent) SetCpu(v string) *TaskComponent {
	o.Properties.Cpu = &v
	return o
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *TaskComponent) GetEnv() []Env {
	if o == nil || utils.IsNil(o.Properties.Env) {
		var ret []Env
		return ret
	}
	return o.Properties.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskComponent) GetEnvOk() ([]Env, bool) {
	if o == nil || utils.IsNil(o.Properties.Env) {
		return nil, false
	}
	return o.Properties.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *TaskComponent) HasEnv() bool {
	if o != nil && !utils.IsNil(o.Properties.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []Env and assigns it to the env field.
// Env:  Define arguments by using environment variables
func (o *TaskComponent) SetEnv(v []Env) *TaskComponent {
	o.Properties.Env = v
	return o
}

// GetImage returns the Image field value
func (o *TaskComponent) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *TaskComponent) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Image, true
}

// SetImage sets field value
func (o *TaskComponent) SetImage(v string) *TaskComponent {
	o.Properties.Image = &v
	return o
}

// GetImagePullPolicy returns the ImagePullPolicy field value if set, zero value otherwise.
func (o *TaskComponent) GetImagePullPolicy() string {
	if o == nil || utils.IsNil(o.Properties.ImagePullPolicy) {
		var ret string
		return ret
	}
	return *o.Properties.ImagePullPolicy
}

// GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskComponent) GetImagePullPolicyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.ImagePullPolicy) {
		return nil, false
	}
	return o.Properties.ImagePullPolicy, true
}

// HasImagePullPolicy returns a boolean if a field has been set.
func (o *TaskComponent) HasImagePullPolicy() bool {
	if o != nil && !utils.IsNil(o.Properties.ImagePullPolicy) {
		return true
	}

	return false
}

// SetImagePullPolicy gets a reference to the given string and assigns it to the imagePullPolicy field.
// ImagePullPolicy:  Specify image pull policy for your service
func (o *TaskComponent) SetImagePullPolicy(v string) *TaskComponent {
	o.Properties.ImagePullPolicy = &v
	return o
}

// GetImagePullSecrets returns the ImagePullSecrets field value if set, zero value otherwise.
func (o *TaskComponent) GetImagePullSecrets() []string {
	if o == nil || utils.IsNil(o.Properties.ImagePullSecrets) {
		var ret []string
		return ret
	}
	return o.Properties.ImagePullSecrets
}

// GetImagePullSecretsOk returns a tuple with the ImagePullSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskComponent) GetImagePullSecretsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.ImagePullSecrets) {
		return nil, false
	}
	return o.Properties.ImagePullSecrets, true
}

// HasImagePullSecrets returns a boolean if a field has been set.
func (o *TaskComponent) HasImagePullSecrets() bool {
	if o != nil && !utils.IsNil(o.Properties.ImagePullSecrets) {
		return true
	}

	return false
}

// SetImagePullSecrets gets a reference to the given []string and assigns it to the imagePullSecrets field.
// ImagePullSecrets:  Specify image pull secrets for your service
func (o *TaskComponent) SetImagePullSecrets(v []string) *TaskComponent {
	o.Properties.ImagePullSecrets = v
	return o
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *TaskComponent) GetLabels() map[string]string {
	if o == nil || utils.IsNil(o.Properties.Labels) {
		var ret map[string]string
		return ret
	}
	return o.Properties.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskComponent) GetLabelsOk() (map[string]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Labels) {
		return nil, false
	}
	return o.Properties.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *TaskComponent) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Properties.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the labels field.
// Labels:  Specify the labels in the workload
func (o *TaskComponent) SetLabels(v map[string]string) *TaskComponent {
	o.Properties.Labels = v
	return o
}

// GetLivenessProbe returns the LivenessProbe field value if set, zero value otherwise.
func (o *TaskComponent) GetLivenessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.Properties.LivenessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.Properties.LivenessProbe
}

// GetLivenessProbeOk returns a tuple with the LivenessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskComponent) GetLivenessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.Properties.LivenessProbe) {
		return nil, false
	}
	return o.Properties.LivenessProbe, true
}

// HasLivenessProbe returns a boolean if a field has been set.
func (o *TaskComponent) HasLivenessProbe() bool {
	if o != nil && !utils.IsNil(o.Properties.LivenessProbe) {
		return true
	}

	return false
}

// SetLivenessProbe gets a reference to the given HealthProbe and assigns it to the livenessProbe field.
// LivenessProbe:
func (o *TaskComponent) SetLivenessProbe(v HealthProbe) *TaskComponent {
	o.Properties.LivenessProbe = &v
	return o
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *TaskComponent) GetMemory() string {
	if o == nil || utils.IsNil(o.Properties.Memory) {
		var ret string
		return ret
	}
	return *o.Properties.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskComponent) GetMemoryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Memory) {
		return nil, false
	}
	return o.Properties.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *TaskComponent) HasMemory() bool {
	if o != nil && !utils.IsNil(o.Properties.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given string and assigns it to the memory field.
// Memory:  Specifies the attributes of the memory resource required for the container.
func (o *TaskComponent) SetMemory(v string) *TaskComponent {
	o.Properties.Memory = &v
	return o
}

// GetReadinessProbe returns the ReadinessProbe field value if set, zero value otherwise.
func (o *TaskComponent) GetReadinessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.Properties.ReadinessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.Properties.ReadinessProbe
}

// GetReadinessProbeOk returns a tuple with the ReadinessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskComponent) GetReadinessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.Properties.ReadinessProbe) {
		return nil, false
	}
	return o.Properties.ReadinessProbe, true
}

// HasReadinessProbe returns a boolean if a field has been set.
func (o *TaskComponent) HasReadinessProbe() bool {
	if o != nil && !utils.IsNil(o.Properties.ReadinessProbe) {
		return true
	}

	return false
}

// SetReadinessProbe gets a reference to the given HealthProbe and assigns it to the readinessProbe field.
// ReadinessProbe:
func (o *TaskComponent) SetReadinessProbe(v HealthProbe) *TaskComponent {
	o.Properties.ReadinessProbe = &v
	return o
}

// GetRestart returns the Restart field value
func (o *TaskComponent) GetRestart() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Restart
}

// GetRestartOk returns a tuple with the Restart field value
// and a boolean to check if the value has been set.
func (o *TaskComponent) GetRestartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Restart, true
}

// SetRestart sets field value
func (o *TaskComponent) SetRestart(v string) *TaskComponent {
	o.Properties.Restart = &v
	return o
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *TaskComponent) GetVolumes() []Volumes {
	if o == nil || utils.IsNil(o.Properties.Volumes) {
		var ret []Volumes
		return ret
	}
	return o.Properties.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskComponent) GetVolumesOk() ([]Volumes, bool) {
	if o == nil || utils.IsNil(o.Properties.Volumes) {
		return nil, false
	}
	return o.Properties.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *TaskComponent) HasVolumes() bool {
	if o != nil && !utils.IsNil(o.Properties.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []Volumes and assigns it to the volumes field.
// Volumes:  Declare volumes and volumeMounts
func (o *TaskComponent) SetVolumes(v []Volumes) *TaskComponent {
	o.Properties.Volumes = v
	return o
}

func (o TaskSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !utils.IsNil(o.Cmd) {
		toSerialize["cmd"] = o.Cmd
	}
	toSerialize["count"] = o.Count
	if !utils.IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !utils.IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	toSerialize["image"] = o.Image
	if !utils.IsNil(o.ImagePullPolicy) {
		toSerialize["imagePullPolicy"] = o.ImagePullPolicy
	}
	if !utils.IsNil(o.ImagePullSecrets) {
		toSerialize["imagePullSecrets"] = o.ImagePullSecrets
	}
	if !utils.IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !utils.IsNil(o.LivenessProbe) {
		toSerialize["livenessProbe"] = o.LivenessProbe
	}
	if !utils.IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !utils.IsNil(o.ReadinessProbe) {
		toSerialize["readinessProbe"] = o.ReadinessProbe
	}
	toSerialize["restart"] = o.Restart
	if !utils.IsNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	return toSerialize, nil
}

type NullableTaskSpec struct {
	value *TaskSpec
	isSet bool
}

func (v *NullableTaskSpec) Get() *TaskSpec {
	return v.value
}

func (v *NullableTaskSpec) Set(val *TaskSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableTaskSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskSpec(val *TaskSpec) *NullableTaskSpec {
	return &NullableTaskSpec{value: val, isSet: true}
}

func (v NullableTaskSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const TaskType = "task"

func init() {
	sdkcommon.RegisterComponent(TaskType, FromComponent)
}

type TaskComponent struct {
	Base       apis.ComponentBase
	Properties TaskSpec
}

func Task(name string) *TaskComponent {
	t := &TaskComponent{Base: apis.ComponentBase{
		Name: name,
		Type: TaskType,
	}}
	return t
}

func (t *TaskComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range t.Base.Traits {
		traits = append(traits, trait.Build())
	}
	res := common.ApplicationComponent{
		DependsOn:  t.Base.DependsOn,
		Inputs:     t.Base.Inputs,
		Name:       t.Base.Name,
		Outputs:    t.Base.Outputs,
		Properties: util.Object2RawExtension(t.Properties),
		Traits:     traits,
		Type:       TaskType,
	}
	return res
}

func (t *TaskComponent) FromComponent(from common.ApplicationComponent) (*TaskComponent, error) {
	for _, trait := range from.Traits {
		_t, err := sdkcommon.FromTrait(&trait)
		if err != nil {
			return nil, err
		}
		t.Base.Traits = append(t.Base.Traits, _t)
	}
	var properties TaskSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	t.Base.Name = from.Name
	t.Base.DependsOn = from.DependsOn
	t.Base.Inputs = from.Inputs
	t.Base.Outputs = from.Outputs
	t.Base.Type = TaskType
	t.Properties = properties
	return t, nil
}

func FromComponent(from common.ApplicationComponent) (apis.Component, error) {
	t := &TaskComponent{}
	return t.FromComponent(from)
}

func (t *TaskComponent) SetTraits(traits ...apis.Trait) *TaskComponent {
	for _, addTrait := range traits {
		found := false
		for i, _t := range t.Base.Traits {
			if _t.DefType() == addTrait.DefType() {
				t.Base.Traits[i] = addTrait
				found = true
				break
			}
			if !found {
				t.Base.Traits = append(t.Base.Traits, addTrait)
			}
		}
	}
	return t
}

func (t *TaskComponent) GetTrait(typ string) apis.Trait {
	for _, _t := range t.Base.Traits {
		if _t.DefType() == typ {
			return _t
		}
	}
	return nil
}

func (t *TaskComponent) GetAllTraits() []apis.Trait {
	return t.Base.Traits
}

func (t *TaskComponent) ComponentName() string {
	return t.Base.Name
}

func (t *TaskComponent) DefType() string {
	return TaskType
}

func (t *TaskComponent) DependsOn(dependsOn []string) *TaskComponent {
	t.Base.DependsOn = dependsOn
	return t
}

func (t *TaskComponent) Inputs(input common.StepInputs) *TaskComponent {
	t.Base.Inputs = input
	return t
}

func (t *TaskComponent) Outputs(output common.StepOutputs) *TaskComponent {
	t.Base.Outputs = output
	return t
}

func (t *TaskComponent) AddDependsOn(dependsOn string) *TaskComponent {
	t.Base.DependsOn = append(t.Base.DependsOn, dependsOn)
	return t
}
