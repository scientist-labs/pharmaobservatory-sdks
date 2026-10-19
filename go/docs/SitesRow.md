# SitesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **NullableString** | The site key: one facility spelling from the registry; the id the site page address carries. | [optional] 
**SiteName** | Pointer to **NullableString** | The facility name as filed in the registry. | [optional] 
**City** | Pointer to **NullableString** | City as filed. | [optional] 
**State** | Pointer to **NullableString** | State or province as filed. | [optional] 
**Country** | Pointer to **NullableString** | Country name as filed. | [optional] 
**Latitude** | Pointer to **NullableFloat32** | Latitude of the finest coordinate the record has. | [optional] 
**Longitude** | Pointer to **NullableFloat32** | Longitude of the finest coordinate the record has. | [optional] 
**GeoPrecision** | Pointer to **NullableString** | How finely the coordinate places the site (city, postal and so on). | [optional] 
**CoordSource** | Pointer to **NullableString** | The source of the coordinate: aact (the registry&#39;s own), geonames (GeoNames) or als (Amazon Location Service). | [optional] 
**IsLocationSourceAssertedOnly** | Pointer to **NullableBool** | True when the city the registry gave is a location that only the registry asserts for the institution. | [optional] 
**SiteOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the site&#39;s institution resolved to (null when unresolved). | [optional] 
**SiteCompanyId** | Pointer to **NullableString** | Company level: the company the site&#39;s institution belongs to (null when unresolved). | [optional] 
**SiteCompanyName** | Pointer to **NullableString** | Display name of the site&#39;s institution&#39;s company. | [optional] 
**SiteGroupId** | Pointer to **NullableString** | Group level: the site&#39;s institution&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SiteGroupName** | Pointer to **NullableString** | Display name of the site&#39;s institution&#39;s group. | [optional] 
**NTrials** | Pointer to **NullableInt64** | Registered trials that list the site. | [optional] 
**NActive** | Pointer to **NullableInt64** | Those trials whose status is ongoing. | [optional] 
**FirstYear** | Pointer to **NullableInt32** | Year of the earliest trial start at the site. | [optional] 
**LastYear** | Pointer to **NullableInt32** | Year of the latest trial start at the site. | [optional] 

## Methods

### NewSitesRow

`func NewSitesRow() *SitesRow`

NewSitesRow instantiates a new SitesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSitesRowWithDefaults

`func NewSitesRowWithDefaults() *SitesRow`

NewSitesRowWithDefaults instantiates a new SitesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *SitesRow) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SitesRow) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SitesRow) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SitesRow) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *SitesRow) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *SitesRow) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetSiteName

`func (o *SitesRow) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *SitesRow) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *SitesRow) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *SitesRow) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### SetSiteNameNil

`func (o *SitesRow) SetSiteNameNil(b bool)`

 SetSiteNameNil sets the value for SiteName to be an explicit nil

### UnsetSiteName
`func (o *SitesRow) UnsetSiteName()`

UnsetSiteName ensures that no value is present for SiteName, not even an explicit nil
### GetCity

`func (o *SitesRow) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SitesRow) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SitesRow) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SitesRow) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *SitesRow) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *SitesRow) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *SitesRow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SitesRow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SitesRow) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SitesRow) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *SitesRow) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *SitesRow) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCountry

`func (o *SitesRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SitesRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SitesRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SitesRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *SitesRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *SitesRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLatitude

`func (o *SitesRow) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *SitesRow) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *SitesRow) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *SitesRow) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *SitesRow) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *SitesRow) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *SitesRow) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *SitesRow) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *SitesRow) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *SitesRow) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *SitesRow) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *SitesRow) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetGeoPrecision

`func (o *SitesRow) GetGeoPrecision() string`

GetGeoPrecision returns the GeoPrecision field if non-nil, zero value otherwise.

### GetGeoPrecisionOk

`func (o *SitesRow) GetGeoPrecisionOk() (*string, bool)`

GetGeoPrecisionOk returns a tuple with the GeoPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoPrecision

`func (o *SitesRow) SetGeoPrecision(v string)`

SetGeoPrecision sets GeoPrecision field to given value.

### HasGeoPrecision

`func (o *SitesRow) HasGeoPrecision() bool`

HasGeoPrecision returns a boolean if a field has been set.

### SetGeoPrecisionNil

`func (o *SitesRow) SetGeoPrecisionNil(b bool)`

 SetGeoPrecisionNil sets the value for GeoPrecision to be an explicit nil

### UnsetGeoPrecision
`func (o *SitesRow) UnsetGeoPrecision()`

UnsetGeoPrecision ensures that no value is present for GeoPrecision, not even an explicit nil
### GetCoordSource

`func (o *SitesRow) GetCoordSource() string`

GetCoordSource returns the CoordSource field if non-nil, zero value otherwise.

### GetCoordSourceOk

`func (o *SitesRow) GetCoordSourceOk() (*string, bool)`

GetCoordSourceOk returns a tuple with the CoordSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordSource

`func (o *SitesRow) SetCoordSource(v string)`

SetCoordSource sets CoordSource field to given value.

### HasCoordSource

`func (o *SitesRow) HasCoordSource() bool`

HasCoordSource returns a boolean if a field has been set.

### SetCoordSourceNil

`func (o *SitesRow) SetCoordSourceNil(b bool)`

 SetCoordSourceNil sets the value for CoordSource to be an explicit nil

### UnsetCoordSource
`func (o *SitesRow) UnsetCoordSource()`

UnsetCoordSource ensures that no value is present for CoordSource, not even an explicit nil
### GetIsLocationSourceAssertedOnly

`func (o *SitesRow) GetIsLocationSourceAssertedOnly() bool`

GetIsLocationSourceAssertedOnly returns the IsLocationSourceAssertedOnly field if non-nil, zero value otherwise.

### GetIsLocationSourceAssertedOnlyOk

`func (o *SitesRow) GetIsLocationSourceAssertedOnlyOk() (*bool, bool)`

GetIsLocationSourceAssertedOnlyOk returns a tuple with the IsLocationSourceAssertedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocationSourceAssertedOnly

`func (o *SitesRow) SetIsLocationSourceAssertedOnly(v bool)`

SetIsLocationSourceAssertedOnly sets IsLocationSourceAssertedOnly field to given value.

### HasIsLocationSourceAssertedOnly

`func (o *SitesRow) HasIsLocationSourceAssertedOnly() bool`

HasIsLocationSourceAssertedOnly returns a boolean if a field has been set.

### SetIsLocationSourceAssertedOnlyNil

`func (o *SitesRow) SetIsLocationSourceAssertedOnlyNil(b bool)`

 SetIsLocationSourceAssertedOnlyNil sets the value for IsLocationSourceAssertedOnly to be an explicit nil

### UnsetIsLocationSourceAssertedOnly
`func (o *SitesRow) UnsetIsLocationSourceAssertedOnly()`

UnsetIsLocationSourceAssertedOnly ensures that no value is present for IsLocationSourceAssertedOnly, not even an explicit nil
### GetSiteOrgId

`func (o *SitesRow) GetSiteOrgId() string`

GetSiteOrgId returns the SiteOrgId field if non-nil, zero value otherwise.

### GetSiteOrgIdOk

`func (o *SitesRow) GetSiteOrgIdOk() (*string, bool)`

GetSiteOrgIdOk returns a tuple with the SiteOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteOrgId

`func (o *SitesRow) SetSiteOrgId(v string)`

SetSiteOrgId sets SiteOrgId field to given value.

### HasSiteOrgId

`func (o *SitesRow) HasSiteOrgId() bool`

HasSiteOrgId returns a boolean if a field has been set.

### SetSiteOrgIdNil

`func (o *SitesRow) SetSiteOrgIdNil(b bool)`

 SetSiteOrgIdNil sets the value for SiteOrgId to be an explicit nil

### UnsetSiteOrgId
`func (o *SitesRow) UnsetSiteOrgId()`

UnsetSiteOrgId ensures that no value is present for SiteOrgId, not even an explicit nil
### GetSiteCompanyId

`func (o *SitesRow) GetSiteCompanyId() string`

GetSiteCompanyId returns the SiteCompanyId field if non-nil, zero value otherwise.

### GetSiteCompanyIdOk

`func (o *SitesRow) GetSiteCompanyIdOk() (*string, bool)`

GetSiteCompanyIdOk returns a tuple with the SiteCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCompanyId

`func (o *SitesRow) SetSiteCompanyId(v string)`

SetSiteCompanyId sets SiteCompanyId field to given value.

### HasSiteCompanyId

`func (o *SitesRow) HasSiteCompanyId() bool`

HasSiteCompanyId returns a boolean if a field has been set.

### SetSiteCompanyIdNil

`func (o *SitesRow) SetSiteCompanyIdNil(b bool)`

 SetSiteCompanyIdNil sets the value for SiteCompanyId to be an explicit nil

### UnsetSiteCompanyId
`func (o *SitesRow) UnsetSiteCompanyId()`

UnsetSiteCompanyId ensures that no value is present for SiteCompanyId, not even an explicit nil
### GetSiteCompanyName

`func (o *SitesRow) GetSiteCompanyName() string`

GetSiteCompanyName returns the SiteCompanyName field if non-nil, zero value otherwise.

### GetSiteCompanyNameOk

`func (o *SitesRow) GetSiteCompanyNameOk() (*string, bool)`

GetSiteCompanyNameOk returns a tuple with the SiteCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCompanyName

`func (o *SitesRow) SetSiteCompanyName(v string)`

SetSiteCompanyName sets SiteCompanyName field to given value.

### HasSiteCompanyName

`func (o *SitesRow) HasSiteCompanyName() bool`

HasSiteCompanyName returns a boolean if a field has been set.

### SetSiteCompanyNameNil

`func (o *SitesRow) SetSiteCompanyNameNil(b bool)`

 SetSiteCompanyNameNil sets the value for SiteCompanyName to be an explicit nil

### UnsetSiteCompanyName
`func (o *SitesRow) UnsetSiteCompanyName()`

UnsetSiteCompanyName ensures that no value is present for SiteCompanyName, not even an explicit nil
### GetSiteGroupId

`func (o *SitesRow) GetSiteGroupId() string`

GetSiteGroupId returns the SiteGroupId field if non-nil, zero value otherwise.

### GetSiteGroupIdOk

`func (o *SitesRow) GetSiteGroupIdOk() (*string, bool)`

GetSiteGroupIdOk returns a tuple with the SiteGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteGroupId

`func (o *SitesRow) SetSiteGroupId(v string)`

SetSiteGroupId sets SiteGroupId field to given value.

### HasSiteGroupId

`func (o *SitesRow) HasSiteGroupId() bool`

HasSiteGroupId returns a boolean if a field has been set.

### SetSiteGroupIdNil

`func (o *SitesRow) SetSiteGroupIdNil(b bool)`

 SetSiteGroupIdNil sets the value for SiteGroupId to be an explicit nil

### UnsetSiteGroupId
`func (o *SitesRow) UnsetSiteGroupId()`

UnsetSiteGroupId ensures that no value is present for SiteGroupId, not even an explicit nil
### GetSiteGroupName

`func (o *SitesRow) GetSiteGroupName() string`

GetSiteGroupName returns the SiteGroupName field if non-nil, zero value otherwise.

### GetSiteGroupNameOk

`func (o *SitesRow) GetSiteGroupNameOk() (*string, bool)`

GetSiteGroupNameOk returns a tuple with the SiteGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteGroupName

`func (o *SitesRow) SetSiteGroupName(v string)`

SetSiteGroupName sets SiteGroupName field to given value.

### HasSiteGroupName

`func (o *SitesRow) HasSiteGroupName() bool`

HasSiteGroupName returns a boolean if a field has been set.

### SetSiteGroupNameNil

`func (o *SitesRow) SetSiteGroupNameNil(b bool)`

 SetSiteGroupNameNil sets the value for SiteGroupName to be an explicit nil

### UnsetSiteGroupName
`func (o *SitesRow) UnsetSiteGroupName()`

UnsetSiteGroupName ensures that no value is present for SiteGroupName, not even an explicit nil
### GetNTrials

`func (o *SitesRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *SitesRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *SitesRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *SitesRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *SitesRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *SitesRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNActive

`func (o *SitesRow) GetNActive() int64`

GetNActive returns the NActive field if non-nil, zero value otherwise.

### GetNActiveOk

`func (o *SitesRow) GetNActiveOk() (*int64, bool)`

GetNActiveOk returns a tuple with the NActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNActive

`func (o *SitesRow) SetNActive(v int64)`

SetNActive sets NActive field to given value.

### HasNActive

`func (o *SitesRow) HasNActive() bool`

HasNActive returns a boolean if a field has been set.

### SetNActiveNil

`func (o *SitesRow) SetNActiveNil(b bool)`

 SetNActiveNil sets the value for NActive to be an explicit nil

### UnsetNActive
`func (o *SitesRow) UnsetNActive()`

UnsetNActive ensures that no value is present for NActive, not even an explicit nil
### GetFirstYear

`func (o *SitesRow) GetFirstYear() int32`

GetFirstYear returns the FirstYear field if non-nil, zero value otherwise.

### GetFirstYearOk

`func (o *SitesRow) GetFirstYearOk() (*int32, bool)`

GetFirstYearOk returns a tuple with the FirstYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstYear

`func (o *SitesRow) SetFirstYear(v int32)`

SetFirstYear sets FirstYear field to given value.

### HasFirstYear

`func (o *SitesRow) HasFirstYear() bool`

HasFirstYear returns a boolean if a field has been set.

### SetFirstYearNil

`func (o *SitesRow) SetFirstYearNil(b bool)`

 SetFirstYearNil sets the value for FirstYear to be an explicit nil

### UnsetFirstYear
`func (o *SitesRow) UnsetFirstYear()`

UnsetFirstYear ensures that no value is present for FirstYear, not even an explicit nil
### GetLastYear

`func (o *SitesRow) GetLastYear() int32`

GetLastYear returns the LastYear field if non-nil, zero value otherwise.

### GetLastYearOk

`func (o *SitesRow) GetLastYearOk() (*int32, bool)`

GetLastYearOk returns a tuple with the LastYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastYear

`func (o *SitesRow) SetLastYear(v int32)`

SetLastYear sets LastYear field to given value.

### HasLastYear

`func (o *SitesRow) HasLastYear() bool`

HasLastYear returns a boolean if a field has been set.

### SetLastYearNil

`func (o *SitesRow) SetLastYearNil(b bool)`

 SetLastYearNil sets the value for LastYear to be an explicit nil

### UnsetLastYear
`func (o *SitesRow) UnsetLastYear()`

UnsetLastYear ensures that no value is present for LastYear, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


