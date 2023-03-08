/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"
	"errors"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Block type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Block{}

// Block struct for Block
type Block struct {
	BlockId  *string    `json:"block_id,omitempty"`
	Elements []Elements `json:"elements,omitempty"`
	Type     *string    `json:"type"`
}

// NewBlockWith instantiates a new Block object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewBlockWith(type_ string) *Block {
	this := Block{}
	this.Type = &type_
	return &this
}

// NewBlockWithDefault instantiates a new Block object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockWithDefault() *Block {
	this := Block{}
	return &this
}

// NewBlock is short for NewBlockWithDefault which instantiates a new Block object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlock() *Block {
	return NewBlockWithDefault()
}

// NewBlockEmpty instantiates a new Block object with no properties set.
// This constructor will not assign any default values to properties.
func NewBlockEmpty() *Block {
	this := Block{}
	return &this
}

// NewBlocks converts a list Block pointers to objects.
// This is helpful when the SetBlock requires a list of objects
func NewBlockList(ps ...*Block) []Block {
	objs := []Block{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Block
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Block) Validate() error {
	if o.Type == nil {
		return errors.New("Type in Block must be set")
	}
	// validate all nested properties
	return nil
}

// GetBlockId returns the BlockId field value if set, zero value otherwise.
func (o *Block) GetBlockId() string {
	if o == nil || utils.IsNil(o.BlockId) {
		var ret string
		return ret
	}
	return *o.BlockId
}

// GetBlockIdOk returns a tuple with the BlockId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Block) GetBlockIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BlockId) {
		return nil, false
	}
	return o.BlockId, true
}

// HasBlockId returns a boolean if a field has been set.
func (o *Block) HasBlockId() bool {
	if o != nil && !utils.IsNil(o.BlockId) {
		return true
	}

	return false
}

// SetBlockId gets a reference to the given string and assigns it to the blockId field.
// BlockId:
func (o *Block) SetBlockId(v string) *Block {
	o.BlockId = &v
	return o
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *Block) GetElements() []Elements {
	if o == nil || utils.IsNil(o.Elements) {
		var ret []Elements
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Block) GetElementsOk() ([]Elements, bool) {
	if o == nil || utils.IsNil(o.Elements) {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *Block) HasElements() bool {
	if o != nil && !utils.IsNil(o.Elements) {
		return true
	}

	return false
}

// SetElements gets a reference to the given []Elements and assigns it to the elements field.
// Elements:
func (o *Block) SetElements(v []Elements) *Block {
	o.Elements = v
	return o
}

// GetType returns the Type field value
func (o *Block) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Block) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *Block) SetType(v string) *Block {
	o.Type = &v
	return o
}

func (o Block) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Block) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.BlockId) {
		toSerialize["block_id"] = o.BlockId
	}
	if !utils.IsNil(o.Elements) {
		toSerialize["elements"] = o.Elements
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableBlock struct {
	value *Block
	isSet bool
}

func (v *NullableBlock) Get() *Block {
	return v.value
}

func (v *NullableBlock) Set(val *Block) {
	v.value = val
	v.isSet = true
}

func (v *NullableBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlock(val *Block) *NullableBlock {
	return &NullableBlock{value: val, isSet: true}
}

func (v NullableBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
