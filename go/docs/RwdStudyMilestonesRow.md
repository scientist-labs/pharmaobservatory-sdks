# RwdStudyMilestonesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StudyId** | Pointer to **NullableString** | The register&#39;s own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional] 
**MilestoneOrdinal** | Pointer to **NullableInt32** | The milestone&#39;s position in the study timeline. | [optional] 
**MilestoneKey** | Pointer to **NullableString** | Which milestone: funding_contract, study_start, data_analysis_start, interim_report or final_report. | [optional] 
**MilestoneLabel** | Pointer to **NullableString** | The milestone&#39;s label. | [optional] 
**PlannedOn** | Pointer to **NullableString** | The planned date. | [optional] 
**ActualOn** | Pointer to **NullableString** | The actual date. | [optional] 
**HasPair** | Pointer to **NullableBool** | True when both dates are present. | [optional] 
**IsBackfilled** | Pointer to **NullableBool** | True when planned and actual are both present and equal to the day (a plan entered after the fact). | [optional] 
**IsComparablePair** | Pointer to **NullableBool** | True when both dates are present and differ, so the slip is a measurement; every delivery distribution is computed over these rows. | [optional] 
**SlipDays** | Pointer to **NullableInt32** | Actual minus planned in days (positive is late); null when either date is missing or the two are equal. | [optional] 
**IsLateOver30d** | Pointer to **NullableBool** | True when the slip is more than 30 days; null when the pair is not comparable. | [optional] 
**RequiredByRegulator** | Pointer to **NullableString** | The registrant&#39;s answer to whether a regulatory body required the study (Yes, No or Unknown). | [optional] 
**MandateStratum** | Pointer to **NullableString** | mandated when a regulator required the study (a post-authorisation obligation), voluntary when the registrant says it did not, unknown when unanswered. Delivery timeliness is only comparable within a stratum. | [optional] 
**RmpRequirement** | Pointer to **NullableString** | The risk management plan (RMP) category the registrant says required the study, as typed. | [optional] 
**StudyStatus** | Pointer to **NullableString** | The study&#39;s status as the register words it (Planned, Ongoing, Finalised, and so on). | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdStudyMilestonesRow

`func NewRwdStudyMilestonesRow() *RwdStudyMilestonesRow`

NewRwdStudyMilestonesRow instantiates a new RwdStudyMilestonesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdStudyMilestonesRowWithDefaults

`func NewRwdStudyMilestonesRowWithDefaults() *RwdStudyMilestonesRow`

NewRwdStudyMilestonesRowWithDefaults instantiates a new RwdStudyMilestonesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStudyId

`func (o *RwdStudyMilestonesRow) GetStudyId() string`

GetStudyId returns the StudyId field if non-nil, zero value otherwise.

### GetStudyIdOk

`func (o *RwdStudyMilestonesRow) GetStudyIdOk() (*string, bool)`

GetStudyIdOk returns a tuple with the StudyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyId

`func (o *RwdStudyMilestonesRow) SetStudyId(v string)`

SetStudyId sets StudyId field to given value.

### HasStudyId

`func (o *RwdStudyMilestonesRow) HasStudyId() bool`

HasStudyId returns a boolean if a field has been set.

### SetStudyIdNil

`func (o *RwdStudyMilestonesRow) SetStudyIdNil(b bool)`

 SetStudyIdNil sets the value for StudyId to be an explicit nil

### UnsetStudyId
`func (o *RwdStudyMilestonesRow) UnsetStudyId()`

UnsetStudyId ensures that no value is present for StudyId, not even an explicit nil
### GetMilestoneOrdinal

`func (o *RwdStudyMilestonesRow) GetMilestoneOrdinal() int32`

GetMilestoneOrdinal returns the MilestoneOrdinal field if non-nil, zero value otherwise.

### GetMilestoneOrdinalOk

`func (o *RwdStudyMilestonesRow) GetMilestoneOrdinalOk() (*int32, bool)`

GetMilestoneOrdinalOk returns a tuple with the MilestoneOrdinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneOrdinal

`func (o *RwdStudyMilestonesRow) SetMilestoneOrdinal(v int32)`

SetMilestoneOrdinal sets MilestoneOrdinal field to given value.

### HasMilestoneOrdinal

`func (o *RwdStudyMilestonesRow) HasMilestoneOrdinal() bool`

HasMilestoneOrdinal returns a boolean if a field has been set.

### SetMilestoneOrdinalNil

`func (o *RwdStudyMilestonesRow) SetMilestoneOrdinalNil(b bool)`

 SetMilestoneOrdinalNil sets the value for MilestoneOrdinal to be an explicit nil

### UnsetMilestoneOrdinal
`func (o *RwdStudyMilestonesRow) UnsetMilestoneOrdinal()`

UnsetMilestoneOrdinal ensures that no value is present for MilestoneOrdinal, not even an explicit nil
### GetMilestoneKey

`func (o *RwdStudyMilestonesRow) GetMilestoneKey() string`

GetMilestoneKey returns the MilestoneKey field if non-nil, zero value otherwise.

### GetMilestoneKeyOk

`func (o *RwdStudyMilestonesRow) GetMilestoneKeyOk() (*string, bool)`

GetMilestoneKeyOk returns a tuple with the MilestoneKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneKey

`func (o *RwdStudyMilestonesRow) SetMilestoneKey(v string)`

SetMilestoneKey sets MilestoneKey field to given value.

### HasMilestoneKey

`func (o *RwdStudyMilestonesRow) HasMilestoneKey() bool`

HasMilestoneKey returns a boolean if a field has been set.

### SetMilestoneKeyNil

`func (o *RwdStudyMilestonesRow) SetMilestoneKeyNil(b bool)`

 SetMilestoneKeyNil sets the value for MilestoneKey to be an explicit nil

### UnsetMilestoneKey
`func (o *RwdStudyMilestonesRow) UnsetMilestoneKey()`

UnsetMilestoneKey ensures that no value is present for MilestoneKey, not even an explicit nil
### GetMilestoneLabel

`func (o *RwdStudyMilestonesRow) GetMilestoneLabel() string`

GetMilestoneLabel returns the MilestoneLabel field if non-nil, zero value otherwise.

### GetMilestoneLabelOk

`func (o *RwdStudyMilestonesRow) GetMilestoneLabelOk() (*string, bool)`

GetMilestoneLabelOk returns a tuple with the MilestoneLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneLabel

`func (o *RwdStudyMilestonesRow) SetMilestoneLabel(v string)`

SetMilestoneLabel sets MilestoneLabel field to given value.

### HasMilestoneLabel

`func (o *RwdStudyMilestonesRow) HasMilestoneLabel() bool`

HasMilestoneLabel returns a boolean if a field has been set.

### SetMilestoneLabelNil

`func (o *RwdStudyMilestonesRow) SetMilestoneLabelNil(b bool)`

 SetMilestoneLabelNil sets the value for MilestoneLabel to be an explicit nil

### UnsetMilestoneLabel
`func (o *RwdStudyMilestonesRow) UnsetMilestoneLabel()`

UnsetMilestoneLabel ensures that no value is present for MilestoneLabel, not even an explicit nil
### GetPlannedOn

`func (o *RwdStudyMilestonesRow) GetPlannedOn() string`

GetPlannedOn returns the PlannedOn field if non-nil, zero value otherwise.

### GetPlannedOnOk

`func (o *RwdStudyMilestonesRow) GetPlannedOnOk() (*string, bool)`

GetPlannedOnOk returns a tuple with the PlannedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedOn

`func (o *RwdStudyMilestonesRow) SetPlannedOn(v string)`

SetPlannedOn sets PlannedOn field to given value.

### HasPlannedOn

`func (o *RwdStudyMilestonesRow) HasPlannedOn() bool`

HasPlannedOn returns a boolean if a field has been set.

### SetPlannedOnNil

`func (o *RwdStudyMilestonesRow) SetPlannedOnNil(b bool)`

 SetPlannedOnNil sets the value for PlannedOn to be an explicit nil

### UnsetPlannedOn
`func (o *RwdStudyMilestonesRow) UnsetPlannedOn()`

UnsetPlannedOn ensures that no value is present for PlannedOn, not even an explicit nil
### GetActualOn

`func (o *RwdStudyMilestonesRow) GetActualOn() string`

GetActualOn returns the ActualOn field if non-nil, zero value otherwise.

### GetActualOnOk

`func (o *RwdStudyMilestonesRow) GetActualOnOk() (*string, bool)`

GetActualOnOk returns a tuple with the ActualOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualOn

`func (o *RwdStudyMilestonesRow) SetActualOn(v string)`

SetActualOn sets ActualOn field to given value.

### HasActualOn

`func (o *RwdStudyMilestonesRow) HasActualOn() bool`

HasActualOn returns a boolean if a field has been set.

### SetActualOnNil

`func (o *RwdStudyMilestonesRow) SetActualOnNil(b bool)`

 SetActualOnNil sets the value for ActualOn to be an explicit nil

### UnsetActualOn
`func (o *RwdStudyMilestonesRow) UnsetActualOn()`

UnsetActualOn ensures that no value is present for ActualOn, not even an explicit nil
### GetHasPair

`func (o *RwdStudyMilestonesRow) GetHasPair() bool`

GetHasPair returns the HasPair field if non-nil, zero value otherwise.

### GetHasPairOk

`func (o *RwdStudyMilestonesRow) GetHasPairOk() (*bool, bool)`

GetHasPairOk returns a tuple with the HasPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPair

`func (o *RwdStudyMilestonesRow) SetHasPair(v bool)`

SetHasPair sets HasPair field to given value.

### HasHasPair

`func (o *RwdStudyMilestonesRow) HasHasPair() bool`

HasHasPair returns a boolean if a field has been set.

### SetHasPairNil

`func (o *RwdStudyMilestonesRow) SetHasPairNil(b bool)`

 SetHasPairNil sets the value for HasPair to be an explicit nil

### UnsetHasPair
`func (o *RwdStudyMilestonesRow) UnsetHasPair()`

UnsetHasPair ensures that no value is present for HasPair, not even an explicit nil
### GetIsBackfilled

`func (o *RwdStudyMilestonesRow) GetIsBackfilled() bool`

GetIsBackfilled returns the IsBackfilled field if non-nil, zero value otherwise.

### GetIsBackfilledOk

`func (o *RwdStudyMilestonesRow) GetIsBackfilledOk() (*bool, bool)`

GetIsBackfilledOk returns a tuple with the IsBackfilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackfilled

`func (o *RwdStudyMilestonesRow) SetIsBackfilled(v bool)`

SetIsBackfilled sets IsBackfilled field to given value.

### HasIsBackfilled

`func (o *RwdStudyMilestonesRow) HasIsBackfilled() bool`

HasIsBackfilled returns a boolean if a field has been set.

### SetIsBackfilledNil

`func (o *RwdStudyMilestonesRow) SetIsBackfilledNil(b bool)`

 SetIsBackfilledNil sets the value for IsBackfilled to be an explicit nil

### UnsetIsBackfilled
`func (o *RwdStudyMilestonesRow) UnsetIsBackfilled()`

UnsetIsBackfilled ensures that no value is present for IsBackfilled, not even an explicit nil
### GetIsComparablePair

`func (o *RwdStudyMilestonesRow) GetIsComparablePair() bool`

GetIsComparablePair returns the IsComparablePair field if non-nil, zero value otherwise.

### GetIsComparablePairOk

`func (o *RwdStudyMilestonesRow) GetIsComparablePairOk() (*bool, bool)`

GetIsComparablePairOk returns a tuple with the IsComparablePair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComparablePair

`func (o *RwdStudyMilestonesRow) SetIsComparablePair(v bool)`

SetIsComparablePair sets IsComparablePair field to given value.

### HasIsComparablePair

`func (o *RwdStudyMilestonesRow) HasIsComparablePair() bool`

HasIsComparablePair returns a boolean if a field has been set.

### SetIsComparablePairNil

`func (o *RwdStudyMilestonesRow) SetIsComparablePairNil(b bool)`

 SetIsComparablePairNil sets the value for IsComparablePair to be an explicit nil

### UnsetIsComparablePair
`func (o *RwdStudyMilestonesRow) UnsetIsComparablePair()`

UnsetIsComparablePair ensures that no value is present for IsComparablePair, not even an explicit nil
### GetSlipDays

`func (o *RwdStudyMilestonesRow) GetSlipDays() int32`

GetSlipDays returns the SlipDays field if non-nil, zero value otherwise.

### GetSlipDaysOk

`func (o *RwdStudyMilestonesRow) GetSlipDaysOk() (*int32, bool)`

GetSlipDaysOk returns a tuple with the SlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlipDays

`func (o *RwdStudyMilestonesRow) SetSlipDays(v int32)`

SetSlipDays sets SlipDays field to given value.

### HasSlipDays

`func (o *RwdStudyMilestonesRow) HasSlipDays() bool`

HasSlipDays returns a boolean if a field has been set.

### SetSlipDaysNil

`func (o *RwdStudyMilestonesRow) SetSlipDaysNil(b bool)`

 SetSlipDaysNil sets the value for SlipDays to be an explicit nil

### UnsetSlipDays
`func (o *RwdStudyMilestonesRow) UnsetSlipDays()`

UnsetSlipDays ensures that no value is present for SlipDays, not even an explicit nil
### GetIsLateOver30d

`func (o *RwdStudyMilestonesRow) GetIsLateOver30d() bool`

GetIsLateOver30d returns the IsLateOver30d field if non-nil, zero value otherwise.

### GetIsLateOver30dOk

`func (o *RwdStudyMilestonesRow) GetIsLateOver30dOk() (*bool, bool)`

GetIsLateOver30dOk returns a tuple with the IsLateOver30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLateOver30d

`func (o *RwdStudyMilestonesRow) SetIsLateOver30d(v bool)`

SetIsLateOver30d sets IsLateOver30d field to given value.

### HasIsLateOver30d

`func (o *RwdStudyMilestonesRow) HasIsLateOver30d() bool`

HasIsLateOver30d returns a boolean if a field has been set.

### SetIsLateOver30dNil

`func (o *RwdStudyMilestonesRow) SetIsLateOver30dNil(b bool)`

 SetIsLateOver30dNil sets the value for IsLateOver30d to be an explicit nil

### UnsetIsLateOver30d
`func (o *RwdStudyMilestonesRow) UnsetIsLateOver30d()`

UnsetIsLateOver30d ensures that no value is present for IsLateOver30d, not even an explicit nil
### GetRequiredByRegulator

`func (o *RwdStudyMilestonesRow) GetRequiredByRegulator() string`

GetRequiredByRegulator returns the RequiredByRegulator field if non-nil, zero value otherwise.

### GetRequiredByRegulatorOk

`func (o *RwdStudyMilestonesRow) GetRequiredByRegulatorOk() (*string, bool)`

GetRequiredByRegulatorOk returns a tuple with the RequiredByRegulator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredByRegulator

`func (o *RwdStudyMilestonesRow) SetRequiredByRegulator(v string)`

SetRequiredByRegulator sets RequiredByRegulator field to given value.

### HasRequiredByRegulator

`func (o *RwdStudyMilestonesRow) HasRequiredByRegulator() bool`

HasRequiredByRegulator returns a boolean if a field has been set.

### SetRequiredByRegulatorNil

`func (o *RwdStudyMilestonesRow) SetRequiredByRegulatorNil(b bool)`

 SetRequiredByRegulatorNil sets the value for RequiredByRegulator to be an explicit nil

### UnsetRequiredByRegulator
`func (o *RwdStudyMilestonesRow) UnsetRequiredByRegulator()`

UnsetRequiredByRegulator ensures that no value is present for RequiredByRegulator, not even an explicit nil
### GetMandateStratum

`func (o *RwdStudyMilestonesRow) GetMandateStratum() string`

GetMandateStratum returns the MandateStratum field if non-nil, zero value otherwise.

### GetMandateStratumOk

`func (o *RwdStudyMilestonesRow) GetMandateStratumOk() (*string, bool)`

GetMandateStratumOk returns a tuple with the MandateStratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateStratum

`func (o *RwdStudyMilestonesRow) SetMandateStratum(v string)`

SetMandateStratum sets MandateStratum field to given value.

### HasMandateStratum

`func (o *RwdStudyMilestonesRow) HasMandateStratum() bool`

HasMandateStratum returns a boolean if a field has been set.

### SetMandateStratumNil

`func (o *RwdStudyMilestonesRow) SetMandateStratumNil(b bool)`

 SetMandateStratumNil sets the value for MandateStratum to be an explicit nil

### UnsetMandateStratum
`func (o *RwdStudyMilestonesRow) UnsetMandateStratum()`

UnsetMandateStratum ensures that no value is present for MandateStratum, not even an explicit nil
### GetRmpRequirement

`func (o *RwdStudyMilestonesRow) GetRmpRequirement() string`

GetRmpRequirement returns the RmpRequirement field if non-nil, zero value otherwise.

### GetRmpRequirementOk

`func (o *RwdStudyMilestonesRow) GetRmpRequirementOk() (*string, bool)`

GetRmpRequirementOk returns a tuple with the RmpRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmpRequirement

`func (o *RwdStudyMilestonesRow) SetRmpRequirement(v string)`

SetRmpRequirement sets RmpRequirement field to given value.

### HasRmpRequirement

`func (o *RwdStudyMilestonesRow) HasRmpRequirement() bool`

HasRmpRequirement returns a boolean if a field has been set.

### SetRmpRequirementNil

`func (o *RwdStudyMilestonesRow) SetRmpRequirementNil(b bool)`

 SetRmpRequirementNil sets the value for RmpRequirement to be an explicit nil

### UnsetRmpRequirement
`func (o *RwdStudyMilestonesRow) UnsetRmpRequirement()`

UnsetRmpRequirement ensures that no value is present for RmpRequirement, not even an explicit nil
### GetStudyStatus

`func (o *RwdStudyMilestonesRow) GetStudyStatus() string`

GetStudyStatus returns the StudyStatus field if non-nil, zero value otherwise.

### GetStudyStatusOk

`func (o *RwdStudyMilestonesRow) GetStudyStatusOk() (*string, bool)`

GetStudyStatusOk returns a tuple with the StudyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyStatus

`func (o *RwdStudyMilestonesRow) SetStudyStatus(v string)`

SetStudyStatus sets StudyStatus field to given value.

### HasStudyStatus

`func (o *RwdStudyMilestonesRow) HasStudyStatus() bool`

HasStudyStatus returns a boolean if a field has been set.

### SetStudyStatusNil

`func (o *RwdStudyMilestonesRow) SetStudyStatusNil(b bool)`

 SetStudyStatusNil sets the value for StudyStatus to be an explicit nil

### UnsetStudyStatus
`func (o *RwdStudyMilestonesRow) UnsetStudyStatus()`

UnsetStudyStatus ensures that no value is present for StudyStatus, not even an explicit nil
### GetSnapshotLabel

`func (o *RwdStudyMilestonesRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *RwdStudyMilestonesRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *RwdStudyMilestonesRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *RwdStudyMilestonesRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *RwdStudyMilestonesRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *RwdStudyMilestonesRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetAccessedMonth

`func (o *RwdStudyMilestonesRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *RwdStudyMilestonesRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *RwdStudyMilestonesRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *RwdStudyMilestonesRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *RwdStudyMilestonesRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *RwdStudyMilestonesRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *RwdStudyMilestonesRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdStudyMilestonesRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdStudyMilestonesRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdStudyMilestonesRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdStudyMilestonesRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdStudyMilestonesRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdStudyMilestonesRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdStudyMilestonesRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdStudyMilestonesRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdStudyMilestonesRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdStudyMilestonesRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdStudyMilestonesRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


