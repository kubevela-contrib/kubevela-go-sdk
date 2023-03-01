/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package container_image

import (
	"encoding/json"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the PatchParams type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PatchParams{}

// PatchParams struct for PatchParams
type PatchParams struct {
	// Specify the name of the target container, if not set, use the component name
	ContainerName *string `json:"containerName,omitempty"`
	// Specify the image of the container
	Image *string `json:"image,omitempty"`
	// Specify the image pull policy of the container
	ImagePullPolicy *string `json:"imagePullPolicy,omitempty"`
}

// NewPatchParamsWith instantiates a new PatchParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchParamsWith() *PatchParams {
	this := PatchParams{}
	var containerName string = ""
	this.ContainerName = &containerName
	var imagePullPolicy string = ""
	this.ImagePullPolicy = &imagePullPolicy
	return &this
}

// NewPatchParams instantiates a new PatchParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchParams() *PatchParams {
	this := PatchParams{}
	var containerName string = ""
	this.ContainerName = &containerName
	var imagePullPolicy string = ""
	this.ImagePullPolicy = &imagePullPolicy
	return &this
}

// NewPatchParamss converts a list PatchParams pointers to objects.
// This is helpful when the SetPatchParams requires a list of objects
func NewPatchParamsList(ps ...*PatchParams) []PatchParams {
	objs := []PatchParams{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetContainerName returns the ContainerName field value if set, zero value otherwise.
func (o *PatchParams) GetContainerName() string {
	if o == nil || utils.IsNil(o.ContainerName) {
		var ret string
		return ret
	}
	return *o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchParams) GetContainerNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ContainerName) {
		return nil, false
	}
	return o.ContainerName, true
}

// HasContainerName returns a boolean if a field has been set.
func (o *PatchParams) HasContainerName() bool {
	if o != nil && !utils.IsNil(o.ContainerName) {
		return true
	}

	return false
}

// SetContainerName gets a reference to the given string and assigns it to the containerName field.
// ContainerName:  Specify the name of the target container, if not set, use the component name
func (o *PatchParams) SetContainerName(v string) *PatchParams {
	o.ContainerName = &v
	return o
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *PatchParams) GetImage() string {
	if o == nil || utils.IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchParams) GetImageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *PatchParams) HasImage() bool {
	if o != nil && !utils.IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the image field.
// Image:  Specify the image of the container
func (o *PatchParams) SetImage(v string) *PatchParams {
	o.Image = &v
	return o
}

// GetImagePullPolicy returns the ImagePullPolicy field value if set, zero value otherwise.
func (o *PatchParams) GetImagePullPolicy() string {
	if o == nil || utils.IsNil(o.ImagePullPolicy) {
		var ret string
		return ret
	}
	return *o.ImagePullPolicy
}

// GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchParams) GetImagePullPolicyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ImagePullPolicy) {
		return nil, false
	}
	return o.ImagePullPolicy, true
}

// HasImagePullPolicy returns a boolean if a field has been set.
func (o *PatchParams) HasImagePullPolicy() bool {
	if o != nil && !utils.IsNil(o.ImagePullPolicy) {
		return true
	}

	return false
}

// SetImagePullPolicy gets a reference to the given string and assigns it to the imagePullPolicy field.
// ImagePullPolicy:  Specify the image pull policy of the container
func (o *PatchParams) SetImagePullPolicy(v string) *PatchParams {
	o.ImagePullPolicy = &v
	return o
}

func (o PatchParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ContainerName) {
		toSerialize["containerName"] = o.ContainerName
	}
	if !utils.IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !utils.IsNil(o.ImagePullPolicy) {
		toSerialize["imagePullPolicy"] = o.ImagePullPolicy
	}
	return toSerialize, nil
}

type NullablePatchParams struct {
	value *PatchParams
	isSet bool
}

func (v NullablePatchParams) Get() *PatchParams {
	return v.value
}

func (v *NullablePatchParams) Set(val *PatchParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchParams(val *PatchParams) *NullablePatchParams {
	return &NullablePatchParams{value: val, isSet: true}
}

func (v NullablePatchParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
