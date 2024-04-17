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

// checks if the Filter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Filter{}

// Filter struct for Filter
type Filter struct {
	HighValue *string `json:"highValue,omitempty"`
	PropertyName string `json:"propertyName"`
	Values []string `json:"values,omitempty"`
	Value *string `json:"value,omitempty"`
	// null
	Operator string `json:"operator"`
}

// NewFilter instantiates a new Filter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilter(propertyName string, operator string) *Filter {
	this := Filter{}
	this.PropertyName = propertyName
	this.Operator = operator
	return &this
}

// NewFilterWithDefaults instantiates a new Filter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterWithDefaults() *Filter {
	this := Filter{}
	return &this
}

// GetHighValue returns the HighValue field value if set, zero value otherwise.
func (o *Filter) GetHighValue() string {
	if o == nil || IsNil(o.HighValue) {
		var ret string
		return ret
	}
	return *o.HighValue
}

// GetHighValueOk returns a tuple with the HighValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetHighValueOk() (*string, bool) {
	if o == nil || IsNil(o.HighValue) {
		return nil, false
	}
	return o.HighValue, true
}

// HasHighValue returns a boolean if a field has been set.
func (o *Filter) HasHighValue() bool {
	if o != nil && !IsNil(o.HighValue) {
		return true
	}

	return false
}

// SetHighValue gets a reference to the given string and assigns it to the HighValue field.
func (o *Filter) SetHighValue(v string) {
	o.HighValue = &v
}

// GetPropertyName returns the PropertyName field value
func (o *Filter) GetPropertyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PropertyName
}

// GetPropertyNameOk returns a tuple with the PropertyName field value
// and a boolean to check if the value has been set.
func (o *Filter) GetPropertyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertyName, true
}

// SetPropertyName sets field value
func (o *Filter) SetPropertyName(v string) {
	o.PropertyName = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *Filter) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *Filter) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *Filter) SetValues(v []string) {
	o.Values = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Filter) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Filter) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Filter) SetValue(v string) {
	o.Value = &v
}

// GetOperator returns the Operator field value
func (o *Filter) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *Filter) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *Filter) SetOperator(v string) {
	o.Operator = v
}

func (o Filter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Filter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HighValue) {
		toSerialize["highValue"] = o.HighValue
	}
	toSerialize["propertyName"] = o.PropertyName
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	toSerialize["operator"] = o.Operator
	return toSerialize, nil
}

type NullableFilter struct {
	value *Filter
	isSet bool
}

func (v NullableFilter) Get() *Filter {
	return v.value
}

func (v *NullableFilter) Set(val *Filter) {
	v.value = val
	v.isSet = true
}

func (v NullableFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilter(val *Filter) *NullableFilter {
	return &NullableFilter{value: val, isSet: true}
}

func (v NullableFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


