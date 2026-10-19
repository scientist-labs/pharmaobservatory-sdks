# PharmaObservatory::TrialEnrollmentHistoryRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **as_of** | **Date** | The registry snapshot date this row describes. | [optional] |
| **enrollment** | **Integer** | Enrollment registered at the snapshot (a planned target or an achieved count, see enrollment_type). | [optional] |
| **enrollment_type** | **String** | ACTUAL or ESTIMATED at the snapshot. | [optional] |
| **overall_status** | **String** | The trial&#39;s overall status code at the snapshot. | [optional] |
| **start_date** | **Date** | Start date at the snapshot. | [optional] |
| **primary_completion_date** | **Date** | Primary completion date at the snapshot. | [optional] |
| **completion_date** | **Date** | Completion date at the snapshot. | [optional] |
| **enrollment_is_plausible** | **Boolean** | False when the registered enrollment is a placeholder value rather than a count. | [optional] |
| **enrollment_flag_reason** | **String** | Why the enrollment was marked a placeholder (null when it was not). | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialEnrollmentHistoryRow.new(
  trial_id: null,
  as_of: null,
  enrollment: null,
  enrollment_type: null,
  overall_status: null,
  start_date: null,
  primary_completion_date: null,
  completion_date: null,
  enrollment_is_plausible: null,
  enrollment_flag_reason: null
)
```

