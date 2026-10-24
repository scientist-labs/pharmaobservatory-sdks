# CompanyCompetitorActivityRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** | Company level: the company the row is about; the entry key for the companies object type. | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**ConditionName** | Pointer to **NullableString** | The MeSH descriptor name of the condition. | [optional] 
**DiseaseRank** | Pointer to **NullableInt32** | The condition&#39;s rank among the company&#39;s disease areas by trial count (1 is the busiest). Only the top 10 are kept. | [optional] 
**CompetitorCompanyId** | Pointer to **NullableString** | Company level: the company active in the same disease area. | [optional] 
**CompetitorCompanyName** | Pointer to **NullableString** | Display name of the competitor&#39;s company. | [optional] 
**CompetitorGroupId** | Pointer to **NullableString** | Group level: the competitor&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**CompetitorGroupName** | Pointer to **NullableString** | Display name of the competitor&#39;s group. | [optional] 
**CompetitorRank** | Pointer to **NullableInt32** | 0 for the company itself, 1 to 6 for the companies ranked by trial count across the company&#39;s top disease areas. | [optional] 
**Trials** | Pointer to **NullableInt32** | Distinct trials the company in this row has in the disease area. | [optional] 

## Methods

### NewCompanyCompetitorActivityRow

`func NewCompanyCompetitorActivityRow() *CompanyCompetitorActivityRow`

NewCompanyCompetitorActivityRow instantiates a new CompanyCompetitorActivityRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyCompetitorActivityRowWithDefaults

`func NewCompanyCompetitorActivityRowWithDefaults() *CompanyCompetitorActivityRow`

NewCompanyCompetitorActivityRowWithDefaults instantiates a new CompanyCompetitorActivityRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CompanyCompetitorActivityRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyCompetitorActivityRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyCompetitorActivityRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CompanyCompetitorActivityRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CompanyCompetitorActivityRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CompanyCompetitorActivityRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *CompanyCompetitorActivityRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyCompetitorActivityRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyCompetitorActivityRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyCompetitorActivityRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompanyCompetitorActivityRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyCompetitorActivityRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *CompanyCompetitorActivityRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CompanyCompetitorActivityRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CompanyCompetitorActivityRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CompanyCompetitorActivityRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CompanyCompetitorActivityRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CompanyCompetitorActivityRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *CompanyCompetitorActivityRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CompanyCompetitorActivityRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CompanyCompetitorActivityRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CompanyCompetitorActivityRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CompanyCompetitorActivityRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CompanyCompetitorActivityRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetConditionId

`func (o *CompanyCompetitorActivityRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *CompanyCompetitorActivityRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *CompanyCompetitorActivityRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *CompanyCompetitorActivityRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *CompanyCompetitorActivityRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *CompanyCompetitorActivityRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *CompanyCompetitorActivityRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *CompanyCompetitorActivityRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *CompanyCompetitorActivityRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *CompanyCompetitorActivityRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *CompanyCompetitorActivityRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *CompanyCompetitorActivityRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetDiseaseRank

`func (o *CompanyCompetitorActivityRow) GetDiseaseRank() int32`

GetDiseaseRank returns the DiseaseRank field if non-nil, zero value otherwise.

### GetDiseaseRankOk

`func (o *CompanyCompetitorActivityRow) GetDiseaseRankOk() (*int32, bool)`

GetDiseaseRankOk returns a tuple with the DiseaseRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiseaseRank

`func (o *CompanyCompetitorActivityRow) SetDiseaseRank(v int32)`

SetDiseaseRank sets DiseaseRank field to given value.

### HasDiseaseRank

`func (o *CompanyCompetitorActivityRow) HasDiseaseRank() bool`

HasDiseaseRank returns a boolean if a field has been set.

### SetDiseaseRankNil

`func (o *CompanyCompetitorActivityRow) SetDiseaseRankNil(b bool)`

 SetDiseaseRankNil sets the value for DiseaseRank to be an explicit nil

### UnsetDiseaseRank
`func (o *CompanyCompetitorActivityRow) UnsetDiseaseRank()`

UnsetDiseaseRank ensures that no value is present for DiseaseRank, not even an explicit nil
### GetCompetitorCompanyId

`func (o *CompanyCompetitorActivityRow) GetCompetitorCompanyId() string`

GetCompetitorCompanyId returns the CompetitorCompanyId field if non-nil, zero value otherwise.

### GetCompetitorCompanyIdOk

`func (o *CompanyCompetitorActivityRow) GetCompetitorCompanyIdOk() (*string, bool)`

GetCompetitorCompanyIdOk returns a tuple with the CompetitorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorCompanyId

`func (o *CompanyCompetitorActivityRow) SetCompetitorCompanyId(v string)`

SetCompetitorCompanyId sets CompetitorCompanyId field to given value.

### HasCompetitorCompanyId

`func (o *CompanyCompetitorActivityRow) HasCompetitorCompanyId() bool`

HasCompetitorCompanyId returns a boolean if a field has been set.

### SetCompetitorCompanyIdNil

`func (o *CompanyCompetitorActivityRow) SetCompetitorCompanyIdNil(b bool)`

 SetCompetitorCompanyIdNil sets the value for CompetitorCompanyId to be an explicit nil

### UnsetCompetitorCompanyId
`func (o *CompanyCompetitorActivityRow) UnsetCompetitorCompanyId()`

UnsetCompetitorCompanyId ensures that no value is present for CompetitorCompanyId, not even an explicit nil
### GetCompetitorCompanyName

`func (o *CompanyCompetitorActivityRow) GetCompetitorCompanyName() string`

GetCompetitorCompanyName returns the CompetitorCompanyName field if non-nil, zero value otherwise.

### GetCompetitorCompanyNameOk

`func (o *CompanyCompetitorActivityRow) GetCompetitorCompanyNameOk() (*string, bool)`

GetCompetitorCompanyNameOk returns a tuple with the CompetitorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorCompanyName

`func (o *CompanyCompetitorActivityRow) SetCompetitorCompanyName(v string)`

SetCompetitorCompanyName sets CompetitorCompanyName field to given value.

### HasCompetitorCompanyName

`func (o *CompanyCompetitorActivityRow) HasCompetitorCompanyName() bool`

HasCompetitorCompanyName returns a boolean if a field has been set.

### SetCompetitorCompanyNameNil

`func (o *CompanyCompetitorActivityRow) SetCompetitorCompanyNameNil(b bool)`

 SetCompetitorCompanyNameNil sets the value for CompetitorCompanyName to be an explicit nil

### UnsetCompetitorCompanyName
`func (o *CompanyCompetitorActivityRow) UnsetCompetitorCompanyName()`

UnsetCompetitorCompanyName ensures that no value is present for CompetitorCompanyName, not even an explicit nil
### GetCompetitorGroupId

`func (o *CompanyCompetitorActivityRow) GetCompetitorGroupId() string`

GetCompetitorGroupId returns the CompetitorGroupId field if non-nil, zero value otherwise.

### GetCompetitorGroupIdOk

`func (o *CompanyCompetitorActivityRow) GetCompetitorGroupIdOk() (*string, bool)`

GetCompetitorGroupIdOk returns a tuple with the CompetitorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorGroupId

`func (o *CompanyCompetitorActivityRow) SetCompetitorGroupId(v string)`

SetCompetitorGroupId sets CompetitorGroupId field to given value.

### HasCompetitorGroupId

`func (o *CompanyCompetitorActivityRow) HasCompetitorGroupId() bool`

HasCompetitorGroupId returns a boolean if a field has been set.

### SetCompetitorGroupIdNil

`func (o *CompanyCompetitorActivityRow) SetCompetitorGroupIdNil(b bool)`

 SetCompetitorGroupIdNil sets the value for CompetitorGroupId to be an explicit nil

### UnsetCompetitorGroupId
`func (o *CompanyCompetitorActivityRow) UnsetCompetitorGroupId()`

UnsetCompetitorGroupId ensures that no value is present for CompetitorGroupId, not even an explicit nil
### GetCompetitorGroupName

`func (o *CompanyCompetitorActivityRow) GetCompetitorGroupName() string`

GetCompetitorGroupName returns the CompetitorGroupName field if non-nil, zero value otherwise.

### GetCompetitorGroupNameOk

`func (o *CompanyCompetitorActivityRow) GetCompetitorGroupNameOk() (*string, bool)`

GetCompetitorGroupNameOk returns a tuple with the CompetitorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorGroupName

`func (o *CompanyCompetitorActivityRow) SetCompetitorGroupName(v string)`

SetCompetitorGroupName sets CompetitorGroupName field to given value.

### HasCompetitorGroupName

`func (o *CompanyCompetitorActivityRow) HasCompetitorGroupName() bool`

HasCompetitorGroupName returns a boolean if a field has been set.

### SetCompetitorGroupNameNil

`func (o *CompanyCompetitorActivityRow) SetCompetitorGroupNameNil(b bool)`

 SetCompetitorGroupNameNil sets the value for CompetitorGroupName to be an explicit nil

### UnsetCompetitorGroupName
`func (o *CompanyCompetitorActivityRow) UnsetCompetitorGroupName()`

UnsetCompetitorGroupName ensures that no value is present for CompetitorGroupName, not even an explicit nil
### GetCompetitorRank

`func (o *CompanyCompetitorActivityRow) GetCompetitorRank() int32`

GetCompetitorRank returns the CompetitorRank field if non-nil, zero value otherwise.

### GetCompetitorRankOk

`func (o *CompanyCompetitorActivityRow) GetCompetitorRankOk() (*int32, bool)`

GetCompetitorRankOk returns a tuple with the CompetitorRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorRank

`func (o *CompanyCompetitorActivityRow) SetCompetitorRank(v int32)`

SetCompetitorRank sets CompetitorRank field to given value.

### HasCompetitorRank

`func (o *CompanyCompetitorActivityRow) HasCompetitorRank() bool`

HasCompetitorRank returns a boolean if a field has been set.

### SetCompetitorRankNil

`func (o *CompanyCompetitorActivityRow) SetCompetitorRankNil(b bool)`

 SetCompetitorRankNil sets the value for CompetitorRank to be an explicit nil

### UnsetCompetitorRank
`func (o *CompanyCompetitorActivityRow) UnsetCompetitorRank()`

UnsetCompetitorRank ensures that no value is present for CompetitorRank, not even an explicit nil
### GetTrials

`func (o *CompanyCompetitorActivityRow) GetTrials() int32`

GetTrials returns the Trials field if non-nil, zero value otherwise.

### GetTrialsOk

`func (o *CompanyCompetitorActivityRow) GetTrialsOk() (*int32, bool)`

GetTrialsOk returns a tuple with the Trials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrials

`func (o *CompanyCompetitorActivityRow) SetTrials(v int32)`

SetTrials sets Trials field to given value.

### HasTrials

`func (o *CompanyCompetitorActivityRow) HasTrials() bool`

HasTrials returns a boolean if a field has been set.

### SetTrialsNil

`func (o *CompanyCompetitorActivityRow) SetTrialsNil(b bool)`

 SetTrialsNil sets the value for Trials to be an explicit nil

### UnsetTrials
`func (o *CompanyCompetitorActivityRow) UnsetTrials()`

UnsetTrials ensures that no value is present for Trials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


