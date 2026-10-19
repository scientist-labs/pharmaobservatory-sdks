# ConditionDeliveryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**condition_id** | Option<**String**> | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional]
**condition_name** | Option<**String**> | The condition's MeSH heading. | [optional]
**n_trials** | Option<**i64**> | Trials in the roll-up, counted under each condition they are tagged with and each condition above it in the MeSH tree. | [optional]
**n_trials_with_observed_plan** | Option<**i64**> | Trials whose first recorded plan is known. | [optional]
**n_trials_with_enrollment_delta** | Option<**i64**> | Trials graded on enrollment. | [optional]
**n_trials_with_completion_delta** | Option<**i64**> | Trials graded on completion date. | [optional]
**n_enrollment_met** | Option<**i64**> | Graded trials that met their planned enrollment. | [optional]
**n_enrollment_missed** | Option<**i64**> | Graded trials that missed their planned enrollment. | [optional]
**enrollment_hit_rate** | Option<**f64**> | Share of graded trials that met planned enrollment (0 to 1). | [optional]
**median_enrollment_attainment_pct** | Option<**f64**> | Median enrollment attained as a percent of the plan. | [optional]
**n_on_time** | Option<**i64**> | Graded trials that completed on or before the planned date. | [optional]
**n_slipped** | Option<**i64**> | Graded trials that completed after the planned date. | [optional]
**on_time_rate** | Option<**f64**> | Share of graded trials that completed on time (0 to 1). | [optional]
**median_completion_slip_days** | Option<**f64**> | Median days between planned and actual primary completion, among graded trials; negative when earlier than planned. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


