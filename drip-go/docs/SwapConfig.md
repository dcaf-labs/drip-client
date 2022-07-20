# SwapConfig

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

### NewSwapConfig

`func NewSwapConfig(vault string, vaultProtoConfig string, vaultTokenAAccount string, vaultTokenBAccount string, tokenAMint string, tokenBMint string, swapTokenMint string, swapTokenAAccount string, swapTokenBAccount string, swapFeeAccount string, swapAuthority string, swap string, ) *SwapConfig`

NewSwapConfig instantiates a new SwapConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwapConfigWithDefaults

`func NewSwapConfigWithDefaults() *SwapConfig`

NewSwapConfigWithDefaults instantiates a new SwapConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVault

`func (o *SwapConfig) GetVault() string`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *SwapConfig) GetVaultOk() (*string, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *SwapConfig) SetVault(v string)`

SetVault sets Vault field to given value.


### GetVaultProtoConfig

`func (o *SwapConfig) GetVaultProtoConfig() string`

GetVaultProtoConfig returns the VaultProtoConfig field if non-nil, zero value otherwise.

### GetVaultProtoConfigOk

`func (o *SwapConfig) GetVaultProtoConfigOk() (*string, bool)`

GetVaultProtoConfigOk returns a tuple with the VaultProtoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultProtoConfig

`func (o *SwapConfig) SetVaultProtoConfig(v string)`

SetVaultProtoConfig sets VaultProtoConfig field to given value.


### GetVaultTokenAAccount

`func (o *SwapConfig) GetVaultTokenAAccount() string`

GetVaultTokenAAccount returns the VaultTokenAAccount field if non-nil, zero value otherwise.

### GetVaultTokenAAccountOk

`func (o *SwapConfig) GetVaultTokenAAccountOk() (*string, bool)`

GetVaultTokenAAccountOk returns a tuple with the VaultTokenAAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultTokenAAccount

`func (o *SwapConfig) SetVaultTokenAAccount(v string)`

SetVaultTokenAAccount sets VaultTokenAAccount field to given value.


### GetVaultTokenBAccount

`func (o *SwapConfig) GetVaultTokenBAccount() string`

GetVaultTokenBAccount returns the VaultTokenBAccount field if non-nil, zero value otherwise.

### GetVaultTokenBAccountOk

`func (o *SwapConfig) GetVaultTokenBAccountOk() (*string, bool)`

GetVaultTokenBAccountOk returns a tuple with the VaultTokenBAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultTokenBAccount

`func (o *SwapConfig) SetVaultTokenBAccount(v string)`

SetVaultTokenBAccount sets VaultTokenBAccount field to given value.


### GetTokenAMint

`func (o *SwapConfig) GetTokenAMint() string`

GetTokenAMint returns the TokenAMint field if non-nil, zero value otherwise.

### GetTokenAMintOk

`func (o *SwapConfig) GetTokenAMintOk() (*string, bool)`

GetTokenAMintOk returns a tuple with the TokenAMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAMint

`func (o *SwapConfig) SetTokenAMint(v string)`

SetTokenAMint sets TokenAMint field to given value.


### GetTokenBMint

`func (o *SwapConfig) GetTokenBMint() string`

GetTokenBMint returns the TokenBMint field if non-nil, zero value otherwise.

### GetTokenBMintOk

`func (o *SwapConfig) GetTokenBMintOk() (*string, bool)`

GetTokenBMintOk returns a tuple with the TokenBMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBMint

`func (o *SwapConfig) SetTokenBMint(v string)`

SetTokenBMint sets TokenBMint field to given value.


### GetSwapTokenMint

`func (o *SwapConfig) GetSwapTokenMint() string`

GetSwapTokenMint returns the SwapTokenMint field if non-nil, zero value otherwise.

### GetSwapTokenMintOk

`func (o *SwapConfig) GetSwapTokenMintOk() (*string, bool)`

GetSwapTokenMintOk returns a tuple with the SwapTokenMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapTokenMint

`func (o *SwapConfig) SetSwapTokenMint(v string)`

SetSwapTokenMint sets SwapTokenMint field to given value.


### GetSwapTokenAAccount

`func (o *SwapConfig) GetSwapTokenAAccount() string`

GetSwapTokenAAccount returns the SwapTokenAAccount field if non-nil, zero value otherwise.

### GetSwapTokenAAccountOk

`func (o *SwapConfig) GetSwapTokenAAccountOk() (*string, bool)`

GetSwapTokenAAccountOk returns a tuple with the SwapTokenAAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapTokenAAccount

`func (o *SwapConfig) SetSwapTokenAAccount(v string)`

SetSwapTokenAAccount sets SwapTokenAAccount field to given value.


### GetSwapTokenBAccount

`func (o *SwapConfig) GetSwapTokenBAccount() string`

GetSwapTokenBAccount returns the SwapTokenBAccount field if non-nil, zero value otherwise.

### GetSwapTokenBAccountOk

`func (o *SwapConfig) GetSwapTokenBAccountOk() (*string, bool)`

GetSwapTokenBAccountOk returns a tuple with the SwapTokenBAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapTokenBAccount

`func (o *SwapConfig) SetSwapTokenBAccount(v string)`

SetSwapTokenBAccount sets SwapTokenBAccount field to given value.


### GetSwapFeeAccount

`func (o *SwapConfig) GetSwapFeeAccount() string`

GetSwapFeeAccount returns the SwapFeeAccount field if non-nil, zero value otherwise.

### GetSwapFeeAccountOk

`func (o *SwapConfig) GetSwapFeeAccountOk() (*string, bool)`

GetSwapFeeAccountOk returns a tuple with the SwapFeeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapFeeAccount

`func (o *SwapConfig) SetSwapFeeAccount(v string)`

SetSwapFeeAccount sets SwapFeeAccount field to given value.


### GetSwapAuthority

`func (o *SwapConfig) GetSwapAuthority() string`

GetSwapAuthority returns the SwapAuthority field if non-nil, zero value otherwise.

### GetSwapAuthorityOk

`func (o *SwapConfig) GetSwapAuthorityOk() (*string, bool)`

GetSwapAuthorityOk returns a tuple with the SwapAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapAuthority

`func (o *SwapConfig) SetSwapAuthority(v string)`

SetSwapAuthority sets SwapAuthority field to given value.


### GetSwap

`func (o *SwapConfig) GetSwap() string`

GetSwap returns the Swap field if non-nil, zero value otherwise.

### GetSwapOk

`func (o *SwapConfig) GetSwapOk() (*string, bool)`

GetSwapOk returns a tuple with the Swap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwap

`func (o *SwapConfig) SetSwap(v string)`

SetSwap sets Swap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


