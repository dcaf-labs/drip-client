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

// DripCommon struct for DripCommon
type DripCommon struct {
	Vault string `json:"vault"`
	VaultProtoConfig string `json:"vaultProtoConfig"`
	VaultTokenAAccount string `json:"vaultTokenAAccount"`
	VaultTokenBAccount string `json:"vaultTokenBAccount"`
	TokenAMint string `json:"tokenAMint"`
	TokenBMint string `json:"tokenBMint"`
}

// NewDripCommon instantiates a new DripCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDripCommon(vault string, vaultProtoConfig string, vaultTokenAAccount string, vaultTokenBAccount string, tokenAMint string, tokenBMint string) *DripCommon {
	this := DripCommon{}
	this.Vault = vault
	this.VaultProtoConfig = vaultProtoConfig
	this.VaultTokenAAccount = vaultTokenAAccount
	this.VaultTokenBAccount = vaultTokenBAccount
	this.TokenAMint = tokenAMint
	this.TokenBMint = tokenBMint
	return &this
}

// NewDripCommonWithDefaults instantiates a new DripCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDripCommonWithDefaults() *DripCommon {
	this := DripCommon{}
	return &this
}

// GetVault returns the Vault field value
func (o *DripCommon) GetVault() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vault
}

// GetVaultOk returns a tuple with the Vault field value
// and a boolean to check if the value has been set.
func (o *DripCommon) GetVaultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vault, true
}

// SetVault sets field value
func (o *DripCommon) SetVault(v string) {
	o.Vault = v
}

// GetVaultProtoConfig returns the VaultProtoConfig field value
func (o *DripCommon) GetVaultProtoConfig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultProtoConfig
}

// GetVaultProtoConfigOk returns a tuple with the VaultProtoConfig field value
// and a boolean to check if the value has been set.
func (o *DripCommon) GetVaultProtoConfigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultProtoConfig, true
}

// SetVaultProtoConfig sets field value
func (o *DripCommon) SetVaultProtoConfig(v string) {
	o.VaultProtoConfig = v
}

// GetVaultTokenAAccount returns the VaultTokenAAccount field value
func (o *DripCommon) GetVaultTokenAAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultTokenAAccount
}

// GetVaultTokenAAccountOk returns a tuple with the VaultTokenAAccount field value
// and a boolean to check if the value has been set.
func (o *DripCommon) GetVaultTokenAAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultTokenAAccount, true
}

// SetVaultTokenAAccount sets field value
func (o *DripCommon) SetVaultTokenAAccount(v string) {
	o.VaultTokenAAccount = v
}

// GetVaultTokenBAccount returns the VaultTokenBAccount field value
func (o *DripCommon) GetVaultTokenBAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultTokenBAccount
}

// GetVaultTokenBAccountOk returns a tuple with the VaultTokenBAccount field value
// and a boolean to check if the value has been set.
func (o *DripCommon) GetVaultTokenBAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultTokenBAccount, true
}

// SetVaultTokenBAccount sets field value
func (o *DripCommon) SetVaultTokenBAccount(v string) {
	o.VaultTokenBAccount = v
}

// GetTokenAMint returns the TokenAMint field value
func (o *DripCommon) GetTokenAMint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenAMint
}

// GetTokenAMintOk returns a tuple with the TokenAMint field value
// and a boolean to check if the value has been set.
func (o *DripCommon) GetTokenAMintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenAMint, true
}

// SetTokenAMint sets field value
func (o *DripCommon) SetTokenAMint(v string) {
	o.TokenAMint = v
}

// GetTokenBMint returns the TokenBMint field value
func (o *DripCommon) GetTokenBMint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenBMint
}

// GetTokenBMintOk returns a tuple with the TokenBMint field value
// and a boolean to check if the value has been set.
func (o *DripCommon) GetTokenBMintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenBMint, true
}

// SetTokenBMint sets field value
func (o *DripCommon) SetTokenBMint(v string) {
	o.TokenBMint = v
}

func (o DripCommon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vault"] = o.Vault
	}
	if true {
		toSerialize["vaultProtoConfig"] = o.VaultProtoConfig
	}
	if true {
		toSerialize["vaultTokenAAccount"] = o.VaultTokenAAccount
	}
	if true {
		toSerialize["vaultTokenBAccount"] = o.VaultTokenBAccount
	}
	if true {
		toSerialize["tokenAMint"] = o.TokenAMint
	}
	if true {
		toSerialize["tokenBMint"] = o.TokenBMint
	}
	return json.Marshal(toSerialize)
}

type NullableDripCommon struct {
	value *DripCommon
	isSet bool
}

func (v NullableDripCommon) Get() *DripCommon {
	return v.value
}

func (v *NullableDripCommon) Set(val *DripCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableDripCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableDripCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDripCommon(val *DripCommon) *NullableDripCommon {
	return &NullableDripCommon{value: val, isSet: true}
}

func (v NullableDripCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDripCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


