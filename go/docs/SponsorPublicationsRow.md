# SponsorPublicationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** | Company level: the company the row is about; the entry key for the companies object type. | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**Pmid** | Pointer to **NullableString** | The PubMed id. | [optional] 
**Citation** | Pointer to **NullableString** | The citation text of the paper as the trial registry cites it. | [optional] 
**Title** | Pointer to **NullableString** | The article title. | [optional] 
**PubYear** | Pointer to **NullableInt32** | The publication year. | [optional] 
**JournalNlmId** | Pointer to **NullableString** | The journal&#39;s NLM (National Library of Medicine) catalog id; null when the paper&#39;s journal is unknown. | [optional] 
**JournalTitle** | Pointer to **NullableString** | The journal title; null when the journal is not in the journal directory. | [optional] 
**PublicationTrials** | Pointer to **NullableInt32** | How many trials of any sponsor cite the paper. | [optional] 
**NTrials** | Pointer to **NullableInt32** | Distinct trials of this company that cite the paper. | [optional] 
**NLeadTrials** | Pointer to **NullableInt32** | Of those, the trials the company leads. | [optional] 

## Methods

### NewSponsorPublicationsRow

`func NewSponsorPublicationsRow() *SponsorPublicationsRow`

NewSponsorPublicationsRow instantiates a new SponsorPublicationsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSponsorPublicationsRowWithDefaults

`func NewSponsorPublicationsRowWithDefaults() *SponsorPublicationsRow`

NewSponsorPublicationsRowWithDefaults instantiates a new SponsorPublicationsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *SponsorPublicationsRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *SponsorPublicationsRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *SponsorPublicationsRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *SponsorPublicationsRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *SponsorPublicationsRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *SponsorPublicationsRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *SponsorPublicationsRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *SponsorPublicationsRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *SponsorPublicationsRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *SponsorPublicationsRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *SponsorPublicationsRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *SponsorPublicationsRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *SponsorPublicationsRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SponsorPublicationsRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SponsorPublicationsRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SponsorPublicationsRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *SponsorPublicationsRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *SponsorPublicationsRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *SponsorPublicationsRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *SponsorPublicationsRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *SponsorPublicationsRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *SponsorPublicationsRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *SponsorPublicationsRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *SponsorPublicationsRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetPmid

`func (o *SponsorPublicationsRow) GetPmid() string`

GetPmid returns the Pmid field if non-nil, zero value otherwise.

### GetPmidOk

`func (o *SponsorPublicationsRow) GetPmidOk() (*string, bool)`

GetPmidOk returns a tuple with the Pmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmid

`func (o *SponsorPublicationsRow) SetPmid(v string)`

SetPmid sets Pmid field to given value.

### HasPmid

`func (o *SponsorPublicationsRow) HasPmid() bool`

HasPmid returns a boolean if a field has been set.

### SetPmidNil

`func (o *SponsorPublicationsRow) SetPmidNil(b bool)`

 SetPmidNil sets the value for Pmid to be an explicit nil

### UnsetPmid
`func (o *SponsorPublicationsRow) UnsetPmid()`

UnsetPmid ensures that no value is present for Pmid, not even an explicit nil
### GetCitation

`func (o *SponsorPublicationsRow) GetCitation() string`

GetCitation returns the Citation field if non-nil, zero value otherwise.

### GetCitationOk

`func (o *SponsorPublicationsRow) GetCitationOk() (*string, bool)`

GetCitationOk returns a tuple with the Citation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitation

`func (o *SponsorPublicationsRow) SetCitation(v string)`

SetCitation sets Citation field to given value.

### HasCitation

`func (o *SponsorPublicationsRow) HasCitation() bool`

HasCitation returns a boolean if a field has been set.

### SetCitationNil

`func (o *SponsorPublicationsRow) SetCitationNil(b bool)`

 SetCitationNil sets the value for Citation to be an explicit nil

### UnsetCitation
`func (o *SponsorPublicationsRow) UnsetCitation()`

UnsetCitation ensures that no value is present for Citation, not even an explicit nil
### GetTitle

`func (o *SponsorPublicationsRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SponsorPublicationsRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SponsorPublicationsRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SponsorPublicationsRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SponsorPublicationsRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SponsorPublicationsRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPubYear

`func (o *SponsorPublicationsRow) GetPubYear() int32`

GetPubYear returns the PubYear field if non-nil, zero value otherwise.

### GetPubYearOk

`func (o *SponsorPublicationsRow) GetPubYearOk() (*int32, bool)`

GetPubYearOk returns a tuple with the PubYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubYear

`func (o *SponsorPublicationsRow) SetPubYear(v int32)`

SetPubYear sets PubYear field to given value.

### HasPubYear

`func (o *SponsorPublicationsRow) HasPubYear() bool`

HasPubYear returns a boolean if a field has been set.

### SetPubYearNil

`func (o *SponsorPublicationsRow) SetPubYearNil(b bool)`

 SetPubYearNil sets the value for PubYear to be an explicit nil

### UnsetPubYear
`func (o *SponsorPublicationsRow) UnsetPubYear()`

UnsetPubYear ensures that no value is present for PubYear, not even an explicit nil
### GetJournalNlmId

`func (o *SponsorPublicationsRow) GetJournalNlmId() string`

GetJournalNlmId returns the JournalNlmId field if non-nil, zero value otherwise.

### GetJournalNlmIdOk

`func (o *SponsorPublicationsRow) GetJournalNlmIdOk() (*string, bool)`

GetJournalNlmIdOk returns a tuple with the JournalNlmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalNlmId

`func (o *SponsorPublicationsRow) SetJournalNlmId(v string)`

SetJournalNlmId sets JournalNlmId field to given value.

### HasJournalNlmId

`func (o *SponsorPublicationsRow) HasJournalNlmId() bool`

HasJournalNlmId returns a boolean if a field has been set.

### SetJournalNlmIdNil

`func (o *SponsorPublicationsRow) SetJournalNlmIdNil(b bool)`

 SetJournalNlmIdNil sets the value for JournalNlmId to be an explicit nil

### UnsetJournalNlmId
`func (o *SponsorPublicationsRow) UnsetJournalNlmId()`

UnsetJournalNlmId ensures that no value is present for JournalNlmId, not even an explicit nil
### GetJournalTitle

`func (o *SponsorPublicationsRow) GetJournalTitle() string`

GetJournalTitle returns the JournalTitle field if non-nil, zero value otherwise.

### GetJournalTitleOk

`func (o *SponsorPublicationsRow) GetJournalTitleOk() (*string, bool)`

GetJournalTitleOk returns a tuple with the JournalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalTitle

`func (o *SponsorPublicationsRow) SetJournalTitle(v string)`

SetJournalTitle sets JournalTitle field to given value.

### HasJournalTitle

`func (o *SponsorPublicationsRow) HasJournalTitle() bool`

HasJournalTitle returns a boolean if a field has been set.

### SetJournalTitleNil

`func (o *SponsorPublicationsRow) SetJournalTitleNil(b bool)`

 SetJournalTitleNil sets the value for JournalTitle to be an explicit nil

### UnsetJournalTitle
`func (o *SponsorPublicationsRow) UnsetJournalTitle()`

UnsetJournalTitle ensures that no value is present for JournalTitle, not even an explicit nil
### GetPublicationTrials

`func (o *SponsorPublicationsRow) GetPublicationTrials() int32`

GetPublicationTrials returns the PublicationTrials field if non-nil, zero value otherwise.

### GetPublicationTrialsOk

`func (o *SponsorPublicationsRow) GetPublicationTrialsOk() (*int32, bool)`

GetPublicationTrialsOk returns a tuple with the PublicationTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationTrials

`func (o *SponsorPublicationsRow) SetPublicationTrials(v int32)`

SetPublicationTrials sets PublicationTrials field to given value.

### HasPublicationTrials

`func (o *SponsorPublicationsRow) HasPublicationTrials() bool`

HasPublicationTrials returns a boolean if a field has been set.

### SetPublicationTrialsNil

`func (o *SponsorPublicationsRow) SetPublicationTrialsNil(b bool)`

 SetPublicationTrialsNil sets the value for PublicationTrials to be an explicit nil

### UnsetPublicationTrials
`func (o *SponsorPublicationsRow) UnsetPublicationTrials()`

UnsetPublicationTrials ensures that no value is present for PublicationTrials, not even an explicit nil
### GetNTrials

`func (o *SponsorPublicationsRow) GetNTrials() int32`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *SponsorPublicationsRow) GetNTrialsOk() (*int32, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *SponsorPublicationsRow) SetNTrials(v int32)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *SponsorPublicationsRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *SponsorPublicationsRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *SponsorPublicationsRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNLeadTrials

`func (o *SponsorPublicationsRow) GetNLeadTrials() int32`

GetNLeadTrials returns the NLeadTrials field if non-nil, zero value otherwise.

### GetNLeadTrialsOk

`func (o *SponsorPublicationsRow) GetNLeadTrialsOk() (*int32, bool)`

GetNLeadTrialsOk returns a tuple with the NLeadTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNLeadTrials

`func (o *SponsorPublicationsRow) SetNLeadTrials(v int32)`

SetNLeadTrials sets NLeadTrials field to given value.

### HasNLeadTrials

`func (o *SponsorPublicationsRow) HasNLeadTrials() bool`

HasNLeadTrials returns a boolean if a field has been set.

### SetNLeadTrialsNil

`func (o *SponsorPublicationsRow) SetNLeadTrialsNil(b bool)`

 SetNLeadTrialsNil sets the value for NLeadTrials to be an explicit nil

### UnsetNLeadTrials
`func (o *SponsorPublicationsRow) UnsetNLeadTrials()`

UnsetNLeadTrials ensures that no value is present for NLeadTrials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


