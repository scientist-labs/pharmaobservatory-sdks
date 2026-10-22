# RwdConductorStudiesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConductorId** | Pointer to **NullableString** | The conductor slug: the register&#39;s normalised institution name key with spaces as hyphens, the id in the web page address (/rwd-conductors/&lt;id&gt;). The entry key of the rwd_conductors object type; injective because no name key carries a hyphen. | [optional] 
**StudyId** | Pointer to **NullableString** | The register&#39;s own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional] 
**ConductorRole** | Pointer to **NullableString** | conducting when the register names the conductor as an institution conducting the study, additional when as an additional institution. | [optional] 
**StudyTitle** | Pointer to **NullableString** | The study&#39;s title as registered. | [optional] 
**StudyStatus** | Pointer to **NullableString** | The study&#39;s status as the register words it (Planned, Ongoing, Finalised, and so on). | [optional] 
**MandateStratum** | Pointer to **NullableString** | mandated when a regulator required the study (a post-authorisation obligation), voluntary when the registrant says it did not, unknown when unanswered. Delivery timeliness is only comparable within a stratum. | [optional] 
**StudyStartActualOn** | Pointer to **NullableString** | Actual study start date. | [optional] 
**FinalReportSlipDays** | Pointer to **NullableInt32** | Actual minus planned final report date in days (positive is late); null when either date is missing or the two dates are equal (a backfill). | [optional] 
**FinalReportIsBackfilled** | Pointer to **NullableBool** | True when the planned and actual final report dates are both present and equal to the day, which reads as a plan entered after the fact rather than perfect delivery. | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdConductorStudiesRow

`func NewRwdConductorStudiesRow() *RwdConductorStudiesRow`

NewRwdConductorStudiesRow instantiates a new RwdConductorStudiesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdConductorStudiesRowWithDefaults

`func NewRwdConductorStudiesRowWithDefaults() *RwdConductorStudiesRow`

NewRwdConductorStudiesRowWithDefaults instantiates a new RwdConductorStudiesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConductorId

`func (o *RwdConductorStudiesRow) GetConductorId() string`

GetConductorId returns the ConductorId field if non-nil, zero value otherwise.

### GetConductorIdOk

`func (o *RwdConductorStudiesRow) GetConductorIdOk() (*string, bool)`

GetConductorIdOk returns a tuple with the ConductorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductorId

`func (o *RwdConductorStudiesRow) SetConductorId(v string)`

SetConductorId sets ConductorId field to given value.

### HasConductorId

`func (o *RwdConductorStudiesRow) HasConductorId() bool`

HasConductorId returns a boolean if a field has been set.

### SetConductorIdNil

`func (o *RwdConductorStudiesRow) SetConductorIdNil(b bool)`

 SetConductorIdNil sets the value for ConductorId to be an explicit nil

### UnsetConductorId
`func (o *RwdConductorStudiesRow) UnsetConductorId()`

UnsetConductorId ensures that no value is present for ConductorId, not even an explicit nil
### GetStudyId

`func (o *RwdConductorStudiesRow) GetStudyId() string`

GetStudyId returns the StudyId field if non-nil, zero value otherwise.

### GetStudyIdOk

`func (o *RwdConductorStudiesRow) GetStudyIdOk() (*string, bool)`

GetStudyIdOk returns a tuple with the StudyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyId

`func (o *RwdConductorStudiesRow) SetStudyId(v string)`

SetStudyId sets StudyId field to given value.

### HasStudyId

`func (o *RwdConductorStudiesRow) HasStudyId() bool`

HasStudyId returns a boolean if a field has been set.

### SetStudyIdNil

`func (o *RwdConductorStudiesRow) SetStudyIdNil(b bool)`

 SetStudyIdNil sets the value for StudyId to be an explicit nil

### UnsetStudyId
`func (o *RwdConductorStudiesRow) UnsetStudyId()`

UnsetStudyId ensures that no value is present for StudyId, not even an explicit nil
### GetConductorRole

`func (o *RwdConductorStudiesRow) GetConductorRole() string`

GetConductorRole returns the ConductorRole field if non-nil, zero value otherwise.

### GetConductorRoleOk

`func (o *RwdConductorStudiesRow) GetConductorRoleOk() (*string, bool)`

GetConductorRoleOk returns a tuple with the ConductorRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductorRole

`func (o *RwdConductorStudiesRow) SetConductorRole(v string)`

SetConductorRole sets ConductorRole field to given value.

### HasConductorRole

`func (o *RwdConductorStudiesRow) HasConductorRole() bool`

HasConductorRole returns a boolean if a field has been set.

### SetConductorRoleNil

`func (o *RwdConductorStudiesRow) SetConductorRoleNil(b bool)`

 SetConductorRoleNil sets the value for ConductorRole to be an explicit nil

### UnsetConductorRole
`func (o *RwdConductorStudiesRow) UnsetConductorRole()`

UnsetConductorRole ensures that no value is present for ConductorRole, not even an explicit nil
### GetStudyTitle

`func (o *RwdConductorStudiesRow) GetStudyTitle() string`

GetStudyTitle returns the StudyTitle field if non-nil, zero value otherwise.

### GetStudyTitleOk

`func (o *RwdConductorStudiesRow) GetStudyTitleOk() (*string, bool)`

GetStudyTitleOk returns a tuple with the StudyTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyTitle

`func (o *RwdConductorStudiesRow) SetStudyTitle(v string)`

SetStudyTitle sets StudyTitle field to given value.

### HasStudyTitle

`func (o *RwdConductorStudiesRow) HasStudyTitle() bool`

HasStudyTitle returns a boolean if a field has been set.

### SetStudyTitleNil

`func (o *RwdConductorStudiesRow) SetStudyTitleNil(b bool)`

 SetStudyTitleNil sets the value for StudyTitle to be an explicit nil

### UnsetStudyTitle
`func (o *RwdConductorStudiesRow) UnsetStudyTitle()`

UnsetStudyTitle ensures that no value is present for StudyTitle, not even an explicit nil
### GetStudyStatus

`func (o *RwdConductorStudiesRow) GetStudyStatus() string`

GetStudyStatus returns the StudyStatus field if non-nil, zero value otherwise.

### GetStudyStatusOk

`func (o *RwdConductorStudiesRow) GetStudyStatusOk() (*string, bool)`

GetStudyStatusOk returns a tuple with the StudyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyStatus

`func (o *RwdConductorStudiesRow) SetStudyStatus(v string)`

SetStudyStatus sets StudyStatus field to given value.

### HasStudyStatus

`func (o *RwdConductorStudiesRow) HasStudyStatus() bool`

HasStudyStatus returns a boolean if a field has been set.

### SetStudyStatusNil

`func (o *RwdConductorStudiesRow) SetStudyStatusNil(b bool)`

 SetStudyStatusNil sets the value for StudyStatus to be an explicit nil

### UnsetStudyStatus
`func (o *RwdConductorStudiesRow) UnsetStudyStatus()`

UnsetStudyStatus ensures that no value is present for StudyStatus, not even an explicit nil
### GetMandateStratum

`func (o *RwdConductorStudiesRow) GetMandateStratum() string`

GetMandateStratum returns the MandateStratum field if non-nil, zero value otherwise.

### GetMandateStratumOk

`func (o *RwdConductorStudiesRow) GetMandateStratumOk() (*string, bool)`

GetMandateStratumOk returns a tuple with the MandateStratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateStratum

`func (o *RwdConductorStudiesRow) SetMandateStratum(v string)`

SetMandateStratum sets MandateStratum field to given value.

### HasMandateStratum

`func (o *RwdConductorStudiesRow) HasMandateStratum() bool`

HasMandateStratum returns a boolean if a field has been set.

### SetMandateStratumNil

`func (o *RwdConductorStudiesRow) SetMandateStratumNil(b bool)`

 SetMandateStratumNil sets the value for MandateStratum to be an explicit nil

### UnsetMandateStratum
`func (o *RwdConductorStudiesRow) UnsetMandateStratum()`

UnsetMandateStratum ensures that no value is present for MandateStratum, not even an explicit nil
### GetStudyStartActualOn

`func (o *RwdConductorStudiesRow) GetStudyStartActualOn() string`

GetStudyStartActualOn returns the StudyStartActualOn field if non-nil, zero value otherwise.

### GetStudyStartActualOnOk

`func (o *RwdConductorStudiesRow) GetStudyStartActualOnOk() (*string, bool)`

GetStudyStartActualOnOk returns a tuple with the StudyStartActualOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyStartActualOn

`func (o *RwdConductorStudiesRow) SetStudyStartActualOn(v string)`

SetStudyStartActualOn sets StudyStartActualOn field to given value.

### HasStudyStartActualOn

`func (o *RwdConductorStudiesRow) HasStudyStartActualOn() bool`

HasStudyStartActualOn returns a boolean if a field has been set.

### SetStudyStartActualOnNil

`func (o *RwdConductorStudiesRow) SetStudyStartActualOnNil(b bool)`

 SetStudyStartActualOnNil sets the value for StudyStartActualOn to be an explicit nil

### UnsetStudyStartActualOn
`func (o *RwdConductorStudiesRow) UnsetStudyStartActualOn()`

UnsetStudyStartActualOn ensures that no value is present for StudyStartActualOn, not even an explicit nil
### GetFinalReportSlipDays

`func (o *RwdConductorStudiesRow) GetFinalReportSlipDays() int32`

GetFinalReportSlipDays returns the FinalReportSlipDays field if non-nil, zero value otherwise.

### GetFinalReportSlipDaysOk

`func (o *RwdConductorStudiesRow) GetFinalReportSlipDaysOk() (*int32, bool)`

GetFinalReportSlipDaysOk returns a tuple with the FinalReportSlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalReportSlipDays

`func (o *RwdConductorStudiesRow) SetFinalReportSlipDays(v int32)`

SetFinalReportSlipDays sets FinalReportSlipDays field to given value.

### HasFinalReportSlipDays

`func (o *RwdConductorStudiesRow) HasFinalReportSlipDays() bool`

HasFinalReportSlipDays returns a boolean if a field has been set.

### SetFinalReportSlipDaysNil

`func (o *RwdConductorStudiesRow) SetFinalReportSlipDaysNil(b bool)`

 SetFinalReportSlipDaysNil sets the value for FinalReportSlipDays to be an explicit nil

### UnsetFinalReportSlipDays
`func (o *RwdConductorStudiesRow) UnsetFinalReportSlipDays()`

UnsetFinalReportSlipDays ensures that no value is present for FinalReportSlipDays, not even an explicit nil
### GetFinalReportIsBackfilled

`func (o *RwdConductorStudiesRow) GetFinalReportIsBackfilled() bool`

GetFinalReportIsBackfilled returns the FinalReportIsBackfilled field if non-nil, zero value otherwise.

### GetFinalReportIsBackfilledOk

`func (o *RwdConductorStudiesRow) GetFinalReportIsBackfilledOk() (*bool, bool)`

GetFinalReportIsBackfilledOk returns a tuple with the FinalReportIsBackfilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalReportIsBackfilled

`func (o *RwdConductorStudiesRow) SetFinalReportIsBackfilled(v bool)`

SetFinalReportIsBackfilled sets FinalReportIsBackfilled field to given value.

### HasFinalReportIsBackfilled

`func (o *RwdConductorStudiesRow) HasFinalReportIsBackfilled() bool`

HasFinalReportIsBackfilled returns a boolean if a field has been set.

### SetFinalReportIsBackfilledNil

`func (o *RwdConductorStudiesRow) SetFinalReportIsBackfilledNil(b bool)`

 SetFinalReportIsBackfilledNil sets the value for FinalReportIsBackfilled to be an explicit nil

### UnsetFinalReportIsBackfilled
`func (o *RwdConductorStudiesRow) UnsetFinalReportIsBackfilled()`

UnsetFinalReportIsBackfilled ensures that no value is present for FinalReportIsBackfilled, not even an explicit nil
### GetSnapshotLabel

`func (o *RwdConductorStudiesRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *RwdConductorStudiesRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *RwdConductorStudiesRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *RwdConductorStudiesRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *RwdConductorStudiesRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *RwdConductorStudiesRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetAccessedMonth

`func (o *RwdConductorStudiesRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *RwdConductorStudiesRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *RwdConductorStudiesRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *RwdConductorStudiesRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *RwdConductorStudiesRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *RwdConductorStudiesRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *RwdConductorStudiesRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdConductorStudiesRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdConductorStudiesRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdConductorStudiesRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdConductorStudiesRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdConductorStudiesRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdConductorStudiesRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdConductorStudiesRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdConductorStudiesRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdConductorStudiesRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdConductorStudiesRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdConductorStudiesRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


