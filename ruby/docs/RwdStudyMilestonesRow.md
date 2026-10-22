# PharmaObservatory::RwdStudyMilestonesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **study_id** | **String** | The register&#39;s own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional] |
| **milestone_ordinal** | **Integer** | The milestone&#39;s position in the study timeline. | [optional] |
| **milestone_key** | **String** | Which milestone: funding_contract, study_start, data_analysis_start, interim_report or final_report. | [optional] |
| **milestone_label** | **String** | The milestone&#39;s label. | [optional] |
| **planned_on** | **Date** | The planned date. | [optional] |
| **actual_on** | **Date** | The actual date. | [optional] |
| **has_pair** | **Boolean** | True when both dates are present. | [optional] |
| **is_backfilled** | **Boolean** | True when planned and actual are both present and equal to the day (a plan entered after the fact). | [optional] |
| **is_comparable_pair** | **Boolean** | True when both dates are present and differ, so the slip is a measurement; every delivery distribution is computed over these rows. | [optional] |
| **slip_days** | **Integer** | Actual minus planned in days (positive is late); null when either date is missing or the two are equal. | [optional] |
| **is_late_over_30d** | **Boolean** | True when the slip is more than 30 days; null when the pair is not comparable. | [optional] |
| **required_by_regulator** | **String** | The registrant&#39;s answer to whether a regulatory body required the study (Yes, No or Unknown). | [optional] |
| **mandate_stratum** | **String** | mandated when a regulator required the study (a post-authorisation obligation), voluntary when the registrant says it did not, unknown when unanswered. Delivery timeliness is only comparable within a stratum. | [optional] |
| **rmp_requirement** | **String** | The risk management plan (RMP) category the registrant says required the study, as typed. | [optional] |
| **study_status** | **String** | The study&#39;s status as the register words it (Planned, Ongoing, Finalised, and so on). | [optional] |
| **snapshot_label** | **String** | The label of the catalogue snapshot the row was read from (a date). | [optional] |
| **accessed_month** | **String** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdStudyMilestonesRow.new(
  study_id: null,
  milestone_ordinal: null,
  milestone_key: null,
  milestone_label: null,
  planned_on: null,
  actual_on: null,
  has_pair: null,
  is_backfilled: null,
  is_comparable_pair: null,
  slip_days: null,
  is_late_over_30d: null,
  required_by_regulator: null,
  mandate_stratum: null,
  rmp_requirement: null,
  study_status: null,
  snapshot_label: null,
  accessed_month: null,
  restriction_tier: null,
  attribution: null
)
```

