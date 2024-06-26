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

// checks if the NextPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NextPage{}

// NextPage struct for NextPage
type NextPage struct {
	Link *string `json:"link,omitempty"`
	After string `json:"after"`
}

// NewNextPage instantiates a new NextPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNextPage(after string) *NextPage {
	this := NextPage{}
	this.After = after
	return &this
}

// NewNextPageWithDefaults instantiates a new NextPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNextPageWithDefaults() *NextPage {
	this := NextPage{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *NextPage) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextPage) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *NextPage) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *NextPage) SetLink(v string) {
	o.Link = &v
}

// GetAfter returns the After field value
func (o *NextPage) GetAfter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.After
}

// GetAfterOk returns a tuple with the After field value
// and a boolean to check if the value has been set.
func (o *NextPage) GetAfterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.After, true
}

// SetAfter sets field value
func (o *NextPage) SetAfter(v string) {
	o.After = v
}

func (o NextPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NextPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	toSerialize["after"] = o.After
	return toSerialize, nil
}

type NullableNextPage struct {
	value *NextPage
	isSet bool
}

func (v NullableNextPage) Get() *NextPage {
	return v.value
}

func (v *NullableNextPage) Set(val *NextPage) {
	v.value = val
	v.isSet = true
}

func (v NullableNextPage) IsSet() bool {
	return v.isSet
}

func (v *NullableNextPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNextPage(val *NextPage) *NullableNextPage {
	return &NullableNextPage{value: val, isSet: true}
}

func (v NullableNextPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNextPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


