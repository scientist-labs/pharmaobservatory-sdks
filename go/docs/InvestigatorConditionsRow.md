# InvestigatorConditionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvestigatorPersonId** | Pointer to **NullableString** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] 
**InvestigatorPersonName** | Pointer to **NullableString** | The investigator&#39;s name as the registry filed it, resolved across the person&#39;s merged name fragments. | [optional] 
**InvestigatorAffiliation** | Pointer to **NullableString** | The investigator&#39;s affiliation as filed (the most-trialled affiliation across the merged fragments). | [optional] 
**ConditionId** | Pointer to **NullableString** | The condition&#39;s MeSH descriptor id, the id in the conditions page address (D001943). | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s MeSH name. | [optional] 
**NTrials** | Pointer to **NullableInt64** | Distinct trials the person is named on in the condition. | [optional] 

## Methods

### NewInvestigatorConditionsRow

`func NewInvestigatorConditionsRow() *InvestigatorConditionsRow`

NewInvestigatorConditionsRow instantiates a new InvestigatorConditionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigatorConditionsRowWithDefaults

`func NewInvestigatorConditionsRowWithDefaults() *InvestigatorConditionsRow`

NewInvestigatorConditionsRowWithDefaults instantiates a new InvestigatorConditionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestigatorPersonId

`func (o *InvestigatorConditionsRow) GetInvestigatorPersonId() string`

GetInvestigatorPersonId returns the InvestigatorPersonId field if non-nil, zero value otherwise.

### GetInvestigatorPersonIdOk

`func (o *InvestigatorConditionsRow) GetInvestigatorPersonIdOk() (*string, bool)`

GetInvestigatorPersonIdOk returns a tuple with the InvestigatorPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonId

`func (o *InvestigatorConditionsRow) SetInvestigatorPersonId(v string)`

SetInvestigatorPersonId sets InvestigatorPersonId field to given value.

### HasInvestigatorPersonId

`func (o *InvestigatorConditionsRow) HasInvestigatorPersonId() bool`

HasInvestigatorPersonId returns a boolean if a field has been set.

### SetInvestigatorPersonIdNil

`func (o *InvestigatorConditionsRow) SetInvestigatorPersonIdNil(b bool)`

 SetInvestigatorPersonIdNil sets the value for InvestigatorPersonId to be an explicit nil

### UnsetInvestigatorPersonId
`func (o *InvestigatorConditionsRow) UnsetInvestigatorPersonId()`

UnsetInvestigatorPersonId ensures that no value is present for InvestigatorPersonId, not even an explicit nil
### GetInvestigatorPersonName

`func (o *InvestigatorConditionsRow) GetInvestigatorPersonName() string`

GetInvestigatorPersonName returns the InvestigatorPersonName field if non-nil, zero value otherwise.

### GetInvestigatorPersonNameOk

`func (o *InvestigatorConditionsRow) GetInvestigatorPersonNameOk() (*string, bool)`

GetInvestigatorPersonNameOk returns a tuple with the InvestigatorPersonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonName

`func (o *InvestigatorConditionsRow) SetInvestigatorPersonName(v string)`

SetInvestigatorPersonName sets InvestigatorPersonName field to given value.

### HasInvestigatorPersonName

`func (o *InvestigatorConditionsRow) HasInvestigatorPersonName() bool`

HasInvestigatorPersonName returns a boolean if a field has been set.

### SetInvestigatorPersonNameNil

`func (o *InvestigatorConditionsRow) SetInvestigatorPersonNameNil(b bool)`

 SetInvestigatorPersonNameNil sets the value for InvestigatorPersonName to be an explicit nil

### UnsetInvestigatorPersonName
`func (o *InvestigatorConditionsRow) UnsetInvestigatorPersonName()`

UnsetInvestigatorPersonName ensures that no value is present for InvestigatorPersonName, not even an explicit nil
### GetInvestigatorAffiliation

`func (o *InvestigatorConditionsRow) GetInvestigatorAffiliation() string`

GetInvestigatorAffiliation returns the InvestigatorAffiliation field if non-nil, zero value otherwise.

### GetInvestigatorAffiliationOk

`func (o *InvestigatorConditionsRow) GetInvestigatorAffiliationOk() (*string, bool)`

GetInvestigatorAffiliationOk returns a tuple with the InvestigatorAffiliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorAffiliation

`func (o *InvestigatorConditionsRow) SetInvestigatorAffiliation(v string)`

SetInvestigatorAffiliation sets InvestigatorAffiliation field to given value.

### HasInvestigatorAffiliation

`func (o *InvestigatorConditionsRow) HasInvestigatorAffiliation() bool`

HasInvestigatorAffiliation returns a boolean if a field has been set.

### SetInvestigatorAffiliationNil

`func (o *InvestigatorConditionsRow) SetInvestigatorAffiliationNil(b bool)`

 SetInvestigatorAffiliationNil sets the value for InvestigatorAffiliation to be an explicit nil

### UnsetInvestigatorAffiliation
`func (o *InvestigatorConditionsRow) UnsetInvestigatorAffiliation()`

UnsetInvestigatorAffiliation ensures that no value is present for InvestigatorAffiliation, not even an explicit nil
### GetConditionId

`func (o *InvestigatorConditionsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *InvestigatorConditionsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *InvestigatorConditionsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *InvestigatorConditionsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *InvestigatorConditionsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *InvestigatorConditionsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *InvestigatorConditionsRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *InvestigatorConditionsRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *InvestigatorConditionsRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *InvestigatorConditionsRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *InvestigatorConditionsRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *InvestigatorConditionsRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetNTrials

`func (o *InvestigatorConditionsRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *InvestigatorConditionsRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *InvestigatorConditionsRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *InvestigatorConditionsRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *InvestigatorConditionsRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *InvestigatorConditionsRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


