# RareDiseasesRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sections** | Option<[**std::collections::HashMap<String, serde_json::Value>**](serde_json::Value.md)> | Sections of the object's web page, by key: patient_organisations, the list the page shows of the organisations that name the disease or condition on their own websites (state, empty_state, coverage, groups of entries carrying the patient_organisation dataset's columns, the per-entry credit and the lane attribution). | [optional]
**object_type** | **String** |  | 
**id** | **String** |  | 
**record** | [**models::RareDiseasesRow**](RareDiseasesRow.md) |  | 
**attributes** | Option<[**models::RareDiseasesRecordAttributes**](RareDiseasesRecord_attributes.md)> |  | [optional]
**datasets** | [**Vec<models::DatasetListing>**](DatasetListing.md) |  | 
**attribution** | Option<[**Vec<models::DatasetAttribution>**](DatasetAttribution.md)> |  | [optional]
**source_credit** | Option<**String**> |  | [optional]
**freshness** | Option<[**models::AnimalWelfareAssurancesResponseFreshness**](AnimalWelfareAssurancesResponse_freshness.md)> |  | [optional]
**build** | Option<[**models::AnimalWelfareAssurancesResponseBuild**](AnimalWelfareAssurancesResponse_build.md)> |  | [optional]
**build_state** | Option<[**models::DatasetBuildState**](DatasetBuildState.md)> |  | [optional]
**restriction** | Option<[**std::collections::HashMap<String, serde_json::Value>**](serde_json::Value.md)> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


