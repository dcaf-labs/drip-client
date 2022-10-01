# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1OrcawhirlpoolQuote**](DefaultApi.md#V1OrcawhirlpoolQuote) | **Post** /v1/orcawhirlpool/quote | Get whirlpool swap quote.



## V1OrcawhirlpoolQuote

> V1OrcawhirlpoolQuote200Response V1OrcawhirlpoolQuote(ctx).V1OrcawhirlpoolQuoteRequest(v1OrcawhirlpoolQuoteRequest).Execute()

Get whirlpool swap quote.

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
    v1OrcawhirlpoolQuoteRequest := *openapiclient.NewV1OrcawhirlpoolQuoteRequest("ConnectionUrl_example", "Whirlpool_example", "InputTokenMint_example", "InputAmount_example") // V1OrcawhirlpoolQuoteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.V1OrcawhirlpoolQuote(context.Background()).V1OrcawhirlpoolQuoteRequest(v1OrcawhirlpoolQuoteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1OrcawhirlpoolQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrcawhirlpoolQuote`: V1OrcawhirlpoolQuote200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1OrcawhirlpoolQuote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OrcawhirlpoolQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1OrcawhirlpoolQuoteRequest** | [**V1OrcawhirlpoolQuoteRequest**](V1OrcawhirlpoolQuoteRequest.md) |  | 

### Return type

[**V1OrcawhirlpoolQuote200Response**](V1OrcawhirlpoolQuote200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

