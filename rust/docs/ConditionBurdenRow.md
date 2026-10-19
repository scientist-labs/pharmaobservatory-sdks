# ConditionBurdenRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**condition_id** | Option<**String**> | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional]
**condition_name** | Option<**String**> | The condition's MeSH heading. | [optional]
**cause_name** | Option<**String**> | The Global Burden of Disease cause matched to the condition. | [optional]
**dalys_millions** | Option<**f64**> | Global disability-adjusted life years for the matched cause, in millions. | [optional]
**deaths_millions** | Option<**f64**> | Global deaths for the matched cause, in millions. | [optional]
**druggability** | Option<**i32**> | Druggability class of the matched cause: 1 highly druggable, 2 druggable, 3 challenging, 4 very challenging. | [optional]
**n_trials** | Option<**i64**> | Trials for the condition, counted under each condition they are tagged with and each condition above it in the MeSH tree. | [optional]
**trials_per_million_dalys** | Option<**f64**> | Trials per million DALYs: n_trials divided by dalys_millions. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


