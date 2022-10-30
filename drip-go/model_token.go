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

// Token struct for Token
type Token struct {
	Pubkey string `json:"pubkey"`
	Decimals int32 `json:"decimals"`
	Symbol *string `json:"symbol,omitempty"`
	IconUrl *string `json:"iconUrl,omitempty"`
	CoinGeckoId *string `json:"coinGeckoId,omitempty"`
}

// NewToken instantiates a new Token object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToken(pubkey string, decimals int32) *Token {
	this := Token{}
	this.Pubkey = pubkey
	this.Decimals = decimals
	return &this
}

// NewTokenWithDefaults instantiates a new Token object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenWithDefaults() *Token {
	this := Token{}
	return &this
}

// GetPubkey returns the Pubkey field value
func (o *Token) GetPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pubkey
}

// GetPubkeyOk returns a tuple with the Pubkey field value
// and a boolean to check if the value has been set.
func (o *Token) GetPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pubkey, true
}

// SetPubkey sets field value
func (o *Token) SetPubkey(v string) {
	o.Pubkey = v
}

// GetDecimals returns the Decimals field value
func (o *Token) GetDecimals() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value
// and a boolean to check if the value has been set.
func (o *Token) GetDecimalsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decimals, true
}

// SetDecimals sets field value
func (o *Token) SetDecimals(v int32) {
	o.Decimals = v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *Token) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *Token) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *Token) SetSymbol(v string) {
	o.Symbol = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *Token) GetIconUrl() string {
	if o == nil || o.IconUrl == nil {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetIconUrlOk() (*string, bool) {
	if o == nil || o.IconUrl == nil {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *Token) HasIconUrl() bool {
	if o != nil && o.IconUrl != nil {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *Token) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetCoinGeckoId returns the CoinGeckoId field value if set, zero value otherwise.
func (o *Token) GetCoinGeckoId() string {
	if o == nil || o.CoinGeckoId == nil {
		var ret string
		return ret
	}
	return *o.CoinGeckoId
}

// GetCoinGeckoIdOk returns a tuple with the CoinGeckoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetCoinGeckoIdOk() (*string, bool) {
	if o == nil || o.CoinGeckoId == nil {
		return nil, false
	}
	return o.CoinGeckoId, true
}

// HasCoinGeckoId returns a boolean if a field has been set.
func (o *Token) HasCoinGeckoId() bool {
	if o != nil && o.CoinGeckoId != nil {
		return true
	}

	return false
}

// SetCoinGeckoId gets a reference to the given string and assigns it to the CoinGeckoId field.
func (o *Token) SetCoinGeckoId(v string) {
	o.CoinGeckoId = &v
}

func (o Token) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pubkey"] = o.Pubkey
	}
	if true {
		toSerialize["decimals"] = o.Decimals
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.IconUrl != nil {
		toSerialize["iconUrl"] = o.IconUrl
	}
	if o.CoinGeckoId != nil {
		toSerialize["coinGeckoId"] = o.CoinGeckoId
	}
	return json.Marshal(toSerialize)
}

type NullableToken struct {
	value *Token
	isSet bool
}

func (v NullableToken) Get() *Token {
	return v.value
}

func (v *NullableToken) Set(val *Token) {
	v.value = val
	v.isSet = true
}

func (v NullableToken) IsSet() bool {
	return v.isSet
}

func (v *NullableToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToken(val *Token) *NullableToken {
	return &NullableToken{value: val, isSet: true}
}

func (v NullableToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


