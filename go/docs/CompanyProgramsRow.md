# CompanyProgramsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the lead sponsor&#39;s organization record. | [optional] 
**SponsorName** | Pointer to **NullableString** | The sponsor&#39;s organization name. | [optional] 
**SponsorCompanyId** | Pointer to **NullableString** | Company level: the sponsor&#39;s company (null when unresolved). | [optional] 
**SponsorCompanyName** | Pointer to **NullableString** | Display name of the sponsor&#39;s company. | [optional] 
**SponsorGroupId** | Pointer to **NullableString** | Group level: the sponsor&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**SponsorGroupName** | Pointer to **NullableString** | Display name of the sponsor&#39;s group. | [optional] 
**DrugId** | Pointer to **NullableString** | The drug key. | [optional] 
**DrugName** | Pointer to **NullableString** | The drug&#39;s name. | [optional] 
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s name. | [optional] 
**MaxPhase** | Pointer to **NullableInt32** | Highest phase reached on the 0 to 4 scale: 0 not reported, 1 through 4 for phase 1 through phase 4 (4 is marketed); a split phase folds up to its higher phase. | [optional] 
**MaxPhaseLabel** | Pointer to **NullableString** | The furthest phase reached with the split phases kept: Early 1, Phase 1, Phase 1/2, Phase 2, Phase 2/3, Phase 3, Phase 4, or Not reported. max_phase is the same value folded onto 0 to 4 (a split phase counts as its higher phase). | [optional] 
**MaxPhaseFineRank** | Pointer to **NullableInt32** | The finest phase rank behind max_phase_label: 6 Phase 4, 5 Phase 3, 4 Phase 2/3, 3 Phase 2, 2 Phase 1/2, 1 Phase 1, 0 Early 1, -1 not reported. Sort on it to order programs by furthest phase with the split phases kept; max_phase is the same value folded onto 0 to 4. | [optional] 
**NTrials** | Pointer to **NullableInt64** | Lead-sponsored trials in the program. | [optional] 
**NActive** | Pointer to **NullableInt64** | Of those, trials in an active status. | [optional] 
**LastYear** | Pointer to **NullableInt32** | Latest trial start year. | [optional] 
**Withdrawn** | Pointer to **NullableBool** | True when the drug has been withdrawn. | [optional] 
**Lifecycle** | Pointer to **NullableString** | approved, active or abandoned. | [optional] 
**DalysMillions** | Pointer to **NullableFloat64** | Disability-adjusted life years (DALYs, millions) of the condition&#39;s global burden, from the Global Burden of Disease (GBD) 2021 study of the Institute for Health Metrics and Evaluation (IHME). Null for a caller without the licence for this column; see withheld_columns. | [optional] 

## Methods

### NewCompanyProgramsRow

`func NewCompanyProgramsRow() *CompanyProgramsRow`

NewCompanyProgramsRow instantiates a new CompanyProgramsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyProgramsRowWithDefaults

`func NewCompanyProgramsRowWithDefaults() *CompanyProgramsRow`

NewCompanyProgramsRowWithDefaults instantiates a new CompanyProgramsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSponsorOrgId

`func (o *CompanyProgramsRow) GetSponsorOrgId() string`

GetSponsorOrgId returns the SponsorOrgId field if non-nil, zero value otherwise.

### GetSponsorOrgIdOk

`func (o *CompanyProgramsRow) GetSponsorOrgIdOk() (*string, bool)`

GetSponsorOrgIdOk returns a tuple with the SponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgId

`func (o *CompanyProgramsRow) SetSponsorOrgId(v string)`

SetSponsorOrgId sets SponsorOrgId field to given value.

### HasSponsorOrgId

`func (o *CompanyProgramsRow) HasSponsorOrgId() bool`

HasSponsorOrgId returns a boolean if a field has been set.

### SetSponsorOrgIdNil

`func (o *CompanyProgramsRow) SetSponsorOrgIdNil(b bool)`

 SetSponsorOrgIdNil sets the value for SponsorOrgId to be an explicit nil

### UnsetSponsorOrgId
`func (o *CompanyProgramsRow) UnsetSponsorOrgId()`

UnsetSponsorOrgId ensures that no value is present for SponsorOrgId, not even an explicit nil
### GetSponsorName

`func (o *CompanyProgramsRow) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *CompanyProgramsRow) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *CompanyProgramsRow) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *CompanyProgramsRow) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### SetSponsorNameNil

`func (o *CompanyProgramsRow) SetSponsorNameNil(b bool)`

 SetSponsorNameNil sets the value for SponsorName to be an explicit nil

### UnsetSponsorName
`func (o *CompanyProgramsRow) UnsetSponsorName()`

UnsetSponsorName ensures that no value is present for SponsorName, not even an explicit nil
### GetSponsorCompanyId

`func (o *CompanyProgramsRow) GetSponsorCompanyId() string`

GetSponsorCompanyId returns the SponsorCompanyId field if non-nil, zero value otherwise.

### GetSponsorCompanyIdOk

`func (o *CompanyProgramsRow) GetSponsorCompanyIdOk() (*string, bool)`

GetSponsorCompanyIdOk returns a tuple with the SponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyId

`func (o *CompanyProgramsRow) SetSponsorCompanyId(v string)`

SetSponsorCompanyId sets SponsorCompanyId field to given value.

### HasSponsorCompanyId

`func (o *CompanyProgramsRow) HasSponsorCompanyId() bool`

HasSponsorCompanyId returns a boolean if a field has been set.

### SetSponsorCompanyIdNil

`func (o *CompanyProgramsRow) SetSponsorCompanyIdNil(b bool)`

 SetSponsorCompanyIdNil sets the value for SponsorCompanyId to be an explicit nil

### UnsetSponsorCompanyId
`func (o *CompanyProgramsRow) UnsetSponsorCompanyId()`

UnsetSponsorCompanyId ensures that no value is present for SponsorCompanyId, not even an explicit nil
### GetSponsorCompanyName

`func (o *CompanyProgramsRow) GetSponsorCompanyName() string`

GetSponsorCompanyName returns the SponsorCompanyName field if non-nil, zero value otherwise.

### GetSponsorCompanyNameOk

`func (o *CompanyProgramsRow) GetSponsorCompanyNameOk() (*string, bool)`

GetSponsorCompanyNameOk returns a tuple with the SponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyName

`func (o *CompanyProgramsRow) SetSponsorCompanyName(v string)`

SetSponsorCompanyName sets SponsorCompanyName field to given value.

### HasSponsorCompanyName

`func (o *CompanyProgramsRow) HasSponsorCompanyName() bool`

HasSponsorCompanyName returns a boolean if a field has been set.

### SetSponsorCompanyNameNil

`func (o *CompanyProgramsRow) SetSponsorCompanyNameNil(b bool)`

 SetSponsorCompanyNameNil sets the value for SponsorCompanyName to be an explicit nil

### UnsetSponsorCompanyName
`func (o *CompanyProgramsRow) UnsetSponsorCompanyName()`

UnsetSponsorCompanyName ensures that no value is present for SponsorCompanyName, not even an explicit nil
### GetSponsorGroupId

`func (o *CompanyProgramsRow) GetSponsorGroupId() string`

GetSponsorGroupId returns the SponsorGroupId field if non-nil, zero value otherwise.

### GetSponsorGroupIdOk

`func (o *CompanyProgramsRow) GetSponsorGroupIdOk() (*string, bool)`

GetSponsorGroupIdOk returns a tuple with the SponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupId

`func (o *CompanyProgramsRow) SetSponsorGroupId(v string)`

SetSponsorGroupId sets SponsorGroupId field to given value.

### HasSponsorGroupId

`func (o *CompanyProgramsRow) HasSponsorGroupId() bool`

HasSponsorGroupId returns a boolean if a field has been set.

### SetSponsorGroupIdNil

`func (o *CompanyProgramsRow) SetSponsorGroupIdNil(b bool)`

 SetSponsorGroupIdNil sets the value for SponsorGroupId to be an explicit nil

### UnsetSponsorGroupId
`func (o *CompanyProgramsRow) UnsetSponsorGroupId()`

UnsetSponsorGroupId ensures that no value is present for SponsorGroupId, not even an explicit nil
### GetSponsorGroupName

`func (o *CompanyProgramsRow) GetSponsorGroupName() string`

GetSponsorGroupName returns the SponsorGroupName field if non-nil, zero value otherwise.

### GetSponsorGroupNameOk

`func (o *CompanyProgramsRow) GetSponsorGroupNameOk() (*string, bool)`

GetSponsorGroupNameOk returns a tuple with the SponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupName

`func (o *CompanyProgramsRow) SetSponsorGroupName(v string)`

SetSponsorGroupName sets SponsorGroupName field to given value.

### HasSponsorGroupName

`func (o *CompanyProgramsRow) HasSponsorGroupName() bool`

HasSponsorGroupName returns a boolean if a field has been set.

### SetSponsorGroupNameNil

`func (o *CompanyProgramsRow) SetSponsorGroupNameNil(b bool)`

 SetSponsorGroupNameNil sets the value for SponsorGroupName to be an explicit nil

### UnsetSponsorGroupName
`func (o *CompanyProgramsRow) UnsetSponsorGroupName()`

UnsetSponsorGroupName ensures that no value is present for SponsorGroupName, not even an explicit nil
### GetDrugId

`func (o *CompanyProgramsRow) GetDrugId() string`

GetDrugId returns the DrugId field if non-nil, zero value otherwise.

### GetDrugIdOk

`func (o *CompanyProgramsRow) GetDrugIdOk() (*string, bool)`

GetDrugIdOk returns a tuple with the DrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugId

`func (o *CompanyProgramsRow) SetDrugId(v string)`

SetDrugId sets DrugId field to given value.

### HasDrugId

`func (o *CompanyProgramsRow) HasDrugId() bool`

HasDrugId returns a boolean if a field has been set.

### SetDrugIdNil

`func (o *CompanyProgramsRow) SetDrugIdNil(b bool)`

 SetDrugIdNil sets the value for DrugId to be an explicit nil

### UnsetDrugId
`func (o *CompanyProgramsRow) UnsetDrugId()`

UnsetDrugId ensures that no value is present for DrugId, not even an explicit nil
### GetDrugName

`func (o *CompanyProgramsRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *CompanyProgramsRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *CompanyProgramsRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *CompanyProgramsRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *CompanyProgramsRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *CompanyProgramsRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetConditionId

`func (o *CompanyProgramsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *CompanyProgramsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *CompanyProgramsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *CompanyProgramsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *CompanyProgramsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *CompanyProgramsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *CompanyProgramsRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *CompanyProgramsRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *CompanyProgramsRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *CompanyProgramsRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *CompanyProgramsRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *CompanyProgramsRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetMaxPhase

`func (o *CompanyProgramsRow) GetMaxPhase() int32`

GetMaxPhase returns the MaxPhase field if non-nil, zero value otherwise.

### GetMaxPhaseOk

`func (o *CompanyProgramsRow) GetMaxPhaseOk() (*int32, bool)`

GetMaxPhaseOk returns a tuple with the MaxPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPhase

`func (o *CompanyProgramsRow) SetMaxPhase(v int32)`

SetMaxPhase sets MaxPhase field to given value.

### HasMaxPhase

`func (o *CompanyProgramsRow) HasMaxPhase() bool`

HasMaxPhase returns a boolean if a field has been set.

### SetMaxPhaseNil

`func (o *CompanyProgramsRow) SetMaxPhaseNil(b bool)`

 SetMaxPhaseNil sets the value for MaxPhase to be an explicit nil

### UnsetMaxPhase
`func (o *CompanyProgramsRow) UnsetMaxPhase()`

UnsetMaxPhase ensures that no value is present for MaxPhase, not even an explicit nil
### GetMaxPhaseLabel

`func (o *CompanyProgramsRow) GetMaxPhaseLabel() string`

GetMaxPhaseLabel returns the MaxPhaseLabel field if non-nil, zero value otherwise.

### GetMaxPhaseLabelOk

`func (o *CompanyProgramsRow) GetMaxPhaseLabelOk() (*string, bool)`

GetMaxPhaseLabelOk returns a tuple with the MaxPhaseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPhaseLabel

`func (o *CompanyProgramsRow) SetMaxPhaseLabel(v string)`

SetMaxPhaseLabel sets MaxPhaseLabel field to given value.

### HasMaxPhaseLabel

`func (o *CompanyProgramsRow) HasMaxPhaseLabel() bool`

HasMaxPhaseLabel returns a boolean if a field has been set.

### SetMaxPhaseLabelNil

`func (o *CompanyProgramsRow) SetMaxPhaseLabelNil(b bool)`

 SetMaxPhaseLabelNil sets the value for MaxPhaseLabel to be an explicit nil

### UnsetMaxPhaseLabel
`func (o *CompanyProgramsRow) UnsetMaxPhaseLabel()`

UnsetMaxPhaseLabel ensures that no value is present for MaxPhaseLabel, not even an explicit nil
### GetMaxPhaseFineRank

`func (o *CompanyProgramsRow) GetMaxPhaseFineRank() int32`

GetMaxPhaseFineRank returns the MaxPhaseFineRank field if non-nil, zero value otherwise.

### GetMaxPhaseFineRankOk

`func (o *CompanyProgramsRow) GetMaxPhaseFineRankOk() (*int32, bool)`

GetMaxPhaseFineRankOk returns a tuple with the MaxPhaseFineRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPhaseFineRank

`func (o *CompanyProgramsRow) SetMaxPhaseFineRank(v int32)`

SetMaxPhaseFineRank sets MaxPhaseFineRank field to given value.

### HasMaxPhaseFineRank

`func (o *CompanyProgramsRow) HasMaxPhaseFineRank() bool`

HasMaxPhaseFineRank returns a boolean if a field has been set.

### SetMaxPhaseFineRankNil

`func (o *CompanyProgramsRow) SetMaxPhaseFineRankNil(b bool)`

 SetMaxPhaseFineRankNil sets the value for MaxPhaseFineRank to be an explicit nil

### UnsetMaxPhaseFineRank
`func (o *CompanyProgramsRow) UnsetMaxPhaseFineRank()`

UnsetMaxPhaseFineRank ensures that no value is present for MaxPhaseFineRank, not even an explicit nil
### GetNTrials

`func (o *CompanyProgramsRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *CompanyProgramsRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *CompanyProgramsRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *CompanyProgramsRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *CompanyProgramsRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *CompanyProgramsRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNActive

`func (o *CompanyProgramsRow) GetNActive() int64`

GetNActive returns the NActive field if non-nil, zero value otherwise.

### GetNActiveOk

`func (o *CompanyProgramsRow) GetNActiveOk() (*int64, bool)`

GetNActiveOk returns a tuple with the NActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNActive

`func (o *CompanyProgramsRow) SetNActive(v int64)`

SetNActive sets NActive field to given value.

### HasNActive

`func (o *CompanyProgramsRow) HasNActive() bool`

HasNActive returns a boolean if a field has been set.

### SetNActiveNil

`func (o *CompanyProgramsRow) SetNActiveNil(b bool)`

 SetNActiveNil sets the value for NActive to be an explicit nil

### UnsetNActive
`func (o *CompanyProgramsRow) UnsetNActive()`

UnsetNActive ensures that no value is present for NActive, not even an explicit nil
### GetLastYear

`func (o *CompanyProgramsRow) GetLastYear() int32`

GetLastYear returns the LastYear field if non-nil, zero value otherwise.

### GetLastYearOk

`func (o *CompanyProgramsRow) GetLastYearOk() (*int32, bool)`

GetLastYearOk returns a tuple with the LastYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastYear

`func (o *CompanyProgramsRow) SetLastYear(v int32)`

SetLastYear sets LastYear field to given value.

### HasLastYear

`func (o *CompanyProgramsRow) HasLastYear() bool`

HasLastYear returns a boolean if a field has been set.

### SetLastYearNil

`func (o *CompanyProgramsRow) SetLastYearNil(b bool)`

 SetLastYearNil sets the value for LastYear to be an explicit nil

### UnsetLastYear
`func (o *CompanyProgramsRow) UnsetLastYear()`

UnsetLastYear ensures that no value is present for LastYear, not even an explicit nil
### GetWithdrawn

`func (o *CompanyProgramsRow) GetWithdrawn() bool`

GetWithdrawn returns the Withdrawn field if non-nil, zero value otherwise.

### GetWithdrawnOk

`func (o *CompanyProgramsRow) GetWithdrawnOk() (*bool, bool)`

GetWithdrawnOk returns a tuple with the Withdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawn

`func (o *CompanyProgramsRow) SetWithdrawn(v bool)`

SetWithdrawn sets Withdrawn field to given value.

### HasWithdrawn

`func (o *CompanyProgramsRow) HasWithdrawn() bool`

HasWithdrawn returns a boolean if a field has been set.

### SetWithdrawnNil

`func (o *CompanyProgramsRow) SetWithdrawnNil(b bool)`

 SetWithdrawnNil sets the value for Withdrawn to be an explicit nil

### UnsetWithdrawn
`func (o *CompanyProgramsRow) UnsetWithdrawn()`

UnsetWithdrawn ensures that no value is present for Withdrawn, not even an explicit nil
### GetLifecycle

`func (o *CompanyProgramsRow) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *CompanyProgramsRow) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *CompanyProgramsRow) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *CompanyProgramsRow) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### SetLifecycleNil

`func (o *CompanyProgramsRow) SetLifecycleNil(b bool)`

 SetLifecycleNil sets the value for Lifecycle to be an explicit nil

### UnsetLifecycle
`func (o *CompanyProgramsRow) UnsetLifecycle()`

UnsetLifecycle ensures that no value is present for Lifecycle, not even an explicit nil
### GetDalysMillions

`func (o *CompanyProgramsRow) GetDalysMillions() float64`

GetDalysMillions returns the DalysMillions field if non-nil, zero value otherwise.

### GetDalysMillionsOk

`func (o *CompanyProgramsRow) GetDalysMillionsOk() (*float64, bool)`

GetDalysMillionsOk returns a tuple with the DalysMillions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDalysMillions

`func (o *CompanyProgramsRow) SetDalysMillions(v float64)`

SetDalysMillions sets DalysMillions field to given value.

### HasDalysMillions

`func (o *CompanyProgramsRow) HasDalysMillions() bool`

HasDalysMillions returns a boolean if a field has been set.

### SetDalysMillionsNil

`func (o *CompanyProgramsRow) SetDalysMillionsNil(b bool)`

 SetDalysMillionsNil sets the value for DalysMillions to be an explicit nil

### UnsetDalysMillions
`func (o *CompanyProgramsRow) UnsetDalysMillions()`

UnsetDalysMillions ensures that no value is present for DalysMillions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


