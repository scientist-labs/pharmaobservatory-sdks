# OrgsGxpSites200ResponseRowsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FacilityName** | **string** |  | 
**City** | **string** |  | 
**Country** | **string** |  | 
**Regulators** | **[]string** |  | 
**CanonicalOrgId** | **string** |  | 
**OrgName** | **string** |  | 
**FacilityKey** | Pointer to **string** |  | [optional] 
**StateRegion** | Pointer to **interface{}** |  | [optional] 
**SeverityLevel** | Pointer to **string** |  | [optional] 
**SeverityRawVerdict** | Pointer to **NullableString** |  | [optional] 
**SeveritySignalDate** | Pointer to **NullableString** |  | [optional] 
**MatchTier** | Pointer to **string** |  | [optional] 
**RestrictionTier** | Pointer to **string** |  | [optional] 
**RegisterDomain** | Pointer to **string** |  | [optional] 
**Latitude** | Pointer to **interface{}** |  | [optional] 
**Longitude** | Pointer to **interface{}** |  | [optional] 
**GeocodePrecision** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewOrgsGxpSites200ResponseRowsInner

`func NewOrgsGxpSites200ResponseRowsInner(facilityName string, city string, country string, regulators []string, canonicalOrgId string, orgName string, ) *OrgsGxpSites200ResponseRowsInner`

NewOrgsGxpSites200ResponseRowsInner instantiates a new OrgsGxpSites200ResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsGxpSites200ResponseRowsInnerWithDefaults

`func NewOrgsGxpSites200ResponseRowsInnerWithDefaults() *OrgsGxpSites200ResponseRowsInner`

NewOrgsGxpSites200ResponseRowsInnerWithDefaults instantiates a new OrgsGxpSites200ResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacilityName

`func (o *OrgsGxpSites200ResponseRowsInner) GetFacilityName() string`

GetFacilityName returns the FacilityName field if non-nil, zero value otherwise.

### GetFacilityNameOk

`func (o *OrgsGxpSites200ResponseRowsInner) GetFacilityNameOk() (*string, bool)`

GetFacilityNameOk returns a tuple with the FacilityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityName

`func (o *OrgsGxpSites200ResponseRowsInner) SetFacilityName(v string)`

SetFacilityName sets FacilityName field to given value.


### GetCity

`func (o *OrgsGxpSites200ResponseRowsInner) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrgsGxpSites200ResponseRowsInner) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrgsGxpSites200ResponseRowsInner) SetCity(v string)`

SetCity sets City field to given value.


### GetCountry

`func (o *OrgsGxpSites200ResponseRowsInner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrgsGxpSites200ResponseRowsInner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrgsGxpSites200ResponseRowsInner) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetRegulators

`func (o *OrgsGxpSites200ResponseRowsInner) GetRegulators() []string`

GetRegulators returns the Regulators field if non-nil, zero value otherwise.

### GetRegulatorsOk

`func (o *OrgsGxpSites200ResponseRowsInner) GetRegulatorsOk() (*[]string, bool)`

GetRegulatorsOk returns a tuple with the Regulators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulators

`func (o *OrgsGxpSites200ResponseRowsInner) SetRegulators(v []string)`

SetRegulators sets Regulators field to given value.


### GetCanonicalOrgId

`func (o *OrgsGxpSites200ResponseRowsInner) GetCanonicalOrgId() string`

GetCanonicalOrgId returns the CanonicalOrgId field if non-nil, zero value otherwise.

### GetCanonicalOrgIdOk

`func (o *OrgsGxpSites200ResponseRowsInner) GetCanonicalOrgIdOk() (*string, bool)`

GetCanonicalOrgIdOk returns a tuple with the CanonicalOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalOrgId

`func (o *OrgsGxpSites200ResponseRowsInner) SetCanonicalOrgId(v string)`

SetCanonicalOrgId sets CanonicalOrgId field to given value.


### GetOrgName

`func (o *OrgsGxpSites200ResponseRowsInner) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *OrgsGxpSites200ResponseRowsInner) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *OrgsGxpSites200ResponseRowsInner) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetFacilityKey

`func (o *OrgsGxpSites200ResponseRowsInner) GetFacilityKey() string`

GetFacilityKey returns the FacilityKey field if non-nil, zero value otherwise.

### GetFacilityKeyOk

`func (o *OrgsGxpSites200ResponseRowsInner) GetFacilityKeyOk() (*string, bool)`

GetFacilityKeyOk returns a tuple with the FacilityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityKey

`func (o *OrgsGxpSites200ResponseRowsInner) SetFacilityKey(v string)`

SetFacilityKey sets FacilityKey field to given value.

### HasFacilityKey

`func (o *OrgsGxpSites200ResponseRowsInner) HasFacilityKey() bool`

HasFacilityKey returns a boolean if a field has been set.

### GetStateRegion

`func (o *OrgsGxpSites200ResponseRowsInner) GetStateRegion() interface{}`

GetStateRegion returns the StateRegion field if non-nil, zero value otherwise.

### GetStateRegionOk

`func (o *OrgsGxpSites200ResponseRowsInner) GetStateRegionOk() (*interface{}, bool)`

GetStateRegionOk returns a tuple with the StateRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRegion

`func (o *OrgsGxpSites200ResponseRowsInner) SetStateRegion(v interface{})`

SetStateRegion sets StateRegion field to given value.

### HasStateRegion

`func (o *OrgsGxpSites200ResponseRowsInner) HasStateRegion() bool`

HasStateRegion returns a boolean if a field has been set.

### SetStateRegionNil

`func (o *OrgsGxpSites200ResponseRowsInner) SetStateRegionNil(b bool)`

 SetStateRegionNil sets the value for StateRegion to be an explicit nil

### UnsetStateRegion
`func (o *OrgsGxpSites200ResponseRowsInner) UnsetStateRegion()`

UnsetStateRegion ensures that no value is present for StateRegion, not even an explicit nil
### GetSeverityLevel

`func (o *OrgsGxpSites200ResponseRowsInner) GetSeverityLevel() string`

GetSeverityLevel returns the SeverityLevel field if non-nil, zero value otherwise.

### GetSeverityLevelOk

`func (o *OrgsGxpSites200ResponseRowsInner) GetSeverityLevelOk() (*string, bool)`

GetSeverityLevelOk returns a tuple with the SeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevel

`func (o *OrgsGxpSites200ResponseRowsInner) SetSeverityLevel(v string)`

SetSeverityLevel sets SeverityLevel field to given value.

### HasSeverityLevel

`func (o *OrgsGxpSites200ResponseRowsInner) HasSeverityLevel() bool`

HasSeverityLevel returns a boolean if a field has been set.

### GetSeverityRawVerdict

`func (o *OrgsGxpSites200ResponseRowsInner) GetSeverityRawVerdict() string`

GetSeverityRawVerdict returns the SeverityRawVerdict field if non-nil, zero value otherwise.

### GetSeverityRawVerdictOk

`func (o *OrgsGxpSites200ResponseRowsInner) GetSeverityRawVerdictOk() (*string, bool)`

GetSeverityRawVerdictOk returns a tuple with the SeverityRawVerdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityRawVerdict

`func (o *OrgsGxpSites200ResponseRowsInner) SetSeverityRawVerdict(v string)`

SetSeverityRawVerdict sets SeverityRawVerdict field to given value.

### HasSeverityRawVerdict

`func (o *OrgsGxpSites200ResponseRowsInner) HasSeverityRawVerdict() bool`

HasSeverityRawVerdict returns a boolean if a field has been set.

### SetSeverityRawVerdictNil

`func (o *OrgsGxpSites200ResponseRowsInner) SetSeverityRawVerdictNil(b bool)`

 SetSeverityRawVerdictNil sets the value for SeverityRawVerdict to be an explicit nil

### UnsetSeverityRawVerdict
`func (o *OrgsGxpSites200ResponseRowsInner) UnsetSeverityRawVerdict()`

UnsetSeverityRawVerdict ensures that no value is present for SeverityRawVerdict, not even an explicit nil
### GetSeveritySignalDate

`func (o *OrgsGxpSites200ResponseRowsInner) GetSeveritySignalDate() string`

GetSeveritySignalDate returns the SeveritySignalDate field if non-nil, zero value otherwise.

### GetSeveritySignalDateOk

`func (o *OrgsGxpSites200ResponseRowsInner) GetSeveritySignalDateOk() (*string, bool)`

GetSeveritySignalDateOk returns a tuple with the SeveritySignalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeveritySignalDate

`func (o *OrgsGxpSites200ResponseRowsInner) SetSeveritySignalDate(v string)`

SetSeveritySignalDate sets SeveritySignalDate field to given value.

### HasSeveritySignalDate

`func (o *OrgsGxpSites200ResponseRowsInner) HasSeveritySignalDate() bool`

HasSeveritySignalDate returns a boolean if a field has been set.

### SetSeveritySignalDateNil

`func (o *OrgsGxpSites200ResponseRowsInner) SetSeveritySignalDateNil(b bool)`

 SetSeveritySignalDateNil sets the value for SeveritySignalDate to be an explicit nil

### UnsetSeveritySignalDate
`func (o *OrgsGxpSites200ResponseRowsInner) UnsetSeveritySignalDate()`

UnsetSeveritySignalDate ensures that no value is present for SeveritySignalDate, not even an explicit nil
### GetMatchTier

`func (o *OrgsGxpSites200ResponseRowsInner) GetMatchTier() string`

GetMatchTier returns the MatchTier field if non-nil, zero value otherwise.

### GetMatchTierOk

`func (o *OrgsGxpSites200ResponseRowsInner) GetMatchTierOk() (*string, bool)`

GetMatchTierOk returns a tuple with the MatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTier

`func (o *OrgsGxpSites200ResponseRowsInner) SetMatchTier(v string)`

SetMatchTier sets MatchTier field to given value.

### HasMatchTier

`func (o *OrgsGxpSites200ResponseRowsInner) HasMatchTier() bool`

HasMatchTier returns a boolean if a field has been set.

### GetRestrictionTier

`func (o *OrgsGxpSites200ResponseRowsInner) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *OrgsGxpSites200ResponseRowsInner) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *OrgsGxpSites200ResponseRowsInner) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *OrgsGxpSites200ResponseRowsInner) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### GetRegisterDomain

`func (o *OrgsGxpSites200ResponseRowsInner) GetRegisterDomain() string`

GetRegisterDomain returns the RegisterDomain field if non-nil, zero value otherwise.

### GetRegisterDomainOk

`func (o *OrgsGxpSites200ResponseRowsInner) GetRegisterDomainOk() (*string, bool)`

GetRegisterDomainOk returns a tuple with the RegisterDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterDomain

`func (o *OrgsGxpSites200ResponseRowsInner) SetRegisterDomain(v string)`

SetRegisterDomain sets RegisterDomain field to given value.

### HasRegisterDomain

`func (o *OrgsGxpSites200ResponseRowsInner) HasRegisterDomain() bool`

HasRegisterDomain returns a boolean if a field has been set.

### GetLatitude

`func (o *OrgsGxpSites200ResponseRowsInner) GetLatitude() interface{}`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *OrgsGxpSites200ResponseRowsInner) GetLatitudeOk() (*interface{}, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *OrgsGxpSites200ResponseRowsInner) SetLatitude(v interface{})`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *OrgsGxpSites200ResponseRowsInner) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *OrgsGxpSites200ResponseRowsInner) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *OrgsGxpSites200ResponseRowsInner) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *OrgsGxpSites200ResponseRowsInner) GetLongitude() interface{}`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *OrgsGxpSites200ResponseRowsInner) GetLongitudeOk() (*interface{}, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *OrgsGxpSites200ResponseRowsInner) SetLongitude(v interface{})`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *OrgsGxpSites200ResponseRowsInner) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *OrgsGxpSites200ResponseRowsInner) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *OrgsGxpSites200ResponseRowsInner) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetGeocodePrecision

`func (o *OrgsGxpSites200ResponseRowsInner) GetGeocodePrecision() interface{}`

GetGeocodePrecision returns the GeocodePrecision field if non-nil, zero value otherwise.

### GetGeocodePrecisionOk

`func (o *OrgsGxpSites200ResponseRowsInner) GetGeocodePrecisionOk() (*interface{}, bool)`

GetGeocodePrecisionOk returns a tuple with the GeocodePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocodePrecision

`func (o *OrgsGxpSites200ResponseRowsInner) SetGeocodePrecision(v interface{})`

SetGeocodePrecision sets GeocodePrecision field to given value.

### HasGeocodePrecision

`func (o *OrgsGxpSites200ResponseRowsInner) HasGeocodePrecision() bool`

HasGeocodePrecision returns a boolean if a field has been set.

### SetGeocodePrecisionNil

`func (o *OrgsGxpSites200ResponseRowsInner) SetGeocodePrecisionNil(b bool)`

 SetGeocodePrecisionNil sets the value for GeocodePrecision to be an explicit nil

### UnsetGeocodePrecision
`func (o *OrgsGxpSites200ResponseRowsInner) UnsetGeocodePrecision()`

UnsetGeocodePrecision ensures that no value is present for GeocodePrecision, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


