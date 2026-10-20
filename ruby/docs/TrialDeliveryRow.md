# PharmaObservatory::TrialDeliveryRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **lead_sponsor_name** | **String** | The lead sponsor name as filed. | [optional] |
| **overall_status_summary** | **String** | A coarse status label (Recruiting, Completed and so on). | [optional] |
| **study_type** | **String** | INTERVENTIONAL, OBSERVATIONAL or EXPANDED_ACCESS. | [optional] |
| **phase** | **String** | The registry&#39;s phase code (PHASE1, PHASE2, PHASE2/PHASE3, NA and so on). Null when the registry gives no phase (observational and expanded-access trials); phase_label carries the display label for those. | [optional] |
| **phase_label** | **String** | The trial phase as a display label (Phase 2, Phase 3, Observational, Not applicable, Expanded access, Unknown). The same trial&#39;s phase column holds the registry code. | [optional] |
| **planned_enrollment** | **Integer** | Planned enrollment from the earliest snapshot typed ESTIMATED. | [optional] |
| **actual_enrollment** | **Integer** | Enrollment from the latest observation typed ACTUAL. | [optional] |
| **enrollment_type** | **String** | Enrollment type on the current record (ACTUAL or ESTIMATED). | [optional] |
| **enrollment_delta** | **Integer** | Actual minus planned enrollment (null without both). | [optional] |
| **enrollment_attainment_pct** | **Float** | Actual as a percentage of planned enrollment (null without both). | [optional] |
| **enrollment_met** | **Boolean** | Whether actual enrollment reached the plan (null without both). | [optional] |
| **enrollment_delta_known** | **Boolean** | Whether a planned and an actual enrollment are both recorded. | [optional] |
| **planned_basis** | **String** | Where the planned enrollment came from: history_earliest_estimated, type_estimated_current or none. | [optional] |
| **actual_basis** | **String** | How the actual enrollment was taken (from the snapshot history, from the current record, or none). | [optional] |
| **planned_primary_completion_date** | **Date** | Planned primary completion date from the earliest snapshot typed ESTIMATED. | [optional] |
| **actual_primary_completion_date** | **Date** | Primary completion date once typed ACTUAL. | [optional] |
| **planned_completion_date** | **Date** | Planned completion date from the earliest snapshot typed ESTIMATED. | [optional] |
| **actual_completion_date** | **Date** | Completion date once typed ACTUAL. | [optional] |
| **primary_completion_date_type** | **String** | ACTUAL (a date that has occurred) or ESTIMATED (a planned date). The registry&#39;s older ANTICIPATED is published as ESTIMATED. | [optional] |
| **completion_date_type** | **String** | ACTUAL (a date that has occurred) or ESTIMATED (a planned date). The registry&#39;s older ANTICIPATED is published as ESTIMATED. | [optional] |
| **primary_completion_slip_days** | **Integer** | Actual minus planned primary completion date; positive is later than planned (null without both). | [optional] |
| **completion_slip_days** | **Integer** | Actual minus planned completion date; positive is later than planned (null without both). | [optional] |
| **completion_slipped** | **Boolean** | Whether the primary completion came later than planned (null without both). | [optional] |
| **has_completion_delta** | **Boolean** | Whether a planned and an actual primary completion date are both recorded. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialDeliveryRow.new(
  trial_id: null,
  lead_sponsor_name: null,
  overall_status_summary: null,
  study_type: null,
  phase: null,
  phase_label: null,
  planned_enrollment: null,
  actual_enrollment: null,
  enrollment_type: null,
  enrollment_delta: null,
  enrollment_attainment_pct: null,
  enrollment_met: null,
  enrollment_delta_known: null,
  planned_basis: null,
  actual_basis: null,
  planned_primary_completion_date: null,
  actual_primary_completion_date: null,
  planned_completion_date: null,
  actual_completion_date: null,
  primary_completion_date_type: null,
  completion_date_type: null,
  primary_completion_slip_days: null,
  completion_slip_days: null,
  completion_slipped: null,
  has_completion_delta: null
)
```

