# ConditionLiteratureRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**condition_id** | Option<**String**> | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional]
**pmid** | Option<**String**> | The PubMed id. | [optional]
**title** | Option<**String**> | The paper's title. | [optional]
**journal_title** | Option<**String**> | The journal's title. | [optional]
**pub_year** | Option<**i32**> | The publication year (2016 or later: the dataset holds recent MeSH-indexed literature). | [optional]
**evidence_type** | Option<**String**> | The paper's study-design label, one normalized publication type by study-design priority. | [optional]
**is_trial_linked** | Option<**bool**> | True when a registered clinical trial cites the paper; false for literature no registered trial cites. | [optional]
**is_major_topic** | Option<**bool**> | True when the condition is a major MeSH topic of the paper. | [optional]
**citation_count** | Option<**i32**> | Citations where iCite (the National Institutes of Health citation-metrics service) carries them; null is not measured, never zero. | [optional]
**relative_citation_ratio** | Option<**f64**> | The iCite Relative Citation Ratio (RCR) where measured; null is not measured. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


