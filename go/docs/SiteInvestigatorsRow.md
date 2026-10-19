# SiteInvestigatorsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **NullableString** | The site key: one facility spelling from the registry. | [optional] 
**InvestigatorPersonName** | Pointer to **NullableString** | The investigator&#39;s name as filed. | [optional] 
**InvestigatorPersonId** | Pointer to **NullableString** | The investigator&#39;s person key. | [optional] 
**InvestigatorAffiliation** | Pointer to **NullableString** | The affiliation as filed. | [optional] 
**NTrials** | Pointer to **NullableInt64** | Trials at the site that name the investigator. | [optional] 
**IsPrincipalInvestigator** | Pointer to **NullableBool** | True when the investigator is named principal investigator on a trial at the site. | [optional] 
**SiteRank** | Pointer to **NullableInt64** | Rank within the site: principal investigators first, then trial count. | [optional] 

## Methods

### NewSiteInvestigatorsRow

`func NewSiteInvestigatorsRow() *SiteInvestigatorsRow`

NewSiteInvestigatorsRow instantiates a new SiteInvestigatorsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteInvestigatorsRowWithDefaults

`func NewSiteInvestigatorsRowWithDefaults() *SiteInvestigatorsRow`

NewSiteInvestigatorsRowWithDefaults instantiates a new SiteInvestigatorsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *SiteInvestigatorsRow) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SiteInvestigatorsRow) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SiteInvestigatorsRow) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SiteInvestigatorsRow) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *SiteInvestigatorsRow) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *SiteInvestigatorsRow) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetInvestigatorPersonName

`func (o *SiteInvestigatorsRow) GetInvestigatorPersonName() string`

GetInvestigatorPersonName returns the InvestigatorPersonName field if non-nil, zero value otherwise.

### GetInvestigatorPersonNameOk

`func (o *SiteInvestigatorsRow) GetInvestigatorPersonNameOk() (*string, bool)`

GetInvestigatorPersonNameOk returns a tuple with the InvestigatorPersonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonName

`func (o *SiteInvestigatorsRow) SetInvestigatorPersonName(v string)`

SetInvestigatorPersonName sets InvestigatorPersonName field to given value.

### HasInvestigatorPersonName

`func (o *SiteInvestigatorsRow) HasInvestigatorPersonName() bool`

HasInvestigatorPersonName returns a boolean if a field has been set.

### SetInvestigatorPersonNameNil

`func (o *SiteInvestigatorsRow) SetInvestigatorPersonNameNil(b bool)`

 SetInvestigatorPersonNameNil sets the value for InvestigatorPersonName to be an explicit nil

### UnsetInvestigatorPersonName
`func (o *SiteInvestigatorsRow) UnsetInvestigatorPersonName()`

UnsetInvestigatorPersonName ensures that no value is present for InvestigatorPersonName, not even an explicit nil
### GetInvestigatorPersonId

`func (o *SiteInvestigatorsRow) GetInvestigatorPersonId() string`

GetInvestigatorPersonId returns the InvestigatorPersonId field if non-nil, zero value otherwise.

### GetInvestigatorPersonIdOk

`func (o *SiteInvestigatorsRow) GetInvestigatorPersonIdOk() (*string, bool)`

GetInvestigatorPersonIdOk returns a tuple with the InvestigatorPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonId

`func (o *SiteInvestigatorsRow) SetInvestigatorPersonId(v string)`

SetInvestigatorPersonId sets InvestigatorPersonId field to given value.

### HasInvestigatorPersonId

`func (o *SiteInvestigatorsRow) HasInvestigatorPersonId() bool`

HasInvestigatorPersonId returns a boolean if a field has been set.

### SetInvestigatorPersonIdNil

`func (o *SiteInvestigatorsRow) SetInvestigatorPersonIdNil(b bool)`

 SetInvestigatorPersonIdNil sets the value for InvestigatorPersonId to be an explicit nil

### UnsetInvestigatorPersonId
`func (o *SiteInvestigatorsRow) UnsetInvestigatorPersonId()`

UnsetInvestigatorPersonId ensures that no value is present for InvestigatorPersonId, not even an explicit nil
### GetInvestigatorAffiliation

`func (o *SiteInvestigatorsRow) GetInvestigatorAffiliation() string`

GetInvestigatorAffiliation returns the InvestigatorAffiliation field if non-nil, zero value otherwise.

### GetInvestigatorAffiliationOk

`func (o *SiteInvestigatorsRow) GetInvestigatorAffiliationOk() (*string, bool)`

GetInvestigatorAffiliationOk returns a tuple with the InvestigatorAffiliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorAffiliation

`func (o *SiteInvestigatorsRow) SetInvestigatorAffiliation(v string)`

SetInvestigatorAffiliation sets InvestigatorAffiliation field to given value.

### HasInvestigatorAffiliation

`func (o *SiteInvestigatorsRow) HasInvestigatorAffiliation() bool`

HasInvestigatorAffiliation returns a boolean if a field has been set.

### SetInvestigatorAffiliationNil

`func (o *SiteInvestigatorsRow) SetInvestigatorAffiliationNil(b bool)`

 SetInvestigatorAffiliationNil sets the value for InvestigatorAffiliation to be an explicit nil

### UnsetInvestigatorAffiliation
`func (o *SiteInvestigatorsRow) UnsetInvestigatorAffiliation()`

UnsetInvestigatorAffiliation ensures that no value is present for InvestigatorAffiliation, not even an explicit nil
### GetNTrials

`func (o *SiteInvestigatorsRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *SiteInvestigatorsRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *SiteInvestigatorsRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *SiteInvestigatorsRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *SiteInvestigatorsRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *SiteInvestigatorsRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetIsPrincipalInvestigator

`func (o *SiteInvestigatorsRow) GetIsPrincipalInvestigator() bool`

GetIsPrincipalInvestigator returns the IsPrincipalInvestigator field if non-nil, zero value otherwise.

### GetIsPrincipalInvestigatorOk

`func (o *SiteInvestigatorsRow) GetIsPrincipalInvestigatorOk() (*bool, bool)`

GetIsPrincipalInvestigatorOk returns a tuple with the IsPrincipalInvestigator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrincipalInvestigator

`func (o *SiteInvestigatorsRow) SetIsPrincipalInvestigator(v bool)`

SetIsPrincipalInvestigator sets IsPrincipalInvestigator field to given value.

### HasIsPrincipalInvestigator

`func (o *SiteInvestigatorsRow) HasIsPrincipalInvestigator() bool`

HasIsPrincipalInvestigator returns a boolean if a field has been set.

### SetIsPrincipalInvestigatorNil

`func (o *SiteInvestigatorsRow) SetIsPrincipalInvestigatorNil(b bool)`

 SetIsPrincipalInvestigatorNil sets the value for IsPrincipalInvestigator to be an explicit nil

### UnsetIsPrincipalInvestigator
`func (o *SiteInvestigatorsRow) UnsetIsPrincipalInvestigator()`

UnsetIsPrincipalInvestigator ensures that no value is present for IsPrincipalInvestigator, not even an explicit nil
### GetSiteRank

`func (o *SiteInvestigatorsRow) GetSiteRank() int64`

GetSiteRank returns the SiteRank field if non-nil, zero value otherwise.

### GetSiteRankOk

`func (o *SiteInvestigatorsRow) GetSiteRankOk() (*int64, bool)`

GetSiteRankOk returns a tuple with the SiteRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteRank

`func (o *SiteInvestigatorsRow) SetSiteRank(v int64)`

SetSiteRank sets SiteRank field to given value.

### HasSiteRank

`func (o *SiteInvestigatorsRow) HasSiteRank() bool`

HasSiteRank returns a boolean if a field has been set.

### SetSiteRankNil

`func (o *SiteInvestigatorsRow) SetSiteRankNil(b bool)`

 SetSiteRankNil sets the value for SiteRank to be an explicit nil

### UnsetSiteRank
`func (o *SiteInvestigatorsRow) UnsetSiteRank()`

UnsetSiteRank ensures that no value is present for SiteRank, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


