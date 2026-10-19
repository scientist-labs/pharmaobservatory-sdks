# TrialEnrollmentHistoryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**as_of** | Option<[**String**](string.md)> | The registry snapshot date this row describes. | [optional]
**enrollment** | Option<**i32**> | Enrollment registered at the snapshot (a planned target or an achieved count, see enrollment_type). | [optional]
**enrollment_type** | Option<**String**> | ACTUAL or ESTIMATED at the snapshot. | [optional]
**overall_status** | Option<**String**> | The trial's overall status code at the snapshot. | [optional]
**start_date** | Option<[**String**](string.md)> | Start date at the snapshot. | [optional]
**primary_completion_date** | Option<[**String**](string.md)> | Primary completion date at the snapshot. | [optional]
**completion_date** | Option<[**String**](string.md)> | Completion date at the snapshot. | [optional]
**enrollment_is_plausible** | Option<**bool**> | False when the registered enrollment is a placeholder value rather than a count. | [optional]
**enrollment_flag_reason** | Option<**String**> | Why the enrollment was marked a placeholder (null when it was not). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


