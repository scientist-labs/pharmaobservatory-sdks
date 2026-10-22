# OwnershipSourceCoverageRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**source_class** | Option<**String**> | The ownership source class (for example sec_def_14a for proxy statements, sec_form_adv for Form ADV). | [optional]
**subject_source_ref** | Option<**String**> | The subject as the source names it: the source's own reference for the owned company. | [optional]
**subject_name** | Option<**String**> | The subject's name as filed. | [optional]
**subject_org_id** | Option<**String**> | Legal-entity level: the organization record of the subject. | [optional]
**subject_company_id** | Option<**String**> | Company level: the company the subject belongs to (null when unresolved). | [optional]
**subject_company_name** | Option<**String**> | Display name of the subject's company. | [optional]
**subject_group_id** | Option<**String**> | Group level: the subject's current ultimate parent company; a root company's group is itself. | [optional]
**subject_group_name** | Option<**String**> | Display name of the subject's group. | [optional]
**source_system** | Option<**String**> | The system the source class is read from (for example sec_edgar). | [optional]
**source_load_state** | Option<**String**> | The whole-source state: loaded, served_by_dedicated_mart (13F is served by the institutional holdings datasets), no_documents_ingested or source_not_ingested. | [optional]
**source_documents_ingested** | Option<**i64**> | Documents the whole source has ingested. | [optional]
**coverage_state** | Option<**String**> | What the source read about this subject: reported, no_reportable_owner, extraction_abstained or read_no_owner. | [optional]
**documents_ingested** | Option<**i64**> | Documents the source ingested about this subject. | [optional]
**observations** | Option<**i64**> | Ownership observations the source recorded about this subject. | [optional]
**current_edges** | Option<**i64**> | Ownership relationships that currently stand. | [optional]
**owners_distinct** | Option<**i64**> | Distinct owners the source named for this subject. | [optional]
**owners_unresolved** | Option<**i64**> | Named owners that did not resolve to an organization or investor. | [optional]
**no_reportable_owner_statements** | Option<**i64**> | Filings that positively stated the subject has no reportable owner. | [optional]
**extraction_abstained** | Option<**i64**> | Documents on which the extractor abstained. | [optional]
**last_filing_date** | Option<[**String**](string.md)> | The latest filing date the source holds for this subject. | [optional]
**max_as_of_date** | Option<[**String**](string.md)> | The latest as-of date among the subject's observations. | [optional]
**source_license** | Option<**String**> | The licence of the source. | [optional]
**attribution** | Option<**String**> | The credit line for the source. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


