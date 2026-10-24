# InvestigatorDebarmentsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvestigatorPersonId** | Pointer to **NullableString** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] 
**DebarmentId** | Pointer to **NullableString** | The debarment record id. | [optional] 
**DebarredName** | Pointer to **NullableString** | The debarred individual&#39;s name as the FDA list spells it. | [optional] 
**EffectiveDate** | Pointer to **NullableString** | The debarment&#39;s effective date. | [optional] 
**Term** | Pointer to **NullableString** | The debarment term as the FDA states it. | [optional] 
**FrCitation** | Pointer to **NullableString** | The Federal Register citation. | [optional] 
**SourceUrl** | Pointer to **NullableString** | A link to the source record. | [optional] 
**MatchTier** | Pointer to **NullableString** | How the debarred individual was matched to the investigator: low is a bare-name match that reaches exactly one person (the only tier published; an ambiguous or unresolved name has no person and no row). | [optional] 

## Methods

### NewInvestigatorDebarmentsRow

`func NewInvestigatorDebarmentsRow() *InvestigatorDebarmentsRow`

NewInvestigatorDebarmentsRow instantiates a new InvestigatorDebarmentsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigatorDebarmentsRowWithDefaults

`func NewInvestigatorDebarmentsRowWithDefaults() *InvestigatorDebarmentsRow`

NewInvestigatorDebarmentsRowWithDefaults instantiates a new InvestigatorDebarmentsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestigatorPersonId

`func (o *InvestigatorDebarmentsRow) GetInvestigatorPersonId() string`

GetInvestigatorPersonId returns the InvestigatorPersonId field if non-nil, zero value otherwise.

### GetInvestigatorPersonIdOk

`func (o *InvestigatorDebarmentsRow) GetInvestigatorPersonIdOk() (*string, bool)`

GetInvestigatorPersonIdOk returns a tuple with the InvestigatorPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonId

`func (o *InvestigatorDebarmentsRow) SetInvestigatorPersonId(v string)`

SetInvestigatorPersonId sets InvestigatorPersonId field to given value.

### HasInvestigatorPersonId

`func (o *InvestigatorDebarmentsRow) HasInvestigatorPersonId() bool`

HasInvestigatorPersonId returns a boolean if a field has been set.

### SetInvestigatorPersonIdNil

`func (o *InvestigatorDebarmentsRow) SetInvestigatorPersonIdNil(b bool)`

 SetInvestigatorPersonIdNil sets the value for InvestigatorPersonId to be an explicit nil

### UnsetInvestigatorPersonId
`func (o *InvestigatorDebarmentsRow) UnsetInvestigatorPersonId()`

UnsetInvestigatorPersonId ensures that no value is present for InvestigatorPersonId, not even an explicit nil
### GetDebarmentId

`func (o *InvestigatorDebarmentsRow) GetDebarmentId() string`

GetDebarmentId returns the DebarmentId field if non-nil, zero value otherwise.

### GetDebarmentIdOk

`func (o *InvestigatorDebarmentsRow) GetDebarmentIdOk() (*string, bool)`

GetDebarmentIdOk returns a tuple with the DebarmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebarmentId

`func (o *InvestigatorDebarmentsRow) SetDebarmentId(v string)`

SetDebarmentId sets DebarmentId field to given value.

### HasDebarmentId

`func (o *InvestigatorDebarmentsRow) HasDebarmentId() bool`

HasDebarmentId returns a boolean if a field has been set.

### SetDebarmentIdNil

`func (o *InvestigatorDebarmentsRow) SetDebarmentIdNil(b bool)`

 SetDebarmentIdNil sets the value for DebarmentId to be an explicit nil

### UnsetDebarmentId
`func (o *InvestigatorDebarmentsRow) UnsetDebarmentId()`

UnsetDebarmentId ensures that no value is present for DebarmentId, not even an explicit nil
### GetDebarredName

`func (o *InvestigatorDebarmentsRow) GetDebarredName() string`

GetDebarredName returns the DebarredName field if non-nil, zero value otherwise.

### GetDebarredNameOk

`func (o *InvestigatorDebarmentsRow) GetDebarredNameOk() (*string, bool)`

GetDebarredNameOk returns a tuple with the DebarredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebarredName

`func (o *InvestigatorDebarmentsRow) SetDebarredName(v string)`

SetDebarredName sets DebarredName field to given value.

### HasDebarredName

`func (o *InvestigatorDebarmentsRow) HasDebarredName() bool`

HasDebarredName returns a boolean if a field has been set.

### SetDebarredNameNil

`func (o *InvestigatorDebarmentsRow) SetDebarredNameNil(b bool)`

 SetDebarredNameNil sets the value for DebarredName to be an explicit nil

### UnsetDebarredName
`func (o *InvestigatorDebarmentsRow) UnsetDebarredName()`

UnsetDebarredName ensures that no value is present for DebarredName, not even an explicit nil
### GetEffectiveDate

`func (o *InvestigatorDebarmentsRow) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *InvestigatorDebarmentsRow) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *InvestigatorDebarmentsRow) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *InvestigatorDebarmentsRow) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### SetEffectiveDateNil

`func (o *InvestigatorDebarmentsRow) SetEffectiveDateNil(b bool)`

 SetEffectiveDateNil sets the value for EffectiveDate to be an explicit nil

### UnsetEffectiveDate
`func (o *InvestigatorDebarmentsRow) UnsetEffectiveDate()`

UnsetEffectiveDate ensures that no value is present for EffectiveDate, not even an explicit nil
### GetTerm

`func (o *InvestigatorDebarmentsRow) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *InvestigatorDebarmentsRow) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *InvestigatorDebarmentsRow) SetTerm(v string)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *InvestigatorDebarmentsRow) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### SetTermNil

`func (o *InvestigatorDebarmentsRow) SetTermNil(b bool)`

 SetTermNil sets the value for Term to be an explicit nil

### UnsetTerm
`func (o *InvestigatorDebarmentsRow) UnsetTerm()`

UnsetTerm ensures that no value is present for Term, not even an explicit nil
### GetFrCitation

`func (o *InvestigatorDebarmentsRow) GetFrCitation() string`

GetFrCitation returns the FrCitation field if non-nil, zero value otherwise.

### GetFrCitationOk

`func (o *InvestigatorDebarmentsRow) GetFrCitationOk() (*string, bool)`

GetFrCitationOk returns a tuple with the FrCitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrCitation

`func (o *InvestigatorDebarmentsRow) SetFrCitation(v string)`

SetFrCitation sets FrCitation field to given value.

### HasFrCitation

`func (o *InvestigatorDebarmentsRow) HasFrCitation() bool`

HasFrCitation returns a boolean if a field has been set.

### SetFrCitationNil

`func (o *InvestigatorDebarmentsRow) SetFrCitationNil(b bool)`

 SetFrCitationNil sets the value for FrCitation to be an explicit nil

### UnsetFrCitation
`func (o *InvestigatorDebarmentsRow) UnsetFrCitation()`

UnsetFrCitation ensures that no value is present for FrCitation, not even an explicit nil
### GetSourceUrl

`func (o *InvestigatorDebarmentsRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *InvestigatorDebarmentsRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *InvestigatorDebarmentsRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *InvestigatorDebarmentsRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *InvestigatorDebarmentsRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *InvestigatorDebarmentsRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetMatchTier

`func (o *InvestigatorDebarmentsRow) GetMatchTier() string`

GetMatchTier returns the MatchTier field if non-nil, zero value otherwise.

### GetMatchTierOk

`func (o *InvestigatorDebarmentsRow) GetMatchTierOk() (*string, bool)`

GetMatchTierOk returns a tuple with the MatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTier

`func (o *InvestigatorDebarmentsRow) SetMatchTier(v string)`

SetMatchTier sets MatchTier field to given value.

### HasMatchTier

`func (o *InvestigatorDebarmentsRow) HasMatchTier() bool`

HasMatchTier returns a boolean if a field has been set.

### SetMatchTierNil

`func (o *InvestigatorDebarmentsRow) SetMatchTierNil(b bool)`

 SetMatchTierNil sets the value for MatchTier to be an explicit nil

### UnsetMatchTier
`func (o *InvestigatorDebarmentsRow) UnsetMatchTier()`

UnsetMatchTier ensures that no value is present for MatchTier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


