# CompanyMembersRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** | The company the legal entity is a member of. | [optional] 
**CompanyName** | Pointer to **NullableString** | The company&#39;s display name. | [optional] 
**LegalEntityId** | Pointer to **NullableString** | A registry-scoped legal entity id (edgar:, ror:, lei: or rx_provider:). A registration, not an organization record id. | [optional] 
**LinkBasis** | Pointer to **NullableString** | The source that placed the legal entity in the company (gleif, edgar, ror, name_key and so on). | [optional] 
**ParentCompanyId** | Pointer to **NullableString** | The company&#39;s direct owner, null for a root. | [optional] 
**ParentCompanyName** | Pointer to **NullableString** | Display name of the direct owner. | [optional] 
**ParentLinkBasis** | Pointer to **NullableString** | The authority of the first ownership hop. | [optional] 
**Since** | Pointer to **NullableString** | Date the ownership link began, when every hop is dated. | [optional] 
**SincePrecision** | Pointer to **NullableString** | Precision of since. | [optional] 
**SinceSource** | Pointer to **NullableString** | Where the date came from. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**HopDepth** | Pointer to **NullableInt32** | Hops from the company to its ultimate parent (0 for a root). | [optional] 

## Methods

### NewCompanyMembersRow

`func NewCompanyMembersRow() *CompanyMembersRow`

NewCompanyMembersRow instantiates a new CompanyMembersRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyMembersRowWithDefaults

`func NewCompanyMembersRowWithDefaults() *CompanyMembersRow`

NewCompanyMembersRowWithDefaults instantiates a new CompanyMembersRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CompanyMembersRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyMembersRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyMembersRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CompanyMembersRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CompanyMembersRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CompanyMembersRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *CompanyMembersRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyMembersRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyMembersRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyMembersRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompanyMembersRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyMembersRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetLegalEntityId

`func (o *CompanyMembersRow) GetLegalEntityId() string`

GetLegalEntityId returns the LegalEntityId field if non-nil, zero value otherwise.

### GetLegalEntityIdOk

`func (o *CompanyMembersRow) GetLegalEntityIdOk() (*string, bool)`

GetLegalEntityIdOk returns a tuple with the LegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityId

`func (o *CompanyMembersRow) SetLegalEntityId(v string)`

SetLegalEntityId sets LegalEntityId field to given value.

### HasLegalEntityId

`func (o *CompanyMembersRow) HasLegalEntityId() bool`

HasLegalEntityId returns a boolean if a field has been set.

### SetLegalEntityIdNil

`func (o *CompanyMembersRow) SetLegalEntityIdNil(b bool)`

 SetLegalEntityIdNil sets the value for LegalEntityId to be an explicit nil

### UnsetLegalEntityId
`func (o *CompanyMembersRow) UnsetLegalEntityId()`

UnsetLegalEntityId ensures that no value is present for LegalEntityId, not even an explicit nil
### GetLinkBasis

`func (o *CompanyMembersRow) GetLinkBasis() string`

GetLinkBasis returns the LinkBasis field if non-nil, zero value otherwise.

### GetLinkBasisOk

`func (o *CompanyMembersRow) GetLinkBasisOk() (*string, bool)`

GetLinkBasisOk returns a tuple with the LinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkBasis

`func (o *CompanyMembersRow) SetLinkBasis(v string)`

SetLinkBasis sets LinkBasis field to given value.

### HasLinkBasis

`func (o *CompanyMembersRow) HasLinkBasis() bool`

HasLinkBasis returns a boolean if a field has been set.

### SetLinkBasisNil

`func (o *CompanyMembersRow) SetLinkBasisNil(b bool)`

 SetLinkBasisNil sets the value for LinkBasis to be an explicit nil

### UnsetLinkBasis
`func (o *CompanyMembersRow) UnsetLinkBasis()`

UnsetLinkBasis ensures that no value is present for LinkBasis, not even an explicit nil
### GetParentCompanyId

`func (o *CompanyMembersRow) GetParentCompanyId() string`

GetParentCompanyId returns the ParentCompanyId field if non-nil, zero value otherwise.

### GetParentCompanyIdOk

`func (o *CompanyMembersRow) GetParentCompanyIdOk() (*string, bool)`

GetParentCompanyIdOk returns a tuple with the ParentCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCompanyId

`func (o *CompanyMembersRow) SetParentCompanyId(v string)`

SetParentCompanyId sets ParentCompanyId field to given value.

### HasParentCompanyId

`func (o *CompanyMembersRow) HasParentCompanyId() bool`

HasParentCompanyId returns a boolean if a field has been set.

### SetParentCompanyIdNil

`func (o *CompanyMembersRow) SetParentCompanyIdNil(b bool)`

 SetParentCompanyIdNil sets the value for ParentCompanyId to be an explicit nil

### UnsetParentCompanyId
`func (o *CompanyMembersRow) UnsetParentCompanyId()`

UnsetParentCompanyId ensures that no value is present for ParentCompanyId, not even an explicit nil
### GetParentCompanyName

`func (o *CompanyMembersRow) GetParentCompanyName() string`

GetParentCompanyName returns the ParentCompanyName field if non-nil, zero value otherwise.

### GetParentCompanyNameOk

`func (o *CompanyMembersRow) GetParentCompanyNameOk() (*string, bool)`

GetParentCompanyNameOk returns a tuple with the ParentCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCompanyName

`func (o *CompanyMembersRow) SetParentCompanyName(v string)`

SetParentCompanyName sets ParentCompanyName field to given value.

### HasParentCompanyName

`func (o *CompanyMembersRow) HasParentCompanyName() bool`

HasParentCompanyName returns a boolean if a field has been set.

### SetParentCompanyNameNil

`func (o *CompanyMembersRow) SetParentCompanyNameNil(b bool)`

 SetParentCompanyNameNil sets the value for ParentCompanyName to be an explicit nil

### UnsetParentCompanyName
`func (o *CompanyMembersRow) UnsetParentCompanyName()`

UnsetParentCompanyName ensures that no value is present for ParentCompanyName, not even an explicit nil
### GetParentLinkBasis

`func (o *CompanyMembersRow) GetParentLinkBasis() string`

GetParentLinkBasis returns the ParentLinkBasis field if non-nil, zero value otherwise.

### GetParentLinkBasisOk

`func (o *CompanyMembersRow) GetParentLinkBasisOk() (*string, bool)`

GetParentLinkBasisOk returns a tuple with the ParentLinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLinkBasis

`func (o *CompanyMembersRow) SetParentLinkBasis(v string)`

SetParentLinkBasis sets ParentLinkBasis field to given value.

### HasParentLinkBasis

`func (o *CompanyMembersRow) HasParentLinkBasis() bool`

HasParentLinkBasis returns a boolean if a field has been set.

### SetParentLinkBasisNil

`func (o *CompanyMembersRow) SetParentLinkBasisNil(b bool)`

 SetParentLinkBasisNil sets the value for ParentLinkBasis to be an explicit nil

### UnsetParentLinkBasis
`func (o *CompanyMembersRow) UnsetParentLinkBasis()`

UnsetParentLinkBasis ensures that no value is present for ParentLinkBasis, not even an explicit nil
### GetSince

`func (o *CompanyMembersRow) GetSince() string`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *CompanyMembersRow) GetSinceOk() (*string, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *CompanyMembersRow) SetSince(v string)`

SetSince sets Since field to given value.

### HasSince

`func (o *CompanyMembersRow) HasSince() bool`

HasSince returns a boolean if a field has been set.

### SetSinceNil

`func (o *CompanyMembersRow) SetSinceNil(b bool)`

 SetSinceNil sets the value for Since to be an explicit nil

### UnsetSince
`func (o *CompanyMembersRow) UnsetSince()`

UnsetSince ensures that no value is present for Since, not even an explicit nil
### GetSincePrecision

`func (o *CompanyMembersRow) GetSincePrecision() string`

GetSincePrecision returns the SincePrecision field if non-nil, zero value otherwise.

### GetSincePrecisionOk

`func (o *CompanyMembersRow) GetSincePrecisionOk() (*string, bool)`

GetSincePrecisionOk returns a tuple with the SincePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSincePrecision

`func (o *CompanyMembersRow) SetSincePrecision(v string)`

SetSincePrecision sets SincePrecision field to given value.

### HasSincePrecision

`func (o *CompanyMembersRow) HasSincePrecision() bool`

HasSincePrecision returns a boolean if a field has been set.

### SetSincePrecisionNil

`func (o *CompanyMembersRow) SetSincePrecisionNil(b bool)`

 SetSincePrecisionNil sets the value for SincePrecision to be an explicit nil

### UnsetSincePrecision
`func (o *CompanyMembersRow) UnsetSincePrecision()`

UnsetSincePrecision ensures that no value is present for SincePrecision, not even an explicit nil
### GetSinceSource

`func (o *CompanyMembersRow) GetSinceSource() string`

GetSinceSource returns the SinceSource field if non-nil, zero value otherwise.

### GetSinceSourceOk

`func (o *CompanyMembersRow) GetSinceSourceOk() (*string, bool)`

GetSinceSourceOk returns a tuple with the SinceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinceSource

`func (o *CompanyMembersRow) SetSinceSource(v string)`

SetSinceSource sets SinceSource field to given value.

### HasSinceSource

`func (o *CompanyMembersRow) HasSinceSource() bool`

HasSinceSource returns a boolean if a field has been set.

### SetSinceSourceNil

`func (o *CompanyMembersRow) SetSinceSourceNil(b bool)`

 SetSinceSourceNil sets the value for SinceSource to be an explicit nil

### UnsetSinceSource
`func (o *CompanyMembersRow) UnsetSinceSource()`

UnsetSinceSource ensures that no value is present for SinceSource, not even an explicit nil
### GetGroupId

`func (o *CompanyMembersRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CompanyMembersRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CompanyMembersRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CompanyMembersRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CompanyMembersRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CompanyMembersRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *CompanyMembersRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CompanyMembersRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CompanyMembersRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CompanyMembersRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CompanyMembersRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CompanyMembersRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetHopDepth

`func (o *CompanyMembersRow) GetHopDepth() int32`

GetHopDepth returns the HopDepth field if non-nil, zero value otherwise.

### GetHopDepthOk

`func (o *CompanyMembersRow) GetHopDepthOk() (*int32, bool)`

GetHopDepthOk returns a tuple with the HopDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHopDepth

`func (o *CompanyMembersRow) SetHopDepth(v int32)`

SetHopDepth sets HopDepth field to given value.

### HasHopDepth

`func (o *CompanyMembersRow) HasHopDepth() bool`

HasHopDepth returns a boolean if a field has been set.

### SetHopDepthNil

`func (o *CompanyMembersRow) SetHopDepthNil(b bool)`

 SetHopDepthNil sets the value for HopDepth to be an explicit nil

### UnsetHopDepth
`func (o *CompanyMembersRow) UnsetHopDepth()`

UnsetHopDepth ensures that no value is present for HopDepth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


