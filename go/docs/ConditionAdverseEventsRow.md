# ConditionAdverseEventsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**AdverseEventTerm** | Pointer to **NullableString** | The adverse event term as the trials report it. | [optional] 
**SubjectsAffected** | Pointer to **NullableFloat32** | Subjects affected, summed across the condition&#39;s trials. | [optional] 
**NTrials** | Pointer to **NullableInt64** | Trials for this condition that report the term. | [optional] 

## Methods

### NewConditionAdverseEventsRow

`func NewConditionAdverseEventsRow() *ConditionAdverseEventsRow`

NewConditionAdverseEventsRow instantiates a new ConditionAdverseEventsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionAdverseEventsRowWithDefaults

`func NewConditionAdverseEventsRowWithDefaults() *ConditionAdverseEventsRow`

NewConditionAdverseEventsRowWithDefaults instantiates a new ConditionAdverseEventsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ConditionAdverseEventsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ConditionAdverseEventsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ConditionAdverseEventsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ConditionAdverseEventsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ConditionAdverseEventsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ConditionAdverseEventsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetAdverseEventTerm

`func (o *ConditionAdverseEventsRow) GetAdverseEventTerm() string`

GetAdverseEventTerm returns the AdverseEventTerm field if non-nil, zero value otherwise.

### GetAdverseEventTermOk

`func (o *ConditionAdverseEventsRow) GetAdverseEventTermOk() (*string, bool)`

GetAdverseEventTermOk returns a tuple with the AdverseEventTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdverseEventTerm

`func (o *ConditionAdverseEventsRow) SetAdverseEventTerm(v string)`

SetAdverseEventTerm sets AdverseEventTerm field to given value.

### HasAdverseEventTerm

`func (o *ConditionAdverseEventsRow) HasAdverseEventTerm() bool`

HasAdverseEventTerm returns a boolean if a field has been set.

### SetAdverseEventTermNil

`func (o *ConditionAdverseEventsRow) SetAdverseEventTermNil(b bool)`

 SetAdverseEventTermNil sets the value for AdverseEventTerm to be an explicit nil

### UnsetAdverseEventTerm
`func (o *ConditionAdverseEventsRow) UnsetAdverseEventTerm()`

UnsetAdverseEventTerm ensures that no value is present for AdverseEventTerm, not even an explicit nil
### GetSubjectsAffected

`func (o *ConditionAdverseEventsRow) GetSubjectsAffected() float32`

GetSubjectsAffected returns the SubjectsAffected field if non-nil, zero value otherwise.

### GetSubjectsAffectedOk

`func (o *ConditionAdverseEventsRow) GetSubjectsAffectedOk() (*float32, bool)`

GetSubjectsAffectedOk returns a tuple with the SubjectsAffected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectsAffected

`func (o *ConditionAdverseEventsRow) SetSubjectsAffected(v float32)`

SetSubjectsAffected sets SubjectsAffected field to given value.

### HasSubjectsAffected

`func (o *ConditionAdverseEventsRow) HasSubjectsAffected() bool`

HasSubjectsAffected returns a boolean if a field has been set.

### SetSubjectsAffectedNil

`func (o *ConditionAdverseEventsRow) SetSubjectsAffectedNil(b bool)`

 SetSubjectsAffectedNil sets the value for SubjectsAffected to be an explicit nil

### UnsetSubjectsAffected
`func (o *ConditionAdverseEventsRow) UnsetSubjectsAffected()`

UnsetSubjectsAffected ensures that no value is present for SubjectsAffected, not even an explicit nil
### GetNTrials

`func (o *ConditionAdverseEventsRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *ConditionAdverseEventsRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *ConditionAdverseEventsRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *ConditionAdverseEventsRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *ConditionAdverseEventsRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *ConditionAdverseEventsRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


