# GxpFacilitiesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FacilityKey** | Pointer to **NullableString** | The canonical facility id. | [optional] 
**RegistryFacilityKeys** | Pointer to **[]string** | The registry records the facility is made of, each as registry:native site id (for example fda_fei:3004123456), sorted. | [optional] 
**FacilityName** | Pointer to **NullableString** | The facility&#39;s name. | [optional] 
**MemberCount** | Pointer to **NullableInt64** | How many registry records are clustered into the facility. | [optional] 
**AddressLine** | Pointer to **NullableString** | The street address. | [optional] 
**City** | Pointer to **NullableString** | The city. | [optional] 
**StateRegion** | Pointer to **NullableString** | The state or region. | [optional] 
**Postal** | Pointer to **NullableString** | The postal code. | [optional] 
**Country** | Pointer to **NullableString** | The two-letter country code. | [optional] 
**Latitude** | Pointer to **NullableFloat32** | Latitude in degrees (WGS84), for the map tool. | [optional] 
**Longitude** | Pointer to **NullableFloat32** | Longitude in degrees (WGS84), for the map tool. | [optional] 
**GeocodePrecision** | Pointer to **NullableString** | What a pin locates: for example a street address, a postal code or only a city. | [optional] 
**CoordSource** | Pointer to **NullableString** | Where the coordinates came from: geonames (GeoNames) or als (Amazon Location Service); it names the attribution to show. | [optional] 
**FirmName** | Pointer to **NullableString** | The operating firm&#39;s name where the facility resolved to an organization; the facility name where it did not. | [optional] 
**FirmOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] 
**FirmCompanyId** | Pointer to **NullableString** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] 
**FirmCompanyName** | Pointer to **NullableString** | Display name of the firm&#39;s company. | [optional] 
**FirmGroupId** | Pointer to **NullableString** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**FirmGroupName** | Pointer to **NullableString** | Display name of the firm&#39;s group. | [optional] 
**IsOrgResolved** | Pointer to **NullableBool** | True when the facility was attached to an operating organization. A false row is kept with null firm ids. | [optional] 
**MatchTier** | Pointer to **NullableString** | How the operating firm was attached: registry (the firm&#39;s own registry record), high, medium or low name-and-geography matches, or null. | [optional] 
**Domains** | Pointer to **[]string** | The regulatory domains of the facility&#39;s signals (drug, biologic, device, veterinary, food, cosmetic, tobacco). | [optional] 
**PrimaryDomain** | Pointer to **NullableString** | The domain with the most signals. | [optional] 
**IsDrugFacility** | Pointer to **NullableBool** | True when the registers say the site makes medicines (an FDA drug registration, or an EU or Canadian medicines register entry). | [optional] 
**IsDeviceFacility** | Pointer to **NullableBool** | True when the registers say the site makes devices. | [optional] 
**RegisterDomain** | Pointer to **NullableString** | The display bucket over the two flags: drug, device, drug_and_device or non_medicine. | [optional] 
**DeviceClassificationStatus** | Pointer to **NullableString** | The registered-product rollup for a device facility: material, unresolved, low_signal or not_applicable. | [optional] 
**IsLowSignalDevice** | Pointer to **NullableBool** | True when the facility&#39;s device products are all low-signal commodity classes. | [optional] 
**Regulators** | Pointer to **[]string** | The regulators holding a registry record for the facility (fda, ema, health_canada), sorted. | [optional] 
**NRegulators** | Pointer to **NullableInt32** | How many regulators hold a registry record. | [optional] 
**IsMultiRegulator** | Pointer to **NullableBool** | True when more than one regulator holds a record: the corroboration flag. | [optional] 
**OrgRegulators** | Pointer to **[]string** | The regulators holding an organization-level record for the operating firm (this carries the UK MHRA leg). | [optional] 
**SeverityRank** | Pointer to **NullableInt32** | The in-force site-scope severity, 0 (no signal) to 4 (non-compliant). | [optional] 
**SeverityLevel** | Pointer to **NullableString** | The in-force site-scope severity: no_signal, clean, advisory, action or non_compliant. | [optional] 
**SeveritySource** | Pointer to **NullableString** | The regulator behind the severity: fda, ema, health_canada or mhra. | [optional] 
**SeveritySignalType** | Pointer to **NullableString** | The kind of signal behind the severity (for example an inspection classification). | [optional] 
**SeverityRawVerdict** | Pointer to **NullableString** | The regulator&#39;s own verdict behind the severity (for example OAI). | [optional] 
**SeveritySignalDate** | Pointer to **NullableString** | The date of the signal behind the severity. | [optional] 
**WorstSeverityRankEver** | Pointer to **NullableInt32** | The highest site-scope severity the facility has ever had, 0 to 4. | [optional] 
**WorstSeverityLevelEver** | Pointer to **NullableString** | The worst severity ever: no_signal, clean, advisory, action or non_compliant. | [optional] 
**HasQualitySignal** | Pointer to **NullableBool** | True when the facility has any quality signal. | [optional] 
**LatestSignalDate** | Pointer to **NullableString** | The date of the latest signal of any scope. | [optional] 
**NSignalsTotal** | Pointer to **NullableInt32** | Signals of any scope, all time. | [optional] 
**NSignals365d** | Pointer to **NullableInt32** | Signals in the last 365 days of the build. | [optional] 
**NInspections** | Pointer to **NullableInt32** | Inspections of the site by any regulator. | [optional] 
**LatestInspectionDate** | Pointer to **NullableString** | The date of the latest inspection by any regulator. | [optional] 
**NFdaInspections** | Pointer to **NullableInt32** | FDA inspections of the site. | [optional] 
**LatestFdaInspectionDate** | Pointer to **NullableString** | The date of the latest FDA inspection. | [optional] 
**LatestInspectionClassification** | Pointer to **NullableString** | The latest FDA inspection outcome (NAI, VAI or OAI). | [optional] 
**NInspectionsNai** | Pointer to **NullableInt32** | FDA inspections classified NAI. | [optional] 
**NInspectionsVai** | Pointer to **NullableInt32** | FDA inspections classified VAI. | [optional] 
**NInspectionsOai** | Pointer to **NullableInt32** | FDA inspections classified OAI. | [optional] 
**N483** | Pointer to **NullableInt32** | FDA inspections that ended with a Form 483. | [optional] 
**NHcInspections** | Pointer to **NullableInt32** | Health Canada inspections of the site. | [optional] 
**LatestHcInspectionDate** | Pointer to **NullableString** | The date of the latest Health Canada inspection. | [optional] 
**LatestHcInspectionRating** | Pointer to **NullableString** | The latest Health Canada rating (for example compliant or non-compliant). | [optional] 
**NHcInspectionsNonCompliant** | Pointer to **NullableInt32** | Health Canada inspections rated non-compliant. | [optional] 
**NImportRefusals** | Pointer to **NullableInt32** | FDA import refusals at the site. | [optional] 
**LatestImportRefusalDate** | Pointer to **NullableString** | The date of the latest import refusal. | [optional] 
**NGmpCertificates** | Pointer to **NullableInt32** | EU GMP certificates of the site. | [optional] 
**LatestGmpCertificateDate** | Pointer to **NullableString** | The date of the latest EU GMP certificate. | [optional] 
**NNoncomplianceStatements** | Pointer to **NullableInt32** | EU GMP non-compliance statements about the site. | [optional] 
**LatestNoncomplianceStatementDate** | Pointer to **NullableString** | The date of the latest EU non-compliance statement. | [optional] 
**NLicences** | Pointer to **NullableInt32** | Health Canada establishment licences of the site. | [optional] 
**LicenceStatus** | Pointer to **NullableString** | The licence status (the strictest licence of the site). | [optional] 
**OrgNLicenceActions** | Pointer to **NullableInt32** | UK MHRA licence suspensions or revocations against the firm (repeated; does not add). | [optional] 
**OrgLatestLicenceActionDate** | Pointer to **NullableString** | The date of the firm&#39;s latest licence action. | [optional] 
**AsOf** | Pointer to **NullableString** | The build date the counts are as of. | [optional] 

## Methods

### NewGxpFacilitiesRow

`func NewGxpFacilitiesRow() *GxpFacilitiesRow`

NewGxpFacilitiesRow instantiates a new GxpFacilitiesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGxpFacilitiesRowWithDefaults

`func NewGxpFacilitiesRowWithDefaults() *GxpFacilitiesRow`

NewGxpFacilitiesRowWithDefaults instantiates a new GxpFacilitiesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacilityKey

`func (o *GxpFacilitiesRow) GetFacilityKey() string`

GetFacilityKey returns the FacilityKey field if non-nil, zero value otherwise.

### GetFacilityKeyOk

`func (o *GxpFacilitiesRow) GetFacilityKeyOk() (*string, bool)`

GetFacilityKeyOk returns a tuple with the FacilityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityKey

`func (o *GxpFacilitiesRow) SetFacilityKey(v string)`

SetFacilityKey sets FacilityKey field to given value.

### HasFacilityKey

`func (o *GxpFacilitiesRow) HasFacilityKey() bool`

HasFacilityKey returns a boolean if a field has been set.

### SetFacilityKeyNil

`func (o *GxpFacilitiesRow) SetFacilityKeyNil(b bool)`

 SetFacilityKeyNil sets the value for FacilityKey to be an explicit nil

### UnsetFacilityKey
`func (o *GxpFacilitiesRow) UnsetFacilityKey()`

UnsetFacilityKey ensures that no value is present for FacilityKey, not even an explicit nil
### GetRegistryFacilityKeys

`func (o *GxpFacilitiesRow) GetRegistryFacilityKeys() []string`

GetRegistryFacilityKeys returns the RegistryFacilityKeys field if non-nil, zero value otherwise.

### GetRegistryFacilityKeysOk

`func (o *GxpFacilitiesRow) GetRegistryFacilityKeysOk() (*[]string, bool)`

GetRegistryFacilityKeysOk returns a tuple with the RegistryFacilityKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryFacilityKeys

`func (o *GxpFacilitiesRow) SetRegistryFacilityKeys(v []string)`

SetRegistryFacilityKeys sets RegistryFacilityKeys field to given value.

### HasRegistryFacilityKeys

`func (o *GxpFacilitiesRow) HasRegistryFacilityKeys() bool`

HasRegistryFacilityKeys returns a boolean if a field has been set.

### SetRegistryFacilityKeysNil

`func (o *GxpFacilitiesRow) SetRegistryFacilityKeysNil(b bool)`

 SetRegistryFacilityKeysNil sets the value for RegistryFacilityKeys to be an explicit nil

### UnsetRegistryFacilityKeys
`func (o *GxpFacilitiesRow) UnsetRegistryFacilityKeys()`

UnsetRegistryFacilityKeys ensures that no value is present for RegistryFacilityKeys, not even an explicit nil
### GetFacilityName

`func (o *GxpFacilitiesRow) GetFacilityName() string`

GetFacilityName returns the FacilityName field if non-nil, zero value otherwise.

### GetFacilityNameOk

`func (o *GxpFacilitiesRow) GetFacilityNameOk() (*string, bool)`

GetFacilityNameOk returns a tuple with the FacilityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityName

`func (o *GxpFacilitiesRow) SetFacilityName(v string)`

SetFacilityName sets FacilityName field to given value.

### HasFacilityName

`func (o *GxpFacilitiesRow) HasFacilityName() bool`

HasFacilityName returns a boolean if a field has been set.

### SetFacilityNameNil

`func (o *GxpFacilitiesRow) SetFacilityNameNil(b bool)`

 SetFacilityNameNil sets the value for FacilityName to be an explicit nil

### UnsetFacilityName
`func (o *GxpFacilitiesRow) UnsetFacilityName()`

UnsetFacilityName ensures that no value is present for FacilityName, not even an explicit nil
### GetMemberCount

`func (o *GxpFacilitiesRow) GetMemberCount() int64`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *GxpFacilitiesRow) GetMemberCountOk() (*int64, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *GxpFacilitiesRow) SetMemberCount(v int64)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *GxpFacilitiesRow) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### SetMemberCountNil

`func (o *GxpFacilitiesRow) SetMemberCountNil(b bool)`

 SetMemberCountNil sets the value for MemberCount to be an explicit nil

### UnsetMemberCount
`func (o *GxpFacilitiesRow) UnsetMemberCount()`

UnsetMemberCount ensures that no value is present for MemberCount, not even an explicit nil
### GetAddressLine

`func (o *GxpFacilitiesRow) GetAddressLine() string`

GetAddressLine returns the AddressLine field if non-nil, zero value otherwise.

### GetAddressLineOk

`func (o *GxpFacilitiesRow) GetAddressLineOk() (*string, bool)`

GetAddressLineOk returns a tuple with the AddressLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine

`func (o *GxpFacilitiesRow) SetAddressLine(v string)`

SetAddressLine sets AddressLine field to given value.

### HasAddressLine

`func (o *GxpFacilitiesRow) HasAddressLine() bool`

HasAddressLine returns a boolean if a field has been set.

### SetAddressLineNil

`func (o *GxpFacilitiesRow) SetAddressLineNil(b bool)`

 SetAddressLineNil sets the value for AddressLine to be an explicit nil

### UnsetAddressLine
`func (o *GxpFacilitiesRow) UnsetAddressLine()`

UnsetAddressLine ensures that no value is present for AddressLine, not even an explicit nil
### GetCity

`func (o *GxpFacilitiesRow) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GxpFacilitiesRow) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GxpFacilitiesRow) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GxpFacilitiesRow) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *GxpFacilitiesRow) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *GxpFacilitiesRow) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetStateRegion

`func (o *GxpFacilitiesRow) GetStateRegion() string`

GetStateRegion returns the StateRegion field if non-nil, zero value otherwise.

### GetStateRegionOk

`func (o *GxpFacilitiesRow) GetStateRegionOk() (*string, bool)`

GetStateRegionOk returns a tuple with the StateRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRegion

`func (o *GxpFacilitiesRow) SetStateRegion(v string)`

SetStateRegion sets StateRegion field to given value.

### HasStateRegion

`func (o *GxpFacilitiesRow) HasStateRegion() bool`

HasStateRegion returns a boolean if a field has been set.

### SetStateRegionNil

`func (o *GxpFacilitiesRow) SetStateRegionNil(b bool)`

 SetStateRegionNil sets the value for StateRegion to be an explicit nil

### UnsetStateRegion
`func (o *GxpFacilitiesRow) UnsetStateRegion()`

UnsetStateRegion ensures that no value is present for StateRegion, not even an explicit nil
### GetPostal

`func (o *GxpFacilitiesRow) GetPostal() string`

GetPostal returns the Postal field if non-nil, zero value otherwise.

### GetPostalOk

`func (o *GxpFacilitiesRow) GetPostalOk() (*string, bool)`

GetPostalOk returns a tuple with the Postal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostal

`func (o *GxpFacilitiesRow) SetPostal(v string)`

SetPostal sets Postal field to given value.

### HasPostal

`func (o *GxpFacilitiesRow) HasPostal() bool`

HasPostal returns a boolean if a field has been set.

### SetPostalNil

`func (o *GxpFacilitiesRow) SetPostalNil(b bool)`

 SetPostalNil sets the value for Postal to be an explicit nil

### UnsetPostal
`func (o *GxpFacilitiesRow) UnsetPostal()`

UnsetPostal ensures that no value is present for Postal, not even an explicit nil
### GetCountry

`func (o *GxpFacilitiesRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GxpFacilitiesRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GxpFacilitiesRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GxpFacilitiesRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *GxpFacilitiesRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *GxpFacilitiesRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLatitude

`func (o *GxpFacilitiesRow) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *GxpFacilitiesRow) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *GxpFacilitiesRow) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *GxpFacilitiesRow) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *GxpFacilitiesRow) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *GxpFacilitiesRow) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *GxpFacilitiesRow) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *GxpFacilitiesRow) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *GxpFacilitiesRow) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *GxpFacilitiesRow) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *GxpFacilitiesRow) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *GxpFacilitiesRow) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetGeocodePrecision

`func (o *GxpFacilitiesRow) GetGeocodePrecision() string`

GetGeocodePrecision returns the GeocodePrecision field if non-nil, zero value otherwise.

### GetGeocodePrecisionOk

`func (o *GxpFacilitiesRow) GetGeocodePrecisionOk() (*string, bool)`

GetGeocodePrecisionOk returns a tuple with the GeocodePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocodePrecision

`func (o *GxpFacilitiesRow) SetGeocodePrecision(v string)`

SetGeocodePrecision sets GeocodePrecision field to given value.

### HasGeocodePrecision

`func (o *GxpFacilitiesRow) HasGeocodePrecision() bool`

HasGeocodePrecision returns a boolean if a field has been set.

### SetGeocodePrecisionNil

`func (o *GxpFacilitiesRow) SetGeocodePrecisionNil(b bool)`

 SetGeocodePrecisionNil sets the value for GeocodePrecision to be an explicit nil

### UnsetGeocodePrecision
`func (o *GxpFacilitiesRow) UnsetGeocodePrecision()`

UnsetGeocodePrecision ensures that no value is present for GeocodePrecision, not even an explicit nil
### GetCoordSource

`func (o *GxpFacilitiesRow) GetCoordSource() string`

GetCoordSource returns the CoordSource field if non-nil, zero value otherwise.

### GetCoordSourceOk

`func (o *GxpFacilitiesRow) GetCoordSourceOk() (*string, bool)`

GetCoordSourceOk returns a tuple with the CoordSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordSource

`func (o *GxpFacilitiesRow) SetCoordSource(v string)`

SetCoordSource sets CoordSource field to given value.

### HasCoordSource

`func (o *GxpFacilitiesRow) HasCoordSource() bool`

HasCoordSource returns a boolean if a field has been set.

### SetCoordSourceNil

`func (o *GxpFacilitiesRow) SetCoordSourceNil(b bool)`

 SetCoordSourceNil sets the value for CoordSource to be an explicit nil

### UnsetCoordSource
`func (o *GxpFacilitiesRow) UnsetCoordSource()`

UnsetCoordSource ensures that no value is present for CoordSource, not even an explicit nil
### GetFirmName

`func (o *GxpFacilitiesRow) GetFirmName() string`

GetFirmName returns the FirmName field if non-nil, zero value otherwise.

### GetFirmNameOk

`func (o *GxpFacilitiesRow) GetFirmNameOk() (*string, bool)`

GetFirmNameOk returns a tuple with the FirmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmName

`func (o *GxpFacilitiesRow) SetFirmName(v string)`

SetFirmName sets FirmName field to given value.

### HasFirmName

`func (o *GxpFacilitiesRow) HasFirmName() bool`

HasFirmName returns a boolean if a field has been set.

### SetFirmNameNil

`func (o *GxpFacilitiesRow) SetFirmNameNil(b bool)`

 SetFirmNameNil sets the value for FirmName to be an explicit nil

### UnsetFirmName
`func (o *GxpFacilitiesRow) UnsetFirmName()`

UnsetFirmName ensures that no value is present for FirmName, not even an explicit nil
### GetFirmOrgId

`func (o *GxpFacilitiesRow) GetFirmOrgId() string`

GetFirmOrgId returns the FirmOrgId field if non-nil, zero value otherwise.

### GetFirmOrgIdOk

`func (o *GxpFacilitiesRow) GetFirmOrgIdOk() (*string, bool)`

GetFirmOrgIdOk returns a tuple with the FirmOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmOrgId

`func (o *GxpFacilitiesRow) SetFirmOrgId(v string)`

SetFirmOrgId sets FirmOrgId field to given value.

### HasFirmOrgId

`func (o *GxpFacilitiesRow) HasFirmOrgId() bool`

HasFirmOrgId returns a boolean if a field has been set.

### SetFirmOrgIdNil

`func (o *GxpFacilitiesRow) SetFirmOrgIdNil(b bool)`

 SetFirmOrgIdNil sets the value for FirmOrgId to be an explicit nil

### UnsetFirmOrgId
`func (o *GxpFacilitiesRow) UnsetFirmOrgId()`

UnsetFirmOrgId ensures that no value is present for FirmOrgId, not even an explicit nil
### GetFirmCompanyId

`func (o *GxpFacilitiesRow) GetFirmCompanyId() string`

GetFirmCompanyId returns the FirmCompanyId field if non-nil, zero value otherwise.

### GetFirmCompanyIdOk

`func (o *GxpFacilitiesRow) GetFirmCompanyIdOk() (*string, bool)`

GetFirmCompanyIdOk returns a tuple with the FirmCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyId

`func (o *GxpFacilitiesRow) SetFirmCompanyId(v string)`

SetFirmCompanyId sets FirmCompanyId field to given value.

### HasFirmCompanyId

`func (o *GxpFacilitiesRow) HasFirmCompanyId() bool`

HasFirmCompanyId returns a boolean if a field has been set.

### SetFirmCompanyIdNil

`func (o *GxpFacilitiesRow) SetFirmCompanyIdNil(b bool)`

 SetFirmCompanyIdNil sets the value for FirmCompanyId to be an explicit nil

### UnsetFirmCompanyId
`func (o *GxpFacilitiesRow) UnsetFirmCompanyId()`

UnsetFirmCompanyId ensures that no value is present for FirmCompanyId, not even an explicit nil
### GetFirmCompanyName

`func (o *GxpFacilitiesRow) GetFirmCompanyName() string`

GetFirmCompanyName returns the FirmCompanyName field if non-nil, zero value otherwise.

### GetFirmCompanyNameOk

`func (o *GxpFacilitiesRow) GetFirmCompanyNameOk() (*string, bool)`

GetFirmCompanyNameOk returns a tuple with the FirmCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyName

`func (o *GxpFacilitiesRow) SetFirmCompanyName(v string)`

SetFirmCompanyName sets FirmCompanyName field to given value.

### HasFirmCompanyName

`func (o *GxpFacilitiesRow) HasFirmCompanyName() bool`

HasFirmCompanyName returns a boolean if a field has been set.

### SetFirmCompanyNameNil

`func (o *GxpFacilitiesRow) SetFirmCompanyNameNil(b bool)`

 SetFirmCompanyNameNil sets the value for FirmCompanyName to be an explicit nil

### UnsetFirmCompanyName
`func (o *GxpFacilitiesRow) UnsetFirmCompanyName()`

UnsetFirmCompanyName ensures that no value is present for FirmCompanyName, not even an explicit nil
### GetFirmGroupId

`func (o *GxpFacilitiesRow) GetFirmGroupId() string`

GetFirmGroupId returns the FirmGroupId field if non-nil, zero value otherwise.

### GetFirmGroupIdOk

`func (o *GxpFacilitiesRow) GetFirmGroupIdOk() (*string, bool)`

GetFirmGroupIdOk returns a tuple with the FirmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupId

`func (o *GxpFacilitiesRow) SetFirmGroupId(v string)`

SetFirmGroupId sets FirmGroupId field to given value.

### HasFirmGroupId

`func (o *GxpFacilitiesRow) HasFirmGroupId() bool`

HasFirmGroupId returns a boolean if a field has been set.

### SetFirmGroupIdNil

`func (o *GxpFacilitiesRow) SetFirmGroupIdNil(b bool)`

 SetFirmGroupIdNil sets the value for FirmGroupId to be an explicit nil

### UnsetFirmGroupId
`func (o *GxpFacilitiesRow) UnsetFirmGroupId()`

UnsetFirmGroupId ensures that no value is present for FirmGroupId, not even an explicit nil
### GetFirmGroupName

`func (o *GxpFacilitiesRow) GetFirmGroupName() string`

GetFirmGroupName returns the FirmGroupName field if non-nil, zero value otherwise.

### GetFirmGroupNameOk

`func (o *GxpFacilitiesRow) GetFirmGroupNameOk() (*string, bool)`

GetFirmGroupNameOk returns a tuple with the FirmGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupName

`func (o *GxpFacilitiesRow) SetFirmGroupName(v string)`

SetFirmGroupName sets FirmGroupName field to given value.

### HasFirmGroupName

`func (o *GxpFacilitiesRow) HasFirmGroupName() bool`

HasFirmGroupName returns a boolean if a field has been set.

### SetFirmGroupNameNil

`func (o *GxpFacilitiesRow) SetFirmGroupNameNil(b bool)`

 SetFirmGroupNameNil sets the value for FirmGroupName to be an explicit nil

### UnsetFirmGroupName
`func (o *GxpFacilitiesRow) UnsetFirmGroupName()`

UnsetFirmGroupName ensures that no value is present for FirmGroupName, not even an explicit nil
### GetIsOrgResolved

`func (o *GxpFacilitiesRow) GetIsOrgResolved() bool`

GetIsOrgResolved returns the IsOrgResolved field if non-nil, zero value otherwise.

### GetIsOrgResolvedOk

`func (o *GxpFacilitiesRow) GetIsOrgResolvedOk() (*bool, bool)`

GetIsOrgResolvedOk returns a tuple with the IsOrgResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgResolved

`func (o *GxpFacilitiesRow) SetIsOrgResolved(v bool)`

SetIsOrgResolved sets IsOrgResolved field to given value.

### HasIsOrgResolved

`func (o *GxpFacilitiesRow) HasIsOrgResolved() bool`

HasIsOrgResolved returns a boolean if a field has been set.

### SetIsOrgResolvedNil

`func (o *GxpFacilitiesRow) SetIsOrgResolvedNil(b bool)`

 SetIsOrgResolvedNil sets the value for IsOrgResolved to be an explicit nil

### UnsetIsOrgResolved
`func (o *GxpFacilitiesRow) UnsetIsOrgResolved()`

UnsetIsOrgResolved ensures that no value is present for IsOrgResolved, not even an explicit nil
### GetMatchTier

`func (o *GxpFacilitiesRow) GetMatchTier() string`

GetMatchTier returns the MatchTier field if non-nil, zero value otherwise.

### GetMatchTierOk

`func (o *GxpFacilitiesRow) GetMatchTierOk() (*string, bool)`

GetMatchTierOk returns a tuple with the MatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTier

`func (o *GxpFacilitiesRow) SetMatchTier(v string)`

SetMatchTier sets MatchTier field to given value.

### HasMatchTier

`func (o *GxpFacilitiesRow) HasMatchTier() bool`

HasMatchTier returns a boolean if a field has been set.

### SetMatchTierNil

`func (o *GxpFacilitiesRow) SetMatchTierNil(b bool)`

 SetMatchTierNil sets the value for MatchTier to be an explicit nil

### UnsetMatchTier
`func (o *GxpFacilitiesRow) UnsetMatchTier()`

UnsetMatchTier ensures that no value is present for MatchTier, not even an explicit nil
### GetDomains

`func (o *GxpFacilitiesRow) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *GxpFacilitiesRow) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *GxpFacilitiesRow) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *GxpFacilitiesRow) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### SetDomainsNil

`func (o *GxpFacilitiesRow) SetDomainsNil(b bool)`

 SetDomainsNil sets the value for Domains to be an explicit nil

### UnsetDomains
`func (o *GxpFacilitiesRow) UnsetDomains()`

UnsetDomains ensures that no value is present for Domains, not even an explicit nil
### GetPrimaryDomain

`func (o *GxpFacilitiesRow) GetPrimaryDomain() string`

GetPrimaryDomain returns the PrimaryDomain field if non-nil, zero value otherwise.

### GetPrimaryDomainOk

`func (o *GxpFacilitiesRow) GetPrimaryDomainOk() (*string, bool)`

GetPrimaryDomainOk returns a tuple with the PrimaryDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDomain

`func (o *GxpFacilitiesRow) SetPrimaryDomain(v string)`

SetPrimaryDomain sets PrimaryDomain field to given value.

### HasPrimaryDomain

`func (o *GxpFacilitiesRow) HasPrimaryDomain() bool`

HasPrimaryDomain returns a boolean if a field has been set.

### SetPrimaryDomainNil

`func (o *GxpFacilitiesRow) SetPrimaryDomainNil(b bool)`

 SetPrimaryDomainNil sets the value for PrimaryDomain to be an explicit nil

### UnsetPrimaryDomain
`func (o *GxpFacilitiesRow) UnsetPrimaryDomain()`

UnsetPrimaryDomain ensures that no value is present for PrimaryDomain, not even an explicit nil
### GetIsDrugFacility

`func (o *GxpFacilitiesRow) GetIsDrugFacility() bool`

GetIsDrugFacility returns the IsDrugFacility field if non-nil, zero value otherwise.

### GetIsDrugFacilityOk

`func (o *GxpFacilitiesRow) GetIsDrugFacilityOk() (*bool, bool)`

GetIsDrugFacilityOk returns a tuple with the IsDrugFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDrugFacility

`func (o *GxpFacilitiesRow) SetIsDrugFacility(v bool)`

SetIsDrugFacility sets IsDrugFacility field to given value.

### HasIsDrugFacility

`func (o *GxpFacilitiesRow) HasIsDrugFacility() bool`

HasIsDrugFacility returns a boolean if a field has been set.

### SetIsDrugFacilityNil

`func (o *GxpFacilitiesRow) SetIsDrugFacilityNil(b bool)`

 SetIsDrugFacilityNil sets the value for IsDrugFacility to be an explicit nil

### UnsetIsDrugFacility
`func (o *GxpFacilitiesRow) UnsetIsDrugFacility()`

UnsetIsDrugFacility ensures that no value is present for IsDrugFacility, not even an explicit nil
### GetIsDeviceFacility

`func (o *GxpFacilitiesRow) GetIsDeviceFacility() bool`

GetIsDeviceFacility returns the IsDeviceFacility field if non-nil, zero value otherwise.

### GetIsDeviceFacilityOk

`func (o *GxpFacilitiesRow) GetIsDeviceFacilityOk() (*bool, bool)`

GetIsDeviceFacilityOk returns a tuple with the IsDeviceFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeviceFacility

`func (o *GxpFacilitiesRow) SetIsDeviceFacility(v bool)`

SetIsDeviceFacility sets IsDeviceFacility field to given value.

### HasIsDeviceFacility

`func (o *GxpFacilitiesRow) HasIsDeviceFacility() bool`

HasIsDeviceFacility returns a boolean if a field has been set.

### SetIsDeviceFacilityNil

`func (o *GxpFacilitiesRow) SetIsDeviceFacilityNil(b bool)`

 SetIsDeviceFacilityNil sets the value for IsDeviceFacility to be an explicit nil

### UnsetIsDeviceFacility
`func (o *GxpFacilitiesRow) UnsetIsDeviceFacility()`

UnsetIsDeviceFacility ensures that no value is present for IsDeviceFacility, not even an explicit nil
### GetRegisterDomain

`func (o *GxpFacilitiesRow) GetRegisterDomain() string`

GetRegisterDomain returns the RegisterDomain field if non-nil, zero value otherwise.

### GetRegisterDomainOk

`func (o *GxpFacilitiesRow) GetRegisterDomainOk() (*string, bool)`

GetRegisterDomainOk returns a tuple with the RegisterDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterDomain

`func (o *GxpFacilitiesRow) SetRegisterDomain(v string)`

SetRegisterDomain sets RegisterDomain field to given value.

### HasRegisterDomain

`func (o *GxpFacilitiesRow) HasRegisterDomain() bool`

HasRegisterDomain returns a boolean if a field has been set.

### SetRegisterDomainNil

`func (o *GxpFacilitiesRow) SetRegisterDomainNil(b bool)`

 SetRegisterDomainNil sets the value for RegisterDomain to be an explicit nil

### UnsetRegisterDomain
`func (o *GxpFacilitiesRow) UnsetRegisterDomain()`

UnsetRegisterDomain ensures that no value is present for RegisterDomain, not even an explicit nil
### GetDeviceClassificationStatus

`func (o *GxpFacilitiesRow) GetDeviceClassificationStatus() string`

GetDeviceClassificationStatus returns the DeviceClassificationStatus field if non-nil, zero value otherwise.

### GetDeviceClassificationStatusOk

`func (o *GxpFacilitiesRow) GetDeviceClassificationStatusOk() (*string, bool)`

GetDeviceClassificationStatusOk returns a tuple with the DeviceClassificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClassificationStatus

`func (o *GxpFacilitiesRow) SetDeviceClassificationStatus(v string)`

SetDeviceClassificationStatus sets DeviceClassificationStatus field to given value.

### HasDeviceClassificationStatus

`func (o *GxpFacilitiesRow) HasDeviceClassificationStatus() bool`

HasDeviceClassificationStatus returns a boolean if a field has been set.

### SetDeviceClassificationStatusNil

`func (o *GxpFacilitiesRow) SetDeviceClassificationStatusNil(b bool)`

 SetDeviceClassificationStatusNil sets the value for DeviceClassificationStatus to be an explicit nil

### UnsetDeviceClassificationStatus
`func (o *GxpFacilitiesRow) UnsetDeviceClassificationStatus()`

UnsetDeviceClassificationStatus ensures that no value is present for DeviceClassificationStatus, not even an explicit nil
### GetIsLowSignalDevice

`func (o *GxpFacilitiesRow) GetIsLowSignalDevice() bool`

GetIsLowSignalDevice returns the IsLowSignalDevice field if non-nil, zero value otherwise.

### GetIsLowSignalDeviceOk

`func (o *GxpFacilitiesRow) GetIsLowSignalDeviceOk() (*bool, bool)`

GetIsLowSignalDeviceOk returns a tuple with the IsLowSignalDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLowSignalDevice

`func (o *GxpFacilitiesRow) SetIsLowSignalDevice(v bool)`

SetIsLowSignalDevice sets IsLowSignalDevice field to given value.

### HasIsLowSignalDevice

`func (o *GxpFacilitiesRow) HasIsLowSignalDevice() bool`

HasIsLowSignalDevice returns a boolean if a field has been set.

### SetIsLowSignalDeviceNil

`func (o *GxpFacilitiesRow) SetIsLowSignalDeviceNil(b bool)`

 SetIsLowSignalDeviceNil sets the value for IsLowSignalDevice to be an explicit nil

### UnsetIsLowSignalDevice
`func (o *GxpFacilitiesRow) UnsetIsLowSignalDevice()`

UnsetIsLowSignalDevice ensures that no value is present for IsLowSignalDevice, not even an explicit nil
### GetRegulators

`func (o *GxpFacilitiesRow) GetRegulators() []string`

GetRegulators returns the Regulators field if non-nil, zero value otherwise.

### GetRegulatorsOk

`func (o *GxpFacilitiesRow) GetRegulatorsOk() (*[]string, bool)`

GetRegulatorsOk returns a tuple with the Regulators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulators

`func (o *GxpFacilitiesRow) SetRegulators(v []string)`

SetRegulators sets Regulators field to given value.

### HasRegulators

`func (o *GxpFacilitiesRow) HasRegulators() bool`

HasRegulators returns a boolean if a field has been set.

### SetRegulatorsNil

`func (o *GxpFacilitiesRow) SetRegulatorsNil(b bool)`

 SetRegulatorsNil sets the value for Regulators to be an explicit nil

### UnsetRegulators
`func (o *GxpFacilitiesRow) UnsetRegulators()`

UnsetRegulators ensures that no value is present for Regulators, not even an explicit nil
### GetNRegulators

`func (o *GxpFacilitiesRow) GetNRegulators() int32`

GetNRegulators returns the NRegulators field if non-nil, zero value otherwise.

### GetNRegulatorsOk

`func (o *GxpFacilitiesRow) GetNRegulatorsOk() (*int32, bool)`

GetNRegulatorsOk returns a tuple with the NRegulators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRegulators

`func (o *GxpFacilitiesRow) SetNRegulators(v int32)`

SetNRegulators sets NRegulators field to given value.

### HasNRegulators

`func (o *GxpFacilitiesRow) HasNRegulators() bool`

HasNRegulators returns a boolean if a field has been set.

### SetNRegulatorsNil

`func (o *GxpFacilitiesRow) SetNRegulatorsNil(b bool)`

 SetNRegulatorsNil sets the value for NRegulators to be an explicit nil

### UnsetNRegulators
`func (o *GxpFacilitiesRow) UnsetNRegulators()`

UnsetNRegulators ensures that no value is present for NRegulators, not even an explicit nil
### GetIsMultiRegulator

`func (o *GxpFacilitiesRow) GetIsMultiRegulator() bool`

GetIsMultiRegulator returns the IsMultiRegulator field if non-nil, zero value otherwise.

### GetIsMultiRegulatorOk

`func (o *GxpFacilitiesRow) GetIsMultiRegulatorOk() (*bool, bool)`

GetIsMultiRegulatorOk returns a tuple with the IsMultiRegulator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiRegulator

`func (o *GxpFacilitiesRow) SetIsMultiRegulator(v bool)`

SetIsMultiRegulator sets IsMultiRegulator field to given value.

### HasIsMultiRegulator

`func (o *GxpFacilitiesRow) HasIsMultiRegulator() bool`

HasIsMultiRegulator returns a boolean if a field has been set.

### SetIsMultiRegulatorNil

`func (o *GxpFacilitiesRow) SetIsMultiRegulatorNil(b bool)`

 SetIsMultiRegulatorNil sets the value for IsMultiRegulator to be an explicit nil

### UnsetIsMultiRegulator
`func (o *GxpFacilitiesRow) UnsetIsMultiRegulator()`

UnsetIsMultiRegulator ensures that no value is present for IsMultiRegulator, not even an explicit nil
### GetOrgRegulators

`func (o *GxpFacilitiesRow) GetOrgRegulators() []string`

GetOrgRegulators returns the OrgRegulators field if non-nil, zero value otherwise.

### GetOrgRegulatorsOk

`func (o *GxpFacilitiesRow) GetOrgRegulatorsOk() (*[]string, bool)`

GetOrgRegulatorsOk returns a tuple with the OrgRegulators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRegulators

`func (o *GxpFacilitiesRow) SetOrgRegulators(v []string)`

SetOrgRegulators sets OrgRegulators field to given value.

### HasOrgRegulators

`func (o *GxpFacilitiesRow) HasOrgRegulators() bool`

HasOrgRegulators returns a boolean if a field has been set.

### SetOrgRegulatorsNil

`func (o *GxpFacilitiesRow) SetOrgRegulatorsNil(b bool)`

 SetOrgRegulatorsNil sets the value for OrgRegulators to be an explicit nil

### UnsetOrgRegulators
`func (o *GxpFacilitiesRow) UnsetOrgRegulators()`

UnsetOrgRegulators ensures that no value is present for OrgRegulators, not even an explicit nil
### GetSeverityRank

`func (o *GxpFacilitiesRow) GetSeverityRank() int32`

GetSeverityRank returns the SeverityRank field if non-nil, zero value otherwise.

### GetSeverityRankOk

`func (o *GxpFacilitiesRow) GetSeverityRankOk() (*int32, bool)`

GetSeverityRankOk returns a tuple with the SeverityRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityRank

`func (o *GxpFacilitiesRow) SetSeverityRank(v int32)`

SetSeverityRank sets SeverityRank field to given value.

### HasSeverityRank

`func (o *GxpFacilitiesRow) HasSeverityRank() bool`

HasSeverityRank returns a boolean if a field has been set.

### SetSeverityRankNil

`func (o *GxpFacilitiesRow) SetSeverityRankNil(b bool)`

 SetSeverityRankNil sets the value for SeverityRank to be an explicit nil

### UnsetSeverityRank
`func (o *GxpFacilitiesRow) UnsetSeverityRank()`

UnsetSeverityRank ensures that no value is present for SeverityRank, not even an explicit nil
### GetSeverityLevel

`func (o *GxpFacilitiesRow) GetSeverityLevel() string`

GetSeverityLevel returns the SeverityLevel field if non-nil, zero value otherwise.

### GetSeverityLevelOk

`func (o *GxpFacilitiesRow) GetSeverityLevelOk() (*string, bool)`

GetSeverityLevelOk returns a tuple with the SeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevel

`func (o *GxpFacilitiesRow) SetSeverityLevel(v string)`

SetSeverityLevel sets SeverityLevel field to given value.

### HasSeverityLevel

`func (o *GxpFacilitiesRow) HasSeverityLevel() bool`

HasSeverityLevel returns a boolean if a field has been set.

### SetSeverityLevelNil

`func (o *GxpFacilitiesRow) SetSeverityLevelNil(b bool)`

 SetSeverityLevelNil sets the value for SeverityLevel to be an explicit nil

### UnsetSeverityLevel
`func (o *GxpFacilitiesRow) UnsetSeverityLevel()`

UnsetSeverityLevel ensures that no value is present for SeverityLevel, not even an explicit nil
### GetSeveritySource

`func (o *GxpFacilitiesRow) GetSeveritySource() string`

GetSeveritySource returns the SeveritySource field if non-nil, zero value otherwise.

### GetSeveritySourceOk

`func (o *GxpFacilitiesRow) GetSeveritySourceOk() (*string, bool)`

GetSeveritySourceOk returns a tuple with the SeveritySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeveritySource

`func (o *GxpFacilitiesRow) SetSeveritySource(v string)`

SetSeveritySource sets SeveritySource field to given value.

### HasSeveritySource

`func (o *GxpFacilitiesRow) HasSeveritySource() bool`

HasSeveritySource returns a boolean if a field has been set.

### SetSeveritySourceNil

`func (o *GxpFacilitiesRow) SetSeveritySourceNil(b bool)`

 SetSeveritySourceNil sets the value for SeveritySource to be an explicit nil

### UnsetSeveritySource
`func (o *GxpFacilitiesRow) UnsetSeveritySource()`

UnsetSeveritySource ensures that no value is present for SeveritySource, not even an explicit nil
### GetSeveritySignalType

`func (o *GxpFacilitiesRow) GetSeveritySignalType() string`

GetSeveritySignalType returns the SeveritySignalType field if non-nil, zero value otherwise.

### GetSeveritySignalTypeOk

`func (o *GxpFacilitiesRow) GetSeveritySignalTypeOk() (*string, bool)`

GetSeveritySignalTypeOk returns a tuple with the SeveritySignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeveritySignalType

`func (o *GxpFacilitiesRow) SetSeveritySignalType(v string)`

SetSeveritySignalType sets SeveritySignalType field to given value.

### HasSeveritySignalType

`func (o *GxpFacilitiesRow) HasSeveritySignalType() bool`

HasSeveritySignalType returns a boolean if a field has been set.

### SetSeveritySignalTypeNil

`func (o *GxpFacilitiesRow) SetSeveritySignalTypeNil(b bool)`

 SetSeveritySignalTypeNil sets the value for SeveritySignalType to be an explicit nil

### UnsetSeveritySignalType
`func (o *GxpFacilitiesRow) UnsetSeveritySignalType()`

UnsetSeveritySignalType ensures that no value is present for SeveritySignalType, not even an explicit nil
### GetSeverityRawVerdict

`func (o *GxpFacilitiesRow) GetSeverityRawVerdict() string`

GetSeverityRawVerdict returns the SeverityRawVerdict field if non-nil, zero value otherwise.

### GetSeverityRawVerdictOk

`func (o *GxpFacilitiesRow) GetSeverityRawVerdictOk() (*string, bool)`

GetSeverityRawVerdictOk returns a tuple with the SeverityRawVerdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityRawVerdict

`func (o *GxpFacilitiesRow) SetSeverityRawVerdict(v string)`

SetSeverityRawVerdict sets SeverityRawVerdict field to given value.

### HasSeverityRawVerdict

`func (o *GxpFacilitiesRow) HasSeverityRawVerdict() bool`

HasSeverityRawVerdict returns a boolean if a field has been set.

### SetSeverityRawVerdictNil

`func (o *GxpFacilitiesRow) SetSeverityRawVerdictNil(b bool)`

 SetSeverityRawVerdictNil sets the value for SeverityRawVerdict to be an explicit nil

### UnsetSeverityRawVerdict
`func (o *GxpFacilitiesRow) UnsetSeverityRawVerdict()`

UnsetSeverityRawVerdict ensures that no value is present for SeverityRawVerdict, not even an explicit nil
### GetSeveritySignalDate

`func (o *GxpFacilitiesRow) GetSeveritySignalDate() string`

GetSeveritySignalDate returns the SeveritySignalDate field if non-nil, zero value otherwise.

### GetSeveritySignalDateOk

`func (o *GxpFacilitiesRow) GetSeveritySignalDateOk() (*string, bool)`

GetSeveritySignalDateOk returns a tuple with the SeveritySignalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeveritySignalDate

`func (o *GxpFacilitiesRow) SetSeveritySignalDate(v string)`

SetSeveritySignalDate sets SeveritySignalDate field to given value.

### HasSeveritySignalDate

`func (o *GxpFacilitiesRow) HasSeveritySignalDate() bool`

HasSeveritySignalDate returns a boolean if a field has been set.

### SetSeveritySignalDateNil

`func (o *GxpFacilitiesRow) SetSeveritySignalDateNil(b bool)`

 SetSeveritySignalDateNil sets the value for SeveritySignalDate to be an explicit nil

### UnsetSeveritySignalDate
`func (o *GxpFacilitiesRow) UnsetSeveritySignalDate()`

UnsetSeveritySignalDate ensures that no value is present for SeveritySignalDate, not even an explicit nil
### GetWorstSeverityRankEver

`func (o *GxpFacilitiesRow) GetWorstSeverityRankEver() int32`

GetWorstSeverityRankEver returns the WorstSeverityRankEver field if non-nil, zero value otherwise.

### GetWorstSeverityRankEverOk

`func (o *GxpFacilitiesRow) GetWorstSeverityRankEverOk() (*int32, bool)`

GetWorstSeverityRankEverOk returns a tuple with the WorstSeverityRankEver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorstSeverityRankEver

`func (o *GxpFacilitiesRow) SetWorstSeverityRankEver(v int32)`

SetWorstSeverityRankEver sets WorstSeverityRankEver field to given value.

### HasWorstSeverityRankEver

`func (o *GxpFacilitiesRow) HasWorstSeverityRankEver() bool`

HasWorstSeverityRankEver returns a boolean if a field has been set.

### SetWorstSeverityRankEverNil

`func (o *GxpFacilitiesRow) SetWorstSeverityRankEverNil(b bool)`

 SetWorstSeverityRankEverNil sets the value for WorstSeverityRankEver to be an explicit nil

### UnsetWorstSeverityRankEver
`func (o *GxpFacilitiesRow) UnsetWorstSeverityRankEver()`

UnsetWorstSeverityRankEver ensures that no value is present for WorstSeverityRankEver, not even an explicit nil
### GetWorstSeverityLevelEver

`func (o *GxpFacilitiesRow) GetWorstSeverityLevelEver() string`

GetWorstSeverityLevelEver returns the WorstSeverityLevelEver field if non-nil, zero value otherwise.

### GetWorstSeverityLevelEverOk

`func (o *GxpFacilitiesRow) GetWorstSeverityLevelEverOk() (*string, bool)`

GetWorstSeverityLevelEverOk returns a tuple with the WorstSeverityLevelEver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorstSeverityLevelEver

`func (o *GxpFacilitiesRow) SetWorstSeverityLevelEver(v string)`

SetWorstSeverityLevelEver sets WorstSeverityLevelEver field to given value.

### HasWorstSeverityLevelEver

`func (o *GxpFacilitiesRow) HasWorstSeverityLevelEver() bool`

HasWorstSeverityLevelEver returns a boolean if a field has been set.

### SetWorstSeverityLevelEverNil

`func (o *GxpFacilitiesRow) SetWorstSeverityLevelEverNil(b bool)`

 SetWorstSeverityLevelEverNil sets the value for WorstSeverityLevelEver to be an explicit nil

### UnsetWorstSeverityLevelEver
`func (o *GxpFacilitiesRow) UnsetWorstSeverityLevelEver()`

UnsetWorstSeverityLevelEver ensures that no value is present for WorstSeverityLevelEver, not even an explicit nil
### GetHasQualitySignal

`func (o *GxpFacilitiesRow) GetHasQualitySignal() bool`

GetHasQualitySignal returns the HasQualitySignal field if non-nil, zero value otherwise.

### GetHasQualitySignalOk

`func (o *GxpFacilitiesRow) GetHasQualitySignalOk() (*bool, bool)`

GetHasQualitySignalOk returns a tuple with the HasQualitySignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasQualitySignal

`func (o *GxpFacilitiesRow) SetHasQualitySignal(v bool)`

SetHasQualitySignal sets HasQualitySignal field to given value.

### HasHasQualitySignal

`func (o *GxpFacilitiesRow) HasHasQualitySignal() bool`

HasHasQualitySignal returns a boolean if a field has been set.

### SetHasQualitySignalNil

`func (o *GxpFacilitiesRow) SetHasQualitySignalNil(b bool)`

 SetHasQualitySignalNil sets the value for HasQualitySignal to be an explicit nil

### UnsetHasQualitySignal
`func (o *GxpFacilitiesRow) UnsetHasQualitySignal()`

UnsetHasQualitySignal ensures that no value is present for HasQualitySignal, not even an explicit nil
### GetLatestSignalDate

`func (o *GxpFacilitiesRow) GetLatestSignalDate() string`

GetLatestSignalDate returns the LatestSignalDate field if non-nil, zero value otherwise.

### GetLatestSignalDateOk

`func (o *GxpFacilitiesRow) GetLatestSignalDateOk() (*string, bool)`

GetLatestSignalDateOk returns a tuple with the LatestSignalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSignalDate

`func (o *GxpFacilitiesRow) SetLatestSignalDate(v string)`

SetLatestSignalDate sets LatestSignalDate field to given value.

### HasLatestSignalDate

`func (o *GxpFacilitiesRow) HasLatestSignalDate() bool`

HasLatestSignalDate returns a boolean if a field has been set.

### SetLatestSignalDateNil

`func (o *GxpFacilitiesRow) SetLatestSignalDateNil(b bool)`

 SetLatestSignalDateNil sets the value for LatestSignalDate to be an explicit nil

### UnsetLatestSignalDate
`func (o *GxpFacilitiesRow) UnsetLatestSignalDate()`

UnsetLatestSignalDate ensures that no value is present for LatestSignalDate, not even an explicit nil
### GetNSignalsTotal

`func (o *GxpFacilitiesRow) GetNSignalsTotal() int32`

GetNSignalsTotal returns the NSignalsTotal field if non-nil, zero value otherwise.

### GetNSignalsTotalOk

`func (o *GxpFacilitiesRow) GetNSignalsTotalOk() (*int32, bool)`

GetNSignalsTotalOk returns a tuple with the NSignalsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSignalsTotal

`func (o *GxpFacilitiesRow) SetNSignalsTotal(v int32)`

SetNSignalsTotal sets NSignalsTotal field to given value.

### HasNSignalsTotal

`func (o *GxpFacilitiesRow) HasNSignalsTotal() bool`

HasNSignalsTotal returns a boolean if a field has been set.

### SetNSignalsTotalNil

`func (o *GxpFacilitiesRow) SetNSignalsTotalNil(b bool)`

 SetNSignalsTotalNil sets the value for NSignalsTotal to be an explicit nil

### UnsetNSignalsTotal
`func (o *GxpFacilitiesRow) UnsetNSignalsTotal()`

UnsetNSignalsTotal ensures that no value is present for NSignalsTotal, not even an explicit nil
### GetNSignals365d

`func (o *GxpFacilitiesRow) GetNSignals365d() int32`

GetNSignals365d returns the NSignals365d field if non-nil, zero value otherwise.

### GetNSignals365dOk

`func (o *GxpFacilitiesRow) GetNSignals365dOk() (*int32, bool)`

GetNSignals365dOk returns a tuple with the NSignals365d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSignals365d

`func (o *GxpFacilitiesRow) SetNSignals365d(v int32)`

SetNSignals365d sets NSignals365d field to given value.

### HasNSignals365d

`func (o *GxpFacilitiesRow) HasNSignals365d() bool`

HasNSignals365d returns a boolean if a field has been set.

### SetNSignals365dNil

`func (o *GxpFacilitiesRow) SetNSignals365dNil(b bool)`

 SetNSignals365dNil sets the value for NSignals365d to be an explicit nil

### UnsetNSignals365d
`func (o *GxpFacilitiesRow) UnsetNSignals365d()`

UnsetNSignals365d ensures that no value is present for NSignals365d, not even an explicit nil
### GetNInspections

`func (o *GxpFacilitiesRow) GetNInspections() int32`

GetNInspections returns the NInspections field if non-nil, zero value otherwise.

### GetNInspectionsOk

`func (o *GxpFacilitiesRow) GetNInspectionsOk() (*int32, bool)`

GetNInspectionsOk returns a tuple with the NInspections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNInspections

`func (o *GxpFacilitiesRow) SetNInspections(v int32)`

SetNInspections sets NInspections field to given value.

### HasNInspections

`func (o *GxpFacilitiesRow) HasNInspections() bool`

HasNInspections returns a boolean if a field has been set.

### SetNInspectionsNil

`func (o *GxpFacilitiesRow) SetNInspectionsNil(b bool)`

 SetNInspectionsNil sets the value for NInspections to be an explicit nil

### UnsetNInspections
`func (o *GxpFacilitiesRow) UnsetNInspections()`

UnsetNInspections ensures that no value is present for NInspections, not even an explicit nil
### GetLatestInspectionDate

`func (o *GxpFacilitiesRow) GetLatestInspectionDate() string`

GetLatestInspectionDate returns the LatestInspectionDate field if non-nil, zero value otherwise.

### GetLatestInspectionDateOk

`func (o *GxpFacilitiesRow) GetLatestInspectionDateOk() (*string, bool)`

GetLatestInspectionDateOk returns a tuple with the LatestInspectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestInspectionDate

`func (o *GxpFacilitiesRow) SetLatestInspectionDate(v string)`

SetLatestInspectionDate sets LatestInspectionDate field to given value.

### HasLatestInspectionDate

`func (o *GxpFacilitiesRow) HasLatestInspectionDate() bool`

HasLatestInspectionDate returns a boolean if a field has been set.

### SetLatestInspectionDateNil

`func (o *GxpFacilitiesRow) SetLatestInspectionDateNil(b bool)`

 SetLatestInspectionDateNil sets the value for LatestInspectionDate to be an explicit nil

### UnsetLatestInspectionDate
`func (o *GxpFacilitiesRow) UnsetLatestInspectionDate()`

UnsetLatestInspectionDate ensures that no value is present for LatestInspectionDate, not even an explicit nil
### GetNFdaInspections

`func (o *GxpFacilitiesRow) GetNFdaInspections() int32`

GetNFdaInspections returns the NFdaInspections field if non-nil, zero value otherwise.

### GetNFdaInspectionsOk

`func (o *GxpFacilitiesRow) GetNFdaInspectionsOk() (*int32, bool)`

GetNFdaInspectionsOk returns a tuple with the NFdaInspections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFdaInspections

`func (o *GxpFacilitiesRow) SetNFdaInspections(v int32)`

SetNFdaInspections sets NFdaInspections field to given value.

### HasNFdaInspections

`func (o *GxpFacilitiesRow) HasNFdaInspections() bool`

HasNFdaInspections returns a boolean if a field has been set.

### SetNFdaInspectionsNil

`func (o *GxpFacilitiesRow) SetNFdaInspectionsNil(b bool)`

 SetNFdaInspectionsNil sets the value for NFdaInspections to be an explicit nil

### UnsetNFdaInspections
`func (o *GxpFacilitiesRow) UnsetNFdaInspections()`

UnsetNFdaInspections ensures that no value is present for NFdaInspections, not even an explicit nil
### GetLatestFdaInspectionDate

`func (o *GxpFacilitiesRow) GetLatestFdaInspectionDate() string`

GetLatestFdaInspectionDate returns the LatestFdaInspectionDate field if non-nil, zero value otherwise.

### GetLatestFdaInspectionDateOk

`func (o *GxpFacilitiesRow) GetLatestFdaInspectionDateOk() (*string, bool)`

GetLatestFdaInspectionDateOk returns a tuple with the LatestFdaInspectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestFdaInspectionDate

`func (o *GxpFacilitiesRow) SetLatestFdaInspectionDate(v string)`

SetLatestFdaInspectionDate sets LatestFdaInspectionDate field to given value.

### HasLatestFdaInspectionDate

`func (o *GxpFacilitiesRow) HasLatestFdaInspectionDate() bool`

HasLatestFdaInspectionDate returns a boolean if a field has been set.

### SetLatestFdaInspectionDateNil

`func (o *GxpFacilitiesRow) SetLatestFdaInspectionDateNil(b bool)`

 SetLatestFdaInspectionDateNil sets the value for LatestFdaInspectionDate to be an explicit nil

### UnsetLatestFdaInspectionDate
`func (o *GxpFacilitiesRow) UnsetLatestFdaInspectionDate()`

UnsetLatestFdaInspectionDate ensures that no value is present for LatestFdaInspectionDate, not even an explicit nil
### GetLatestInspectionClassification

`func (o *GxpFacilitiesRow) GetLatestInspectionClassification() string`

GetLatestInspectionClassification returns the LatestInspectionClassification field if non-nil, zero value otherwise.

### GetLatestInspectionClassificationOk

`func (o *GxpFacilitiesRow) GetLatestInspectionClassificationOk() (*string, bool)`

GetLatestInspectionClassificationOk returns a tuple with the LatestInspectionClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestInspectionClassification

`func (o *GxpFacilitiesRow) SetLatestInspectionClassification(v string)`

SetLatestInspectionClassification sets LatestInspectionClassification field to given value.

### HasLatestInspectionClassification

`func (o *GxpFacilitiesRow) HasLatestInspectionClassification() bool`

HasLatestInspectionClassification returns a boolean if a field has been set.

### SetLatestInspectionClassificationNil

`func (o *GxpFacilitiesRow) SetLatestInspectionClassificationNil(b bool)`

 SetLatestInspectionClassificationNil sets the value for LatestInspectionClassification to be an explicit nil

### UnsetLatestInspectionClassification
`func (o *GxpFacilitiesRow) UnsetLatestInspectionClassification()`

UnsetLatestInspectionClassification ensures that no value is present for LatestInspectionClassification, not even an explicit nil
### GetNInspectionsNai

`func (o *GxpFacilitiesRow) GetNInspectionsNai() int32`

GetNInspectionsNai returns the NInspectionsNai field if non-nil, zero value otherwise.

### GetNInspectionsNaiOk

`func (o *GxpFacilitiesRow) GetNInspectionsNaiOk() (*int32, bool)`

GetNInspectionsNaiOk returns a tuple with the NInspectionsNai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNInspectionsNai

`func (o *GxpFacilitiesRow) SetNInspectionsNai(v int32)`

SetNInspectionsNai sets NInspectionsNai field to given value.

### HasNInspectionsNai

`func (o *GxpFacilitiesRow) HasNInspectionsNai() bool`

HasNInspectionsNai returns a boolean if a field has been set.

### SetNInspectionsNaiNil

`func (o *GxpFacilitiesRow) SetNInspectionsNaiNil(b bool)`

 SetNInspectionsNaiNil sets the value for NInspectionsNai to be an explicit nil

### UnsetNInspectionsNai
`func (o *GxpFacilitiesRow) UnsetNInspectionsNai()`

UnsetNInspectionsNai ensures that no value is present for NInspectionsNai, not even an explicit nil
### GetNInspectionsVai

`func (o *GxpFacilitiesRow) GetNInspectionsVai() int32`

GetNInspectionsVai returns the NInspectionsVai field if non-nil, zero value otherwise.

### GetNInspectionsVaiOk

`func (o *GxpFacilitiesRow) GetNInspectionsVaiOk() (*int32, bool)`

GetNInspectionsVaiOk returns a tuple with the NInspectionsVai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNInspectionsVai

`func (o *GxpFacilitiesRow) SetNInspectionsVai(v int32)`

SetNInspectionsVai sets NInspectionsVai field to given value.

### HasNInspectionsVai

`func (o *GxpFacilitiesRow) HasNInspectionsVai() bool`

HasNInspectionsVai returns a boolean if a field has been set.

### SetNInspectionsVaiNil

`func (o *GxpFacilitiesRow) SetNInspectionsVaiNil(b bool)`

 SetNInspectionsVaiNil sets the value for NInspectionsVai to be an explicit nil

### UnsetNInspectionsVai
`func (o *GxpFacilitiesRow) UnsetNInspectionsVai()`

UnsetNInspectionsVai ensures that no value is present for NInspectionsVai, not even an explicit nil
### GetNInspectionsOai

`func (o *GxpFacilitiesRow) GetNInspectionsOai() int32`

GetNInspectionsOai returns the NInspectionsOai field if non-nil, zero value otherwise.

### GetNInspectionsOaiOk

`func (o *GxpFacilitiesRow) GetNInspectionsOaiOk() (*int32, bool)`

GetNInspectionsOaiOk returns a tuple with the NInspectionsOai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNInspectionsOai

`func (o *GxpFacilitiesRow) SetNInspectionsOai(v int32)`

SetNInspectionsOai sets NInspectionsOai field to given value.

### HasNInspectionsOai

`func (o *GxpFacilitiesRow) HasNInspectionsOai() bool`

HasNInspectionsOai returns a boolean if a field has been set.

### SetNInspectionsOaiNil

`func (o *GxpFacilitiesRow) SetNInspectionsOaiNil(b bool)`

 SetNInspectionsOaiNil sets the value for NInspectionsOai to be an explicit nil

### UnsetNInspectionsOai
`func (o *GxpFacilitiesRow) UnsetNInspectionsOai()`

UnsetNInspectionsOai ensures that no value is present for NInspectionsOai, not even an explicit nil
### GetN483

`func (o *GxpFacilitiesRow) GetN483() int32`

GetN483 returns the N483 field if non-nil, zero value otherwise.

### GetN483Ok

`func (o *GxpFacilitiesRow) GetN483Ok() (*int32, bool)`

GetN483Ok returns a tuple with the N483 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN483

`func (o *GxpFacilitiesRow) SetN483(v int32)`

SetN483 sets N483 field to given value.

### HasN483

`func (o *GxpFacilitiesRow) HasN483() bool`

HasN483 returns a boolean if a field has been set.

### SetN483Nil

`func (o *GxpFacilitiesRow) SetN483Nil(b bool)`

 SetN483Nil sets the value for N483 to be an explicit nil

### UnsetN483
`func (o *GxpFacilitiesRow) UnsetN483()`

UnsetN483 ensures that no value is present for N483, not even an explicit nil
### GetNHcInspections

`func (o *GxpFacilitiesRow) GetNHcInspections() int32`

GetNHcInspections returns the NHcInspections field if non-nil, zero value otherwise.

### GetNHcInspectionsOk

`func (o *GxpFacilitiesRow) GetNHcInspectionsOk() (*int32, bool)`

GetNHcInspectionsOk returns a tuple with the NHcInspections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNHcInspections

`func (o *GxpFacilitiesRow) SetNHcInspections(v int32)`

SetNHcInspections sets NHcInspections field to given value.

### HasNHcInspections

`func (o *GxpFacilitiesRow) HasNHcInspections() bool`

HasNHcInspections returns a boolean if a field has been set.

### SetNHcInspectionsNil

`func (o *GxpFacilitiesRow) SetNHcInspectionsNil(b bool)`

 SetNHcInspectionsNil sets the value for NHcInspections to be an explicit nil

### UnsetNHcInspections
`func (o *GxpFacilitiesRow) UnsetNHcInspections()`

UnsetNHcInspections ensures that no value is present for NHcInspections, not even an explicit nil
### GetLatestHcInspectionDate

`func (o *GxpFacilitiesRow) GetLatestHcInspectionDate() string`

GetLatestHcInspectionDate returns the LatestHcInspectionDate field if non-nil, zero value otherwise.

### GetLatestHcInspectionDateOk

`func (o *GxpFacilitiesRow) GetLatestHcInspectionDateOk() (*string, bool)`

GetLatestHcInspectionDateOk returns a tuple with the LatestHcInspectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestHcInspectionDate

`func (o *GxpFacilitiesRow) SetLatestHcInspectionDate(v string)`

SetLatestHcInspectionDate sets LatestHcInspectionDate field to given value.

### HasLatestHcInspectionDate

`func (o *GxpFacilitiesRow) HasLatestHcInspectionDate() bool`

HasLatestHcInspectionDate returns a boolean if a field has been set.

### SetLatestHcInspectionDateNil

`func (o *GxpFacilitiesRow) SetLatestHcInspectionDateNil(b bool)`

 SetLatestHcInspectionDateNil sets the value for LatestHcInspectionDate to be an explicit nil

### UnsetLatestHcInspectionDate
`func (o *GxpFacilitiesRow) UnsetLatestHcInspectionDate()`

UnsetLatestHcInspectionDate ensures that no value is present for LatestHcInspectionDate, not even an explicit nil
### GetLatestHcInspectionRating

`func (o *GxpFacilitiesRow) GetLatestHcInspectionRating() string`

GetLatestHcInspectionRating returns the LatestHcInspectionRating field if non-nil, zero value otherwise.

### GetLatestHcInspectionRatingOk

`func (o *GxpFacilitiesRow) GetLatestHcInspectionRatingOk() (*string, bool)`

GetLatestHcInspectionRatingOk returns a tuple with the LatestHcInspectionRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestHcInspectionRating

`func (o *GxpFacilitiesRow) SetLatestHcInspectionRating(v string)`

SetLatestHcInspectionRating sets LatestHcInspectionRating field to given value.

### HasLatestHcInspectionRating

`func (o *GxpFacilitiesRow) HasLatestHcInspectionRating() bool`

HasLatestHcInspectionRating returns a boolean if a field has been set.

### SetLatestHcInspectionRatingNil

`func (o *GxpFacilitiesRow) SetLatestHcInspectionRatingNil(b bool)`

 SetLatestHcInspectionRatingNil sets the value for LatestHcInspectionRating to be an explicit nil

### UnsetLatestHcInspectionRating
`func (o *GxpFacilitiesRow) UnsetLatestHcInspectionRating()`

UnsetLatestHcInspectionRating ensures that no value is present for LatestHcInspectionRating, not even an explicit nil
### GetNHcInspectionsNonCompliant

`func (o *GxpFacilitiesRow) GetNHcInspectionsNonCompliant() int32`

GetNHcInspectionsNonCompliant returns the NHcInspectionsNonCompliant field if non-nil, zero value otherwise.

### GetNHcInspectionsNonCompliantOk

`func (o *GxpFacilitiesRow) GetNHcInspectionsNonCompliantOk() (*int32, bool)`

GetNHcInspectionsNonCompliantOk returns a tuple with the NHcInspectionsNonCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNHcInspectionsNonCompliant

`func (o *GxpFacilitiesRow) SetNHcInspectionsNonCompliant(v int32)`

SetNHcInspectionsNonCompliant sets NHcInspectionsNonCompliant field to given value.

### HasNHcInspectionsNonCompliant

`func (o *GxpFacilitiesRow) HasNHcInspectionsNonCompliant() bool`

HasNHcInspectionsNonCompliant returns a boolean if a field has been set.

### SetNHcInspectionsNonCompliantNil

`func (o *GxpFacilitiesRow) SetNHcInspectionsNonCompliantNil(b bool)`

 SetNHcInspectionsNonCompliantNil sets the value for NHcInspectionsNonCompliant to be an explicit nil

### UnsetNHcInspectionsNonCompliant
`func (o *GxpFacilitiesRow) UnsetNHcInspectionsNonCompliant()`

UnsetNHcInspectionsNonCompliant ensures that no value is present for NHcInspectionsNonCompliant, not even an explicit nil
### GetNImportRefusals

`func (o *GxpFacilitiesRow) GetNImportRefusals() int32`

GetNImportRefusals returns the NImportRefusals field if non-nil, zero value otherwise.

### GetNImportRefusalsOk

`func (o *GxpFacilitiesRow) GetNImportRefusalsOk() (*int32, bool)`

GetNImportRefusalsOk returns a tuple with the NImportRefusals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNImportRefusals

`func (o *GxpFacilitiesRow) SetNImportRefusals(v int32)`

SetNImportRefusals sets NImportRefusals field to given value.

### HasNImportRefusals

`func (o *GxpFacilitiesRow) HasNImportRefusals() bool`

HasNImportRefusals returns a boolean if a field has been set.

### SetNImportRefusalsNil

`func (o *GxpFacilitiesRow) SetNImportRefusalsNil(b bool)`

 SetNImportRefusalsNil sets the value for NImportRefusals to be an explicit nil

### UnsetNImportRefusals
`func (o *GxpFacilitiesRow) UnsetNImportRefusals()`

UnsetNImportRefusals ensures that no value is present for NImportRefusals, not even an explicit nil
### GetLatestImportRefusalDate

`func (o *GxpFacilitiesRow) GetLatestImportRefusalDate() string`

GetLatestImportRefusalDate returns the LatestImportRefusalDate field if non-nil, zero value otherwise.

### GetLatestImportRefusalDateOk

`func (o *GxpFacilitiesRow) GetLatestImportRefusalDateOk() (*string, bool)`

GetLatestImportRefusalDateOk returns a tuple with the LatestImportRefusalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestImportRefusalDate

`func (o *GxpFacilitiesRow) SetLatestImportRefusalDate(v string)`

SetLatestImportRefusalDate sets LatestImportRefusalDate field to given value.

### HasLatestImportRefusalDate

`func (o *GxpFacilitiesRow) HasLatestImportRefusalDate() bool`

HasLatestImportRefusalDate returns a boolean if a field has been set.

### SetLatestImportRefusalDateNil

`func (o *GxpFacilitiesRow) SetLatestImportRefusalDateNil(b bool)`

 SetLatestImportRefusalDateNil sets the value for LatestImportRefusalDate to be an explicit nil

### UnsetLatestImportRefusalDate
`func (o *GxpFacilitiesRow) UnsetLatestImportRefusalDate()`

UnsetLatestImportRefusalDate ensures that no value is present for LatestImportRefusalDate, not even an explicit nil
### GetNGmpCertificates

`func (o *GxpFacilitiesRow) GetNGmpCertificates() int32`

GetNGmpCertificates returns the NGmpCertificates field if non-nil, zero value otherwise.

### GetNGmpCertificatesOk

`func (o *GxpFacilitiesRow) GetNGmpCertificatesOk() (*int32, bool)`

GetNGmpCertificatesOk returns a tuple with the NGmpCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNGmpCertificates

`func (o *GxpFacilitiesRow) SetNGmpCertificates(v int32)`

SetNGmpCertificates sets NGmpCertificates field to given value.

### HasNGmpCertificates

`func (o *GxpFacilitiesRow) HasNGmpCertificates() bool`

HasNGmpCertificates returns a boolean if a field has been set.

### SetNGmpCertificatesNil

`func (o *GxpFacilitiesRow) SetNGmpCertificatesNil(b bool)`

 SetNGmpCertificatesNil sets the value for NGmpCertificates to be an explicit nil

### UnsetNGmpCertificates
`func (o *GxpFacilitiesRow) UnsetNGmpCertificates()`

UnsetNGmpCertificates ensures that no value is present for NGmpCertificates, not even an explicit nil
### GetLatestGmpCertificateDate

`func (o *GxpFacilitiesRow) GetLatestGmpCertificateDate() string`

GetLatestGmpCertificateDate returns the LatestGmpCertificateDate field if non-nil, zero value otherwise.

### GetLatestGmpCertificateDateOk

`func (o *GxpFacilitiesRow) GetLatestGmpCertificateDateOk() (*string, bool)`

GetLatestGmpCertificateDateOk returns a tuple with the LatestGmpCertificateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestGmpCertificateDate

`func (o *GxpFacilitiesRow) SetLatestGmpCertificateDate(v string)`

SetLatestGmpCertificateDate sets LatestGmpCertificateDate field to given value.

### HasLatestGmpCertificateDate

`func (o *GxpFacilitiesRow) HasLatestGmpCertificateDate() bool`

HasLatestGmpCertificateDate returns a boolean if a field has been set.

### SetLatestGmpCertificateDateNil

`func (o *GxpFacilitiesRow) SetLatestGmpCertificateDateNil(b bool)`

 SetLatestGmpCertificateDateNil sets the value for LatestGmpCertificateDate to be an explicit nil

### UnsetLatestGmpCertificateDate
`func (o *GxpFacilitiesRow) UnsetLatestGmpCertificateDate()`

UnsetLatestGmpCertificateDate ensures that no value is present for LatestGmpCertificateDate, not even an explicit nil
### GetNNoncomplianceStatements

`func (o *GxpFacilitiesRow) GetNNoncomplianceStatements() int32`

GetNNoncomplianceStatements returns the NNoncomplianceStatements field if non-nil, zero value otherwise.

### GetNNoncomplianceStatementsOk

`func (o *GxpFacilitiesRow) GetNNoncomplianceStatementsOk() (*int32, bool)`

GetNNoncomplianceStatementsOk returns a tuple with the NNoncomplianceStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNNoncomplianceStatements

`func (o *GxpFacilitiesRow) SetNNoncomplianceStatements(v int32)`

SetNNoncomplianceStatements sets NNoncomplianceStatements field to given value.

### HasNNoncomplianceStatements

`func (o *GxpFacilitiesRow) HasNNoncomplianceStatements() bool`

HasNNoncomplianceStatements returns a boolean if a field has been set.

### SetNNoncomplianceStatementsNil

`func (o *GxpFacilitiesRow) SetNNoncomplianceStatementsNil(b bool)`

 SetNNoncomplianceStatementsNil sets the value for NNoncomplianceStatements to be an explicit nil

### UnsetNNoncomplianceStatements
`func (o *GxpFacilitiesRow) UnsetNNoncomplianceStatements()`

UnsetNNoncomplianceStatements ensures that no value is present for NNoncomplianceStatements, not even an explicit nil
### GetLatestNoncomplianceStatementDate

`func (o *GxpFacilitiesRow) GetLatestNoncomplianceStatementDate() string`

GetLatestNoncomplianceStatementDate returns the LatestNoncomplianceStatementDate field if non-nil, zero value otherwise.

### GetLatestNoncomplianceStatementDateOk

`func (o *GxpFacilitiesRow) GetLatestNoncomplianceStatementDateOk() (*string, bool)`

GetLatestNoncomplianceStatementDateOk returns a tuple with the LatestNoncomplianceStatementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestNoncomplianceStatementDate

`func (o *GxpFacilitiesRow) SetLatestNoncomplianceStatementDate(v string)`

SetLatestNoncomplianceStatementDate sets LatestNoncomplianceStatementDate field to given value.

### HasLatestNoncomplianceStatementDate

`func (o *GxpFacilitiesRow) HasLatestNoncomplianceStatementDate() bool`

HasLatestNoncomplianceStatementDate returns a boolean if a field has been set.

### SetLatestNoncomplianceStatementDateNil

`func (o *GxpFacilitiesRow) SetLatestNoncomplianceStatementDateNil(b bool)`

 SetLatestNoncomplianceStatementDateNil sets the value for LatestNoncomplianceStatementDate to be an explicit nil

### UnsetLatestNoncomplianceStatementDate
`func (o *GxpFacilitiesRow) UnsetLatestNoncomplianceStatementDate()`

UnsetLatestNoncomplianceStatementDate ensures that no value is present for LatestNoncomplianceStatementDate, not even an explicit nil
### GetNLicences

`func (o *GxpFacilitiesRow) GetNLicences() int32`

GetNLicences returns the NLicences field if non-nil, zero value otherwise.

### GetNLicencesOk

`func (o *GxpFacilitiesRow) GetNLicencesOk() (*int32, bool)`

GetNLicencesOk returns a tuple with the NLicences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNLicences

`func (o *GxpFacilitiesRow) SetNLicences(v int32)`

SetNLicences sets NLicences field to given value.

### HasNLicences

`func (o *GxpFacilitiesRow) HasNLicences() bool`

HasNLicences returns a boolean if a field has been set.

### SetNLicencesNil

`func (o *GxpFacilitiesRow) SetNLicencesNil(b bool)`

 SetNLicencesNil sets the value for NLicences to be an explicit nil

### UnsetNLicences
`func (o *GxpFacilitiesRow) UnsetNLicences()`

UnsetNLicences ensures that no value is present for NLicences, not even an explicit nil
### GetLicenceStatus

`func (o *GxpFacilitiesRow) GetLicenceStatus() string`

GetLicenceStatus returns the LicenceStatus field if non-nil, zero value otherwise.

### GetLicenceStatusOk

`func (o *GxpFacilitiesRow) GetLicenceStatusOk() (*string, bool)`

GetLicenceStatusOk returns a tuple with the LicenceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenceStatus

`func (o *GxpFacilitiesRow) SetLicenceStatus(v string)`

SetLicenceStatus sets LicenceStatus field to given value.

### HasLicenceStatus

`func (o *GxpFacilitiesRow) HasLicenceStatus() bool`

HasLicenceStatus returns a boolean if a field has been set.

### SetLicenceStatusNil

`func (o *GxpFacilitiesRow) SetLicenceStatusNil(b bool)`

 SetLicenceStatusNil sets the value for LicenceStatus to be an explicit nil

### UnsetLicenceStatus
`func (o *GxpFacilitiesRow) UnsetLicenceStatus()`

UnsetLicenceStatus ensures that no value is present for LicenceStatus, not even an explicit nil
### GetOrgNLicenceActions

`func (o *GxpFacilitiesRow) GetOrgNLicenceActions() int32`

GetOrgNLicenceActions returns the OrgNLicenceActions field if non-nil, zero value otherwise.

### GetOrgNLicenceActionsOk

`func (o *GxpFacilitiesRow) GetOrgNLicenceActionsOk() (*int32, bool)`

GetOrgNLicenceActionsOk returns a tuple with the OrgNLicenceActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgNLicenceActions

`func (o *GxpFacilitiesRow) SetOrgNLicenceActions(v int32)`

SetOrgNLicenceActions sets OrgNLicenceActions field to given value.

### HasOrgNLicenceActions

`func (o *GxpFacilitiesRow) HasOrgNLicenceActions() bool`

HasOrgNLicenceActions returns a boolean if a field has been set.

### SetOrgNLicenceActionsNil

`func (o *GxpFacilitiesRow) SetOrgNLicenceActionsNil(b bool)`

 SetOrgNLicenceActionsNil sets the value for OrgNLicenceActions to be an explicit nil

### UnsetOrgNLicenceActions
`func (o *GxpFacilitiesRow) UnsetOrgNLicenceActions()`

UnsetOrgNLicenceActions ensures that no value is present for OrgNLicenceActions, not even an explicit nil
### GetOrgLatestLicenceActionDate

`func (o *GxpFacilitiesRow) GetOrgLatestLicenceActionDate() string`

GetOrgLatestLicenceActionDate returns the OrgLatestLicenceActionDate field if non-nil, zero value otherwise.

### GetOrgLatestLicenceActionDateOk

`func (o *GxpFacilitiesRow) GetOrgLatestLicenceActionDateOk() (*string, bool)`

GetOrgLatestLicenceActionDateOk returns a tuple with the OrgLatestLicenceActionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgLatestLicenceActionDate

`func (o *GxpFacilitiesRow) SetOrgLatestLicenceActionDate(v string)`

SetOrgLatestLicenceActionDate sets OrgLatestLicenceActionDate field to given value.

### HasOrgLatestLicenceActionDate

`func (o *GxpFacilitiesRow) HasOrgLatestLicenceActionDate() bool`

HasOrgLatestLicenceActionDate returns a boolean if a field has been set.

### SetOrgLatestLicenceActionDateNil

`func (o *GxpFacilitiesRow) SetOrgLatestLicenceActionDateNil(b bool)`

 SetOrgLatestLicenceActionDateNil sets the value for OrgLatestLicenceActionDate to be an explicit nil

### UnsetOrgLatestLicenceActionDate
`func (o *GxpFacilitiesRow) UnsetOrgLatestLicenceActionDate()`

UnsetOrgLatestLicenceActionDate ensures that no value is present for OrgLatestLicenceActionDate, not even an explicit nil
### GetAsOf

`func (o *GxpFacilitiesRow) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *GxpFacilitiesRow) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *GxpFacilitiesRow) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *GxpFacilitiesRow) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *GxpFacilitiesRow) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *GxpFacilitiesRow) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


