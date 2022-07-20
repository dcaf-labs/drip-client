# TokenSwap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | **string** |  | 
**Mint** | **string** |  | 
**Authority** | **string** |  | 
**FeeAccount** | **string** |  | 
**TokenAAccount** | **string** |  | 
**TokenBAccount** | **string** |  | 
**Pair** | **string** | token pair reference identifier | 

## Methods

### NewTokenSwap

`func NewTokenSwap(pubkey string, mint string, authority string, feeAccount string, tokenAAccount string, tokenBAccount string, pair string, ) *TokenSwap`

NewTokenSwap instantiates a new TokenSwap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenSwapWithDefaults

`func NewTokenSwapWithDefaults() *TokenSwap`

NewTokenSwapWithDefaults instantiates a new TokenSwap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *TokenSwap) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *TokenSwap) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *TokenSwap) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.


### GetMint

`func (o *TokenSwap) GetMint() string`

GetMint returns the Mint field if non-nil, zero value otherwise.

### GetMintOk

`func (o *TokenSwap) GetMintOk() (*string, bool)`

GetMintOk returns a tuple with the Mint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMint

`func (o *TokenSwap) SetMint(v string)`

SetMint sets Mint field to given value.


### GetAuthority

`func (o *TokenSwap) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *TokenSwap) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *TokenSwap) SetAuthority(v string)`

SetAuthority sets Authority field to given value.


### GetFeeAccount

`func (o *TokenSwap) GetFeeAccount() string`

GetFeeAccount returns the FeeAccount field if non-nil, zero value otherwise.

### GetFeeAccountOk

`func (o *TokenSwap) GetFeeAccountOk() (*string, bool)`

GetFeeAccountOk returns a tuple with the FeeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAccount

`func (o *TokenSwap) SetFeeAccount(v string)`

SetFeeAccount sets FeeAccount field to given value.


### GetTokenAAccount

`func (o *TokenSwap) GetTokenAAccount() string`

GetTokenAAccount returns the TokenAAccount field if non-nil, zero value otherwise.

### GetTokenAAccountOk

`func (o *TokenSwap) GetTokenAAccountOk() (*string, bool)`

GetTokenAAccountOk returns a tuple with the TokenAAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAAccount

`func (o *TokenSwap) SetTokenAAccount(v string)`

SetTokenAAccount sets TokenAAccount field to given value.


### GetTokenBAccount

`func (o *TokenSwap) GetTokenBAccount() string`

GetTokenBAccount returns the TokenBAccount field if non-nil, zero value otherwise.

### GetTokenBAccountOk

`func (o *TokenSwap) GetTokenBAccountOk() (*string, bool)`

GetTokenBAccountOk returns a tuple with the TokenBAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBAccount

`func (o *TokenSwap) SetTokenBAccount(v string)`

SetTokenBAccount sets TokenBAccount field to given value.


### GetPair

`func (o *TokenSwap) GetPair() string`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *TokenSwap) GetPairOk() (*string, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *TokenSwap) SetPair(v string)`

SetPair sets Pair field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


