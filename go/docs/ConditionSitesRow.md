# ConditionSitesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**SiteId** | Pointer to **NullableString** | The site key (a normalized facility, city and country). | [optional] 
**SiteName** | Pointer to **NullableString** | The facility name as the registry lists it. | [optional] 
**City** | Pointer to **NullableString** | City as listed. | [optional] 
**Country** | Pointer to **NullableString** | Country as listed. | [optional] 
**NTrials** | Pointer to **NullableInt64** | Trials for this condition listed at the site. | [optional] 

## Methods

### NewConditionSitesRow

`func NewConditionSitesRow() *ConditionSitesRow`

NewConditionSitesRow instantiates a new ConditionSitesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionSitesRowWithDefaults

`func NewConditionSitesRowWithDefaults() *ConditionSitesRow`

NewConditionSitesRowWithDefaults instantiates a new ConditionSitesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ConditionSitesRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ConditionSitesRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ConditionSitesRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ConditionSitesRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ConditionSitesRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ConditionSitesRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetSiteId

`func (o *ConditionSitesRow) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ConditionSitesRow) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ConditionSitesRow) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ConditionSitesRow) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *ConditionSitesRow) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *ConditionSitesRow) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetSiteName

`func (o *ConditionSitesRow) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *ConditionSitesRow) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *ConditionSitesRow) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *ConditionSitesRow) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### SetSiteNameNil

`func (o *ConditionSitesRow) SetSiteNameNil(b bool)`

 SetSiteNameNil sets the value for SiteName to be an explicit nil

### UnsetSiteName
`func (o *ConditionSitesRow) UnsetSiteName()`

UnsetSiteName ensures that no value is present for SiteName, not even an explicit nil
### GetCity

`func (o *ConditionSitesRow) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ConditionSitesRow) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ConditionSitesRow) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ConditionSitesRow) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ConditionSitesRow) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ConditionSitesRow) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *ConditionSitesRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ConditionSitesRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ConditionSitesRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ConditionSitesRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *ConditionSitesRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *ConditionSitesRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetNTrials

`func (o *ConditionSitesRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *ConditionSitesRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *ConditionSitesRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *ConditionSitesRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *ConditionSitesRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *ConditionSitesRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


