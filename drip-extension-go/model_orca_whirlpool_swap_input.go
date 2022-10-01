/*
Drip Cloud Functions

Drip auxilary backend.

API version: 1.0.0
Contact: team@dcaf.so
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dripextension

import (
	"encoding/json"
)

// OrcaWhirlpoolSwapInput struct for OrcaWhirlpoolSwapInput
type OrcaWhirlpoolSwapInput struct {
	Amount string `json:"amount"`
	OtherAmountThreshold string `json:"otherAmountThreshold"`
	SqrtPriceLimit string `json:"sqrtPriceLimit"`
	AmountSpecifiedIsInput bool `json:"amountSpecifiedIsInput"`
	AToB bool `json:"aToB"`
	TickArray0 string `json:"tickArray0"`
	TickArray1 string `json:"tickArray1"`
	TickArray2 string `json:"tickArray2"`
}

// NewOrcaWhirlpoolSwapInput instantiates a new OrcaWhirlpoolSwapInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrcaWhirlpoolSwapInput(amount string, otherAmountThreshold string, sqrtPriceLimit string, amountSpecifiedIsInput bool, aToB bool, tickArray0 string, tickArray1 string, tickArray2 string) *OrcaWhirlpoolSwapInput {
	this := OrcaWhirlpoolSwapInput{}
	this.Amount = amount
	this.OtherAmountThreshold = otherAmountThreshold
	this.SqrtPriceLimit = sqrtPriceLimit
	this.AmountSpecifiedIsInput = amountSpecifiedIsInput
	this.AToB = aToB
	this.TickArray0 = tickArray0
	this.TickArray1 = tickArray1
	this.TickArray2 = tickArray2
	return &this
}

// NewOrcaWhirlpoolSwapInputWithDefaults instantiates a new OrcaWhirlpoolSwapInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrcaWhirlpoolSwapInputWithDefaults() *OrcaWhirlpoolSwapInput {
	this := OrcaWhirlpoolSwapInput{}
	return &this
}

// GetAmount returns the Amount field value
func (o *OrcaWhirlpoolSwapInput) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolSwapInput) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *OrcaWhirlpoolSwapInput) SetAmount(v string) {
	o.Amount = v
}

// GetOtherAmountThreshold returns the OtherAmountThreshold field value
func (o *OrcaWhirlpoolSwapInput) GetOtherAmountThreshold() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OtherAmountThreshold
}

// GetOtherAmountThresholdOk returns a tuple with the OtherAmountThreshold field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolSwapInput) GetOtherAmountThresholdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OtherAmountThreshold, true
}

// SetOtherAmountThreshold sets field value
func (o *OrcaWhirlpoolSwapInput) SetOtherAmountThreshold(v string) {
	o.OtherAmountThreshold = v
}

// GetSqrtPriceLimit returns the SqrtPriceLimit field value
func (o *OrcaWhirlpoolSwapInput) GetSqrtPriceLimit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SqrtPriceLimit
}

// GetSqrtPriceLimitOk returns a tuple with the SqrtPriceLimit field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolSwapInput) GetSqrtPriceLimitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SqrtPriceLimit, true
}

// SetSqrtPriceLimit sets field value
func (o *OrcaWhirlpoolSwapInput) SetSqrtPriceLimit(v string) {
	o.SqrtPriceLimit = v
}

// GetAmountSpecifiedIsInput returns the AmountSpecifiedIsInput field value
func (o *OrcaWhirlpoolSwapInput) GetAmountSpecifiedIsInput() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AmountSpecifiedIsInput
}

// GetAmountSpecifiedIsInputOk returns a tuple with the AmountSpecifiedIsInput field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolSwapInput) GetAmountSpecifiedIsInputOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountSpecifiedIsInput, true
}

// SetAmountSpecifiedIsInput sets field value
func (o *OrcaWhirlpoolSwapInput) SetAmountSpecifiedIsInput(v bool) {
	o.AmountSpecifiedIsInput = v
}

// GetAToB returns the AToB field value
func (o *OrcaWhirlpoolSwapInput) GetAToB() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AToB
}

// GetAToBOk returns a tuple with the AToB field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolSwapInput) GetAToBOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AToB, true
}

// SetAToB sets field value
func (o *OrcaWhirlpoolSwapInput) SetAToB(v bool) {
	o.AToB = v
}

// GetTickArray0 returns the TickArray0 field value
func (o *OrcaWhirlpoolSwapInput) GetTickArray0() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TickArray0
}

// GetTickArray0Ok returns a tuple with the TickArray0 field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolSwapInput) GetTickArray0Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TickArray0, true
}

// SetTickArray0 sets field value
func (o *OrcaWhirlpoolSwapInput) SetTickArray0(v string) {
	o.TickArray0 = v
}

// GetTickArray1 returns the TickArray1 field value
func (o *OrcaWhirlpoolSwapInput) GetTickArray1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TickArray1
}

// GetTickArray1Ok returns a tuple with the TickArray1 field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolSwapInput) GetTickArray1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TickArray1, true
}

// SetTickArray1 sets field value
func (o *OrcaWhirlpoolSwapInput) SetTickArray1(v string) {
	o.TickArray1 = v
}

// GetTickArray2 returns the TickArray2 field value
func (o *OrcaWhirlpoolSwapInput) GetTickArray2() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TickArray2
}

// GetTickArray2Ok returns a tuple with the TickArray2 field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolSwapInput) GetTickArray2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TickArray2, true
}

// SetTickArray2 sets field value
func (o *OrcaWhirlpoolSwapInput) SetTickArray2(v string) {
	o.TickArray2 = v
}

func (o OrcaWhirlpoolSwapInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["otherAmountThreshold"] = o.OtherAmountThreshold
	}
	if true {
		toSerialize["sqrtPriceLimit"] = o.SqrtPriceLimit
	}
	if true {
		toSerialize["amountSpecifiedIsInput"] = o.AmountSpecifiedIsInput
	}
	if true {
		toSerialize["aToB"] = o.AToB
	}
	if true {
		toSerialize["tickArray0"] = o.TickArray0
	}
	if true {
		toSerialize["tickArray1"] = o.TickArray1
	}
	if true {
		toSerialize["tickArray2"] = o.TickArray2
	}
	return json.Marshal(toSerialize)
}

type NullableOrcaWhirlpoolSwapInput struct {
	value *OrcaWhirlpoolSwapInput
	isSet bool
}

func (v NullableOrcaWhirlpoolSwapInput) Get() *OrcaWhirlpoolSwapInput {
	return v.value
}

func (v *NullableOrcaWhirlpoolSwapInput) Set(val *OrcaWhirlpoolSwapInput) {
	v.value = val
	v.isSet = true
}

func (v NullableOrcaWhirlpoolSwapInput) IsSet() bool {
	return v.isSet
}

func (v *NullableOrcaWhirlpoolSwapInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrcaWhirlpoolSwapInput(val *OrcaWhirlpoolSwapInput) *NullableOrcaWhirlpoolSwapInput {
	return &NullableOrcaWhirlpoolSwapInput{value: val, isSet: true}
}

func (v NullableOrcaWhirlpoolSwapInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrcaWhirlpoolSwapInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


