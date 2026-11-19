# ProviderRwdDeliveryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderRwdDeliveryId** | Pointer to **NullableString** | The row id: a hash of provider_uuid, conductor_name_key and mandate_stratum. | [optional] 
**ProviderId** | Pointer to **NullableString** | The supplier id: the Scientist.com provider id as text, the id in the supplier page address (/suppliers/&lt;id&gt;). The entry key of the suppliers object type. A supplier is not a company. | [optional] 
**ProviderUuid** | Pointer to **NullableString** | The Scientist.com public provider uuid, the join key to the RX marketplace. Present on every row, including providers the marketplace does not publish. | [optional] 
**ProviderSlug** | Pointer to **NullableString** | The supplier&#39;s marketplace slug. | [optional] 
**ProviderName** | Pointer to **NullableString** | The supplier&#39;s name. | [optional] 
**InstitutionId** | Pointer to **NullableString** | The register Institution ID of the winning route to the conductor (strongest method, then lowest id). | [optional] 
**MatchMethod** | Pointer to **NullableString** | match_method says how the provider was tied to the register institution: org_member_exact (through the supplier&#39;s resolved organization, match_confidence carried from that match), rx_name_key_exact (an exact match of the register institution&#39;s name key to the supplier&#39;s, match_confidence null, a name match and not a verified identity) or org_unit_of_entry (an organization unit of the supplier&#39;s organization, one-directional, match_confidence the bridge&#39;s own figure). Where a provider reaches one conductor by several routes the strongest wins, in that order, and match_methods lists every route. | [optional] 
**MatchConfidence** | Pointer to **NullableFloat32** | The bridge&#39;s own confidence for the winning route, 0 to 1; null on an rx_name_key_exact row, whose exact name match defines none. Passed through, never computed here. | [optional] 
**ConductorNameKey** | Pointer to **NullableString** | The normalised institution name key the conductor rollup is keyed on; every register institution whose name folds to it is one conductor. Not a vendor family. | [optional] 
**MatchedInstitutionIds** | Pointer to **[]string** | Every register Institution ID that routes the provider to the conductor name key (more than one when the register lists the name twice or two methods reach it). | [optional] 
**MatchMethods** | Pointer to **[]string** | Every match method that routes the provider to this conductor name key; match_method is the strongest of them. | [optional] 
**ConductorName** | Pointer to **NullableString** | The conductor&#39;s display name: one register spelling of the key. | [optional] 
**MandateStratum** | Pointer to **NullableString** | all (every study of the conductor), mandated (studies a regulator required) or voluntary (studies the registrant says it did not require); studies with an unknown mandate are in all only. The mandated and voluntary books slip differently: choose a stratum. | [optional] 
**LifetimeStudies** | Pointer to **NullableInt32** | Distinct registered studies the conductor is named as conducting institution on. | [optional] 
**FinalReportPairs** | Pointer to **NullableInt32** | Conducting studies with both a planned and an actual final report date. | [optional] 
**ComparablePairs** | Pointer to **NullableInt32** | Conducting studies with a final report slip (both dates present and different); the N of the figures and the denominator of the floor. | [optional] 
**BackfilledPairs** | Pointer to **NullableInt32** | Of the final_report_pairs, the studies whose planned and actual date are equal to the day (a likely backfill, a form artefact). Excluded from every figure, counted here. | [optional] 
**BackfilledShare** | Pointer to **NullableFloat32** | backfilled_pairs over final_report_pairs; null when there is no pair. | [optional] 
**MandatedShare** | Pointer to **NullableFloat32** | The share of the conductor&#39;s conducting studies a regulator required; read it beside any figure to see how mandated the book is. Null when the conductor has no conducting study. | [optional] 
**NFloor** | Pointer to **NullableInt32** | The floor on comparable_pairs below which the four distribution figures are null (suppressed, not zero): 10. | [optional] 
**MeetsNFloor** | Pointer to **NullableBool** | True when comparable_pairs reaches n_floor, so the four figures are shown; false means they are null and the studies are in provider_rwd_delivery_study. | [optional] 
**MedianSlipDays** | Pointer to **NullableFloat64** | Median final report slip in days against the planned date (positive is late) over the comparable pairs; null below the floor. | [optional] 
**P25SlipDays** | Pointer to **NullableFloat64** | First quartile of the final report slip in days; null below the floor. | [optional] 
**P75SlipDays** | Pointer to **NullableFloat64** | Third quartile of the final report slip in days; null below the floor. | [optional] 
**LateOver30d** | Pointer to **NullableInt32** | Comparable studies whose final report slipped more than 30 days; null below the floor. | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewProviderRwdDeliveryRow

`func NewProviderRwdDeliveryRow() *ProviderRwdDeliveryRow`

NewProviderRwdDeliveryRow instantiates a new ProviderRwdDeliveryRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderRwdDeliveryRowWithDefaults

`func NewProviderRwdDeliveryRowWithDefaults() *ProviderRwdDeliveryRow`

NewProviderRwdDeliveryRowWithDefaults instantiates a new ProviderRwdDeliveryRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderRwdDeliveryId

`func (o *ProviderRwdDeliveryRow) GetProviderRwdDeliveryId() string`

GetProviderRwdDeliveryId returns the ProviderRwdDeliveryId field if non-nil, zero value otherwise.

### GetProviderRwdDeliveryIdOk

`func (o *ProviderRwdDeliveryRow) GetProviderRwdDeliveryIdOk() (*string, bool)`

GetProviderRwdDeliveryIdOk returns a tuple with the ProviderRwdDeliveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderRwdDeliveryId

`func (o *ProviderRwdDeliveryRow) SetProviderRwdDeliveryId(v string)`

SetProviderRwdDeliveryId sets ProviderRwdDeliveryId field to given value.

### HasProviderRwdDeliveryId

`func (o *ProviderRwdDeliveryRow) HasProviderRwdDeliveryId() bool`

HasProviderRwdDeliveryId returns a boolean if a field has been set.

### SetProviderRwdDeliveryIdNil

`func (o *ProviderRwdDeliveryRow) SetProviderRwdDeliveryIdNil(b bool)`

 SetProviderRwdDeliveryIdNil sets the value for ProviderRwdDeliveryId to be an explicit nil

### UnsetProviderRwdDeliveryId
`func (o *ProviderRwdDeliveryRow) UnsetProviderRwdDeliveryId()`

UnsetProviderRwdDeliveryId ensures that no value is present for ProviderRwdDeliveryId, not even an explicit nil
### GetProviderId

`func (o *ProviderRwdDeliveryRow) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ProviderRwdDeliveryRow) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ProviderRwdDeliveryRow) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ProviderRwdDeliveryRow) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *ProviderRwdDeliveryRow) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *ProviderRwdDeliveryRow) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetProviderUuid

`func (o *ProviderRwdDeliveryRow) GetProviderUuid() string`

GetProviderUuid returns the ProviderUuid field if non-nil, zero value otherwise.

### GetProviderUuidOk

`func (o *ProviderRwdDeliveryRow) GetProviderUuidOk() (*string, bool)`

GetProviderUuidOk returns a tuple with the ProviderUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUuid

`func (o *ProviderRwdDeliveryRow) SetProviderUuid(v string)`

SetProviderUuid sets ProviderUuid field to given value.

### HasProviderUuid

`func (o *ProviderRwdDeliveryRow) HasProviderUuid() bool`

HasProviderUuid returns a boolean if a field has been set.

### SetProviderUuidNil

`func (o *ProviderRwdDeliveryRow) SetProviderUuidNil(b bool)`

 SetProviderUuidNil sets the value for ProviderUuid to be an explicit nil

### UnsetProviderUuid
`func (o *ProviderRwdDeliveryRow) UnsetProviderUuid()`

UnsetProviderUuid ensures that no value is present for ProviderUuid, not even an explicit nil
### GetProviderSlug

`func (o *ProviderRwdDeliveryRow) GetProviderSlug() string`

GetProviderSlug returns the ProviderSlug field if non-nil, zero value otherwise.

### GetProviderSlugOk

`func (o *ProviderRwdDeliveryRow) GetProviderSlugOk() (*string, bool)`

GetProviderSlugOk returns a tuple with the ProviderSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSlug

`func (o *ProviderRwdDeliveryRow) SetProviderSlug(v string)`

SetProviderSlug sets ProviderSlug field to given value.

### HasProviderSlug

`func (o *ProviderRwdDeliveryRow) HasProviderSlug() bool`

HasProviderSlug returns a boolean if a field has been set.

### SetProviderSlugNil

`func (o *ProviderRwdDeliveryRow) SetProviderSlugNil(b bool)`

 SetProviderSlugNil sets the value for ProviderSlug to be an explicit nil

### UnsetProviderSlug
`func (o *ProviderRwdDeliveryRow) UnsetProviderSlug()`

UnsetProviderSlug ensures that no value is present for ProviderSlug, not even an explicit nil
### GetProviderName

`func (o *ProviderRwdDeliveryRow) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ProviderRwdDeliveryRow) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ProviderRwdDeliveryRow) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ProviderRwdDeliveryRow) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *ProviderRwdDeliveryRow) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *ProviderRwdDeliveryRow) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetInstitutionId

`func (o *ProviderRwdDeliveryRow) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *ProviderRwdDeliveryRow) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *ProviderRwdDeliveryRow) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.

### HasInstitutionId

`func (o *ProviderRwdDeliveryRow) HasInstitutionId() bool`

HasInstitutionId returns a boolean if a field has been set.

### SetInstitutionIdNil

`func (o *ProviderRwdDeliveryRow) SetInstitutionIdNil(b bool)`

 SetInstitutionIdNil sets the value for InstitutionId to be an explicit nil

### UnsetInstitutionId
`func (o *ProviderRwdDeliveryRow) UnsetInstitutionId()`

UnsetInstitutionId ensures that no value is present for InstitutionId, not even an explicit nil
### GetMatchMethod

`func (o *ProviderRwdDeliveryRow) GetMatchMethod() string`

GetMatchMethod returns the MatchMethod field if non-nil, zero value otherwise.

### GetMatchMethodOk

`func (o *ProviderRwdDeliveryRow) GetMatchMethodOk() (*string, bool)`

GetMatchMethodOk returns a tuple with the MatchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMethod

`func (o *ProviderRwdDeliveryRow) SetMatchMethod(v string)`

SetMatchMethod sets MatchMethod field to given value.

### HasMatchMethod

`func (o *ProviderRwdDeliveryRow) HasMatchMethod() bool`

HasMatchMethod returns a boolean if a field has been set.

### SetMatchMethodNil

`func (o *ProviderRwdDeliveryRow) SetMatchMethodNil(b bool)`

 SetMatchMethodNil sets the value for MatchMethod to be an explicit nil

### UnsetMatchMethod
`func (o *ProviderRwdDeliveryRow) UnsetMatchMethod()`

UnsetMatchMethod ensures that no value is present for MatchMethod, not even an explicit nil
### GetMatchConfidence

`func (o *ProviderRwdDeliveryRow) GetMatchConfidence() float32`

GetMatchConfidence returns the MatchConfidence field if non-nil, zero value otherwise.

### GetMatchConfidenceOk

`func (o *ProviderRwdDeliveryRow) GetMatchConfidenceOk() (*float32, bool)`

GetMatchConfidenceOk returns a tuple with the MatchConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchConfidence

`func (o *ProviderRwdDeliveryRow) SetMatchConfidence(v float32)`

SetMatchConfidence sets MatchConfidence field to given value.

### HasMatchConfidence

`func (o *ProviderRwdDeliveryRow) HasMatchConfidence() bool`

HasMatchConfidence returns a boolean if a field has been set.

### SetMatchConfidenceNil

`func (o *ProviderRwdDeliveryRow) SetMatchConfidenceNil(b bool)`

 SetMatchConfidenceNil sets the value for MatchConfidence to be an explicit nil

### UnsetMatchConfidence
`func (o *ProviderRwdDeliveryRow) UnsetMatchConfidence()`

UnsetMatchConfidence ensures that no value is present for MatchConfidence, not even an explicit nil
### GetConductorNameKey

`func (o *ProviderRwdDeliveryRow) GetConductorNameKey() string`

GetConductorNameKey returns the ConductorNameKey field if non-nil, zero value otherwise.

### GetConductorNameKeyOk

`func (o *ProviderRwdDeliveryRow) GetConductorNameKeyOk() (*string, bool)`

GetConductorNameKeyOk returns a tuple with the ConductorNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductorNameKey

`func (o *ProviderRwdDeliveryRow) SetConductorNameKey(v string)`

SetConductorNameKey sets ConductorNameKey field to given value.

### HasConductorNameKey

`func (o *ProviderRwdDeliveryRow) HasConductorNameKey() bool`

HasConductorNameKey returns a boolean if a field has been set.

### SetConductorNameKeyNil

`func (o *ProviderRwdDeliveryRow) SetConductorNameKeyNil(b bool)`

 SetConductorNameKeyNil sets the value for ConductorNameKey to be an explicit nil

### UnsetConductorNameKey
`func (o *ProviderRwdDeliveryRow) UnsetConductorNameKey()`

UnsetConductorNameKey ensures that no value is present for ConductorNameKey, not even an explicit nil
### GetMatchedInstitutionIds

`func (o *ProviderRwdDeliveryRow) GetMatchedInstitutionIds() []string`

GetMatchedInstitutionIds returns the MatchedInstitutionIds field if non-nil, zero value otherwise.

### GetMatchedInstitutionIdsOk

`func (o *ProviderRwdDeliveryRow) GetMatchedInstitutionIdsOk() (*[]string, bool)`

GetMatchedInstitutionIdsOk returns a tuple with the MatchedInstitutionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedInstitutionIds

`func (o *ProviderRwdDeliveryRow) SetMatchedInstitutionIds(v []string)`

SetMatchedInstitutionIds sets MatchedInstitutionIds field to given value.

### HasMatchedInstitutionIds

`func (o *ProviderRwdDeliveryRow) HasMatchedInstitutionIds() bool`

HasMatchedInstitutionIds returns a boolean if a field has been set.

### SetMatchedInstitutionIdsNil

`func (o *ProviderRwdDeliveryRow) SetMatchedInstitutionIdsNil(b bool)`

 SetMatchedInstitutionIdsNil sets the value for MatchedInstitutionIds to be an explicit nil

### UnsetMatchedInstitutionIds
`func (o *ProviderRwdDeliveryRow) UnsetMatchedInstitutionIds()`

UnsetMatchedInstitutionIds ensures that no value is present for MatchedInstitutionIds, not even an explicit nil
### GetMatchMethods

`func (o *ProviderRwdDeliveryRow) GetMatchMethods() []string`

GetMatchMethods returns the MatchMethods field if non-nil, zero value otherwise.

### GetMatchMethodsOk

`func (o *ProviderRwdDeliveryRow) GetMatchMethodsOk() (*[]string, bool)`

GetMatchMethodsOk returns a tuple with the MatchMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMethods

`func (o *ProviderRwdDeliveryRow) SetMatchMethods(v []string)`

SetMatchMethods sets MatchMethods field to given value.

### HasMatchMethods

`func (o *ProviderRwdDeliveryRow) HasMatchMethods() bool`

HasMatchMethods returns a boolean if a field has been set.

### SetMatchMethodsNil

`func (o *ProviderRwdDeliveryRow) SetMatchMethodsNil(b bool)`

 SetMatchMethodsNil sets the value for MatchMethods to be an explicit nil

### UnsetMatchMethods
`func (o *ProviderRwdDeliveryRow) UnsetMatchMethods()`

UnsetMatchMethods ensures that no value is present for MatchMethods, not even an explicit nil
### GetConductorName

`func (o *ProviderRwdDeliveryRow) GetConductorName() string`

GetConductorName returns the ConductorName field if non-nil, zero value otherwise.

### GetConductorNameOk

`func (o *ProviderRwdDeliveryRow) GetConductorNameOk() (*string, bool)`

GetConductorNameOk returns a tuple with the ConductorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductorName

`func (o *ProviderRwdDeliveryRow) SetConductorName(v string)`

SetConductorName sets ConductorName field to given value.

### HasConductorName

`func (o *ProviderRwdDeliveryRow) HasConductorName() bool`

HasConductorName returns a boolean if a field has been set.

### SetConductorNameNil

`func (o *ProviderRwdDeliveryRow) SetConductorNameNil(b bool)`

 SetConductorNameNil sets the value for ConductorName to be an explicit nil

### UnsetConductorName
`func (o *ProviderRwdDeliveryRow) UnsetConductorName()`

UnsetConductorName ensures that no value is present for ConductorName, not even an explicit nil
### GetMandateStratum

`func (o *ProviderRwdDeliveryRow) GetMandateStratum() string`

GetMandateStratum returns the MandateStratum field if non-nil, zero value otherwise.

### GetMandateStratumOk

`func (o *ProviderRwdDeliveryRow) GetMandateStratumOk() (*string, bool)`

GetMandateStratumOk returns a tuple with the MandateStratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateStratum

`func (o *ProviderRwdDeliveryRow) SetMandateStratum(v string)`

SetMandateStratum sets MandateStratum field to given value.

### HasMandateStratum

`func (o *ProviderRwdDeliveryRow) HasMandateStratum() bool`

HasMandateStratum returns a boolean if a field has been set.

### SetMandateStratumNil

`func (o *ProviderRwdDeliveryRow) SetMandateStratumNil(b bool)`

 SetMandateStratumNil sets the value for MandateStratum to be an explicit nil

### UnsetMandateStratum
`func (o *ProviderRwdDeliveryRow) UnsetMandateStratum()`

UnsetMandateStratum ensures that no value is present for MandateStratum, not even an explicit nil
### GetLifetimeStudies

`func (o *ProviderRwdDeliveryRow) GetLifetimeStudies() int32`

GetLifetimeStudies returns the LifetimeStudies field if non-nil, zero value otherwise.

### GetLifetimeStudiesOk

`func (o *ProviderRwdDeliveryRow) GetLifetimeStudiesOk() (*int32, bool)`

GetLifetimeStudiesOk returns a tuple with the LifetimeStudies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeStudies

`func (o *ProviderRwdDeliveryRow) SetLifetimeStudies(v int32)`

SetLifetimeStudies sets LifetimeStudies field to given value.

### HasLifetimeStudies

`func (o *ProviderRwdDeliveryRow) HasLifetimeStudies() bool`

HasLifetimeStudies returns a boolean if a field has been set.

### SetLifetimeStudiesNil

`func (o *ProviderRwdDeliveryRow) SetLifetimeStudiesNil(b bool)`

 SetLifetimeStudiesNil sets the value for LifetimeStudies to be an explicit nil

### UnsetLifetimeStudies
`func (o *ProviderRwdDeliveryRow) UnsetLifetimeStudies()`

UnsetLifetimeStudies ensures that no value is present for LifetimeStudies, not even an explicit nil
### GetFinalReportPairs

`func (o *ProviderRwdDeliveryRow) GetFinalReportPairs() int32`

GetFinalReportPairs returns the FinalReportPairs field if non-nil, zero value otherwise.

### GetFinalReportPairsOk

`func (o *ProviderRwdDeliveryRow) GetFinalReportPairsOk() (*int32, bool)`

GetFinalReportPairsOk returns a tuple with the FinalReportPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalReportPairs

`func (o *ProviderRwdDeliveryRow) SetFinalReportPairs(v int32)`

SetFinalReportPairs sets FinalReportPairs field to given value.

### HasFinalReportPairs

`func (o *ProviderRwdDeliveryRow) HasFinalReportPairs() bool`

HasFinalReportPairs returns a boolean if a field has been set.

### SetFinalReportPairsNil

`func (o *ProviderRwdDeliveryRow) SetFinalReportPairsNil(b bool)`

 SetFinalReportPairsNil sets the value for FinalReportPairs to be an explicit nil

### UnsetFinalReportPairs
`func (o *ProviderRwdDeliveryRow) UnsetFinalReportPairs()`

UnsetFinalReportPairs ensures that no value is present for FinalReportPairs, not even an explicit nil
### GetComparablePairs

`func (o *ProviderRwdDeliveryRow) GetComparablePairs() int32`

GetComparablePairs returns the ComparablePairs field if non-nil, zero value otherwise.

### GetComparablePairsOk

`func (o *ProviderRwdDeliveryRow) GetComparablePairsOk() (*int32, bool)`

GetComparablePairsOk returns a tuple with the ComparablePairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparablePairs

`func (o *ProviderRwdDeliveryRow) SetComparablePairs(v int32)`

SetComparablePairs sets ComparablePairs field to given value.

### HasComparablePairs

`func (o *ProviderRwdDeliveryRow) HasComparablePairs() bool`

HasComparablePairs returns a boolean if a field has been set.

### SetComparablePairsNil

`func (o *ProviderRwdDeliveryRow) SetComparablePairsNil(b bool)`

 SetComparablePairsNil sets the value for ComparablePairs to be an explicit nil

### UnsetComparablePairs
`func (o *ProviderRwdDeliveryRow) UnsetComparablePairs()`

UnsetComparablePairs ensures that no value is present for ComparablePairs, not even an explicit nil
### GetBackfilledPairs

`func (o *ProviderRwdDeliveryRow) GetBackfilledPairs() int32`

GetBackfilledPairs returns the BackfilledPairs field if non-nil, zero value otherwise.

### GetBackfilledPairsOk

`func (o *ProviderRwdDeliveryRow) GetBackfilledPairsOk() (*int32, bool)`

GetBackfilledPairsOk returns a tuple with the BackfilledPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfilledPairs

`func (o *ProviderRwdDeliveryRow) SetBackfilledPairs(v int32)`

SetBackfilledPairs sets BackfilledPairs field to given value.

### HasBackfilledPairs

`func (o *ProviderRwdDeliveryRow) HasBackfilledPairs() bool`

HasBackfilledPairs returns a boolean if a field has been set.

### SetBackfilledPairsNil

`func (o *ProviderRwdDeliveryRow) SetBackfilledPairsNil(b bool)`

 SetBackfilledPairsNil sets the value for BackfilledPairs to be an explicit nil

### UnsetBackfilledPairs
`func (o *ProviderRwdDeliveryRow) UnsetBackfilledPairs()`

UnsetBackfilledPairs ensures that no value is present for BackfilledPairs, not even an explicit nil
### GetBackfilledShare

`func (o *ProviderRwdDeliveryRow) GetBackfilledShare() float32`

GetBackfilledShare returns the BackfilledShare field if non-nil, zero value otherwise.

### GetBackfilledShareOk

`func (o *ProviderRwdDeliveryRow) GetBackfilledShareOk() (*float32, bool)`

GetBackfilledShareOk returns a tuple with the BackfilledShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfilledShare

`func (o *ProviderRwdDeliveryRow) SetBackfilledShare(v float32)`

SetBackfilledShare sets BackfilledShare field to given value.

### HasBackfilledShare

`func (o *ProviderRwdDeliveryRow) HasBackfilledShare() bool`

HasBackfilledShare returns a boolean if a field has been set.

### SetBackfilledShareNil

`func (o *ProviderRwdDeliveryRow) SetBackfilledShareNil(b bool)`

 SetBackfilledShareNil sets the value for BackfilledShare to be an explicit nil

### UnsetBackfilledShare
`func (o *ProviderRwdDeliveryRow) UnsetBackfilledShare()`

UnsetBackfilledShare ensures that no value is present for BackfilledShare, not even an explicit nil
### GetMandatedShare

`func (o *ProviderRwdDeliveryRow) GetMandatedShare() float32`

GetMandatedShare returns the MandatedShare field if non-nil, zero value otherwise.

### GetMandatedShareOk

`func (o *ProviderRwdDeliveryRow) GetMandatedShareOk() (*float32, bool)`

GetMandatedShareOk returns a tuple with the MandatedShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatedShare

`func (o *ProviderRwdDeliveryRow) SetMandatedShare(v float32)`

SetMandatedShare sets MandatedShare field to given value.

### HasMandatedShare

`func (o *ProviderRwdDeliveryRow) HasMandatedShare() bool`

HasMandatedShare returns a boolean if a field has been set.

### SetMandatedShareNil

`func (o *ProviderRwdDeliveryRow) SetMandatedShareNil(b bool)`

 SetMandatedShareNil sets the value for MandatedShare to be an explicit nil

### UnsetMandatedShare
`func (o *ProviderRwdDeliveryRow) UnsetMandatedShare()`

UnsetMandatedShare ensures that no value is present for MandatedShare, not even an explicit nil
### GetNFloor

`func (o *ProviderRwdDeliveryRow) GetNFloor() int32`

GetNFloor returns the NFloor field if non-nil, zero value otherwise.

### GetNFloorOk

`func (o *ProviderRwdDeliveryRow) GetNFloorOk() (*int32, bool)`

GetNFloorOk returns a tuple with the NFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFloor

`func (o *ProviderRwdDeliveryRow) SetNFloor(v int32)`

SetNFloor sets NFloor field to given value.

### HasNFloor

`func (o *ProviderRwdDeliveryRow) HasNFloor() bool`

HasNFloor returns a boolean if a field has been set.

### SetNFloorNil

`func (o *ProviderRwdDeliveryRow) SetNFloorNil(b bool)`

 SetNFloorNil sets the value for NFloor to be an explicit nil

### UnsetNFloor
`func (o *ProviderRwdDeliveryRow) UnsetNFloor()`

UnsetNFloor ensures that no value is present for NFloor, not even an explicit nil
### GetMeetsNFloor

`func (o *ProviderRwdDeliveryRow) GetMeetsNFloor() bool`

GetMeetsNFloor returns the MeetsNFloor field if non-nil, zero value otherwise.

### GetMeetsNFloorOk

`func (o *ProviderRwdDeliveryRow) GetMeetsNFloorOk() (*bool, bool)`

GetMeetsNFloorOk returns a tuple with the MeetsNFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetsNFloor

`func (o *ProviderRwdDeliveryRow) SetMeetsNFloor(v bool)`

SetMeetsNFloor sets MeetsNFloor field to given value.

### HasMeetsNFloor

`func (o *ProviderRwdDeliveryRow) HasMeetsNFloor() bool`

HasMeetsNFloor returns a boolean if a field has been set.

### SetMeetsNFloorNil

`func (o *ProviderRwdDeliveryRow) SetMeetsNFloorNil(b bool)`

 SetMeetsNFloorNil sets the value for MeetsNFloor to be an explicit nil

### UnsetMeetsNFloor
`func (o *ProviderRwdDeliveryRow) UnsetMeetsNFloor()`

UnsetMeetsNFloor ensures that no value is present for MeetsNFloor, not even an explicit nil
### GetMedianSlipDays

`func (o *ProviderRwdDeliveryRow) GetMedianSlipDays() float64`

GetMedianSlipDays returns the MedianSlipDays field if non-nil, zero value otherwise.

### GetMedianSlipDaysOk

`func (o *ProviderRwdDeliveryRow) GetMedianSlipDaysOk() (*float64, bool)`

GetMedianSlipDaysOk returns a tuple with the MedianSlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianSlipDays

`func (o *ProviderRwdDeliveryRow) SetMedianSlipDays(v float64)`

SetMedianSlipDays sets MedianSlipDays field to given value.

### HasMedianSlipDays

`func (o *ProviderRwdDeliveryRow) HasMedianSlipDays() bool`

HasMedianSlipDays returns a boolean if a field has been set.

### SetMedianSlipDaysNil

`func (o *ProviderRwdDeliveryRow) SetMedianSlipDaysNil(b bool)`

 SetMedianSlipDaysNil sets the value for MedianSlipDays to be an explicit nil

### UnsetMedianSlipDays
`func (o *ProviderRwdDeliveryRow) UnsetMedianSlipDays()`

UnsetMedianSlipDays ensures that no value is present for MedianSlipDays, not even an explicit nil
### GetP25SlipDays

`func (o *ProviderRwdDeliveryRow) GetP25SlipDays() float64`

GetP25SlipDays returns the P25SlipDays field if non-nil, zero value otherwise.

### GetP25SlipDaysOk

`func (o *ProviderRwdDeliveryRow) GetP25SlipDaysOk() (*float64, bool)`

GetP25SlipDaysOk returns a tuple with the P25SlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP25SlipDays

`func (o *ProviderRwdDeliveryRow) SetP25SlipDays(v float64)`

SetP25SlipDays sets P25SlipDays field to given value.

### HasP25SlipDays

`func (o *ProviderRwdDeliveryRow) HasP25SlipDays() bool`

HasP25SlipDays returns a boolean if a field has been set.

### SetP25SlipDaysNil

`func (o *ProviderRwdDeliveryRow) SetP25SlipDaysNil(b bool)`

 SetP25SlipDaysNil sets the value for P25SlipDays to be an explicit nil

### UnsetP25SlipDays
`func (o *ProviderRwdDeliveryRow) UnsetP25SlipDays()`

UnsetP25SlipDays ensures that no value is present for P25SlipDays, not even an explicit nil
### GetP75SlipDays

`func (o *ProviderRwdDeliveryRow) GetP75SlipDays() float64`

GetP75SlipDays returns the P75SlipDays field if non-nil, zero value otherwise.

### GetP75SlipDaysOk

`func (o *ProviderRwdDeliveryRow) GetP75SlipDaysOk() (*float64, bool)`

GetP75SlipDaysOk returns a tuple with the P75SlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP75SlipDays

`func (o *ProviderRwdDeliveryRow) SetP75SlipDays(v float64)`

SetP75SlipDays sets P75SlipDays field to given value.

### HasP75SlipDays

`func (o *ProviderRwdDeliveryRow) HasP75SlipDays() bool`

HasP75SlipDays returns a boolean if a field has been set.

### SetP75SlipDaysNil

`func (o *ProviderRwdDeliveryRow) SetP75SlipDaysNil(b bool)`

 SetP75SlipDaysNil sets the value for P75SlipDays to be an explicit nil

### UnsetP75SlipDays
`func (o *ProviderRwdDeliveryRow) UnsetP75SlipDays()`

UnsetP75SlipDays ensures that no value is present for P75SlipDays, not even an explicit nil
### GetLateOver30d

`func (o *ProviderRwdDeliveryRow) GetLateOver30d() int32`

GetLateOver30d returns the LateOver30d field if non-nil, zero value otherwise.

### GetLateOver30dOk

`func (o *ProviderRwdDeliveryRow) GetLateOver30dOk() (*int32, bool)`

GetLateOver30dOk returns a tuple with the LateOver30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateOver30d

`func (o *ProviderRwdDeliveryRow) SetLateOver30d(v int32)`

SetLateOver30d sets LateOver30d field to given value.

### HasLateOver30d

`func (o *ProviderRwdDeliveryRow) HasLateOver30d() bool`

HasLateOver30d returns a boolean if a field has been set.

### SetLateOver30dNil

`func (o *ProviderRwdDeliveryRow) SetLateOver30dNil(b bool)`

 SetLateOver30dNil sets the value for LateOver30d to be an explicit nil

### UnsetLateOver30d
`func (o *ProviderRwdDeliveryRow) UnsetLateOver30d()`

UnsetLateOver30d ensures that no value is present for LateOver30d, not even an explicit nil
### GetSnapshotLabel

`func (o *ProviderRwdDeliveryRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *ProviderRwdDeliveryRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *ProviderRwdDeliveryRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *ProviderRwdDeliveryRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *ProviderRwdDeliveryRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *ProviderRwdDeliveryRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetAccessedMonth

`func (o *ProviderRwdDeliveryRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *ProviderRwdDeliveryRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *ProviderRwdDeliveryRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *ProviderRwdDeliveryRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *ProviderRwdDeliveryRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *ProviderRwdDeliveryRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *ProviderRwdDeliveryRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *ProviderRwdDeliveryRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *ProviderRwdDeliveryRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *ProviderRwdDeliveryRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *ProviderRwdDeliveryRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *ProviderRwdDeliveryRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *ProviderRwdDeliveryRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *ProviderRwdDeliveryRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *ProviderRwdDeliveryRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *ProviderRwdDeliveryRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *ProviderRwdDeliveryRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *ProviderRwdDeliveryRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


