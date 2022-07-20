# \MintApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MintPost**](MintApi.md#MintPost) | **Post** /mint | Mint tokens (DEVNET ONLY)



## MintPost

> MintResponse MintPost(ctx).MintRequest(mintRequest).Execute()

Mint tokens (DEVNET ONLY)



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
    mintRequest := *openapiclient.NewMintRequest("Mint_example", "Wallet_example", "Amount_example") // MintRequest | Pet to add to the store

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MintApi.MintPost(context.Background()).MintRequest(mintRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintApi.MintPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MintPost`: MintResponse
    fmt.Fprintf(os.Stdout, "Response from `MintApi.MintPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMintPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintRequest** | [**MintRequest**](MintRequest.md) | Pet to add to the store | 

### Return type

[**MintResponse**](MintResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

