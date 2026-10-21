# PharmaObservatory::ConditionTrialStartForecastRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **condition_name** | **String** | The condition&#39;s MeSH heading. | [optional] |
| **target_year** | **Integer** | The calendar year forecast. | [optional] |
| **horizon_offset** | **Integer** | Years ahead of the build year (0 is the current year). | [optional] |
| **months_observed** | **Integer** | Months of the target year already observed. | [optional] |
| **months_forecast** | **Integer** | Months of the target year forecast. | [optional] |
| **input_observed_through_cutoff** | **Integer** | Trials started in the target year up to the training cutoff. | [optional] |
| **input_planned_remaining** | **Integer** | Trials registered as planned to start in the rest of the target year. | [optional] |
| **forecast_point** | **Float** | The forecast trial starts for the whole target year. A forecast of the eventual count, never an actual count, and it carries no interval. | [optional] |
| **forecast_status** | **String** | The series&#39; forecast status. | [optional] |
| **input_trials_window_60m** | **Integer** | Trials of the condition (as in condition_trials) with a start date in the 60 calendar months ending with the cutoff month. | [optional] |
| **input_starts_per_year_60m** | **Float** | Average trial starts per year over the last 60 months. | [optional] |
| **size_bucket** | **String** | The series size bucket the backtest error is measured in. | [optional] |
| **backtest_horizon** | **String** | The horizon the backtest error was measured at. | [optional] |
| **backtest_smape_pct** | **Float** | Backtest symmetric mean absolute percentage error. | [optional] |
| **backtest_rel_bias** | **Float** | Backtest relative bias. | [optional] |
| **backtest_relmae_vs_snaive** | **Float** | Backtest mean absolute error relative to the seasonal naive baseline. | [optional] |
| **backtest_n_series** | **Integer** | Series the backtest covered. | [optional] |
| **backtest_series_kind** | **String** | The kind of series the backtest error was measured on. | [optional] |
| **backtest_ref** | **String** | The backtest run reference. | [optional] |
| **planned_start_coefficient** | **Float** | The multiple of input_planned_remaining the planned-start method adds to the months after the training cutoff; null on a row built with the flat mean alone. | [optional] |
| **planned_start_stricter_check_agrees** | **Boolean** | On a planned-start row, whether the stricter re-check (trained only on target years at least two before the cutoff year) adopted the same cell; null on a flat-mean row. | [optional] |
| **planned_start_stricter_check_rel_bias** | **Float** | On a planned-start row, the relative bias (forecast minus actual over actual) of the stricter re-check; null on a flat-mean row. | [optional] |
| **model_name** | **String** | The forecasting model. | [optional] |
| **model_version** | **String** | The model version. | [optional] |
| **training_cutoff** | **Date** | The date the model&#39;s training data ends. | [optional] |
| **method** | **String** | How the forecast was produced. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ConditionTrialStartForecastRow.new(
  condition_id: null,
  condition_name: null,
  target_year: null,
  horizon_offset: null,
  months_observed: null,
  months_forecast: null,
  input_observed_through_cutoff: null,
  input_planned_remaining: null,
  forecast_point: null,
  forecast_status: null,
  input_trials_window_60m: null,
  input_starts_per_year_60m: null,
  size_bucket: null,
  backtest_horizon: null,
  backtest_smape_pct: null,
  backtest_rel_bias: null,
  backtest_relmae_vs_snaive: null,
  backtest_n_series: null,
  backtest_series_kind: null,
  backtest_ref: null,
  planned_start_coefficient: null,
  planned_start_stricter_check_agrees: null,
  planned_start_stricter_check_rel_bias: null,
  model_name: null,
  model_version: null,
  training_cutoff: null,
  method: null
)
```

