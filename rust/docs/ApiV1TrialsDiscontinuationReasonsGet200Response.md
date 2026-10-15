# ApiV1TrialsDiscontinuationReasonsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**error** | Option<**String**> |  | [optional]
**message** | Option<**String**> |  | [optional]
**build_state** | [**models::ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](_api_v1_biomarker_assay_screen_get_200_response_build_state.md) |  | 
**filters** | Option<[**models::ApiV1TrialsDiscontinuationReasonsGet200ResponseFilters**](_api_v1_trials_discontinuation_reasons_get_200_response_filters.md)> |  | [optional]
**matched_count** | Option<**i32**> |  | [optional]
**trials_with_why_stopped** | Option<**i32**> |  | [optional]
**why_stopped_coverage_rate** | Option<**f32**> |  | [optional]
**coverage_caveat** | Option<**String**> |  | [optional]
**trials** | Option<[**Vec<models::ApiV1TrialsDiscontinuationReasonsGet200ResponseTrialsInner>**](_api_v1_trials_discontinuation_reasons_get_200_response_trials_inner.md)> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


