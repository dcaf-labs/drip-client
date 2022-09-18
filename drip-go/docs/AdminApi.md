# \AdminApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AdminPositionsGet**](AdminApi.md#V1AdminPositionsGet) | **Get** /v1/admin/positions | Get All Positions
[**V1AdminSummaryActivewalletsGet**](AdminApi.md#V1AdminSummaryActivewalletsGet) | **Get** /v1/admin/summary/activewallets | Get All Active Wallet Addresses
[**V1AdminVaultPubkeyPathEnablePut**](AdminApi.md#V1AdminVaultPubkeyPathEnablePut) | **Put** /v1/admin/vault/{pubkeyPath}/enable | Toggle the &#39;enabled&#39; flag on a vault
[**V1AdminVaultsGet**](AdminApi.md#V1AdminVaultsGet) | **Get** /v1/admin/vaults | Get All Vaults



## V1AdminPositionsGet

> []ExpandedAdminPosition V1AdminPositionsGet(ctx).TokenId(tokenId).Expand(expand).Enabled(enabled).IsClosed(isClosed).Offset(offset).Limit(limit).Execute()

Get All Positions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenId := "tokenId_example" // string | 
    expand := []string{"Expand_example"} // []string |  (optional)
    enabled := true // bool |  (optional)
    isClosed := true // bool |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.V1AdminPositionsGet(context.Background()).TokenId(tokenId).Expand(expand).Enabled(enabled).IsClosed(isClosed).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.V1AdminPositionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AdminPositionsGet`: []ExpandedAdminPosition
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.V1AdminPositionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AdminPositionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenId** | **string** |  | 
 **expand** | **[]string** |  | 
 **enabled** | **bool** |  | 
 **isClosed** | **bool** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**[]ExpandedAdminPosition**](ExpandedAdminPosition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AdminSummaryActivewalletsGet

> []ActiveWallet V1AdminSummaryActivewalletsGet(ctx).TokenId(tokenId).Vault(vault).IsClosed(isClosed).Owner(owner).Execute()

Get All Active Wallet Addresses



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenId := "tokenId_example" // string | 
    vault := "vault_example" // string |  (optional)
    isClosed := true // bool |  (optional)
    owner := "owner_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.V1AdminSummaryActivewalletsGet(context.Background()).TokenId(tokenId).Vault(vault).IsClosed(isClosed).Owner(owner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.V1AdminSummaryActivewalletsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AdminSummaryActivewalletsGet`: []ActiveWallet
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.V1AdminSummaryActivewalletsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AdminSummaryActivewalletsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenId** | **string** |  | 
 **vault** | **string** |  | 
 **isClosed** | **bool** |  | 
 **owner** | **string** |  | 

### Return type

[**[]ActiveWallet**](ActiveWallet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AdminVaultPubkeyPathEnablePut

> Vault V1AdminVaultPubkeyPathEnablePut(ctx, pubkeyPath).TokenId(tokenId).Execute()

Toggle the 'enabled' flag on a vault



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pubkeyPath := "pubkeyPath_example" // string | 
    tokenId := "tokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.V1AdminVaultPubkeyPathEnablePut(context.Background(), pubkeyPath).TokenId(tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.V1AdminVaultPubkeyPathEnablePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AdminVaultPubkeyPathEnablePut`: Vault
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.V1AdminVaultPubkeyPathEnablePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkeyPath** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AdminVaultPubkeyPathEnablePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenId** | **string** |  | 

### Return type

[**Vault**](Vault.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AdminVaultsGet

> []ExpandedAdminVault V1AdminVaultsGet(ctx).TokenId(tokenId).Expand(expand).Vault(vault).TokenA(tokenA).TokenB(tokenB).Enabled(enabled).Offset(offset).Limit(limit).Execute()

Get All Vaults



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenId := "tokenId_example" // string | 
    expand := []string{"Expand_example"} // []string |  (optional)
    vault := "vault_example" // string |  (optional)
    tokenA := "tokenA_example" // string |  (optional)
    tokenB := "tokenB_example" // string |  (optional)
    enabled := true // bool |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.V1AdminVaultsGet(context.Background()).TokenId(tokenId).Expand(expand).Vault(vault).TokenA(tokenA).TokenB(tokenB).Enabled(enabled).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.V1AdminVaultsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AdminVaultsGet`: []ExpandedAdminVault
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.V1AdminVaultsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AdminVaultsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenId** | **string** |  | 
 **expand** | **[]string** |  | 
 **vault** | **string** |  | 
 **tokenA** | **string** |  | 
 **tokenB** | **string** |  | 
 **enabled** | **bool** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**[]ExpandedAdminVault**](ExpandedAdminVault.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

