# CompanyUsPatentsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** | Company level: the company the row is about; the entry key for the companies object type. | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**PatentId** | Pointer to **NullableString** | The PatentsView patent id. | [optional] 
**PatentNumber** | Pointer to **NullableString** | The US patent number. | [optional] 
**PatentTitle** | Pointer to **NullableString** | The patent title. | [optional] 
**PatentDate** | Pointer to **NullableString** | The grant date. | [optional] 
**IsBiomedical** | Pointer to **NullableBool** | True when the patent carries a biomedical CPC class; the split is never a filter on what belongs to the portfolio. | [optional] 

## Methods

### NewCompanyUsPatentsRow

`func NewCompanyUsPatentsRow() *CompanyUsPatentsRow`

NewCompanyUsPatentsRow instantiates a new CompanyUsPatentsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyUsPatentsRowWithDefaults

`func NewCompanyUsPatentsRowWithDefaults() *CompanyUsPatentsRow`

NewCompanyUsPatentsRowWithDefaults instantiates a new CompanyUsPatentsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CompanyUsPatentsRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyUsPatentsRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyUsPatentsRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CompanyUsPatentsRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CompanyUsPatentsRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CompanyUsPatentsRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *CompanyUsPatentsRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyUsPatentsRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyUsPatentsRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyUsPatentsRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompanyUsPatentsRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyUsPatentsRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *CompanyUsPatentsRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CompanyUsPatentsRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CompanyUsPatentsRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CompanyUsPatentsRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CompanyUsPatentsRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CompanyUsPatentsRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *CompanyUsPatentsRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CompanyUsPatentsRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CompanyUsPatentsRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CompanyUsPatentsRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CompanyUsPatentsRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CompanyUsPatentsRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetPatentId

`func (o *CompanyUsPatentsRow) GetPatentId() string`

GetPatentId returns the PatentId field if non-nil, zero value otherwise.

### GetPatentIdOk

`func (o *CompanyUsPatentsRow) GetPatentIdOk() (*string, bool)`

GetPatentIdOk returns a tuple with the PatentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatentId

`func (o *CompanyUsPatentsRow) SetPatentId(v string)`

SetPatentId sets PatentId field to given value.

### HasPatentId

`func (o *CompanyUsPatentsRow) HasPatentId() bool`

HasPatentId returns a boolean if a field has been set.

### SetPatentIdNil

`func (o *CompanyUsPatentsRow) SetPatentIdNil(b bool)`

 SetPatentIdNil sets the value for PatentId to be an explicit nil

### UnsetPatentId
`func (o *CompanyUsPatentsRow) UnsetPatentId()`

UnsetPatentId ensures that no value is present for PatentId, not even an explicit nil
### GetPatentNumber

`func (o *CompanyUsPatentsRow) GetPatentNumber() string`

GetPatentNumber returns the PatentNumber field if non-nil, zero value otherwise.

### GetPatentNumberOk

`func (o *CompanyUsPatentsRow) GetPatentNumberOk() (*string, bool)`

GetPatentNumberOk returns a tuple with the PatentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatentNumber

`func (o *CompanyUsPatentsRow) SetPatentNumber(v string)`

SetPatentNumber sets PatentNumber field to given value.

### HasPatentNumber

`func (o *CompanyUsPatentsRow) HasPatentNumber() bool`

HasPatentNumber returns a boolean if a field has been set.

### SetPatentNumberNil

`func (o *CompanyUsPatentsRow) SetPatentNumberNil(b bool)`

 SetPatentNumberNil sets the value for PatentNumber to be an explicit nil

### UnsetPatentNumber
`func (o *CompanyUsPatentsRow) UnsetPatentNumber()`

UnsetPatentNumber ensures that no value is present for PatentNumber, not even an explicit nil
### GetPatentTitle

`func (o *CompanyUsPatentsRow) GetPatentTitle() string`

GetPatentTitle returns the PatentTitle field if non-nil, zero value otherwise.

### GetPatentTitleOk

`func (o *CompanyUsPatentsRow) GetPatentTitleOk() (*string, bool)`

GetPatentTitleOk returns a tuple with the PatentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatentTitle

`func (o *CompanyUsPatentsRow) SetPatentTitle(v string)`

SetPatentTitle sets PatentTitle field to given value.

### HasPatentTitle

`func (o *CompanyUsPatentsRow) HasPatentTitle() bool`

HasPatentTitle returns a boolean if a field has been set.

### SetPatentTitleNil

`func (o *CompanyUsPatentsRow) SetPatentTitleNil(b bool)`

 SetPatentTitleNil sets the value for PatentTitle to be an explicit nil

### UnsetPatentTitle
`func (o *CompanyUsPatentsRow) UnsetPatentTitle()`

UnsetPatentTitle ensures that no value is present for PatentTitle, not even an explicit nil
### GetPatentDate

`func (o *CompanyUsPatentsRow) GetPatentDate() string`

GetPatentDate returns the PatentDate field if non-nil, zero value otherwise.

### GetPatentDateOk

`func (o *CompanyUsPatentsRow) GetPatentDateOk() (*string, bool)`

GetPatentDateOk returns a tuple with the PatentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatentDate

`func (o *CompanyUsPatentsRow) SetPatentDate(v string)`

SetPatentDate sets PatentDate field to given value.

### HasPatentDate

`func (o *CompanyUsPatentsRow) HasPatentDate() bool`

HasPatentDate returns a boolean if a field has been set.

### SetPatentDateNil

`func (o *CompanyUsPatentsRow) SetPatentDateNil(b bool)`

 SetPatentDateNil sets the value for PatentDate to be an explicit nil

### UnsetPatentDate
`func (o *CompanyUsPatentsRow) UnsetPatentDate()`

UnsetPatentDate ensures that no value is present for PatentDate, not even an explicit nil
### GetIsBiomedical

`func (o *CompanyUsPatentsRow) GetIsBiomedical() bool`

GetIsBiomedical returns the IsBiomedical field if non-nil, zero value otherwise.

### GetIsBiomedicalOk

`func (o *CompanyUsPatentsRow) GetIsBiomedicalOk() (*bool, bool)`

GetIsBiomedicalOk returns a tuple with the IsBiomedical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBiomedical

`func (o *CompanyUsPatentsRow) SetIsBiomedical(v bool)`

SetIsBiomedical sets IsBiomedical field to given value.

### HasIsBiomedical

`func (o *CompanyUsPatentsRow) HasIsBiomedical() bool`

HasIsBiomedical returns a boolean if a field has been set.

### SetIsBiomedicalNil

`func (o *CompanyUsPatentsRow) SetIsBiomedicalNil(b bool)`

 SetIsBiomedicalNil sets the value for IsBiomedical to be an explicit nil

### UnsetIsBiomedical
`func (o *CompanyUsPatentsRow) UnsetIsBiomedical()`

UnsetIsBiomedical ensures that no value is present for IsBiomedical, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


