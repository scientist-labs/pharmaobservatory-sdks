# DrugTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**Title** | Pointer to **NullableString** | The trial&#39;s brief title. | [optional] 
**Phase** | Pointer to **NullableString** | The registry&#39;s phase code. Null when the registry gives no phase; phase_label carries the display label for those. | [optional] 
**PhaseLabel** | Pointer to **NullableString** | The trial phase as a display label (Phase 2, Phase 3, Observational, Not applicable, Expanded access, Unknown). | [optional] 
**PhaseRank** | Pointer to **NullableInt32** | The development phase as a number for max and sort: 4 is Phase 4, 3 is Phase 3 or 2/3, 2 is Phase 2 or 1/2, 1 is Phase 1 or Early Phase 1, 0 is not reported. | [optional] 
**OverallStatus** | Pointer to **NullableString** | The registry&#39;s overall status. | [optional] 
**OverallStatusSummary** | Pointer to **NullableString** | The overall status grouped as Completed, Ongoing, Terminated or Other. | [optional] 
**StudyType** | Pointer to **NullableString** | INTERVENTIONAL, OBSERVATIONAL or EXPANDED_ACCESS. | [optional] 
**StartDate** | Pointer to **NullableString** | Trial start date. | [optional] 
**PrimaryCompletionDate** | Pointer to **NullableString** | The date primary data collection ends (estimated for an ongoing trial). | [optional] 
**CompletionDate** | Pointer to **NullableString** | The trial&#39;s completion date (estimated for an ongoing trial). | [optional] 
**Enrollment** | Pointer to **NullableInt32** | Enrollment as registered for the whole trial. | [optional] 
**HasResults** | Pointer to **NullableBool** | True when the sponsor has posted results to the registry. | [optional] 
**LeadSponsorName** | Pointer to **NullableString** | The lead sponsor&#39;s name as filed. | [optional] 
**LeadSponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the lead sponsor resolved to (null when unresolved). | [optional] 
**LeadSponsorCompanyId** | Pointer to **NullableString** | Company level: the lead sponsor&#39;s company (null when unresolved). | [optional] 
**LeadSponsorCompanyName** | Pointer to **NullableString** | Display name of the lead sponsor&#39;s company. | [optional] 
**LeadSponsorGroupId** | Pointer to **NullableString** | Group level: the lead sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**LeadSponsorGroupName** | Pointer to **NullableString** | Display name of the lead sponsor&#39;s group. | [optional] 
**ConditionIds** | Pointer to **[]string** | The condition keys (MeSH, Medical Subject Headings, descriptor ids) the trial studies; drug_trial_conditions has one row per specific indication. | [optional] 

## Methods

### NewDrugTrialsRow

`func NewDrugTrialsRow() *DrugTrialsRow`

NewDrugTrialsRow instantiates a new DrugTrialsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugTrialsRowWithDefaults

`func NewDrugTrialsRowWithDefaults() *DrugTrialsRow`

NewDrugTrialsRowWithDefaults instantiates a new DrugTrialsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugTrialsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugTrialsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugTrialsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugTrialsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugTrialsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugTrialsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetTrialId

`func (o *DrugTrialsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *DrugTrialsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *DrugTrialsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *DrugTrialsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *DrugTrialsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *DrugTrialsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetTitle

`func (o *DrugTrialsRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DrugTrialsRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DrugTrialsRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DrugTrialsRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *DrugTrialsRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *DrugTrialsRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPhase

`func (o *DrugTrialsRow) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *DrugTrialsRow) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *DrugTrialsRow) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *DrugTrialsRow) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *DrugTrialsRow) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *DrugTrialsRow) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetPhaseLabel

`func (o *DrugTrialsRow) GetPhaseLabel() string`

GetPhaseLabel returns the PhaseLabel field if non-nil, zero value otherwise.

### GetPhaseLabelOk

`func (o *DrugTrialsRow) GetPhaseLabelOk() (*string, bool)`

GetPhaseLabelOk returns a tuple with the PhaseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseLabel

`func (o *DrugTrialsRow) SetPhaseLabel(v string)`

SetPhaseLabel sets PhaseLabel field to given value.

### HasPhaseLabel

`func (o *DrugTrialsRow) HasPhaseLabel() bool`

HasPhaseLabel returns a boolean if a field has been set.

### SetPhaseLabelNil

`func (o *DrugTrialsRow) SetPhaseLabelNil(b bool)`

 SetPhaseLabelNil sets the value for PhaseLabel to be an explicit nil

### UnsetPhaseLabel
`func (o *DrugTrialsRow) UnsetPhaseLabel()`

UnsetPhaseLabel ensures that no value is present for PhaseLabel, not even an explicit nil
### GetPhaseRank

`func (o *DrugTrialsRow) GetPhaseRank() int32`

GetPhaseRank returns the PhaseRank field if non-nil, zero value otherwise.

### GetPhaseRankOk

`func (o *DrugTrialsRow) GetPhaseRankOk() (*int32, bool)`

GetPhaseRankOk returns a tuple with the PhaseRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseRank

`func (o *DrugTrialsRow) SetPhaseRank(v int32)`

SetPhaseRank sets PhaseRank field to given value.

### HasPhaseRank

`func (o *DrugTrialsRow) HasPhaseRank() bool`

HasPhaseRank returns a boolean if a field has been set.

### SetPhaseRankNil

`func (o *DrugTrialsRow) SetPhaseRankNil(b bool)`

 SetPhaseRankNil sets the value for PhaseRank to be an explicit nil

### UnsetPhaseRank
`func (o *DrugTrialsRow) UnsetPhaseRank()`

UnsetPhaseRank ensures that no value is present for PhaseRank, not even an explicit nil
### GetOverallStatus

`func (o *DrugTrialsRow) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *DrugTrialsRow) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *DrugTrialsRow) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *DrugTrialsRow) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### SetOverallStatusNil

`func (o *DrugTrialsRow) SetOverallStatusNil(b bool)`

 SetOverallStatusNil sets the value for OverallStatus to be an explicit nil

### UnsetOverallStatus
`func (o *DrugTrialsRow) UnsetOverallStatus()`

UnsetOverallStatus ensures that no value is present for OverallStatus, not even an explicit nil
### GetOverallStatusSummary

`func (o *DrugTrialsRow) GetOverallStatusSummary() string`

GetOverallStatusSummary returns the OverallStatusSummary field if non-nil, zero value otherwise.

### GetOverallStatusSummaryOk

`func (o *DrugTrialsRow) GetOverallStatusSummaryOk() (*string, bool)`

GetOverallStatusSummaryOk returns a tuple with the OverallStatusSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatusSummary

`func (o *DrugTrialsRow) SetOverallStatusSummary(v string)`

SetOverallStatusSummary sets OverallStatusSummary field to given value.

### HasOverallStatusSummary

`func (o *DrugTrialsRow) HasOverallStatusSummary() bool`

HasOverallStatusSummary returns a boolean if a field has been set.

### SetOverallStatusSummaryNil

`func (o *DrugTrialsRow) SetOverallStatusSummaryNil(b bool)`

 SetOverallStatusSummaryNil sets the value for OverallStatusSummary to be an explicit nil

### UnsetOverallStatusSummary
`func (o *DrugTrialsRow) UnsetOverallStatusSummary()`

UnsetOverallStatusSummary ensures that no value is present for OverallStatusSummary, not even an explicit nil
### GetStudyType

`func (o *DrugTrialsRow) GetStudyType() string`

GetStudyType returns the StudyType field if non-nil, zero value otherwise.

### GetStudyTypeOk

`func (o *DrugTrialsRow) GetStudyTypeOk() (*string, bool)`

GetStudyTypeOk returns a tuple with the StudyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyType

`func (o *DrugTrialsRow) SetStudyType(v string)`

SetStudyType sets StudyType field to given value.

### HasStudyType

`func (o *DrugTrialsRow) HasStudyType() bool`

HasStudyType returns a boolean if a field has been set.

### SetStudyTypeNil

`func (o *DrugTrialsRow) SetStudyTypeNil(b bool)`

 SetStudyTypeNil sets the value for StudyType to be an explicit nil

### UnsetStudyType
`func (o *DrugTrialsRow) UnsetStudyType()`

UnsetStudyType ensures that no value is present for StudyType, not even an explicit nil
### GetStartDate

`func (o *DrugTrialsRow) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DrugTrialsRow) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DrugTrialsRow) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DrugTrialsRow) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *DrugTrialsRow) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *DrugTrialsRow) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetPrimaryCompletionDate

`func (o *DrugTrialsRow) GetPrimaryCompletionDate() string`

GetPrimaryCompletionDate returns the PrimaryCompletionDate field if non-nil, zero value otherwise.

### GetPrimaryCompletionDateOk

`func (o *DrugTrialsRow) GetPrimaryCompletionDateOk() (*string, bool)`

GetPrimaryCompletionDateOk returns a tuple with the PrimaryCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCompletionDate

`func (o *DrugTrialsRow) SetPrimaryCompletionDate(v string)`

SetPrimaryCompletionDate sets PrimaryCompletionDate field to given value.

### HasPrimaryCompletionDate

`func (o *DrugTrialsRow) HasPrimaryCompletionDate() bool`

HasPrimaryCompletionDate returns a boolean if a field has been set.

### SetPrimaryCompletionDateNil

`func (o *DrugTrialsRow) SetPrimaryCompletionDateNil(b bool)`

 SetPrimaryCompletionDateNil sets the value for PrimaryCompletionDate to be an explicit nil

### UnsetPrimaryCompletionDate
`func (o *DrugTrialsRow) UnsetPrimaryCompletionDate()`

UnsetPrimaryCompletionDate ensures that no value is present for PrimaryCompletionDate, not even an explicit nil
### GetCompletionDate

`func (o *DrugTrialsRow) GetCompletionDate() string`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *DrugTrialsRow) GetCompletionDateOk() (*string, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *DrugTrialsRow) SetCompletionDate(v string)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *DrugTrialsRow) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### SetCompletionDateNil

`func (o *DrugTrialsRow) SetCompletionDateNil(b bool)`

 SetCompletionDateNil sets the value for CompletionDate to be an explicit nil

### UnsetCompletionDate
`func (o *DrugTrialsRow) UnsetCompletionDate()`

UnsetCompletionDate ensures that no value is present for CompletionDate, not even an explicit nil
### GetEnrollment

`func (o *DrugTrialsRow) GetEnrollment() int32`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *DrugTrialsRow) GetEnrollmentOk() (*int32, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *DrugTrialsRow) SetEnrollment(v int32)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *DrugTrialsRow) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### SetEnrollmentNil

`func (o *DrugTrialsRow) SetEnrollmentNil(b bool)`

 SetEnrollmentNil sets the value for Enrollment to be an explicit nil

### UnsetEnrollment
`func (o *DrugTrialsRow) UnsetEnrollment()`

UnsetEnrollment ensures that no value is present for Enrollment, not even an explicit nil
### GetHasResults

`func (o *DrugTrialsRow) GetHasResults() bool`

GetHasResults returns the HasResults field if non-nil, zero value otherwise.

### GetHasResultsOk

`func (o *DrugTrialsRow) GetHasResultsOk() (*bool, bool)`

GetHasResultsOk returns a tuple with the HasResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasResults

`func (o *DrugTrialsRow) SetHasResults(v bool)`

SetHasResults sets HasResults field to given value.

### HasHasResults

`func (o *DrugTrialsRow) HasHasResults() bool`

HasHasResults returns a boolean if a field has been set.

### SetHasResultsNil

`func (o *DrugTrialsRow) SetHasResultsNil(b bool)`

 SetHasResultsNil sets the value for HasResults to be an explicit nil

### UnsetHasResults
`func (o *DrugTrialsRow) UnsetHasResults()`

UnsetHasResults ensures that no value is present for HasResults, not even an explicit nil
### GetLeadSponsorName

`func (o *DrugTrialsRow) GetLeadSponsorName() string`

GetLeadSponsorName returns the LeadSponsorName field if non-nil, zero value otherwise.

### GetLeadSponsorNameOk

`func (o *DrugTrialsRow) GetLeadSponsorNameOk() (*string, bool)`

GetLeadSponsorNameOk returns a tuple with the LeadSponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorName

`func (o *DrugTrialsRow) SetLeadSponsorName(v string)`

SetLeadSponsorName sets LeadSponsorName field to given value.

### HasLeadSponsorName

`func (o *DrugTrialsRow) HasLeadSponsorName() bool`

HasLeadSponsorName returns a boolean if a field has been set.

### SetLeadSponsorNameNil

`func (o *DrugTrialsRow) SetLeadSponsorNameNil(b bool)`

 SetLeadSponsorNameNil sets the value for LeadSponsorName to be an explicit nil

### UnsetLeadSponsorName
`func (o *DrugTrialsRow) UnsetLeadSponsorName()`

UnsetLeadSponsorName ensures that no value is present for LeadSponsorName, not even an explicit nil
### GetLeadSponsorOrgId

`func (o *DrugTrialsRow) GetLeadSponsorOrgId() string`

GetLeadSponsorOrgId returns the LeadSponsorOrgId field if non-nil, zero value otherwise.

### GetLeadSponsorOrgIdOk

`func (o *DrugTrialsRow) GetLeadSponsorOrgIdOk() (*string, bool)`

GetLeadSponsorOrgIdOk returns a tuple with the LeadSponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorOrgId

`func (o *DrugTrialsRow) SetLeadSponsorOrgId(v string)`

SetLeadSponsorOrgId sets LeadSponsorOrgId field to given value.

### HasLeadSponsorOrgId

`func (o *DrugTrialsRow) HasLeadSponsorOrgId() bool`

HasLeadSponsorOrgId returns a boolean if a field has been set.

### SetLeadSponsorOrgIdNil

`func (o *DrugTrialsRow) SetLeadSponsorOrgIdNil(b bool)`

 SetLeadSponsorOrgIdNil sets the value for LeadSponsorOrgId to be an explicit nil

### UnsetLeadSponsorOrgId
`func (o *DrugTrialsRow) UnsetLeadSponsorOrgId()`

UnsetLeadSponsorOrgId ensures that no value is present for LeadSponsorOrgId, not even an explicit nil
### GetLeadSponsorCompanyId

`func (o *DrugTrialsRow) GetLeadSponsorCompanyId() string`

GetLeadSponsorCompanyId returns the LeadSponsorCompanyId field if non-nil, zero value otherwise.

### GetLeadSponsorCompanyIdOk

`func (o *DrugTrialsRow) GetLeadSponsorCompanyIdOk() (*string, bool)`

GetLeadSponsorCompanyIdOk returns a tuple with the LeadSponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorCompanyId

`func (o *DrugTrialsRow) SetLeadSponsorCompanyId(v string)`

SetLeadSponsorCompanyId sets LeadSponsorCompanyId field to given value.

### HasLeadSponsorCompanyId

`func (o *DrugTrialsRow) HasLeadSponsorCompanyId() bool`

HasLeadSponsorCompanyId returns a boolean if a field has been set.

### SetLeadSponsorCompanyIdNil

`func (o *DrugTrialsRow) SetLeadSponsorCompanyIdNil(b bool)`

 SetLeadSponsorCompanyIdNil sets the value for LeadSponsorCompanyId to be an explicit nil

### UnsetLeadSponsorCompanyId
`func (o *DrugTrialsRow) UnsetLeadSponsorCompanyId()`

UnsetLeadSponsorCompanyId ensures that no value is present for LeadSponsorCompanyId, not even an explicit nil
### GetLeadSponsorCompanyName

`func (o *DrugTrialsRow) GetLeadSponsorCompanyName() string`

GetLeadSponsorCompanyName returns the LeadSponsorCompanyName field if non-nil, zero value otherwise.

### GetLeadSponsorCompanyNameOk

`func (o *DrugTrialsRow) GetLeadSponsorCompanyNameOk() (*string, bool)`

GetLeadSponsorCompanyNameOk returns a tuple with the LeadSponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorCompanyName

`func (o *DrugTrialsRow) SetLeadSponsorCompanyName(v string)`

SetLeadSponsorCompanyName sets LeadSponsorCompanyName field to given value.

### HasLeadSponsorCompanyName

`func (o *DrugTrialsRow) HasLeadSponsorCompanyName() bool`

HasLeadSponsorCompanyName returns a boolean if a field has been set.

### SetLeadSponsorCompanyNameNil

`func (o *DrugTrialsRow) SetLeadSponsorCompanyNameNil(b bool)`

 SetLeadSponsorCompanyNameNil sets the value for LeadSponsorCompanyName to be an explicit nil

### UnsetLeadSponsorCompanyName
`func (o *DrugTrialsRow) UnsetLeadSponsorCompanyName()`

UnsetLeadSponsorCompanyName ensures that no value is present for LeadSponsorCompanyName, not even an explicit nil
### GetLeadSponsorGroupId

`func (o *DrugTrialsRow) GetLeadSponsorGroupId() string`

GetLeadSponsorGroupId returns the LeadSponsorGroupId field if non-nil, zero value otherwise.

### GetLeadSponsorGroupIdOk

`func (o *DrugTrialsRow) GetLeadSponsorGroupIdOk() (*string, bool)`

GetLeadSponsorGroupIdOk returns a tuple with the LeadSponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorGroupId

`func (o *DrugTrialsRow) SetLeadSponsorGroupId(v string)`

SetLeadSponsorGroupId sets LeadSponsorGroupId field to given value.

### HasLeadSponsorGroupId

`func (o *DrugTrialsRow) HasLeadSponsorGroupId() bool`

HasLeadSponsorGroupId returns a boolean if a field has been set.

### SetLeadSponsorGroupIdNil

`func (o *DrugTrialsRow) SetLeadSponsorGroupIdNil(b bool)`

 SetLeadSponsorGroupIdNil sets the value for LeadSponsorGroupId to be an explicit nil

### UnsetLeadSponsorGroupId
`func (o *DrugTrialsRow) UnsetLeadSponsorGroupId()`

UnsetLeadSponsorGroupId ensures that no value is present for LeadSponsorGroupId, not even an explicit nil
### GetLeadSponsorGroupName

`func (o *DrugTrialsRow) GetLeadSponsorGroupName() string`

GetLeadSponsorGroupName returns the LeadSponsorGroupName field if non-nil, zero value otherwise.

### GetLeadSponsorGroupNameOk

`func (o *DrugTrialsRow) GetLeadSponsorGroupNameOk() (*string, bool)`

GetLeadSponsorGroupNameOk returns a tuple with the LeadSponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorGroupName

`func (o *DrugTrialsRow) SetLeadSponsorGroupName(v string)`

SetLeadSponsorGroupName sets LeadSponsorGroupName field to given value.

### HasLeadSponsorGroupName

`func (o *DrugTrialsRow) HasLeadSponsorGroupName() bool`

HasLeadSponsorGroupName returns a boolean if a field has been set.

### SetLeadSponsorGroupNameNil

`func (o *DrugTrialsRow) SetLeadSponsorGroupNameNil(b bool)`

 SetLeadSponsorGroupNameNil sets the value for LeadSponsorGroupName to be an explicit nil

### UnsetLeadSponsorGroupName
`func (o *DrugTrialsRow) UnsetLeadSponsorGroupName()`

UnsetLeadSponsorGroupName ensures that no value is present for LeadSponsorGroupName, not even an explicit nil
### GetConditionIds

`func (o *DrugTrialsRow) GetConditionIds() []string`

GetConditionIds returns the ConditionIds field if non-nil, zero value otherwise.

### GetConditionIdsOk

`func (o *DrugTrialsRow) GetConditionIdsOk() (*[]string, bool)`

GetConditionIdsOk returns a tuple with the ConditionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionIds

`func (o *DrugTrialsRow) SetConditionIds(v []string)`

SetConditionIds sets ConditionIds field to given value.

### HasConditionIds

`func (o *DrugTrialsRow) HasConditionIds() bool`

HasConditionIds returns a boolean if a field has been set.

### SetConditionIdsNil

`func (o *DrugTrialsRow) SetConditionIdsNil(b bool)`

 SetConditionIdsNil sets the value for ConditionIds to be an explicit nil

### UnsetConditionIds
`func (o *DrugTrialsRow) UnsetConditionIds()`

UnsetConditionIds ensures that no value is present for ConditionIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


