# TrialPublicationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**pmid** | Option<**String**> | The PubMed id of the paper. | [optional]
**title** | Option<**String**> | The paper's title. | [optional]
**journal_title** | Option<**String**> | The journal's title (null when the paper's journal is not in the journal table). | [optional]
**pub_year** | Option<**i32**> | The publication year (null when undated). | [optional]
**evidence_type** | Option<**String**> | The paper's study-design label, one normalized publication type by study-design priority; the page's design column. | [optional]
**citation** | Option<**String**> | The citation for the paper: one citation per paper, the longest among the trials that cite the PubMed id. | [optional]
**reference_types** | Option<**Vec<String>**> | Every reference kind the registry filed for this paper on this trial, lower-cased (background, result, derived). | [optional]
**results** | Option<**bool**> | True when any reference kind of the row is a results reference; the page's Results label. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


