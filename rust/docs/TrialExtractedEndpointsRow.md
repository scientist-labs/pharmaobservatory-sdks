# TrialExtractedEndpointsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**result_reported** | Option<**bool**> | Whether the extraction found a formal primary-endpoint result. | [optional]
**signal** | Option<**String**> | The extracted reading, verbatim: met, missed, mixed, reported_indeterminate or not_reported. | [optional]
**outcome_category** | Option<**String**> | The signal grouped as met, missed, mixed or unclear. | [optional]
**has_clean_outcome** | Option<**bool**> | Whether outcome_category is met, missed or mixed. | [optional]
**primary_endpoints** | Option<**Vec<String>**> | The primary endpoint descriptions the extraction found. | [optional]
**primary_endpoint_count** | Option<**i32**> | Number of primary endpoints extracted. | [optional]
**confidence** | Option<**f64**> | The extraction's confidence, from 0 to 1. | [optional]
**evidence_quote** | Option<**String**> | The text quoted as evidence for the reading. | [optional]
**source_url** | Option<**String**> | The registry page the reading was taken from. | [optional]
**model** | Option<**String**> | The model that produced the reading. | [optional]
**extracted_at** | Option<**String**> | When the extraction ran. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


