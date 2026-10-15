# ApiV1RwdStudyScreenGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**filters** | [**models::ApiV1RwdStudyScreenGet200ResponseFilters**](_api_v1_rwd_study_screen_get_200_response_filters.md) |  | 
**matched_studies** | Option<**i32**> |  | 
**register_total** | Option<**i32**> |  | 
**delivery_distribution** | Option<[**models::ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution**](_api_v1_rwd_study_screen_get_200_response_delivery_distribution.md)> |  | 
**coverage_caveat** | **String** |  | 
**tier_provisional_note** | **String** |  | 
**build_state** | [**models::ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](_api_v1_biomarker_assay_screen_get_200_response_build_state.md) |  | 
**restriction** | [**models::ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction**](_api_v1_conditions__condition_id__burden_vs_evidence_get_200_response_restriction.md) |  | 
**condition_filter_built** | Option<**bool**> | false means a condition was named and the condition mart is not built in this warehouse, so the filter could not run and matched_studies, register_total and delivery_distribution are unknown (null), not zero. true means the filter ran, or no condition was named. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


