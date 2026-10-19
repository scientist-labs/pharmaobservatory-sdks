# PharmaObservatory::TrialEnrollmentRateRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **phase** | **String** | The registry&#39;s phase code. | [optional] |
| **lead_sponsor_name** | **String** | The lead sponsor name as filed. | [optional] |
| **overall_status_summary** | **String** | A coarse status label (Completed and so on). | [optional] |
| **actual_enrollment** | **Integer** | Enrollment typed ACTUAL. | [optional] |
| **enrollment_type** | **String** | ACTUAL or ESTIMATED. | [optional] |
| **enrollment_is_plausible** | **Boolean** | False when the registered enrollment is a placeholder value rather than a count. | [optional] |
| **active_sites** | **Integer** | Number of sites the rate is computed over. | [optional] |
| **facility_country_count** | **Integer** | Number of countries with a listed site. | [optional] |
| **has_us_facility** | **Boolean** | Whether any listed site is in the United States. | [optional] |
| **start_date** | **Date** | Trial start date. | [optional] |
| **primary_completion_date** | **Date** | Primary completion date. | [optional] |
| **completion_date** | **Date** | Completion date. | [optional] |
| **primary_completion_date_type** | **String** | ACTUAL or ESTIMATED. | [optional] |
| **completion_date_type** | **String** | ACTUAL or ESTIMATED. | [optional] |
| **is_rate_computable** | **Boolean** | Whether the inputs support a rate. | [optional] |
| **rate_flag_reason** | **String** | Why a rate could not be computed (null when it could). | [optional] |
| **duration_months** | **Float** | Months from start to primary completion. | [optional] |
| **duration_months_overall** | **Float** | Months from start to completion. | [optional] |
| **pts_per_month** | **Float** | Participants enrolled per month. | [optional] |
| **pts_per_site_month** | **Float** | Participants enrolled per site per month. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialEnrollmentRateRow.new(
  trial_id: null,
  phase: null,
  lead_sponsor_name: null,
  overall_status_summary: null,
  actual_enrollment: null,
  enrollment_type: null,
  enrollment_is_plausible: null,
  active_sites: null,
  facility_country_count: null,
  has_us_facility: null,
  start_date: null,
  primary_completion_date: null,
  completion_date: null,
  primary_completion_date_type: null,
  completion_date_type: null,
  is_rate_computable: null,
  rate_flag_reason: null,
  duration_months: null,
  duration_months_overall: null,
  pts_per_month: null,
  pts_per_site_month: null
)
```

