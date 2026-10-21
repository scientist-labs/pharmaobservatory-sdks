# DrugCombinationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**Relation** | Pointer to **NullableString** | has_component when the related drug is a component of this drug; component_of when this drug is a component of the related drug. | [optional] 
**RelatedDrugId** | Pointer to **NullableString** | The molecule key of the related drug. | [optional] 
**RelatedDrugName** | Pointer to **NullableString** | The related drug&#39;s display name. | [optional] 
**CombinationId** | Pointer to **NullableString** | A stable id for this drug, direction and related drug. | [optional] 
**NComponents** | Pointer to **NullableInt32** | Components in the combination product. | [optional] 
**ComboUniiMixture** | Pointer to **NullableBool** | True when the registry itself lists the combination as a mixture substance. | [optional] 
**HasIngredientConfirmed** | Pointer to **NullableBool** | True when RxNorm&#39;s ingredient list confirms the components. | [optional] 

## Methods

### NewDrugCombinationsRow

`func NewDrugCombinationsRow() *DrugCombinationsRow`

NewDrugCombinationsRow instantiates a new DrugCombinationsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugCombinationsRowWithDefaults

`func NewDrugCombinationsRowWithDefaults() *DrugCombinationsRow`

NewDrugCombinationsRowWithDefaults instantiates a new DrugCombinationsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugCombinationsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugCombinationsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugCombinationsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugCombinationsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugCombinationsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugCombinationsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetRelation

`func (o *DrugCombinationsRow) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *DrugCombinationsRow) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *DrugCombinationsRow) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *DrugCombinationsRow) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### SetRelationNil

`func (o *DrugCombinationsRow) SetRelationNil(b bool)`

 SetRelationNil sets the value for Relation to be an explicit nil

### UnsetRelation
`func (o *DrugCombinationsRow) UnsetRelation()`

UnsetRelation ensures that no value is present for Relation, not even an explicit nil
### GetRelatedDrugId

`func (o *DrugCombinationsRow) GetRelatedDrugId() string`

GetRelatedDrugId returns the RelatedDrugId field if non-nil, zero value otherwise.

### GetRelatedDrugIdOk

`func (o *DrugCombinationsRow) GetRelatedDrugIdOk() (*string, bool)`

GetRelatedDrugIdOk returns a tuple with the RelatedDrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDrugId

`func (o *DrugCombinationsRow) SetRelatedDrugId(v string)`

SetRelatedDrugId sets RelatedDrugId field to given value.

### HasRelatedDrugId

`func (o *DrugCombinationsRow) HasRelatedDrugId() bool`

HasRelatedDrugId returns a boolean if a field has been set.

### SetRelatedDrugIdNil

`func (o *DrugCombinationsRow) SetRelatedDrugIdNil(b bool)`

 SetRelatedDrugIdNil sets the value for RelatedDrugId to be an explicit nil

### UnsetRelatedDrugId
`func (o *DrugCombinationsRow) UnsetRelatedDrugId()`

UnsetRelatedDrugId ensures that no value is present for RelatedDrugId, not even an explicit nil
### GetRelatedDrugName

`func (o *DrugCombinationsRow) GetRelatedDrugName() string`

GetRelatedDrugName returns the RelatedDrugName field if non-nil, zero value otherwise.

### GetRelatedDrugNameOk

`func (o *DrugCombinationsRow) GetRelatedDrugNameOk() (*string, bool)`

GetRelatedDrugNameOk returns a tuple with the RelatedDrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDrugName

`func (o *DrugCombinationsRow) SetRelatedDrugName(v string)`

SetRelatedDrugName sets RelatedDrugName field to given value.

### HasRelatedDrugName

`func (o *DrugCombinationsRow) HasRelatedDrugName() bool`

HasRelatedDrugName returns a boolean if a field has been set.

### SetRelatedDrugNameNil

`func (o *DrugCombinationsRow) SetRelatedDrugNameNil(b bool)`

 SetRelatedDrugNameNil sets the value for RelatedDrugName to be an explicit nil

### UnsetRelatedDrugName
`func (o *DrugCombinationsRow) UnsetRelatedDrugName()`

UnsetRelatedDrugName ensures that no value is present for RelatedDrugName, not even an explicit nil
### GetCombinationId

`func (o *DrugCombinationsRow) GetCombinationId() string`

GetCombinationId returns the CombinationId field if non-nil, zero value otherwise.

### GetCombinationIdOk

`func (o *DrugCombinationsRow) GetCombinationIdOk() (*string, bool)`

GetCombinationIdOk returns a tuple with the CombinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinationId

`func (o *DrugCombinationsRow) SetCombinationId(v string)`

SetCombinationId sets CombinationId field to given value.

### HasCombinationId

`func (o *DrugCombinationsRow) HasCombinationId() bool`

HasCombinationId returns a boolean if a field has been set.

### SetCombinationIdNil

`func (o *DrugCombinationsRow) SetCombinationIdNil(b bool)`

 SetCombinationIdNil sets the value for CombinationId to be an explicit nil

### UnsetCombinationId
`func (o *DrugCombinationsRow) UnsetCombinationId()`

UnsetCombinationId ensures that no value is present for CombinationId, not even an explicit nil
### GetNComponents

`func (o *DrugCombinationsRow) GetNComponents() int32`

GetNComponents returns the NComponents field if non-nil, zero value otherwise.

### GetNComponentsOk

`func (o *DrugCombinationsRow) GetNComponentsOk() (*int32, bool)`

GetNComponentsOk returns a tuple with the NComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNComponents

`func (o *DrugCombinationsRow) SetNComponents(v int32)`

SetNComponents sets NComponents field to given value.

### HasNComponents

`func (o *DrugCombinationsRow) HasNComponents() bool`

HasNComponents returns a boolean if a field has been set.

### SetNComponentsNil

`func (o *DrugCombinationsRow) SetNComponentsNil(b bool)`

 SetNComponentsNil sets the value for NComponents to be an explicit nil

### UnsetNComponents
`func (o *DrugCombinationsRow) UnsetNComponents()`

UnsetNComponents ensures that no value is present for NComponents, not even an explicit nil
### GetComboUniiMixture

`func (o *DrugCombinationsRow) GetComboUniiMixture() bool`

GetComboUniiMixture returns the ComboUniiMixture field if non-nil, zero value otherwise.

### GetComboUniiMixtureOk

`func (o *DrugCombinationsRow) GetComboUniiMixtureOk() (*bool, bool)`

GetComboUniiMixtureOk returns a tuple with the ComboUniiMixture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComboUniiMixture

`func (o *DrugCombinationsRow) SetComboUniiMixture(v bool)`

SetComboUniiMixture sets ComboUniiMixture field to given value.

### HasComboUniiMixture

`func (o *DrugCombinationsRow) HasComboUniiMixture() bool`

HasComboUniiMixture returns a boolean if a field has been set.

### SetComboUniiMixtureNil

`func (o *DrugCombinationsRow) SetComboUniiMixtureNil(b bool)`

 SetComboUniiMixtureNil sets the value for ComboUniiMixture to be an explicit nil

### UnsetComboUniiMixture
`func (o *DrugCombinationsRow) UnsetComboUniiMixture()`

UnsetComboUniiMixture ensures that no value is present for ComboUniiMixture, not even an explicit nil
### GetHasIngredientConfirmed

`func (o *DrugCombinationsRow) GetHasIngredientConfirmed() bool`

GetHasIngredientConfirmed returns the HasIngredientConfirmed field if non-nil, zero value otherwise.

### GetHasIngredientConfirmedOk

`func (o *DrugCombinationsRow) GetHasIngredientConfirmedOk() (*bool, bool)`

GetHasIngredientConfirmedOk returns a tuple with the HasIngredientConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIngredientConfirmed

`func (o *DrugCombinationsRow) SetHasIngredientConfirmed(v bool)`

SetHasIngredientConfirmed sets HasIngredientConfirmed field to given value.

### HasHasIngredientConfirmed

`func (o *DrugCombinationsRow) HasHasIngredientConfirmed() bool`

HasHasIngredientConfirmed returns a boolean if a field has been set.

### SetHasIngredientConfirmedNil

`func (o *DrugCombinationsRow) SetHasIngredientConfirmedNil(b bool)`

 SetHasIngredientConfirmedNil sets the value for HasIngredientConfirmed to be an explicit nil

### UnsetHasIngredientConfirmed
`func (o *DrugCombinationsRow) UnsetHasIngredientConfirmed()`

UnsetHasIngredientConfirmed ensures that no value is present for HasIngredientConfirmed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


