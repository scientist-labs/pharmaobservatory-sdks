# TargetsRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**object_type** | **String** |  | 
**id** | **String** |  | 
**record** | [**models::TargetsRow**](TargetsRow.md) |  | 
**attributes** | Option<[**serde_json::Value**](.md)> |  | [optional]
**datasets** | [**Vec<models::DatasetListing>**](DatasetListing.md) |  | 
**attribution** | Option<[**Vec<models::DatasetAttribution>**](DatasetAttribution.md)> |  | [optional]
**source_credit** | Option<**String**> |  | [optional]
**freshness** | Option<[**models::BiomarkerAssayEvidenceResponseFreshness**](BiomarkerAssayEvidenceResponse_freshness.md)> |  | [optional]
**build** | Option<[**models::BiomarkerAssayEvidenceResponseBuild**](BiomarkerAssayEvidenceResponse_build.md)> |  | [optional]
**build_state** | Option<[**models::DatasetBuildState**](DatasetBuildState.md)> |  | [optional]
**restriction** | Option<[**std::collections::HashMap<String, serde_json::Value>**](serde_json::Value.md)> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


