# DebarmentsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebarmentId** | Pointer to **NullableString** | The debarment id. | [optional] 
**FirmName** | Pointer to **NullableString** | The firm&#39;s name as filed with the regulator (the debarred entity). | [optional] 
**FirmOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] 
**FirmCompanyId** | Pointer to **NullableString** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] 
**FirmCompanyName** | Pointer to **NullableString** | Display name of the firm&#39;s company. | [optional] 
**FirmGroupId** | Pointer to **NullableString** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**FirmGroupName** | Pointer to **NullableString** | Display name of the firm&#39;s group. | [optional] 
**IsOrgResolved** | Pointer to **NullableBool** | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional] 
**MatchTier** | Pointer to **NullableString** | How well the firm name was attached to an organization: high, medium or low name-and-geography matches, registry (the operating firm of a registered facility), ambiguous (several candidates, none chosen) or unresolved. Warning letters, untitled letters and debarments attach by firm name alone. | [optional] 
**EffectiveDate** | Pointer to **NullableString** | The date the debarment took effect. | [optional] 
**Term** | Pointer to **NullableString** | The debarment term (for example permanent or a number of years). | [optional] 
**FrCitation** | Pointer to **NullableString** | The Federal Register citation of the notice. | [optional] 
**FrDate** | Pointer to **NullableString** | The Federal Register publication date. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The FDA debarment list page. | [optional] 

## Methods

### NewDebarmentsRow

`func NewDebarmentsRow() *DebarmentsRow`

NewDebarmentsRow instantiates a new DebarmentsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebarmentsRowWithDefaults

`func NewDebarmentsRowWithDefaults() *DebarmentsRow`

NewDebarmentsRowWithDefaults instantiates a new DebarmentsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebarmentId

`func (o *DebarmentsRow) GetDebarmentId() string`

GetDebarmentId returns the DebarmentId field if non-nil, zero value otherwise.

### GetDebarmentIdOk

`func (o *DebarmentsRow) GetDebarmentIdOk() (*string, bool)`

GetDebarmentIdOk returns a tuple with the DebarmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebarmentId

`func (o *DebarmentsRow) SetDebarmentId(v string)`

SetDebarmentId sets DebarmentId field to given value.

### HasDebarmentId

`func (o *DebarmentsRow) HasDebarmentId() bool`

HasDebarmentId returns a boolean if a field has been set.

### SetDebarmentIdNil

`func (o *DebarmentsRow) SetDebarmentIdNil(b bool)`

 SetDebarmentIdNil sets the value for DebarmentId to be an explicit nil

### UnsetDebarmentId
`func (o *DebarmentsRow) UnsetDebarmentId()`

UnsetDebarmentId ensures that no value is present for DebarmentId, not even an explicit nil
### GetFirmName

`func (o *DebarmentsRow) GetFirmName() string`

GetFirmName returns the FirmName field if non-nil, zero value otherwise.

### GetFirmNameOk

`func (o *DebarmentsRow) GetFirmNameOk() (*string, bool)`

GetFirmNameOk returns a tuple with the FirmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmName

`func (o *DebarmentsRow) SetFirmName(v string)`

SetFirmName sets FirmName field to given value.

### HasFirmName

`func (o *DebarmentsRow) HasFirmName() bool`

HasFirmName returns a boolean if a field has been set.

### SetFirmNameNil

`func (o *DebarmentsRow) SetFirmNameNil(b bool)`

 SetFirmNameNil sets the value for FirmName to be an explicit nil

### UnsetFirmName
`func (o *DebarmentsRow) UnsetFirmName()`

UnsetFirmName ensures that no value is present for FirmName, not even an explicit nil
### GetFirmOrgId

`func (o *DebarmentsRow) GetFirmOrgId() string`

GetFirmOrgId returns the FirmOrgId field if non-nil, zero value otherwise.

### GetFirmOrgIdOk

`func (o *DebarmentsRow) GetFirmOrgIdOk() (*string, bool)`

GetFirmOrgIdOk returns a tuple with the FirmOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmOrgId

`func (o *DebarmentsRow) SetFirmOrgId(v string)`

SetFirmOrgId sets FirmOrgId field to given value.

### HasFirmOrgId

`func (o *DebarmentsRow) HasFirmOrgId() bool`

HasFirmOrgId returns a boolean if a field has been set.

### SetFirmOrgIdNil

`func (o *DebarmentsRow) SetFirmOrgIdNil(b bool)`

 SetFirmOrgIdNil sets the value for FirmOrgId to be an explicit nil

### UnsetFirmOrgId
`func (o *DebarmentsRow) UnsetFirmOrgId()`

UnsetFirmOrgId ensures that no value is present for FirmOrgId, not even an explicit nil
### GetFirmCompanyId

`func (o *DebarmentsRow) GetFirmCompanyId() string`

GetFirmCompanyId returns the FirmCompanyId field if non-nil, zero value otherwise.

### GetFirmCompanyIdOk

`func (o *DebarmentsRow) GetFirmCompanyIdOk() (*string, bool)`

GetFirmCompanyIdOk returns a tuple with the FirmCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyId

`func (o *DebarmentsRow) SetFirmCompanyId(v string)`

SetFirmCompanyId sets FirmCompanyId field to given value.

### HasFirmCompanyId

`func (o *DebarmentsRow) HasFirmCompanyId() bool`

HasFirmCompanyId returns a boolean if a field has been set.

### SetFirmCompanyIdNil

`func (o *DebarmentsRow) SetFirmCompanyIdNil(b bool)`

 SetFirmCompanyIdNil sets the value for FirmCompanyId to be an explicit nil

### UnsetFirmCompanyId
`func (o *DebarmentsRow) UnsetFirmCompanyId()`

UnsetFirmCompanyId ensures that no value is present for FirmCompanyId, not even an explicit nil
### GetFirmCompanyName

`func (o *DebarmentsRow) GetFirmCompanyName() string`

GetFirmCompanyName returns the FirmCompanyName field if non-nil, zero value otherwise.

### GetFirmCompanyNameOk

`func (o *DebarmentsRow) GetFirmCompanyNameOk() (*string, bool)`

GetFirmCompanyNameOk returns a tuple with the FirmCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyName

`func (o *DebarmentsRow) SetFirmCompanyName(v string)`

SetFirmCompanyName sets FirmCompanyName field to given value.

### HasFirmCompanyName

`func (o *DebarmentsRow) HasFirmCompanyName() bool`

HasFirmCompanyName returns a boolean if a field has been set.

### SetFirmCompanyNameNil

`func (o *DebarmentsRow) SetFirmCompanyNameNil(b bool)`

 SetFirmCompanyNameNil sets the value for FirmCompanyName to be an explicit nil

### UnsetFirmCompanyName
`func (o *DebarmentsRow) UnsetFirmCompanyName()`

UnsetFirmCompanyName ensures that no value is present for FirmCompanyName, not even an explicit nil
### GetFirmGroupId

`func (o *DebarmentsRow) GetFirmGroupId() string`

GetFirmGroupId returns the FirmGroupId field if non-nil, zero value otherwise.

### GetFirmGroupIdOk

`func (o *DebarmentsRow) GetFirmGroupIdOk() (*string, bool)`

GetFirmGroupIdOk returns a tuple with the FirmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupId

`func (o *DebarmentsRow) SetFirmGroupId(v string)`

SetFirmGroupId sets FirmGroupId field to given value.

### HasFirmGroupId

`func (o *DebarmentsRow) HasFirmGroupId() bool`

HasFirmGroupId returns a boolean if a field has been set.

### SetFirmGroupIdNil

`func (o *DebarmentsRow) SetFirmGroupIdNil(b bool)`

 SetFirmGroupIdNil sets the value for FirmGroupId to be an explicit nil

### UnsetFirmGroupId
`func (o *DebarmentsRow) UnsetFirmGroupId()`

UnsetFirmGroupId ensures that no value is present for FirmGroupId, not even an explicit nil
### GetFirmGroupName

`func (o *DebarmentsRow) GetFirmGroupName() string`

GetFirmGroupName returns the FirmGroupName field if non-nil, zero value otherwise.

### GetFirmGroupNameOk

`func (o *DebarmentsRow) GetFirmGroupNameOk() (*string, bool)`

GetFirmGroupNameOk returns a tuple with the FirmGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupName

`func (o *DebarmentsRow) SetFirmGroupName(v string)`

SetFirmGroupName sets FirmGroupName field to given value.

### HasFirmGroupName

`func (o *DebarmentsRow) HasFirmGroupName() bool`

HasFirmGroupName returns a boolean if a field has been set.

### SetFirmGroupNameNil

`func (o *DebarmentsRow) SetFirmGroupNameNil(b bool)`

 SetFirmGroupNameNil sets the value for FirmGroupName to be an explicit nil

### UnsetFirmGroupName
`func (o *DebarmentsRow) UnsetFirmGroupName()`

UnsetFirmGroupName ensures that no value is present for FirmGroupName, not even an explicit nil
### GetIsOrgResolved

`func (o *DebarmentsRow) GetIsOrgResolved() bool`

GetIsOrgResolved returns the IsOrgResolved field if non-nil, zero value otherwise.

### GetIsOrgResolvedOk

`func (o *DebarmentsRow) GetIsOrgResolvedOk() (*bool, bool)`

GetIsOrgResolvedOk returns a tuple with the IsOrgResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgResolved

`func (o *DebarmentsRow) SetIsOrgResolved(v bool)`

SetIsOrgResolved sets IsOrgResolved field to given value.

### HasIsOrgResolved

`func (o *DebarmentsRow) HasIsOrgResolved() bool`

HasIsOrgResolved returns a boolean if a field has been set.

### SetIsOrgResolvedNil

`func (o *DebarmentsRow) SetIsOrgResolvedNil(b bool)`

 SetIsOrgResolvedNil sets the value for IsOrgResolved to be an explicit nil

### UnsetIsOrgResolved
`func (o *DebarmentsRow) UnsetIsOrgResolved()`

UnsetIsOrgResolved ensures that no value is present for IsOrgResolved, not even an explicit nil
### GetMatchTier

`func (o *DebarmentsRow) GetMatchTier() string`

GetMatchTier returns the MatchTier field if non-nil, zero value otherwise.

### GetMatchTierOk

`func (o *DebarmentsRow) GetMatchTierOk() (*string, bool)`

GetMatchTierOk returns a tuple with the MatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTier

`func (o *DebarmentsRow) SetMatchTier(v string)`

SetMatchTier sets MatchTier field to given value.

### HasMatchTier

`func (o *DebarmentsRow) HasMatchTier() bool`

HasMatchTier returns a boolean if a field has been set.

### SetMatchTierNil

`func (o *DebarmentsRow) SetMatchTierNil(b bool)`

 SetMatchTierNil sets the value for MatchTier to be an explicit nil

### UnsetMatchTier
`func (o *DebarmentsRow) UnsetMatchTier()`

UnsetMatchTier ensures that no value is present for MatchTier, not even an explicit nil
### GetEffectiveDate

`func (o *DebarmentsRow) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *DebarmentsRow) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *DebarmentsRow) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *DebarmentsRow) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### SetEffectiveDateNil

`func (o *DebarmentsRow) SetEffectiveDateNil(b bool)`

 SetEffectiveDateNil sets the value for EffectiveDate to be an explicit nil

### UnsetEffectiveDate
`func (o *DebarmentsRow) UnsetEffectiveDate()`

UnsetEffectiveDate ensures that no value is present for EffectiveDate, not even an explicit nil
### GetTerm

`func (o *DebarmentsRow) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *DebarmentsRow) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *DebarmentsRow) SetTerm(v string)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *DebarmentsRow) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### SetTermNil

`func (o *DebarmentsRow) SetTermNil(b bool)`

 SetTermNil sets the value for Term to be an explicit nil

### UnsetTerm
`func (o *DebarmentsRow) UnsetTerm()`

UnsetTerm ensures that no value is present for Term, not even an explicit nil
### GetFrCitation

`func (o *DebarmentsRow) GetFrCitation() string`

GetFrCitation returns the FrCitation field if non-nil, zero value otherwise.

### GetFrCitationOk

`func (o *DebarmentsRow) GetFrCitationOk() (*string, bool)`

GetFrCitationOk returns a tuple with the FrCitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrCitation

`func (o *DebarmentsRow) SetFrCitation(v string)`

SetFrCitation sets FrCitation field to given value.

### HasFrCitation

`func (o *DebarmentsRow) HasFrCitation() bool`

HasFrCitation returns a boolean if a field has been set.

### SetFrCitationNil

`func (o *DebarmentsRow) SetFrCitationNil(b bool)`

 SetFrCitationNil sets the value for FrCitation to be an explicit nil

### UnsetFrCitation
`func (o *DebarmentsRow) UnsetFrCitation()`

UnsetFrCitation ensures that no value is present for FrCitation, not even an explicit nil
### GetFrDate

`func (o *DebarmentsRow) GetFrDate() string`

GetFrDate returns the FrDate field if non-nil, zero value otherwise.

### GetFrDateOk

`func (o *DebarmentsRow) GetFrDateOk() (*string, bool)`

GetFrDateOk returns a tuple with the FrDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrDate

`func (o *DebarmentsRow) SetFrDate(v string)`

SetFrDate sets FrDate field to given value.

### HasFrDate

`func (o *DebarmentsRow) HasFrDate() bool`

HasFrDate returns a boolean if a field has been set.

### SetFrDateNil

`func (o *DebarmentsRow) SetFrDateNil(b bool)`

 SetFrDateNil sets the value for FrDate to be an explicit nil

### UnsetFrDate
`func (o *DebarmentsRow) UnsetFrDate()`

UnsetFrDate ensures that no value is present for FrDate, not even an explicit nil
### GetSourceUrl

`func (o *DebarmentsRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *DebarmentsRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *DebarmentsRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *DebarmentsRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *DebarmentsRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *DebarmentsRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


