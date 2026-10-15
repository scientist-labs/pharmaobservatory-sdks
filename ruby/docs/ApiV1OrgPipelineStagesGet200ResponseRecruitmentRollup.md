# PharmaObservatory::ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollup

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **kind** | **String** |  |  |
| **state** | **String** |  |  |
| **realized_state** | **String** |  |  |
| **cutoff** | **String** |  |  |
| **min_n** | **Integer** |  |  |
| **cap_months** | **Integer** |  |  |
| **model_version** | **String** |  |  |
| **method** | **String** |  |  |
| **n_trials** | **Integer** |  |  |
| **n_evaluated** | **Integer** |  |  |
| **n_not_evaluated** | **Integer** |  |  |
| **statuses** | [**ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupStatuses**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupStatuses.md) |  |  |
| **n_superseded** | **Integer** |  |  |
| **n_open** | **Integer** |  |  |
| **n_at_cap** | **Integer** |  |  |
| **n_overdue** | **Integer** |  |  |
| **n_error_measured** | **Integer** |  |  |
| **predicted** | [**ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPredicted**](ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPredicted.md) |  |  |
| **n_realized** | **Integer** |  |  |
| **n_realized_over_cap** | **Integer** |  |  |
| **realized** | [**ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupRealized**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupRealized.md) |  |  |
| **phases** | [**Array&lt;ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner&gt;**](ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollupPhasesInner.md) |  |  |
| **scope** | **String** |  |  |
| **title** | **Object** |  |  |
| **body** | **String** |  |  |
| **figures** | [**Array&lt;ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupFiguresInner&gt;**](ApiV1ConditionsConditionIdSponsorLandscapeGet200ResponseRecruitmentRollupFiguresInner.md) |  |  |
| **status_counts** | [**Array&lt;ApiV1SponsorLandscapeRecruitmentRollupStatusCountsInner&gt;**](ApiV1SponsorLandscapeRecruitmentRollupStatusCountsInner.md) |  |  |
| **source** | **String** |  |  |
| **predicted_at_horizon_cap** | [**ApiV1SponsorLandscapeRecruitmentRollupPredictedAtHorizonCap**](ApiV1SponsorLandscapeRecruitmentRollupPredictedAtHorizonCap.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1OrgPipelineStagesGet200ResponseRecruitmentRollup.new(
  kind: null,
  state: null,
  realized_state: null,
  cutoff: null,
  min_n: null,
  cap_months: null,
  model_version: null,
  method: null,
  n_trials: null,
  n_evaluated: null,
  n_not_evaluated: null,
  statuses: null,
  n_superseded: null,
  n_open: null,
  n_at_cap: null,
  n_overdue: null,
  n_error_measured: null,
  predicted: null,
  n_realized: null,
  n_realized_over_cap: null,
  realized: null,
  phases: null,
  scope: null,
  title: null,
  body: null,
  figures: null,
  status_counts: null,
  source: null,
  predicted_at_horizon_cap: null
)
```

