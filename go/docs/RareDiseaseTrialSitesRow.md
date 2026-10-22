# RareDiseaseTrialSitesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrphaCode** | Pointer to **NullableString** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] 
**MappingArm** | Pointer to **NullableString** | exact or unspecified: the mapping arm the site&#39;s trials belong to. | [optional] 
**SiteId** | Pointer to **NullableString** | The site key: one facility spelling from the registry; joins the sites dataset. | [optional] 
**SiteRank** | Pointer to **NullableInt32** | The site&#39;s rank by trial count within its disease and arm (1 is the most trials); a position in a list, not a quality score. | [optional] 
**SiteName** | Pointer to **NullableString** | The facility name as filed in the registry. | [optional] 
**City** | Pointer to **NullableString** | The facility&#39;s city as filed. | [optional] 
**State** | Pointer to **NullableString** | The facility&#39;s state or region as filed. | [optional] 
**Country** | Pointer to **NullableString** | The facility&#39;s country as filed. | [optional] 
**Latitude** | Pointer to **NullableFloat32** | The site&#39;s latitude in degrees; null when no coordinate is known. | [optional] 
**Longitude** | Pointer to **NullableFloat32** | The site&#39;s longitude in degrees; null when no coordinate is known. | [optional] 
**CoordSource** | Pointer to **NullableString** | The source of the coordinate: aact (the registry&#39;s own), geonames (GeoNames) or als (Amazon Location Service). | [optional] 
**ConditionIds** | Pointer to **[]string** | The MeSH condition ids through which the disease reaches the site&#39;s trials. | [optional] 
**NTrials** | Pointer to **NullableInt32** | Registered trials of this disease and arm that list the site. | [optional] 
**NTrialsDirect** | Pointer to **NullableInt32** | Those trials whose own condition list names the mapped condition. | [optional] 
**NActiveTrials** | Pointer to **NullableInt32** | Those trials that are recruiting, not yet recruiting, active or enrolling by invitation. | [optional] 
**FirstTrialYear** | Pointer to **NullableInt32** | The earliest start year among those trials. | [optional] 
**LastTrialYear** | Pointer to **NullableInt32** | The latest start year among those trials. | [optional] 
**LastUpdatePostedDate** | Pointer to **NullableString** | The latest registry update among those trials. | [optional] 

## Methods

### NewRareDiseaseTrialSitesRow

`func NewRareDiseaseTrialSitesRow() *RareDiseaseTrialSitesRow`

NewRareDiseaseTrialSitesRow instantiates a new RareDiseaseTrialSitesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRareDiseaseTrialSitesRowWithDefaults

`func NewRareDiseaseTrialSitesRowWithDefaults() *RareDiseaseTrialSitesRow`

NewRareDiseaseTrialSitesRowWithDefaults instantiates a new RareDiseaseTrialSitesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrphaCode

`func (o *RareDiseaseTrialSitesRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *RareDiseaseTrialSitesRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *RareDiseaseTrialSitesRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *RareDiseaseTrialSitesRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *RareDiseaseTrialSitesRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *RareDiseaseTrialSitesRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetMappingArm

`func (o *RareDiseaseTrialSitesRow) GetMappingArm() string`

GetMappingArm returns the MappingArm field if non-nil, zero value otherwise.

### GetMappingArmOk

`func (o *RareDiseaseTrialSitesRow) GetMappingArmOk() (*string, bool)`

GetMappingArmOk returns a tuple with the MappingArm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingArm

`func (o *RareDiseaseTrialSitesRow) SetMappingArm(v string)`

SetMappingArm sets MappingArm field to given value.

### HasMappingArm

`func (o *RareDiseaseTrialSitesRow) HasMappingArm() bool`

HasMappingArm returns a boolean if a field has been set.

### SetMappingArmNil

`func (o *RareDiseaseTrialSitesRow) SetMappingArmNil(b bool)`

 SetMappingArmNil sets the value for MappingArm to be an explicit nil

### UnsetMappingArm
`func (o *RareDiseaseTrialSitesRow) UnsetMappingArm()`

UnsetMappingArm ensures that no value is present for MappingArm, not even an explicit nil
### GetSiteId

`func (o *RareDiseaseTrialSitesRow) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *RareDiseaseTrialSitesRow) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *RareDiseaseTrialSitesRow) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *RareDiseaseTrialSitesRow) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *RareDiseaseTrialSitesRow) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *RareDiseaseTrialSitesRow) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetSiteRank

`func (o *RareDiseaseTrialSitesRow) GetSiteRank() int32`

GetSiteRank returns the SiteRank field if non-nil, zero value otherwise.

### GetSiteRankOk

`func (o *RareDiseaseTrialSitesRow) GetSiteRankOk() (*int32, bool)`

GetSiteRankOk returns a tuple with the SiteRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteRank

`func (o *RareDiseaseTrialSitesRow) SetSiteRank(v int32)`

SetSiteRank sets SiteRank field to given value.

### HasSiteRank

`func (o *RareDiseaseTrialSitesRow) HasSiteRank() bool`

HasSiteRank returns a boolean if a field has been set.

### SetSiteRankNil

`func (o *RareDiseaseTrialSitesRow) SetSiteRankNil(b bool)`

 SetSiteRankNil sets the value for SiteRank to be an explicit nil

### UnsetSiteRank
`func (o *RareDiseaseTrialSitesRow) UnsetSiteRank()`

UnsetSiteRank ensures that no value is present for SiteRank, not even an explicit nil
### GetSiteName

`func (o *RareDiseaseTrialSitesRow) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *RareDiseaseTrialSitesRow) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *RareDiseaseTrialSitesRow) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *RareDiseaseTrialSitesRow) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### SetSiteNameNil

`func (o *RareDiseaseTrialSitesRow) SetSiteNameNil(b bool)`

 SetSiteNameNil sets the value for SiteName to be an explicit nil

### UnsetSiteName
`func (o *RareDiseaseTrialSitesRow) UnsetSiteName()`

UnsetSiteName ensures that no value is present for SiteName, not even an explicit nil
### GetCity

`func (o *RareDiseaseTrialSitesRow) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *RareDiseaseTrialSitesRow) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *RareDiseaseTrialSitesRow) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *RareDiseaseTrialSitesRow) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *RareDiseaseTrialSitesRow) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *RareDiseaseTrialSitesRow) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *RareDiseaseTrialSitesRow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RareDiseaseTrialSitesRow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RareDiseaseTrialSitesRow) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RareDiseaseTrialSitesRow) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *RareDiseaseTrialSitesRow) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *RareDiseaseTrialSitesRow) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCountry

`func (o *RareDiseaseTrialSitesRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RareDiseaseTrialSitesRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RareDiseaseTrialSitesRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RareDiseaseTrialSitesRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *RareDiseaseTrialSitesRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *RareDiseaseTrialSitesRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLatitude

`func (o *RareDiseaseTrialSitesRow) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *RareDiseaseTrialSitesRow) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *RareDiseaseTrialSitesRow) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *RareDiseaseTrialSitesRow) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *RareDiseaseTrialSitesRow) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *RareDiseaseTrialSitesRow) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *RareDiseaseTrialSitesRow) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *RareDiseaseTrialSitesRow) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *RareDiseaseTrialSitesRow) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *RareDiseaseTrialSitesRow) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *RareDiseaseTrialSitesRow) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *RareDiseaseTrialSitesRow) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetCoordSource

`func (o *RareDiseaseTrialSitesRow) GetCoordSource() string`

GetCoordSource returns the CoordSource field if non-nil, zero value otherwise.

### GetCoordSourceOk

`func (o *RareDiseaseTrialSitesRow) GetCoordSourceOk() (*string, bool)`

GetCoordSourceOk returns a tuple with the CoordSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordSource

`func (o *RareDiseaseTrialSitesRow) SetCoordSource(v string)`

SetCoordSource sets CoordSource field to given value.

### HasCoordSource

`func (o *RareDiseaseTrialSitesRow) HasCoordSource() bool`

HasCoordSource returns a boolean if a field has been set.

### SetCoordSourceNil

`func (o *RareDiseaseTrialSitesRow) SetCoordSourceNil(b bool)`

 SetCoordSourceNil sets the value for CoordSource to be an explicit nil

### UnsetCoordSource
`func (o *RareDiseaseTrialSitesRow) UnsetCoordSource()`

UnsetCoordSource ensures that no value is present for CoordSource, not even an explicit nil
### GetConditionIds

`func (o *RareDiseaseTrialSitesRow) GetConditionIds() []string`

GetConditionIds returns the ConditionIds field if non-nil, zero value otherwise.

### GetConditionIdsOk

`func (o *RareDiseaseTrialSitesRow) GetConditionIdsOk() (*[]string, bool)`

GetConditionIdsOk returns a tuple with the ConditionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionIds

`func (o *RareDiseaseTrialSitesRow) SetConditionIds(v []string)`

SetConditionIds sets ConditionIds field to given value.

### HasConditionIds

`func (o *RareDiseaseTrialSitesRow) HasConditionIds() bool`

HasConditionIds returns a boolean if a field has been set.

### SetConditionIdsNil

`func (o *RareDiseaseTrialSitesRow) SetConditionIdsNil(b bool)`

 SetConditionIdsNil sets the value for ConditionIds to be an explicit nil

### UnsetConditionIds
`func (o *RareDiseaseTrialSitesRow) UnsetConditionIds()`

UnsetConditionIds ensures that no value is present for ConditionIds, not even an explicit nil
### GetNTrials

`func (o *RareDiseaseTrialSitesRow) GetNTrials() int32`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *RareDiseaseTrialSitesRow) GetNTrialsOk() (*int32, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *RareDiseaseTrialSitesRow) SetNTrials(v int32)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *RareDiseaseTrialSitesRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *RareDiseaseTrialSitesRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *RareDiseaseTrialSitesRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNTrialsDirect

`func (o *RareDiseaseTrialSitesRow) GetNTrialsDirect() int32`

GetNTrialsDirect returns the NTrialsDirect field if non-nil, zero value otherwise.

### GetNTrialsDirectOk

`func (o *RareDiseaseTrialSitesRow) GetNTrialsDirectOk() (*int32, bool)`

GetNTrialsDirectOk returns a tuple with the NTrialsDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsDirect

`func (o *RareDiseaseTrialSitesRow) SetNTrialsDirect(v int32)`

SetNTrialsDirect sets NTrialsDirect field to given value.

### HasNTrialsDirect

`func (o *RareDiseaseTrialSitesRow) HasNTrialsDirect() bool`

HasNTrialsDirect returns a boolean if a field has been set.

### SetNTrialsDirectNil

`func (o *RareDiseaseTrialSitesRow) SetNTrialsDirectNil(b bool)`

 SetNTrialsDirectNil sets the value for NTrialsDirect to be an explicit nil

### UnsetNTrialsDirect
`func (o *RareDiseaseTrialSitesRow) UnsetNTrialsDirect()`

UnsetNTrialsDirect ensures that no value is present for NTrialsDirect, not even an explicit nil
### GetNActiveTrials

`func (o *RareDiseaseTrialSitesRow) GetNActiveTrials() int32`

GetNActiveTrials returns the NActiveTrials field if non-nil, zero value otherwise.

### GetNActiveTrialsOk

`func (o *RareDiseaseTrialSitesRow) GetNActiveTrialsOk() (*int32, bool)`

GetNActiveTrialsOk returns a tuple with the NActiveTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNActiveTrials

`func (o *RareDiseaseTrialSitesRow) SetNActiveTrials(v int32)`

SetNActiveTrials sets NActiveTrials field to given value.

### HasNActiveTrials

`func (o *RareDiseaseTrialSitesRow) HasNActiveTrials() bool`

HasNActiveTrials returns a boolean if a field has been set.

### SetNActiveTrialsNil

`func (o *RareDiseaseTrialSitesRow) SetNActiveTrialsNil(b bool)`

 SetNActiveTrialsNil sets the value for NActiveTrials to be an explicit nil

### UnsetNActiveTrials
`func (o *RareDiseaseTrialSitesRow) UnsetNActiveTrials()`

UnsetNActiveTrials ensures that no value is present for NActiveTrials, not even an explicit nil
### GetFirstTrialYear

`func (o *RareDiseaseTrialSitesRow) GetFirstTrialYear() int32`

GetFirstTrialYear returns the FirstTrialYear field if non-nil, zero value otherwise.

### GetFirstTrialYearOk

`func (o *RareDiseaseTrialSitesRow) GetFirstTrialYearOk() (*int32, bool)`

GetFirstTrialYearOk returns a tuple with the FirstTrialYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTrialYear

`func (o *RareDiseaseTrialSitesRow) SetFirstTrialYear(v int32)`

SetFirstTrialYear sets FirstTrialYear field to given value.

### HasFirstTrialYear

`func (o *RareDiseaseTrialSitesRow) HasFirstTrialYear() bool`

HasFirstTrialYear returns a boolean if a field has been set.

### SetFirstTrialYearNil

`func (o *RareDiseaseTrialSitesRow) SetFirstTrialYearNil(b bool)`

 SetFirstTrialYearNil sets the value for FirstTrialYear to be an explicit nil

### UnsetFirstTrialYear
`func (o *RareDiseaseTrialSitesRow) UnsetFirstTrialYear()`

UnsetFirstTrialYear ensures that no value is present for FirstTrialYear, not even an explicit nil
### GetLastTrialYear

`func (o *RareDiseaseTrialSitesRow) GetLastTrialYear() int32`

GetLastTrialYear returns the LastTrialYear field if non-nil, zero value otherwise.

### GetLastTrialYearOk

`func (o *RareDiseaseTrialSitesRow) GetLastTrialYearOk() (*int32, bool)`

GetLastTrialYearOk returns a tuple with the LastTrialYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrialYear

`func (o *RareDiseaseTrialSitesRow) SetLastTrialYear(v int32)`

SetLastTrialYear sets LastTrialYear field to given value.

### HasLastTrialYear

`func (o *RareDiseaseTrialSitesRow) HasLastTrialYear() bool`

HasLastTrialYear returns a boolean if a field has been set.

### SetLastTrialYearNil

`func (o *RareDiseaseTrialSitesRow) SetLastTrialYearNil(b bool)`

 SetLastTrialYearNil sets the value for LastTrialYear to be an explicit nil

### UnsetLastTrialYear
`func (o *RareDiseaseTrialSitesRow) UnsetLastTrialYear()`

UnsetLastTrialYear ensures that no value is present for LastTrialYear, not even an explicit nil
### GetLastUpdatePostedDate

`func (o *RareDiseaseTrialSitesRow) GetLastUpdatePostedDate() string`

GetLastUpdatePostedDate returns the LastUpdatePostedDate field if non-nil, zero value otherwise.

### GetLastUpdatePostedDateOk

`func (o *RareDiseaseTrialSitesRow) GetLastUpdatePostedDateOk() (*string, bool)`

GetLastUpdatePostedDateOk returns a tuple with the LastUpdatePostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatePostedDate

`func (o *RareDiseaseTrialSitesRow) SetLastUpdatePostedDate(v string)`

SetLastUpdatePostedDate sets LastUpdatePostedDate field to given value.

### HasLastUpdatePostedDate

`func (o *RareDiseaseTrialSitesRow) HasLastUpdatePostedDate() bool`

HasLastUpdatePostedDate returns a boolean if a field has been set.

### SetLastUpdatePostedDateNil

`func (o *RareDiseaseTrialSitesRow) SetLastUpdatePostedDateNil(b bool)`

 SetLastUpdatePostedDateNil sets the value for LastUpdatePostedDate to be an explicit nil

### UnsetLastUpdatePostedDate
`func (o *RareDiseaseTrialSitesRow) UnsetLastUpdatePostedDate()`

UnsetLastUpdatePostedDate ensures that no value is present for LastUpdatePostedDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


