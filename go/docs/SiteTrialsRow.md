# SiteTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **NullableString** | The site key: one facility spelling from the registry. | [optional] 
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**TrialTitle** | Pointer to **NullableString** | The trial&#39;s brief title. | [optional] 
**StudyType** | Pointer to **NullableString** | INTERVENTIONAL or OBSERVATIONAL. | [optional] 
**Phase** | Pointer to **NullableString** | The registry&#39;s phase code (PHASE1, PHASE2, PHASE2/PHASE3, NA and so on). Null when the registry gives no phase (observational and expanded-access trials); phase_label carries the display label for those. | [optional] 
**PhaseLabel** | Pointer to **NullableString** | The trial phase as a display label (Phase 2, Phase 3, Observational, Not applicable, Expanded access, Unknown). The same trial&#39;s phase column holds the registry code. | [optional] 
**OverallStatus** | Pointer to **NullableString** | The registry&#39;s overall status. | [optional] 
**StartDate** | Pointer to **NullableString** | Trial start date. | [optional] 
**StartYear** | Pointer to **NullableInt32** | Year of the trial start. | [optional] 
**Enrollment** | Pointer to **NullableInt32** | Enrollment as registered for the whole trial. | [optional] 
**LeadSponsorName** | Pointer to **NullableString** | The lead sponsor&#39;s name as filed. | [optional] 
**LeadSponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the lead sponsor resolved to (null when unresolved). | [optional] 
**LeadSponsorCompanyId** | Pointer to **NullableString** | Company level: the company the lead sponsor belongs to (null when unresolved). | [optional] 
**LeadSponsorCompanyName** | Pointer to **NullableString** | Display name of the lead sponsor&#39;s company. | [optional] 
**LeadSponsorGroupId** | Pointer to **NullableString** | Group level: the lead sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**LeadSponsorGroupName** | Pointer to **NullableString** | Display name of the lead sponsor&#39;s group. | [optional] 

## Methods

### NewSiteTrialsRow

`func NewSiteTrialsRow() *SiteTrialsRow`

NewSiteTrialsRow instantiates a new SiteTrialsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteTrialsRowWithDefaults

`func NewSiteTrialsRowWithDefaults() *SiteTrialsRow`

NewSiteTrialsRowWithDefaults instantiates a new SiteTrialsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *SiteTrialsRow) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SiteTrialsRow) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SiteTrialsRow) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SiteTrialsRow) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *SiteTrialsRow) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *SiteTrialsRow) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetTrialId

`func (o *SiteTrialsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *SiteTrialsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *SiteTrialsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *SiteTrialsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *SiteTrialsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *SiteTrialsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetTrialTitle

`func (o *SiteTrialsRow) GetTrialTitle() string`

GetTrialTitle returns the TrialTitle field if non-nil, zero value otherwise.

### GetTrialTitleOk

`func (o *SiteTrialsRow) GetTrialTitleOk() (*string, bool)`

GetTrialTitleOk returns a tuple with the TrialTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialTitle

`func (o *SiteTrialsRow) SetTrialTitle(v string)`

SetTrialTitle sets TrialTitle field to given value.

### HasTrialTitle

`func (o *SiteTrialsRow) HasTrialTitle() bool`

HasTrialTitle returns a boolean if a field has been set.

### SetTrialTitleNil

`func (o *SiteTrialsRow) SetTrialTitleNil(b bool)`

 SetTrialTitleNil sets the value for TrialTitle to be an explicit nil

### UnsetTrialTitle
`func (o *SiteTrialsRow) UnsetTrialTitle()`

UnsetTrialTitle ensures that no value is present for TrialTitle, not even an explicit nil
### GetStudyType

`func (o *SiteTrialsRow) GetStudyType() string`

GetStudyType returns the StudyType field if non-nil, zero value otherwise.

### GetStudyTypeOk

`func (o *SiteTrialsRow) GetStudyTypeOk() (*string, bool)`

GetStudyTypeOk returns a tuple with the StudyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyType

`func (o *SiteTrialsRow) SetStudyType(v string)`

SetStudyType sets StudyType field to given value.

### HasStudyType

`func (o *SiteTrialsRow) HasStudyType() bool`

HasStudyType returns a boolean if a field has been set.

### SetStudyTypeNil

`func (o *SiteTrialsRow) SetStudyTypeNil(b bool)`

 SetStudyTypeNil sets the value for StudyType to be an explicit nil

### UnsetStudyType
`func (o *SiteTrialsRow) UnsetStudyType()`

UnsetStudyType ensures that no value is present for StudyType, not even an explicit nil
### GetPhase

`func (o *SiteTrialsRow) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *SiteTrialsRow) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *SiteTrialsRow) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *SiteTrialsRow) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *SiteTrialsRow) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *SiteTrialsRow) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetPhaseLabel

`func (o *SiteTrialsRow) GetPhaseLabel() string`

GetPhaseLabel returns the PhaseLabel field if non-nil, zero value otherwise.

### GetPhaseLabelOk

`func (o *SiteTrialsRow) GetPhaseLabelOk() (*string, bool)`

GetPhaseLabelOk returns a tuple with the PhaseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseLabel

`func (o *SiteTrialsRow) SetPhaseLabel(v string)`

SetPhaseLabel sets PhaseLabel field to given value.

### HasPhaseLabel

`func (o *SiteTrialsRow) HasPhaseLabel() bool`

HasPhaseLabel returns a boolean if a field has been set.

### SetPhaseLabelNil

`func (o *SiteTrialsRow) SetPhaseLabelNil(b bool)`

 SetPhaseLabelNil sets the value for PhaseLabel to be an explicit nil

### UnsetPhaseLabel
`func (o *SiteTrialsRow) UnsetPhaseLabel()`

UnsetPhaseLabel ensures that no value is present for PhaseLabel, not even an explicit nil
### GetOverallStatus

`func (o *SiteTrialsRow) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *SiteTrialsRow) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *SiteTrialsRow) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *SiteTrialsRow) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### SetOverallStatusNil

`func (o *SiteTrialsRow) SetOverallStatusNil(b bool)`

 SetOverallStatusNil sets the value for OverallStatus to be an explicit nil

### UnsetOverallStatus
`func (o *SiteTrialsRow) UnsetOverallStatus()`

UnsetOverallStatus ensures that no value is present for OverallStatus, not even an explicit nil
### GetStartDate

`func (o *SiteTrialsRow) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SiteTrialsRow) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SiteTrialsRow) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SiteTrialsRow) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *SiteTrialsRow) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *SiteTrialsRow) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetStartYear

`func (o *SiteTrialsRow) GetStartYear() int32`

GetStartYear returns the StartYear field if non-nil, zero value otherwise.

### GetStartYearOk

`func (o *SiteTrialsRow) GetStartYearOk() (*int32, bool)`

GetStartYearOk returns a tuple with the StartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartYear

`func (o *SiteTrialsRow) SetStartYear(v int32)`

SetStartYear sets StartYear field to given value.

### HasStartYear

`func (o *SiteTrialsRow) HasStartYear() bool`

HasStartYear returns a boolean if a field has been set.

### SetStartYearNil

`func (o *SiteTrialsRow) SetStartYearNil(b bool)`

 SetStartYearNil sets the value for StartYear to be an explicit nil

### UnsetStartYear
`func (o *SiteTrialsRow) UnsetStartYear()`

UnsetStartYear ensures that no value is present for StartYear, not even an explicit nil
### GetEnrollment

`func (o *SiteTrialsRow) GetEnrollment() int32`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *SiteTrialsRow) GetEnrollmentOk() (*int32, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *SiteTrialsRow) SetEnrollment(v int32)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *SiteTrialsRow) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### SetEnrollmentNil

`func (o *SiteTrialsRow) SetEnrollmentNil(b bool)`

 SetEnrollmentNil sets the value for Enrollment to be an explicit nil

### UnsetEnrollment
`func (o *SiteTrialsRow) UnsetEnrollment()`

UnsetEnrollment ensures that no value is present for Enrollment, not even an explicit nil
### GetLeadSponsorName

`func (o *SiteTrialsRow) GetLeadSponsorName() string`

GetLeadSponsorName returns the LeadSponsorName field if non-nil, zero value otherwise.

### GetLeadSponsorNameOk

`func (o *SiteTrialsRow) GetLeadSponsorNameOk() (*string, bool)`

GetLeadSponsorNameOk returns a tuple with the LeadSponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorName

`func (o *SiteTrialsRow) SetLeadSponsorName(v string)`

SetLeadSponsorName sets LeadSponsorName field to given value.

### HasLeadSponsorName

`func (o *SiteTrialsRow) HasLeadSponsorName() bool`

HasLeadSponsorName returns a boolean if a field has been set.

### SetLeadSponsorNameNil

`func (o *SiteTrialsRow) SetLeadSponsorNameNil(b bool)`

 SetLeadSponsorNameNil sets the value for LeadSponsorName to be an explicit nil

### UnsetLeadSponsorName
`func (o *SiteTrialsRow) UnsetLeadSponsorName()`

UnsetLeadSponsorName ensures that no value is present for LeadSponsorName, not even an explicit nil
### GetLeadSponsorOrgId

`func (o *SiteTrialsRow) GetLeadSponsorOrgId() string`

GetLeadSponsorOrgId returns the LeadSponsorOrgId field if non-nil, zero value otherwise.

### GetLeadSponsorOrgIdOk

`func (o *SiteTrialsRow) GetLeadSponsorOrgIdOk() (*string, bool)`

GetLeadSponsorOrgIdOk returns a tuple with the LeadSponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorOrgId

`func (o *SiteTrialsRow) SetLeadSponsorOrgId(v string)`

SetLeadSponsorOrgId sets LeadSponsorOrgId field to given value.

### HasLeadSponsorOrgId

`func (o *SiteTrialsRow) HasLeadSponsorOrgId() bool`

HasLeadSponsorOrgId returns a boolean if a field has been set.

### SetLeadSponsorOrgIdNil

`func (o *SiteTrialsRow) SetLeadSponsorOrgIdNil(b bool)`

 SetLeadSponsorOrgIdNil sets the value for LeadSponsorOrgId to be an explicit nil

### UnsetLeadSponsorOrgId
`func (o *SiteTrialsRow) UnsetLeadSponsorOrgId()`

UnsetLeadSponsorOrgId ensures that no value is present for LeadSponsorOrgId, not even an explicit nil
### GetLeadSponsorCompanyId

`func (o *SiteTrialsRow) GetLeadSponsorCompanyId() string`

GetLeadSponsorCompanyId returns the LeadSponsorCompanyId field if non-nil, zero value otherwise.

### GetLeadSponsorCompanyIdOk

`func (o *SiteTrialsRow) GetLeadSponsorCompanyIdOk() (*string, bool)`

GetLeadSponsorCompanyIdOk returns a tuple with the LeadSponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorCompanyId

`func (o *SiteTrialsRow) SetLeadSponsorCompanyId(v string)`

SetLeadSponsorCompanyId sets LeadSponsorCompanyId field to given value.

### HasLeadSponsorCompanyId

`func (o *SiteTrialsRow) HasLeadSponsorCompanyId() bool`

HasLeadSponsorCompanyId returns a boolean if a field has been set.

### SetLeadSponsorCompanyIdNil

`func (o *SiteTrialsRow) SetLeadSponsorCompanyIdNil(b bool)`

 SetLeadSponsorCompanyIdNil sets the value for LeadSponsorCompanyId to be an explicit nil

### UnsetLeadSponsorCompanyId
`func (o *SiteTrialsRow) UnsetLeadSponsorCompanyId()`

UnsetLeadSponsorCompanyId ensures that no value is present for LeadSponsorCompanyId, not even an explicit nil
### GetLeadSponsorCompanyName

`func (o *SiteTrialsRow) GetLeadSponsorCompanyName() string`

GetLeadSponsorCompanyName returns the LeadSponsorCompanyName field if non-nil, zero value otherwise.

### GetLeadSponsorCompanyNameOk

`func (o *SiteTrialsRow) GetLeadSponsorCompanyNameOk() (*string, bool)`

GetLeadSponsorCompanyNameOk returns a tuple with the LeadSponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorCompanyName

`func (o *SiteTrialsRow) SetLeadSponsorCompanyName(v string)`

SetLeadSponsorCompanyName sets LeadSponsorCompanyName field to given value.

### HasLeadSponsorCompanyName

`func (o *SiteTrialsRow) HasLeadSponsorCompanyName() bool`

HasLeadSponsorCompanyName returns a boolean if a field has been set.

### SetLeadSponsorCompanyNameNil

`func (o *SiteTrialsRow) SetLeadSponsorCompanyNameNil(b bool)`

 SetLeadSponsorCompanyNameNil sets the value for LeadSponsorCompanyName to be an explicit nil

### UnsetLeadSponsorCompanyName
`func (o *SiteTrialsRow) UnsetLeadSponsorCompanyName()`

UnsetLeadSponsorCompanyName ensures that no value is present for LeadSponsorCompanyName, not even an explicit nil
### GetLeadSponsorGroupId

`func (o *SiteTrialsRow) GetLeadSponsorGroupId() string`

GetLeadSponsorGroupId returns the LeadSponsorGroupId field if non-nil, zero value otherwise.

### GetLeadSponsorGroupIdOk

`func (o *SiteTrialsRow) GetLeadSponsorGroupIdOk() (*string, bool)`

GetLeadSponsorGroupIdOk returns a tuple with the LeadSponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorGroupId

`func (o *SiteTrialsRow) SetLeadSponsorGroupId(v string)`

SetLeadSponsorGroupId sets LeadSponsorGroupId field to given value.

### HasLeadSponsorGroupId

`func (o *SiteTrialsRow) HasLeadSponsorGroupId() bool`

HasLeadSponsorGroupId returns a boolean if a field has been set.

### SetLeadSponsorGroupIdNil

`func (o *SiteTrialsRow) SetLeadSponsorGroupIdNil(b bool)`

 SetLeadSponsorGroupIdNil sets the value for LeadSponsorGroupId to be an explicit nil

### UnsetLeadSponsorGroupId
`func (o *SiteTrialsRow) UnsetLeadSponsorGroupId()`

UnsetLeadSponsorGroupId ensures that no value is present for LeadSponsorGroupId, not even an explicit nil
### GetLeadSponsorGroupName

`func (o *SiteTrialsRow) GetLeadSponsorGroupName() string`

GetLeadSponsorGroupName returns the LeadSponsorGroupName field if non-nil, zero value otherwise.

### GetLeadSponsorGroupNameOk

`func (o *SiteTrialsRow) GetLeadSponsorGroupNameOk() (*string, bool)`

GetLeadSponsorGroupNameOk returns a tuple with the LeadSponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorGroupName

`func (o *SiteTrialsRow) SetLeadSponsorGroupName(v string)`

SetLeadSponsorGroupName sets LeadSponsorGroupName field to given value.

### HasLeadSponsorGroupName

`func (o *SiteTrialsRow) HasLeadSponsorGroupName() bool`

HasLeadSponsorGroupName returns a boolean if a field has been set.

### SetLeadSponsorGroupNameNil

`func (o *SiteTrialsRow) SetLeadSponsorGroupNameNil(b bool)`

 SetLeadSponsorGroupNameNil sets the value for LeadSponsorGroupName to be an explicit nil

### UnsetLeadSponsorGroupName
`func (o *SiteTrialsRow) UnsetLeadSponsorGroupName()`

UnsetLeadSponsorGroupName ensures that no value is present for LeadSponsorGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


