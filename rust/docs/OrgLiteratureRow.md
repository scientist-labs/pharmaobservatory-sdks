# OrgLiteratureRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**org_id** | Option<**String**> | Legal-entity level: the resolved parent organization record the byline names (edgar:..., ror:... or sponsor:...), the entry key for the legal_entities object type. | [optional]
**pmid** | Option<**String**> | The PubMed id. | [optional]
**title** | Option<**String**> | The paper's title. | [optional]
**journal_title** | Option<**String**> | The journal's title. | [optional]
**pub_year** | Option<**i32**> | The publication year (2016 or later: the dataset holds recent MeSH-indexed literature). | [optional]
**evidence_type** | Option<**String**> | The paper's study-design label, one normalized publication type by study-design priority. | [optional]
**is_trial_linked** | Option<**bool**> | True when a registered clinical trial cites the paper; false for literature no registered trial cites. | [optional]
**org_arm** | Option<**String**> | company when the organization is company-typed (matched on a legal-form token in the affiliation); institution for every other organization. | [optional]
**match_tier** | Option<**String**> | The strongest tier that attributed a byline to the organization: legal_form or both (company arm), high or medium (institution arm). | [optional]
**n_bylines** | Option<**i32**> | The number of distinct author positions on the paper whose affiliation named the organization. | [optional]
**citation_count** | Option<**i32**> | Citations where iCite (the National Institutes of Health citation-metrics service) carries them; null is not measured, never zero. | [optional]
**relative_citation_ratio** | Option<**f64**> | The iCite Relative Citation Ratio (RCR) where measured; null is not measured. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


