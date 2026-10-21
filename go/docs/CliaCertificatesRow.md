# CliaCertificatesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CliaNumber** | Pointer to **NullableString** | The ten-character CLIA certificate number. | [optional] 
**LabName** | Pointer to **NullableString** | The laboratory&#39;s name as listed on the certificate. | [optional] 
**FirmName** | Pointer to **NullableString** | The organization the certificate resolved to (its canonical name). | [optional] 
**FirmOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] 
**FirmCompanyId** | Pointer to **NullableString** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] 
**FirmCompanyName** | Pointer to **NullableString** | Display name of the firm&#39;s company. | [optional] 
**FirmGroupId** | Pointer to **NullableString** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**FirmGroupName** | Pointer to **NullableString** | Display name of the firm&#39;s group. | [optional] 
**MatchTier** | Pointer to **NullableString** | How the certificate was attached to the organization: high, medium or low. | [optional] 
**MatchConfidence** | Pointer to **NullableFloat32** | The matcher&#39;s confidence, 0 to 1. | [optional] 
**City** | Pointer to **NullableString** | The laboratory&#39;s city. | [optional] 
**StateRegion** | Pointer to **NullableString** | The laboratory&#39;s state. | [optional] 
**Postal5** | Pointer to **NullableString** | The five-digit postal code. | [optional] 
**Country** | Pointer to **NullableString** | The country code. | [optional] 
**Latitude** | Pointer to **NullableFloat32** | Latitude in degrees (WGS84), for the map tool. | [optional] 
**Longitude** | Pointer to **NullableFloat32** | Longitude in degrees (WGS84), for the map tool. | [optional] 
**GeocodePrecision** | Pointer to **NullableString** | What a pin locates (an address, a postal code or a city). | [optional] 
**CoordSource** | Pointer to **NullableString** | Where the coordinates came from. | [optional] 
**IsActive** | Pointer to **NullableBool** | True when the certificate is active on the file. | [optional] 
**CertificateType** | Pointer to **NullableString** | The CLIA certificate type (for example Compliance, Accreditation or Waiver). | [optional] 
**ComplexityCeiling** | Pointer to **NullableString** | The highest test complexity the certificate permits. | [optional] 
**AccreditorClaimed** | Pointer to **[]string** | The accreditation body the file names. | [optional] 
**AccreditorConfirmed** | Pointer to **[]string** | The accreditation body confirmed against the accreditor&#39;s own list. | [optional] 
**HasCapAccreditation** | Pointer to **NullableBool** | True when the laboratory holds College of American Pathologists (CAP) accreditation. | [optional] 
**HasAshiAccreditation** | Pointer to **NullableBool** | True when the laboratory holds American Society for Histocompatibility and Immunogenetics (ASHI) accreditation. | [optional] 
**HasAabbAccreditation** | Pointer to **NullableBool** | True when the laboratory holds AABB (formerly American Association of Blood Banks) accreditation. | [optional] 
**FacilityType** | Pointer to **NullableString** | The laboratory&#39;s facility type on the file. | [optional] 
**AnnualTestVolume** | Pointer to **NullableInt64** | Reported annual test volume. | [optional] 
**TestVolumeSource** | Pointer to **NullableString** | Where the volume figure came from. | [optional] 
**LabSiteCount** | Pointer to **NullableInt32** | How many sites the laboratory reports. | [optional] 
**PosVintage** | Pointer to **NullableString** | The vintage of the CMS Provider of Services file. | [optional] 
**AsOf** | Pointer to **NullableString** | The build date of the figures. | [optional] 

## Methods

### NewCliaCertificatesRow

`func NewCliaCertificatesRow() *CliaCertificatesRow`

NewCliaCertificatesRow instantiates a new CliaCertificatesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCliaCertificatesRowWithDefaults

`func NewCliaCertificatesRowWithDefaults() *CliaCertificatesRow`

NewCliaCertificatesRowWithDefaults instantiates a new CliaCertificatesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCliaNumber

`func (o *CliaCertificatesRow) GetCliaNumber() string`

GetCliaNumber returns the CliaNumber field if non-nil, zero value otherwise.

### GetCliaNumberOk

`func (o *CliaCertificatesRow) GetCliaNumberOk() (*string, bool)`

GetCliaNumberOk returns a tuple with the CliaNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliaNumber

`func (o *CliaCertificatesRow) SetCliaNumber(v string)`

SetCliaNumber sets CliaNumber field to given value.

### HasCliaNumber

`func (o *CliaCertificatesRow) HasCliaNumber() bool`

HasCliaNumber returns a boolean if a field has been set.

### SetCliaNumberNil

`func (o *CliaCertificatesRow) SetCliaNumberNil(b bool)`

 SetCliaNumberNil sets the value for CliaNumber to be an explicit nil

### UnsetCliaNumber
`func (o *CliaCertificatesRow) UnsetCliaNumber()`

UnsetCliaNumber ensures that no value is present for CliaNumber, not even an explicit nil
### GetLabName

`func (o *CliaCertificatesRow) GetLabName() string`

GetLabName returns the LabName field if non-nil, zero value otherwise.

### GetLabNameOk

`func (o *CliaCertificatesRow) GetLabNameOk() (*string, bool)`

GetLabNameOk returns a tuple with the LabName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabName

`func (o *CliaCertificatesRow) SetLabName(v string)`

SetLabName sets LabName field to given value.

### HasLabName

`func (o *CliaCertificatesRow) HasLabName() bool`

HasLabName returns a boolean if a field has been set.

### SetLabNameNil

`func (o *CliaCertificatesRow) SetLabNameNil(b bool)`

 SetLabNameNil sets the value for LabName to be an explicit nil

### UnsetLabName
`func (o *CliaCertificatesRow) UnsetLabName()`

UnsetLabName ensures that no value is present for LabName, not even an explicit nil
### GetFirmName

`func (o *CliaCertificatesRow) GetFirmName() string`

GetFirmName returns the FirmName field if non-nil, zero value otherwise.

### GetFirmNameOk

`func (o *CliaCertificatesRow) GetFirmNameOk() (*string, bool)`

GetFirmNameOk returns a tuple with the FirmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmName

`func (o *CliaCertificatesRow) SetFirmName(v string)`

SetFirmName sets FirmName field to given value.

### HasFirmName

`func (o *CliaCertificatesRow) HasFirmName() bool`

HasFirmName returns a boolean if a field has been set.

### SetFirmNameNil

`func (o *CliaCertificatesRow) SetFirmNameNil(b bool)`

 SetFirmNameNil sets the value for FirmName to be an explicit nil

### UnsetFirmName
`func (o *CliaCertificatesRow) UnsetFirmName()`

UnsetFirmName ensures that no value is present for FirmName, not even an explicit nil
### GetFirmOrgId

`func (o *CliaCertificatesRow) GetFirmOrgId() string`

GetFirmOrgId returns the FirmOrgId field if non-nil, zero value otherwise.

### GetFirmOrgIdOk

`func (o *CliaCertificatesRow) GetFirmOrgIdOk() (*string, bool)`

GetFirmOrgIdOk returns a tuple with the FirmOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmOrgId

`func (o *CliaCertificatesRow) SetFirmOrgId(v string)`

SetFirmOrgId sets FirmOrgId field to given value.

### HasFirmOrgId

`func (o *CliaCertificatesRow) HasFirmOrgId() bool`

HasFirmOrgId returns a boolean if a field has been set.

### SetFirmOrgIdNil

`func (o *CliaCertificatesRow) SetFirmOrgIdNil(b bool)`

 SetFirmOrgIdNil sets the value for FirmOrgId to be an explicit nil

### UnsetFirmOrgId
`func (o *CliaCertificatesRow) UnsetFirmOrgId()`

UnsetFirmOrgId ensures that no value is present for FirmOrgId, not even an explicit nil
### GetFirmCompanyId

`func (o *CliaCertificatesRow) GetFirmCompanyId() string`

GetFirmCompanyId returns the FirmCompanyId field if non-nil, zero value otherwise.

### GetFirmCompanyIdOk

`func (o *CliaCertificatesRow) GetFirmCompanyIdOk() (*string, bool)`

GetFirmCompanyIdOk returns a tuple with the FirmCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyId

`func (o *CliaCertificatesRow) SetFirmCompanyId(v string)`

SetFirmCompanyId sets FirmCompanyId field to given value.

### HasFirmCompanyId

`func (o *CliaCertificatesRow) HasFirmCompanyId() bool`

HasFirmCompanyId returns a boolean if a field has been set.

### SetFirmCompanyIdNil

`func (o *CliaCertificatesRow) SetFirmCompanyIdNil(b bool)`

 SetFirmCompanyIdNil sets the value for FirmCompanyId to be an explicit nil

### UnsetFirmCompanyId
`func (o *CliaCertificatesRow) UnsetFirmCompanyId()`

UnsetFirmCompanyId ensures that no value is present for FirmCompanyId, not even an explicit nil
### GetFirmCompanyName

`func (o *CliaCertificatesRow) GetFirmCompanyName() string`

GetFirmCompanyName returns the FirmCompanyName field if non-nil, zero value otherwise.

### GetFirmCompanyNameOk

`func (o *CliaCertificatesRow) GetFirmCompanyNameOk() (*string, bool)`

GetFirmCompanyNameOk returns a tuple with the FirmCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyName

`func (o *CliaCertificatesRow) SetFirmCompanyName(v string)`

SetFirmCompanyName sets FirmCompanyName field to given value.

### HasFirmCompanyName

`func (o *CliaCertificatesRow) HasFirmCompanyName() bool`

HasFirmCompanyName returns a boolean if a field has been set.

### SetFirmCompanyNameNil

`func (o *CliaCertificatesRow) SetFirmCompanyNameNil(b bool)`

 SetFirmCompanyNameNil sets the value for FirmCompanyName to be an explicit nil

### UnsetFirmCompanyName
`func (o *CliaCertificatesRow) UnsetFirmCompanyName()`

UnsetFirmCompanyName ensures that no value is present for FirmCompanyName, not even an explicit nil
### GetFirmGroupId

`func (o *CliaCertificatesRow) GetFirmGroupId() string`

GetFirmGroupId returns the FirmGroupId field if non-nil, zero value otherwise.

### GetFirmGroupIdOk

`func (o *CliaCertificatesRow) GetFirmGroupIdOk() (*string, bool)`

GetFirmGroupIdOk returns a tuple with the FirmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupId

`func (o *CliaCertificatesRow) SetFirmGroupId(v string)`

SetFirmGroupId sets FirmGroupId field to given value.

### HasFirmGroupId

`func (o *CliaCertificatesRow) HasFirmGroupId() bool`

HasFirmGroupId returns a boolean if a field has been set.

### SetFirmGroupIdNil

`func (o *CliaCertificatesRow) SetFirmGroupIdNil(b bool)`

 SetFirmGroupIdNil sets the value for FirmGroupId to be an explicit nil

### UnsetFirmGroupId
`func (o *CliaCertificatesRow) UnsetFirmGroupId()`

UnsetFirmGroupId ensures that no value is present for FirmGroupId, not even an explicit nil
### GetFirmGroupName

`func (o *CliaCertificatesRow) GetFirmGroupName() string`

GetFirmGroupName returns the FirmGroupName field if non-nil, zero value otherwise.

### GetFirmGroupNameOk

`func (o *CliaCertificatesRow) GetFirmGroupNameOk() (*string, bool)`

GetFirmGroupNameOk returns a tuple with the FirmGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupName

`func (o *CliaCertificatesRow) SetFirmGroupName(v string)`

SetFirmGroupName sets FirmGroupName field to given value.

### HasFirmGroupName

`func (o *CliaCertificatesRow) HasFirmGroupName() bool`

HasFirmGroupName returns a boolean if a field has been set.

### SetFirmGroupNameNil

`func (o *CliaCertificatesRow) SetFirmGroupNameNil(b bool)`

 SetFirmGroupNameNil sets the value for FirmGroupName to be an explicit nil

### UnsetFirmGroupName
`func (o *CliaCertificatesRow) UnsetFirmGroupName()`

UnsetFirmGroupName ensures that no value is present for FirmGroupName, not even an explicit nil
### GetMatchTier

`func (o *CliaCertificatesRow) GetMatchTier() string`

GetMatchTier returns the MatchTier field if non-nil, zero value otherwise.

### GetMatchTierOk

`func (o *CliaCertificatesRow) GetMatchTierOk() (*string, bool)`

GetMatchTierOk returns a tuple with the MatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTier

`func (o *CliaCertificatesRow) SetMatchTier(v string)`

SetMatchTier sets MatchTier field to given value.

### HasMatchTier

`func (o *CliaCertificatesRow) HasMatchTier() bool`

HasMatchTier returns a boolean if a field has been set.

### SetMatchTierNil

`func (o *CliaCertificatesRow) SetMatchTierNil(b bool)`

 SetMatchTierNil sets the value for MatchTier to be an explicit nil

### UnsetMatchTier
`func (o *CliaCertificatesRow) UnsetMatchTier()`

UnsetMatchTier ensures that no value is present for MatchTier, not even an explicit nil
### GetMatchConfidence

`func (o *CliaCertificatesRow) GetMatchConfidence() float32`

GetMatchConfidence returns the MatchConfidence field if non-nil, zero value otherwise.

### GetMatchConfidenceOk

`func (o *CliaCertificatesRow) GetMatchConfidenceOk() (*float32, bool)`

GetMatchConfidenceOk returns a tuple with the MatchConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchConfidence

`func (o *CliaCertificatesRow) SetMatchConfidence(v float32)`

SetMatchConfidence sets MatchConfidence field to given value.

### HasMatchConfidence

`func (o *CliaCertificatesRow) HasMatchConfidence() bool`

HasMatchConfidence returns a boolean if a field has been set.

### SetMatchConfidenceNil

`func (o *CliaCertificatesRow) SetMatchConfidenceNil(b bool)`

 SetMatchConfidenceNil sets the value for MatchConfidence to be an explicit nil

### UnsetMatchConfidence
`func (o *CliaCertificatesRow) UnsetMatchConfidence()`

UnsetMatchConfidence ensures that no value is present for MatchConfidence, not even an explicit nil
### GetCity

`func (o *CliaCertificatesRow) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CliaCertificatesRow) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CliaCertificatesRow) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CliaCertificatesRow) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *CliaCertificatesRow) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *CliaCertificatesRow) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetStateRegion

`func (o *CliaCertificatesRow) GetStateRegion() string`

GetStateRegion returns the StateRegion field if non-nil, zero value otherwise.

### GetStateRegionOk

`func (o *CliaCertificatesRow) GetStateRegionOk() (*string, bool)`

GetStateRegionOk returns a tuple with the StateRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRegion

`func (o *CliaCertificatesRow) SetStateRegion(v string)`

SetStateRegion sets StateRegion field to given value.

### HasStateRegion

`func (o *CliaCertificatesRow) HasStateRegion() bool`

HasStateRegion returns a boolean if a field has been set.

### SetStateRegionNil

`func (o *CliaCertificatesRow) SetStateRegionNil(b bool)`

 SetStateRegionNil sets the value for StateRegion to be an explicit nil

### UnsetStateRegion
`func (o *CliaCertificatesRow) UnsetStateRegion()`

UnsetStateRegion ensures that no value is present for StateRegion, not even an explicit nil
### GetPostal5

`func (o *CliaCertificatesRow) GetPostal5() string`

GetPostal5 returns the Postal5 field if non-nil, zero value otherwise.

### GetPostal5Ok

`func (o *CliaCertificatesRow) GetPostal5Ok() (*string, bool)`

GetPostal5Ok returns a tuple with the Postal5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostal5

`func (o *CliaCertificatesRow) SetPostal5(v string)`

SetPostal5 sets Postal5 field to given value.

### HasPostal5

`func (o *CliaCertificatesRow) HasPostal5() bool`

HasPostal5 returns a boolean if a field has been set.

### SetPostal5Nil

`func (o *CliaCertificatesRow) SetPostal5Nil(b bool)`

 SetPostal5Nil sets the value for Postal5 to be an explicit nil

### UnsetPostal5
`func (o *CliaCertificatesRow) UnsetPostal5()`

UnsetPostal5 ensures that no value is present for Postal5, not even an explicit nil
### GetCountry

`func (o *CliaCertificatesRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CliaCertificatesRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CliaCertificatesRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CliaCertificatesRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *CliaCertificatesRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *CliaCertificatesRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLatitude

`func (o *CliaCertificatesRow) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *CliaCertificatesRow) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *CliaCertificatesRow) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *CliaCertificatesRow) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *CliaCertificatesRow) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *CliaCertificatesRow) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *CliaCertificatesRow) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *CliaCertificatesRow) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *CliaCertificatesRow) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *CliaCertificatesRow) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *CliaCertificatesRow) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *CliaCertificatesRow) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetGeocodePrecision

`func (o *CliaCertificatesRow) GetGeocodePrecision() string`

GetGeocodePrecision returns the GeocodePrecision field if non-nil, zero value otherwise.

### GetGeocodePrecisionOk

`func (o *CliaCertificatesRow) GetGeocodePrecisionOk() (*string, bool)`

GetGeocodePrecisionOk returns a tuple with the GeocodePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocodePrecision

`func (o *CliaCertificatesRow) SetGeocodePrecision(v string)`

SetGeocodePrecision sets GeocodePrecision field to given value.

### HasGeocodePrecision

`func (o *CliaCertificatesRow) HasGeocodePrecision() bool`

HasGeocodePrecision returns a boolean if a field has been set.

### SetGeocodePrecisionNil

`func (o *CliaCertificatesRow) SetGeocodePrecisionNil(b bool)`

 SetGeocodePrecisionNil sets the value for GeocodePrecision to be an explicit nil

### UnsetGeocodePrecision
`func (o *CliaCertificatesRow) UnsetGeocodePrecision()`

UnsetGeocodePrecision ensures that no value is present for GeocodePrecision, not even an explicit nil
### GetCoordSource

`func (o *CliaCertificatesRow) GetCoordSource() string`

GetCoordSource returns the CoordSource field if non-nil, zero value otherwise.

### GetCoordSourceOk

`func (o *CliaCertificatesRow) GetCoordSourceOk() (*string, bool)`

GetCoordSourceOk returns a tuple with the CoordSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordSource

`func (o *CliaCertificatesRow) SetCoordSource(v string)`

SetCoordSource sets CoordSource field to given value.

### HasCoordSource

`func (o *CliaCertificatesRow) HasCoordSource() bool`

HasCoordSource returns a boolean if a field has been set.

### SetCoordSourceNil

`func (o *CliaCertificatesRow) SetCoordSourceNil(b bool)`

 SetCoordSourceNil sets the value for CoordSource to be an explicit nil

### UnsetCoordSource
`func (o *CliaCertificatesRow) UnsetCoordSource()`

UnsetCoordSource ensures that no value is present for CoordSource, not even an explicit nil
### GetIsActive

`func (o *CliaCertificatesRow) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CliaCertificatesRow) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CliaCertificatesRow) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CliaCertificatesRow) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *CliaCertificatesRow) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *CliaCertificatesRow) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetCertificateType

`func (o *CliaCertificatesRow) GetCertificateType() string`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *CliaCertificatesRow) GetCertificateTypeOk() (*string, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *CliaCertificatesRow) SetCertificateType(v string)`

SetCertificateType sets CertificateType field to given value.

### HasCertificateType

`func (o *CliaCertificatesRow) HasCertificateType() bool`

HasCertificateType returns a boolean if a field has been set.

### SetCertificateTypeNil

`func (o *CliaCertificatesRow) SetCertificateTypeNil(b bool)`

 SetCertificateTypeNil sets the value for CertificateType to be an explicit nil

### UnsetCertificateType
`func (o *CliaCertificatesRow) UnsetCertificateType()`

UnsetCertificateType ensures that no value is present for CertificateType, not even an explicit nil
### GetComplexityCeiling

`func (o *CliaCertificatesRow) GetComplexityCeiling() string`

GetComplexityCeiling returns the ComplexityCeiling field if non-nil, zero value otherwise.

### GetComplexityCeilingOk

`func (o *CliaCertificatesRow) GetComplexityCeilingOk() (*string, bool)`

GetComplexityCeilingOk returns a tuple with the ComplexityCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexityCeiling

`func (o *CliaCertificatesRow) SetComplexityCeiling(v string)`

SetComplexityCeiling sets ComplexityCeiling field to given value.

### HasComplexityCeiling

`func (o *CliaCertificatesRow) HasComplexityCeiling() bool`

HasComplexityCeiling returns a boolean if a field has been set.

### SetComplexityCeilingNil

`func (o *CliaCertificatesRow) SetComplexityCeilingNil(b bool)`

 SetComplexityCeilingNil sets the value for ComplexityCeiling to be an explicit nil

### UnsetComplexityCeiling
`func (o *CliaCertificatesRow) UnsetComplexityCeiling()`

UnsetComplexityCeiling ensures that no value is present for ComplexityCeiling, not even an explicit nil
### GetAccreditorClaimed

`func (o *CliaCertificatesRow) GetAccreditorClaimed() []string`

GetAccreditorClaimed returns the AccreditorClaimed field if non-nil, zero value otherwise.

### GetAccreditorClaimedOk

`func (o *CliaCertificatesRow) GetAccreditorClaimedOk() (*[]string, bool)`

GetAccreditorClaimedOk returns a tuple with the AccreditorClaimed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccreditorClaimed

`func (o *CliaCertificatesRow) SetAccreditorClaimed(v []string)`

SetAccreditorClaimed sets AccreditorClaimed field to given value.

### HasAccreditorClaimed

`func (o *CliaCertificatesRow) HasAccreditorClaimed() bool`

HasAccreditorClaimed returns a boolean if a field has been set.

### SetAccreditorClaimedNil

`func (o *CliaCertificatesRow) SetAccreditorClaimedNil(b bool)`

 SetAccreditorClaimedNil sets the value for AccreditorClaimed to be an explicit nil

### UnsetAccreditorClaimed
`func (o *CliaCertificatesRow) UnsetAccreditorClaimed()`

UnsetAccreditorClaimed ensures that no value is present for AccreditorClaimed, not even an explicit nil
### GetAccreditorConfirmed

`func (o *CliaCertificatesRow) GetAccreditorConfirmed() []string`

GetAccreditorConfirmed returns the AccreditorConfirmed field if non-nil, zero value otherwise.

### GetAccreditorConfirmedOk

`func (o *CliaCertificatesRow) GetAccreditorConfirmedOk() (*[]string, bool)`

GetAccreditorConfirmedOk returns a tuple with the AccreditorConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccreditorConfirmed

`func (o *CliaCertificatesRow) SetAccreditorConfirmed(v []string)`

SetAccreditorConfirmed sets AccreditorConfirmed field to given value.

### HasAccreditorConfirmed

`func (o *CliaCertificatesRow) HasAccreditorConfirmed() bool`

HasAccreditorConfirmed returns a boolean if a field has been set.

### SetAccreditorConfirmedNil

`func (o *CliaCertificatesRow) SetAccreditorConfirmedNil(b bool)`

 SetAccreditorConfirmedNil sets the value for AccreditorConfirmed to be an explicit nil

### UnsetAccreditorConfirmed
`func (o *CliaCertificatesRow) UnsetAccreditorConfirmed()`

UnsetAccreditorConfirmed ensures that no value is present for AccreditorConfirmed, not even an explicit nil
### GetHasCapAccreditation

`func (o *CliaCertificatesRow) GetHasCapAccreditation() bool`

GetHasCapAccreditation returns the HasCapAccreditation field if non-nil, zero value otherwise.

### GetHasCapAccreditationOk

`func (o *CliaCertificatesRow) GetHasCapAccreditationOk() (*bool, bool)`

GetHasCapAccreditationOk returns a tuple with the HasCapAccreditation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCapAccreditation

`func (o *CliaCertificatesRow) SetHasCapAccreditation(v bool)`

SetHasCapAccreditation sets HasCapAccreditation field to given value.

### HasHasCapAccreditation

`func (o *CliaCertificatesRow) HasHasCapAccreditation() bool`

HasHasCapAccreditation returns a boolean if a field has been set.

### SetHasCapAccreditationNil

`func (o *CliaCertificatesRow) SetHasCapAccreditationNil(b bool)`

 SetHasCapAccreditationNil sets the value for HasCapAccreditation to be an explicit nil

### UnsetHasCapAccreditation
`func (o *CliaCertificatesRow) UnsetHasCapAccreditation()`

UnsetHasCapAccreditation ensures that no value is present for HasCapAccreditation, not even an explicit nil
### GetHasAshiAccreditation

`func (o *CliaCertificatesRow) GetHasAshiAccreditation() bool`

GetHasAshiAccreditation returns the HasAshiAccreditation field if non-nil, zero value otherwise.

### GetHasAshiAccreditationOk

`func (o *CliaCertificatesRow) GetHasAshiAccreditationOk() (*bool, bool)`

GetHasAshiAccreditationOk returns a tuple with the HasAshiAccreditation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAshiAccreditation

`func (o *CliaCertificatesRow) SetHasAshiAccreditation(v bool)`

SetHasAshiAccreditation sets HasAshiAccreditation field to given value.

### HasHasAshiAccreditation

`func (o *CliaCertificatesRow) HasHasAshiAccreditation() bool`

HasHasAshiAccreditation returns a boolean if a field has been set.

### SetHasAshiAccreditationNil

`func (o *CliaCertificatesRow) SetHasAshiAccreditationNil(b bool)`

 SetHasAshiAccreditationNil sets the value for HasAshiAccreditation to be an explicit nil

### UnsetHasAshiAccreditation
`func (o *CliaCertificatesRow) UnsetHasAshiAccreditation()`

UnsetHasAshiAccreditation ensures that no value is present for HasAshiAccreditation, not even an explicit nil
### GetHasAabbAccreditation

`func (o *CliaCertificatesRow) GetHasAabbAccreditation() bool`

GetHasAabbAccreditation returns the HasAabbAccreditation field if non-nil, zero value otherwise.

### GetHasAabbAccreditationOk

`func (o *CliaCertificatesRow) GetHasAabbAccreditationOk() (*bool, bool)`

GetHasAabbAccreditationOk returns a tuple with the HasAabbAccreditation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAabbAccreditation

`func (o *CliaCertificatesRow) SetHasAabbAccreditation(v bool)`

SetHasAabbAccreditation sets HasAabbAccreditation field to given value.

### HasHasAabbAccreditation

`func (o *CliaCertificatesRow) HasHasAabbAccreditation() bool`

HasHasAabbAccreditation returns a boolean if a field has been set.

### SetHasAabbAccreditationNil

`func (o *CliaCertificatesRow) SetHasAabbAccreditationNil(b bool)`

 SetHasAabbAccreditationNil sets the value for HasAabbAccreditation to be an explicit nil

### UnsetHasAabbAccreditation
`func (o *CliaCertificatesRow) UnsetHasAabbAccreditation()`

UnsetHasAabbAccreditation ensures that no value is present for HasAabbAccreditation, not even an explicit nil
### GetFacilityType

`func (o *CliaCertificatesRow) GetFacilityType() string`

GetFacilityType returns the FacilityType field if non-nil, zero value otherwise.

### GetFacilityTypeOk

`func (o *CliaCertificatesRow) GetFacilityTypeOk() (*string, bool)`

GetFacilityTypeOk returns a tuple with the FacilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityType

`func (o *CliaCertificatesRow) SetFacilityType(v string)`

SetFacilityType sets FacilityType field to given value.

### HasFacilityType

`func (o *CliaCertificatesRow) HasFacilityType() bool`

HasFacilityType returns a boolean if a field has been set.

### SetFacilityTypeNil

`func (o *CliaCertificatesRow) SetFacilityTypeNil(b bool)`

 SetFacilityTypeNil sets the value for FacilityType to be an explicit nil

### UnsetFacilityType
`func (o *CliaCertificatesRow) UnsetFacilityType()`

UnsetFacilityType ensures that no value is present for FacilityType, not even an explicit nil
### GetAnnualTestVolume

`func (o *CliaCertificatesRow) GetAnnualTestVolume() int64`

GetAnnualTestVolume returns the AnnualTestVolume field if non-nil, zero value otherwise.

### GetAnnualTestVolumeOk

`func (o *CliaCertificatesRow) GetAnnualTestVolumeOk() (*int64, bool)`

GetAnnualTestVolumeOk returns a tuple with the AnnualTestVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualTestVolume

`func (o *CliaCertificatesRow) SetAnnualTestVolume(v int64)`

SetAnnualTestVolume sets AnnualTestVolume field to given value.

### HasAnnualTestVolume

`func (o *CliaCertificatesRow) HasAnnualTestVolume() bool`

HasAnnualTestVolume returns a boolean if a field has been set.

### SetAnnualTestVolumeNil

`func (o *CliaCertificatesRow) SetAnnualTestVolumeNil(b bool)`

 SetAnnualTestVolumeNil sets the value for AnnualTestVolume to be an explicit nil

### UnsetAnnualTestVolume
`func (o *CliaCertificatesRow) UnsetAnnualTestVolume()`

UnsetAnnualTestVolume ensures that no value is present for AnnualTestVolume, not even an explicit nil
### GetTestVolumeSource

`func (o *CliaCertificatesRow) GetTestVolumeSource() string`

GetTestVolumeSource returns the TestVolumeSource field if non-nil, zero value otherwise.

### GetTestVolumeSourceOk

`func (o *CliaCertificatesRow) GetTestVolumeSourceOk() (*string, bool)`

GetTestVolumeSourceOk returns a tuple with the TestVolumeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestVolumeSource

`func (o *CliaCertificatesRow) SetTestVolumeSource(v string)`

SetTestVolumeSource sets TestVolumeSource field to given value.

### HasTestVolumeSource

`func (o *CliaCertificatesRow) HasTestVolumeSource() bool`

HasTestVolumeSource returns a boolean if a field has been set.

### SetTestVolumeSourceNil

`func (o *CliaCertificatesRow) SetTestVolumeSourceNil(b bool)`

 SetTestVolumeSourceNil sets the value for TestVolumeSource to be an explicit nil

### UnsetTestVolumeSource
`func (o *CliaCertificatesRow) UnsetTestVolumeSource()`

UnsetTestVolumeSource ensures that no value is present for TestVolumeSource, not even an explicit nil
### GetLabSiteCount

`func (o *CliaCertificatesRow) GetLabSiteCount() int32`

GetLabSiteCount returns the LabSiteCount field if non-nil, zero value otherwise.

### GetLabSiteCountOk

`func (o *CliaCertificatesRow) GetLabSiteCountOk() (*int32, bool)`

GetLabSiteCountOk returns a tuple with the LabSiteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabSiteCount

`func (o *CliaCertificatesRow) SetLabSiteCount(v int32)`

SetLabSiteCount sets LabSiteCount field to given value.

### HasLabSiteCount

`func (o *CliaCertificatesRow) HasLabSiteCount() bool`

HasLabSiteCount returns a boolean if a field has been set.

### SetLabSiteCountNil

`func (o *CliaCertificatesRow) SetLabSiteCountNil(b bool)`

 SetLabSiteCountNil sets the value for LabSiteCount to be an explicit nil

### UnsetLabSiteCount
`func (o *CliaCertificatesRow) UnsetLabSiteCount()`

UnsetLabSiteCount ensures that no value is present for LabSiteCount, not even an explicit nil
### GetPosVintage

`func (o *CliaCertificatesRow) GetPosVintage() string`

GetPosVintage returns the PosVintage field if non-nil, zero value otherwise.

### GetPosVintageOk

`func (o *CliaCertificatesRow) GetPosVintageOk() (*string, bool)`

GetPosVintageOk returns a tuple with the PosVintage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosVintage

`func (o *CliaCertificatesRow) SetPosVintage(v string)`

SetPosVintage sets PosVintage field to given value.

### HasPosVintage

`func (o *CliaCertificatesRow) HasPosVintage() bool`

HasPosVintage returns a boolean if a field has been set.

### SetPosVintageNil

`func (o *CliaCertificatesRow) SetPosVintageNil(b bool)`

 SetPosVintageNil sets the value for PosVintage to be an explicit nil

### UnsetPosVintage
`func (o *CliaCertificatesRow) UnsetPosVintage()`

UnsetPosVintage ensures that no value is present for PosVintage, not even an explicit nil
### GetAsOf

`func (o *CliaCertificatesRow) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *CliaCertificatesRow) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *CliaCertificatesRow) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *CliaCertificatesRow) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *CliaCertificatesRow) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *CliaCertificatesRow) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


