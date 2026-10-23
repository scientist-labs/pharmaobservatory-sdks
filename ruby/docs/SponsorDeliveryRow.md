# PharmaObservatory::SponsorDeliveryRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **sponsor_org_id** | **String** | Legal-entity level: the lead sponsor organization record. | [optional] |
| **sponsor_name** | **String** | Display name of the sponsor organization record. | [optional] |
| **sponsor_company_id** | **String** | Company level: the company the sponsor belongs to (null when unresolved). | [optional] |
| **sponsor_company_name** | **String** | Display name of the sponsor&#39;s company. | [optional] |
| **sponsor_group_id** | **String** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **sponsor_group_name** | **String** | Display name of the sponsor&#39;s group. | [optional] |
| **n_trials_total** | **Integer** | Lead-sponsored trials of the organization. | [optional] |
| **n_trials_with_observed_plan** | **Integer** | Trials whose planned dates come from the earliest estimate in the registry&#39;s own history. | [optional] |
| **n_trials_with_enrollment_delta** | **Integer** | Trials with a measured difference between planned and actual enrollment; the base of the enrollment rate. | [optional] |
| **n_trials_with_completion_delta** | **Integer** | Trials with a measured difference between planned and actual primary completion; the base of the on-time rate. | [optional] |
| **n_enrollment_met** | **Integer** | Trials with an enrollment delta that reached their planned enrollment. | [optional] |
| **n_enrollment_missed** | **Integer** | Trials with an enrollment delta that did not reach their planned enrollment. | [optional] |
| **enrollment_hit_rate** | **Float** | n_enrollment_met divided by n_trials_with_enrollment_delta; null when that base is zero. | [optional] |
| **median_enrollment_attainment_pct** | **Float** | Median of actual enrollment as a percentage of planned, over trials with an enrollment delta. | [optional] |
| **n_ontime** | **Integer** | Trials with a completion delta that were not delayed past their planned primary completion. | [optional] |
| **n_slipped** | **Integer** | Trials with a completion delta that completed or are projected after their planned primary completion. | [optional] |
| **on_time_rate** | **Float** | n_ontime divided by n_trials_with_completion_delta; null when that base is zero. | [optional] |
| **median_completion_slip_days** | **Float** | Median days between planned and actual primary completion, over trials with a completion delta. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SponsorDeliveryRow.new(
  sponsor_org_id: null,
  sponsor_name: null,
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null,
  n_trials_total: null,
  n_trials_with_observed_plan: null,
  n_trials_with_enrollment_delta: null,
  n_trials_with_completion_delta: null,
  n_enrollment_met: null,
  n_enrollment_missed: null,
  enrollment_hit_rate: null,
  median_enrollment_attainment_pct: null,
  n_ontime: null,
  n_slipped: null,
  on_time_rate: null,
  median_completion_slip_days: null
)
```

