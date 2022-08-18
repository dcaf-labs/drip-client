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

// TokenMetadata struct for TokenMetadata
type TokenMetadata struct {
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	Description string `json:"description"`
	Image string `json:"image"`
	ExternalUrl string `json:"external_url"`
	Collection TokenMetadataCollection `json:"collection"`
}

// NewTokenMetadata instantiates a new TokenMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenMetadata(name string, symbol string, description string, image string, externalUrl string, collection TokenMetadataCollection) *TokenMetadata {
	this := TokenMetadata{}
	this.Name = name
	this.Symbol = symbol
	this.Description = description
	this.Image = image
	this.ExternalUrl = externalUrl
	this.Collection = collection
	return &this
}

// NewTokenMetadataWithDefaults instantiates a new TokenMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenMetadataWithDefaults() *TokenMetadata {
	this := TokenMetadata{}
	return &this
}

// GetName returns the Name field value
func (o *TokenMetadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TokenMetadata) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TokenMetadata) SetName(v string) {
	o.Name = v
}

// GetSymbol returns the Symbol field value
func (o *TokenMetadata) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *TokenMetadata) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *TokenMetadata) SetSymbol(v string) {
	o.Symbol = v
}

// GetDescription returns the Description field value
func (o *TokenMetadata) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TokenMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TokenMetadata) SetDescription(v string) {
	o.Description = v
}

// GetImage returns the Image field value
func (o *TokenMetadata) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *TokenMetadata) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *TokenMetadata) SetImage(v string) {
	o.Image = v
}

// GetExternalUrl returns the ExternalUrl field value
func (o *TokenMetadata) GetExternalUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalUrl
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value
// and a boolean to check if the value has been set.
func (o *TokenMetadata) GetExternalUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUrl, true
}

// SetExternalUrl sets field value
func (o *TokenMetadata) SetExternalUrl(v string) {
	o.ExternalUrl = v
}

// GetCollection returns the Collection field value
func (o *TokenMetadata) GetCollection() TokenMetadataCollection {
	if o == nil {
		var ret TokenMetadataCollection
		return ret
	}

	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value
// and a boolean to check if the value has been set.
func (o *TokenMetadata) GetCollectionOk() (*TokenMetadataCollection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collection, true
}

// SetCollection sets field value
func (o *TokenMetadata) SetCollection(v TokenMetadataCollection) {
	o.Collection = v
}

func (o TokenMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["image"] = o.Image
	}
	if true {
		toSerialize["external_url"] = o.ExternalUrl
	}
	if true {
		toSerialize["collection"] = o.Collection
	}
	return json.Marshal(toSerialize)
}

type NullableTokenMetadata struct {
	value *TokenMetadata
	isSet bool
}

func (v NullableTokenMetadata) Get() *TokenMetadata {
	return v.value
}

func (v *NullableTokenMetadata) Set(val *TokenMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenMetadata(val *TokenMetadata) *NullableTokenMetadata {
	return &NullableTokenMetadata{value: val, isSet: true}
}

func (v NullableTokenMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


