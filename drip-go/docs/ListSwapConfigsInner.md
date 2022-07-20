# ListSwapConfigsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vault** | **string** |  | 
**VaultProtoConfig** | **string** |  | 
**VaultTokenAAccount** | **string** |  | 
**VaultTokenBAccount** | **string** |  | 
**TokenAMint** | **string** |  | 
**TokenBMint** | **string** |  | 
**SwapTokenMint** | **string** |  | 
**SwapTokenAAccount** | **string** |  | 
**SwapTokenBAccount** | **string** |  | 
**SwapFeeAccount** | **string** |  | 
**SwapAuthority** | **string** |  | 
**Swap** | **string** |  | 

## Methods

### NewListSwapConfigsInner

`func NewListSwapConfigsInner(vault string, vaultProtoConfig string, vaultTokenAAccount string, vaultTokenBAccount string, tokenAMint string, tokenBMint string, swapTokenMint string, swapTokenAAccount string, swapTokenBAccount string, swapFeeAccount string, swapAuthority string, swap string, ) *ListSwapConfigsInner`

NewListSwapConfigsInner instantiates a new ListSwapConfigsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSwapConfigsInnerWithDefaults

`func NewListSwapConfigsInnerWithDefaults() *ListSwapConfigsInner`

NewListSwapConfigsInnerWithDefaults instantiates a new ListSwapConfigsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVault

`func (o *ListSwapConfigsInner) GetVault() string`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *ListSwapConfigsInner) GetVaultOk() (*string, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *ListSwapConfigsInner) SetVault(v string)`

SetVault sets Vault field to given value.


### GetVaultProtoConfig

`func (o *ListSwapConfigsInner) GetVaultProtoConfig() string`

GetVaultProtoConfig returns the VaultProtoConfig field if non-nil, zero value otherwise.

### GetVaultProtoConfigOk

`func (o *ListSwapConfigsInner) GetVaultProtoConfigOk() (*string, bool)`

GetVaultProtoConfigOk returns a tuple with the VaultProtoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultProtoConfig

`func (o *ListSwapConfigsInner) SetVaultProtoConfig(v string)`

SetVaultProtoConfig sets VaultProtoConfig field to given value.


### GetVaultTokenAAccount

`func (o *ListSwapConfigsInner) GetVaultTokenAAccount() string`

GetVaultTokenAAccount returns the VaultTokenAAccount field if non-nil, zero value otherwise.

### GetVaultTokenAAccountOk

`func (o *ListSwapConfigsInner) GetVaultTokenAAccountOk() (*string, bool)`

GetVaultTokenAAccountOk returns a tuple with the VaultTokenAAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultTokenAAccount

`func (o *ListSwapConfigsInner) SetVaultTokenAAccount(v string)`

SetVaultTokenAAccount sets VaultTokenAAccount field to given value.


### GetVaultTokenBAccount

`func (o *ListSwapConfigsInner) GetVaultTokenBAccount() string`

GetVaultTokenBAccount returns the VaultTokenBAccount field if non-nil, zero value otherwise.

### GetVaultTokenBAccountOk

`func (o *ListSwapConfigsInner) GetVaultTokenBAccountOk() (*string, bool)`

GetVaultTokenBAccountOk returns a tuple with the VaultTokenBAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultTokenBAccount

`func (o *ListSwapConfigsInner) SetVaultTokenBAccount(v string)`

SetVaultTokenBAccount sets VaultTokenBAccount field to given value.


### GetTokenAMint

`func (o *ListSwapConfigsInner) GetTokenAMint() string`

GetTokenAMint returns the TokenAMint field if non-nil, zero value otherwise.

### GetTokenAMintOk

`func (o *ListSwapConfigsInner) GetTokenAMintOk() (*string, bool)`

GetTokenAMintOk returns a tuple with the TokenAMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAMint

`func (o *ListSwapConfigsInner) SetTokenAMint(v string)`

SetTokenAMint sets TokenAMint field to given value.


### GetTokenBMint

`func (o *ListSwapConfigsInner) GetTokenBMint() string`

GetTokenBMint returns the TokenBMint field if non-nil, zero value otherwise.

### GetTokenBMintOk

`func (o *ListSwapConfigsInner) GetTokenBMintOk() (*string, bool)`

GetTokenBMintOk returns a tuple with the TokenBMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBMint

`func (o *ListSwapConfigsInner) SetTokenBMint(v string)`

SetTokenBMint sets TokenBMint field to given value.


### GetSwapTokenMint

`func (o *ListSwapConfigsInner) GetSwapTokenMint() string`

GetSwapTokenMint returns the SwapTokenMint field if non-nil, zero value otherwise.

### GetSwapTokenMintOk

`func (o *ListSwapConfigsInner) GetSwapTokenMintOk() (*string, bool)`

GetSwapTokenMintOk returns a tuple with the SwapTokenMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapTokenMint

`func (o *ListSwapConfigsInner) SetSwapTokenMint(v string)`

SetSwapTokenMint sets SwapTokenMint field to given value.


### GetSwapTokenAAccount

`func (o *ListSwapConfigsInner) GetSwapTokenAAccount() string`

GetSwapTokenAAccount returns the SwapTokenAAccount field if non-nil, zero value otherwise.

### GetSwapTokenAAccountOk

`func (o *ListSwapConfigsInner) GetSwapTokenAAccountOk() (*string, bool)`

GetSwapTokenAAccountOk returns a tuple with the SwapTokenAAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapTokenAAccount

`func (o *ListSwapConfigsInner) SetSwapTokenAAccount(v string)`

SetSwapTokenAAccount sets SwapTokenAAccount field to given value.


### GetSwapTokenBAccount

`func (o *ListSwapConfigsInner) GetSwapTokenBAccount() string`

GetSwapTokenBAccount returns the SwapTokenBAccount field if non-nil, zero value otherwise.

### GetSwapTokenBAccountOk

`func (o *ListSwapConfigsInner) GetSwapTokenBAccountOk() (*string, bool)`

GetSwapTokenBAccountOk returns a tuple with the SwapTokenBAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapTokenBAccount

`func (o *ListSwapConfigsInner) SetSwapTokenBAccount(v string)`

SetSwapTokenBAccount sets SwapTokenBAccount field to given value.


### GetSwapFeeAccount

`func (o *ListSwapConfigsInner) GetSwapFeeAccount() string`

GetSwapFeeAccount returns the SwapFeeAccount field if non-nil, zero value otherwise.

### GetSwapFeeAccountOk

`func (o *ListSwapConfigsInner) GetSwapFeeAccountOk() (*string, bool)`

GetSwapFeeAccountOk returns a tuple with the SwapFeeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapFeeAccount

`func (o *ListSwapConfigsInner) SetSwapFeeAccount(v string)`

SetSwapFeeAccount sets SwapFeeAccount field to given value.


### GetSwapAuthority

`func (o *ListSwapConfigsInner) GetSwapAuthority() string`

GetSwapAuthority returns the SwapAuthority field if non-nil, zero value otherwise.

### GetSwapAuthorityOk

`func (o *ListSwapConfigsInner) GetSwapAuthorityOk() (*string, bool)`

GetSwapAuthorityOk returns a tuple with the SwapAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapAuthority

`func (o *ListSwapConfigsInner) SetSwapAuthority(v string)`

SetSwapAuthority sets SwapAuthority field to given value.


### GetSwap

`func (o *ListSwapConfigsInner) GetSwap() string`

GetSwap returns the Swap field if non-nil, zero value otherwise.

### GetSwapOk

`func (o *ListSwapConfigsInner) GetSwapOk() (*string, bool)`

GetSwapOk returns a tuple with the Swap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwap

`func (o *ListSwapConfigsInner) SetSwap(v string)`

SetSwap sets Swap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


