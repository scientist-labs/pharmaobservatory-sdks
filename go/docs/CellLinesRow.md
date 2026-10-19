# CellLinesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CellLineId** | Pointer to **NullableString** | The catalog listing&#39;s model id (scientist_model_id). | [optional] 
**AlternateCellLineIds** | Pointer to **NullableString** | Other catalog model ids for the same line. | [optional] 
**ModelLotId** | Pointer to **NullableString** | The lot id. | [optional] 
**CellType** | Pointer to **NullableString** | The cell line&#39;s name. | [optional] 
**ProviderId** | Pointer to **NullableString** | The supplier&#39;s catalog id. | [optional] 
**Supplier** | Pointer to **NullableString** | The supplier as listed. | [optional] 
**SampleSpecies** | Pointer to **NullableString** | The species. | [optional] 
**StorefrontEnable** | Pointer to **NullableBool** | Whether the listing is enabled in the storefront. | [optional] 
**MutationData** | Pointer to **NullableString** | Mutation data as listed. | [optional] 
**Status** | Pointer to **NullableString** | The listing status. | [optional] 
**SafetyStatus** | Pointer to **NullableString** | confirmed_problematic, review_candidate, no_problematic_match or missing_cell_type. | [optional] 
**ResolutionStatus** | Pointer to **NullableString** | resolved, unresolved_review, no_match or missing_input. | [optional] 
**ConfirmedProblematic** | Pointer to **NullableBool** | Whether the line is confirmed problematic in Cellosaurus. | [optional] 
**CellosaurusAccession** | Pointer to **NullableString** | The displayed Cellosaurus accession. | [optional] 
**CellosaurusUrl** | Pointer to **NullableString** | Link to the Cellosaurus entry. | [optional] 
**ProblemCategory** | Pointer to **NullableString** | The problem category, for a problematic line. | [optional] 
**ProblematicComment** | Pointer to **NullableString** | The Cellosaurus comment, for a problematic line. | [optional] 
**IclacRegistrationIds** | Pointer to **[]string** | International Cell Line Authentication Committee (ICLAC) registration ids. | [optional] 
**ReleaseVersion** | Pointer to **NullableString** | The Cellosaurus release the check ran against. | [optional] 

## Methods

### NewCellLinesRow

`func NewCellLinesRow() *CellLinesRow`

NewCellLinesRow instantiates a new CellLinesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellLinesRowWithDefaults

`func NewCellLinesRowWithDefaults() *CellLinesRow`

NewCellLinesRowWithDefaults instantiates a new CellLinesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCellLineId

`func (o *CellLinesRow) GetCellLineId() string`

GetCellLineId returns the CellLineId field if non-nil, zero value otherwise.

### GetCellLineIdOk

`func (o *CellLinesRow) GetCellLineIdOk() (*string, bool)`

GetCellLineIdOk returns a tuple with the CellLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellLineId

`func (o *CellLinesRow) SetCellLineId(v string)`

SetCellLineId sets CellLineId field to given value.

### HasCellLineId

`func (o *CellLinesRow) HasCellLineId() bool`

HasCellLineId returns a boolean if a field has been set.

### SetCellLineIdNil

`func (o *CellLinesRow) SetCellLineIdNil(b bool)`

 SetCellLineIdNil sets the value for CellLineId to be an explicit nil

### UnsetCellLineId
`func (o *CellLinesRow) UnsetCellLineId()`

UnsetCellLineId ensures that no value is present for CellLineId, not even an explicit nil
### GetAlternateCellLineIds

`func (o *CellLinesRow) GetAlternateCellLineIds() string`

GetAlternateCellLineIds returns the AlternateCellLineIds field if non-nil, zero value otherwise.

### GetAlternateCellLineIdsOk

`func (o *CellLinesRow) GetAlternateCellLineIdsOk() (*string, bool)`

GetAlternateCellLineIdsOk returns a tuple with the AlternateCellLineIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateCellLineIds

`func (o *CellLinesRow) SetAlternateCellLineIds(v string)`

SetAlternateCellLineIds sets AlternateCellLineIds field to given value.

### HasAlternateCellLineIds

`func (o *CellLinesRow) HasAlternateCellLineIds() bool`

HasAlternateCellLineIds returns a boolean if a field has been set.

### SetAlternateCellLineIdsNil

`func (o *CellLinesRow) SetAlternateCellLineIdsNil(b bool)`

 SetAlternateCellLineIdsNil sets the value for AlternateCellLineIds to be an explicit nil

### UnsetAlternateCellLineIds
`func (o *CellLinesRow) UnsetAlternateCellLineIds()`

UnsetAlternateCellLineIds ensures that no value is present for AlternateCellLineIds, not even an explicit nil
### GetModelLotId

`func (o *CellLinesRow) GetModelLotId() string`

GetModelLotId returns the ModelLotId field if non-nil, zero value otherwise.

### GetModelLotIdOk

`func (o *CellLinesRow) GetModelLotIdOk() (*string, bool)`

GetModelLotIdOk returns a tuple with the ModelLotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelLotId

`func (o *CellLinesRow) SetModelLotId(v string)`

SetModelLotId sets ModelLotId field to given value.

### HasModelLotId

`func (o *CellLinesRow) HasModelLotId() bool`

HasModelLotId returns a boolean if a field has been set.

### SetModelLotIdNil

`func (o *CellLinesRow) SetModelLotIdNil(b bool)`

 SetModelLotIdNil sets the value for ModelLotId to be an explicit nil

### UnsetModelLotId
`func (o *CellLinesRow) UnsetModelLotId()`

UnsetModelLotId ensures that no value is present for ModelLotId, not even an explicit nil
### GetCellType

`func (o *CellLinesRow) GetCellType() string`

GetCellType returns the CellType field if non-nil, zero value otherwise.

### GetCellTypeOk

`func (o *CellLinesRow) GetCellTypeOk() (*string, bool)`

GetCellTypeOk returns a tuple with the CellType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellType

`func (o *CellLinesRow) SetCellType(v string)`

SetCellType sets CellType field to given value.

### HasCellType

`func (o *CellLinesRow) HasCellType() bool`

HasCellType returns a boolean if a field has been set.

### SetCellTypeNil

`func (o *CellLinesRow) SetCellTypeNil(b bool)`

 SetCellTypeNil sets the value for CellType to be an explicit nil

### UnsetCellType
`func (o *CellLinesRow) UnsetCellType()`

UnsetCellType ensures that no value is present for CellType, not even an explicit nil
### GetProviderId

`func (o *CellLinesRow) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CellLinesRow) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CellLinesRow) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *CellLinesRow) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *CellLinesRow) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *CellLinesRow) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetSupplier

`func (o *CellLinesRow) GetSupplier() string`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *CellLinesRow) GetSupplierOk() (*string, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *CellLinesRow) SetSupplier(v string)`

SetSupplier sets Supplier field to given value.

### HasSupplier

`func (o *CellLinesRow) HasSupplier() bool`

HasSupplier returns a boolean if a field has been set.

### SetSupplierNil

`func (o *CellLinesRow) SetSupplierNil(b bool)`

 SetSupplierNil sets the value for Supplier to be an explicit nil

### UnsetSupplier
`func (o *CellLinesRow) UnsetSupplier()`

UnsetSupplier ensures that no value is present for Supplier, not even an explicit nil
### GetSampleSpecies

`func (o *CellLinesRow) GetSampleSpecies() string`

GetSampleSpecies returns the SampleSpecies field if non-nil, zero value otherwise.

### GetSampleSpeciesOk

`func (o *CellLinesRow) GetSampleSpeciesOk() (*string, bool)`

GetSampleSpeciesOk returns a tuple with the SampleSpecies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSpecies

`func (o *CellLinesRow) SetSampleSpecies(v string)`

SetSampleSpecies sets SampleSpecies field to given value.

### HasSampleSpecies

`func (o *CellLinesRow) HasSampleSpecies() bool`

HasSampleSpecies returns a boolean if a field has been set.

### SetSampleSpeciesNil

`func (o *CellLinesRow) SetSampleSpeciesNil(b bool)`

 SetSampleSpeciesNil sets the value for SampleSpecies to be an explicit nil

### UnsetSampleSpecies
`func (o *CellLinesRow) UnsetSampleSpecies()`

UnsetSampleSpecies ensures that no value is present for SampleSpecies, not even an explicit nil
### GetStorefrontEnable

`func (o *CellLinesRow) GetStorefrontEnable() bool`

GetStorefrontEnable returns the StorefrontEnable field if non-nil, zero value otherwise.

### GetStorefrontEnableOk

`func (o *CellLinesRow) GetStorefrontEnableOk() (*bool, bool)`

GetStorefrontEnableOk returns a tuple with the StorefrontEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorefrontEnable

`func (o *CellLinesRow) SetStorefrontEnable(v bool)`

SetStorefrontEnable sets StorefrontEnable field to given value.

### HasStorefrontEnable

`func (o *CellLinesRow) HasStorefrontEnable() bool`

HasStorefrontEnable returns a boolean if a field has been set.

### SetStorefrontEnableNil

`func (o *CellLinesRow) SetStorefrontEnableNil(b bool)`

 SetStorefrontEnableNil sets the value for StorefrontEnable to be an explicit nil

### UnsetStorefrontEnable
`func (o *CellLinesRow) UnsetStorefrontEnable()`

UnsetStorefrontEnable ensures that no value is present for StorefrontEnable, not even an explicit nil
### GetMutationData

`func (o *CellLinesRow) GetMutationData() string`

GetMutationData returns the MutationData field if non-nil, zero value otherwise.

### GetMutationDataOk

`func (o *CellLinesRow) GetMutationDataOk() (*string, bool)`

GetMutationDataOk returns a tuple with the MutationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutationData

`func (o *CellLinesRow) SetMutationData(v string)`

SetMutationData sets MutationData field to given value.

### HasMutationData

`func (o *CellLinesRow) HasMutationData() bool`

HasMutationData returns a boolean if a field has been set.

### SetMutationDataNil

`func (o *CellLinesRow) SetMutationDataNil(b bool)`

 SetMutationDataNil sets the value for MutationData to be an explicit nil

### UnsetMutationData
`func (o *CellLinesRow) UnsetMutationData()`

UnsetMutationData ensures that no value is present for MutationData, not even an explicit nil
### GetStatus

`func (o *CellLinesRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CellLinesRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CellLinesRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CellLinesRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CellLinesRow) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CellLinesRow) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSafetyStatus

`func (o *CellLinesRow) GetSafetyStatus() string`

GetSafetyStatus returns the SafetyStatus field if non-nil, zero value otherwise.

### GetSafetyStatusOk

`func (o *CellLinesRow) GetSafetyStatusOk() (*string, bool)`

GetSafetyStatusOk returns a tuple with the SafetyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyStatus

`func (o *CellLinesRow) SetSafetyStatus(v string)`

SetSafetyStatus sets SafetyStatus field to given value.

### HasSafetyStatus

`func (o *CellLinesRow) HasSafetyStatus() bool`

HasSafetyStatus returns a boolean if a field has been set.

### SetSafetyStatusNil

`func (o *CellLinesRow) SetSafetyStatusNil(b bool)`

 SetSafetyStatusNil sets the value for SafetyStatus to be an explicit nil

### UnsetSafetyStatus
`func (o *CellLinesRow) UnsetSafetyStatus()`

UnsetSafetyStatus ensures that no value is present for SafetyStatus, not even an explicit nil
### GetResolutionStatus

`func (o *CellLinesRow) GetResolutionStatus() string`

GetResolutionStatus returns the ResolutionStatus field if non-nil, zero value otherwise.

### GetResolutionStatusOk

`func (o *CellLinesRow) GetResolutionStatusOk() (*string, bool)`

GetResolutionStatusOk returns a tuple with the ResolutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionStatus

`func (o *CellLinesRow) SetResolutionStatus(v string)`

SetResolutionStatus sets ResolutionStatus field to given value.

### HasResolutionStatus

`func (o *CellLinesRow) HasResolutionStatus() bool`

HasResolutionStatus returns a boolean if a field has been set.

### SetResolutionStatusNil

`func (o *CellLinesRow) SetResolutionStatusNil(b bool)`

 SetResolutionStatusNil sets the value for ResolutionStatus to be an explicit nil

### UnsetResolutionStatus
`func (o *CellLinesRow) UnsetResolutionStatus()`

UnsetResolutionStatus ensures that no value is present for ResolutionStatus, not even an explicit nil
### GetConfirmedProblematic

`func (o *CellLinesRow) GetConfirmedProblematic() bool`

GetConfirmedProblematic returns the ConfirmedProblematic field if non-nil, zero value otherwise.

### GetConfirmedProblematicOk

`func (o *CellLinesRow) GetConfirmedProblematicOk() (*bool, bool)`

GetConfirmedProblematicOk returns a tuple with the ConfirmedProblematic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedProblematic

`func (o *CellLinesRow) SetConfirmedProblematic(v bool)`

SetConfirmedProblematic sets ConfirmedProblematic field to given value.

### HasConfirmedProblematic

`func (o *CellLinesRow) HasConfirmedProblematic() bool`

HasConfirmedProblematic returns a boolean if a field has been set.

### SetConfirmedProblematicNil

`func (o *CellLinesRow) SetConfirmedProblematicNil(b bool)`

 SetConfirmedProblematicNil sets the value for ConfirmedProblematic to be an explicit nil

### UnsetConfirmedProblematic
`func (o *CellLinesRow) UnsetConfirmedProblematic()`

UnsetConfirmedProblematic ensures that no value is present for ConfirmedProblematic, not even an explicit nil
### GetCellosaurusAccession

`func (o *CellLinesRow) GetCellosaurusAccession() string`

GetCellosaurusAccession returns the CellosaurusAccession field if non-nil, zero value otherwise.

### GetCellosaurusAccessionOk

`func (o *CellLinesRow) GetCellosaurusAccessionOk() (*string, bool)`

GetCellosaurusAccessionOk returns a tuple with the CellosaurusAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellosaurusAccession

`func (o *CellLinesRow) SetCellosaurusAccession(v string)`

SetCellosaurusAccession sets CellosaurusAccession field to given value.

### HasCellosaurusAccession

`func (o *CellLinesRow) HasCellosaurusAccession() bool`

HasCellosaurusAccession returns a boolean if a field has been set.

### SetCellosaurusAccessionNil

`func (o *CellLinesRow) SetCellosaurusAccessionNil(b bool)`

 SetCellosaurusAccessionNil sets the value for CellosaurusAccession to be an explicit nil

### UnsetCellosaurusAccession
`func (o *CellLinesRow) UnsetCellosaurusAccession()`

UnsetCellosaurusAccession ensures that no value is present for CellosaurusAccession, not even an explicit nil
### GetCellosaurusUrl

`func (o *CellLinesRow) GetCellosaurusUrl() string`

GetCellosaurusUrl returns the CellosaurusUrl field if non-nil, zero value otherwise.

### GetCellosaurusUrlOk

`func (o *CellLinesRow) GetCellosaurusUrlOk() (*string, bool)`

GetCellosaurusUrlOk returns a tuple with the CellosaurusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellosaurusUrl

`func (o *CellLinesRow) SetCellosaurusUrl(v string)`

SetCellosaurusUrl sets CellosaurusUrl field to given value.

### HasCellosaurusUrl

`func (o *CellLinesRow) HasCellosaurusUrl() bool`

HasCellosaurusUrl returns a boolean if a field has been set.

### SetCellosaurusUrlNil

`func (o *CellLinesRow) SetCellosaurusUrlNil(b bool)`

 SetCellosaurusUrlNil sets the value for CellosaurusUrl to be an explicit nil

### UnsetCellosaurusUrl
`func (o *CellLinesRow) UnsetCellosaurusUrl()`

UnsetCellosaurusUrl ensures that no value is present for CellosaurusUrl, not even an explicit nil
### GetProblemCategory

`func (o *CellLinesRow) GetProblemCategory() string`

GetProblemCategory returns the ProblemCategory field if non-nil, zero value otherwise.

### GetProblemCategoryOk

`func (o *CellLinesRow) GetProblemCategoryOk() (*string, bool)`

GetProblemCategoryOk returns a tuple with the ProblemCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemCategory

`func (o *CellLinesRow) SetProblemCategory(v string)`

SetProblemCategory sets ProblemCategory field to given value.

### HasProblemCategory

`func (o *CellLinesRow) HasProblemCategory() bool`

HasProblemCategory returns a boolean if a field has been set.

### SetProblemCategoryNil

`func (o *CellLinesRow) SetProblemCategoryNil(b bool)`

 SetProblemCategoryNil sets the value for ProblemCategory to be an explicit nil

### UnsetProblemCategory
`func (o *CellLinesRow) UnsetProblemCategory()`

UnsetProblemCategory ensures that no value is present for ProblemCategory, not even an explicit nil
### GetProblematicComment

`func (o *CellLinesRow) GetProblematicComment() string`

GetProblematicComment returns the ProblematicComment field if non-nil, zero value otherwise.

### GetProblematicCommentOk

`func (o *CellLinesRow) GetProblematicCommentOk() (*string, bool)`

GetProblematicCommentOk returns a tuple with the ProblematicComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblematicComment

`func (o *CellLinesRow) SetProblematicComment(v string)`

SetProblematicComment sets ProblematicComment field to given value.

### HasProblematicComment

`func (o *CellLinesRow) HasProblematicComment() bool`

HasProblematicComment returns a boolean if a field has been set.

### SetProblematicCommentNil

`func (o *CellLinesRow) SetProblematicCommentNil(b bool)`

 SetProblematicCommentNil sets the value for ProblematicComment to be an explicit nil

### UnsetProblematicComment
`func (o *CellLinesRow) UnsetProblematicComment()`

UnsetProblematicComment ensures that no value is present for ProblematicComment, not even an explicit nil
### GetIclacRegistrationIds

`func (o *CellLinesRow) GetIclacRegistrationIds() []string`

GetIclacRegistrationIds returns the IclacRegistrationIds field if non-nil, zero value otherwise.

### GetIclacRegistrationIdsOk

`func (o *CellLinesRow) GetIclacRegistrationIdsOk() (*[]string, bool)`

GetIclacRegistrationIdsOk returns a tuple with the IclacRegistrationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIclacRegistrationIds

`func (o *CellLinesRow) SetIclacRegistrationIds(v []string)`

SetIclacRegistrationIds sets IclacRegistrationIds field to given value.

### HasIclacRegistrationIds

`func (o *CellLinesRow) HasIclacRegistrationIds() bool`

HasIclacRegistrationIds returns a boolean if a field has been set.

### SetIclacRegistrationIdsNil

`func (o *CellLinesRow) SetIclacRegistrationIdsNil(b bool)`

 SetIclacRegistrationIdsNil sets the value for IclacRegistrationIds to be an explicit nil

### UnsetIclacRegistrationIds
`func (o *CellLinesRow) UnsetIclacRegistrationIds()`

UnsetIclacRegistrationIds ensures that no value is present for IclacRegistrationIds, not even an explicit nil
### GetReleaseVersion

`func (o *CellLinesRow) GetReleaseVersion() string`

GetReleaseVersion returns the ReleaseVersion field if non-nil, zero value otherwise.

### GetReleaseVersionOk

`func (o *CellLinesRow) GetReleaseVersionOk() (*string, bool)`

GetReleaseVersionOk returns a tuple with the ReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersion

`func (o *CellLinesRow) SetReleaseVersion(v string)`

SetReleaseVersion sets ReleaseVersion field to given value.

### HasReleaseVersion

`func (o *CellLinesRow) HasReleaseVersion() bool`

HasReleaseVersion returns a boolean if a field has been set.

### SetReleaseVersionNil

`func (o *CellLinesRow) SetReleaseVersionNil(b bool)`

 SetReleaseVersionNil sets the value for ReleaseVersion to be an explicit nil

### UnsetReleaseVersion
`func (o *CellLinesRow) UnsetReleaseVersion()`

UnsetReleaseVersion ensures that no value is present for ReleaseVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


