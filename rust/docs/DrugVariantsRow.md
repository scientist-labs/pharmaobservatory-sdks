# DrugVariantsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**member_id** | Option<**String**> | The id of the folded drug record: the drug_id other datasets carry. | [optional]
**member_name** | Option<**String**> | The folded record's display name. | [optional]
**role** | Option<**String**> | single for a one-to-one fold; component_of_compound for a combination arm listing the molecule. | [optional]
**method** | Option<**String**> | How the fold was made (for example rxnorm_synonym, unii_synonym, brand, override, registered_substance). | [optional]
**tier** | Option<**i32**> | The resolver tier that made the fold (1 to 4); 4 is the model backstop. | [optional]
**confidence** | Option<**f64**> | The resolver's confidence for the method, from 0 to 1. | [optional]
**member_trials** | Option<**i32**> | Registered trials of the member record. | [optional]
**member_max_phase** | Option<**i32**> | The member's highest development phase. | [optional]
**member_approved** | Option<**bool**> | True when the member is approved somewhere. | [optional]
**member_biologic** | Option<**bool**> | True when the member is a biologic substance. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


