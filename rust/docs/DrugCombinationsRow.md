# DrugCombinationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**relation** | Option<**String**> | has_component when the related drug is a component of this drug; component_of when this drug is a component of the related drug. | [optional]
**related_drug_id** | Option<**String**> | The molecule key of the related drug. | [optional]
**related_drug_name** | Option<**String**> | The related drug's display name. | [optional]
**combination_id** | Option<**String**> | A stable id for this drug, direction and related drug. | [optional]
**n_components** | Option<**i32**> | Components in the combination product. | [optional]
**combo_unii_mixture** | Option<**bool**> | True when the registry itself lists the combination as a mixture substance. | [optional]
**has_ingredient_confirmed** | Option<**bool**> | True when RxNorm's ingredient list confirms the components. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


