/*
Hyperledger Cactus Plugin - HTLC-ETH Besu

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-htlc-eth-besu

import (
	"encoding/json"
)

// checks if the RunTransactionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunTransactionResponse{}

// RunTransactionResponse struct for RunTransactionResponse
type RunTransactionResponse struct {
	TransactionReceipt Web3TransactionReceipt `json:"transactionReceipt"`
}

// NewRunTransactionResponse instantiates a new RunTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunTransactionResponse(transactionReceipt Web3TransactionReceipt) *RunTransactionResponse {
	this := RunTransactionResponse{}
	this.TransactionReceipt = transactionReceipt
	return &this
}

// NewRunTransactionResponseWithDefaults instantiates a new RunTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunTransactionResponseWithDefaults() *RunTransactionResponse {
	this := RunTransactionResponse{}
	return &this
}

// GetTransactionReceipt returns the TransactionReceipt field value
func (o *RunTransactionResponse) GetTransactionReceipt() Web3TransactionReceipt {
	if o == nil {
		var ret Web3TransactionReceipt
		return ret
	}

	return o.TransactionReceipt
}

// GetTransactionReceiptOk returns a tuple with the TransactionReceipt field value
// and a boolean to check if the value has been set.
func (o *RunTransactionResponse) GetTransactionReceiptOk() (*Web3TransactionReceipt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionReceipt, true
}

// SetTransactionReceipt sets field value
func (o *RunTransactionResponse) SetTransactionReceipt(v Web3TransactionReceipt) {
	o.TransactionReceipt = v
}

func (o RunTransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunTransactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transactionReceipt"] = o.TransactionReceipt
	return toSerialize, nil
}

type NullableRunTransactionResponse struct {
	value *RunTransactionResponse
	isSet bool
}

func (v NullableRunTransactionResponse) Get() *RunTransactionResponse {
	return v.value
}

func (v *NullableRunTransactionResponse) Set(val *RunTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRunTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRunTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunTransactionResponse(val *RunTransactionResponse) *NullableRunTransactionResponse {
	return &NullableRunTransactionResponse{value: val, isSet: true}
}

func (v NullableRunTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


