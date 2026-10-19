# BiomarkerVariantMentionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteKey** | Pointer to **NullableString** | The route key of the gene the variants belong to, the entry key for the biomarkers object type. | [optional] 
**GeneSymbol** | Pointer to **NullableString** | The gene symbol. | [optional] 
**VariantBiomarkerId** | Pointer to **NullableString** | The variant&#39;s own biomarker key. | [optional] 
**VariantRouteKey** | Pointer to **NullableString** | The variant&#39;s own route key (the id its web page uses). | [optional] 
**DisplayName** | Pointer to **NullableString** | The variant&#39;s display name. | [optional] 
**AminoAcidChange** | Pointer to **NullableString** | The protein change, for example V600E. | [optional] 
**BiomarkerKind** | Pointer to **NullableString** | The kind of biomarker the row describes (a protein variant). | [optional] 
**EvidenceType** | Pointer to **NullableString** | The evidence type (a literature mention, never an actionability claim). | [optional] 
**Source** | Pointer to **NullableString** | The source of the evidence. | [optional] 
**NPapers** | Pointer to **NullableInt64** | Distinct papers that name the variant. | [optional] 
**FirstPublicationYear** | Pointer to **NullableInt32** | Earliest publication year. | [optional] 
**LastPublicationYear** | Pointer to **NullableInt32** | Latest publication year. | [optional] 
**SamplePmids** | Pointer to **[]string** | Up to five sample PubMed ids, newest first. | [optional] 
**SampleCitations** | Pointer to **[]string** | Up to five sample citations, newest first. | [optional] 

## Methods

### NewBiomarkerVariantMentionsRow

`func NewBiomarkerVariantMentionsRow() *BiomarkerVariantMentionsRow`

NewBiomarkerVariantMentionsRow instantiates a new BiomarkerVariantMentionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerVariantMentionsRowWithDefaults

`func NewBiomarkerVariantMentionsRowWithDefaults() *BiomarkerVariantMentionsRow`

NewBiomarkerVariantMentionsRowWithDefaults instantiates a new BiomarkerVariantMentionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteKey

`func (o *BiomarkerVariantMentionsRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerVariantMentionsRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerVariantMentionsRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerVariantMentionsRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerVariantMentionsRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerVariantMentionsRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetGeneSymbol

`func (o *BiomarkerVariantMentionsRow) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *BiomarkerVariantMentionsRow) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *BiomarkerVariantMentionsRow) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *BiomarkerVariantMentionsRow) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbolNil

`func (o *BiomarkerVariantMentionsRow) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *BiomarkerVariantMentionsRow) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetVariantBiomarkerId

`func (o *BiomarkerVariantMentionsRow) GetVariantBiomarkerId() string`

GetVariantBiomarkerId returns the VariantBiomarkerId field if non-nil, zero value otherwise.

### GetVariantBiomarkerIdOk

`func (o *BiomarkerVariantMentionsRow) GetVariantBiomarkerIdOk() (*string, bool)`

GetVariantBiomarkerIdOk returns a tuple with the VariantBiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantBiomarkerId

`func (o *BiomarkerVariantMentionsRow) SetVariantBiomarkerId(v string)`

SetVariantBiomarkerId sets VariantBiomarkerId field to given value.

### HasVariantBiomarkerId

`func (o *BiomarkerVariantMentionsRow) HasVariantBiomarkerId() bool`

HasVariantBiomarkerId returns a boolean if a field has been set.

### SetVariantBiomarkerIdNil

`func (o *BiomarkerVariantMentionsRow) SetVariantBiomarkerIdNil(b bool)`

 SetVariantBiomarkerIdNil sets the value for VariantBiomarkerId to be an explicit nil

### UnsetVariantBiomarkerId
`func (o *BiomarkerVariantMentionsRow) UnsetVariantBiomarkerId()`

UnsetVariantBiomarkerId ensures that no value is present for VariantBiomarkerId, not even an explicit nil
### GetVariantRouteKey

`func (o *BiomarkerVariantMentionsRow) GetVariantRouteKey() string`

GetVariantRouteKey returns the VariantRouteKey field if non-nil, zero value otherwise.

### GetVariantRouteKeyOk

`func (o *BiomarkerVariantMentionsRow) GetVariantRouteKeyOk() (*string, bool)`

GetVariantRouteKeyOk returns a tuple with the VariantRouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantRouteKey

`func (o *BiomarkerVariantMentionsRow) SetVariantRouteKey(v string)`

SetVariantRouteKey sets VariantRouteKey field to given value.

### HasVariantRouteKey

`func (o *BiomarkerVariantMentionsRow) HasVariantRouteKey() bool`

HasVariantRouteKey returns a boolean if a field has been set.

### SetVariantRouteKeyNil

`func (o *BiomarkerVariantMentionsRow) SetVariantRouteKeyNil(b bool)`

 SetVariantRouteKeyNil sets the value for VariantRouteKey to be an explicit nil

### UnsetVariantRouteKey
`func (o *BiomarkerVariantMentionsRow) UnsetVariantRouteKey()`

UnsetVariantRouteKey ensures that no value is present for VariantRouteKey, not even an explicit nil
### GetDisplayName

`func (o *BiomarkerVariantMentionsRow) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BiomarkerVariantMentionsRow) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BiomarkerVariantMentionsRow) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BiomarkerVariantMentionsRow) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *BiomarkerVariantMentionsRow) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *BiomarkerVariantMentionsRow) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetAminoAcidChange

`func (o *BiomarkerVariantMentionsRow) GetAminoAcidChange() string`

GetAminoAcidChange returns the AminoAcidChange field if non-nil, zero value otherwise.

### GetAminoAcidChangeOk

`func (o *BiomarkerVariantMentionsRow) GetAminoAcidChangeOk() (*string, bool)`

GetAminoAcidChangeOk returns a tuple with the AminoAcidChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcidChange

`func (o *BiomarkerVariantMentionsRow) SetAminoAcidChange(v string)`

SetAminoAcidChange sets AminoAcidChange field to given value.

### HasAminoAcidChange

`func (o *BiomarkerVariantMentionsRow) HasAminoAcidChange() bool`

HasAminoAcidChange returns a boolean if a field has been set.

### SetAminoAcidChangeNil

`func (o *BiomarkerVariantMentionsRow) SetAminoAcidChangeNil(b bool)`

 SetAminoAcidChangeNil sets the value for AminoAcidChange to be an explicit nil

### UnsetAminoAcidChange
`func (o *BiomarkerVariantMentionsRow) UnsetAminoAcidChange()`

UnsetAminoAcidChange ensures that no value is present for AminoAcidChange, not even an explicit nil
### GetBiomarkerKind

`func (o *BiomarkerVariantMentionsRow) GetBiomarkerKind() string`

GetBiomarkerKind returns the BiomarkerKind field if non-nil, zero value otherwise.

### GetBiomarkerKindOk

`func (o *BiomarkerVariantMentionsRow) GetBiomarkerKindOk() (*string, bool)`

GetBiomarkerKindOk returns a tuple with the BiomarkerKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerKind

`func (o *BiomarkerVariantMentionsRow) SetBiomarkerKind(v string)`

SetBiomarkerKind sets BiomarkerKind field to given value.

### HasBiomarkerKind

`func (o *BiomarkerVariantMentionsRow) HasBiomarkerKind() bool`

HasBiomarkerKind returns a boolean if a field has been set.

### SetBiomarkerKindNil

`func (o *BiomarkerVariantMentionsRow) SetBiomarkerKindNil(b bool)`

 SetBiomarkerKindNil sets the value for BiomarkerKind to be an explicit nil

### UnsetBiomarkerKind
`func (o *BiomarkerVariantMentionsRow) UnsetBiomarkerKind()`

UnsetBiomarkerKind ensures that no value is present for BiomarkerKind, not even an explicit nil
### GetEvidenceType

`func (o *BiomarkerVariantMentionsRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *BiomarkerVariantMentionsRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *BiomarkerVariantMentionsRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *BiomarkerVariantMentionsRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *BiomarkerVariantMentionsRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *BiomarkerVariantMentionsRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetSource

`func (o *BiomarkerVariantMentionsRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BiomarkerVariantMentionsRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BiomarkerVariantMentionsRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BiomarkerVariantMentionsRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *BiomarkerVariantMentionsRow) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *BiomarkerVariantMentionsRow) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetNPapers

`func (o *BiomarkerVariantMentionsRow) GetNPapers() int64`

GetNPapers returns the NPapers field if non-nil, zero value otherwise.

### GetNPapersOk

`func (o *BiomarkerVariantMentionsRow) GetNPapersOk() (*int64, bool)`

GetNPapersOk returns a tuple with the NPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPapers

`func (o *BiomarkerVariantMentionsRow) SetNPapers(v int64)`

SetNPapers sets NPapers field to given value.

### HasNPapers

`func (o *BiomarkerVariantMentionsRow) HasNPapers() bool`

HasNPapers returns a boolean if a field has been set.

### SetNPapersNil

`func (o *BiomarkerVariantMentionsRow) SetNPapersNil(b bool)`

 SetNPapersNil sets the value for NPapers to be an explicit nil

### UnsetNPapers
`func (o *BiomarkerVariantMentionsRow) UnsetNPapers()`

UnsetNPapers ensures that no value is present for NPapers, not even an explicit nil
### GetFirstPublicationYear

`func (o *BiomarkerVariantMentionsRow) GetFirstPublicationYear() int32`

GetFirstPublicationYear returns the FirstPublicationYear field if non-nil, zero value otherwise.

### GetFirstPublicationYearOk

`func (o *BiomarkerVariantMentionsRow) GetFirstPublicationYearOk() (*int32, bool)`

GetFirstPublicationYearOk returns a tuple with the FirstPublicationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPublicationYear

`func (o *BiomarkerVariantMentionsRow) SetFirstPublicationYear(v int32)`

SetFirstPublicationYear sets FirstPublicationYear field to given value.

### HasFirstPublicationYear

`func (o *BiomarkerVariantMentionsRow) HasFirstPublicationYear() bool`

HasFirstPublicationYear returns a boolean if a field has been set.

### SetFirstPublicationYearNil

`func (o *BiomarkerVariantMentionsRow) SetFirstPublicationYearNil(b bool)`

 SetFirstPublicationYearNil sets the value for FirstPublicationYear to be an explicit nil

### UnsetFirstPublicationYear
`func (o *BiomarkerVariantMentionsRow) UnsetFirstPublicationYear()`

UnsetFirstPublicationYear ensures that no value is present for FirstPublicationYear, not even an explicit nil
### GetLastPublicationYear

`func (o *BiomarkerVariantMentionsRow) GetLastPublicationYear() int32`

GetLastPublicationYear returns the LastPublicationYear field if non-nil, zero value otherwise.

### GetLastPublicationYearOk

`func (o *BiomarkerVariantMentionsRow) GetLastPublicationYearOk() (*int32, bool)`

GetLastPublicationYearOk returns a tuple with the LastPublicationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPublicationYear

`func (o *BiomarkerVariantMentionsRow) SetLastPublicationYear(v int32)`

SetLastPublicationYear sets LastPublicationYear field to given value.

### HasLastPublicationYear

`func (o *BiomarkerVariantMentionsRow) HasLastPublicationYear() bool`

HasLastPublicationYear returns a boolean if a field has been set.

### SetLastPublicationYearNil

`func (o *BiomarkerVariantMentionsRow) SetLastPublicationYearNil(b bool)`

 SetLastPublicationYearNil sets the value for LastPublicationYear to be an explicit nil

### UnsetLastPublicationYear
`func (o *BiomarkerVariantMentionsRow) UnsetLastPublicationYear()`

UnsetLastPublicationYear ensures that no value is present for LastPublicationYear, not even an explicit nil
### GetSamplePmids

`func (o *BiomarkerVariantMentionsRow) GetSamplePmids() []string`

GetSamplePmids returns the SamplePmids field if non-nil, zero value otherwise.

### GetSamplePmidsOk

`func (o *BiomarkerVariantMentionsRow) GetSamplePmidsOk() (*[]string, bool)`

GetSamplePmidsOk returns a tuple with the SamplePmids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplePmids

`func (o *BiomarkerVariantMentionsRow) SetSamplePmids(v []string)`

SetSamplePmids sets SamplePmids field to given value.

### HasSamplePmids

`func (o *BiomarkerVariantMentionsRow) HasSamplePmids() bool`

HasSamplePmids returns a boolean if a field has been set.

### SetSamplePmidsNil

`func (o *BiomarkerVariantMentionsRow) SetSamplePmidsNil(b bool)`

 SetSamplePmidsNil sets the value for SamplePmids to be an explicit nil

### UnsetSamplePmids
`func (o *BiomarkerVariantMentionsRow) UnsetSamplePmids()`

UnsetSamplePmids ensures that no value is present for SamplePmids, not even an explicit nil
### GetSampleCitations

`func (o *BiomarkerVariantMentionsRow) GetSampleCitations() []string`

GetSampleCitations returns the SampleCitations field if non-nil, zero value otherwise.

### GetSampleCitationsOk

`func (o *BiomarkerVariantMentionsRow) GetSampleCitationsOk() (*[]string, bool)`

GetSampleCitationsOk returns a tuple with the SampleCitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleCitations

`func (o *BiomarkerVariantMentionsRow) SetSampleCitations(v []string)`

SetSampleCitations sets SampleCitations field to given value.

### HasSampleCitations

`func (o *BiomarkerVariantMentionsRow) HasSampleCitations() bool`

HasSampleCitations returns a boolean if a field has been set.

### SetSampleCitationsNil

`func (o *BiomarkerVariantMentionsRow) SetSampleCitationsNil(b bool)`

 SetSampleCitationsNil sets the value for SampleCitations to be an explicit nil

### UnsetSampleCitations
`func (o *BiomarkerVariantMentionsRow) UnsetSampleCitations()`

UnsetSampleCitations ensures that no value is present for SampleCitations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


