# DrugPricingRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**pricing_unit** | Option<**String**> | The unit the costs are per (EA each, ML milliliter, GM gram). | [optional]
**priced_ndcs** | Option<**i32**> | Packages (National Drug Codes) priced. | [optional]
**min_unit_cost** | Option<**f64**> | The lowest cost per unit. | [optional]
**median_unit_cost** | Option<**f64**> | The median cost per unit. | [optional]
**max_unit_cost** | Option<**f64**> | The highest cost per unit. | [optional]
**has_brand_priced** | Option<**bool**> | True when a brand package is priced. | [optional]
**has_generic_priced** | Option<**bool**> | True when a generic package is priced. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


