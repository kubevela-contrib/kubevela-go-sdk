/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mysql_exporter_server

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

// checks if the MysqlExporterServerSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MysqlExporterServerSpec{}

// MysqlExporterServerSpec struct for MysqlExporterServerSpec
type MysqlExporterServerSpec struct {
	// Specify the CPU capacity of the Exporter collector.
	Cpu *string `json:"cpu,omitempty"`
	// Disable annotation means do not add the annotations for the exporter pod, and the Prometheus can not scrape it.
	DisableAnnotation *bool `json:"disableAnnotation"`
	// Specify the Memory capacity of the Exporter collector.
	Memory *string `json:"memory,omitempty"`
	// Specify the host of the target Mysql server, maybe you could set the mysql component name.
	MysqlHost *string `json:"mysqlHost"`
	// Specify the port of the target Mysql server.
	MysqlPort *int32 `json:"mysqlPort"`
	// Specify the name of the Exporter.
	Name *string `json:"name"`
	// Specify the password of the target Mysql server.
	Password *string `json:"password"`
	// Specify the username of the target Mysql server.
	Username *string `json:"username"`
	// Specify the version of the Exporter collector.
	Version *string `json:"version"`
}

// NewMysqlExporterServerSpecWith instantiates a new MysqlExporterServerSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewMysqlExporterServerSpecWith(disableAnnotation bool, mysqlHost string, mysqlPort int32, name string, password string, username string, version string) *MysqlExporterServerSpec {
	this := MysqlExporterServerSpec{}
	this.DisableAnnotation = &disableAnnotation
	this.MysqlHost = &mysqlHost
	this.MysqlPort = &mysqlPort
	this.Name = &name
	this.Password = &password
	this.Username = &username
	this.Version = &version
	return &this
}

// NewMysqlExporterServerSpecWithDefault instantiates a new MysqlExporterServerSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlExporterServerSpecWithDefault() *MysqlExporterServerSpec {
	this := MysqlExporterServerSpec{}
	var disableAnnotation bool = false
	this.DisableAnnotation = &disableAnnotation
	var mysqlHost string = "mysql-server"
	this.MysqlHost = &mysqlHost
	var mysqlPort int32 = 3306
	this.MysqlPort = &mysqlPort
	var name string = "mysql-server-exporter"
	this.Name = &name
	var version string = "v0.14.0"
	this.Version = &version
	return &this
}

// NewMysqlExporterServerSpec is short for NewMysqlExporterServerSpecWithDefault which instantiates a new MysqlExporterServerSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlExporterServerSpec() *MysqlExporterServerSpec {
	return NewMysqlExporterServerSpecWithDefault()
}

// NewMysqlExporterServerSpecEmpty instantiates a new MysqlExporterServerSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewMysqlExporterServerSpecEmpty() *MysqlExporterServerSpec {
	this := MysqlExporterServerSpec{}
	return &this
}

// NewMysqlExporterServerSpecs converts a list MysqlExporterServerSpec pointers to objects.
// This is helpful when the SetMysqlExporterServerSpec requires a list of objects
func NewMysqlExporterServerSpecList(ps ...*MysqlExporterServerSpec) []MysqlExporterServerSpec {
	objs := []MysqlExporterServerSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this MysqlExporterServerSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *MysqlExporterServerComponent) Validate() error {
	if o.Properties.DisableAnnotation == nil {
		return errors.New("DisableAnnotation in MysqlExporterServerSpec must be set")
	}
	if o.Properties.MysqlHost == nil {
		return errors.New("MysqlHost in MysqlExporterServerSpec must be set")
	}
	if o.Properties.MysqlPort == nil {
		return errors.New("MysqlPort in MysqlExporterServerSpec must be set")
	}
	if o.Properties.Name == nil {
		return errors.New("Name in MysqlExporterServerSpec must be set")
	}
	if o.Properties.Password == nil {
		return errors.New("Password in MysqlExporterServerSpec must be set")
	}
	if o.Properties.Username == nil {
		return errors.New("Username in MysqlExporterServerSpec must be set")
	}
	if o.Properties.Version == nil {
		return errors.New("Version in MysqlExporterServerSpec must be set")
	}
	// validate all nested properties

	for i, v := range o.Base.Traits {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("traits[%d] %s in %s component is invalid: %w", i, v.DefType(), MysqlExporterServerType, err)
		}
	}
	return nil
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *MysqlExporterServerComponent) GetCpu() string {
	if o == nil || utils.IsNil(o.Properties.Cpu) {
		var ret string
		return ret
	}
	return *o.Properties.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlExporterServerComponent) GetCpuOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Cpu) {
		return nil, false
	}
	return o.Properties.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *MysqlExporterServerComponent) HasCpu() bool {
	if o != nil && !utils.IsNil(o.Properties.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given string and assigns it to the cpu field.
// Cpu:  Specify the CPU capacity of the Exporter collector.
func (o *MysqlExporterServerComponent) SetCpu(v string) *MysqlExporterServerComponent {
	o.Properties.Cpu = &v
	return o
}

// GetDisableAnnotation returns the DisableAnnotation field value
func (o *MysqlExporterServerComponent) GetDisableAnnotation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return *o.Properties.DisableAnnotation
}

// GetDisableAnnotationOk returns a tuple with the DisableAnnotation field value
// and a boolean to check if the value has been set.
func (o *MysqlExporterServerComponent) GetDisableAnnotationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.DisableAnnotation, true
}

// SetDisableAnnotation sets field value
func (o *MysqlExporterServerComponent) SetDisableAnnotation(v bool) *MysqlExporterServerComponent {
	o.Properties.DisableAnnotation = &v
	return o
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *MysqlExporterServerComponent) GetMemory() string {
	if o == nil || utils.IsNil(o.Properties.Memory) {
		var ret string
		return ret
	}
	return *o.Properties.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlExporterServerComponent) GetMemoryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Memory) {
		return nil, false
	}
	return o.Properties.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *MysqlExporterServerComponent) HasMemory() bool {
	if o != nil && !utils.IsNil(o.Properties.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given string and assigns it to the memory field.
// Memory:  Specify the Memory capacity of the Exporter collector.
func (o *MysqlExporterServerComponent) SetMemory(v string) *MysqlExporterServerComponent {
	o.Properties.Memory = &v
	return o
}

// GetMysqlHost returns the MysqlHost field value
func (o *MysqlExporterServerComponent) GetMysqlHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.MysqlHost
}

// GetMysqlHostOk returns a tuple with the MysqlHost field value
// and a boolean to check if the value has been set.
func (o *MysqlExporterServerComponent) GetMysqlHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.MysqlHost, true
}

// SetMysqlHost sets field value
func (o *MysqlExporterServerComponent) SetMysqlHost(v string) *MysqlExporterServerComponent {
	o.Properties.MysqlHost = &v
	return o
}

// GetMysqlPort returns the MysqlPort field value
func (o *MysqlExporterServerComponent) GetMysqlPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Properties.MysqlPort
}

// GetMysqlPortOk returns a tuple with the MysqlPort field value
// and a boolean to check if the value has been set.
func (o *MysqlExporterServerComponent) GetMysqlPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.MysqlPort, true
}

// SetMysqlPort sets field value
func (o *MysqlExporterServerComponent) SetMysqlPort(v int32) *MysqlExporterServerComponent {
	o.Properties.MysqlPort = &v
	return o
}

// GetName returns the Name field value
func (o *MysqlExporterServerComponent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MysqlExporterServerComponent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Name, true
}

// SetName sets field value
func (o *MysqlExporterServerComponent) SetName(v string) *MysqlExporterServerComponent {
	o.Properties.Name = &v
	return o
}

// GetPassword returns the Password field value
func (o *MysqlExporterServerComponent) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *MysqlExporterServerComponent) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Password, true
}

// SetPassword sets field value
func (o *MysqlExporterServerComponent) SetPassword(v string) *MysqlExporterServerComponent {
	o.Properties.Password = &v
	return o
}

// GetUsername returns the Username field value
func (o *MysqlExporterServerComponent) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *MysqlExporterServerComponent) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Username, true
}

// SetUsername sets field value
func (o *MysqlExporterServerComponent) SetUsername(v string) *MysqlExporterServerComponent {
	o.Properties.Username = &v
	return o
}

// GetVersion returns the Version field value
func (o *MysqlExporterServerComponent) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *MysqlExporterServerComponent) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Version, true
}

// SetVersion sets field value
func (o *MysqlExporterServerComponent) SetVersion(v string) *MysqlExporterServerComponent {
	o.Properties.Version = &v
	return o
}

func (o MysqlExporterServerSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MysqlExporterServerSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	toSerialize["disableAnnotation"] = o.DisableAnnotation
	if !utils.IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	toSerialize["mysqlHost"] = o.MysqlHost
	toSerialize["mysqlPort"] = o.MysqlPort
	toSerialize["name"] = o.Name
	toSerialize["password"] = o.Password
	toSerialize["username"] = o.Username
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableMysqlExporterServerSpec struct {
	value *MysqlExporterServerSpec
	isSet bool
}

func (v *NullableMysqlExporterServerSpec) Get() *MysqlExporterServerSpec {
	return v.value
}

func (v *NullableMysqlExporterServerSpec) Set(val *MysqlExporterServerSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableMysqlExporterServerSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlExporterServerSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlExporterServerSpec(val *MysqlExporterServerSpec) *NullableMysqlExporterServerSpec {
	return &NullableMysqlExporterServerSpec{value: val, isSet: true}
}

func (v NullableMysqlExporterServerSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlExporterServerSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const MysqlExporterServerType = "mysql-exporter-server"

func init() {
	sdkcommon.RegisterComponent(MysqlExporterServerType, FromComponent)
}

type MysqlExporterServerComponent struct {
	Base       apis.ComponentBase
	Properties MysqlExporterServerSpec
}

func MysqlExporterServer(name string) *MysqlExporterServerComponent {
	m := &MysqlExporterServerComponent{Base: apis.ComponentBase{
		Name: name,
		Type: MysqlExporterServerType,
	}}
	return m
}

func (m *MysqlExporterServerComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range m.Base.Traits {
		traits = append(traits, trait.Build())
	}
	res := common.ApplicationComponent{
		DependsOn:  m.Base.DependsOn,
		Inputs:     m.Base.Inputs,
		Name:       m.Base.Name,
		Outputs:    m.Base.Outputs,
		Properties: util.Object2RawExtension(m.Properties),
		Traits:     traits,
		Type:       MysqlExporterServerType,
	}
	return res
}

func (m *MysqlExporterServerComponent) FromComponent(from common.ApplicationComponent) (*MysqlExporterServerComponent, error) {
	for _, trait := range from.Traits {
		_t, err := sdkcommon.FromTrait(&trait)
		if err != nil {
			return nil, err
		}
		m.Base.Traits = append(m.Base.Traits, _t)
	}
	var properties MysqlExporterServerSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	m.Base.Name = from.Name
	m.Base.DependsOn = from.DependsOn
	m.Base.Inputs = from.Inputs
	m.Base.Outputs = from.Outputs
	m.Base.Type = MysqlExporterServerType
	m.Properties = properties
	return m, nil
}

func FromComponent(from common.ApplicationComponent) (apis.Component, error) {
	m := &MysqlExporterServerComponent{}
	return m.FromComponent(from)
}

func (m *MysqlExporterServerComponent) SetTraits(traits ...apis.Trait) *MysqlExporterServerComponent {
	for _, addTrait := range traits {
		found := false
		for i, _t := range m.Base.Traits {
			if _t.DefType() == addTrait.DefType() {
				m.Base.Traits[i] = addTrait
				found = true
				break
			}
			if !found {
				m.Base.Traits = append(m.Base.Traits, addTrait)
			}
		}
	}
	return m
}

func (m *MysqlExporterServerComponent) GetTrait(typ string) apis.Trait {
	for _, _t := range m.Base.Traits {
		if _t.DefType() == typ {
			return _t
		}
	}
	return nil
}

func (m *MysqlExporterServerComponent) GetAllTraits() []apis.Trait {
	return m.Base.Traits
}

func (m *MysqlExporterServerComponent) ComponentName() string {
	return m.Base.Name
}

func (m *MysqlExporterServerComponent) DefType() string {
	return MysqlExporterServerType
}

func (m *MysqlExporterServerComponent) DependsOn(dependsOn []string) *MysqlExporterServerComponent {
	m.Base.DependsOn = dependsOn
	return m
}

func (m *MysqlExporterServerComponent) Inputs(input common.StepInputs) *MysqlExporterServerComponent {
	m.Base.Inputs = input
	return m
}

func (m *MysqlExporterServerComponent) Outputs(output common.StepOutputs) *MysqlExporterServerComponent {
	m.Base.Outputs = output
	return m
}

func (m *MysqlExporterServerComponent) AddDependsOn(dependsOn string) *MysqlExporterServerComponent {
	m.Base.DependsOn = append(m.Base.DependsOn, dependsOn)
	return m
}
