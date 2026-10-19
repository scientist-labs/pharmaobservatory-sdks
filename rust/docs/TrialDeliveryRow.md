# TrialDeliveryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**lead_sponsor_name** | Option<**String**> | The lead sponsor name as filed. | [optional]
**overall_status_summary** | Option<**String**> | A coarse status label (Recruiting, Completed and so on). | [optional]
**study_type** | Option<**String**> | INTERVENTIONAL, OBSERVATIONAL or EXPANDED_ACCESS. | [optional]
**phase** | Option<**String**> | The registry's phase code. | [optional]
**planned_enrollment** | Option<**i32**> | Planned enrollment from the earliest snapshot typed ESTIMATED. | [optional]
**actual_enrollment** | Option<**i32**> | Enrollment from the latest observation typed ACTUAL. | [optional]
**enrollment_type** | Option<**String**> | Enrollment type on the current record (ACTUAL or ESTIMATED). | [optional]
**enrollment_delta** | Option<**i32**> | Actual minus planned enrollment (null without both). | [optional]
**enrollment_attainment_pct** | Option<**f64**> | Actual as a percentage of planned enrollment (null without both). | [optional]
**enrollment_met** | Option<**bool**> | Whether actual enrollment reached the plan (null without both). | [optional]
**has_enrollment_delta** | Option<**bool**> | Whether a planned and an actual enrollment are both recorded. | [optional]
**planned_basis** | Option<**String**> | Where the planned enrollment came from: history_earliest_estimated, type_estimated_current or none. | [optional]
**actual_basis** | Option<**String**> | How the actual enrollment was taken (from the snapshot history, from the current record, or none). | [optional]
**planned_primary_completion_date** | Option<[**String**](string.md)> | Planned primary completion date from the earliest snapshot typed ESTIMATED. | [optional]
**actual_primary_completion_date** | Option<[**String**](string.md)> | Primary completion date once typed ACTUAL. | [optional]
**planned_completion_date** | Option<[**String**](string.md)> | Planned completion date from the earliest snapshot typed ESTIMATED. | [optional]
**actual_completion_date** | Option<[**String**](string.md)> | Completion date once typed ACTUAL. | [optional]
**primary_completion_date_type** | Option<**String**> | Type of the current primary completion date (ACTUAL or ESTIMATED). | [optional]
**completion_date_type** | Option<**String**> | Type of the current completion date (ACTUAL or ESTIMATED). | [optional]
**primary_completion_slip_days** | Option<**i32**> | Actual minus planned primary completion date; positive is later than planned (null without both). | [optional]
**completion_slip_days** | Option<**i32**> | Actual minus planned completion date; positive is later than planned (null without both). | [optional]
**completion_slipped** | Option<**bool**> | Whether the primary completion came later than planned (null without both). | [optional]
**has_completion_delta** | Option<**bool**> | Whether a planned and an actual primary completion date are both recorded. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


