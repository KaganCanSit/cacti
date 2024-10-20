/*
Hyperledger Cacti Plugin - Connector Corda

Can perform basic tasks on a Corda ledger

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-corda

import (
	"encoding/json"
	"time"
)

// checks if the FlowStatusV1ResponsesFlowStatusResponsesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowStatusV1ResponsesFlowStatusResponsesInner{}

// FlowStatusV1ResponsesFlowStatusResponsesInner struct for FlowStatusV1ResponsesFlowStatusResponsesInner
type FlowStatusV1ResponsesFlowStatusResponsesInner struct {
	ClientRequestId NullableString `json:"clientRequestId,omitempty"`
	FlowError *FlowV1Error `json:"flowError,omitempty"`
	FlowId NullableString `json:"flowId,omitempty"`
	FlowResult NullableString `json:"flowResult,omitempty"`
	FlowStatus *string `json:"flowStatus,omitempty"`
	HoldingIDShortHash *string `json:"holdingIDShortHash,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewFlowStatusV1ResponsesFlowStatusResponsesInner instantiates a new FlowStatusV1ResponsesFlowStatusResponsesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowStatusV1ResponsesFlowStatusResponsesInner() *FlowStatusV1ResponsesFlowStatusResponsesInner {
	this := FlowStatusV1ResponsesFlowStatusResponsesInner{}
	return &this
}

// NewFlowStatusV1ResponsesFlowStatusResponsesInnerWithDefaults instantiates a new FlowStatusV1ResponsesFlowStatusResponsesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowStatusV1ResponsesFlowStatusResponsesInnerWithDefaults() *FlowStatusV1ResponsesFlowStatusResponsesInner {
	this := FlowStatusV1ResponsesFlowStatusResponsesInner{}
	return &this
}

// GetClientRequestId returns the ClientRequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) GetClientRequestId() string {
	if o == nil || IsNil(o.ClientRequestId.Get()) {
		var ret string
		return ret
	}
	return *o.ClientRequestId.Get()
}

// GetClientRequestIdOk returns a tuple with the ClientRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) GetClientRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientRequestId.Get(), o.ClientRequestId.IsSet()
}

// HasClientRequestId returns a boolean if a field has been set.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) HasClientRequestId() bool {
	if o != nil && o.ClientRequestId.IsSet() {
		return true
	}

	return false
}

// SetClientRequestId gets a reference to the given NullableString and assigns it to the ClientRequestId field.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) SetClientRequestId(v string) {
	o.ClientRequestId.Set(&v)
}
// SetClientRequestIdNil sets the value for ClientRequestId to be an explicit nil
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) SetClientRequestIdNil() {
	o.ClientRequestId.Set(nil)
}

// UnsetClientRequestId ensures that no value is present for ClientRequestId, not even an explicit nil
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) UnsetClientRequestId() {
	o.ClientRequestId.Unset()
}

// GetFlowError returns the FlowError field value if set, zero value otherwise.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) GetFlowError() FlowV1Error {
	if o == nil || IsNil(o.FlowError) {
		var ret FlowV1Error
		return ret
	}
	return *o.FlowError
}

// GetFlowErrorOk returns a tuple with the FlowError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) GetFlowErrorOk() (*FlowV1Error, bool) {
	if o == nil || IsNil(o.FlowError) {
		return nil, false
	}
	return o.FlowError, true
}

// HasFlowError returns a boolean if a field has been set.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) HasFlowError() bool {
	if o != nil && !IsNil(o.FlowError) {
		return true
	}

	return false
}

// SetFlowError gets a reference to the given FlowV1Error and assigns it to the FlowError field.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) SetFlowError(v FlowV1Error) {
	o.FlowError = &v
}

// GetFlowId returns the FlowId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) GetFlowId() string {
	if o == nil || IsNil(o.FlowId.Get()) {
		var ret string
		return ret
	}
	return *o.FlowId.Get()
}

// GetFlowIdOk returns a tuple with the FlowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) GetFlowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowId.Get(), o.FlowId.IsSet()
}

// HasFlowId returns a boolean if a field has been set.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) HasFlowId() bool {
	if o != nil && o.FlowId.IsSet() {
		return true
	}

	return false
}

// SetFlowId gets a reference to the given NullableString and assigns it to the FlowId field.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) SetFlowId(v string) {
	o.FlowId.Set(&v)
}
// SetFlowIdNil sets the value for FlowId to be an explicit nil
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) SetFlowIdNil() {
	o.FlowId.Set(nil)
}

// UnsetFlowId ensures that no value is present for FlowId, not even an explicit nil
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) UnsetFlowId() {
	o.FlowId.Unset()
}

// GetFlowResult returns the FlowResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) GetFlowResult() string {
	if o == nil || IsNil(o.FlowResult.Get()) {
		var ret string
		return ret
	}
	return *o.FlowResult.Get()
}

// GetFlowResultOk returns a tuple with the FlowResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) GetFlowResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowResult.Get(), o.FlowResult.IsSet()
}

// HasFlowResult returns a boolean if a field has been set.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) HasFlowResult() bool {
	if o != nil && o.FlowResult.IsSet() {
		return true
	}

	return false
}

// SetFlowResult gets a reference to the given NullableString and assigns it to the FlowResult field.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) SetFlowResult(v string) {
	o.FlowResult.Set(&v)
}
// SetFlowResultNil sets the value for FlowResult to be an explicit nil
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) SetFlowResultNil() {
	o.FlowResult.Set(nil)
}

// UnsetFlowResult ensures that no value is present for FlowResult, not even an explicit nil
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) UnsetFlowResult() {
	o.FlowResult.Unset()
}

// GetFlowStatus returns the FlowStatus field value if set, zero value otherwise.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) GetFlowStatus() string {
	if o == nil || IsNil(o.FlowStatus) {
		var ret string
		return ret
	}
	return *o.FlowStatus
}

// GetFlowStatusOk returns a tuple with the FlowStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) GetFlowStatusOk() (*string, bool) {
	if o == nil || IsNil(o.FlowStatus) {
		return nil, false
	}
	return o.FlowStatus, true
}

// HasFlowStatus returns a boolean if a field has been set.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) HasFlowStatus() bool {
	if o != nil && !IsNil(o.FlowStatus) {
		return true
	}

	return false
}

// SetFlowStatus gets a reference to the given string and assigns it to the FlowStatus field.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) SetFlowStatus(v string) {
	o.FlowStatus = &v
}

// GetHoldingIDShortHash returns the HoldingIDShortHash field value if set, zero value otherwise.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) GetHoldingIDShortHash() string {
	if o == nil || IsNil(o.HoldingIDShortHash) {
		var ret string
		return ret
	}
	return *o.HoldingIDShortHash
}

// GetHoldingIDShortHashOk returns a tuple with the HoldingIDShortHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) GetHoldingIDShortHashOk() (*string, bool) {
	if o == nil || IsNil(o.HoldingIDShortHash) {
		return nil, false
	}
	return o.HoldingIDShortHash, true
}

// HasHoldingIDShortHash returns a boolean if a field has been set.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) HasHoldingIDShortHash() bool {
	if o != nil && !IsNil(o.HoldingIDShortHash) {
		return true
	}

	return false
}

// SetHoldingIDShortHash gets a reference to the given string and assigns it to the HoldingIDShortHash field.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) SetHoldingIDShortHash(v string) {
	o.HoldingIDShortHash = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *FlowStatusV1ResponsesFlowStatusResponsesInner) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o FlowStatusV1ResponsesFlowStatusResponsesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowStatusV1ResponsesFlowStatusResponsesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientRequestId.IsSet() {
		toSerialize["clientRequestId"] = o.ClientRequestId.Get()
	}
	if !IsNil(o.FlowError) {
		toSerialize["flowError"] = o.FlowError
	}
	if o.FlowId.IsSet() {
		toSerialize["flowId"] = o.FlowId.Get()
	}
	if o.FlowResult.IsSet() {
		toSerialize["flowResult"] = o.FlowResult.Get()
	}
	if !IsNil(o.FlowStatus) {
		toSerialize["flowStatus"] = o.FlowStatus
	}
	if !IsNil(o.HoldingIDShortHash) {
		toSerialize["holdingIDShortHash"] = o.HoldingIDShortHash
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableFlowStatusV1ResponsesFlowStatusResponsesInner struct {
	value *FlowStatusV1ResponsesFlowStatusResponsesInner
	isSet bool
}

func (v NullableFlowStatusV1ResponsesFlowStatusResponsesInner) Get() *FlowStatusV1ResponsesFlowStatusResponsesInner {
	return v.value
}

func (v *NullableFlowStatusV1ResponsesFlowStatusResponsesInner) Set(val *FlowStatusV1ResponsesFlowStatusResponsesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowStatusV1ResponsesFlowStatusResponsesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowStatusV1ResponsesFlowStatusResponsesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowStatusV1ResponsesFlowStatusResponsesInner(val *FlowStatusV1ResponsesFlowStatusResponsesInner) *NullableFlowStatusV1ResponsesFlowStatusResponsesInner {
	return &NullableFlowStatusV1ResponsesFlowStatusResponsesInner{value: val, isSet: true}
}

func (v NullableFlowStatusV1ResponsesFlowStatusResponsesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowStatusV1ResponsesFlowStatusResponsesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


