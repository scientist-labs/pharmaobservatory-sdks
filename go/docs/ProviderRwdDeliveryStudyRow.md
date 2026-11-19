# ProviderRwdDeliveryStudyRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderRwdDeliveryStudyId** | Pointer to **NullableString** | The row id: a hash of provider_uuid, conductor_name_key and study_id. | [optional] 
**ProviderId** | Pointer to **NullableString** | The supplier id: the Scientist.com provider id as text, the id in the supplier page address (/suppliers/&lt;id&gt;). The entry key of the suppliers object type. A supplier is not a company. | [optional] 
**ProviderUuid** | Pointer to **NullableString** | The Scientist.com public provider uuid, the join key to the RX marketplace. Present on every row, including providers the marketplace does not publish. | [optional] 
**ProviderSlug** | Pointer to **NullableString** | The supplier&#39;s marketplace slug. | [optional] 
**ProviderName** | Pointer to **NullableString** | The supplier&#39;s name. | [optional] 
**InstitutionId** | Pointer to **NullableString** | The register Institution ID of the winning route to the conductor (strongest method, then lowest id). | [optional] 
**MatchMethod** | Pointer to **NullableString** | match_method says how the provider was tied to the register institution: org_member_exact (through the supplier&#39;s resolved organization, match_confidence carried from that match), rx_name_key_exact (an exact match of the register institution&#39;s name key to the supplier&#39;s, match_confidence null, a name match and not a verified identity) or org_unit_of_entry (an organization unit of the supplier&#39;s organization, one-directional, match_confidence the bridge&#39;s own figure). Where a provider reaches one conductor by several routes the strongest wins, in that order, and match_methods lists every route. | [optional] 
**MatchConfidence** | Pointer to **NullableFloat32** | The bridge&#39;s own confidence for the winning route, 0 to 1; null on an rx_name_key_exact row, whose exact name match defines none. Passed through, never computed here. | [optional] 
**ConductorNameKey** | Pointer to **NullableString** | The normalised institution name key the conductor rollup is keyed on; every register institution whose name folds to it is one conductor. Not a vendor family. | [optional] 
**StudyId** | Pointer to **NullableString** | The register&#39;s own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional] 
**EuPasNumber** | Pointer to **NullableString** | The EU PAS (post-authorisation study) register number of the study. | [optional] 
**StudyMandateStratum** | Pointer to **NullableString** | The study&#39;s own mandate: mandated (a regulator required it), voluntary, or unknown (the register does not say). | [optional] 
**FinalReportPlannedOn** | Pointer to **NullableString** | Planned date of the final study report. | [optional] 
**FinalReportActualOn** | Pointer to **NullableString** | Actual date of the final study report. | [optional] 
**FinalReportSlipDays** | Pointer to **NullableInt32** | Actual minus planned final report date in days (positive is late, negative is early). A plain fact about one study, never a statistic. | [optional] 
**AllMeetsNFloor** | Pointer to **NullableBool** | True when the provider and conductor&#39;s all row meets the floor of 10 comparable pairs. | [optional] 
**StratumMeetsNFloor** | Pointer to **NullableBool** | True when the row of the study&#39;s own mandate stratum meets the floor; null for a study whose mandate is unknown (it belongs to the all row only). | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewProviderRwdDeliveryStudyRow

`func NewProviderRwdDeliveryStudyRow() *ProviderRwdDeliveryStudyRow`

NewProviderRwdDeliveryStudyRow instantiates a new ProviderRwdDeliveryStudyRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderRwdDeliveryStudyRowWithDefaults

`func NewProviderRwdDeliveryStudyRowWithDefaults() *ProviderRwdDeliveryStudyRow`

NewProviderRwdDeliveryStudyRowWithDefaults instantiates a new ProviderRwdDeliveryStudyRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderRwdDeliveryStudyId

`func (o *ProviderRwdDeliveryStudyRow) GetProviderRwdDeliveryStudyId() string`

GetProviderRwdDeliveryStudyId returns the ProviderRwdDeliveryStudyId field if non-nil, zero value otherwise.

### GetProviderRwdDeliveryStudyIdOk

`func (o *ProviderRwdDeliveryStudyRow) GetProviderRwdDeliveryStudyIdOk() (*string, bool)`

GetProviderRwdDeliveryStudyIdOk returns a tuple with the ProviderRwdDeliveryStudyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderRwdDeliveryStudyId

`func (o *ProviderRwdDeliveryStudyRow) SetProviderRwdDeliveryStudyId(v string)`

SetProviderRwdDeliveryStudyId sets ProviderRwdDeliveryStudyId field to given value.

### HasProviderRwdDeliveryStudyId

`func (o *ProviderRwdDeliveryStudyRow) HasProviderRwdDeliveryStudyId() bool`

HasProviderRwdDeliveryStudyId returns a boolean if a field has been set.

### SetProviderRwdDeliveryStudyIdNil

`func (o *ProviderRwdDeliveryStudyRow) SetProviderRwdDeliveryStudyIdNil(b bool)`

 SetProviderRwdDeliveryStudyIdNil sets the value for ProviderRwdDeliveryStudyId to be an explicit nil

### UnsetProviderRwdDeliveryStudyId
`func (o *ProviderRwdDeliveryStudyRow) UnsetProviderRwdDeliveryStudyId()`

UnsetProviderRwdDeliveryStudyId ensures that no value is present for ProviderRwdDeliveryStudyId, not even an explicit nil
### GetProviderId

`func (o *ProviderRwdDeliveryStudyRow) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ProviderRwdDeliveryStudyRow) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ProviderRwdDeliveryStudyRow) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ProviderRwdDeliveryStudyRow) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *ProviderRwdDeliveryStudyRow) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *ProviderRwdDeliveryStudyRow) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetProviderUuid

`func (o *ProviderRwdDeliveryStudyRow) GetProviderUuid() string`

GetProviderUuid returns the ProviderUuid field if non-nil, zero value otherwise.

### GetProviderUuidOk

`func (o *ProviderRwdDeliveryStudyRow) GetProviderUuidOk() (*string, bool)`

GetProviderUuidOk returns a tuple with the ProviderUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUuid

`func (o *ProviderRwdDeliveryStudyRow) SetProviderUuid(v string)`

SetProviderUuid sets ProviderUuid field to given value.

### HasProviderUuid

`func (o *ProviderRwdDeliveryStudyRow) HasProviderUuid() bool`

HasProviderUuid returns a boolean if a field has been set.

### SetProviderUuidNil

`func (o *ProviderRwdDeliveryStudyRow) SetProviderUuidNil(b bool)`

 SetProviderUuidNil sets the value for ProviderUuid to be an explicit nil

### UnsetProviderUuid
`func (o *ProviderRwdDeliveryStudyRow) UnsetProviderUuid()`

UnsetProviderUuid ensures that no value is present for ProviderUuid, not even an explicit nil
### GetProviderSlug

`func (o *ProviderRwdDeliveryStudyRow) GetProviderSlug() string`

GetProviderSlug returns the ProviderSlug field if non-nil, zero value otherwise.

### GetProviderSlugOk

`func (o *ProviderRwdDeliveryStudyRow) GetProviderSlugOk() (*string, bool)`

GetProviderSlugOk returns a tuple with the ProviderSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSlug

`func (o *ProviderRwdDeliveryStudyRow) SetProviderSlug(v string)`

SetProviderSlug sets ProviderSlug field to given value.

### HasProviderSlug

`func (o *ProviderRwdDeliveryStudyRow) HasProviderSlug() bool`

HasProviderSlug returns a boolean if a field has been set.

### SetProviderSlugNil

`func (o *ProviderRwdDeliveryStudyRow) SetProviderSlugNil(b bool)`

 SetProviderSlugNil sets the value for ProviderSlug to be an explicit nil

### UnsetProviderSlug
`func (o *ProviderRwdDeliveryStudyRow) UnsetProviderSlug()`

UnsetProviderSlug ensures that no value is present for ProviderSlug, not even an explicit nil
### GetProviderName

`func (o *ProviderRwdDeliveryStudyRow) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ProviderRwdDeliveryStudyRow) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ProviderRwdDeliveryStudyRow) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ProviderRwdDeliveryStudyRow) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *ProviderRwdDeliveryStudyRow) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *ProviderRwdDeliveryStudyRow) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetInstitutionId

`func (o *ProviderRwdDeliveryStudyRow) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *ProviderRwdDeliveryStudyRow) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *ProviderRwdDeliveryStudyRow) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.

### HasInstitutionId

`func (o *ProviderRwdDeliveryStudyRow) HasInstitutionId() bool`

HasInstitutionId returns a boolean if a field has been set.

### SetInstitutionIdNil

`func (o *ProviderRwdDeliveryStudyRow) SetInstitutionIdNil(b bool)`

 SetInstitutionIdNil sets the value for InstitutionId to be an explicit nil

### UnsetInstitutionId
`func (o *ProviderRwdDeliveryStudyRow) UnsetInstitutionId()`

UnsetInstitutionId ensures that no value is present for InstitutionId, not even an explicit nil
### GetMatchMethod

`func (o *ProviderRwdDeliveryStudyRow) GetMatchMethod() string`

GetMatchMethod returns the MatchMethod field if non-nil, zero value otherwise.

### GetMatchMethodOk

`func (o *ProviderRwdDeliveryStudyRow) GetMatchMethodOk() (*string, bool)`

GetMatchMethodOk returns a tuple with the MatchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMethod

`func (o *ProviderRwdDeliveryStudyRow) SetMatchMethod(v string)`

SetMatchMethod sets MatchMethod field to given value.

### HasMatchMethod

`func (o *ProviderRwdDeliveryStudyRow) HasMatchMethod() bool`

HasMatchMethod returns a boolean if a field has been set.

### SetMatchMethodNil

`func (o *ProviderRwdDeliveryStudyRow) SetMatchMethodNil(b bool)`

 SetMatchMethodNil sets the value for MatchMethod to be an explicit nil

### UnsetMatchMethod
`func (o *ProviderRwdDeliveryStudyRow) UnsetMatchMethod()`

UnsetMatchMethod ensures that no value is present for MatchMethod, not even an explicit nil
### GetMatchConfidence

`func (o *ProviderRwdDeliveryStudyRow) GetMatchConfidence() float32`

GetMatchConfidence returns the MatchConfidence field if non-nil, zero value otherwise.

### GetMatchConfidenceOk

`func (o *ProviderRwdDeliveryStudyRow) GetMatchConfidenceOk() (*float32, bool)`

GetMatchConfidenceOk returns a tuple with the MatchConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchConfidence

`func (o *ProviderRwdDeliveryStudyRow) SetMatchConfidence(v float32)`

SetMatchConfidence sets MatchConfidence field to given value.

### HasMatchConfidence

`func (o *ProviderRwdDeliveryStudyRow) HasMatchConfidence() bool`

HasMatchConfidence returns a boolean if a field has been set.

### SetMatchConfidenceNil

`func (o *ProviderRwdDeliveryStudyRow) SetMatchConfidenceNil(b bool)`

 SetMatchConfidenceNil sets the value for MatchConfidence to be an explicit nil

### UnsetMatchConfidence
`func (o *ProviderRwdDeliveryStudyRow) UnsetMatchConfidence()`

UnsetMatchConfidence ensures that no value is present for MatchConfidence, not even an explicit nil
### GetConductorNameKey

`func (o *ProviderRwdDeliveryStudyRow) GetConductorNameKey() string`

GetConductorNameKey returns the ConductorNameKey field if non-nil, zero value otherwise.

### GetConductorNameKeyOk

`func (o *ProviderRwdDeliveryStudyRow) GetConductorNameKeyOk() (*string, bool)`

GetConductorNameKeyOk returns a tuple with the ConductorNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductorNameKey

`func (o *ProviderRwdDeliveryStudyRow) SetConductorNameKey(v string)`

SetConductorNameKey sets ConductorNameKey field to given value.

### HasConductorNameKey

`func (o *ProviderRwdDeliveryStudyRow) HasConductorNameKey() bool`

HasConductorNameKey returns a boolean if a field has been set.

### SetConductorNameKeyNil

`func (o *ProviderRwdDeliveryStudyRow) SetConductorNameKeyNil(b bool)`

 SetConductorNameKeyNil sets the value for ConductorNameKey to be an explicit nil

### UnsetConductorNameKey
`func (o *ProviderRwdDeliveryStudyRow) UnsetConductorNameKey()`

UnsetConductorNameKey ensures that no value is present for ConductorNameKey, not even an explicit nil
### GetStudyId

`func (o *ProviderRwdDeliveryStudyRow) GetStudyId() string`

GetStudyId returns the StudyId field if non-nil, zero value otherwise.

### GetStudyIdOk

`func (o *ProviderRwdDeliveryStudyRow) GetStudyIdOk() (*string, bool)`

GetStudyIdOk returns a tuple with the StudyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyId

`func (o *ProviderRwdDeliveryStudyRow) SetStudyId(v string)`

SetStudyId sets StudyId field to given value.

### HasStudyId

`func (o *ProviderRwdDeliveryStudyRow) HasStudyId() bool`

HasStudyId returns a boolean if a field has been set.

### SetStudyIdNil

`func (o *ProviderRwdDeliveryStudyRow) SetStudyIdNil(b bool)`

 SetStudyIdNil sets the value for StudyId to be an explicit nil

### UnsetStudyId
`func (o *ProviderRwdDeliveryStudyRow) UnsetStudyId()`

UnsetStudyId ensures that no value is present for StudyId, not even an explicit nil
### GetEuPasNumber

`func (o *ProviderRwdDeliveryStudyRow) GetEuPasNumber() string`

GetEuPasNumber returns the EuPasNumber field if non-nil, zero value otherwise.

### GetEuPasNumberOk

`func (o *ProviderRwdDeliveryStudyRow) GetEuPasNumberOk() (*string, bool)`

GetEuPasNumberOk returns a tuple with the EuPasNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEuPasNumber

`func (o *ProviderRwdDeliveryStudyRow) SetEuPasNumber(v string)`

SetEuPasNumber sets EuPasNumber field to given value.

### HasEuPasNumber

`func (o *ProviderRwdDeliveryStudyRow) HasEuPasNumber() bool`

HasEuPasNumber returns a boolean if a field has been set.

### SetEuPasNumberNil

`func (o *ProviderRwdDeliveryStudyRow) SetEuPasNumberNil(b bool)`

 SetEuPasNumberNil sets the value for EuPasNumber to be an explicit nil

### UnsetEuPasNumber
`func (o *ProviderRwdDeliveryStudyRow) UnsetEuPasNumber()`

UnsetEuPasNumber ensures that no value is present for EuPasNumber, not even an explicit nil
### GetStudyMandateStratum

`func (o *ProviderRwdDeliveryStudyRow) GetStudyMandateStratum() string`

GetStudyMandateStratum returns the StudyMandateStratum field if non-nil, zero value otherwise.

### GetStudyMandateStratumOk

`func (o *ProviderRwdDeliveryStudyRow) GetStudyMandateStratumOk() (*string, bool)`

GetStudyMandateStratumOk returns a tuple with the StudyMandateStratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyMandateStratum

`func (o *ProviderRwdDeliveryStudyRow) SetStudyMandateStratum(v string)`

SetStudyMandateStratum sets StudyMandateStratum field to given value.

### HasStudyMandateStratum

`func (o *ProviderRwdDeliveryStudyRow) HasStudyMandateStratum() bool`

HasStudyMandateStratum returns a boolean if a field has been set.

### SetStudyMandateStratumNil

`func (o *ProviderRwdDeliveryStudyRow) SetStudyMandateStratumNil(b bool)`

 SetStudyMandateStratumNil sets the value for StudyMandateStratum to be an explicit nil

### UnsetStudyMandateStratum
`func (o *ProviderRwdDeliveryStudyRow) UnsetStudyMandateStratum()`

UnsetStudyMandateStratum ensures that no value is present for StudyMandateStratum, not even an explicit nil
### GetFinalReportPlannedOn

`func (o *ProviderRwdDeliveryStudyRow) GetFinalReportPlannedOn() string`

GetFinalReportPlannedOn returns the FinalReportPlannedOn field if non-nil, zero value otherwise.

### GetFinalReportPlannedOnOk

`func (o *ProviderRwdDeliveryStudyRow) GetFinalReportPlannedOnOk() (*string, bool)`

GetFinalReportPlannedOnOk returns a tuple with the FinalReportPlannedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalReportPlannedOn

`func (o *ProviderRwdDeliveryStudyRow) SetFinalReportPlannedOn(v string)`

SetFinalReportPlannedOn sets FinalReportPlannedOn field to given value.

### HasFinalReportPlannedOn

`func (o *ProviderRwdDeliveryStudyRow) HasFinalReportPlannedOn() bool`

HasFinalReportPlannedOn returns a boolean if a field has been set.

### SetFinalReportPlannedOnNil

`func (o *ProviderRwdDeliveryStudyRow) SetFinalReportPlannedOnNil(b bool)`

 SetFinalReportPlannedOnNil sets the value for FinalReportPlannedOn to be an explicit nil

### UnsetFinalReportPlannedOn
`func (o *ProviderRwdDeliveryStudyRow) UnsetFinalReportPlannedOn()`

UnsetFinalReportPlannedOn ensures that no value is present for FinalReportPlannedOn, not even an explicit nil
### GetFinalReportActualOn

`func (o *ProviderRwdDeliveryStudyRow) GetFinalReportActualOn() string`

GetFinalReportActualOn returns the FinalReportActualOn field if non-nil, zero value otherwise.

### GetFinalReportActualOnOk

`func (o *ProviderRwdDeliveryStudyRow) GetFinalReportActualOnOk() (*string, bool)`

GetFinalReportActualOnOk returns a tuple with the FinalReportActualOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalReportActualOn

`func (o *ProviderRwdDeliveryStudyRow) SetFinalReportActualOn(v string)`

SetFinalReportActualOn sets FinalReportActualOn field to given value.

### HasFinalReportActualOn

`func (o *ProviderRwdDeliveryStudyRow) HasFinalReportActualOn() bool`

HasFinalReportActualOn returns a boolean if a field has been set.

### SetFinalReportActualOnNil

`func (o *ProviderRwdDeliveryStudyRow) SetFinalReportActualOnNil(b bool)`

 SetFinalReportActualOnNil sets the value for FinalReportActualOn to be an explicit nil

### UnsetFinalReportActualOn
`func (o *ProviderRwdDeliveryStudyRow) UnsetFinalReportActualOn()`

UnsetFinalReportActualOn ensures that no value is present for FinalReportActualOn, not even an explicit nil
### GetFinalReportSlipDays

`func (o *ProviderRwdDeliveryStudyRow) GetFinalReportSlipDays() int32`

GetFinalReportSlipDays returns the FinalReportSlipDays field if non-nil, zero value otherwise.

### GetFinalReportSlipDaysOk

`func (o *ProviderRwdDeliveryStudyRow) GetFinalReportSlipDaysOk() (*int32, bool)`

GetFinalReportSlipDaysOk returns a tuple with the FinalReportSlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalReportSlipDays

`func (o *ProviderRwdDeliveryStudyRow) SetFinalReportSlipDays(v int32)`

SetFinalReportSlipDays sets FinalReportSlipDays field to given value.

### HasFinalReportSlipDays

`func (o *ProviderRwdDeliveryStudyRow) HasFinalReportSlipDays() bool`

HasFinalReportSlipDays returns a boolean if a field has been set.

### SetFinalReportSlipDaysNil

`func (o *ProviderRwdDeliveryStudyRow) SetFinalReportSlipDaysNil(b bool)`

 SetFinalReportSlipDaysNil sets the value for FinalReportSlipDays to be an explicit nil

### UnsetFinalReportSlipDays
`func (o *ProviderRwdDeliveryStudyRow) UnsetFinalReportSlipDays()`

UnsetFinalReportSlipDays ensures that no value is present for FinalReportSlipDays, not even an explicit nil
### GetAllMeetsNFloor

`func (o *ProviderRwdDeliveryStudyRow) GetAllMeetsNFloor() bool`

GetAllMeetsNFloor returns the AllMeetsNFloor field if non-nil, zero value otherwise.

### GetAllMeetsNFloorOk

`func (o *ProviderRwdDeliveryStudyRow) GetAllMeetsNFloorOk() (*bool, bool)`

GetAllMeetsNFloorOk returns a tuple with the AllMeetsNFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMeetsNFloor

`func (o *ProviderRwdDeliveryStudyRow) SetAllMeetsNFloor(v bool)`

SetAllMeetsNFloor sets AllMeetsNFloor field to given value.

### HasAllMeetsNFloor

`func (o *ProviderRwdDeliveryStudyRow) HasAllMeetsNFloor() bool`

HasAllMeetsNFloor returns a boolean if a field has been set.

### SetAllMeetsNFloorNil

`func (o *ProviderRwdDeliveryStudyRow) SetAllMeetsNFloorNil(b bool)`

 SetAllMeetsNFloorNil sets the value for AllMeetsNFloor to be an explicit nil

### UnsetAllMeetsNFloor
`func (o *ProviderRwdDeliveryStudyRow) UnsetAllMeetsNFloor()`

UnsetAllMeetsNFloor ensures that no value is present for AllMeetsNFloor, not even an explicit nil
### GetStratumMeetsNFloor

`func (o *ProviderRwdDeliveryStudyRow) GetStratumMeetsNFloor() bool`

GetStratumMeetsNFloor returns the StratumMeetsNFloor field if non-nil, zero value otherwise.

### GetStratumMeetsNFloorOk

`func (o *ProviderRwdDeliveryStudyRow) GetStratumMeetsNFloorOk() (*bool, bool)`

GetStratumMeetsNFloorOk returns a tuple with the StratumMeetsNFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStratumMeetsNFloor

`func (o *ProviderRwdDeliveryStudyRow) SetStratumMeetsNFloor(v bool)`

SetStratumMeetsNFloor sets StratumMeetsNFloor field to given value.

### HasStratumMeetsNFloor

`func (o *ProviderRwdDeliveryStudyRow) HasStratumMeetsNFloor() bool`

HasStratumMeetsNFloor returns a boolean if a field has been set.

### SetStratumMeetsNFloorNil

`func (o *ProviderRwdDeliveryStudyRow) SetStratumMeetsNFloorNil(b bool)`

 SetStratumMeetsNFloorNil sets the value for StratumMeetsNFloor to be an explicit nil

### UnsetStratumMeetsNFloor
`func (o *ProviderRwdDeliveryStudyRow) UnsetStratumMeetsNFloor()`

UnsetStratumMeetsNFloor ensures that no value is present for StratumMeetsNFloor, not even an explicit nil
### GetSnapshotLabel

`func (o *ProviderRwdDeliveryStudyRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *ProviderRwdDeliveryStudyRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *ProviderRwdDeliveryStudyRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *ProviderRwdDeliveryStudyRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *ProviderRwdDeliveryStudyRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *ProviderRwdDeliveryStudyRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetAccessedMonth

`func (o *ProviderRwdDeliveryStudyRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *ProviderRwdDeliveryStudyRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *ProviderRwdDeliveryStudyRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *ProviderRwdDeliveryStudyRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *ProviderRwdDeliveryStudyRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *ProviderRwdDeliveryStudyRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *ProviderRwdDeliveryStudyRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *ProviderRwdDeliveryStudyRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *ProviderRwdDeliveryStudyRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *ProviderRwdDeliveryStudyRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *ProviderRwdDeliveryStudyRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *ProviderRwdDeliveryStudyRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *ProviderRwdDeliveryStudyRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *ProviderRwdDeliveryStudyRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *ProviderRwdDeliveryStudyRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *ProviderRwdDeliveryStudyRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *ProviderRwdDeliveryStudyRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *ProviderRwdDeliveryStudyRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


