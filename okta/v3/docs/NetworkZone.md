# NetworkZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asns** | Pointer to **[]string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Gateways** | Pointer to [**[]NetworkZoneAddress**](NetworkZoneAddress.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Locations** | Pointer to [**[]NetworkZoneLocation**](NetworkZoneLocation.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Proxies** | Pointer to [**[]NetworkZoneAddress**](NetworkZoneAddress.md) |  | [optional] 
**ProxyType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**System** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to **string** |  | [optional] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewNetworkZone

`func NewNetworkZone() *NetworkZone`

NewNetworkZone instantiates a new NetworkZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkZoneWithDefaults

`func NewNetworkZoneWithDefaults() *NetworkZone`

NewNetworkZoneWithDefaults instantiates a new NetworkZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsns

`func (o *NetworkZone) GetAsns() []string`

GetAsns returns the Asns field if non-nil, zero value otherwise.

### GetAsnsOk

`func (o *NetworkZone) GetAsnsOk() (*[]string, bool)`

GetAsnsOk returns a tuple with the Asns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsns

`func (o *NetworkZone) SetAsns(v []string)`

SetAsns sets Asns field to given value.

### HasAsns

`func (o *NetworkZone) HasAsns() bool`

HasAsns returns a boolean if a field has been set.

### GetCreated

`func (o *NetworkZone) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NetworkZone) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NetworkZone) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *NetworkZone) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetGateways

`func (o *NetworkZone) GetGateways() []NetworkZoneAddress`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *NetworkZone) GetGatewaysOk() (*[]NetworkZoneAddress, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *NetworkZone) SetGateways(v []NetworkZoneAddress)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *NetworkZone) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetId

`func (o *NetworkZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NetworkZone) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NetworkZone) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NetworkZone) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NetworkZone) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLocations

`func (o *NetworkZone) GetLocations() []NetworkZoneLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *NetworkZone) GetLocationsOk() (*[]NetworkZoneLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *NetworkZone) SetLocations(v []NetworkZoneLocation)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *NetworkZone) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetName

`func (o *NetworkZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProxies

`func (o *NetworkZone) GetProxies() []NetworkZoneAddress`

GetProxies returns the Proxies field if non-nil, zero value otherwise.

### GetProxiesOk

`func (o *NetworkZone) GetProxiesOk() (*[]NetworkZoneAddress, bool)`

GetProxiesOk returns a tuple with the Proxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxies

`func (o *NetworkZone) SetProxies(v []NetworkZoneAddress)`

SetProxies sets Proxies field to given value.

### HasProxies

`func (o *NetworkZone) HasProxies() bool`

HasProxies returns a boolean if a field has been set.

### GetProxyType

`func (o *NetworkZone) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *NetworkZone) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *NetworkZone) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *NetworkZone) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkZone) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkZone) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkZone) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkZone) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystem

`func (o *NetworkZone) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *NetworkZone) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *NetworkZone) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *NetworkZone) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetType

`func (o *NetworkZone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkZone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkZone) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkZone) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsage

`func (o *NetworkZone) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *NetworkZone) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *NetworkZone) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *NetworkZone) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetLinks

`func (o *NetworkZone) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkZone) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkZone) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkZone) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


