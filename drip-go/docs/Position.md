# Position

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | **string** |  | 
**Vault** | **string** |  | 
**Authority** | **string** |  | 
**DepositedTokenAAmount** | **string** |  | 
**WithdrawnTokenBAmount** | **string** |  | 
**DepositTimestamp** | **string** |  | 
**DcaPeriodIdBeforeDeposit** | **string** |  | 
**NumberOfSwaps** | **string** |  | 
**PeriodicDripAmount** | **string** |  | 
**IsClosed** | **bool** |  | 

## Methods

### NewPosition

`func NewPosition(pubkey string, vault string, authority string, depositedTokenAAmount string, withdrawnTokenBAmount string, depositTimestamp string, dcaPeriodIdBeforeDeposit string, numberOfSwaps string, periodicDripAmount string, isClosed bool, ) *Position`

NewPosition instantiates a new Position object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionWithDefaults

`func NewPositionWithDefaults() *Position`

NewPositionWithDefaults instantiates a new Position object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *Position) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *Position) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *Position) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.


### GetVault

`func (o *Position) GetVault() string`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *Position) GetVaultOk() (*string, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *Position) SetVault(v string)`

SetVault sets Vault field to given value.


### GetAuthority

`func (o *Position) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *Position) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *Position) SetAuthority(v string)`

SetAuthority sets Authority field to given value.


### GetDepositedTokenAAmount

`func (o *Position) GetDepositedTokenAAmount() string`

GetDepositedTokenAAmount returns the DepositedTokenAAmount field if non-nil, zero value otherwise.

### GetDepositedTokenAAmountOk

`func (o *Position) GetDepositedTokenAAmountOk() (*string, bool)`

GetDepositedTokenAAmountOk returns a tuple with the DepositedTokenAAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositedTokenAAmount

`func (o *Position) SetDepositedTokenAAmount(v string)`

SetDepositedTokenAAmount sets DepositedTokenAAmount field to given value.


### GetWithdrawnTokenBAmount

`func (o *Position) GetWithdrawnTokenBAmount() string`

GetWithdrawnTokenBAmount returns the WithdrawnTokenBAmount field if non-nil, zero value otherwise.

### GetWithdrawnTokenBAmountOk

`func (o *Position) GetWithdrawnTokenBAmountOk() (*string, bool)`

GetWithdrawnTokenBAmountOk returns a tuple with the WithdrawnTokenBAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawnTokenBAmount

`func (o *Position) SetWithdrawnTokenBAmount(v string)`

SetWithdrawnTokenBAmount sets WithdrawnTokenBAmount field to given value.


### GetDepositTimestamp

`func (o *Position) GetDepositTimestamp() string`

GetDepositTimestamp returns the DepositTimestamp field if non-nil, zero value otherwise.

### GetDepositTimestampOk

`func (o *Position) GetDepositTimestampOk() (*string, bool)`

GetDepositTimestampOk returns a tuple with the DepositTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositTimestamp

`func (o *Position) SetDepositTimestamp(v string)`

SetDepositTimestamp sets DepositTimestamp field to given value.


### GetDcaPeriodIdBeforeDeposit

`func (o *Position) GetDcaPeriodIdBeforeDeposit() string`

GetDcaPeriodIdBeforeDeposit returns the DcaPeriodIdBeforeDeposit field if non-nil, zero value otherwise.

### GetDcaPeriodIdBeforeDepositOk

`func (o *Position) GetDcaPeriodIdBeforeDepositOk() (*string, bool)`

GetDcaPeriodIdBeforeDepositOk returns a tuple with the DcaPeriodIdBeforeDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcaPeriodIdBeforeDeposit

`func (o *Position) SetDcaPeriodIdBeforeDeposit(v string)`

SetDcaPeriodIdBeforeDeposit sets DcaPeriodIdBeforeDeposit field to given value.


### GetNumberOfSwaps

`func (o *Position) GetNumberOfSwaps() string`

GetNumberOfSwaps returns the NumberOfSwaps field if non-nil, zero value otherwise.

### GetNumberOfSwapsOk

`func (o *Position) GetNumberOfSwapsOk() (*string, bool)`

GetNumberOfSwapsOk returns a tuple with the NumberOfSwaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSwaps

`func (o *Position) SetNumberOfSwaps(v string)`

SetNumberOfSwaps sets NumberOfSwaps field to given value.


### GetPeriodicDripAmount

`func (o *Position) GetPeriodicDripAmount() string`

GetPeriodicDripAmount returns the PeriodicDripAmount field if non-nil, zero value otherwise.

### GetPeriodicDripAmountOk

`func (o *Position) GetPeriodicDripAmountOk() (*string, bool)`

GetPeriodicDripAmountOk returns a tuple with the PeriodicDripAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicDripAmount

`func (o *Position) SetPeriodicDripAmount(v string)`

SetPeriodicDripAmount sets PeriodicDripAmount field to given value.


### GetIsClosed

`func (o *Position) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *Position) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *Position) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


