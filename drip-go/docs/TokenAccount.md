# TokenAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | **string** |  | 
**Mint** | **string** |  | 
**Owner** | **string** |  | 
**Amount** | **string** |  | 
**State** | **string** |  | 

## Methods

### NewTokenAccount

`func NewTokenAccount(pubkey string, mint string, owner string, amount string, state string, ) *TokenAccount`

NewTokenAccount instantiates a new TokenAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenAccountWithDefaults

`func NewTokenAccountWithDefaults() *TokenAccount`

NewTokenAccountWithDefaults instantiates a new TokenAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *TokenAccount) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *TokenAccount) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *TokenAccount) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.


### GetMint

`func (o *TokenAccount) GetMint() string`

GetMint returns the Mint field if non-nil, zero value otherwise.

### GetMintOk

`func (o *TokenAccount) GetMintOk() (*string, bool)`

GetMintOk returns a tuple with the Mint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMint

`func (o *TokenAccount) SetMint(v string)`

SetMint sets Mint field to given value.


### GetOwner

`func (o *TokenAccount) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *TokenAccount) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *TokenAccount) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetAmount

`func (o *TokenAccount) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TokenAccount) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TokenAccount) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetState

`func (o *TokenAccount) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TokenAccount) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TokenAccount) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


