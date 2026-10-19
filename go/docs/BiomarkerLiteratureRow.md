# BiomarkerLiteratureRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteKey** | Pointer to **NullableString** | The route key of the gene the paper names, the entry key for the biomarkers object type. | [optional] 
**GeneSymbol** | Pointer to **NullableString** | The gene symbol. | [optional] 
**Pmid** | Pointer to **NullableString** | The PubMed id. | [optional] 
**EvidenceType** | Pointer to **NullableString** | The evidence type (a publication mention). | [optional] 
**Title** | Pointer to **NullableString** | The paper&#39;s title. | [optional] 
**JournalNlmId** | Pointer to **NullableString** | The journal&#39;s NLM id. | [optional] 
**JournalTitle** | Pointer to **NullableString** | The journal&#39;s title. | [optional] 
**PubYear** | Pointer to **NullableInt32** | The publication year. | [optional] 
**CitationCount** | Pointer to **NullableInt32** | Citations, where iCite carries them (null is not measured, never zero). | [optional] 
**RelativeCitationRatio** | Pointer to **NullableFloat32** | The iCite relative citation ratio, where measured. | [optional] 
**DisplayOrder** | Pointer to **NullableInt64** | The page&#39;s fixed order: newest year first, then citations. | [optional] 

## Methods

### NewBiomarkerLiteratureRow

`func NewBiomarkerLiteratureRow() *BiomarkerLiteratureRow`

NewBiomarkerLiteratureRow instantiates a new BiomarkerLiteratureRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerLiteratureRowWithDefaults

`func NewBiomarkerLiteratureRowWithDefaults() *BiomarkerLiteratureRow`

NewBiomarkerLiteratureRowWithDefaults instantiates a new BiomarkerLiteratureRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteKey

`func (o *BiomarkerLiteratureRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerLiteratureRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerLiteratureRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerLiteratureRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerLiteratureRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerLiteratureRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetGeneSymbol

`func (o *BiomarkerLiteratureRow) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *BiomarkerLiteratureRow) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *BiomarkerLiteratureRow) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *BiomarkerLiteratureRow) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbolNil

`func (o *BiomarkerLiteratureRow) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *BiomarkerLiteratureRow) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetPmid

`func (o *BiomarkerLiteratureRow) GetPmid() string`

GetPmid returns the Pmid field if non-nil, zero value otherwise.

### GetPmidOk

`func (o *BiomarkerLiteratureRow) GetPmidOk() (*string, bool)`

GetPmidOk returns a tuple with the Pmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmid

`func (o *BiomarkerLiteratureRow) SetPmid(v string)`

SetPmid sets Pmid field to given value.

### HasPmid

`func (o *BiomarkerLiteratureRow) HasPmid() bool`

HasPmid returns a boolean if a field has been set.

### SetPmidNil

`func (o *BiomarkerLiteratureRow) SetPmidNil(b bool)`

 SetPmidNil sets the value for Pmid to be an explicit nil

### UnsetPmid
`func (o *BiomarkerLiteratureRow) UnsetPmid()`

UnsetPmid ensures that no value is present for Pmid, not even an explicit nil
### GetEvidenceType

`func (o *BiomarkerLiteratureRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *BiomarkerLiteratureRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *BiomarkerLiteratureRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *BiomarkerLiteratureRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *BiomarkerLiteratureRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *BiomarkerLiteratureRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetTitle

`func (o *BiomarkerLiteratureRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BiomarkerLiteratureRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BiomarkerLiteratureRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BiomarkerLiteratureRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BiomarkerLiteratureRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BiomarkerLiteratureRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetJournalNlmId

`func (o *BiomarkerLiteratureRow) GetJournalNlmId() string`

GetJournalNlmId returns the JournalNlmId field if non-nil, zero value otherwise.

### GetJournalNlmIdOk

`func (o *BiomarkerLiteratureRow) GetJournalNlmIdOk() (*string, bool)`

GetJournalNlmIdOk returns a tuple with the JournalNlmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalNlmId

`func (o *BiomarkerLiteratureRow) SetJournalNlmId(v string)`

SetJournalNlmId sets JournalNlmId field to given value.

### HasJournalNlmId

`func (o *BiomarkerLiteratureRow) HasJournalNlmId() bool`

HasJournalNlmId returns a boolean if a field has been set.

### SetJournalNlmIdNil

`func (o *BiomarkerLiteratureRow) SetJournalNlmIdNil(b bool)`

 SetJournalNlmIdNil sets the value for JournalNlmId to be an explicit nil

### UnsetJournalNlmId
`func (o *BiomarkerLiteratureRow) UnsetJournalNlmId()`

UnsetJournalNlmId ensures that no value is present for JournalNlmId, not even an explicit nil
### GetJournalTitle

`func (o *BiomarkerLiteratureRow) GetJournalTitle() string`

GetJournalTitle returns the JournalTitle field if non-nil, zero value otherwise.

### GetJournalTitleOk

`func (o *BiomarkerLiteratureRow) GetJournalTitleOk() (*string, bool)`

GetJournalTitleOk returns a tuple with the JournalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalTitle

`func (o *BiomarkerLiteratureRow) SetJournalTitle(v string)`

SetJournalTitle sets JournalTitle field to given value.

### HasJournalTitle

`func (o *BiomarkerLiteratureRow) HasJournalTitle() bool`

HasJournalTitle returns a boolean if a field has been set.

### SetJournalTitleNil

`func (o *BiomarkerLiteratureRow) SetJournalTitleNil(b bool)`

 SetJournalTitleNil sets the value for JournalTitle to be an explicit nil

### UnsetJournalTitle
`func (o *BiomarkerLiteratureRow) UnsetJournalTitle()`

UnsetJournalTitle ensures that no value is present for JournalTitle, not even an explicit nil
### GetPubYear

`func (o *BiomarkerLiteratureRow) GetPubYear() int32`

GetPubYear returns the PubYear field if non-nil, zero value otherwise.

### GetPubYearOk

`func (o *BiomarkerLiteratureRow) GetPubYearOk() (*int32, bool)`

GetPubYearOk returns a tuple with the PubYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubYear

`func (o *BiomarkerLiteratureRow) SetPubYear(v int32)`

SetPubYear sets PubYear field to given value.

### HasPubYear

`func (o *BiomarkerLiteratureRow) HasPubYear() bool`

HasPubYear returns a boolean if a field has been set.

### SetPubYearNil

`func (o *BiomarkerLiteratureRow) SetPubYearNil(b bool)`

 SetPubYearNil sets the value for PubYear to be an explicit nil

### UnsetPubYear
`func (o *BiomarkerLiteratureRow) UnsetPubYear()`

UnsetPubYear ensures that no value is present for PubYear, not even an explicit nil
### GetCitationCount

`func (o *BiomarkerLiteratureRow) GetCitationCount() int32`

GetCitationCount returns the CitationCount field if non-nil, zero value otherwise.

### GetCitationCountOk

`func (o *BiomarkerLiteratureRow) GetCitationCountOk() (*int32, bool)`

GetCitationCountOk returns a tuple with the CitationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitationCount

`func (o *BiomarkerLiteratureRow) SetCitationCount(v int32)`

SetCitationCount sets CitationCount field to given value.

### HasCitationCount

`func (o *BiomarkerLiteratureRow) HasCitationCount() bool`

HasCitationCount returns a boolean if a field has been set.

### SetCitationCountNil

`func (o *BiomarkerLiteratureRow) SetCitationCountNil(b bool)`

 SetCitationCountNil sets the value for CitationCount to be an explicit nil

### UnsetCitationCount
`func (o *BiomarkerLiteratureRow) UnsetCitationCount()`

UnsetCitationCount ensures that no value is present for CitationCount, not even an explicit nil
### GetRelativeCitationRatio

`func (o *BiomarkerLiteratureRow) GetRelativeCitationRatio() float32`

GetRelativeCitationRatio returns the RelativeCitationRatio field if non-nil, zero value otherwise.

### GetRelativeCitationRatioOk

`func (o *BiomarkerLiteratureRow) GetRelativeCitationRatioOk() (*float32, bool)`

GetRelativeCitationRatioOk returns a tuple with the RelativeCitationRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeCitationRatio

`func (o *BiomarkerLiteratureRow) SetRelativeCitationRatio(v float32)`

SetRelativeCitationRatio sets RelativeCitationRatio field to given value.

### HasRelativeCitationRatio

`func (o *BiomarkerLiteratureRow) HasRelativeCitationRatio() bool`

HasRelativeCitationRatio returns a boolean if a field has been set.

### SetRelativeCitationRatioNil

`func (o *BiomarkerLiteratureRow) SetRelativeCitationRatioNil(b bool)`

 SetRelativeCitationRatioNil sets the value for RelativeCitationRatio to be an explicit nil

### UnsetRelativeCitationRatio
`func (o *BiomarkerLiteratureRow) UnsetRelativeCitationRatio()`

UnsetRelativeCitationRatio ensures that no value is present for RelativeCitationRatio, not even an explicit nil
### GetDisplayOrder

`func (o *BiomarkerLiteratureRow) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *BiomarkerLiteratureRow) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *BiomarkerLiteratureRow) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *BiomarkerLiteratureRow) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### SetDisplayOrderNil

`func (o *BiomarkerLiteratureRow) SetDisplayOrderNil(b bool)`

 SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil

### UnsetDisplayOrder
`func (o *BiomarkerLiteratureRow) UnsetDisplayOrder()`

UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


