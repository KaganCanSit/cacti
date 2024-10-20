/*
Hyperledger Cactus Plugin - Connector Besu

Can perform basic tasks on a Besu ledger

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-besu

import (
	"encoding/json"
)

// checks if the GetBalanceV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBalanceV1Response{}

// GetBalanceV1Response struct for GetBalanceV1Response
type GetBalanceV1Response struct {
	Balance string `json:"balance"`
}

// NewGetBalanceV1Response instantiates a new GetBalanceV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalanceV1Response(balance string) *GetBalanceV1Response {
	this := GetBalanceV1Response{}
	this.Balance = balance
	return &this
}

// NewGetBalanceV1ResponseWithDefaults instantiates a new GetBalanceV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalanceV1ResponseWithDefaults() *GetBalanceV1Response {
	this := GetBalanceV1Response{}
	return &this
}

// GetBalance returns the Balance field value
func (o *GetBalanceV1Response) GetBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *GetBalanceV1Response) GetBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *GetBalanceV1Response) SetBalance(v string) {
	o.Balance = v
}

func (o GetBalanceV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBalanceV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["balance"] = o.Balance
	return toSerialize, nil
}

type NullableGetBalanceV1Response struct {
	value *GetBalanceV1Response
	isSet bool
}

func (v NullableGetBalanceV1Response) Get() *GetBalanceV1Response {
	return v.value
}

func (v *NullableGetBalanceV1Response) Set(val *GetBalanceV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalanceV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalanceV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalanceV1Response(val *GetBalanceV1Response) *NullableGetBalanceV1Response {
	return &NullableGetBalanceV1Response{value: val, isSet: true}
}

func (v NullableGetBalanceV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalanceV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


