# ConditionSiteYearsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**SiteId** | Pointer to **NullableString** | The site key (a normalized facility, city and country). | [optional] 
**StartYear** | Pointer to **NullableInt32** | Year the trials started. Null for a trial with no start date. | [optional] 
**SiteName** | Pointer to **NullableString** | The facility name as the registry lists it. | [optional] 
**City** | Pointer to **NullableString** | City as listed. | [optional] 
**State** | Pointer to **NullableString** | State or province as listed. | [optional] 
**Country** | Pointer to **NullableString** | Country as listed. | [optional] 
**Latitude** | Pointer to **NullableFloat32** | Latitude of the finest coordinate the site record has. | [optional] 
**Longitude** | Pointer to **NullableFloat32** | Longitude of the finest coordinate the site record has. | [optional] 
**NTrials** | Pointer to **NullableInt32** | Trials of the condition that started that year and list the site. | [optional] 

## Methods

### NewConditionSiteYearsRow

`func NewConditionSiteYearsRow() *ConditionSiteYearsRow`

NewConditionSiteYearsRow instantiates a new ConditionSiteYearsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionSiteYearsRowWithDefaults

`func NewConditionSiteYearsRowWithDefaults() *ConditionSiteYearsRow`

NewConditionSiteYearsRowWithDefaults instantiates a new ConditionSiteYearsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ConditionSiteYearsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ConditionSiteYearsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ConditionSiteYearsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ConditionSiteYearsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ConditionSiteYearsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ConditionSiteYearsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetSiteId

`func (o *ConditionSiteYearsRow) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ConditionSiteYearsRow) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ConditionSiteYearsRow) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ConditionSiteYearsRow) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *ConditionSiteYearsRow) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *ConditionSiteYearsRow) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetStartYear

`func (o *ConditionSiteYearsRow) GetStartYear() int32`

GetStartYear returns the StartYear field if non-nil, zero value otherwise.

### GetStartYearOk

`func (o *ConditionSiteYearsRow) GetStartYearOk() (*int32, bool)`

GetStartYearOk returns a tuple with the StartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartYear

`func (o *ConditionSiteYearsRow) SetStartYear(v int32)`

SetStartYear sets StartYear field to given value.

### HasStartYear

`func (o *ConditionSiteYearsRow) HasStartYear() bool`

HasStartYear returns a boolean if a field has been set.

### SetStartYearNil

`func (o *ConditionSiteYearsRow) SetStartYearNil(b bool)`

 SetStartYearNil sets the value for StartYear to be an explicit nil

### UnsetStartYear
`func (o *ConditionSiteYearsRow) UnsetStartYear()`

UnsetStartYear ensures that no value is present for StartYear, not even an explicit nil
### GetSiteName

`func (o *ConditionSiteYearsRow) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *ConditionSiteYearsRow) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *ConditionSiteYearsRow) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *ConditionSiteYearsRow) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### SetSiteNameNil

`func (o *ConditionSiteYearsRow) SetSiteNameNil(b bool)`

 SetSiteNameNil sets the value for SiteName to be an explicit nil

### UnsetSiteName
`func (o *ConditionSiteYearsRow) UnsetSiteName()`

UnsetSiteName ensures that no value is present for SiteName, not even an explicit nil
### GetCity

`func (o *ConditionSiteYearsRow) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ConditionSiteYearsRow) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ConditionSiteYearsRow) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ConditionSiteYearsRow) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ConditionSiteYearsRow) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ConditionSiteYearsRow) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *ConditionSiteYearsRow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConditionSiteYearsRow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConditionSiteYearsRow) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ConditionSiteYearsRow) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ConditionSiteYearsRow) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ConditionSiteYearsRow) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCountry

`func (o *ConditionSiteYearsRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ConditionSiteYearsRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ConditionSiteYearsRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ConditionSiteYearsRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *ConditionSiteYearsRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *ConditionSiteYearsRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLatitude

`func (o *ConditionSiteYearsRow) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *ConditionSiteYearsRow) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *ConditionSiteYearsRow) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *ConditionSiteYearsRow) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *ConditionSiteYearsRow) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *ConditionSiteYearsRow) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *ConditionSiteYearsRow) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *ConditionSiteYearsRow) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *ConditionSiteYearsRow) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *ConditionSiteYearsRow) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *ConditionSiteYearsRow) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *ConditionSiteYearsRow) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetNTrials

`func (o *ConditionSiteYearsRow) GetNTrials() int32`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *ConditionSiteYearsRow) GetNTrialsOk() (*int32, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *ConditionSiteYearsRow) SetNTrials(v int32)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *ConditionSiteYearsRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *ConditionSiteYearsRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *ConditionSiteYearsRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


