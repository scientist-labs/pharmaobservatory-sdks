# TrialDrugClassesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**DrugId** | Pointer to **NullableString** | The canonical drug key the trial&#39;s intervention names matched. | [optional] 
**DrugName** | Pointer to **NullableString** | The drug&#39;s display name. | [optional] 
**ResolvedDrugId** | Pointer to **NullableString** | The drug key after same-substance drug records are merged. | [optional] 
**ApprovedAnywhere** | Pointer to **NullableBool** | True when the drug is marked approved by any regulator (false means investigational or unknown). | [optional] 
**AtcClass** | Pointer to **NullableString** | The ATC first-level class of the drug (null when it has none). | [optional] 

## Methods

### NewTrialDrugClassesRow

`func NewTrialDrugClassesRow() *TrialDrugClassesRow`

NewTrialDrugClassesRow instantiates a new TrialDrugClassesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialDrugClassesRowWithDefaults

`func NewTrialDrugClassesRowWithDefaults() *TrialDrugClassesRow`

NewTrialDrugClassesRowWithDefaults instantiates a new TrialDrugClassesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *TrialDrugClassesRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TrialDrugClassesRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TrialDrugClassesRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TrialDrugClassesRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TrialDrugClassesRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TrialDrugClassesRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetDrugId

`func (o *TrialDrugClassesRow) GetDrugId() string`

GetDrugId returns the DrugId field if non-nil, zero value otherwise.

### GetDrugIdOk

`func (o *TrialDrugClassesRow) GetDrugIdOk() (*string, bool)`

GetDrugIdOk returns a tuple with the DrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugId

`func (o *TrialDrugClassesRow) SetDrugId(v string)`

SetDrugId sets DrugId field to given value.

### HasDrugId

`func (o *TrialDrugClassesRow) HasDrugId() bool`

HasDrugId returns a boolean if a field has been set.

### SetDrugIdNil

`func (o *TrialDrugClassesRow) SetDrugIdNil(b bool)`

 SetDrugIdNil sets the value for DrugId to be an explicit nil

### UnsetDrugId
`func (o *TrialDrugClassesRow) UnsetDrugId()`

UnsetDrugId ensures that no value is present for DrugId, not even an explicit nil
### GetDrugName

`func (o *TrialDrugClassesRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *TrialDrugClassesRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *TrialDrugClassesRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *TrialDrugClassesRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *TrialDrugClassesRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *TrialDrugClassesRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetResolvedDrugId

`func (o *TrialDrugClassesRow) GetResolvedDrugId() string`

GetResolvedDrugId returns the ResolvedDrugId field if non-nil, zero value otherwise.

### GetResolvedDrugIdOk

`func (o *TrialDrugClassesRow) GetResolvedDrugIdOk() (*string, bool)`

GetResolvedDrugIdOk returns a tuple with the ResolvedDrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDrugId

`func (o *TrialDrugClassesRow) SetResolvedDrugId(v string)`

SetResolvedDrugId sets ResolvedDrugId field to given value.

### HasResolvedDrugId

`func (o *TrialDrugClassesRow) HasResolvedDrugId() bool`

HasResolvedDrugId returns a boolean if a field has been set.

### SetResolvedDrugIdNil

`func (o *TrialDrugClassesRow) SetResolvedDrugIdNil(b bool)`

 SetResolvedDrugIdNil sets the value for ResolvedDrugId to be an explicit nil

### UnsetResolvedDrugId
`func (o *TrialDrugClassesRow) UnsetResolvedDrugId()`

UnsetResolvedDrugId ensures that no value is present for ResolvedDrugId, not even an explicit nil
### GetApprovedAnywhere

`func (o *TrialDrugClassesRow) GetApprovedAnywhere() bool`

GetApprovedAnywhere returns the ApprovedAnywhere field if non-nil, zero value otherwise.

### GetApprovedAnywhereOk

`func (o *TrialDrugClassesRow) GetApprovedAnywhereOk() (*bool, bool)`

GetApprovedAnywhereOk returns a tuple with the ApprovedAnywhere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAnywhere

`func (o *TrialDrugClassesRow) SetApprovedAnywhere(v bool)`

SetApprovedAnywhere sets ApprovedAnywhere field to given value.

### HasApprovedAnywhere

`func (o *TrialDrugClassesRow) HasApprovedAnywhere() bool`

HasApprovedAnywhere returns a boolean if a field has been set.

### SetApprovedAnywhereNil

`func (o *TrialDrugClassesRow) SetApprovedAnywhereNil(b bool)`

 SetApprovedAnywhereNil sets the value for ApprovedAnywhere to be an explicit nil

### UnsetApprovedAnywhere
`func (o *TrialDrugClassesRow) UnsetApprovedAnywhere()`

UnsetApprovedAnywhere ensures that no value is present for ApprovedAnywhere, not even an explicit nil
### GetAtcClass

`func (o *TrialDrugClassesRow) GetAtcClass() string`

GetAtcClass returns the AtcClass field if non-nil, zero value otherwise.

### GetAtcClassOk

`func (o *TrialDrugClassesRow) GetAtcClassOk() (*string, bool)`

GetAtcClassOk returns a tuple with the AtcClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtcClass

`func (o *TrialDrugClassesRow) SetAtcClass(v string)`

SetAtcClass sets AtcClass field to given value.

### HasAtcClass

`func (o *TrialDrugClassesRow) HasAtcClass() bool`

HasAtcClass returns a boolean if a field has been set.

### SetAtcClassNil

`func (o *TrialDrugClassesRow) SetAtcClassNil(b bool)`

 SetAtcClassNil sets the value for AtcClass to be an explicit nil

### UnsetAtcClass
`func (o *TrialDrugClassesRow) UnsetAtcClass()`

UnsetAtcClass ensures that no value is present for AtcClass, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


