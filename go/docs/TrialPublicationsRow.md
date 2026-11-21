# TrialPublicationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**Pmid** | Pointer to **NullableString** | The PubMed id of the paper. | [optional] 
**Title** | Pointer to **NullableString** | The paper&#39;s title. | [optional] 
**JournalTitle** | Pointer to **NullableString** | The journal&#39;s title (null when the paper&#39;s journal is not in the journal table). | [optional] 
**PubYear** | Pointer to **NullableInt32** | The publication year (null when undated). | [optional] 
**EvidenceType** | Pointer to **NullableString** | The paper&#39;s study-design label, one normalized publication type by study-design priority; the page&#39;s design column. | [optional] 
**Citation** | Pointer to **NullableString** | The citation for the paper: one citation per paper, the longest among the trials that cite the PubMed id. | [optional] 
**ReferenceTypes** | Pointer to **[]string** | Every reference kind the registry filed for this paper on this trial, lower-cased (background, result, derived). | [optional] 
**Results** | Pointer to **NullableBool** | True when any reference kind of the row is a results reference; the page&#39;s Results label. | [optional] 

## Methods

### NewTrialPublicationsRow

`func NewTrialPublicationsRow() *TrialPublicationsRow`

NewTrialPublicationsRow instantiates a new TrialPublicationsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialPublicationsRowWithDefaults

`func NewTrialPublicationsRowWithDefaults() *TrialPublicationsRow`

NewTrialPublicationsRowWithDefaults instantiates a new TrialPublicationsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *TrialPublicationsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TrialPublicationsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TrialPublicationsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TrialPublicationsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TrialPublicationsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TrialPublicationsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetPmid

`func (o *TrialPublicationsRow) GetPmid() string`

GetPmid returns the Pmid field if non-nil, zero value otherwise.

### GetPmidOk

`func (o *TrialPublicationsRow) GetPmidOk() (*string, bool)`

GetPmidOk returns a tuple with the Pmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmid

`func (o *TrialPublicationsRow) SetPmid(v string)`

SetPmid sets Pmid field to given value.

### HasPmid

`func (o *TrialPublicationsRow) HasPmid() bool`

HasPmid returns a boolean if a field has been set.

### SetPmidNil

`func (o *TrialPublicationsRow) SetPmidNil(b bool)`

 SetPmidNil sets the value for Pmid to be an explicit nil

### UnsetPmid
`func (o *TrialPublicationsRow) UnsetPmid()`

UnsetPmid ensures that no value is present for Pmid, not even an explicit nil
### GetTitle

`func (o *TrialPublicationsRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TrialPublicationsRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TrialPublicationsRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TrialPublicationsRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TrialPublicationsRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TrialPublicationsRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetJournalTitle

`func (o *TrialPublicationsRow) GetJournalTitle() string`

GetJournalTitle returns the JournalTitle field if non-nil, zero value otherwise.

### GetJournalTitleOk

`func (o *TrialPublicationsRow) GetJournalTitleOk() (*string, bool)`

GetJournalTitleOk returns a tuple with the JournalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalTitle

`func (o *TrialPublicationsRow) SetJournalTitle(v string)`

SetJournalTitle sets JournalTitle field to given value.

### HasJournalTitle

`func (o *TrialPublicationsRow) HasJournalTitle() bool`

HasJournalTitle returns a boolean if a field has been set.

### SetJournalTitleNil

`func (o *TrialPublicationsRow) SetJournalTitleNil(b bool)`

 SetJournalTitleNil sets the value for JournalTitle to be an explicit nil

### UnsetJournalTitle
`func (o *TrialPublicationsRow) UnsetJournalTitle()`

UnsetJournalTitle ensures that no value is present for JournalTitle, not even an explicit nil
### GetPubYear

`func (o *TrialPublicationsRow) GetPubYear() int32`

GetPubYear returns the PubYear field if non-nil, zero value otherwise.

### GetPubYearOk

`func (o *TrialPublicationsRow) GetPubYearOk() (*int32, bool)`

GetPubYearOk returns a tuple with the PubYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubYear

`func (o *TrialPublicationsRow) SetPubYear(v int32)`

SetPubYear sets PubYear field to given value.

### HasPubYear

`func (o *TrialPublicationsRow) HasPubYear() bool`

HasPubYear returns a boolean if a field has been set.

### SetPubYearNil

`func (o *TrialPublicationsRow) SetPubYearNil(b bool)`

 SetPubYearNil sets the value for PubYear to be an explicit nil

### UnsetPubYear
`func (o *TrialPublicationsRow) UnsetPubYear()`

UnsetPubYear ensures that no value is present for PubYear, not even an explicit nil
### GetEvidenceType

`func (o *TrialPublicationsRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *TrialPublicationsRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *TrialPublicationsRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *TrialPublicationsRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *TrialPublicationsRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *TrialPublicationsRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetCitation

`func (o *TrialPublicationsRow) GetCitation() string`

GetCitation returns the Citation field if non-nil, zero value otherwise.

### GetCitationOk

`func (o *TrialPublicationsRow) GetCitationOk() (*string, bool)`

GetCitationOk returns a tuple with the Citation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitation

`func (o *TrialPublicationsRow) SetCitation(v string)`

SetCitation sets Citation field to given value.

### HasCitation

`func (o *TrialPublicationsRow) HasCitation() bool`

HasCitation returns a boolean if a field has been set.

### SetCitationNil

`func (o *TrialPublicationsRow) SetCitationNil(b bool)`

 SetCitationNil sets the value for Citation to be an explicit nil

### UnsetCitation
`func (o *TrialPublicationsRow) UnsetCitation()`

UnsetCitation ensures that no value is present for Citation, not even an explicit nil
### GetReferenceTypes

`func (o *TrialPublicationsRow) GetReferenceTypes() []string`

GetReferenceTypes returns the ReferenceTypes field if non-nil, zero value otherwise.

### GetReferenceTypesOk

`func (o *TrialPublicationsRow) GetReferenceTypesOk() (*[]string, bool)`

GetReferenceTypesOk returns a tuple with the ReferenceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTypes

`func (o *TrialPublicationsRow) SetReferenceTypes(v []string)`

SetReferenceTypes sets ReferenceTypes field to given value.

### HasReferenceTypes

`func (o *TrialPublicationsRow) HasReferenceTypes() bool`

HasReferenceTypes returns a boolean if a field has been set.

### SetReferenceTypesNil

`func (o *TrialPublicationsRow) SetReferenceTypesNil(b bool)`

 SetReferenceTypesNil sets the value for ReferenceTypes to be an explicit nil

### UnsetReferenceTypes
`func (o *TrialPublicationsRow) UnsetReferenceTypes()`

UnsetReferenceTypes ensures that no value is present for ReferenceTypes, not even an explicit nil
### GetResults

`func (o *TrialPublicationsRow) GetResults() bool`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TrialPublicationsRow) GetResultsOk() (*bool, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TrialPublicationsRow) SetResults(v bool)`

SetResults sets Results field to given value.

### HasResults

`func (o *TrialPublicationsRow) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *TrialPublicationsRow) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *TrialPublicationsRow) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


