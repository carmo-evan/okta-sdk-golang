# \OrgSettingApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkRemoveEmailAddressBounces**](OrgSettingApi.md#BulkRemoveEmailAddressBounces) | **Post** /api/v1/org/email/bounces/remove-list | Remove Emails from Email Provider Bounce List
[**ExtendOktaSupport**](OrgSettingApi.md#ExtendOktaSupport) | **Post** /api/v1/org/privacy/oktaSupport/extend | Extend Okta Support Access
[**GetOktaCommunicationSettings**](OrgSettingApi.md#GetOktaCommunicationSettings) | **Get** /api/v1/org/privacy/oktaCommunication | Retreive the Okta Communication Settings
[**GetOrgContactTypes**](OrgSettingApi.md#GetOrgContactTypes) | **Get** /api/v1/org/contacts | Retrieve the Org Contact Types
[**GetOrgContactUser**](OrgSettingApi.md#GetOrgContactUser) | **Get** /api/v1/org/contacts/{contactType} | Retrieve the User of the Contact Type
[**GetOrgOktaSupportSettings**](OrgSettingApi.md#GetOrgOktaSupportSettings) | **Get** /api/v1/org/privacy/oktaSupport | Retrieve the Okta Support Settings
[**GetOrgPreferences**](OrgSettingApi.md#GetOrgPreferences) | **Get** /api/v1/org/preferences | Retrieve the Org Preferences
[**GetOrgSettings**](OrgSettingApi.md#GetOrgSettings) | **Get** /api/v1/org | Retrieve the Org Settings
[**GrantOktaSupport**](OrgSettingApi.md#GrantOktaSupport) | **Post** /api/v1/org/privacy/oktaSupport/grant | Grant Okta Support Access to your Org
[**HideOktaUIFooter**](OrgSettingApi.md#HideOktaUIFooter) | **Post** /api/v1/org/preferences/hideEndUserFooter | Update the Preference to Hide the Okta Dashboard Footer
[**OptInUsersToOktaCommunicationEmails**](OrgSettingApi.md#OptInUsersToOktaCommunicationEmails) | **Post** /api/v1/org/privacy/oktaCommunication/optIn | Opt in all Users to Okta Communication emails
[**OptOutUsersFromOktaCommunicationEmails**](OrgSettingApi.md#OptOutUsersFromOktaCommunicationEmails) | **Post** /api/v1/org/privacy/oktaCommunication/optOut | Opt out all Users from Okta Communication emails
[**PartialUpdateOrgSetting**](OrgSettingApi.md#PartialUpdateOrgSetting) | **Post** /api/v1/org | Update the Org Settings
[**RevokeOktaSupport**](OrgSettingApi.md#RevokeOktaSupport) | **Post** /api/v1/org/privacy/oktaSupport/revoke | Revoke Okta Support Access
[**ShowOktaUIFooter**](OrgSettingApi.md#ShowOktaUIFooter) | **Post** /api/v1/org/preferences/showEndUserFooter | Update the Preference to Show the Okta Dashboard Footer
[**UpdateOrgContactUser**](OrgSettingApi.md#UpdateOrgContactUser) | **Put** /api/v1/org/contacts/{contactType} | Replace the User of the Contact Type
[**UpdateOrgLogo**](OrgSettingApi.md#UpdateOrgLogo) | **Post** /api/v1/org/logo | Upload the Org Logo
[**UpdateOrgSetting**](OrgSettingApi.md#UpdateOrgSetting) | **Put** /api/v1/org | Replace the Org Settings



## BulkRemoveEmailAddressBounces

> BouncesRemoveListResult BulkRemoveEmailAddressBounces(ctx).BouncesRemoveListObj(bouncesRemoveListObj).Execute()

Remove Emails from Email Provider Bounce List



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
    bouncesRemoveListObj := *openapiclient.NewBouncesRemoveListObj() // BouncesRemoveListObj |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingApi.BulkRemoveEmailAddressBounces(context.Background()).BouncesRemoveListObj(bouncesRemoveListObj).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.BulkRemoveEmailAddressBounces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkRemoveEmailAddressBounces`: BouncesRemoveListResult
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingApi.BulkRemoveEmailAddressBounces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkRemoveEmailAddressBouncesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bouncesRemoveListObj** | [**BouncesRemoveListObj**](BouncesRemoveListObj.md) |  | 

### Return type

[**BouncesRemoveListResult**](BouncesRemoveListResult.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtendOktaSupport

> OrgOktaSupportSettingsObj ExtendOktaSupport(ctx).Execute()

Extend Okta Support Access



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
    resp, r, err := apiClient.OrgSettingApi.ExtendOktaSupport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.ExtendOktaSupport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtendOktaSupport`: OrgOktaSupportSettingsObj
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingApi.ExtendOktaSupport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiExtendOktaSupportRequest struct via the builder pattern


### Return type

[**OrgOktaSupportSettingsObj**](OrgOktaSupportSettingsObj.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOktaCommunicationSettings

> OrgOktaCommunicationSetting GetOktaCommunicationSettings(ctx).Execute()

Retreive the Okta Communication Settings



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
    resp, r, err := apiClient.OrgSettingApi.GetOktaCommunicationSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.GetOktaCommunicationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOktaCommunicationSettings`: OrgOktaCommunicationSetting
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingApi.GetOktaCommunicationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOktaCommunicationSettingsRequest struct via the builder pattern


### Return type

[**OrgOktaCommunicationSetting**](OrgOktaCommunicationSetting.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgContactTypes

> []OrgContactTypeObj GetOrgContactTypes(ctx).Execute()

Retrieve the Org Contact Types



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
    resp, r, err := apiClient.OrgSettingApi.GetOrgContactTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.GetOrgContactTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgContactTypes`: []OrgContactTypeObj
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingApi.GetOrgContactTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgContactTypesRequest struct via the builder pattern


### Return type

[**[]OrgContactTypeObj**](OrgContactTypeObj.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgContactUser

> OrgContactUser GetOrgContactUser(ctx, contactType).Execute()

Retrieve the User of the Contact Type



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
    contactType := "contactType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingApi.GetOrgContactUser(context.Background(), contactType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.GetOrgContactUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgContactUser`: OrgContactUser
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingApi.GetOrgContactUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgContactUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgContactUser**](OrgContactUser.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgOktaSupportSettings

> OrgOktaSupportSettingsObj GetOrgOktaSupportSettings(ctx).Execute()

Retrieve the Okta Support Settings



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
    resp, r, err := apiClient.OrgSettingApi.GetOrgOktaSupportSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.GetOrgOktaSupportSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgOktaSupportSettings`: OrgOktaSupportSettingsObj
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingApi.GetOrgOktaSupportSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgOktaSupportSettingsRequest struct via the builder pattern


### Return type

[**OrgOktaSupportSettingsObj**](OrgOktaSupportSettingsObj.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgPreferences

> OrgPreferences GetOrgPreferences(ctx).Execute()

Retrieve the Org Preferences



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
    resp, r, err := apiClient.OrgSettingApi.GetOrgPreferences(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.GetOrgPreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgPreferences`: OrgPreferences
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingApi.GetOrgPreferences`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgPreferencesRequest struct via the builder pattern


### Return type

[**OrgPreferences**](OrgPreferences.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgSettings

> OrgSetting GetOrgSettings(ctx).Execute()

Retrieve the Org Settings



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
    resp, r, err := apiClient.OrgSettingApi.GetOrgSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.GetOrgSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgSettings`: OrgSetting
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingApi.GetOrgSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSettingsRequest struct via the builder pattern


### Return type

[**OrgSetting**](OrgSetting.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantOktaSupport

> OrgOktaSupportSettingsObj GrantOktaSupport(ctx).Execute()

Grant Okta Support Access to your Org



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
    resp, r, err := apiClient.OrgSettingApi.GrantOktaSupport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.GrantOktaSupport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GrantOktaSupport`: OrgOktaSupportSettingsObj
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingApi.GrantOktaSupport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGrantOktaSupportRequest struct via the builder pattern


### Return type

[**OrgOktaSupportSettingsObj**](OrgOktaSupportSettingsObj.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HideOktaUIFooter

> OrgPreferences HideOktaUIFooter(ctx).Execute()

Update the Preference to Hide the Okta Dashboard Footer



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
    resp, r, err := apiClient.OrgSettingApi.HideOktaUIFooter(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.HideOktaUIFooter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HideOktaUIFooter`: OrgPreferences
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingApi.HideOktaUIFooter`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHideOktaUIFooterRequest struct via the builder pattern


### Return type

[**OrgPreferences**](OrgPreferences.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptInUsersToOktaCommunicationEmails

> OrgOktaCommunicationSetting OptInUsersToOktaCommunicationEmails(ctx).Execute()

Opt in all Users to Okta Communication emails



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
    resp, r, err := apiClient.OrgSettingApi.OptInUsersToOktaCommunicationEmails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.OptInUsersToOktaCommunicationEmails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OptInUsersToOktaCommunicationEmails`: OrgOktaCommunicationSetting
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingApi.OptInUsersToOktaCommunicationEmails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptInUsersToOktaCommunicationEmailsRequest struct via the builder pattern


### Return type

[**OrgOktaCommunicationSetting**](OrgOktaCommunicationSetting.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptOutUsersFromOktaCommunicationEmails

> OrgOktaCommunicationSetting OptOutUsersFromOktaCommunicationEmails(ctx).Execute()

Opt out all Users from Okta Communication emails



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
    resp, r, err := apiClient.OrgSettingApi.OptOutUsersFromOktaCommunicationEmails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.OptOutUsersFromOktaCommunicationEmails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OptOutUsersFromOktaCommunicationEmails`: OrgOktaCommunicationSetting
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingApi.OptOutUsersFromOktaCommunicationEmails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptOutUsersFromOktaCommunicationEmailsRequest struct via the builder pattern


### Return type

[**OrgOktaCommunicationSetting**](OrgOktaCommunicationSetting.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateOrgSetting

> OrgSetting PartialUpdateOrgSetting(ctx).OrgSetting(orgSetting).Execute()

Update the Org Settings



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
    orgSetting := *openapiclient.NewOrgSetting() // OrgSetting |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingApi.PartialUpdateOrgSetting(context.Background()).OrgSetting(orgSetting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.PartialUpdateOrgSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartialUpdateOrgSetting`: OrgSetting
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingApi.PartialUpdateOrgSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateOrgSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgSetting** | [**OrgSetting**](OrgSetting.md) |  | 

### Return type

[**OrgSetting**](OrgSetting.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeOktaSupport

> OrgOktaSupportSettingsObj RevokeOktaSupport(ctx).Execute()

Revoke Okta Support Access



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
    resp, r, err := apiClient.OrgSettingApi.RevokeOktaSupport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.RevokeOktaSupport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeOktaSupport`: OrgOktaSupportSettingsObj
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingApi.RevokeOktaSupport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeOktaSupportRequest struct via the builder pattern


### Return type

[**OrgOktaSupportSettingsObj**](OrgOktaSupportSettingsObj.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowOktaUIFooter

> OrgPreferences ShowOktaUIFooter(ctx).Execute()

Update the Preference to Show the Okta Dashboard Footer



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
    resp, r, err := apiClient.OrgSettingApi.ShowOktaUIFooter(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.ShowOktaUIFooter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowOktaUIFooter`: OrgPreferences
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingApi.ShowOktaUIFooter`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowOktaUIFooterRequest struct via the builder pattern


### Return type

[**OrgPreferences**](OrgPreferences.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgContactUser

> OrgContactUser UpdateOrgContactUser(ctx, contactType).OrgContactUser(orgContactUser).Execute()

Replace the User of the Contact Type



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
    contactType := "contactType_example" // string | 
    orgContactUser := *openapiclient.NewOrgContactUser() // OrgContactUser | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingApi.UpdateOrgContactUser(context.Background(), contactType).OrgContactUser(orgContactUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.UpdateOrgContactUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgContactUser`: OrgContactUser
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingApi.UpdateOrgContactUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgContactUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgContactUser** | [**OrgContactUser**](OrgContactUser.md) |  | 

### Return type

[**OrgContactUser**](OrgContactUser.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgLogo

> UpdateOrgLogo(ctx).File(file).Execute()

Upload the Org Logo



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
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingApi.UpdateOrgLogo(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.UpdateOrgLogo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgSetting

> OrgSetting UpdateOrgSetting(ctx).OrgSetting(orgSetting).Execute()

Replace the Org Settings



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
    orgSetting := *openapiclient.NewOrgSetting() // OrgSetting | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingApi.UpdateOrgSetting(context.Background()).OrgSetting(orgSetting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingApi.UpdateOrgSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgSetting`: OrgSetting
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingApi.UpdateOrgSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgSetting** | [**OrgSetting**](OrgSetting.md) |  | 

### Return type

[**OrgSetting**](OrgSetting.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

