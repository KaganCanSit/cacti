/*
Hyperledger Cacti Plugin - Connector Corda

Can perform basic tasks on a Corda ledger

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-corda

import (
	"encoding/json"
)

// checks if the PublicKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicKey{}

// PublicKey An instance of a java.security.PublicKey (which is an interface) implementation such as org.hyperledger.cactus.plugin.ledger.connector.corda.server.impl.PublicKeyImpl
type PublicKey struct {
	Algorithm string `json:"algorithm"`
	Format string `json:"format"`
	Encoded string `json:"encoded"`
}

// NewPublicKey instantiates a new PublicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicKey(algorithm string, format string, encoded string) *PublicKey {
	this := PublicKey{}
	this.Algorithm = algorithm
	this.Format = format
	this.Encoded = encoded
	return &this
}

// NewPublicKeyWithDefaults instantiates a new PublicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicKeyWithDefaults() *PublicKey {
	this := PublicKey{}
	return &this
}

// GetAlgorithm returns the Algorithm field value
func (o *PublicKey) GetAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *PublicKey) GetAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value
func (o *PublicKey) SetAlgorithm(v string) {
	o.Algorithm = v
}

// GetFormat returns the Format field value
func (o *PublicKey) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *PublicKey) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *PublicKey) SetFormat(v string) {
	o.Format = v
}

// GetEncoded returns the Encoded field value
func (o *PublicKey) GetEncoded() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Encoded
}

// GetEncodedOk returns a tuple with the Encoded field value
// and a boolean to check if the value has been set.
func (o *PublicKey) GetEncodedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encoded, true
}

// SetEncoded sets field value
func (o *PublicKey) SetEncoded(v string) {
	o.Encoded = v
}

func (o PublicKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["algorithm"] = o.Algorithm
	toSerialize["format"] = o.Format
	toSerialize["encoded"] = o.Encoded
	return toSerialize, nil
}

type NullablePublicKey struct {
	value *PublicKey
	isSet bool
}

func (v NullablePublicKey) Get() *PublicKey {
	return v.value
}

func (v *NullablePublicKey) Set(val *PublicKey) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicKey) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicKey(val *PublicKey) *NullablePublicKey {
	return &NullablePublicKey{value: val, isSet: true}
}

func (v NullablePublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


