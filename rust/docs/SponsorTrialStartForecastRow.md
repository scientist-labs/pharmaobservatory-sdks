# SponsorTrialStartForecastRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sponsor_org_id** | Option<**String**> | Legal-entity level: the organization record the forecast series is for. | [optional]
**sponsor_name** | Option<**String**> | Display name of the organization record. | [optional]
**sponsor_company_id** | Option<**String**> | Company level: the company the organization belongs to (null when unresolved). | [optional]
**sponsor_company_name** | Option<**String**> | Display name of the sponsor's company. | [optional]
**sponsor_group_id** | Option<**String**> | Group level: the sponsor's current ultimate parent company; a root company's group is itself. | [optional]
**sponsor_group_name** | Option<**String**> | Display name of the sponsor's group. | [optional]
**target_year** | Option<**i32**> | The calendar year forecast. | [optional]
**horizon_offset** | Option<**i32**> | Years ahead of the build year (0 is the current year). | [optional]
**months_observed** | Option<**i32**> | Months of the target year already observed. | [optional]
**months_forecast** | Option<**i32**> | Months of the target year forecast. | [optional]
**input_observed_through_cutoff** | Option<**i32**> | Trials started in the target year up to the training cutoff. | [optional]
**input_planned_remaining** | Option<**i32**> | Trials registered as planned to start in the rest of the target year. | [optional]
**forecast_point** | Option<**f64**> | The forecast trial starts for the whole target year. A forecast of the eventual count, never an actual count, and it carries no interval. | [optional]
**forecast_status** | Option<**String**> | The series' forecast status. | [optional]
**input_trials_window_60m** | Option<**i32**> | Trials the organization started in the last 60 months. | [optional]
**input_starts_per_year_60m** | Option<**f64**> | Average trial starts per year over the last 60 months. | [optional]
**size_bucket** | Option<**String**> | The series size bucket the backtest error is measured in. | [optional]
**backtest_horizon** | Option<**String**> | The horizon the backtest error was measured at. | [optional]
**backtest_smape_pct** | Option<**f64**> | Backtest symmetric mean absolute percentage error. | [optional]
**backtest_rel_bias** | Option<**f64**> | Backtest relative bias. | [optional]
**backtest_relmae_vs_snaive** | Option<**f64**> | Backtest mean absolute error relative to the seasonal naive baseline. | [optional]
**backtest_n_series** | Option<**i32**> | Series the backtest covered. | [optional]
**backtest_ref** | Option<**String**> | The backtest run reference. | [optional]
**model_name** | Option<**String**> | The forecasting model. | [optional]
**model_version** | Option<**String**> | The model version. | [optional]
**training_cutoff** | Option<[**String**](string.md)> | The date the model's training data ends. | [optional]
**method** | Option<**String**> | How the forecast was produced. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


