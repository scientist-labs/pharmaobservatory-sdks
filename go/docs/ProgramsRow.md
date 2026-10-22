# ProgramsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramId** | Pointer to **NullableString** | The program key: a hash of drug_id and condition_id, one value per row. | [optional] 
**DrugId** | Pointer to **NullableString** | The drug key of the program, the drug record the trials name (a salt form or a trial arm is its own record). | [optional] 
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug the program&#39;s drug record folds to, and the entry key for the drugs object type. | [optional] 
**DrugName** | Pointer to **NullableString** | The drug&#39;s name. | [optional] 
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s name. | [optional] 
**IsIndication** | Pointer to **NullableBool** | True for a specific disease, false for a broad MeSH umbrella condition. | [optional] 
**MaxPhase** | Pointer to **NullableInt32** | Highest phase reached across the program&#39;s trials: 0 not reported, 1 through 4 for phase 1 through phase 4 (the 0 to 4 scale shared with company_programs; see caveats). | [optional] 
**NTrials** | Pointer to **NullableInt64** | Distinct trials of the program, across all sponsors. | [optional] 
**NCompleted** | Pointer to **NullableInt64** | Of those, trials with status completed. | [optional] 
**NTerminated** | Pointer to **NullableInt64** | Of those, trials with status terminated. | [optional] 
**FirstYear** | Pointer to **NullableInt32** | Earliest plausible trial start year. | [optional] 
**LastYear** | Pointer to **NullableInt32** | Latest plausible trial completion or start year. | [optional] 
**NSponsorCompanies** | Pointer to **NullableInt32** | Distinct lead-sponsor companies across the program&#39;s trials; an unresolved lead sponsor counts as itself. | [optional] 
**DalysMillions** | Pointer to **NullableFloat64** | Disability-adjusted life years (DALYs, millions) of the condition&#39;s global burden, from the Global Burden of Disease (GBD) 2021 study of the Institute for Health Metrics and Evaluation (IHME). Null for a caller without the licence for this column; see withheld_columns. | [optional] 

## Methods

### NewProgramsRow

`func NewProgramsRow() *ProgramsRow`

NewProgramsRow instantiates a new ProgramsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramsRowWithDefaults

`func NewProgramsRowWithDefaults() *ProgramsRow`

NewProgramsRowWithDefaults instantiates a new ProgramsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramId

`func (o *ProgramsRow) GetProgramId() string`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *ProgramsRow) GetProgramIdOk() (*string, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *ProgramsRow) SetProgramId(v string)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *ProgramsRow) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramIdNil

`func (o *ProgramsRow) SetProgramIdNil(b bool)`

 SetProgramIdNil sets the value for ProgramId to be an explicit nil

### UnsetProgramId
`func (o *ProgramsRow) UnsetProgramId()`

UnsetProgramId ensures that no value is present for ProgramId, not even an explicit nil
### GetDrugId

`func (o *ProgramsRow) GetDrugId() string`

GetDrugId returns the DrugId field if non-nil, zero value otherwise.

### GetDrugIdOk

`func (o *ProgramsRow) GetDrugIdOk() (*string, bool)`

GetDrugIdOk returns a tuple with the DrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugId

`func (o *ProgramsRow) SetDrugId(v string)`

SetDrugId sets DrugId field to given value.

### HasDrugId

`func (o *ProgramsRow) HasDrugId() bool`

HasDrugId returns a boolean if a field has been set.

### SetDrugIdNil

`func (o *ProgramsRow) SetDrugIdNil(b bool)`

 SetDrugIdNil sets the value for DrugId to be an explicit nil

### UnsetDrugId
`func (o *ProgramsRow) UnsetDrugId()`

UnsetDrugId ensures that no value is present for DrugId, not even an explicit nil
### GetMoleculeId

`func (o *ProgramsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *ProgramsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *ProgramsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *ProgramsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *ProgramsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *ProgramsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetDrugName

`func (o *ProgramsRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *ProgramsRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *ProgramsRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *ProgramsRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *ProgramsRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *ProgramsRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetConditionId

`func (o *ProgramsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ProgramsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ProgramsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ProgramsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ProgramsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ProgramsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *ProgramsRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *ProgramsRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *ProgramsRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *ProgramsRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *ProgramsRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *ProgramsRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetIsIndication

`func (o *ProgramsRow) GetIsIndication() bool`

GetIsIndication returns the IsIndication field if non-nil, zero value otherwise.

### GetIsIndicationOk

`func (o *ProgramsRow) GetIsIndicationOk() (*bool, bool)`

GetIsIndicationOk returns a tuple with the IsIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIndication

`func (o *ProgramsRow) SetIsIndication(v bool)`

SetIsIndication sets IsIndication field to given value.

### HasIsIndication

`func (o *ProgramsRow) HasIsIndication() bool`

HasIsIndication returns a boolean if a field has been set.

### SetIsIndicationNil

`func (o *ProgramsRow) SetIsIndicationNil(b bool)`

 SetIsIndicationNil sets the value for IsIndication to be an explicit nil

### UnsetIsIndication
`func (o *ProgramsRow) UnsetIsIndication()`

UnsetIsIndication ensures that no value is present for IsIndication, not even an explicit nil
### GetMaxPhase

`func (o *ProgramsRow) GetMaxPhase() int32`

GetMaxPhase returns the MaxPhase field if non-nil, zero value otherwise.

### GetMaxPhaseOk

`func (o *ProgramsRow) GetMaxPhaseOk() (*int32, bool)`

GetMaxPhaseOk returns a tuple with the MaxPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPhase

`func (o *ProgramsRow) SetMaxPhase(v int32)`

SetMaxPhase sets MaxPhase field to given value.

### HasMaxPhase

`func (o *ProgramsRow) HasMaxPhase() bool`

HasMaxPhase returns a boolean if a field has been set.

### SetMaxPhaseNil

`func (o *ProgramsRow) SetMaxPhaseNil(b bool)`

 SetMaxPhaseNil sets the value for MaxPhase to be an explicit nil

### UnsetMaxPhase
`func (o *ProgramsRow) UnsetMaxPhase()`

UnsetMaxPhase ensures that no value is present for MaxPhase, not even an explicit nil
### GetNTrials

`func (o *ProgramsRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *ProgramsRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *ProgramsRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *ProgramsRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *ProgramsRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *ProgramsRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNCompleted

`func (o *ProgramsRow) GetNCompleted() int64`

GetNCompleted returns the NCompleted field if non-nil, zero value otherwise.

### GetNCompletedOk

`func (o *ProgramsRow) GetNCompletedOk() (*int64, bool)`

GetNCompletedOk returns a tuple with the NCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCompleted

`func (o *ProgramsRow) SetNCompleted(v int64)`

SetNCompleted sets NCompleted field to given value.

### HasNCompleted

`func (o *ProgramsRow) HasNCompleted() bool`

HasNCompleted returns a boolean if a field has been set.

### SetNCompletedNil

`func (o *ProgramsRow) SetNCompletedNil(b bool)`

 SetNCompletedNil sets the value for NCompleted to be an explicit nil

### UnsetNCompleted
`func (o *ProgramsRow) UnsetNCompleted()`

UnsetNCompleted ensures that no value is present for NCompleted, not even an explicit nil
### GetNTerminated

`func (o *ProgramsRow) GetNTerminated() int64`

GetNTerminated returns the NTerminated field if non-nil, zero value otherwise.

### GetNTerminatedOk

`func (o *ProgramsRow) GetNTerminatedOk() (*int64, bool)`

GetNTerminatedOk returns a tuple with the NTerminated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTerminated

`func (o *ProgramsRow) SetNTerminated(v int64)`

SetNTerminated sets NTerminated field to given value.

### HasNTerminated

`func (o *ProgramsRow) HasNTerminated() bool`

HasNTerminated returns a boolean if a field has been set.

### SetNTerminatedNil

`func (o *ProgramsRow) SetNTerminatedNil(b bool)`

 SetNTerminatedNil sets the value for NTerminated to be an explicit nil

### UnsetNTerminated
`func (o *ProgramsRow) UnsetNTerminated()`

UnsetNTerminated ensures that no value is present for NTerminated, not even an explicit nil
### GetFirstYear

`func (o *ProgramsRow) GetFirstYear() int32`

GetFirstYear returns the FirstYear field if non-nil, zero value otherwise.

### GetFirstYearOk

`func (o *ProgramsRow) GetFirstYearOk() (*int32, bool)`

GetFirstYearOk returns a tuple with the FirstYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstYear

`func (o *ProgramsRow) SetFirstYear(v int32)`

SetFirstYear sets FirstYear field to given value.

### HasFirstYear

`func (o *ProgramsRow) HasFirstYear() bool`

HasFirstYear returns a boolean if a field has been set.

### SetFirstYearNil

`func (o *ProgramsRow) SetFirstYearNil(b bool)`

 SetFirstYearNil sets the value for FirstYear to be an explicit nil

### UnsetFirstYear
`func (o *ProgramsRow) UnsetFirstYear()`

UnsetFirstYear ensures that no value is present for FirstYear, not even an explicit nil
### GetLastYear

`func (o *ProgramsRow) GetLastYear() int32`

GetLastYear returns the LastYear field if non-nil, zero value otherwise.

### GetLastYearOk

`func (o *ProgramsRow) GetLastYearOk() (*int32, bool)`

GetLastYearOk returns a tuple with the LastYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastYear

`func (o *ProgramsRow) SetLastYear(v int32)`

SetLastYear sets LastYear field to given value.

### HasLastYear

`func (o *ProgramsRow) HasLastYear() bool`

HasLastYear returns a boolean if a field has been set.

### SetLastYearNil

`func (o *ProgramsRow) SetLastYearNil(b bool)`

 SetLastYearNil sets the value for LastYear to be an explicit nil

### UnsetLastYear
`func (o *ProgramsRow) UnsetLastYear()`

UnsetLastYear ensures that no value is present for LastYear, not even an explicit nil
### GetNSponsorCompanies

`func (o *ProgramsRow) GetNSponsorCompanies() int32`

GetNSponsorCompanies returns the NSponsorCompanies field if non-nil, zero value otherwise.

### GetNSponsorCompaniesOk

`func (o *ProgramsRow) GetNSponsorCompaniesOk() (*int32, bool)`

GetNSponsorCompaniesOk returns a tuple with the NSponsorCompanies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSponsorCompanies

`func (o *ProgramsRow) SetNSponsorCompanies(v int32)`

SetNSponsorCompanies sets NSponsorCompanies field to given value.

### HasNSponsorCompanies

`func (o *ProgramsRow) HasNSponsorCompanies() bool`

HasNSponsorCompanies returns a boolean if a field has been set.

### SetNSponsorCompaniesNil

`func (o *ProgramsRow) SetNSponsorCompaniesNil(b bool)`

 SetNSponsorCompaniesNil sets the value for NSponsorCompanies to be an explicit nil

### UnsetNSponsorCompanies
`func (o *ProgramsRow) UnsetNSponsorCompanies()`

UnsetNSponsorCompanies ensures that no value is present for NSponsorCompanies, not even an explicit nil
### GetDalysMillions

`func (o *ProgramsRow) GetDalysMillions() float64`

GetDalysMillions returns the DalysMillions field if non-nil, zero value otherwise.

### GetDalysMillionsOk

`func (o *ProgramsRow) GetDalysMillionsOk() (*float64, bool)`

GetDalysMillionsOk returns a tuple with the DalysMillions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDalysMillions

`func (o *ProgramsRow) SetDalysMillions(v float64)`

SetDalysMillions sets DalysMillions field to given value.

### HasDalysMillions

`func (o *ProgramsRow) HasDalysMillions() bool`

HasDalysMillions returns a boolean if a field has been set.

### SetDalysMillionsNil

`func (o *ProgramsRow) SetDalysMillionsNil(b bool)`

 SetDalysMillionsNil sets the value for DalysMillions to be an explicit nil

### UnsetDalysMillions
`func (o *ProgramsRow) UnsetDalysMillions()`

UnsetDalysMillions ensures that no value is present for DalysMillions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


