# PharmaObservatory::ConditionDeliveryRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **condition_name** | **String** | The condition&#39;s MeSH heading. | [optional] |
| **n_trials** | **Integer** | Trials in the roll-up, counted under each condition they are tagged with and each condition above it in the MeSH tree. | [optional] |
| **n_trials_with_observed_plan** | **Integer** | Trials whose first recorded plan is known. | [optional] |
| **n_trials_with_enrollment_delta** | **Integer** | Trials graded on enrollment. | [optional] |
| **n_trials_with_completion_delta** | **Integer** | Trials graded on completion date. | [optional] |
| **n_enrollment_met** | **Integer** | Graded trials that met their planned enrollment. | [optional] |
| **n_enrollment_missed** | **Integer** | Graded trials that missed their planned enrollment. | [optional] |
| **enrollment_hit_rate** | **Float** | Share of graded trials that met planned enrollment (0 to 1). | [optional] |
| **median_enrollment_attainment_pct** | **Float** | Median enrollment attained as a percent of the plan. | [optional] |
| **n_on_time** | **Integer** | Graded trials that completed on or before the planned date. | [optional] |
| **n_slipped** | **Integer** | Graded trials that completed after the planned date. | [optional] |
| **on_time_rate** | **Float** | Share of graded trials that completed on time (0 to 1). | [optional] |
| **median_completion_slip_days** | **Float** | Median days between planned and actual primary completion, among graded trials; negative when earlier than planned. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ConditionDeliveryRow.new(
  condition_id: null,
  condition_name: null,
  n_trials: null,
  n_trials_with_observed_plan: null,
  n_trials_with_enrollment_delta: null,
  n_trials_with_completion_delta: null,
  n_enrollment_met: null,
  n_enrollment_missed: null,
  enrollment_hit_rate: null,
  median_enrollment_attainment_pct: null,
  n_on_time: null,
  n_slipped: null,
  on_time_rate: null,
  median_completion_slip_days: null
)
```

