# SupplierRwdDeclarationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderRwdDeclaredSourceId** | Pointer to **NullableString** | The fact id: a hash of the supplier uuid, the fact kind and the fact reference. | [optional] 
**ProviderId** | Pointer to **NullableString** | The supplier id: the Scientist.com provider id as text, the entry key of the suppliers object type. A supplier is not a company. | [optional] 
**ProviderUuid** | Pointer to **NullableString** | The Scientist.com public provider uuid, the join key to the RX marketplace. | [optional] 
**ProviderSlug** | Pointer to **NullableString** | The supplier&#39;s marketplace slug. | [optional] 
**ProviderName** | Pointer to **NullableString** | The supplier&#39;s name. | [optional] 
**Role** | Pointer to **NullableString** | Always supplier_declared: the supplier said it, on the Scientist.com questionnaire. It is never a register role. | [optional] 
**EvidenceRefKind** | Pointer to **NullableString** | Which kind of fact the row is: rx_rfi_owned_database (a database the supplier lists), rx_rfi_data_access (countries and regions it declares data access in, naming no source) or rx_rfi_therapeutic_experience (rare-disease and therapeutic-area experience). | [optional] 
**EvidenceRef** | Pointer to **NullableString** | The fact reference, per evidence_ref_kind: the questionnaire id with the item position for a database, the questionnaire id otherwise. Read it with evidence_ref_kind. | [optional] 
**EvidenceLabel** | Pointer to **NullableString** | The database short name the supplier gave (capped at 120 characters), or a fixed label for the two other kinds. | [optional] 
**MatchMethod** | Pointer to **NullableString** | How the fact was tied to the supplier: rx_provider_direct, the questionnaire carries the supplier id and nothing was inferred. | [optional] 
**CountryNames** | Pointer to **[]string** | Database rows: the database country. Access rows: the countries the supplier declares access in. Null on therapeutic rows. | [optional] 
**DeclaredRegions** | Pointer to **[]string** | Access rows: the regions the supplier declares access in, as the questionnaire writes them. Null elsewhere. | [optional] 
**DataTypes** | Pointer to **[]string** | The data types, mapped to the current label of each option (a label with no mapping is kept as ticked). Database rows: types the source holds. Access rows: types the supplier declares experience with. | [optional] 
**DataSourceType** | Pointer to **NullableString** | data_types joined with &#39; | &#39;. | [optional] 
**DataTypesBasis** | Pointer to **NullableString** | What data_types means: source_holds (database rows) or supplier_experience (access rows, which name no source). Null where there are no data types. | [optional] 
**DeclaredValidation** | Pointer to **NullableString** | Database rows: the supplier&#39;s own statement of validation, declared_validated, declared_not_validated or declared_validation_in_progress. | [optional] 
**DeclaredFirstCollectionYear** | Pointer to **NullableInt32** | Database rows: the first-collection year the supplier gave, kept where it is between 1950 and the build year, else null. | [optional] 
**YearImplausible** | Pointer to **NullableBool** | Database rows: true where the supplier gave a year outside 1950 to the build year (declared_first_collection_year is then null). | [optional] 
**DeclaredLinkable** | Pointer to **NullableBool** | Database rows: whether the supplier says the database can be linked; null where it did not answer. | [optional] 
**TherapeuticAreas** | Pointer to **[]string** | Therapeutic rows: the therapeutic areas ticked, current labels. An empty list is a declared negative, not missing data. | [optional] 
**DeclaredRareDiseaseExperience** | Pointer to **NullableBool** | Therapeutic rows: true where the supplier says it has rare-disease experience, false where it says it has none (a declared negative), null where it did not answer. | [optional] 
**AsOf** | Pointer to **NullableString** | The date the questionnaire was submitted. | [optional] 
**IsExpired** | Pointer to **NullableBool** | Whether the questionnaire is past its 730-day validity at build time. | [optional] 
**ExpiredAt** | Pointer to **NullableTime** | When the questionnaire validity ends. | [optional] 
**QuestionnaireKey** | Pointer to **NullableString** | The questionnaire version the fact came from: rwe_current, or rwe_2019 on a database row. | [optional] 
**QuestionnaireLabel** | Pointer to **NullableString** | The questionnaire version as a label. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the source: none, a first-party source with no license restriction. | [optional] 
**Attribution** | Pointer to **NullableString** | The plain source line to show with a row. | [optional] 
**AuthorityClass** | Pointer to **NullableString** | Always declared: the supplier&#39;s own statement, not verified. | [optional] 

## Methods

### NewSupplierRwdDeclarationsRow

`func NewSupplierRwdDeclarationsRow() *SupplierRwdDeclarationsRow`

NewSupplierRwdDeclarationsRow instantiates a new SupplierRwdDeclarationsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplierRwdDeclarationsRowWithDefaults

`func NewSupplierRwdDeclarationsRowWithDefaults() *SupplierRwdDeclarationsRow`

NewSupplierRwdDeclarationsRowWithDefaults instantiates a new SupplierRwdDeclarationsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderRwdDeclaredSourceId

`func (o *SupplierRwdDeclarationsRow) GetProviderRwdDeclaredSourceId() string`

GetProviderRwdDeclaredSourceId returns the ProviderRwdDeclaredSourceId field if non-nil, zero value otherwise.

### GetProviderRwdDeclaredSourceIdOk

`func (o *SupplierRwdDeclarationsRow) GetProviderRwdDeclaredSourceIdOk() (*string, bool)`

GetProviderRwdDeclaredSourceIdOk returns a tuple with the ProviderRwdDeclaredSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderRwdDeclaredSourceId

`func (o *SupplierRwdDeclarationsRow) SetProviderRwdDeclaredSourceId(v string)`

SetProviderRwdDeclaredSourceId sets ProviderRwdDeclaredSourceId field to given value.

### HasProviderRwdDeclaredSourceId

`func (o *SupplierRwdDeclarationsRow) HasProviderRwdDeclaredSourceId() bool`

HasProviderRwdDeclaredSourceId returns a boolean if a field has been set.

### SetProviderRwdDeclaredSourceIdNil

`func (o *SupplierRwdDeclarationsRow) SetProviderRwdDeclaredSourceIdNil(b bool)`

 SetProviderRwdDeclaredSourceIdNil sets the value for ProviderRwdDeclaredSourceId to be an explicit nil

### UnsetProviderRwdDeclaredSourceId
`func (o *SupplierRwdDeclarationsRow) UnsetProviderRwdDeclaredSourceId()`

UnsetProviderRwdDeclaredSourceId ensures that no value is present for ProviderRwdDeclaredSourceId, not even an explicit nil
### GetProviderId

`func (o *SupplierRwdDeclarationsRow) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *SupplierRwdDeclarationsRow) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *SupplierRwdDeclarationsRow) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *SupplierRwdDeclarationsRow) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *SupplierRwdDeclarationsRow) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *SupplierRwdDeclarationsRow) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetProviderUuid

`func (o *SupplierRwdDeclarationsRow) GetProviderUuid() string`

GetProviderUuid returns the ProviderUuid field if non-nil, zero value otherwise.

### GetProviderUuidOk

`func (o *SupplierRwdDeclarationsRow) GetProviderUuidOk() (*string, bool)`

GetProviderUuidOk returns a tuple with the ProviderUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUuid

`func (o *SupplierRwdDeclarationsRow) SetProviderUuid(v string)`

SetProviderUuid sets ProviderUuid field to given value.

### HasProviderUuid

`func (o *SupplierRwdDeclarationsRow) HasProviderUuid() bool`

HasProviderUuid returns a boolean if a field has been set.

### SetProviderUuidNil

`func (o *SupplierRwdDeclarationsRow) SetProviderUuidNil(b bool)`

 SetProviderUuidNil sets the value for ProviderUuid to be an explicit nil

### UnsetProviderUuid
`func (o *SupplierRwdDeclarationsRow) UnsetProviderUuid()`

UnsetProviderUuid ensures that no value is present for ProviderUuid, not even an explicit nil
### GetProviderSlug

`func (o *SupplierRwdDeclarationsRow) GetProviderSlug() string`

GetProviderSlug returns the ProviderSlug field if non-nil, zero value otherwise.

### GetProviderSlugOk

`func (o *SupplierRwdDeclarationsRow) GetProviderSlugOk() (*string, bool)`

GetProviderSlugOk returns a tuple with the ProviderSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSlug

`func (o *SupplierRwdDeclarationsRow) SetProviderSlug(v string)`

SetProviderSlug sets ProviderSlug field to given value.

### HasProviderSlug

`func (o *SupplierRwdDeclarationsRow) HasProviderSlug() bool`

HasProviderSlug returns a boolean if a field has been set.

### SetProviderSlugNil

`func (o *SupplierRwdDeclarationsRow) SetProviderSlugNil(b bool)`

 SetProviderSlugNil sets the value for ProviderSlug to be an explicit nil

### UnsetProviderSlug
`func (o *SupplierRwdDeclarationsRow) UnsetProviderSlug()`

UnsetProviderSlug ensures that no value is present for ProviderSlug, not even an explicit nil
### GetProviderName

`func (o *SupplierRwdDeclarationsRow) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *SupplierRwdDeclarationsRow) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *SupplierRwdDeclarationsRow) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *SupplierRwdDeclarationsRow) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *SupplierRwdDeclarationsRow) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *SupplierRwdDeclarationsRow) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetRole

`func (o *SupplierRwdDeclarationsRow) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SupplierRwdDeclarationsRow) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SupplierRwdDeclarationsRow) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *SupplierRwdDeclarationsRow) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *SupplierRwdDeclarationsRow) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *SupplierRwdDeclarationsRow) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetEvidenceRefKind

`func (o *SupplierRwdDeclarationsRow) GetEvidenceRefKind() string`

GetEvidenceRefKind returns the EvidenceRefKind field if non-nil, zero value otherwise.

### GetEvidenceRefKindOk

`func (o *SupplierRwdDeclarationsRow) GetEvidenceRefKindOk() (*string, bool)`

GetEvidenceRefKindOk returns a tuple with the EvidenceRefKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceRefKind

`func (o *SupplierRwdDeclarationsRow) SetEvidenceRefKind(v string)`

SetEvidenceRefKind sets EvidenceRefKind field to given value.

### HasEvidenceRefKind

`func (o *SupplierRwdDeclarationsRow) HasEvidenceRefKind() bool`

HasEvidenceRefKind returns a boolean if a field has been set.

### SetEvidenceRefKindNil

`func (o *SupplierRwdDeclarationsRow) SetEvidenceRefKindNil(b bool)`

 SetEvidenceRefKindNil sets the value for EvidenceRefKind to be an explicit nil

### UnsetEvidenceRefKind
`func (o *SupplierRwdDeclarationsRow) UnsetEvidenceRefKind()`

UnsetEvidenceRefKind ensures that no value is present for EvidenceRefKind, not even an explicit nil
### GetEvidenceRef

`func (o *SupplierRwdDeclarationsRow) GetEvidenceRef() string`

GetEvidenceRef returns the EvidenceRef field if non-nil, zero value otherwise.

### GetEvidenceRefOk

`func (o *SupplierRwdDeclarationsRow) GetEvidenceRefOk() (*string, bool)`

GetEvidenceRefOk returns a tuple with the EvidenceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceRef

`func (o *SupplierRwdDeclarationsRow) SetEvidenceRef(v string)`

SetEvidenceRef sets EvidenceRef field to given value.

### HasEvidenceRef

`func (o *SupplierRwdDeclarationsRow) HasEvidenceRef() bool`

HasEvidenceRef returns a boolean if a field has been set.

### SetEvidenceRefNil

`func (o *SupplierRwdDeclarationsRow) SetEvidenceRefNil(b bool)`

 SetEvidenceRefNil sets the value for EvidenceRef to be an explicit nil

### UnsetEvidenceRef
`func (o *SupplierRwdDeclarationsRow) UnsetEvidenceRef()`

UnsetEvidenceRef ensures that no value is present for EvidenceRef, not even an explicit nil
### GetEvidenceLabel

`func (o *SupplierRwdDeclarationsRow) GetEvidenceLabel() string`

GetEvidenceLabel returns the EvidenceLabel field if non-nil, zero value otherwise.

### GetEvidenceLabelOk

`func (o *SupplierRwdDeclarationsRow) GetEvidenceLabelOk() (*string, bool)`

GetEvidenceLabelOk returns a tuple with the EvidenceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceLabel

`func (o *SupplierRwdDeclarationsRow) SetEvidenceLabel(v string)`

SetEvidenceLabel sets EvidenceLabel field to given value.

### HasEvidenceLabel

`func (o *SupplierRwdDeclarationsRow) HasEvidenceLabel() bool`

HasEvidenceLabel returns a boolean if a field has been set.

### SetEvidenceLabelNil

`func (o *SupplierRwdDeclarationsRow) SetEvidenceLabelNil(b bool)`

 SetEvidenceLabelNil sets the value for EvidenceLabel to be an explicit nil

### UnsetEvidenceLabel
`func (o *SupplierRwdDeclarationsRow) UnsetEvidenceLabel()`

UnsetEvidenceLabel ensures that no value is present for EvidenceLabel, not even an explicit nil
### GetMatchMethod

`func (o *SupplierRwdDeclarationsRow) GetMatchMethod() string`

GetMatchMethod returns the MatchMethod field if non-nil, zero value otherwise.

### GetMatchMethodOk

`func (o *SupplierRwdDeclarationsRow) GetMatchMethodOk() (*string, bool)`

GetMatchMethodOk returns a tuple with the MatchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMethod

`func (o *SupplierRwdDeclarationsRow) SetMatchMethod(v string)`

SetMatchMethod sets MatchMethod field to given value.

### HasMatchMethod

`func (o *SupplierRwdDeclarationsRow) HasMatchMethod() bool`

HasMatchMethod returns a boolean if a field has been set.

### SetMatchMethodNil

`func (o *SupplierRwdDeclarationsRow) SetMatchMethodNil(b bool)`

 SetMatchMethodNil sets the value for MatchMethod to be an explicit nil

### UnsetMatchMethod
`func (o *SupplierRwdDeclarationsRow) UnsetMatchMethod()`

UnsetMatchMethod ensures that no value is present for MatchMethod, not even an explicit nil
### GetCountryNames

`func (o *SupplierRwdDeclarationsRow) GetCountryNames() []string`

GetCountryNames returns the CountryNames field if non-nil, zero value otherwise.

### GetCountryNamesOk

`func (o *SupplierRwdDeclarationsRow) GetCountryNamesOk() (*[]string, bool)`

GetCountryNamesOk returns a tuple with the CountryNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryNames

`func (o *SupplierRwdDeclarationsRow) SetCountryNames(v []string)`

SetCountryNames sets CountryNames field to given value.

### HasCountryNames

`func (o *SupplierRwdDeclarationsRow) HasCountryNames() bool`

HasCountryNames returns a boolean if a field has been set.

### SetCountryNamesNil

`func (o *SupplierRwdDeclarationsRow) SetCountryNamesNil(b bool)`

 SetCountryNamesNil sets the value for CountryNames to be an explicit nil

### UnsetCountryNames
`func (o *SupplierRwdDeclarationsRow) UnsetCountryNames()`

UnsetCountryNames ensures that no value is present for CountryNames, not even an explicit nil
### GetDeclaredRegions

`func (o *SupplierRwdDeclarationsRow) GetDeclaredRegions() []string`

GetDeclaredRegions returns the DeclaredRegions field if non-nil, zero value otherwise.

### GetDeclaredRegionsOk

`func (o *SupplierRwdDeclarationsRow) GetDeclaredRegionsOk() (*[]string, bool)`

GetDeclaredRegionsOk returns a tuple with the DeclaredRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredRegions

`func (o *SupplierRwdDeclarationsRow) SetDeclaredRegions(v []string)`

SetDeclaredRegions sets DeclaredRegions field to given value.

### HasDeclaredRegions

`func (o *SupplierRwdDeclarationsRow) HasDeclaredRegions() bool`

HasDeclaredRegions returns a boolean if a field has been set.

### SetDeclaredRegionsNil

`func (o *SupplierRwdDeclarationsRow) SetDeclaredRegionsNil(b bool)`

 SetDeclaredRegionsNil sets the value for DeclaredRegions to be an explicit nil

### UnsetDeclaredRegions
`func (o *SupplierRwdDeclarationsRow) UnsetDeclaredRegions()`

UnsetDeclaredRegions ensures that no value is present for DeclaredRegions, not even an explicit nil
### GetDataTypes

`func (o *SupplierRwdDeclarationsRow) GetDataTypes() []string`

GetDataTypes returns the DataTypes field if non-nil, zero value otherwise.

### GetDataTypesOk

`func (o *SupplierRwdDeclarationsRow) GetDataTypesOk() (*[]string, bool)`

GetDataTypesOk returns a tuple with the DataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypes

`func (o *SupplierRwdDeclarationsRow) SetDataTypes(v []string)`

SetDataTypes sets DataTypes field to given value.

### HasDataTypes

`func (o *SupplierRwdDeclarationsRow) HasDataTypes() bool`

HasDataTypes returns a boolean if a field has been set.

### SetDataTypesNil

`func (o *SupplierRwdDeclarationsRow) SetDataTypesNil(b bool)`

 SetDataTypesNil sets the value for DataTypes to be an explicit nil

### UnsetDataTypes
`func (o *SupplierRwdDeclarationsRow) UnsetDataTypes()`

UnsetDataTypes ensures that no value is present for DataTypes, not even an explicit nil
### GetDataSourceType

`func (o *SupplierRwdDeclarationsRow) GetDataSourceType() string`

GetDataSourceType returns the DataSourceType field if non-nil, zero value otherwise.

### GetDataSourceTypeOk

`func (o *SupplierRwdDeclarationsRow) GetDataSourceTypeOk() (*string, bool)`

GetDataSourceTypeOk returns a tuple with the DataSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceType

`func (o *SupplierRwdDeclarationsRow) SetDataSourceType(v string)`

SetDataSourceType sets DataSourceType field to given value.

### HasDataSourceType

`func (o *SupplierRwdDeclarationsRow) HasDataSourceType() bool`

HasDataSourceType returns a boolean if a field has been set.

### SetDataSourceTypeNil

`func (o *SupplierRwdDeclarationsRow) SetDataSourceTypeNil(b bool)`

 SetDataSourceTypeNil sets the value for DataSourceType to be an explicit nil

### UnsetDataSourceType
`func (o *SupplierRwdDeclarationsRow) UnsetDataSourceType()`

UnsetDataSourceType ensures that no value is present for DataSourceType, not even an explicit nil
### GetDataTypesBasis

`func (o *SupplierRwdDeclarationsRow) GetDataTypesBasis() string`

GetDataTypesBasis returns the DataTypesBasis field if non-nil, zero value otherwise.

### GetDataTypesBasisOk

`func (o *SupplierRwdDeclarationsRow) GetDataTypesBasisOk() (*string, bool)`

GetDataTypesBasisOk returns a tuple with the DataTypesBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypesBasis

`func (o *SupplierRwdDeclarationsRow) SetDataTypesBasis(v string)`

SetDataTypesBasis sets DataTypesBasis field to given value.

### HasDataTypesBasis

`func (o *SupplierRwdDeclarationsRow) HasDataTypesBasis() bool`

HasDataTypesBasis returns a boolean if a field has been set.

### SetDataTypesBasisNil

`func (o *SupplierRwdDeclarationsRow) SetDataTypesBasisNil(b bool)`

 SetDataTypesBasisNil sets the value for DataTypesBasis to be an explicit nil

### UnsetDataTypesBasis
`func (o *SupplierRwdDeclarationsRow) UnsetDataTypesBasis()`

UnsetDataTypesBasis ensures that no value is present for DataTypesBasis, not even an explicit nil
### GetDeclaredValidation

`func (o *SupplierRwdDeclarationsRow) GetDeclaredValidation() string`

GetDeclaredValidation returns the DeclaredValidation field if non-nil, zero value otherwise.

### GetDeclaredValidationOk

`func (o *SupplierRwdDeclarationsRow) GetDeclaredValidationOk() (*string, bool)`

GetDeclaredValidationOk returns a tuple with the DeclaredValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValidation

`func (o *SupplierRwdDeclarationsRow) SetDeclaredValidation(v string)`

SetDeclaredValidation sets DeclaredValidation field to given value.

### HasDeclaredValidation

`func (o *SupplierRwdDeclarationsRow) HasDeclaredValidation() bool`

HasDeclaredValidation returns a boolean if a field has been set.

### SetDeclaredValidationNil

`func (o *SupplierRwdDeclarationsRow) SetDeclaredValidationNil(b bool)`

 SetDeclaredValidationNil sets the value for DeclaredValidation to be an explicit nil

### UnsetDeclaredValidation
`func (o *SupplierRwdDeclarationsRow) UnsetDeclaredValidation()`

UnsetDeclaredValidation ensures that no value is present for DeclaredValidation, not even an explicit nil
### GetDeclaredFirstCollectionYear

`func (o *SupplierRwdDeclarationsRow) GetDeclaredFirstCollectionYear() int32`

GetDeclaredFirstCollectionYear returns the DeclaredFirstCollectionYear field if non-nil, zero value otherwise.

### GetDeclaredFirstCollectionYearOk

`func (o *SupplierRwdDeclarationsRow) GetDeclaredFirstCollectionYearOk() (*int32, bool)`

GetDeclaredFirstCollectionYearOk returns a tuple with the DeclaredFirstCollectionYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredFirstCollectionYear

`func (o *SupplierRwdDeclarationsRow) SetDeclaredFirstCollectionYear(v int32)`

SetDeclaredFirstCollectionYear sets DeclaredFirstCollectionYear field to given value.

### HasDeclaredFirstCollectionYear

`func (o *SupplierRwdDeclarationsRow) HasDeclaredFirstCollectionYear() bool`

HasDeclaredFirstCollectionYear returns a boolean if a field has been set.

### SetDeclaredFirstCollectionYearNil

`func (o *SupplierRwdDeclarationsRow) SetDeclaredFirstCollectionYearNil(b bool)`

 SetDeclaredFirstCollectionYearNil sets the value for DeclaredFirstCollectionYear to be an explicit nil

### UnsetDeclaredFirstCollectionYear
`func (o *SupplierRwdDeclarationsRow) UnsetDeclaredFirstCollectionYear()`

UnsetDeclaredFirstCollectionYear ensures that no value is present for DeclaredFirstCollectionYear, not even an explicit nil
### GetYearImplausible

`func (o *SupplierRwdDeclarationsRow) GetYearImplausible() bool`

GetYearImplausible returns the YearImplausible field if non-nil, zero value otherwise.

### GetYearImplausibleOk

`func (o *SupplierRwdDeclarationsRow) GetYearImplausibleOk() (*bool, bool)`

GetYearImplausibleOk returns a tuple with the YearImplausible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearImplausible

`func (o *SupplierRwdDeclarationsRow) SetYearImplausible(v bool)`

SetYearImplausible sets YearImplausible field to given value.

### HasYearImplausible

`func (o *SupplierRwdDeclarationsRow) HasYearImplausible() bool`

HasYearImplausible returns a boolean if a field has been set.

### SetYearImplausibleNil

`func (o *SupplierRwdDeclarationsRow) SetYearImplausibleNil(b bool)`

 SetYearImplausibleNil sets the value for YearImplausible to be an explicit nil

### UnsetYearImplausible
`func (o *SupplierRwdDeclarationsRow) UnsetYearImplausible()`

UnsetYearImplausible ensures that no value is present for YearImplausible, not even an explicit nil
### GetDeclaredLinkable

`func (o *SupplierRwdDeclarationsRow) GetDeclaredLinkable() bool`

GetDeclaredLinkable returns the DeclaredLinkable field if non-nil, zero value otherwise.

### GetDeclaredLinkableOk

`func (o *SupplierRwdDeclarationsRow) GetDeclaredLinkableOk() (*bool, bool)`

GetDeclaredLinkableOk returns a tuple with the DeclaredLinkable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredLinkable

`func (o *SupplierRwdDeclarationsRow) SetDeclaredLinkable(v bool)`

SetDeclaredLinkable sets DeclaredLinkable field to given value.

### HasDeclaredLinkable

`func (o *SupplierRwdDeclarationsRow) HasDeclaredLinkable() bool`

HasDeclaredLinkable returns a boolean if a field has been set.

### SetDeclaredLinkableNil

`func (o *SupplierRwdDeclarationsRow) SetDeclaredLinkableNil(b bool)`

 SetDeclaredLinkableNil sets the value for DeclaredLinkable to be an explicit nil

### UnsetDeclaredLinkable
`func (o *SupplierRwdDeclarationsRow) UnsetDeclaredLinkable()`

UnsetDeclaredLinkable ensures that no value is present for DeclaredLinkable, not even an explicit nil
### GetTherapeuticAreas

`func (o *SupplierRwdDeclarationsRow) GetTherapeuticAreas() []string`

GetTherapeuticAreas returns the TherapeuticAreas field if non-nil, zero value otherwise.

### GetTherapeuticAreasOk

`func (o *SupplierRwdDeclarationsRow) GetTherapeuticAreasOk() (*[]string, bool)`

GetTherapeuticAreasOk returns a tuple with the TherapeuticAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTherapeuticAreas

`func (o *SupplierRwdDeclarationsRow) SetTherapeuticAreas(v []string)`

SetTherapeuticAreas sets TherapeuticAreas field to given value.

### HasTherapeuticAreas

`func (o *SupplierRwdDeclarationsRow) HasTherapeuticAreas() bool`

HasTherapeuticAreas returns a boolean if a field has been set.

### SetTherapeuticAreasNil

`func (o *SupplierRwdDeclarationsRow) SetTherapeuticAreasNil(b bool)`

 SetTherapeuticAreasNil sets the value for TherapeuticAreas to be an explicit nil

### UnsetTherapeuticAreas
`func (o *SupplierRwdDeclarationsRow) UnsetTherapeuticAreas()`

UnsetTherapeuticAreas ensures that no value is present for TherapeuticAreas, not even an explicit nil
### GetDeclaredRareDiseaseExperience

`func (o *SupplierRwdDeclarationsRow) GetDeclaredRareDiseaseExperience() bool`

GetDeclaredRareDiseaseExperience returns the DeclaredRareDiseaseExperience field if non-nil, zero value otherwise.

### GetDeclaredRareDiseaseExperienceOk

`func (o *SupplierRwdDeclarationsRow) GetDeclaredRareDiseaseExperienceOk() (*bool, bool)`

GetDeclaredRareDiseaseExperienceOk returns a tuple with the DeclaredRareDiseaseExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredRareDiseaseExperience

`func (o *SupplierRwdDeclarationsRow) SetDeclaredRareDiseaseExperience(v bool)`

SetDeclaredRareDiseaseExperience sets DeclaredRareDiseaseExperience field to given value.

### HasDeclaredRareDiseaseExperience

`func (o *SupplierRwdDeclarationsRow) HasDeclaredRareDiseaseExperience() bool`

HasDeclaredRareDiseaseExperience returns a boolean if a field has been set.

### SetDeclaredRareDiseaseExperienceNil

`func (o *SupplierRwdDeclarationsRow) SetDeclaredRareDiseaseExperienceNil(b bool)`

 SetDeclaredRareDiseaseExperienceNil sets the value for DeclaredRareDiseaseExperience to be an explicit nil

### UnsetDeclaredRareDiseaseExperience
`func (o *SupplierRwdDeclarationsRow) UnsetDeclaredRareDiseaseExperience()`

UnsetDeclaredRareDiseaseExperience ensures that no value is present for DeclaredRareDiseaseExperience, not even an explicit nil
### GetAsOf

`func (o *SupplierRwdDeclarationsRow) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *SupplierRwdDeclarationsRow) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *SupplierRwdDeclarationsRow) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *SupplierRwdDeclarationsRow) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *SupplierRwdDeclarationsRow) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *SupplierRwdDeclarationsRow) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil
### GetIsExpired

`func (o *SupplierRwdDeclarationsRow) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *SupplierRwdDeclarationsRow) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *SupplierRwdDeclarationsRow) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *SupplierRwdDeclarationsRow) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### SetIsExpiredNil

`func (o *SupplierRwdDeclarationsRow) SetIsExpiredNil(b bool)`

 SetIsExpiredNil sets the value for IsExpired to be an explicit nil

### UnsetIsExpired
`func (o *SupplierRwdDeclarationsRow) UnsetIsExpired()`

UnsetIsExpired ensures that no value is present for IsExpired, not even an explicit nil
### GetExpiredAt

`func (o *SupplierRwdDeclarationsRow) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *SupplierRwdDeclarationsRow) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *SupplierRwdDeclarationsRow) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *SupplierRwdDeclarationsRow) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### SetExpiredAtNil

`func (o *SupplierRwdDeclarationsRow) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *SupplierRwdDeclarationsRow) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetQuestionnaireKey

`func (o *SupplierRwdDeclarationsRow) GetQuestionnaireKey() string`

GetQuestionnaireKey returns the QuestionnaireKey field if non-nil, zero value otherwise.

### GetQuestionnaireKeyOk

`func (o *SupplierRwdDeclarationsRow) GetQuestionnaireKeyOk() (*string, bool)`

GetQuestionnaireKeyOk returns a tuple with the QuestionnaireKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaireKey

`func (o *SupplierRwdDeclarationsRow) SetQuestionnaireKey(v string)`

SetQuestionnaireKey sets QuestionnaireKey field to given value.

### HasQuestionnaireKey

`func (o *SupplierRwdDeclarationsRow) HasQuestionnaireKey() bool`

HasQuestionnaireKey returns a boolean if a field has been set.

### SetQuestionnaireKeyNil

`func (o *SupplierRwdDeclarationsRow) SetQuestionnaireKeyNil(b bool)`

 SetQuestionnaireKeyNil sets the value for QuestionnaireKey to be an explicit nil

### UnsetQuestionnaireKey
`func (o *SupplierRwdDeclarationsRow) UnsetQuestionnaireKey()`

UnsetQuestionnaireKey ensures that no value is present for QuestionnaireKey, not even an explicit nil
### GetQuestionnaireLabel

`func (o *SupplierRwdDeclarationsRow) GetQuestionnaireLabel() string`

GetQuestionnaireLabel returns the QuestionnaireLabel field if non-nil, zero value otherwise.

### GetQuestionnaireLabelOk

`func (o *SupplierRwdDeclarationsRow) GetQuestionnaireLabelOk() (*string, bool)`

GetQuestionnaireLabelOk returns a tuple with the QuestionnaireLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaireLabel

`func (o *SupplierRwdDeclarationsRow) SetQuestionnaireLabel(v string)`

SetQuestionnaireLabel sets QuestionnaireLabel field to given value.

### HasQuestionnaireLabel

`func (o *SupplierRwdDeclarationsRow) HasQuestionnaireLabel() bool`

HasQuestionnaireLabel returns a boolean if a field has been set.

### SetQuestionnaireLabelNil

`func (o *SupplierRwdDeclarationsRow) SetQuestionnaireLabelNil(b bool)`

 SetQuestionnaireLabelNil sets the value for QuestionnaireLabel to be an explicit nil

### UnsetQuestionnaireLabel
`func (o *SupplierRwdDeclarationsRow) UnsetQuestionnaireLabel()`

UnsetQuestionnaireLabel ensures that no value is present for QuestionnaireLabel, not even an explicit nil
### GetRestrictionTier

`func (o *SupplierRwdDeclarationsRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *SupplierRwdDeclarationsRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *SupplierRwdDeclarationsRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *SupplierRwdDeclarationsRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *SupplierRwdDeclarationsRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *SupplierRwdDeclarationsRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *SupplierRwdDeclarationsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *SupplierRwdDeclarationsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *SupplierRwdDeclarationsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *SupplierRwdDeclarationsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *SupplierRwdDeclarationsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *SupplierRwdDeclarationsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil
### GetAuthorityClass

`func (o *SupplierRwdDeclarationsRow) GetAuthorityClass() string`

GetAuthorityClass returns the AuthorityClass field if non-nil, zero value otherwise.

### GetAuthorityClassOk

`func (o *SupplierRwdDeclarationsRow) GetAuthorityClassOk() (*string, bool)`

GetAuthorityClassOk returns a tuple with the AuthorityClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityClass

`func (o *SupplierRwdDeclarationsRow) SetAuthorityClass(v string)`

SetAuthorityClass sets AuthorityClass field to given value.

### HasAuthorityClass

`func (o *SupplierRwdDeclarationsRow) HasAuthorityClass() bool`

HasAuthorityClass returns a boolean if a field has been set.

### SetAuthorityClassNil

`func (o *SupplierRwdDeclarationsRow) SetAuthorityClassNil(b bool)`

 SetAuthorityClassNil sets the value for AuthorityClass to be an explicit nil

### UnsetAuthorityClass
`func (o *SupplierRwdDeclarationsRow) UnsetAuthorityClass()`

UnsetAuthorityClass ensures that no value is present for AuthorityClass, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


