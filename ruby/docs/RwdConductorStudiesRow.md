# PharmaObservatory::RwdConductorStudiesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **conductor_id** | **String** | The conductor slug: the register&#39;s normalised institution name key with spaces as hyphens, the id in the web page address (/rwd-conductors/&lt;id&gt;). The entry key of the rwd_conductors object type; injective because no name key carries a hyphen. | [optional] |
| **study_id** | **String** | The register&#39;s own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional] |
| **conductor_role** | **String** | conducting when the register names the conductor as an institution conducting the study, additional when as an additional institution. | [optional] |
| **study_title** | **String** | The study&#39;s title as registered. | [optional] |
| **study_status** | **String** | The study&#39;s status as the register words it (Planned, Ongoing, Finalised, and so on). | [optional] |
| **mandate_stratum** | **String** | mandated when a regulator required the study (a post-authorisation obligation), voluntary when the registrant says it did not, unknown when unanswered. Delivery timeliness is only comparable within a stratum. | [optional] |
| **study_start_actual_on** | **Date** | Actual study start date. | [optional] |
| **final_report_slip_days** | **Integer** | Actual minus planned final report date in days (positive is late); null when either date is missing or the two dates are equal (a backfill). | [optional] |
| **final_report_is_backfilled** | **Boolean** | True when the planned and actual final report dates are both present and equal to the day, which reads as a plan entered after the fact rather than perfect delivery. | [optional] |
| **snapshot_label** | **String** | The label of the catalogue snapshot the row was read from (a date). | [optional] |
| **accessed_month** | **String** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdConductorStudiesRow.new(
  conductor_id: null,
  study_id: null,
  conductor_role: null,
  study_title: null,
  study_status: null,
  mandate_stratum: null,
  study_start_actual_on: null,
  final_report_slip_days: null,
  final_report_is_backfilled: null,
  snapshot_label: null,
  accessed_month: null,
  restriction_tier: null,
  attribution: null
)
```

