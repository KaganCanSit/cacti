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

// checks if the LocalLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalLog{}

// LocalLog struct for LocalLog
type LocalLog struct {
	Key *string `json:"key,omitempty"`
	SessionID string `json:"sessionID"`
	Data *string `json:"data,omitempty"`
	Type string `json:"type"`
	Operation string `json:"operation"`
	Timestamp *string `json:"timestamp,omitempty"`
}

// NewLocalLog instantiates a new LocalLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalLog(sessionID string, type_ string, operation string) *LocalLog {
	this := LocalLog{}
	this.SessionID = sessionID
	this.Type = type_
	this.Operation = operation
	return &this
}

// NewLocalLogWithDefaults instantiates a new LocalLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalLogWithDefaults() *LocalLog {
	this := LocalLog{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *LocalLog) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalLog) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *LocalLog) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *LocalLog) SetKey(v string) {
	o.Key = &v
}

// GetSessionID returns the SessionID field value
func (o *LocalLog) GetSessionID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionID
}

// GetSessionIDOk returns a tuple with the SessionID field value
// and a boolean to check if the value has been set.
func (o *LocalLog) GetSessionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionID, true
}

// SetSessionID sets field value
func (o *LocalLog) SetSessionID(v string) {
	o.SessionID = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LocalLog) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalLog) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LocalLog) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *LocalLog) SetData(v string) {
	o.Data = &v
}

// GetType returns the Type field value
func (o *LocalLog) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LocalLog) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LocalLog) SetType(v string) {
	o.Type = v
}

// GetOperation returns the Operation field value
func (o *LocalLog) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *LocalLog) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *LocalLog) SetOperation(v string) {
	o.Operation = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *LocalLog) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalLog) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *LocalLog) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *LocalLog) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o LocalLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	toSerialize["sessionID"] = o.SessionID
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["type"] = o.Type
	toSerialize["operation"] = o.Operation
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableLocalLog struct {
	value *LocalLog
	isSet bool
}

func (v NullableLocalLog) Get() *LocalLog {
	return v.value
}

func (v *NullableLocalLog) Set(val *LocalLog) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalLog) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalLog(val *LocalLog) *NullableLocalLog {
	return &NullableLocalLog{value: val, isSet: true}
}

func (v NullableLocalLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


