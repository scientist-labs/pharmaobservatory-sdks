# ApiV1OrgPipelineStagesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**company_id** | Option<**String**> |  | [optional]
**name** | Option<**String**> |  | [optional]
**approved_marketed** | [**Vec<models::ApiV1OrgPipelineStagesGet200ResponseApprovedMarketedInner>**](_api_v1_org_pipeline_stages_get_200_response_approved_marketed_inner.md) |  | 
**in_development** | [**Vec<models::ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner>**](_api_v1_org_pipeline_stages_get_200_response_in_development_inner.md) |  | 
**abandoned_dormant** | [**Vec<models::ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner>**](_api_v1_org_pipeline_stages_get_200_response_in_development_inner.md) |  | 
**build_state** | [**models::ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState**](_api_v1_conditions__condition_id__evidence_freshness_get_200_response_build_state.md) |  | 
**sections** | [**models::ApiV1OrgPipelineStagesGet200ResponseSections**](_api_v1_org_pipeline_stages_get_200_response_sections.md) |  | 
**error** | Option<**String**> |  | [optional]
**message** | Option<**String**> |  | [optional]
**name_withheld_note** | Option<**String**> |  | [optional]
**name_built** | Option<**bool**> |  | [optional]
**trials_by_year** | Option<[**Vec<models::ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialsByYearInner>**](_api_v1_conditions__condition_id__sponsor_landscape_get_200_response_trials_by_year_inner.md)> |  | [optional]
**trial_start_forecast** | Option<[**models::ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast**](_api_v1_org_pipeline_stages_get_200_response_trial_start_forecast.md)> |  | [optional]
**statement** | Option<**String**> |  | [optional]
**recruitment_rollup** | Option<[**models::ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollup**](_api_v1_org_pipeline_stages_get_200_response_recruitment_rollup.md)> |  | [optional]
**recruitment_statement** | Option<**String**> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


