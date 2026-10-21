# ConditionTrialStartForecastRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**condition_id** | Option<**String**> | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional]
**condition_name** | Option<**String**> | The condition's MeSH heading. | [optional]
**target_year** | Option<**i32**> | The calendar year forecast. | [optional]
**horizon_offset** | Option<**i32**> | Years ahead of the build year (0 is the current year). | [optional]
**months_observed** | Option<**i32**> | Months of the target year already observed. | [optional]
**months_forecast** | Option<**i32**> | Months of the target year forecast. | [optional]
**input_observed_through_cutoff** | Option<**i32**> | Trials started in the target year up to the training cutoff. | [optional]
**input_planned_remaining** | Option<**i32**> | Trials registered as planned to start in the rest of the target year. | [optional]
**forecast_point** | Option<**f64**> | The forecast trial starts for the whole target year. A forecast of the eventual count, never an actual count, and it carries no interval. | [optional]
**forecast_status** | Option<**String**> | The series' forecast status. | [optional]
**input_trials_window_60m** | Option<**i32**> | Trials of the condition (as in condition_trials) with a start date in the 60 calendar months ending with the cutoff month. | [optional]
**input_starts_per_year_60m** | Option<**f64**> | Average trial starts per year over the last 60 months. | [optional]
**size_bucket** | Option<**String**> | The series size bucket the backtest error is measured in. | [optional]
**backtest_horizon** | Option<**String**> | The horizon the backtest error was measured at. | [optional]
**backtest_smape_pct** | Option<**f64**> | Backtest symmetric mean absolute percentage error. | [optional]
**backtest_rel_bias** | Option<**f64**> | Backtest relative bias. | [optional]
**backtest_relmae_vs_snaive** | Option<**f64**> | Backtest mean absolute error relative to the seasonal naive baseline. | [optional]
**backtest_n_series** | Option<**i32**> | Series the backtest covered. | [optional]
**backtest_series_kind** | Option<**String**> | The kind of series the backtest error was measured on. | [optional]
**backtest_ref** | Option<**String**> | The backtest run reference. | [optional]
**planned_start_coefficient** | Option<**f64**> | The multiple of input_planned_remaining the planned-start method adds to the months after the training cutoff; null on a row built with the flat mean alone. | [optional]
**planned_start_stricter_check_agrees** | Option<**bool**> | On a planned-start row, whether the stricter re-check (trained only on target years at least two before the cutoff year) adopted the same cell; null on a flat-mean row. | [optional]
**planned_start_stricter_check_rel_bias** | Option<**f64**> | On a planned-start row, the relative bias (forecast minus actual over actual) of the stricter re-check; null on a flat-mean row. | [optional]
**model_name** | Option<**String**> | The forecasting model. | [optional]
**model_version** | Option<**String**> | The model version. | [optional]
**training_cutoff** | Option<[**String**](string.md)> | The date the model's training data ends. | [optional]
**method** | Option<**String**> | How the forecast was produced. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


