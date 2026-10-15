# ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**kind** | **String** |  | 
**state** | **String** |  | 
**realized_state** | **String** |  | 
**cutoff** | **String** |  | 
**min_n** | **i32** |  | 
**cap_months** | **i32** |  | 
**model_version** | **String** |  | 
**method** | **String** |  | 
**n_trials** | **i32** |  | 
**n_evaluated** | **i32** |  | 
**n_not_evaluated** | **i32** |  | 
**statuses** | [**models::ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupStatuses**](_api_v1_conditions__condition_id__sponsor_landscape_get_200_response_recruitment_rollup_statuses.md) |  | 
**n_superseded** | **i32** |  | 
**n_open** | **i32** |  | 
**n_at_cap** | **i32** |  | 
**n_overdue** | **i32** |  | 
**n_error_measured** | **i32** |  | 
**predicted** | [**models::ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPredicted**](_api_v1_org_pipeline_stages_get_200_response_recruitment_rollup_predicted.md) |  | 
**n_realized** | **i32** |  | 
**n_realized_over_cap** | **i32** |  | 
**realized** | [**models::ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupRealized**](_api_v1_conditions__condition_id__sponsor_landscape_get_200_response_recruitment_rollup_realized.md) |  | 
**phases** | [**Vec<models::ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner>**](_api_v1_org_pipeline_stages_get_200_response_recruitment_rollup_phases_inner.md) |  | 
**scope** | **String** |  | 
**title** | Option<[**serde_json::Value**](.md)> |  | 
**body** | **String** |  | 
**figures** | [**Vec<models::ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupFiguresInner>**](_api_v1_conditions__condition_id__sponsor_landscape_get_200_response_recruitment_rollup_figures_inner.md) |  | 
**status_counts** | [**Vec<models::ApiV1SponsorLandscapeRecruitmentRollupStatusCountsInner>**](ApiV1SponsorLandscapeRecruitmentRollupStatusCountsInner.md) |  | 
**source** | **String** |  | 
**predicted_at_horizon_cap** | Option<[**models::ApiV1SponsorLandscapeRecruitmentRollupPredictedAtHorizonCap**](ApiV1SponsorLandscapeRecruitmentRollupPredictedAtHorizonCap.md)> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


