# TargetLiteratureRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_id** | Option<**String**> | The target's Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. | [optional]
**target_symbol** | Option<**String**> | The target's HGNC gene symbol. | [optional]
**pmid** | Option<**String**> | The PubMed id. | [optional]
**title** | Option<**String**> | The paper's title. | [optional]
**journal_title** | Option<**String**> | The journal's title. | [optional]
**pub_year** | Option<**i32**> | The publication year (2016 or later: the dataset holds recent MeSH-indexed literature). | [optional]
**evidence_type** | Option<**String**> | The paper's study-design label, one normalized publication type by study-design priority. | [optional]
**is_trial_linked** | Option<**bool**> | True when a registered clinical trial cites the paper; false for literature no registered trial cites. | [optional]
**citation_count** | Option<**i32**> | Citations where iCite (the National Institutes of Health citation-metrics service) carries them; null is not measured, never zero. | [optional]
**relative_citation_ratio** | Option<**f64**> | The iCite Relative Citation Ratio (RCR) where measured; null is not measured. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


