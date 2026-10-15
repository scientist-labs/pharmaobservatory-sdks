# PharmaObservatory::ApiV1ConditionsConditionIdSponsorLandscapeGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** |  |  |
| **name** | **String** |  |  |
| **sponsors** | [**Array&lt;ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseSponsorsInner&gt;**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseSponsorsInner.md) |  |  |
| **phase_breakdown** | [**Array&lt;ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponsePhaseBreakdownInner&gt;**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponsePhaseBreakdownInner.md) |  |  |
| **active_year_span** | [**ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseActiveYearSpan**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseActiveYearSpan.md) |  |  |
| **build_state** | [**ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState**](ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState.md) |  | [optional] |
| **sections** | [**ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseSections**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseSections.md) |  | [optional] |
| **trials_by_year** | [**Array&lt;ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialsByYearInner&gt;**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialsByYearInner.md) |  | [optional] |
| **trial_start_forecast** | [**ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialStartForecast**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseTrialStartForecast.md) |  | [optional] |
| **statement** | **String** |  | [optional] |
| **recruitment_rollup** | [**ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollup.md) |  | [optional] |
| **recruitment_statement** | **String** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1ConditionsConditionIdSponsorLandscapeGet200Response.new(
  condition_id: null,
  name: null,
  sponsors: null,
  phase_breakdown: null,
  active_year_span: null,
  build_state: null,
  sections: null,
  trials_by_year: null,
  trial_start_forecast: null,
  statement: null,
  recruitment_rollup: null,
  recruitment_statement: null
)
```

