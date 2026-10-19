# PharmaObservatory::SponsorRecruitmentByPhaseRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **sponsor_org_id** | **String** | Legal-entity level: the organization record the series is for (the lead sponsor). | [optional] |
| **sponsor_name** | **String** | Display name of the organization record. | [optional] |
| **sponsor_company_id** | **String** | Company level: the company the organization belongs to (null when unresolved). | [optional] |
| **sponsor_company_name** | **String** | Display name of the sponsor&#39;s company. | [optional] |
| **sponsor_group_id** | **String** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **sponsor_group_name** | **String** | Display name of the sponsor&#39;s group. | [optional] |
| **phase_bucket** | **String** | The phase bucket of the trial (one vocabulary for predicted and realized rows). | [optional] |
| **n_predicted_open** | **Integer** | Per phase bucket, the same count as the series column of that name. | [optional] |
| **n_predicted_at_cap** | **Integer** | Per phase bucket, the same count as the series column of that name. | [optional] |
| **n_predicted_overdue** | **Integer** | Per phase bucket, the same count as the series column of that name. | [optional] |
| **predicted_months_median** | **Float** | Median over the bucket&#39;s rows; NULL under min_n. | [optional] |
| **baseline_months_median** | **Float** | Median over the bucket&#39;s rows; NULL under min_n. | [optional] |
| **n_realized** | **Integer** | Per phase bucket, the same count as the series column of that name. | [optional] |
| **realized_months_median** | **Float** | Median over the bucket&#39;s rows; NULL under min_n. | [optional] |
| **error_phase_mae_months** | **Float** | Backtest mean absolute error in months for this phase bucket; NULL on an unmeasured bucket or one with no open predicted trial. | [optional] |
| **error_phase_relmae_vs_baseline** | **Float** | The model&#39;s mean absolute error (MAE) over the baseline&#39;s MAE in this phase bucket. | [optional] |
| **error_phase_baseline_mae_months** | **Float** | The baseline&#39;s mean absolute error (MAE) in months in this phase bucket. | [optional] |
| **error_phase_n_trials** | **Integer** | Backtest trials behind the phase figure. | [optional] |
| **error_phase_n_years** | **Integer** | Backtest years behind the phase figure. | [optional] |
| **min_n** | **Integer** | The floor applied: a median is withheld below this many trials, carried so a reader sees the floor used. | [optional] |
| **training_cutoff** | **Date** | The date the model&#39;s training data ends. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SponsorRecruitmentByPhaseRow.new(
  sponsor_org_id: null,
  sponsor_name: null,
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null,
  phase_bucket: null,
  n_predicted_open: null,
  n_predicted_at_cap: null,
  n_predicted_overdue: null,
  predicted_months_median: null,
  baseline_months_median: null,
  n_realized: null,
  realized_months_median: null,
  error_phase_mae_months: null,
  error_phase_relmae_vs_baseline: null,
  error_phase_baseline_mae_months: null,
  error_phase_n_trials: null,
  error_phase_n_years: null,
  min_n: null,
  training_cutoff: null
)
```

