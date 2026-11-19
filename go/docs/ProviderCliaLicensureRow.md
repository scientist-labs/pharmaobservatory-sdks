# ProviderCliaLicensureRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | Pointer to **NullableString** | The supplier id: the Scientist.com provider id as text, the id in the supplier page address (/suppliers/&lt;id&gt;). The entry key for the suppliers object type. | [optional] 
**ProviderUuid** | Pointer to **NullableString** | The supplier&#39;s public Scientist.com uuid, the join key to the RX marketplace. Present on every row, including providers the marketplace does not publish. | [optional] 
**ProviderSlug** | Pointer to **NullableString** | The supplier&#39;s marketplace slug. | [optional] 
**ProviderName** | Pointer to **NullableString** | The supplier&#39;s name as listed on the marketplace. | [optional] 
**CliaNumber** | Pointer to **NullableString** | The ten-character CLIA certificate number. | [optional] 
**LabName** | Pointer to **NullableString** | The laboratory&#39;s name as listed on the certificate. | [optional] 
**MatchMethod** | Pointer to **NullableString** | How the certificate was attached to the supplier: org_member_exact (through the supplier&#39;s resolved organization) or rx_name_key_exact (an exact normalized-name match, not a verified identity). | [optional] 
**MatchConfidence** | Pointer to **NullableFloat32** | The organization match&#39;s confidence, 0 to 1, on an org_member_exact row; null on an rx_name_key_exact row, whose exact name match defines none. | [optional] 
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
**SourceLicense** | Pointer to **NullableString** | The license of the source file on every row. | [optional] 
**Attribution** | Pointer to **NullableString** | The source credit line on every row. | [optional] 

## Methods

### NewProviderCliaLicensureRow

`func NewProviderCliaLicensureRow() *ProviderCliaLicensureRow`

NewProviderCliaLicensureRow instantiates a new ProviderCliaLicensureRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderCliaLicensureRowWithDefaults

`func NewProviderCliaLicensureRowWithDefaults() *ProviderCliaLicensureRow`

NewProviderCliaLicensureRowWithDefaults instantiates a new ProviderCliaLicensureRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *ProviderCliaLicensureRow) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ProviderCliaLicensureRow) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ProviderCliaLicensureRow) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ProviderCliaLicensureRow) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *ProviderCliaLicensureRow) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *ProviderCliaLicensureRow) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetProviderUuid

`func (o *ProviderCliaLicensureRow) GetProviderUuid() string`

GetProviderUuid returns the ProviderUuid field if non-nil, zero value otherwise.

### GetProviderUuidOk

`func (o *ProviderCliaLicensureRow) GetProviderUuidOk() (*string, bool)`

GetProviderUuidOk returns a tuple with the ProviderUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUuid

`func (o *ProviderCliaLicensureRow) SetProviderUuid(v string)`

SetProviderUuid sets ProviderUuid field to given value.

### HasProviderUuid

`func (o *ProviderCliaLicensureRow) HasProviderUuid() bool`

HasProviderUuid returns a boolean if a field has been set.

### SetProviderUuidNil

`func (o *ProviderCliaLicensureRow) SetProviderUuidNil(b bool)`

 SetProviderUuidNil sets the value for ProviderUuid to be an explicit nil

### UnsetProviderUuid
`func (o *ProviderCliaLicensureRow) UnsetProviderUuid()`

UnsetProviderUuid ensures that no value is present for ProviderUuid, not even an explicit nil
### GetProviderSlug

`func (o *ProviderCliaLicensureRow) GetProviderSlug() string`

GetProviderSlug returns the ProviderSlug field if non-nil, zero value otherwise.

### GetProviderSlugOk

`func (o *ProviderCliaLicensureRow) GetProviderSlugOk() (*string, bool)`

GetProviderSlugOk returns a tuple with the ProviderSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSlug

`func (o *ProviderCliaLicensureRow) SetProviderSlug(v string)`

SetProviderSlug sets ProviderSlug field to given value.

### HasProviderSlug

`func (o *ProviderCliaLicensureRow) HasProviderSlug() bool`

HasProviderSlug returns a boolean if a field has been set.

### SetProviderSlugNil

`func (o *ProviderCliaLicensureRow) SetProviderSlugNil(b bool)`

 SetProviderSlugNil sets the value for ProviderSlug to be an explicit nil

### UnsetProviderSlug
`func (o *ProviderCliaLicensureRow) UnsetProviderSlug()`

UnsetProviderSlug ensures that no value is present for ProviderSlug, not even an explicit nil
### GetProviderName

`func (o *ProviderCliaLicensureRow) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ProviderCliaLicensureRow) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ProviderCliaLicensureRow) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ProviderCliaLicensureRow) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *ProviderCliaLicensureRow) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *ProviderCliaLicensureRow) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetCliaNumber

`func (o *ProviderCliaLicensureRow) GetCliaNumber() string`

GetCliaNumber returns the CliaNumber field if non-nil, zero value otherwise.

### GetCliaNumberOk

`func (o *ProviderCliaLicensureRow) GetCliaNumberOk() (*string, bool)`

GetCliaNumberOk returns a tuple with the CliaNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliaNumber

`func (o *ProviderCliaLicensureRow) SetCliaNumber(v string)`

SetCliaNumber sets CliaNumber field to given value.

### HasCliaNumber

`func (o *ProviderCliaLicensureRow) HasCliaNumber() bool`

HasCliaNumber returns a boolean if a field has been set.

### SetCliaNumberNil

`func (o *ProviderCliaLicensureRow) SetCliaNumberNil(b bool)`

 SetCliaNumberNil sets the value for CliaNumber to be an explicit nil

### UnsetCliaNumber
`func (o *ProviderCliaLicensureRow) UnsetCliaNumber()`

UnsetCliaNumber ensures that no value is present for CliaNumber, not even an explicit nil
### GetLabName

`func (o *ProviderCliaLicensureRow) GetLabName() string`

GetLabName returns the LabName field if non-nil, zero value otherwise.

### GetLabNameOk

`func (o *ProviderCliaLicensureRow) GetLabNameOk() (*string, bool)`

GetLabNameOk returns a tuple with the LabName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabName

`func (o *ProviderCliaLicensureRow) SetLabName(v string)`

SetLabName sets LabName field to given value.

### HasLabName

`func (o *ProviderCliaLicensureRow) HasLabName() bool`

HasLabName returns a boolean if a field has been set.

### SetLabNameNil

`func (o *ProviderCliaLicensureRow) SetLabNameNil(b bool)`

 SetLabNameNil sets the value for LabName to be an explicit nil

### UnsetLabName
`func (o *ProviderCliaLicensureRow) UnsetLabName()`

UnsetLabName ensures that no value is present for LabName, not even an explicit nil
### GetMatchMethod

`func (o *ProviderCliaLicensureRow) GetMatchMethod() string`

GetMatchMethod returns the MatchMethod field if non-nil, zero value otherwise.

### GetMatchMethodOk

`func (o *ProviderCliaLicensureRow) GetMatchMethodOk() (*string, bool)`

GetMatchMethodOk returns a tuple with the MatchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMethod

`func (o *ProviderCliaLicensureRow) SetMatchMethod(v string)`

SetMatchMethod sets MatchMethod field to given value.

### HasMatchMethod

`func (o *ProviderCliaLicensureRow) HasMatchMethod() bool`

HasMatchMethod returns a boolean if a field has been set.

### SetMatchMethodNil

`func (o *ProviderCliaLicensureRow) SetMatchMethodNil(b bool)`

 SetMatchMethodNil sets the value for MatchMethod to be an explicit nil

### UnsetMatchMethod
`func (o *ProviderCliaLicensureRow) UnsetMatchMethod()`

UnsetMatchMethod ensures that no value is present for MatchMethod, not even an explicit nil
### GetMatchConfidence

`func (o *ProviderCliaLicensureRow) GetMatchConfidence() float32`

GetMatchConfidence returns the MatchConfidence field if non-nil, zero value otherwise.

### GetMatchConfidenceOk

`func (o *ProviderCliaLicensureRow) GetMatchConfidenceOk() (*float32, bool)`

GetMatchConfidenceOk returns a tuple with the MatchConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchConfidence

`func (o *ProviderCliaLicensureRow) SetMatchConfidence(v float32)`

SetMatchConfidence sets MatchConfidence field to given value.

### HasMatchConfidence

`func (o *ProviderCliaLicensureRow) HasMatchConfidence() bool`

HasMatchConfidence returns a boolean if a field has been set.

### SetMatchConfidenceNil

`func (o *ProviderCliaLicensureRow) SetMatchConfidenceNil(b bool)`

 SetMatchConfidenceNil sets the value for MatchConfidence to be an explicit nil

### UnsetMatchConfidence
`func (o *ProviderCliaLicensureRow) UnsetMatchConfidence()`

UnsetMatchConfidence ensures that no value is present for MatchConfidence, not even an explicit nil
### GetCity

`func (o *ProviderCliaLicensureRow) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ProviderCliaLicensureRow) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ProviderCliaLicensureRow) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ProviderCliaLicensureRow) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ProviderCliaLicensureRow) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ProviderCliaLicensureRow) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetStateRegion

`func (o *ProviderCliaLicensureRow) GetStateRegion() string`

GetStateRegion returns the StateRegion field if non-nil, zero value otherwise.

### GetStateRegionOk

`func (o *ProviderCliaLicensureRow) GetStateRegionOk() (*string, bool)`

GetStateRegionOk returns a tuple with the StateRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRegion

`func (o *ProviderCliaLicensureRow) SetStateRegion(v string)`

SetStateRegion sets StateRegion field to given value.

### HasStateRegion

`func (o *ProviderCliaLicensureRow) HasStateRegion() bool`

HasStateRegion returns a boolean if a field has been set.

### SetStateRegionNil

`func (o *ProviderCliaLicensureRow) SetStateRegionNil(b bool)`

 SetStateRegionNil sets the value for StateRegion to be an explicit nil

### UnsetStateRegion
`func (o *ProviderCliaLicensureRow) UnsetStateRegion()`

UnsetStateRegion ensures that no value is present for StateRegion, not even an explicit nil
### GetPostal5

`func (o *ProviderCliaLicensureRow) GetPostal5() string`

GetPostal5 returns the Postal5 field if non-nil, zero value otherwise.

### GetPostal5Ok

`func (o *ProviderCliaLicensureRow) GetPostal5Ok() (*string, bool)`

GetPostal5Ok returns a tuple with the Postal5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostal5

`func (o *ProviderCliaLicensureRow) SetPostal5(v string)`

SetPostal5 sets Postal5 field to given value.

### HasPostal5

`func (o *ProviderCliaLicensureRow) HasPostal5() bool`

HasPostal5 returns a boolean if a field has been set.

### SetPostal5Nil

`func (o *ProviderCliaLicensureRow) SetPostal5Nil(b bool)`

 SetPostal5Nil sets the value for Postal5 to be an explicit nil

### UnsetPostal5
`func (o *ProviderCliaLicensureRow) UnsetPostal5()`

UnsetPostal5 ensures that no value is present for Postal5, not even an explicit nil
### GetCountry

`func (o *ProviderCliaLicensureRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ProviderCliaLicensureRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ProviderCliaLicensureRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ProviderCliaLicensureRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *ProviderCliaLicensureRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *ProviderCliaLicensureRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLatitude

`func (o *ProviderCliaLicensureRow) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *ProviderCliaLicensureRow) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *ProviderCliaLicensureRow) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *ProviderCliaLicensureRow) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *ProviderCliaLicensureRow) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *ProviderCliaLicensureRow) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *ProviderCliaLicensureRow) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *ProviderCliaLicensureRow) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *ProviderCliaLicensureRow) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *ProviderCliaLicensureRow) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *ProviderCliaLicensureRow) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *ProviderCliaLicensureRow) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetGeocodePrecision

`func (o *ProviderCliaLicensureRow) GetGeocodePrecision() string`

GetGeocodePrecision returns the GeocodePrecision field if non-nil, zero value otherwise.

### GetGeocodePrecisionOk

`func (o *ProviderCliaLicensureRow) GetGeocodePrecisionOk() (*string, bool)`

GetGeocodePrecisionOk returns a tuple with the GeocodePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocodePrecision

`func (o *ProviderCliaLicensureRow) SetGeocodePrecision(v string)`

SetGeocodePrecision sets GeocodePrecision field to given value.

### HasGeocodePrecision

`func (o *ProviderCliaLicensureRow) HasGeocodePrecision() bool`

HasGeocodePrecision returns a boolean if a field has been set.

### SetGeocodePrecisionNil

`func (o *ProviderCliaLicensureRow) SetGeocodePrecisionNil(b bool)`

 SetGeocodePrecisionNil sets the value for GeocodePrecision to be an explicit nil

### UnsetGeocodePrecision
`func (o *ProviderCliaLicensureRow) UnsetGeocodePrecision()`

UnsetGeocodePrecision ensures that no value is present for GeocodePrecision, not even an explicit nil
### GetCoordSource

`func (o *ProviderCliaLicensureRow) GetCoordSource() string`

GetCoordSource returns the CoordSource field if non-nil, zero value otherwise.

### GetCoordSourceOk

`func (o *ProviderCliaLicensureRow) GetCoordSourceOk() (*string, bool)`

GetCoordSourceOk returns a tuple with the CoordSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordSource

`func (o *ProviderCliaLicensureRow) SetCoordSource(v string)`

SetCoordSource sets CoordSource field to given value.

### HasCoordSource

`func (o *ProviderCliaLicensureRow) HasCoordSource() bool`

HasCoordSource returns a boolean if a field has been set.

### SetCoordSourceNil

`func (o *ProviderCliaLicensureRow) SetCoordSourceNil(b bool)`

 SetCoordSourceNil sets the value for CoordSource to be an explicit nil

### UnsetCoordSource
`func (o *ProviderCliaLicensureRow) UnsetCoordSource()`

UnsetCoordSource ensures that no value is present for CoordSource, not even an explicit nil
### GetIsActive

`func (o *ProviderCliaLicensureRow) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProviderCliaLicensureRow) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProviderCliaLicensureRow) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ProviderCliaLicensureRow) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *ProviderCliaLicensureRow) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ProviderCliaLicensureRow) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetCertificateType

`func (o *ProviderCliaLicensureRow) GetCertificateType() string`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *ProviderCliaLicensureRow) GetCertificateTypeOk() (*string, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *ProviderCliaLicensureRow) SetCertificateType(v string)`

SetCertificateType sets CertificateType field to given value.

### HasCertificateType

`func (o *ProviderCliaLicensureRow) HasCertificateType() bool`

HasCertificateType returns a boolean if a field has been set.

### SetCertificateTypeNil

`func (o *ProviderCliaLicensureRow) SetCertificateTypeNil(b bool)`

 SetCertificateTypeNil sets the value for CertificateType to be an explicit nil

### UnsetCertificateType
`func (o *ProviderCliaLicensureRow) UnsetCertificateType()`

UnsetCertificateType ensures that no value is present for CertificateType, not even an explicit nil
### GetComplexityCeiling

`func (o *ProviderCliaLicensureRow) GetComplexityCeiling() string`

GetComplexityCeiling returns the ComplexityCeiling field if non-nil, zero value otherwise.

### GetComplexityCeilingOk

`func (o *ProviderCliaLicensureRow) GetComplexityCeilingOk() (*string, bool)`

GetComplexityCeilingOk returns a tuple with the ComplexityCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexityCeiling

`func (o *ProviderCliaLicensureRow) SetComplexityCeiling(v string)`

SetComplexityCeiling sets ComplexityCeiling field to given value.

### HasComplexityCeiling

`func (o *ProviderCliaLicensureRow) HasComplexityCeiling() bool`

HasComplexityCeiling returns a boolean if a field has been set.

### SetComplexityCeilingNil

`func (o *ProviderCliaLicensureRow) SetComplexityCeilingNil(b bool)`

 SetComplexityCeilingNil sets the value for ComplexityCeiling to be an explicit nil

### UnsetComplexityCeiling
`func (o *ProviderCliaLicensureRow) UnsetComplexityCeiling()`

UnsetComplexityCeiling ensures that no value is present for ComplexityCeiling, not even an explicit nil
### GetAccreditorClaimed

`func (o *ProviderCliaLicensureRow) GetAccreditorClaimed() []string`

GetAccreditorClaimed returns the AccreditorClaimed field if non-nil, zero value otherwise.

### GetAccreditorClaimedOk

`func (o *ProviderCliaLicensureRow) GetAccreditorClaimedOk() (*[]string, bool)`

GetAccreditorClaimedOk returns a tuple with the AccreditorClaimed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccreditorClaimed

`func (o *ProviderCliaLicensureRow) SetAccreditorClaimed(v []string)`

SetAccreditorClaimed sets AccreditorClaimed field to given value.

### HasAccreditorClaimed

`func (o *ProviderCliaLicensureRow) HasAccreditorClaimed() bool`

HasAccreditorClaimed returns a boolean if a field has been set.

### SetAccreditorClaimedNil

`func (o *ProviderCliaLicensureRow) SetAccreditorClaimedNil(b bool)`

 SetAccreditorClaimedNil sets the value for AccreditorClaimed to be an explicit nil

### UnsetAccreditorClaimed
`func (o *ProviderCliaLicensureRow) UnsetAccreditorClaimed()`

UnsetAccreditorClaimed ensures that no value is present for AccreditorClaimed, not even an explicit nil
### GetAccreditorConfirmed

`func (o *ProviderCliaLicensureRow) GetAccreditorConfirmed() []string`

GetAccreditorConfirmed returns the AccreditorConfirmed field if non-nil, zero value otherwise.

### GetAccreditorConfirmedOk

`func (o *ProviderCliaLicensureRow) GetAccreditorConfirmedOk() (*[]string, bool)`

GetAccreditorConfirmedOk returns a tuple with the AccreditorConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccreditorConfirmed

`func (o *ProviderCliaLicensureRow) SetAccreditorConfirmed(v []string)`

SetAccreditorConfirmed sets AccreditorConfirmed field to given value.

### HasAccreditorConfirmed

`func (o *ProviderCliaLicensureRow) HasAccreditorConfirmed() bool`

HasAccreditorConfirmed returns a boolean if a field has been set.

### SetAccreditorConfirmedNil

`func (o *ProviderCliaLicensureRow) SetAccreditorConfirmedNil(b bool)`

 SetAccreditorConfirmedNil sets the value for AccreditorConfirmed to be an explicit nil

### UnsetAccreditorConfirmed
`func (o *ProviderCliaLicensureRow) UnsetAccreditorConfirmed()`

UnsetAccreditorConfirmed ensures that no value is present for AccreditorConfirmed, not even an explicit nil
### GetHasCapAccreditation

`func (o *ProviderCliaLicensureRow) GetHasCapAccreditation() bool`

GetHasCapAccreditation returns the HasCapAccreditation field if non-nil, zero value otherwise.

### GetHasCapAccreditationOk

`func (o *ProviderCliaLicensureRow) GetHasCapAccreditationOk() (*bool, bool)`

GetHasCapAccreditationOk returns a tuple with the HasCapAccreditation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCapAccreditation

`func (o *ProviderCliaLicensureRow) SetHasCapAccreditation(v bool)`

SetHasCapAccreditation sets HasCapAccreditation field to given value.

### HasHasCapAccreditation

`func (o *ProviderCliaLicensureRow) HasHasCapAccreditation() bool`

HasHasCapAccreditation returns a boolean if a field has been set.

### SetHasCapAccreditationNil

`func (o *ProviderCliaLicensureRow) SetHasCapAccreditationNil(b bool)`

 SetHasCapAccreditationNil sets the value for HasCapAccreditation to be an explicit nil

### UnsetHasCapAccreditation
`func (o *ProviderCliaLicensureRow) UnsetHasCapAccreditation()`

UnsetHasCapAccreditation ensures that no value is present for HasCapAccreditation, not even an explicit nil
### GetHasAshiAccreditation

`func (o *ProviderCliaLicensureRow) GetHasAshiAccreditation() bool`

GetHasAshiAccreditation returns the HasAshiAccreditation field if non-nil, zero value otherwise.

### GetHasAshiAccreditationOk

`func (o *ProviderCliaLicensureRow) GetHasAshiAccreditationOk() (*bool, bool)`

GetHasAshiAccreditationOk returns a tuple with the HasAshiAccreditation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAshiAccreditation

`func (o *ProviderCliaLicensureRow) SetHasAshiAccreditation(v bool)`

SetHasAshiAccreditation sets HasAshiAccreditation field to given value.

### HasHasAshiAccreditation

`func (o *ProviderCliaLicensureRow) HasHasAshiAccreditation() bool`

HasHasAshiAccreditation returns a boolean if a field has been set.

### SetHasAshiAccreditationNil

`func (o *ProviderCliaLicensureRow) SetHasAshiAccreditationNil(b bool)`

 SetHasAshiAccreditationNil sets the value for HasAshiAccreditation to be an explicit nil

### UnsetHasAshiAccreditation
`func (o *ProviderCliaLicensureRow) UnsetHasAshiAccreditation()`

UnsetHasAshiAccreditation ensures that no value is present for HasAshiAccreditation, not even an explicit nil
### GetHasAabbAccreditation

`func (o *ProviderCliaLicensureRow) GetHasAabbAccreditation() bool`

GetHasAabbAccreditation returns the HasAabbAccreditation field if non-nil, zero value otherwise.

### GetHasAabbAccreditationOk

`func (o *ProviderCliaLicensureRow) GetHasAabbAccreditationOk() (*bool, bool)`

GetHasAabbAccreditationOk returns a tuple with the HasAabbAccreditation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAabbAccreditation

`func (o *ProviderCliaLicensureRow) SetHasAabbAccreditation(v bool)`

SetHasAabbAccreditation sets HasAabbAccreditation field to given value.

### HasHasAabbAccreditation

`func (o *ProviderCliaLicensureRow) HasHasAabbAccreditation() bool`

HasHasAabbAccreditation returns a boolean if a field has been set.

### SetHasAabbAccreditationNil

`func (o *ProviderCliaLicensureRow) SetHasAabbAccreditationNil(b bool)`

 SetHasAabbAccreditationNil sets the value for HasAabbAccreditation to be an explicit nil

### UnsetHasAabbAccreditation
`func (o *ProviderCliaLicensureRow) UnsetHasAabbAccreditation()`

UnsetHasAabbAccreditation ensures that no value is present for HasAabbAccreditation, not even an explicit nil
### GetFacilityType

`func (o *ProviderCliaLicensureRow) GetFacilityType() string`

GetFacilityType returns the FacilityType field if non-nil, zero value otherwise.

### GetFacilityTypeOk

`func (o *ProviderCliaLicensureRow) GetFacilityTypeOk() (*string, bool)`

GetFacilityTypeOk returns a tuple with the FacilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityType

`func (o *ProviderCliaLicensureRow) SetFacilityType(v string)`

SetFacilityType sets FacilityType field to given value.

### HasFacilityType

`func (o *ProviderCliaLicensureRow) HasFacilityType() bool`

HasFacilityType returns a boolean if a field has been set.

### SetFacilityTypeNil

`func (o *ProviderCliaLicensureRow) SetFacilityTypeNil(b bool)`

 SetFacilityTypeNil sets the value for FacilityType to be an explicit nil

### UnsetFacilityType
`func (o *ProviderCliaLicensureRow) UnsetFacilityType()`

UnsetFacilityType ensures that no value is present for FacilityType, not even an explicit nil
### GetAnnualTestVolume

`func (o *ProviderCliaLicensureRow) GetAnnualTestVolume() int64`

GetAnnualTestVolume returns the AnnualTestVolume field if non-nil, zero value otherwise.

### GetAnnualTestVolumeOk

`func (o *ProviderCliaLicensureRow) GetAnnualTestVolumeOk() (*int64, bool)`

GetAnnualTestVolumeOk returns a tuple with the AnnualTestVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualTestVolume

`func (o *ProviderCliaLicensureRow) SetAnnualTestVolume(v int64)`

SetAnnualTestVolume sets AnnualTestVolume field to given value.

### HasAnnualTestVolume

`func (o *ProviderCliaLicensureRow) HasAnnualTestVolume() bool`

HasAnnualTestVolume returns a boolean if a field has been set.

### SetAnnualTestVolumeNil

`func (o *ProviderCliaLicensureRow) SetAnnualTestVolumeNil(b bool)`

 SetAnnualTestVolumeNil sets the value for AnnualTestVolume to be an explicit nil

### UnsetAnnualTestVolume
`func (o *ProviderCliaLicensureRow) UnsetAnnualTestVolume()`

UnsetAnnualTestVolume ensures that no value is present for AnnualTestVolume, not even an explicit nil
### GetTestVolumeSource

`func (o *ProviderCliaLicensureRow) GetTestVolumeSource() string`

GetTestVolumeSource returns the TestVolumeSource field if non-nil, zero value otherwise.

### GetTestVolumeSourceOk

`func (o *ProviderCliaLicensureRow) GetTestVolumeSourceOk() (*string, bool)`

GetTestVolumeSourceOk returns a tuple with the TestVolumeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestVolumeSource

`func (o *ProviderCliaLicensureRow) SetTestVolumeSource(v string)`

SetTestVolumeSource sets TestVolumeSource field to given value.

### HasTestVolumeSource

`func (o *ProviderCliaLicensureRow) HasTestVolumeSource() bool`

HasTestVolumeSource returns a boolean if a field has been set.

### SetTestVolumeSourceNil

`func (o *ProviderCliaLicensureRow) SetTestVolumeSourceNil(b bool)`

 SetTestVolumeSourceNil sets the value for TestVolumeSource to be an explicit nil

### UnsetTestVolumeSource
`func (o *ProviderCliaLicensureRow) UnsetTestVolumeSource()`

UnsetTestVolumeSource ensures that no value is present for TestVolumeSource, not even an explicit nil
### GetLabSiteCount

`func (o *ProviderCliaLicensureRow) GetLabSiteCount() int32`

GetLabSiteCount returns the LabSiteCount field if non-nil, zero value otherwise.

### GetLabSiteCountOk

`func (o *ProviderCliaLicensureRow) GetLabSiteCountOk() (*int32, bool)`

GetLabSiteCountOk returns a tuple with the LabSiteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabSiteCount

`func (o *ProviderCliaLicensureRow) SetLabSiteCount(v int32)`

SetLabSiteCount sets LabSiteCount field to given value.

### HasLabSiteCount

`func (o *ProviderCliaLicensureRow) HasLabSiteCount() bool`

HasLabSiteCount returns a boolean if a field has been set.

### SetLabSiteCountNil

`func (o *ProviderCliaLicensureRow) SetLabSiteCountNil(b bool)`

 SetLabSiteCountNil sets the value for LabSiteCount to be an explicit nil

### UnsetLabSiteCount
`func (o *ProviderCliaLicensureRow) UnsetLabSiteCount()`

UnsetLabSiteCount ensures that no value is present for LabSiteCount, not even an explicit nil
### GetPosVintage

`func (o *ProviderCliaLicensureRow) GetPosVintage() string`

GetPosVintage returns the PosVintage field if non-nil, zero value otherwise.

### GetPosVintageOk

`func (o *ProviderCliaLicensureRow) GetPosVintageOk() (*string, bool)`

GetPosVintageOk returns a tuple with the PosVintage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosVintage

`func (o *ProviderCliaLicensureRow) SetPosVintage(v string)`

SetPosVintage sets PosVintage field to given value.

### HasPosVintage

`func (o *ProviderCliaLicensureRow) HasPosVintage() bool`

HasPosVintage returns a boolean if a field has been set.

### SetPosVintageNil

`func (o *ProviderCliaLicensureRow) SetPosVintageNil(b bool)`

 SetPosVintageNil sets the value for PosVintage to be an explicit nil

### UnsetPosVintage
`func (o *ProviderCliaLicensureRow) UnsetPosVintage()`

UnsetPosVintage ensures that no value is present for PosVintage, not even an explicit nil
### GetAsOf

`func (o *ProviderCliaLicensureRow) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *ProviderCliaLicensureRow) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *ProviderCliaLicensureRow) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *ProviderCliaLicensureRow) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *ProviderCliaLicensureRow) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *ProviderCliaLicensureRow) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil
### GetSourceLicense

`func (o *ProviderCliaLicensureRow) GetSourceLicense() string`

GetSourceLicense returns the SourceLicense field if non-nil, zero value otherwise.

### GetSourceLicenseOk

`func (o *ProviderCliaLicensureRow) GetSourceLicenseOk() (*string, bool)`

GetSourceLicenseOk returns a tuple with the SourceLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLicense

`func (o *ProviderCliaLicensureRow) SetSourceLicense(v string)`

SetSourceLicense sets SourceLicense field to given value.

### HasSourceLicense

`func (o *ProviderCliaLicensureRow) HasSourceLicense() bool`

HasSourceLicense returns a boolean if a field has been set.

### SetSourceLicenseNil

`func (o *ProviderCliaLicensureRow) SetSourceLicenseNil(b bool)`

 SetSourceLicenseNil sets the value for SourceLicense to be an explicit nil

### UnsetSourceLicense
`func (o *ProviderCliaLicensureRow) UnsetSourceLicense()`

UnsetSourceLicense ensures that no value is present for SourceLicense, not even an explicit nil
### GetAttribution

`func (o *ProviderCliaLicensureRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *ProviderCliaLicensureRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *ProviderCliaLicensureRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *ProviderCliaLicensureRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *ProviderCliaLicensureRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *ProviderCliaLicensureRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


