/*
Hyperledger Cactus Plugin - Connector Fabric

Can perform basic tasks on a fabric ledger

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-fabric

import (
	"encoding/json"
)

// checks if the DeploymentTargetOrganization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentTargetOrganization{}

// DeploymentTargetOrganization struct for DeploymentTargetOrganization
type DeploymentTargetOrganization struct {
	// Mapped to environment variables of the Fabric CLI container.
	CORE_PEER_LOCALMSPID string `json:"CORE_PEER_LOCALMSPID"`
	// Mapped to environment variables of the Fabric CLI container.
	CORE_PEER_ADDRESS string `json:"CORE_PEER_ADDRESS"`
	// Mapped to environment variables of the Fabric CLI container.
	CORE_PEER_MSPCONFIGPATH string `json:"CORE_PEER_MSPCONFIGPATH"`
	// Mapped to environment variables of the Fabric CLI container.
	CORE_PEER_TLS_ROOTCERT_FILE string `json:"CORE_PEER_TLS_ROOTCERT_FILE"`
	// Mapped to environment variables of the Fabric CLI container.
	ORDERER_TLS_ROOTCERT_FILE string `json:"ORDERER_TLS_ROOTCERT_FILE"`
}

// NewDeploymentTargetOrganization instantiates a new DeploymentTargetOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentTargetOrganization(cOREPEERLOCALMSPID string, cOREPEERADDRESS string, cOREPEERMSPCONFIGPATH string, cOREPEERTLSROOTCERTFILE string, oRDERERTLSROOTCERTFILE string) *DeploymentTargetOrganization {
	this := DeploymentTargetOrganization{}
	this.CORE_PEER_LOCALMSPID = cOREPEERLOCALMSPID
	this.CORE_PEER_ADDRESS = cOREPEERADDRESS
	this.CORE_PEER_MSPCONFIGPATH = cOREPEERMSPCONFIGPATH
	this.CORE_PEER_TLS_ROOTCERT_FILE = cOREPEERTLSROOTCERTFILE
	this.ORDERER_TLS_ROOTCERT_FILE = oRDERERTLSROOTCERTFILE
	return &this
}

// NewDeploymentTargetOrganizationWithDefaults instantiates a new DeploymentTargetOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentTargetOrganizationWithDefaults() *DeploymentTargetOrganization {
	this := DeploymentTargetOrganization{}
	return &this
}

// GetCORE_PEER_LOCALMSPID returns the CORE_PEER_LOCALMSPID field value
func (o *DeploymentTargetOrganization) GetCORE_PEER_LOCALMSPID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CORE_PEER_LOCALMSPID
}

// GetCORE_PEER_LOCALMSPIDOk returns a tuple with the CORE_PEER_LOCALMSPID field value
// and a boolean to check if the value has been set.
func (o *DeploymentTargetOrganization) GetCORE_PEER_LOCALMSPIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CORE_PEER_LOCALMSPID, true
}

// SetCORE_PEER_LOCALMSPID sets field value
func (o *DeploymentTargetOrganization) SetCORE_PEER_LOCALMSPID(v string) {
	o.CORE_PEER_LOCALMSPID = v
}

// GetCORE_PEER_ADDRESS returns the CORE_PEER_ADDRESS field value
func (o *DeploymentTargetOrganization) GetCORE_PEER_ADDRESS() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CORE_PEER_ADDRESS
}

// GetCORE_PEER_ADDRESSOk returns a tuple with the CORE_PEER_ADDRESS field value
// and a boolean to check if the value has been set.
func (o *DeploymentTargetOrganization) GetCORE_PEER_ADDRESSOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CORE_PEER_ADDRESS, true
}

// SetCORE_PEER_ADDRESS sets field value
func (o *DeploymentTargetOrganization) SetCORE_PEER_ADDRESS(v string) {
	o.CORE_PEER_ADDRESS = v
}

// GetCORE_PEER_MSPCONFIGPATH returns the CORE_PEER_MSPCONFIGPATH field value
func (o *DeploymentTargetOrganization) GetCORE_PEER_MSPCONFIGPATH() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CORE_PEER_MSPCONFIGPATH
}

// GetCORE_PEER_MSPCONFIGPATHOk returns a tuple with the CORE_PEER_MSPCONFIGPATH field value
// and a boolean to check if the value has been set.
func (o *DeploymentTargetOrganization) GetCORE_PEER_MSPCONFIGPATHOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CORE_PEER_MSPCONFIGPATH, true
}

// SetCORE_PEER_MSPCONFIGPATH sets field value
func (o *DeploymentTargetOrganization) SetCORE_PEER_MSPCONFIGPATH(v string) {
	o.CORE_PEER_MSPCONFIGPATH = v
}

// GetCORE_PEER_TLS_ROOTCERT_FILE returns the CORE_PEER_TLS_ROOTCERT_FILE field value
func (o *DeploymentTargetOrganization) GetCORE_PEER_TLS_ROOTCERT_FILE() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CORE_PEER_TLS_ROOTCERT_FILE
}

// GetCORE_PEER_TLS_ROOTCERT_FILEOk returns a tuple with the CORE_PEER_TLS_ROOTCERT_FILE field value
// and a boolean to check if the value has been set.
func (o *DeploymentTargetOrganization) GetCORE_PEER_TLS_ROOTCERT_FILEOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CORE_PEER_TLS_ROOTCERT_FILE, true
}

// SetCORE_PEER_TLS_ROOTCERT_FILE sets field value
func (o *DeploymentTargetOrganization) SetCORE_PEER_TLS_ROOTCERT_FILE(v string) {
	o.CORE_PEER_TLS_ROOTCERT_FILE = v
}

// GetORDERER_TLS_ROOTCERT_FILE returns the ORDERER_TLS_ROOTCERT_FILE field value
func (o *DeploymentTargetOrganization) GetORDERER_TLS_ROOTCERT_FILE() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ORDERER_TLS_ROOTCERT_FILE
}

// GetORDERER_TLS_ROOTCERT_FILEOk returns a tuple with the ORDERER_TLS_ROOTCERT_FILE field value
// and a boolean to check if the value has been set.
func (o *DeploymentTargetOrganization) GetORDERER_TLS_ROOTCERT_FILEOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ORDERER_TLS_ROOTCERT_FILE, true
}

// SetORDERER_TLS_ROOTCERT_FILE sets field value
func (o *DeploymentTargetOrganization) SetORDERER_TLS_ROOTCERT_FILE(v string) {
	o.ORDERER_TLS_ROOTCERT_FILE = v
}

func (o DeploymentTargetOrganization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentTargetOrganization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CORE_PEER_LOCALMSPID"] = o.CORE_PEER_LOCALMSPID
	toSerialize["CORE_PEER_ADDRESS"] = o.CORE_PEER_ADDRESS
	toSerialize["CORE_PEER_MSPCONFIGPATH"] = o.CORE_PEER_MSPCONFIGPATH
	toSerialize["CORE_PEER_TLS_ROOTCERT_FILE"] = o.CORE_PEER_TLS_ROOTCERT_FILE
	toSerialize["ORDERER_TLS_ROOTCERT_FILE"] = o.ORDERER_TLS_ROOTCERT_FILE
	return toSerialize, nil
}

type NullableDeploymentTargetOrganization struct {
	value *DeploymentTargetOrganization
	isSet bool
}

func (v NullableDeploymentTargetOrganization) Get() *DeploymentTargetOrganization {
	return v.value
}

func (v *NullableDeploymentTargetOrganization) Set(val *DeploymentTargetOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentTargetOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentTargetOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentTargetOrganization(val *DeploymentTargetOrganization) *NullableDeploymentTargetOrganization {
	return &NullableDeploymentTargetOrganization{value: val, isSet: true}
}

func (v NullableDeploymentTargetOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentTargetOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


