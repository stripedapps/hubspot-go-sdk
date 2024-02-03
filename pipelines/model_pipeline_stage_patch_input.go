/*
Pipelines

Pipelines represent distinct stages in a workflow, like closing a deal or servicing a support ticket. These endpoints provide access to read and modify pipelines in HubSpot. Pipelines support `deals` and `tickets` object types.  ## Pipeline ID validation  When calling endpoints that take pipelineId as a parameter, that ID must correspond to an existing, un-archived pipeline. Otherwise the request will fail with a `404 Not Found` response.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PipelineStagePatchInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PipelineStagePatchInput{}

// PipelineStagePatchInput An input used to update some properties on a pipeline definition.
type PipelineStagePatchInput struct {
	// Whether the pipeline is archived.
	Archived *bool `json:"archived,omitempty"`
	// A JSON object containing properties that are not present on all object pipelines.  For `deals` pipelines, the `probability` field is required (`{ \"probability\": 0.5 }`), and represents the likelihood a deal will close. Possible values are between 0.0 and 1.0 in increments of 0.1.  For `tickets` pipelines, the `ticketState` field is optional (`{ \"ticketState\": \"OPEN\" }`), and represents whether the ticket remains open or has been closed by a member of your Support team. Possible values are `OPEN` or `CLOSED`.
	Metadata map[string]string `json:"metadata"`
	// The order for displaying this pipeline stage. If two pipeline stages have a matching `displayOrder`, they will be sorted alphabetically by label.
	DisplayOrder *int32 `json:"displayOrder,omitempty"`
	// A label used to organize pipeline stages in HubSpot's UI. Each pipeline stage's label must be unique within that pipeline.
	Label *string `json:"label,omitempty"`
}

// NewPipelineStagePatchInput instantiates a new PipelineStagePatchInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStagePatchInput(metadata map[string]string) *PipelineStagePatchInput {
	this := PipelineStagePatchInput{}
	this.Metadata = metadata
	return &this
}

// NewPipelineStagePatchInputWithDefaults instantiates a new PipelineStagePatchInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStagePatchInputWithDefaults() *PipelineStagePatchInput {
	this := PipelineStagePatchInput{}
	return &this
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *PipelineStagePatchInput) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStagePatchInput) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *PipelineStagePatchInput) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *PipelineStagePatchInput) SetArchived(v bool) {
	o.Archived = &v
}

// GetMetadata returns the Metadata field value
func (o *PipelineStagePatchInput) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *PipelineStagePatchInput) GetMetadataOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *PipelineStagePatchInput) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *PipelineStagePatchInput) GetDisplayOrder() int32 {
	if o == nil || IsNil(o.DisplayOrder) {
		var ret int32
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStagePatchInput) GetDisplayOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.DisplayOrder) {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *PipelineStagePatchInput) HasDisplayOrder() bool {
	if o != nil && !IsNil(o.DisplayOrder) {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int32 and assigns it to the DisplayOrder field.
func (o *PipelineStagePatchInput) SetDisplayOrder(v int32) {
	o.DisplayOrder = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PipelineStagePatchInput) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStagePatchInput) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PipelineStagePatchInput) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PipelineStagePatchInput) SetLabel(v string) {
	o.Label = &v
}

func (o PipelineStagePatchInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PipelineStagePatchInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	toSerialize["metadata"] = o.Metadata
	if !IsNil(o.DisplayOrder) {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullablePipelineStagePatchInput struct {
	value *PipelineStagePatchInput
	isSet bool
}

func (v NullablePipelineStagePatchInput) Get() *PipelineStagePatchInput {
	return v.value
}

func (v *NullablePipelineStagePatchInput) Set(val *PipelineStagePatchInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStagePatchInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStagePatchInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStagePatchInput(val *PipelineStagePatchInput) *NullablePipelineStagePatchInput {
	return &NullablePipelineStagePatchInput{value: val, isSet: true}
}

func (v NullablePipelineStagePatchInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStagePatchInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


