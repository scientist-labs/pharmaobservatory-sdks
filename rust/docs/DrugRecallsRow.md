# DrugRecallsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**recall_number** | Option<**String**> | The FDA recall number. | [optional]
**classification** | Option<**String**> | The FDA recall class (Class I, Class II or Class III). | [optional]
**status** | Option<**String**> | The recall status (for example Ongoing, Completed, Terminated). | [optional]
**recalling_firm** | Option<**String**> | The firm that recalled the product, as filed. | [optional]
**product_description** | Option<**String**> | The recalled product as described in the report. | [optional]
**reason_for_recall** | Option<**String**> | The reason given for the recall. | [optional]
**recall_initiation_date** | Option<[**String**](string.md)> | The date the recall began. | [optional]
**report_date** | Option<[**String**](string.md)> | The date the FDA posted the report. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


