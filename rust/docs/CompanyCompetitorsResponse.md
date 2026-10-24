# CompanyCompetitorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**dataset** | **String** |  | 
**grain** | Option<**String**> |  | [optional]
**object_type** | Option<**String**> |  | [optional]
**id** | Option<**String**> |  | [optional]
**scope** | Option<[**models::AnimalWelfareAssurancesResponseScope**](AnimalWelfareAssurancesResponse_scope.md)> |  | [optional]
**columns** | [**Vec<models::DatasetColumn>**](DatasetColumn.md) |  | 
**rows** | [**Vec<models::CompanyCompetitorsRow>**](CompanyCompetitorsRow.md) |  | 
**total** | Option<**i32**> | The exact count of the filtered rows on an ungrouped read, carried by the cursor; null when grouped. | [optional]
**next_cursor** | Option<**String**> |  | [optional]
**sort** | Option<**Vec<String>**> |  | [optional]
**filters** | Option<[**Vec<std::collections::HashMap<String, serde_json::Value>>**](std::collections::HashMap.md)> |  | [optional]
**group_by** | Option<**Vec<String>**> |  | [optional]
**measures** | Option<**Vec<String>**> |  | [optional]
**attribution** | Option<[**Vec<models::DatasetAttribution>**](DatasetAttribution.md)> |  | [optional]
**source_credit** | Option<**String**> |  | [optional]
**withheld_tiers** | Option<**Vec<String>**> |  | [optional]
**withheld_columns** | Option<**Vec<String>**> |  | [optional]
**freshness** | Option<[**models::AnimalWelfareAssurancesResponseFreshness**](AnimalWelfareAssurancesResponse_freshness.md)> |  | [optional]
**build** | Option<[**models::AnimalWelfareAssurancesResponseBuild**](AnimalWelfareAssurancesResponse_build.md)> |  | [optional]
**build_state** | Option<[**models::DatasetBuildState**](DatasetBuildState.md)> |  | [optional]
**restriction** | Option<[**std::collections::HashMap<String, serde_json::Value>**](serde_json::Value.md)> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


