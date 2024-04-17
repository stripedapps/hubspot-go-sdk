/*
Line Items

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Error type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Error{}

// Error struct for Error
type Error struct {
	// A specific category that contains more specific detail about the error
	SubCategory *string `json:"subCategory,omitempty"`
	// Context about the error condition
	Context *map[string][]string `json:"context,omitempty"`
	// A unique identifier for the request. Include this value with any error reports or support tickets
	CorrelationId string `json:"correlationId"`
	// A map of link names to associated URIs containing documentation about the error or recommended remediation steps
	Links *map[string]string `json:"links,omitempty"`
	// A human readable message describing the error along with remediation steps where appropriate
	Message string `json:"message"`
	// The error category
	Category string `json:"category"`
	// further information about the error
	Errors []ErrorDetail `json:"errors,omitempty"`
}

// NewError instantiates a new Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError(correlationId string, message string, category string) *Error {
	this := Error{}
	this.CorrelationId = correlationId
	this.Message = message
	this.Category = category
	return &this
}

// NewErrorWithDefaults instantiates a new Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorWithDefaults() *Error {
	this := Error{}
	return &this
}

// GetSubCategory returns the SubCategory field value if set, zero value otherwise.
func (o *Error) GetSubCategory() string {
	if o == nil || IsNil(o.SubCategory) {
		var ret string
		return ret
	}
	return *o.SubCategory
}

// GetSubCategoryOk returns a tuple with the SubCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetSubCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.SubCategory) {
		return nil, false
	}
	return o.SubCategory, true
}

// HasSubCategory returns a boolean if a field has been set.
func (o *Error) HasSubCategory() bool {
	if o != nil && !IsNil(o.SubCategory) {
		return true
	}

	return false
}

// SetSubCategory gets a reference to the given string and assigns it to the SubCategory field.
func (o *Error) SetSubCategory(v string) {
	o.SubCategory = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *Error) GetContext() map[string][]string {
	if o == nil || IsNil(o.Context) {
		var ret map[string][]string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetContextOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *Error) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string][]string and assigns it to the Context field.
func (o *Error) SetContext(v map[string][]string) {
	o.Context = &v
}

// GetCorrelationId returns the CorrelationId field value
func (o *Error) GetCorrelationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value
// and a boolean to check if the value has been set.
func (o *Error) GetCorrelationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CorrelationId, true
}

// SetCorrelationId sets field value
func (o *Error) SetCorrelationId(v string) {
	o.CorrelationId = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Error) GetLinks() map[string]string {
	if o == nil || IsNil(o.Links) {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetLinksOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Error) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *Error) SetLinks(v map[string]string) {
	o.Links = &v
}

// GetMessage returns the Message field value
func (o *Error) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Error) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Error) SetMessage(v string) {
	o.Message = v
}

// GetCategory returns the Category field value
func (o *Error) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *Error) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *Error) SetCategory(v string) {
	o.Category = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Error) GetErrors() []ErrorDetail {
	if o == nil || IsNil(o.Errors) {
		var ret []ErrorDetail
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetErrorsOk() ([]ErrorDetail, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Error) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ErrorDetail and assigns it to the Errors field.
func (o *Error) SetErrors(v []ErrorDetail) {
	o.Errors = v
}

func (o Error) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Error) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubCategory) {
		toSerialize["subCategory"] = o.SubCategory
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	toSerialize["correlationId"] = o.CorrelationId
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["message"] = o.Message
	toSerialize["category"] = o.Category
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableError struct {
	value *Error
	isSet bool
}

func (v NullableError) Get() *Error {
	return v.value
}

func (v *NullableError) Set(val *Error) {
	v.value = val
	v.isSet = true
}

func (v NullableError) IsSet() bool {
	return v.isSet
}

func (v *NullableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError(val *Error) *NullableError {
	return &NullableError{value: val, isSet: true}
}

func (v NullableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


