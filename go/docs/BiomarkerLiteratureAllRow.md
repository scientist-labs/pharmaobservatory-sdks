# BiomarkerLiteratureAllRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteKey** | Pointer to **NullableString** | The route key of the gene the paper names, the entry key for the biomarkers object type. | [optional] 
**GeneSymbol** | Pointer to **NullableString** | The gene symbol. | [optional] 
**Pmid** | Pointer to **NullableString** | The PubMed id. | [optional] 
**Title** | Pointer to **NullableString** | The paper&#39;s title. | [optional] 
**JournalTitle** | Pointer to **NullableString** | The journal&#39;s title. | [optional] 
**PubYear** | Pointer to **NullableInt32** | The publication year (2016 or later: the dataset holds recent MeSH-indexed literature). | [optional] 
**EvidenceType** | Pointer to **NullableString** | The paper&#39;s study-design label, one normalized publication type by study-design priority. | [optional] 
**IsTrialLinked** | Pointer to **NullableBool** | True when a registered clinical trial cites the paper; false for literature no registered trial cites. | [optional] 
**CitationCount** | Pointer to **NullableInt32** | Citations where iCite (the National Institutes of Health citation-metrics service) carries them; null is not measured, never zero. | [optional] 
**RelativeCitationRatio** | Pointer to **NullableFloat32** | The iCite Relative Citation Ratio (RCR) where measured; null is not measured. | [optional] 

## Methods

### NewBiomarkerLiteratureAllRow

`func NewBiomarkerLiteratureAllRow() *BiomarkerLiteratureAllRow`

NewBiomarkerLiteratureAllRow instantiates a new BiomarkerLiteratureAllRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerLiteratureAllRowWithDefaults

`func NewBiomarkerLiteratureAllRowWithDefaults() *BiomarkerLiteratureAllRow`

NewBiomarkerLiteratureAllRowWithDefaults instantiates a new BiomarkerLiteratureAllRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteKey

`func (o *BiomarkerLiteratureAllRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerLiteratureAllRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerLiteratureAllRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerLiteratureAllRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerLiteratureAllRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerLiteratureAllRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetGeneSymbol

`func (o *BiomarkerLiteratureAllRow) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *BiomarkerLiteratureAllRow) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *BiomarkerLiteratureAllRow) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *BiomarkerLiteratureAllRow) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbolNil

`func (o *BiomarkerLiteratureAllRow) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *BiomarkerLiteratureAllRow) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetPmid

`func (o *BiomarkerLiteratureAllRow) GetPmid() string`

GetPmid returns the Pmid field if non-nil, zero value otherwise.

### GetPmidOk

`func (o *BiomarkerLiteratureAllRow) GetPmidOk() (*string, bool)`

GetPmidOk returns a tuple with the Pmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmid

`func (o *BiomarkerLiteratureAllRow) SetPmid(v string)`

SetPmid sets Pmid field to given value.

### HasPmid

`func (o *BiomarkerLiteratureAllRow) HasPmid() bool`

HasPmid returns a boolean if a field has been set.

### SetPmidNil

`func (o *BiomarkerLiteratureAllRow) SetPmidNil(b bool)`

 SetPmidNil sets the value for Pmid to be an explicit nil

### UnsetPmid
`func (o *BiomarkerLiteratureAllRow) UnsetPmid()`

UnsetPmid ensures that no value is present for Pmid, not even an explicit nil
### GetTitle

`func (o *BiomarkerLiteratureAllRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BiomarkerLiteratureAllRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BiomarkerLiteratureAllRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BiomarkerLiteratureAllRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BiomarkerLiteratureAllRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BiomarkerLiteratureAllRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetJournalTitle

`func (o *BiomarkerLiteratureAllRow) GetJournalTitle() string`

GetJournalTitle returns the JournalTitle field if non-nil, zero value otherwise.

### GetJournalTitleOk

`func (o *BiomarkerLiteratureAllRow) GetJournalTitleOk() (*string, bool)`

GetJournalTitleOk returns a tuple with the JournalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalTitle

`func (o *BiomarkerLiteratureAllRow) SetJournalTitle(v string)`

SetJournalTitle sets JournalTitle field to given value.

### HasJournalTitle

`func (o *BiomarkerLiteratureAllRow) HasJournalTitle() bool`

HasJournalTitle returns a boolean if a field has been set.

### SetJournalTitleNil

`func (o *BiomarkerLiteratureAllRow) SetJournalTitleNil(b bool)`

 SetJournalTitleNil sets the value for JournalTitle to be an explicit nil

### UnsetJournalTitle
`func (o *BiomarkerLiteratureAllRow) UnsetJournalTitle()`

UnsetJournalTitle ensures that no value is present for JournalTitle, not even an explicit nil
### GetPubYear

`func (o *BiomarkerLiteratureAllRow) GetPubYear() int32`

GetPubYear returns the PubYear field if non-nil, zero value otherwise.

### GetPubYearOk

`func (o *BiomarkerLiteratureAllRow) GetPubYearOk() (*int32, bool)`

GetPubYearOk returns a tuple with the PubYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubYear

`func (o *BiomarkerLiteratureAllRow) SetPubYear(v int32)`

SetPubYear sets PubYear field to given value.

### HasPubYear

`func (o *BiomarkerLiteratureAllRow) HasPubYear() bool`

HasPubYear returns a boolean if a field has been set.

### SetPubYearNil

`func (o *BiomarkerLiteratureAllRow) SetPubYearNil(b bool)`

 SetPubYearNil sets the value for PubYear to be an explicit nil

### UnsetPubYear
`func (o *BiomarkerLiteratureAllRow) UnsetPubYear()`

UnsetPubYear ensures that no value is present for PubYear, not even an explicit nil
### GetEvidenceType

`func (o *BiomarkerLiteratureAllRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *BiomarkerLiteratureAllRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *BiomarkerLiteratureAllRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *BiomarkerLiteratureAllRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *BiomarkerLiteratureAllRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *BiomarkerLiteratureAllRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetIsTrialLinked

`func (o *BiomarkerLiteratureAllRow) GetIsTrialLinked() bool`

GetIsTrialLinked returns the IsTrialLinked field if non-nil, zero value otherwise.

### GetIsTrialLinkedOk

`func (o *BiomarkerLiteratureAllRow) GetIsTrialLinkedOk() (*bool, bool)`

GetIsTrialLinkedOk returns a tuple with the IsTrialLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrialLinked

`func (o *BiomarkerLiteratureAllRow) SetIsTrialLinked(v bool)`

SetIsTrialLinked sets IsTrialLinked field to given value.

### HasIsTrialLinked

`func (o *BiomarkerLiteratureAllRow) HasIsTrialLinked() bool`

HasIsTrialLinked returns a boolean if a field has been set.

### SetIsTrialLinkedNil

`func (o *BiomarkerLiteratureAllRow) SetIsTrialLinkedNil(b bool)`

 SetIsTrialLinkedNil sets the value for IsTrialLinked to be an explicit nil

### UnsetIsTrialLinked
`func (o *BiomarkerLiteratureAllRow) UnsetIsTrialLinked()`

UnsetIsTrialLinked ensures that no value is present for IsTrialLinked, not even an explicit nil
### GetCitationCount

`func (o *BiomarkerLiteratureAllRow) GetCitationCount() int32`

GetCitationCount returns the CitationCount field if non-nil, zero value otherwise.

### GetCitationCountOk

`func (o *BiomarkerLiteratureAllRow) GetCitationCountOk() (*int32, bool)`

GetCitationCountOk returns a tuple with the CitationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitationCount

`func (o *BiomarkerLiteratureAllRow) SetCitationCount(v int32)`

SetCitationCount sets CitationCount field to given value.

### HasCitationCount

`func (o *BiomarkerLiteratureAllRow) HasCitationCount() bool`

HasCitationCount returns a boolean if a field has been set.

### SetCitationCountNil

`func (o *BiomarkerLiteratureAllRow) SetCitationCountNil(b bool)`

 SetCitationCountNil sets the value for CitationCount to be an explicit nil

### UnsetCitationCount
`func (o *BiomarkerLiteratureAllRow) UnsetCitationCount()`

UnsetCitationCount ensures that no value is present for CitationCount, not even an explicit nil
### GetRelativeCitationRatio

`func (o *BiomarkerLiteratureAllRow) GetRelativeCitationRatio() float32`

GetRelativeCitationRatio returns the RelativeCitationRatio field if non-nil, zero value otherwise.

### GetRelativeCitationRatioOk

`func (o *BiomarkerLiteratureAllRow) GetRelativeCitationRatioOk() (*float32, bool)`

GetRelativeCitationRatioOk returns a tuple with the RelativeCitationRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeCitationRatio

`func (o *BiomarkerLiteratureAllRow) SetRelativeCitationRatio(v float32)`

SetRelativeCitationRatio sets RelativeCitationRatio field to given value.

### HasRelativeCitationRatio

`func (o *BiomarkerLiteratureAllRow) HasRelativeCitationRatio() bool`

HasRelativeCitationRatio returns a boolean if a field has been set.

### SetRelativeCitationRatioNil

`func (o *BiomarkerLiteratureAllRow) SetRelativeCitationRatioNil(b bool)`

 SetRelativeCitationRatioNil sets the value for RelativeCitationRatio to be an explicit nil

### UnsetRelativeCitationRatio
`func (o *BiomarkerLiteratureAllRow) UnsetRelativeCitationRatio()`

UnsetRelativeCitationRatio ensures that no value is present for RelativeCitationRatio, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


