/*
Hyperledger Cactus Plugin - Odap Hermes

Implementation for Odap and Hermes

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-satp-hermes

import (
	"encoding/json"
)

// checks if the RollbackV1Message type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RollbackV1Message{}

// RollbackV1Message struct for RollbackV1Message
type RollbackV1Message struct {
	SessionID string `json:"sessionID"`
	Success bool `json:"success"`
	ActionPerformed []string `json:"actionPerformed"`
	Proofs []string `json:"proofs"`
	Signature string `json:"signature"`
}

// NewRollbackV1Message instantiates a new RollbackV1Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRollbackV1Message(sessionID string, success bool, actionPerformed []string, proofs []string, signature string) *RollbackV1Message {
	this := RollbackV1Message{}
	this.SessionID = sessionID
	this.Success = success
	this.ActionPerformed = actionPerformed
	this.Proofs = proofs
	this.Signature = signature
	return &this
}

// NewRollbackV1MessageWithDefaults instantiates a new RollbackV1Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollbackV1MessageWithDefaults() *RollbackV1Message {
	this := RollbackV1Message{}
	return &this
}

// GetSessionID returns the SessionID field value
func (o *RollbackV1Message) GetSessionID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionID
}

// GetSessionIDOk returns a tuple with the SessionID field value
// and a boolean to check if the value has been set.
func (o *RollbackV1Message) GetSessionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionID, true
}

// SetSessionID sets field value
func (o *RollbackV1Message) SetSessionID(v string) {
	o.SessionID = v
}

// GetSuccess returns the Success field value
func (o *RollbackV1Message) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *RollbackV1Message) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *RollbackV1Message) SetSuccess(v bool) {
	o.Success = v
}

// GetActionPerformed returns the ActionPerformed field value
func (o *RollbackV1Message) GetActionPerformed() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ActionPerformed
}

// GetActionPerformedOk returns a tuple with the ActionPerformed field value
// and a boolean to check if the value has been set.
func (o *RollbackV1Message) GetActionPerformedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionPerformed, true
}

// SetActionPerformed sets field value
func (o *RollbackV1Message) SetActionPerformed(v []string) {
	o.ActionPerformed = v
}

// GetProofs returns the Proofs field value
func (o *RollbackV1Message) GetProofs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Proofs
}

// GetProofsOk returns a tuple with the Proofs field value
// and a boolean to check if the value has been set.
func (o *RollbackV1Message) GetProofsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proofs, true
}

// SetProofs sets field value
func (o *RollbackV1Message) SetProofs(v []string) {
	o.Proofs = v
}

// GetSignature returns the Signature field value
func (o *RollbackV1Message) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *RollbackV1Message) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *RollbackV1Message) SetSignature(v string) {
	o.Signature = v
}

func (o RollbackV1Message) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RollbackV1Message) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionID"] = o.SessionID
	toSerialize["success"] = o.Success
	toSerialize["actionPerformed"] = o.ActionPerformed
	toSerialize["proofs"] = o.Proofs
	toSerialize["signature"] = o.Signature
	return toSerialize, nil
}

type NullableRollbackV1Message struct {
	value *RollbackV1Message
	isSet bool
}

func (v NullableRollbackV1Message) Get() *RollbackV1Message {
	return v.value
}

func (v *NullableRollbackV1Message) Set(val *RollbackV1Message) {
	v.value = val
	v.isSet = true
}

func (v NullableRollbackV1Message) IsSet() bool {
	return v.isSet
}

func (v *NullableRollbackV1Message) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRollbackV1Message(val *RollbackV1Message) *NullableRollbackV1Message {
	return &NullableRollbackV1Message{value: val, isSet: true}
}

func (v NullableRollbackV1Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRollbackV1Message) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


