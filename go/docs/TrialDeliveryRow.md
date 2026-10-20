# TrialDeliveryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**LeadSponsorName** | Pointer to **NullableString** | The lead sponsor name as filed. | [optional] 
**OverallStatusSummary** | Pointer to **NullableString** | A coarse status label (Recruiting, Completed and so on). | [optional] 
**StudyType** | Pointer to **NullableString** | INTERVENTIONAL, OBSERVATIONAL or EXPANDED_ACCESS. | [optional] 
**Phase** | Pointer to **NullableString** | The registry&#39;s phase code (PHASE1, PHASE2, PHASE2/PHASE3, NA and so on). Null when the registry gives no phase (observational and expanded-access trials); phase_label carries the display label for those. | [optional] 
**PhaseLabel** | Pointer to **NullableString** | The trial phase as a display label (Phase 2, Phase 3, Observational, Not applicable, Expanded access, Unknown). The same trial&#39;s phase column holds the registry code. | [optional] 
**PlannedEnrollment** | Pointer to **NullableInt32** | Planned enrollment from the earliest snapshot typed ESTIMATED. | [optional] 
**ActualEnrollment** | Pointer to **NullableInt32** | Enrollment from the latest observation typed ACTUAL. | [optional] 
**EnrollmentType** | Pointer to **NullableString** | Enrollment type on the current record (ACTUAL or ESTIMATED). | [optional] 
**EnrollmentDelta** | Pointer to **NullableInt32** | Actual minus planned enrollment (null without both). | [optional] 
**EnrollmentAttainmentPct** | Pointer to **NullableFloat32** | Actual as a percentage of planned enrollment (null without both). | [optional] 
**EnrollmentMet** | Pointer to **NullableBool** | Whether actual enrollment reached the plan (null without both). | [optional] 
**EnrollmentDeltaKnown** | Pointer to **NullableBool** | Whether a planned and an actual enrollment are both recorded. | [optional] 
**PlannedBasis** | Pointer to **NullableString** | Where the planned enrollment came from: history_earliest_estimated, type_estimated_current or none. | [optional] 
**ActualBasis** | Pointer to **NullableString** | How the actual enrollment was taken (from the snapshot history, from the current record, or none). | [optional] 
**PlannedPrimaryCompletionDate** | Pointer to **NullableString** | Planned primary completion date from the earliest snapshot typed ESTIMATED. | [optional] 
**ActualPrimaryCompletionDate** | Pointer to **NullableString** | Primary completion date once typed ACTUAL. | [optional] 
**PlannedCompletionDate** | Pointer to **NullableString** | Planned completion date from the earliest snapshot typed ESTIMATED. | [optional] 
**ActualCompletionDate** | Pointer to **NullableString** | Completion date once typed ACTUAL. | [optional] 
**PrimaryCompletionDateType** | Pointer to **NullableString** | ACTUAL (a date that has occurred) or ESTIMATED (a planned date). The registry&#39;s older ANTICIPATED is published as ESTIMATED. | [optional] 
**CompletionDateType** | Pointer to **NullableString** | ACTUAL (a date that has occurred) or ESTIMATED (a planned date). The registry&#39;s older ANTICIPATED is published as ESTIMATED. | [optional] 
**PrimaryCompletionSlipDays** | Pointer to **NullableInt32** | Actual minus planned primary completion date; positive is later than planned (null without both). | [optional] 
**CompletionSlipDays** | Pointer to **NullableInt32** | Actual minus planned completion date; positive is later than planned (null without both). | [optional] 
**CompletionSlipped** | Pointer to **NullableBool** | Whether the primary completion came later than planned (null without both). | [optional] 
**HasCompletionDelta** | Pointer to **NullableBool** | Whether a planned and an actual primary completion date are both recorded. | [optional] 

## Methods

### NewTrialDeliveryRow

`func NewTrialDeliveryRow() *TrialDeliveryRow`

NewTrialDeliveryRow instantiates a new TrialDeliveryRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialDeliveryRowWithDefaults

`func NewTrialDeliveryRowWithDefaults() *TrialDeliveryRow`

NewTrialDeliveryRowWithDefaults instantiates a new TrialDeliveryRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *TrialDeliveryRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TrialDeliveryRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TrialDeliveryRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TrialDeliveryRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TrialDeliveryRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TrialDeliveryRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetLeadSponsorName

`func (o *TrialDeliveryRow) GetLeadSponsorName() string`

GetLeadSponsorName returns the LeadSponsorName field if non-nil, zero value otherwise.

### GetLeadSponsorNameOk

`func (o *TrialDeliveryRow) GetLeadSponsorNameOk() (*string, bool)`

GetLeadSponsorNameOk returns a tuple with the LeadSponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorName

`func (o *TrialDeliveryRow) SetLeadSponsorName(v string)`

SetLeadSponsorName sets LeadSponsorName field to given value.

### HasLeadSponsorName

`func (o *TrialDeliveryRow) HasLeadSponsorName() bool`

HasLeadSponsorName returns a boolean if a field has been set.

### SetLeadSponsorNameNil

`func (o *TrialDeliveryRow) SetLeadSponsorNameNil(b bool)`

 SetLeadSponsorNameNil sets the value for LeadSponsorName to be an explicit nil

### UnsetLeadSponsorName
`func (o *TrialDeliveryRow) UnsetLeadSponsorName()`

UnsetLeadSponsorName ensures that no value is present for LeadSponsorName, not even an explicit nil
### GetOverallStatusSummary

`func (o *TrialDeliveryRow) GetOverallStatusSummary() string`

GetOverallStatusSummary returns the OverallStatusSummary field if non-nil, zero value otherwise.

### GetOverallStatusSummaryOk

`func (o *TrialDeliveryRow) GetOverallStatusSummaryOk() (*string, bool)`

GetOverallStatusSummaryOk returns a tuple with the OverallStatusSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatusSummary

`func (o *TrialDeliveryRow) SetOverallStatusSummary(v string)`

SetOverallStatusSummary sets OverallStatusSummary field to given value.

### HasOverallStatusSummary

`func (o *TrialDeliveryRow) HasOverallStatusSummary() bool`

HasOverallStatusSummary returns a boolean if a field has been set.

### SetOverallStatusSummaryNil

`func (o *TrialDeliveryRow) SetOverallStatusSummaryNil(b bool)`

 SetOverallStatusSummaryNil sets the value for OverallStatusSummary to be an explicit nil

### UnsetOverallStatusSummary
`func (o *TrialDeliveryRow) UnsetOverallStatusSummary()`

UnsetOverallStatusSummary ensures that no value is present for OverallStatusSummary, not even an explicit nil
### GetStudyType

`func (o *TrialDeliveryRow) GetStudyType() string`

GetStudyType returns the StudyType field if non-nil, zero value otherwise.

### GetStudyTypeOk

`func (o *TrialDeliveryRow) GetStudyTypeOk() (*string, bool)`

GetStudyTypeOk returns a tuple with the StudyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyType

`func (o *TrialDeliveryRow) SetStudyType(v string)`

SetStudyType sets StudyType field to given value.

### HasStudyType

`func (o *TrialDeliveryRow) HasStudyType() bool`

HasStudyType returns a boolean if a field has been set.

### SetStudyTypeNil

`func (o *TrialDeliveryRow) SetStudyTypeNil(b bool)`

 SetStudyTypeNil sets the value for StudyType to be an explicit nil

### UnsetStudyType
`func (o *TrialDeliveryRow) UnsetStudyType()`

UnsetStudyType ensures that no value is present for StudyType, not even an explicit nil
### GetPhase

`func (o *TrialDeliveryRow) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *TrialDeliveryRow) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *TrialDeliveryRow) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *TrialDeliveryRow) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *TrialDeliveryRow) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *TrialDeliveryRow) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetPhaseLabel

`func (o *TrialDeliveryRow) GetPhaseLabel() string`

GetPhaseLabel returns the PhaseLabel field if non-nil, zero value otherwise.

### GetPhaseLabelOk

`func (o *TrialDeliveryRow) GetPhaseLabelOk() (*string, bool)`

GetPhaseLabelOk returns a tuple with the PhaseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseLabel

`func (o *TrialDeliveryRow) SetPhaseLabel(v string)`

SetPhaseLabel sets PhaseLabel field to given value.

### HasPhaseLabel

`func (o *TrialDeliveryRow) HasPhaseLabel() bool`

HasPhaseLabel returns a boolean if a field has been set.

### SetPhaseLabelNil

`func (o *TrialDeliveryRow) SetPhaseLabelNil(b bool)`

 SetPhaseLabelNil sets the value for PhaseLabel to be an explicit nil

### UnsetPhaseLabel
`func (o *TrialDeliveryRow) UnsetPhaseLabel()`

UnsetPhaseLabel ensures that no value is present for PhaseLabel, not even an explicit nil
### GetPlannedEnrollment

`func (o *TrialDeliveryRow) GetPlannedEnrollment() int32`

GetPlannedEnrollment returns the PlannedEnrollment field if non-nil, zero value otherwise.

### GetPlannedEnrollmentOk

`func (o *TrialDeliveryRow) GetPlannedEnrollmentOk() (*int32, bool)`

GetPlannedEnrollmentOk returns a tuple with the PlannedEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedEnrollment

`func (o *TrialDeliveryRow) SetPlannedEnrollment(v int32)`

SetPlannedEnrollment sets PlannedEnrollment field to given value.

### HasPlannedEnrollment

`func (o *TrialDeliveryRow) HasPlannedEnrollment() bool`

HasPlannedEnrollment returns a boolean if a field has been set.

### SetPlannedEnrollmentNil

`func (o *TrialDeliveryRow) SetPlannedEnrollmentNil(b bool)`

 SetPlannedEnrollmentNil sets the value for PlannedEnrollment to be an explicit nil

### UnsetPlannedEnrollment
`func (o *TrialDeliveryRow) UnsetPlannedEnrollment()`

UnsetPlannedEnrollment ensures that no value is present for PlannedEnrollment, not even an explicit nil
### GetActualEnrollment

`func (o *TrialDeliveryRow) GetActualEnrollment() int32`

GetActualEnrollment returns the ActualEnrollment field if non-nil, zero value otherwise.

### GetActualEnrollmentOk

`func (o *TrialDeliveryRow) GetActualEnrollmentOk() (*int32, bool)`

GetActualEnrollmentOk returns a tuple with the ActualEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualEnrollment

`func (o *TrialDeliveryRow) SetActualEnrollment(v int32)`

SetActualEnrollment sets ActualEnrollment field to given value.

### HasActualEnrollment

`func (o *TrialDeliveryRow) HasActualEnrollment() bool`

HasActualEnrollment returns a boolean if a field has been set.

### SetActualEnrollmentNil

`func (o *TrialDeliveryRow) SetActualEnrollmentNil(b bool)`

 SetActualEnrollmentNil sets the value for ActualEnrollment to be an explicit nil

### UnsetActualEnrollment
`func (o *TrialDeliveryRow) UnsetActualEnrollment()`

UnsetActualEnrollment ensures that no value is present for ActualEnrollment, not even an explicit nil
### GetEnrollmentType

`func (o *TrialDeliveryRow) GetEnrollmentType() string`

GetEnrollmentType returns the EnrollmentType field if non-nil, zero value otherwise.

### GetEnrollmentTypeOk

`func (o *TrialDeliveryRow) GetEnrollmentTypeOk() (*string, bool)`

GetEnrollmentTypeOk returns a tuple with the EnrollmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentType

`func (o *TrialDeliveryRow) SetEnrollmentType(v string)`

SetEnrollmentType sets EnrollmentType field to given value.

### HasEnrollmentType

`func (o *TrialDeliveryRow) HasEnrollmentType() bool`

HasEnrollmentType returns a boolean if a field has been set.

### SetEnrollmentTypeNil

`func (o *TrialDeliveryRow) SetEnrollmentTypeNil(b bool)`

 SetEnrollmentTypeNil sets the value for EnrollmentType to be an explicit nil

### UnsetEnrollmentType
`func (o *TrialDeliveryRow) UnsetEnrollmentType()`

UnsetEnrollmentType ensures that no value is present for EnrollmentType, not even an explicit nil
### GetEnrollmentDelta

`func (o *TrialDeliveryRow) GetEnrollmentDelta() int32`

GetEnrollmentDelta returns the EnrollmentDelta field if non-nil, zero value otherwise.

### GetEnrollmentDeltaOk

`func (o *TrialDeliveryRow) GetEnrollmentDeltaOk() (*int32, bool)`

GetEnrollmentDeltaOk returns a tuple with the EnrollmentDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentDelta

`func (o *TrialDeliveryRow) SetEnrollmentDelta(v int32)`

SetEnrollmentDelta sets EnrollmentDelta field to given value.

### HasEnrollmentDelta

`func (o *TrialDeliveryRow) HasEnrollmentDelta() bool`

HasEnrollmentDelta returns a boolean if a field has been set.

### SetEnrollmentDeltaNil

`func (o *TrialDeliveryRow) SetEnrollmentDeltaNil(b bool)`

 SetEnrollmentDeltaNil sets the value for EnrollmentDelta to be an explicit nil

### UnsetEnrollmentDelta
`func (o *TrialDeliveryRow) UnsetEnrollmentDelta()`

UnsetEnrollmentDelta ensures that no value is present for EnrollmentDelta, not even an explicit nil
### GetEnrollmentAttainmentPct

`func (o *TrialDeliveryRow) GetEnrollmentAttainmentPct() float32`

GetEnrollmentAttainmentPct returns the EnrollmentAttainmentPct field if non-nil, zero value otherwise.

### GetEnrollmentAttainmentPctOk

`func (o *TrialDeliveryRow) GetEnrollmentAttainmentPctOk() (*float32, bool)`

GetEnrollmentAttainmentPctOk returns a tuple with the EnrollmentAttainmentPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentAttainmentPct

`func (o *TrialDeliveryRow) SetEnrollmentAttainmentPct(v float32)`

SetEnrollmentAttainmentPct sets EnrollmentAttainmentPct field to given value.

### HasEnrollmentAttainmentPct

`func (o *TrialDeliveryRow) HasEnrollmentAttainmentPct() bool`

HasEnrollmentAttainmentPct returns a boolean if a field has been set.

### SetEnrollmentAttainmentPctNil

`func (o *TrialDeliveryRow) SetEnrollmentAttainmentPctNil(b bool)`

 SetEnrollmentAttainmentPctNil sets the value for EnrollmentAttainmentPct to be an explicit nil

### UnsetEnrollmentAttainmentPct
`func (o *TrialDeliveryRow) UnsetEnrollmentAttainmentPct()`

UnsetEnrollmentAttainmentPct ensures that no value is present for EnrollmentAttainmentPct, not even an explicit nil
### GetEnrollmentMet

`func (o *TrialDeliveryRow) GetEnrollmentMet() bool`

GetEnrollmentMet returns the EnrollmentMet field if non-nil, zero value otherwise.

### GetEnrollmentMetOk

`func (o *TrialDeliveryRow) GetEnrollmentMetOk() (*bool, bool)`

GetEnrollmentMetOk returns a tuple with the EnrollmentMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentMet

`func (o *TrialDeliveryRow) SetEnrollmentMet(v bool)`

SetEnrollmentMet sets EnrollmentMet field to given value.

### HasEnrollmentMet

`func (o *TrialDeliveryRow) HasEnrollmentMet() bool`

HasEnrollmentMet returns a boolean if a field has been set.

### SetEnrollmentMetNil

`func (o *TrialDeliveryRow) SetEnrollmentMetNil(b bool)`

 SetEnrollmentMetNil sets the value for EnrollmentMet to be an explicit nil

### UnsetEnrollmentMet
`func (o *TrialDeliveryRow) UnsetEnrollmentMet()`

UnsetEnrollmentMet ensures that no value is present for EnrollmentMet, not even an explicit nil
### GetEnrollmentDeltaKnown

`func (o *TrialDeliveryRow) GetEnrollmentDeltaKnown() bool`

GetEnrollmentDeltaKnown returns the EnrollmentDeltaKnown field if non-nil, zero value otherwise.

### GetEnrollmentDeltaKnownOk

`func (o *TrialDeliveryRow) GetEnrollmentDeltaKnownOk() (*bool, bool)`

GetEnrollmentDeltaKnownOk returns a tuple with the EnrollmentDeltaKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentDeltaKnown

`func (o *TrialDeliveryRow) SetEnrollmentDeltaKnown(v bool)`

SetEnrollmentDeltaKnown sets EnrollmentDeltaKnown field to given value.

### HasEnrollmentDeltaKnown

`func (o *TrialDeliveryRow) HasEnrollmentDeltaKnown() bool`

HasEnrollmentDeltaKnown returns a boolean if a field has been set.

### SetEnrollmentDeltaKnownNil

`func (o *TrialDeliveryRow) SetEnrollmentDeltaKnownNil(b bool)`

 SetEnrollmentDeltaKnownNil sets the value for EnrollmentDeltaKnown to be an explicit nil

### UnsetEnrollmentDeltaKnown
`func (o *TrialDeliveryRow) UnsetEnrollmentDeltaKnown()`

UnsetEnrollmentDeltaKnown ensures that no value is present for EnrollmentDeltaKnown, not even an explicit nil
### GetPlannedBasis

`func (o *TrialDeliveryRow) GetPlannedBasis() string`

GetPlannedBasis returns the PlannedBasis field if non-nil, zero value otherwise.

### GetPlannedBasisOk

`func (o *TrialDeliveryRow) GetPlannedBasisOk() (*string, bool)`

GetPlannedBasisOk returns a tuple with the PlannedBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedBasis

`func (o *TrialDeliveryRow) SetPlannedBasis(v string)`

SetPlannedBasis sets PlannedBasis field to given value.

### HasPlannedBasis

`func (o *TrialDeliveryRow) HasPlannedBasis() bool`

HasPlannedBasis returns a boolean if a field has been set.

### SetPlannedBasisNil

`func (o *TrialDeliveryRow) SetPlannedBasisNil(b bool)`

 SetPlannedBasisNil sets the value for PlannedBasis to be an explicit nil

### UnsetPlannedBasis
`func (o *TrialDeliveryRow) UnsetPlannedBasis()`

UnsetPlannedBasis ensures that no value is present for PlannedBasis, not even an explicit nil
### GetActualBasis

`func (o *TrialDeliveryRow) GetActualBasis() string`

GetActualBasis returns the ActualBasis field if non-nil, zero value otherwise.

### GetActualBasisOk

`func (o *TrialDeliveryRow) GetActualBasisOk() (*string, bool)`

GetActualBasisOk returns a tuple with the ActualBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualBasis

`func (o *TrialDeliveryRow) SetActualBasis(v string)`

SetActualBasis sets ActualBasis field to given value.

### HasActualBasis

`func (o *TrialDeliveryRow) HasActualBasis() bool`

HasActualBasis returns a boolean if a field has been set.

### SetActualBasisNil

`func (o *TrialDeliveryRow) SetActualBasisNil(b bool)`

 SetActualBasisNil sets the value for ActualBasis to be an explicit nil

### UnsetActualBasis
`func (o *TrialDeliveryRow) UnsetActualBasis()`

UnsetActualBasis ensures that no value is present for ActualBasis, not even an explicit nil
### GetPlannedPrimaryCompletionDate

`func (o *TrialDeliveryRow) GetPlannedPrimaryCompletionDate() string`

GetPlannedPrimaryCompletionDate returns the PlannedPrimaryCompletionDate field if non-nil, zero value otherwise.

### GetPlannedPrimaryCompletionDateOk

`func (o *TrialDeliveryRow) GetPlannedPrimaryCompletionDateOk() (*string, bool)`

GetPlannedPrimaryCompletionDateOk returns a tuple with the PlannedPrimaryCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedPrimaryCompletionDate

`func (o *TrialDeliveryRow) SetPlannedPrimaryCompletionDate(v string)`

SetPlannedPrimaryCompletionDate sets PlannedPrimaryCompletionDate field to given value.

### HasPlannedPrimaryCompletionDate

`func (o *TrialDeliveryRow) HasPlannedPrimaryCompletionDate() bool`

HasPlannedPrimaryCompletionDate returns a boolean if a field has been set.

### SetPlannedPrimaryCompletionDateNil

`func (o *TrialDeliveryRow) SetPlannedPrimaryCompletionDateNil(b bool)`

 SetPlannedPrimaryCompletionDateNil sets the value for PlannedPrimaryCompletionDate to be an explicit nil

### UnsetPlannedPrimaryCompletionDate
`func (o *TrialDeliveryRow) UnsetPlannedPrimaryCompletionDate()`

UnsetPlannedPrimaryCompletionDate ensures that no value is present for PlannedPrimaryCompletionDate, not even an explicit nil
### GetActualPrimaryCompletionDate

`func (o *TrialDeliveryRow) GetActualPrimaryCompletionDate() string`

GetActualPrimaryCompletionDate returns the ActualPrimaryCompletionDate field if non-nil, zero value otherwise.

### GetActualPrimaryCompletionDateOk

`func (o *TrialDeliveryRow) GetActualPrimaryCompletionDateOk() (*string, bool)`

GetActualPrimaryCompletionDateOk returns a tuple with the ActualPrimaryCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPrimaryCompletionDate

`func (o *TrialDeliveryRow) SetActualPrimaryCompletionDate(v string)`

SetActualPrimaryCompletionDate sets ActualPrimaryCompletionDate field to given value.

### HasActualPrimaryCompletionDate

`func (o *TrialDeliveryRow) HasActualPrimaryCompletionDate() bool`

HasActualPrimaryCompletionDate returns a boolean if a field has been set.

### SetActualPrimaryCompletionDateNil

`func (o *TrialDeliveryRow) SetActualPrimaryCompletionDateNil(b bool)`

 SetActualPrimaryCompletionDateNil sets the value for ActualPrimaryCompletionDate to be an explicit nil

### UnsetActualPrimaryCompletionDate
`func (o *TrialDeliveryRow) UnsetActualPrimaryCompletionDate()`

UnsetActualPrimaryCompletionDate ensures that no value is present for ActualPrimaryCompletionDate, not even an explicit nil
### GetPlannedCompletionDate

`func (o *TrialDeliveryRow) GetPlannedCompletionDate() string`

GetPlannedCompletionDate returns the PlannedCompletionDate field if non-nil, zero value otherwise.

### GetPlannedCompletionDateOk

`func (o *TrialDeliveryRow) GetPlannedCompletionDateOk() (*string, bool)`

GetPlannedCompletionDateOk returns a tuple with the PlannedCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedCompletionDate

`func (o *TrialDeliveryRow) SetPlannedCompletionDate(v string)`

SetPlannedCompletionDate sets PlannedCompletionDate field to given value.

### HasPlannedCompletionDate

`func (o *TrialDeliveryRow) HasPlannedCompletionDate() bool`

HasPlannedCompletionDate returns a boolean if a field has been set.

### SetPlannedCompletionDateNil

`func (o *TrialDeliveryRow) SetPlannedCompletionDateNil(b bool)`

 SetPlannedCompletionDateNil sets the value for PlannedCompletionDate to be an explicit nil

### UnsetPlannedCompletionDate
`func (o *TrialDeliveryRow) UnsetPlannedCompletionDate()`

UnsetPlannedCompletionDate ensures that no value is present for PlannedCompletionDate, not even an explicit nil
### GetActualCompletionDate

`func (o *TrialDeliveryRow) GetActualCompletionDate() string`

GetActualCompletionDate returns the ActualCompletionDate field if non-nil, zero value otherwise.

### GetActualCompletionDateOk

`func (o *TrialDeliveryRow) GetActualCompletionDateOk() (*string, bool)`

GetActualCompletionDateOk returns a tuple with the ActualCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualCompletionDate

`func (o *TrialDeliveryRow) SetActualCompletionDate(v string)`

SetActualCompletionDate sets ActualCompletionDate field to given value.

### HasActualCompletionDate

`func (o *TrialDeliveryRow) HasActualCompletionDate() bool`

HasActualCompletionDate returns a boolean if a field has been set.

### SetActualCompletionDateNil

`func (o *TrialDeliveryRow) SetActualCompletionDateNil(b bool)`

 SetActualCompletionDateNil sets the value for ActualCompletionDate to be an explicit nil

### UnsetActualCompletionDate
`func (o *TrialDeliveryRow) UnsetActualCompletionDate()`

UnsetActualCompletionDate ensures that no value is present for ActualCompletionDate, not even an explicit nil
### GetPrimaryCompletionDateType

`func (o *TrialDeliveryRow) GetPrimaryCompletionDateType() string`

GetPrimaryCompletionDateType returns the PrimaryCompletionDateType field if non-nil, zero value otherwise.

### GetPrimaryCompletionDateTypeOk

`func (o *TrialDeliveryRow) GetPrimaryCompletionDateTypeOk() (*string, bool)`

GetPrimaryCompletionDateTypeOk returns a tuple with the PrimaryCompletionDateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCompletionDateType

`func (o *TrialDeliveryRow) SetPrimaryCompletionDateType(v string)`

SetPrimaryCompletionDateType sets PrimaryCompletionDateType field to given value.

### HasPrimaryCompletionDateType

`func (o *TrialDeliveryRow) HasPrimaryCompletionDateType() bool`

HasPrimaryCompletionDateType returns a boolean if a field has been set.

### SetPrimaryCompletionDateTypeNil

`func (o *TrialDeliveryRow) SetPrimaryCompletionDateTypeNil(b bool)`

 SetPrimaryCompletionDateTypeNil sets the value for PrimaryCompletionDateType to be an explicit nil

### UnsetPrimaryCompletionDateType
`func (o *TrialDeliveryRow) UnsetPrimaryCompletionDateType()`

UnsetPrimaryCompletionDateType ensures that no value is present for PrimaryCompletionDateType, not even an explicit nil
### GetCompletionDateType

`func (o *TrialDeliveryRow) GetCompletionDateType() string`

GetCompletionDateType returns the CompletionDateType field if non-nil, zero value otherwise.

### GetCompletionDateTypeOk

`func (o *TrialDeliveryRow) GetCompletionDateTypeOk() (*string, bool)`

GetCompletionDateTypeOk returns a tuple with the CompletionDateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateType

`func (o *TrialDeliveryRow) SetCompletionDateType(v string)`

SetCompletionDateType sets CompletionDateType field to given value.

### HasCompletionDateType

`func (o *TrialDeliveryRow) HasCompletionDateType() bool`

HasCompletionDateType returns a boolean if a field has been set.

### SetCompletionDateTypeNil

`func (o *TrialDeliveryRow) SetCompletionDateTypeNil(b bool)`

 SetCompletionDateTypeNil sets the value for CompletionDateType to be an explicit nil

### UnsetCompletionDateType
`func (o *TrialDeliveryRow) UnsetCompletionDateType()`

UnsetCompletionDateType ensures that no value is present for CompletionDateType, not even an explicit nil
### GetPrimaryCompletionSlipDays

`func (o *TrialDeliveryRow) GetPrimaryCompletionSlipDays() int32`

GetPrimaryCompletionSlipDays returns the PrimaryCompletionSlipDays field if non-nil, zero value otherwise.

### GetPrimaryCompletionSlipDaysOk

`func (o *TrialDeliveryRow) GetPrimaryCompletionSlipDaysOk() (*int32, bool)`

GetPrimaryCompletionSlipDaysOk returns a tuple with the PrimaryCompletionSlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCompletionSlipDays

`func (o *TrialDeliveryRow) SetPrimaryCompletionSlipDays(v int32)`

SetPrimaryCompletionSlipDays sets PrimaryCompletionSlipDays field to given value.

### HasPrimaryCompletionSlipDays

`func (o *TrialDeliveryRow) HasPrimaryCompletionSlipDays() bool`

HasPrimaryCompletionSlipDays returns a boolean if a field has been set.

### SetPrimaryCompletionSlipDaysNil

`func (o *TrialDeliveryRow) SetPrimaryCompletionSlipDaysNil(b bool)`

 SetPrimaryCompletionSlipDaysNil sets the value for PrimaryCompletionSlipDays to be an explicit nil

### UnsetPrimaryCompletionSlipDays
`func (o *TrialDeliveryRow) UnsetPrimaryCompletionSlipDays()`

UnsetPrimaryCompletionSlipDays ensures that no value is present for PrimaryCompletionSlipDays, not even an explicit nil
### GetCompletionSlipDays

`func (o *TrialDeliveryRow) GetCompletionSlipDays() int32`

GetCompletionSlipDays returns the CompletionSlipDays field if non-nil, zero value otherwise.

### GetCompletionSlipDaysOk

`func (o *TrialDeliveryRow) GetCompletionSlipDaysOk() (*int32, bool)`

GetCompletionSlipDaysOk returns a tuple with the CompletionSlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionSlipDays

`func (o *TrialDeliveryRow) SetCompletionSlipDays(v int32)`

SetCompletionSlipDays sets CompletionSlipDays field to given value.

### HasCompletionSlipDays

`func (o *TrialDeliveryRow) HasCompletionSlipDays() bool`

HasCompletionSlipDays returns a boolean if a field has been set.

### SetCompletionSlipDaysNil

`func (o *TrialDeliveryRow) SetCompletionSlipDaysNil(b bool)`

 SetCompletionSlipDaysNil sets the value for CompletionSlipDays to be an explicit nil

### UnsetCompletionSlipDays
`func (o *TrialDeliveryRow) UnsetCompletionSlipDays()`

UnsetCompletionSlipDays ensures that no value is present for CompletionSlipDays, not even an explicit nil
### GetCompletionSlipped

`func (o *TrialDeliveryRow) GetCompletionSlipped() bool`

GetCompletionSlipped returns the CompletionSlipped field if non-nil, zero value otherwise.

### GetCompletionSlippedOk

`func (o *TrialDeliveryRow) GetCompletionSlippedOk() (*bool, bool)`

GetCompletionSlippedOk returns a tuple with the CompletionSlipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionSlipped

`func (o *TrialDeliveryRow) SetCompletionSlipped(v bool)`

SetCompletionSlipped sets CompletionSlipped field to given value.

### HasCompletionSlipped

`func (o *TrialDeliveryRow) HasCompletionSlipped() bool`

HasCompletionSlipped returns a boolean if a field has been set.

### SetCompletionSlippedNil

`func (o *TrialDeliveryRow) SetCompletionSlippedNil(b bool)`

 SetCompletionSlippedNil sets the value for CompletionSlipped to be an explicit nil

### UnsetCompletionSlipped
`func (o *TrialDeliveryRow) UnsetCompletionSlipped()`

UnsetCompletionSlipped ensures that no value is present for CompletionSlipped, not even an explicit nil
### GetHasCompletionDelta

`func (o *TrialDeliveryRow) GetHasCompletionDelta() bool`

GetHasCompletionDelta returns the HasCompletionDelta field if non-nil, zero value otherwise.

### GetHasCompletionDeltaOk

`func (o *TrialDeliveryRow) GetHasCompletionDeltaOk() (*bool, bool)`

GetHasCompletionDeltaOk returns a tuple with the HasCompletionDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCompletionDelta

`func (o *TrialDeliveryRow) SetHasCompletionDelta(v bool)`

SetHasCompletionDelta sets HasCompletionDelta field to given value.

### HasHasCompletionDelta

`func (o *TrialDeliveryRow) HasHasCompletionDelta() bool`

HasHasCompletionDelta returns a boolean if a field has been set.

### SetHasCompletionDeltaNil

`func (o *TrialDeliveryRow) SetHasCompletionDeltaNil(b bool)`

 SetHasCompletionDeltaNil sets the value for HasCompletionDelta to be an explicit nil

### UnsetHasCompletionDelta
`func (o *TrialDeliveryRow) UnsetHasCompletionDelta()`

UnsetHasCompletionDelta ensures that no value is present for HasCompletionDelta, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


