# \TrustedOriginApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateOrigin**](TrustedOriginApi.md#ActivateOrigin) | **Post** /api/v1/trustedOrigins/{trustedOriginId}/lifecycle/activate | Activate a Trusted Origin
[**CreateOrigin**](TrustedOriginApi.md#CreateOrigin) | **Post** /api/v1/trustedOrigins | Create a Trusted Origin
[**DeactivateOrigin**](TrustedOriginApi.md#DeactivateOrigin) | **Post** /api/v1/trustedOrigins/{trustedOriginId}/lifecycle/deactivate | Deactivate a Trusted Origin
[**DeleteOrigin**](TrustedOriginApi.md#DeleteOrigin) | **Delete** /api/v1/trustedOrigins/{trustedOriginId} | Delete a Trusted Origin
[**GetOrigin**](TrustedOriginApi.md#GetOrigin) | **Get** /api/v1/trustedOrigins/{trustedOriginId} | Retrieve a Trusted Origin
[**ListOrigins**](TrustedOriginApi.md#ListOrigins) | **Get** /api/v1/trustedOrigins | List all Trusted Origins
[**UpdateOrigin**](TrustedOriginApi.md#UpdateOrigin) | **Put** /api/v1/trustedOrigins/{trustedOriginId} | Replace a Trusted Origin



## ActivateOrigin

> TrustedOrigin ActivateOrigin(ctx, trustedOriginId).Execute()

Activate a Trusted Origin



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
    trustedOriginId := "trustedOriginId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginApi.ActivateOrigin(context.Background(), trustedOriginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginApi.ActivateOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateOrigin`: TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginApi.ActivateOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedOriginId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrustedOrigin**](TrustedOrigin.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrigin

> TrustedOrigin CreateOrigin(ctx).TrustedOrigin(trustedOrigin).Execute()

Create a Trusted Origin



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
    trustedOrigin := *openapiclient.NewTrustedOrigin() // TrustedOrigin | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginApi.CreateOrigin(context.Background()).TrustedOrigin(trustedOrigin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginApi.CreateOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrigin`: TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginApi.CreateOrigin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trustedOrigin** | [**TrustedOrigin**](TrustedOrigin.md) |  | 

### Return type

[**TrustedOrigin**](TrustedOrigin.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateOrigin

> TrustedOrigin DeactivateOrigin(ctx, trustedOriginId).Execute()

Deactivate a Trusted Origin



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
    trustedOriginId := "trustedOriginId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginApi.DeactivateOrigin(context.Background(), trustedOriginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginApi.DeactivateOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateOrigin`: TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginApi.DeactivateOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedOriginId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrustedOrigin**](TrustedOrigin.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrigin

> DeleteOrigin(ctx, trustedOriginId).Execute()

Delete a Trusted Origin



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
    trustedOriginId := "trustedOriginId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginApi.DeleteOrigin(context.Background(), trustedOriginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginApi.DeleteOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedOriginId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrigin

> TrustedOrigin GetOrigin(ctx, trustedOriginId).Execute()

Retrieve a Trusted Origin



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
    trustedOriginId := "trustedOriginId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginApi.GetOrigin(context.Background(), trustedOriginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginApi.GetOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrigin`: TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginApi.GetOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedOriginId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrustedOrigin**](TrustedOrigin.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrigins

> []TrustedOrigin ListOrigins(ctx).Q(q).Filter(filter).After(after).Limit(limit).Execute()

List all Trusted Origins



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
    q := "q_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginApi.ListOrigins(context.Background()).Q(q).Filter(filter).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginApi.ListOrigins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrigins`: []TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginApi.ListOrigins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOriginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **filter** | **string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | [default to -1]

### Return type

[**[]TrustedOrigin**](TrustedOrigin.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrigin

> TrustedOrigin UpdateOrigin(ctx, trustedOriginId).TrustedOrigin(trustedOrigin).Execute()

Replace a Trusted Origin



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
    trustedOriginId := "trustedOriginId_example" // string | 
    trustedOrigin := *openapiclient.NewTrustedOrigin() // TrustedOrigin | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginApi.UpdateOrigin(context.Background(), trustedOriginId).TrustedOrigin(trustedOrigin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginApi.UpdateOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrigin`: TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginApi.UpdateOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedOriginId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trustedOrigin** | [**TrustedOrigin**](TrustedOrigin.md) |  | 

### Return type

[**TrustedOrigin**](TrustedOrigin.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

