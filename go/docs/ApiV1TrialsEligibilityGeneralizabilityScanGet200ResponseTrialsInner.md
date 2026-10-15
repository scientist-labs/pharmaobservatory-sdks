# ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NctId** | **string** |  | 
**BriefTitle** | **string** |  | 
**Gender** | **string** |  | 
**MinimumAge** | **NullableString** |  | 
**MaximumAge** | **NullableString** |  | 
**HealthyVolunteers** | **NullableBool** |  | 
**FacilityCountryCount** | **NullableInt32** |  | 
**FacilityCountries** | **[]string** |  | 

## Methods

### NewApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner

`func NewApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner(nctId string, briefTitle string, gender string, minimumAge NullableString, maximumAge NullableString, healthyVolunteers NullableBool, facilityCountryCount NullableInt32, facilityCountries []string, ) *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner`

NewApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner instantiates a new ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInnerWithDefaults

`func NewApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInnerWithDefaults() *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner`

NewApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInnerWithDefaults instantiates a new ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNctId

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) GetNctId() string`

GetNctId returns the NctId field if non-nil, zero value otherwise.

### GetNctIdOk

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) GetNctIdOk() (*string, bool)`

GetNctIdOk returns a tuple with the NctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNctId

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) SetNctId(v string)`

SetNctId sets NctId field to given value.


### GetBriefTitle

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) GetBriefTitle() string`

GetBriefTitle returns the BriefTitle field if non-nil, zero value otherwise.

### GetBriefTitleOk

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) GetBriefTitleOk() (*string, bool)`

GetBriefTitleOk returns a tuple with the BriefTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBriefTitle

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) SetBriefTitle(v string)`

SetBriefTitle sets BriefTitle field to given value.


### GetGender

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) SetGender(v string)`

SetGender sets Gender field to given value.


### GetMinimumAge

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) GetMinimumAge() string`

GetMinimumAge returns the MinimumAge field if non-nil, zero value otherwise.

### GetMinimumAgeOk

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) GetMinimumAgeOk() (*string, bool)`

GetMinimumAgeOk returns a tuple with the MinimumAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAge

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) SetMinimumAge(v string)`

SetMinimumAge sets MinimumAge field to given value.


### SetMinimumAgeNil

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) SetMinimumAgeNil(b bool)`

 SetMinimumAgeNil sets the value for MinimumAge to be an explicit nil

### UnsetMinimumAge
`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) UnsetMinimumAge()`

UnsetMinimumAge ensures that no value is present for MinimumAge, not even an explicit nil
### GetMaximumAge

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) GetMaximumAge() string`

GetMaximumAge returns the MaximumAge field if non-nil, zero value otherwise.

### GetMaximumAgeOk

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) GetMaximumAgeOk() (*string, bool)`

GetMaximumAgeOk returns a tuple with the MaximumAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAge

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) SetMaximumAge(v string)`

SetMaximumAge sets MaximumAge field to given value.


### SetMaximumAgeNil

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) SetMaximumAgeNil(b bool)`

 SetMaximumAgeNil sets the value for MaximumAge to be an explicit nil

### UnsetMaximumAge
`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) UnsetMaximumAge()`

UnsetMaximumAge ensures that no value is present for MaximumAge, not even an explicit nil
### GetHealthyVolunteers

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) GetHealthyVolunteers() bool`

GetHealthyVolunteers returns the HealthyVolunteers field if non-nil, zero value otherwise.

### GetHealthyVolunteersOk

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) GetHealthyVolunteersOk() (*bool, bool)`

GetHealthyVolunteersOk returns a tuple with the HealthyVolunteers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyVolunteers

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) SetHealthyVolunteers(v bool)`

SetHealthyVolunteers sets HealthyVolunteers field to given value.


### SetHealthyVolunteersNil

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) SetHealthyVolunteersNil(b bool)`

 SetHealthyVolunteersNil sets the value for HealthyVolunteers to be an explicit nil

### UnsetHealthyVolunteers
`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) UnsetHealthyVolunteers()`

UnsetHealthyVolunteers ensures that no value is present for HealthyVolunteers, not even an explicit nil
### GetFacilityCountryCount

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) GetFacilityCountryCount() int32`

GetFacilityCountryCount returns the FacilityCountryCount field if non-nil, zero value otherwise.

### GetFacilityCountryCountOk

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) GetFacilityCountryCountOk() (*int32, bool)`

GetFacilityCountryCountOk returns a tuple with the FacilityCountryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCountryCount

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) SetFacilityCountryCount(v int32)`

SetFacilityCountryCount sets FacilityCountryCount field to given value.


### SetFacilityCountryCountNil

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) SetFacilityCountryCountNil(b bool)`

 SetFacilityCountryCountNil sets the value for FacilityCountryCount to be an explicit nil

### UnsetFacilityCountryCount
`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) UnsetFacilityCountryCount()`

UnsetFacilityCountryCount ensures that no value is present for FacilityCountryCount, not even an explicit nil
### GetFacilityCountries

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) GetFacilityCountries() []string`

GetFacilityCountries returns the FacilityCountries field if non-nil, zero value otherwise.

### GetFacilityCountriesOk

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) GetFacilityCountriesOk() (*[]string, bool)`

GetFacilityCountriesOk returns a tuple with the FacilityCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCountries

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) SetFacilityCountries(v []string)`

SetFacilityCountries sets FacilityCountries field to given value.


### SetFacilityCountriesNil

`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) SetFacilityCountriesNil(b bool)`

 SetFacilityCountriesNil sets the value for FacilityCountries to be an explicit nil

### UnsetFacilityCountries
`func (o *ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner) UnsetFacilityCountries()`

UnsetFacilityCountries ensures that no value is present for FacilityCountries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


