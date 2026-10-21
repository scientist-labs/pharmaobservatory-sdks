# DrugApprovalsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**authority** | Option<**String**> | The approving authority. | [optional]
**jurisdiction** | Option<**String**> | The jurisdiction the authority covers. | [optional]
**status** | Option<**String**> | The approval status the source states. | [optional]
**first_approval_date** | Option<[**String**](string.md)> | The earliest authorisation or first-marketed date the source gives. | [optional]
**approval_id** | Option<**String**> | The product number the authority uses (EMA product number or Health Canada drug identification number). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


