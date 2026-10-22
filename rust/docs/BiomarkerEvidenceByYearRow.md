# BiomarkerEvidenceByYearRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**route_key** | Option<**String**> | The gene's route key, the entry key for the biomarkers object type; the gene symbol. | [optional]
**gene_symbol** | Option<**String**> | The gene symbol. | [optional]
**evidence_type** | Option<**String**> | The evidence arm. trial_mention counts trials that mention the gene (the year is the trial start year); publication_mention counts papers cited by registered trials that mention it (the year is the publication year). | [optional]
**evidence_year** | Option<**i32**> | The year of the mentions. Null for the undated class. | [optional]
**year_class** | Option<**String**> | How the year is treated. plotted: inside the window the web chart draws. pre_window: before 1990. future_dated: after the current year. corpus_incomplete: after the latest year the arm's corpus covers. undated: the source has no year. | [optional]
**n_mentions** | Option<**i32**> | Mentions of the gene in that arm, year and class. A document that mentions the gene more than once counts once per mention. | [optional]
**n_documents** | Option<**i32**> | Distinct documents (trials or papers) that mention the gene in that arm, year and class. | [optional]
**year_denominator** | Option<**i32**> | Distinct documents in that arm and year that mention any gene, the denominator of the gene's share of the year. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


