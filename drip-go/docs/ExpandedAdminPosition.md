# ExpandedAdminPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | **string** |  | 
**Vault** | [**Vault**](Vault.md) |  | 
**Authority** | **string** |  | 
**DepositedTokenAAmount** | **string** |  | 
**WithdrawnTokenBAmount** | **string** |  | 
**DepositTimestamp** | **string** |  | 
**DcaPeriodIdBeforeDeposit** | **string** |  | 
**NumberOfSwaps** | **string** |  | 
**PeriodicDripAmount** | **string** |  | 
**IsClosed** | **bool** |  | 
**ProtoConfig** | Pointer to [**ProtoConfig**](ProtoConfig.md) |  | [optional] 
**TokenA** | Pointer to [**Token**](Token.md) |  | [optional] 
**TokenB** | Pointer to [**Token**](Token.md) |  | [optional] 

## Methods

### NewExpandedAdminPosition

`func NewExpandedAdminPosition(pubkey string, vault Vault, authority string, depositedTokenAAmount string, withdrawnTokenBAmount string, depositTimestamp string, dcaPeriodIdBeforeDeposit string, numberOfSwaps string, periodicDripAmount string, isClosed bool, ) *ExpandedAdminPosition`

NewExpandedAdminPosition instantiates a new ExpandedAdminPosition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandedAdminPositionWithDefaults

`func NewExpandedAdminPositionWithDefaults() *ExpandedAdminPosition`

NewExpandedAdminPositionWithDefaults instantiates a new ExpandedAdminPosition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *ExpandedAdminPosition) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *ExpandedAdminPosition) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *ExpandedAdminPosition) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.


### GetVault

`func (o *ExpandedAdminPosition) GetVault() Vault`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *ExpandedAdminPosition) GetVaultOk() (*Vault, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *ExpandedAdminPosition) SetVault(v Vault)`

SetVault sets Vault field to given value.


### GetAuthority

`func (o *ExpandedAdminPosition) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *ExpandedAdminPosition) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *ExpandedAdminPosition) SetAuthority(v string)`

SetAuthority sets Authority field to given value.


### GetDepositedTokenAAmount

`func (o *ExpandedAdminPosition) GetDepositedTokenAAmount() string`

GetDepositedTokenAAmount returns the DepositedTokenAAmount field if non-nil, zero value otherwise.

### GetDepositedTokenAAmountOk

`func (o *ExpandedAdminPosition) GetDepositedTokenAAmountOk() (*string, bool)`

GetDepositedTokenAAmountOk returns a tuple with the DepositedTokenAAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositedTokenAAmount

`func (o *ExpandedAdminPosition) SetDepositedTokenAAmount(v string)`

SetDepositedTokenAAmount sets DepositedTokenAAmount field to given value.


### GetWithdrawnTokenBAmount

`func (o *ExpandedAdminPosition) GetWithdrawnTokenBAmount() string`

GetWithdrawnTokenBAmount returns the WithdrawnTokenBAmount field if non-nil, zero value otherwise.

### GetWithdrawnTokenBAmountOk

`func (o *ExpandedAdminPosition) GetWithdrawnTokenBAmountOk() (*string, bool)`

GetWithdrawnTokenBAmountOk returns a tuple with the WithdrawnTokenBAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawnTokenBAmount

`func (o *ExpandedAdminPosition) SetWithdrawnTokenBAmount(v string)`

SetWithdrawnTokenBAmount sets WithdrawnTokenBAmount field to given value.


### GetDepositTimestamp

`func (o *ExpandedAdminPosition) GetDepositTimestamp() string`

GetDepositTimestamp returns the DepositTimestamp field if non-nil, zero value otherwise.

### GetDepositTimestampOk

`func (o *ExpandedAdminPosition) GetDepositTimestampOk() (*string, bool)`

GetDepositTimestampOk returns a tuple with the DepositTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositTimestamp

`func (o *ExpandedAdminPosition) SetDepositTimestamp(v string)`

SetDepositTimestamp sets DepositTimestamp field to given value.


### GetDcaPeriodIdBeforeDeposit

`func (o *ExpandedAdminPosition) GetDcaPeriodIdBeforeDeposit() string`

GetDcaPeriodIdBeforeDeposit returns the DcaPeriodIdBeforeDeposit field if non-nil, zero value otherwise.

### GetDcaPeriodIdBeforeDepositOk

`func (o *ExpandedAdminPosition) GetDcaPeriodIdBeforeDepositOk() (*string, bool)`

GetDcaPeriodIdBeforeDepositOk returns a tuple with the DcaPeriodIdBeforeDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcaPeriodIdBeforeDeposit

`func (o *ExpandedAdminPosition) SetDcaPeriodIdBeforeDeposit(v string)`

SetDcaPeriodIdBeforeDeposit sets DcaPeriodIdBeforeDeposit field to given value.


### GetNumberOfSwaps

`func (o *ExpandedAdminPosition) GetNumberOfSwaps() string`

GetNumberOfSwaps returns the NumberOfSwaps field if non-nil, zero value otherwise.

### GetNumberOfSwapsOk

`func (o *ExpandedAdminPosition) GetNumberOfSwapsOk() (*string, bool)`

GetNumberOfSwapsOk returns a tuple with the NumberOfSwaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSwaps

`func (o *ExpandedAdminPosition) SetNumberOfSwaps(v string)`

SetNumberOfSwaps sets NumberOfSwaps field to given value.


### GetPeriodicDripAmount

`func (o *ExpandedAdminPosition) GetPeriodicDripAmount() string`

GetPeriodicDripAmount returns the PeriodicDripAmount field if non-nil, zero value otherwise.

### GetPeriodicDripAmountOk

`func (o *ExpandedAdminPosition) GetPeriodicDripAmountOk() (*string, bool)`

GetPeriodicDripAmountOk returns a tuple with the PeriodicDripAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicDripAmount

`func (o *ExpandedAdminPosition) SetPeriodicDripAmount(v string)`

SetPeriodicDripAmount sets PeriodicDripAmount field to given value.


### GetIsClosed

`func (o *ExpandedAdminPosition) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *ExpandedAdminPosition) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *ExpandedAdminPosition) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.


### GetProtoConfig

`func (o *ExpandedAdminPosition) GetProtoConfig() ProtoConfig`

GetProtoConfig returns the ProtoConfig field if non-nil, zero value otherwise.

### GetProtoConfigOk

`func (o *ExpandedAdminPosition) GetProtoConfigOk() (*ProtoConfig, bool)`

GetProtoConfigOk returns a tuple with the ProtoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoConfig

`func (o *ExpandedAdminPosition) SetProtoConfig(v ProtoConfig)`

SetProtoConfig sets ProtoConfig field to given value.

### HasProtoConfig

`func (o *ExpandedAdminPosition) HasProtoConfig() bool`

HasProtoConfig returns a boolean if a field has been set.

### GetTokenA

`func (o *ExpandedAdminPosition) GetTokenA() Token`

GetTokenA returns the TokenA field if non-nil, zero value otherwise.

### GetTokenAOk

`func (o *ExpandedAdminPosition) GetTokenAOk() (*Token, bool)`

GetTokenAOk returns a tuple with the TokenA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenA

`func (o *ExpandedAdminPosition) SetTokenA(v Token)`

SetTokenA sets TokenA field to given value.

### HasTokenA

`func (o *ExpandedAdminPosition) HasTokenA() bool`

HasTokenA returns a boolean if a field has been set.

### GetTokenB

`func (o *ExpandedAdminPosition) GetTokenB() Token`

GetTokenB returns the TokenB field if non-nil, zero value otherwise.

### GetTokenBOk

`func (o *ExpandedAdminPosition) GetTokenBOk() (*Token, bool)`

GetTokenBOk returns a tuple with the TokenB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenB

`func (o *ExpandedAdminPosition) SetTokenB(v Token)`

SetTokenB sets TokenB field to given value.

### HasTokenB

`func (o *ExpandedAdminPosition) HasTokenB() bool`

HasTokenB returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


