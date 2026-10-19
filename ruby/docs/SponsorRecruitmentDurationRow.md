# PharmaObservatory::SponsorRecruitmentDurationRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **sponsor_org_id** | **String** | Legal-entity level: the organization record the series is for (the lead sponsor). | [optional] |
| **sponsor_name** | **String** | Display name of the organization record. | [optional] |
| **sponsor_company_id** | **String** | Company level: the company the organization belongs to (null when unresolved). | [optional] |
| **sponsor_company_name** | **String** | Display name of the sponsor&#39;s company. | [optional] |
| **sponsor_group_id** | **String** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **sponsor_group_name** | **String** | Display name of the sponsor&#39;s group. | [optional] |
| **n_trials_in_series** | **Integer** | Trials the organization leads. | [optional] |
| **n_evaluated** | **Integer** | Of those, trials with a row in the per-trial predictions at the cutoff. | [optional] |
| **n_not_evaluated** | **Integer** | n_trials_in_series - n_evaluated: registered after the cutoff or a coverage gap; the data cannot tell which. | [optional] |
| **n_out_of_population** | **Integer** | Evaluated trials with prediction_status out_of_population. | [optional] |
| **n_completed** | **Integer** | Evaluated trials with prediction_status completed. | [optional] |
| **n_stopped** | **Integer** | Evaluated trials with prediction_status stopped. | [optional] |
| **n_no_registration_record** | **Integer** | Evaluated trials with prediction_status no_registration_record. | [optional] |
| **n_retrospective_registration** | **Integer** | Evaluated trials with prediction_status retrospective_registration. | [optional] |
| **n_start_not_actual** | **Integer** | Evaluated trials with prediction_status start_not_actual. | [optional] |
| **n_fit_failed** | **Integer** | Evaluated trials with prediction_status fit_failed. | [optional] |
| **n_predicted** | **Integer** | Evaluated trials with prediction_status predicted, superseded or not. | [optional] |
| **n_predicted_superseded** | **Integer** | Predicted trials whose CURRENT record shows the outcome (COMPLETED, TERMINATED, WITHDRAWN or an ACTUAL primary completion); kept out of every predicted figure, as the trial page withholds them. | [optional] |
| **n_predicted_open** | **Integer** | n_predicted - n_predicted_superseded; the denominator of every predicted figure. | [optional] |
| **n_predicted_at_cap** | **Integer** | Open predicted trials whose figure sits at the horizon cap (an at-cap row contributes the cap to the median). | [optional] |
| **n_predicted_overdue** | **Integer** | Open predicted trials running past their predicted end at the cutoff. | [optional] |
| **n_predicted_error_measured** | **Integer** | Open predicted trials in a phase bucket the backtest measured (error_phase_mae_months not null). | [optional] |
| **predicted_months_median** | **Float** | Median of the carried predicted months over open predicted trials (capped at horizon_cap_months). NULL under min_n. | [optional] |
| **predicted_months_p25** | **Float** | 25th percentile, same rows and floor. | [optional] |
| **predicted_months_p75** | **Float** | 75th percentile, same rows and floor. | [optional] |
| **baseline_months_median** | **Float** | Median of the carried baseline months over the same open predicted trials. NULL under min_n. | [optional] |
| **n_realized** | **Integer** | Trials with a registered span (primary completion typed ACTUAL today, plausible start and completion), from the realized-duration input. | [optional] |
| **realized_months_median** | **Float** | Median registered start to primary completion in months over n_realized trials. NULL under min_n. | [optional] |
| **realized_months_p25** | **Float** | 25th percentile, same rows and floor. | [optional] |
| **realized_months_p75** | **Float** | 75th percentile, same rows and floor. | [optional] |
| **n_realized_over_cap** | **Integer** | Realized spans above horizon_cap_months, so a reader sees how many would have hit the cap. | [optional] |
| **min_n** | **Integer** | The floor applied: a median is withheld below this many trials, carried so a reader sees the floor used. | [optional] |
| **rollup_status** | **String** | rolled_up (open predicted &gt;&#x3D; min_n), below_floor (1 &lt;&#x3D; open &lt; min_n, counts only), none_predicted (evaluated, no open prediction), none_evaluated. | [optional] |
| **realized_status** | **String** | rolled_up (n_realized &gt;&#x3D; min_n), below_floor, or none. | [optional] |
| **training_cutoff** | **Date** | The date the model&#39;s training data ends. | [optional] |
| **model_name** | **String** | The recruitment-duration model. | [optional] |
| **model_version** | **String** | The model version. | [optional] |
| **horizon_cap_months** | **Integer** | The cap on predicted months (36 today), from the prediction rows. | [optional] |
| **method** | **String** | The run&#39;s method sentence from the run provenance, never re-authored. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SponsorRecruitmentDurationRow.new(
  sponsor_org_id: null,
  sponsor_name: null,
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null,
  n_trials_in_series: null,
  n_evaluated: null,
  n_not_evaluated: null,
  n_out_of_population: null,
  n_completed: null,
  n_stopped: null,
  n_no_registration_record: null,
  n_retrospective_registration: null,
  n_start_not_actual: null,
  n_fit_failed: null,
  n_predicted: null,
  n_predicted_superseded: null,
  n_predicted_open: null,
  n_predicted_at_cap: null,
  n_predicted_overdue: null,
  n_predicted_error_measured: null,
  predicted_months_median: null,
  predicted_months_p25: null,
  predicted_months_p75: null,
  baseline_months_median: null,
  n_realized: null,
  realized_months_median: null,
  realized_months_p25: null,
  realized_months_p75: null,
  n_realized_over_cap: null,
  min_n: null,
  rollup_status: null,
  realized_status: null,
  training_cutoff: null,
  model_name: null,
  model_version: null,
  horizon_cap_months: null,
  method: null
)
```

