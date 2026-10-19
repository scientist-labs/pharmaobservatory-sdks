# ConditionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s MeSH heading. | [optional] 
**PrimaryCategory** | Pointer to **NullableString** | The condition&#39;s primary MeSH disease category. | [optional] 
**Categories** | Pointer to **[]string** | The MeSH tree branches the condition sits under (C and F03 codes). | [optional] 
**IsNeoplasm** | Pointer to **NullableBool** | True when the condition sits under the MeSH neoplasms branch (C04). | [optional] 
**TreeNumbers** | Pointer to **[]string** | The condition&#39;s MeSH tree numbers. | [optional] 
**NTrials** | Pointer to **NullableInt64** | Trials tagged with the condition or with a condition beneath it in the MeSH tree. | [optional] 
**NDirectTrials** | Pointer to **NullableInt64** | Trials tagged with the condition itself. | [optional] 

## Methods

### NewConditionsRow

`func NewConditionsRow() *ConditionsRow`

NewConditionsRow instantiates a new ConditionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionsRowWithDefaults

`func NewConditionsRowWithDefaults() *ConditionsRow`

NewConditionsRowWithDefaults instantiates a new ConditionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ConditionsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ConditionsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ConditionsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ConditionsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ConditionsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ConditionsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *ConditionsRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *ConditionsRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *ConditionsRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *ConditionsRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *ConditionsRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *ConditionsRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetPrimaryCategory

`func (o *ConditionsRow) GetPrimaryCategory() string`

GetPrimaryCategory returns the PrimaryCategory field if non-nil, zero value otherwise.

### GetPrimaryCategoryOk

`func (o *ConditionsRow) GetPrimaryCategoryOk() (*string, bool)`

GetPrimaryCategoryOk returns a tuple with the PrimaryCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCategory

`func (o *ConditionsRow) SetPrimaryCategory(v string)`

SetPrimaryCategory sets PrimaryCategory field to given value.

### HasPrimaryCategory

`func (o *ConditionsRow) HasPrimaryCategory() bool`

HasPrimaryCategory returns a boolean if a field has been set.

### SetPrimaryCategoryNil

`func (o *ConditionsRow) SetPrimaryCategoryNil(b bool)`

 SetPrimaryCategoryNil sets the value for PrimaryCategory to be an explicit nil

### UnsetPrimaryCategory
`func (o *ConditionsRow) UnsetPrimaryCategory()`

UnsetPrimaryCategory ensures that no value is present for PrimaryCategory, not even an explicit nil
### GetCategories

`func (o *ConditionsRow) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ConditionsRow) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ConditionsRow) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ConditionsRow) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *ConditionsRow) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *ConditionsRow) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetIsNeoplasm

`func (o *ConditionsRow) GetIsNeoplasm() bool`

GetIsNeoplasm returns the IsNeoplasm field if non-nil, zero value otherwise.

### GetIsNeoplasmOk

`func (o *ConditionsRow) GetIsNeoplasmOk() (*bool, bool)`

GetIsNeoplasmOk returns a tuple with the IsNeoplasm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNeoplasm

`func (o *ConditionsRow) SetIsNeoplasm(v bool)`

SetIsNeoplasm sets IsNeoplasm field to given value.

### HasIsNeoplasm

`func (o *ConditionsRow) HasIsNeoplasm() bool`

HasIsNeoplasm returns a boolean if a field has been set.

### SetIsNeoplasmNil

`func (o *ConditionsRow) SetIsNeoplasmNil(b bool)`

 SetIsNeoplasmNil sets the value for IsNeoplasm to be an explicit nil

### UnsetIsNeoplasm
`func (o *ConditionsRow) UnsetIsNeoplasm()`

UnsetIsNeoplasm ensures that no value is present for IsNeoplasm, not even an explicit nil
### GetTreeNumbers

`func (o *ConditionsRow) GetTreeNumbers() []string`

GetTreeNumbers returns the TreeNumbers field if non-nil, zero value otherwise.

### GetTreeNumbersOk

`func (o *ConditionsRow) GetTreeNumbersOk() (*[]string, bool)`

GetTreeNumbersOk returns a tuple with the TreeNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreeNumbers

`func (o *ConditionsRow) SetTreeNumbers(v []string)`

SetTreeNumbers sets TreeNumbers field to given value.

### HasTreeNumbers

`func (o *ConditionsRow) HasTreeNumbers() bool`

HasTreeNumbers returns a boolean if a field has been set.

### SetTreeNumbersNil

`func (o *ConditionsRow) SetTreeNumbersNil(b bool)`

 SetTreeNumbersNil sets the value for TreeNumbers to be an explicit nil

### UnsetTreeNumbers
`func (o *ConditionsRow) UnsetTreeNumbers()`

UnsetTreeNumbers ensures that no value is present for TreeNumbers, not even an explicit nil
### GetNTrials

`func (o *ConditionsRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *ConditionsRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *ConditionsRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *ConditionsRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *ConditionsRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *ConditionsRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNDirectTrials

`func (o *ConditionsRow) GetNDirectTrials() int64`

GetNDirectTrials returns the NDirectTrials field if non-nil, zero value otherwise.

### GetNDirectTrialsOk

`func (o *ConditionsRow) GetNDirectTrialsOk() (*int64, bool)`

GetNDirectTrialsOk returns a tuple with the NDirectTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNDirectTrials

`func (o *ConditionsRow) SetNDirectTrials(v int64)`

SetNDirectTrials sets NDirectTrials field to given value.

### HasNDirectTrials

`func (o *ConditionsRow) HasNDirectTrials() bool`

HasNDirectTrials returns a boolean if a field has been set.

### SetNDirectTrialsNil

`func (o *ConditionsRow) SetNDirectTrialsNil(b bool)`

 SetNDirectTrialsNil sets the value for NDirectTrials to be an explicit nil

### UnsetNDirectTrials
`func (o *ConditionsRow) UnsetNDirectTrials()`

UnsetNDirectTrials ensures that no value is present for NDirectTrials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


