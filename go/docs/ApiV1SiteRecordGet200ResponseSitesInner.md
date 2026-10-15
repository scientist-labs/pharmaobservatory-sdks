# ApiV1SiteRecordGet200ResponseSitesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteKey** | **string** |  | 
**FacilityName** | **string** |  | 
**City** | **string** |  | 
**State** | **NullableString** |  | 
**Country** | **string** |  | 
**NTrials** | **int32** |  | 
**NActive** | **int32** |  | 
**ActiveYears** | [**ApiV1SiteRecordGet200ResponseSitesInnerActiveYears**](ApiV1SiteRecordGet200ResponseSitesInnerActiveYears.md) |  | 
**Latitude** | Pointer to **string** |  | [optional] 
**Longitude** | Pointer to **string** |  | [optional] 
**Trials** | Pointer to [**[]ApiV1SiteRecordGet200ResponseSitesInnerTrialsInner**](ApiV1SiteRecordGet200ResponseSitesInnerTrialsInner.md) |  | [optional] 
**Indications** | Pointer to [**[]ApiV1SiteRecordGet200ResponseSitesInnerIndicationsInner**](ApiV1SiteRecordGet200ResponseSitesInnerIndicationsInner.md) |  | [optional] 
**Drugs** | Pointer to [**[]ApiV1SiteRecordGet200ResponseSitesInnerDrugsInner**](ApiV1SiteRecordGet200ResponseSitesInnerDrugsInner.md) |  | [optional] 
**EstimatedEnrollment** | Pointer to [**NullableApiV1SiteRecordGet200ResponseSitesInnerEstimatedEnrollment**](ApiV1SiteRecordGet200ResponseSitesInnerEstimatedEnrollment.md) |  | [optional] 
**Investigators** | Pointer to [**[]ApiV1SiteRecordGet200ResponseSitesInnerInvestigatorsInner**](ApiV1SiteRecordGet200ResponseSitesInnerInvestigatorsInner.md) |  | [optional] 

## Methods

### NewApiV1SiteRecordGet200ResponseSitesInner

`func NewApiV1SiteRecordGet200ResponseSitesInner(siteKey string, facilityName string, city string, state NullableString, country string, nTrials int32, nActive int32, activeYears ApiV1SiteRecordGet200ResponseSitesInnerActiveYears, ) *ApiV1SiteRecordGet200ResponseSitesInner`

NewApiV1SiteRecordGet200ResponseSitesInner instantiates a new ApiV1SiteRecordGet200ResponseSitesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1SiteRecordGet200ResponseSitesInnerWithDefaults

`func NewApiV1SiteRecordGet200ResponseSitesInnerWithDefaults() *ApiV1SiteRecordGet200ResponseSitesInner`

NewApiV1SiteRecordGet200ResponseSitesInnerWithDefaults instantiates a new ApiV1SiteRecordGet200ResponseSitesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteKey

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetSiteKey() string`

GetSiteKey returns the SiteKey field if non-nil, zero value otherwise.

### GetSiteKeyOk

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetSiteKeyOk() (*string, bool)`

GetSiteKeyOk returns a tuple with the SiteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteKey

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) SetSiteKey(v string)`

SetSiteKey sets SiteKey field to given value.


### GetFacilityName

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetFacilityName() string`

GetFacilityName returns the FacilityName field if non-nil, zero value otherwise.

### GetFacilityNameOk

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetFacilityNameOk() (*string, bool)`

GetFacilityNameOk returns a tuple with the FacilityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityName

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) SetFacilityName(v string)`

SetFacilityName sets FacilityName field to given value.


### GetCity

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ApiV1SiteRecordGet200ResponseSitesInner) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCountry

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetNTrials

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetNTrials() int32`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetNTrialsOk() (*int32, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) SetNTrials(v int32)`

SetNTrials sets NTrials field to given value.


### GetNActive

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetNActive() int32`

GetNActive returns the NActive field if non-nil, zero value otherwise.

### GetNActiveOk

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetNActiveOk() (*int32, bool)`

GetNActiveOk returns a tuple with the NActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNActive

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) SetNActive(v int32)`

SetNActive sets NActive field to given value.


### GetActiveYears

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetActiveYears() ApiV1SiteRecordGet200ResponseSitesInnerActiveYears`

GetActiveYears returns the ActiveYears field if non-nil, zero value otherwise.

### GetActiveYearsOk

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetActiveYearsOk() (*ApiV1SiteRecordGet200ResponseSitesInnerActiveYears, bool)`

GetActiveYearsOk returns a tuple with the ActiveYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveYears

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) SetActiveYears(v ApiV1SiteRecordGet200ResponseSitesInnerActiveYears)`

SetActiveYears sets ActiveYears field to given value.


### GetLatitude

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetTrials

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetTrials() []ApiV1SiteRecordGet200ResponseSitesInnerTrialsInner`

GetTrials returns the Trials field if non-nil, zero value otherwise.

### GetTrialsOk

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetTrialsOk() (*[]ApiV1SiteRecordGet200ResponseSitesInnerTrialsInner, bool)`

GetTrialsOk returns a tuple with the Trials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrials

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) SetTrials(v []ApiV1SiteRecordGet200ResponseSitesInnerTrialsInner)`

SetTrials sets Trials field to given value.

### HasTrials

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) HasTrials() bool`

HasTrials returns a boolean if a field has been set.

### GetIndications

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetIndications() []ApiV1SiteRecordGet200ResponseSitesInnerIndicationsInner`

GetIndications returns the Indications field if non-nil, zero value otherwise.

### GetIndicationsOk

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetIndicationsOk() (*[]ApiV1SiteRecordGet200ResponseSitesInnerIndicationsInner, bool)`

GetIndicationsOk returns a tuple with the Indications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndications

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) SetIndications(v []ApiV1SiteRecordGet200ResponseSitesInnerIndicationsInner)`

SetIndications sets Indications field to given value.

### HasIndications

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) HasIndications() bool`

HasIndications returns a boolean if a field has been set.

### GetDrugs

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetDrugs() []ApiV1SiteRecordGet200ResponseSitesInnerDrugsInner`

GetDrugs returns the Drugs field if non-nil, zero value otherwise.

### GetDrugsOk

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetDrugsOk() (*[]ApiV1SiteRecordGet200ResponseSitesInnerDrugsInner, bool)`

GetDrugsOk returns a tuple with the Drugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugs

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) SetDrugs(v []ApiV1SiteRecordGet200ResponseSitesInnerDrugsInner)`

SetDrugs sets Drugs field to given value.

### HasDrugs

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) HasDrugs() bool`

HasDrugs returns a boolean if a field has been set.

### GetEstimatedEnrollment

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetEstimatedEnrollment() ApiV1SiteRecordGet200ResponseSitesInnerEstimatedEnrollment`

GetEstimatedEnrollment returns the EstimatedEnrollment field if non-nil, zero value otherwise.

### GetEstimatedEnrollmentOk

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetEstimatedEnrollmentOk() (*ApiV1SiteRecordGet200ResponseSitesInnerEstimatedEnrollment, bool)`

GetEstimatedEnrollmentOk returns a tuple with the EstimatedEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedEnrollment

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) SetEstimatedEnrollment(v ApiV1SiteRecordGet200ResponseSitesInnerEstimatedEnrollment)`

SetEstimatedEnrollment sets EstimatedEnrollment field to given value.

### HasEstimatedEnrollment

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) HasEstimatedEnrollment() bool`

HasEstimatedEnrollment returns a boolean if a field has been set.

### SetEstimatedEnrollmentNil

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) SetEstimatedEnrollmentNil(b bool)`

 SetEstimatedEnrollmentNil sets the value for EstimatedEnrollment to be an explicit nil

### UnsetEstimatedEnrollment
`func (o *ApiV1SiteRecordGet200ResponseSitesInner) UnsetEstimatedEnrollment()`

UnsetEstimatedEnrollment ensures that no value is present for EstimatedEnrollment, not even an explicit nil
### GetInvestigators

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetInvestigators() []ApiV1SiteRecordGet200ResponseSitesInnerInvestigatorsInner`

GetInvestigators returns the Investigators field if non-nil, zero value otherwise.

### GetInvestigatorsOk

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) GetInvestigatorsOk() (*[]ApiV1SiteRecordGet200ResponseSitesInnerInvestigatorsInner, bool)`

GetInvestigatorsOk returns a tuple with the Investigators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigators

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) SetInvestigators(v []ApiV1SiteRecordGet200ResponseSitesInnerInvestigatorsInner)`

SetInvestigators sets Investigators field to given value.

### HasInvestigators

`func (o *ApiV1SiteRecordGet200ResponseSitesInner) HasInvestigators() bool`

HasInvestigators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


