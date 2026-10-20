# TrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**BriefTitle** | Pointer to **NullableString** | The trial&#39;s brief title. | [optional] 
**OfficialTitle** | Pointer to **NullableString** | The trial&#39;s official title. | [optional] 
**StudyType** | Pointer to **NullableString** | INTERVENTIONAL, OBSERVATIONAL or EXPANDED_ACCESS. | [optional] 
**OverallStatus** | Pointer to **NullableString** | The registry&#39;s overall status code (RECRUITING, COMPLETED, TERMINATED and so on). | [optional] 
**OverallStatusSummary** | Pointer to **NullableString** | A coarse status label derived from the overall status (Recruiting, Completed, Discontinued and so on). | [optional] 
**Phase** | Pointer to **NullableString** | The registry&#39;s phase code (PHASE1, PHASE2, PHASE2/PHASE3, NA and so on). Null when the registry gives no phase (observational and expanded-access trials); phase_label carries the display label for those. | [optional] 
**PhaseLabel** | Pointer to **NullableString** | The trial phase as a label (Phase 2, Phase 3, Observational, Not applicable). | [optional] 
**PhaseByStatus** | Pointer to **NullableString** | The phase code and the status summary joined with a colon (PHASE2:Completed). | [optional] 
**Enrollment** | Pointer to **NullableInt32** | Enrollment as registered: a planned target or an achieved count (see enrollment_type). | [optional] 
**EnrollmentType** | Pointer to **NullableString** | ACTUAL (a count of enrolled participants) or ESTIMATED (a planned target). | [optional] 
**EnrollmentIsPlausible** | Pointer to **NullableBool** | False when the registered enrollment is a placeholder value rather than a count. | [optional] 
**EnrollmentFlagReason** | Pointer to **NullableString** | Why the registered enrollment was marked a placeholder (null when it was not). | [optional] 
**StartDate** | Pointer to **NullableString** | Trial start date. | [optional] 
**StartYear** | Pointer to **NullableInt32** | Year of the trial start. | [optional] 
**PrimaryCompletionDate** | Pointer to **NullableString** | Primary completion date. | [optional] 
**PrimaryCompletionDateType** | Pointer to **NullableString** | ACTUAL (a date that has occurred) or ESTIMATED (a planned date). The registry&#39;s older ANTICIPATED is published as ESTIMATED. | [optional] 
**CompletionDate** | Pointer to **NullableString** | Completion date. | [optional] 
**CompletionDateType** | Pointer to **NullableString** | ACTUAL (a date that has occurred) or ESTIMATED (a planned date). The registry&#39;s older ANTICIPATED is published as ESTIMATED. | [optional] 
**StudyFirstPostedDate** | Pointer to **NullableString** | Date the study was first posted to the registry. | [optional] 
**LastUpdatePostedDate** | Pointer to **NullableString** | Date the registry record was last updated. | [optional] 
**WhyStopped** | Pointer to **NullableString** | The sponsor&#39;s stated reason a trial was stopped early, as filed. | [optional] 
**NumberOfArms** | Pointer to **NullableInt32** | Number of study arms. | [optional] 
**WereResultsReported** | Pointer to **NullableBool** | Whether results are posted to the registry. | [optional] 
**Gender** | Pointer to **NullableString** | Eligible sex: ALL, FEMALE or MALE. | [optional] 
**MinimumAge** | Pointer to **NullableString** | Minimum eligible age, as filed (for example 18 Years). | [optional] 
**MaximumAge** | Pointer to **NullableString** | Maximum eligible age, as filed (for example 65 Years). | [optional] 
**HealthyVolunteers** | Pointer to **NullableBool** | Whether healthy volunteers are accepted. | [optional] 
**EligibilityCriteria** | Pointer to **NullableString** | The inclusion and exclusion criteria text, as filed. | [optional] 
**LeadSponsorName** | Pointer to **NullableString** | The lead sponsor name as filed. | [optional] 
**LeadSponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the lead sponsor resolved to (null when unresolved). | [optional] 
**LeadSponsorCompanyId** | Pointer to **NullableString** | Company level: the company the lead sponsor belongs to (null when unresolved). | [optional] 
**LeadSponsorCompanyName** | Pointer to **NullableString** | Display name of the lead sponsor&#39;s company. | [optional] 
**LeadSponsorGroupId** | Pointer to **NullableString** | Group level: the lead sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**LeadSponsorGroupName** | Pointer to **NullableString** | Display name of the lead sponsor&#39;s group. | [optional] 
**Collaborators** | Pointer to **[]string** | Collaborating sponsor names, as filed. | [optional] 
**ParticipatingSponsors** | Pointer to **[]string** | Names of the lead sponsor and the collaborators, as filed. | [optional] 
**Conditions** | Pointer to **[]string** | Conditions studied, as registered. | [optional] 
**ConditionMesh** | Pointer to **[]string** | Conditions studied, as Medical Subject Headings (MeSH) terms the registry assigns. | [optional] 
**Interventions** | Pointer to **[]string** | Interventions, as registered. | [optional] 
**InterventionMesh** | Pointer to **[]string** | Interventions as Medical Subject Headings (MeSH) terms the registry assigns. | [optional] 
**InterventionTypes** | Pointer to **[]string** | Intervention types (DRUG, BIOLOGICAL, DEVICE and so on). | [optional] 
**NumberOfFacilities** | Pointer to **NullableInt32** | Number of study sites the registry record states. | [optional] 
**FacilityCount** | Pointer to **NullableInt32** | Number of study sites listed on the record. | [optional] 
**HasUsFacility** | Pointer to **NullableBool** | Whether any listed site is in the United States. | [optional] 
**FacilityCountryCount** | Pointer to **NullableInt32** | Number of countries with a listed site. | [optional] 
**FacilityCountries** | Pointer to **[]string** | Countries with a listed site. | [optional] 
**ConditionIds** | Pointer to **[]string** | MeSH condition ids the trial is linked to, direct and through the MeSH tree. Filter on one id to match the trials that carry it. | [optional] 
**DrugIds** | Pointer to **[]string** | Drug ids the trial&#39;s interventions resolved to. Filter on one id to match the trials that carry it. | [optional] 
**SponsorOrgIds** | Pointer to **[]string** | Organization record ids of the sponsors on the trial, lead and collaborator. Filter on one id to match the trials that carry it. | [optional] 

## Methods

### NewTrialsRow

`func NewTrialsRow() *TrialsRow`

NewTrialsRow instantiates a new TrialsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialsRowWithDefaults

`func NewTrialsRowWithDefaults() *TrialsRow`

NewTrialsRowWithDefaults instantiates a new TrialsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *TrialsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TrialsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TrialsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TrialsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TrialsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TrialsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetBriefTitle

`func (o *TrialsRow) GetBriefTitle() string`

GetBriefTitle returns the BriefTitle field if non-nil, zero value otherwise.

### GetBriefTitleOk

`func (o *TrialsRow) GetBriefTitleOk() (*string, bool)`

GetBriefTitleOk returns a tuple with the BriefTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBriefTitle

`func (o *TrialsRow) SetBriefTitle(v string)`

SetBriefTitle sets BriefTitle field to given value.

### HasBriefTitle

`func (o *TrialsRow) HasBriefTitle() bool`

HasBriefTitle returns a boolean if a field has been set.

### SetBriefTitleNil

`func (o *TrialsRow) SetBriefTitleNil(b bool)`

 SetBriefTitleNil sets the value for BriefTitle to be an explicit nil

### UnsetBriefTitle
`func (o *TrialsRow) UnsetBriefTitle()`

UnsetBriefTitle ensures that no value is present for BriefTitle, not even an explicit nil
### GetOfficialTitle

`func (o *TrialsRow) GetOfficialTitle() string`

GetOfficialTitle returns the OfficialTitle field if non-nil, zero value otherwise.

### GetOfficialTitleOk

`func (o *TrialsRow) GetOfficialTitleOk() (*string, bool)`

GetOfficialTitleOk returns a tuple with the OfficialTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialTitle

`func (o *TrialsRow) SetOfficialTitle(v string)`

SetOfficialTitle sets OfficialTitle field to given value.

### HasOfficialTitle

`func (o *TrialsRow) HasOfficialTitle() bool`

HasOfficialTitle returns a boolean if a field has been set.

### SetOfficialTitleNil

`func (o *TrialsRow) SetOfficialTitleNil(b bool)`

 SetOfficialTitleNil sets the value for OfficialTitle to be an explicit nil

### UnsetOfficialTitle
`func (o *TrialsRow) UnsetOfficialTitle()`

UnsetOfficialTitle ensures that no value is present for OfficialTitle, not even an explicit nil
### GetStudyType

`func (o *TrialsRow) GetStudyType() string`

GetStudyType returns the StudyType field if non-nil, zero value otherwise.

### GetStudyTypeOk

`func (o *TrialsRow) GetStudyTypeOk() (*string, bool)`

GetStudyTypeOk returns a tuple with the StudyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyType

`func (o *TrialsRow) SetStudyType(v string)`

SetStudyType sets StudyType field to given value.

### HasStudyType

`func (o *TrialsRow) HasStudyType() bool`

HasStudyType returns a boolean if a field has been set.

### SetStudyTypeNil

`func (o *TrialsRow) SetStudyTypeNil(b bool)`

 SetStudyTypeNil sets the value for StudyType to be an explicit nil

### UnsetStudyType
`func (o *TrialsRow) UnsetStudyType()`

UnsetStudyType ensures that no value is present for StudyType, not even an explicit nil
### GetOverallStatus

`func (o *TrialsRow) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *TrialsRow) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *TrialsRow) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *TrialsRow) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### SetOverallStatusNil

`func (o *TrialsRow) SetOverallStatusNil(b bool)`

 SetOverallStatusNil sets the value for OverallStatus to be an explicit nil

### UnsetOverallStatus
`func (o *TrialsRow) UnsetOverallStatus()`

UnsetOverallStatus ensures that no value is present for OverallStatus, not even an explicit nil
### GetOverallStatusSummary

`func (o *TrialsRow) GetOverallStatusSummary() string`

GetOverallStatusSummary returns the OverallStatusSummary field if non-nil, zero value otherwise.

### GetOverallStatusSummaryOk

`func (o *TrialsRow) GetOverallStatusSummaryOk() (*string, bool)`

GetOverallStatusSummaryOk returns a tuple with the OverallStatusSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatusSummary

`func (o *TrialsRow) SetOverallStatusSummary(v string)`

SetOverallStatusSummary sets OverallStatusSummary field to given value.

### HasOverallStatusSummary

`func (o *TrialsRow) HasOverallStatusSummary() bool`

HasOverallStatusSummary returns a boolean if a field has been set.

### SetOverallStatusSummaryNil

`func (o *TrialsRow) SetOverallStatusSummaryNil(b bool)`

 SetOverallStatusSummaryNil sets the value for OverallStatusSummary to be an explicit nil

### UnsetOverallStatusSummary
`func (o *TrialsRow) UnsetOverallStatusSummary()`

UnsetOverallStatusSummary ensures that no value is present for OverallStatusSummary, not even an explicit nil
### GetPhase

`func (o *TrialsRow) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *TrialsRow) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *TrialsRow) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *TrialsRow) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *TrialsRow) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *TrialsRow) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetPhaseLabel

`func (o *TrialsRow) GetPhaseLabel() string`

GetPhaseLabel returns the PhaseLabel field if non-nil, zero value otherwise.

### GetPhaseLabelOk

`func (o *TrialsRow) GetPhaseLabelOk() (*string, bool)`

GetPhaseLabelOk returns a tuple with the PhaseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseLabel

`func (o *TrialsRow) SetPhaseLabel(v string)`

SetPhaseLabel sets PhaseLabel field to given value.

### HasPhaseLabel

`func (o *TrialsRow) HasPhaseLabel() bool`

HasPhaseLabel returns a boolean if a field has been set.

### SetPhaseLabelNil

`func (o *TrialsRow) SetPhaseLabelNil(b bool)`

 SetPhaseLabelNil sets the value for PhaseLabel to be an explicit nil

### UnsetPhaseLabel
`func (o *TrialsRow) UnsetPhaseLabel()`

UnsetPhaseLabel ensures that no value is present for PhaseLabel, not even an explicit nil
### GetPhaseByStatus

`func (o *TrialsRow) GetPhaseByStatus() string`

GetPhaseByStatus returns the PhaseByStatus field if non-nil, zero value otherwise.

### GetPhaseByStatusOk

`func (o *TrialsRow) GetPhaseByStatusOk() (*string, bool)`

GetPhaseByStatusOk returns a tuple with the PhaseByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseByStatus

`func (o *TrialsRow) SetPhaseByStatus(v string)`

SetPhaseByStatus sets PhaseByStatus field to given value.

### HasPhaseByStatus

`func (o *TrialsRow) HasPhaseByStatus() bool`

HasPhaseByStatus returns a boolean if a field has been set.

### SetPhaseByStatusNil

`func (o *TrialsRow) SetPhaseByStatusNil(b bool)`

 SetPhaseByStatusNil sets the value for PhaseByStatus to be an explicit nil

### UnsetPhaseByStatus
`func (o *TrialsRow) UnsetPhaseByStatus()`

UnsetPhaseByStatus ensures that no value is present for PhaseByStatus, not even an explicit nil
### GetEnrollment

`func (o *TrialsRow) GetEnrollment() int32`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *TrialsRow) GetEnrollmentOk() (*int32, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *TrialsRow) SetEnrollment(v int32)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *TrialsRow) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### SetEnrollmentNil

`func (o *TrialsRow) SetEnrollmentNil(b bool)`

 SetEnrollmentNil sets the value for Enrollment to be an explicit nil

### UnsetEnrollment
`func (o *TrialsRow) UnsetEnrollment()`

UnsetEnrollment ensures that no value is present for Enrollment, not even an explicit nil
### GetEnrollmentType

`func (o *TrialsRow) GetEnrollmentType() string`

GetEnrollmentType returns the EnrollmentType field if non-nil, zero value otherwise.

### GetEnrollmentTypeOk

`func (o *TrialsRow) GetEnrollmentTypeOk() (*string, bool)`

GetEnrollmentTypeOk returns a tuple with the EnrollmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentType

`func (o *TrialsRow) SetEnrollmentType(v string)`

SetEnrollmentType sets EnrollmentType field to given value.

### HasEnrollmentType

`func (o *TrialsRow) HasEnrollmentType() bool`

HasEnrollmentType returns a boolean if a field has been set.

### SetEnrollmentTypeNil

`func (o *TrialsRow) SetEnrollmentTypeNil(b bool)`

 SetEnrollmentTypeNil sets the value for EnrollmentType to be an explicit nil

### UnsetEnrollmentType
`func (o *TrialsRow) UnsetEnrollmentType()`

UnsetEnrollmentType ensures that no value is present for EnrollmentType, not even an explicit nil
### GetEnrollmentIsPlausible

`func (o *TrialsRow) GetEnrollmentIsPlausible() bool`

GetEnrollmentIsPlausible returns the EnrollmentIsPlausible field if non-nil, zero value otherwise.

### GetEnrollmentIsPlausibleOk

`func (o *TrialsRow) GetEnrollmentIsPlausibleOk() (*bool, bool)`

GetEnrollmentIsPlausibleOk returns a tuple with the EnrollmentIsPlausible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentIsPlausible

`func (o *TrialsRow) SetEnrollmentIsPlausible(v bool)`

SetEnrollmentIsPlausible sets EnrollmentIsPlausible field to given value.

### HasEnrollmentIsPlausible

`func (o *TrialsRow) HasEnrollmentIsPlausible() bool`

HasEnrollmentIsPlausible returns a boolean if a field has been set.

### SetEnrollmentIsPlausibleNil

`func (o *TrialsRow) SetEnrollmentIsPlausibleNil(b bool)`

 SetEnrollmentIsPlausibleNil sets the value for EnrollmentIsPlausible to be an explicit nil

### UnsetEnrollmentIsPlausible
`func (o *TrialsRow) UnsetEnrollmentIsPlausible()`

UnsetEnrollmentIsPlausible ensures that no value is present for EnrollmentIsPlausible, not even an explicit nil
### GetEnrollmentFlagReason

`func (o *TrialsRow) GetEnrollmentFlagReason() string`

GetEnrollmentFlagReason returns the EnrollmentFlagReason field if non-nil, zero value otherwise.

### GetEnrollmentFlagReasonOk

`func (o *TrialsRow) GetEnrollmentFlagReasonOk() (*string, bool)`

GetEnrollmentFlagReasonOk returns a tuple with the EnrollmentFlagReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlagReason

`func (o *TrialsRow) SetEnrollmentFlagReason(v string)`

SetEnrollmentFlagReason sets EnrollmentFlagReason field to given value.

### HasEnrollmentFlagReason

`func (o *TrialsRow) HasEnrollmentFlagReason() bool`

HasEnrollmentFlagReason returns a boolean if a field has been set.

### SetEnrollmentFlagReasonNil

`func (o *TrialsRow) SetEnrollmentFlagReasonNil(b bool)`

 SetEnrollmentFlagReasonNil sets the value for EnrollmentFlagReason to be an explicit nil

### UnsetEnrollmentFlagReason
`func (o *TrialsRow) UnsetEnrollmentFlagReason()`

UnsetEnrollmentFlagReason ensures that no value is present for EnrollmentFlagReason, not even an explicit nil
### GetStartDate

`func (o *TrialsRow) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TrialsRow) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TrialsRow) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TrialsRow) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *TrialsRow) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *TrialsRow) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetStartYear

`func (o *TrialsRow) GetStartYear() int32`

GetStartYear returns the StartYear field if non-nil, zero value otherwise.

### GetStartYearOk

`func (o *TrialsRow) GetStartYearOk() (*int32, bool)`

GetStartYearOk returns a tuple with the StartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartYear

`func (o *TrialsRow) SetStartYear(v int32)`

SetStartYear sets StartYear field to given value.

### HasStartYear

`func (o *TrialsRow) HasStartYear() bool`

HasStartYear returns a boolean if a field has been set.

### SetStartYearNil

`func (o *TrialsRow) SetStartYearNil(b bool)`

 SetStartYearNil sets the value for StartYear to be an explicit nil

### UnsetStartYear
`func (o *TrialsRow) UnsetStartYear()`

UnsetStartYear ensures that no value is present for StartYear, not even an explicit nil
### GetPrimaryCompletionDate

`func (o *TrialsRow) GetPrimaryCompletionDate() string`

GetPrimaryCompletionDate returns the PrimaryCompletionDate field if non-nil, zero value otherwise.

### GetPrimaryCompletionDateOk

`func (o *TrialsRow) GetPrimaryCompletionDateOk() (*string, bool)`

GetPrimaryCompletionDateOk returns a tuple with the PrimaryCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCompletionDate

`func (o *TrialsRow) SetPrimaryCompletionDate(v string)`

SetPrimaryCompletionDate sets PrimaryCompletionDate field to given value.

### HasPrimaryCompletionDate

`func (o *TrialsRow) HasPrimaryCompletionDate() bool`

HasPrimaryCompletionDate returns a boolean if a field has been set.

### SetPrimaryCompletionDateNil

`func (o *TrialsRow) SetPrimaryCompletionDateNil(b bool)`

 SetPrimaryCompletionDateNil sets the value for PrimaryCompletionDate to be an explicit nil

### UnsetPrimaryCompletionDate
`func (o *TrialsRow) UnsetPrimaryCompletionDate()`

UnsetPrimaryCompletionDate ensures that no value is present for PrimaryCompletionDate, not even an explicit nil
### GetPrimaryCompletionDateType

`func (o *TrialsRow) GetPrimaryCompletionDateType() string`

GetPrimaryCompletionDateType returns the PrimaryCompletionDateType field if non-nil, zero value otherwise.

### GetPrimaryCompletionDateTypeOk

`func (o *TrialsRow) GetPrimaryCompletionDateTypeOk() (*string, bool)`

GetPrimaryCompletionDateTypeOk returns a tuple with the PrimaryCompletionDateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCompletionDateType

`func (o *TrialsRow) SetPrimaryCompletionDateType(v string)`

SetPrimaryCompletionDateType sets PrimaryCompletionDateType field to given value.

### HasPrimaryCompletionDateType

`func (o *TrialsRow) HasPrimaryCompletionDateType() bool`

HasPrimaryCompletionDateType returns a boolean if a field has been set.

### SetPrimaryCompletionDateTypeNil

`func (o *TrialsRow) SetPrimaryCompletionDateTypeNil(b bool)`

 SetPrimaryCompletionDateTypeNil sets the value for PrimaryCompletionDateType to be an explicit nil

### UnsetPrimaryCompletionDateType
`func (o *TrialsRow) UnsetPrimaryCompletionDateType()`

UnsetPrimaryCompletionDateType ensures that no value is present for PrimaryCompletionDateType, not even an explicit nil
### GetCompletionDate

`func (o *TrialsRow) GetCompletionDate() string`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *TrialsRow) GetCompletionDateOk() (*string, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *TrialsRow) SetCompletionDate(v string)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *TrialsRow) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### SetCompletionDateNil

`func (o *TrialsRow) SetCompletionDateNil(b bool)`

 SetCompletionDateNil sets the value for CompletionDate to be an explicit nil

### UnsetCompletionDate
`func (o *TrialsRow) UnsetCompletionDate()`

UnsetCompletionDate ensures that no value is present for CompletionDate, not even an explicit nil
### GetCompletionDateType

`func (o *TrialsRow) GetCompletionDateType() string`

GetCompletionDateType returns the CompletionDateType field if non-nil, zero value otherwise.

### GetCompletionDateTypeOk

`func (o *TrialsRow) GetCompletionDateTypeOk() (*string, bool)`

GetCompletionDateTypeOk returns a tuple with the CompletionDateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateType

`func (o *TrialsRow) SetCompletionDateType(v string)`

SetCompletionDateType sets CompletionDateType field to given value.

### HasCompletionDateType

`func (o *TrialsRow) HasCompletionDateType() bool`

HasCompletionDateType returns a boolean if a field has been set.

### SetCompletionDateTypeNil

`func (o *TrialsRow) SetCompletionDateTypeNil(b bool)`

 SetCompletionDateTypeNil sets the value for CompletionDateType to be an explicit nil

### UnsetCompletionDateType
`func (o *TrialsRow) UnsetCompletionDateType()`

UnsetCompletionDateType ensures that no value is present for CompletionDateType, not even an explicit nil
### GetStudyFirstPostedDate

`func (o *TrialsRow) GetStudyFirstPostedDate() string`

GetStudyFirstPostedDate returns the StudyFirstPostedDate field if non-nil, zero value otherwise.

### GetStudyFirstPostedDateOk

`func (o *TrialsRow) GetStudyFirstPostedDateOk() (*string, bool)`

GetStudyFirstPostedDateOk returns a tuple with the StudyFirstPostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyFirstPostedDate

`func (o *TrialsRow) SetStudyFirstPostedDate(v string)`

SetStudyFirstPostedDate sets StudyFirstPostedDate field to given value.

### HasStudyFirstPostedDate

`func (o *TrialsRow) HasStudyFirstPostedDate() bool`

HasStudyFirstPostedDate returns a boolean if a field has been set.

### SetStudyFirstPostedDateNil

`func (o *TrialsRow) SetStudyFirstPostedDateNil(b bool)`

 SetStudyFirstPostedDateNil sets the value for StudyFirstPostedDate to be an explicit nil

### UnsetStudyFirstPostedDate
`func (o *TrialsRow) UnsetStudyFirstPostedDate()`

UnsetStudyFirstPostedDate ensures that no value is present for StudyFirstPostedDate, not even an explicit nil
### GetLastUpdatePostedDate

`func (o *TrialsRow) GetLastUpdatePostedDate() string`

GetLastUpdatePostedDate returns the LastUpdatePostedDate field if non-nil, zero value otherwise.

### GetLastUpdatePostedDateOk

`func (o *TrialsRow) GetLastUpdatePostedDateOk() (*string, bool)`

GetLastUpdatePostedDateOk returns a tuple with the LastUpdatePostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatePostedDate

`func (o *TrialsRow) SetLastUpdatePostedDate(v string)`

SetLastUpdatePostedDate sets LastUpdatePostedDate field to given value.

### HasLastUpdatePostedDate

`func (o *TrialsRow) HasLastUpdatePostedDate() bool`

HasLastUpdatePostedDate returns a boolean if a field has been set.

### SetLastUpdatePostedDateNil

`func (o *TrialsRow) SetLastUpdatePostedDateNil(b bool)`

 SetLastUpdatePostedDateNil sets the value for LastUpdatePostedDate to be an explicit nil

### UnsetLastUpdatePostedDate
`func (o *TrialsRow) UnsetLastUpdatePostedDate()`

UnsetLastUpdatePostedDate ensures that no value is present for LastUpdatePostedDate, not even an explicit nil
### GetWhyStopped

`func (o *TrialsRow) GetWhyStopped() string`

GetWhyStopped returns the WhyStopped field if non-nil, zero value otherwise.

### GetWhyStoppedOk

`func (o *TrialsRow) GetWhyStoppedOk() (*string, bool)`

GetWhyStoppedOk returns a tuple with the WhyStopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhyStopped

`func (o *TrialsRow) SetWhyStopped(v string)`

SetWhyStopped sets WhyStopped field to given value.

### HasWhyStopped

`func (o *TrialsRow) HasWhyStopped() bool`

HasWhyStopped returns a boolean if a field has been set.

### SetWhyStoppedNil

`func (o *TrialsRow) SetWhyStoppedNil(b bool)`

 SetWhyStoppedNil sets the value for WhyStopped to be an explicit nil

### UnsetWhyStopped
`func (o *TrialsRow) UnsetWhyStopped()`

UnsetWhyStopped ensures that no value is present for WhyStopped, not even an explicit nil
### GetNumberOfArms

`func (o *TrialsRow) GetNumberOfArms() int32`

GetNumberOfArms returns the NumberOfArms field if non-nil, zero value otherwise.

### GetNumberOfArmsOk

`func (o *TrialsRow) GetNumberOfArmsOk() (*int32, bool)`

GetNumberOfArmsOk returns a tuple with the NumberOfArms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfArms

`func (o *TrialsRow) SetNumberOfArms(v int32)`

SetNumberOfArms sets NumberOfArms field to given value.

### HasNumberOfArms

`func (o *TrialsRow) HasNumberOfArms() bool`

HasNumberOfArms returns a boolean if a field has been set.

### SetNumberOfArmsNil

`func (o *TrialsRow) SetNumberOfArmsNil(b bool)`

 SetNumberOfArmsNil sets the value for NumberOfArms to be an explicit nil

### UnsetNumberOfArms
`func (o *TrialsRow) UnsetNumberOfArms()`

UnsetNumberOfArms ensures that no value is present for NumberOfArms, not even an explicit nil
### GetWereResultsReported

`func (o *TrialsRow) GetWereResultsReported() bool`

GetWereResultsReported returns the WereResultsReported field if non-nil, zero value otherwise.

### GetWereResultsReportedOk

`func (o *TrialsRow) GetWereResultsReportedOk() (*bool, bool)`

GetWereResultsReportedOk returns a tuple with the WereResultsReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWereResultsReported

`func (o *TrialsRow) SetWereResultsReported(v bool)`

SetWereResultsReported sets WereResultsReported field to given value.

### HasWereResultsReported

`func (o *TrialsRow) HasWereResultsReported() bool`

HasWereResultsReported returns a boolean if a field has been set.

### SetWereResultsReportedNil

`func (o *TrialsRow) SetWereResultsReportedNil(b bool)`

 SetWereResultsReportedNil sets the value for WereResultsReported to be an explicit nil

### UnsetWereResultsReported
`func (o *TrialsRow) UnsetWereResultsReported()`

UnsetWereResultsReported ensures that no value is present for WereResultsReported, not even an explicit nil
### GetGender

`func (o *TrialsRow) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *TrialsRow) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *TrialsRow) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *TrialsRow) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *TrialsRow) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *TrialsRow) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetMinimumAge

`func (o *TrialsRow) GetMinimumAge() string`

GetMinimumAge returns the MinimumAge field if non-nil, zero value otherwise.

### GetMinimumAgeOk

`func (o *TrialsRow) GetMinimumAgeOk() (*string, bool)`

GetMinimumAgeOk returns a tuple with the MinimumAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAge

`func (o *TrialsRow) SetMinimumAge(v string)`

SetMinimumAge sets MinimumAge field to given value.

### HasMinimumAge

`func (o *TrialsRow) HasMinimumAge() bool`

HasMinimumAge returns a boolean if a field has been set.

### SetMinimumAgeNil

`func (o *TrialsRow) SetMinimumAgeNil(b bool)`

 SetMinimumAgeNil sets the value for MinimumAge to be an explicit nil

### UnsetMinimumAge
`func (o *TrialsRow) UnsetMinimumAge()`

UnsetMinimumAge ensures that no value is present for MinimumAge, not even an explicit nil
### GetMaximumAge

`func (o *TrialsRow) GetMaximumAge() string`

GetMaximumAge returns the MaximumAge field if non-nil, zero value otherwise.

### GetMaximumAgeOk

`func (o *TrialsRow) GetMaximumAgeOk() (*string, bool)`

GetMaximumAgeOk returns a tuple with the MaximumAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAge

`func (o *TrialsRow) SetMaximumAge(v string)`

SetMaximumAge sets MaximumAge field to given value.

### HasMaximumAge

`func (o *TrialsRow) HasMaximumAge() bool`

HasMaximumAge returns a boolean if a field has been set.

### SetMaximumAgeNil

`func (o *TrialsRow) SetMaximumAgeNil(b bool)`

 SetMaximumAgeNil sets the value for MaximumAge to be an explicit nil

### UnsetMaximumAge
`func (o *TrialsRow) UnsetMaximumAge()`

UnsetMaximumAge ensures that no value is present for MaximumAge, not even an explicit nil
### GetHealthyVolunteers

`func (o *TrialsRow) GetHealthyVolunteers() bool`

GetHealthyVolunteers returns the HealthyVolunteers field if non-nil, zero value otherwise.

### GetHealthyVolunteersOk

`func (o *TrialsRow) GetHealthyVolunteersOk() (*bool, bool)`

GetHealthyVolunteersOk returns a tuple with the HealthyVolunteers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyVolunteers

`func (o *TrialsRow) SetHealthyVolunteers(v bool)`

SetHealthyVolunteers sets HealthyVolunteers field to given value.

### HasHealthyVolunteers

`func (o *TrialsRow) HasHealthyVolunteers() bool`

HasHealthyVolunteers returns a boolean if a field has been set.

### SetHealthyVolunteersNil

`func (o *TrialsRow) SetHealthyVolunteersNil(b bool)`

 SetHealthyVolunteersNil sets the value for HealthyVolunteers to be an explicit nil

### UnsetHealthyVolunteers
`func (o *TrialsRow) UnsetHealthyVolunteers()`

UnsetHealthyVolunteers ensures that no value is present for HealthyVolunteers, not even an explicit nil
### GetEligibilityCriteria

`func (o *TrialsRow) GetEligibilityCriteria() string`

GetEligibilityCriteria returns the EligibilityCriteria field if non-nil, zero value otherwise.

### GetEligibilityCriteriaOk

`func (o *TrialsRow) GetEligibilityCriteriaOk() (*string, bool)`

GetEligibilityCriteriaOk returns a tuple with the EligibilityCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibilityCriteria

`func (o *TrialsRow) SetEligibilityCriteria(v string)`

SetEligibilityCriteria sets EligibilityCriteria field to given value.

### HasEligibilityCriteria

`func (o *TrialsRow) HasEligibilityCriteria() bool`

HasEligibilityCriteria returns a boolean if a field has been set.

### SetEligibilityCriteriaNil

`func (o *TrialsRow) SetEligibilityCriteriaNil(b bool)`

 SetEligibilityCriteriaNil sets the value for EligibilityCriteria to be an explicit nil

### UnsetEligibilityCriteria
`func (o *TrialsRow) UnsetEligibilityCriteria()`

UnsetEligibilityCriteria ensures that no value is present for EligibilityCriteria, not even an explicit nil
### GetLeadSponsorName

`func (o *TrialsRow) GetLeadSponsorName() string`

GetLeadSponsorName returns the LeadSponsorName field if non-nil, zero value otherwise.

### GetLeadSponsorNameOk

`func (o *TrialsRow) GetLeadSponsorNameOk() (*string, bool)`

GetLeadSponsorNameOk returns a tuple with the LeadSponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorName

`func (o *TrialsRow) SetLeadSponsorName(v string)`

SetLeadSponsorName sets LeadSponsorName field to given value.

### HasLeadSponsorName

`func (o *TrialsRow) HasLeadSponsorName() bool`

HasLeadSponsorName returns a boolean if a field has been set.

### SetLeadSponsorNameNil

`func (o *TrialsRow) SetLeadSponsorNameNil(b bool)`

 SetLeadSponsorNameNil sets the value for LeadSponsorName to be an explicit nil

### UnsetLeadSponsorName
`func (o *TrialsRow) UnsetLeadSponsorName()`

UnsetLeadSponsorName ensures that no value is present for LeadSponsorName, not even an explicit nil
### GetLeadSponsorOrgId

`func (o *TrialsRow) GetLeadSponsorOrgId() string`

GetLeadSponsorOrgId returns the LeadSponsorOrgId field if non-nil, zero value otherwise.

### GetLeadSponsorOrgIdOk

`func (o *TrialsRow) GetLeadSponsorOrgIdOk() (*string, bool)`

GetLeadSponsorOrgIdOk returns a tuple with the LeadSponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorOrgId

`func (o *TrialsRow) SetLeadSponsorOrgId(v string)`

SetLeadSponsorOrgId sets LeadSponsorOrgId field to given value.

### HasLeadSponsorOrgId

`func (o *TrialsRow) HasLeadSponsorOrgId() bool`

HasLeadSponsorOrgId returns a boolean if a field has been set.

### SetLeadSponsorOrgIdNil

`func (o *TrialsRow) SetLeadSponsorOrgIdNil(b bool)`

 SetLeadSponsorOrgIdNil sets the value for LeadSponsorOrgId to be an explicit nil

### UnsetLeadSponsorOrgId
`func (o *TrialsRow) UnsetLeadSponsorOrgId()`

UnsetLeadSponsorOrgId ensures that no value is present for LeadSponsorOrgId, not even an explicit nil
### GetLeadSponsorCompanyId

`func (o *TrialsRow) GetLeadSponsorCompanyId() string`

GetLeadSponsorCompanyId returns the LeadSponsorCompanyId field if non-nil, zero value otherwise.

### GetLeadSponsorCompanyIdOk

`func (o *TrialsRow) GetLeadSponsorCompanyIdOk() (*string, bool)`

GetLeadSponsorCompanyIdOk returns a tuple with the LeadSponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorCompanyId

`func (o *TrialsRow) SetLeadSponsorCompanyId(v string)`

SetLeadSponsorCompanyId sets LeadSponsorCompanyId field to given value.

### HasLeadSponsorCompanyId

`func (o *TrialsRow) HasLeadSponsorCompanyId() bool`

HasLeadSponsorCompanyId returns a boolean if a field has been set.

### SetLeadSponsorCompanyIdNil

`func (o *TrialsRow) SetLeadSponsorCompanyIdNil(b bool)`

 SetLeadSponsorCompanyIdNil sets the value for LeadSponsorCompanyId to be an explicit nil

### UnsetLeadSponsorCompanyId
`func (o *TrialsRow) UnsetLeadSponsorCompanyId()`

UnsetLeadSponsorCompanyId ensures that no value is present for LeadSponsorCompanyId, not even an explicit nil
### GetLeadSponsorCompanyName

`func (o *TrialsRow) GetLeadSponsorCompanyName() string`

GetLeadSponsorCompanyName returns the LeadSponsorCompanyName field if non-nil, zero value otherwise.

### GetLeadSponsorCompanyNameOk

`func (o *TrialsRow) GetLeadSponsorCompanyNameOk() (*string, bool)`

GetLeadSponsorCompanyNameOk returns a tuple with the LeadSponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorCompanyName

`func (o *TrialsRow) SetLeadSponsorCompanyName(v string)`

SetLeadSponsorCompanyName sets LeadSponsorCompanyName field to given value.

### HasLeadSponsorCompanyName

`func (o *TrialsRow) HasLeadSponsorCompanyName() bool`

HasLeadSponsorCompanyName returns a boolean if a field has been set.

### SetLeadSponsorCompanyNameNil

`func (o *TrialsRow) SetLeadSponsorCompanyNameNil(b bool)`

 SetLeadSponsorCompanyNameNil sets the value for LeadSponsorCompanyName to be an explicit nil

### UnsetLeadSponsorCompanyName
`func (o *TrialsRow) UnsetLeadSponsorCompanyName()`

UnsetLeadSponsorCompanyName ensures that no value is present for LeadSponsorCompanyName, not even an explicit nil
### GetLeadSponsorGroupId

`func (o *TrialsRow) GetLeadSponsorGroupId() string`

GetLeadSponsorGroupId returns the LeadSponsorGroupId field if non-nil, zero value otherwise.

### GetLeadSponsorGroupIdOk

`func (o *TrialsRow) GetLeadSponsorGroupIdOk() (*string, bool)`

GetLeadSponsorGroupIdOk returns a tuple with the LeadSponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorGroupId

`func (o *TrialsRow) SetLeadSponsorGroupId(v string)`

SetLeadSponsorGroupId sets LeadSponsorGroupId field to given value.

### HasLeadSponsorGroupId

`func (o *TrialsRow) HasLeadSponsorGroupId() bool`

HasLeadSponsorGroupId returns a boolean if a field has been set.

### SetLeadSponsorGroupIdNil

`func (o *TrialsRow) SetLeadSponsorGroupIdNil(b bool)`

 SetLeadSponsorGroupIdNil sets the value for LeadSponsorGroupId to be an explicit nil

### UnsetLeadSponsorGroupId
`func (o *TrialsRow) UnsetLeadSponsorGroupId()`

UnsetLeadSponsorGroupId ensures that no value is present for LeadSponsorGroupId, not even an explicit nil
### GetLeadSponsorGroupName

`func (o *TrialsRow) GetLeadSponsorGroupName() string`

GetLeadSponsorGroupName returns the LeadSponsorGroupName field if non-nil, zero value otherwise.

### GetLeadSponsorGroupNameOk

`func (o *TrialsRow) GetLeadSponsorGroupNameOk() (*string, bool)`

GetLeadSponsorGroupNameOk returns a tuple with the LeadSponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorGroupName

`func (o *TrialsRow) SetLeadSponsorGroupName(v string)`

SetLeadSponsorGroupName sets LeadSponsorGroupName field to given value.

### HasLeadSponsorGroupName

`func (o *TrialsRow) HasLeadSponsorGroupName() bool`

HasLeadSponsorGroupName returns a boolean if a field has been set.

### SetLeadSponsorGroupNameNil

`func (o *TrialsRow) SetLeadSponsorGroupNameNil(b bool)`

 SetLeadSponsorGroupNameNil sets the value for LeadSponsorGroupName to be an explicit nil

### UnsetLeadSponsorGroupName
`func (o *TrialsRow) UnsetLeadSponsorGroupName()`

UnsetLeadSponsorGroupName ensures that no value is present for LeadSponsorGroupName, not even an explicit nil
### GetCollaborators

`func (o *TrialsRow) GetCollaborators() []string`

GetCollaborators returns the Collaborators field if non-nil, zero value otherwise.

### GetCollaboratorsOk

`func (o *TrialsRow) GetCollaboratorsOk() (*[]string, bool)`

GetCollaboratorsOk returns a tuple with the Collaborators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborators

`func (o *TrialsRow) SetCollaborators(v []string)`

SetCollaborators sets Collaborators field to given value.

### HasCollaborators

`func (o *TrialsRow) HasCollaborators() bool`

HasCollaborators returns a boolean if a field has been set.

### SetCollaboratorsNil

`func (o *TrialsRow) SetCollaboratorsNil(b bool)`

 SetCollaboratorsNil sets the value for Collaborators to be an explicit nil

### UnsetCollaborators
`func (o *TrialsRow) UnsetCollaborators()`

UnsetCollaborators ensures that no value is present for Collaborators, not even an explicit nil
### GetParticipatingSponsors

`func (o *TrialsRow) GetParticipatingSponsors() []string`

GetParticipatingSponsors returns the ParticipatingSponsors field if non-nil, zero value otherwise.

### GetParticipatingSponsorsOk

`func (o *TrialsRow) GetParticipatingSponsorsOk() (*[]string, bool)`

GetParticipatingSponsorsOk returns a tuple with the ParticipatingSponsors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipatingSponsors

`func (o *TrialsRow) SetParticipatingSponsors(v []string)`

SetParticipatingSponsors sets ParticipatingSponsors field to given value.

### HasParticipatingSponsors

`func (o *TrialsRow) HasParticipatingSponsors() bool`

HasParticipatingSponsors returns a boolean if a field has been set.

### SetParticipatingSponsorsNil

`func (o *TrialsRow) SetParticipatingSponsorsNil(b bool)`

 SetParticipatingSponsorsNil sets the value for ParticipatingSponsors to be an explicit nil

### UnsetParticipatingSponsors
`func (o *TrialsRow) UnsetParticipatingSponsors()`

UnsetParticipatingSponsors ensures that no value is present for ParticipatingSponsors, not even an explicit nil
### GetConditions

`func (o *TrialsRow) GetConditions() []string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *TrialsRow) GetConditionsOk() (*[]string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *TrialsRow) SetConditions(v []string)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *TrialsRow) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *TrialsRow) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *TrialsRow) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetConditionMesh

`func (o *TrialsRow) GetConditionMesh() []string`

GetConditionMesh returns the ConditionMesh field if non-nil, zero value otherwise.

### GetConditionMeshOk

`func (o *TrialsRow) GetConditionMeshOk() (*[]string, bool)`

GetConditionMeshOk returns a tuple with the ConditionMesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionMesh

`func (o *TrialsRow) SetConditionMesh(v []string)`

SetConditionMesh sets ConditionMesh field to given value.

### HasConditionMesh

`func (o *TrialsRow) HasConditionMesh() bool`

HasConditionMesh returns a boolean if a field has been set.

### SetConditionMeshNil

`func (o *TrialsRow) SetConditionMeshNil(b bool)`

 SetConditionMeshNil sets the value for ConditionMesh to be an explicit nil

### UnsetConditionMesh
`func (o *TrialsRow) UnsetConditionMesh()`

UnsetConditionMesh ensures that no value is present for ConditionMesh, not even an explicit nil
### GetInterventions

`func (o *TrialsRow) GetInterventions() []string`

GetInterventions returns the Interventions field if non-nil, zero value otherwise.

### GetInterventionsOk

`func (o *TrialsRow) GetInterventionsOk() (*[]string, bool)`

GetInterventionsOk returns a tuple with the Interventions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterventions

`func (o *TrialsRow) SetInterventions(v []string)`

SetInterventions sets Interventions field to given value.

### HasInterventions

`func (o *TrialsRow) HasInterventions() bool`

HasInterventions returns a boolean if a field has been set.

### SetInterventionsNil

`func (o *TrialsRow) SetInterventionsNil(b bool)`

 SetInterventionsNil sets the value for Interventions to be an explicit nil

### UnsetInterventions
`func (o *TrialsRow) UnsetInterventions()`

UnsetInterventions ensures that no value is present for Interventions, not even an explicit nil
### GetInterventionMesh

`func (o *TrialsRow) GetInterventionMesh() []string`

GetInterventionMesh returns the InterventionMesh field if non-nil, zero value otherwise.

### GetInterventionMeshOk

`func (o *TrialsRow) GetInterventionMeshOk() (*[]string, bool)`

GetInterventionMeshOk returns a tuple with the InterventionMesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterventionMesh

`func (o *TrialsRow) SetInterventionMesh(v []string)`

SetInterventionMesh sets InterventionMesh field to given value.

### HasInterventionMesh

`func (o *TrialsRow) HasInterventionMesh() bool`

HasInterventionMesh returns a boolean if a field has been set.

### SetInterventionMeshNil

`func (o *TrialsRow) SetInterventionMeshNil(b bool)`

 SetInterventionMeshNil sets the value for InterventionMesh to be an explicit nil

### UnsetInterventionMesh
`func (o *TrialsRow) UnsetInterventionMesh()`

UnsetInterventionMesh ensures that no value is present for InterventionMesh, not even an explicit nil
### GetInterventionTypes

`func (o *TrialsRow) GetInterventionTypes() []string`

GetInterventionTypes returns the InterventionTypes field if non-nil, zero value otherwise.

### GetInterventionTypesOk

`func (o *TrialsRow) GetInterventionTypesOk() (*[]string, bool)`

GetInterventionTypesOk returns a tuple with the InterventionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterventionTypes

`func (o *TrialsRow) SetInterventionTypes(v []string)`

SetInterventionTypes sets InterventionTypes field to given value.

### HasInterventionTypes

`func (o *TrialsRow) HasInterventionTypes() bool`

HasInterventionTypes returns a boolean if a field has been set.

### SetInterventionTypesNil

`func (o *TrialsRow) SetInterventionTypesNil(b bool)`

 SetInterventionTypesNil sets the value for InterventionTypes to be an explicit nil

### UnsetInterventionTypes
`func (o *TrialsRow) UnsetInterventionTypes()`

UnsetInterventionTypes ensures that no value is present for InterventionTypes, not even an explicit nil
### GetNumberOfFacilities

`func (o *TrialsRow) GetNumberOfFacilities() int32`

GetNumberOfFacilities returns the NumberOfFacilities field if non-nil, zero value otherwise.

### GetNumberOfFacilitiesOk

`func (o *TrialsRow) GetNumberOfFacilitiesOk() (*int32, bool)`

GetNumberOfFacilitiesOk returns a tuple with the NumberOfFacilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFacilities

`func (o *TrialsRow) SetNumberOfFacilities(v int32)`

SetNumberOfFacilities sets NumberOfFacilities field to given value.

### HasNumberOfFacilities

`func (o *TrialsRow) HasNumberOfFacilities() bool`

HasNumberOfFacilities returns a boolean if a field has been set.

### SetNumberOfFacilitiesNil

`func (o *TrialsRow) SetNumberOfFacilitiesNil(b bool)`

 SetNumberOfFacilitiesNil sets the value for NumberOfFacilities to be an explicit nil

### UnsetNumberOfFacilities
`func (o *TrialsRow) UnsetNumberOfFacilities()`

UnsetNumberOfFacilities ensures that no value is present for NumberOfFacilities, not even an explicit nil
### GetFacilityCount

`func (o *TrialsRow) GetFacilityCount() int32`

GetFacilityCount returns the FacilityCount field if non-nil, zero value otherwise.

### GetFacilityCountOk

`func (o *TrialsRow) GetFacilityCountOk() (*int32, bool)`

GetFacilityCountOk returns a tuple with the FacilityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCount

`func (o *TrialsRow) SetFacilityCount(v int32)`

SetFacilityCount sets FacilityCount field to given value.

### HasFacilityCount

`func (o *TrialsRow) HasFacilityCount() bool`

HasFacilityCount returns a boolean if a field has been set.

### SetFacilityCountNil

`func (o *TrialsRow) SetFacilityCountNil(b bool)`

 SetFacilityCountNil sets the value for FacilityCount to be an explicit nil

### UnsetFacilityCount
`func (o *TrialsRow) UnsetFacilityCount()`

UnsetFacilityCount ensures that no value is present for FacilityCount, not even an explicit nil
### GetHasUsFacility

`func (o *TrialsRow) GetHasUsFacility() bool`

GetHasUsFacility returns the HasUsFacility field if non-nil, zero value otherwise.

### GetHasUsFacilityOk

`func (o *TrialsRow) GetHasUsFacilityOk() (*bool, bool)`

GetHasUsFacilityOk returns a tuple with the HasUsFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUsFacility

`func (o *TrialsRow) SetHasUsFacility(v bool)`

SetHasUsFacility sets HasUsFacility field to given value.

### HasHasUsFacility

`func (o *TrialsRow) HasHasUsFacility() bool`

HasHasUsFacility returns a boolean if a field has been set.

### SetHasUsFacilityNil

`func (o *TrialsRow) SetHasUsFacilityNil(b bool)`

 SetHasUsFacilityNil sets the value for HasUsFacility to be an explicit nil

### UnsetHasUsFacility
`func (o *TrialsRow) UnsetHasUsFacility()`

UnsetHasUsFacility ensures that no value is present for HasUsFacility, not even an explicit nil
### GetFacilityCountryCount

`func (o *TrialsRow) GetFacilityCountryCount() int32`

GetFacilityCountryCount returns the FacilityCountryCount field if non-nil, zero value otherwise.

### GetFacilityCountryCountOk

`func (o *TrialsRow) GetFacilityCountryCountOk() (*int32, bool)`

GetFacilityCountryCountOk returns a tuple with the FacilityCountryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCountryCount

`func (o *TrialsRow) SetFacilityCountryCount(v int32)`

SetFacilityCountryCount sets FacilityCountryCount field to given value.

### HasFacilityCountryCount

`func (o *TrialsRow) HasFacilityCountryCount() bool`

HasFacilityCountryCount returns a boolean if a field has been set.

### SetFacilityCountryCountNil

`func (o *TrialsRow) SetFacilityCountryCountNil(b bool)`

 SetFacilityCountryCountNil sets the value for FacilityCountryCount to be an explicit nil

### UnsetFacilityCountryCount
`func (o *TrialsRow) UnsetFacilityCountryCount()`

UnsetFacilityCountryCount ensures that no value is present for FacilityCountryCount, not even an explicit nil
### GetFacilityCountries

`func (o *TrialsRow) GetFacilityCountries() []string`

GetFacilityCountries returns the FacilityCountries field if non-nil, zero value otherwise.

### GetFacilityCountriesOk

`func (o *TrialsRow) GetFacilityCountriesOk() (*[]string, bool)`

GetFacilityCountriesOk returns a tuple with the FacilityCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCountries

`func (o *TrialsRow) SetFacilityCountries(v []string)`

SetFacilityCountries sets FacilityCountries field to given value.

### HasFacilityCountries

`func (o *TrialsRow) HasFacilityCountries() bool`

HasFacilityCountries returns a boolean if a field has been set.

### SetFacilityCountriesNil

`func (o *TrialsRow) SetFacilityCountriesNil(b bool)`

 SetFacilityCountriesNil sets the value for FacilityCountries to be an explicit nil

### UnsetFacilityCountries
`func (o *TrialsRow) UnsetFacilityCountries()`

UnsetFacilityCountries ensures that no value is present for FacilityCountries, not even an explicit nil
### GetConditionIds

`func (o *TrialsRow) GetConditionIds() []string`

GetConditionIds returns the ConditionIds field if non-nil, zero value otherwise.

### GetConditionIdsOk

`func (o *TrialsRow) GetConditionIdsOk() (*[]string, bool)`

GetConditionIdsOk returns a tuple with the ConditionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionIds

`func (o *TrialsRow) SetConditionIds(v []string)`

SetConditionIds sets ConditionIds field to given value.

### HasConditionIds

`func (o *TrialsRow) HasConditionIds() bool`

HasConditionIds returns a boolean if a field has been set.

### SetConditionIdsNil

`func (o *TrialsRow) SetConditionIdsNil(b bool)`

 SetConditionIdsNil sets the value for ConditionIds to be an explicit nil

### UnsetConditionIds
`func (o *TrialsRow) UnsetConditionIds()`

UnsetConditionIds ensures that no value is present for ConditionIds, not even an explicit nil
### GetDrugIds

`func (o *TrialsRow) GetDrugIds() []string`

GetDrugIds returns the DrugIds field if non-nil, zero value otherwise.

### GetDrugIdsOk

`func (o *TrialsRow) GetDrugIdsOk() (*[]string, bool)`

GetDrugIdsOk returns a tuple with the DrugIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugIds

`func (o *TrialsRow) SetDrugIds(v []string)`

SetDrugIds sets DrugIds field to given value.

### HasDrugIds

`func (o *TrialsRow) HasDrugIds() bool`

HasDrugIds returns a boolean if a field has been set.

### SetDrugIdsNil

`func (o *TrialsRow) SetDrugIdsNil(b bool)`

 SetDrugIdsNil sets the value for DrugIds to be an explicit nil

### UnsetDrugIds
`func (o *TrialsRow) UnsetDrugIds()`

UnsetDrugIds ensures that no value is present for DrugIds, not even an explicit nil
### GetSponsorOrgIds

`func (o *TrialsRow) GetSponsorOrgIds() []string`

GetSponsorOrgIds returns the SponsorOrgIds field if non-nil, zero value otherwise.

### GetSponsorOrgIdsOk

`func (o *TrialsRow) GetSponsorOrgIdsOk() (*[]string, bool)`

GetSponsorOrgIdsOk returns a tuple with the SponsorOrgIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgIds

`func (o *TrialsRow) SetSponsorOrgIds(v []string)`

SetSponsorOrgIds sets SponsorOrgIds field to given value.

### HasSponsorOrgIds

`func (o *TrialsRow) HasSponsorOrgIds() bool`

HasSponsorOrgIds returns a boolean if a field has been set.

### SetSponsorOrgIdsNil

`func (o *TrialsRow) SetSponsorOrgIdsNil(b bool)`

 SetSponsorOrgIdsNil sets the value for SponsorOrgIds to be an explicit nil

### UnsetSponsorOrgIds
`func (o *TrialsRow) UnsetSponsorOrgIds()`

UnsetSponsorOrgIds ensures that no value is present for SponsorOrgIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


