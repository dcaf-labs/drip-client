# ListVaultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | **string** |  | 
**ProtoConfig** | **string** |  | 
**TokenAAccount** | **string** |  | 
**TokenBAccount** | **string** |  | 
**TreasuryTokenBAccount** | **string** |  | 
**TokenAMint** | **string** |  | 
**TokenBMint** | **string** |  | 
**LastDcaPeriod** | **string** |  | 
**DripAmount** | **string** |  | 
**DcaActivationTimestamp** | **string** | unix timestamp | 

## Methods

### NewListVaultsInner

`func NewListVaultsInner(pubkey string, protoConfig string, tokenAAccount string, tokenBAccount string, treasuryTokenBAccount string, tokenAMint string, tokenBMint string, lastDcaPeriod string, dripAmount string, dcaActivationTimestamp string, ) *ListVaultsInner`

NewListVaultsInner instantiates a new ListVaultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVaultsInnerWithDefaults

`func NewListVaultsInnerWithDefaults() *ListVaultsInner`

NewListVaultsInnerWithDefaults instantiates a new ListVaultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *ListVaultsInner) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *ListVaultsInner) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *ListVaultsInner) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.


### GetProtoConfig

`func (o *ListVaultsInner) GetProtoConfig() string`

GetProtoConfig returns the ProtoConfig field if non-nil, zero value otherwise.

### GetProtoConfigOk

`func (o *ListVaultsInner) GetProtoConfigOk() (*string, bool)`

GetProtoConfigOk returns a tuple with the ProtoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoConfig

`func (o *ListVaultsInner) SetProtoConfig(v string)`

SetProtoConfig sets ProtoConfig field to given value.


### GetTokenAAccount

`func (o *ListVaultsInner) GetTokenAAccount() string`

GetTokenAAccount returns the TokenAAccount field if non-nil, zero value otherwise.

### GetTokenAAccountOk

`func (o *ListVaultsInner) GetTokenAAccountOk() (*string, bool)`

GetTokenAAccountOk returns a tuple with the TokenAAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAAccount

`func (o *ListVaultsInner) SetTokenAAccount(v string)`

SetTokenAAccount sets TokenAAccount field to given value.


### GetTokenBAccount

`func (o *ListVaultsInner) GetTokenBAccount() string`

GetTokenBAccount returns the TokenBAccount field if non-nil, zero value otherwise.

### GetTokenBAccountOk

`func (o *ListVaultsInner) GetTokenBAccountOk() (*string, bool)`

GetTokenBAccountOk returns a tuple with the TokenBAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBAccount

`func (o *ListVaultsInner) SetTokenBAccount(v string)`

SetTokenBAccount sets TokenBAccount field to given value.


### GetTreasuryTokenBAccount

`func (o *ListVaultsInner) GetTreasuryTokenBAccount() string`

GetTreasuryTokenBAccount returns the TreasuryTokenBAccount field if non-nil, zero value otherwise.

### GetTreasuryTokenBAccountOk

`func (o *ListVaultsInner) GetTreasuryTokenBAccountOk() (*string, bool)`

GetTreasuryTokenBAccountOk returns a tuple with the TreasuryTokenBAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreasuryTokenBAccount

`func (o *ListVaultsInner) SetTreasuryTokenBAccount(v string)`

SetTreasuryTokenBAccount sets TreasuryTokenBAccount field to given value.


### GetTokenAMint

`func (o *ListVaultsInner) GetTokenAMint() string`

GetTokenAMint returns the TokenAMint field if non-nil, zero value otherwise.

### GetTokenAMintOk

`func (o *ListVaultsInner) GetTokenAMintOk() (*string, bool)`

GetTokenAMintOk returns a tuple with the TokenAMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAMint

`func (o *ListVaultsInner) SetTokenAMint(v string)`

SetTokenAMint sets TokenAMint field to given value.


### GetTokenBMint

`func (o *ListVaultsInner) GetTokenBMint() string`

GetTokenBMint returns the TokenBMint field if non-nil, zero value otherwise.

### GetTokenBMintOk

`func (o *ListVaultsInner) GetTokenBMintOk() (*string, bool)`

GetTokenBMintOk returns a tuple with the TokenBMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBMint

`func (o *ListVaultsInner) SetTokenBMint(v string)`

SetTokenBMint sets TokenBMint field to given value.


### GetLastDcaPeriod

`func (o *ListVaultsInner) GetLastDcaPeriod() string`

GetLastDcaPeriod returns the LastDcaPeriod field if non-nil, zero value otherwise.

### GetLastDcaPeriodOk

`func (o *ListVaultsInner) GetLastDcaPeriodOk() (*string, bool)`

GetLastDcaPeriodOk returns a tuple with the LastDcaPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDcaPeriod

`func (o *ListVaultsInner) SetLastDcaPeriod(v string)`

SetLastDcaPeriod sets LastDcaPeriod field to given value.


### GetDripAmount

`func (o *ListVaultsInner) GetDripAmount() string`

GetDripAmount returns the DripAmount field if non-nil, zero value otherwise.

### GetDripAmountOk

`func (o *ListVaultsInner) GetDripAmountOk() (*string, bool)`

GetDripAmountOk returns a tuple with the DripAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDripAmount

`func (o *ListVaultsInner) SetDripAmount(v string)`

SetDripAmount sets DripAmount field to given value.


### GetDcaActivationTimestamp

`func (o *ListVaultsInner) GetDcaActivationTimestamp() string`

GetDcaActivationTimestamp returns the DcaActivationTimestamp field if non-nil, zero value otherwise.

### GetDcaActivationTimestampOk

`func (o *ListVaultsInner) GetDcaActivationTimestampOk() (*string, bool)`

GetDcaActivationTimestampOk returns a tuple with the DcaActivationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcaActivationTimestamp

`func (o *ListVaultsInner) SetDcaActivationTimestamp(v string)`

SetDcaActivationTimestamp sets DcaActivationTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


