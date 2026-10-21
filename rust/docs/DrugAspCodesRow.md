# DrugAspCodesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**hcpcs** | Option<**String**> | The HCPCS billing code. | [optional]
**hcpcs_desc** | Option<**String**> | The code's short description. | [optional]
**hcpcs_dosage** | Option<**String**> | The billing unit the code is priced per. | [optional]
**asp_payment_limit** | Option<**f64**> | The Medicare payment limit per billing unit. | [optional]
**priced_ndcs** | Option<**i32**> | Packages priced under the code. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


