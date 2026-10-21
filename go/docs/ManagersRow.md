# ManagersRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagerId** | Pointer to **NullableString** | The manager key, the id the web page uses: edgar_cik:&lt;10-digit filer CIK&gt;, or party:sec-cik:&lt;ref&gt; for a filer whose CIK is malformed. The same value as holder_identity_key on institutional_holdings. | [optional] 
**Name** | Pointer to **NullableString** | The manager&#39;s name as filed (the longest variant), falling back to the manager key. | [optional] 
**NameVariantCount** | Pointer to **NullableInt32** | How many spellings of the name the filings carry. | [optional] 
**InvestorId** | Pointer to **NullableString** | The manager&#39;s investor key where the link resolves; null when the manager resolves to no investor or its rows disagree (see investor_link_conflict). | [optional] 
**InvestorLinkConflict** | Pointer to **NullableBool** | True when the manager&#39;s own holding rows resolve to different investors, so no investor is linked. | [optional] 
**OrgId** | Pointer to **NullableString** | Legal-entity level: the manager&#39;s organization record when all its rows agree on one (null otherwise). | [optional] 
**CompanyId** | Pointer to **NullableString** | Company level: the manager&#39;s company (null when unresolved). | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the manager&#39;s company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the manager&#39;s company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the manager&#39;s group. | [optional] 
**PositionCount** | Pointer to **NullableInt32** | Current positions the manager reports, one per issuer security. | [optional] 
**IssuerCount** | Pointer to **NullableInt32** | Distinct issuer securities (CUSIPs) across those positions. | [optional] 
**ValuedPositionCount** | Pointer to **NullableInt32** | Of those positions, the ones with a reported value (a position whose filing scale was ambiguous has none). | [optional] 
**FirstObservedDate** | Pointer to **NullableString** | The earliest filing date a current position was first observed. | [optional] 
**LastFilingDate** | Pointer to **NullableString** | The latest filing date among the manager&#39;s positions. | [optional] 

## Methods

### NewManagersRow

`func NewManagersRow() *ManagersRow`

NewManagersRow instantiates a new ManagersRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagersRowWithDefaults

`func NewManagersRowWithDefaults() *ManagersRow`

NewManagersRowWithDefaults instantiates a new ManagersRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagerId

`func (o *ManagersRow) GetManagerId() string`

GetManagerId returns the ManagerId field if non-nil, zero value otherwise.

### GetManagerIdOk

`func (o *ManagersRow) GetManagerIdOk() (*string, bool)`

GetManagerIdOk returns a tuple with the ManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerId

`func (o *ManagersRow) SetManagerId(v string)`

SetManagerId sets ManagerId field to given value.

### HasManagerId

`func (o *ManagersRow) HasManagerId() bool`

HasManagerId returns a boolean if a field has been set.

### SetManagerIdNil

`func (o *ManagersRow) SetManagerIdNil(b bool)`

 SetManagerIdNil sets the value for ManagerId to be an explicit nil

### UnsetManagerId
`func (o *ManagersRow) UnsetManagerId()`

UnsetManagerId ensures that no value is present for ManagerId, not even an explicit nil
### GetName

`func (o *ManagersRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagersRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagersRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagersRow) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ManagersRow) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ManagersRow) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNameVariantCount

`func (o *ManagersRow) GetNameVariantCount() int32`

GetNameVariantCount returns the NameVariantCount field if non-nil, zero value otherwise.

### GetNameVariantCountOk

`func (o *ManagersRow) GetNameVariantCountOk() (*int32, bool)`

GetNameVariantCountOk returns a tuple with the NameVariantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameVariantCount

`func (o *ManagersRow) SetNameVariantCount(v int32)`

SetNameVariantCount sets NameVariantCount field to given value.

### HasNameVariantCount

`func (o *ManagersRow) HasNameVariantCount() bool`

HasNameVariantCount returns a boolean if a field has been set.

### SetNameVariantCountNil

`func (o *ManagersRow) SetNameVariantCountNil(b bool)`

 SetNameVariantCountNil sets the value for NameVariantCount to be an explicit nil

### UnsetNameVariantCount
`func (o *ManagersRow) UnsetNameVariantCount()`

UnsetNameVariantCount ensures that no value is present for NameVariantCount, not even an explicit nil
### GetInvestorId

`func (o *ManagersRow) GetInvestorId() string`

GetInvestorId returns the InvestorId field if non-nil, zero value otherwise.

### GetInvestorIdOk

`func (o *ManagersRow) GetInvestorIdOk() (*string, bool)`

GetInvestorIdOk returns a tuple with the InvestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestorId

`func (o *ManagersRow) SetInvestorId(v string)`

SetInvestorId sets InvestorId field to given value.

### HasInvestorId

`func (o *ManagersRow) HasInvestorId() bool`

HasInvestorId returns a boolean if a field has been set.

### SetInvestorIdNil

`func (o *ManagersRow) SetInvestorIdNil(b bool)`

 SetInvestorIdNil sets the value for InvestorId to be an explicit nil

### UnsetInvestorId
`func (o *ManagersRow) UnsetInvestorId()`

UnsetInvestorId ensures that no value is present for InvestorId, not even an explicit nil
### GetInvestorLinkConflict

`func (o *ManagersRow) GetInvestorLinkConflict() bool`

GetInvestorLinkConflict returns the InvestorLinkConflict field if non-nil, zero value otherwise.

### GetInvestorLinkConflictOk

`func (o *ManagersRow) GetInvestorLinkConflictOk() (*bool, bool)`

GetInvestorLinkConflictOk returns a tuple with the InvestorLinkConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestorLinkConflict

`func (o *ManagersRow) SetInvestorLinkConflict(v bool)`

SetInvestorLinkConflict sets InvestorLinkConflict field to given value.

### HasInvestorLinkConflict

`func (o *ManagersRow) HasInvestorLinkConflict() bool`

HasInvestorLinkConflict returns a boolean if a field has been set.

### SetInvestorLinkConflictNil

`func (o *ManagersRow) SetInvestorLinkConflictNil(b bool)`

 SetInvestorLinkConflictNil sets the value for InvestorLinkConflict to be an explicit nil

### UnsetInvestorLinkConflict
`func (o *ManagersRow) UnsetInvestorLinkConflict()`

UnsetInvestorLinkConflict ensures that no value is present for InvestorLinkConflict, not even an explicit nil
### GetOrgId

`func (o *ManagersRow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ManagersRow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ManagersRow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ManagersRow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *ManagersRow) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *ManagersRow) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetCompanyId

`func (o *ManagersRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ManagersRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ManagersRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ManagersRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *ManagersRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *ManagersRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *ManagersRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *ManagersRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *ManagersRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *ManagersRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *ManagersRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *ManagersRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *ManagersRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ManagersRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ManagersRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ManagersRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *ManagersRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *ManagersRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *ManagersRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ManagersRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ManagersRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ManagersRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *ManagersRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *ManagersRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetPositionCount

`func (o *ManagersRow) GetPositionCount() int32`

GetPositionCount returns the PositionCount field if non-nil, zero value otherwise.

### GetPositionCountOk

`func (o *ManagersRow) GetPositionCountOk() (*int32, bool)`

GetPositionCountOk returns a tuple with the PositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionCount

`func (o *ManagersRow) SetPositionCount(v int32)`

SetPositionCount sets PositionCount field to given value.

### HasPositionCount

`func (o *ManagersRow) HasPositionCount() bool`

HasPositionCount returns a boolean if a field has been set.

### SetPositionCountNil

`func (o *ManagersRow) SetPositionCountNil(b bool)`

 SetPositionCountNil sets the value for PositionCount to be an explicit nil

### UnsetPositionCount
`func (o *ManagersRow) UnsetPositionCount()`

UnsetPositionCount ensures that no value is present for PositionCount, not even an explicit nil
### GetIssuerCount

`func (o *ManagersRow) GetIssuerCount() int32`

GetIssuerCount returns the IssuerCount field if non-nil, zero value otherwise.

### GetIssuerCountOk

`func (o *ManagersRow) GetIssuerCountOk() (*int32, bool)`

GetIssuerCountOk returns a tuple with the IssuerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCount

`func (o *ManagersRow) SetIssuerCount(v int32)`

SetIssuerCount sets IssuerCount field to given value.

### HasIssuerCount

`func (o *ManagersRow) HasIssuerCount() bool`

HasIssuerCount returns a boolean if a field has been set.

### SetIssuerCountNil

`func (o *ManagersRow) SetIssuerCountNil(b bool)`

 SetIssuerCountNil sets the value for IssuerCount to be an explicit nil

### UnsetIssuerCount
`func (o *ManagersRow) UnsetIssuerCount()`

UnsetIssuerCount ensures that no value is present for IssuerCount, not even an explicit nil
### GetValuedPositionCount

`func (o *ManagersRow) GetValuedPositionCount() int32`

GetValuedPositionCount returns the ValuedPositionCount field if non-nil, zero value otherwise.

### GetValuedPositionCountOk

`func (o *ManagersRow) GetValuedPositionCountOk() (*int32, bool)`

GetValuedPositionCountOk returns a tuple with the ValuedPositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuedPositionCount

`func (o *ManagersRow) SetValuedPositionCount(v int32)`

SetValuedPositionCount sets ValuedPositionCount field to given value.

### HasValuedPositionCount

`func (o *ManagersRow) HasValuedPositionCount() bool`

HasValuedPositionCount returns a boolean if a field has been set.

### SetValuedPositionCountNil

`func (o *ManagersRow) SetValuedPositionCountNil(b bool)`

 SetValuedPositionCountNil sets the value for ValuedPositionCount to be an explicit nil

### UnsetValuedPositionCount
`func (o *ManagersRow) UnsetValuedPositionCount()`

UnsetValuedPositionCount ensures that no value is present for ValuedPositionCount, not even an explicit nil
### GetFirstObservedDate

`func (o *ManagersRow) GetFirstObservedDate() string`

GetFirstObservedDate returns the FirstObservedDate field if non-nil, zero value otherwise.

### GetFirstObservedDateOk

`func (o *ManagersRow) GetFirstObservedDateOk() (*string, bool)`

GetFirstObservedDateOk returns a tuple with the FirstObservedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstObservedDate

`func (o *ManagersRow) SetFirstObservedDate(v string)`

SetFirstObservedDate sets FirstObservedDate field to given value.

### HasFirstObservedDate

`func (o *ManagersRow) HasFirstObservedDate() bool`

HasFirstObservedDate returns a boolean if a field has been set.

### SetFirstObservedDateNil

`func (o *ManagersRow) SetFirstObservedDateNil(b bool)`

 SetFirstObservedDateNil sets the value for FirstObservedDate to be an explicit nil

### UnsetFirstObservedDate
`func (o *ManagersRow) UnsetFirstObservedDate()`

UnsetFirstObservedDate ensures that no value is present for FirstObservedDate, not even an explicit nil
### GetLastFilingDate

`func (o *ManagersRow) GetLastFilingDate() string`

GetLastFilingDate returns the LastFilingDate field if non-nil, zero value otherwise.

### GetLastFilingDateOk

`func (o *ManagersRow) GetLastFilingDateOk() (*string, bool)`

GetLastFilingDateOk returns a tuple with the LastFilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFilingDate

`func (o *ManagersRow) SetLastFilingDate(v string)`

SetLastFilingDate sets LastFilingDate field to given value.

### HasLastFilingDate

`func (o *ManagersRow) HasLastFilingDate() bool`

HasLastFilingDate returns a boolean if a field has been set.

### SetLastFilingDateNil

`func (o *ManagersRow) SetLastFilingDateNil(b bool)`

 SetLastFilingDateNil sets the value for LastFilingDate to be an explicit nil

### UnsetLastFilingDate
`func (o *ManagersRow) UnsetLastFilingDate()`

UnsetLastFilingDate ensures that no value is present for LastFilingDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


