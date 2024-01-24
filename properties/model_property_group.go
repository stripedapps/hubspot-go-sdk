/*
Properties

All HubSpot objects store data in default and custom properties. These endpoints provide access to read and modify object properties in HubSpot.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PropertyGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyGroup{}

// PropertyGroup An ID for a group of properties
type PropertyGroup struct {
	Archived bool `json:"archived"`
	// The internal property group name, which must be used when referencing the property group via the API.
	Name string `json:"name"`
	// Property groups are displayed in order starting with the lowest positive integer value. Values of -1 will cause the property group to be displayed after any positive values.
	DisplayOrder int32 `json:"displayOrder"`
	// A human-readable label that will be shown in HubSpot.
	Label string `json:"label"`
}

// NewPropertyGroup instantiates a new PropertyGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyGroup(archived bool, name string, displayOrder int32, label string) *PropertyGroup {
	this := PropertyGroup{}
	this.Archived = archived
	this.Name = name
	this.DisplayOrder = displayOrder
	this.Label = label
	return &this
}

// NewPropertyGroupWithDefaults instantiates a new PropertyGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyGroupWithDefaults() *PropertyGroup {
	this := PropertyGroup{}
	return &this
}

// GetArchived returns the Archived field value
func (o *PropertyGroup) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *PropertyGroup) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *PropertyGroup) SetArchived(v bool) {
	o.Archived = v
}

// GetName returns the Name field value
func (o *PropertyGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PropertyGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PropertyGroup) SetName(v string) {
	o.Name = v
}

// GetDisplayOrder returns the DisplayOrder field value
func (o *PropertyGroup) GetDisplayOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value
// and a boolean to check if the value has been set.
func (o *PropertyGroup) GetDisplayOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayOrder, true
}

// SetDisplayOrder sets field value
func (o *PropertyGroup) SetDisplayOrder(v int32) {
	o.DisplayOrder = v
}

// GetLabel returns the Label field value
func (o *PropertyGroup) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *PropertyGroup) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *PropertyGroup) SetLabel(v string) {
	o.Label = v
}

func (o PropertyGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["archived"] = o.Archived
	toSerialize["name"] = o.Name
	toSerialize["displayOrder"] = o.DisplayOrder
	toSerialize["label"] = o.Label
	return toSerialize, nil
}

type NullablePropertyGroup struct {
	value *PropertyGroup
	isSet bool
}

func (v NullablePropertyGroup) Get() *PropertyGroup {
	return v.value
}

func (v *NullablePropertyGroup) Set(val *PropertyGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyGroup(val *PropertyGroup) *NullablePropertyGroup {
	return &NullablePropertyGroup{value: val, isSet: true}
}

func (v NullablePropertyGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


