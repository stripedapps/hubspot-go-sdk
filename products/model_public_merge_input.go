/*
Products

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PublicMergeInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicMergeInput{}

// PublicMergeInput struct for PublicMergeInput
type PublicMergeInput struct {
	ObjectIdToMerge string `json:"objectIdToMerge"`
	PrimaryObjectId string `json:"primaryObjectId"`
}

// NewPublicMergeInput instantiates a new PublicMergeInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicMergeInput(objectIdToMerge string, primaryObjectId string) *PublicMergeInput {
	this := PublicMergeInput{}
	this.ObjectIdToMerge = objectIdToMerge
	this.PrimaryObjectId = primaryObjectId
	return &this
}

// NewPublicMergeInputWithDefaults instantiates a new PublicMergeInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicMergeInputWithDefaults() *PublicMergeInput {
	this := PublicMergeInput{}
	return &this
}

// GetObjectIdToMerge returns the ObjectIdToMerge field value
func (o *PublicMergeInput) GetObjectIdToMerge() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectIdToMerge
}

// GetObjectIdToMergeOk returns a tuple with the ObjectIdToMerge field value
// and a boolean to check if the value has been set.
func (o *PublicMergeInput) GetObjectIdToMergeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectIdToMerge, true
}

// SetObjectIdToMerge sets field value
func (o *PublicMergeInput) SetObjectIdToMerge(v string) {
	o.ObjectIdToMerge = v
}

// GetPrimaryObjectId returns the PrimaryObjectId field value
func (o *PublicMergeInput) GetPrimaryObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryObjectId
}

// GetPrimaryObjectIdOk returns a tuple with the PrimaryObjectId field value
// and a boolean to check if the value has been set.
func (o *PublicMergeInput) GetPrimaryObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryObjectId, true
}

// SetPrimaryObjectId sets field value
func (o *PublicMergeInput) SetPrimaryObjectId(v string) {
	o.PrimaryObjectId = v
}

func (o PublicMergeInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicMergeInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objectIdToMerge"] = o.ObjectIdToMerge
	toSerialize["primaryObjectId"] = o.PrimaryObjectId
	return toSerialize, nil
}

type NullablePublicMergeInput struct {
	value *PublicMergeInput
	isSet bool
}

func (v NullablePublicMergeInput) Get() *PublicMergeInput {
	return v.value
}

func (v *NullablePublicMergeInput) Set(val *PublicMergeInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicMergeInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicMergeInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicMergeInput(val *PublicMergeInput) *NullablePublicMergeInput {
	return &NullablePublicMergeInput{value: val, isSet: true}
}

func (v NullablePublicMergeInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicMergeInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

