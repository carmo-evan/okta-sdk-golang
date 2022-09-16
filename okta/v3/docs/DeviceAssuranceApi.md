# \DeviceAssuranceApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceAssurancePolicy**](DeviceAssuranceApi.md#CreateDeviceAssurancePolicy) | **Post** /api/v1/device-assurances | Create a Device Assurance Policy
[**DeleteDeviceAssurancePolicy**](DeviceAssuranceApi.md#DeleteDeviceAssurancePolicy) | **Delete** /api/v1/device-assurances/{deviceAssuranceId} | Delete a Device Assurance Policy
[**GetDeviceAssurancePolicy**](DeviceAssuranceApi.md#GetDeviceAssurancePolicy) | **Get** /api/v1/device-assurances/{deviceAssuranceId} | Retrieve a Device Assurance Policy
[**ListDeviceAssurancePolicies**](DeviceAssuranceApi.md#ListDeviceAssurancePolicies) | **Get** /api/v1/device-assurances | List all Device Assurance Policies
[**UpdateDeviceAssurancePolicy**](DeviceAssuranceApi.md#UpdateDeviceAssurancePolicy) | **Put** /api/v1/device-assurances/{deviceAssuranceId} | Replace a Device Assurance Policy



## CreateDeviceAssurancePolicy

> DeviceAssurance CreateDeviceAssurancePolicy(ctx).DeviceAssurance(deviceAssurance).Execute()

Create a Device Assurance Policy



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
    deviceAssurance := *openapiclient.NewDeviceAssurance() // DeviceAssurance | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAssuranceApi.CreateDeviceAssurancePolicy(context.Background()).DeviceAssurance(deviceAssurance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssuranceApi.CreateDeviceAssurancePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceAssurancePolicy`: DeviceAssurance
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssuranceApi.CreateDeviceAssurancePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceAssurancePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceAssurance** | [**DeviceAssurance**](DeviceAssurance.md) |  | 

### Return type

[**DeviceAssurance**](DeviceAssurance.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceAssurancePolicy

> DeleteDeviceAssurancePolicy(ctx, deviceAssuranceId).Execute()

Delete a Device Assurance Policy



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
    deviceAssuranceId := "deviceAssuranceId_example" // string | Id of the Device Assurance Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAssuranceApi.DeleteDeviceAssurancePolicy(context.Background(), deviceAssuranceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssuranceApi.DeleteDeviceAssurancePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceAssuranceId** | **string** | Id of the Device Assurance Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceAssurancePolicyRequest struct via the builder pattern


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


## GetDeviceAssurancePolicy

> DeviceAssurance GetDeviceAssurancePolicy(ctx, deviceAssuranceId).Execute()

Retrieve a Device Assurance Policy



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
    deviceAssuranceId := "deviceAssuranceId_example" // string | Id of the Device Assurance Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAssuranceApi.GetDeviceAssurancePolicy(context.Background(), deviceAssuranceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssuranceApi.GetDeviceAssurancePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceAssurancePolicy`: DeviceAssurance
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssuranceApi.GetDeviceAssurancePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceAssuranceId** | **string** | Id of the Device Assurance Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAssurancePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceAssurance**](DeviceAssurance.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceAssurancePolicies

> []DeviceAssurance ListDeviceAssurancePolicies(ctx).Execute()

List all Device Assurance Policies



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
    resp, r, err := apiClient.DeviceAssuranceApi.ListDeviceAssurancePolicies(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssuranceApi.ListDeviceAssurancePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeviceAssurancePolicies`: []DeviceAssurance
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssuranceApi.ListDeviceAssurancePolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceAssurancePoliciesRequest struct via the builder pattern


### Return type

[**[]DeviceAssurance**](DeviceAssurance.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceAssurancePolicy

> DeviceAssurance UpdateDeviceAssurancePolicy(ctx, deviceAssuranceId).DeviceAssurance(deviceAssurance).Execute()

Replace a Device Assurance Policy



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
    deviceAssuranceId := "deviceAssuranceId_example" // string | Id of the Device Assurance Policy
    deviceAssurance := *openapiclient.NewDeviceAssurance() // DeviceAssurance | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAssuranceApi.UpdateDeviceAssurancePolicy(context.Background(), deviceAssuranceId).DeviceAssurance(deviceAssurance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssuranceApi.UpdateDeviceAssurancePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceAssurancePolicy`: DeviceAssurance
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssuranceApi.UpdateDeviceAssurancePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceAssuranceId** | **string** | Id of the Device Assurance Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceAssurancePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceAssurance** | [**DeviceAssurance**](DeviceAssurance.md) |  | 

### Return type

[**DeviceAssurance**](DeviceAssurance.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

