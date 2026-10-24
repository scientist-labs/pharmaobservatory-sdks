# CompanyCollaboratorsRow

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
**CollaboratorCompanyId** | Pointer to **NullableString** | Company level: the company that co-sponsors the trials. | [optional] 
**CollaboratorCompanyName** | Pointer to **NullableString** | Display name of the collaborator&#39;s company. | [optional] 
**CollaboratorGroupId** | Pointer to **NullableString** | Group level: the collaborator&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**CollaboratorGroupName** | Pointer to **NullableString** | Display name of the collaborator&#39;s group. | [optional] 
**CollaboratorRank** | Pointer to **NullableInt32** | 0 for the company itself, 1 to 6 for its collaborators ranked by shared trials across its top disease areas. | [optional] 
**SharedTrials** | Pointer to **NullableInt32** | Distinct trials the company shares with the collaborator in this disease area (for rank 0, the company&#39;s own trials in it). | [optional] 

## Methods

### NewCompanyCollaboratorsRow

`func NewCompanyCollaboratorsRow() *CompanyCollaboratorsRow`

NewCompanyCollaboratorsRow instantiates a new CompanyCollaboratorsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyCollaboratorsRowWithDefaults

`func NewCompanyCollaboratorsRowWithDefaults() *CompanyCollaboratorsRow`

NewCompanyCollaboratorsRowWithDefaults instantiates a new CompanyCollaboratorsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CompanyCollaboratorsRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyCollaboratorsRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyCollaboratorsRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CompanyCollaboratorsRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CompanyCollaboratorsRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CompanyCollaboratorsRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *CompanyCollaboratorsRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyCollaboratorsRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyCollaboratorsRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyCollaboratorsRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompanyCollaboratorsRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyCollaboratorsRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *CompanyCollaboratorsRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CompanyCollaboratorsRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CompanyCollaboratorsRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CompanyCollaboratorsRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CompanyCollaboratorsRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CompanyCollaboratorsRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *CompanyCollaboratorsRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CompanyCollaboratorsRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CompanyCollaboratorsRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CompanyCollaboratorsRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CompanyCollaboratorsRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CompanyCollaboratorsRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetConditionId

`func (o *CompanyCollaboratorsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *CompanyCollaboratorsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *CompanyCollaboratorsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *CompanyCollaboratorsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *CompanyCollaboratorsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *CompanyCollaboratorsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *CompanyCollaboratorsRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *CompanyCollaboratorsRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *CompanyCollaboratorsRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *CompanyCollaboratorsRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *CompanyCollaboratorsRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *CompanyCollaboratorsRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetDiseaseRank

`func (o *CompanyCollaboratorsRow) GetDiseaseRank() int32`

GetDiseaseRank returns the DiseaseRank field if non-nil, zero value otherwise.

### GetDiseaseRankOk

`func (o *CompanyCollaboratorsRow) GetDiseaseRankOk() (*int32, bool)`

GetDiseaseRankOk returns a tuple with the DiseaseRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiseaseRank

`func (o *CompanyCollaboratorsRow) SetDiseaseRank(v int32)`

SetDiseaseRank sets DiseaseRank field to given value.

### HasDiseaseRank

`func (o *CompanyCollaboratorsRow) HasDiseaseRank() bool`

HasDiseaseRank returns a boolean if a field has been set.

### SetDiseaseRankNil

`func (o *CompanyCollaboratorsRow) SetDiseaseRankNil(b bool)`

 SetDiseaseRankNil sets the value for DiseaseRank to be an explicit nil

### UnsetDiseaseRank
`func (o *CompanyCollaboratorsRow) UnsetDiseaseRank()`

UnsetDiseaseRank ensures that no value is present for DiseaseRank, not even an explicit nil
### GetCollaboratorCompanyId

`func (o *CompanyCollaboratorsRow) GetCollaboratorCompanyId() string`

GetCollaboratorCompanyId returns the CollaboratorCompanyId field if non-nil, zero value otherwise.

### GetCollaboratorCompanyIdOk

`func (o *CompanyCollaboratorsRow) GetCollaboratorCompanyIdOk() (*string, bool)`

GetCollaboratorCompanyIdOk returns a tuple with the CollaboratorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaboratorCompanyId

`func (o *CompanyCollaboratorsRow) SetCollaboratorCompanyId(v string)`

SetCollaboratorCompanyId sets CollaboratorCompanyId field to given value.

### HasCollaboratorCompanyId

`func (o *CompanyCollaboratorsRow) HasCollaboratorCompanyId() bool`

HasCollaboratorCompanyId returns a boolean if a field has been set.

### SetCollaboratorCompanyIdNil

`func (o *CompanyCollaboratorsRow) SetCollaboratorCompanyIdNil(b bool)`

 SetCollaboratorCompanyIdNil sets the value for CollaboratorCompanyId to be an explicit nil

### UnsetCollaboratorCompanyId
`func (o *CompanyCollaboratorsRow) UnsetCollaboratorCompanyId()`

UnsetCollaboratorCompanyId ensures that no value is present for CollaboratorCompanyId, not even an explicit nil
### GetCollaboratorCompanyName

`func (o *CompanyCollaboratorsRow) GetCollaboratorCompanyName() string`

GetCollaboratorCompanyName returns the CollaboratorCompanyName field if non-nil, zero value otherwise.

### GetCollaboratorCompanyNameOk

`func (o *CompanyCollaboratorsRow) GetCollaboratorCompanyNameOk() (*string, bool)`

GetCollaboratorCompanyNameOk returns a tuple with the CollaboratorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaboratorCompanyName

`func (o *CompanyCollaboratorsRow) SetCollaboratorCompanyName(v string)`

SetCollaboratorCompanyName sets CollaboratorCompanyName field to given value.

### HasCollaboratorCompanyName

`func (o *CompanyCollaboratorsRow) HasCollaboratorCompanyName() bool`

HasCollaboratorCompanyName returns a boolean if a field has been set.

### SetCollaboratorCompanyNameNil

`func (o *CompanyCollaboratorsRow) SetCollaboratorCompanyNameNil(b bool)`

 SetCollaboratorCompanyNameNil sets the value for CollaboratorCompanyName to be an explicit nil

### UnsetCollaboratorCompanyName
`func (o *CompanyCollaboratorsRow) UnsetCollaboratorCompanyName()`

UnsetCollaboratorCompanyName ensures that no value is present for CollaboratorCompanyName, not even an explicit nil
### GetCollaboratorGroupId

`func (o *CompanyCollaboratorsRow) GetCollaboratorGroupId() string`

GetCollaboratorGroupId returns the CollaboratorGroupId field if non-nil, zero value otherwise.

### GetCollaboratorGroupIdOk

`func (o *CompanyCollaboratorsRow) GetCollaboratorGroupIdOk() (*string, bool)`

GetCollaboratorGroupIdOk returns a tuple with the CollaboratorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaboratorGroupId

`func (o *CompanyCollaboratorsRow) SetCollaboratorGroupId(v string)`

SetCollaboratorGroupId sets CollaboratorGroupId field to given value.

### HasCollaboratorGroupId

`func (o *CompanyCollaboratorsRow) HasCollaboratorGroupId() bool`

HasCollaboratorGroupId returns a boolean if a field has been set.

### SetCollaboratorGroupIdNil

`func (o *CompanyCollaboratorsRow) SetCollaboratorGroupIdNil(b bool)`

 SetCollaboratorGroupIdNil sets the value for CollaboratorGroupId to be an explicit nil

### UnsetCollaboratorGroupId
`func (o *CompanyCollaboratorsRow) UnsetCollaboratorGroupId()`

UnsetCollaboratorGroupId ensures that no value is present for CollaboratorGroupId, not even an explicit nil
### GetCollaboratorGroupName

`func (o *CompanyCollaboratorsRow) GetCollaboratorGroupName() string`

GetCollaboratorGroupName returns the CollaboratorGroupName field if non-nil, zero value otherwise.

### GetCollaboratorGroupNameOk

`func (o *CompanyCollaboratorsRow) GetCollaboratorGroupNameOk() (*string, bool)`

GetCollaboratorGroupNameOk returns a tuple with the CollaboratorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaboratorGroupName

`func (o *CompanyCollaboratorsRow) SetCollaboratorGroupName(v string)`

SetCollaboratorGroupName sets CollaboratorGroupName field to given value.

### HasCollaboratorGroupName

`func (o *CompanyCollaboratorsRow) HasCollaboratorGroupName() bool`

HasCollaboratorGroupName returns a boolean if a field has been set.

### SetCollaboratorGroupNameNil

`func (o *CompanyCollaboratorsRow) SetCollaboratorGroupNameNil(b bool)`

 SetCollaboratorGroupNameNil sets the value for CollaboratorGroupName to be an explicit nil

### UnsetCollaboratorGroupName
`func (o *CompanyCollaboratorsRow) UnsetCollaboratorGroupName()`

UnsetCollaboratorGroupName ensures that no value is present for CollaboratorGroupName, not even an explicit nil
### GetCollaboratorRank

`func (o *CompanyCollaboratorsRow) GetCollaboratorRank() int32`

GetCollaboratorRank returns the CollaboratorRank field if non-nil, zero value otherwise.

### GetCollaboratorRankOk

`func (o *CompanyCollaboratorsRow) GetCollaboratorRankOk() (*int32, bool)`

GetCollaboratorRankOk returns a tuple with the CollaboratorRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaboratorRank

`func (o *CompanyCollaboratorsRow) SetCollaboratorRank(v int32)`

SetCollaboratorRank sets CollaboratorRank field to given value.

### HasCollaboratorRank

`func (o *CompanyCollaboratorsRow) HasCollaboratorRank() bool`

HasCollaboratorRank returns a boolean if a field has been set.

### SetCollaboratorRankNil

`func (o *CompanyCollaboratorsRow) SetCollaboratorRankNil(b bool)`

 SetCollaboratorRankNil sets the value for CollaboratorRank to be an explicit nil

### UnsetCollaboratorRank
`func (o *CompanyCollaboratorsRow) UnsetCollaboratorRank()`

UnsetCollaboratorRank ensures that no value is present for CollaboratorRank, not even an explicit nil
### GetSharedTrials

`func (o *CompanyCollaboratorsRow) GetSharedTrials() int32`

GetSharedTrials returns the SharedTrials field if non-nil, zero value otherwise.

### GetSharedTrialsOk

`func (o *CompanyCollaboratorsRow) GetSharedTrialsOk() (*int32, bool)`

GetSharedTrialsOk returns a tuple with the SharedTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedTrials

`func (o *CompanyCollaboratorsRow) SetSharedTrials(v int32)`

SetSharedTrials sets SharedTrials field to given value.

### HasSharedTrials

`func (o *CompanyCollaboratorsRow) HasSharedTrials() bool`

HasSharedTrials returns a boolean if a field has been set.

### SetSharedTrialsNil

`func (o *CompanyCollaboratorsRow) SetSharedTrialsNil(b bool)`

 SetSharedTrialsNil sets the value for SharedTrials to be an explicit nil

### UnsetSharedTrials
`func (o *CompanyCollaboratorsRow) UnsetSharedTrials()`

UnsetSharedTrials ensures that no value is present for SharedTrials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


