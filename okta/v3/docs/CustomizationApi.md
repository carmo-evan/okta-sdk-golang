# \CustomizationApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailCustomization**](CustomizationApi.md#CreateEmailCustomization) | **Post** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations | Create an Email Customization
[**DeleteAllCustomizations**](CustomizationApi.md#DeleteAllCustomizations) | **Delete** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations | Delete all Email Customizations
[**DeleteBrandThemeBackgroundImage**](CustomizationApi.md#DeleteBrandThemeBackgroundImage) | **Delete** /api/v1/brands/{brandId}/themes/{themeId}/background-image | Delete the Background Image
[**DeleteBrandThemeFavicon**](CustomizationApi.md#DeleteBrandThemeFavicon) | **Delete** /api/v1/brands/{brandId}/themes/{themeId}/favicon | Delete the Favicon
[**DeleteBrandThemeLogo**](CustomizationApi.md#DeleteBrandThemeLogo) | **Delete** /api/v1/brands/{brandId}/themes/{themeId}/logo | Delete the Logo
[**DeleteEmailCustomization**](CustomizationApi.md#DeleteEmailCustomization) | **Delete** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId} | Delete an Email Customization
[**GetBrand**](CustomizationApi.md#GetBrand) | **Get** /api/v1/brands/{brandId} | Retrieve a Brand
[**GetBrandTheme**](CustomizationApi.md#GetBrandTheme) | **Get** /api/v1/brands/{brandId}/themes/{themeId} | Retrieve a Theme
[**GetCustomizationPreview**](CustomizationApi.md#GetCustomizationPreview) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId}/preview | Preview an Email Customization
[**GetEmailCustomization**](CustomizationApi.md#GetEmailCustomization) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId} | Retrieve an Email Customization
[**GetEmailDefaultContent**](CustomizationApi.md#GetEmailDefaultContent) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/default-content | Retrieve an Email Template Default Content
[**GetEmailDefaultPreview**](CustomizationApi.md#GetEmailDefaultPreview) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/default-content/preview | Preview the Email Template Default Content
[**GetEmailSettings**](CustomizationApi.md#GetEmailSettings) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/settings | Retrieve the Email Template Settings
[**GetEmailTemplate**](CustomizationApi.md#GetEmailTemplate) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName} | Retrieve an Email Template
[**ListAllSignInWidgetVersions**](CustomizationApi.md#ListAllSignInWidgetVersions) | **Get** /api/v1/brands/{brandId}/pages/sign-in/widget-versions | List all Sign-in Widget Versions
[**ListBrandThemes**](CustomizationApi.md#ListBrandThemes) | **Get** /api/v1/brands/{brandId}/themes | List all Themes
[**ListBrands**](CustomizationApi.md#ListBrands) | **Get** /api/v1/brands | List all Brands
[**ListEmailCustomizations**](CustomizationApi.md#ListEmailCustomizations) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations | List all Email Customizations
[**ListEmailTemplates**](CustomizationApi.md#ListEmailTemplates) | **Get** /api/v1/brands/{brandId}/templates/email | List all Email Templates
[**PreviewErrorPage**](CustomizationApi.md#PreviewErrorPage) | **Post** /api/v1/brands/{brandId}/pages/error/preview | Preview the Error Page
[**ReplaceErrorPage**](CustomizationApi.md#ReplaceErrorPage) | **Put** /api/v1/brands/{brandId}/pages/error | Replace the Error Page
[**ReplaceSignInPage**](CustomizationApi.md#ReplaceSignInPage) | **Put** /api/v1/brands/{brandId}/pages/sign-in | Replace the Sign-in Page
[**ReplaceSignInPagePreview**](CustomizationApi.md#ReplaceSignInPagePreview) | **Put** /api/v1/brands/{brandId}/pages/sign-in/preview | Replace the Sign-in Page Preview
[**ReplaceSignOutPageSettings**](CustomizationApi.md#ReplaceSignOutPageSettings) | **Put** /api/v1/brands/{brandId}/pages/sign-out | Replace the Sign-out Page Settings
[**ResetErrorPage**](CustomizationApi.md#ResetErrorPage) | **Delete** /api/v1/brands/{brandId}/pages/error | Reset the Error Page
[**ResetSignInPage**](CustomizationApi.md#ResetSignInPage) | **Delete** /api/v1/brands/{brandId}/pages/sign-in | Reset the Sign-in Page
[**ResetSignInPagePreview**](CustomizationApi.md#ResetSignInPagePreview) | **Delete** /api/v1/brands/{brandId}/pages/sign-in/preview | Reset the Sign-in Page Preview
[**RetrieveErrorPage**](CustomizationApi.md#RetrieveErrorPage) | **Get** /api/v1/brands/{brandId}/pages/error | Retrieve the Error Page
[**RetrieveSignInPage**](CustomizationApi.md#RetrieveSignInPage) | **Get** /api/v1/brands/{brandId}/pages/sign-in | Retrieve the Sign-in Page
[**RetrieveSignInPagePreview**](CustomizationApi.md#RetrieveSignInPagePreview) | **Get** /api/v1/brands/{brandId}/pages/sign-in/preview | Retrieve the Sign-in Page Preview
[**RetrieveSignOutPageSettings**](CustomizationApi.md#RetrieveSignOutPageSettings) | **Get** /api/v1/brands/{brandId}/pages/sign-out | Retrieve the Sign-out Page Settings
[**SendTestEmail**](CustomizationApi.md#SendTestEmail) | **Post** /api/v1/brands/{brandId}/templates/email/{templateName}/test | Send a Test Email
[**UpdateBrand**](CustomizationApi.md#UpdateBrand) | **Put** /api/v1/brands/{brandId} | Replace a Brand
[**UpdateBrandTheme**](CustomizationApi.md#UpdateBrandTheme) | **Put** /api/v1/brands/{brandId}/themes/{themeId} | Replace a Theme
[**UpdateEmailCustomization**](CustomizationApi.md#UpdateEmailCustomization) | **Put** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId} | Replace an Email Customization
[**UpdateEmailSettings**](CustomizationApi.md#UpdateEmailSettings) | **Put** /api/v1/brands/{brandId}/templates/email/{templateName}/settings | Replace the Email Template Settings
[**UploadBrandThemeBackgroundImage**](CustomizationApi.md#UploadBrandThemeBackgroundImage) | **Post** /api/v1/brands/{brandId}/themes/{themeId}/background-image | Upload the Background Image
[**UploadBrandThemeFavicon**](CustomizationApi.md#UploadBrandThemeFavicon) | **Post** /api/v1/brands/{brandId}/themes/{themeId}/favicon | Upload the Favicon
[**UploadBrandThemeLogo**](CustomizationApi.md#UploadBrandThemeLogo) | **Post** /api/v1/brands/{brandId}/themes/{themeId}/logo | Upload the Logo



## CreateEmailCustomization

> EmailCustomization CreateEmailCustomization(ctx, brandId, templateName).Instance(instance).Execute()

Create an Email Customization



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
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    instance := *openapiclient.NewEmailCustomization("Body_example", "Subject_example", "Language_example") // EmailCustomization |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.CreateEmailCustomization(context.Background(), brandId, templateName).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.CreateEmailCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmailCustomization`: EmailCustomization
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.CreateEmailCustomization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailCustomizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instance** | [**EmailCustomization**](EmailCustomization.md) |  | 

### Return type

[**EmailCustomization**](EmailCustomization.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllCustomizations

> DeleteAllCustomizations(ctx, brandId, templateName).Execute()

Delete all Email Customizations



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
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.DeleteAllCustomizations(context.Background(), brandId, templateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.DeleteAllCustomizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllCustomizationsRequest struct via the builder pattern


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


## DeleteBrandThemeBackgroundImage

> DeleteBrandThemeBackgroundImage(ctx, brandId, themeId).Execute()

Delete the Background Image



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
    brandId := "brandId_example" // string | 
    themeId := "themeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.DeleteBrandThemeBackgroundImage(context.Background(), brandId, themeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.DeleteBrandThemeBackgroundImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 
**themeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBrandThemeBackgroundImageRequest struct via the builder pattern


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


## DeleteBrandThemeFavicon

> DeleteBrandThemeFavicon(ctx, brandId, themeId).Execute()

Delete the Favicon



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
    brandId := "brandId_example" // string | 
    themeId := "themeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.DeleteBrandThemeFavicon(context.Background(), brandId, themeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.DeleteBrandThemeFavicon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 
**themeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBrandThemeFaviconRequest struct via the builder pattern


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


## DeleteBrandThemeLogo

> DeleteBrandThemeLogo(ctx, brandId, themeId).Execute()

Delete the Logo



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
    brandId := "brandId_example" // string | 
    themeId := "themeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.DeleteBrandThemeLogo(context.Background(), brandId, themeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.DeleteBrandThemeLogo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 
**themeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBrandThemeLogoRequest struct via the builder pattern


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


## DeleteEmailCustomization

> DeleteEmailCustomization(ctx, brandId, templateName, customizationId).Execute()

Delete an Email Customization



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
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    customizationId := "customizationId_example" // string | The ID of the email customization.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.DeleteEmailCustomization(context.Background(), brandId, templateName, customizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.DeleteEmailCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 
**customizationId** | **string** | The ID of the email customization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailCustomizationRequest struct via the builder pattern


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


## GetBrand

> Brand GetBrand(ctx, brandId).Execute()

Retrieve a Brand



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
    brandId := "brandId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetBrand(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrand`: Brand
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Brand**](Brand.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandTheme

> ThemeResponse GetBrandTheme(ctx, brandId, themeId).Execute()

Retrieve a Theme



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
    brandId := "brandId_example" // string | 
    themeId := "themeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetBrandTheme(context.Background(), brandId, themeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetBrandTheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandTheme`: ThemeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetBrandTheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 
**themeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandThemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ThemeResponse**](ThemeResponse.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomizationPreview

> EmailPreview GetCustomizationPreview(ctx, brandId, templateName, customizationId).Execute()

Preview an Email Customization



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
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    customizationId := "customizationId_example" // string | The ID of the email customization.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetCustomizationPreview(context.Background(), brandId, templateName, customizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetCustomizationPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomizationPreview`: EmailPreview
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetCustomizationPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 
**customizationId** | **string** | The ID of the email customization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomizationPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EmailPreview**](EmailPreview.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailCustomization

> EmailCustomization GetEmailCustomization(ctx, brandId, templateName, customizationId).Execute()

Retrieve an Email Customization



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
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    customizationId := "customizationId_example" // string | The ID of the email customization.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetEmailCustomization(context.Background(), brandId, templateName, customizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetEmailCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailCustomization`: EmailCustomization
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetEmailCustomization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 
**customizationId** | **string** | The ID of the email customization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailCustomizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EmailCustomization**](EmailCustomization.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailDefaultContent

> EmailDefaultContent GetEmailDefaultContent(ctx, brandId, templateName).Language(language).Execute()

Retrieve an Email Template Default Content



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
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    language := "language_example" // string | The language to use for the email. Defaults to the current user's language if unspecified. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetEmailDefaultContent(context.Background(), brandId, templateName).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetEmailDefaultContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailDefaultContent`: EmailDefaultContent
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetEmailDefaultContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailDefaultContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **language** | **string** | The language to use for the email. Defaults to the current user&#39;s language if unspecified. | 

### Return type

[**EmailDefaultContent**](EmailDefaultContent.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailDefaultPreview

> EmailPreview GetEmailDefaultPreview(ctx, brandId, templateName).Language(language).Execute()

Preview the Email Template Default Content



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
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    language := "language_example" // string | The language to use for the email. Defaults to the current user's language if unspecified. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetEmailDefaultPreview(context.Background(), brandId, templateName).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetEmailDefaultPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailDefaultPreview`: EmailPreview
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetEmailDefaultPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailDefaultPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **language** | **string** | The language to use for the email. Defaults to the current user&#39;s language if unspecified. | 

### Return type

[**EmailPreview**](EmailPreview.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailSettings

> EmailSettings GetEmailSettings(ctx, brandId, templateName).Execute()

Retrieve the Email Template Settings



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
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetEmailSettings(context.Background(), brandId, templateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetEmailSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailSettings`: EmailSettings
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetEmailSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EmailSettings**](EmailSettings.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailTemplate

> EmailTemplate GetEmailTemplate(ctx, brandId, templateName).Expand(expand).Execute()

Retrieve an Email Template



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
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetEmailTemplate(context.Background(), brandId, templateName).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetEmailTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailTemplate`: EmailTemplate
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetEmailTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **[]string** | Specifies additional metadata to be included in the response. | 

### Return type

[**EmailTemplate**](EmailTemplate.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllSignInWidgetVersions

> []string ListAllSignInWidgetVersions(ctx, brandId).Execute()

List all Sign-in Widget Versions



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
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ListAllSignInWidgetVersions(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ListAllSignInWidgetVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllSignInWidgetVersions`: []string
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ListAllSignInWidgetVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllSignInWidgetVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBrandThemes

> []ThemeResponse ListBrandThemes(ctx, brandId).Execute()

List all Themes



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
    brandId := "brandId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ListBrandThemes(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ListBrandThemes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBrandThemes`: []ThemeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ListBrandThemes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBrandThemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ThemeResponse**](ThemeResponse.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBrands

> []Brand ListBrands(ctx).Execute()

List all Brands



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
    resp, r, err := apiClient.CustomizationApi.ListBrands(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ListBrands``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBrands`: []Brand
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ListBrands`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBrandsRequest struct via the builder pattern


### Return type

[**[]Brand**](Brand.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailCustomizations

> []EmailCustomization ListEmailCustomizations(ctx, brandId, templateName).After(after).Limit(limit).Execute()

List all Email Customizations



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
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/reference/core-okta-api/#pagination) for more information. (optional)
    limit := int32(56) // int32 | A limit on the number of objects to return. (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ListEmailCustomizations(context.Background(), brandId, templateName).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ListEmailCustomizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailCustomizations`: []EmailCustomization
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ListEmailCustomizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEmailCustomizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/reference/core-okta-api/#pagination) for more information. | 
 **limit** | **int32** | A limit on the number of objects to return. | [default to 20]

### Return type

[**[]EmailCustomization**](EmailCustomization.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailTemplates

> []EmailTemplate ListEmailTemplates(ctx, brandId).After(after).Limit(limit).Expand(expand).Execute()

List all Email Templates



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
    brandId := "brandId_example" // string | The ID of the brand.
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/reference/core-okta-api/#pagination) for more information. (optional)
    limit := int32(56) // int32 | A limit on the number of objects to return. (optional) (default to 20)
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ListEmailTemplates(context.Background(), brandId).After(after).Limit(limit).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ListEmailTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailTemplates`: []EmailTemplate
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ListEmailTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEmailTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/reference/core-okta-api/#pagination) for more information. | 
 **limit** | **int32** | A limit on the number of objects to return. | [default to 20]
 **expand** | **[]string** | Specifies additional metadata to be included in the response. | 

### Return type

[**[]EmailTemplate**](EmailTemplate.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewErrorPage

> string PreviewErrorPage(ctx, brandId).CustomizablePage(customizablePage).Execute()

Preview the Error Page



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
    brandId := "brandId_example" // string | The ID of the brand.
    customizablePage := *openapiclient.NewCustomizablePage("PageContent_example") // CustomizablePage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.PreviewErrorPage(context.Background(), brandId).CustomizablePage(customizablePage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.PreviewErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewErrorPage`: string
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.PreviewErrorPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewErrorPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customizablePage** | [**CustomizablePage**](CustomizablePage.md) |  | 

### Return type

**string**

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/html, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceErrorPage

> CustomizablePage ReplaceErrorPage(ctx, brandId).CustomizablePage(customizablePage).Execute()

Replace the Error Page



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
    brandId := "brandId_example" // string | The ID of the brand.
    customizablePage := *openapiclient.NewCustomizablePage("PageContent_example") // CustomizablePage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ReplaceErrorPage(context.Background(), brandId).CustomizablePage(customizablePage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ReplaceErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceErrorPage`: CustomizablePage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ReplaceErrorPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceErrorPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customizablePage** | [**CustomizablePage**](CustomizablePage.md) |  | 

### Return type

[**CustomizablePage**](CustomizablePage.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSignInPage

> SignInPage ReplaceSignInPage(ctx, brandId).SignInPage(signInPage).Execute()

Replace the Sign-in Page



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
    brandId := "brandId_example" // string | The ID of the brand.
    signInPage := *openapiclient.NewSignInPage("PageContent_example", "Type_example") // SignInPage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ReplaceSignInPage(context.Background(), brandId).SignInPage(signInPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ReplaceSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceSignInPage`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ReplaceSignInPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSignInPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signInPage** | [**SignInPage**](SignInPage.md) |  | 

### Return type

[**SignInPage**](SignInPage.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSignInPagePreview

> ReplaceSignInPagePreview(ctx, brandId).SignInPage(signInPage).Execute()

Replace the Sign-in Page Preview



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
    brandId := "brandId_example" // string | The ID of the brand.
    signInPage := *openapiclient.NewSignInPage("PageContent_example", "Type_example") // SignInPage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ReplaceSignInPagePreview(context.Background(), brandId).SignInPage(signInPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ReplaceSignInPagePreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSignInPagePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signInPage** | [**SignInPage**](SignInPage.md) |  | 

### Return type

 (empty response body)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSignOutPageSettings

> HostedPage ReplaceSignOutPageSettings(ctx, brandId).HostedPage(hostedPage).Execute()

Replace the Sign-out Page Settings



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
    brandId := "brandId_example" // string | The ID of the brand.
    hostedPage := *openapiclient.NewHostedPage("Type_example") // HostedPage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ReplaceSignOutPageSettings(context.Background(), brandId).HostedPage(hostedPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ReplaceSignOutPageSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceSignOutPageSettings`: HostedPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ReplaceSignOutPageSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSignOutPageSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostedPage** | [**HostedPage**](HostedPage.md) |  | 

### Return type

[**HostedPage**](HostedPage.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetErrorPage

> ResetErrorPage(ctx, brandId).Execute()

Reset the Error Page



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
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ResetErrorPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ResetErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetErrorPageRequest struct via the builder pattern


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


## ResetSignInPage

> ResetSignInPage(ctx, brandId).Execute()

Reset the Sign-in Page



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
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ResetSignInPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ResetSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetSignInPageRequest struct via the builder pattern


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


## ResetSignInPagePreview

> ResetSignInPagePreview(ctx, brandId).Execute()

Reset the Sign-in Page Preview



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
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ResetSignInPagePreview(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ResetSignInPagePreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetSignInPagePreviewRequest struct via the builder pattern


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


## RetrieveErrorPage

> CustomizablePage RetrieveErrorPage(ctx, brandId).Execute()

Retrieve the Error Page



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
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.RetrieveErrorPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.RetrieveErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveErrorPage`: CustomizablePage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.RetrieveErrorPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveErrorPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomizablePage**](CustomizablePage.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSignInPage

> SignInPage RetrieveSignInPage(ctx, brandId).Execute()

Retrieve the Sign-in Page



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
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.RetrieveSignInPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.RetrieveSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSignInPage`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.RetrieveSignInPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSignInPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SignInPage**](SignInPage.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSignInPagePreview

> SignInPage RetrieveSignInPagePreview(ctx, brandId).Execute()

Retrieve the Sign-in Page Preview



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
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.RetrieveSignInPagePreview(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.RetrieveSignInPagePreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSignInPagePreview`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.RetrieveSignInPagePreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSignInPagePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SignInPage**](SignInPage.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSignOutPageSettings

> HostedPage RetrieveSignOutPageSettings(ctx, brandId).Execute()

Retrieve the Sign-out Page Settings



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
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.RetrieveSignOutPageSettings(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.RetrieveSignOutPageSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSignOutPageSettings`: HostedPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.RetrieveSignOutPageSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSignOutPageSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostedPage**](HostedPage.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTestEmail

> SendTestEmail(ctx, brandId, templateName).Language(language).Execute()

Send a Test Email



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
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    language := "language_example" // string | The language to use for the email. Defaults to the current user's language if unspecified. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.SendTestEmail(context.Background(), brandId, templateName).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.SendTestEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendTestEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **language** | **string** | The language to use for the email. Defaults to the current user&#39;s language if unspecified. | 

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


## UpdateBrand

> Brand UpdateBrand(ctx, brandId).Brand(brand).Execute()

Replace a Brand



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
    brandId := "brandId_example" // string | 
    brand := *openapiclient.NewBrand() // Brand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.UpdateBrand(context.Background(), brandId).Brand(brand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.UpdateBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBrand`: Brand
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.UpdateBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **brand** | [**Brand**](Brand.md) |  | 

### Return type

[**Brand**](Brand.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBrandTheme

> ThemeResponse UpdateBrandTheme(ctx, brandId, themeId).Theme(theme).Execute()

Replace a Theme



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
    brandId := "brandId_example" // string | 
    themeId := "themeId_example" // string | 
    theme := *openapiclient.NewTheme() // Theme | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.UpdateBrandTheme(context.Background(), brandId, themeId).Theme(theme).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.UpdateBrandTheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBrandTheme`: ThemeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.UpdateBrandTheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 
**themeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBrandThemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **theme** | [**Theme**](Theme.md) |  | 

### Return type

[**ThemeResponse**](ThemeResponse.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailCustomization

> EmailCustomization UpdateEmailCustomization(ctx, brandId, templateName, customizationId).Instance(instance).Execute()

Replace an Email Customization



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
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    customizationId := "customizationId_example" // string | The ID of the email customization.
    instance := *openapiclient.NewEmailCustomization("Body_example", "Subject_example", "Language_example") // EmailCustomization | Request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.UpdateEmailCustomization(context.Background(), brandId, templateName, customizationId).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.UpdateEmailCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEmailCustomization`: EmailCustomization
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.UpdateEmailCustomization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 
**customizationId** | **string** | The ID of the email customization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailCustomizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **instance** | [**EmailCustomization**](EmailCustomization.md) | Request | 

### Return type

[**EmailCustomization**](EmailCustomization.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailSettings

> UpdateEmailSettings(ctx, brandId, templateName).EmailSettings(emailSettings).Execute()

Replace the Email Template Settings



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
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    emailSettings := *openapiclient.NewEmailSettings("Recipients_example") // EmailSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.UpdateEmailSettings(context.Background(), brandId, templateName).EmailSettings(emailSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.UpdateEmailSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **emailSettings** | [**EmailSettings**](EmailSettings.md) |  | 

### Return type

 (empty response body)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadBrandThemeBackgroundImage

> ImageUploadResponse UploadBrandThemeBackgroundImage(ctx, brandId, themeId).File(file).Execute()

Upload the Background Image



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
    brandId := "brandId_example" // string | 
    themeId := "themeId_example" // string | 
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.UploadBrandThemeBackgroundImage(context.Background(), brandId, themeId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.UploadBrandThemeBackgroundImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadBrandThemeBackgroundImage`: ImageUploadResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.UploadBrandThemeBackgroundImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 
**themeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadBrandThemeBackgroundImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** |  | 

### Return type

[**ImageUploadResponse**](ImageUploadResponse.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadBrandThemeFavicon

> ImageUploadResponse UploadBrandThemeFavicon(ctx, brandId, themeId).File(file).Execute()

Upload the Favicon



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
    brandId := "brandId_example" // string | 
    themeId := "themeId_example" // string | 
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.UploadBrandThemeFavicon(context.Background(), brandId, themeId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.UploadBrandThemeFavicon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadBrandThemeFavicon`: ImageUploadResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.UploadBrandThemeFavicon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 
**themeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadBrandThemeFaviconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** |  | 

### Return type

[**ImageUploadResponse**](ImageUploadResponse.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadBrandThemeLogo

> ImageUploadResponse UploadBrandThemeLogo(ctx, brandId, themeId).File(file).Execute()

Upload the Logo



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
    brandId := "brandId_example" // string | 
    themeId := "themeId_example" // string | 
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.UploadBrandThemeLogo(context.Background(), brandId, themeId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.UploadBrandThemeLogo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadBrandThemeLogo`: ImageUploadResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.UploadBrandThemeLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 
**themeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadBrandThemeLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** |  | 

### Return type

[**ImageUploadResponse**](ImageUploadResponse.md)

### Authorization

[API_Token](../README.md#API_Token), [OAuth_2.0](../README.md#OAuth_2.0)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

