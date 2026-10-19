# TrialDemographicsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**has_sex_data** | Option<**bool**> | Whether a sex baseline breakdown was posted. | [optional]
**has_race_data** | Option<**bool**> | Whether a race baseline breakdown was posted. | [optional]
**has_ethnicity_data** | Option<**bool**> | Whether an ethnicity baseline breakdown was posted. | [optional]
**has_age_data** | Option<**bool**> | Whether a categorical age baseline breakdown was posted. | [optional]
**pct_female** | Option<**f64**> | Percent of baseline participants recorded as female (null when sex was not posted). | [optional]
**race_breakdown** | Option<**String**> | A JSON object of race category to percent of baseline participants, as posted. | [optional]
**n_analyzed** | Option<**i32**> | Baseline participants analyzed. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


