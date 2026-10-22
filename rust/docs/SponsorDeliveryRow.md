# SponsorDeliveryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sponsor_org_id** | Option<**String**> | Legal-entity level: the lead sponsor organization record. | [optional]
**sponsor_name** | Option<**String**> | Display name of the sponsor organization record. | [optional]
**sponsor_company_id** | Option<**String**> | Company level: the company the sponsor belongs to (null when unresolved). | [optional]
**sponsor_company_name** | Option<**String**> | Display name of the sponsor's company. | [optional]
**sponsor_group_id** | Option<**String**> | Group level: the sponsor's current ultimate parent company; a root company's group is itself. | [optional]
**sponsor_group_name** | Option<**String**> | Display name of the sponsor's group. | [optional]
**n_trials_total** | Option<**i64**> | Lead-sponsored trials of the organization. | [optional]
**n_trials_with_observed_plan** | Option<**i64**> | Trials whose planned dates come from the earliest estimate in the registry's own history. | [optional]
**n_trials_with_enrollment_delta** | Option<**i64**> | Trials with a measured difference between planned and actual enrollment; the base of the enrollment rate. | [optional]
**n_trials_with_completion_delta** | Option<**i64**> | Trials with a measured difference between planned and actual primary completion; the base of the on-time rate. | [optional]
**n_enrollment_met** | Option<**i64**> | Trials with an enrollment delta that reached their planned enrollment. | [optional]
**n_enrollment_missed** | Option<**i64**> | Trials with an enrollment delta that did not reach their planned enrollment. | [optional]
**enrollment_hit_rate** | Option<**f64**> | n_enrollment_met divided by n_trials_with_enrollment_delta; null when that base is zero. | [optional]
**median_enrollment_attainment_pct** | Option<**f64**> | Median of actual enrollment as a percentage of planned, over trials with an enrollment delta. | [optional]
**n_ontime** | Option<**i64**> | Trials with a completion delta that were not delayed past their planned primary completion. | [optional]
**n_slipped** | Option<**i64**> | Trials with a completion delta that completed or are projected after their planned primary completion. | [optional]
**on_time_rate** | Option<**f64**> | n_ontime divided by n_trials_with_completion_delta; null when that base is zero. | [optional]
**median_completion_slip_days** | Option<**f64**> | Median days between planned and actual primary completion, over trials with a completion delta. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


