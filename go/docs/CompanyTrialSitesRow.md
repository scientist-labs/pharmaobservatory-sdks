# CompanyTrialSitesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** | The company key (company:&lt;key&gt;) of the lead sponsor. | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**RosterRootId** | Pointer to **NullableString** | The Clinical Research Map&#39;s organization key for these rows: the group when it is a public roster root, else the company when it is one, else null. | [optional] 
**SiteId** | Pointer to **NullableString** | The site key: one facility spelling from the registry; the id the site page address carries. | [optional] 
**FacilityName** | Pointer to **NullableString** | The facility name as filed in the registry. | [optional] 
**City** | Pointer to **NullableString** | City as filed. | [optional] 
**State** | Pointer to **NullableString** | State or province as filed. | [optional] 
**Country** | Pointer to **NullableString** | Country name as filed. | [optional] 
**SiteCountryCode** | Pointer to **NullableString** | ISO 3166-1 alpha-2 code of the site&#39;s country, from the registry&#39;s country name; null when the name resolves to no country. | [optional] 
**Latitude** | Pointer to **NullableFloat32** | Latitude of the finest coordinate the site record has. | [optional] 
**Longitude** | Pointer to **NullableFloat32** | Longitude of the finest coordinate the site record has. | [optional] 
**GeoPrecision** | Pointer to **NullableString** | How finely the coordinate places the site (city, postal and so on). | [optional] 
**CoordSource** | Pointer to **NullableString** | The source of the coordinate: aact (the registry&#39;s own), geonames (GeoNames) or als (Amazon Location Service). | [optional] 
**NTrials** | Pointer to **NullableInt32** | Distinct trials the company leads that list the site. | [optional] 
**SiteRank** | Pointer to **NullableInt32** | 1 is the company&#39;s busiest located site (most led trials, then site key); dense per company. | [optional] 
**NSites** | Pointer to **NullableInt32** | The company&#39;s located site count, repeated on each of its rows. | [optional] 

## Methods

### NewCompanyTrialSitesRow

`func NewCompanyTrialSitesRow() *CompanyTrialSitesRow`

NewCompanyTrialSitesRow instantiates a new CompanyTrialSitesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyTrialSitesRowWithDefaults

`func NewCompanyTrialSitesRowWithDefaults() *CompanyTrialSitesRow`

NewCompanyTrialSitesRowWithDefaults instantiates a new CompanyTrialSitesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CompanyTrialSitesRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyTrialSitesRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyTrialSitesRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CompanyTrialSitesRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CompanyTrialSitesRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CompanyTrialSitesRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *CompanyTrialSitesRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyTrialSitesRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyTrialSitesRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyTrialSitesRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompanyTrialSitesRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyTrialSitesRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *CompanyTrialSitesRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CompanyTrialSitesRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CompanyTrialSitesRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CompanyTrialSitesRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CompanyTrialSitesRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CompanyTrialSitesRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *CompanyTrialSitesRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CompanyTrialSitesRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CompanyTrialSitesRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CompanyTrialSitesRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CompanyTrialSitesRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CompanyTrialSitesRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetRosterRootId

`func (o *CompanyTrialSitesRow) GetRosterRootId() string`

GetRosterRootId returns the RosterRootId field if non-nil, zero value otherwise.

### GetRosterRootIdOk

`func (o *CompanyTrialSitesRow) GetRosterRootIdOk() (*string, bool)`

GetRosterRootIdOk returns a tuple with the RosterRootId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRosterRootId

`func (o *CompanyTrialSitesRow) SetRosterRootId(v string)`

SetRosterRootId sets RosterRootId field to given value.

### HasRosterRootId

`func (o *CompanyTrialSitesRow) HasRosterRootId() bool`

HasRosterRootId returns a boolean if a field has been set.

### SetRosterRootIdNil

`func (o *CompanyTrialSitesRow) SetRosterRootIdNil(b bool)`

 SetRosterRootIdNil sets the value for RosterRootId to be an explicit nil

### UnsetRosterRootId
`func (o *CompanyTrialSitesRow) UnsetRosterRootId()`

UnsetRosterRootId ensures that no value is present for RosterRootId, not even an explicit nil
### GetSiteId

`func (o *CompanyTrialSitesRow) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CompanyTrialSitesRow) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CompanyTrialSitesRow) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *CompanyTrialSitesRow) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *CompanyTrialSitesRow) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *CompanyTrialSitesRow) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetFacilityName

`func (o *CompanyTrialSitesRow) GetFacilityName() string`

GetFacilityName returns the FacilityName field if non-nil, zero value otherwise.

### GetFacilityNameOk

`func (o *CompanyTrialSitesRow) GetFacilityNameOk() (*string, bool)`

GetFacilityNameOk returns a tuple with the FacilityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityName

`func (o *CompanyTrialSitesRow) SetFacilityName(v string)`

SetFacilityName sets FacilityName field to given value.

### HasFacilityName

`func (o *CompanyTrialSitesRow) HasFacilityName() bool`

HasFacilityName returns a boolean if a field has been set.

### SetFacilityNameNil

`func (o *CompanyTrialSitesRow) SetFacilityNameNil(b bool)`

 SetFacilityNameNil sets the value for FacilityName to be an explicit nil

### UnsetFacilityName
`func (o *CompanyTrialSitesRow) UnsetFacilityName()`

UnsetFacilityName ensures that no value is present for FacilityName, not even an explicit nil
### GetCity

`func (o *CompanyTrialSitesRow) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CompanyTrialSitesRow) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CompanyTrialSitesRow) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CompanyTrialSitesRow) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *CompanyTrialSitesRow) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *CompanyTrialSitesRow) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *CompanyTrialSitesRow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CompanyTrialSitesRow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CompanyTrialSitesRow) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CompanyTrialSitesRow) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *CompanyTrialSitesRow) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *CompanyTrialSitesRow) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCountry

`func (o *CompanyTrialSitesRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CompanyTrialSitesRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CompanyTrialSitesRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CompanyTrialSitesRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *CompanyTrialSitesRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *CompanyTrialSitesRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetSiteCountryCode

`func (o *CompanyTrialSitesRow) GetSiteCountryCode() string`

GetSiteCountryCode returns the SiteCountryCode field if non-nil, zero value otherwise.

### GetSiteCountryCodeOk

`func (o *CompanyTrialSitesRow) GetSiteCountryCodeOk() (*string, bool)`

GetSiteCountryCodeOk returns a tuple with the SiteCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCountryCode

`func (o *CompanyTrialSitesRow) SetSiteCountryCode(v string)`

SetSiteCountryCode sets SiteCountryCode field to given value.

### HasSiteCountryCode

`func (o *CompanyTrialSitesRow) HasSiteCountryCode() bool`

HasSiteCountryCode returns a boolean if a field has been set.

### SetSiteCountryCodeNil

`func (o *CompanyTrialSitesRow) SetSiteCountryCodeNil(b bool)`

 SetSiteCountryCodeNil sets the value for SiteCountryCode to be an explicit nil

### UnsetSiteCountryCode
`func (o *CompanyTrialSitesRow) UnsetSiteCountryCode()`

UnsetSiteCountryCode ensures that no value is present for SiteCountryCode, not even an explicit nil
### GetLatitude

`func (o *CompanyTrialSitesRow) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *CompanyTrialSitesRow) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *CompanyTrialSitesRow) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *CompanyTrialSitesRow) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *CompanyTrialSitesRow) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *CompanyTrialSitesRow) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *CompanyTrialSitesRow) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *CompanyTrialSitesRow) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *CompanyTrialSitesRow) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *CompanyTrialSitesRow) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *CompanyTrialSitesRow) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *CompanyTrialSitesRow) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetGeoPrecision

`func (o *CompanyTrialSitesRow) GetGeoPrecision() string`

GetGeoPrecision returns the GeoPrecision field if non-nil, zero value otherwise.

### GetGeoPrecisionOk

`func (o *CompanyTrialSitesRow) GetGeoPrecisionOk() (*string, bool)`

GetGeoPrecisionOk returns a tuple with the GeoPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoPrecision

`func (o *CompanyTrialSitesRow) SetGeoPrecision(v string)`

SetGeoPrecision sets GeoPrecision field to given value.

### HasGeoPrecision

`func (o *CompanyTrialSitesRow) HasGeoPrecision() bool`

HasGeoPrecision returns a boolean if a field has been set.

### SetGeoPrecisionNil

`func (o *CompanyTrialSitesRow) SetGeoPrecisionNil(b bool)`

 SetGeoPrecisionNil sets the value for GeoPrecision to be an explicit nil

### UnsetGeoPrecision
`func (o *CompanyTrialSitesRow) UnsetGeoPrecision()`

UnsetGeoPrecision ensures that no value is present for GeoPrecision, not even an explicit nil
### GetCoordSource

`func (o *CompanyTrialSitesRow) GetCoordSource() string`

GetCoordSource returns the CoordSource field if non-nil, zero value otherwise.

### GetCoordSourceOk

`func (o *CompanyTrialSitesRow) GetCoordSourceOk() (*string, bool)`

GetCoordSourceOk returns a tuple with the CoordSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordSource

`func (o *CompanyTrialSitesRow) SetCoordSource(v string)`

SetCoordSource sets CoordSource field to given value.

### HasCoordSource

`func (o *CompanyTrialSitesRow) HasCoordSource() bool`

HasCoordSource returns a boolean if a field has been set.

### SetCoordSourceNil

`func (o *CompanyTrialSitesRow) SetCoordSourceNil(b bool)`

 SetCoordSourceNil sets the value for CoordSource to be an explicit nil

### UnsetCoordSource
`func (o *CompanyTrialSitesRow) UnsetCoordSource()`

UnsetCoordSource ensures that no value is present for CoordSource, not even an explicit nil
### GetNTrials

`func (o *CompanyTrialSitesRow) GetNTrials() int32`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *CompanyTrialSitesRow) GetNTrialsOk() (*int32, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *CompanyTrialSitesRow) SetNTrials(v int32)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *CompanyTrialSitesRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *CompanyTrialSitesRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *CompanyTrialSitesRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetSiteRank

`func (o *CompanyTrialSitesRow) GetSiteRank() int32`

GetSiteRank returns the SiteRank field if non-nil, zero value otherwise.

### GetSiteRankOk

`func (o *CompanyTrialSitesRow) GetSiteRankOk() (*int32, bool)`

GetSiteRankOk returns a tuple with the SiteRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteRank

`func (o *CompanyTrialSitesRow) SetSiteRank(v int32)`

SetSiteRank sets SiteRank field to given value.

### HasSiteRank

`func (o *CompanyTrialSitesRow) HasSiteRank() bool`

HasSiteRank returns a boolean if a field has been set.

### SetSiteRankNil

`func (o *CompanyTrialSitesRow) SetSiteRankNil(b bool)`

 SetSiteRankNil sets the value for SiteRank to be an explicit nil

### UnsetSiteRank
`func (o *CompanyTrialSitesRow) UnsetSiteRank()`

UnsetSiteRank ensures that no value is present for SiteRank, not even an explicit nil
### GetNSites

`func (o *CompanyTrialSitesRow) GetNSites() int32`

GetNSites returns the NSites field if non-nil, zero value otherwise.

### GetNSitesOk

`func (o *CompanyTrialSitesRow) GetNSitesOk() (*int32, bool)`

GetNSitesOk returns a tuple with the NSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSites

`func (o *CompanyTrialSitesRow) SetNSites(v int32)`

SetNSites sets NSites field to given value.

### HasNSites

`func (o *CompanyTrialSitesRow) HasNSites() bool`

HasNSites returns a boolean if a field has been set.

### SetNSitesNil

`func (o *CompanyTrialSitesRow) SetNSitesNil(b bool)`

 SetNSitesNil sets the value for NSites to be an explicit nil

### UnsetNSites
`func (o *CompanyTrialSitesRow) UnsetNSites()`

UnsetNSites ensures that no value is present for NSites, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


