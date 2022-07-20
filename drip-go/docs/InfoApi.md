# \InfoApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PositionsGet**](InfoApi.md#PositionsGet) | **Get** /positions | Get User Positions
[**ProtoconfigsGet**](InfoApi.md#ProtoconfigsGet) | **Get** /protoconfigs | Get Proto Configs
[**RootGet**](InfoApi.md#RootGet) | **Get** / | Health Check
[**SwaggerJsonGet**](InfoApi.md#SwaggerJsonGet) | **Get** /swagger.json | Swagger spec
[**SwapConfigsGet**](InfoApi.md#SwapConfigsGet) | **Get** /swapConfigs | Get Token Swaps Configs
[**SwapsGet**](InfoApi.md#SwapsGet) | **Get** /swaps | Get Token Swaps
[**TokenpairsGet**](InfoApi.md#TokenpairsGet) | **Get** /tokenpairs | Get Token Pairs
[**TokensGet**](InfoApi.md#TokensGet) | **Get** /tokens | Get Tokens
[**VaultperiodsGet**](InfoApi.md#VaultperiodsGet) | **Get** /vaultperiods | Get Vault Periods
[**VaultsGet**](InfoApi.md#VaultsGet) | **Get** /vaults | Get Supported Vaults



## PositionsGet

> []ListPositionsInner PositionsGet(ctx).Wallet(wallet).Execute()

Get User Positions



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
    wallet := "wallet_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfoApi.PositionsGet(context.Background()).Wallet(wallet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.PositionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PositionsGet`: []ListPositionsInner
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.PositionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPositionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wallet** | **string** |  | 

### Return type

[**[]ListPositionsInner**](ListPositionsInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProtoconfigsGet

> []ListProtoConfigsInner ProtoconfigsGet(ctx).TokenA(tokenA).TokenB(tokenB).Execute()

Get Proto Configs



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
    tokenA := "tokenA_example" // string |  (optional)
    tokenB := "tokenB_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfoApi.ProtoconfigsGet(context.Background()).TokenA(tokenA).TokenB(tokenB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.ProtoconfigsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProtoconfigsGet`: []ListProtoConfigsInner
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.ProtoconfigsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProtoconfigsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenA** | **string** |  | 
 **tokenB** | **string** |  | 

### Return type

[**[]ListProtoConfigsInner**](ListProtoConfigsInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RootGet

> PingResponse RootGet(ctx).Execute()

Health Check



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfoApi.RootGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.RootGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RootGet`: PingResponse
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.RootGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRootGetRequest struct via the builder pattern


### Return type

[**PingResponse**](PingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwaggerJsonGet

> map[string]interface{} SwaggerJsonGet(ctx).Execute()

Swagger spec

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfoApi.SwaggerJsonGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.SwaggerJsonGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwaggerJsonGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.SwaggerJsonGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSwaggerJsonGetRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwapConfigsGet

> []ListSwapConfigsInner SwapConfigsGet(ctx).Vault(vault).Execute()

Get Token Swaps Configs



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
    vault := "vault_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfoApi.SwapConfigsGet(context.Background()).Vault(vault).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.SwapConfigsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwapConfigsGet`: []ListSwapConfigsInner
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.SwapConfigsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSwapConfigsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vault** | **string** |  | 

### Return type

[**[]ListSwapConfigsInner**](ListSwapConfigsInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwapsGet

> []ListTokenSwapsInner SwapsGet(ctx).TokenPair(tokenPair).Execute()

Get Token Swaps



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
    tokenPair := "tokenPair_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfoApi.SwapsGet(context.Background()).TokenPair(tokenPair).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.SwapsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwapsGet`: []ListTokenSwapsInner
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.SwapsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSwapsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenPair** | **string** |  | 

### Return type

[**[]ListTokenSwapsInner**](ListTokenSwapsInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenpairsGet

> []ListTokenPairsInner TokenpairsGet(ctx).TokenA(tokenA).TokenB(tokenB).Execute()

Get Token Pairs



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
    tokenA := "tokenA_example" // string |  (optional)
    tokenB := "tokenB_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfoApi.TokenpairsGet(context.Background()).TokenA(tokenA).TokenB(tokenB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.TokenpairsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokenpairsGet`: []ListTokenPairsInner
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.TokenpairsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenpairsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenA** | **string** |  | 
 **tokenB** | **string** |  | 

### Return type

[**[]ListTokenPairsInner**](ListTokenPairsInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokensGet

> []ListTokensInner TokensGet(ctx).TokenA(tokenA).TokenB(tokenB).Execute()

Get Tokens



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
    tokenA := "tokenA_example" // string |  (optional)
    tokenB := "tokenB_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfoApi.TokensGet(context.Background()).TokenA(tokenA).TokenB(tokenB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.TokensGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokensGet`: []ListTokensInner
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.TokensGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokensGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenA** | **string** |  | 
 **tokenB** | **string** |  | 

### Return type

[**[]ListTokensInner**](ListTokensInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultperiodsGet

> []ListVaultPeriodsInner VaultperiodsGet(ctx).Vault(vault).VaultPeriod(vaultPeriod).Offset(offset).Limit(limit).Execute()

Get Vault Periods



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
    vault := "vault_example" // string | 
    vaultPeriod := "vaultPeriod_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfoApi.VaultperiodsGet(context.Background()).Vault(vault).VaultPeriod(vaultPeriod).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.VaultperiodsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VaultperiodsGet`: []ListVaultPeriodsInner
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.VaultperiodsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVaultperiodsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vault** | **string** |  | 
 **vaultPeriod** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**[]ListVaultPeriodsInner**](ListVaultPeriodsInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultsGet

> []ListVaultsInner VaultsGet(ctx).TokenA(tokenA).TokenB(tokenB).ProtoConfig(protoConfig).Execute()

Get Supported Vaults



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
    tokenA := "tokenA_example" // string |  (optional)
    tokenB := "tokenB_example" // string |  (optional)
    protoConfig := "protoConfig_example" // string | Vault proto config public key. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfoApi.VaultsGet(context.Background()).TokenA(tokenA).TokenB(tokenB).ProtoConfig(protoConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.VaultsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VaultsGet`: []ListVaultsInner
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.VaultsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVaultsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenA** | **string** |  | 
 **tokenB** | **string** |  | 
 **protoConfig** | **string** | Vault proto config public key. | 

### Return type

[**[]ListVaultsInner**](ListVaultsInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

