# ExpandedAdminVault

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
**MaxSlippageBps** | **int32** |  | 
**MaxPriceDeviationBps** | **int32** |  | 
**OracleConfig** | Pointer to **string** |  | [optional] 
**ProtoConfigValue** | Pointer to [**ProtoConfig**](ProtoConfig.md) |  | [optional] 
**TokenAMintValue** | Pointer to [**Token**](Token.md) |  | [optional] 
**TokenBMintValue** | Pointer to [**Token**](Token.md) |  | [optional] 
**TokenAAccountValue** | Pointer to [**TokenAccount**](TokenAccount.md) |  | [optional] 
**TokenBAccountValue** | Pointer to [**TokenAccount**](TokenAccount.md) |  | [optional] 
**TreasuryTokenBAccountValue** | Pointer to [**TokenAccount**](TokenAccount.md) |  | [optional] 

## Methods

### NewExpandedAdminVault

`func NewExpandedAdminVault(pubkey string, protoConfig string, tokenAAccount string, tokenBAccount string, treasuryTokenBAccount string, tokenAMint string, tokenBMint string, lastDcaPeriod string, dripAmount string, dcaActivationTimestamp string, enabled bool, maxSlippageBps int32, maxPriceDeviationBps int32, ) *ExpandedAdminVault`

NewExpandedAdminVault instantiates a new ExpandedAdminVault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandedAdminVaultWithDefaults

`func NewExpandedAdminVaultWithDefaults() *ExpandedAdminVault`

NewExpandedAdminVaultWithDefaults instantiates a new ExpandedAdminVault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *ExpandedAdminVault) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *ExpandedAdminVault) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *ExpandedAdminVault) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.


### GetProtoConfig

`func (o *ExpandedAdminVault) GetProtoConfig() string`

GetProtoConfig returns the ProtoConfig field if non-nil, zero value otherwise.

### GetProtoConfigOk

`func (o *ExpandedAdminVault) GetProtoConfigOk() (*string, bool)`

GetProtoConfigOk returns a tuple with the ProtoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoConfig

`func (o *ExpandedAdminVault) SetProtoConfig(v string)`

SetProtoConfig sets ProtoConfig field to given value.


### GetTokenAAccount

`func (o *ExpandedAdminVault) GetTokenAAccount() string`

GetTokenAAccount returns the TokenAAccount field if non-nil, zero value otherwise.

### GetTokenAAccountOk

`func (o *ExpandedAdminVault) GetTokenAAccountOk() (*string, bool)`

GetTokenAAccountOk returns a tuple with the TokenAAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAAccount

`func (o *ExpandedAdminVault) SetTokenAAccount(v string)`

SetTokenAAccount sets TokenAAccount field to given value.


### GetTokenBAccount

`func (o *ExpandedAdminVault) GetTokenBAccount() string`

GetTokenBAccount returns the TokenBAccount field if non-nil, zero value otherwise.

### GetTokenBAccountOk

`func (o *ExpandedAdminVault) GetTokenBAccountOk() (*string, bool)`

GetTokenBAccountOk returns a tuple with the TokenBAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBAccount

`func (o *ExpandedAdminVault) SetTokenBAccount(v string)`

SetTokenBAccount sets TokenBAccount field to given value.


### GetTreasuryTokenBAccount

`func (o *ExpandedAdminVault) GetTreasuryTokenBAccount() string`

GetTreasuryTokenBAccount returns the TreasuryTokenBAccount field if non-nil, zero value otherwise.

### GetTreasuryTokenBAccountOk

`func (o *ExpandedAdminVault) GetTreasuryTokenBAccountOk() (*string, bool)`

GetTreasuryTokenBAccountOk returns a tuple with the TreasuryTokenBAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreasuryTokenBAccount

`func (o *ExpandedAdminVault) SetTreasuryTokenBAccount(v string)`

SetTreasuryTokenBAccount sets TreasuryTokenBAccount field to given value.


### GetTokenAMint

`func (o *ExpandedAdminVault) GetTokenAMint() string`

GetTokenAMint returns the TokenAMint field if non-nil, zero value otherwise.

### GetTokenAMintOk

`func (o *ExpandedAdminVault) GetTokenAMintOk() (*string, bool)`

GetTokenAMintOk returns a tuple with the TokenAMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAMint

`func (o *ExpandedAdminVault) SetTokenAMint(v string)`

SetTokenAMint sets TokenAMint field to given value.


### GetTokenBMint

`func (o *ExpandedAdminVault) GetTokenBMint() string`

GetTokenBMint returns the TokenBMint field if non-nil, zero value otherwise.

### GetTokenBMintOk

`func (o *ExpandedAdminVault) GetTokenBMintOk() (*string, bool)`

GetTokenBMintOk returns a tuple with the TokenBMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBMint

`func (o *ExpandedAdminVault) SetTokenBMint(v string)`

SetTokenBMint sets TokenBMint field to given value.


### GetLastDcaPeriod

`func (o *ExpandedAdminVault) GetLastDcaPeriod() string`

GetLastDcaPeriod returns the LastDcaPeriod field if non-nil, zero value otherwise.

### GetLastDcaPeriodOk

`func (o *ExpandedAdminVault) GetLastDcaPeriodOk() (*string, bool)`

GetLastDcaPeriodOk returns a tuple with the LastDcaPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDcaPeriod

`func (o *ExpandedAdminVault) SetLastDcaPeriod(v string)`

SetLastDcaPeriod sets LastDcaPeriod field to given value.


### GetDripAmount

`func (o *ExpandedAdminVault) GetDripAmount() string`

GetDripAmount returns the DripAmount field if non-nil, zero value otherwise.

### GetDripAmountOk

`func (o *ExpandedAdminVault) GetDripAmountOk() (*string, bool)`

GetDripAmountOk returns a tuple with the DripAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDripAmount

`func (o *ExpandedAdminVault) SetDripAmount(v string)`

SetDripAmount sets DripAmount field to given value.


### GetDcaActivationTimestamp

`func (o *ExpandedAdminVault) GetDcaActivationTimestamp() string`

GetDcaActivationTimestamp returns the DcaActivationTimestamp field if non-nil, zero value otherwise.

### GetDcaActivationTimestampOk

`func (o *ExpandedAdminVault) GetDcaActivationTimestampOk() (*string, bool)`

GetDcaActivationTimestampOk returns a tuple with the DcaActivationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcaActivationTimestamp

`func (o *ExpandedAdminVault) SetDcaActivationTimestamp(v string)`

SetDcaActivationTimestamp sets DcaActivationTimestamp field to given value.


### GetEnabled

`func (o *ExpandedAdminVault) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExpandedAdminVault) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExpandedAdminVault) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMaxSlippageBps

`func (o *ExpandedAdminVault) GetMaxSlippageBps() int32`

GetMaxSlippageBps returns the MaxSlippageBps field if non-nil, zero value otherwise.

### GetMaxSlippageBpsOk

`func (o *ExpandedAdminVault) GetMaxSlippageBpsOk() (*int32, bool)`

GetMaxSlippageBpsOk returns a tuple with the MaxSlippageBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSlippageBps

`func (o *ExpandedAdminVault) SetMaxSlippageBps(v int32)`

SetMaxSlippageBps sets MaxSlippageBps field to given value.


### GetMaxPriceDeviationBps

`func (o *ExpandedAdminVault) GetMaxPriceDeviationBps() int32`

GetMaxPriceDeviationBps returns the MaxPriceDeviationBps field if non-nil, zero value otherwise.

### GetMaxPriceDeviationBpsOk

`func (o *ExpandedAdminVault) GetMaxPriceDeviationBpsOk() (*int32, bool)`

GetMaxPriceDeviationBpsOk returns a tuple with the MaxPriceDeviationBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriceDeviationBps

`func (o *ExpandedAdminVault) SetMaxPriceDeviationBps(v int32)`

SetMaxPriceDeviationBps sets MaxPriceDeviationBps field to given value.


### GetOracleConfig

`func (o *ExpandedAdminVault) GetOracleConfig() string`

GetOracleConfig returns the OracleConfig field if non-nil, zero value otherwise.

### GetOracleConfigOk

`func (o *ExpandedAdminVault) GetOracleConfigOk() (*string, bool)`

GetOracleConfigOk returns a tuple with the OracleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleConfig

`func (o *ExpandedAdminVault) SetOracleConfig(v string)`

SetOracleConfig sets OracleConfig field to given value.

### HasOracleConfig

`func (o *ExpandedAdminVault) HasOracleConfig() bool`

HasOracleConfig returns a boolean if a field has been set.

### GetProtoConfigValue

`func (o *ExpandedAdminVault) GetProtoConfigValue() ProtoConfig`

GetProtoConfigValue returns the ProtoConfigValue field if non-nil, zero value otherwise.

### GetProtoConfigValueOk

`func (o *ExpandedAdminVault) GetProtoConfigValueOk() (*ProtoConfig, bool)`

GetProtoConfigValueOk returns a tuple with the ProtoConfigValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoConfigValue

`func (o *ExpandedAdminVault) SetProtoConfigValue(v ProtoConfig)`

SetProtoConfigValue sets ProtoConfigValue field to given value.

### HasProtoConfigValue

`func (o *ExpandedAdminVault) HasProtoConfigValue() bool`

HasProtoConfigValue returns a boolean if a field has been set.

### GetTokenAMintValue

`func (o *ExpandedAdminVault) GetTokenAMintValue() Token`

GetTokenAMintValue returns the TokenAMintValue field if non-nil, zero value otherwise.

### GetTokenAMintValueOk

`func (o *ExpandedAdminVault) GetTokenAMintValueOk() (*Token, bool)`

GetTokenAMintValueOk returns a tuple with the TokenAMintValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAMintValue

`func (o *ExpandedAdminVault) SetTokenAMintValue(v Token)`

SetTokenAMintValue sets TokenAMintValue field to given value.

### HasTokenAMintValue

`func (o *ExpandedAdminVault) HasTokenAMintValue() bool`

HasTokenAMintValue returns a boolean if a field has been set.

### GetTokenBMintValue

`func (o *ExpandedAdminVault) GetTokenBMintValue() Token`

GetTokenBMintValue returns the TokenBMintValue field if non-nil, zero value otherwise.

### GetTokenBMintValueOk

`func (o *ExpandedAdminVault) GetTokenBMintValueOk() (*Token, bool)`

GetTokenBMintValueOk returns a tuple with the TokenBMintValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBMintValue

`func (o *ExpandedAdminVault) SetTokenBMintValue(v Token)`

SetTokenBMintValue sets TokenBMintValue field to given value.

### HasTokenBMintValue

`func (o *ExpandedAdminVault) HasTokenBMintValue() bool`

HasTokenBMintValue returns a boolean if a field has been set.

### GetTokenAAccountValue

`func (o *ExpandedAdminVault) GetTokenAAccountValue() TokenAccount`

GetTokenAAccountValue returns the TokenAAccountValue field if non-nil, zero value otherwise.

### GetTokenAAccountValueOk

`func (o *ExpandedAdminVault) GetTokenAAccountValueOk() (*TokenAccount, bool)`

GetTokenAAccountValueOk returns a tuple with the TokenAAccountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAAccountValue

`func (o *ExpandedAdminVault) SetTokenAAccountValue(v TokenAccount)`

SetTokenAAccountValue sets TokenAAccountValue field to given value.

### HasTokenAAccountValue

`func (o *ExpandedAdminVault) HasTokenAAccountValue() bool`

HasTokenAAccountValue returns a boolean if a field has been set.

### GetTokenBAccountValue

`func (o *ExpandedAdminVault) GetTokenBAccountValue() TokenAccount`

GetTokenBAccountValue returns the TokenBAccountValue field if non-nil, zero value otherwise.

### GetTokenBAccountValueOk

`func (o *ExpandedAdminVault) GetTokenBAccountValueOk() (*TokenAccount, bool)`

GetTokenBAccountValueOk returns a tuple with the TokenBAccountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBAccountValue

`func (o *ExpandedAdminVault) SetTokenBAccountValue(v TokenAccount)`

SetTokenBAccountValue sets TokenBAccountValue field to given value.

### HasTokenBAccountValue

`func (o *ExpandedAdminVault) HasTokenBAccountValue() bool`

HasTokenBAccountValue returns a boolean if a field has been set.

### GetTreasuryTokenBAccountValue

`func (o *ExpandedAdminVault) GetTreasuryTokenBAccountValue() TokenAccount`

GetTreasuryTokenBAccountValue returns the TreasuryTokenBAccountValue field if non-nil, zero value otherwise.

### GetTreasuryTokenBAccountValueOk

`func (o *ExpandedAdminVault) GetTreasuryTokenBAccountValueOk() (*TokenAccount, bool)`

GetTreasuryTokenBAccountValueOk returns a tuple with the TreasuryTokenBAccountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreasuryTokenBAccountValue

`func (o *ExpandedAdminVault) SetTreasuryTokenBAccountValue(v TokenAccount)`

SetTreasuryTokenBAccountValue sets TreasuryTokenBAccountValue field to given value.

### HasTreasuryTokenBAccountValue

`func (o *ExpandedAdminVault) HasTreasuryTokenBAccountValue() bool`

HasTreasuryTokenBAccountValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


