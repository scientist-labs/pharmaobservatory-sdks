# UntitledLettersRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LetterId** | Pointer to **NullableString** | The letter id. | [optional] 
**LetterUrl** | Pointer to **NullableString** | The FDA page of the letter. | [optional] 
**FirmName** | Pointer to **NullableString** | The firm&#39;s name as filed with the regulator. | [optional] 
**FirmOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] 
**FirmCompanyId** | Pointer to **NullableString** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] 
**FirmCompanyName** | Pointer to **NullableString** | Display name of the firm&#39;s company. | [optional] 
**FirmGroupId** | Pointer to **NullableString** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**FirmGroupName** | Pointer to **NullableString** | Display name of the firm&#39;s group. | [optional] 
**IsOrgResolved** | Pointer to **NullableBool** | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional] 
**MatchTier** | Pointer to **NullableString** | How well the firm name was attached to an organization: high, medium or low name-and-geography matches, registry (the operating firm of a registered facility), ambiguous (several candidates, none chosen) or unresolved. Warning letters, untitled letters and debarments attach by firm name alone. | [optional] 
**ProductIssue** | Pointer to **NullableString** | The product and the issue the letter raises (for example the promotional claim). | [optional] 
**IssueDate** | Pointer to **NullableString** | The date the letter was issued. | [optional] 
**ResponseLetterUrl** | Pointer to **NullableString** | The firm&#39;s response letter on the FDA site, where one is posted. | [optional] 
**CloseoutLetterUrl** | Pointer to **NullableString** | The FDA close-out letter, where one is posted. | [optional] 

## Methods

### NewUntitledLettersRow

`func NewUntitledLettersRow() *UntitledLettersRow`

NewUntitledLettersRow instantiates a new UntitledLettersRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUntitledLettersRowWithDefaults

`func NewUntitledLettersRowWithDefaults() *UntitledLettersRow`

NewUntitledLettersRowWithDefaults instantiates a new UntitledLettersRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLetterId

`func (o *UntitledLettersRow) GetLetterId() string`

GetLetterId returns the LetterId field if non-nil, zero value otherwise.

### GetLetterIdOk

`func (o *UntitledLettersRow) GetLetterIdOk() (*string, bool)`

GetLetterIdOk returns a tuple with the LetterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetterId

`func (o *UntitledLettersRow) SetLetterId(v string)`

SetLetterId sets LetterId field to given value.

### HasLetterId

`func (o *UntitledLettersRow) HasLetterId() bool`

HasLetterId returns a boolean if a field has been set.

### SetLetterIdNil

`func (o *UntitledLettersRow) SetLetterIdNil(b bool)`

 SetLetterIdNil sets the value for LetterId to be an explicit nil

### UnsetLetterId
`func (o *UntitledLettersRow) UnsetLetterId()`

UnsetLetterId ensures that no value is present for LetterId, not even an explicit nil
### GetLetterUrl

`func (o *UntitledLettersRow) GetLetterUrl() string`

GetLetterUrl returns the LetterUrl field if non-nil, zero value otherwise.

### GetLetterUrlOk

`func (o *UntitledLettersRow) GetLetterUrlOk() (*string, bool)`

GetLetterUrlOk returns a tuple with the LetterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetterUrl

`func (o *UntitledLettersRow) SetLetterUrl(v string)`

SetLetterUrl sets LetterUrl field to given value.

### HasLetterUrl

`func (o *UntitledLettersRow) HasLetterUrl() bool`

HasLetterUrl returns a boolean if a field has been set.

### SetLetterUrlNil

`func (o *UntitledLettersRow) SetLetterUrlNil(b bool)`

 SetLetterUrlNil sets the value for LetterUrl to be an explicit nil

### UnsetLetterUrl
`func (o *UntitledLettersRow) UnsetLetterUrl()`

UnsetLetterUrl ensures that no value is present for LetterUrl, not even an explicit nil
### GetFirmName

`func (o *UntitledLettersRow) GetFirmName() string`

GetFirmName returns the FirmName field if non-nil, zero value otherwise.

### GetFirmNameOk

`func (o *UntitledLettersRow) GetFirmNameOk() (*string, bool)`

GetFirmNameOk returns a tuple with the FirmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmName

`func (o *UntitledLettersRow) SetFirmName(v string)`

SetFirmName sets FirmName field to given value.

### HasFirmName

`func (o *UntitledLettersRow) HasFirmName() bool`

HasFirmName returns a boolean if a field has been set.

### SetFirmNameNil

`func (o *UntitledLettersRow) SetFirmNameNil(b bool)`

 SetFirmNameNil sets the value for FirmName to be an explicit nil

### UnsetFirmName
`func (o *UntitledLettersRow) UnsetFirmName()`

UnsetFirmName ensures that no value is present for FirmName, not even an explicit nil
### GetFirmOrgId

`func (o *UntitledLettersRow) GetFirmOrgId() string`

GetFirmOrgId returns the FirmOrgId field if non-nil, zero value otherwise.

### GetFirmOrgIdOk

`func (o *UntitledLettersRow) GetFirmOrgIdOk() (*string, bool)`

GetFirmOrgIdOk returns a tuple with the FirmOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmOrgId

`func (o *UntitledLettersRow) SetFirmOrgId(v string)`

SetFirmOrgId sets FirmOrgId field to given value.

### HasFirmOrgId

`func (o *UntitledLettersRow) HasFirmOrgId() bool`

HasFirmOrgId returns a boolean if a field has been set.

### SetFirmOrgIdNil

`func (o *UntitledLettersRow) SetFirmOrgIdNil(b bool)`

 SetFirmOrgIdNil sets the value for FirmOrgId to be an explicit nil

### UnsetFirmOrgId
`func (o *UntitledLettersRow) UnsetFirmOrgId()`

UnsetFirmOrgId ensures that no value is present for FirmOrgId, not even an explicit nil
### GetFirmCompanyId

`func (o *UntitledLettersRow) GetFirmCompanyId() string`

GetFirmCompanyId returns the FirmCompanyId field if non-nil, zero value otherwise.

### GetFirmCompanyIdOk

`func (o *UntitledLettersRow) GetFirmCompanyIdOk() (*string, bool)`

GetFirmCompanyIdOk returns a tuple with the FirmCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyId

`func (o *UntitledLettersRow) SetFirmCompanyId(v string)`

SetFirmCompanyId sets FirmCompanyId field to given value.

### HasFirmCompanyId

`func (o *UntitledLettersRow) HasFirmCompanyId() bool`

HasFirmCompanyId returns a boolean if a field has been set.

### SetFirmCompanyIdNil

`func (o *UntitledLettersRow) SetFirmCompanyIdNil(b bool)`

 SetFirmCompanyIdNil sets the value for FirmCompanyId to be an explicit nil

### UnsetFirmCompanyId
`func (o *UntitledLettersRow) UnsetFirmCompanyId()`

UnsetFirmCompanyId ensures that no value is present for FirmCompanyId, not even an explicit nil
### GetFirmCompanyName

`func (o *UntitledLettersRow) GetFirmCompanyName() string`

GetFirmCompanyName returns the FirmCompanyName field if non-nil, zero value otherwise.

### GetFirmCompanyNameOk

`func (o *UntitledLettersRow) GetFirmCompanyNameOk() (*string, bool)`

GetFirmCompanyNameOk returns a tuple with the FirmCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyName

`func (o *UntitledLettersRow) SetFirmCompanyName(v string)`

SetFirmCompanyName sets FirmCompanyName field to given value.

### HasFirmCompanyName

`func (o *UntitledLettersRow) HasFirmCompanyName() bool`

HasFirmCompanyName returns a boolean if a field has been set.

### SetFirmCompanyNameNil

`func (o *UntitledLettersRow) SetFirmCompanyNameNil(b bool)`

 SetFirmCompanyNameNil sets the value for FirmCompanyName to be an explicit nil

### UnsetFirmCompanyName
`func (o *UntitledLettersRow) UnsetFirmCompanyName()`

UnsetFirmCompanyName ensures that no value is present for FirmCompanyName, not even an explicit nil
### GetFirmGroupId

`func (o *UntitledLettersRow) GetFirmGroupId() string`

GetFirmGroupId returns the FirmGroupId field if non-nil, zero value otherwise.

### GetFirmGroupIdOk

`func (o *UntitledLettersRow) GetFirmGroupIdOk() (*string, bool)`

GetFirmGroupIdOk returns a tuple with the FirmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupId

`func (o *UntitledLettersRow) SetFirmGroupId(v string)`

SetFirmGroupId sets FirmGroupId field to given value.

### HasFirmGroupId

`func (o *UntitledLettersRow) HasFirmGroupId() bool`

HasFirmGroupId returns a boolean if a field has been set.

### SetFirmGroupIdNil

`func (o *UntitledLettersRow) SetFirmGroupIdNil(b bool)`

 SetFirmGroupIdNil sets the value for FirmGroupId to be an explicit nil

### UnsetFirmGroupId
`func (o *UntitledLettersRow) UnsetFirmGroupId()`

UnsetFirmGroupId ensures that no value is present for FirmGroupId, not even an explicit nil
### GetFirmGroupName

`func (o *UntitledLettersRow) GetFirmGroupName() string`

GetFirmGroupName returns the FirmGroupName field if non-nil, zero value otherwise.

### GetFirmGroupNameOk

`func (o *UntitledLettersRow) GetFirmGroupNameOk() (*string, bool)`

GetFirmGroupNameOk returns a tuple with the FirmGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupName

`func (o *UntitledLettersRow) SetFirmGroupName(v string)`

SetFirmGroupName sets FirmGroupName field to given value.

### HasFirmGroupName

`func (o *UntitledLettersRow) HasFirmGroupName() bool`

HasFirmGroupName returns a boolean if a field has been set.

### SetFirmGroupNameNil

`func (o *UntitledLettersRow) SetFirmGroupNameNil(b bool)`

 SetFirmGroupNameNil sets the value for FirmGroupName to be an explicit nil

### UnsetFirmGroupName
`func (o *UntitledLettersRow) UnsetFirmGroupName()`

UnsetFirmGroupName ensures that no value is present for FirmGroupName, not even an explicit nil
### GetIsOrgResolved

`func (o *UntitledLettersRow) GetIsOrgResolved() bool`

GetIsOrgResolved returns the IsOrgResolved field if non-nil, zero value otherwise.

### GetIsOrgResolvedOk

`func (o *UntitledLettersRow) GetIsOrgResolvedOk() (*bool, bool)`

GetIsOrgResolvedOk returns a tuple with the IsOrgResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgResolved

`func (o *UntitledLettersRow) SetIsOrgResolved(v bool)`

SetIsOrgResolved sets IsOrgResolved field to given value.

### HasIsOrgResolved

`func (o *UntitledLettersRow) HasIsOrgResolved() bool`

HasIsOrgResolved returns a boolean if a field has been set.

### SetIsOrgResolvedNil

`func (o *UntitledLettersRow) SetIsOrgResolvedNil(b bool)`

 SetIsOrgResolvedNil sets the value for IsOrgResolved to be an explicit nil

### UnsetIsOrgResolved
`func (o *UntitledLettersRow) UnsetIsOrgResolved()`

UnsetIsOrgResolved ensures that no value is present for IsOrgResolved, not even an explicit nil
### GetMatchTier

`func (o *UntitledLettersRow) GetMatchTier() string`

GetMatchTier returns the MatchTier field if non-nil, zero value otherwise.

### GetMatchTierOk

`func (o *UntitledLettersRow) GetMatchTierOk() (*string, bool)`

GetMatchTierOk returns a tuple with the MatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTier

`func (o *UntitledLettersRow) SetMatchTier(v string)`

SetMatchTier sets MatchTier field to given value.

### HasMatchTier

`func (o *UntitledLettersRow) HasMatchTier() bool`

HasMatchTier returns a boolean if a field has been set.

### SetMatchTierNil

`func (o *UntitledLettersRow) SetMatchTierNil(b bool)`

 SetMatchTierNil sets the value for MatchTier to be an explicit nil

### UnsetMatchTier
`func (o *UntitledLettersRow) UnsetMatchTier()`

UnsetMatchTier ensures that no value is present for MatchTier, not even an explicit nil
### GetProductIssue

`func (o *UntitledLettersRow) GetProductIssue() string`

GetProductIssue returns the ProductIssue field if non-nil, zero value otherwise.

### GetProductIssueOk

`func (o *UntitledLettersRow) GetProductIssueOk() (*string, bool)`

GetProductIssueOk returns a tuple with the ProductIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIssue

`func (o *UntitledLettersRow) SetProductIssue(v string)`

SetProductIssue sets ProductIssue field to given value.

### HasProductIssue

`func (o *UntitledLettersRow) HasProductIssue() bool`

HasProductIssue returns a boolean if a field has been set.

### SetProductIssueNil

`func (o *UntitledLettersRow) SetProductIssueNil(b bool)`

 SetProductIssueNil sets the value for ProductIssue to be an explicit nil

### UnsetProductIssue
`func (o *UntitledLettersRow) UnsetProductIssue()`

UnsetProductIssue ensures that no value is present for ProductIssue, not even an explicit nil
### GetIssueDate

`func (o *UntitledLettersRow) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *UntitledLettersRow) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *UntitledLettersRow) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *UntitledLettersRow) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *UntitledLettersRow) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *UntitledLettersRow) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetResponseLetterUrl

`func (o *UntitledLettersRow) GetResponseLetterUrl() string`

GetResponseLetterUrl returns the ResponseLetterUrl field if non-nil, zero value otherwise.

### GetResponseLetterUrlOk

`func (o *UntitledLettersRow) GetResponseLetterUrlOk() (*string, bool)`

GetResponseLetterUrlOk returns a tuple with the ResponseLetterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterUrl

`func (o *UntitledLettersRow) SetResponseLetterUrl(v string)`

SetResponseLetterUrl sets ResponseLetterUrl field to given value.

### HasResponseLetterUrl

`func (o *UntitledLettersRow) HasResponseLetterUrl() bool`

HasResponseLetterUrl returns a boolean if a field has been set.

### SetResponseLetterUrlNil

`func (o *UntitledLettersRow) SetResponseLetterUrlNil(b bool)`

 SetResponseLetterUrlNil sets the value for ResponseLetterUrl to be an explicit nil

### UnsetResponseLetterUrl
`func (o *UntitledLettersRow) UnsetResponseLetterUrl()`

UnsetResponseLetterUrl ensures that no value is present for ResponseLetterUrl, not even an explicit nil
### GetCloseoutLetterUrl

`func (o *UntitledLettersRow) GetCloseoutLetterUrl() string`

GetCloseoutLetterUrl returns the CloseoutLetterUrl field if non-nil, zero value otherwise.

### GetCloseoutLetterUrlOk

`func (o *UntitledLettersRow) GetCloseoutLetterUrlOk() (*string, bool)`

GetCloseoutLetterUrlOk returns a tuple with the CloseoutLetterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseoutLetterUrl

`func (o *UntitledLettersRow) SetCloseoutLetterUrl(v string)`

SetCloseoutLetterUrl sets CloseoutLetterUrl field to given value.

### HasCloseoutLetterUrl

`func (o *UntitledLettersRow) HasCloseoutLetterUrl() bool`

HasCloseoutLetterUrl returns a boolean if a field has been set.

### SetCloseoutLetterUrlNil

`func (o *UntitledLettersRow) SetCloseoutLetterUrlNil(b bool)`

 SetCloseoutLetterUrlNil sets the value for CloseoutLetterUrl to be an explicit nil

### UnsetCloseoutLetterUrl
`func (o *UntitledLettersRow) UnsetCloseoutLetterUrl()`

UnsetCloseoutLetterUrl ensures that no value is present for CloseoutLetterUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


