# TargetTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **NullableString** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] 
**TargetSymbol** | Pointer to **NullableString** | The target&#39;s HGNC gene symbol (the id the target page address carries). | [optional] 
**TrialId** | Pointer to **NullableString** | The ClinicalTrials.gov NCT id, the id the trials datasets use. | [optional] 
**EvidenceSource** | Pointer to **NullableString** | The best evidence tier among the trial&#39;s drugs that act on the target. | [optional] 
**BriefTitle** | Pointer to **NullableString** | The trial&#39;s brief title. | [optional] 
**Phase** | Pointer to **NullableString** | The registry&#39;s phase code (PHASE1, PHASE2, PHASE2/PHASE3, NA and so on). Null when the registry gives no phase (observational and expanded-access trials); phase_label carries the display label for those. | [optional] 
**PhaseLabel** | Pointer to **NullableString** | The trial phase as a display label (Phase 2, Phase 3, Observational, Not applicable, Expanded access, Unknown). The same trial&#39;s phase column holds the registry code. | [optional] 
**OverallStatus** | Pointer to **NullableString** | The registry&#39;s overall recruitment status. | [optional] 
**StartYear** | Pointer to **NullableInt32** | The year the trial started. | [optional] 
**LeadSponsorName** | Pointer to **NullableString** | The lead sponsor&#39;s name as filed. | [optional] 
**LeadSponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the lead sponsor organization; null when it does not resolve. | [optional] 
**LeadSponsorCompanyId** | Pointer to **NullableString** | Company level: the lead sponsor&#39;s company; null when none. | [optional] 
**LeadSponsorCompanyName** | Pointer to **NullableString** | Display name of the lead sponsor&#39;s company. | [optional] 
**LeadSponsorGroupId** | Pointer to **NullableString** | Group level: the lead sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**LeadSponsorGroupName** | Pointer to **NullableString** | Display name of the lead sponsor&#39;s group. | [optional] 

## Methods

### NewTargetTrialsRow

`func NewTargetTrialsRow() *TargetTrialsRow`

NewTargetTrialsRow instantiates a new TargetTrialsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetTrialsRowWithDefaults

`func NewTargetTrialsRowWithDefaults() *TargetTrialsRow`

NewTargetTrialsRowWithDefaults instantiates a new TargetTrialsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *TargetTrialsRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TargetTrialsRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TargetTrialsRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *TargetTrialsRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *TargetTrialsRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *TargetTrialsRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetSymbol

`func (o *TargetTrialsRow) GetTargetSymbol() string`

GetTargetSymbol returns the TargetSymbol field if non-nil, zero value otherwise.

### GetTargetSymbolOk

`func (o *TargetTrialsRow) GetTargetSymbolOk() (*string, bool)`

GetTargetSymbolOk returns a tuple with the TargetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSymbol

`func (o *TargetTrialsRow) SetTargetSymbol(v string)`

SetTargetSymbol sets TargetSymbol field to given value.

### HasTargetSymbol

`func (o *TargetTrialsRow) HasTargetSymbol() bool`

HasTargetSymbol returns a boolean if a field has been set.

### SetTargetSymbolNil

`func (o *TargetTrialsRow) SetTargetSymbolNil(b bool)`

 SetTargetSymbolNil sets the value for TargetSymbol to be an explicit nil

### UnsetTargetSymbol
`func (o *TargetTrialsRow) UnsetTargetSymbol()`

UnsetTargetSymbol ensures that no value is present for TargetSymbol, not even an explicit nil
### GetTrialId

`func (o *TargetTrialsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TargetTrialsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TargetTrialsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TargetTrialsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TargetTrialsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TargetTrialsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetEvidenceSource

`func (o *TargetTrialsRow) GetEvidenceSource() string`

GetEvidenceSource returns the EvidenceSource field if non-nil, zero value otherwise.

### GetEvidenceSourceOk

`func (o *TargetTrialsRow) GetEvidenceSourceOk() (*string, bool)`

GetEvidenceSourceOk returns a tuple with the EvidenceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceSource

`func (o *TargetTrialsRow) SetEvidenceSource(v string)`

SetEvidenceSource sets EvidenceSource field to given value.

### HasEvidenceSource

`func (o *TargetTrialsRow) HasEvidenceSource() bool`

HasEvidenceSource returns a boolean if a field has been set.

### SetEvidenceSourceNil

`func (o *TargetTrialsRow) SetEvidenceSourceNil(b bool)`

 SetEvidenceSourceNil sets the value for EvidenceSource to be an explicit nil

### UnsetEvidenceSource
`func (o *TargetTrialsRow) UnsetEvidenceSource()`

UnsetEvidenceSource ensures that no value is present for EvidenceSource, not even an explicit nil
### GetBriefTitle

`func (o *TargetTrialsRow) GetBriefTitle() string`

GetBriefTitle returns the BriefTitle field if non-nil, zero value otherwise.

### GetBriefTitleOk

`func (o *TargetTrialsRow) GetBriefTitleOk() (*string, bool)`

GetBriefTitleOk returns a tuple with the BriefTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBriefTitle

`func (o *TargetTrialsRow) SetBriefTitle(v string)`

SetBriefTitle sets BriefTitle field to given value.

### HasBriefTitle

`func (o *TargetTrialsRow) HasBriefTitle() bool`

HasBriefTitle returns a boolean if a field has been set.

### SetBriefTitleNil

`func (o *TargetTrialsRow) SetBriefTitleNil(b bool)`

 SetBriefTitleNil sets the value for BriefTitle to be an explicit nil

### UnsetBriefTitle
`func (o *TargetTrialsRow) UnsetBriefTitle()`

UnsetBriefTitle ensures that no value is present for BriefTitle, not even an explicit nil
### GetPhase

`func (o *TargetTrialsRow) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *TargetTrialsRow) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *TargetTrialsRow) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *TargetTrialsRow) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *TargetTrialsRow) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *TargetTrialsRow) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetPhaseLabel

`func (o *TargetTrialsRow) GetPhaseLabel() string`

GetPhaseLabel returns the PhaseLabel field if non-nil, zero value otherwise.

### GetPhaseLabelOk

`func (o *TargetTrialsRow) GetPhaseLabelOk() (*string, bool)`

GetPhaseLabelOk returns a tuple with the PhaseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseLabel

`func (o *TargetTrialsRow) SetPhaseLabel(v string)`

SetPhaseLabel sets PhaseLabel field to given value.

### HasPhaseLabel

`func (o *TargetTrialsRow) HasPhaseLabel() bool`

HasPhaseLabel returns a boolean if a field has been set.

### SetPhaseLabelNil

`func (o *TargetTrialsRow) SetPhaseLabelNil(b bool)`

 SetPhaseLabelNil sets the value for PhaseLabel to be an explicit nil

### UnsetPhaseLabel
`func (o *TargetTrialsRow) UnsetPhaseLabel()`

UnsetPhaseLabel ensures that no value is present for PhaseLabel, not even an explicit nil
### GetOverallStatus

`func (o *TargetTrialsRow) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *TargetTrialsRow) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *TargetTrialsRow) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *TargetTrialsRow) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### SetOverallStatusNil

`func (o *TargetTrialsRow) SetOverallStatusNil(b bool)`

 SetOverallStatusNil sets the value for OverallStatus to be an explicit nil

### UnsetOverallStatus
`func (o *TargetTrialsRow) UnsetOverallStatus()`

UnsetOverallStatus ensures that no value is present for OverallStatus, not even an explicit nil
### GetStartYear

`func (o *TargetTrialsRow) GetStartYear() int32`

GetStartYear returns the StartYear field if non-nil, zero value otherwise.

### GetStartYearOk

`func (o *TargetTrialsRow) GetStartYearOk() (*int32, bool)`

GetStartYearOk returns a tuple with the StartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartYear

`func (o *TargetTrialsRow) SetStartYear(v int32)`

SetStartYear sets StartYear field to given value.

### HasStartYear

`func (o *TargetTrialsRow) HasStartYear() bool`

HasStartYear returns a boolean if a field has been set.

### SetStartYearNil

`func (o *TargetTrialsRow) SetStartYearNil(b bool)`

 SetStartYearNil sets the value for StartYear to be an explicit nil

### UnsetStartYear
`func (o *TargetTrialsRow) UnsetStartYear()`

UnsetStartYear ensures that no value is present for StartYear, not even an explicit nil
### GetLeadSponsorName

`func (o *TargetTrialsRow) GetLeadSponsorName() string`

GetLeadSponsorName returns the LeadSponsorName field if non-nil, zero value otherwise.

### GetLeadSponsorNameOk

`func (o *TargetTrialsRow) GetLeadSponsorNameOk() (*string, bool)`

GetLeadSponsorNameOk returns a tuple with the LeadSponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorName

`func (o *TargetTrialsRow) SetLeadSponsorName(v string)`

SetLeadSponsorName sets LeadSponsorName field to given value.

### HasLeadSponsorName

`func (o *TargetTrialsRow) HasLeadSponsorName() bool`

HasLeadSponsorName returns a boolean if a field has been set.

### SetLeadSponsorNameNil

`func (o *TargetTrialsRow) SetLeadSponsorNameNil(b bool)`

 SetLeadSponsorNameNil sets the value for LeadSponsorName to be an explicit nil

### UnsetLeadSponsorName
`func (o *TargetTrialsRow) UnsetLeadSponsorName()`

UnsetLeadSponsorName ensures that no value is present for LeadSponsorName, not even an explicit nil
### GetLeadSponsorOrgId

`func (o *TargetTrialsRow) GetLeadSponsorOrgId() string`

GetLeadSponsorOrgId returns the LeadSponsorOrgId field if non-nil, zero value otherwise.

### GetLeadSponsorOrgIdOk

`func (o *TargetTrialsRow) GetLeadSponsorOrgIdOk() (*string, bool)`

GetLeadSponsorOrgIdOk returns a tuple with the LeadSponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorOrgId

`func (o *TargetTrialsRow) SetLeadSponsorOrgId(v string)`

SetLeadSponsorOrgId sets LeadSponsorOrgId field to given value.

### HasLeadSponsorOrgId

`func (o *TargetTrialsRow) HasLeadSponsorOrgId() bool`

HasLeadSponsorOrgId returns a boolean if a field has been set.

### SetLeadSponsorOrgIdNil

`func (o *TargetTrialsRow) SetLeadSponsorOrgIdNil(b bool)`

 SetLeadSponsorOrgIdNil sets the value for LeadSponsorOrgId to be an explicit nil

### UnsetLeadSponsorOrgId
`func (o *TargetTrialsRow) UnsetLeadSponsorOrgId()`

UnsetLeadSponsorOrgId ensures that no value is present for LeadSponsorOrgId, not even an explicit nil
### GetLeadSponsorCompanyId

`func (o *TargetTrialsRow) GetLeadSponsorCompanyId() string`

GetLeadSponsorCompanyId returns the LeadSponsorCompanyId field if non-nil, zero value otherwise.

### GetLeadSponsorCompanyIdOk

`func (o *TargetTrialsRow) GetLeadSponsorCompanyIdOk() (*string, bool)`

GetLeadSponsorCompanyIdOk returns a tuple with the LeadSponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorCompanyId

`func (o *TargetTrialsRow) SetLeadSponsorCompanyId(v string)`

SetLeadSponsorCompanyId sets LeadSponsorCompanyId field to given value.

### HasLeadSponsorCompanyId

`func (o *TargetTrialsRow) HasLeadSponsorCompanyId() bool`

HasLeadSponsorCompanyId returns a boolean if a field has been set.

### SetLeadSponsorCompanyIdNil

`func (o *TargetTrialsRow) SetLeadSponsorCompanyIdNil(b bool)`

 SetLeadSponsorCompanyIdNil sets the value for LeadSponsorCompanyId to be an explicit nil

### UnsetLeadSponsorCompanyId
`func (o *TargetTrialsRow) UnsetLeadSponsorCompanyId()`

UnsetLeadSponsorCompanyId ensures that no value is present for LeadSponsorCompanyId, not even an explicit nil
### GetLeadSponsorCompanyName

`func (o *TargetTrialsRow) GetLeadSponsorCompanyName() string`

GetLeadSponsorCompanyName returns the LeadSponsorCompanyName field if non-nil, zero value otherwise.

### GetLeadSponsorCompanyNameOk

`func (o *TargetTrialsRow) GetLeadSponsorCompanyNameOk() (*string, bool)`

GetLeadSponsorCompanyNameOk returns a tuple with the LeadSponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorCompanyName

`func (o *TargetTrialsRow) SetLeadSponsorCompanyName(v string)`

SetLeadSponsorCompanyName sets LeadSponsorCompanyName field to given value.

### HasLeadSponsorCompanyName

`func (o *TargetTrialsRow) HasLeadSponsorCompanyName() bool`

HasLeadSponsorCompanyName returns a boolean if a field has been set.

### SetLeadSponsorCompanyNameNil

`func (o *TargetTrialsRow) SetLeadSponsorCompanyNameNil(b bool)`

 SetLeadSponsorCompanyNameNil sets the value for LeadSponsorCompanyName to be an explicit nil

### UnsetLeadSponsorCompanyName
`func (o *TargetTrialsRow) UnsetLeadSponsorCompanyName()`

UnsetLeadSponsorCompanyName ensures that no value is present for LeadSponsorCompanyName, not even an explicit nil
### GetLeadSponsorGroupId

`func (o *TargetTrialsRow) GetLeadSponsorGroupId() string`

GetLeadSponsorGroupId returns the LeadSponsorGroupId field if non-nil, zero value otherwise.

### GetLeadSponsorGroupIdOk

`func (o *TargetTrialsRow) GetLeadSponsorGroupIdOk() (*string, bool)`

GetLeadSponsorGroupIdOk returns a tuple with the LeadSponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorGroupId

`func (o *TargetTrialsRow) SetLeadSponsorGroupId(v string)`

SetLeadSponsorGroupId sets LeadSponsorGroupId field to given value.

### HasLeadSponsorGroupId

`func (o *TargetTrialsRow) HasLeadSponsorGroupId() bool`

HasLeadSponsorGroupId returns a boolean if a field has been set.

### SetLeadSponsorGroupIdNil

`func (o *TargetTrialsRow) SetLeadSponsorGroupIdNil(b bool)`

 SetLeadSponsorGroupIdNil sets the value for LeadSponsorGroupId to be an explicit nil

### UnsetLeadSponsorGroupId
`func (o *TargetTrialsRow) UnsetLeadSponsorGroupId()`

UnsetLeadSponsorGroupId ensures that no value is present for LeadSponsorGroupId, not even an explicit nil
### GetLeadSponsorGroupName

`func (o *TargetTrialsRow) GetLeadSponsorGroupName() string`

GetLeadSponsorGroupName returns the LeadSponsorGroupName field if non-nil, zero value otherwise.

### GetLeadSponsorGroupNameOk

`func (o *TargetTrialsRow) GetLeadSponsorGroupNameOk() (*string, bool)`

GetLeadSponsorGroupNameOk returns a tuple with the LeadSponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorGroupName

`func (o *TargetTrialsRow) SetLeadSponsorGroupName(v string)`

SetLeadSponsorGroupName sets LeadSponsorGroupName field to given value.

### HasLeadSponsorGroupName

`func (o *TargetTrialsRow) HasLeadSponsorGroupName() bool`

HasLeadSponsorGroupName returns a boolean if a field has been set.

### SetLeadSponsorGroupNameNil

`func (o *TargetTrialsRow) SetLeadSponsorGroupNameNil(b bool)`

 SetLeadSponsorGroupNameNil sets the value for LeadSponsorGroupName to be an explicit nil

### UnsetLeadSponsorGroupName
`func (o *TargetTrialsRow) UnsetLeadSponsorGroupName()`

UnsetLeadSponsorGroupName ensures that no value is present for LeadSponsorGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


