# HospitalsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccn** | Pointer to **NullableString** | CMS Certification Number, the six-character Medicare provider id the hospital page address carries. | [optional] 
**HospitalName** | Pointer to **NullableString** | The hospital&#39;s facility name in the CMS registry. | [optional] 
**City** | Pointer to **NullableString** | The hospital&#39;s city. | [optional] 
**State** | Pointer to **NullableString** | The hospital&#39;s two-letter state or territory code. | [optional] 
**BedCount** | Pointer to **NullableInt32** | Certified bed count in the latest registry file. Null when the file has no bed record. | [optional] 
**HospitalType** | Pointer to **NullableString** | The CMS hospital type. | [optional] 
**Ownership** | Pointer to **NullableString** | The CMS ownership category. | [optional] 
**IsTeaching** | Pointer to **NullableBool** | Whether the hospital is a teaching hospital. | [optional] 
**EmergencyServices** | Pointer to **NullableString** | Whether the hospital provides emergency services, as CMS files it. Null where CMS left it blank. | [optional] 
**Latitude** | Pointer to **NullableFloat32** | Latitude of the hospital&#39;s own coordinate. Null exactly when geocode_method is unmatched. | [optional] 
**Longitude** | Pointer to **NullableFloat32** | Longitude of the hospital&#39;s own coordinate. Null exactly when geocode_method is unmatched. | [optional] 
**GeocodeMethod** | Pointer to **NullableString** | How the coordinate was placed: als (Amazon Location Service address match), zip_centroid (coarse ZIP-code centroid) or unmatched (no coordinate). | [optional] 
**HrrNum** | Pointer to **NullableInt32** | Dartmouth Atlas Hospital Referral Region number. Null when hrr_match is not zip. | [optional] 
**HrrName** | Pointer to **NullableString** | The Hospital Referral Region, named for its hub city and state. The hub state can differ from the hospital&#39;s own state. | [optional] 
**HsaNum** | Pointer to **NullableInt32** | Dartmouth Atlas Hospital Service Area number. Null when hrr_match is not zip. | [optional] 
**HsaName** | Pointer to **NullableString** | The Hospital Service Area, named for its hub city and state. | [optional] 
**HrrMatch** | Pointer to **NullableString** | zip (region found by ZIP), territory_uncovered (the Atlas does not cover US territories) or zip_unmatched (a mainland ZIP absent from the 2019 crosswalk). | [optional] 

## Methods

### NewHospitalsRow

`func NewHospitalsRow() *HospitalsRow`

NewHospitalsRow instantiates a new HospitalsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHospitalsRowWithDefaults

`func NewHospitalsRowWithDefaults() *HospitalsRow`

NewHospitalsRowWithDefaults instantiates a new HospitalsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcn

`func (o *HospitalsRow) GetCcn() string`

GetCcn returns the Ccn field if non-nil, zero value otherwise.

### GetCcnOk

`func (o *HospitalsRow) GetCcnOk() (*string, bool)`

GetCcnOk returns a tuple with the Ccn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcn

`func (o *HospitalsRow) SetCcn(v string)`

SetCcn sets Ccn field to given value.

### HasCcn

`func (o *HospitalsRow) HasCcn() bool`

HasCcn returns a boolean if a field has been set.

### SetCcnNil

`func (o *HospitalsRow) SetCcnNil(b bool)`

 SetCcnNil sets the value for Ccn to be an explicit nil

### UnsetCcn
`func (o *HospitalsRow) UnsetCcn()`

UnsetCcn ensures that no value is present for Ccn, not even an explicit nil
### GetHospitalName

`func (o *HospitalsRow) GetHospitalName() string`

GetHospitalName returns the HospitalName field if non-nil, zero value otherwise.

### GetHospitalNameOk

`func (o *HospitalsRow) GetHospitalNameOk() (*string, bool)`

GetHospitalNameOk returns a tuple with the HospitalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHospitalName

`func (o *HospitalsRow) SetHospitalName(v string)`

SetHospitalName sets HospitalName field to given value.

### HasHospitalName

`func (o *HospitalsRow) HasHospitalName() bool`

HasHospitalName returns a boolean if a field has been set.

### SetHospitalNameNil

`func (o *HospitalsRow) SetHospitalNameNil(b bool)`

 SetHospitalNameNil sets the value for HospitalName to be an explicit nil

### UnsetHospitalName
`func (o *HospitalsRow) UnsetHospitalName()`

UnsetHospitalName ensures that no value is present for HospitalName, not even an explicit nil
### GetCity

`func (o *HospitalsRow) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *HospitalsRow) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *HospitalsRow) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *HospitalsRow) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *HospitalsRow) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *HospitalsRow) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *HospitalsRow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HospitalsRow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HospitalsRow) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HospitalsRow) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *HospitalsRow) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *HospitalsRow) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetBedCount

`func (o *HospitalsRow) GetBedCount() int32`

GetBedCount returns the BedCount field if non-nil, zero value otherwise.

### GetBedCountOk

`func (o *HospitalsRow) GetBedCountOk() (*int32, bool)`

GetBedCountOk returns a tuple with the BedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBedCount

`func (o *HospitalsRow) SetBedCount(v int32)`

SetBedCount sets BedCount field to given value.

### HasBedCount

`func (o *HospitalsRow) HasBedCount() bool`

HasBedCount returns a boolean if a field has been set.

### SetBedCountNil

`func (o *HospitalsRow) SetBedCountNil(b bool)`

 SetBedCountNil sets the value for BedCount to be an explicit nil

### UnsetBedCount
`func (o *HospitalsRow) UnsetBedCount()`

UnsetBedCount ensures that no value is present for BedCount, not even an explicit nil
### GetHospitalType

`func (o *HospitalsRow) GetHospitalType() string`

GetHospitalType returns the HospitalType field if non-nil, zero value otherwise.

### GetHospitalTypeOk

`func (o *HospitalsRow) GetHospitalTypeOk() (*string, bool)`

GetHospitalTypeOk returns a tuple with the HospitalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHospitalType

`func (o *HospitalsRow) SetHospitalType(v string)`

SetHospitalType sets HospitalType field to given value.

### HasHospitalType

`func (o *HospitalsRow) HasHospitalType() bool`

HasHospitalType returns a boolean if a field has been set.

### SetHospitalTypeNil

`func (o *HospitalsRow) SetHospitalTypeNil(b bool)`

 SetHospitalTypeNil sets the value for HospitalType to be an explicit nil

### UnsetHospitalType
`func (o *HospitalsRow) UnsetHospitalType()`

UnsetHospitalType ensures that no value is present for HospitalType, not even an explicit nil
### GetOwnership

`func (o *HospitalsRow) GetOwnership() string`

GetOwnership returns the Ownership field if non-nil, zero value otherwise.

### GetOwnershipOk

`func (o *HospitalsRow) GetOwnershipOk() (*string, bool)`

GetOwnershipOk returns a tuple with the Ownership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnership

`func (o *HospitalsRow) SetOwnership(v string)`

SetOwnership sets Ownership field to given value.

### HasOwnership

`func (o *HospitalsRow) HasOwnership() bool`

HasOwnership returns a boolean if a field has been set.

### SetOwnershipNil

`func (o *HospitalsRow) SetOwnershipNil(b bool)`

 SetOwnershipNil sets the value for Ownership to be an explicit nil

### UnsetOwnership
`func (o *HospitalsRow) UnsetOwnership()`

UnsetOwnership ensures that no value is present for Ownership, not even an explicit nil
### GetIsTeaching

`func (o *HospitalsRow) GetIsTeaching() bool`

GetIsTeaching returns the IsTeaching field if non-nil, zero value otherwise.

### GetIsTeachingOk

`func (o *HospitalsRow) GetIsTeachingOk() (*bool, bool)`

GetIsTeachingOk returns a tuple with the IsTeaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTeaching

`func (o *HospitalsRow) SetIsTeaching(v bool)`

SetIsTeaching sets IsTeaching field to given value.

### HasIsTeaching

`func (o *HospitalsRow) HasIsTeaching() bool`

HasIsTeaching returns a boolean if a field has been set.

### SetIsTeachingNil

`func (o *HospitalsRow) SetIsTeachingNil(b bool)`

 SetIsTeachingNil sets the value for IsTeaching to be an explicit nil

### UnsetIsTeaching
`func (o *HospitalsRow) UnsetIsTeaching()`

UnsetIsTeaching ensures that no value is present for IsTeaching, not even an explicit nil
### GetEmergencyServices

`func (o *HospitalsRow) GetEmergencyServices() string`

GetEmergencyServices returns the EmergencyServices field if non-nil, zero value otherwise.

### GetEmergencyServicesOk

`func (o *HospitalsRow) GetEmergencyServicesOk() (*string, bool)`

GetEmergencyServicesOk returns a tuple with the EmergencyServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyServices

`func (o *HospitalsRow) SetEmergencyServices(v string)`

SetEmergencyServices sets EmergencyServices field to given value.

### HasEmergencyServices

`func (o *HospitalsRow) HasEmergencyServices() bool`

HasEmergencyServices returns a boolean if a field has been set.

### SetEmergencyServicesNil

`func (o *HospitalsRow) SetEmergencyServicesNil(b bool)`

 SetEmergencyServicesNil sets the value for EmergencyServices to be an explicit nil

### UnsetEmergencyServices
`func (o *HospitalsRow) UnsetEmergencyServices()`

UnsetEmergencyServices ensures that no value is present for EmergencyServices, not even an explicit nil
### GetLatitude

`func (o *HospitalsRow) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *HospitalsRow) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *HospitalsRow) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *HospitalsRow) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *HospitalsRow) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *HospitalsRow) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *HospitalsRow) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *HospitalsRow) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *HospitalsRow) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *HospitalsRow) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *HospitalsRow) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *HospitalsRow) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetGeocodeMethod

`func (o *HospitalsRow) GetGeocodeMethod() string`

GetGeocodeMethod returns the GeocodeMethod field if non-nil, zero value otherwise.

### GetGeocodeMethodOk

`func (o *HospitalsRow) GetGeocodeMethodOk() (*string, bool)`

GetGeocodeMethodOk returns a tuple with the GeocodeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocodeMethod

`func (o *HospitalsRow) SetGeocodeMethod(v string)`

SetGeocodeMethod sets GeocodeMethod field to given value.

### HasGeocodeMethod

`func (o *HospitalsRow) HasGeocodeMethod() bool`

HasGeocodeMethod returns a boolean if a field has been set.

### SetGeocodeMethodNil

`func (o *HospitalsRow) SetGeocodeMethodNil(b bool)`

 SetGeocodeMethodNil sets the value for GeocodeMethod to be an explicit nil

### UnsetGeocodeMethod
`func (o *HospitalsRow) UnsetGeocodeMethod()`

UnsetGeocodeMethod ensures that no value is present for GeocodeMethod, not even an explicit nil
### GetHrrNum

`func (o *HospitalsRow) GetHrrNum() int32`

GetHrrNum returns the HrrNum field if non-nil, zero value otherwise.

### GetHrrNumOk

`func (o *HospitalsRow) GetHrrNumOk() (*int32, bool)`

GetHrrNumOk returns a tuple with the HrrNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHrrNum

`func (o *HospitalsRow) SetHrrNum(v int32)`

SetHrrNum sets HrrNum field to given value.

### HasHrrNum

`func (o *HospitalsRow) HasHrrNum() bool`

HasHrrNum returns a boolean if a field has been set.

### SetHrrNumNil

`func (o *HospitalsRow) SetHrrNumNil(b bool)`

 SetHrrNumNil sets the value for HrrNum to be an explicit nil

### UnsetHrrNum
`func (o *HospitalsRow) UnsetHrrNum()`

UnsetHrrNum ensures that no value is present for HrrNum, not even an explicit nil
### GetHrrName

`func (o *HospitalsRow) GetHrrName() string`

GetHrrName returns the HrrName field if non-nil, zero value otherwise.

### GetHrrNameOk

`func (o *HospitalsRow) GetHrrNameOk() (*string, bool)`

GetHrrNameOk returns a tuple with the HrrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHrrName

`func (o *HospitalsRow) SetHrrName(v string)`

SetHrrName sets HrrName field to given value.

### HasHrrName

`func (o *HospitalsRow) HasHrrName() bool`

HasHrrName returns a boolean if a field has been set.

### SetHrrNameNil

`func (o *HospitalsRow) SetHrrNameNil(b bool)`

 SetHrrNameNil sets the value for HrrName to be an explicit nil

### UnsetHrrName
`func (o *HospitalsRow) UnsetHrrName()`

UnsetHrrName ensures that no value is present for HrrName, not even an explicit nil
### GetHsaNum

`func (o *HospitalsRow) GetHsaNum() int32`

GetHsaNum returns the HsaNum field if non-nil, zero value otherwise.

### GetHsaNumOk

`func (o *HospitalsRow) GetHsaNumOk() (*int32, bool)`

GetHsaNumOk returns a tuple with the HsaNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsaNum

`func (o *HospitalsRow) SetHsaNum(v int32)`

SetHsaNum sets HsaNum field to given value.

### HasHsaNum

`func (o *HospitalsRow) HasHsaNum() bool`

HasHsaNum returns a boolean if a field has been set.

### SetHsaNumNil

`func (o *HospitalsRow) SetHsaNumNil(b bool)`

 SetHsaNumNil sets the value for HsaNum to be an explicit nil

### UnsetHsaNum
`func (o *HospitalsRow) UnsetHsaNum()`

UnsetHsaNum ensures that no value is present for HsaNum, not even an explicit nil
### GetHsaName

`func (o *HospitalsRow) GetHsaName() string`

GetHsaName returns the HsaName field if non-nil, zero value otherwise.

### GetHsaNameOk

`func (o *HospitalsRow) GetHsaNameOk() (*string, bool)`

GetHsaNameOk returns a tuple with the HsaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsaName

`func (o *HospitalsRow) SetHsaName(v string)`

SetHsaName sets HsaName field to given value.

### HasHsaName

`func (o *HospitalsRow) HasHsaName() bool`

HasHsaName returns a boolean if a field has been set.

### SetHsaNameNil

`func (o *HospitalsRow) SetHsaNameNil(b bool)`

 SetHsaNameNil sets the value for HsaName to be an explicit nil

### UnsetHsaName
`func (o *HospitalsRow) UnsetHsaName()`

UnsetHsaName ensures that no value is present for HsaName, not even an explicit nil
### GetHrrMatch

`func (o *HospitalsRow) GetHrrMatch() string`

GetHrrMatch returns the HrrMatch field if non-nil, zero value otherwise.

### GetHrrMatchOk

`func (o *HospitalsRow) GetHrrMatchOk() (*string, bool)`

GetHrrMatchOk returns a tuple with the HrrMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHrrMatch

`func (o *HospitalsRow) SetHrrMatch(v string)`

SetHrrMatch sets HrrMatch field to given value.

### HasHrrMatch

`func (o *HospitalsRow) HasHrrMatch() bool`

HasHrrMatch returns a boolean if a field has been set.

### SetHrrMatchNil

`func (o *HospitalsRow) SetHrrMatchNil(b bool)`

 SetHrrMatchNil sets the value for HrrMatch to be an explicit nil

### UnsetHrrMatch
`func (o *HospitalsRow) UnsetHrrMatch()`

UnsetHrrMatch ensures that no value is present for HrrMatch, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


