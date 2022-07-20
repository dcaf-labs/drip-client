# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MintPost**](DefaultApi.md#MintPost) | **Post** /mint | Mint tokens (DEVNET ONLY)
[**PositionsGet**](DefaultApi.md#PositionsGet) | **Get** /positions | Get User Positions
[**ProtoconfigsGet**](DefaultApi.md#ProtoconfigsGet) | **Get** /protoconfigs | Get Proto Configs
[**RootGet**](DefaultApi.md#RootGet) | **Get** / | Health Check
[**SwaggerJsonGet**](DefaultApi.md#SwaggerJsonGet) | **Get** /swagger.json | Swagger spec
[**SwapConfigsGet**](DefaultApi.md#SwapConfigsGet) | **Get** /swapConfigs | Get Token Swaps Configs
[**SwapsGet**](DefaultApi.md#SwapsGet) | **Get** /swaps | Get Token Swaps
[**TokenpairsGet**](DefaultApi.md#TokenpairsGet) | **Get** /tokenpairs | Get Token Pairs
[**TokensGet**](DefaultApi.md#TokensGet) | **Get** /tokens | Get Tokens
[**VaultperiodsGet**](DefaultApi.md#VaultperiodsGet) | **Get** /vaultperiods | Get Vault Periods
[**VaultsGet**](DefaultApi.md#VaultsGet) | **Get** /vaults | Get Supported Vaults



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
    resp, r, err := apiClient.DefaultApi.MintPost(context.Background()).MintRequest(mintRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MintPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MintPost`: MintResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MintPost`: %v\n", resp)
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


## PositionsGet

> []Position PositionsGet(ctx).Wallet(wallet).Execute()

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
    resp, r, err := apiClient.DefaultApi.PositionsGet(context.Background()).Wallet(wallet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PositionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PositionsGet`: []Position
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PositionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPositionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wallet** | **string** |  | 

### Return type

[**[]Position**](Position.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProtoconfigsGet

> []ProtoConfig ProtoconfigsGet(ctx).TokenA(tokenA).TokenB(tokenB).Execute()

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
    resp, r, err := apiClient.DefaultApi.ProtoconfigsGet(context.Background()).TokenA(tokenA).TokenB(tokenB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ProtoconfigsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProtoconfigsGet`: []ProtoConfig
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ProtoconfigsGet`: %v\n", resp)
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

[**[]ProtoConfig**](ProtoConfig.md)

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
    resp, r, err := apiClient.DefaultApi.RootGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RootGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RootGet`: PingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RootGet`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.SwaggerJsonGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SwaggerJsonGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwaggerJsonGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SwaggerJsonGet`: %v\n", resp)
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

> []SwapConfig SwapConfigsGet(ctx).Vault(vault).Execute()

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
    resp, r, err := apiClient.DefaultApi.SwapConfigsGet(context.Background()).Vault(vault).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SwapConfigsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwapConfigsGet`: []SwapConfig
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SwapConfigsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSwapConfigsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vault** | **string** |  | 

### Return type

[**[]SwapConfig**](SwapConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwapsGet

> []TokenSwap SwapsGet(ctx).TokenPair(tokenPair).Execute()

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
    resp, r, err := apiClient.DefaultApi.SwapsGet(context.Background()).TokenPair(tokenPair).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SwapsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwapsGet`: []TokenSwap
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SwapsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSwapsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenPair** | **string** |  | 

### Return type

[**[]TokenSwap**](TokenSwap.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenpairsGet

> []TokenPair TokenpairsGet(ctx).TokenA(tokenA).TokenB(tokenB).Execute()

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
    resp, r, err := apiClient.DefaultApi.TokenpairsGet(context.Background()).TokenA(tokenA).TokenB(tokenB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TokenpairsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokenpairsGet`: []TokenPair
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TokenpairsGet`: %v\n", resp)
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

[**[]TokenPair**](TokenPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokensGet

> []Token TokensGet(ctx).TokenA(tokenA).TokenB(tokenB).Execute()

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
    resp, r, err := apiClient.DefaultApi.TokensGet(context.Background()).TokenA(tokenA).TokenB(tokenB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TokensGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokensGet`: []Token
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TokensGet`: %v\n", resp)
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

[**[]Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultperiodsGet

> []VaultPeriod VaultperiodsGet(ctx).Vault(vault).VaultPeriod(vaultPeriod).Offset(offset).Limit(limit).Execute()

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
    resp, r, err := apiClient.DefaultApi.VaultperiodsGet(context.Background()).Vault(vault).VaultPeriod(vaultPeriod).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VaultperiodsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VaultperiodsGet`: []VaultPeriod
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VaultperiodsGet`: %v\n", resp)
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

[**[]VaultPeriod**](VaultPeriod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultsGet

> []Vault VaultsGet(ctx).TokenA(tokenA).TokenB(tokenB).ProtoConfig(protoConfig).Execute()

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
    resp, r, err := apiClient.DefaultApi.VaultsGet(context.Background()).TokenA(tokenA).TokenB(tokenB).ProtoConfig(protoConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VaultsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VaultsGet`: []Vault
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VaultsGet`: %v\n", resp)
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

[**[]Vault**](Vault.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

