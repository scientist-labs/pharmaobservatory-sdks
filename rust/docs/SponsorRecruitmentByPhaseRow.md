# SponsorRecruitmentByPhaseRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sponsor_org_id** | Option<**String**> | Legal-entity level: the organization record the series is for (the lead sponsor). | [optional]
**sponsor_name** | Option<**String**> | Display name of the organization record. | [optional]
**sponsor_company_id** | Option<**String**> | Company level: the company the organization belongs to (null when unresolved). | [optional]
**sponsor_company_name** | Option<**String**> | Display name of the sponsor's company. | [optional]
**sponsor_group_id** | Option<**String**> | Group level: the sponsor's current ultimate parent company; a root company's group is itself. | [optional]
**sponsor_group_name** | Option<**String**> | Display name of the sponsor's group. | [optional]
**phase_bucket** | Option<**String**> | The phase bucket of the trial (one vocabulary for predicted and realized rows). | [optional]
**n_predicted_open** | Option<**i32**> | Per phase bucket, the same count as the series column of that name. | [optional]
**n_predicted_at_cap** | Option<**i32**> | Per phase bucket, the same count as the series column of that name. | [optional]
**n_predicted_overdue** | Option<**i32**> | Per phase bucket, the same count as the series column of that name. | [optional]
**predicted_months_median** | Option<**f64**> | Median over the bucket's rows; NULL under min_n. | [optional]
**baseline_months_median** | Option<**f64**> | Median over the bucket's rows; NULL under min_n. | [optional]
**n_realized** | Option<**i32**> | Per phase bucket, the same count as the series column of that name. | [optional]
**realized_months_median** | Option<**f64**> | Median over the bucket's rows; NULL under min_n. | [optional]
**error_phase_mae_months** | Option<**f64**> | Backtest mean absolute error in months for this phase bucket; NULL on an unmeasured bucket or one with no open predicted trial. | [optional]
**error_phase_relmae_vs_baseline** | Option<**f64**> | The model's mean absolute error (MAE) over the baseline's MAE in this phase bucket. | [optional]
**error_phase_baseline_mae_months** | Option<**f64**> | The baseline's mean absolute error (MAE) in months in this phase bucket. | [optional]
**error_phase_n_trials** | Option<**i32**> | Backtest trials behind the phase figure. | [optional]
**error_phase_n_years** | Option<**i32**> | Backtest years behind the phase figure. | [optional]
**min_n** | Option<**i32**> | The floor applied: a median is withheld below this many trials, carried so a reader sees the floor used. | [optional]
**training_cutoff** | Option<[**String**](string.md)> | The date the model's training data ends. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


