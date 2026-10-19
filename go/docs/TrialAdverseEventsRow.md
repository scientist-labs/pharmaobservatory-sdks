# TrialAdverseEventsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**EventType** | Pointer to **NullableString** | serious or other. | [optional] 
**OrganSystem** | Pointer to **NullableString** | The organ system class, as filed ((unspecified) where the filing gives none). | [optional] 
**AdverseEventTerm** | Pointer to **NullableString** | The adverse event term, as filed. | [optional] 
**SubjectsAffected** | Pointer to **NullableInt64** | Participants affected, summed over the trial&#39;s groups. | [optional] 
**SubjectsAtRisk** | Pointer to **NullableInt64** | Participants at risk, summed over the trial&#39;s groups. | [optional] 

## Methods

### NewTrialAdverseEventsRow

`func NewTrialAdverseEventsRow() *TrialAdverseEventsRow`

NewTrialAdverseEventsRow instantiates a new TrialAdverseEventsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialAdverseEventsRowWithDefaults

`func NewTrialAdverseEventsRowWithDefaults() *TrialAdverseEventsRow`

NewTrialAdverseEventsRowWithDefaults instantiates a new TrialAdverseEventsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *TrialAdverseEventsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TrialAdverseEventsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TrialAdverseEventsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TrialAdverseEventsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TrialAdverseEventsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TrialAdverseEventsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetEventType

`func (o *TrialAdverseEventsRow) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TrialAdverseEventsRow) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TrialAdverseEventsRow) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *TrialAdverseEventsRow) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventTypeNil

`func (o *TrialAdverseEventsRow) SetEventTypeNil(b bool)`

 SetEventTypeNil sets the value for EventType to be an explicit nil

### UnsetEventType
`func (o *TrialAdverseEventsRow) UnsetEventType()`

UnsetEventType ensures that no value is present for EventType, not even an explicit nil
### GetOrganSystem

`func (o *TrialAdverseEventsRow) GetOrganSystem() string`

GetOrganSystem returns the OrganSystem field if non-nil, zero value otherwise.

### GetOrganSystemOk

`func (o *TrialAdverseEventsRow) GetOrganSystemOk() (*string, bool)`

GetOrganSystemOk returns a tuple with the OrganSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganSystem

`func (o *TrialAdverseEventsRow) SetOrganSystem(v string)`

SetOrganSystem sets OrganSystem field to given value.

### HasOrganSystem

`func (o *TrialAdverseEventsRow) HasOrganSystem() bool`

HasOrganSystem returns a boolean if a field has been set.

### SetOrganSystemNil

`func (o *TrialAdverseEventsRow) SetOrganSystemNil(b bool)`

 SetOrganSystemNil sets the value for OrganSystem to be an explicit nil

### UnsetOrganSystem
`func (o *TrialAdverseEventsRow) UnsetOrganSystem()`

UnsetOrganSystem ensures that no value is present for OrganSystem, not even an explicit nil
### GetAdverseEventTerm

`func (o *TrialAdverseEventsRow) GetAdverseEventTerm() string`

GetAdverseEventTerm returns the AdverseEventTerm field if non-nil, zero value otherwise.

### GetAdverseEventTermOk

`func (o *TrialAdverseEventsRow) GetAdverseEventTermOk() (*string, bool)`

GetAdverseEventTermOk returns a tuple with the AdverseEventTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdverseEventTerm

`func (o *TrialAdverseEventsRow) SetAdverseEventTerm(v string)`

SetAdverseEventTerm sets AdverseEventTerm field to given value.

### HasAdverseEventTerm

`func (o *TrialAdverseEventsRow) HasAdverseEventTerm() bool`

HasAdverseEventTerm returns a boolean if a field has been set.

### SetAdverseEventTermNil

`func (o *TrialAdverseEventsRow) SetAdverseEventTermNil(b bool)`

 SetAdverseEventTermNil sets the value for AdverseEventTerm to be an explicit nil

### UnsetAdverseEventTerm
`func (o *TrialAdverseEventsRow) UnsetAdverseEventTerm()`

UnsetAdverseEventTerm ensures that no value is present for AdverseEventTerm, not even an explicit nil
### GetSubjectsAffected

`func (o *TrialAdverseEventsRow) GetSubjectsAffected() int64`

GetSubjectsAffected returns the SubjectsAffected field if non-nil, zero value otherwise.

### GetSubjectsAffectedOk

`func (o *TrialAdverseEventsRow) GetSubjectsAffectedOk() (*int64, bool)`

GetSubjectsAffectedOk returns a tuple with the SubjectsAffected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectsAffected

`func (o *TrialAdverseEventsRow) SetSubjectsAffected(v int64)`

SetSubjectsAffected sets SubjectsAffected field to given value.

### HasSubjectsAffected

`func (o *TrialAdverseEventsRow) HasSubjectsAffected() bool`

HasSubjectsAffected returns a boolean if a field has been set.

### SetSubjectsAffectedNil

`func (o *TrialAdverseEventsRow) SetSubjectsAffectedNil(b bool)`

 SetSubjectsAffectedNil sets the value for SubjectsAffected to be an explicit nil

### UnsetSubjectsAffected
`func (o *TrialAdverseEventsRow) UnsetSubjectsAffected()`

UnsetSubjectsAffected ensures that no value is present for SubjectsAffected, not even an explicit nil
### GetSubjectsAtRisk

`func (o *TrialAdverseEventsRow) GetSubjectsAtRisk() int64`

GetSubjectsAtRisk returns the SubjectsAtRisk field if non-nil, zero value otherwise.

### GetSubjectsAtRiskOk

`func (o *TrialAdverseEventsRow) GetSubjectsAtRiskOk() (*int64, bool)`

GetSubjectsAtRiskOk returns a tuple with the SubjectsAtRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectsAtRisk

`func (o *TrialAdverseEventsRow) SetSubjectsAtRisk(v int64)`

SetSubjectsAtRisk sets SubjectsAtRisk field to given value.

### HasSubjectsAtRisk

`func (o *TrialAdverseEventsRow) HasSubjectsAtRisk() bool`

HasSubjectsAtRisk returns a boolean if a field has been set.

### SetSubjectsAtRiskNil

`func (o *TrialAdverseEventsRow) SetSubjectsAtRiskNil(b bool)`

 SetSubjectsAtRiskNil sets the value for SubjectsAtRisk to be an explicit nil

### UnsetSubjectsAtRisk
`func (o *TrialAdverseEventsRow) UnsetSubjectsAtRisk()`

UnsetSubjectsAtRisk ensures that no value is present for SubjectsAtRisk, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


