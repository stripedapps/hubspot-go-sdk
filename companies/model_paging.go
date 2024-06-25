/*
Companies

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Paging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Paging{}

// Paging struct for Paging
type Paging struct {
	Next *NextPage `json:"next,omitempty"`
	Prev *PreviousPage `json:"prev,omitempty"`
}

// NewPaging instantiates a new Paging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaging() *Paging {
	this := Paging{}
	return &this
}

// NewPagingWithDefaults instantiates a new Paging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagingWithDefaults() *Paging {
	this := Paging{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *Paging) GetNext() NextPage {
	if o == nil || IsNil(o.Next) {
		var ret NextPage
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paging) GetNextOk() (*NextPage, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *Paging) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given NextPage and assigns it to the Next field.
func (o *Paging) SetNext(v NextPage) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *Paging) GetPrev() PreviousPage {
	if o == nil || IsNil(o.Prev) {
		var ret PreviousPage
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paging) GetPrevOk() (*PreviousPage, bool) {
	if o == nil || IsNil(o.Prev) {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *Paging) HasPrev() bool {
	if o != nil && !IsNil(o.Prev) {
		return true
	}

	return false
}

// SetPrev gets a reference to the given PreviousPage and assigns it to the Prev field.
func (o *Paging) SetPrev(v PreviousPage) {
	o.Prev = &v
}

func (o Paging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Paging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Prev) {
		toSerialize["prev"] = o.Prev
	}
	return toSerialize, nil
}

type NullablePaging struct {
	value *Paging
	isSet bool
}

func (v NullablePaging) Get() *Paging {
	return v.value
}

func (v *NullablePaging) Set(val *Paging) {
	v.value = val
	v.isSet = true
}

func (v NullablePaging) IsSet() bool {
	return v.isSet
}

func (v *NullablePaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaging(val *Paging) *NullablePaging {
	return &NullablePaging{value: val, isSet: true}
}

func (v NullablePaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


