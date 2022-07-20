# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | **string** |  | 
**Symbol** | Pointer to **string** |  | [optional] 
**Decimals** | **int32** |  | 

## Methods

### NewToken

`func NewToken(pubkey string, decimals int32, ) *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *Token) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *Token) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *Token) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.


### GetSymbol

`func (o *Token) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Token) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Token) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Token) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDecimals

`func (o *Token) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *Token) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *Token) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


