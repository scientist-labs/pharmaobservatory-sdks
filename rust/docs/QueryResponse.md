# QueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**columns** | [**Vec<models::QueryResponseColumnsInner>**](QueryResponse_columns_inner.md) |  | 
**rows** | [**Vec<Vec<serde_json::Value>>**](Vec.md) |  | 
**total_rows** | Option<**i32**> |  | [optional]
**truncated** | Option<**bool**> |  | [optional]
**result_id** | **String** |  | 
**next_cursor** | Option<**String**> |  | [optional]
**attribution** | Option<[**Vec<models::DatasetAttribution>**](DatasetAttribution.md)> |  | [optional]
**source_credit** | Option<**String**> |  | [optional]
**withheld_tiers** | Option<**Vec<String>**> |  | [optional]
**withheld_columns** | Option<**Vec<String>**> |  | [optional]
**build_state** | Option<[**models::DatasetBuildState**](DatasetBuildState.md)> |  | [optional]
**dataset** | **String** |  | 
**notice** | Option<**String**> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


