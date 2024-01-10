/*
Webhooks

Provides a way for apps to subscribe to certain change events in HubSpot. Once configured, apps will receive event payloads containing details about the changes at a specified target URL. There can only be one target URL for receiving event notifications per app.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SubscriptionBatchUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionBatchUpdateRequest{}

// SubscriptionBatchUpdateRequest struct for SubscriptionBatchUpdateRequest
type SubscriptionBatchUpdateRequest struct {
	Active bool `json:"active"`
	Id int32 `json:"id"`
}

// NewSubscriptionBatchUpdateRequest instantiates a new SubscriptionBatchUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionBatchUpdateRequest(active bool, id int32) *SubscriptionBatchUpdateRequest {
	this := SubscriptionBatchUpdateRequest{}
	this.Active = active
	this.Id = id
	return &this
}

// NewSubscriptionBatchUpdateRequestWithDefaults instantiates a new SubscriptionBatchUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionBatchUpdateRequestWithDefaults() *SubscriptionBatchUpdateRequest {
	this := SubscriptionBatchUpdateRequest{}
	return &this
}

// GetActive returns the Active field value
func (o *SubscriptionBatchUpdateRequest) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *SubscriptionBatchUpdateRequest) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *SubscriptionBatchUpdateRequest) SetActive(v bool) {
	o.Active = v
}

// GetId returns the Id field value
func (o *SubscriptionBatchUpdateRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionBatchUpdateRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionBatchUpdateRequest) SetId(v int32) {
	o.Id = v
}

func (o SubscriptionBatchUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionBatchUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableSubscriptionBatchUpdateRequest struct {
	value *SubscriptionBatchUpdateRequest
	isSet bool
}

func (v NullableSubscriptionBatchUpdateRequest) Get() *SubscriptionBatchUpdateRequest {
	return v.value
}

func (v *NullableSubscriptionBatchUpdateRequest) Set(val *SubscriptionBatchUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionBatchUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionBatchUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionBatchUpdateRequest(val *SubscriptionBatchUpdateRequest) *NullableSubscriptionBatchUpdateRequest {
	return &NullableSubscriptionBatchUpdateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionBatchUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionBatchUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


