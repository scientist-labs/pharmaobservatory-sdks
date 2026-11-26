# TargetLiteraturePeopleRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **NullableString** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. | [optional] 
**InvestigatorPersonId** | Pointer to **NullableString** | The author&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] 
**PersonName** | Pointer to **NullableString** | The person&#39;s name, resolved across the person&#39;s merged name fragments. | [optional] 
**PersonAffiliation** | Pointer to **NullableString** | The person&#39;s affiliation as the person record holds it (null when not recorded). | [optional] 
**Orcid** | Pointer to **NullableString** | The person&#39;s ORCID iD (Open Researcher and Contributor ID) when the person record carries one (null when none). | [optional] 
**NPapers** | Pointer to **NullableInt32** | Distinct papers on which a ledger decision links a byline to the person. | [optional] 
**NTrialLinkedPapers** | Pointer to **NullableInt32** | The part of n_papers a registered clinical trial cites. | [optional] 
**LatestPubYear** | Pointer to **NullableInt32** | The newest publication year among those papers. | [optional] 
**NFirstAuthor** | Pointer to **NullableInt32** | The part of n_papers on which the person&#39;s linked byline is the first author position. | [optional] 
**NLastAuthor** | Pointer to **NullableInt32** | The part of n_papers on which the person&#39;s linked byline is the last author position. | [optional] 

## Methods

### NewTargetLiteraturePeopleRow

`func NewTargetLiteraturePeopleRow() *TargetLiteraturePeopleRow`

NewTargetLiteraturePeopleRow instantiates a new TargetLiteraturePeopleRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetLiteraturePeopleRowWithDefaults

`func NewTargetLiteraturePeopleRowWithDefaults() *TargetLiteraturePeopleRow`

NewTargetLiteraturePeopleRowWithDefaults instantiates a new TargetLiteraturePeopleRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *TargetLiteraturePeopleRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TargetLiteraturePeopleRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TargetLiteraturePeopleRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *TargetLiteraturePeopleRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *TargetLiteraturePeopleRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *TargetLiteraturePeopleRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetInvestigatorPersonId

`func (o *TargetLiteraturePeopleRow) GetInvestigatorPersonId() string`

GetInvestigatorPersonId returns the InvestigatorPersonId field if non-nil, zero value otherwise.

### GetInvestigatorPersonIdOk

`func (o *TargetLiteraturePeopleRow) GetInvestigatorPersonIdOk() (*string, bool)`

GetInvestigatorPersonIdOk returns a tuple with the InvestigatorPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonId

`func (o *TargetLiteraturePeopleRow) SetInvestigatorPersonId(v string)`

SetInvestigatorPersonId sets InvestigatorPersonId field to given value.

### HasInvestigatorPersonId

`func (o *TargetLiteraturePeopleRow) HasInvestigatorPersonId() bool`

HasInvestigatorPersonId returns a boolean if a field has been set.

### SetInvestigatorPersonIdNil

`func (o *TargetLiteraturePeopleRow) SetInvestigatorPersonIdNil(b bool)`

 SetInvestigatorPersonIdNil sets the value for InvestigatorPersonId to be an explicit nil

### UnsetInvestigatorPersonId
`func (o *TargetLiteraturePeopleRow) UnsetInvestigatorPersonId()`

UnsetInvestigatorPersonId ensures that no value is present for InvestigatorPersonId, not even an explicit nil
### GetPersonName

`func (o *TargetLiteraturePeopleRow) GetPersonName() string`

GetPersonName returns the PersonName field if non-nil, zero value otherwise.

### GetPersonNameOk

`func (o *TargetLiteraturePeopleRow) GetPersonNameOk() (*string, bool)`

GetPersonNameOk returns a tuple with the PersonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonName

`func (o *TargetLiteraturePeopleRow) SetPersonName(v string)`

SetPersonName sets PersonName field to given value.

### HasPersonName

`func (o *TargetLiteraturePeopleRow) HasPersonName() bool`

HasPersonName returns a boolean if a field has been set.

### SetPersonNameNil

`func (o *TargetLiteraturePeopleRow) SetPersonNameNil(b bool)`

 SetPersonNameNil sets the value for PersonName to be an explicit nil

### UnsetPersonName
`func (o *TargetLiteraturePeopleRow) UnsetPersonName()`

UnsetPersonName ensures that no value is present for PersonName, not even an explicit nil
### GetPersonAffiliation

`func (o *TargetLiteraturePeopleRow) GetPersonAffiliation() string`

GetPersonAffiliation returns the PersonAffiliation field if non-nil, zero value otherwise.

### GetPersonAffiliationOk

`func (o *TargetLiteraturePeopleRow) GetPersonAffiliationOk() (*string, bool)`

GetPersonAffiliationOk returns a tuple with the PersonAffiliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonAffiliation

`func (o *TargetLiteraturePeopleRow) SetPersonAffiliation(v string)`

SetPersonAffiliation sets PersonAffiliation field to given value.

### HasPersonAffiliation

`func (o *TargetLiteraturePeopleRow) HasPersonAffiliation() bool`

HasPersonAffiliation returns a boolean if a field has been set.

### SetPersonAffiliationNil

`func (o *TargetLiteraturePeopleRow) SetPersonAffiliationNil(b bool)`

 SetPersonAffiliationNil sets the value for PersonAffiliation to be an explicit nil

### UnsetPersonAffiliation
`func (o *TargetLiteraturePeopleRow) UnsetPersonAffiliation()`

UnsetPersonAffiliation ensures that no value is present for PersonAffiliation, not even an explicit nil
### GetOrcid

`func (o *TargetLiteraturePeopleRow) GetOrcid() string`

GetOrcid returns the Orcid field if non-nil, zero value otherwise.

### GetOrcidOk

`func (o *TargetLiteraturePeopleRow) GetOrcidOk() (*string, bool)`

GetOrcidOk returns a tuple with the Orcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrcid

`func (o *TargetLiteraturePeopleRow) SetOrcid(v string)`

SetOrcid sets Orcid field to given value.

### HasOrcid

`func (o *TargetLiteraturePeopleRow) HasOrcid() bool`

HasOrcid returns a boolean if a field has been set.

### SetOrcidNil

`func (o *TargetLiteraturePeopleRow) SetOrcidNil(b bool)`

 SetOrcidNil sets the value for Orcid to be an explicit nil

### UnsetOrcid
`func (o *TargetLiteraturePeopleRow) UnsetOrcid()`

UnsetOrcid ensures that no value is present for Orcid, not even an explicit nil
### GetNPapers

`func (o *TargetLiteraturePeopleRow) GetNPapers() int32`

GetNPapers returns the NPapers field if non-nil, zero value otherwise.

### GetNPapersOk

`func (o *TargetLiteraturePeopleRow) GetNPapersOk() (*int32, bool)`

GetNPapersOk returns a tuple with the NPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPapers

`func (o *TargetLiteraturePeopleRow) SetNPapers(v int32)`

SetNPapers sets NPapers field to given value.

### HasNPapers

`func (o *TargetLiteraturePeopleRow) HasNPapers() bool`

HasNPapers returns a boolean if a field has been set.

### SetNPapersNil

`func (o *TargetLiteraturePeopleRow) SetNPapersNil(b bool)`

 SetNPapersNil sets the value for NPapers to be an explicit nil

### UnsetNPapers
`func (o *TargetLiteraturePeopleRow) UnsetNPapers()`

UnsetNPapers ensures that no value is present for NPapers, not even an explicit nil
### GetNTrialLinkedPapers

`func (o *TargetLiteraturePeopleRow) GetNTrialLinkedPapers() int32`

GetNTrialLinkedPapers returns the NTrialLinkedPapers field if non-nil, zero value otherwise.

### GetNTrialLinkedPapersOk

`func (o *TargetLiteraturePeopleRow) GetNTrialLinkedPapersOk() (*int32, bool)`

GetNTrialLinkedPapersOk returns a tuple with the NTrialLinkedPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialLinkedPapers

`func (o *TargetLiteraturePeopleRow) SetNTrialLinkedPapers(v int32)`

SetNTrialLinkedPapers sets NTrialLinkedPapers field to given value.

### HasNTrialLinkedPapers

`func (o *TargetLiteraturePeopleRow) HasNTrialLinkedPapers() bool`

HasNTrialLinkedPapers returns a boolean if a field has been set.

### SetNTrialLinkedPapersNil

`func (o *TargetLiteraturePeopleRow) SetNTrialLinkedPapersNil(b bool)`

 SetNTrialLinkedPapersNil sets the value for NTrialLinkedPapers to be an explicit nil

### UnsetNTrialLinkedPapers
`func (o *TargetLiteraturePeopleRow) UnsetNTrialLinkedPapers()`

UnsetNTrialLinkedPapers ensures that no value is present for NTrialLinkedPapers, not even an explicit nil
### GetLatestPubYear

`func (o *TargetLiteraturePeopleRow) GetLatestPubYear() int32`

GetLatestPubYear returns the LatestPubYear field if non-nil, zero value otherwise.

### GetLatestPubYearOk

`func (o *TargetLiteraturePeopleRow) GetLatestPubYearOk() (*int32, bool)`

GetLatestPubYearOk returns a tuple with the LatestPubYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestPubYear

`func (o *TargetLiteraturePeopleRow) SetLatestPubYear(v int32)`

SetLatestPubYear sets LatestPubYear field to given value.

### HasLatestPubYear

`func (o *TargetLiteraturePeopleRow) HasLatestPubYear() bool`

HasLatestPubYear returns a boolean if a field has been set.

### SetLatestPubYearNil

`func (o *TargetLiteraturePeopleRow) SetLatestPubYearNil(b bool)`

 SetLatestPubYearNil sets the value for LatestPubYear to be an explicit nil

### UnsetLatestPubYear
`func (o *TargetLiteraturePeopleRow) UnsetLatestPubYear()`

UnsetLatestPubYear ensures that no value is present for LatestPubYear, not even an explicit nil
### GetNFirstAuthor

`func (o *TargetLiteraturePeopleRow) GetNFirstAuthor() int32`

GetNFirstAuthor returns the NFirstAuthor field if non-nil, zero value otherwise.

### GetNFirstAuthorOk

`func (o *TargetLiteraturePeopleRow) GetNFirstAuthorOk() (*int32, bool)`

GetNFirstAuthorOk returns a tuple with the NFirstAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFirstAuthor

`func (o *TargetLiteraturePeopleRow) SetNFirstAuthor(v int32)`

SetNFirstAuthor sets NFirstAuthor field to given value.

### HasNFirstAuthor

`func (o *TargetLiteraturePeopleRow) HasNFirstAuthor() bool`

HasNFirstAuthor returns a boolean if a field has been set.

### SetNFirstAuthorNil

`func (o *TargetLiteraturePeopleRow) SetNFirstAuthorNil(b bool)`

 SetNFirstAuthorNil sets the value for NFirstAuthor to be an explicit nil

### UnsetNFirstAuthor
`func (o *TargetLiteraturePeopleRow) UnsetNFirstAuthor()`

UnsetNFirstAuthor ensures that no value is present for NFirstAuthor, not even an explicit nil
### GetNLastAuthor

`func (o *TargetLiteraturePeopleRow) GetNLastAuthor() int32`

GetNLastAuthor returns the NLastAuthor field if non-nil, zero value otherwise.

### GetNLastAuthorOk

`func (o *TargetLiteraturePeopleRow) GetNLastAuthorOk() (*int32, bool)`

GetNLastAuthorOk returns a tuple with the NLastAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNLastAuthor

`func (o *TargetLiteraturePeopleRow) SetNLastAuthor(v int32)`

SetNLastAuthor sets NLastAuthor field to given value.

### HasNLastAuthor

`func (o *TargetLiteraturePeopleRow) HasNLastAuthor() bool`

HasNLastAuthor returns a boolean if a field has been set.

### SetNLastAuthorNil

`func (o *TargetLiteraturePeopleRow) SetNLastAuthorNil(b bool)`

 SetNLastAuthorNil sets the value for NLastAuthor to be an explicit nil

### UnsetNLastAuthor
`func (o *TargetLiteraturePeopleRow) UnsetNLastAuthor()`

UnsetNLastAuthor ensures that no value is present for NLastAuthor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


