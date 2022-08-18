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

// TokenMetadataCollection struct for TokenMetadataCollection
type TokenMetadataCollection struct {
	Name string `json:"name"`
	Family string `json:"family"`
}

// NewTokenMetadataCollection instantiates a new TokenMetadataCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenMetadataCollection(name string, family string) *TokenMetadataCollection {
	this := TokenMetadataCollection{}
	this.Name = name
	this.Family = family
	return &this
}

// NewTokenMetadataCollectionWithDefaults instantiates a new TokenMetadataCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenMetadataCollectionWithDefaults() *TokenMetadataCollection {
	this := TokenMetadataCollection{}
	return &this
}

// GetName returns the Name field value
func (o *TokenMetadataCollection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TokenMetadataCollection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TokenMetadataCollection) SetName(v string) {
	o.Name = v
}

// GetFamily returns the Family field value
func (o *TokenMetadataCollection) GetFamily() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *TokenMetadataCollection) GetFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *TokenMetadataCollection) SetFamily(v string) {
	o.Family = v
}

func (o TokenMetadataCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["family"] = o.Family
	}
	return json.Marshal(toSerialize)
}

type NullableTokenMetadataCollection struct {
	value *TokenMetadataCollection
	isSet bool
}

func (v NullableTokenMetadataCollection) Get() *TokenMetadataCollection {
	return v.value
}

func (v *NullableTokenMetadataCollection) Set(val *TokenMetadataCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenMetadataCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenMetadataCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenMetadataCollection(val *TokenMetadataCollection) *NullableTokenMetadataCollection {
	return &NullableTokenMetadataCollection{value: val, isSet: true}
}

func (v NullableTokenMetadataCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenMetadataCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


