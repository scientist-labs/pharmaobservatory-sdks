# ConditionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**condition_id** | Option<**String**> | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional]
**condition_name** | Option<**String**> | The condition's MeSH heading. | [optional]
**primary_category** | Option<**String**> | The condition's primary MeSH disease category. | [optional]
**categories** | Option<**Vec<String>**> | The MeSH tree branches the condition sits under (C and F03 codes). | [optional]
**is_neoplasm** | Option<**bool**> | True when the condition sits under the MeSH neoplasms branch (C04). | [optional]
**tree_numbers** | Option<**Vec<String>**> | The condition's MeSH tree numbers. | [optional]
**n_trials** | Option<**i64**> | Trials tagged with the condition or with a condition beneath it in the MeSH tree. | [optional]
**n_direct_trials** | Option<**i64**> | Trials tagged with the condition itself. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


