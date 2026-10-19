# TrialSitesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**SiteKey** | Pointer to **NullableString** | A site key built from the normalized facility name, city, state and country. | [optional] 
**ListingOrdinal** | Pointer to **NullableInt32** | 1 for the first listing of the site on the trial; higher for a repeat listing of the same site on the same trial. | [optional] 
**FacilityName** | Pointer to **NullableString** | The facility name as filed. | [optional] 
**City** | Pointer to **NullableString** | City. | [optional] 
**State** | Pointer to **NullableString** | State or region. | [optional] 
**Country** | Pointer to **NullableString** | Country. | [optional] 
**Postal** | Pointer to **NullableString** | Postal code. | [optional] 
**Latitude** | Pointer to **NullableFloat64** | Latitude of the site, in degrees. | [optional] 
**Longitude** | Pointer to **NullableFloat64** | Longitude of the site, in degrees. | [optional] 
**GeoPrecision** | Pointer to **NullableString** | How exact the coordinate is (postal, city and so on). | [optional] 
**CoordSource** | Pointer to **NullableString** | Where the coordinate came from: aact, geonames or als. | [optional] 
**TrialStatus** | Pointer to **NullableString** | The trial&#39;s overall status code. | [optional] 
**TrialStartYear** | Pointer to **NullableInt32** | Year the trial started. | [optional] 
**TrialTitle** | Pointer to **NullableString** | The trial&#39;s brief title. | [optional] 

## Methods

### NewTrialSitesRow

`func NewTrialSitesRow() *TrialSitesRow`

NewTrialSitesRow instantiates a new TrialSitesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialSitesRowWithDefaults

`func NewTrialSitesRowWithDefaults() *TrialSitesRow`

NewTrialSitesRowWithDefaults instantiates a new TrialSitesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *TrialSitesRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TrialSitesRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TrialSitesRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TrialSitesRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TrialSitesRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TrialSitesRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetSiteKey

`func (o *TrialSitesRow) GetSiteKey() string`

GetSiteKey returns the SiteKey field if non-nil, zero value otherwise.

### GetSiteKeyOk

`func (o *TrialSitesRow) GetSiteKeyOk() (*string, bool)`

GetSiteKeyOk returns a tuple with the SiteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteKey

`func (o *TrialSitesRow) SetSiteKey(v string)`

SetSiteKey sets SiteKey field to given value.

### HasSiteKey

`func (o *TrialSitesRow) HasSiteKey() bool`

HasSiteKey returns a boolean if a field has been set.

### SetSiteKeyNil

`func (o *TrialSitesRow) SetSiteKeyNil(b bool)`

 SetSiteKeyNil sets the value for SiteKey to be an explicit nil

### UnsetSiteKey
`func (o *TrialSitesRow) UnsetSiteKey()`

UnsetSiteKey ensures that no value is present for SiteKey, not even an explicit nil
### GetListingOrdinal

`func (o *TrialSitesRow) GetListingOrdinal() int32`

GetListingOrdinal returns the ListingOrdinal field if non-nil, zero value otherwise.

### GetListingOrdinalOk

`func (o *TrialSitesRow) GetListingOrdinalOk() (*int32, bool)`

GetListingOrdinalOk returns a tuple with the ListingOrdinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingOrdinal

`func (o *TrialSitesRow) SetListingOrdinal(v int32)`

SetListingOrdinal sets ListingOrdinal field to given value.

### HasListingOrdinal

`func (o *TrialSitesRow) HasListingOrdinal() bool`

HasListingOrdinal returns a boolean if a field has been set.

### SetListingOrdinalNil

`func (o *TrialSitesRow) SetListingOrdinalNil(b bool)`

 SetListingOrdinalNil sets the value for ListingOrdinal to be an explicit nil

### UnsetListingOrdinal
`func (o *TrialSitesRow) UnsetListingOrdinal()`

UnsetListingOrdinal ensures that no value is present for ListingOrdinal, not even an explicit nil
### GetFacilityName

`func (o *TrialSitesRow) GetFacilityName() string`

GetFacilityName returns the FacilityName field if non-nil, zero value otherwise.

### GetFacilityNameOk

`func (o *TrialSitesRow) GetFacilityNameOk() (*string, bool)`

GetFacilityNameOk returns a tuple with the FacilityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityName

`func (o *TrialSitesRow) SetFacilityName(v string)`

SetFacilityName sets FacilityName field to given value.

### HasFacilityName

`func (o *TrialSitesRow) HasFacilityName() bool`

HasFacilityName returns a boolean if a field has been set.

### SetFacilityNameNil

`func (o *TrialSitesRow) SetFacilityNameNil(b bool)`

 SetFacilityNameNil sets the value for FacilityName to be an explicit nil

### UnsetFacilityName
`func (o *TrialSitesRow) UnsetFacilityName()`

UnsetFacilityName ensures that no value is present for FacilityName, not even an explicit nil
### GetCity

`func (o *TrialSitesRow) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *TrialSitesRow) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *TrialSitesRow) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *TrialSitesRow) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *TrialSitesRow) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *TrialSitesRow) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *TrialSitesRow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TrialSitesRow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TrialSitesRow) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TrialSitesRow) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *TrialSitesRow) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *TrialSitesRow) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCountry

`func (o *TrialSitesRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TrialSitesRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TrialSitesRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TrialSitesRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *TrialSitesRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *TrialSitesRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetPostal

`func (o *TrialSitesRow) GetPostal() string`

GetPostal returns the Postal field if non-nil, zero value otherwise.

### GetPostalOk

`func (o *TrialSitesRow) GetPostalOk() (*string, bool)`

GetPostalOk returns a tuple with the Postal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostal

`func (o *TrialSitesRow) SetPostal(v string)`

SetPostal sets Postal field to given value.

### HasPostal

`func (o *TrialSitesRow) HasPostal() bool`

HasPostal returns a boolean if a field has been set.

### SetPostalNil

`func (o *TrialSitesRow) SetPostalNil(b bool)`

 SetPostalNil sets the value for Postal to be an explicit nil

### UnsetPostal
`func (o *TrialSitesRow) UnsetPostal()`

UnsetPostal ensures that no value is present for Postal, not even an explicit nil
### GetLatitude

`func (o *TrialSitesRow) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *TrialSitesRow) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *TrialSitesRow) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *TrialSitesRow) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *TrialSitesRow) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *TrialSitesRow) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *TrialSitesRow) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *TrialSitesRow) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *TrialSitesRow) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *TrialSitesRow) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *TrialSitesRow) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *TrialSitesRow) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetGeoPrecision

`func (o *TrialSitesRow) GetGeoPrecision() string`

GetGeoPrecision returns the GeoPrecision field if non-nil, zero value otherwise.

### GetGeoPrecisionOk

`func (o *TrialSitesRow) GetGeoPrecisionOk() (*string, bool)`

GetGeoPrecisionOk returns a tuple with the GeoPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoPrecision

`func (o *TrialSitesRow) SetGeoPrecision(v string)`

SetGeoPrecision sets GeoPrecision field to given value.

### HasGeoPrecision

`func (o *TrialSitesRow) HasGeoPrecision() bool`

HasGeoPrecision returns a boolean if a field has been set.

### SetGeoPrecisionNil

`func (o *TrialSitesRow) SetGeoPrecisionNil(b bool)`

 SetGeoPrecisionNil sets the value for GeoPrecision to be an explicit nil

### UnsetGeoPrecision
`func (o *TrialSitesRow) UnsetGeoPrecision()`

UnsetGeoPrecision ensures that no value is present for GeoPrecision, not even an explicit nil
### GetCoordSource

`func (o *TrialSitesRow) GetCoordSource() string`

GetCoordSource returns the CoordSource field if non-nil, zero value otherwise.

### GetCoordSourceOk

`func (o *TrialSitesRow) GetCoordSourceOk() (*string, bool)`

GetCoordSourceOk returns a tuple with the CoordSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordSource

`func (o *TrialSitesRow) SetCoordSource(v string)`

SetCoordSource sets CoordSource field to given value.

### HasCoordSource

`func (o *TrialSitesRow) HasCoordSource() bool`

HasCoordSource returns a boolean if a field has been set.

### SetCoordSourceNil

`func (o *TrialSitesRow) SetCoordSourceNil(b bool)`

 SetCoordSourceNil sets the value for CoordSource to be an explicit nil

### UnsetCoordSource
`func (o *TrialSitesRow) UnsetCoordSource()`

UnsetCoordSource ensures that no value is present for CoordSource, not even an explicit nil
### GetTrialStatus

`func (o *TrialSitesRow) GetTrialStatus() string`

GetTrialStatus returns the TrialStatus field if non-nil, zero value otherwise.

### GetTrialStatusOk

`func (o *TrialSitesRow) GetTrialStatusOk() (*string, bool)`

GetTrialStatusOk returns a tuple with the TrialStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStatus

`func (o *TrialSitesRow) SetTrialStatus(v string)`

SetTrialStatus sets TrialStatus field to given value.

### HasTrialStatus

`func (o *TrialSitesRow) HasTrialStatus() bool`

HasTrialStatus returns a boolean if a field has been set.

### SetTrialStatusNil

`func (o *TrialSitesRow) SetTrialStatusNil(b bool)`

 SetTrialStatusNil sets the value for TrialStatus to be an explicit nil

### UnsetTrialStatus
`func (o *TrialSitesRow) UnsetTrialStatus()`

UnsetTrialStatus ensures that no value is present for TrialStatus, not even an explicit nil
### GetTrialStartYear

`func (o *TrialSitesRow) GetTrialStartYear() int32`

GetTrialStartYear returns the TrialStartYear field if non-nil, zero value otherwise.

### GetTrialStartYearOk

`func (o *TrialSitesRow) GetTrialStartYearOk() (*int32, bool)`

GetTrialStartYearOk returns a tuple with the TrialStartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStartYear

`func (o *TrialSitesRow) SetTrialStartYear(v int32)`

SetTrialStartYear sets TrialStartYear field to given value.

### HasTrialStartYear

`func (o *TrialSitesRow) HasTrialStartYear() bool`

HasTrialStartYear returns a boolean if a field has been set.

### SetTrialStartYearNil

`func (o *TrialSitesRow) SetTrialStartYearNil(b bool)`

 SetTrialStartYearNil sets the value for TrialStartYear to be an explicit nil

### UnsetTrialStartYear
`func (o *TrialSitesRow) UnsetTrialStartYear()`

UnsetTrialStartYear ensures that no value is present for TrialStartYear, not even an explicit nil
### GetTrialTitle

`func (o *TrialSitesRow) GetTrialTitle() string`

GetTrialTitle returns the TrialTitle field if non-nil, zero value otherwise.

### GetTrialTitleOk

`func (o *TrialSitesRow) GetTrialTitleOk() (*string, bool)`

GetTrialTitleOk returns a tuple with the TrialTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialTitle

`func (o *TrialSitesRow) SetTrialTitle(v string)`

SetTrialTitle sets TrialTitle field to given value.

### HasTrialTitle

`func (o *TrialSitesRow) HasTrialTitle() bool`

HasTrialTitle returns a boolean if a field has been set.

### SetTrialTitleNil

`func (o *TrialSitesRow) SetTrialTitleNil(b bool)`

 SetTrialTitleNil sets the value for TrialTitle to be an explicit nil

### UnsetTrialTitle
`func (o *TrialSitesRow) UnsetTrialTitle()`

UnsetTrialTitle ensures that no value is present for TrialTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


