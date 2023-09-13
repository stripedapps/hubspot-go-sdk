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

// checks if the CollectionResponseWithTotalSimplePublicObjectForwardPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResponseWithTotalSimplePublicObjectForwardPaging{}

// CollectionResponseWithTotalSimplePublicObjectForwardPaging struct for CollectionResponseWithTotalSimplePublicObjectForwardPaging
type CollectionResponseWithTotalSimplePublicObjectForwardPaging struct {
	Total int32 `json:"total"`
	Results []SimplePublicObject `json:"results"`
	Paging *ForwardPaging `json:"paging,omitempty"`
}

// NewCollectionResponseWithTotalSimplePublicObjectForwardPaging instantiates a new CollectionResponseWithTotalSimplePublicObjectForwardPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseWithTotalSimplePublicObjectForwardPaging(total int32, results []SimplePublicObject) *CollectionResponseWithTotalSimplePublicObjectForwardPaging {
	this := CollectionResponseWithTotalSimplePublicObjectForwardPaging{}
	this.Total = total
	this.Results = results
	return &this
}

// NewCollectionResponseWithTotalSimplePublicObjectForwardPagingWithDefaults instantiates a new CollectionResponseWithTotalSimplePublicObjectForwardPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseWithTotalSimplePublicObjectForwardPagingWithDefaults() *CollectionResponseWithTotalSimplePublicObjectForwardPaging {
	this := CollectionResponseWithTotalSimplePublicObjectForwardPaging{}
	return &this
}

// GetTotal returns the Total field value
func (o *CollectionResponseWithTotalSimplePublicObjectForwardPaging) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalSimplePublicObjectForwardPaging) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CollectionResponseWithTotalSimplePublicObjectForwardPaging) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *CollectionResponseWithTotalSimplePublicObjectForwardPaging) GetResults() []SimplePublicObject {
	if o == nil {
		var ret []SimplePublicObject
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalSimplePublicObjectForwardPaging) GetResultsOk() ([]SimplePublicObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseWithTotalSimplePublicObjectForwardPaging) SetResults(v []SimplePublicObject) {
	o.Results = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponseWithTotalSimplePublicObjectForwardPaging) GetPaging() ForwardPaging {
	if o == nil || IsNil(o.Paging) {
		var ret ForwardPaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalSimplePublicObjectForwardPaging) GetPagingOk() (*ForwardPaging, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponseWithTotalSimplePublicObjectForwardPaging) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ForwardPaging and assigns it to the Paging field.
func (o *CollectionResponseWithTotalSimplePublicObjectForwardPaging) SetPaging(v ForwardPaging) {
	o.Paging = &v
}

func (o CollectionResponseWithTotalSimplePublicObjectForwardPaging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResponseWithTotalSimplePublicObjectForwardPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return toSerialize, nil
}

type NullableCollectionResponseWithTotalSimplePublicObjectForwardPaging struct {
	value *CollectionResponseWithTotalSimplePublicObjectForwardPaging
	isSet bool
}

func (v NullableCollectionResponseWithTotalSimplePublicObjectForwardPaging) Get() *CollectionResponseWithTotalSimplePublicObjectForwardPaging {
	return v.value
}

func (v *NullableCollectionResponseWithTotalSimplePublicObjectForwardPaging) Set(val *CollectionResponseWithTotalSimplePublicObjectForwardPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseWithTotalSimplePublicObjectForwardPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseWithTotalSimplePublicObjectForwardPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseWithTotalSimplePublicObjectForwardPaging(val *CollectionResponseWithTotalSimplePublicObjectForwardPaging) *NullableCollectionResponseWithTotalSimplePublicObjectForwardPaging {
	return &NullableCollectionResponseWithTotalSimplePublicObjectForwardPaging{value: val, isSet: true}
}

func (v NullableCollectionResponseWithTotalSimplePublicObjectForwardPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseWithTotalSimplePublicObjectForwardPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


