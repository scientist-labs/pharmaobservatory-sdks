# CroFootprintRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorCompanyId** | Pointer to **NullableString** | Company level: the contract research organization&#39;s company. | [optional] 
**VendorCompanyName** | Pointer to **NullableString** | Display name of the contract research organization&#39;s company. | [optional] 
**VendorGroupId** | Pointer to **NullableString** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**VendorGroupName** | Pointer to **NullableString** | Display name of the company&#39;s group. | [optional] 
**Country** | Pointer to **NullableString** | The country name as the registry&#39;s facility roster spells it. | [optional] 
**NTrials** | Pointer to **NullableInt32** | Distinct trials the company delivered with a facility in the country. | [optional] 
**NSites** | Pointer to **NullableInt32** | Distinct sites (facility keys) in the country across those trials. | [optional] 
**Latitude** | Pointer to **NullableFloat32** | Latitude of the country&#39;s approximate centroid. Null for a country with no centroid. | [optional] 
**Longitude** | Pointer to **NullableFloat32** | Longitude of the country&#39;s approximate centroid. Null for a country with no centroid. | [optional] 

## Methods

### NewCroFootprintRow

`func NewCroFootprintRow() *CroFootprintRow`

NewCroFootprintRow instantiates a new CroFootprintRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCroFootprintRowWithDefaults

`func NewCroFootprintRowWithDefaults() *CroFootprintRow`

NewCroFootprintRowWithDefaults instantiates a new CroFootprintRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendorCompanyId

`func (o *CroFootprintRow) GetVendorCompanyId() string`

GetVendorCompanyId returns the VendorCompanyId field if non-nil, zero value otherwise.

### GetVendorCompanyIdOk

`func (o *CroFootprintRow) GetVendorCompanyIdOk() (*string, bool)`

GetVendorCompanyIdOk returns a tuple with the VendorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCompanyId

`func (o *CroFootprintRow) SetVendorCompanyId(v string)`

SetVendorCompanyId sets VendorCompanyId field to given value.

### HasVendorCompanyId

`func (o *CroFootprintRow) HasVendorCompanyId() bool`

HasVendorCompanyId returns a boolean if a field has been set.

### SetVendorCompanyIdNil

`func (o *CroFootprintRow) SetVendorCompanyIdNil(b bool)`

 SetVendorCompanyIdNil sets the value for VendorCompanyId to be an explicit nil

### UnsetVendorCompanyId
`func (o *CroFootprintRow) UnsetVendorCompanyId()`

UnsetVendorCompanyId ensures that no value is present for VendorCompanyId, not even an explicit nil
### GetVendorCompanyName

`func (o *CroFootprintRow) GetVendorCompanyName() string`

GetVendorCompanyName returns the VendorCompanyName field if non-nil, zero value otherwise.

### GetVendorCompanyNameOk

`func (o *CroFootprintRow) GetVendorCompanyNameOk() (*string, bool)`

GetVendorCompanyNameOk returns a tuple with the VendorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCompanyName

`func (o *CroFootprintRow) SetVendorCompanyName(v string)`

SetVendorCompanyName sets VendorCompanyName field to given value.

### HasVendorCompanyName

`func (o *CroFootprintRow) HasVendorCompanyName() bool`

HasVendorCompanyName returns a boolean if a field has been set.

### SetVendorCompanyNameNil

`func (o *CroFootprintRow) SetVendorCompanyNameNil(b bool)`

 SetVendorCompanyNameNil sets the value for VendorCompanyName to be an explicit nil

### UnsetVendorCompanyName
`func (o *CroFootprintRow) UnsetVendorCompanyName()`

UnsetVendorCompanyName ensures that no value is present for VendorCompanyName, not even an explicit nil
### GetVendorGroupId

`func (o *CroFootprintRow) GetVendorGroupId() string`

GetVendorGroupId returns the VendorGroupId field if non-nil, zero value otherwise.

### GetVendorGroupIdOk

`func (o *CroFootprintRow) GetVendorGroupIdOk() (*string, bool)`

GetVendorGroupIdOk returns a tuple with the VendorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorGroupId

`func (o *CroFootprintRow) SetVendorGroupId(v string)`

SetVendorGroupId sets VendorGroupId field to given value.

### HasVendorGroupId

`func (o *CroFootprintRow) HasVendorGroupId() bool`

HasVendorGroupId returns a boolean if a field has been set.

### SetVendorGroupIdNil

`func (o *CroFootprintRow) SetVendorGroupIdNil(b bool)`

 SetVendorGroupIdNil sets the value for VendorGroupId to be an explicit nil

### UnsetVendorGroupId
`func (o *CroFootprintRow) UnsetVendorGroupId()`

UnsetVendorGroupId ensures that no value is present for VendorGroupId, not even an explicit nil
### GetVendorGroupName

`func (o *CroFootprintRow) GetVendorGroupName() string`

GetVendorGroupName returns the VendorGroupName field if non-nil, zero value otherwise.

### GetVendorGroupNameOk

`func (o *CroFootprintRow) GetVendorGroupNameOk() (*string, bool)`

GetVendorGroupNameOk returns a tuple with the VendorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorGroupName

`func (o *CroFootprintRow) SetVendorGroupName(v string)`

SetVendorGroupName sets VendorGroupName field to given value.

### HasVendorGroupName

`func (o *CroFootprintRow) HasVendorGroupName() bool`

HasVendorGroupName returns a boolean if a field has been set.

### SetVendorGroupNameNil

`func (o *CroFootprintRow) SetVendorGroupNameNil(b bool)`

 SetVendorGroupNameNil sets the value for VendorGroupName to be an explicit nil

### UnsetVendorGroupName
`func (o *CroFootprintRow) UnsetVendorGroupName()`

UnsetVendorGroupName ensures that no value is present for VendorGroupName, not even an explicit nil
### GetCountry

`func (o *CroFootprintRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CroFootprintRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CroFootprintRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CroFootprintRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *CroFootprintRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *CroFootprintRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetNTrials

`func (o *CroFootprintRow) GetNTrials() int32`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *CroFootprintRow) GetNTrialsOk() (*int32, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *CroFootprintRow) SetNTrials(v int32)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *CroFootprintRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *CroFootprintRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *CroFootprintRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNSites

`func (o *CroFootprintRow) GetNSites() int32`

GetNSites returns the NSites field if non-nil, zero value otherwise.

### GetNSitesOk

`func (o *CroFootprintRow) GetNSitesOk() (*int32, bool)`

GetNSitesOk returns a tuple with the NSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSites

`func (o *CroFootprintRow) SetNSites(v int32)`

SetNSites sets NSites field to given value.

### HasNSites

`func (o *CroFootprintRow) HasNSites() bool`

HasNSites returns a boolean if a field has been set.

### SetNSitesNil

`func (o *CroFootprintRow) SetNSitesNil(b bool)`

 SetNSitesNil sets the value for NSites to be an explicit nil

### UnsetNSites
`func (o *CroFootprintRow) UnsetNSites()`

UnsetNSites ensures that no value is present for NSites, not even an explicit nil
### GetLatitude

`func (o *CroFootprintRow) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *CroFootprintRow) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *CroFootprintRow) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *CroFootprintRow) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *CroFootprintRow) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *CroFootprintRow) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *CroFootprintRow) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *CroFootprintRow) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *CroFootprintRow) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *CroFootprintRow) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *CroFootprintRow) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *CroFootprintRow) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


