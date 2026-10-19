# PharmaObservatory::SponsorTrialStartForecastRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **sponsor_org_id** | **String** | Legal-entity level: the organization record the forecast series is for. | [optional] |
| **sponsor_name** | **String** | Display name of the organization record. | [optional] |
| **sponsor_company_id** | **String** | Company level: the company the organization belongs to (null when unresolved). | [optional] |
| **sponsor_company_name** | **String** | Display name of the sponsor&#39;s company. | [optional] |
| **sponsor_group_id** | **String** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **sponsor_group_name** | **String** | Display name of the sponsor&#39;s group. | [optional] |
| **target_year** | **Integer** | The calendar year forecast. | [optional] |
| **horizon_offset** | **Integer** | Years ahead of the build year (0 is the current year). | [optional] |
| **months_observed** | **Integer** | Months of the target year already observed. | [optional] |
| **months_forecast** | **Integer** | Months of the target year forecast. | [optional] |
| **input_observed_through_cutoff** | **Integer** | Trials started in the target year up to the training cutoff. | [optional] |
| **input_planned_remaining** | **Integer** | Trials registered as planned to start in the rest of the target year. | [optional] |
| **forecast_point** | **Float** | The forecast trial starts for the whole target year. A forecast of the eventual count, never an actual count, and it carries no interval. | [optional] |
| **forecast_status** | **String** | The series&#39; forecast status. | [optional] |
| **input_trials_window_60m** | **Integer** | Trials the organization started in the last 60 months. | [optional] |
| **input_starts_per_year_60m** | **Float** | Average trial starts per year over the last 60 months. | [optional] |
| **size_bucket** | **String** | The series size bucket the backtest error is measured in. | [optional] |
| **backtest_horizon** | **String** | The horizon the backtest error was measured at. | [optional] |
| **backtest_smape_pct** | **Float** | Backtest symmetric mean absolute percentage error. | [optional] |
| **backtest_rel_bias** | **Float** | Backtest relative bias. | [optional] |
| **backtest_relmae_vs_snaive** | **Float** | Backtest mean absolute error relative to the seasonal naive baseline. | [optional] |
| **backtest_n_series** | **Integer** | Series the backtest covered. | [optional] |
| **backtest_ref** | **String** | The backtest run reference. | [optional] |
| **model_name** | **String** | The forecasting model. | [optional] |
| **model_version** | **String** | The model version. | [optional] |
| **training_cutoff** | **Date** | The date the model&#39;s training data ends. | [optional] |
| **method** | **String** | How the forecast was produced. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SponsorTrialStartForecastRow.new(
  sponsor_org_id: null,
  sponsor_name: null,
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null,
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
  backtest_ref: null,
  model_name: null,
  model_version: null,
  training_cutoff: null,
  method: null
)
```

