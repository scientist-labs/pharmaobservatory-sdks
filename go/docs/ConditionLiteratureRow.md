# ConditionLiteratureRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**Pmid** | Pointer to **NullableString** | The PubMed id. | [optional] 
**Title** | Pointer to **NullableString** | The paper&#39;s title. | [optional] 
**JournalTitle** | Pointer to **NullableString** | The journal&#39;s title. | [optional] 
**PubYear** | Pointer to **NullableInt32** | The publication year (2016 or later: the dataset holds recent MeSH-indexed literature). | [optional] 
**EvidenceType** | Pointer to **NullableString** | The paper&#39;s study-design label, one normalized publication type by study-design priority. | [optional] 
**IsTrialLinked** | Pointer to **NullableBool** | True when a registered clinical trial cites the paper; false for literature no registered trial cites. | [optional] 
**IsMajorTopic** | Pointer to **NullableBool** | True when the condition is a major MeSH topic of the paper. | [optional] 
**CitationCount** | Pointer to **NullableInt32** | Citations where iCite (the National Institutes of Health citation-metrics service) carries them; null is not measured, never zero. | [optional] 
**RelativeCitationRatio** | Pointer to **NullableFloat32** | The iCite Relative Citation Ratio (RCR) where measured; null is not measured. | [optional] 

## Methods

### NewConditionLiteratureRow

`func NewConditionLiteratureRow() *ConditionLiteratureRow`

NewConditionLiteratureRow instantiates a new ConditionLiteratureRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionLiteratureRowWithDefaults

`func NewConditionLiteratureRowWithDefaults() *ConditionLiteratureRow`

NewConditionLiteratureRowWithDefaults instantiates a new ConditionLiteratureRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ConditionLiteratureRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ConditionLiteratureRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ConditionLiteratureRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ConditionLiteratureRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ConditionLiteratureRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ConditionLiteratureRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetPmid

`func (o *ConditionLiteratureRow) GetPmid() string`

GetPmid returns the Pmid field if non-nil, zero value otherwise.

### GetPmidOk

`func (o *ConditionLiteratureRow) GetPmidOk() (*string, bool)`

GetPmidOk returns a tuple with the Pmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmid

`func (o *ConditionLiteratureRow) SetPmid(v string)`

SetPmid sets Pmid field to given value.

### HasPmid

`func (o *ConditionLiteratureRow) HasPmid() bool`

HasPmid returns a boolean if a field has been set.

### SetPmidNil

`func (o *ConditionLiteratureRow) SetPmidNil(b bool)`

 SetPmidNil sets the value for Pmid to be an explicit nil

### UnsetPmid
`func (o *ConditionLiteratureRow) UnsetPmid()`

UnsetPmid ensures that no value is present for Pmid, not even an explicit nil
### GetTitle

`func (o *ConditionLiteratureRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConditionLiteratureRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConditionLiteratureRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ConditionLiteratureRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ConditionLiteratureRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ConditionLiteratureRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetJournalTitle

`func (o *ConditionLiteratureRow) GetJournalTitle() string`

GetJournalTitle returns the JournalTitle field if non-nil, zero value otherwise.

### GetJournalTitleOk

`func (o *ConditionLiteratureRow) GetJournalTitleOk() (*string, bool)`

GetJournalTitleOk returns a tuple with the JournalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalTitle

`func (o *ConditionLiteratureRow) SetJournalTitle(v string)`

SetJournalTitle sets JournalTitle field to given value.

### HasJournalTitle

`func (o *ConditionLiteratureRow) HasJournalTitle() bool`

HasJournalTitle returns a boolean if a field has been set.

### SetJournalTitleNil

`func (o *ConditionLiteratureRow) SetJournalTitleNil(b bool)`

 SetJournalTitleNil sets the value for JournalTitle to be an explicit nil

### UnsetJournalTitle
`func (o *ConditionLiteratureRow) UnsetJournalTitle()`

UnsetJournalTitle ensures that no value is present for JournalTitle, not even an explicit nil
### GetPubYear

`func (o *ConditionLiteratureRow) GetPubYear() int32`

GetPubYear returns the PubYear field if non-nil, zero value otherwise.

### GetPubYearOk

`func (o *ConditionLiteratureRow) GetPubYearOk() (*int32, bool)`

GetPubYearOk returns a tuple with the PubYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubYear

`func (o *ConditionLiteratureRow) SetPubYear(v int32)`

SetPubYear sets PubYear field to given value.

### HasPubYear

`func (o *ConditionLiteratureRow) HasPubYear() bool`

HasPubYear returns a boolean if a field has been set.

### SetPubYearNil

`func (o *ConditionLiteratureRow) SetPubYearNil(b bool)`

 SetPubYearNil sets the value for PubYear to be an explicit nil

### UnsetPubYear
`func (o *ConditionLiteratureRow) UnsetPubYear()`

UnsetPubYear ensures that no value is present for PubYear, not even an explicit nil
### GetEvidenceType

`func (o *ConditionLiteratureRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *ConditionLiteratureRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *ConditionLiteratureRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *ConditionLiteratureRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *ConditionLiteratureRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *ConditionLiteratureRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetIsTrialLinked

`func (o *ConditionLiteratureRow) GetIsTrialLinked() bool`

GetIsTrialLinked returns the IsTrialLinked field if non-nil, zero value otherwise.

### GetIsTrialLinkedOk

`func (o *ConditionLiteratureRow) GetIsTrialLinkedOk() (*bool, bool)`

GetIsTrialLinkedOk returns a tuple with the IsTrialLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrialLinked

`func (o *ConditionLiteratureRow) SetIsTrialLinked(v bool)`

SetIsTrialLinked sets IsTrialLinked field to given value.

### HasIsTrialLinked

`func (o *ConditionLiteratureRow) HasIsTrialLinked() bool`

HasIsTrialLinked returns a boolean if a field has been set.

### SetIsTrialLinkedNil

`func (o *ConditionLiteratureRow) SetIsTrialLinkedNil(b bool)`

 SetIsTrialLinkedNil sets the value for IsTrialLinked to be an explicit nil

### UnsetIsTrialLinked
`func (o *ConditionLiteratureRow) UnsetIsTrialLinked()`

UnsetIsTrialLinked ensures that no value is present for IsTrialLinked, not even an explicit nil
### GetIsMajorTopic

`func (o *ConditionLiteratureRow) GetIsMajorTopic() bool`

GetIsMajorTopic returns the IsMajorTopic field if non-nil, zero value otherwise.

### GetIsMajorTopicOk

`func (o *ConditionLiteratureRow) GetIsMajorTopicOk() (*bool, bool)`

GetIsMajorTopicOk returns a tuple with the IsMajorTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMajorTopic

`func (o *ConditionLiteratureRow) SetIsMajorTopic(v bool)`

SetIsMajorTopic sets IsMajorTopic field to given value.

### HasIsMajorTopic

`func (o *ConditionLiteratureRow) HasIsMajorTopic() bool`

HasIsMajorTopic returns a boolean if a field has been set.

### SetIsMajorTopicNil

`func (o *ConditionLiteratureRow) SetIsMajorTopicNil(b bool)`

 SetIsMajorTopicNil sets the value for IsMajorTopic to be an explicit nil

### UnsetIsMajorTopic
`func (o *ConditionLiteratureRow) UnsetIsMajorTopic()`

UnsetIsMajorTopic ensures that no value is present for IsMajorTopic, not even an explicit nil
### GetCitationCount

`func (o *ConditionLiteratureRow) GetCitationCount() int32`

GetCitationCount returns the CitationCount field if non-nil, zero value otherwise.

### GetCitationCountOk

`func (o *ConditionLiteratureRow) GetCitationCountOk() (*int32, bool)`

GetCitationCountOk returns a tuple with the CitationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitationCount

`func (o *ConditionLiteratureRow) SetCitationCount(v int32)`

SetCitationCount sets CitationCount field to given value.

### HasCitationCount

`func (o *ConditionLiteratureRow) HasCitationCount() bool`

HasCitationCount returns a boolean if a field has been set.

### SetCitationCountNil

`func (o *ConditionLiteratureRow) SetCitationCountNil(b bool)`

 SetCitationCountNil sets the value for CitationCount to be an explicit nil

### UnsetCitationCount
`func (o *ConditionLiteratureRow) UnsetCitationCount()`

UnsetCitationCount ensures that no value is present for CitationCount, not even an explicit nil
### GetRelativeCitationRatio

`func (o *ConditionLiteratureRow) GetRelativeCitationRatio() float32`

GetRelativeCitationRatio returns the RelativeCitationRatio field if non-nil, zero value otherwise.

### GetRelativeCitationRatioOk

`func (o *ConditionLiteratureRow) GetRelativeCitationRatioOk() (*float32, bool)`

GetRelativeCitationRatioOk returns a tuple with the RelativeCitationRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeCitationRatio

`func (o *ConditionLiteratureRow) SetRelativeCitationRatio(v float32)`

SetRelativeCitationRatio sets RelativeCitationRatio field to given value.

### HasRelativeCitationRatio

`func (o *ConditionLiteratureRow) HasRelativeCitationRatio() bool`

HasRelativeCitationRatio returns a boolean if a field has been set.

### SetRelativeCitationRatioNil

`func (o *ConditionLiteratureRow) SetRelativeCitationRatioNil(b bool)`

 SetRelativeCitationRatioNil sets the value for RelativeCitationRatio to be an explicit nil

### UnsetRelativeCitationRatio
`func (o *ConditionLiteratureRow) UnsetRelativeCitationRatio()`

UnsetRelativeCitationRatio ensures that no value is present for RelativeCitationRatio, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


