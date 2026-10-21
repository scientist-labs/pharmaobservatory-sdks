# DrugPricingRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**PricingUnit** | Pointer to **NullableString** | The unit the costs are per (EA each, ML milliliter, GM gram). | [optional] 
**PricedNdcs** | Pointer to **NullableInt32** | Packages (National Drug Codes) priced. | [optional] 
**MinUnitCost** | Pointer to **NullableFloat32** | The lowest cost per unit. | [optional] 
**MedianUnitCost** | Pointer to **NullableFloat32** | The median cost per unit. | [optional] 
**MaxUnitCost** | Pointer to **NullableFloat32** | The highest cost per unit. | [optional] 
**HasBrandPriced** | Pointer to **NullableBool** | True when a brand package is priced. | [optional] 
**HasGenericPriced** | Pointer to **NullableBool** | True when a generic package is priced. | [optional] 

## Methods

### NewDrugPricingRow

`func NewDrugPricingRow() *DrugPricingRow`

NewDrugPricingRow instantiates a new DrugPricingRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugPricingRowWithDefaults

`func NewDrugPricingRowWithDefaults() *DrugPricingRow`

NewDrugPricingRowWithDefaults instantiates a new DrugPricingRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugPricingRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugPricingRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugPricingRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugPricingRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugPricingRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugPricingRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetPricingUnit

`func (o *DrugPricingRow) GetPricingUnit() string`

GetPricingUnit returns the PricingUnit field if non-nil, zero value otherwise.

### GetPricingUnitOk

`func (o *DrugPricingRow) GetPricingUnitOk() (*string, bool)`

GetPricingUnitOk returns a tuple with the PricingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingUnit

`func (o *DrugPricingRow) SetPricingUnit(v string)`

SetPricingUnit sets PricingUnit field to given value.

### HasPricingUnit

`func (o *DrugPricingRow) HasPricingUnit() bool`

HasPricingUnit returns a boolean if a field has been set.

### SetPricingUnitNil

`func (o *DrugPricingRow) SetPricingUnitNil(b bool)`

 SetPricingUnitNil sets the value for PricingUnit to be an explicit nil

### UnsetPricingUnit
`func (o *DrugPricingRow) UnsetPricingUnit()`

UnsetPricingUnit ensures that no value is present for PricingUnit, not even an explicit nil
### GetPricedNdcs

`func (o *DrugPricingRow) GetPricedNdcs() int32`

GetPricedNdcs returns the PricedNdcs field if non-nil, zero value otherwise.

### GetPricedNdcsOk

`func (o *DrugPricingRow) GetPricedNdcsOk() (*int32, bool)`

GetPricedNdcsOk returns a tuple with the PricedNdcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricedNdcs

`func (o *DrugPricingRow) SetPricedNdcs(v int32)`

SetPricedNdcs sets PricedNdcs field to given value.

### HasPricedNdcs

`func (o *DrugPricingRow) HasPricedNdcs() bool`

HasPricedNdcs returns a boolean if a field has been set.

### SetPricedNdcsNil

`func (o *DrugPricingRow) SetPricedNdcsNil(b bool)`

 SetPricedNdcsNil sets the value for PricedNdcs to be an explicit nil

### UnsetPricedNdcs
`func (o *DrugPricingRow) UnsetPricedNdcs()`

UnsetPricedNdcs ensures that no value is present for PricedNdcs, not even an explicit nil
### GetMinUnitCost

`func (o *DrugPricingRow) GetMinUnitCost() float32`

GetMinUnitCost returns the MinUnitCost field if non-nil, zero value otherwise.

### GetMinUnitCostOk

`func (o *DrugPricingRow) GetMinUnitCostOk() (*float32, bool)`

GetMinUnitCostOk returns a tuple with the MinUnitCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUnitCost

`func (o *DrugPricingRow) SetMinUnitCost(v float32)`

SetMinUnitCost sets MinUnitCost field to given value.

### HasMinUnitCost

`func (o *DrugPricingRow) HasMinUnitCost() bool`

HasMinUnitCost returns a boolean if a field has been set.

### SetMinUnitCostNil

`func (o *DrugPricingRow) SetMinUnitCostNil(b bool)`

 SetMinUnitCostNil sets the value for MinUnitCost to be an explicit nil

### UnsetMinUnitCost
`func (o *DrugPricingRow) UnsetMinUnitCost()`

UnsetMinUnitCost ensures that no value is present for MinUnitCost, not even an explicit nil
### GetMedianUnitCost

`func (o *DrugPricingRow) GetMedianUnitCost() float32`

GetMedianUnitCost returns the MedianUnitCost field if non-nil, zero value otherwise.

### GetMedianUnitCostOk

`func (o *DrugPricingRow) GetMedianUnitCostOk() (*float32, bool)`

GetMedianUnitCostOk returns a tuple with the MedianUnitCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianUnitCost

`func (o *DrugPricingRow) SetMedianUnitCost(v float32)`

SetMedianUnitCost sets MedianUnitCost field to given value.

### HasMedianUnitCost

`func (o *DrugPricingRow) HasMedianUnitCost() bool`

HasMedianUnitCost returns a boolean if a field has been set.

### SetMedianUnitCostNil

`func (o *DrugPricingRow) SetMedianUnitCostNil(b bool)`

 SetMedianUnitCostNil sets the value for MedianUnitCost to be an explicit nil

### UnsetMedianUnitCost
`func (o *DrugPricingRow) UnsetMedianUnitCost()`

UnsetMedianUnitCost ensures that no value is present for MedianUnitCost, not even an explicit nil
### GetMaxUnitCost

`func (o *DrugPricingRow) GetMaxUnitCost() float32`

GetMaxUnitCost returns the MaxUnitCost field if non-nil, zero value otherwise.

### GetMaxUnitCostOk

`func (o *DrugPricingRow) GetMaxUnitCostOk() (*float32, bool)`

GetMaxUnitCostOk returns a tuple with the MaxUnitCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnitCost

`func (o *DrugPricingRow) SetMaxUnitCost(v float32)`

SetMaxUnitCost sets MaxUnitCost field to given value.

### HasMaxUnitCost

`func (o *DrugPricingRow) HasMaxUnitCost() bool`

HasMaxUnitCost returns a boolean if a field has been set.

### SetMaxUnitCostNil

`func (o *DrugPricingRow) SetMaxUnitCostNil(b bool)`

 SetMaxUnitCostNil sets the value for MaxUnitCost to be an explicit nil

### UnsetMaxUnitCost
`func (o *DrugPricingRow) UnsetMaxUnitCost()`

UnsetMaxUnitCost ensures that no value is present for MaxUnitCost, not even an explicit nil
### GetHasBrandPriced

`func (o *DrugPricingRow) GetHasBrandPriced() bool`

GetHasBrandPriced returns the HasBrandPriced field if non-nil, zero value otherwise.

### GetHasBrandPricedOk

`func (o *DrugPricingRow) GetHasBrandPricedOk() (*bool, bool)`

GetHasBrandPricedOk returns a tuple with the HasBrandPriced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBrandPriced

`func (o *DrugPricingRow) SetHasBrandPriced(v bool)`

SetHasBrandPriced sets HasBrandPriced field to given value.

### HasHasBrandPriced

`func (o *DrugPricingRow) HasHasBrandPriced() bool`

HasHasBrandPriced returns a boolean if a field has been set.

### SetHasBrandPricedNil

`func (o *DrugPricingRow) SetHasBrandPricedNil(b bool)`

 SetHasBrandPricedNil sets the value for HasBrandPriced to be an explicit nil

### UnsetHasBrandPriced
`func (o *DrugPricingRow) UnsetHasBrandPriced()`

UnsetHasBrandPriced ensures that no value is present for HasBrandPriced, not even an explicit nil
### GetHasGenericPriced

`func (o *DrugPricingRow) GetHasGenericPriced() bool`

GetHasGenericPriced returns the HasGenericPriced field if non-nil, zero value otherwise.

### GetHasGenericPricedOk

`func (o *DrugPricingRow) GetHasGenericPricedOk() (*bool, bool)`

GetHasGenericPricedOk returns a tuple with the HasGenericPriced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGenericPriced

`func (o *DrugPricingRow) SetHasGenericPriced(v bool)`

SetHasGenericPriced sets HasGenericPriced field to given value.

### HasHasGenericPriced

`func (o *DrugPricingRow) HasHasGenericPriced() bool`

HasHasGenericPriced returns a boolean if a field has been set.

### SetHasGenericPricedNil

`func (o *DrugPricingRow) SetHasGenericPricedNil(b bool)`

 SetHasGenericPricedNil sets the value for HasGenericPriced to be an explicit nil

### UnsetHasGenericPriced
`func (o *DrugPricingRow) UnsetHasGenericPriced()`

UnsetHasGenericPriced ensures that no value is present for HasGenericPriced, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


