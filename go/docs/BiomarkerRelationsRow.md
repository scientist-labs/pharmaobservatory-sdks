# BiomarkerRelationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteKey** | Pointer to **NullableString** | The biomarker&#39;s own route key, the entry key for the biomarkers object type (the gene symbol for a gene, the variant key for a protein variant). | [optional] 
**BiomarkerId** | Pointer to **NullableString** | The biomarker&#39;s key. | [optional] 
**BiomarkerKind** | Pointer to **NullableString** | Whether the row describes a gene or a protein variant. | [optional] 
**GeneSymbol** | Pointer to **NullableString** | The gene symbol. | [optional] 
**AminoAcidChange** | Pointer to **NullableString** | The protein change, for a variant. | [optional] 
**RelatedEntityKind** | Pointer to **NullableString** | The kind of the related entity (gene, drug or condition). | [optional] 
**RelatedEntityId** | Pointer to **NullableString** | The related entity&#39;s key in its own vocabulary (a gene symbol, a drug id such as unii:..., or a condition key). | [optional] 
**RelatedDisplayName** | Pointer to **NullableString** | The related entity&#39;s name. | [optional] 
**RelationType** | Pointer to **NullableString** | The source&#39;s relation label, as the source states it. | [optional] 
**Pmid** | Pointer to **NullableString** | The PubMed id of the paper the relation was extracted from. | [optional] 
**Title** | Pointer to **NullableString** | The paper&#39;s title. | [optional] 
**JournalNlmId** | Pointer to **NullableString** | The journal&#39;s National Library of Medicine (NLM) id. | [optional] 
**JournalTitle** | Pointer to **NullableString** | The journal&#39;s title. | [optional] 
**PubYear** | Pointer to **NullableInt32** | The paper&#39;s publication year. | [optional] 
**NSourceAssertions** | Pointer to **NullableInt64** | Source assertions of this relation in the paper. | [optional] 
**NRelationPapers** | Pointer to **NullableInt64** | Papers that carry this biomarker, related entity and relation label. | [optional] 
**EvidenceType** | Pointer to **NullableString** | The evidence type (a text-mined relation, never a clinical claim). | [optional] 
**Source** | Pointer to **NullableString** | The source of the evidence. | [optional] 
**DisplayOrder** | Pointer to **NullableInt64** | The order the web page lists the relation in within its kind. | [optional] 

## Methods

### NewBiomarkerRelationsRow

`func NewBiomarkerRelationsRow() *BiomarkerRelationsRow`

NewBiomarkerRelationsRow instantiates a new BiomarkerRelationsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerRelationsRowWithDefaults

`func NewBiomarkerRelationsRowWithDefaults() *BiomarkerRelationsRow`

NewBiomarkerRelationsRowWithDefaults instantiates a new BiomarkerRelationsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteKey

`func (o *BiomarkerRelationsRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerRelationsRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerRelationsRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerRelationsRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerRelationsRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerRelationsRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetBiomarkerId

`func (o *BiomarkerRelationsRow) GetBiomarkerId() string`

GetBiomarkerId returns the BiomarkerId field if non-nil, zero value otherwise.

### GetBiomarkerIdOk

`func (o *BiomarkerRelationsRow) GetBiomarkerIdOk() (*string, bool)`

GetBiomarkerIdOk returns a tuple with the BiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerId

`func (o *BiomarkerRelationsRow) SetBiomarkerId(v string)`

SetBiomarkerId sets BiomarkerId field to given value.

### HasBiomarkerId

`func (o *BiomarkerRelationsRow) HasBiomarkerId() bool`

HasBiomarkerId returns a boolean if a field has been set.

### SetBiomarkerIdNil

`func (o *BiomarkerRelationsRow) SetBiomarkerIdNil(b bool)`

 SetBiomarkerIdNil sets the value for BiomarkerId to be an explicit nil

### UnsetBiomarkerId
`func (o *BiomarkerRelationsRow) UnsetBiomarkerId()`

UnsetBiomarkerId ensures that no value is present for BiomarkerId, not even an explicit nil
### GetBiomarkerKind

`func (o *BiomarkerRelationsRow) GetBiomarkerKind() string`

GetBiomarkerKind returns the BiomarkerKind field if non-nil, zero value otherwise.

### GetBiomarkerKindOk

`func (o *BiomarkerRelationsRow) GetBiomarkerKindOk() (*string, bool)`

GetBiomarkerKindOk returns a tuple with the BiomarkerKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerKind

`func (o *BiomarkerRelationsRow) SetBiomarkerKind(v string)`

SetBiomarkerKind sets BiomarkerKind field to given value.

### HasBiomarkerKind

`func (o *BiomarkerRelationsRow) HasBiomarkerKind() bool`

HasBiomarkerKind returns a boolean if a field has been set.

### SetBiomarkerKindNil

`func (o *BiomarkerRelationsRow) SetBiomarkerKindNil(b bool)`

 SetBiomarkerKindNil sets the value for BiomarkerKind to be an explicit nil

### UnsetBiomarkerKind
`func (o *BiomarkerRelationsRow) UnsetBiomarkerKind()`

UnsetBiomarkerKind ensures that no value is present for BiomarkerKind, not even an explicit nil
### GetGeneSymbol

`func (o *BiomarkerRelationsRow) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *BiomarkerRelationsRow) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *BiomarkerRelationsRow) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *BiomarkerRelationsRow) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbolNil

`func (o *BiomarkerRelationsRow) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *BiomarkerRelationsRow) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetAminoAcidChange

`func (o *BiomarkerRelationsRow) GetAminoAcidChange() string`

GetAminoAcidChange returns the AminoAcidChange field if non-nil, zero value otherwise.

### GetAminoAcidChangeOk

`func (o *BiomarkerRelationsRow) GetAminoAcidChangeOk() (*string, bool)`

GetAminoAcidChangeOk returns a tuple with the AminoAcidChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcidChange

`func (o *BiomarkerRelationsRow) SetAminoAcidChange(v string)`

SetAminoAcidChange sets AminoAcidChange field to given value.

### HasAminoAcidChange

`func (o *BiomarkerRelationsRow) HasAminoAcidChange() bool`

HasAminoAcidChange returns a boolean if a field has been set.

### SetAminoAcidChangeNil

`func (o *BiomarkerRelationsRow) SetAminoAcidChangeNil(b bool)`

 SetAminoAcidChangeNil sets the value for AminoAcidChange to be an explicit nil

### UnsetAminoAcidChange
`func (o *BiomarkerRelationsRow) UnsetAminoAcidChange()`

UnsetAminoAcidChange ensures that no value is present for AminoAcidChange, not even an explicit nil
### GetRelatedEntityKind

`func (o *BiomarkerRelationsRow) GetRelatedEntityKind() string`

GetRelatedEntityKind returns the RelatedEntityKind field if non-nil, zero value otherwise.

### GetRelatedEntityKindOk

`func (o *BiomarkerRelationsRow) GetRelatedEntityKindOk() (*string, bool)`

GetRelatedEntityKindOk returns a tuple with the RelatedEntityKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedEntityKind

`func (o *BiomarkerRelationsRow) SetRelatedEntityKind(v string)`

SetRelatedEntityKind sets RelatedEntityKind field to given value.

### HasRelatedEntityKind

`func (o *BiomarkerRelationsRow) HasRelatedEntityKind() bool`

HasRelatedEntityKind returns a boolean if a field has been set.

### SetRelatedEntityKindNil

`func (o *BiomarkerRelationsRow) SetRelatedEntityKindNil(b bool)`

 SetRelatedEntityKindNil sets the value for RelatedEntityKind to be an explicit nil

### UnsetRelatedEntityKind
`func (o *BiomarkerRelationsRow) UnsetRelatedEntityKind()`

UnsetRelatedEntityKind ensures that no value is present for RelatedEntityKind, not even an explicit nil
### GetRelatedEntityId

`func (o *BiomarkerRelationsRow) GetRelatedEntityId() string`

GetRelatedEntityId returns the RelatedEntityId field if non-nil, zero value otherwise.

### GetRelatedEntityIdOk

`func (o *BiomarkerRelationsRow) GetRelatedEntityIdOk() (*string, bool)`

GetRelatedEntityIdOk returns a tuple with the RelatedEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedEntityId

`func (o *BiomarkerRelationsRow) SetRelatedEntityId(v string)`

SetRelatedEntityId sets RelatedEntityId field to given value.

### HasRelatedEntityId

`func (o *BiomarkerRelationsRow) HasRelatedEntityId() bool`

HasRelatedEntityId returns a boolean if a field has been set.

### SetRelatedEntityIdNil

`func (o *BiomarkerRelationsRow) SetRelatedEntityIdNil(b bool)`

 SetRelatedEntityIdNil sets the value for RelatedEntityId to be an explicit nil

### UnsetRelatedEntityId
`func (o *BiomarkerRelationsRow) UnsetRelatedEntityId()`

UnsetRelatedEntityId ensures that no value is present for RelatedEntityId, not even an explicit nil
### GetRelatedDisplayName

`func (o *BiomarkerRelationsRow) GetRelatedDisplayName() string`

GetRelatedDisplayName returns the RelatedDisplayName field if non-nil, zero value otherwise.

### GetRelatedDisplayNameOk

`func (o *BiomarkerRelationsRow) GetRelatedDisplayNameOk() (*string, bool)`

GetRelatedDisplayNameOk returns a tuple with the RelatedDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDisplayName

`func (o *BiomarkerRelationsRow) SetRelatedDisplayName(v string)`

SetRelatedDisplayName sets RelatedDisplayName field to given value.

### HasRelatedDisplayName

`func (o *BiomarkerRelationsRow) HasRelatedDisplayName() bool`

HasRelatedDisplayName returns a boolean if a field has been set.

### SetRelatedDisplayNameNil

`func (o *BiomarkerRelationsRow) SetRelatedDisplayNameNil(b bool)`

 SetRelatedDisplayNameNil sets the value for RelatedDisplayName to be an explicit nil

### UnsetRelatedDisplayName
`func (o *BiomarkerRelationsRow) UnsetRelatedDisplayName()`

UnsetRelatedDisplayName ensures that no value is present for RelatedDisplayName, not even an explicit nil
### GetRelationType

`func (o *BiomarkerRelationsRow) GetRelationType() string`

GetRelationType returns the RelationType field if non-nil, zero value otherwise.

### GetRelationTypeOk

`func (o *BiomarkerRelationsRow) GetRelationTypeOk() (*string, bool)`

GetRelationTypeOk returns a tuple with the RelationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationType

`func (o *BiomarkerRelationsRow) SetRelationType(v string)`

SetRelationType sets RelationType field to given value.

### HasRelationType

`func (o *BiomarkerRelationsRow) HasRelationType() bool`

HasRelationType returns a boolean if a field has been set.

### SetRelationTypeNil

`func (o *BiomarkerRelationsRow) SetRelationTypeNil(b bool)`

 SetRelationTypeNil sets the value for RelationType to be an explicit nil

### UnsetRelationType
`func (o *BiomarkerRelationsRow) UnsetRelationType()`

UnsetRelationType ensures that no value is present for RelationType, not even an explicit nil
### GetPmid

`func (o *BiomarkerRelationsRow) GetPmid() string`

GetPmid returns the Pmid field if non-nil, zero value otherwise.

### GetPmidOk

`func (o *BiomarkerRelationsRow) GetPmidOk() (*string, bool)`

GetPmidOk returns a tuple with the Pmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmid

`func (o *BiomarkerRelationsRow) SetPmid(v string)`

SetPmid sets Pmid field to given value.

### HasPmid

`func (o *BiomarkerRelationsRow) HasPmid() bool`

HasPmid returns a boolean if a field has been set.

### SetPmidNil

`func (o *BiomarkerRelationsRow) SetPmidNil(b bool)`

 SetPmidNil sets the value for Pmid to be an explicit nil

### UnsetPmid
`func (o *BiomarkerRelationsRow) UnsetPmid()`

UnsetPmid ensures that no value is present for Pmid, not even an explicit nil
### GetTitle

`func (o *BiomarkerRelationsRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BiomarkerRelationsRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BiomarkerRelationsRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BiomarkerRelationsRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BiomarkerRelationsRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BiomarkerRelationsRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetJournalNlmId

`func (o *BiomarkerRelationsRow) GetJournalNlmId() string`

GetJournalNlmId returns the JournalNlmId field if non-nil, zero value otherwise.

### GetJournalNlmIdOk

`func (o *BiomarkerRelationsRow) GetJournalNlmIdOk() (*string, bool)`

GetJournalNlmIdOk returns a tuple with the JournalNlmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalNlmId

`func (o *BiomarkerRelationsRow) SetJournalNlmId(v string)`

SetJournalNlmId sets JournalNlmId field to given value.

### HasJournalNlmId

`func (o *BiomarkerRelationsRow) HasJournalNlmId() bool`

HasJournalNlmId returns a boolean if a field has been set.

### SetJournalNlmIdNil

`func (o *BiomarkerRelationsRow) SetJournalNlmIdNil(b bool)`

 SetJournalNlmIdNil sets the value for JournalNlmId to be an explicit nil

### UnsetJournalNlmId
`func (o *BiomarkerRelationsRow) UnsetJournalNlmId()`

UnsetJournalNlmId ensures that no value is present for JournalNlmId, not even an explicit nil
### GetJournalTitle

`func (o *BiomarkerRelationsRow) GetJournalTitle() string`

GetJournalTitle returns the JournalTitle field if non-nil, zero value otherwise.

### GetJournalTitleOk

`func (o *BiomarkerRelationsRow) GetJournalTitleOk() (*string, bool)`

GetJournalTitleOk returns a tuple with the JournalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalTitle

`func (o *BiomarkerRelationsRow) SetJournalTitle(v string)`

SetJournalTitle sets JournalTitle field to given value.

### HasJournalTitle

`func (o *BiomarkerRelationsRow) HasJournalTitle() bool`

HasJournalTitle returns a boolean if a field has been set.

### SetJournalTitleNil

`func (o *BiomarkerRelationsRow) SetJournalTitleNil(b bool)`

 SetJournalTitleNil sets the value for JournalTitle to be an explicit nil

### UnsetJournalTitle
`func (o *BiomarkerRelationsRow) UnsetJournalTitle()`

UnsetJournalTitle ensures that no value is present for JournalTitle, not even an explicit nil
### GetPubYear

`func (o *BiomarkerRelationsRow) GetPubYear() int32`

GetPubYear returns the PubYear field if non-nil, zero value otherwise.

### GetPubYearOk

`func (o *BiomarkerRelationsRow) GetPubYearOk() (*int32, bool)`

GetPubYearOk returns a tuple with the PubYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubYear

`func (o *BiomarkerRelationsRow) SetPubYear(v int32)`

SetPubYear sets PubYear field to given value.

### HasPubYear

`func (o *BiomarkerRelationsRow) HasPubYear() bool`

HasPubYear returns a boolean if a field has been set.

### SetPubYearNil

`func (o *BiomarkerRelationsRow) SetPubYearNil(b bool)`

 SetPubYearNil sets the value for PubYear to be an explicit nil

### UnsetPubYear
`func (o *BiomarkerRelationsRow) UnsetPubYear()`

UnsetPubYear ensures that no value is present for PubYear, not even an explicit nil
### GetNSourceAssertions

`func (o *BiomarkerRelationsRow) GetNSourceAssertions() int64`

GetNSourceAssertions returns the NSourceAssertions field if non-nil, zero value otherwise.

### GetNSourceAssertionsOk

`func (o *BiomarkerRelationsRow) GetNSourceAssertionsOk() (*int64, bool)`

GetNSourceAssertionsOk returns a tuple with the NSourceAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSourceAssertions

`func (o *BiomarkerRelationsRow) SetNSourceAssertions(v int64)`

SetNSourceAssertions sets NSourceAssertions field to given value.

### HasNSourceAssertions

`func (o *BiomarkerRelationsRow) HasNSourceAssertions() bool`

HasNSourceAssertions returns a boolean if a field has been set.

### SetNSourceAssertionsNil

`func (o *BiomarkerRelationsRow) SetNSourceAssertionsNil(b bool)`

 SetNSourceAssertionsNil sets the value for NSourceAssertions to be an explicit nil

### UnsetNSourceAssertions
`func (o *BiomarkerRelationsRow) UnsetNSourceAssertions()`

UnsetNSourceAssertions ensures that no value is present for NSourceAssertions, not even an explicit nil
### GetNRelationPapers

`func (o *BiomarkerRelationsRow) GetNRelationPapers() int64`

GetNRelationPapers returns the NRelationPapers field if non-nil, zero value otherwise.

### GetNRelationPapersOk

`func (o *BiomarkerRelationsRow) GetNRelationPapersOk() (*int64, bool)`

GetNRelationPapersOk returns a tuple with the NRelationPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRelationPapers

`func (o *BiomarkerRelationsRow) SetNRelationPapers(v int64)`

SetNRelationPapers sets NRelationPapers field to given value.

### HasNRelationPapers

`func (o *BiomarkerRelationsRow) HasNRelationPapers() bool`

HasNRelationPapers returns a boolean if a field has been set.

### SetNRelationPapersNil

`func (o *BiomarkerRelationsRow) SetNRelationPapersNil(b bool)`

 SetNRelationPapersNil sets the value for NRelationPapers to be an explicit nil

### UnsetNRelationPapers
`func (o *BiomarkerRelationsRow) UnsetNRelationPapers()`

UnsetNRelationPapers ensures that no value is present for NRelationPapers, not even an explicit nil
### GetEvidenceType

`func (o *BiomarkerRelationsRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *BiomarkerRelationsRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *BiomarkerRelationsRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *BiomarkerRelationsRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *BiomarkerRelationsRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *BiomarkerRelationsRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetSource

`func (o *BiomarkerRelationsRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BiomarkerRelationsRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BiomarkerRelationsRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BiomarkerRelationsRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *BiomarkerRelationsRow) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *BiomarkerRelationsRow) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetDisplayOrder

`func (o *BiomarkerRelationsRow) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *BiomarkerRelationsRow) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *BiomarkerRelationsRow) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *BiomarkerRelationsRow) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### SetDisplayOrderNil

`func (o *BiomarkerRelationsRow) SetDisplayOrderNil(b bool)`

 SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil

### UnsetDisplayOrder
`func (o *BiomarkerRelationsRow) UnsetDisplayOrder()`

UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


