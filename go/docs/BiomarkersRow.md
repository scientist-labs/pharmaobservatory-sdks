# BiomarkersRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BiomarkerId** | Pointer to **NullableString** | The biomarker key. | [optional] 
**RouteKey** | Pointer to **NullableString** | The biomarker&#39;s route key, the id the web page uses and the one entry key every biomarker dataset takes (the gene symbol for a gene biomarker). | [optional] 
**BiomarkerKind** | Pointer to **NullableString** | gene, protein_variant or another feature kind. | [optional] 
**IdentityStatus** | Pointer to **NullableString** | Whether the biomarker&#39;s identity is confirmed. | [optional] 
**DisplayName** | Pointer to **NullableString** | The biomarker&#39;s display name. | [optional] 
**ParentBiomarkerId** | Pointer to **NullableString** | The parent biomarker (a variant&#39;s gene). | [optional] 
**GeneSymbol** | Pointer to **NullableString** | The gene symbol. | [optional] 
**AminoAcidChange** | Pointer to **NullableString** | The amino-acid change, for a variant. | [optional] 
**TargetId** | Pointer to **NullableString** | The target&#39;s Ensembl gene id. | [optional] 
**HgncId** | Pointer to **NullableString** | The HUGO Gene Nomenclature Committee (HGNC) gene id. | [optional] 
**UniprotAccession** | Pointer to **NullableString** | The UniProt accession. | [optional] 
**TargetProteinName** | Pointer to **NullableString** | The recommended name of the protein the biomarker&#39;s gene encodes: the UniProt protein of the accession, else the first UniProt protein of the gene symbol, else the Open Targets target name. Null when none names it. | [optional] 
**FeatureCode** | Pointer to **NullableString** | Feature code for a non-gene feature. | [optional] 
**FeatureType** | Pointer to **NullableString** | Feature type for a non-gene feature. | [optional] 
**EvidenceTier** | Pointer to **NullableInt32** | Strongest evidence tier: 1 companion diagnostic of the US Food and Drug Administration (FDA) through 6 no evidence. | [optional] 
**EvidenceTierLabel** | Pointer to **NullableString** | The tier&#39;s name. | [optional] 
**NFdaCdxAssertions** | Pointer to **NullableInt32** | FDA companion diagnostic assertions. | [optional] 
**NFdaBqpAssertions** | Pointer to **NullableInt32** | FDA biomarker qualification project assertions. | [optional] 
**NCancerActionableCgi** | Pointer to **NullableInt32** | Cancer Genome Interpreter actionable rows. | [optional] 
**NOtAssociation** | Pointer to **NullableInt32** | Open Targets target-disease associations. | [optional] 
**NTrialMention** | Pointer to **NullableInt32** | Trials that name the biomarker. | [optional] 
**NPublicationVariantMentions** | Pointer to **NullableInt32** | Publication variant mentions. | [optional] 
**HasClinicalValidationEvidence** | Pointer to **NullableBool** | Whether clinical validation evidence exists. | [optional] 
**HasAnalyticalValidationEvidence** | Pointer to **NullableBool** | Whether analytical validation evidence exists. | [optional] 
**NAssayProviderPapers** | Pointer to **NullableInt32** | Papers naming an assay provider for the biomarker. | [optional] 
**NAssayProviderSuppliers** | Pointer to **NullableInt32** | Assay-provider suppliers named. | [optional] 
**NFdaEvidenceProviders** | Pointer to **NullableInt32** | Of those, providers with an FDA record found for the biomarker (a companion diagnostic, a cleared or approved device the paper names, a cleared or approved IVD under the analyte&#39;s product code, or a CLIA-certified lab that ran the test). US only; no record found is not a statement about validation. Null when no assay provider is named for the biomarker. | [optional] 

## Methods

### NewBiomarkersRow

`func NewBiomarkersRow() *BiomarkersRow`

NewBiomarkersRow instantiates a new BiomarkersRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkersRowWithDefaults

`func NewBiomarkersRowWithDefaults() *BiomarkersRow`

NewBiomarkersRowWithDefaults instantiates a new BiomarkersRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBiomarkerId

`func (o *BiomarkersRow) GetBiomarkerId() string`

GetBiomarkerId returns the BiomarkerId field if non-nil, zero value otherwise.

### GetBiomarkerIdOk

`func (o *BiomarkersRow) GetBiomarkerIdOk() (*string, bool)`

GetBiomarkerIdOk returns a tuple with the BiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerId

`func (o *BiomarkersRow) SetBiomarkerId(v string)`

SetBiomarkerId sets BiomarkerId field to given value.

### HasBiomarkerId

`func (o *BiomarkersRow) HasBiomarkerId() bool`

HasBiomarkerId returns a boolean if a field has been set.

### SetBiomarkerIdNil

`func (o *BiomarkersRow) SetBiomarkerIdNil(b bool)`

 SetBiomarkerIdNil sets the value for BiomarkerId to be an explicit nil

### UnsetBiomarkerId
`func (o *BiomarkersRow) UnsetBiomarkerId()`

UnsetBiomarkerId ensures that no value is present for BiomarkerId, not even an explicit nil
### GetRouteKey

`func (o *BiomarkersRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkersRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkersRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkersRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkersRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkersRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetBiomarkerKind

`func (o *BiomarkersRow) GetBiomarkerKind() string`

GetBiomarkerKind returns the BiomarkerKind field if non-nil, zero value otherwise.

### GetBiomarkerKindOk

`func (o *BiomarkersRow) GetBiomarkerKindOk() (*string, bool)`

GetBiomarkerKindOk returns a tuple with the BiomarkerKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerKind

`func (o *BiomarkersRow) SetBiomarkerKind(v string)`

SetBiomarkerKind sets BiomarkerKind field to given value.

### HasBiomarkerKind

`func (o *BiomarkersRow) HasBiomarkerKind() bool`

HasBiomarkerKind returns a boolean if a field has been set.

### SetBiomarkerKindNil

`func (o *BiomarkersRow) SetBiomarkerKindNil(b bool)`

 SetBiomarkerKindNil sets the value for BiomarkerKind to be an explicit nil

### UnsetBiomarkerKind
`func (o *BiomarkersRow) UnsetBiomarkerKind()`

UnsetBiomarkerKind ensures that no value is present for BiomarkerKind, not even an explicit nil
### GetIdentityStatus

`func (o *BiomarkersRow) GetIdentityStatus() string`

GetIdentityStatus returns the IdentityStatus field if non-nil, zero value otherwise.

### GetIdentityStatusOk

`func (o *BiomarkersRow) GetIdentityStatusOk() (*string, bool)`

GetIdentityStatusOk returns a tuple with the IdentityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStatus

`func (o *BiomarkersRow) SetIdentityStatus(v string)`

SetIdentityStatus sets IdentityStatus field to given value.

### HasIdentityStatus

`func (o *BiomarkersRow) HasIdentityStatus() bool`

HasIdentityStatus returns a boolean if a field has been set.

### SetIdentityStatusNil

`func (o *BiomarkersRow) SetIdentityStatusNil(b bool)`

 SetIdentityStatusNil sets the value for IdentityStatus to be an explicit nil

### UnsetIdentityStatus
`func (o *BiomarkersRow) UnsetIdentityStatus()`

UnsetIdentityStatus ensures that no value is present for IdentityStatus, not even an explicit nil
### GetDisplayName

`func (o *BiomarkersRow) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BiomarkersRow) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BiomarkersRow) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BiomarkersRow) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *BiomarkersRow) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *BiomarkersRow) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetParentBiomarkerId

`func (o *BiomarkersRow) GetParentBiomarkerId() string`

GetParentBiomarkerId returns the ParentBiomarkerId field if non-nil, zero value otherwise.

### GetParentBiomarkerIdOk

`func (o *BiomarkersRow) GetParentBiomarkerIdOk() (*string, bool)`

GetParentBiomarkerIdOk returns a tuple with the ParentBiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBiomarkerId

`func (o *BiomarkersRow) SetParentBiomarkerId(v string)`

SetParentBiomarkerId sets ParentBiomarkerId field to given value.

### HasParentBiomarkerId

`func (o *BiomarkersRow) HasParentBiomarkerId() bool`

HasParentBiomarkerId returns a boolean if a field has been set.

### SetParentBiomarkerIdNil

`func (o *BiomarkersRow) SetParentBiomarkerIdNil(b bool)`

 SetParentBiomarkerIdNil sets the value for ParentBiomarkerId to be an explicit nil

### UnsetParentBiomarkerId
`func (o *BiomarkersRow) UnsetParentBiomarkerId()`

UnsetParentBiomarkerId ensures that no value is present for ParentBiomarkerId, not even an explicit nil
### GetGeneSymbol

`func (o *BiomarkersRow) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *BiomarkersRow) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *BiomarkersRow) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *BiomarkersRow) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbolNil

`func (o *BiomarkersRow) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *BiomarkersRow) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetAminoAcidChange

`func (o *BiomarkersRow) GetAminoAcidChange() string`

GetAminoAcidChange returns the AminoAcidChange field if non-nil, zero value otherwise.

### GetAminoAcidChangeOk

`func (o *BiomarkersRow) GetAminoAcidChangeOk() (*string, bool)`

GetAminoAcidChangeOk returns a tuple with the AminoAcidChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcidChange

`func (o *BiomarkersRow) SetAminoAcidChange(v string)`

SetAminoAcidChange sets AminoAcidChange field to given value.

### HasAminoAcidChange

`func (o *BiomarkersRow) HasAminoAcidChange() bool`

HasAminoAcidChange returns a boolean if a field has been set.

### SetAminoAcidChangeNil

`func (o *BiomarkersRow) SetAminoAcidChangeNil(b bool)`

 SetAminoAcidChangeNil sets the value for AminoAcidChange to be an explicit nil

### UnsetAminoAcidChange
`func (o *BiomarkersRow) UnsetAminoAcidChange()`

UnsetAminoAcidChange ensures that no value is present for AminoAcidChange, not even an explicit nil
### GetTargetId

`func (o *BiomarkersRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *BiomarkersRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *BiomarkersRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *BiomarkersRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *BiomarkersRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *BiomarkersRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetHgncId

`func (o *BiomarkersRow) GetHgncId() string`

GetHgncId returns the HgncId field if non-nil, zero value otherwise.

### GetHgncIdOk

`func (o *BiomarkersRow) GetHgncIdOk() (*string, bool)`

GetHgncIdOk returns a tuple with the HgncId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgncId

`func (o *BiomarkersRow) SetHgncId(v string)`

SetHgncId sets HgncId field to given value.

### HasHgncId

`func (o *BiomarkersRow) HasHgncId() bool`

HasHgncId returns a boolean if a field has been set.

### SetHgncIdNil

`func (o *BiomarkersRow) SetHgncIdNil(b bool)`

 SetHgncIdNil sets the value for HgncId to be an explicit nil

### UnsetHgncId
`func (o *BiomarkersRow) UnsetHgncId()`

UnsetHgncId ensures that no value is present for HgncId, not even an explicit nil
### GetUniprotAccession

`func (o *BiomarkersRow) GetUniprotAccession() string`

GetUniprotAccession returns the UniprotAccession field if non-nil, zero value otherwise.

### GetUniprotAccessionOk

`func (o *BiomarkersRow) GetUniprotAccessionOk() (*string, bool)`

GetUniprotAccessionOk returns a tuple with the UniprotAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniprotAccession

`func (o *BiomarkersRow) SetUniprotAccession(v string)`

SetUniprotAccession sets UniprotAccession field to given value.

### HasUniprotAccession

`func (o *BiomarkersRow) HasUniprotAccession() bool`

HasUniprotAccession returns a boolean if a field has been set.

### SetUniprotAccessionNil

`func (o *BiomarkersRow) SetUniprotAccessionNil(b bool)`

 SetUniprotAccessionNil sets the value for UniprotAccession to be an explicit nil

### UnsetUniprotAccession
`func (o *BiomarkersRow) UnsetUniprotAccession()`

UnsetUniprotAccession ensures that no value is present for UniprotAccession, not even an explicit nil
### GetTargetProteinName

`func (o *BiomarkersRow) GetTargetProteinName() string`

GetTargetProteinName returns the TargetProteinName field if non-nil, zero value otherwise.

### GetTargetProteinNameOk

`func (o *BiomarkersRow) GetTargetProteinNameOk() (*string, bool)`

GetTargetProteinNameOk returns a tuple with the TargetProteinName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProteinName

`func (o *BiomarkersRow) SetTargetProteinName(v string)`

SetTargetProteinName sets TargetProteinName field to given value.

### HasTargetProteinName

`func (o *BiomarkersRow) HasTargetProteinName() bool`

HasTargetProteinName returns a boolean if a field has been set.

### SetTargetProteinNameNil

`func (o *BiomarkersRow) SetTargetProteinNameNil(b bool)`

 SetTargetProteinNameNil sets the value for TargetProteinName to be an explicit nil

### UnsetTargetProteinName
`func (o *BiomarkersRow) UnsetTargetProteinName()`

UnsetTargetProteinName ensures that no value is present for TargetProteinName, not even an explicit nil
### GetFeatureCode

`func (o *BiomarkersRow) GetFeatureCode() string`

GetFeatureCode returns the FeatureCode field if non-nil, zero value otherwise.

### GetFeatureCodeOk

`func (o *BiomarkersRow) GetFeatureCodeOk() (*string, bool)`

GetFeatureCodeOk returns a tuple with the FeatureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureCode

`func (o *BiomarkersRow) SetFeatureCode(v string)`

SetFeatureCode sets FeatureCode field to given value.

### HasFeatureCode

`func (o *BiomarkersRow) HasFeatureCode() bool`

HasFeatureCode returns a boolean if a field has been set.

### SetFeatureCodeNil

`func (o *BiomarkersRow) SetFeatureCodeNil(b bool)`

 SetFeatureCodeNil sets the value for FeatureCode to be an explicit nil

### UnsetFeatureCode
`func (o *BiomarkersRow) UnsetFeatureCode()`

UnsetFeatureCode ensures that no value is present for FeatureCode, not even an explicit nil
### GetFeatureType

`func (o *BiomarkersRow) GetFeatureType() string`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *BiomarkersRow) GetFeatureTypeOk() (*string, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *BiomarkersRow) SetFeatureType(v string)`

SetFeatureType sets FeatureType field to given value.

### HasFeatureType

`func (o *BiomarkersRow) HasFeatureType() bool`

HasFeatureType returns a boolean if a field has been set.

### SetFeatureTypeNil

`func (o *BiomarkersRow) SetFeatureTypeNil(b bool)`

 SetFeatureTypeNil sets the value for FeatureType to be an explicit nil

### UnsetFeatureType
`func (o *BiomarkersRow) UnsetFeatureType()`

UnsetFeatureType ensures that no value is present for FeatureType, not even an explicit nil
### GetEvidenceTier

`func (o *BiomarkersRow) GetEvidenceTier() int32`

GetEvidenceTier returns the EvidenceTier field if non-nil, zero value otherwise.

### GetEvidenceTierOk

`func (o *BiomarkersRow) GetEvidenceTierOk() (*int32, bool)`

GetEvidenceTierOk returns a tuple with the EvidenceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceTier

`func (o *BiomarkersRow) SetEvidenceTier(v int32)`

SetEvidenceTier sets EvidenceTier field to given value.

### HasEvidenceTier

`func (o *BiomarkersRow) HasEvidenceTier() bool`

HasEvidenceTier returns a boolean if a field has been set.

### SetEvidenceTierNil

`func (o *BiomarkersRow) SetEvidenceTierNil(b bool)`

 SetEvidenceTierNil sets the value for EvidenceTier to be an explicit nil

### UnsetEvidenceTier
`func (o *BiomarkersRow) UnsetEvidenceTier()`

UnsetEvidenceTier ensures that no value is present for EvidenceTier, not even an explicit nil
### GetEvidenceTierLabel

`func (o *BiomarkersRow) GetEvidenceTierLabel() string`

GetEvidenceTierLabel returns the EvidenceTierLabel field if non-nil, zero value otherwise.

### GetEvidenceTierLabelOk

`func (o *BiomarkersRow) GetEvidenceTierLabelOk() (*string, bool)`

GetEvidenceTierLabelOk returns a tuple with the EvidenceTierLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceTierLabel

`func (o *BiomarkersRow) SetEvidenceTierLabel(v string)`

SetEvidenceTierLabel sets EvidenceTierLabel field to given value.

### HasEvidenceTierLabel

`func (o *BiomarkersRow) HasEvidenceTierLabel() bool`

HasEvidenceTierLabel returns a boolean if a field has been set.

### SetEvidenceTierLabelNil

`func (o *BiomarkersRow) SetEvidenceTierLabelNil(b bool)`

 SetEvidenceTierLabelNil sets the value for EvidenceTierLabel to be an explicit nil

### UnsetEvidenceTierLabel
`func (o *BiomarkersRow) UnsetEvidenceTierLabel()`

UnsetEvidenceTierLabel ensures that no value is present for EvidenceTierLabel, not even an explicit nil
### GetNFdaCdxAssertions

`func (o *BiomarkersRow) GetNFdaCdxAssertions() int32`

GetNFdaCdxAssertions returns the NFdaCdxAssertions field if non-nil, zero value otherwise.

### GetNFdaCdxAssertionsOk

`func (o *BiomarkersRow) GetNFdaCdxAssertionsOk() (*int32, bool)`

GetNFdaCdxAssertionsOk returns a tuple with the NFdaCdxAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFdaCdxAssertions

`func (o *BiomarkersRow) SetNFdaCdxAssertions(v int32)`

SetNFdaCdxAssertions sets NFdaCdxAssertions field to given value.

### HasNFdaCdxAssertions

`func (o *BiomarkersRow) HasNFdaCdxAssertions() bool`

HasNFdaCdxAssertions returns a boolean if a field has been set.

### SetNFdaCdxAssertionsNil

`func (o *BiomarkersRow) SetNFdaCdxAssertionsNil(b bool)`

 SetNFdaCdxAssertionsNil sets the value for NFdaCdxAssertions to be an explicit nil

### UnsetNFdaCdxAssertions
`func (o *BiomarkersRow) UnsetNFdaCdxAssertions()`

UnsetNFdaCdxAssertions ensures that no value is present for NFdaCdxAssertions, not even an explicit nil
### GetNFdaBqpAssertions

`func (o *BiomarkersRow) GetNFdaBqpAssertions() int32`

GetNFdaBqpAssertions returns the NFdaBqpAssertions field if non-nil, zero value otherwise.

### GetNFdaBqpAssertionsOk

`func (o *BiomarkersRow) GetNFdaBqpAssertionsOk() (*int32, bool)`

GetNFdaBqpAssertionsOk returns a tuple with the NFdaBqpAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFdaBqpAssertions

`func (o *BiomarkersRow) SetNFdaBqpAssertions(v int32)`

SetNFdaBqpAssertions sets NFdaBqpAssertions field to given value.

### HasNFdaBqpAssertions

`func (o *BiomarkersRow) HasNFdaBqpAssertions() bool`

HasNFdaBqpAssertions returns a boolean if a field has been set.

### SetNFdaBqpAssertionsNil

`func (o *BiomarkersRow) SetNFdaBqpAssertionsNil(b bool)`

 SetNFdaBqpAssertionsNil sets the value for NFdaBqpAssertions to be an explicit nil

### UnsetNFdaBqpAssertions
`func (o *BiomarkersRow) UnsetNFdaBqpAssertions()`

UnsetNFdaBqpAssertions ensures that no value is present for NFdaBqpAssertions, not even an explicit nil
### GetNCancerActionableCgi

`func (o *BiomarkersRow) GetNCancerActionableCgi() int32`

GetNCancerActionableCgi returns the NCancerActionableCgi field if non-nil, zero value otherwise.

### GetNCancerActionableCgiOk

`func (o *BiomarkersRow) GetNCancerActionableCgiOk() (*int32, bool)`

GetNCancerActionableCgiOk returns a tuple with the NCancerActionableCgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCancerActionableCgi

`func (o *BiomarkersRow) SetNCancerActionableCgi(v int32)`

SetNCancerActionableCgi sets NCancerActionableCgi field to given value.

### HasNCancerActionableCgi

`func (o *BiomarkersRow) HasNCancerActionableCgi() bool`

HasNCancerActionableCgi returns a boolean if a field has been set.

### SetNCancerActionableCgiNil

`func (o *BiomarkersRow) SetNCancerActionableCgiNil(b bool)`

 SetNCancerActionableCgiNil sets the value for NCancerActionableCgi to be an explicit nil

### UnsetNCancerActionableCgi
`func (o *BiomarkersRow) UnsetNCancerActionableCgi()`

UnsetNCancerActionableCgi ensures that no value is present for NCancerActionableCgi, not even an explicit nil
### GetNOtAssociation

`func (o *BiomarkersRow) GetNOtAssociation() int32`

GetNOtAssociation returns the NOtAssociation field if non-nil, zero value otherwise.

### GetNOtAssociationOk

`func (o *BiomarkersRow) GetNOtAssociationOk() (*int32, bool)`

GetNOtAssociationOk returns a tuple with the NOtAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOtAssociation

`func (o *BiomarkersRow) SetNOtAssociation(v int32)`

SetNOtAssociation sets NOtAssociation field to given value.

### HasNOtAssociation

`func (o *BiomarkersRow) HasNOtAssociation() bool`

HasNOtAssociation returns a boolean if a field has been set.

### SetNOtAssociationNil

`func (o *BiomarkersRow) SetNOtAssociationNil(b bool)`

 SetNOtAssociationNil sets the value for NOtAssociation to be an explicit nil

### UnsetNOtAssociation
`func (o *BiomarkersRow) UnsetNOtAssociation()`

UnsetNOtAssociation ensures that no value is present for NOtAssociation, not even an explicit nil
### GetNTrialMention

`func (o *BiomarkersRow) GetNTrialMention() int32`

GetNTrialMention returns the NTrialMention field if non-nil, zero value otherwise.

### GetNTrialMentionOk

`func (o *BiomarkersRow) GetNTrialMentionOk() (*int32, bool)`

GetNTrialMentionOk returns a tuple with the NTrialMention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialMention

`func (o *BiomarkersRow) SetNTrialMention(v int32)`

SetNTrialMention sets NTrialMention field to given value.

### HasNTrialMention

`func (o *BiomarkersRow) HasNTrialMention() bool`

HasNTrialMention returns a boolean if a field has been set.

### SetNTrialMentionNil

`func (o *BiomarkersRow) SetNTrialMentionNil(b bool)`

 SetNTrialMentionNil sets the value for NTrialMention to be an explicit nil

### UnsetNTrialMention
`func (o *BiomarkersRow) UnsetNTrialMention()`

UnsetNTrialMention ensures that no value is present for NTrialMention, not even an explicit nil
### GetNPublicationVariantMentions

`func (o *BiomarkersRow) GetNPublicationVariantMentions() int32`

GetNPublicationVariantMentions returns the NPublicationVariantMentions field if non-nil, zero value otherwise.

### GetNPublicationVariantMentionsOk

`func (o *BiomarkersRow) GetNPublicationVariantMentionsOk() (*int32, bool)`

GetNPublicationVariantMentionsOk returns a tuple with the NPublicationVariantMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPublicationVariantMentions

`func (o *BiomarkersRow) SetNPublicationVariantMentions(v int32)`

SetNPublicationVariantMentions sets NPublicationVariantMentions field to given value.

### HasNPublicationVariantMentions

`func (o *BiomarkersRow) HasNPublicationVariantMentions() bool`

HasNPublicationVariantMentions returns a boolean if a field has been set.

### SetNPublicationVariantMentionsNil

`func (o *BiomarkersRow) SetNPublicationVariantMentionsNil(b bool)`

 SetNPublicationVariantMentionsNil sets the value for NPublicationVariantMentions to be an explicit nil

### UnsetNPublicationVariantMentions
`func (o *BiomarkersRow) UnsetNPublicationVariantMentions()`

UnsetNPublicationVariantMentions ensures that no value is present for NPublicationVariantMentions, not even an explicit nil
### GetHasClinicalValidationEvidence

`func (o *BiomarkersRow) GetHasClinicalValidationEvidence() bool`

GetHasClinicalValidationEvidence returns the HasClinicalValidationEvidence field if non-nil, zero value otherwise.

### GetHasClinicalValidationEvidenceOk

`func (o *BiomarkersRow) GetHasClinicalValidationEvidenceOk() (*bool, bool)`

GetHasClinicalValidationEvidenceOk returns a tuple with the HasClinicalValidationEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasClinicalValidationEvidence

`func (o *BiomarkersRow) SetHasClinicalValidationEvidence(v bool)`

SetHasClinicalValidationEvidence sets HasClinicalValidationEvidence field to given value.

### HasHasClinicalValidationEvidence

`func (o *BiomarkersRow) HasHasClinicalValidationEvidence() bool`

HasHasClinicalValidationEvidence returns a boolean if a field has been set.

### SetHasClinicalValidationEvidenceNil

`func (o *BiomarkersRow) SetHasClinicalValidationEvidenceNil(b bool)`

 SetHasClinicalValidationEvidenceNil sets the value for HasClinicalValidationEvidence to be an explicit nil

### UnsetHasClinicalValidationEvidence
`func (o *BiomarkersRow) UnsetHasClinicalValidationEvidence()`

UnsetHasClinicalValidationEvidence ensures that no value is present for HasClinicalValidationEvidence, not even an explicit nil
### GetHasAnalyticalValidationEvidence

`func (o *BiomarkersRow) GetHasAnalyticalValidationEvidence() bool`

GetHasAnalyticalValidationEvidence returns the HasAnalyticalValidationEvidence field if non-nil, zero value otherwise.

### GetHasAnalyticalValidationEvidenceOk

`func (o *BiomarkersRow) GetHasAnalyticalValidationEvidenceOk() (*bool, bool)`

GetHasAnalyticalValidationEvidenceOk returns a tuple with the HasAnalyticalValidationEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAnalyticalValidationEvidence

`func (o *BiomarkersRow) SetHasAnalyticalValidationEvidence(v bool)`

SetHasAnalyticalValidationEvidence sets HasAnalyticalValidationEvidence field to given value.

### HasHasAnalyticalValidationEvidence

`func (o *BiomarkersRow) HasHasAnalyticalValidationEvidence() bool`

HasHasAnalyticalValidationEvidence returns a boolean if a field has been set.

### SetHasAnalyticalValidationEvidenceNil

`func (o *BiomarkersRow) SetHasAnalyticalValidationEvidenceNil(b bool)`

 SetHasAnalyticalValidationEvidenceNil sets the value for HasAnalyticalValidationEvidence to be an explicit nil

### UnsetHasAnalyticalValidationEvidence
`func (o *BiomarkersRow) UnsetHasAnalyticalValidationEvidence()`

UnsetHasAnalyticalValidationEvidence ensures that no value is present for HasAnalyticalValidationEvidence, not even an explicit nil
### GetNAssayProviderPapers

`func (o *BiomarkersRow) GetNAssayProviderPapers() int32`

GetNAssayProviderPapers returns the NAssayProviderPapers field if non-nil, zero value otherwise.

### GetNAssayProviderPapersOk

`func (o *BiomarkersRow) GetNAssayProviderPapersOk() (*int32, bool)`

GetNAssayProviderPapersOk returns a tuple with the NAssayProviderPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNAssayProviderPapers

`func (o *BiomarkersRow) SetNAssayProviderPapers(v int32)`

SetNAssayProviderPapers sets NAssayProviderPapers field to given value.

### HasNAssayProviderPapers

`func (o *BiomarkersRow) HasNAssayProviderPapers() bool`

HasNAssayProviderPapers returns a boolean if a field has been set.

### SetNAssayProviderPapersNil

`func (o *BiomarkersRow) SetNAssayProviderPapersNil(b bool)`

 SetNAssayProviderPapersNil sets the value for NAssayProviderPapers to be an explicit nil

### UnsetNAssayProviderPapers
`func (o *BiomarkersRow) UnsetNAssayProviderPapers()`

UnsetNAssayProviderPapers ensures that no value is present for NAssayProviderPapers, not even an explicit nil
### GetNAssayProviderSuppliers

`func (o *BiomarkersRow) GetNAssayProviderSuppliers() int32`

GetNAssayProviderSuppliers returns the NAssayProviderSuppliers field if non-nil, zero value otherwise.

### GetNAssayProviderSuppliersOk

`func (o *BiomarkersRow) GetNAssayProviderSuppliersOk() (*int32, bool)`

GetNAssayProviderSuppliersOk returns a tuple with the NAssayProviderSuppliers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNAssayProviderSuppliers

`func (o *BiomarkersRow) SetNAssayProviderSuppliers(v int32)`

SetNAssayProviderSuppliers sets NAssayProviderSuppliers field to given value.

### HasNAssayProviderSuppliers

`func (o *BiomarkersRow) HasNAssayProviderSuppliers() bool`

HasNAssayProviderSuppliers returns a boolean if a field has been set.

### SetNAssayProviderSuppliersNil

`func (o *BiomarkersRow) SetNAssayProviderSuppliersNil(b bool)`

 SetNAssayProviderSuppliersNil sets the value for NAssayProviderSuppliers to be an explicit nil

### UnsetNAssayProviderSuppliers
`func (o *BiomarkersRow) UnsetNAssayProviderSuppliers()`

UnsetNAssayProviderSuppliers ensures that no value is present for NAssayProviderSuppliers, not even an explicit nil
### GetNFdaEvidenceProviders

`func (o *BiomarkersRow) GetNFdaEvidenceProviders() int32`

GetNFdaEvidenceProviders returns the NFdaEvidenceProviders field if non-nil, zero value otherwise.

### GetNFdaEvidenceProvidersOk

`func (o *BiomarkersRow) GetNFdaEvidenceProvidersOk() (*int32, bool)`

GetNFdaEvidenceProvidersOk returns a tuple with the NFdaEvidenceProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFdaEvidenceProviders

`func (o *BiomarkersRow) SetNFdaEvidenceProviders(v int32)`

SetNFdaEvidenceProviders sets NFdaEvidenceProviders field to given value.

### HasNFdaEvidenceProviders

`func (o *BiomarkersRow) HasNFdaEvidenceProviders() bool`

HasNFdaEvidenceProviders returns a boolean if a field has been set.

### SetNFdaEvidenceProvidersNil

`func (o *BiomarkersRow) SetNFdaEvidenceProvidersNil(b bool)`

 SetNFdaEvidenceProvidersNil sets the value for NFdaEvidenceProviders to be an explicit nil

### UnsetNFdaEvidenceProviders
`func (o *BiomarkersRow) UnsetNFdaEvidenceProviders()`

UnsetNFdaEvidenceProviders ensures that no value is present for NFdaEvidenceProviders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


