# DrugLabelSafetyRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**has_boxed_warning** | Option<**bool**> | True when the label carries a boxed warning. | [optional]
**boxed_warning_text** | Option<**String**> | The boxed warning text, null when there is none. | [optional]
**warnings** | Option<**String**> | The warnings and precautions text. | [optional]
**label_effective_date** | Option<[**String**](string.md)> | The effective date of the label. | [optional]
**label_count** | Option<**i32**> | Labels matched to the drug. | [optional]
**set_id** | Option<**String**> | The label set id. The DailyMed page is https://dailymed.nlm.nih.gov/dailymed/drugInfo.cfm?setid= followed by this id. | [optional]
**match_method** | Option<**String**> | How the label was matched to the drug. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


