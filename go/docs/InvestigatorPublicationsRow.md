# InvestigatorPublicationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvestigatorPersonId** | Pointer to **NullableString** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] 
**Pmid** | Pointer to **NullableString** | The PubMed id of the publication. | [optional] 
**LinkBasis** | Pointer to **NullableString** | authored when the person is linked to the paper as an author; trial_reference when a trial the person is on registers the paper (not authorship). | [optional] 
**AuthorPosition** | Pointer to **NullableString** | first, middle or last for an authored row; null for a trial_reference row. | [optional] 
**AuthorMatchMethod** | Pointer to **NullableString** | How the author was matched to the person (orcid, name_on_linked_trial and so on); null for a trial_reference row. | [optional] 
**Citation** | Pointer to **NullableString** | The citation as a registered trial cites the paper (null when no trial cites it). | [optional] 
**Title** | Pointer to **NullableString** | The paper&#39;s title (null when no registered trial cites it). | [optional] 
**PubYear** | Pointer to **NullableInt32** | The publication year (null when no registered trial cites the paper). | [optional] 
**NTrialsCiting** | Pointer to **NullableInt32** | Registered trials that cite the paper (null when none does). | [optional] 

## Methods

### NewInvestigatorPublicationsRow

`func NewInvestigatorPublicationsRow() *InvestigatorPublicationsRow`

NewInvestigatorPublicationsRow instantiates a new InvestigatorPublicationsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigatorPublicationsRowWithDefaults

`func NewInvestigatorPublicationsRowWithDefaults() *InvestigatorPublicationsRow`

NewInvestigatorPublicationsRowWithDefaults instantiates a new InvestigatorPublicationsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestigatorPersonId

`func (o *InvestigatorPublicationsRow) GetInvestigatorPersonId() string`

GetInvestigatorPersonId returns the InvestigatorPersonId field if non-nil, zero value otherwise.

### GetInvestigatorPersonIdOk

`func (o *InvestigatorPublicationsRow) GetInvestigatorPersonIdOk() (*string, bool)`

GetInvestigatorPersonIdOk returns a tuple with the InvestigatorPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonId

`func (o *InvestigatorPublicationsRow) SetInvestigatorPersonId(v string)`

SetInvestigatorPersonId sets InvestigatorPersonId field to given value.

### HasInvestigatorPersonId

`func (o *InvestigatorPublicationsRow) HasInvestigatorPersonId() bool`

HasInvestigatorPersonId returns a boolean if a field has been set.

### SetInvestigatorPersonIdNil

`func (o *InvestigatorPublicationsRow) SetInvestigatorPersonIdNil(b bool)`

 SetInvestigatorPersonIdNil sets the value for InvestigatorPersonId to be an explicit nil

### UnsetInvestigatorPersonId
`func (o *InvestigatorPublicationsRow) UnsetInvestigatorPersonId()`

UnsetInvestigatorPersonId ensures that no value is present for InvestigatorPersonId, not even an explicit nil
### GetPmid

`func (o *InvestigatorPublicationsRow) GetPmid() string`

GetPmid returns the Pmid field if non-nil, zero value otherwise.

### GetPmidOk

`func (o *InvestigatorPublicationsRow) GetPmidOk() (*string, bool)`

GetPmidOk returns a tuple with the Pmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmid

`func (o *InvestigatorPublicationsRow) SetPmid(v string)`

SetPmid sets Pmid field to given value.

### HasPmid

`func (o *InvestigatorPublicationsRow) HasPmid() bool`

HasPmid returns a boolean if a field has been set.

### SetPmidNil

`func (o *InvestigatorPublicationsRow) SetPmidNil(b bool)`

 SetPmidNil sets the value for Pmid to be an explicit nil

### UnsetPmid
`func (o *InvestigatorPublicationsRow) UnsetPmid()`

UnsetPmid ensures that no value is present for Pmid, not even an explicit nil
### GetLinkBasis

`func (o *InvestigatorPublicationsRow) GetLinkBasis() string`

GetLinkBasis returns the LinkBasis field if non-nil, zero value otherwise.

### GetLinkBasisOk

`func (o *InvestigatorPublicationsRow) GetLinkBasisOk() (*string, bool)`

GetLinkBasisOk returns a tuple with the LinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkBasis

`func (o *InvestigatorPublicationsRow) SetLinkBasis(v string)`

SetLinkBasis sets LinkBasis field to given value.

### HasLinkBasis

`func (o *InvestigatorPublicationsRow) HasLinkBasis() bool`

HasLinkBasis returns a boolean if a field has been set.

### SetLinkBasisNil

`func (o *InvestigatorPublicationsRow) SetLinkBasisNil(b bool)`

 SetLinkBasisNil sets the value for LinkBasis to be an explicit nil

### UnsetLinkBasis
`func (o *InvestigatorPublicationsRow) UnsetLinkBasis()`

UnsetLinkBasis ensures that no value is present for LinkBasis, not even an explicit nil
### GetAuthorPosition

`func (o *InvestigatorPublicationsRow) GetAuthorPosition() string`

GetAuthorPosition returns the AuthorPosition field if non-nil, zero value otherwise.

### GetAuthorPositionOk

`func (o *InvestigatorPublicationsRow) GetAuthorPositionOk() (*string, bool)`

GetAuthorPositionOk returns a tuple with the AuthorPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorPosition

`func (o *InvestigatorPublicationsRow) SetAuthorPosition(v string)`

SetAuthorPosition sets AuthorPosition field to given value.

### HasAuthorPosition

`func (o *InvestigatorPublicationsRow) HasAuthorPosition() bool`

HasAuthorPosition returns a boolean if a field has been set.

### SetAuthorPositionNil

`func (o *InvestigatorPublicationsRow) SetAuthorPositionNil(b bool)`

 SetAuthorPositionNil sets the value for AuthorPosition to be an explicit nil

### UnsetAuthorPosition
`func (o *InvestigatorPublicationsRow) UnsetAuthorPosition()`

UnsetAuthorPosition ensures that no value is present for AuthorPosition, not even an explicit nil
### GetAuthorMatchMethod

`func (o *InvestigatorPublicationsRow) GetAuthorMatchMethod() string`

GetAuthorMatchMethod returns the AuthorMatchMethod field if non-nil, zero value otherwise.

### GetAuthorMatchMethodOk

`func (o *InvestigatorPublicationsRow) GetAuthorMatchMethodOk() (*string, bool)`

GetAuthorMatchMethodOk returns a tuple with the AuthorMatchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorMatchMethod

`func (o *InvestigatorPublicationsRow) SetAuthorMatchMethod(v string)`

SetAuthorMatchMethod sets AuthorMatchMethod field to given value.

### HasAuthorMatchMethod

`func (o *InvestigatorPublicationsRow) HasAuthorMatchMethod() bool`

HasAuthorMatchMethod returns a boolean if a field has been set.

### SetAuthorMatchMethodNil

`func (o *InvestigatorPublicationsRow) SetAuthorMatchMethodNil(b bool)`

 SetAuthorMatchMethodNil sets the value for AuthorMatchMethod to be an explicit nil

### UnsetAuthorMatchMethod
`func (o *InvestigatorPublicationsRow) UnsetAuthorMatchMethod()`

UnsetAuthorMatchMethod ensures that no value is present for AuthorMatchMethod, not even an explicit nil
### GetCitation

`func (o *InvestigatorPublicationsRow) GetCitation() string`

GetCitation returns the Citation field if non-nil, zero value otherwise.

### GetCitationOk

`func (o *InvestigatorPublicationsRow) GetCitationOk() (*string, bool)`

GetCitationOk returns a tuple with the Citation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitation

`func (o *InvestigatorPublicationsRow) SetCitation(v string)`

SetCitation sets Citation field to given value.

### HasCitation

`func (o *InvestigatorPublicationsRow) HasCitation() bool`

HasCitation returns a boolean if a field has been set.

### SetCitationNil

`func (o *InvestigatorPublicationsRow) SetCitationNil(b bool)`

 SetCitationNil sets the value for Citation to be an explicit nil

### UnsetCitation
`func (o *InvestigatorPublicationsRow) UnsetCitation()`

UnsetCitation ensures that no value is present for Citation, not even an explicit nil
### GetTitle

`func (o *InvestigatorPublicationsRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InvestigatorPublicationsRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InvestigatorPublicationsRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InvestigatorPublicationsRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *InvestigatorPublicationsRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *InvestigatorPublicationsRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPubYear

`func (o *InvestigatorPublicationsRow) GetPubYear() int32`

GetPubYear returns the PubYear field if non-nil, zero value otherwise.

### GetPubYearOk

`func (o *InvestigatorPublicationsRow) GetPubYearOk() (*int32, bool)`

GetPubYearOk returns a tuple with the PubYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubYear

`func (o *InvestigatorPublicationsRow) SetPubYear(v int32)`

SetPubYear sets PubYear field to given value.

### HasPubYear

`func (o *InvestigatorPublicationsRow) HasPubYear() bool`

HasPubYear returns a boolean if a field has been set.

### SetPubYearNil

`func (o *InvestigatorPublicationsRow) SetPubYearNil(b bool)`

 SetPubYearNil sets the value for PubYear to be an explicit nil

### UnsetPubYear
`func (o *InvestigatorPublicationsRow) UnsetPubYear()`

UnsetPubYear ensures that no value is present for PubYear, not even an explicit nil
### GetNTrialsCiting

`func (o *InvestigatorPublicationsRow) GetNTrialsCiting() int32`

GetNTrialsCiting returns the NTrialsCiting field if non-nil, zero value otherwise.

### GetNTrialsCitingOk

`func (o *InvestigatorPublicationsRow) GetNTrialsCitingOk() (*int32, bool)`

GetNTrialsCitingOk returns a tuple with the NTrialsCiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsCiting

`func (o *InvestigatorPublicationsRow) SetNTrialsCiting(v int32)`

SetNTrialsCiting sets NTrialsCiting field to given value.

### HasNTrialsCiting

`func (o *InvestigatorPublicationsRow) HasNTrialsCiting() bool`

HasNTrialsCiting returns a boolean if a field has been set.

### SetNTrialsCitingNil

`func (o *InvestigatorPublicationsRow) SetNTrialsCitingNil(b bool)`

 SetNTrialsCitingNil sets the value for NTrialsCiting to be an explicit nil

### UnsetNTrialsCiting
`func (o *InvestigatorPublicationsRow) UnsetNTrialsCiting()`

UnsetNTrialsCiting ensures that no value is present for NTrialsCiting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


