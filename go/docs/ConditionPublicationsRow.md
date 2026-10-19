# ConditionPublicationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**Pmid** | Pointer to **NullableString** | The PubMed id. | [optional] 
**Title** | Pointer to **NullableString** | The article title. | [optional] 
**PubYear** | Pointer to **NullableInt32** | Publication year. | [optional] 
**JournalTitle** | Pointer to **NullableString** | The journal title. | [optional] 
**EvidenceType** | Pointer to **NullableString** | The article&#39;s PubMed publication type, one per article, picked by study design (Randomized Controlled Trial, Journal Article and so on). | [optional] 
**IsMajorTopic** | Pointer to **NullableBool** | True when the condition is a major topic of the article. | [optional] 
**NLinkedTrials** | Pointer to **NullableInt64** | Registry trials linked to the article. | [optional] 

## Methods

### NewConditionPublicationsRow

`func NewConditionPublicationsRow() *ConditionPublicationsRow`

NewConditionPublicationsRow instantiates a new ConditionPublicationsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionPublicationsRowWithDefaults

`func NewConditionPublicationsRowWithDefaults() *ConditionPublicationsRow`

NewConditionPublicationsRowWithDefaults instantiates a new ConditionPublicationsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ConditionPublicationsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ConditionPublicationsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ConditionPublicationsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ConditionPublicationsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ConditionPublicationsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ConditionPublicationsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetPmid

`func (o *ConditionPublicationsRow) GetPmid() string`

GetPmid returns the Pmid field if non-nil, zero value otherwise.

### GetPmidOk

`func (o *ConditionPublicationsRow) GetPmidOk() (*string, bool)`

GetPmidOk returns a tuple with the Pmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmid

`func (o *ConditionPublicationsRow) SetPmid(v string)`

SetPmid sets Pmid field to given value.

### HasPmid

`func (o *ConditionPublicationsRow) HasPmid() bool`

HasPmid returns a boolean if a field has been set.

### SetPmidNil

`func (o *ConditionPublicationsRow) SetPmidNil(b bool)`

 SetPmidNil sets the value for Pmid to be an explicit nil

### UnsetPmid
`func (o *ConditionPublicationsRow) UnsetPmid()`

UnsetPmid ensures that no value is present for Pmid, not even an explicit nil
### GetTitle

`func (o *ConditionPublicationsRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConditionPublicationsRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConditionPublicationsRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ConditionPublicationsRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ConditionPublicationsRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ConditionPublicationsRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPubYear

`func (o *ConditionPublicationsRow) GetPubYear() int32`

GetPubYear returns the PubYear field if non-nil, zero value otherwise.

### GetPubYearOk

`func (o *ConditionPublicationsRow) GetPubYearOk() (*int32, bool)`

GetPubYearOk returns a tuple with the PubYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubYear

`func (o *ConditionPublicationsRow) SetPubYear(v int32)`

SetPubYear sets PubYear field to given value.

### HasPubYear

`func (o *ConditionPublicationsRow) HasPubYear() bool`

HasPubYear returns a boolean if a field has been set.

### SetPubYearNil

`func (o *ConditionPublicationsRow) SetPubYearNil(b bool)`

 SetPubYearNil sets the value for PubYear to be an explicit nil

### UnsetPubYear
`func (o *ConditionPublicationsRow) UnsetPubYear()`

UnsetPubYear ensures that no value is present for PubYear, not even an explicit nil
### GetJournalTitle

`func (o *ConditionPublicationsRow) GetJournalTitle() string`

GetJournalTitle returns the JournalTitle field if non-nil, zero value otherwise.

### GetJournalTitleOk

`func (o *ConditionPublicationsRow) GetJournalTitleOk() (*string, bool)`

GetJournalTitleOk returns a tuple with the JournalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalTitle

`func (o *ConditionPublicationsRow) SetJournalTitle(v string)`

SetJournalTitle sets JournalTitle field to given value.

### HasJournalTitle

`func (o *ConditionPublicationsRow) HasJournalTitle() bool`

HasJournalTitle returns a boolean if a field has been set.

### SetJournalTitleNil

`func (o *ConditionPublicationsRow) SetJournalTitleNil(b bool)`

 SetJournalTitleNil sets the value for JournalTitle to be an explicit nil

### UnsetJournalTitle
`func (o *ConditionPublicationsRow) UnsetJournalTitle()`

UnsetJournalTitle ensures that no value is present for JournalTitle, not even an explicit nil
### GetEvidenceType

`func (o *ConditionPublicationsRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *ConditionPublicationsRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *ConditionPublicationsRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *ConditionPublicationsRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *ConditionPublicationsRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *ConditionPublicationsRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetIsMajorTopic

`func (o *ConditionPublicationsRow) GetIsMajorTopic() bool`

GetIsMajorTopic returns the IsMajorTopic field if non-nil, zero value otherwise.

### GetIsMajorTopicOk

`func (o *ConditionPublicationsRow) GetIsMajorTopicOk() (*bool, bool)`

GetIsMajorTopicOk returns a tuple with the IsMajorTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMajorTopic

`func (o *ConditionPublicationsRow) SetIsMajorTopic(v bool)`

SetIsMajorTopic sets IsMajorTopic field to given value.

### HasIsMajorTopic

`func (o *ConditionPublicationsRow) HasIsMajorTopic() bool`

HasIsMajorTopic returns a boolean if a field has been set.

### SetIsMajorTopicNil

`func (o *ConditionPublicationsRow) SetIsMajorTopicNil(b bool)`

 SetIsMajorTopicNil sets the value for IsMajorTopic to be an explicit nil

### UnsetIsMajorTopic
`func (o *ConditionPublicationsRow) UnsetIsMajorTopic()`

UnsetIsMajorTopic ensures that no value is present for IsMajorTopic, not even an explicit nil
### GetNLinkedTrials

`func (o *ConditionPublicationsRow) GetNLinkedTrials() int64`

GetNLinkedTrials returns the NLinkedTrials field if non-nil, zero value otherwise.

### GetNLinkedTrialsOk

`func (o *ConditionPublicationsRow) GetNLinkedTrialsOk() (*int64, bool)`

GetNLinkedTrialsOk returns a tuple with the NLinkedTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNLinkedTrials

`func (o *ConditionPublicationsRow) SetNLinkedTrials(v int64)`

SetNLinkedTrials sets NLinkedTrials field to given value.

### HasNLinkedTrials

`func (o *ConditionPublicationsRow) HasNLinkedTrials() bool`

HasNLinkedTrials returns a boolean if a field has been set.

### SetNLinkedTrialsNil

`func (o *ConditionPublicationsRow) SetNLinkedTrialsNil(b bool)`

 SetNLinkedTrialsNil sets the value for NLinkedTrials to be an explicit nil

### UnsetNLinkedTrials
`func (o *ConditionPublicationsRow) UnsetNLinkedTrials()`

UnsetNLinkedTrials ensures that no value is present for NLinkedTrials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


