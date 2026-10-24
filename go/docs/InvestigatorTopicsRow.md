# InvestigatorTopicsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvestigatorPersonId** | Pointer to **NullableString** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] 
**DescriptorId** | Pointer to **NullableString** | The MeSH descriptor id. | [optional] 
**DescriptorName** | Pointer to **NullableString** | The MeSH descriptor name. | [optional] 
**IsDisease** | Pointer to **NullableBool** | True when the descriptor is also a condition. | [optional] 
**ConditionId** | Pointer to **NullableString** | The descriptor id as a conditions id when the descriptor is a condition (null otherwise). | [optional] 
**NPublications** | Pointer to **NullableInt64** | Distinct papers the person authored on the descriptor. | [optional] 
**NMajorTopicPublications** | Pointer to **NullableInt64** | Those papers where the descriptor is a major topic. | [optional] 
**TopicWeight** | Pointer to **NullableInt64** | n_publications plus n_major_topic_publications, so a major-topic paper counts twice. | [optional] 
**TopicRank** | Pointer to **NullableInt32** | The person&#39;s own topic order by weight (1 is their strongest topic); not a ranking of people. | [optional] 

## Methods

### NewInvestigatorTopicsRow

`func NewInvestigatorTopicsRow() *InvestigatorTopicsRow`

NewInvestigatorTopicsRow instantiates a new InvestigatorTopicsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigatorTopicsRowWithDefaults

`func NewInvestigatorTopicsRowWithDefaults() *InvestigatorTopicsRow`

NewInvestigatorTopicsRowWithDefaults instantiates a new InvestigatorTopicsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestigatorPersonId

`func (o *InvestigatorTopicsRow) GetInvestigatorPersonId() string`

GetInvestigatorPersonId returns the InvestigatorPersonId field if non-nil, zero value otherwise.

### GetInvestigatorPersonIdOk

`func (o *InvestigatorTopicsRow) GetInvestigatorPersonIdOk() (*string, bool)`

GetInvestigatorPersonIdOk returns a tuple with the InvestigatorPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonId

`func (o *InvestigatorTopicsRow) SetInvestigatorPersonId(v string)`

SetInvestigatorPersonId sets InvestigatorPersonId field to given value.

### HasInvestigatorPersonId

`func (o *InvestigatorTopicsRow) HasInvestigatorPersonId() bool`

HasInvestigatorPersonId returns a boolean if a field has been set.

### SetInvestigatorPersonIdNil

`func (o *InvestigatorTopicsRow) SetInvestigatorPersonIdNil(b bool)`

 SetInvestigatorPersonIdNil sets the value for InvestigatorPersonId to be an explicit nil

### UnsetInvestigatorPersonId
`func (o *InvestigatorTopicsRow) UnsetInvestigatorPersonId()`

UnsetInvestigatorPersonId ensures that no value is present for InvestigatorPersonId, not even an explicit nil
### GetDescriptorId

`func (o *InvestigatorTopicsRow) GetDescriptorId() string`

GetDescriptorId returns the DescriptorId field if non-nil, zero value otherwise.

### GetDescriptorIdOk

`func (o *InvestigatorTopicsRow) GetDescriptorIdOk() (*string, bool)`

GetDescriptorIdOk returns a tuple with the DescriptorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorId

`func (o *InvestigatorTopicsRow) SetDescriptorId(v string)`

SetDescriptorId sets DescriptorId field to given value.

### HasDescriptorId

`func (o *InvestigatorTopicsRow) HasDescriptorId() bool`

HasDescriptorId returns a boolean if a field has been set.

### SetDescriptorIdNil

`func (o *InvestigatorTopicsRow) SetDescriptorIdNil(b bool)`

 SetDescriptorIdNil sets the value for DescriptorId to be an explicit nil

### UnsetDescriptorId
`func (o *InvestigatorTopicsRow) UnsetDescriptorId()`

UnsetDescriptorId ensures that no value is present for DescriptorId, not even an explicit nil
### GetDescriptorName

`func (o *InvestigatorTopicsRow) GetDescriptorName() string`

GetDescriptorName returns the DescriptorName field if non-nil, zero value otherwise.

### GetDescriptorNameOk

`func (o *InvestigatorTopicsRow) GetDescriptorNameOk() (*string, bool)`

GetDescriptorNameOk returns a tuple with the DescriptorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorName

`func (o *InvestigatorTopicsRow) SetDescriptorName(v string)`

SetDescriptorName sets DescriptorName field to given value.

### HasDescriptorName

`func (o *InvestigatorTopicsRow) HasDescriptorName() bool`

HasDescriptorName returns a boolean if a field has been set.

### SetDescriptorNameNil

`func (o *InvestigatorTopicsRow) SetDescriptorNameNil(b bool)`

 SetDescriptorNameNil sets the value for DescriptorName to be an explicit nil

### UnsetDescriptorName
`func (o *InvestigatorTopicsRow) UnsetDescriptorName()`

UnsetDescriptorName ensures that no value is present for DescriptorName, not even an explicit nil
### GetIsDisease

`func (o *InvestigatorTopicsRow) GetIsDisease() bool`

GetIsDisease returns the IsDisease field if non-nil, zero value otherwise.

### GetIsDiseaseOk

`func (o *InvestigatorTopicsRow) GetIsDiseaseOk() (*bool, bool)`

GetIsDiseaseOk returns a tuple with the IsDisease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisease

`func (o *InvestigatorTopicsRow) SetIsDisease(v bool)`

SetIsDisease sets IsDisease field to given value.

### HasIsDisease

`func (o *InvestigatorTopicsRow) HasIsDisease() bool`

HasIsDisease returns a boolean if a field has been set.

### SetIsDiseaseNil

`func (o *InvestigatorTopicsRow) SetIsDiseaseNil(b bool)`

 SetIsDiseaseNil sets the value for IsDisease to be an explicit nil

### UnsetIsDisease
`func (o *InvestigatorTopicsRow) UnsetIsDisease()`

UnsetIsDisease ensures that no value is present for IsDisease, not even an explicit nil
### GetConditionId

`func (o *InvestigatorTopicsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *InvestigatorTopicsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *InvestigatorTopicsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *InvestigatorTopicsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *InvestigatorTopicsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *InvestigatorTopicsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetNPublications

`func (o *InvestigatorTopicsRow) GetNPublications() int64`

GetNPublications returns the NPublications field if non-nil, zero value otherwise.

### GetNPublicationsOk

`func (o *InvestigatorTopicsRow) GetNPublicationsOk() (*int64, bool)`

GetNPublicationsOk returns a tuple with the NPublications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPublications

`func (o *InvestigatorTopicsRow) SetNPublications(v int64)`

SetNPublications sets NPublications field to given value.

### HasNPublications

`func (o *InvestigatorTopicsRow) HasNPublications() bool`

HasNPublications returns a boolean if a field has been set.

### SetNPublicationsNil

`func (o *InvestigatorTopicsRow) SetNPublicationsNil(b bool)`

 SetNPublicationsNil sets the value for NPublications to be an explicit nil

### UnsetNPublications
`func (o *InvestigatorTopicsRow) UnsetNPublications()`

UnsetNPublications ensures that no value is present for NPublications, not even an explicit nil
### GetNMajorTopicPublications

`func (o *InvestigatorTopicsRow) GetNMajorTopicPublications() int64`

GetNMajorTopicPublications returns the NMajorTopicPublications field if non-nil, zero value otherwise.

### GetNMajorTopicPublicationsOk

`func (o *InvestigatorTopicsRow) GetNMajorTopicPublicationsOk() (*int64, bool)`

GetNMajorTopicPublicationsOk returns a tuple with the NMajorTopicPublications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNMajorTopicPublications

`func (o *InvestigatorTopicsRow) SetNMajorTopicPublications(v int64)`

SetNMajorTopicPublications sets NMajorTopicPublications field to given value.

### HasNMajorTopicPublications

`func (o *InvestigatorTopicsRow) HasNMajorTopicPublications() bool`

HasNMajorTopicPublications returns a boolean if a field has been set.

### SetNMajorTopicPublicationsNil

`func (o *InvestigatorTopicsRow) SetNMajorTopicPublicationsNil(b bool)`

 SetNMajorTopicPublicationsNil sets the value for NMajorTopicPublications to be an explicit nil

### UnsetNMajorTopicPublications
`func (o *InvestigatorTopicsRow) UnsetNMajorTopicPublications()`

UnsetNMajorTopicPublications ensures that no value is present for NMajorTopicPublications, not even an explicit nil
### GetTopicWeight

`func (o *InvestigatorTopicsRow) GetTopicWeight() int64`

GetTopicWeight returns the TopicWeight field if non-nil, zero value otherwise.

### GetTopicWeightOk

`func (o *InvestigatorTopicsRow) GetTopicWeightOk() (*int64, bool)`

GetTopicWeightOk returns a tuple with the TopicWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicWeight

`func (o *InvestigatorTopicsRow) SetTopicWeight(v int64)`

SetTopicWeight sets TopicWeight field to given value.

### HasTopicWeight

`func (o *InvestigatorTopicsRow) HasTopicWeight() bool`

HasTopicWeight returns a boolean if a field has been set.

### SetTopicWeightNil

`func (o *InvestigatorTopicsRow) SetTopicWeightNil(b bool)`

 SetTopicWeightNil sets the value for TopicWeight to be an explicit nil

### UnsetTopicWeight
`func (o *InvestigatorTopicsRow) UnsetTopicWeight()`

UnsetTopicWeight ensures that no value is present for TopicWeight, not even an explicit nil
### GetTopicRank

`func (o *InvestigatorTopicsRow) GetTopicRank() int32`

GetTopicRank returns the TopicRank field if non-nil, zero value otherwise.

### GetTopicRankOk

`func (o *InvestigatorTopicsRow) GetTopicRankOk() (*int32, bool)`

GetTopicRankOk returns a tuple with the TopicRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicRank

`func (o *InvestigatorTopicsRow) SetTopicRank(v int32)`

SetTopicRank sets TopicRank field to given value.

### HasTopicRank

`func (o *InvestigatorTopicsRow) HasTopicRank() bool`

HasTopicRank returns a boolean if a field has been set.

### SetTopicRankNil

`func (o *InvestigatorTopicsRow) SetTopicRankNil(b bool)`

 SetTopicRankNil sets the value for TopicRank to be an explicit nil

### UnsetTopicRank
`func (o *InvestigatorTopicsRow) UnsetTopicRank()`

UnsetTopicRank ensures that no value is present for TopicRank, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


