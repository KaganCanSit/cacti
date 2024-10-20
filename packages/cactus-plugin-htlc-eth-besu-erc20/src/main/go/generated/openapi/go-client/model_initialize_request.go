/*
Hyperledger Cactus Plugin - HTLC ETH BESU ERC20

Allows Cactus nodes to interact with HTLC contracts with ERC-20 Tokens

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-htlc-eth-besu-erc20

import (
	"encoding/json"
)

// checks if the InitializeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InitializeRequest{}

// InitializeRequest struct for InitializeRequest
type InitializeRequest struct {
	// connectorId for the connector besu plugin
	ConnectorId string `json:"connectorId"`
	// keychainId for the keychain plugin
	KeychainId string `json:"keychainId"`
	ConstructorArgs []interface{} `json:"constructorArgs"`
	Web3SigningCredential Web3SigningCredential `json:"web3SigningCredential"`
	Gas *float32 `json:"gas,omitempty"`
}

// NewInitializeRequest instantiates a new InitializeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitializeRequest(connectorId string, keychainId string, constructorArgs []interface{}, web3SigningCredential Web3SigningCredential) *InitializeRequest {
	this := InitializeRequest{}
	this.ConnectorId = connectorId
	this.KeychainId = keychainId
	this.ConstructorArgs = constructorArgs
	this.Web3SigningCredential = web3SigningCredential
	return &this
}

// NewInitializeRequestWithDefaults instantiates a new InitializeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitializeRequestWithDefaults() *InitializeRequest {
	this := InitializeRequest{}
	return &this
}

// GetConnectorId returns the ConnectorId field value
func (o *InitializeRequest) GetConnectorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value
// and a boolean to check if the value has been set.
func (o *InitializeRequest) GetConnectorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorId, true
}

// SetConnectorId sets field value
func (o *InitializeRequest) SetConnectorId(v string) {
	o.ConnectorId = v
}

// GetKeychainId returns the KeychainId field value
func (o *InitializeRequest) GetKeychainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeychainId
}

// GetKeychainIdOk returns a tuple with the KeychainId field value
// and a boolean to check if the value has been set.
func (o *InitializeRequest) GetKeychainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeychainId, true
}

// SetKeychainId sets field value
func (o *InitializeRequest) SetKeychainId(v string) {
	o.KeychainId = v
}

// GetConstructorArgs returns the ConstructorArgs field value
func (o *InitializeRequest) GetConstructorArgs() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.ConstructorArgs
}

// GetConstructorArgsOk returns a tuple with the ConstructorArgs field value
// and a boolean to check if the value has been set.
func (o *InitializeRequest) GetConstructorArgsOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConstructorArgs, true
}

// SetConstructorArgs sets field value
func (o *InitializeRequest) SetConstructorArgs(v []interface{}) {
	o.ConstructorArgs = v
}

// GetWeb3SigningCredential returns the Web3SigningCredential field value
func (o *InitializeRequest) GetWeb3SigningCredential() Web3SigningCredential {
	if o == nil {
		var ret Web3SigningCredential
		return ret
	}

	return o.Web3SigningCredential
}

// GetWeb3SigningCredentialOk returns a tuple with the Web3SigningCredential field value
// and a boolean to check if the value has been set.
func (o *InitializeRequest) GetWeb3SigningCredentialOk() (*Web3SigningCredential, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Web3SigningCredential, true
}

// SetWeb3SigningCredential sets field value
func (o *InitializeRequest) SetWeb3SigningCredential(v Web3SigningCredential) {
	o.Web3SigningCredential = v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *InitializeRequest) GetGas() float32 {
	if o == nil || IsNil(o.Gas) {
		var ret float32
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitializeRequest) GetGasOk() (*float32, bool) {
	if o == nil || IsNil(o.Gas) {
		return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *InitializeRequest) HasGas() bool {
	if o != nil && !IsNil(o.Gas) {
		return true
	}

	return false
}

// SetGas gets a reference to the given float32 and assigns it to the Gas field.
func (o *InitializeRequest) SetGas(v float32) {
	o.Gas = &v
}

func (o InitializeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InitializeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connectorId"] = o.ConnectorId
	toSerialize["keychainId"] = o.KeychainId
	toSerialize["constructorArgs"] = o.ConstructorArgs
	toSerialize["web3SigningCredential"] = o.Web3SigningCredential
	if !IsNil(o.Gas) {
		toSerialize["gas"] = o.Gas
	}
	return toSerialize, nil
}

type NullableInitializeRequest struct {
	value *InitializeRequest
	isSet bool
}

func (v NullableInitializeRequest) Get() *InitializeRequest {
	return v.value
}

func (v *NullableInitializeRequest) Set(val *InitializeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInitializeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInitializeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitializeRequest(val *InitializeRequest) *NullableInitializeRequest {
	return &NullableInitializeRequest{value: val, isSet: true}
}

func (v NullableInitializeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitializeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


