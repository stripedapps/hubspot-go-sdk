/*
Notes

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the BatchResponseSimplePublicObjectWithErrors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchResponseSimplePublicObjectWithErrors{}

// BatchResponseSimplePublicObjectWithErrors struct for BatchResponseSimplePublicObjectWithErrors
type BatchResponseSimplePublicObjectWithErrors struct {
	CompletedAt time.Time `json:"completedAt"`
	NumErrors *int32 `json:"numErrors,omitempty"`
	RequestedAt *time.Time `json:"requestedAt,omitempty"`
	StartedAt time.Time `json:"startedAt"`
	Links *map[string]string `json:"links,omitempty"`
	Results []SimplePublicObject `json:"results"`
	Errors []StandardError `json:"errors,omitempty"`
	Status string `json:"status"`
}

type _BatchResponseSimplePublicObjectWithErrors BatchResponseSimplePublicObjectWithErrors

// NewBatchResponseSimplePublicObjectWithErrors instantiates a new BatchResponseSimplePublicObjectWithErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchResponseSimplePublicObjectWithErrors(completedAt time.Time, startedAt time.Time, results []SimplePublicObject, status string) *BatchResponseSimplePublicObjectWithErrors {
	this := BatchResponseSimplePublicObjectWithErrors{}
	this.CompletedAt = completedAt
	this.StartedAt = startedAt
	this.Results = results
	this.Status = status
	return &this
}

// NewBatchResponseSimplePublicObjectWithErrorsWithDefaults instantiates a new BatchResponseSimplePublicObjectWithErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchResponseSimplePublicObjectWithErrorsWithDefaults() *BatchResponseSimplePublicObjectWithErrors {
	this := BatchResponseSimplePublicObjectWithErrors{}
	return &this
}

// GetCompletedAt returns the CompletedAt field value
func (o *BatchResponseSimplePublicObjectWithErrors) GetCompletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSimplePublicObjectWithErrors) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedAt, true
}

// SetCompletedAt sets field value
func (o *BatchResponseSimplePublicObjectWithErrors) SetCompletedAt(v time.Time) {
	o.CompletedAt = v
}

// GetNumErrors returns the NumErrors field value if set, zero value otherwise.
func (o *BatchResponseSimplePublicObjectWithErrors) GetNumErrors() int32 {
	if o == nil || IsNil(o.NumErrors) {
		var ret int32
		return ret
	}
	return *o.NumErrors
}

// GetNumErrorsOk returns a tuple with the NumErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseSimplePublicObjectWithErrors) GetNumErrorsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumErrors) {
		return nil, false
	}
	return o.NumErrors, true
}

// HasNumErrors returns a boolean if a field has been set.
func (o *BatchResponseSimplePublicObjectWithErrors) HasNumErrors() bool {
	if o != nil && !IsNil(o.NumErrors) {
		return true
	}

	return false
}

// SetNumErrors gets a reference to the given int32 and assigns it to the NumErrors field.
func (o *BatchResponseSimplePublicObjectWithErrors) SetNumErrors(v int32) {
	o.NumErrors = &v
}

// GetRequestedAt returns the RequestedAt field value if set, zero value otherwise.
func (o *BatchResponseSimplePublicObjectWithErrors) GetRequestedAt() time.Time {
	if o == nil || IsNil(o.RequestedAt) {
		var ret time.Time
		return ret
	}
	return *o.RequestedAt
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseSimplePublicObjectWithErrors) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RequestedAt) {
		return nil, false
	}
	return o.RequestedAt, true
}

// HasRequestedAt returns a boolean if a field has been set.
func (o *BatchResponseSimplePublicObjectWithErrors) HasRequestedAt() bool {
	if o != nil && !IsNil(o.RequestedAt) {
		return true
	}

	return false
}

// SetRequestedAt gets a reference to the given time.Time and assigns it to the RequestedAt field.
func (o *BatchResponseSimplePublicObjectWithErrors) SetRequestedAt(v time.Time) {
	o.RequestedAt = &v
}

// GetStartedAt returns the StartedAt field value
func (o *BatchResponseSimplePublicObjectWithErrors) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSimplePublicObjectWithErrors) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *BatchResponseSimplePublicObjectWithErrors) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BatchResponseSimplePublicObjectWithErrors) GetLinks() map[string]string {
	if o == nil || IsNil(o.Links) {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseSimplePublicObjectWithErrors) GetLinksOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BatchResponseSimplePublicObjectWithErrors) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *BatchResponseSimplePublicObjectWithErrors) SetLinks(v map[string]string) {
	o.Links = &v
}

// GetResults returns the Results field value
func (o *BatchResponseSimplePublicObjectWithErrors) GetResults() []SimplePublicObject {
	if o == nil {
		var ret []SimplePublicObject
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSimplePublicObjectWithErrors) GetResultsOk() ([]SimplePublicObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *BatchResponseSimplePublicObjectWithErrors) SetResults(v []SimplePublicObject) {
	o.Results = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BatchResponseSimplePublicObjectWithErrors) GetErrors() []StandardError {
	if o == nil || IsNil(o.Errors) {
		var ret []StandardError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseSimplePublicObjectWithErrors) GetErrorsOk() ([]StandardError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BatchResponseSimplePublicObjectWithErrors) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []StandardError and assigns it to the Errors field.
func (o *BatchResponseSimplePublicObjectWithErrors) SetErrors(v []StandardError) {
	o.Errors = v
}

// GetStatus returns the Status field value
func (o *BatchResponseSimplePublicObjectWithErrors) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSimplePublicObjectWithErrors) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BatchResponseSimplePublicObjectWithErrors) SetStatus(v string) {
	o.Status = v
}

func (o BatchResponseSimplePublicObjectWithErrors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchResponseSimplePublicObjectWithErrors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["completedAt"] = o.CompletedAt
	if !IsNil(o.NumErrors) {
		toSerialize["numErrors"] = o.NumErrors
	}
	if !IsNil(o.RequestedAt) {
		toSerialize["requestedAt"] = o.RequestedAt
	}
	toSerialize["startedAt"] = o.StartedAt
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["results"] = o.Results
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *BatchResponseSimplePublicObjectWithErrors) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"completedAt",
		"startedAt",
		"results",
		"status",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBatchResponseSimplePublicObjectWithErrors := _BatchResponseSimplePublicObjectWithErrors{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchResponseSimplePublicObjectWithErrors)

	if err != nil {
		return err
	}

	*o = BatchResponseSimplePublicObjectWithErrors(varBatchResponseSimplePublicObjectWithErrors)

	return err
}

type NullableBatchResponseSimplePublicObjectWithErrors struct {
	value *BatchResponseSimplePublicObjectWithErrors
	isSet bool
}

func (v NullableBatchResponseSimplePublicObjectWithErrors) Get() *BatchResponseSimplePublicObjectWithErrors {
	return v.value
}

func (v *NullableBatchResponseSimplePublicObjectWithErrors) Set(val *BatchResponseSimplePublicObjectWithErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchResponseSimplePublicObjectWithErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchResponseSimplePublicObjectWithErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchResponseSimplePublicObjectWithErrors(val *BatchResponseSimplePublicObjectWithErrors) *NullableBatchResponseSimplePublicObjectWithErrors {
	return &NullableBatchResponseSimplePublicObjectWithErrors{value: val, isSet: true}
}

func (v NullableBatchResponseSimplePublicObjectWithErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchResponseSimplePublicObjectWithErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


