/*
Drip Backend

Drip backend service. All API's have a IP rate limit of 10 requests per second. 

API version: 1.0.0
Contact: mocha@dcaf.so
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package drip

import (
	"encoding/json"
)

// ExpandedAdminPosition Get Position with Expanded Data
type ExpandedAdminPosition struct {
	Pubkey string `json:"pubkey"`
	Vault Vault `json:"vault"`
	Authority string `json:"authority"`
	DepositedTokenAAmount string `json:"depositedTokenAAmount"`
	WithdrawnTokenBAmount string `json:"withdrawnTokenBAmount"`
	DepositTimestamp string `json:"depositTimestamp"`
	DcaPeriodIdBeforeDeposit string `json:"dcaPeriodIdBeforeDeposit"`
	NumberOfSwaps string `json:"numberOfSwaps"`
	PeriodicDripAmount string `json:"periodicDripAmount"`
	IsClosed bool `json:"isClosed"`
	ProtoConfig *ProtoConfig `json:"protoConfig,omitempty"`
	TokenA *Token `json:"tokenA,omitempty"`
	TokenB *Token `json:"tokenB,omitempty"`
}

// NewExpandedAdminPosition instantiates a new ExpandedAdminPosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpandedAdminPosition(pubkey string, vault Vault, authority string, depositedTokenAAmount string, withdrawnTokenBAmount string, depositTimestamp string, dcaPeriodIdBeforeDeposit string, numberOfSwaps string, periodicDripAmount string, isClosed bool) *ExpandedAdminPosition {
	this := ExpandedAdminPosition{}
	this.Pubkey = pubkey
	this.Vault = vault
	this.Authority = authority
	this.DepositedTokenAAmount = depositedTokenAAmount
	this.WithdrawnTokenBAmount = withdrawnTokenBAmount
	this.DepositTimestamp = depositTimestamp
	this.DcaPeriodIdBeforeDeposit = dcaPeriodIdBeforeDeposit
	this.NumberOfSwaps = numberOfSwaps
	this.PeriodicDripAmount = periodicDripAmount
	this.IsClosed = isClosed
	return &this
}

// NewExpandedAdminPositionWithDefaults instantiates a new ExpandedAdminPosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpandedAdminPositionWithDefaults() *ExpandedAdminPosition {
	this := ExpandedAdminPosition{}
	return &this
}

// GetPubkey returns the Pubkey field value
func (o *ExpandedAdminPosition) GetPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pubkey
}

// GetPubkeyOk returns a tuple with the Pubkey field value
// and a boolean to check if the value has been set.
func (o *ExpandedAdminPosition) GetPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pubkey, true
}

// SetPubkey sets field value
func (o *ExpandedAdminPosition) SetPubkey(v string) {
	o.Pubkey = v
}

// GetVault returns the Vault field value
func (o *ExpandedAdminPosition) GetVault() Vault {
	if o == nil {
		var ret Vault
		return ret
	}

	return o.Vault
}

// GetVaultOk returns a tuple with the Vault field value
// and a boolean to check if the value has been set.
func (o *ExpandedAdminPosition) GetVaultOk() (*Vault, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vault, true
}

// SetVault sets field value
func (o *ExpandedAdminPosition) SetVault(v Vault) {
	o.Vault = v
}

// GetAuthority returns the Authority field value
func (o *ExpandedAdminPosition) GetAuthority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Authority
}

// GetAuthorityOk returns a tuple with the Authority field value
// and a boolean to check if the value has been set.
func (o *ExpandedAdminPosition) GetAuthorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authority, true
}

// SetAuthority sets field value
func (o *ExpandedAdminPosition) SetAuthority(v string) {
	o.Authority = v
}

// GetDepositedTokenAAmount returns the DepositedTokenAAmount field value
func (o *ExpandedAdminPosition) GetDepositedTokenAAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepositedTokenAAmount
}

// GetDepositedTokenAAmountOk returns a tuple with the DepositedTokenAAmount field value
// and a boolean to check if the value has been set.
func (o *ExpandedAdminPosition) GetDepositedTokenAAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepositedTokenAAmount, true
}

// SetDepositedTokenAAmount sets field value
func (o *ExpandedAdminPosition) SetDepositedTokenAAmount(v string) {
	o.DepositedTokenAAmount = v
}

// GetWithdrawnTokenBAmount returns the WithdrawnTokenBAmount field value
func (o *ExpandedAdminPosition) GetWithdrawnTokenBAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WithdrawnTokenBAmount
}

// GetWithdrawnTokenBAmountOk returns a tuple with the WithdrawnTokenBAmount field value
// and a boolean to check if the value has been set.
func (o *ExpandedAdminPosition) GetWithdrawnTokenBAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WithdrawnTokenBAmount, true
}

// SetWithdrawnTokenBAmount sets field value
func (o *ExpandedAdminPosition) SetWithdrawnTokenBAmount(v string) {
	o.WithdrawnTokenBAmount = v
}

// GetDepositTimestamp returns the DepositTimestamp field value
func (o *ExpandedAdminPosition) GetDepositTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepositTimestamp
}

// GetDepositTimestampOk returns a tuple with the DepositTimestamp field value
// and a boolean to check if the value has been set.
func (o *ExpandedAdminPosition) GetDepositTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepositTimestamp, true
}

// SetDepositTimestamp sets field value
func (o *ExpandedAdminPosition) SetDepositTimestamp(v string) {
	o.DepositTimestamp = v
}

// GetDcaPeriodIdBeforeDeposit returns the DcaPeriodIdBeforeDeposit field value
func (o *ExpandedAdminPosition) GetDcaPeriodIdBeforeDeposit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DcaPeriodIdBeforeDeposit
}

// GetDcaPeriodIdBeforeDepositOk returns a tuple with the DcaPeriodIdBeforeDeposit field value
// and a boolean to check if the value has been set.
func (o *ExpandedAdminPosition) GetDcaPeriodIdBeforeDepositOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DcaPeriodIdBeforeDeposit, true
}

// SetDcaPeriodIdBeforeDeposit sets field value
func (o *ExpandedAdminPosition) SetDcaPeriodIdBeforeDeposit(v string) {
	o.DcaPeriodIdBeforeDeposit = v
}

// GetNumberOfSwaps returns the NumberOfSwaps field value
func (o *ExpandedAdminPosition) GetNumberOfSwaps() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NumberOfSwaps
}

// GetNumberOfSwapsOk returns a tuple with the NumberOfSwaps field value
// and a boolean to check if the value has been set.
func (o *ExpandedAdminPosition) GetNumberOfSwapsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfSwaps, true
}

// SetNumberOfSwaps sets field value
func (o *ExpandedAdminPosition) SetNumberOfSwaps(v string) {
	o.NumberOfSwaps = v
}

// GetPeriodicDripAmount returns the PeriodicDripAmount field value
func (o *ExpandedAdminPosition) GetPeriodicDripAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PeriodicDripAmount
}

// GetPeriodicDripAmountOk returns a tuple with the PeriodicDripAmount field value
// and a boolean to check if the value has been set.
func (o *ExpandedAdminPosition) GetPeriodicDripAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodicDripAmount, true
}

// SetPeriodicDripAmount sets field value
func (o *ExpandedAdminPosition) SetPeriodicDripAmount(v string) {
	o.PeriodicDripAmount = v
}

// GetIsClosed returns the IsClosed field value
func (o *ExpandedAdminPosition) GetIsClosed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsClosed
}

// GetIsClosedOk returns a tuple with the IsClosed field value
// and a boolean to check if the value has been set.
func (o *ExpandedAdminPosition) GetIsClosedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsClosed, true
}

// SetIsClosed sets field value
func (o *ExpandedAdminPosition) SetIsClosed(v bool) {
	o.IsClosed = v
}

// GetProtoConfig returns the ProtoConfig field value if set, zero value otherwise.
func (o *ExpandedAdminPosition) GetProtoConfig() ProtoConfig {
	if o == nil || o.ProtoConfig == nil {
		var ret ProtoConfig
		return ret
	}
	return *o.ProtoConfig
}

// GetProtoConfigOk returns a tuple with the ProtoConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandedAdminPosition) GetProtoConfigOk() (*ProtoConfig, bool) {
	if o == nil || o.ProtoConfig == nil {
		return nil, false
	}
	return o.ProtoConfig, true
}

// HasProtoConfig returns a boolean if a field has been set.
func (o *ExpandedAdminPosition) HasProtoConfig() bool {
	if o != nil && o.ProtoConfig != nil {
		return true
	}

	return false
}

// SetProtoConfig gets a reference to the given ProtoConfig and assigns it to the ProtoConfig field.
func (o *ExpandedAdminPosition) SetProtoConfig(v ProtoConfig) {
	o.ProtoConfig = &v
}

// GetTokenA returns the TokenA field value if set, zero value otherwise.
func (o *ExpandedAdminPosition) GetTokenA() Token {
	if o == nil || o.TokenA == nil {
		var ret Token
		return ret
	}
	return *o.TokenA
}

// GetTokenAOk returns a tuple with the TokenA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandedAdminPosition) GetTokenAOk() (*Token, bool) {
	if o == nil || o.TokenA == nil {
		return nil, false
	}
	return o.TokenA, true
}

// HasTokenA returns a boolean if a field has been set.
func (o *ExpandedAdminPosition) HasTokenA() bool {
	if o != nil && o.TokenA != nil {
		return true
	}

	return false
}

// SetTokenA gets a reference to the given Token and assigns it to the TokenA field.
func (o *ExpandedAdminPosition) SetTokenA(v Token) {
	o.TokenA = &v
}

// GetTokenB returns the TokenB field value if set, zero value otherwise.
func (o *ExpandedAdminPosition) GetTokenB() Token {
	if o == nil || o.TokenB == nil {
		var ret Token
		return ret
	}
	return *o.TokenB
}

// GetTokenBOk returns a tuple with the TokenB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandedAdminPosition) GetTokenBOk() (*Token, bool) {
	if o == nil || o.TokenB == nil {
		return nil, false
	}
	return o.TokenB, true
}

// HasTokenB returns a boolean if a field has been set.
func (o *ExpandedAdminPosition) HasTokenB() bool {
	if o != nil && o.TokenB != nil {
		return true
	}

	return false
}

// SetTokenB gets a reference to the given Token and assigns it to the TokenB field.
func (o *ExpandedAdminPosition) SetTokenB(v Token) {
	o.TokenB = &v
}

func (o ExpandedAdminPosition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pubkey"] = o.Pubkey
	}
	if true {
		toSerialize["vault"] = o.Vault
	}
	if true {
		toSerialize["authority"] = o.Authority
	}
	if true {
		toSerialize["depositedTokenAAmount"] = o.DepositedTokenAAmount
	}
	if true {
		toSerialize["withdrawnTokenBAmount"] = o.WithdrawnTokenBAmount
	}
	if true {
		toSerialize["depositTimestamp"] = o.DepositTimestamp
	}
	if true {
		toSerialize["dcaPeriodIdBeforeDeposit"] = o.DcaPeriodIdBeforeDeposit
	}
	if true {
		toSerialize["numberOfSwaps"] = o.NumberOfSwaps
	}
	if true {
		toSerialize["periodicDripAmount"] = o.PeriodicDripAmount
	}
	if true {
		toSerialize["isClosed"] = o.IsClosed
	}
	if o.ProtoConfig != nil {
		toSerialize["protoConfig"] = o.ProtoConfig
	}
	if o.TokenA != nil {
		toSerialize["tokenA"] = o.TokenA
	}
	if o.TokenB != nil {
		toSerialize["tokenB"] = o.TokenB
	}
	return json.Marshal(toSerialize)
}

type NullableExpandedAdminPosition struct {
	value *ExpandedAdminPosition
	isSet bool
}

func (v NullableExpandedAdminPosition) Get() *ExpandedAdminPosition {
	return v.value
}

func (v *NullableExpandedAdminPosition) Set(val *ExpandedAdminPosition) {
	v.value = val
	v.isSet = true
}

func (v NullableExpandedAdminPosition) IsSet() bool {
	return v.isSet
}

func (v *NullableExpandedAdminPosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpandedAdminPosition(val *ExpandedAdminPosition) *NullableExpandedAdminPosition {
	return &NullableExpandedAdminPosition{value: val, isSet: true}
}

func (v NullableExpandedAdminPosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpandedAdminPosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

