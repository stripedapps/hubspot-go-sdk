/*
Schemas

The CRM uses schemas to define how custom objects should store and represent information in the HubSpot CRM. Schemas define details about an object's type, properties, and associations. The schema can be uniquely identified by its **object type ID**.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ObjectTypePropertyCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectTypePropertyCreate{}

// ObjectTypePropertyCreate Defines a property to create.
type ObjectTypePropertyCreate struct {
	Hidden *bool `json:"hidden,omitempty"`
	// Controls how the property options will be sorted in the HubSpot UI.
	OptionSortStrategy *string `json:"optionSortStrategy,omitempty"`
	// The order that this property should be displayed in the HubSpot UI relative to other properties for this object type. Properties are displayed in order starting with the lowest positive integer value. A value of -1 will cause the property to be displayed **after** any positive values.
	DisplayOrder *int32 `json:"displayOrder,omitempty"`
	// A description of the property that will be shown as help text in HubSpot.
	Description *string `json:"description,omitempty"`
	// Whether the property will display the currency symbol in the HubSpot UI.
	ShowCurrencySymbol *bool `json:"showCurrencySymbol,omitempty"`
	// A human-readable property label that will be shown in HubSpot.
	Label string `json:"label"`
	// The data type of the property.
	Type string `json:"type"`
	// Whether the property can be used in a HubSpot form.
	FormField *bool `json:"formField,omitempty"`
	// The name of the group this property belongs to.
	GroupName *string `json:"groupName,omitempty"`
	// Defines the options this property will return, e.g. OWNER would return name of users on the portal.
	ReferencedObjectType *string `json:"referencedObjectType,omitempty"`
	// Controls how text properties are formatted in the HubSpot UI
	TextDisplayHint *string `json:"textDisplayHint,omitempty"`
	// The internal property name, which must be used when referencing the property from the API.
	Name string `json:"name"`
	// A list of available options for the property. This field is only required for enumerated properties.
	Options []OptionInput `json:"options,omitempty"`
	// Allow users to search for information entered to this field (limited to 3 properties)
	SearchableInGlobalSearch *bool `json:"searchableInGlobalSearch,omitempty"`
	// Controls how numeric properties are formatted in the HubSpot UI
	NumberDisplayHint *string `json:"numberDisplayHint,omitempty"`
	// Whether or not the property's value must be unique. Once set, this can't be changed.
	HasUniqueValue *bool `json:"hasUniqueValue,omitempty"`
	// Controls how the property appears in HubSpot.
	FieldType string `json:"fieldType"`
}

type _ObjectTypePropertyCreate ObjectTypePropertyCreate

// NewObjectTypePropertyCreate instantiates a new ObjectTypePropertyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectTypePropertyCreate(label string, type_ string, name string, fieldType string) *ObjectTypePropertyCreate {
	this := ObjectTypePropertyCreate{}
	this.Label = label
	this.Type = type_
	this.Name = name
	this.FieldType = fieldType
	return &this
}

// NewObjectTypePropertyCreateWithDefaults instantiates a new ObjectTypePropertyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectTypePropertyCreateWithDefaults() *ObjectTypePropertyCreate {
	this := ObjectTypePropertyCreate{}
	return &this
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *ObjectTypePropertyCreate) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectTypePropertyCreate) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *ObjectTypePropertyCreate) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *ObjectTypePropertyCreate) SetHidden(v bool) {
	o.Hidden = &v
}

// GetOptionSortStrategy returns the OptionSortStrategy field value if set, zero value otherwise.
func (o *ObjectTypePropertyCreate) GetOptionSortStrategy() string {
	if o == nil || IsNil(o.OptionSortStrategy) {
		var ret string
		return ret
	}
	return *o.OptionSortStrategy
}

// GetOptionSortStrategyOk returns a tuple with the OptionSortStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectTypePropertyCreate) GetOptionSortStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.OptionSortStrategy) {
		return nil, false
	}
	return o.OptionSortStrategy, true
}

// HasOptionSortStrategy returns a boolean if a field has been set.
func (o *ObjectTypePropertyCreate) HasOptionSortStrategy() bool {
	if o != nil && !IsNil(o.OptionSortStrategy) {
		return true
	}

	return false
}

// SetOptionSortStrategy gets a reference to the given string and assigns it to the OptionSortStrategy field.
func (o *ObjectTypePropertyCreate) SetOptionSortStrategy(v string) {
	o.OptionSortStrategy = &v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *ObjectTypePropertyCreate) GetDisplayOrder() int32 {
	if o == nil || IsNil(o.DisplayOrder) {
		var ret int32
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectTypePropertyCreate) GetDisplayOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.DisplayOrder) {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *ObjectTypePropertyCreate) HasDisplayOrder() bool {
	if o != nil && !IsNil(o.DisplayOrder) {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int32 and assigns it to the DisplayOrder field.
func (o *ObjectTypePropertyCreate) SetDisplayOrder(v int32) {
	o.DisplayOrder = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ObjectTypePropertyCreate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectTypePropertyCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ObjectTypePropertyCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ObjectTypePropertyCreate) SetDescription(v string) {
	o.Description = &v
}

// GetShowCurrencySymbol returns the ShowCurrencySymbol field value if set, zero value otherwise.
func (o *ObjectTypePropertyCreate) GetShowCurrencySymbol() bool {
	if o == nil || IsNil(o.ShowCurrencySymbol) {
		var ret bool
		return ret
	}
	return *o.ShowCurrencySymbol
}

// GetShowCurrencySymbolOk returns a tuple with the ShowCurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectTypePropertyCreate) GetShowCurrencySymbolOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowCurrencySymbol) {
		return nil, false
	}
	return o.ShowCurrencySymbol, true
}

// HasShowCurrencySymbol returns a boolean if a field has been set.
func (o *ObjectTypePropertyCreate) HasShowCurrencySymbol() bool {
	if o != nil && !IsNil(o.ShowCurrencySymbol) {
		return true
	}

	return false
}

// SetShowCurrencySymbol gets a reference to the given bool and assigns it to the ShowCurrencySymbol field.
func (o *ObjectTypePropertyCreate) SetShowCurrencySymbol(v bool) {
	o.ShowCurrencySymbol = &v
}

// GetLabel returns the Label field value
func (o *ObjectTypePropertyCreate) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ObjectTypePropertyCreate) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ObjectTypePropertyCreate) SetLabel(v string) {
	o.Label = v
}

// GetType returns the Type field value
func (o *ObjectTypePropertyCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObjectTypePropertyCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ObjectTypePropertyCreate) SetType(v string) {
	o.Type = v
}

// GetFormField returns the FormField field value if set, zero value otherwise.
func (o *ObjectTypePropertyCreate) GetFormField() bool {
	if o == nil || IsNil(o.FormField) {
		var ret bool
		return ret
	}
	return *o.FormField
}

// GetFormFieldOk returns a tuple with the FormField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectTypePropertyCreate) GetFormFieldOk() (*bool, bool) {
	if o == nil || IsNil(o.FormField) {
		return nil, false
	}
	return o.FormField, true
}

// HasFormField returns a boolean if a field has been set.
func (o *ObjectTypePropertyCreate) HasFormField() bool {
	if o != nil && !IsNil(o.FormField) {
		return true
	}

	return false
}

// SetFormField gets a reference to the given bool and assigns it to the FormField field.
func (o *ObjectTypePropertyCreate) SetFormField(v bool) {
	o.FormField = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *ObjectTypePropertyCreate) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectTypePropertyCreate) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *ObjectTypePropertyCreate) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *ObjectTypePropertyCreate) SetGroupName(v string) {
	o.GroupName = &v
}

// GetReferencedObjectType returns the ReferencedObjectType field value if set, zero value otherwise.
func (o *ObjectTypePropertyCreate) GetReferencedObjectType() string {
	if o == nil || IsNil(o.ReferencedObjectType) {
		var ret string
		return ret
	}
	return *o.ReferencedObjectType
}

// GetReferencedObjectTypeOk returns a tuple with the ReferencedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectTypePropertyCreate) GetReferencedObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReferencedObjectType) {
		return nil, false
	}
	return o.ReferencedObjectType, true
}

// HasReferencedObjectType returns a boolean if a field has been set.
func (o *ObjectTypePropertyCreate) HasReferencedObjectType() bool {
	if o != nil && !IsNil(o.ReferencedObjectType) {
		return true
	}

	return false
}

// SetReferencedObjectType gets a reference to the given string and assigns it to the ReferencedObjectType field.
func (o *ObjectTypePropertyCreate) SetReferencedObjectType(v string) {
	o.ReferencedObjectType = &v
}

// GetTextDisplayHint returns the TextDisplayHint field value if set, zero value otherwise.
func (o *ObjectTypePropertyCreate) GetTextDisplayHint() string {
	if o == nil || IsNil(o.TextDisplayHint) {
		var ret string
		return ret
	}
	return *o.TextDisplayHint
}

// GetTextDisplayHintOk returns a tuple with the TextDisplayHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectTypePropertyCreate) GetTextDisplayHintOk() (*string, bool) {
	if o == nil || IsNil(o.TextDisplayHint) {
		return nil, false
	}
	return o.TextDisplayHint, true
}

// HasTextDisplayHint returns a boolean if a field has been set.
func (o *ObjectTypePropertyCreate) HasTextDisplayHint() bool {
	if o != nil && !IsNil(o.TextDisplayHint) {
		return true
	}

	return false
}

// SetTextDisplayHint gets a reference to the given string and assigns it to the TextDisplayHint field.
func (o *ObjectTypePropertyCreate) SetTextDisplayHint(v string) {
	o.TextDisplayHint = &v
}

// GetName returns the Name field value
func (o *ObjectTypePropertyCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObjectTypePropertyCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ObjectTypePropertyCreate) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ObjectTypePropertyCreate) GetOptions() []OptionInput {
	if o == nil || IsNil(o.Options) {
		var ret []OptionInput
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectTypePropertyCreate) GetOptionsOk() ([]OptionInput, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ObjectTypePropertyCreate) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []OptionInput and assigns it to the Options field.
func (o *ObjectTypePropertyCreate) SetOptions(v []OptionInput) {
	o.Options = v
}

// GetSearchableInGlobalSearch returns the SearchableInGlobalSearch field value if set, zero value otherwise.
func (o *ObjectTypePropertyCreate) GetSearchableInGlobalSearch() bool {
	if o == nil || IsNil(o.SearchableInGlobalSearch) {
		var ret bool
		return ret
	}
	return *o.SearchableInGlobalSearch
}

// GetSearchableInGlobalSearchOk returns a tuple with the SearchableInGlobalSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectTypePropertyCreate) GetSearchableInGlobalSearchOk() (*bool, bool) {
	if o == nil || IsNil(o.SearchableInGlobalSearch) {
		return nil, false
	}
	return o.SearchableInGlobalSearch, true
}

// HasSearchableInGlobalSearch returns a boolean if a field has been set.
func (o *ObjectTypePropertyCreate) HasSearchableInGlobalSearch() bool {
	if o != nil && !IsNil(o.SearchableInGlobalSearch) {
		return true
	}

	return false
}

// SetSearchableInGlobalSearch gets a reference to the given bool and assigns it to the SearchableInGlobalSearch field.
func (o *ObjectTypePropertyCreate) SetSearchableInGlobalSearch(v bool) {
	o.SearchableInGlobalSearch = &v
}

// GetNumberDisplayHint returns the NumberDisplayHint field value if set, zero value otherwise.
func (o *ObjectTypePropertyCreate) GetNumberDisplayHint() string {
	if o == nil || IsNil(o.NumberDisplayHint) {
		var ret string
		return ret
	}
	return *o.NumberDisplayHint
}

// GetNumberDisplayHintOk returns a tuple with the NumberDisplayHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectTypePropertyCreate) GetNumberDisplayHintOk() (*string, bool) {
	if o == nil || IsNil(o.NumberDisplayHint) {
		return nil, false
	}
	return o.NumberDisplayHint, true
}

// HasNumberDisplayHint returns a boolean if a field has been set.
func (o *ObjectTypePropertyCreate) HasNumberDisplayHint() bool {
	if o != nil && !IsNil(o.NumberDisplayHint) {
		return true
	}

	return false
}

// SetNumberDisplayHint gets a reference to the given string and assigns it to the NumberDisplayHint field.
func (o *ObjectTypePropertyCreate) SetNumberDisplayHint(v string) {
	o.NumberDisplayHint = &v
}

// GetHasUniqueValue returns the HasUniqueValue field value if set, zero value otherwise.
func (o *ObjectTypePropertyCreate) GetHasUniqueValue() bool {
	if o == nil || IsNil(o.HasUniqueValue) {
		var ret bool
		return ret
	}
	return *o.HasUniqueValue
}

// GetHasUniqueValueOk returns a tuple with the HasUniqueValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectTypePropertyCreate) GetHasUniqueValueOk() (*bool, bool) {
	if o == nil || IsNil(o.HasUniqueValue) {
		return nil, false
	}
	return o.HasUniqueValue, true
}

// HasHasUniqueValue returns a boolean if a field has been set.
func (o *ObjectTypePropertyCreate) HasHasUniqueValue() bool {
	if o != nil && !IsNil(o.HasUniqueValue) {
		return true
	}

	return false
}

// SetHasUniqueValue gets a reference to the given bool and assigns it to the HasUniqueValue field.
func (o *ObjectTypePropertyCreate) SetHasUniqueValue(v bool) {
	o.HasUniqueValue = &v
}

// GetFieldType returns the FieldType field value
func (o *ObjectTypePropertyCreate) GetFieldType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value
// and a boolean to check if the value has been set.
func (o *ObjectTypePropertyCreate) GetFieldTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldType, true
}

// SetFieldType sets field value
func (o *ObjectTypePropertyCreate) SetFieldType(v string) {
	o.FieldType = v
}

func (o ObjectTypePropertyCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectTypePropertyCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	if !IsNil(o.OptionSortStrategy) {
		toSerialize["optionSortStrategy"] = o.OptionSortStrategy
	}
	if !IsNil(o.DisplayOrder) {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ShowCurrencySymbol) {
		toSerialize["showCurrencySymbol"] = o.ShowCurrencySymbol
	}
	toSerialize["label"] = o.Label
	toSerialize["type"] = o.Type
	if !IsNil(o.FormField) {
		toSerialize["formField"] = o.FormField
	}
	if !IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	if !IsNil(o.ReferencedObjectType) {
		toSerialize["referencedObjectType"] = o.ReferencedObjectType
	}
	if !IsNil(o.TextDisplayHint) {
		toSerialize["textDisplayHint"] = o.TextDisplayHint
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.SearchableInGlobalSearch) {
		toSerialize["searchableInGlobalSearch"] = o.SearchableInGlobalSearch
	}
	if !IsNil(o.NumberDisplayHint) {
		toSerialize["numberDisplayHint"] = o.NumberDisplayHint
	}
	if !IsNil(o.HasUniqueValue) {
		toSerialize["hasUniqueValue"] = o.HasUniqueValue
	}
	toSerialize["fieldType"] = o.FieldType
	return toSerialize, nil
}

func (o *ObjectTypePropertyCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"type",
		"name",
		"fieldType",
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

	varObjectTypePropertyCreate := _ObjectTypePropertyCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varObjectTypePropertyCreate)

	if err != nil {
		return err
	}

	*o = ObjectTypePropertyCreate(varObjectTypePropertyCreate)

	return err
}

type NullableObjectTypePropertyCreate struct {
	value *ObjectTypePropertyCreate
	isSet bool
}

func (v NullableObjectTypePropertyCreate) Get() *ObjectTypePropertyCreate {
	return v.value
}

func (v *NullableObjectTypePropertyCreate) Set(val *ObjectTypePropertyCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectTypePropertyCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectTypePropertyCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectTypePropertyCreate(val *ObjectTypePropertyCreate) *NullableObjectTypePropertyCreate {
	return &NullableObjectTypePropertyCreate{value: val, isSet: true}
}

func (v NullableObjectTypePropertyCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectTypePropertyCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


