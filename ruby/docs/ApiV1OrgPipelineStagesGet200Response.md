# PharmaObservatory::ApiV1OrgPipelineStagesGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** |  | [optional] |
| **name** | **String** |  | [optional] |
| **approved_marketed** | [**Array&lt;ApiV1OrgPipelineStagesGet200ResponseApprovedMarketedInner&gt;**](ApiV1OrgPipelineStagesGet200ResponseApprovedMarketedInner.md) |  |  |
| **in_development** | [**Array&lt;ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner&gt;**](ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner.md) |  |  |
| **abandoned_dormant** | [**Array&lt;ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner&gt;**](ApiV1OrgPipelineStagesGet200ResponseInDevelopmentInner.md) |  |  |
| **build_state** | [**ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState**](ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState.md) |  |  |
| **sections** | [**ApiV1OrgPipelineStagesGet200ResponseSections**](ApiV1OrgPipelineStagesGet200ResponseSections.md) |  |  |
| **error** | **String** |  | [optional] |
| **message** | **String** |  | [optional] |
| **name_withheld_note** | **String** |  | [optional] |
| **name_built** | **Boolean** |  | [optional] |
| **trials_by_year** | [**Array&lt;ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialsByYearInner&gt;**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialsByYearInner.md) |  | [optional] |
| **trial_start_forecast** | [**ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast**](ApiV1OrgPipelineStagesGet200ResponseTrialStartForecast.md) |  | [optional] |
| **statement** | **String** |  | [optional] |
| **recruitment_rollup** | [**ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollup**](ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollup.md) |  | [optional] |
| **recruitment_statement** | **String** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1OrgPipelineStagesGet200Response.new(
  company_id: null,
  name: null,
  approved_marketed: null,
  in_development: null,
  abandoned_dormant: null,
  build_state: null,
  sections: null,
  error: null,
  message: null,
  name_withheld_note: null,
  name_built: null,
  trials_by_year: null,
  trial_start_forecast: null,
  statement: null,
  recruitment_rollup: null,
  recruitment_statement: null
)
```

