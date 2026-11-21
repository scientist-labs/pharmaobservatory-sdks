# ConditionSiteCoverageRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**StartYear** | Pointer to **NullableInt32** | Year the trials started. Null for a trial with no start date. | [optional] 
**NTrials** | Pointer to **NullableInt32** | Trials of the condition that started that year. | [optional] 
**NTrialsLocated** | Pointer to **NullableInt32** | Of those trials, how many list at least one facility with a coordinate. | [optional] 

## Methods

### NewConditionSiteCoverageRow

`func NewConditionSiteCoverageRow() *ConditionSiteCoverageRow`

NewConditionSiteCoverageRow instantiates a new ConditionSiteCoverageRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionSiteCoverageRowWithDefaults

`func NewConditionSiteCoverageRowWithDefaults() *ConditionSiteCoverageRow`

NewConditionSiteCoverageRowWithDefaults instantiates a new ConditionSiteCoverageRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ConditionSiteCoverageRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ConditionSiteCoverageRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ConditionSiteCoverageRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ConditionSiteCoverageRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ConditionSiteCoverageRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ConditionSiteCoverageRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetStartYear

`func (o *ConditionSiteCoverageRow) GetStartYear() int32`

GetStartYear returns the StartYear field if non-nil, zero value otherwise.

### GetStartYearOk

`func (o *ConditionSiteCoverageRow) GetStartYearOk() (*int32, bool)`

GetStartYearOk returns a tuple with the StartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartYear

`func (o *ConditionSiteCoverageRow) SetStartYear(v int32)`

SetStartYear sets StartYear field to given value.

### HasStartYear

`func (o *ConditionSiteCoverageRow) HasStartYear() bool`

HasStartYear returns a boolean if a field has been set.

### SetStartYearNil

`func (o *ConditionSiteCoverageRow) SetStartYearNil(b bool)`

 SetStartYearNil sets the value for StartYear to be an explicit nil

### UnsetStartYear
`func (o *ConditionSiteCoverageRow) UnsetStartYear()`

UnsetStartYear ensures that no value is present for StartYear, not even an explicit nil
### GetNTrials

`func (o *ConditionSiteCoverageRow) GetNTrials() int32`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *ConditionSiteCoverageRow) GetNTrialsOk() (*int32, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *ConditionSiteCoverageRow) SetNTrials(v int32)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *ConditionSiteCoverageRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *ConditionSiteCoverageRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *ConditionSiteCoverageRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNTrialsLocated

`func (o *ConditionSiteCoverageRow) GetNTrialsLocated() int32`

GetNTrialsLocated returns the NTrialsLocated field if non-nil, zero value otherwise.

### GetNTrialsLocatedOk

`func (o *ConditionSiteCoverageRow) GetNTrialsLocatedOk() (*int32, bool)`

GetNTrialsLocatedOk returns a tuple with the NTrialsLocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsLocated

`func (o *ConditionSiteCoverageRow) SetNTrialsLocated(v int32)`

SetNTrialsLocated sets NTrialsLocated field to given value.

### HasNTrialsLocated

`func (o *ConditionSiteCoverageRow) HasNTrialsLocated() bool`

HasNTrialsLocated returns a boolean if a field has been set.

### SetNTrialsLocatedNil

`func (o *ConditionSiteCoverageRow) SetNTrialsLocatedNil(b bool)`

 SetNTrialsLocatedNil sets the value for NTrialsLocated to be an explicit nil

### UnsetNTrialsLocated
`func (o *ConditionSiteCoverageRow) UnsetNTrialsLocated()`

UnsetNTrialsLocated ensures that no value is present for NTrialsLocated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


