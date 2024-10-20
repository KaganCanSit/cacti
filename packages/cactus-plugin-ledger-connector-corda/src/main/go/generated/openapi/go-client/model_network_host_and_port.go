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

// checks if the NetworkHostAndPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkHostAndPort{}

// NetworkHostAndPort struct for NetworkHostAndPort
type NetworkHostAndPort struct {
	Host string `json:"host"`
	Port float32 `json:"port"`
}

// NewNetworkHostAndPort instantiates a new NetworkHostAndPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkHostAndPort(host string, port float32) *NetworkHostAndPort {
	this := NetworkHostAndPort{}
	this.Host = host
	this.Port = port
	return &this
}

// NewNetworkHostAndPortWithDefaults instantiates a new NetworkHostAndPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkHostAndPortWithDefaults() *NetworkHostAndPort {
	this := NetworkHostAndPort{}
	return &this
}

// GetHost returns the Host field value
func (o *NetworkHostAndPort) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *NetworkHostAndPort) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *NetworkHostAndPort) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *NetworkHostAndPort) GetPort() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *NetworkHostAndPort) GetPortOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *NetworkHostAndPort) SetPort(v float32) {
	o.Port = v
}

func (o NetworkHostAndPort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkHostAndPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["port"] = o.Port
	return toSerialize, nil
}

type NullableNetworkHostAndPort struct {
	value *NetworkHostAndPort
	isSet bool
}

func (v NullableNetworkHostAndPort) Get() *NetworkHostAndPort {
	return v.value
}

func (v *NullableNetworkHostAndPort) Set(val *NetworkHostAndPort) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkHostAndPort) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkHostAndPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkHostAndPort(val *NetworkHostAndPort) *NullableNetworkHostAndPort {
	return &NullableNetworkHostAndPort{value: val, isSet: true}
}

func (v NullableNetworkHostAndPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkHostAndPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


