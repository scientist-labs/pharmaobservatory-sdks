# RareDiseaseTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrphaCode** | Pointer to **NullableString** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] 
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**MappingArm** | Pointer to **NullableString** | exact when the link rests on an asserted equivalence, unspecified when it rests on a MONDO cross-reference that states no relation. | [optional] 
**EdgeScope** | Pointer to **NullableString** | direct when the trial&#39;s own condition list names the mapped condition, broader_family when only a broader condition does. | [optional] 
**ViaExact** | Pointer to **NullableBool** | True when an exact mapping links the trial to the disease. | [optional] 
**ViaUnspecified** | Pointer to **NullableBool** | True when an unspecified-relation mapping links the trial to the disease. | [optional] 
**ViaDirect** | Pointer to **NullableBool** | True when a direct edge links the trial to the disease. | [optional] 
**ViaBroader** | Pointer to **NullableBool** | True when a broader-family edge links the trial to the disease. | [optional] 
**MappingTargetCurie** | Pointer to **NullableString** | The MeSH condition (as a prefixed identifier) the mapping reaches the trial through. | [optional] 
**MappingSources** | Pointer to **[]string** | The files that assert the mapping (ordo, orphadata_product1 or mondo_legacy_bridge). | [optional] 
**TrialTitle** | Pointer to **NullableString** | The trial&#39;s brief title. | [optional] 
**OverallStatus** | Pointer to **NullableString** | The registry&#39;s overall recruitment status (COMPLETED, RECRUITING, TERMINATED and so on). | [optional] 
**StudyType** | Pointer to **NullableString** | The registry&#39;s study type (INTERVENTIONAL, OBSERVATIONAL or EXPANDED_ACCESS). | [optional] 
**PhaseLabel** | Pointer to **NullableString** | The trial phase as a display label (Phase 2, Phase 3, Observational and so on). | [optional] 
**Enrollment** | Pointer to **NullableInt32** | The trial&#39;s total enrollment as registered. | [optional] 
**EnrollmentType** | Pointer to **NullableString** | ACTUAL when the registered enrollment is the number enrolled, ESTIMATED when it is the planned number. | [optional] 
**StartDate** | Pointer to **NullableString** | The trial&#39;s start date as registered. | [optional] 
**StartYear** | Pointer to **NullableInt32** | The year of start_date. | [optional] 
**CompletionDate** | Pointer to **NullableString** | The trial&#39;s completion date as registered. | [optional] 
**WhyStopped** | Pointer to **NullableString** | The reason the sponsor gave for stopping the trial, where it stopped early. | [optional] 
**WereResultsReported** | Pointer to **NullableBool** | True when the sponsor has posted results to the registry. | [optional] 
**LastUpdatePostedDate** | Pointer to **NullableString** | The date the registry record was last updated. | [optional] 
**FacilityCount** | Pointer to **NullableInt32** | Facilities the registry lists for the trial. | [optional] 
**FacilityCountryCount** | Pointer to **NullableInt32** | Countries the listed facilities are in. | [optional] 
**FacilityCountries** | Pointer to **[]string** | The countries the listed facilities are in. | [optional] 
**DaysSinceUpdate** | Pointer to **NullableInt32** | Days from the last registry update to the registry snapshot date. | [optional] 
**IsUnverified** | Pointer to **NullableBool** | True when the registry record was last updated two or more years before the snapshot date. | [optional] 
**SponsorName** | Pointer to **NullableString** | The sponsor&#39;s name as filed in the registry. | [optional] 
**SponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the sponsor resolved to (null when unresolved). | [optional] 
**SponsorCompanyId** | Pointer to **NullableString** | Company level: the sponsor&#39;s company (null when unresolved). | [optional] 
**SponsorCompanyName** | Pointer to **NullableString** | Display name of the sponsor&#39;s company. | [optional] 
**SponsorGroupId** | Pointer to **NullableString** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SponsorGroupName** | Pointer to **NullableString** | Display name of the sponsor&#39;s group. | [optional] 
**SponsorAgencyClass** | Pointer to **NullableString** | The funder class the registry declared for the lead sponsor; null when the trial has no lead sponsor row. | [optional] 
**SponsorResolutionConfidence** | Pointer to **NullableString** | How firmly the lead sponsor resolved to an organization record (high, medium or unresolved). | [optional] 
**LicenseName** | Pointer to **NullableString** | The licence of the trial registry data. | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRareDiseaseTrialsRow

`func NewRareDiseaseTrialsRow() *RareDiseaseTrialsRow`

NewRareDiseaseTrialsRow instantiates a new RareDiseaseTrialsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRareDiseaseTrialsRowWithDefaults

`func NewRareDiseaseTrialsRowWithDefaults() *RareDiseaseTrialsRow`

NewRareDiseaseTrialsRowWithDefaults instantiates a new RareDiseaseTrialsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrphaCode

`func (o *RareDiseaseTrialsRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *RareDiseaseTrialsRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *RareDiseaseTrialsRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *RareDiseaseTrialsRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *RareDiseaseTrialsRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *RareDiseaseTrialsRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetTrialId

`func (o *RareDiseaseTrialsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *RareDiseaseTrialsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *RareDiseaseTrialsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *RareDiseaseTrialsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *RareDiseaseTrialsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *RareDiseaseTrialsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetMappingArm

`func (o *RareDiseaseTrialsRow) GetMappingArm() string`

GetMappingArm returns the MappingArm field if non-nil, zero value otherwise.

### GetMappingArmOk

`func (o *RareDiseaseTrialsRow) GetMappingArmOk() (*string, bool)`

GetMappingArmOk returns a tuple with the MappingArm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingArm

`func (o *RareDiseaseTrialsRow) SetMappingArm(v string)`

SetMappingArm sets MappingArm field to given value.

### HasMappingArm

`func (o *RareDiseaseTrialsRow) HasMappingArm() bool`

HasMappingArm returns a boolean if a field has been set.

### SetMappingArmNil

`func (o *RareDiseaseTrialsRow) SetMappingArmNil(b bool)`

 SetMappingArmNil sets the value for MappingArm to be an explicit nil

### UnsetMappingArm
`func (o *RareDiseaseTrialsRow) UnsetMappingArm()`

UnsetMappingArm ensures that no value is present for MappingArm, not even an explicit nil
### GetEdgeScope

`func (o *RareDiseaseTrialsRow) GetEdgeScope() string`

GetEdgeScope returns the EdgeScope field if non-nil, zero value otherwise.

### GetEdgeScopeOk

`func (o *RareDiseaseTrialsRow) GetEdgeScopeOk() (*string, bool)`

GetEdgeScopeOk returns a tuple with the EdgeScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeScope

`func (o *RareDiseaseTrialsRow) SetEdgeScope(v string)`

SetEdgeScope sets EdgeScope field to given value.

### HasEdgeScope

`func (o *RareDiseaseTrialsRow) HasEdgeScope() bool`

HasEdgeScope returns a boolean if a field has been set.

### SetEdgeScopeNil

`func (o *RareDiseaseTrialsRow) SetEdgeScopeNil(b bool)`

 SetEdgeScopeNil sets the value for EdgeScope to be an explicit nil

### UnsetEdgeScope
`func (o *RareDiseaseTrialsRow) UnsetEdgeScope()`

UnsetEdgeScope ensures that no value is present for EdgeScope, not even an explicit nil
### GetViaExact

`func (o *RareDiseaseTrialsRow) GetViaExact() bool`

GetViaExact returns the ViaExact field if non-nil, zero value otherwise.

### GetViaExactOk

`func (o *RareDiseaseTrialsRow) GetViaExactOk() (*bool, bool)`

GetViaExactOk returns a tuple with the ViaExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaExact

`func (o *RareDiseaseTrialsRow) SetViaExact(v bool)`

SetViaExact sets ViaExact field to given value.

### HasViaExact

`func (o *RareDiseaseTrialsRow) HasViaExact() bool`

HasViaExact returns a boolean if a field has been set.

### SetViaExactNil

`func (o *RareDiseaseTrialsRow) SetViaExactNil(b bool)`

 SetViaExactNil sets the value for ViaExact to be an explicit nil

### UnsetViaExact
`func (o *RareDiseaseTrialsRow) UnsetViaExact()`

UnsetViaExact ensures that no value is present for ViaExact, not even an explicit nil
### GetViaUnspecified

`func (o *RareDiseaseTrialsRow) GetViaUnspecified() bool`

GetViaUnspecified returns the ViaUnspecified field if non-nil, zero value otherwise.

### GetViaUnspecifiedOk

`func (o *RareDiseaseTrialsRow) GetViaUnspecifiedOk() (*bool, bool)`

GetViaUnspecifiedOk returns a tuple with the ViaUnspecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaUnspecified

`func (o *RareDiseaseTrialsRow) SetViaUnspecified(v bool)`

SetViaUnspecified sets ViaUnspecified field to given value.

### HasViaUnspecified

`func (o *RareDiseaseTrialsRow) HasViaUnspecified() bool`

HasViaUnspecified returns a boolean if a field has been set.

### SetViaUnspecifiedNil

`func (o *RareDiseaseTrialsRow) SetViaUnspecifiedNil(b bool)`

 SetViaUnspecifiedNil sets the value for ViaUnspecified to be an explicit nil

### UnsetViaUnspecified
`func (o *RareDiseaseTrialsRow) UnsetViaUnspecified()`

UnsetViaUnspecified ensures that no value is present for ViaUnspecified, not even an explicit nil
### GetViaDirect

`func (o *RareDiseaseTrialsRow) GetViaDirect() bool`

GetViaDirect returns the ViaDirect field if non-nil, zero value otherwise.

### GetViaDirectOk

`func (o *RareDiseaseTrialsRow) GetViaDirectOk() (*bool, bool)`

GetViaDirectOk returns a tuple with the ViaDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaDirect

`func (o *RareDiseaseTrialsRow) SetViaDirect(v bool)`

SetViaDirect sets ViaDirect field to given value.

### HasViaDirect

`func (o *RareDiseaseTrialsRow) HasViaDirect() bool`

HasViaDirect returns a boolean if a field has been set.

### SetViaDirectNil

`func (o *RareDiseaseTrialsRow) SetViaDirectNil(b bool)`

 SetViaDirectNil sets the value for ViaDirect to be an explicit nil

### UnsetViaDirect
`func (o *RareDiseaseTrialsRow) UnsetViaDirect()`

UnsetViaDirect ensures that no value is present for ViaDirect, not even an explicit nil
### GetViaBroader

`func (o *RareDiseaseTrialsRow) GetViaBroader() bool`

GetViaBroader returns the ViaBroader field if non-nil, zero value otherwise.

### GetViaBroaderOk

`func (o *RareDiseaseTrialsRow) GetViaBroaderOk() (*bool, bool)`

GetViaBroaderOk returns a tuple with the ViaBroader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaBroader

`func (o *RareDiseaseTrialsRow) SetViaBroader(v bool)`

SetViaBroader sets ViaBroader field to given value.

### HasViaBroader

`func (o *RareDiseaseTrialsRow) HasViaBroader() bool`

HasViaBroader returns a boolean if a field has been set.

### SetViaBroaderNil

`func (o *RareDiseaseTrialsRow) SetViaBroaderNil(b bool)`

 SetViaBroaderNil sets the value for ViaBroader to be an explicit nil

### UnsetViaBroader
`func (o *RareDiseaseTrialsRow) UnsetViaBroader()`

UnsetViaBroader ensures that no value is present for ViaBroader, not even an explicit nil
### GetMappingTargetCurie

`func (o *RareDiseaseTrialsRow) GetMappingTargetCurie() string`

GetMappingTargetCurie returns the MappingTargetCurie field if non-nil, zero value otherwise.

### GetMappingTargetCurieOk

`func (o *RareDiseaseTrialsRow) GetMappingTargetCurieOk() (*string, bool)`

GetMappingTargetCurieOk returns a tuple with the MappingTargetCurie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingTargetCurie

`func (o *RareDiseaseTrialsRow) SetMappingTargetCurie(v string)`

SetMappingTargetCurie sets MappingTargetCurie field to given value.

### HasMappingTargetCurie

`func (o *RareDiseaseTrialsRow) HasMappingTargetCurie() bool`

HasMappingTargetCurie returns a boolean if a field has been set.

### SetMappingTargetCurieNil

`func (o *RareDiseaseTrialsRow) SetMappingTargetCurieNil(b bool)`

 SetMappingTargetCurieNil sets the value for MappingTargetCurie to be an explicit nil

### UnsetMappingTargetCurie
`func (o *RareDiseaseTrialsRow) UnsetMappingTargetCurie()`

UnsetMappingTargetCurie ensures that no value is present for MappingTargetCurie, not even an explicit nil
### GetMappingSources

`func (o *RareDiseaseTrialsRow) GetMappingSources() []string`

GetMappingSources returns the MappingSources field if non-nil, zero value otherwise.

### GetMappingSourcesOk

`func (o *RareDiseaseTrialsRow) GetMappingSourcesOk() (*[]string, bool)`

GetMappingSourcesOk returns a tuple with the MappingSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingSources

`func (o *RareDiseaseTrialsRow) SetMappingSources(v []string)`

SetMappingSources sets MappingSources field to given value.

### HasMappingSources

`func (o *RareDiseaseTrialsRow) HasMappingSources() bool`

HasMappingSources returns a boolean if a field has been set.

### SetMappingSourcesNil

`func (o *RareDiseaseTrialsRow) SetMappingSourcesNil(b bool)`

 SetMappingSourcesNil sets the value for MappingSources to be an explicit nil

### UnsetMappingSources
`func (o *RareDiseaseTrialsRow) UnsetMappingSources()`

UnsetMappingSources ensures that no value is present for MappingSources, not even an explicit nil
### GetTrialTitle

`func (o *RareDiseaseTrialsRow) GetTrialTitle() string`

GetTrialTitle returns the TrialTitle field if non-nil, zero value otherwise.

### GetTrialTitleOk

`func (o *RareDiseaseTrialsRow) GetTrialTitleOk() (*string, bool)`

GetTrialTitleOk returns a tuple with the TrialTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialTitle

`func (o *RareDiseaseTrialsRow) SetTrialTitle(v string)`

SetTrialTitle sets TrialTitle field to given value.

### HasTrialTitle

`func (o *RareDiseaseTrialsRow) HasTrialTitle() bool`

HasTrialTitle returns a boolean if a field has been set.

### SetTrialTitleNil

`func (o *RareDiseaseTrialsRow) SetTrialTitleNil(b bool)`

 SetTrialTitleNil sets the value for TrialTitle to be an explicit nil

### UnsetTrialTitle
`func (o *RareDiseaseTrialsRow) UnsetTrialTitle()`

UnsetTrialTitle ensures that no value is present for TrialTitle, not even an explicit nil
### GetOverallStatus

`func (o *RareDiseaseTrialsRow) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *RareDiseaseTrialsRow) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *RareDiseaseTrialsRow) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *RareDiseaseTrialsRow) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### SetOverallStatusNil

`func (o *RareDiseaseTrialsRow) SetOverallStatusNil(b bool)`

 SetOverallStatusNil sets the value for OverallStatus to be an explicit nil

### UnsetOverallStatus
`func (o *RareDiseaseTrialsRow) UnsetOverallStatus()`

UnsetOverallStatus ensures that no value is present for OverallStatus, not even an explicit nil
### GetStudyType

`func (o *RareDiseaseTrialsRow) GetStudyType() string`

GetStudyType returns the StudyType field if non-nil, zero value otherwise.

### GetStudyTypeOk

`func (o *RareDiseaseTrialsRow) GetStudyTypeOk() (*string, bool)`

GetStudyTypeOk returns a tuple with the StudyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyType

`func (o *RareDiseaseTrialsRow) SetStudyType(v string)`

SetStudyType sets StudyType field to given value.

### HasStudyType

`func (o *RareDiseaseTrialsRow) HasStudyType() bool`

HasStudyType returns a boolean if a field has been set.

### SetStudyTypeNil

`func (o *RareDiseaseTrialsRow) SetStudyTypeNil(b bool)`

 SetStudyTypeNil sets the value for StudyType to be an explicit nil

### UnsetStudyType
`func (o *RareDiseaseTrialsRow) UnsetStudyType()`

UnsetStudyType ensures that no value is present for StudyType, not even an explicit nil
### GetPhaseLabel

`func (o *RareDiseaseTrialsRow) GetPhaseLabel() string`

GetPhaseLabel returns the PhaseLabel field if non-nil, zero value otherwise.

### GetPhaseLabelOk

`func (o *RareDiseaseTrialsRow) GetPhaseLabelOk() (*string, bool)`

GetPhaseLabelOk returns a tuple with the PhaseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseLabel

`func (o *RareDiseaseTrialsRow) SetPhaseLabel(v string)`

SetPhaseLabel sets PhaseLabel field to given value.

### HasPhaseLabel

`func (o *RareDiseaseTrialsRow) HasPhaseLabel() bool`

HasPhaseLabel returns a boolean if a field has been set.

### SetPhaseLabelNil

`func (o *RareDiseaseTrialsRow) SetPhaseLabelNil(b bool)`

 SetPhaseLabelNil sets the value for PhaseLabel to be an explicit nil

### UnsetPhaseLabel
`func (o *RareDiseaseTrialsRow) UnsetPhaseLabel()`

UnsetPhaseLabel ensures that no value is present for PhaseLabel, not even an explicit nil
### GetEnrollment

`func (o *RareDiseaseTrialsRow) GetEnrollment() int32`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *RareDiseaseTrialsRow) GetEnrollmentOk() (*int32, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *RareDiseaseTrialsRow) SetEnrollment(v int32)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *RareDiseaseTrialsRow) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### SetEnrollmentNil

`func (o *RareDiseaseTrialsRow) SetEnrollmentNil(b bool)`

 SetEnrollmentNil sets the value for Enrollment to be an explicit nil

### UnsetEnrollment
`func (o *RareDiseaseTrialsRow) UnsetEnrollment()`

UnsetEnrollment ensures that no value is present for Enrollment, not even an explicit nil
### GetEnrollmentType

`func (o *RareDiseaseTrialsRow) GetEnrollmentType() string`

GetEnrollmentType returns the EnrollmentType field if non-nil, zero value otherwise.

### GetEnrollmentTypeOk

`func (o *RareDiseaseTrialsRow) GetEnrollmentTypeOk() (*string, bool)`

GetEnrollmentTypeOk returns a tuple with the EnrollmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentType

`func (o *RareDiseaseTrialsRow) SetEnrollmentType(v string)`

SetEnrollmentType sets EnrollmentType field to given value.

### HasEnrollmentType

`func (o *RareDiseaseTrialsRow) HasEnrollmentType() bool`

HasEnrollmentType returns a boolean if a field has been set.

### SetEnrollmentTypeNil

`func (o *RareDiseaseTrialsRow) SetEnrollmentTypeNil(b bool)`

 SetEnrollmentTypeNil sets the value for EnrollmentType to be an explicit nil

### UnsetEnrollmentType
`func (o *RareDiseaseTrialsRow) UnsetEnrollmentType()`

UnsetEnrollmentType ensures that no value is present for EnrollmentType, not even an explicit nil
### GetStartDate

`func (o *RareDiseaseTrialsRow) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *RareDiseaseTrialsRow) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *RareDiseaseTrialsRow) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *RareDiseaseTrialsRow) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *RareDiseaseTrialsRow) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *RareDiseaseTrialsRow) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetStartYear

`func (o *RareDiseaseTrialsRow) GetStartYear() int32`

GetStartYear returns the StartYear field if non-nil, zero value otherwise.

### GetStartYearOk

`func (o *RareDiseaseTrialsRow) GetStartYearOk() (*int32, bool)`

GetStartYearOk returns a tuple with the StartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartYear

`func (o *RareDiseaseTrialsRow) SetStartYear(v int32)`

SetStartYear sets StartYear field to given value.

### HasStartYear

`func (o *RareDiseaseTrialsRow) HasStartYear() bool`

HasStartYear returns a boolean if a field has been set.

### SetStartYearNil

`func (o *RareDiseaseTrialsRow) SetStartYearNil(b bool)`

 SetStartYearNil sets the value for StartYear to be an explicit nil

### UnsetStartYear
`func (o *RareDiseaseTrialsRow) UnsetStartYear()`

UnsetStartYear ensures that no value is present for StartYear, not even an explicit nil
### GetCompletionDate

`func (o *RareDiseaseTrialsRow) GetCompletionDate() string`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *RareDiseaseTrialsRow) GetCompletionDateOk() (*string, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *RareDiseaseTrialsRow) SetCompletionDate(v string)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *RareDiseaseTrialsRow) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### SetCompletionDateNil

`func (o *RareDiseaseTrialsRow) SetCompletionDateNil(b bool)`

 SetCompletionDateNil sets the value for CompletionDate to be an explicit nil

### UnsetCompletionDate
`func (o *RareDiseaseTrialsRow) UnsetCompletionDate()`

UnsetCompletionDate ensures that no value is present for CompletionDate, not even an explicit nil
### GetWhyStopped

`func (o *RareDiseaseTrialsRow) GetWhyStopped() string`

GetWhyStopped returns the WhyStopped field if non-nil, zero value otherwise.

### GetWhyStoppedOk

`func (o *RareDiseaseTrialsRow) GetWhyStoppedOk() (*string, bool)`

GetWhyStoppedOk returns a tuple with the WhyStopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhyStopped

`func (o *RareDiseaseTrialsRow) SetWhyStopped(v string)`

SetWhyStopped sets WhyStopped field to given value.

### HasWhyStopped

`func (o *RareDiseaseTrialsRow) HasWhyStopped() bool`

HasWhyStopped returns a boolean if a field has been set.

### SetWhyStoppedNil

`func (o *RareDiseaseTrialsRow) SetWhyStoppedNil(b bool)`

 SetWhyStoppedNil sets the value for WhyStopped to be an explicit nil

### UnsetWhyStopped
`func (o *RareDiseaseTrialsRow) UnsetWhyStopped()`

UnsetWhyStopped ensures that no value is present for WhyStopped, not even an explicit nil
### GetWereResultsReported

`func (o *RareDiseaseTrialsRow) GetWereResultsReported() bool`

GetWereResultsReported returns the WereResultsReported field if non-nil, zero value otherwise.

### GetWereResultsReportedOk

`func (o *RareDiseaseTrialsRow) GetWereResultsReportedOk() (*bool, bool)`

GetWereResultsReportedOk returns a tuple with the WereResultsReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWereResultsReported

`func (o *RareDiseaseTrialsRow) SetWereResultsReported(v bool)`

SetWereResultsReported sets WereResultsReported field to given value.

### HasWereResultsReported

`func (o *RareDiseaseTrialsRow) HasWereResultsReported() bool`

HasWereResultsReported returns a boolean if a field has been set.

### SetWereResultsReportedNil

`func (o *RareDiseaseTrialsRow) SetWereResultsReportedNil(b bool)`

 SetWereResultsReportedNil sets the value for WereResultsReported to be an explicit nil

### UnsetWereResultsReported
`func (o *RareDiseaseTrialsRow) UnsetWereResultsReported()`

UnsetWereResultsReported ensures that no value is present for WereResultsReported, not even an explicit nil
### GetLastUpdatePostedDate

`func (o *RareDiseaseTrialsRow) GetLastUpdatePostedDate() string`

GetLastUpdatePostedDate returns the LastUpdatePostedDate field if non-nil, zero value otherwise.

### GetLastUpdatePostedDateOk

`func (o *RareDiseaseTrialsRow) GetLastUpdatePostedDateOk() (*string, bool)`

GetLastUpdatePostedDateOk returns a tuple with the LastUpdatePostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatePostedDate

`func (o *RareDiseaseTrialsRow) SetLastUpdatePostedDate(v string)`

SetLastUpdatePostedDate sets LastUpdatePostedDate field to given value.

### HasLastUpdatePostedDate

`func (o *RareDiseaseTrialsRow) HasLastUpdatePostedDate() bool`

HasLastUpdatePostedDate returns a boolean if a field has been set.

### SetLastUpdatePostedDateNil

`func (o *RareDiseaseTrialsRow) SetLastUpdatePostedDateNil(b bool)`

 SetLastUpdatePostedDateNil sets the value for LastUpdatePostedDate to be an explicit nil

### UnsetLastUpdatePostedDate
`func (o *RareDiseaseTrialsRow) UnsetLastUpdatePostedDate()`

UnsetLastUpdatePostedDate ensures that no value is present for LastUpdatePostedDate, not even an explicit nil
### GetFacilityCount

`func (o *RareDiseaseTrialsRow) GetFacilityCount() int32`

GetFacilityCount returns the FacilityCount field if non-nil, zero value otherwise.

### GetFacilityCountOk

`func (o *RareDiseaseTrialsRow) GetFacilityCountOk() (*int32, bool)`

GetFacilityCountOk returns a tuple with the FacilityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCount

`func (o *RareDiseaseTrialsRow) SetFacilityCount(v int32)`

SetFacilityCount sets FacilityCount field to given value.

### HasFacilityCount

`func (o *RareDiseaseTrialsRow) HasFacilityCount() bool`

HasFacilityCount returns a boolean if a field has been set.

### SetFacilityCountNil

`func (o *RareDiseaseTrialsRow) SetFacilityCountNil(b bool)`

 SetFacilityCountNil sets the value for FacilityCount to be an explicit nil

### UnsetFacilityCount
`func (o *RareDiseaseTrialsRow) UnsetFacilityCount()`

UnsetFacilityCount ensures that no value is present for FacilityCount, not even an explicit nil
### GetFacilityCountryCount

`func (o *RareDiseaseTrialsRow) GetFacilityCountryCount() int32`

GetFacilityCountryCount returns the FacilityCountryCount field if non-nil, zero value otherwise.

### GetFacilityCountryCountOk

`func (o *RareDiseaseTrialsRow) GetFacilityCountryCountOk() (*int32, bool)`

GetFacilityCountryCountOk returns a tuple with the FacilityCountryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCountryCount

`func (o *RareDiseaseTrialsRow) SetFacilityCountryCount(v int32)`

SetFacilityCountryCount sets FacilityCountryCount field to given value.

### HasFacilityCountryCount

`func (o *RareDiseaseTrialsRow) HasFacilityCountryCount() bool`

HasFacilityCountryCount returns a boolean if a field has been set.

### SetFacilityCountryCountNil

`func (o *RareDiseaseTrialsRow) SetFacilityCountryCountNil(b bool)`

 SetFacilityCountryCountNil sets the value for FacilityCountryCount to be an explicit nil

### UnsetFacilityCountryCount
`func (o *RareDiseaseTrialsRow) UnsetFacilityCountryCount()`

UnsetFacilityCountryCount ensures that no value is present for FacilityCountryCount, not even an explicit nil
### GetFacilityCountries

`func (o *RareDiseaseTrialsRow) GetFacilityCountries() []string`

GetFacilityCountries returns the FacilityCountries field if non-nil, zero value otherwise.

### GetFacilityCountriesOk

`func (o *RareDiseaseTrialsRow) GetFacilityCountriesOk() (*[]string, bool)`

GetFacilityCountriesOk returns a tuple with the FacilityCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCountries

`func (o *RareDiseaseTrialsRow) SetFacilityCountries(v []string)`

SetFacilityCountries sets FacilityCountries field to given value.

### HasFacilityCountries

`func (o *RareDiseaseTrialsRow) HasFacilityCountries() bool`

HasFacilityCountries returns a boolean if a field has been set.

### SetFacilityCountriesNil

`func (o *RareDiseaseTrialsRow) SetFacilityCountriesNil(b bool)`

 SetFacilityCountriesNil sets the value for FacilityCountries to be an explicit nil

### UnsetFacilityCountries
`func (o *RareDiseaseTrialsRow) UnsetFacilityCountries()`

UnsetFacilityCountries ensures that no value is present for FacilityCountries, not even an explicit nil
### GetDaysSinceUpdate

`func (o *RareDiseaseTrialsRow) GetDaysSinceUpdate() int32`

GetDaysSinceUpdate returns the DaysSinceUpdate field if non-nil, zero value otherwise.

### GetDaysSinceUpdateOk

`func (o *RareDiseaseTrialsRow) GetDaysSinceUpdateOk() (*int32, bool)`

GetDaysSinceUpdateOk returns a tuple with the DaysSinceUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysSinceUpdate

`func (o *RareDiseaseTrialsRow) SetDaysSinceUpdate(v int32)`

SetDaysSinceUpdate sets DaysSinceUpdate field to given value.

### HasDaysSinceUpdate

`func (o *RareDiseaseTrialsRow) HasDaysSinceUpdate() bool`

HasDaysSinceUpdate returns a boolean if a field has been set.

### SetDaysSinceUpdateNil

`func (o *RareDiseaseTrialsRow) SetDaysSinceUpdateNil(b bool)`

 SetDaysSinceUpdateNil sets the value for DaysSinceUpdate to be an explicit nil

### UnsetDaysSinceUpdate
`func (o *RareDiseaseTrialsRow) UnsetDaysSinceUpdate()`

UnsetDaysSinceUpdate ensures that no value is present for DaysSinceUpdate, not even an explicit nil
### GetIsUnverified

`func (o *RareDiseaseTrialsRow) GetIsUnverified() bool`

GetIsUnverified returns the IsUnverified field if non-nil, zero value otherwise.

### GetIsUnverifiedOk

`func (o *RareDiseaseTrialsRow) GetIsUnverifiedOk() (*bool, bool)`

GetIsUnverifiedOk returns a tuple with the IsUnverified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnverified

`func (o *RareDiseaseTrialsRow) SetIsUnverified(v bool)`

SetIsUnverified sets IsUnverified field to given value.

### HasIsUnverified

`func (o *RareDiseaseTrialsRow) HasIsUnverified() bool`

HasIsUnverified returns a boolean if a field has been set.

### SetIsUnverifiedNil

`func (o *RareDiseaseTrialsRow) SetIsUnverifiedNil(b bool)`

 SetIsUnverifiedNil sets the value for IsUnverified to be an explicit nil

### UnsetIsUnverified
`func (o *RareDiseaseTrialsRow) UnsetIsUnverified()`

UnsetIsUnverified ensures that no value is present for IsUnverified, not even an explicit nil
### GetSponsorName

`func (o *RareDiseaseTrialsRow) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *RareDiseaseTrialsRow) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *RareDiseaseTrialsRow) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *RareDiseaseTrialsRow) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### SetSponsorNameNil

`func (o *RareDiseaseTrialsRow) SetSponsorNameNil(b bool)`

 SetSponsorNameNil sets the value for SponsorName to be an explicit nil

### UnsetSponsorName
`func (o *RareDiseaseTrialsRow) UnsetSponsorName()`

UnsetSponsorName ensures that no value is present for SponsorName, not even an explicit nil
### GetSponsorOrgId

`func (o *RareDiseaseTrialsRow) GetSponsorOrgId() string`

GetSponsorOrgId returns the SponsorOrgId field if non-nil, zero value otherwise.

### GetSponsorOrgIdOk

`func (o *RareDiseaseTrialsRow) GetSponsorOrgIdOk() (*string, bool)`

GetSponsorOrgIdOk returns a tuple with the SponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgId

`func (o *RareDiseaseTrialsRow) SetSponsorOrgId(v string)`

SetSponsorOrgId sets SponsorOrgId field to given value.

### HasSponsorOrgId

`func (o *RareDiseaseTrialsRow) HasSponsorOrgId() bool`

HasSponsorOrgId returns a boolean if a field has been set.

### SetSponsorOrgIdNil

`func (o *RareDiseaseTrialsRow) SetSponsorOrgIdNil(b bool)`

 SetSponsorOrgIdNil sets the value for SponsorOrgId to be an explicit nil

### UnsetSponsorOrgId
`func (o *RareDiseaseTrialsRow) UnsetSponsorOrgId()`

UnsetSponsorOrgId ensures that no value is present for SponsorOrgId, not even an explicit nil
### GetSponsorCompanyId

`func (o *RareDiseaseTrialsRow) GetSponsorCompanyId() string`

GetSponsorCompanyId returns the SponsorCompanyId field if non-nil, zero value otherwise.

### GetSponsorCompanyIdOk

`func (o *RareDiseaseTrialsRow) GetSponsorCompanyIdOk() (*string, bool)`

GetSponsorCompanyIdOk returns a tuple with the SponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyId

`func (o *RareDiseaseTrialsRow) SetSponsorCompanyId(v string)`

SetSponsorCompanyId sets SponsorCompanyId field to given value.

### HasSponsorCompanyId

`func (o *RareDiseaseTrialsRow) HasSponsorCompanyId() bool`

HasSponsorCompanyId returns a boolean if a field has been set.

### SetSponsorCompanyIdNil

`func (o *RareDiseaseTrialsRow) SetSponsorCompanyIdNil(b bool)`

 SetSponsorCompanyIdNil sets the value for SponsorCompanyId to be an explicit nil

### UnsetSponsorCompanyId
`func (o *RareDiseaseTrialsRow) UnsetSponsorCompanyId()`

UnsetSponsorCompanyId ensures that no value is present for SponsorCompanyId, not even an explicit nil
### GetSponsorCompanyName

`func (o *RareDiseaseTrialsRow) GetSponsorCompanyName() string`

GetSponsorCompanyName returns the SponsorCompanyName field if non-nil, zero value otherwise.

### GetSponsorCompanyNameOk

`func (o *RareDiseaseTrialsRow) GetSponsorCompanyNameOk() (*string, bool)`

GetSponsorCompanyNameOk returns a tuple with the SponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyName

`func (o *RareDiseaseTrialsRow) SetSponsorCompanyName(v string)`

SetSponsorCompanyName sets SponsorCompanyName field to given value.

### HasSponsorCompanyName

`func (o *RareDiseaseTrialsRow) HasSponsorCompanyName() bool`

HasSponsorCompanyName returns a boolean if a field has been set.

### SetSponsorCompanyNameNil

`func (o *RareDiseaseTrialsRow) SetSponsorCompanyNameNil(b bool)`

 SetSponsorCompanyNameNil sets the value for SponsorCompanyName to be an explicit nil

### UnsetSponsorCompanyName
`func (o *RareDiseaseTrialsRow) UnsetSponsorCompanyName()`

UnsetSponsorCompanyName ensures that no value is present for SponsorCompanyName, not even an explicit nil
### GetSponsorGroupId

`func (o *RareDiseaseTrialsRow) GetSponsorGroupId() string`

GetSponsorGroupId returns the SponsorGroupId field if non-nil, zero value otherwise.

### GetSponsorGroupIdOk

`func (o *RareDiseaseTrialsRow) GetSponsorGroupIdOk() (*string, bool)`

GetSponsorGroupIdOk returns a tuple with the SponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupId

`func (o *RareDiseaseTrialsRow) SetSponsorGroupId(v string)`

SetSponsorGroupId sets SponsorGroupId field to given value.

### HasSponsorGroupId

`func (o *RareDiseaseTrialsRow) HasSponsorGroupId() bool`

HasSponsorGroupId returns a boolean if a field has been set.

### SetSponsorGroupIdNil

`func (o *RareDiseaseTrialsRow) SetSponsorGroupIdNil(b bool)`

 SetSponsorGroupIdNil sets the value for SponsorGroupId to be an explicit nil

### UnsetSponsorGroupId
`func (o *RareDiseaseTrialsRow) UnsetSponsorGroupId()`

UnsetSponsorGroupId ensures that no value is present for SponsorGroupId, not even an explicit nil
### GetSponsorGroupName

`func (o *RareDiseaseTrialsRow) GetSponsorGroupName() string`

GetSponsorGroupName returns the SponsorGroupName field if non-nil, zero value otherwise.

### GetSponsorGroupNameOk

`func (o *RareDiseaseTrialsRow) GetSponsorGroupNameOk() (*string, bool)`

GetSponsorGroupNameOk returns a tuple with the SponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupName

`func (o *RareDiseaseTrialsRow) SetSponsorGroupName(v string)`

SetSponsorGroupName sets SponsorGroupName field to given value.

### HasSponsorGroupName

`func (o *RareDiseaseTrialsRow) HasSponsorGroupName() bool`

HasSponsorGroupName returns a boolean if a field has been set.

### SetSponsorGroupNameNil

`func (o *RareDiseaseTrialsRow) SetSponsorGroupNameNil(b bool)`

 SetSponsorGroupNameNil sets the value for SponsorGroupName to be an explicit nil

### UnsetSponsorGroupName
`func (o *RareDiseaseTrialsRow) UnsetSponsorGroupName()`

UnsetSponsorGroupName ensures that no value is present for SponsorGroupName, not even an explicit nil
### GetSponsorAgencyClass

`func (o *RareDiseaseTrialsRow) GetSponsorAgencyClass() string`

GetSponsorAgencyClass returns the SponsorAgencyClass field if non-nil, zero value otherwise.

### GetSponsorAgencyClassOk

`func (o *RareDiseaseTrialsRow) GetSponsorAgencyClassOk() (*string, bool)`

GetSponsorAgencyClassOk returns a tuple with the SponsorAgencyClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorAgencyClass

`func (o *RareDiseaseTrialsRow) SetSponsorAgencyClass(v string)`

SetSponsorAgencyClass sets SponsorAgencyClass field to given value.

### HasSponsorAgencyClass

`func (o *RareDiseaseTrialsRow) HasSponsorAgencyClass() bool`

HasSponsorAgencyClass returns a boolean if a field has been set.

### SetSponsorAgencyClassNil

`func (o *RareDiseaseTrialsRow) SetSponsorAgencyClassNil(b bool)`

 SetSponsorAgencyClassNil sets the value for SponsorAgencyClass to be an explicit nil

### UnsetSponsorAgencyClass
`func (o *RareDiseaseTrialsRow) UnsetSponsorAgencyClass()`

UnsetSponsorAgencyClass ensures that no value is present for SponsorAgencyClass, not even an explicit nil
### GetSponsorResolutionConfidence

`func (o *RareDiseaseTrialsRow) GetSponsorResolutionConfidence() string`

GetSponsorResolutionConfidence returns the SponsorResolutionConfidence field if non-nil, zero value otherwise.

### GetSponsorResolutionConfidenceOk

`func (o *RareDiseaseTrialsRow) GetSponsorResolutionConfidenceOk() (*string, bool)`

GetSponsorResolutionConfidenceOk returns a tuple with the SponsorResolutionConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorResolutionConfidence

`func (o *RareDiseaseTrialsRow) SetSponsorResolutionConfidence(v string)`

SetSponsorResolutionConfidence sets SponsorResolutionConfidence field to given value.

### HasSponsorResolutionConfidence

`func (o *RareDiseaseTrialsRow) HasSponsorResolutionConfidence() bool`

HasSponsorResolutionConfidence returns a boolean if a field has been set.

### SetSponsorResolutionConfidenceNil

`func (o *RareDiseaseTrialsRow) SetSponsorResolutionConfidenceNil(b bool)`

 SetSponsorResolutionConfidenceNil sets the value for SponsorResolutionConfidence to be an explicit nil

### UnsetSponsorResolutionConfidence
`func (o *RareDiseaseTrialsRow) UnsetSponsorResolutionConfidence()`

UnsetSponsorResolutionConfidence ensures that no value is present for SponsorResolutionConfidence, not even an explicit nil
### GetLicenseName

`func (o *RareDiseaseTrialsRow) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *RareDiseaseTrialsRow) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *RareDiseaseTrialsRow) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *RareDiseaseTrialsRow) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### SetLicenseNameNil

`func (o *RareDiseaseTrialsRow) SetLicenseNameNil(b bool)`

 SetLicenseNameNil sets the value for LicenseName to be an explicit nil

### UnsetLicenseName
`func (o *RareDiseaseTrialsRow) UnsetLicenseName()`

UnsetLicenseName ensures that no value is present for LicenseName, not even an explicit nil
### GetAttribution

`func (o *RareDiseaseTrialsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RareDiseaseTrialsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RareDiseaseTrialsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RareDiseaseTrialsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RareDiseaseTrialsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RareDiseaseTrialsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


