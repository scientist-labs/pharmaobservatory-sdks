# RareDiseaseGenesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneAssertionId** | Pointer to **NullableString** | The identifier of the disease-gene assertion. | [optional] 
**OrphaCode** | Pointer to **NullableString** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] 
**GeneSymbol** | Pointer to **NullableString** | The gene symbol as Orphadata gives it. | [optional] 
**GeneName** | Pointer to **NullableString** | The gene&#39;s full name. | [optional] 
**GeneType** | Pointer to **NullableString** | The gene type Orphadata records (for example gene_with_protein_product). | [optional] 
**AssociationType** | Pointer to **NullableString** | The association Orphadata asserts between the gene and the disease, for example a disease-causing germline mutation or a modifying germline mutation. | [optional] 
**AssociationStatus** | Pointer to **NullableString** | assessed when Orphanet has assessed the association, not_yet_assessed otherwise. | [optional] 
**IsKnownCausal** | Pointer to **NullableBool** | True for an assessed disease-causing mutation assertion. | [optional] 
**EnsemblId** | Pointer to **NullableString** | The Ensembl gene id Orphadata gives for the gene. | [optional] 
**NcbiGeneId** | Pointer to **NullableString** | The NCBI Gene id Orphadata gives for the gene. | [optional] 
**TargetId** | Pointer to **NullableString** | The Ensembl gene id of the matching row in the targets dataset; null when the gene has no row there. | [optional] 
**GeneResolutionStatus** | Pointer to **NullableString** | How the gene&#39;s Ensembl and NCBI identifiers resolved against the gene spine. | [optional] 
**GeneIdentifierConflict** | Pointer to **NullableBool** | True when the identifiers on the assertion disagree with the gene spine. | [optional] 
**SourceRelease** | Pointer to **NullableString** | The Orphadata Science release the row was read from. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The URL of the source file the row was read from. | [optional] 
**LicenseName** | Pointer to **NullableString** | The licence of the Orphadata Science file. | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRareDiseaseGenesRow

`func NewRareDiseaseGenesRow() *RareDiseaseGenesRow`

NewRareDiseaseGenesRow instantiates a new RareDiseaseGenesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRareDiseaseGenesRowWithDefaults

`func NewRareDiseaseGenesRowWithDefaults() *RareDiseaseGenesRow`

NewRareDiseaseGenesRowWithDefaults instantiates a new RareDiseaseGenesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneAssertionId

`func (o *RareDiseaseGenesRow) GetGeneAssertionId() string`

GetGeneAssertionId returns the GeneAssertionId field if non-nil, zero value otherwise.

### GetGeneAssertionIdOk

`func (o *RareDiseaseGenesRow) GetGeneAssertionIdOk() (*string, bool)`

GetGeneAssertionIdOk returns a tuple with the GeneAssertionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneAssertionId

`func (o *RareDiseaseGenesRow) SetGeneAssertionId(v string)`

SetGeneAssertionId sets GeneAssertionId field to given value.

### HasGeneAssertionId

`func (o *RareDiseaseGenesRow) HasGeneAssertionId() bool`

HasGeneAssertionId returns a boolean if a field has been set.

### SetGeneAssertionIdNil

`func (o *RareDiseaseGenesRow) SetGeneAssertionIdNil(b bool)`

 SetGeneAssertionIdNil sets the value for GeneAssertionId to be an explicit nil

### UnsetGeneAssertionId
`func (o *RareDiseaseGenesRow) UnsetGeneAssertionId()`

UnsetGeneAssertionId ensures that no value is present for GeneAssertionId, not even an explicit nil
### GetOrphaCode

`func (o *RareDiseaseGenesRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *RareDiseaseGenesRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *RareDiseaseGenesRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *RareDiseaseGenesRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *RareDiseaseGenesRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *RareDiseaseGenesRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetGeneSymbol

`func (o *RareDiseaseGenesRow) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *RareDiseaseGenesRow) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *RareDiseaseGenesRow) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *RareDiseaseGenesRow) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbolNil

`func (o *RareDiseaseGenesRow) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *RareDiseaseGenesRow) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetGeneName

`func (o *RareDiseaseGenesRow) GetGeneName() string`

GetGeneName returns the GeneName field if non-nil, zero value otherwise.

### GetGeneNameOk

`func (o *RareDiseaseGenesRow) GetGeneNameOk() (*string, bool)`

GetGeneNameOk returns a tuple with the GeneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneName

`func (o *RareDiseaseGenesRow) SetGeneName(v string)`

SetGeneName sets GeneName field to given value.

### HasGeneName

`func (o *RareDiseaseGenesRow) HasGeneName() bool`

HasGeneName returns a boolean if a field has been set.

### SetGeneNameNil

`func (o *RareDiseaseGenesRow) SetGeneNameNil(b bool)`

 SetGeneNameNil sets the value for GeneName to be an explicit nil

### UnsetGeneName
`func (o *RareDiseaseGenesRow) UnsetGeneName()`

UnsetGeneName ensures that no value is present for GeneName, not even an explicit nil
### GetGeneType

`func (o *RareDiseaseGenesRow) GetGeneType() string`

GetGeneType returns the GeneType field if non-nil, zero value otherwise.

### GetGeneTypeOk

`func (o *RareDiseaseGenesRow) GetGeneTypeOk() (*string, bool)`

GetGeneTypeOk returns a tuple with the GeneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneType

`func (o *RareDiseaseGenesRow) SetGeneType(v string)`

SetGeneType sets GeneType field to given value.

### HasGeneType

`func (o *RareDiseaseGenesRow) HasGeneType() bool`

HasGeneType returns a boolean if a field has been set.

### SetGeneTypeNil

`func (o *RareDiseaseGenesRow) SetGeneTypeNil(b bool)`

 SetGeneTypeNil sets the value for GeneType to be an explicit nil

### UnsetGeneType
`func (o *RareDiseaseGenesRow) UnsetGeneType()`

UnsetGeneType ensures that no value is present for GeneType, not even an explicit nil
### GetAssociationType

`func (o *RareDiseaseGenesRow) GetAssociationType() string`

GetAssociationType returns the AssociationType field if non-nil, zero value otherwise.

### GetAssociationTypeOk

`func (o *RareDiseaseGenesRow) GetAssociationTypeOk() (*string, bool)`

GetAssociationTypeOk returns a tuple with the AssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationType

`func (o *RareDiseaseGenesRow) SetAssociationType(v string)`

SetAssociationType sets AssociationType field to given value.

### HasAssociationType

`func (o *RareDiseaseGenesRow) HasAssociationType() bool`

HasAssociationType returns a boolean if a field has been set.

### SetAssociationTypeNil

`func (o *RareDiseaseGenesRow) SetAssociationTypeNil(b bool)`

 SetAssociationTypeNil sets the value for AssociationType to be an explicit nil

### UnsetAssociationType
`func (o *RareDiseaseGenesRow) UnsetAssociationType()`

UnsetAssociationType ensures that no value is present for AssociationType, not even an explicit nil
### GetAssociationStatus

`func (o *RareDiseaseGenesRow) GetAssociationStatus() string`

GetAssociationStatus returns the AssociationStatus field if non-nil, zero value otherwise.

### GetAssociationStatusOk

`func (o *RareDiseaseGenesRow) GetAssociationStatusOk() (*string, bool)`

GetAssociationStatusOk returns a tuple with the AssociationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationStatus

`func (o *RareDiseaseGenesRow) SetAssociationStatus(v string)`

SetAssociationStatus sets AssociationStatus field to given value.

### HasAssociationStatus

`func (o *RareDiseaseGenesRow) HasAssociationStatus() bool`

HasAssociationStatus returns a boolean if a field has been set.

### SetAssociationStatusNil

`func (o *RareDiseaseGenesRow) SetAssociationStatusNil(b bool)`

 SetAssociationStatusNil sets the value for AssociationStatus to be an explicit nil

### UnsetAssociationStatus
`func (o *RareDiseaseGenesRow) UnsetAssociationStatus()`

UnsetAssociationStatus ensures that no value is present for AssociationStatus, not even an explicit nil
### GetIsKnownCausal

`func (o *RareDiseaseGenesRow) GetIsKnownCausal() bool`

GetIsKnownCausal returns the IsKnownCausal field if non-nil, zero value otherwise.

### GetIsKnownCausalOk

`func (o *RareDiseaseGenesRow) GetIsKnownCausalOk() (*bool, bool)`

GetIsKnownCausalOk returns a tuple with the IsKnownCausal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKnownCausal

`func (o *RareDiseaseGenesRow) SetIsKnownCausal(v bool)`

SetIsKnownCausal sets IsKnownCausal field to given value.

### HasIsKnownCausal

`func (o *RareDiseaseGenesRow) HasIsKnownCausal() bool`

HasIsKnownCausal returns a boolean if a field has been set.

### SetIsKnownCausalNil

`func (o *RareDiseaseGenesRow) SetIsKnownCausalNil(b bool)`

 SetIsKnownCausalNil sets the value for IsKnownCausal to be an explicit nil

### UnsetIsKnownCausal
`func (o *RareDiseaseGenesRow) UnsetIsKnownCausal()`

UnsetIsKnownCausal ensures that no value is present for IsKnownCausal, not even an explicit nil
### GetEnsemblId

`func (o *RareDiseaseGenesRow) GetEnsemblId() string`

GetEnsemblId returns the EnsemblId field if non-nil, zero value otherwise.

### GetEnsemblIdOk

`func (o *RareDiseaseGenesRow) GetEnsemblIdOk() (*string, bool)`

GetEnsemblIdOk returns a tuple with the EnsemblId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsemblId

`func (o *RareDiseaseGenesRow) SetEnsemblId(v string)`

SetEnsemblId sets EnsemblId field to given value.

### HasEnsemblId

`func (o *RareDiseaseGenesRow) HasEnsemblId() bool`

HasEnsemblId returns a boolean if a field has been set.

### SetEnsemblIdNil

`func (o *RareDiseaseGenesRow) SetEnsemblIdNil(b bool)`

 SetEnsemblIdNil sets the value for EnsemblId to be an explicit nil

### UnsetEnsemblId
`func (o *RareDiseaseGenesRow) UnsetEnsemblId()`

UnsetEnsemblId ensures that no value is present for EnsemblId, not even an explicit nil
### GetNcbiGeneId

`func (o *RareDiseaseGenesRow) GetNcbiGeneId() string`

GetNcbiGeneId returns the NcbiGeneId field if non-nil, zero value otherwise.

### GetNcbiGeneIdOk

`func (o *RareDiseaseGenesRow) GetNcbiGeneIdOk() (*string, bool)`

GetNcbiGeneIdOk returns a tuple with the NcbiGeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcbiGeneId

`func (o *RareDiseaseGenesRow) SetNcbiGeneId(v string)`

SetNcbiGeneId sets NcbiGeneId field to given value.

### HasNcbiGeneId

`func (o *RareDiseaseGenesRow) HasNcbiGeneId() bool`

HasNcbiGeneId returns a boolean if a field has been set.

### SetNcbiGeneIdNil

`func (o *RareDiseaseGenesRow) SetNcbiGeneIdNil(b bool)`

 SetNcbiGeneIdNil sets the value for NcbiGeneId to be an explicit nil

### UnsetNcbiGeneId
`func (o *RareDiseaseGenesRow) UnsetNcbiGeneId()`

UnsetNcbiGeneId ensures that no value is present for NcbiGeneId, not even an explicit nil
### GetTargetId

`func (o *RareDiseaseGenesRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *RareDiseaseGenesRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *RareDiseaseGenesRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *RareDiseaseGenesRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *RareDiseaseGenesRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *RareDiseaseGenesRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetGeneResolutionStatus

`func (o *RareDiseaseGenesRow) GetGeneResolutionStatus() string`

GetGeneResolutionStatus returns the GeneResolutionStatus field if non-nil, zero value otherwise.

### GetGeneResolutionStatusOk

`func (o *RareDiseaseGenesRow) GetGeneResolutionStatusOk() (*string, bool)`

GetGeneResolutionStatusOk returns a tuple with the GeneResolutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneResolutionStatus

`func (o *RareDiseaseGenesRow) SetGeneResolutionStatus(v string)`

SetGeneResolutionStatus sets GeneResolutionStatus field to given value.

### HasGeneResolutionStatus

`func (o *RareDiseaseGenesRow) HasGeneResolutionStatus() bool`

HasGeneResolutionStatus returns a boolean if a field has been set.

### SetGeneResolutionStatusNil

`func (o *RareDiseaseGenesRow) SetGeneResolutionStatusNil(b bool)`

 SetGeneResolutionStatusNil sets the value for GeneResolutionStatus to be an explicit nil

### UnsetGeneResolutionStatus
`func (o *RareDiseaseGenesRow) UnsetGeneResolutionStatus()`

UnsetGeneResolutionStatus ensures that no value is present for GeneResolutionStatus, not even an explicit nil
### GetGeneIdentifierConflict

`func (o *RareDiseaseGenesRow) GetGeneIdentifierConflict() bool`

GetGeneIdentifierConflict returns the GeneIdentifierConflict field if non-nil, zero value otherwise.

### GetGeneIdentifierConflictOk

`func (o *RareDiseaseGenesRow) GetGeneIdentifierConflictOk() (*bool, bool)`

GetGeneIdentifierConflictOk returns a tuple with the GeneIdentifierConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneIdentifierConflict

`func (o *RareDiseaseGenesRow) SetGeneIdentifierConflict(v bool)`

SetGeneIdentifierConflict sets GeneIdentifierConflict field to given value.

### HasGeneIdentifierConflict

`func (o *RareDiseaseGenesRow) HasGeneIdentifierConflict() bool`

HasGeneIdentifierConflict returns a boolean if a field has been set.

### SetGeneIdentifierConflictNil

`func (o *RareDiseaseGenesRow) SetGeneIdentifierConflictNil(b bool)`

 SetGeneIdentifierConflictNil sets the value for GeneIdentifierConflict to be an explicit nil

### UnsetGeneIdentifierConflict
`func (o *RareDiseaseGenesRow) UnsetGeneIdentifierConflict()`

UnsetGeneIdentifierConflict ensures that no value is present for GeneIdentifierConflict, not even an explicit nil
### GetSourceRelease

`func (o *RareDiseaseGenesRow) GetSourceRelease() string`

GetSourceRelease returns the SourceRelease field if non-nil, zero value otherwise.

### GetSourceReleaseOk

`func (o *RareDiseaseGenesRow) GetSourceReleaseOk() (*string, bool)`

GetSourceReleaseOk returns a tuple with the SourceRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRelease

`func (o *RareDiseaseGenesRow) SetSourceRelease(v string)`

SetSourceRelease sets SourceRelease field to given value.

### HasSourceRelease

`func (o *RareDiseaseGenesRow) HasSourceRelease() bool`

HasSourceRelease returns a boolean if a field has been set.

### SetSourceReleaseNil

`func (o *RareDiseaseGenesRow) SetSourceReleaseNil(b bool)`

 SetSourceReleaseNil sets the value for SourceRelease to be an explicit nil

### UnsetSourceRelease
`func (o *RareDiseaseGenesRow) UnsetSourceRelease()`

UnsetSourceRelease ensures that no value is present for SourceRelease, not even an explicit nil
### GetSourceUrl

`func (o *RareDiseaseGenesRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *RareDiseaseGenesRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *RareDiseaseGenesRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *RareDiseaseGenesRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *RareDiseaseGenesRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *RareDiseaseGenesRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetLicenseName

`func (o *RareDiseaseGenesRow) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *RareDiseaseGenesRow) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *RareDiseaseGenesRow) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *RareDiseaseGenesRow) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### SetLicenseNameNil

`func (o *RareDiseaseGenesRow) SetLicenseNameNil(b bool)`

 SetLicenseNameNil sets the value for LicenseName to be an explicit nil

### UnsetLicenseName
`func (o *RareDiseaseGenesRow) UnsetLicenseName()`

UnsetLicenseName ensures that no value is present for LicenseName, not even an explicit nil
### GetAttribution

`func (o *RareDiseaseGenesRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RareDiseaseGenesRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RareDiseaseGenesRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RareDiseaseGenesRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RareDiseaseGenesRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RareDiseaseGenesRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


