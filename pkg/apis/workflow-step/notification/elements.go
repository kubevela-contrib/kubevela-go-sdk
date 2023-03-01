/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Elements type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Elements{}

// Elements struct for Elements
type Elements struct {
	ActionId             *string               `json:"action_id,omitempty"`
	AltText              *string               `json:"alt_text,omitempty"`
	Confirm              *Confirm              `json:"confirm,omitempty"`
	DispatchActionConfig *DispatchActionConfig `json:"dispatch_action_config,omitempty"`
	ImageUrl             *string               `json:"image_url,omitempty"`
	InitialDate          *string               `json:"initial_date,omitempty"`
	InitialOptions       []Option              `json:"initial_options,omitempty"`
	InitialTime          *string               `json:"initial_time,omitempty"`
	InitialValue         *string               `json:"initial_value,omitempty"`
	MaxLength            *int32                `json:"max_length,omitempty"`
	MaxSelectedItems     *int32                `json:"max_selected_items,omitempty"`
	MinLength            *int32                `json:"min_length,omitempty"`
	Multiline            *bool                 `json:"multiline,omitempty"`
	OptionGroups         []Option              `json:"option_groups,omitempty"`
	Options              []Option              `json:"options,omitempty"`
	Placeholder          *TextType             `json:"placeholder,omitempty"`
	Style                *string               `json:"style,omitempty"`
	Text                 *TextType             `json:"text,omitempty"`
	Type                 *string               `json:"type,omitempty"`
	Url                  *string               `json:"url,omitempty"`
	Value                *string               `json:"value,omitempty"`
}

// NewElementsWith instantiates a new Elements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewElementsWith() *Elements {
	this := Elements{}
	return &this
}

// NewElements instantiates a new Elements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewElements() *Elements {
	this := Elements{}
	return &this
}

// NewElementss converts a list Elements pointers to objects.
// This is helpful when the SetElements requires a list of objects
func NewElementsList(ps ...*Elements) []Elements {
	objs := []Elements{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetActionId returns the ActionId field value if set, zero value otherwise.
func (o *Elements) GetActionId() string {
	if o == nil || utils.IsNil(o.ActionId) {
		var ret string
		return ret
	}
	return *o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetActionIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ActionId) {
		return nil, false
	}
	return o.ActionId, true
}

// HasActionId returns a boolean if a field has been set.
func (o *Elements) HasActionId() bool {
	if o != nil && !utils.IsNil(o.ActionId) {
		return true
	}

	return false
}

// SetActionId gets a reference to the given string and assigns it to the actionId field.
// ActionId:
func (o *Elements) SetActionId(v string) *Elements {
	o.ActionId = &v
	return o
}

// GetAltText returns the AltText field value if set, zero value otherwise.
func (o *Elements) GetAltText() string {
	if o == nil || utils.IsNil(o.AltText) {
		var ret string
		return ret
	}
	return *o.AltText
}

// GetAltTextOk returns a tuple with the AltText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetAltTextOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AltText) {
		return nil, false
	}
	return o.AltText, true
}

// HasAltText returns a boolean if a field has been set.
func (o *Elements) HasAltText() bool {
	if o != nil && !utils.IsNil(o.AltText) {
		return true
	}

	return false
}

// SetAltText gets a reference to the given string and assigns it to the altText field.
// AltText:
func (o *Elements) SetAltText(v string) *Elements {
	o.AltText = &v
	return o
}

// GetConfirm returns the Confirm field value if set, zero value otherwise.
func (o *Elements) GetConfirm() Confirm {
	if o == nil || utils.IsNil(o.Confirm) {
		var ret Confirm
		return ret
	}
	return *o.Confirm
}

// GetConfirmOk returns a tuple with the Confirm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetConfirmOk() (*Confirm, bool) {
	if o == nil || utils.IsNil(o.Confirm) {
		return nil, false
	}
	return o.Confirm, true
}

// HasConfirm returns a boolean if a field has been set.
func (o *Elements) HasConfirm() bool {
	if o != nil && !utils.IsNil(o.Confirm) {
		return true
	}

	return false
}

// SetConfirm gets a reference to the given Confirm and assigns it to the confirm field.
// Confirm:
func (o *Elements) SetConfirm(v Confirm) *Elements {
	o.Confirm = &v
	return o
}

// GetDispatchActionConfig returns the DispatchActionConfig field value if set, zero value otherwise.
func (o *Elements) GetDispatchActionConfig() DispatchActionConfig {
	if o == nil || utils.IsNil(o.DispatchActionConfig) {
		var ret DispatchActionConfig
		return ret
	}
	return *o.DispatchActionConfig
}

// GetDispatchActionConfigOk returns a tuple with the DispatchActionConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetDispatchActionConfigOk() (*DispatchActionConfig, bool) {
	if o == nil || utils.IsNil(o.DispatchActionConfig) {
		return nil, false
	}
	return o.DispatchActionConfig, true
}

// HasDispatchActionConfig returns a boolean if a field has been set.
func (o *Elements) HasDispatchActionConfig() bool {
	if o != nil && !utils.IsNil(o.DispatchActionConfig) {
		return true
	}

	return false
}

// SetDispatchActionConfig gets a reference to the given DispatchActionConfig and assigns it to the dispatchActionConfig field.
// DispatchActionConfig:
func (o *Elements) SetDispatchActionConfig(v DispatchActionConfig) *Elements {
	o.DispatchActionConfig = &v
	return o
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *Elements) GetImageUrl() string {
	if o == nil || utils.IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetImageUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *Elements) HasImageUrl() bool {
	if o != nil && !utils.IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the imageUrl field.
// ImageUrl:
func (o *Elements) SetImageUrl(v string) *Elements {
	o.ImageUrl = &v
	return o
}

// GetInitialDate returns the InitialDate field value if set, zero value otherwise.
func (o *Elements) GetInitialDate() string {
	if o == nil || utils.IsNil(o.InitialDate) {
		var ret string
		return ret
	}
	return *o.InitialDate
}

// GetInitialDateOk returns a tuple with the InitialDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetInitialDateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.InitialDate) {
		return nil, false
	}
	return o.InitialDate, true
}

// HasInitialDate returns a boolean if a field has been set.
func (o *Elements) HasInitialDate() bool {
	if o != nil && !utils.IsNil(o.InitialDate) {
		return true
	}

	return false
}

// SetInitialDate gets a reference to the given string and assigns it to the initialDate field.
// InitialDate:
func (o *Elements) SetInitialDate(v string) *Elements {
	o.InitialDate = &v
	return o
}

// GetInitialOptions returns the InitialOptions field value if set, zero value otherwise.
func (o *Elements) GetInitialOptions() []Option {
	if o == nil || utils.IsNil(o.InitialOptions) {
		var ret []Option
		return ret
	}
	return o.InitialOptions
}

// GetInitialOptionsOk returns a tuple with the InitialOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetInitialOptionsOk() ([]Option, bool) {
	if o == nil || utils.IsNil(o.InitialOptions) {
		return nil, false
	}
	return o.InitialOptions, true
}

// HasInitialOptions returns a boolean if a field has been set.
func (o *Elements) HasInitialOptions() bool {
	if o != nil && !utils.IsNil(o.InitialOptions) {
		return true
	}

	return false
}

// SetInitialOptions gets a reference to the given []Option and assigns it to the initialOptions field.
// InitialOptions:
func (o *Elements) SetInitialOptions(v []Option) *Elements {
	o.InitialOptions = v
	return o
}

// GetInitialTime returns the InitialTime field value if set, zero value otherwise.
func (o *Elements) GetInitialTime() string {
	if o == nil || utils.IsNil(o.InitialTime) {
		var ret string
		return ret
	}
	return *o.InitialTime
}

// GetInitialTimeOk returns a tuple with the InitialTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetInitialTimeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.InitialTime) {
		return nil, false
	}
	return o.InitialTime, true
}

// HasInitialTime returns a boolean if a field has been set.
func (o *Elements) HasInitialTime() bool {
	if o != nil && !utils.IsNil(o.InitialTime) {
		return true
	}

	return false
}

// SetInitialTime gets a reference to the given string and assigns it to the initialTime field.
// InitialTime:
func (o *Elements) SetInitialTime(v string) *Elements {
	o.InitialTime = &v
	return o
}

// GetInitialValue returns the InitialValue field value if set, zero value otherwise.
func (o *Elements) GetInitialValue() string {
	if o == nil || utils.IsNil(o.InitialValue) {
		var ret string
		return ret
	}
	return *o.InitialValue
}

// GetInitialValueOk returns a tuple with the InitialValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetInitialValueOk() (*string, bool) {
	if o == nil || utils.IsNil(o.InitialValue) {
		return nil, false
	}
	return o.InitialValue, true
}

// HasInitialValue returns a boolean if a field has been set.
func (o *Elements) HasInitialValue() bool {
	if o != nil && !utils.IsNil(o.InitialValue) {
		return true
	}

	return false
}

// SetInitialValue gets a reference to the given string and assigns it to the initialValue field.
// InitialValue:
func (o *Elements) SetInitialValue(v string) *Elements {
	o.InitialValue = &v
	return o
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise.
func (o *Elements) GetMaxLength() int32 {
	if o == nil || utils.IsNil(o.MaxLength) {
		var ret int32
		return ret
	}
	return *o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetMaxLengthOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.MaxLength) {
		return nil, false
	}
	return o.MaxLength, true
}

// HasMaxLength returns a boolean if a field has been set.
func (o *Elements) HasMaxLength() bool {
	if o != nil && !utils.IsNil(o.MaxLength) {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given int32 and assigns it to the maxLength field.
// MaxLength:
func (o *Elements) SetMaxLength(v int32) *Elements {
	o.MaxLength = &v
	return o
}

// GetMaxSelectedItems returns the MaxSelectedItems field value if set, zero value otherwise.
func (o *Elements) GetMaxSelectedItems() int32 {
	if o == nil || utils.IsNil(o.MaxSelectedItems) {
		var ret int32
		return ret
	}
	return *o.MaxSelectedItems
}

// GetMaxSelectedItemsOk returns a tuple with the MaxSelectedItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetMaxSelectedItemsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.MaxSelectedItems) {
		return nil, false
	}
	return o.MaxSelectedItems, true
}

// HasMaxSelectedItems returns a boolean if a field has been set.
func (o *Elements) HasMaxSelectedItems() bool {
	if o != nil && !utils.IsNil(o.MaxSelectedItems) {
		return true
	}

	return false
}

// SetMaxSelectedItems gets a reference to the given int32 and assigns it to the maxSelectedItems field.
// MaxSelectedItems:
func (o *Elements) SetMaxSelectedItems(v int32) *Elements {
	o.MaxSelectedItems = &v
	return o
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *Elements) GetMinLength() int32 {
	if o == nil || utils.IsNil(o.MinLength) {
		var ret int32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetMinLengthOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.MinLength) {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *Elements) HasMinLength() bool {
	if o != nil && !utils.IsNil(o.MinLength) {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given int32 and assigns it to the minLength field.
// MinLength:
func (o *Elements) SetMinLength(v int32) *Elements {
	o.MinLength = &v
	return o
}

// GetMultiline returns the Multiline field value if set, zero value otherwise.
func (o *Elements) GetMultiline() bool {
	if o == nil || utils.IsNil(o.Multiline) {
		var ret bool
		return ret
	}
	return *o.Multiline
}

// GetMultilineOk returns a tuple with the Multiline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetMultilineOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Multiline) {
		return nil, false
	}
	return o.Multiline, true
}

// HasMultiline returns a boolean if a field has been set.
func (o *Elements) HasMultiline() bool {
	if o != nil && !utils.IsNil(o.Multiline) {
		return true
	}

	return false
}

// SetMultiline gets a reference to the given bool and assigns it to the multiline field.
// Multiline:
func (o *Elements) SetMultiline(v bool) *Elements {
	o.Multiline = &v
	return o
}

// GetOptionGroups returns the OptionGroups field value if set, zero value otherwise.
func (o *Elements) GetOptionGroups() []Option {
	if o == nil || utils.IsNil(o.OptionGroups) {
		var ret []Option
		return ret
	}
	return o.OptionGroups
}

// GetOptionGroupsOk returns a tuple with the OptionGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetOptionGroupsOk() ([]Option, bool) {
	if o == nil || utils.IsNil(o.OptionGroups) {
		return nil, false
	}
	return o.OptionGroups, true
}

// HasOptionGroups returns a boolean if a field has been set.
func (o *Elements) HasOptionGroups() bool {
	if o != nil && !utils.IsNil(o.OptionGroups) {
		return true
	}

	return false
}

// SetOptionGroups gets a reference to the given []Option and assigns it to the optionGroups field.
// OptionGroups:
func (o *Elements) SetOptionGroups(v []Option) *Elements {
	o.OptionGroups = v
	return o
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *Elements) GetOptions() []Option {
	if o == nil || utils.IsNil(o.Options) {
		var ret []Option
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetOptionsOk() ([]Option, bool) {
	if o == nil || utils.IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *Elements) HasOptions() bool {
	if o != nil && !utils.IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []Option and assigns it to the options field.
// Options:
func (o *Elements) SetOptions(v []Option) *Elements {
	o.Options = v
	return o
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *Elements) GetPlaceholder() TextType {
	if o == nil || utils.IsNil(o.Placeholder) {
		var ret TextType
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetPlaceholderOk() (*TextType, bool) {
	if o == nil || utils.IsNil(o.Placeholder) {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *Elements) HasPlaceholder() bool {
	if o != nil && !utils.IsNil(o.Placeholder) {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given TextType and assigns it to the placeholder field.
// Placeholder:
func (o *Elements) SetPlaceholder(v TextType) *Elements {
	o.Placeholder = &v
	return o
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *Elements) GetStyle() string {
	if o == nil || utils.IsNil(o.Style) {
		var ret string
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetStyleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Style) {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *Elements) HasStyle() bool {
	if o != nil && !utils.IsNil(o.Style) {
		return true
	}

	return false
}

// SetStyle gets a reference to the given string and assigns it to the style field.
// Style:
func (o *Elements) SetStyle(v string) *Elements {
	o.Style = &v
	return o
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Elements) GetText() TextType {
	if o == nil || utils.IsNil(o.Text) {
		var ret TextType
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetTextOk() (*TextType, bool) {
	if o == nil || utils.IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Elements) HasText() bool {
	if o != nil && !utils.IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given TextType and assigns it to the text field.
// Text:
func (o *Elements) SetText(v TextType) *Elements {
	o.Text = &v
	return o
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Elements) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Elements) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the type_ field.
// Type:
func (o *Elements) SetType(v string) *Elements {
	o.Type = &v
	return o
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Elements) GetUrl() string {
	if o == nil || utils.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Elements) HasUrl() bool {
	if o != nil && !utils.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the url field.
// Url:
func (o *Elements) SetUrl(v string) *Elements {
	o.Url = &v
	return o
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Elements) GetValue() string {
	if o == nil || utils.IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elements) GetValueOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Elements) HasValue() bool {
	if o != nil && !utils.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the value field.
// Value:
func (o *Elements) SetValue(v string) *Elements {
	o.Value = &v
	return o
}

func (o Elements) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Elements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ActionId) {
		toSerialize["action_id"] = o.ActionId
	}
	if !utils.IsNil(o.AltText) {
		toSerialize["alt_text"] = o.AltText
	}
	if !utils.IsNil(o.Confirm) {
		toSerialize["confirm"] = o.Confirm
	}
	if !utils.IsNil(o.DispatchActionConfig) {
		toSerialize["dispatch_action_config"] = o.DispatchActionConfig
	}
	if !utils.IsNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	if !utils.IsNil(o.InitialDate) {
		toSerialize["initial_date"] = o.InitialDate
	}
	if !utils.IsNil(o.InitialOptions) {
		toSerialize["initial_options"] = o.InitialOptions
	}
	if !utils.IsNil(o.InitialTime) {
		toSerialize["initial_time"] = o.InitialTime
	}
	if !utils.IsNil(o.InitialValue) {
		toSerialize["initial_value"] = o.InitialValue
	}
	if !utils.IsNil(o.MaxLength) {
		toSerialize["max_length"] = o.MaxLength
	}
	if !utils.IsNil(o.MaxSelectedItems) {
		toSerialize["max_selected_items"] = o.MaxSelectedItems
	}
	if !utils.IsNil(o.MinLength) {
		toSerialize["min_length"] = o.MinLength
	}
	if !utils.IsNil(o.Multiline) {
		toSerialize["multiline"] = o.Multiline
	}
	if !utils.IsNil(o.OptionGroups) {
		toSerialize["option_groups"] = o.OptionGroups
	}
	if !utils.IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !utils.IsNil(o.Placeholder) {
		toSerialize["placeholder"] = o.Placeholder
	}
	if !utils.IsNil(o.Style) {
		toSerialize["style"] = o.Style
	}
	if !utils.IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !utils.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableElements struct {
	value *Elements
	isSet bool
}

func (v NullableElements) Get() *Elements {
	return v.value
}

func (v *NullableElements) Set(val *Elements) {
	v.value = val
	v.isSet = true
}

func (v NullableElements) IsSet() bool {
	return v.isSet
}

func (v *NullableElements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableElements(val *Elements) *NullableElements {
	return &NullableElements{value: val, isSet: true}
}

func (v NullableElements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableElements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
