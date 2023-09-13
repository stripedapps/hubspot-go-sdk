/*
Contacts

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the ValueWithTimestamp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueWithTimestamp{}

// ValueWithTimestamp struct for ValueWithTimestamp
type ValueWithTimestamp struct {
	Value string `json:"value"`
	Timestamp time.Time `json:"timestamp"`
	SourceType string `json:"sourceType"`
	SourceId *string `json:"sourceId,omitempty"`
	SourceLabel *string `json:"sourceLabel,omitempty"`
	UpdatedByUserId *int32 `json:"updatedByUserId,omitempty"`
}

// NewValueWithTimestamp instantiates a new ValueWithTimestamp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueWithTimestamp(value string, timestamp time.Time, sourceType string) *ValueWithTimestamp {
	this := ValueWithTimestamp{}
	this.Value = value
	this.Timestamp = timestamp
	this.SourceType = sourceType
	return &this
}

// NewValueWithTimestampWithDefaults instantiates a new ValueWithTimestamp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueWithTimestampWithDefaults() *ValueWithTimestamp {
	this := ValueWithTimestamp{}
	return &this
}

// GetValue returns the Value field value
func (o *ValueWithTimestamp) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ValueWithTimestamp) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ValueWithTimestamp) SetValue(v string) {
	o.Value = v
}

// GetTimestamp returns the Timestamp field value
func (o *ValueWithTimestamp) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ValueWithTimestamp) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ValueWithTimestamp) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetSourceType returns the SourceType field value
func (o *ValueWithTimestamp) GetSourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *ValueWithTimestamp) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *ValueWithTimestamp) SetSourceType(v string) {
	o.SourceType = v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *ValueWithTimestamp) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueWithTimestamp) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *ValueWithTimestamp) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *ValueWithTimestamp) SetSourceId(v string) {
	o.SourceId = &v
}

// GetSourceLabel returns the SourceLabel field value if set, zero value otherwise.
func (o *ValueWithTimestamp) GetSourceLabel() string {
	if o == nil || IsNil(o.SourceLabel) {
		var ret string
		return ret
	}
	return *o.SourceLabel
}

// GetSourceLabelOk returns a tuple with the SourceLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueWithTimestamp) GetSourceLabelOk() (*string, bool) {
	if o == nil || IsNil(o.SourceLabel) {
		return nil, false
	}
	return o.SourceLabel, true
}

// HasSourceLabel returns a boolean if a field has been set.
func (o *ValueWithTimestamp) HasSourceLabel() bool {
	if o != nil && !IsNil(o.SourceLabel) {
		return true
	}

	return false
}

// SetSourceLabel gets a reference to the given string and assigns it to the SourceLabel field.
func (o *ValueWithTimestamp) SetSourceLabel(v string) {
	o.SourceLabel = &v
}

// GetUpdatedByUserId returns the UpdatedByUserId field value if set, zero value otherwise.
func (o *ValueWithTimestamp) GetUpdatedByUserId() int32 {
	if o == nil || IsNil(o.UpdatedByUserId) {
		var ret int32
		return ret
	}
	return *o.UpdatedByUserId
}

// GetUpdatedByUserIdOk returns a tuple with the UpdatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueWithTimestamp) GetUpdatedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedByUserId) {
		return nil, false
	}
	return o.UpdatedByUserId, true
}

// HasUpdatedByUserId returns a boolean if a field has been set.
func (o *ValueWithTimestamp) HasUpdatedByUserId() bool {
	if o != nil && !IsNil(o.UpdatedByUserId) {
		return true
	}

	return false
}

// SetUpdatedByUserId gets a reference to the given int32 and assigns it to the UpdatedByUserId field.
func (o *ValueWithTimestamp) SetUpdatedByUserId(v int32) {
	o.UpdatedByUserId = &v
}

func (o ValueWithTimestamp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValueWithTimestamp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["sourceType"] = o.SourceType
	if !IsNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !IsNil(o.SourceLabel) {
		toSerialize["sourceLabel"] = o.SourceLabel
	}
	if !IsNil(o.UpdatedByUserId) {
		toSerialize["updatedByUserId"] = o.UpdatedByUserId
	}
	return toSerialize, nil
}

type NullableValueWithTimestamp struct {
	value *ValueWithTimestamp
	isSet bool
}

func (v NullableValueWithTimestamp) Get() *ValueWithTimestamp {
	return v.value
}

func (v *NullableValueWithTimestamp) Set(val *ValueWithTimestamp) {
	v.value = val
	v.isSet = true
}

func (v NullableValueWithTimestamp) IsSet() bool {
	return v.isSet
}

func (v *NullableValueWithTimestamp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueWithTimestamp(val *ValueWithTimestamp) *NullableValueWithTimestamp {
	return &NullableValueWithTimestamp{value: val, isSet: true}
}

func (v NullableValueWithTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueWithTimestamp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


