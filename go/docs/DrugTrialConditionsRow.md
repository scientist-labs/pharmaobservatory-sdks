# DrugTrialConditionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**ConditionId** | Pointer to **NullableString** | The condition key (a MeSH, Medical Subject Headings, descriptor id). | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s name. | [optional] 
**Phase** | Pointer to **NullableString** | The registry&#39;s phase code of the trial. | [optional] 
**PhaseRank** | Pointer to **NullableInt32** | The trial&#39;s phase as a number for max and sort (4 is Phase 4, 0 is not reported). | [optional] 
**OverallStatusSummary** | Pointer to **NullableString** | The trial&#39;s overall status grouped as Completed, Ongoing, Terminated or Other. | [optional] 

## Methods

### NewDrugTrialConditionsRow

`func NewDrugTrialConditionsRow() *DrugTrialConditionsRow`

NewDrugTrialConditionsRow instantiates a new DrugTrialConditionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugTrialConditionsRowWithDefaults

`func NewDrugTrialConditionsRowWithDefaults() *DrugTrialConditionsRow`

NewDrugTrialConditionsRowWithDefaults instantiates a new DrugTrialConditionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugTrialConditionsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugTrialConditionsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugTrialConditionsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugTrialConditionsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugTrialConditionsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugTrialConditionsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetTrialId

`func (o *DrugTrialConditionsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *DrugTrialConditionsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *DrugTrialConditionsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *DrugTrialConditionsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *DrugTrialConditionsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *DrugTrialConditionsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetConditionId

`func (o *DrugTrialConditionsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *DrugTrialConditionsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *DrugTrialConditionsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *DrugTrialConditionsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *DrugTrialConditionsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *DrugTrialConditionsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *DrugTrialConditionsRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *DrugTrialConditionsRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *DrugTrialConditionsRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *DrugTrialConditionsRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *DrugTrialConditionsRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *DrugTrialConditionsRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetPhase

`func (o *DrugTrialConditionsRow) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *DrugTrialConditionsRow) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *DrugTrialConditionsRow) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *DrugTrialConditionsRow) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *DrugTrialConditionsRow) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *DrugTrialConditionsRow) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetPhaseRank

`func (o *DrugTrialConditionsRow) GetPhaseRank() int32`

GetPhaseRank returns the PhaseRank field if non-nil, zero value otherwise.

### GetPhaseRankOk

`func (o *DrugTrialConditionsRow) GetPhaseRankOk() (*int32, bool)`

GetPhaseRankOk returns a tuple with the PhaseRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseRank

`func (o *DrugTrialConditionsRow) SetPhaseRank(v int32)`

SetPhaseRank sets PhaseRank field to given value.

### HasPhaseRank

`func (o *DrugTrialConditionsRow) HasPhaseRank() bool`

HasPhaseRank returns a boolean if a field has been set.

### SetPhaseRankNil

`func (o *DrugTrialConditionsRow) SetPhaseRankNil(b bool)`

 SetPhaseRankNil sets the value for PhaseRank to be an explicit nil

### UnsetPhaseRank
`func (o *DrugTrialConditionsRow) UnsetPhaseRank()`

UnsetPhaseRank ensures that no value is present for PhaseRank, not even an explicit nil
### GetOverallStatusSummary

`func (o *DrugTrialConditionsRow) GetOverallStatusSummary() string`

GetOverallStatusSummary returns the OverallStatusSummary field if non-nil, zero value otherwise.

### GetOverallStatusSummaryOk

`func (o *DrugTrialConditionsRow) GetOverallStatusSummaryOk() (*string, bool)`

GetOverallStatusSummaryOk returns a tuple with the OverallStatusSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatusSummary

`func (o *DrugTrialConditionsRow) SetOverallStatusSummary(v string)`

SetOverallStatusSummary sets OverallStatusSummary field to given value.

### HasOverallStatusSummary

`func (o *DrugTrialConditionsRow) HasOverallStatusSummary() bool`

HasOverallStatusSummary returns a boolean if a field has been set.

### SetOverallStatusSummaryNil

`func (o *DrugTrialConditionsRow) SetOverallStatusSummaryNil(b bool)`

 SetOverallStatusSummaryNil sets the value for OverallStatusSummary to be an explicit nil

### UnsetOverallStatusSummary
`func (o *DrugTrialConditionsRow) UnsetOverallStatusSummary()`

UnsetOverallStatusSummary ensures that no value is present for OverallStatusSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


