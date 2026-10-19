# SiteHospitalsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **NullableString** | The site key: one facility spelling from the registry. | [optional] 
**Ccn** | Pointer to **NullableString** | CMS Certification Number of the matched hospital. | [optional] 
**HospitalName** | Pointer to **NullableString** | The hospital&#39;s name in the CMS registry. | [optional] 
**HospitalType** | Pointer to **NullableString** | The CMS hospital type. | [optional] 
**Ownership** | Pointer to **NullableString** | The CMS ownership category. | [optional] 
**BedCount** | Pointer to **NullableInt32** | Certified bed count in the latest registry file. | [optional] 
**IsTeaching** | Pointer to **NullableBool** | Whether the hospital is a teaching hospital. | [optional] 
**HospitalLatitude** | Pointer to **NullableFloat32** | Latitude of the hospital. | [optional] 
**HospitalLongitude** | Pointer to **NullableFloat32** | Longitude of the hospital. | [optional] 
**GeocodeMethod** | Pointer to **NullableString** | How the hospital&#39;s coordinate was placed. | [optional] 
**CcnNTrials** | Pointer to **NullableInt64** | Registered trials across the hospital&#39;s matched site records. | [optional] 
**CcnNSiteVariants** | Pointer to **NullableInt64** | Site records matched to the same hospital. | [optional] 
**IsPrimarySite** | Pointer to **NullableBool** | True for the matched site record with the most trials for this hospital. | [optional] 

## Methods

### NewSiteHospitalsRow

`func NewSiteHospitalsRow() *SiteHospitalsRow`

NewSiteHospitalsRow instantiates a new SiteHospitalsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteHospitalsRowWithDefaults

`func NewSiteHospitalsRowWithDefaults() *SiteHospitalsRow`

NewSiteHospitalsRowWithDefaults instantiates a new SiteHospitalsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *SiteHospitalsRow) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SiteHospitalsRow) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SiteHospitalsRow) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SiteHospitalsRow) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *SiteHospitalsRow) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *SiteHospitalsRow) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetCcn

`func (o *SiteHospitalsRow) GetCcn() string`

GetCcn returns the Ccn field if non-nil, zero value otherwise.

### GetCcnOk

`func (o *SiteHospitalsRow) GetCcnOk() (*string, bool)`

GetCcnOk returns a tuple with the Ccn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcn

`func (o *SiteHospitalsRow) SetCcn(v string)`

SetCcn sets Ccn field to given value.

### HasCcn

`func (o *SiteHospitalsRow) HasCcn() bool`

HasCcn returns a boolean if a field has been set.

### SetCcnNil

`func (o *SiteHospitalsRow) SetCcnNil(b bool)`

 SetCcnNil sets the value for Ccn to be an explicit nil

### UnsetCcn
`func (o *SiteHospitalsRow) UnsetCcn()`

UnsetCcn ensures that no value is present for Ccn, not even an explicit nil
### GetHospitalName

`func (o *SiteHospitalsRow) GetHospitalName() string`

GetHospitalName returns the HospitalName field if non-nil, zero value otherwise.

### GetHospitalNameOk

`func (o *SiteHospitalsRow) GetHospitalNameOk() (*string, bool)`

GetHospitalNameOk returns a tuple with the HospitalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHospitalName

`func (o *SiteHospitalsRow) SetHospitalName(v string)`

SetHospitalName sets HospitalName field to given value.

### HasHospitalName

`func (o *SiteHospitalsRow) HasHospitalName() bool`

HasHospitalName returns a boolean if a field has been set.

### SetHospitalNameNil

`func (o *SiteHospitalsRow) SetHospitalNameNil(b bool)`

 SetHospitalNameNil sets the value for HospitalName to be an explicit nil

### UnsetHospitalName
`func (o *SiteHospitalsRow) UnsetHospitalName()`

UnsetHospitalName ensures that no value is present for HospitalName, not even an explicit nil
### GetHospitalType

`func (o *SiteHospitalsRow) GetHospitalType() string`

GetHospitalType returns the HospitalType field if non-nil, zero value otherwise.

### GetHospitalTypeOk

`func (o *SiteHospitalsRow) GetHospitalTypeOk() (*string, bool)`

GetHospitalTypeOk returns a tuple with the HospitalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHospitalType

`func (o *SiteHospitalsRow) SetHospitalType(v string)`

SetHospitalType sets HospitalType field to given value.

### HasHospitalType

`func (o *SiteHospitalsRow) HasHospitalType() bool`

HasHospitalType returns a boolean if a field has been set.

### SetHospitalTypeNil

`func (o *SiteHospitalsRow) SetHospitalTypeNil(b bool)`

 SetHospitalTypeNil sets the value for HospitalType to be an explicit nil

### UnsetHospitalType
`func (o *SiteHospitalsRow) UnsetHospitalType()`

UnsetHospitalType ensures that no value is present for HospitalType, not even an explicit nil
### GetOwnership

`func (o *SiteHospitalsRow) GetOwnership() string`

GetOwnership returns the Ownership field if non-nil, zero value otherwise.

### GetOwnershipOk

`func (o *SiteHospitalsRow) GetOwnershipOk() (*string, bool)`

GetOwnershipOk returns a tuple with the Ownership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnership

`func (o *SiteHospitalsRow) SetOwnership(v string)`

SetOwnership sets Ownership field to given value.

### HasOwnership

`func (o *SiteHospitalsRow) HasOwnership() bool`

HasOwnership returns a boolean if a field has been set.

### SetOwnershipNil

`func (o *SiteHospitalsRow) SetOwnershipNil(b bool)`

 SetOwnershipNil sets the value for Ownership to be an explicit nil

### UnsetOwnership
`func (o *SiteHospitalsRow) UnsetOwnership()`

UnsetOwnership ensures that no value is present for Ownership, not even an explicit nil
### GetBedCount

`func (o *SiteHospitalsRow) GetBedCount() int32`

GetBedCount returns the BedCount field if non-nil, zero value otherwise.

### GetBedCountOk

`func (o *SiteHospitalsRow) GetBedCountOk() (*int32, bool)`

GetBedCountOk returns a tuple with the BedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBedCount

`func (o *SiteHospitalsRow) SetBedCount(v int32)`

SetBedCount sets BedCount field to given value.

### HasBedCount

`func (o *SiteHospitalsRow) HasBedCount() bool`

HasBedCount returns a boolean if a field has been set.

### SetBedCountNil

`func (o *SiteHospitalsRow) SetBedCountNil(b bool)`

 SetBedCountNil sets the value for BedCount to be an explicit nil

### UnsetBedCount
`func (o *SiteHospitalsRow) UnsetBedCount()`

UnsetBedCount ensures that no value is present for BedCount, not even an explicit nil
### GetIsTeaching

`func (o *SiteHospitalsRow) GetIsTeaching() bool`

GetIsTeaching returns the IsTeaching field if non-nil, zero value otherwise.

### GetIsTeachingOk

`func (o *SiteHospitalsRow) GetIsTeachingOk() (*bool, bool)`

GetIsTeachingOk returns a tuple with the IsTeaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTeaching

`func (o *SiteHospitalsRow) SetIsTeaching(v bool)`

SetIsTeaching sets IsTeaching field to given value.

### HasIsTeaching

`func (o *SiteHospitalsRow) HasIsTeaching() bool`

HasIsTeaching returns a boolean if a field has been set.

### SetIsTeachingNil

`func (o *SiteHospitalsRow) SetIsTeachingNil(b bool)`

 SetIsTeachingNil sets the value for IsTeaching to be an explicit nil

### UnsetIsTeaching
`func (o *SiteHospitalsRow) UnsetIsTeaching()`

UnsetIsTeaching ensures that no value is present for IsTeaching, not even an explicit nil
### GetHospitalLatitude

`func (o *SiteHospitalsRow) GetHospitalLatitude() float32`

GetHospitalLatitude returns the HospitalLatitude field if non-nil, zero value otherwise.

### GetHospitalLatitudeOk

`func (o *SiteHospitalsRow) GetHospitalLatitudeOk() (*float32, bool)`

GetHospitalLatitudeOk returns a tuple with the HospitalLatitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHospitalLatitude

`func (o *SiteHospitalsRow) SetHospitalLatitude(v float32)`

SetHospitalLatitude sets HospitalLatitude field to given value.

### HasHospitalLatitude

`func (o *SiteHospitalsRow) HasHospitalLatitude() bool`

HasHospitalLatitude returns a boolean if a field has been set.

### SetHospitalLatitudeNil

`func (o *SiteHospitalsRow) SetHospitalLatitudeNil(b bool)`

 SetHospitalLatitudeNil sets the value for HospitalLatitude to be an explicit nil

### UnsetHospitalLatitude
`func (o *SiteHospitalsRow) UnsetHospitalLatitude()`

UnsetHospitalLatitude ensures that no value is present for HospitalLatitude, not even an explicit nil
### GetHospitalLongitude

`func (o *SiteHospitalsRow) GetHospitalLongitude() float32`

GetHospitalLongitude returns the HospitalLongitude field if non-nil, zero value otherwise.

### GetHospitalLongitudeOk

`func (o *SiteHospitalsRow) GetHospitalLongitudeOk() (*float32, bool)`

GetHospitalLongitudeOk returns a tuple with the HospitalLongitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHospitalLongitude

`func (o *SiteHospitalsRow) SetHospitalLongitude(v float32)`

SetHospitalLongitude sets HospitalLongitude field to given value.

### HasHospitalLongitude

`func (o *SiteHospitalsRow) HasHospitalLongitude() bool`

HasHospitalLongitude returns a boolean if a field has been set.

### SetHospitalLongitudeNil

`func (o *SiteHospitalsRow) SetHospitalLongitudeNil(b bool)`

 SetHospitalLongitudeNil sets the value for HospitalLongitude to be an explicit nil

### UnsetHospitalLongitude
`func (o *SiteHospitalsRow) UnsetHospitalLongitude()`

UnsetHospitalLongitude ensures that no value is present for HospitalLongitude, not even an explicit nil
### GetGeocodeMethod

`func (o *SiteHospitalsRow) GetGeocodeMethod() string`

GetGeocodeMethod returns the GeocodeMethod field if non-nil, zero value otherwise.

### GetGeocodeMethodOk

`func (o *SiteHospitalsRow) GetGeocodeMethodOk() (*string, bool)`

GetGeocodeMethodOk returns a tuple with the GeocodeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocodeMethod

`func (o *SiteHospitalsRow) SetGeocodeMethod(v string)`

SetGeocodeMethod sets GeocodeMethod field to given value.

### HasGeocodeMethod

`func (o *SiteHospitalsRow) HasGeocodeMethod() bool`

HasGeocodeMethod returns a boolean if a field has been set.

### SetGeocodeMethodNil

`func (o *SiteHospitalsRow) SetGeocodeMethodNil(b bool)`

 SetGeocodeMethodNil sets the value for GeocodeMethod to be an explicit nil

### UnsetGeocodeMethod
`func (o *SiteHospitalsRow) UnsetGeocodeMethod()`

UnsetGeocodeMethod ensures that no value is present for GeocodeMethod, not even an explicit nil
### GetCcnNTrials

`func (o *SiteHospitalsRow) GetCcnNTrials() int64`

GetCcnNTrials returns the CcnNTrials field if non-nil, zero value otherwise.

### GetCcnNTrialsOk

`func (o *SiteHospitalsRow) GetCcnNTrialsOk() (*int64, bool)`

GetCcnNTrialsOk returns a tuple with the CcnNTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcnNTrials

`func (o *SiteHospitalsRow) SetCcnNTrials(v int64)`

SetCcnNTrials sets CcnNTrials field to given value.

### HasCcnNTrials

`func (o *SiteHospitalsRow) HasCcnNTrials() bool`

HasCcnNTrials returns a boolean if a field has been set.

### SetCcnNTrialsNil

`func (o *SiteHospitalsRow) SetCcnNTrialsNil(b bool)`

 SetCcnNTrialsNil sets the value for CcnNTrials to be an explicit nil

### UnsetCcnNTrials
`func (o *SiteHospitalsRow) UnsetCcnNTrials()`

UnsetCcnNTrials ensures that no value is present for CcnNTrials, not even an explicit nil
### GetCcnNSiteVariants

`func (o *SiteHospitalsRow) GetCcnNSiteVariants() int64`

GetCcnNSiteVariants returns the CcnNSiteVariants field if non-nil, zero value otherwise.

### GetCcnNSiteVariantsOk

`func (o *SiteHospitalsRow) GetCcnNSiteVariantsOk() (*int64, bool)`

GetCcnNSiteVariantsOk returns a tuple with the CcnNSiteVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcnNSiteVariants

`func (o *SiteHospitalsRow) SetCcnNSiteVariants(v int64)`

SetCcnNSiteVariants sets CcnNSiteVariants field to given value.

### HasCcnNSiteVariants

`func (o *SiteHospitalsRow) HasCcnNSiteVariants() bool`

HasCcnNSiteVariants returns a boolean if a field has been set.

### SetCcnNSiteVariantsNil

`func (o *SiteHospitalsRow) SetCcnNSiteVariantsNil(b bool)`

 SetCcnNSiteVariantsNil sets the value for CcnNSiteVariants to be an explicit nil

### UnsetCcnNSiteVariants
`func (o *SiteHospitalsRow) UnsetCcnNSiteVariants()`

UnsetCcnNSiteVariants ensures that no value is present for CcnNSiteVariants, not even an explicit nil
### GetIsPrimarySite

`func (o *SiteHospitalsRow) GetIsPrimarySite() bool`

GetIsPrimarySite returns the IsPrimarySite field if non-nil, zero value otherwise.

### GetIsPrimarySiteOk

`func (o *SiteHospitalsRow) GetIsPrimarySiteOk() (*bool, bool)`

GetIsPrimarySiteOk returns a tuple with the IsPrimarySite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimarySite

`func (o *SiteHospitalsRow) SetIsPrimarySite(v bool)`

SetIsPrimarySite sets IsPrimarySite field to given value.

### HasIsPrimarySite

`func (o *SiteHospitalsRow) HasIsPrimarySite() bool`

HasIsPrimarySite returns a boolean if a field has been set.

### SetIsPrimarySiteNil

`func (o *SiteHospitalsRow) SetIsPrimarySiteNil(b bool)`

 SetIsPrimarySiteNil sets the value for IsPrimarySite to be an explicit nil

### UnsetIsPrimarySite
`func (o *SiteHospitalsRow) UnsetIsPrimarySite()`

UnsetIsPrimarySite ensures that no value is present for IsPrimarySite, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


