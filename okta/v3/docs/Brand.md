# Brand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgreeToCustomPrivacyPolicy** | Pointer to **bool** |  | [optional] 
**CustomPrivacyPolicyUrl** | Pointer to **string** |  | [optional] 
**DefaultApp** | Pointer to [**BrandDefaultApp**](BrandDefaultApp.md) |  | [optional] 
**DisplayLanguage** | Pointer to **string** | The language specified as an [IETF BCP 47 language tag](https://datatracker.ietf.org/doc/html/rfc5646). | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**OptOutOfUserCommunications** | Pointer to **bool** |  | [optional] 
**RemovePoweredByOkta** | Pointer to **bool** |  | [optional] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewBrand

`func NewBrand() *Brand`

NewBrand instantiates a new Brand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandWithDefaults

`func NewBrandWithDefaults() *Brand`

NewBrandWithDefaults instantiates a new Brand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreeToCustomPrivacyPolicy

`func (o *Brand) GetAgreeToCustomPrivacyPolicy() bool`

GetAgreeToCustomPrivacyPolicy returns the AgreeToCustomPrivacyPolicy field if non-nil, zero value otherwise.

### GetAgreeToCustomPrivacyPolicyOk

`func (o *Brand) GetAgreeToCustomPrivacyPolicyOk() (*bool, bool)`

GetAgreeToCustomPrivacyPolicyOk returns a tuple with the AgreeToCustomPrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreeToCustomPrivacyPolicy

`func (o *Brand) SetAgreeToCustomPrivacyPolicy(v bool)`

SetAgreeToCustomPrivacyPolicy sets AgreeToCustomPrivacyPolicy field to given value.

### HasAgreeToCustomPrivacyPolicy

`func (o *Brand) HasAgreeToCustomPrivacyPolicy() bool`

HasAgreeToCustomPrivacyPolicy returns a boolean if a field has been set.

### GetCustomPrivacyPolicyUrl

`func (o *Brand) GetCustomPrivacyPolicyUrl() string`

GetCustomPrivacyPolicyUrl returns the CustomPrivacyPolicyUrl field if non-nil, zero value otherwise.

### GetCustomPrivacyPolicyUrlOk

`func (o *Brand) GetCustomPrivacyPolicyUrlOk() (*string, bool)`

GetCustomPrivacyPolicyUrlOk returns a tuple with the CustomPrivacyPolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrivacyPolicyUrl

`func (o *Brand) SetCustomPrivacyPolicyUrl(v string)`

SetCustomPrivacyPolicyUrl sets CustomPrivacyPolicyUrl field to given value.

### HasCustomPrivacyPolicyUrl

`func (o *Brand) HasCustomPrivacyPolicyUrl() bool`

HasCustomPrivacyPolicyUrl returns a boolean if a field has been set.

### GetDefaultApp

`func (o *Brand) GetDefaultApp() BrandDefaultApp`

GetDefaultApp returns the DefaultApp field if non-nil, zero value otherwise.

### GetDefaultAppOk

`func (o *Brand) GetDefaultAppOk() (*BrandDefaultApp, bool)`

GetDefaultAppOk returns a tuple with the DefaultApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApp

`func (o *Brand) SetDefaultApp(v BrandDefaultApp)`

SetDefaultApp sets DefaultApp field to given value.

### HasDefaultApp

`func (o *Brand) HasDefaultApp() bool`

HasDefaultApp returns a boolean if a field has been set.

### GetDisplayLanguage

`func (o *Brand) GetDisplayLanguage() string`

GetDisplayLanguage returns the DisplayLanguage field if non-nil, zero value otherwise.

### GetDisplayLanguageOk

`func (o *Brand) GetDisplayLanguageOk() (*string, bool)`

GetDisplayLanguageOk returns a tuple with the DisplayLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLanguage

`func (o *Brand) SetDisplayLanguage(v string)`

SetDisplayLanguage sets DisplayLanguage field to given value.

### HasDisplayLanguage

`func (o *Brand) HasDisplayLanguage() bool`

HasDisplayLanguage returns a boolean if a field has been set.

### GetId

`func (o *Brand) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Brand) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Brand) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Brand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOptOutOfUserCommunications

`func (o *Brand) GetOptOutOfUserCommunications() bool`

GetOptOutOfUserCommunications returns the OptOutOfUserCommunications field if non-nil, zero value otherwise.

### GetOptOutOfUserCommunicationsOk

`func (o *Brand) GetOptOutOfUserCommunicationsOk() (*bool, bool)`

GetOptOutOfUserCommunicationsOk returns a tuple with the OptOutOfUserCommunications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptOutOfUserCommunications

`func (o *Brand) SetOptOutOfUserCommunications(v bool)`

SetOptOutOfUserCommunications sets OptOutOfUserCommunications field to given value.

### HasOptOutOfUserCommunications

`func (o *Brand) HasOptOutOfUserCommunications() bool`

HasOptOutOfUserCommunications returns a boolean if a field has been set.

### GetRemovePoweredByOkta

`func (o *Brand) GetRemovePoweredByOkta() bool`

GetRemovePoweredByOkta returns the RemovePoweredByOkta field if non-nil, zero value otherwise.

### GetRemovePoweredByOktaOk

`func (o *Brand) GetRemovePoweredByOktaOk() (*bool, bool)`

GetRemovePoweredByOktaOk returns a tuple with the RemovePoweredByOkta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovePoweredByOkta

`func (o *Brand) SetRemovePoweredByOkta(v bool)`

SetRemovePoweredByOkta sets RemovePoweredByOkta field to given value.

### HasRemovePoweredByOkta

`func (o *Brand) HasRemovePoweredByOkta() bool`

HasRemovePoweredByOkta returns a boolean if a field has been set.

### GetLinks

`func (o *Brand) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Brand) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Brand) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Brand) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


