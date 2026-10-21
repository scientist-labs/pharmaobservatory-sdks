# DrugMedicareSpendRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**year** | Option<**i32**> | The calendar year of the spending. | [optional]
**total_spend** | Option<**f64**> | Medicare spending in US dollars. | [optional]
**claims** | Option<**f64**> | Claims or services. | [optional]
**beneficiaries** | Option<**f64**> | Beneficiaries, summed over payers. | [optional]
**payers** | Option<**Vec<String>**> | The Medicare programs summed (Part B, Part D). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


