# ConditionPublicationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**condition_id** | Option<**String**> | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional]
**pmid** | Option<**String**> | The PubMed id. | [optional]
**title** | Option<**String**> | The article title. | [optional]
**pub_year** | Option<**i32**> | Publication year. | [optional]
**journal_title** | Option<**String**> | The journal title. | [optional]
**evidence_type** | Option<**String**> | The article's PubMed publication type, one per article, picked by study design (Randomized Controlled Trial, Journal Article and so on). | [optional]
**is_major_topic** | Option<**bool**> | True when the condition is a major topic of the article. | [optional]
**n_linked_trials** | Option<**i64**> | Registry trials linked to the article. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


