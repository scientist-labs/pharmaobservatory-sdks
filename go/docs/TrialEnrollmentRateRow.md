# TrialEnrollmentRateRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**Phase** | Pointer to **NullableString** | The registry&#39;s phase code (PHASE1, PHASE2, PHASE2/PHASE3, NA and so on). Null when the registry gives no phase (observational and expanded-access trials); phase_label carries the display label for those. | [optional] 
**PhaseLabel** | Pointer to **NullableString** | The trial phase as a display label (Phase 2, Phase 3, Observational, Not applicable, Expanded access, Unknown). The same trial&#39;s phase column holds the registry code. | [optional] 
**LeadSponsorName** | Pointer to **NullableString** | The lead sponsor name as filed. | [optional] 
**OverallStatusSummary** | Pointer to **NullableString** | A coarse status label (Completed and so on). | [optional] 
**ActualEnrollment** | Pointer to **NullableInt32** | Enrollment typed ACTUAL. | [optional] 
**EnrollmentType** | Pointer to **NullableString** | ACTUAL or ESTIMATED. | [optional] 
**EnrollmentIsPlausible** | Pointer to **NullableBool** | False when the registered enrollment is a placeholder value rather than a count. | [optional] 
**ActiveSites** | Pointer to **NullableInt32** | Number of sites the rate is computed over. | [optional] 
**FacilityCountryCount** | Pointer to **NullableInt32** | Number of countries with a listed site. | [optional] 
**HasUsFacility** | Pointer to **NullableBool** | Whether any listed site is in the United States. | [optional] 
**StartDate** | Pointer to **NullableString** | Trial start date. | [optional] 
**PrimaryCompletionDate** | Pointer to **NullableString** | Primary completion date. | [optional] 
**CompletionDate** | Pointer to **NullableString** | Completion date. | [optional] 
**PrimaryCompletionDateType** | Pointer to **NullableString** | ACTUAL (a date that has occurred) or ESTIMATED (a planned date). The registry&#39;s older ANTICIPATED is published as ESTIMATED. | [optional] 
**CompletionDateType** | Pointer to **NullableString** | ACTUAL (a date that has occurred) or ESTIMATED (a planned date). The registry&#39;s older ANTICIPATED is published as ESTIMATED. | [optional] 
**IsRateComputable** | Pointer to **NullableBool** | Whether the inputs support a rate. | [optional] 
**RateFlagReason** | Pointer to **NullableString** | Why a rate could not be computed (null when it could). | [optional] 
**DurationMonths** | Pointer to **NullableFloat32** | Months from start to primary completion. | [optional] 
**DurationMonthsOverall** | Pointer to **NullableFloat32** | Months from start to completion. | [optional] 
**PtsPerMonth** | Pointer to **NullableFloat32** | Participants enrolled per month. | [optional] 
**PtsPerSiteMonth** | Pointer to **NullableFloat32** | Participants enrolled per site per month. | [optional] 

## Methods

### NewTrialEnrollmentRateRow

`func NewTrialEnrollmentRateRow() *TrialEnrollmentRateRow`

NewTrialEnrollmentRateRow instantiates a new TrialEnrollmentRateRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialEnrollmentRateRowWithDefaults

`func NewTrialEnrollmentRateRowWithDefaults() *TrialEnrollmentRateRow`

NewTrialEnrollmentRateRowWithDefaults instantiates a new TrialEnrollmentRateRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *TrialEnrollmentRateRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TrialEnrollmentRateRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TrialEnrollmentRateRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TrialEnrollmentRateRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TrialEnrollmentRateRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TrialEnrollmentRateRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetPhase

`func (o *TrialEnrollmentRateRow) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *TrialEnrollmentRateRow) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *TrialEnrollmentRateRow) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *TrialEnrollmentRateRow) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *TrialEnrollmentRateRow) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *TrialEnrollmentRateRow) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetPhaseLabel

`func (o *TrialEnrollmentRateRow) GetPhaseLabel() string`

GetPhaseLabel returns the PhaseLabel field if non-nil, zero value otherwise.

### GetPhaseLabelOk

`func (o *TrialEnrollmentRateRow) GetPhaseLabelOk() (*string, bool)`

GetPhaseLabelOk returns a tuple with the PhaseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseLabel

`func (o *TrialEnrollmentRateRow) SetPhaseLabel(v string)`

SetPhaseLabel sets PhaseLabel field to given value.

### HasPhaseLabel

`func (o *TrialEnrollmentRateRow) HasPhaseLabel() bool`

HasPhaseLabel returns a boolean if a field has been set.

### SetPhaseLabelNil

`func (o *TrialEnrollmentRateRow) SetPhaseLabelNil(b bool)`

 SetPhaseLabelNil sets the value for PhaseLabel to be an explicit nil

### UnsetPhaseLabel
`func (o *TrialEnrollmentRateRow) UnsetPhaseLabel()`

UnsetPhaseLabel ensures that no value is present for PhaseLabel, not even an explicit nil
### GetLeadSponsorName

`func (o *TrialEnrollmentRateRow) GetLeadSponsorName() string`

GetLeadSponsorName returns the LeadSponsorName field if non-nil, zero value otherwise.

### GetLeadSponsorNameOk

`func (o *TrialEnrollmentRateRow) GetLeadSponsorNameOk() (*string, bool)`

GetLeadSponsorNameOk returns a tuple with the LeadSponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorName

`func (o *TrialEnrollmentRateRow) SetLeadSponsorName(v string)`

SetLeadSponsorName sets LeadSponsorName field to given value.

### HasLeadSponsorName

`func (o *TrialEnrollmentRateRow) HasLeadSponsorName() bool`

HasLeadSponsorName returns a boolean if a field has been set.

### SetLeadSponsorNameNil

`func (o *TrialEnrollmentRateRow) SetLeadSponsorNameNil(b bool)`

 SetLeadSponsorNameNil sets the value for LeadSponsorName to be an explicit nil

### UnsetLeadSponsorName
`func (o *TrialEnrollmentRateRow) UnsetLeadSponsorName()`

UnsetLeadSponsorName ensures that no value is present for LeadSponsorName, not even an explicit nil
### GetOverallStatusSummary

`func (o *TrialEnrollmentRateRow) GetOverallStatusSummary() string`

GetOverallStatusSummary returns the OverallStatusSummary field if non-nil, zero value otherwise.

### GetOverallStatusSummaryOk

`func (o *TrialEnrollmentRateRow) GetOverallStatusSummaryOk() (*string, bool)`

GetOverallStatusSummaryOk returns a tuple with the OverallStatusSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatusSummary

`func (o *TrialEnrollmentRateRow) SetOverallStatusSummary(v string)`

SetOverallStatusSummary sets OverallStatusSummary field to given value.

### HasOverallStatusSummary

`func (o *TrialEnrollmentRateRow) HasOverallStatusSummary() bool`

HasOverallStatusSummary returns a boolean if a field has been set.

### SetOverallStatusSummaryNil

`func (o *TrialEnrollmentRateRow) SetOverallStatusSummaryNil(b bool)`

 SetOverallStatusSummaryNil sets the value for OverallStatusSummary to be an explicit nil

### UnsetOverallStatusSummary
`func (o *TrialEnrollmentRateRow) UnsetOverallStatusSummary()`

UnsetOverallStatusSummary ensures that no value is present for OverallStatusSummary, not even an explicit nil
### GetActualEnrollment

`func (o *TrialEnrollmentRateRow) GetActualEnrollment() int32`

GetActualEnrollment returns the ActualEnrollment field if non-nil, zero value otherwise.

### GetActualEnrollmentOk

`func (o *TrialEnrollmentRateRow) GetActualEnrollmentOk() (*int32, bool)`

GetActualEnrollmentOk returns a tuple with the ActualEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualEnrollment

`func (o *TrialEnrollmentRateRow) SetActualEnrollment(v int32)`

SetActualEnrollment sets ActualEnrollment field to given value.

### HasActualEnrollment

`func (o *TrialEnrollmentRateRow) HasActualEnrollment() bool`

HasActualEnrollment returns a boolean if a field has been set.

### SetActualEnrollmentNil

`func (o *TrialEnrollmentRateRow) SetActualEnrollmentNil(b bool)`

 SetActualEnrollmentNil sets the value for ActualEnrollment to be an explicit nil

### UnsetActualEnrollment
`func (o *TrialEnrollmentRateRow) UnsetActualEnrollment()`

UnsetActualEnrollment ensures that no value is present for ActualEnrollment, not even an explicit nil
### GetEnrollmentType

`func (o *TrialEnrollmentRateRow) GetEnrollmentType() string`

GetEnrollmentType returns the EnrollmentType field if non-nil, zero value otherwise.

### GetEnrollmentTypeOk

`func (o *TrialEnrollmentRateRow) GetEnrollmentTypeOk() (*string, bool)`

GetEnrollmentTypeOk returns a tuple with the EnrollmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentType

`func (o *TrialEnrollmentRateRow) SetEnrollmentType(v string)`

SetEnrollmentType sets EnrollmentType field to given value.

### HasEnrollmentType

`func (o *TrialEnrollmentRateRow) HasEnrollmentType() bool`

HasEnrollmentType returns a boolean if a field has been set.

### SetEnrollmentTypeNil

`func (o *TrialEnrollmentRateRow) SetEnrollmentTypeNil(b bool)`

 SetEnrollmentTypeNil sets the value for EnrollmentType to be an explicit nil

### UnsetEnrollmentType
`func (o *TrialEnrollmentRateRow) UnsetEnrollmentType()`

UnsetEnrollmentType ensures that no value is present for EnrollmentType, not even an explicit nil
### GetEnrollmentIsPlausible

`func (o *TrialEnrollmentRateRow) GetEnrollmentIsPlausible() bool`

GetEnrollmentIsPlausible returns the EnrollmentIsPlausible field if non-nil, zero value otherwise.

### GetEnrollmentIsPlausibleOk

`func (o *TrialEnrollmentRateRow) GetEnrollmentIsPlausibleOk() (*bool, bool)`

GetEnrollmentIsPlausibleOk returns a tuple with the EnrollmentIsPlausible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentIsPlausible

`func (o *TrialEnrollmentRateRow) SetEnrollmentIsPlausible(v bool)`

SetEnrollmentIsPlausible sets EnrollmentIsPlausible field to given value.

### HasEnrollmentIsPlausible

`func (o *TrialEnrollmentRateRow) HasEnrollmentIsPlausible() bool`

HasEnrollmentIsPlausible returns a boolean if a field has been set.

### SetEnrollmentIsPlausibleNil

`func (o *TrialEnrollmentRateRow) SetEnrollmentIsPlausibleNil(b bool)`

 SetEnrollmentIsPlausibleNil sets the value for EnrollmentIsPlausible to be an explicit nil

### UnsetEnrollmentIsPlausible
`func (o *TrialEnrollmentRateRow) UnsetEnrollmentIsPlausible()`

UnsetEnrollmentIsPlausible ensures that no value is present for EnrollmentIsPlausible, not even an explicit nil
### GetActiveSites

`func (o *TrialEnrollmentRateRow) GetActiveSites() int32`

GetActiveSites returns the ActiveSites field if non-nil, zero value otherwise.

### GetActiveSitesOk

`func (o *TrialEnrollmentRateRow) GetActiveSitesOk() (*int32, bool)`

GetActiveSitesOk returns a tuple with the ActiveSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSites

`func (o *TrialEnrollmentRateRow) SetActiveSites(v int32)`

SetActiveSites sets ActiveSites field to given value.

### HasActiveSites

`func (o *TrialEnrollmentRateRow) HasActiveSites() bool`

HasActiveSites returns a boolean if a field has been set.

### SetActiveSitesNil

`func (o *TrialEnrollmentRateRow) SetActiveSitesNil(b bool)`

 SetActiveSitesNil sets the value for ActiveSites to be an explicit nil

### UnsetActiveSites
`func (o *TrialEnrollmentRateRow) UnsetActiveSites()`

UnsetActiveSites ensures that no value is present for ActiveSites, not even an explicit nil
### GetFacilityCountryCount

`func (o *TrialEnrollmentRateRow) GetFacilityCountryCount() int32`

GetFacilityCountryCount returns the FacilityCountryCount field if non-nil, zero value otherwise.

### GetFacilityCountryCountOk

`func (o *TrialEnrollmentRateRow) GetFacilityCountryCountOk() (*int32, bool)`

GetFacilityCountryCountOk returns a tuple with the FacilityCountryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCountryCount

`func (o *TrialEnrollmentRateRow) SetFacilityCountryCount(v int32)`

SetFacilityCountryCount sets FacilityCountryCount field to given value.

### HasFacilityCountryCount

`func (o *TrialEnrollmentRateRow) HasFacilityCountryCount() bool`

HasFacilityCountryCount returns a boolean if a field has been set.

### SetFacilityCountryCountNil

`func (o *TrialEnrollmentRateRow) SetFacilityCountryCountNil(b bool)`

 SetFacilityCountryCountNil sets the value for FacilityCountryCount to be an explicit nil

### UnsetFacilityCountryCount
`func (o *TrialEnrollmentRateRow) UnsetFacilityCountryCount()`

UnsetFacilityCountryCount ensures that no value is present for FacilityCountryCount, not even an explicit nil
### GetHasUsFacility

`func (o *TrialEnrollmentRateRow) GetHasUsFacility() bool`

GetHasUsFacility returns the HasUsFacility field if non-nil, zero value otherwise.

### GetHasUsFacilityOk

`func (o *TrialEnrollmentRateRow) GetHasUsFacilityOk() (*bool, bool)`

GetHasUsFacilityOk returns a tuple with the HasUsFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUsFacility

`func (o *TrialEnrollmentRateRow) SetHasUsFacility(v bool)`

SetHasUsFacility sets HasUsFacility field to given value.

### HasHasUsFacility

`func (o *TrialEnrollmentRateRow) HasHasUsFacility() bool`

HasHasUsFacility returns a boolean if a field has been set.

### SetHasUsFacilityNil

`func (o *TrialEnrollmentRateRow) SetHasUsFacilityNil(b bool)`

 SetHasUsFacilityNil sets the value for HasUsFacility to be an explicit nil

### UnsetHasUsFacility
`func (o *TrialEnrollmentRateRow) UnsetHasUsFacility()`

UnsetHasUsFacility ensures that no value is present for HasUsFacility, not even an explicit nil
### GetStartDate

`func (o *TrialEnrollmentRateRow) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TrialEnrollmentRateRow) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TrialEnrollmentRateRow) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TrialEnrollmentRateRow) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *TrialEnrollmentRateRow) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *TrialEnrollmentRateRow) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetPrimaryCompletionDate

`func (o *TrialEnrollmentRateRow) GetPrimaryCompletionDate() string`

GetPrimaryCompletionDate returns the PrimaryCompletionDate field if non-nil, zero value otherwise.

### GetPrimaryCompletionDateOk

`func (o *TrialEnrollmentRateRow) GetPrimaryCompletionDateOk() (*string, bool)`

GetPrimaryCompletionDateOk returns a tuple with the PrimaryCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCompletionDate

`func (o *TrialEnrollmentRateRow) SetPrimaryCompletionDate(v string)`

SetPrimaryCompletionDate sets PrimaryCompletionDate field to given value.

### HasPrimaryCompletionDate

`func (o *TrialEnrollmentRateRow) HasPrimaryCompletionDate() bool`

HasPrimaryCompletionDate returns a boolean if a field has been set.

### SetPrimaryCompletionDateNil

`func (o *TrialEnrollmentRateRow) SetPrimaryCompletionDateNil(b bool)`

 SetPrimaryCompletionDateNil sets the value for PrimaryCompletionDate to be an explicit nil

### UnsetPrimaryCompletionDate
`func (o *TrialEnrollmentRateRow) UnsetPrimaryCompletionDate()`

UnsetPrimaryCompletionDate ensures that no value is present for PrimaryCompletionDate, not even an explicit nil
### GetCompletionDate

`func (o *TrialEnrollmentRateRow) GetCompletionDate() string`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *TrialEnrollmentRateRow) GetCompletionDateOk() (*string, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *TrialEnrollmentRateRow) SetCompletionDate(v string)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *TrialEnrollmentRateRow) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### SetCompletionDateNil

`func (o *TrialEnrollmentRateRow) SetCompletionDateNil(b bool)`

 SetCompletionDateNil sets the value for CompletionDate to be an explicit nil

### UnsetCompletionDate
`func (o *TrialEnrollmentRateRow) UnsetCompletionDate()`

UnsetCompletionDate ensures that no value is present for CompletionDate, not even an explicit nil
### GetPrimaryCompletionDateType

`func (o *TrialEnrollmentRateRow) GetPrimaryCompletionDateType() string`

GetPrimaryCompletionDateType returns the PrimaryCompletionDateType field if non-nil, zero value otherwise.

### GetPrimaryCompletionDateTypeOk

`func (o *TrialEnrollmentRateRow) GetPrimaryCompletionDateTypeOk() (*string, bool)`

GetPrimaryCompletionDateTypeOk returns a tuple with the PrimaryCompletionDateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCompletionDateType

`func (o *TrialEnrollmentRateRow) SetPrimaryCompletionDateType(v string)`

SetPrimaryCompletionDateType sets PrimaryCompletionDateType field to given value.

### HasPrimaryCompletionDateType

`func (o *TrialEnrollmentRateRow) HasPrimaryCompletionDateType() bool`

HasPrimaryCompletionDateType returns a boolean if a field has been set.

### SetPrimaryCompletionDateTypeNil

`func (o *TrialEnrollmentRateRow) SetPrimaryCompletionDateTypeNil(b bool)`

 SetPrimaryCompletionDateTypeNil sets the value for PrimaryCompletionDateType to be an explicit nil

### UnsetPrimaryCompletionDateType
`func (o *TrialEnrollmentRateRow) UnsetPrimaryCompletionDateType()`

UnsetPrimaryCompletionDateType ensures that no value is present for PrimaryCompletionDateType, not even an explicit nil
### GetCompletionDateType

`func (o *TrialEnrollmentRateRow) GetCompletionDateType() string`

GetCompletionDateType returns the CompletionDateType field if non-nil, zero value otherwise.

### GetCompletionDateTypeOk

`func (o *TrialEnrollmentRateRow) GetCompletionDateTypeOk() (*string, bool)`

GetCompletionDateTypeOk returns a tuple with the CompletionDateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateType

`func (o *TrialEnrollmentRateRow) SetCompletionDateType(v string)`

SetCompletionDateType sets CompletionDateType field to given value.

### HasCompletionDateType

`func (o *TrialEnrollmentRateRow) HasCompletionDateType() bool`

HasCompletionDateType returns a boolean if a field has been set.

### SetCompletionDateTypeNil

`func (o *TrialEnrollmentRateRow) SetCompletionDateTypeNil(b bool)`

 SetCompletionDateTypeNil sets the value for CompletionDateType to be an explicit nil

### UnsetCompletionDateType
`func (o *TrialEnrollmentRateRow) UnsetCompletionDateType()`

UnsetCompletionDateType ensures that no value is present for CompletionDateType, not even an explicit nil
### GetIsRateComputable

`func (o *TrialEnrollmentRateRow) GetIsRateComputable() bool`

GetIsRateComputable returns the IsRateComputable field if non-nil, zero value otherwise.

### GetIsRateComputableOk

`func (o *TrialEnrollmentRateRow) GetIsRateComputableOk() (*bool, bool)`

GetIsRateComputableOk returns a tuple with the IsRateComputable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRateComputable

`func (o *TrialEnrollmentRateRow) SetIsRateComputable(v bool)`

SetIsRateComputable sets IsRateComputable field to given value.

### HasIsRateComputable

`func (o *TrialEnrollmentRateRow) HasIsRateComputable() bool`

HasIsRateComputable returns a boolean if a field has been set.

### SetIsRateComputableNil

`func (o *TrialEnrollmentRateRow) SetIsRateComputableNil(b bool)`

 SetIsRateComputableNil sets the value for IsRateComputable to be an explicit nil

### UnsetIsRateComputable
`func (o *TrialEnrollmentRateRow) UnsetIsRateComputable()`

UnsetIsRateComputable ensures that no value is present for IsRateComputable, not even an explicit nil
### GetRateFlagReason

`func (o *TrialEnrollmentRateRow) GetRateFlagReason() string`

GetRateFlagReason returns the RateFlagReason field if non-nil, zero value otherwise.

### GetRateFlagReasonOk

`func (o *TrialEnrollmentRateRow) GetRateFlagReasonOk() (*string, bool)`

GetRateFlagReasonOk returns a tuple with the RateFlagReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateFlagReason

`func (o *TrialEnrollmentRateRow) SetRateFlagReason(v string)`

SetRateFlagReason sets RateFlagReason field to given value.

### HasRateFlagReason

`func (o *TrialEnrollmentRateRow) HasRateFlagReason() bool`

HasRateFlagReason returns a boolean if a field has been set.

### SetRateFlagReasonNil

`func (o *TrialEnrollmentRateRow) SetRateFlagReasonNil(b bool)`

 SetRateFlagReasonNil sets the value for RateFlagReason to be an explicit nil

### UnsetRateFlagReason
`func (o *TrialEnrollmentRateRow) UnsetRateFlagReason()`

UnsetRateFlagReason ensures that no value is present for RateFlagReason, not even an explicit nil
### GetDurationMonths

`func (o *TrialEnrollmentRateRow) GetDurationMonths() float32`

GetDurationMonths returns the DurationMonths field if non-nil, zero value otherwise.

### GetDurationMonthsOk

`func (o *TrialEnrollmentRateRow) GetDurationMonthsOk() (*float32, bool)`

GetDurationMonthsOk returns a tuple with the DurationMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMonths

`func (o *TrialEnrollmentRateRow) SetDurationMonths(v float32)`

SetDurationMonths sets DurationMonths field to given value.

### HasDurationMonths

`func (o *TrialEnrollmentRateRow) HasDurationMonths() bool`

HasDurationMonths returns a boolean if a field has been set.

### SetDurationMonthsNil

`func (o *TrialEnrollmentRateRow) SetDurationMonthsNil(b bool)`

 SetDurationMonthsNil sets the value for DurationMonths to be an explicit nil

### UnsetDurationMonths
`func (o *TrialEnrollmentRateRow) UnsetDurationMonths()`

UnsetDurationMonths ensures that no value is present for DurationMonths, not even an explicit nil
### GetDurationMonthsOverall

`func (o *TrialEnrollmentRateRow) GetDurationMonthsOverall() float32`

GetDurationMonthsOverall returns the DurationMonthsOverall field if non-nil, zero value otherwise.

### GetDurationMonthsOverallOk

`func (o *TrialEnrollmentRateRow) GetDurationMonthsOverallOk() (*float32, bool)`

GetDurationMonthsOverallOk returns a tuple with the DurationMonthsOverall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMonthsOverall

`func (o *TrialEnrollmentRateRow) SetDurationMonthsOverall(v float32)`

SetDurationMonthsOverall sets DurationMonthsOverall field to given value.

### HasDurationMonthsOverall

`func (o *TrialEnrollmentRateRow) HasDurationMonthsOverall() bool`

HasDurationMonthsOverall returns a boolean if a field has been set.

### SetDurationMonthsOverallNil

`func (o *TrialEnrollmentRateRow) SetDurationMonthsOverallNil(b bool)`

 SetDurationMonthsOverallNil sets the value for DurationMonthsOverall to be an explicit nil

### UnsetDurationMonthsOverall
`func (o *TrialEnrollmentRateRow) UnsetDurationMonthsOverall()`

UnsetDurationMonthsOverall ensures that no value is present for DurationMonthsOverall, not even an explicit nil
### GetPtsPerMonth

`func (o *TrialEnrollmentRateRow) GetPtsPerMonth() float32`

GetPtsPerMonth returns the PtsPerMonth field if non-nil, zero value otherwise.

### GetPtsPerMonthOk

`func (o *TrialEnrollmentRateRow) GetPtsPerMonthOk() (*float32, bool)`

GetPtsPerMonthOk returns a tuple with the PtsPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtsPerMonth

`func (o *TrialEnrollmentRateRow) SetPtsPerMonth(v float32)`

SetPtsPerMonth sets PtsPerMonth field to given value.

### HasPtsPerMonth

`func (o *TrialEnrollmentRateRow) HasPtsPerMonth() bool`

HasPtsPerMonth returns a boolean if a field has been set.

### SetPtsPerMonthNil

`func (o *TrialEnrollmentRateRow) SetPtsPerMonthNil(b bool)`

 SetPtsPerMonthNil sets the value for PtsPerMonth to be an explicit nil

### UnsetPtsPerMonth
`func (o *TrialEnrollmentRateRow) UnsetPtsPerMonth()`

UnsetPtsPerMonth ensures that no value is present for PtsPerMonth, not even an explicit nil
### GetPtsPerSiteMonth

`func (o *TrialEnrollmentRateRow) GetPtsPerSiteMonth() float32`

GetPtsPerSiteMonth returns the PtsPerSiteMonth field if non-nil, zero value otherwise.

### GetPtsPerSiteMonthOk

`func (o *TrialEnrollmentRateRow) GetPtsPerSiteMonthOk() (*float32, bool)`

GetPtsPerSiteMonthOk returns a tuple with the PtsPerSiteMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtsPerSiteMonth

`func (o *TrialEnrollmentRateRow) SetPtsPerSiteMonth(v float32)`

SetPtsPerSiteMonth sets PtsPerSiteMonth field to given value.

### HasPtsPerSiteMonth

`func (o *TrialEnrollmentRateRow) HasPtsPerSiteMonth() bool`

HasPtsPerSiteMonth returns a boolean if a field has been set.

### SetPtsPerSiteMonthNil

`func (o *TrialEnrollmentRateRow) SetPtsPerSiteMonthNil(b bool)`

 SetPtsPerSiteMonthNil sets the value for PtsPerSiteMonth to be an explicit nil

### UnsetPtsPerSiteMonth
`func (o *TrialEnrollmentRateRow) UnsetPtsPerSiteMonth()`

UnsetPtsPerSiteMonth ensures that no value is present for PtsPerSiteMonth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


