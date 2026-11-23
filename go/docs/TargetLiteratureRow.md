# TargetLiteratureRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **NullableString** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. | [optional] 
**TargetSymbol** | Pointer to **NullableString** | The target&#39;s HGNC gene symbol. | [optional] 
**Pmid** | Pointer to **NullableString** | The PubMed id. | [optional] 
**Title** | Pointer to **NullableString** | The paper&#39;s title. | [optional] 
**JournalTitle** | Pointer to **NullableString** | The journal&#39;s title. | [optional] 
**PubYear** | Pointer to **NullableInt32** | The publication year (2016 or later: the dataset holds recent MeSH-indexed literature). | [optional] 
**EvidenceType** | Pointer to **NullableString** | The paper&#39;s study-design label, one normalized publication type by study-design priority. | [optional] 
**IsTrialLinked** | Pointer to **NullableBool** | True when a registered clinical trial cites the paper; false for literature no registered trial cites. | [optional] 
**CitationCount** | Pointer to **NullableInt32** | Citations where iCite (the National Institutes of Health citation-metrics service) carries them; null is not measured, never zero. | [optional] 
**RelativeCitationRatio** | Pointer to **NullableFloat32** | The iCite Relative Citation Ratio (RCR) where measured; null is not measured. | [optional] 

## Methods

### NewTargetLiteratureRow

`func NewTargetLiteratureRow() *TargetLiteratureRow`

NewTargetLiteratureRow instantiates a new TargetLiteratureRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetLiteratureRowWithDefaults

`func NewTargetLiteratureRowWithDefaults() *TargetLiteratureRow`

NewTargetLiteratureRowWithDefaults instantiates a new TargetLiteratureRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *TargetLiteratureRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TargetLiteratureRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TargetLiteratureRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *TargetLiteratureRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *TargetLiteratureRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *TargetLiteratureRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetSymbol

`func (o *TargetLiteratureRow) GetTargetSymbol() string`

GetTargetSymbol returns the TargetSymbol field if non-nil, zero value otherwise.

### GetTargetSymbolOk

`func (o *TargetLiteratureRow) GetTargetSymbolOk() (*string, bool)`

GetTargetSymbolOk returns a tuple with the TargetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSymbol

`func (o *TargetLiteratureRow) SetTargetSymbol(v string)`

SetTargetSymbol sets TargetSymbol field to given value.

### HasTargetSymbol

`func (o *TargetLiteratureRow) HasTargetSymbol() bool`

HasTargetSymbol returns a boolean if a field has been set.

### SetTargetSymbolNil

`func (o *TargetLiteratureRow) SetTargetSymbolNil(b bool)`

 SetTargetSymbolNil sets the value for TargetSymbol to be an explicit nil

### UnsetTargetSymbol
`func (o *TargetLiteratureRow) UnsetTargetSymbol()`

UnsetTargetSymbol ensures that no value is present for TargetSymbol, not even an explicit nil
### GetPmid

`func (o *TargetLiteratureRow) GetPmid() string`

GetPmid returns the Pmid field if non-nil, zero value otherwise.

### GetPmidOk

`func (o *TargetLiteratureRow) GetPmidOk() (*string, bool)`

GetPmidOk returns a tuple with the Pmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmid

`func (o *TargetLiteratureRow) SetPmid(v string)`

SetPmid sets Pmid field to given value.

### HasPmid

`func (o *TargetLiteratureRow) HasPmid() bool`

HasPmid returns a boolean if a field has been set.

### SetPmidNil

`func (o *TargetLiteratureRow) SetPmidNil(b bool)`

 SetPmidNil sets the value for Pmid to be an explicit nil

### UnsetPmid
`func (o *TargetLiteratureRow) UnsetPmid()`

UnsetPmid ensures that no value is present for Pmid, not even an explicit nil
### GetTitle

`func (o *TargetLiteratureRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TargetLiteratureRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TargetLiteratureRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TargetLiteratureRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TargetLiteratureRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TargetLiteratureRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetJournalTitle

`func (o *TargetLiteratureRow) GetJournalTitle() string`

GetJournalTitle returns the JournalTitle field if non-nil, zero value otherwise.

### GetJournalTitleOk

`func (o *TargetLiteratureRow) GetJournalTitleOk() (*string, bool)`

GetJournalTitleOk returns a tuple with the JournalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalTitle

`func (o *TargetLiteratureRow) SetJournalTitle(v string)`

SetJournalTitle sets JournalTitle field to given value.

### HasJournalTitle

`func (o *TargetLiteratureRow) HasJournalTitle() bool`

HasJournalTitle returns a boolean if a field has been set.

### SetJournalTitleNil

`func (o *TargetLiteratureRow) SetJournalTitleNil(b bool)`

 SetJournalTitleNil sets the value for JournalTitle to be an explicit nil

### UnsetJournalTitle
`func (o *TargetLiteratureRow) UnsetJournalTitle()`

UnsetJournalTitle ensures that no value is present for JournalTitle, not even an explicit nil
### GetPubYear

`func (o *TargetLiteratureRow) GetPubYear() int32`

GetPubYear returns the PubYear field if non-nil, zero value otherwise.

### GetPubYearOk

`func (o *TargetLiteratureRow) GetPubYearOk() (*int32, bool)`

GetPubYearOk returns a tuple with the PubYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubYear

`func (o *TargetLiteratureRow) SetPubYear(v int32)`

SetPubYear sets PubYear field to given value.

### HasPubYear

`func (o *TargetLiteratureRow) HasPubYear() bool`

HasPubYear returns a boolean if a field has been set.

### SetPubYearNil

`func (o *TargetLiteratureRow) SetPubYearNil(b bool)`

 SetPubYearNil sets the value for PubYear to be an explicit nil

### UnsetPubYear
`func (o *TargetLiteratureRow) UnsetPubYear()`

UnsetPubYear ensures that no value is present for PubYear, not even an explicit nil
### GetEvidenceType

`func (o *TargetLiteratureRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *TargetLiteratureRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *TargetLiteratureRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *TargetLiteratureRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *TargetLiteratureRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *TargetLiteratureRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetIsTrialLinked

`func (o *TargetLiteratureRow) GetIsTrialLinked() bool`

GetIsTrialLinked returns the IsTrialLinked field if non-nil, zero value otherwise.

### GetIsTrialLinkedOk

`func (o *TargetLiteratureRow) GetIsTrialLinkedOk() (*bool, bool)`

GetIsTrialLinkedOk returns a tuple with the IsTrialLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrialLinked

`func (o *TargetLiteratureRow) SetIsTrialLinked(v bool)`

SetIsTrialLinked sets IsTrialLinked field to given value.

### HasIsTrialLinked

`func (o *TargetLiteratureRow) HasIsTrialLinked() bool`

HasIsTrialLinked returns a boolean if a field has been set.

### SetIsTrialLinkedNil

`func (o *TargetLiteratureRow) SetIsTrialLinkedNil(b bool)`

 SetIsTrialLinkedNil sets the value for IsTrialLinked to be an explicit nil

### UnsetIsTrialLinked
`func (o *TargetLiteratureRow) UnsetIsTrialLinked()`

UnsetIsTrialLinked ensures that no value is present for IsTrialLinked, not even an explicit nil
### GetCitationCount

`func (o *TargetLiteratureRow) GetCitationCount() int32`

GetCitationCount returns the CitationCount field if non-nil, zero value otherwise.

### GetCitationCountOk

`func (o *TargetLiteratureRow) GetCitationCountOk() (*int32, bool)`

GetCitationCountOk returns a tuple with the CitationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitationCount

`func (o *TargetLiteratureRow) SetCitationCount(v int32)`

SetCitationCount sets CitationCount field to given value.

### HasCitationCount

`func (o *TargetLiteratureRow) HasCitationCount() bool`

HasCitationCount returns a boolean if a field has been set.

### SetCitationCountNil

`func (o *TargetLiteratureRow) SetCitationCountNil(b bool)`

 SetCitationCountNil sets the value for CitationCount to be an explicit nil

### UnsetCitationCount
`func (o *TargetLiteratureRow) UnsetCitationCount()`

UnsetCitationCount ensures that no value is present for CitationCount, not even an explicit nil
### GetRelativeCitationRatio

`func (o *TargetLiteratureRow) GetRelativeCitationRatio() float32`

GetRelativeCitationRatio returns the RelativeCitationRatio field if non-nil, zero value otherwise.

### GetRelativeCitationRatioOk

`func (o *TargetLiteratureRow) GetRelativeCitationRatioOk() (*float32, bool)`

GetRelativeCitationRatioOk returns a tuple with the RelativeCitationRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeCitationRatio

`func (o *TargetLiteratureRow) SetRelativeCitationRatio(v float32)`

SetRelativeCitationRatio sets RelativeCitationRatio field to given value.

### HasRelativeCitationRatio

`func (o *TargetLiteratureRow) HasRelativeCitationRatio() bool`

HasRelativeCitationRatio returns a boolean if a field has been set.

### SetRelativeCitationRatioNil

`func (o *TargetLiteratureRow) SetRelativeCitationRatioNil(b bool)`

 SetRelativeCitationRatioNil sets the value for RelativeCitationRatio to be an explicit nil

### UnsetRelativeCitationRatio
`func (o *TargetLiteratureRow) UnsetRelativeCitationRatio()`

UnsetRelativeCitationRatio ensures that no value is present for RelativeCitationRatio, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


