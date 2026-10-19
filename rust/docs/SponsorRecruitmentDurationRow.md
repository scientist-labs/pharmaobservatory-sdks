# SponsorRecruitmentDurationRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sponsor_org_id** | Option<**String**> | Legal-entity level: the organization record the series is for (the lead sponsor). | [optional]
**sponsor_name** | Option<**String**> | Display name of the organization record. | [optional]
**sponsor_company_id** | Option<**String**> | Company level: the company the organization belongs to (null when unresolved). | [optional]
**sponsor_company_name** | Option<**String**> | Display name of the sponsor's company. | [optional]
**sponsor_group_id** | Option<**String**> | Group level: the sponsor's current ultimate parent company; a root company's group is itself. | [optional]
**sponsor_group_name** | Option<**String**> | Display name of the sponsor's group. | [optional]
**n_trials_in_series** | Option<**i32**> | Trials the organization leads. | [optional]
**n_evaluated** | Option<**i32**> | Of those, trials with a row in the per-trial predictions at the cutoff. | [optional]
**n_not_evaluated** | Option<**i32**> | n_trials_in_series - n_evaluated: registered after the cutoff or a coverage gap; the data cannot tell which. | [optional]
**n_out_of_population** | Option<**i32**> | Evaluated trials with prediction_status out_of_population. | [optional]
**n_completed** | Option<**i32**> | Evaluated trials with prediction_status completed. | [optional]
**n_stopped** | Option<**i32**> | Evaluated trials with prediction_status stopped. | [optional]
**n_no_registration_record** | Option<**i32**> | Evaluated trials with prediction_status no_registration_record. | [optional]
**n_retrospective_registration** | Option<**i32**> | Evaluated trials with prediction_status retrospective_registration. | [optional]
**n_start_not_actual** | Option<**i32**> | Evaluated trials with prediction_status start_not_actual. | [optional]
**n_fit_failed** | Option<**i32**> | Evaluated trials with prediction_status fit_failed. | [optional]
**n_predicted** | Option<**i32**> | Evaluated trials with prediction_status predicted, superseded or not. | [optional]
**n_predicted_superseded** | Option<**i32**> | Predicted trials whose CURRENT record shows the outcome (COMPLETED, TERMINATED, WITHDRAWN or an ACTUAL primary completion); kept out of every predicted figure. | [optional]
**n_predicted_open** | Option<**i32**> | n_predicted - n_predicted_superseded; the denominator of every predicted figure. | [optional]
**n_predicted_at_cap** | Option<**i32**> | Open predicted trials whose figure sits at the horizon cap (an at-cap row contributes the cap to the median). | [optional]
**n_predicted_overdue** | Option<**i32**> | Open predicted trials running past their predicted end at the cutoff. | [optional]
**n_predicted_error_measured** | Option<**i32**> | Open predicted trials in a phase bucket the backtest measured (error_phase_mae_months not null). | [optional]
**predicted_months_median** | Option<**f64**> | Median of the carried predicted months over open predicted trials (capped at horizon_cap_months). NULL under min_n. | [optional]
**predicted_months_p25** | Option<**f64**> | 25th percentile, same rows and floor. | [optional]
**predicted_months_p75** | Option<**f64**> | 75th percentile, same rows and floor. | [optional]
**baseline_months_median** | Option<**f64**> | Median of the carried baseline months over the same open predicted trials. NULL under min_n. | [optional]
**n_realized** | Option<**i32**> | Trials with a registered span (primary completion typed ACTUAL today, plausible start and completion), from the realized-duration input. | [optional]
**realized_months_median** | Option<**f64**> | Median registered start to primary completion in months over n_realized trials. NULL under min_n. | [optional]
**realized_months_p25** | Option<**f64**> | 25th percentile, same rows and floor. | [optional]
**realized_months_p75** | Option<**f64**> | 75th percentile, same rows and floor. | [optional]
**n_realized_over_cap** | Option<**i32**> | Realized spans above horizon_cap_months, so a reader sees how many would have hit the cap. | [optional]
**min_n** | Option<**i32**> | The floor applied: a median is withheld below this many trials, carried so a reader sees the floor used. | [optional]
**rollup_status** | Option<**String**> | rolled_up (open predicted >= min_n), below_floor (1 <= open < min_n, counts only), none_predicted (evaluated, no open prediction), none_evaluated. | [optional]
**realized_status** | Option<**String**> | rolled_up (n_realized >= min_n), below_floor, or none. | [optional]
**training_cutoff** | Option<[**String**](string.md)> | The date the model's training data ends. | [optional]
**model_name** | Option<**String**> | The recruitment-duration model. | [optional]
**model_version** | Option<**String**> | The model version. | [optional]
**horizon_cap_months** | Option<**i32**> | The cap on predicted months (36 today), from the prediction rows. | [optional]
**method** | Option<**String**> | The run's method sentence from the run provenance, never re-authored. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


