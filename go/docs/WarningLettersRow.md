# WarningLettersRow

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
**IssuingOffice** | Pointer to **NullableString** | The FDA office that issued the letter (a center or a district office). | [optional] 
**Subject** | Pointer to **NullableString** | The letter&#39;s subject line. | [optional] 
**Excerpt** | Pointer to **NullableString** | An excerpt of the letter text. | [optional] 
**IssueDate** | Pointer to **NullableString** | The date the letter was issued. | [optional] 
**PostedDate** | Pointer to **NullableString** | The date the letter was posted on the FDA site. | [optional] 
**ResponseLetterUrl** | Pointer to **NullableString** | The firm&#39;s response letter on the FDA site, where one is posted. | [optional] 
**CloseoutLetterUrl** | Pointer to **NullableString** | The FDA close-out letter, where one is posted. | [optional] 

## Methods

### NewWarningLettersRow

`func NewWarningLettersRow() *WarningLettersRow`

NewWarningLettersRow instantiates a new WarningLettersRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarningLettersRowWithDefaults

`func NewWarningLettersRowWithDefaults() *WarningLettersRow`

NewWarningLettersRowWithDefaults instantiates a new WarningLettersRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLetterId

`func (o *WarningLettersRow) GetLetterId() string`

GetLetterId returns the LetterId field if non-nil, zero value otherwise.

### GetLetterIdOk

`func (o *WarningLettersRow) GetLetterIdOk() (*string, bool)`

GetLetterIdOk returns a tuple with the LetterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetterId

`func (o *WarningLettersRow) SetLetterId(v string)`

SetLetterId sets LetterId field to given value.

### HasLetterId

`func (o *WarningLettersRow) HasLetterId() bool`

HasLetterId returns a boolean if a field has been set.

### SetLetterIdNil

`func (o *WarningLettersRow) SetLetterIdNil(b bool)`

 SetLetterIdNil sets the value for LetterId to be an explicit nil

### UnsetLetterId
`func (o *WarningLettersRow) UnsetLetterId()`

UnsetLetterId ensures that no value is present for LetterId, not even an explicit nil
### GetLetterUrl

`func (o *WarningLettersRow) GetLetterUrl() string`

GetLetterUrl returns the LetterUrl field if non-nil, zero value otherwise.

### GetLetterUrlOk

`func (o *WarningLettersRow) GetLetterUrlOk() (*string, bool)`

GetLetterUrlOk returns a tuple with the LetterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetterUrl

`func (o *WarningLettersRow) SetLetterUrl(v string)`

SetLetterUrl sets LetterUrl field to given value.

### HasLetterUrl

`func (o *WarningLettersRow) HasLetterUrl() bool`

HasLetterUrl returns a boolean if a field has been set.

### SetLetterUrlNil

`func (o *WarningLettersRow) SetLetterUrlNil(b bool)`

 SetLetterUrlNil sets the value for LetterUrl to be an explicit nil

### UnsetLetterUrl
`func (o *WarningLettersRow) UnsetLetterUrl()`

UnsetLetterUrl ensures that no value is present for LetterUrl, not even an explicit nil
### GetFirmName

`func (o *WarningLettersRow) GetFirmName() string`

GetFirmName returns the FirmName field if non-nil, zero value otherwise.

### GetFirmNameOk

`func (o *WarningLettersRow) GetFirmNameOk() (*string, bool)`

GetFirmNameOk returns a tuple with the FirmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmName

`func (o *WarningLettersRow) SetFirmName(v string)`

SetFirmName sets FirmName field to given value.

### HasFirmName

`func (o *WarningLettersRow) HasFirmName() bool`

HasFirmName returns a boolean if a field has been set.

### SetFirmNameNil

`func (o *WarningLettersRow) SetFirmNameNil(b bool)`

 SetFirmNameNil sets the value for FirmName to be an explicit nil

### UnsetFirmName
`func (o *WarningLettersRow) UnsetFirmName()`

UnsetFirmName ensures that no value is present for FirmName, not even an explicit nil
### GetFirmOrgId

`func (o *WarningLettersRow) GetFirmOrgId() string`

GetFirmOrgId returns the FirmOrgId field if non-nil, zero value otherwise.

### GetFirmOrgIdOk

`func (o *WarningLettersRow) GetFirmOrgIdOk() (*string, bool)`

GetFirmOrgIdOk returns a tuple with the FirmOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmOrgId

`func (o *WarningLettersRow) SetFirmOrgId(v string)`

SetFirmOrgId sets FirmOrgId field to given value.

### HasFirmOrgId

`func (o *WarningLettersRow) HasFirmOrgId() bool`

HasFirmOrgId returns a boolean if a field has been set.

### SetFirmOrgIdNil

`func (o *WarningLettersRow) SetFirmOrgIdNil(b bool)`

 SetFirmOrgIdNil sets the value for FirmOrgId to be an explicit nil

### UnsetFirmOrgId
`func (o *WarningLettersRow) UnsetFirmOrgId()`

UnsetFirmOrgId ensures that no value is present for FirmOrgId, not even an explicit nil
### GetFirmCompanyId

`func (o *WarningLettersRow) GetFirmCompanyId() string`

GetFirmCompanyId returns the FirmCompanyId field if non-nil, zero value otherwise.

### GetFirmCompanyIdOk

`func (o *WarningLettersRow) GetFirmCompanyIdOk() (*string, bool)`

GetFirmCompanyIdOk returns a tuple with the FirmCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyId

`func (o *WarningLettersRow) SetFirmCompanyId(v string)`

SetFirmCompanyId sets FirmCompanyId field to given value.

### HasFirmCompanyId

`func (o *WarningLettersRow) HasFirmCompanyId() bool`

HasFirmCompanyId returns a boolean if a field has been set.

### SetFirmCompanyIdNil

`func (o *WarningLettersRow) SetFirmCompanyIdNil(b bool)`

 SetFirmCompanyIdNil sets the value for FirmCompanyId to be an explicit nil

### UnsetFirmCompanyId
`func (o *WarningLettersRow) UnsetFirmCompanyId()`

UnsetFirmCompanyId ensures that no value is present for FirmCompanyId, not even an explicit nil
### GetFirmCompanyName

`func (o *WarningLettersRow) GetFirmCompanyName() string`

GetFirmCompanyName returns the FirmCompanyName field if non-nil, zero value otherwise.

### GetFirmCompanyNameOk

`func (o *WarningLettersRow) GetFirmCompanyNameOk() (*string, bool)`

GetFirmCompanyNameOk returns a tuple with the FirmCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyName

`func (o *WarningLettersRow) SetFirmCompanyName(v string)`

SetFirmCompanyName sets FirmCompanyName field to given value.

### HasFirmCompanyName

`func (o *WarningLettersRow) HasFirmCompanyName() bool`

HasFirmCompanyName returns a boolean if a field has been set.

### SetFirmCompanyNameNil

`func (o *WarningLettersRow) SetFirmCompanyNameNil(b bool)`

 SetFirmCompanyNameNil sets the value for FirmCompanyName to be an explicit nil

### UnsetFirmCompanyName
`func (o *WarningLettersRow) UnsetFirmCompanyName()`

UnsetFirmCompanyName ensures that no value is present for FirmCompanyName, not even an explicit nil
### GetFirmGroupId

`func (o *WarningLettersRow) GetFirmGroupId() string`

GetFirmGroupId returns the FirmGroupId field if non-nil, zero value otherwise.

### GetFirmGroupIdOk

`func (o *WarningLettersRow) GetFirmGroupIdOk() (*string, bool)`

GetFirmGroupIdOk returns a tuple with the FirmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupId

`func (o *WarningLettersRow) SetFirmGroupId(v string)`

SetFirmGroupId sets FirmGroupId field to given value.

### HasFirmGroupId

`func (o *WarningLettersRow) HasFirmGroupId() bool`

HasFirmGroupId returns a boolean if a field has been set.

### SetFirmGroupIdNil

`func (o *WarningLettersRow) SetFirmGroupIdNil(b bool)`

 SetFirmGroupIdNil sets the value for FirmGroupId to be an explicit nil

### UnsetFirmGroupId
`func (o *WarningLettersRow) UnsetFirmGroupId()`

UnsetFirmGroupId ensures that no value is present for FirmGroupId, not even an explicit nil
### GetFirmGroupName

`func (o *WarningLettersRow) GetFirmGroupName() string`

GetFirmGroupName returns the FirmGroupName field if non-nil, zero value otherwise.

### GetFirmGroupNameOk

`func (o *WarningLettersRow) GetFirmGroupNameOk() (*string, bool)`

GetFirmGroupNameOk returns a tuple with the FirmGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupName

`func (o *WarningLettersRow) SetFirmGroupName(v string)`

SetFirmGroupName sets FirmGroupName field to given value.

### HasFirmGroupName

`func (o *WarningLettersRow) HasFirmGroupName() bool`

HasFirmGroupName returns a boolean if a field has been set.

### SetFirmGroupNameNil

`func (o *WarningLettersRow) SetFirmGroupNameNil(b bool)`

 SetFirmGroupNameNil sets the value for FirmGroupName to be an explicit nil

### UnsetFirmGroupName
`func (o *WarningLettersRow) UnsetFirmGroupName()`

UnsetFirmGroupName ensures that no value is present for FirmGroupName, not even an explicit nil
### GetIsOrgResolved

`func (o *WarningLettersRow) GetIsOrgResolved() bool`

GetIsOrgResolved returns the IsOrgResolved field if non-nil, zero value otherwise.

### GetIsOrgResolvedOk

`func (o *WarningLettersRow) GetIsOrgResolvedOk() (*bool, bool)`

GetIsOrgResolvedOk returns a tuple with the IsOrgResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgResolved

`func (o *WarningLettersRow) SetIsOrgResolved(v bool)`

SetIsOrgResolved sets IsOrgResolved field to given value.

### HasIsOrgResolved

`func (o *WarningLettersRow) HasIsOrgResolved() bool`

HasIsOrgResolved returns a boolean if a field has been set.

### SetIsOrgResolvedNil

`func (o *WarningLettersRow) SetIsOrgResolvedNil(b bool)`

 SetIsOrgResolvedNil sets the value for IsOrgResolved to be an explicit nil

### UnsetIsOrgResolved
`func (o *WarningLettersRow) UnsetIsOrgResolved()`

UnsetIsOrgResolved ensures that no value is present for IsOrgResolved, not even an explicit nil
### GetMatchTier

`func (o *WarningLettersRow) GetMatchTier() string`

GetMatchTier returns the MatchTier field if non-nil, zero value otherwise.

### GetMatchTierOk

`func (o *WarningLettersRow) GetMatchTierOk() (*string, bool)`

GetMatchTierOk returns a tuple with the MatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTier

`func (o *WarningLettersRow) SetMatchTier(v string)`

SetMatchTier sets MatchTier field to given value.

### HasMatchTier

`func (o *WarningLettersRow) HasMatchTier() bool`

HasMatchTier returns a boolean if a field has been set.

### SetMatchTierNil

`func (o *WarningLettersRow) SetMatchTierNil(b bool)`

 SetMatchTierNil sets the value for MatchTier to be an explicit nil

### UnsetMatchTier
`func (o *WarningLettersRow) UnsetMatchTier()`

UnsetMatchTier ensures that no value is present for MatchTier, not even an explicit nil
### GetIssuingOffice

`func (o *WarningLettersRow) GetIssuingOffice() string`

GetIssuingOffice returns the IssuingOffice field if non-nil, zero value otherwise.

### GetIssuingOfficeOk

`func (o *WarningLettersRow) GetIssuingOfficeOk() (*string, bool)`

GetIssuingOfficeOk returns a tuple with the IssuingOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingOffice

`func (o *WarningLettersRow) SetIssuingOffice(v string)`

SetIssuingOffice sets IssuingOffice field to given value.

### HasIssuingOffice

`func (o *WarningLettersRow) HasIssuingOffice() bool`

HasIssuingOffice returns a boolean if a field has been set.

### SetIssuingOfficeNil

`func (o *WarningLettersRow) SetIssuingOfficeNil(b bool)`

 SetIssuingOfficeNil sets the value for IssuingOffice to be an explicit nil

### UnsetIssuingOffice
`func (o *WarningLettersRow) UnsetIssuingOffice()`

UnsetIssuingOffice ensures that no value is present for IssuingOffice, not even an explicit nil
### GetSubject

`func (o *WarningLettersRow) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *WarningLettersRow) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *WarningLettersRow) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *WarningLettersRow) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *WarningLettersRow) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *WarningLettersRow) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetExcerpt

`func (o *WarningLettersRow) GetExcerpt() string`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *WarningLettersRow) GetExcerptOk() (*string, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *WarningLettersRow) SetExcerpt(v string)`

SetExcerpt sets Excerpt field to given value.

### HasExcerpt

`func (o *WarningLettersRow) HasExcerpt() bool`

HasExcerpt returns a boolean if a field has been set.

### SetExcerptNil

`func (o *WarningLettersRow) SetExcerptNil(b bool)`

 SetExcerptNil sets the value for Excerpt to be an explicit nil

### UnsetExcerpt
`func (o *WarningLettersRow) UnsetExcerpt()`

UnsetExcerpt ensures that no value is present for Excerpt, not even an explicit nil
### GetIssueDate

`func (o *WarningLettersRow) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *WarningLettersRow) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *WarningLettersRow) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *WarningLettersRow) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *WarningLettersRow) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *WarningLettersRow) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetPostedDate

`func (o *WarningLettersRow) GetPostedDate() string`

GetPostedDate returns the PostedDate field if non-nil, zero value otherwise.

### GetPostedDateOk

`func (o *WarningLettersRow) GetPostedDateOk() (*string, bool)`

GetPostedDateOk returns a tuple with the PostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedDate

`func (o *WarningLettersRow) SetPostedDate(v string)`

SetPostedDate sets PostedDate field to given value.

### HasPostedDate

`func (o *WarningLettersRow) HasPostedDate() bool`

HasPostedDate returns a boolean if a field has been set.

### SetPostedDateNil

`func (o *WarningLettersRow) SetPostedDateNil(b bool)`

 SetPostedDateNil sets the value for PostedDate to be an explicit nil

### UnsetPostedDate
`func (o *WarningLettersRow) UnsetPostedDate()`

UnsetPostedDate ensures that no value is present for PostedDate, not even an explicit nil
### GetResponseLetterUrl

`func (o *WarningLettersRow) GetResponseLetterUrl() string`

GetResponseLetterUrl returns the ResponseLetterUrl field if non-nil, zero value otherwise.

### GetResponseLetterUrlOk

`func (o *WarningLettersRow) GetResponseLetterUrlOk() (*string, bool)`

GetResponseLetterUrlOk returns a tuple with the ResponseLetterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterUrl

`func (o *WarningLettersRow) SetResponseLetterUrl(v string)`

SetResponseLetterUrl sets ResponseLetterUrl field to given value.

### HasResponseLetterUrl

`func (o *WarningLettersRow) HasResponseLetterUrl() bool`

HasResponseLetterUrl returns a boolean if a field has been set.

### SetResponseLetterUrlNil

`func (o *WarningLettersRow) SetResponseLetterUrlNil(b bool)`

 SetResponseLetterUrlNil sets the value for ResponseLetterUrl to be an explicit nil

### UnsetResponseLetterUrl
`func (o *WarningLettersRow) UnsetResponseLetterUrl()`

UnsetResponseLetterUrl ensures that no value is present for ResponseLetterUrl, not even an explicit nil
### GetCloseoutLetterUrl

`func (o *WarningLettersRow) GetCloseoutLetterUrl() string`

GetCloseoutLetterUrl returns the CloseoutLetterUrl field if non-nil, zero value otherwise.

### GetCloseoutLetterUrlOk

`func (o *WarningLettersRow) GetCloseoutLetterUrlOk() (*string, bool)`

GetCloseoutLetterUrlOk returns a tuple with the CloseoutLetterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseoutLetterUrl

`func (o *WarningLettersRow) SetCloseoutLetterUrl(v string)`

SetCloseoutLetterUrl sets CloseoutLetterUrl field to given value.

### HasCloseoutLetterUrl

`func (o *WarningLettersRow) HasCloseoutLetterUrl() bool`

HasCloseoutLetterUrl returns a boolean if a field has been set.

### SetCloseoutLetterUrlNil

`func (o *WarningLettersRow) SetCloseoutLetterUrlNil(b bool)`

 SetCloseoutLetterUrlNil sets the value for CloseoutLetterUrl to be an explicit nil

### UnsetCloseoutLetterUrl
`func (o *WarningLettersRow) UnsetCloseoutLetterUrl()`

UnsetCloseoutLetterUrl ensures that no value is present for CloseoutLetterUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


