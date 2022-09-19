# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MintPost**](DefaultApi.md#MintPost) | **Post** /mint | Mint tokens (DEVNET ONLY)
[**RootGet**](DefaultApi.md#RootGet) | **Get** / | Health Check
[**SwaggerJsonGet**](DefaultApi.md#SwaggerJsonGet) | **Get** /swagger.json | Swagger spec
[**V1DripOrcawhirlpoolconfigsGet**](DefaultApi.md#V1DripOrcawhirlpoolconfigsGet) | **Get** /v1/drip/orcawhirlpoolconfigs | Get Orca Whirlpool Swap Configs
[**V1DripPositionPubkeyPathMetadataGet**](DefaultApi.md#V1DripPositionPubkeyPathMetadataGet) | **Get** /v1/drip/position/{pubkeyPath}/metadata | Get Drip Position Metadata
[**V1DripPubkeyPathTokenmetadataGet**](DefaultApi.md#V1DripPubkeyPathTokenmetadataGet) | **Get** /v1/drip/{pubkeyPath}/tokenmetadata | Get TokenMetadata for Devnet Mints.
[**V1DripSpltokenswapconfigsGet**](DefaultApi.md#V1DripSpltokenswapconfigsGet) | **Get** /v1/drip/spltokenswapconfigs | Get Token Swaps Configs
[**V1PositionsGet**](DefaultApi.md#V1PositionsGet) | **Get** /v1/positions | Get User Positions
[**V1ProtoconfigsGet**](DefaultApi.md#V1ProtoconfigsGet) | **Get** /v1/protoconfigs | Get Proto Configs
[**V1VaultTokenpairsGet**](DefaultApi.md#V1VaultTokenpairsGet) | **Get** /v1/vault/tokenpairs | Get all Supported Token Pairs
[**V1VaultTokensGet**](DefaultApi.md#V1VaultTokensGet) | **Get** /v1/vault/tokens | Get all Supported Tokens
[**V1VaultperiodsGet**](DefaultApi.md#V1VaultperiodsGet) | **Get** /v1/vaultperiods | Get Vault Periods
[**V1VaultsGet**](DefaultApi.md#V1VaultsGet) | **Get** /v1/vaults | Get Supported Vaults



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


## V1DripOrcawhirlpoolconfigsGet

> []OrcaWhirlpoolConfig V1DripOrcawhirlpoolconfigsGet(ctx).Vault(vault).Execute()

Get Orca Whirlpool Swap Configs



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
    resp, r, err := apiClient.DefaultApi.V1DripOrcawhirlpoolconfigsGet(context.Background()).Vault(vault).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1DripOrcawhirlpoolconfigsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DripOrcawhirlpoolconfigsGet`: []OrcaWhirlpoolConfig
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1DripOrcawhirlpoolconfigsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DripOrcawhirlpoolconfigsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vault** | **string** |  | 

### Return type

[**[]OrcaWhirlpoolConfig**](OrcaWhirlpoolConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DripPositionPubkeyPathMetadataGet

> TokenMetadata V1DripPositionPubkeyPathMetadataGet(ctx, pubkeyPath).Execute()

Get Drip Position Metadata

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.V1DripPositionPubkeyPathMetadataGet(context.Background(), pubkeyPath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1DripPositionPubkeyPathMetadataGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DripPositionPubkeyPathMetadataGet`: TokenMetadata
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1DripPositionPubkeyPathMetadataGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkeyPath** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DripPositionPubkeyPathMetadataGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenMetadata**](TokenMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DripPubkeyPathTokenmetadataGet

> TokenMetadata V1DripPubkeyPathTokenmetadataGet(ctx, pubkeyPath).Execute()

Get TokenMetadata for Devnet Mints.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.V1DripPubkeyPathTokenmetadataGet(context.Background(), pubkeyPath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1DripPubkeyPathTokenmetadataGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DripPubkeyPathTokenmetadataGet`: TokenMetadata
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1DripPubkeyPathTokenmetadataGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkeyPath** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DripPubkeyPathTokenmetadataGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenMetadata**](TokenMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DripSpltokenswapconfigsGet

> []SplTokenSwapConfig V1DripSpltokenswapconfigsGet(ctx).Vault(vault).Execute()

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
    resp, r, err := apiClient.DefaultApi.V1DripSpltokenswapconfigsGet(context.Background()).Vault(vault).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1DripSpltokenswapconfigsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DripSpltokenswapconfigsGet`: []SplTokenSwapConfig
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1DripSpltokenswapconfigsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DripSpltokenswapconfigsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vault** | **string** |  | 

### Return type

[**[]SplTokenSwapConfig**](SplTokenSwapConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PositionsGet

> []Position V1PositionsGet(ctx).Wallet(wallet).IsClosed(isClosed).Offset(offset).Limit(limit).Execute()

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
    isClosed := true // bool |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.V1PositionsGet(context.Background()).Wallet(wallet).IsClosed(isClosed).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1PositionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PositionsGet`: []Position
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1PositionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PositionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wallet** | **string** |  | 
 **isClosed** | **bool** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

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


## V1ProtoconfigsGet

> []ProtoConfig V1ProtoconfigsGet(ctx).TokenA(tokenA).TokenB(tokenB).Execute()

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
    resp, r, err := apiClient.DefaultApi.V1ProtoconfigsGet(context.Background()).TokenA(tokenA).TokenB(tokenB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1ProtoconfigsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProtoconfigsGet`: []ProtoConfig
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1ProtoconfigsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ProtoconfigsGetRequest struct via the builder pattern


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


## V1VaultTokenpairsGet

> []TokenPair V1VaultTokenpairsGet(ctx).TokenA(tokenA).TokenB(tokenB).Execute()

Get all Supported Token Pairs



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
    resp, r, err := apiClient.DefaultApi.V1VaultTokenpairsGet(context.Background()).TokenA(tokenA).TokenB(tokenB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1VaultTokenpairsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1VaultTokenpairsGet`: []TokenPair
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1VaultTokenpairsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1VaultTokenpairsGetRequest struct via the builder pattern


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


## V1VaultTokensGet

> []Token V1VaultTokensGet(ctx).TokenA(tokenA).TokenB(tokenB).Execute()

Get all Supported Tokens



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
    resp, r, err := apiClient.DefaultApi.V1VaultTokensGet(context.Background()).TokenA(tokenA).TokenB(tokenB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1VaultTokensGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1VaultTokensGet`: []Token
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1VaultTokensGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1VaultTokensGetRequest struct via the builder pattern


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


## V1VaultperiodsGet

> []VaultPeriod V1VaultperiodsGet(ctx).Vault(vault).VaultPeriod(vaultPeriod).Offset(offset).Limit(limit).Execute()

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
    resp, r, err := apiClient.DefaultApi.V1VaultperiodsGet(context.Background()).Vault(vault).VaultPeriod(vaultPeriod).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1VaultperiodsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1VaultperiodsGet`: []VaultPeriod
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1VaultperiodsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1VaultperiodsGetRequest struct via the builder pattern


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


## V1VaultsGet

> []Vault V1VaultsGet(ctx).TokenA(tokenA).TokenB(tokenB).ProtoConfig(protoConfig).Execute()

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
    resp, r, err := apiClient.DefaultApi.V1VaultsGet(context.Background()).TokenA(tokenA).TokenB(tokenB).ProtoConfig(protoConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1VaultsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1VaultsGet`: []Vault
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1VaultsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1VaultsGetRequest struct via the builder pattern


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

