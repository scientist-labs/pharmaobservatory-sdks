# ProviderRwdFootprintRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderRwdFootprintId** | Pointer to **NullableString** | The assertion id: a hash of the provider uuid, role, register entity kind and id, match method and institution. | [optional] 
**ProviderId** | Pointer to **NullableString** | The supplier id: the Scientist.com provider id as text, the entry key of the suppliers object type. A supplier is not a company. | [optional] 
**ProviderUuid** | Pointer to **NullableString** | The Scientist.com public provider uuid, the join key to the RX marketplace. Present on every row, including providers the marketplace does not publish. | [optional] 
**ProviderSlug** | Pointer to **NullableString** | The supplier&#39;s marketplace slug. | [optional] 
**ProviderName** | Pointer to **NullableString** | The supplier&#39;s name. | [optional] 
**InstitutionId** | Pointer to **NullableString** | The register institution the provider was matched THROUGH. The entry key of the rwd_institutions object type. | [optional] 
**InstitutionName** | Pointer to **NullableString** | The register&#39;s own name for that institution. | [optional] 
**Role** | Pointer to **NullableString** | What the register entity is to the institution: conductor or additional_institution (a study), custodian (a data source&#39;s holder) or network_member (a network&#39;s member institution). | [optional] 
**EvidenceRefKind** | Pointer to **NullableString** | Which register entity evidence_ref names: rwd_study_id, rwd_data_source_id or rwd_network_id. | [optional] 
**EvidenceRef** | Pointer to **NullableString** | The register id of the study, data source or network, per evidence_ref_kind. Ids of different kinds can coincide; read it with evidence_ref_kind. | [optional] 
**EvidenceLabel** | Pointer to **NullableString** | The register&#39;s own entity name (study title, data source name or network name), never prose. | [optional] 
**CatalogueUrl** | Pointer to **NullableString** | The address of the register entity on the HMA-EMA catalogue site. | [optional] 
**StudyStartActualOn** | Pointer to **NullableString** | Study rows only: the registered actual start date; null otherwise and where the study has none. | [optional] 
**StudyStatus** | Pointer to **NullableString** | Study rows only: the study&#39;s registered status. | [optional] 
**MandateStratum** | Pointer to **NullableString** | Study rows only: mandated, voluntary or unknown. | [optional] 
**IsDarwinEuStudy** | Pointer to **NullableBool** | Study rows only: whether the study is a DARWIN EU (the EMA&#39;s Data Analysis and Real World Interrogation Network) study. | [optional] 
**DataSourceIds** | Pointer to **[]string** | Study rows only: the register data source ids the study names (the names that match a registered source). | [optional] 
**DataSourceNamesUnmatched** | Pointer to **[]string** | Study rows only: the data source names the study gives that match no registered source. | [optional] 
**StudyCountryNames** | Pointer to **[]string** | Study rows only: the countries the study names. | [optional] 
**DataSourceType** | Pointer to **NullableString** | Data source rows only: the data source type(s) the holder ticked, as the register exports them. | [optional] 
**CountryNames** | Pointer to **[]string** | Data source rows only: the countries the source covers. | [optional] 
**MatchMethod** | Pointer to **NullableString** | How the provider was tied to the institution: org_member_exact (through the supplier&#39;s resolved organization), rx_name_key_exact (an exact name match, not a verified identity) or org_unit_of_entry (the register entry&#39;s own statement of its parent organization). | [optional] 
**AlsoMatchedByName** | Pointer to **NullableBool** | Whether the same provider and institution were also tied by an exact name match. | [optional] 
**MatchConfidence** | Pointer to **NullableFloat32** | The upstream bridge&#39;s confidence figure, passed through unchanged; null on rx_name_key_exact rows. | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewProviderRwdFootprintRow

`func NewProviderRwdFootprintRow() *ProviderRwdFootprintRow`

NewProviderRwdFootprintRow instantiates a new ProviderRwdFootprintRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderRwdFootprintRowWithDefaults

`func NewProviderRwdFootprintRowWithDefaults() *ProviderRwdFootprintRow`

NewProviderRwdFootprintRowWithDefaults instantiates a new ProviderRwdFootprintRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderRwdFootprintId

`func (o *ProviderRwdFootprintRow) GetProviderRwdFootprintId() string`

GetProviderRwdFootprintId returns the ProviderRwdFootprintId field if non-nil, zero value otherwise.

### GetProviderRwdFootprintIdOk

`func (o *ProviderRwdFootprintRow) GetProviderRwdFootprintIdOk() (*string, bool)`

GetProviderRwdFootprintIdOk returns a tuple with the ProviderRwdFootprintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderRwdFootprintId

`func (o *ProviderRwdFootprintRow) SetProviderRwdFootprintId(v string)`

SetProviderRwdFootprintId sets ProviderRwdFootprintId field to given value.

### HasProviderRwdFootprintId

`func (o *ProviderRwdFootprintRow) HasProviderRwdFootprintId() bool`

HasProviderRwdFootprintId returns a boolean if a field has been set.

### SetProviderRwdFootprintIdNil

`func (o *ProviderRwdFootprintRow) SetProviderRwdFootprintIdNil(b bool)`

 SetProviderRwdFootprintIdNil sets the value for ProviderRwdFootprintId to be an explicit nil

### UnsetProviderRwdFootprintId
`func (o *ProviderRwdFootprintRow) UnsetProviderRwdFootprintId()`

UnsetProviderRwdFootprintId ensures that no value is present for ProviderRwdFootprintId, not even an explicit nil
### GetProviderId

`func (o *ProviderRwdFootprintRow) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ProviderRwdFootprintRow) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ProviderRwdFootprintRow) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ProviderRwdFootprintRow) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *ProviderRwdFootprintRow) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *ProviderRwdFootprintRow) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetProviderUuid

`func (o *ProviderRwdFootprintRow) GetProviderUuid() string`

GetProviderUuid returns the ProviderUuid field if non-nil, zero value otherwise.

### GetProviderUuidOk

`func (o *ProviderRwdFootprintRow) GetProviderUuidOk() (*string, bool)`

GetProviderUuidOk returns a tuple with the ProviderUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUuid

`func (o *ProviderRwdFootprintRow) SetProviderUuid(v string)`

SetProviderUuid sets ProviderUuid field to given value.

### HasProviderUuid

`func (o *ProviderRwdFootprintRow) HasProviderUuid() bool`

HasProviderUuid returns a boolean if a field has been set.

### SetProviderUuidNil

`func (o *ProviderRwdFootprintRow) SetProviderUuidNil(b bool)`

 SetProviderUuidNil sets the value for ProviderUuid to be an explicit nil

### UnsetProviderUuid
`func (o *ProviderRwdFootprintRow) UnsetProviderUuid()`

UnsetProviderUuid ensures that no value is present for ProviderUuid, not even an explicit nil
### GetProviderSlug

`func (o *ProviderRwdFootprintRow) GetProviderSlug() string`

GetProviderSlug returns the ProviderSlug field if non-nil, zero value otherwise.

### GetProviderSlugOk

`func (o *ProviderRwdFootprintRow) GetProviderSlugOk() (*string, bool)`

GetProviderSlugOk returns a tuple with the ProviderSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSlug

`func (o *ProviderRwdFootprintRow) SetProviderSlug(v string)`

SetProviderSlug sets ProviderSlug field to given value.

### HasProviderSlug

`func (o *ProviderRwdFootprintRow) HasProviderSlug() bool`

HasProviderSlug returns a boolean if a field has been set.

### SetProviderSlugNil

`func (o *ProviderRwdFootprintRow) SetProviderSlugNil(b bool)`

 SetProviderSlugNil sets the value for ProviderSlug to be an explicit nil

### UnsetProviderSlug
`func (o *ProviderRwdFootprintRow) UnsetProviderSlug()`

UnsetProviderSlug ensures that no value is present for ProviderSlug, not even an explicit nil
### GetProviderName

`func (o *ProviderRwdFootprintRow) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ProviderRwdFootprintRow) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ProviderRwdFootprintRow) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ProviderRwdFootprintRow) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *ProviderRwdFootprintRow) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *ProviderRwdFootprintRow) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetInstitutionId

`func (o *ProviderRwdFootprintRow) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *ProviderRwdFootprintRow) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *ProviderRwdFootprintRow) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.

### HasInstitutionId

`func (o *ProviderRwdFootprintRow) HasInstitutionId() bool`

HasInstitutionId returns a boolean if a field has been set.

### SetInstitutionIdNil

`func (o *ProviderRwdFootprintRow) SetInstitutionIdNil(b bool)`

 SetInstitutionIdNil sets the value for InstitutionId to be an explicit nil

### UnsetInstitutionId
`func (o *ProviderRwdFootprintRow) UnsetInstitutionId()`

UnsetInstitutionId ensures that no value is present for InstitutionId, not even an explicit nil
### GetInstitutionName

`func (o *ProviderRwdFootprintRow) GetInstitutionName() string`

GetInstitutionName returns the InstitutionName field if non-nil, zero value otherwise.

### GetInstitutionNameOk

`func (o *ProviderRwdFootprintRow) GetInstitutionNameOk() (*string, bool)`

GetInstitutionNameOk returns a tuple with the InstitutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionName

`func (o *ProviderRwdFootprintRow) SetInstitutionName(v string)`

SetInstitutionName sets InstitutionName field to given value.

### HasInstitutionName

`func (o *ProviderRwdFootprintRow) HasInstitutionName() bool`

HasInstitutionName returns a boolean if a field has been set.

### SetInstitutionNameNil

`func (o *ProviderRwdFootprintRow) SetInstitutionNameNil(b bool)`

 SetInstitutionNameNil sets the value for InstitutionName to be an explicit nil

### UnsetInstitutionName
`func (o *ProviderRwdFootprintRow) UnsetInstitutionName()`

UnsetInstitutionName ensures that no value is present for InstitutionName, not even an explicit nil
### GetRole

`func (o *ProviderRwdFootprintRow) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProviderRwdFootprintRow) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProviderRwdFootprintRow) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ProviderRwdFootprintRow) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *ProviderRwdFootprintRow) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *ProviderRwdFootprintRow) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetEvidenceRefKind

`func (o *ProviderRwdFootprintRow) GetEvidenceRefKind() string`

GetEvidenceRefKind returns the EvidenceRefKind field if non-nil, zero value otherwise.

### GetEvidenceRefKindOk

`func (o *ProviderRwdFootprintRow) GetEvidenceRefKindOk() (*string, bool)`

GetEvidenceRefKindOk returns a tuple with the EvidenceRefKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceRefKind

`func (o *ProviderRwdFootprintRow) SetEvidenceRefKind(v string)`

SetEvidenceRefKind sets EvidenceRefKind field to given value.

### HasEvidenceRefKind

`func (o *ProviderRwdFootprintRow) HasEvidenceRefKind() bool`

HasEvidenceRefKind returns a boolean if a field has been set.

### SetEvidenceRefKindNil

`func (o *ProviderRwdFootprintRow) SetEvidenceRefKindNil(b bool)`

 SetEvidenceRefKindNil sets the value for EvidenceRefKind to be an explicit nil

### UnsetEvidenceRefKind
`func (o *ProviderRwdFootprintRow) UnsetEvidenceRefKind()`

UnsetEvidenceRefKind ensures that no value is present for EvidenceRefKind, not even an explicit nil
### GetEvidenceRef

`func (o *ProviderRwdFootprintRow) GetEvidenceRef() string`

GetEvidenceRef returns the EvidenceRef field if non-nil, zero value otherwise.

### GetEvidenceRefOk

`func (o *ProviderRwdFootprintRow) GetEvidenceRefOk() (*string, bool)`

GetEvidenceRefOk returns a tuple with the EvidenceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceRef

`func (o *ProviderRwdFootprintRow) SetEvidenceRef(v string)`

SetEvidenceRef sets EvidenceRef field to given value.

### HasEvidenceRef

`func (o *ProviderRwdFootprintRow) HasEvidenceRef() bool`

HasEvidenceRef returns a boolean if a field has been set.

### SetEvidenceRefNil

`func (o *ProviderRwdFootprintRow) SetEvidenceRefNil(b bool)`

 SetEvidenceRefNil sets the value for EvidenceRef to be an explicit nil

### UnsetEvidenceRef
`func (o *ProviderRwdFootprintRow) UnsetEvidenceRef()`

UnsetEvidenceRef ensures that no value is present for EvidenceRef, not even an explicit nil
### GetEvidenceLabel

`func (o *ProviderRwdFootprintRow) GetEvidenceLabel() string`

GetEvidenceLabel returns the EvidenceLabel field if non-nil, zero value otherwise.

### GetEvidenceLabelOk

`func (o *ProviderRwdFootprintRow) GetEvidenceLabelOk() (*string, bool)`

GetEvidenceLabelOk returns a tuple with the EvidenceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceLabel

`func (o *ProviderRwdFootprintRow) SetEvidenceLabel(v string)`

SetEvidenceLabel sets EvidenceLabel field to given value.

### HasEvidenceLabel

`func (o *ProviderRwdFootprintRow) HasEvidenceLabel() bool`

HasEvidenceLabel returns a boolean if a field has been set.

### SetEvidenceLabelNil

`func (o *ProviderRwdFootprintRow) SetEvidenceLabelNil(b bool)`

 SetEvidenceLabelNil sets the value for EvidenceLabel to be an explicit nil

### UnsetEvidenceLabel
`func (o *ProviderRwdFootprintRow) UnsetEvidenceLabel()`

UnsetEvidenceLabel ensures that no value is present for EvidenceLabel, not even an explicit nil
### GetCatalogueUrl

`func (o *ProviderRwdFootprintRow) GetCatalogueUrl() string`

GetCatalogueUrl returns the CatalogueUrl field if non-nil, zero value otherwise.

### GetCatalogueUrlOk

`func (o *ProviderRwdFootprintRow) GetCatalogueUrlOk() (*string, bool)`

GetCatalogueUrlOk returns a tuple with the CatalogueUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueUrl

`func (o *ProviderRwdFootprintRow) SetCatalogueUrl(v string)`

SetCatalogueUrl sets CatalogueUrl field to given value.

### HasCatalogueUrl

`func (o *ProviderRwdFootprintRow) HasCatalogueUrl() bool`

HasCatalogueUrl returns a boolean if a field has been set.

### SetCatalogueUrlNil

`func (o *ProviderRwdFootprintRow) SetCatalogueUrlNil(b bool)`

 SetCatalogueUrlNil sets the value for CatalogueUrl to be an explicit nil

### UnsetCatalogueUrl
`func (o *ProviderRwdFootprintRow) UnsetCatalogueUrl()`

UnsetCatalogueUrl ensures that no value is present for CatalogueUrl, not even an explicit nil
### GetStudyStartActualOn

`func (o *ProviderRwdFootprintRow) GetStudyStartActualOn() string`

GetStudyStartActualOn returns the StudyStartActualOn field if non-nil, zero value otherwise.

### GetStudyStartActualOnOk

`func (o *ProviderRwdFootprintRow) GetStudyStartActualOnOk() (*string, bool)`

GetStudyStartActualOnOk returns a tuple with the StudyStartActualOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyStartActualOn

`func (o *ProviderRwdFootprintRow) SetStudyStartActualOn(v string)`

SetStudyStartActualOn sets StudyStartActualOn field to given value.

### HasStudyStartActualOn

`func (o *ProviderRwdFootprintRow) HasStudyStartActualOn() bool`

HasStudyStartActualOn returns a boolean if a field has been set.

### SetStudyStartActualOnNil

`func (o *ProviderRwdFootprintRow) SetStudyStartActualOnNil(b bool)`

 SetStudyStartActualOnNil sets the value for StudyStartActualOn to be an explicit nil

### UnsetStudyStartActualOn
`func (o *ProviderRwdFootprintRow) UnsetStudyStartActualOn()`

UnsetStudyStartActualOn ensures that no value is present for StudyStartActualOn, not even an explicit nil
### GetStudyStatus

`func (o *ProviderRwdFootprintRow) GetStudyStatus() string`

GetStudyStatus returns the StudyStatus field if non-nil, zero value otherwise.

### GetStudyStatusOk

`func (o *ProviderRwdFootprintRow) GetStudyStatusOk() (*string, bool)`

GetStudyStatusOk returns a tuple with the StudyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyStatus

`func (o *ProviderRwdFootprintRow) SetStudyStatus(v string)`

SetStudyStatus sets StudyStatus field to given value.

### HasStudyStatus

`func (o *ProviderRwdFootprintRow) HasStudyStatus() bool`

HasStudyStatus returns a boolean if a field has been set.

### SetStudyStatusNil

`func (o *ProviderRwdFootprintRow) SetStudyStatusNil(b bool)`

 SetStudyStatusNil sets the value for StudyStatus to be an explicit nil

### UnsetStudyStatus
`func (o *ProviderRwdFootprintRow) UnsetStudyStatus()`

UnsetStudyStatus ensures that no value is present for StudyStatus, not even an explicit nil
### GetMandateStratum

`func (o *ProviderRwdFootprintRow) GetMandateStratum() string`

GetMandateStratum returns the MandateStratum field if non-nil, zero value otherwise.

### GetMandateStratumOk

`func (o *ProviderRwdFootprintRow) GetMandateStratumOk() (*string, bool)`

GetMandateStratumOk returns a tuple with the MandateStratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateStratum

`func (o *ProviderRwdFootprintRow) SetMandateStratum(v string)`

SetMandateStratum sets MandateStratum field to given value.

### HasMandateStratum

`func (o *ProviderRwdFootprintRow) HasMandateStratum() bool`

HasMandateStratum returns a boolean if a field has been set.

### SetMandateStratumNil

`func (o *ProviderRwdFootprintRow) SetMandateStratumNil(b bool)`

 SetMandateStratumNil sets the value for MandateStratum to be an explicit nil

### UnsetMandateStratum
`func (o *ProviderRwdFootprintRow) UnsetMandateStratum()`

UnsetMandateStratum ensures that no value is present for MandateStratum, not even an explicit nil
### GetIsDarwinEuStudy

`func (o *ProviderRwdFootprintRow) GetIsDarwinEuStudy() bool`

GetIsDarwinEuStudy returns the IsDarwinEuStudy field if non-nil, zero value otherwise.

### GetIsDarwinEuStudyOk

`func (o *ProviderRwdFootprintRow) GetIsDarwinEuStudyOk() (*bool, bool)`

GetIsDarwinEuStudyOk returns a tuple with the IsDarwinEuStudy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDarwinEuStudy

`func (o *ProviderRwdFootprintRow) SetIsDarwinEuStudy(v bool)`

SetIsDarwinEuStudy sets IsDarwinEuStudy field to given value.

### HasIsDarwinEuStudy

`func (o *ProviderRwdFootprintRow) HasIsDarwinEuStudy() bool`

HasIsDarwinEuStudy returns a boolean if a field has been set.

### SetIsDarwinEuStudyNil

`func (o *ProviderRwdFootprintRow) SetIsDarwinEuStudyNil(b bool)`

 SetIsDarwinEuStudyNil sets the value for IsDarwinEuStudy to be an explicit nil

### UnsetIsDarwinEuStudy
`func (o *ProviderRwdFootprintRow) UnsetIsDarwinEuStudy()`

UnsetIsDarwinEuStudy ensures that no value is present for IsDarwinEuStudy, not even an explicit nil
### GetDataSourceIds

`func (o *ProviderRwdFootprintRow) GetDataSourceIds() []string`

GetDataSourceIds returns the DataSourceIds field if non-nil, zero value otherwise.

### GetDataSourceIdsOk

`func (o *ProviderRwdFootprintRow) GetDataSourceIdsOk() (*[]string, bool)`

GetDataSourceIdsOk returns a tuple with the DataSourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceIds

`func (o *ProviderRwdFootprintRow) SetDataSourceIds(v []string)`

SetDataSourceIds sets DataSourceIds field to given value.

### HasDataSourceIds

`func (o *ProviderRwdFootprintRow) HasDataSourceIds() bool`

HasDataSourceIds returns a boolean if a field has been set.

### SetDataSourceIdsNil

`func (o *ProviderRwdFootprintRow) SetDataSourceIdsNil(b bool)`

 SetDataSourceIdsNil sets the value for DataSourceIds to be an explicit nil

### UnsetDataSourceIds
`func (o *ProviderRwdFootprintRow) UnsetDataSourceIds()`

UnsetDataSourceIds ensures that no value is present for DataSourceIds, not even an explicit nil
### GetDataSourceNamesUnmatched

`func (o *ProviderRwdFootprintRow) GetDataSourceNamesUnmatched() []string`

GetDataSourceNamesUnmatched returns the DataSourceNamesUnmatched field if non-nil, zero value otherwise.

### GetDataSourceNamesUnmatchedOk

`func (o *ProviderRwdFootprintRow) GetDataSourceNamesUnmatchedOk() (*[]string, bool)`

GetDataSourceNamesUnmatchedOk returns a tuple with the DataSourceNamesUnmatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceNamesUnmatched

`func (o *ProviderRwdFootprintRow) SetDataSourceNamesUnmatched(v []string)`

SetDataSourceNamesUnmatched sets DataSourceNamesUnmatched field to given value.

### HasDataSourceNamesUnmatched

`func (o *ProviderRwdFootprintRow) HasDataSourceNamesUnmatched() bool`

HasDataSourceNamesUnmatched returns a boolean if a field has been set.

### SetDataSourceNamesUnmatchedNil

`func (o *ProviderRwdFootprintRow) SetDataSourceNamesUnmatchedNil(b bool)`

 SetDataSourceNamesUnmatchedNil sets the value for DataSourceNamesUnmatched to be an explicit nil

### UnsetDataSourceNamesUnmatched
`func (o *ProviderRwdFootprintRow) UnsetDataSourceNamesUnmatched()`

UnsetDataSourceNamesUnmatched ensures that no value is present for DataSourceNamesUnmatched, not even an explicit nil
### GetStudyCountryNames

`func (o *ProviderRwdFootprintRow) GetStudyCountryNames() []string`

GetStudyCountryNames returns the StudyCountryNames field if non-nil, zero value otherwise.

### GetStudyCountryNamesOk

`func (o *ProviderRwdFootprintRow) GetStudyCountryNamesOk() (*[]string, bool)`

GetStudyCountryNamesOk returns a tuple with the StudyCountryNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyCountryNames

`func (o *ProviderRwdFootprintRow) SetStudyCountryNames(v []string)`

SetStudyCountryNames sets StudyCountryNames field to given value.

### HasStudyCountryNames

`func (o *ProviderRwdFootprintRow) HasStudyCountryNames() bool`

HasStudyCountryNames returns a boolean if a field has been set.

### SetStudyCountryNamesNil

`func (o *ProviderRwdFootprintRow) SetStudyCountryNamesNil(b bool)`

 SetStudyCountryNamesNil sets the value for StudyCountryNames to be an explicit nil

### UnsetStudyCountryNames
`func (o *ProviderRwdFootprintRow) UnsetStudyCountryNames()`

UnsetStudyCountryNames ensures that no value is present for StudyCountryNames, not even an explicit nil
### GetDataSourceType

`func (o *ProviderRwdFootprintRow) GetDataSourceType() string`

GetDataSourceType returns the DataSourceType field if non-nil, zero value otherwise.

### GetDataSourceTypeOk

`func (o *ProviderRwdFootprintRow) GetDataSourceTypeOk() (*string, bool)`

GetDataSourceTypeOk returns a tuple with the DataSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceType

`func (o *ProviderRwdFootprintRow) SetDataSourceType(v string)`

SetDataSourceType sets DataSourceType field to given value.

### HasDataSourceType

`func (o *ProviderRwdFootprintRow) HasDataSourceType() bool`

HasDataSourceType returns a boolean if a field has been set.

### SetDataSourceTypeNil

`func (o *ProviderRwdFootprintRow) SetDataSourceTypeNil(b bool)`

 SetDataSourceTypeNil sets the value for DataSourceType to be an explicit nil

### UnsetDataSourceType
`func (o *ProviderRwdFootprintRow) UnsetDataSourceType()`

UnsetDataSourceType ensures that no value is present for DataSourceType, not even an explicit nil
### GetCountryNames

`func (o *ProviderRwdFootprintRow) GetCountryNames() []string`

GetCountryNames returns the CountryNames field if non-nil, zero value otherwise.

### GetCountryNamesOk

`func (o *ProviderRwdFootprintRow) GetCountryNamesOk() (*[]string, bool)`

GetCountryNamesOk returns a tuple with the CountryNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryNames

`func (o *ProviderRwdFootprintRow) SetCountryNames(v []string)`

SetCountryNames sets CountryNames field to given value.

### HasCountryNames

`func (o *ProviderRwdFootprintRow) HasCountryNames() bool`

HasCountryNames returns a boolean if a field has been set.

### SetCountryNamesNil

`func (o *ProviderRwdFootprintRow) SetCountryNamesNil(b bool)`

 SetCountryNamesNil sets the value for CountryNames to be an explicit nil

### UnsetCountryNames
`func (o *ProviderRwdFootprintRow) UnsetCountryNames()`

UnsetCountryNames ensures that no value is present for CountryNames, not even an explicit nil
### GetMatchMethod

`func (o *ProviderRwdFootprintRow) GetMatchMethod() string`

GetMatchMethod returns the MatchMethod field if non-nil, zero value otherwise.

### GetMatchMethodOk

`func (o *ProviderRwdFootprintRow) GetMatchMethodOk() (*string, bool)`

GetMatchMethodOk returns a tuple with the MatchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMethod

`func (o *ProviderRwdFootprintRow) SetMatchMethod(v string)`

SetMatchMethod sets MatchMethod field to given value.

### HasMatchMethod

`func (o *ProviderRwdFootprintRow) HasMatchMethod() bool`

HasMatchMethod returns a boolean if a field has been set.

### SetMatchMethodNil

`func (o *ProviderRwdFootprintRow) SetMatchMethodNil(b bool)`

 SetMatchMethodNil sets the value for MatchMethod to be an explicit nil

### UnsetMatchMethod
`func (o *ProviderRwdFootprintRow) UnsetMatchMethod()`

UnsetMatchMethod ensures that no value is present for MatchMethod, not even an explicit nil
### GetAlsoMatchedByName

`func (o *ProviderRwdFootprintRow) GetAlsoMatchedByName() bool`

GetAlsoMatchedByName returns the AlsoMatchedByName field if non-nil, zero value otherwise.

### GetAlsoMatchedByNameOk

`func (o *ProviderRwdFootprintRow) GetAlsoMatchedByNameOk() (*bool, bool)`

GetAlsoMatchedByNameOk returns a tuple with the AlsoMatchedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlsoMatchedByName

`func (o *ProviderRwdFootprintRow) SetAlsoMatchedByName(v bool)`

SetAlsoMatchedByName sets AlsoMatchedByName field to given value.

### HasAlsoMatchedByName

`func (o *ProviderRwdFootprintRow) HasAlsoMatchedByName() bool`

HasAlsoMatchedByName returns a boolean if a field has been set.

### SetAlsoMatchedByNameNil

`func (o *ProviderRwdFootprintRow) SetAlsoMatchedByNameNil(b bool)`

 SetAlsoMatchedByNameNil sets the value for AlsoMatchedByName to be an explicit nil

### UnsetAlsoMatchedByName
`func (o *ProviderRwdFootprintRow) UnsetAlsoMatchedByName()`

UnsetAlsoMatchedByName ensures that no value is present for AlsoMatchedByName, not even an explicit nil
### GetMatchConfidence

`func (o *ProviderRwdFootprintRow) GetMatchConfidence() float32`

GetMatchConfidence returns the MatchConfidence field if non-nil, zero value otherwise.

### GetMatchConfidenceOk

`func (o *ProviderRwdFootprintRow) GetMatchConfidenceOk() (*float32, bool)`

GetMatchConfidenceOk returns a tuple with the MatchConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchConfidence

`func (o *ProviderRwdFootprintRow) SetMatchConfidence(v float32)`

SetMatchConfidence sets MatchConfidence field to given value.

### HasMatchConfidence

`func (o *ProviderRwdFootprintRow) HasMatchConfidence() bool`

HasMatchConfidence returns a boolean if a field has been set.

### SetMatchConfidenceNil

`func (o *ProviderRwdFootprintRow) SetMatchConfidenceNil(b bool)`

 SetMatchConfidenceNil sets the value for MatchConfidence to be an explicit nil

### UnsetMatchConfidence
`func (o *ProviderRwdFootprintRow) UnsetMatchConfidence()`

UnsetMatchConfidence ensures that no value is present for MatchConfidence, not even an explicit nil
### GetSnapshotLabel

`func (o *ProviderRwdFootprintRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *ProviderRwdFootprintRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *ProviderRwdFootprintRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *ProviderRwdFootprintRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *ProviderRwdFootprintRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *ProviderRwdFootprintRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetAccessedMonth

`func (o *ProviderRwdFootprintRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *ProviderRwdFootprintRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *ProviderRwdFootprintRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *ProviderRwdFootprintRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *ProviderRwdFootprintRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *ProviderRwdFootprintRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *ProviderRwdFootprintRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *ProviderRwdFootprintRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *ProviderRwdFootprintRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *ProviderRwdFootprintRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *ProviderRwdFootprintRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *ProviderRwdFootprintRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *ProviderRwdFootprintRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *ProviderRwdFootprintRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *ProviderRwdFootprintRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *ProviderRwdFootprintRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *ProviderRwdFootprintRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *ProviderRwdFootprintRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


