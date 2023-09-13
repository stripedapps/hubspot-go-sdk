/*
Contacts

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SimplePublicObjectInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplePublicObjectInput{}

// SimplePublicObjectInput struct for SimplePublicObjectInput
type SimplePublicObjectInput struct {
	Properties map[string]string `json:"properties"`
}

// NewSimplePublicObjectInput instantiates a new SimplePublicObjectInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplePublicObjectInput(properties map[string]string) *SimplePublicObjectInput {
	this := SimplePublicObjectInput{}
	this.Properties = properties
	return &this
}

// NewSimplePublicObjectInputWithDefaults instantiates a new SimplePublicObjectInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplePublicObjectInputWithDefaults() *SimplePublicObjectInput {
	this := SimplePublicObjectInput{}
	return &this
}

// GetProperties returns the Properties field value
func (o *SimplePublicObjectInput) GetProperties() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectInput) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *SimplePublicObjectInput) SetProperties(v map[string]string) {
	o.Properties = v
}

func (o SimplePublicObjectInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplePublicObjectInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableSimplePublicObjectInput struct {
	value *SimplePublicObjectInput
	isSet bool
}

func (v NullableSimplePublicObjectInput) Get() *SimplePublicObjectInput {
	return v.value
}

func (v *NullableSimplePublicObjectInput) Set(val *SimplePublicObjectInput) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplePublicObjectInput) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplePublicObjectInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplePublicObjectInput(val *SimplePublicObjectInput) *NullableSimplePublicObjectInput {
	return &NullableSimplePublicObjectInput{value: val, isSet: true}
}

func (v NullableSimplePublicObjectInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplePublicObjectInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


