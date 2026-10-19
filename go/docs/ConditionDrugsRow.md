# ConditionDrugsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**DrugId** | Pointer to **NullableString** | The drug key (a canonical drug id; the same value as molecule_id on this dataset). | [optional] 
**MoleculeId** | Pointer to **NullableString** | The molecule key the drug resolves to. | [optional] 
**DrugName** | Pointer to **NullableString** | The molecule&#39;s display name. | [optional] 
**NTrials** | Pointer to **NullableInt64** | Trials for this condition that test the molecule. | [optional] 

## Methods

### NewConditionDrugsRow

`func NewConditionDrugsRow() *ConditionDrugsRow`

NewConditionDrugsRow instantiates a new ConditionDrugsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionDrugsRowWithDefaults

`func NewConditionDrugsRowWithDefaults() *ConditionDrugsRow`

NewConditionDrugsRowWithDefaults instantiates a new ConditionDrugsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ConditionDrugsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ConditionDrugsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ConditionDrugsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ConditionDrugsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ConditionDrugsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ConditionDrugsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetDrugId

`func (o *ConditionDrugsRow) GetDrugId() string`

GetDrugId returns the DrugId field if non-nil, zero value otherwise.

### GetDrugIdOk

`func (o *ConditionDrugsRow) GetDrugIdOk() (*string, bool)`

GetDrugIdOk returns a tuple with the DrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugId

`func (o *ConditionDrugsRow) SetDrugId(v string)`

SetDrugId sets DrugId field to given value.

### HasDrugId

`func (o *ConditionDrugsRow) HasDrugId() bool`

HasDrugId returns a boolean if a field has been set.

### SetDrugIdNil

`func (o *ConditionDrugsRow) SetDrugIdNil(b bool)`

 SetDrugIdNil sets the value for DrugId to be an explicit nil

### UnsetDrugId
`func (o *ConditionDrugsRow) UnsetDrugId()`

UnsetDrugId ensures that no value is present for DrugId, not even an explicit nil
### GetMoleculeId

`func (o *ConditionDrugsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *ConditionDrugsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *ConditionDrugsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *ConditionDrugsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *ConditionDrugsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *ConditionDrugsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetDrugName

`func (o *ConditionDrugsRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *ConditionDrugsRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *ConditionDrugsRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *ConditionDrugsRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *ConditionDrugsRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *ConditionDrugsRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetNTrials

`func (o *ConditionDrugsRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *ConditionDrugsRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *ConditionDrugsRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *ConditionDrugsRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *ConditionDrugsRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *ConditionDrugsRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


