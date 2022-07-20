# \AdminApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminVaultPubkeyPathEnablePut**](AdminApi.md#AdminVaultPubkeyPathEnablePut) | **Put** /admin/vault/{pubkeyPath}/enable | Enable a vault
[**AdminVaultsGet**](AdminApi.md#AdminVaultsGet) | **Get** /admin/vaults | Get All Vaults



## AdminVaultPubkeyPathEnablePut

> Vault AdminVaultPubkeyPathEnablePut(ctx, pubkeyPath).TokenId(tokenId).Execute()

Enable a vault



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
    resp, r, err := apiClient.AdminApi.AdminVaultPubkeyPathEnablePut(context.Background(), pubkeyPath).TokenId(tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminVaultPubkeyPathEnablePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminVaultPubkeyPathEnablePut`: Vault
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminVaultPubkeyPathEnablePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkeyPath** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminVaultPubkeyPathEnablePutRequest struct via the builder pattern


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


## AdminVaultsGet

> []ExpandedAdminVault AdminVaultsGet(ctx).TokenId(tokenId).Expand(expand).Enabled(enabled).Offset(offset).Limit(limit).Execute()

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
    enabled := true // bool |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminVaultsGet(context.Background()).TokenId(tokenId).Expand(expand).Enabled(enabled).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminVaultsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminVaultsGet`: []ExpandedAdminVault
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminVaultsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminVaultsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenId** | **string** |  | 
 **expand** | **[]string** |  | 
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

