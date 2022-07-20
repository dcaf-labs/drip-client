# Vault

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
**Enabled** | **bool** |  | 

## Methods

### NewVault

`func NewVault(pubkey string, protoConfig string, tokenAAccount string, tokenBAccount string, treasuryTokenBAccount string, tokenAMint string, tokenBMint string, lastDcaPeriod string, dripAmount string, dcaActivationTimestamp string, enabled bool, ) *Vault`

NewVault instantiates a new Vault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultWithDefaults

`func NewVaultWithDefaults() *Vault`

NewVaultWithDefaults instantiates a new Vault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *Vault) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *Vault) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *Vault) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.


### GetProtoConfig

`func (o *Vault) GetProtoConfig() string`

GetProtoConfig returns the ProtoConfig field if non-nil, zero value otherwise.

### GetProtoConfigOk

`func (o *Vault) GetProtoConfigOk() (*string, bool)`

GetProtoConfigOk returns a tuple with the ProtoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoConfig

`func (o *Vault) SetProtoConfig(v string)`

SetProtoConfig sets ProtoConfig field to given value.


### GetTokenAAccount

`func (o *Vault) GetTokenAAccount() string`

GetTokenAAccount returns the TokenAAccount field if non-nil, zero value otherwise.

### GetTokenAAccountOk

`func (o *Vault) GetTokenAAccountOk() (*string, bool)`

GetTokenAAccountOk returns a tuple with the TokenAAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAAccount

`func (o *Vault) SetTokenAAccount(v string)`

SetTokenAAccount sets TokenAAccount field to given value.


### GetTokenBAccount

`func (o *Vault) GetTokenBAccount() string`

GetTokenBAccount returns the TokenBAccount field if non-nil, zero value otherwise.

### GetTokenBAccountOk

`func (o *Vault) GetTokenBAccountOk() (*string, bool)`

GetTokenBAccountOk returns a tuple with the TokenBAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBAccount

`func (o *Vault) SetTokenBAccount(v string)`

SetTokenBAccount sets TokenBAccount field to given value.


### GetTreasuryTokenBAccount

`func (o *Vault) GetTreasuryTokenBAccount() string`

GetTreasuryTokenBAccount returns the TreasuryTokenBAccount field if non-nil, zero value otherwise.

### GetTreasuryTokenBAccountOk

`func (o *Vault) GetTreasuryTokenBAccountOk() (*string, bool)`

GetTreasuryTokenBAccountOk returns a tuple with the TreasuryTokenBAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreasuryTokenBAccount

`func (o *Vault) SetTreasuryTokenBAccount(v string)`

SetTreasuryTokenBAccount sets TreasuryTokenBAccount field to given value.


### GetTokenAMint

`func (o *Vault) GetTokenAMint() string`

GetTokenAMint returns the TokenAMint field if non-nil, zero value otherwise.

### GetTokenAMintOk

`func (o *Vault) GetTokenAMintOk() (*string, bool)`

GetTokenAMintOk returns a tuple with the TokenAMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAMint

`func (o *Vault) SetTokenAMint(v string)`

SetTokenAMint sets TokenAMint field to given value.


### GetTokenBMint

`func (o *Vault) GetTokenBMint() string`

GetTokenBMint returns the TokenBMint field if non-nil, zero value otherwise.

### GetTokenBMintOk

`func (o *Vault) GetTokenBMintOk() (*string, bool)`

GetTokenBMintOk returns a tuple with the TokenBMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBMint

`func (o *Vault) SetTokenBMint(v string)`

SetTokenBMint sets TokenBMint field to given value.


### GetLastDcaPeriod

`func (o *Vault) GetLastDcaPeriod() string`

GetLastDcaPeriod returns the LastDcaPeriod field if non-nil, zero value otherwise.

### GetLastDcaPeriodOk

`func (o *Vault) GetLastDcaPeriodOk() (*string, bool)`

GetLastDcaPeriodOk returns a tuple with the LastDcaPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDcaPeriod

`func (o *Vault) SetLastDcaPeriod(v string)`

SetLastDcaPeriod sets LastDcaPeriod field to given value.


### GetDripAmount

`func (o *Vault) GetDripAmount() string`

GetDripAmount returns the DripAmount field if non-nil, zero value otherwise.

### GetDripAmountOk

`func (o *Vault) GetDripAmountOk() (*string, bool)`

GetDripAmountOk returns a tuple with the DripAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDripAmount

`func (o *Vault) SetDripAmount(v string)`

SetDripAmount sets DripAmount field to given value.


### GetDcaActivationTimestamp

`func (o *Vault) GetDcaActivationTimestamp() string`

GetDcaActivationTimestamp returns the DcaActivationTimestamp field if non-nil, zero value otherwise.

### GetDcaActivationTimestampOk

`func (o *Vault) GetDcaActivationTimestampOk() (*string, bool)`

GetDcaActivationTimestampOk returns a tuple with the DcaActivationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcaActivationTimestamp

`func (o *Vault) SetDcaActivationTimestamp(v string)`

SetDcaActivationTimestamp sets DcaActivationTimestamp field to given value.


### GetEnabled

`func (o *Vault) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Vault) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Vault) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


