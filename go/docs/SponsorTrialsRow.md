# SponsorTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**SponsorRole** | Pointer to **NullableString** | lead or collaborator. | [optional] 
**SponsorName** | Pointer to **NullableString** | The sponsor name as filed. | [optional] 
**SponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the sponsor resolved to (null when unresolved). | [optional] 
**SponsorCompanyId** | Pointer to **NullableString** | Company level: the company the sponsor belongs to (null when unresolved). | [optional] 
**SponsorCompanyName** | Pointer to **NullableString** | Display name of the sponsor&#39;s company. | [optional] 
**SponsorGroupId** | Pointer to **NullableString** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SponsorGroupName** | Pointer to **NullableString** | Display name of the sponsor&#39;s group. | [optional] 
**SponsorAgencyClass** | Pointer to **NullableString** | The registry&#39;s agency class for the sponsor (INDUSTRY, NIH for the US National Institutes of Health, OTHER ...). | [optional] 
**SponsorResolutionMethod** | Pointer to **NullableString** | How the sponsor name was resolved to an organization record. | [optional] 
**TrialTitle** | Pointer to **NullableString** | The trial&#39;s brief title. | [optional] 
**StudyType** | Pointer to **NullableString** | INTERVENTIONAL or OBSERVATIONAL. | [optional] 
**Phase** | Pointer to **NullableString** | The trial phase label. | [optional] 
**OverallStatus** | Pointer to **NullableString** | The registry&#39;s overall status. | [optional] 
**StartDate** | Pointer to **NullableString** | Trial start date. | [optional] 
**StartYear** | Pointer to **NullableInt32** | Year of the trial start. | [optional] 
**PrimaryCompletionDate** | Pointer to **NullableString** | Primary completion date. | [optional] 
**Enrollment** | Pointer to **NullableInt32** | Enrollment as registered. | [optional] 
**WereResultsReported** | Pointer to **NullableBool** | Whether results are posted. | [optional] 
**Conditions** | Pointer to **[]string** | Conditions studied, as registered. | [optional] 
**Interventions** | Pointer to **[]string** | Interventions, as registered. | [optional] 

## Methods

### NewSponsorTrialsRow

`func NewSponsorTrialsRow() *SponsorTrialsRow`

NewSponsorTrialsRow instantiates a new SponsorTrialsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSponsorTrialsRowWithDefaults

`func NewSponsorTrialsRowWithDefaults() *SponsorTrialsRow`

NewSponsorTrialsRowWithDefaults instantiates a new SponsorTrialsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *SponsorTrialsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *SponsorTrialsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *SponsorTrialsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *SponsorTrialsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *SponsorTrialsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *SponsorTrialsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetSponsorRole

`func (o *SponsorTrialsRow) GetSponsorRole() string`

GetSponsorRole returns the SponsorRole field if non-nil, zero value otherwise.

### GetSponsorRoleOk

`func (o *SponsorTrialsRow) GetSponsorRoleOk() (*string, bool)`

GetSponsorRoleOk returns a tuple with the SponsorRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorRole

`func (o *SponsorTrialsRow) SetSponsorRole(v string)`

SetSponsorRole sets SponsorRole field to given value.

### HasSponsorRole

`func (o *SponsorTrialsRow) HasSponsorRole() bool`

HasSponsorRole returns a boolean if a field has been set.

### SetSponsorRoleNil

`func (o *SponsorTrialsRow) SetSponsorRoleNil(b bool)`

 SetSponsorRoleNil sets the value for SponsorRole to be an explicit nil

### UnsetSponsorRole
`func (o *SponsorTrialsRow) UnsetSponsorRole()`

UnsetSponsorRole ensures that no value is present for SponsorRole, not even an explicit nil
### GetSponsorName

`func (o *SponsorTrialsRow) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *SponsorTrialsRow) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *SponsorTrialsRow) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *SponsorTrialsRow) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### SetSponsorNameNil

`func (o *SponsorTrialsRow) SetSponsorNameNil(b bool)`

 SetSponsorNameNil sets the value for SponsorName to be an explicit nil

### UnsetSponsorName
`func (o *SponsorTrialsRow) UnsetSponsorName()`

UnsetSponsorName ensures that no value is present for SponsorName, not even an explicit nil
### GetSponsorOrgId

`func (o *SponsorTrialsRow) GetSponsorOrgId() string`

GetSponsorOrgId returns the SponsorOrgId field if non-nil, zero value otherwise.

### GetSponsorOrgIdOk

`func (o *SponsorTrialsRow) GetSponsorOrgIdOk() (*string, bool)`

GetSponsorOrgIdOk returns a tuple with the SponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgId

`func (o *SponsorTrialsRow) SetSponsorOrgId(v string)`

SetSponsorOrgId sets SponsorOrgId field to given value.

### HasSponsorOrgId

`func (o *SponsorTrialsRow) HasSponsorOrgId() bool`

HasSponsorOrgId returns a boolean if a field has been set.

### SetSponsorOrgIdNil

`func (o *SponsorTrialsRow) SetSponsorOrgIdNil(b bool)`

 SetSponsorOrgIdNil sets the value for SponsorOrgId to be an explicit nil

### UnsetSponsorOrgId
`func (o *SponsorTrialsRow) UnsetSponsorOrgId()`

UnsetSponsorOrgId ensures that no value is present for SponsorOrgId, not even an explicit nil
### GetSponsorCompanyId

`func (o *SponsorTrialsRow) GetSponsorCompanyId() string`

GetSponsorCompanyId returns the SponsorCompanyId field if non-nil, zero value otherwise.

### GetSponsorCompanyIdOk

`func (o *SponsorTrialsRow) GetSponsorCompanyIdOk() (*string, bool)`

GetSponsorCompanyIdOk returns a tuple with the SponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyId

`func (o *SponsorTrialsRow) SetSponsorCompanyId(v string)`

SetSponsorCompanyId sets SponsorCompanyId field to given value.

### HasSponsorCompanyId

`func (o *SponsorTrialsRow) HasSponsorCompanyId() bool`

HasSponsorCompanyId returns a boolean if a field has been set.

### SetSponsorCompanyIdNil

`func (o *SponsorTrialsRow) SetSponsorCompanyIdNil(b bool)`

 SetSponsorCompanyIdNil sets the value for SponsorCompanyId to be an explicit nil

### UnsetSponsorCompanyId
`func (o *SponsorTrialsRow) UnsetSponsorCompanyId()`

UnsetSponsorCompanyId ensures that no value is present for SponsorCompanyId, not even an explicit nil
### GetSponsorCompanyName

`func (o *SponsorTrialsRow) GetSponsorCompanyName() string`

GetSponsorCompanyName returns the SponsorCompanyName field if non-nil, zero value otherwise.

### GetSponsorCompanyNameOk

`func (o *SponsorTrialsRow) GetSponsorCompanyNameOk() (*string, bool)`

GetSponsorCompanyNameOk returns a tuple with the SponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyName

`func (o *SponsorTrialsRow) SetSponsorCompanyName(v string)`

SetSponsorCompanyName sets SponsorCompanyName field to given value.

### HasSponsorCompanyName

`func (o *SponsorTrialsRow) HasSponsorCompanyName() bool`

HasSponsorCompanyName returns a boolean if a field has been set.

### SetSponsorCompanyNameNil

`func (o *SponsorTrialsRow) SetSponsorCompanyNameNil(b bool)`

 SetSponsorCompanyNameNil sets the value for SponsorCompanyName to be an explicit nil

### UnsetSponsorCompanyName
`func (o *SponsorTrialsRow) UnsetSponsorCompanyName()`

UnsetSponsorCompanyName ensures that no value is present for SponsorCompanyName, not even an explicit nil
### GetSponsorGroupId

`func (o *SponsorTrialsRow) GetSponsorGroupId() string`

GetSponsorGroupId returns the SponsorGroupId field if non-nil, zero value otherwise.

### GetSponsorGroupIdOk

`func (o *SponsorTrialsRow) GetSponsorGroupIdOk() (*string, bool)`

GetSponsorGroupIdOk returns a tuple with the SponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupId

`func (o *SponsorTrialsRow) SetSponsorGroupId(v string)`

SetSponsorGroupId sets SponsorGroupId field to given value.

### HasSponsorGroupId

`func (o *SponsorTrialsRow) HasSponsorGroupId() bool`

HasSponsorGroupId returns a boolean if a field has been set.

### SetSponsorGroupIdNil

`func (o *SponsorTrialsRow) SetSponsorGroupIdNil(b bool)`

 SetSponsorGroupIdNil sets the value for SponsorGroupId to be an explicit nil

### UnsetSponsorGroupId
`func (o *SponsorTrialsRow) UnsetSponsorGroupId()`

UnsetSponsorGroupId ensures that no value is present for SponsorGroupId, not even an explicit nil
### GetSponsorGroupName

`func (o *SponsorTrialsRow) GetSponsorGroupName() string`

GetSponsorGroupName returns the SponsorGroupName field if non-nil, zero value otherwise.

### GetSponsorGroupNameOk

`func (o *SponsorTrialsRow) GetSponsorGroupNameOk() (*string, bool)`

GetSponsorGroupNameOk returns a tuple with the SponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupName

`func (o *SponsorTrialsRow) SetSponsorGroupName(v string)`

SetSponsorGroupName sets SponsorGroupName field to given value.

### HasSponsorGroupName

`func (o *SponsorTrialsRow) HasSponsorGroupName() bool`

HasSponsorGroupName returns a boolean if a field has been set.

### SetSponsorGroupNameNil

`func (o *SponsorTrialsRow) SetSponsorGroupNameNil(b bool)`

 SetSponsorGroupNameNil sets the value for SponsorGroupName to be an explicit nil

### UnsetSponsorGroupName
`func (o *SponsorTrialsRow) UnsetSponsorGroupName()`

UnsetSponsorGroupName ensures that no value is present for SponsorGroupName, not even an explicit nil
### GetSponsorAgencyClass

`func (o *SponsorTrialsRow) GetSponsorAgencyClass() string`

GetSponsorAgencyClass returns the SponsorAgencyClass field if non-nil, zero value otherwise.

### GetSponsorAgencyClassOk

`func (o *SponsorTrialsRow) GetSponsorAgencyClassOk() (*string, bool)`

GetSponsorAgencyClassOk returns a tuple with the SponsorAgencyClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorAgencyClass

`func (o *SponsorTrialsRow) SetSponsorAgencyClass(v string)`

SetSponsorAgencyClass sets SponsorAgencyClass field to given value.

### HasSponsorAgencyClass

`func (o *SponsorTrialsRow) HasSponsorAgencyClass() bool`

HasSponsorAgencyClass returns a boolean if a field has been set.

### SetSponsorAgencyClassNil

`func (o *SponsorTrialsRow) SetSponsorAgencyClassNil(b bool)`

 SetSponsorAgencyClassNil sets the value for SponsorAgencyClass to be an explicit nil

### UnsetSponsorAgencyClass
`func (o *SponsorTrialsRow) UnsetSponsorAgencyClass()`

UnsetSponsorAgencyClass ensures that no value is present for SponsorAgencyClass, not even an explicit nil
### GetSponsorResolutionMethod

`func (o *SponsorTrialsRow) GetSponsorResolutionMethod() string`

GetSponsorResolutionMethod returns the SponsorResolutionMethod field if non-nil, zero value otherwise.

### GetSponsorResolutionMethodOk

`func (o *SponsorTrialsRow) GetSponsorResolutionMethodOk() (*string, bool)`

GetSponsorResolutionMethodOk returns a tuple with the SponsorResolutionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorResolutionMethod

`func (o *SponsorTrialsRow) SetSponsorResolutionMethod(v string)`

SetSponsorResolutionMethod sets SponsorResolutionMethod field to given value.

### HasSponsorResolutionMethod

`func (o *SponsorTrialsRow) HasSponsorResolutionMethod() bool`

HasSponsorResolutionMethod returns a boolean if a field has been set.

### SetSponsorResolutionMethodNil

`func (o *SponsorTrialsRow) SetSponsorResolutionMethodNil(b bool)`

 SetSponsorResolutionMethodNil sets the value for SponsorResolutionMethod to be an explicit nil

### UnsetSponsorResolutionMethod
`func (o *SponsorTrialsRow) UnsetSponsorResolutionMethod()`

UnsetSponsorResolutionMethod ensures that no value is present for SponsorResolutionMethod, not even an explicit nil
### GetTrialTitle

`func (o *SponsorTrialsRow) GetTrialTitle() string`

GetTrialTitle returns the TrialTitle field if non-nil, zero value otherwise.

### GetTrialTitleOk

`func (o *SponsorTrialsRow) GetTrialTitleOk() (*string, bool)`

GetTrialTitleOk returns a tuple with the TrialTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialTitle

`func (o *SponsorTrialsRow) SetTrialTitle(v string)`

SetTrialTitle sets TrialTitle field to given value.

### HasTrialTitle

`func (o *SponsorTrialsRow) HasTrialTitle() bool`

HasTrialTitle returns a boolean if a field has been set.

### SetTrialTitleNil

`func (o *SponsorTrialsRow) SetTrialTitleNil(b bool)`

 SetTrialTitleNil sets the value for TrialTitle to be an explicit nil

### UnsetTrialTitle
`func (o *SponsorTrialsRow) UnsetTrialTitle()`

UnsetTrialTitle ensures that no value is present for TrialTitle, not even an explicit nil
### GetStudyType

`func (o *SponsorTrialsRow) GetStudyType() string`

GetStudyType returns the StudyType field if non-nil, zero value otherwise.

### GetStudyTypeOk

`func (o *SponsorTrialsRow) GetStudyTypeOk() (*string, bool)`

GetStudyTypeOk returns a tuple with the StudyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyType

`func (o *SponsorTrialsRow) SetStudyType(v string)`

SetStudyType sets StudyType field to given value.

### HasStudyType

`func (o *SponsorTrialsRow) HasStudyType() bool`

HasStudyType returns a boolean if a field has been set.

### SetStudyTypeNil

`func (o *SponsorTrialsRow) SetStudyTypeNil(b bool)`

 SetStudyTypeNil sets the value for StudyType to be an explicit nil

### UnsetStudyType
`func (o *SponsorTrialsRow) UnsetStudyType()`

UnsetStudyType ensures that no value is present for StudyType, not even an explicit nil
### GetPhase

`func (o *SponsorTrialsRow) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *SponsorTrialsRow) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *SponsorTrialsRow) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *SponsorTrialsRow) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *SponsorTrialsRow) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *SponsorTrialsRow) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetOverallStatus

`func (o *SponsorTrialsRow) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *SponsorTrialsRow) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *SponsorTrialsRow) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *SponsorTrialsRow) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### SetOverallStatusNil

`func (o *SponsorTrialsRow) SetOverallStatusNil(b bool)`

 SetOverallStatusNil sets the value for OverallStatus to be an explicit nil

### UnsetOverallStatus
`func (o *SponsorTrialsRow) UnsetOverallStatus()`

UnsetOverallStatus ensures that no value is present for OverallStatus, not even an explicit nil
### GetStartDate

`func (o *SponsorTrialsRow) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SponsorTrialsRow) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SponsorTrialsRow) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SponsorTrialsRow) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *SponsorTrialsRow) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *SponsorTrialsRow) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetStartYear

`func (o *SponsorTrialsRow) GetStartYear() int32`

GetStartYear returns the StartYear field if non-nil, zero value otherwise.

### GetStartYearOk

`func (o *SponsorTrialsRow) GetStartYearOk() (*int32, bool)`

GetStartYearOk returns a tuple with the StartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartYear

`func (o *SponsorTrialsRow) SetStartYear(v int32)`

SetStartYear sets StartYear field to given value.

### HasStartYear

`func (o *SponsorTrialsRow) HasStartYear() bool`

HasStartYear returns a boolean if a field has been set.

### SetStartYearNil

`func (o *SponsorTrialsRow) SetStartYearNil(b bool)`

 SetStartYearNil sets the value for StartYear to be an explicit nil

### UnsetStartYear
`func (o *SponsorTrialsRow) UnsetStartYear()`

UnsetStartYear ensures that no value is present for StartYear, not even an explicit nil
### GetPrimaryCompletionDate

`func (o *SponsorTrialsRow) GetPrimaryCompletionDate() string`

GetPrimaryCompletionDate returns the PrimaryCompletionDate field if non-nil, zero value otherwise.

### GetPrimaryCompletionDateOk

`func (o *SponsorTrialsRow) GetPrimaryCompletionDateOk() (*string, bool)`

GetPrimaryCompletionDateOk returns a tuple with the PrimaryCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCompletionDate

`func (o *SponsorTrialsRow) SetPrimaryCompletionDate(v string)`

SetPrimaryCompletionDate sets PrimaryCompletionDate field to given value.

### HasPrimaryCompletionDate

`func (o *SponsorTrialsRow) HasPrimaryCompletionDate() bool`

HasPrimaryCompletionDate returns a boolean if a field has been set.

### SetPrimaryCompletionDateNil

`func (o *SponsorTrialsRow) SetPrimaryCompletionDateNil(b bool)`

 SetPrimaryCompletionDateNil sets the value for PrimaryCompletionDate to be an explicit nil

### UnsetPrimaryCompletionDate
`func (o *SponsorTrialsRow) UnsetPrimaryCompletionDate()`

UnsetPrimaryCompletionDate ensures that no value is present for PrimaryCompletionDate, not even an explicit nil
### GetEnrollment

`func (o *SponsorTrialsRow) GetEnrollment() int32`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *SponsorTrialsRow) GetEnrollmentOk() (*int32, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *SponsorTrialsRow) SetEnrollment(v int32)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *SponsorTrialsRow) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### SetEnrollmentNil

`func (o *SponsorTrialsRow) SetEnrollmentNil(b bool)`

 SetEnrollmentNil sets the value for Enrollment to be an explicit nil

### UnsetEnrollment
`func (o *SponsorTrialsRow) UnsetEnrollment()`

UnsetEnrollment ensures that no value is present for Enrollment, not even an explicit nil
### GetWereResultsReported

`func (o *SponsorTrialsRow) GetWereResultsReported() bool`

GetWereResultsReported returns the WereResultsReported field if non-nil, zero value otherwise.

### GetWereResultsReportedOk

`func (o *SponsorTrialsRow) GetWereResultsReportedOk() (*bool, bool)`

GetWereResultsReportedOk returns a tuple with the WereResultsReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWereResultsReported

`func (o *SponsorTrialsRow) SetWereResultsReported(v bool)`

SetWereResultsReported sets WereResultsReported field to given value.

### HasWereResultsReported

`func (o *SponsorTrialsRow) HasWereResultsReported() bool`

HasWereResultsReported returns a boolean if a field has been set.

### SetWereResultsReportedNil

`func (o *SponsorTrialsRow) SetWereResultsReportedNil(b bool)`

 SetWereResultsReportedNil sets the value for WereResultsReported to be an explicit nil

### UnsetWereResultsReported
`func (o *SponsorTrialsRow) UnsetWereResultsReported()`

UnsetWereResultsReported ensures that no value is present for WereResultsReported, not even an explicit nil
### GetConditions

`func (o *SponsorTrialsRow) GetConditions() []string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *SponsorTrialsRow) GetConditionsOk() (*[]string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *SponsorTrialsRow) SetConditions(v []string)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *SponsorTrialsRow) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *SponsorTrialsRow) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *SponsorTrialsRow) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetInterventions

`func (o *SponsorTrialsRow) GetInterventions() []string`

GetInterventions returns the Interventions field if non-nil, zero value otherwise.

### GetInterventionsOk

`func (o *SponsorTrialsRow) GetInterventionsOk() (*[]string, bool)`

GetInterventionsOk returns a tuple with the Interventions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterventions

`func (o *SponsorTrialsRow) SetInterventions(v []string)`

SetInterventions sets Interventions field to given value.

### HasInterventions

`func (o *SponsorTrialsRow) HasInterventions() bool`

HasInterventions returns a boolean if a field has been set.

### SetInterventionsNil

`func (o *SponsorTrialsRow) SetInterventionsNil(b bool)`

 SetInterventionsNil sets the value for Interventions to be an explicit nil

### UnsetInterventions
`func (o *SponsorTrialsRow) UnsetInterventions()`

UnsetInterventions ensures that no value is present for Interventions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


