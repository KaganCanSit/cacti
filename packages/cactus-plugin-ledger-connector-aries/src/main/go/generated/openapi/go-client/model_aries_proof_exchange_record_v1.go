/*
Hyperledger Cacti Plugin - Connector Aries

Can communicate with other Aries agents and Cacti Aries connectors

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-aries

import (
	"encoding/json"
)

// checks if the AriesProofExchangeRecordV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AriesProofExchangeRecordV1{}

// AriesProofExchangeRecordV1 Proof exchange record from Aries framework (simplified)
type AriesProofExchangeRecordV1 struct {
	Id string `json:"id"`
	ConnectionId *string `json:"connectionId,omitempty"`
	ThreadId string `json:"threadId"`
	State string `json:"state"`
	ProtocolVersion string `json:"protocolVersion"`
	IsVerified *bool `json:"isVerified,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AriesProofExchangeRecordV1 AriesProofExchangeRecordV1

// NewAriesProofExchangeRecordV1 instantiates a new AriesProofExchangeRecordV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAriesProofExchangeRecordV1(id string, threadId string, state string, protocolVersion string) *AriesProofExchangeRecordV1 {
	this := AriesProofExchangeRecordV1{}
	this.Id = id
	this.ThreadId = threadId
	this.State = state
	this.ProtocolVersion = protocolVersion
	return &this
}

// NewAriesProofExchangeRecordV1WithDefaults instantiates a new AriesProofExchangeRecordV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAriesProofExchangeRecordV1WithDefaults() *AriesProofExchangeRecordV1 {
	this := AriesProofExchangeRecordV1{}
	return &this
}

// GetId returns the Id field value
func (o *AriesProofExchangeRecordV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AriesProofExchangeRecordV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AriesProofExchangeRecordV1) SetId(v string) {
	o.Id = v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *AriesProofExchangeRecordV1) GetConnectionId() string {
	if o == nil || IsNil(o.ConnectionId) {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AriesProofExchangeRecordV1) GetConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionId) {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *AriesProofExchangeRecordV1) HasConnectionId() bool {
	if o != nil && !IsNil(o.ConnectionId) {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *AriesProofExchangeRecordV1) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetThreadId returns the ThreadId field value
func (o *AriesProofExchangeRecordV1) GetThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *AriesProofExchangeRecordV1) GetThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *AriesProofExchangeRecordV1) SetThreadId(v string) {
	o.ThreadId = v
}

// GetState returns the State field value
func (o *AriesProofExchangeRecordV1) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *AriesProofExchangeRecordV1) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *AriesProofExchangeRecordV1) SetState(v string) {
	o.State = v
}

// GetProtocolVersion returns the ProtocolVersion field value
func (o *AriesProofExchangeRecordV1) GetProtocolVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProtocolVersion
}

// GetProtocolVersionOk returns a tuple with the ProtocolVersion field value
// and a boolean to check if the value has been set.
func (o *AriesProofExchangeRecordV1) GetProtocolVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProtocolVersion, true
}

// SetProtocolVersion sets field value
func (o *AriesProofExchangeRecordV1) SetProtocolVersion(v string) {
	o.ProtocolVersion = v
}

// GetIsVerified returns the IsVerified field value if set, zero value otherwise.
func (o *AriesProofExchangeRecordV1) GetIsVerified() bool {
	if o == nil || IsNil(o.IsVerified) {
		var ret bool
		return ret
	}
	return *o.IsVerified
}

// GetIsVerifiedOk returns a tuple with the IsVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AriesProofExchangeRecordV1) GetIsVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsVerified) {
		return nil, false
	}
	return o.IsVerified, true
}

// HasIsVerified returns a boolean if a field has been set.
func (o *AriesProofExchangeRecordV1) HasIsVerified() bool {
	if o != nil && !IsNil(o.IsVerified) {
		return true
	}

	return false
}

// SetIsVerified gets a reference to the given bool and assigns it to the IsVerified field.
func (o *AriesProofExchangeRecordV1) SetIsVerified(v bool) {
	o.IsVerified = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *AriesProofExchangeRecordV1) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AriesProofExchangeRecordV1) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *AriesProofExchangeRecordV1) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *AriesProofExchangeRecordV1) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o AriesProofExchangeRecordV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AriesProofExchangeRecordV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.ConnectionId) {
		toSerialize["connectionId"] = o.ConnectionId
	}
	toSerialize["threadId"] = o.ThreadId
	toSerialize["state"] = o.State
	toSerialize["protocolVersion"] = o.ProtocolVersion
	if !IsNil(o.IsVerified) {
		toSerialize["isVerified"] = o.IsVerified
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AriesProofExchangeRecordV1) UnmarshalJSON(bytes []byte) (err error) {
	varAriesProofExchangeRecordV1 := _AriesProofExchangeRecordV1{}

	if err = json.Unmarshal(bytes, &varAriesProofExchangeRecordV1); err == nil {
		*o = AriesProofExchangeRecordV1(varAriesProofExchangeRecordV1)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "connectionId")
		delete(additionalProperties, "threadId")
		delete(additionalProperties, "state")
		delete(additionalProperties, "protocolVersion")
		delete(additionalProperties, "isVerified")
		delete(additionalProperties, "errorMessage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAriesProofExchangeRecordV1 struct {
	value *AriesProofExchangeRecordV1
	isSet bool
}

func (v NullableAriesProofExchangeRecordV1) Get() *AriesProofExchangeRecordV1 {
	return v.value
}

func (v *NullableAriesProofExchangeRecordV1) Set(val *AriesProofExchangeRecordV1) {
	v.value = val
	v.isSet = true
}

func (v NullableAriesProofExchangeRecordV1) IsSet() bool {
	return v.isSet
}

func (v *NullableAriesProofExchangeRecordV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAriesProofExchangeRecordV1(val *AriesProofExchangeRecordV1) *NullableAriesProofExchangeRecordV1 {
	return &NullableAriesProofExchangeRecordV1{value: val, isSet: true}
}

func (v NullableAriesProofExchangeRecordV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAriesProofExchangeRecordV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


