# ConditionBurdenRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s MeSH heading. | [optional] 
**CauseName** | Pointer to **NullableString** | The Global Burden of Disease cause matched to the condition. | [optional] 
**DalysMillions** | Pointer to **NullableFloat64** | Global disability-adjusted life years for the matched cause, in millions. | [optional] 
**DeathsMillions** | Pointer to **NullableFloat64** | Global deaths for the matched cause, in millions. | [optional] 
**Druggability** | Pointer to **NullableInt32** | Druggability class of the matched cause: 1 highly druggable, 2 druggable, 3 challenging, 4 very challenging. | [optional] 
**NTrials** | Pointer to **NullableInt64** | Trials for the condition, counted under each condition they are tagged with and each condition above it in the MeSH tree. | [optional] 
**TrialsPerMillionDalys** | Pointer to **NullableFloat32** | Trials per million DALYs: n_trials divided by dalys_millions. | [optional] 

## Methods

### NewConditionBurdenRow

`func NewConditionBurdenRow() *ConditionBurdenRow`

NewConditionBurdenRow instantiates a new ConditionBurdenRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionBurdenRowWithDefaults

`func NewConditionBurdenRowWithDefaults() *ConditionBurdenRow`

NewConditionBurdenRowWithDefaults instantiates a new ConditionBurdenRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ConditionBurdenRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ConditionBurdenRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ConditionBurdenRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ConditionBurdenRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ConditionBurdenRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ConditionBurdenRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *ConditionBurdenRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *ConditionBurdenRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *ConditionBurdenRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *ConditionBurdenRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *ConditionBurdenRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *ConditionBurdenRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetCauseName

`func (o *ConditionBurdenRow) GetCauseName() string`

GetCauseName returns the CauseName field if non-nil, zero value otherwise.

### GetCauseNameOk

`func (o *ConditionBurdenRow) GetCauseNameOk() (*string, bool)`

GetCauseNameOk returns a tuple with the CauseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseName

`func (o *ConditionBurdenRow) SetCauseName(v string)`

SetCauseName sets CauseName field to given value.

### HasCauseName

`func (o *ConditionBurdenRow) HasCauseName() bool`

HasCauseName returns a boolean if a field has been set.

### SetCauseNameNil

`func (o *ConditionBurdenRow) SetCauseNameNil(b bool)`

 SetCauseNameNil sets the value for CauseName to be an explicit nil

### UnsetCauseName
`func (o *ConditionBurdenRow) UnsetCauseName()`

UnsetCauseName ensures that no value is present for CauseName, not even an explicit nil
### GetDalysMillions

`func (o *ConditionBurdenRow) GetDalysMillions() float64`

GetDalysMillions returns the DalysMillions field if non-nil, zero value otherwise.

### GetDalysMillionsOk

`func (o *ConditionBurdenRow) GetDalysMillionsOk() (*float64, bool)`

GetDalysMillionsOk returns a tuple with the DalysMillions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDalysMillions

`func (o *ConditionBurdenRow) SetDalysMillions(v float64)`

SetDalysMillions sets DalysMillions field to given value.

### HasDalysMillions

`func (o *ConditionBurdenRow) HasDalysMillions() bool`

HasDalysMillions returns a boolean if a field has been set.

### SetDalysMillionsNil

`func (o *ConditionBurdenRow) SetDalysMillionsNil(b bool)`

 SetDalysMillionsNil sets the value for DalysMillions to be an explicit nil

### UnsetDalysMillions
`func (o *ConditionBurdenRow) UnsetDalysMillions()`

UnsetDalysMillions ensures that no value is present for DalysMillions, not even an explicit nil
### GetDeathsMillions

`func (o *ConditionBurdenRow) GetDeathsMillions() float64`

GetDeathsMillions returns the DeathsMillions field if non-nil, zero value otherwise.

### GetDeathsMillionsOk

`func (o *ConditionBurdenRow) GetDeathsMillionsOk() (*float64, bool)`

GetDeathsMillionsOk returns a tuple with the DeathsMillions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeathsMillions

`func (o *ConditionBurdenRow) SetDeathsMillions(v float64)`

SetDeathsMillions sets DeathsMillions field to given value.

### HasDeathsMillions

`func (o *ConditionBurdenRow) HasDeathsMillions() bool`

HasDeathsMillions returns a boolean if a field has been set.

### SetDeathsMillionsNil

`func (o *ConditionBurdenRow) SetDeathsMillionsNil(b bool)`

 SetDeathsMillionsNil sets the value for DeathsMillions to be an explicit nil

### UnsetDeathsMillions
`func (o *ConditionBurdenRow) UnsetDeathsMillions()`

UnsetDeathsMillions ensures that no value is present for DeathsMillions, not even an explicit nil
### GetDruggability

`func (o *ConditionBurdenRow) GetDruggability() int32`

GetDruggability returns the Druggability field if non-nil, zero value otherwise.

### GetDruggabilityOk

`func (o *ConditionBurdenRow) GetDruggabilityOk() (*int32, bool)`

GetDruggabilityOk returns a tuple with the Druggability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDruggability

`func (o *ConditionBurdenRow) SetDruggability(v int32)`

SetDruggability sets Druggability field to given value.

### HasDruggability

`func (o *ConditionBurdenRow) HasDruggability() bool`

HasDruggability returns a boolean if a field has been set.

### SetDruggabilityNil

`func (o *ConditionBurdenRow) SetDruggabilityNil(b bool)`

 SetDruggabilityNil sets the value for Druggability to be an explicit nil

### UnsetDruggability
`func (o *ConditionBurdenRow) UnsetDruggability()`

UnsetDruggability ensures that no value is present for Druggability, not even an explicit nil
### GetNTrials

`func (o *ConditionBurdenRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *ConditionBurdenRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *ConditionBurdenRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *ConditionBurdenRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *ConditionBurdenRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *ConditionBurdenRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetTrialsPerMillionDalys

`func (o *ConditionBurdenRow) GetTrialsPerMillionDalys() float32`

GetTrialsPerMillionDalys returns the TrialsPerMillionDalys field if non-nil, zero value otherwise.

### GetTrialsPerMillionDalysOk

`func (o *ConditionBurdenRow) GetTrialsPerMillionDalysOk() (*float32, bool)`

GetTrialsPerMillionDalysOk returns a tuple with the TrialsPerMillionDalys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialsPerMillionDalys

`func (o *ConditionBurdenRow) SetTrialsPerMillionDalys(v float32)`

SetTrialsPerMillionDalys sets TrialsPerMillionDalys field to given value.

### HasTrialsPerMillionDalys

`func (o *ConditionBurdenRow) HasTrialsPerMillionDalys() bool`

HasTrialsPerMillionDalys returns a boolean if a field has been set.

### SetTrialsPerMillionDalysNil

`func (o *ConditionBurdenRow) SetTrialsPerMillionDalysNil(b bool)`

 SetTrialsPerMillionDalysNil sets the value for TrialsPerMillionDalys to be an explicit nil

### UnsetTrialsPerMillionDalys
`func (o *ConditionBurdenRow) UnsetTrialsPerMillionDalys()`

UnsetTrialsPerMillionDalys ensures that no value is present for TrialsPerMillionDalys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


