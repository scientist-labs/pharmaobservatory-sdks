# ConditionAdverseEventsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**condition_id** | Option<**String**> | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional]
**adverse_event_term** | Option<**String**> | The adverse event term as the trials report it. | [optional]
**subjects_affected** | Option<**f64**> | Subjects affected, summed across the condition's trials. | [optional]
**n_trials** | Option<**i64**> | Trials for this condition that report the term. | [optional]
**n_terms_total** | Option<**i64**> | Distinct serious adverse event terms across the condition's trials. The rows are only the top 10 of these, so a row count is not the term count; use this. | [optional]
**n_trials_total** | Option<**i64**> | Results-posting trials of the condition that report at least one serious adverse event term. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


