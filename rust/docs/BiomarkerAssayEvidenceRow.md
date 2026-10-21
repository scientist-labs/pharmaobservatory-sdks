# BiomarkerAssayEvidenceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**biomarker_assay_evidence_id** | Option<**String**> | The extracted mention's id. | [optional]
**route_key** | Option<**String**> | The route key of the gene the mention resolved to, the entry key for the biomarkers object type. Null when the gene did not resolve. | [optional]
**biomarker_id** | Option<**String**> | The resolved gene's biomarker key. Null when the gene did not resolve. | [optional]
**document_id** | Option<**String**> | The id of the source document the facts were extracted from. | [optional]
**pmid** | Option<**String**> | The PubMed id of the paper. | [optional]
**biomarker_ordinal** | Option<**i32**> | The position of the mention in the paper's extracted biomarker list, from 1. | [optional]
**biomarker_name_as_extracted** | Option<**String**> | The biomarker name as extracted. | [optional]
**gene_symbol_as_extracted** | Option<**String**> | The gene symbol as extracted, before resolution. Null when the mention names no gene. | [optional]
**biomarker_synonyms** | Option<**Vec<String>**> | The synonyms the paper gives, a document-level list repeated on each of its rows. | [optional]
**gene_resolution_status** | Option<**String**> | How the mention resolved to a gene. resolved_direct: exact symbol match. resolved_alias: an unambiguous dictionary alias. unresolved: a symbol was extracted that matches nothing or an ambiguous alias; not guessed. not_applicable: no gene symbol was extracted. | [optional]
**gene_resolution_method** | Option<**String**> | exact_gene_symbol or dictionary_alias; null when the gene did not resolve. | [optional]
**assay_name** | Option<**String**> | The assay, trade name or platform the paper names. A document-level value. | [optional]
**specimen_type** | Option<**String**> | The specimen type (blood, plasma, tissue). A document-level value. | [optional]
**comparator_method** | Option<**String**> | The reference method the performance was measured against, when reported. | [optional]
**reported_performance** | Option<**String**> | The performance as the paper states it (sensitivity, specificity, area under the curve), as text and not computed. | [optional]
**validation_type_mentioned** | Option<**String**> | What the paper claims about the validation type: a fact about the text, not a tier. | [optional]
**assay_provider_company** | Option<**Vec<String>**> | Organization names as extracted that run or provide the assay. Not resolved to a company. | [optional]
**funding_organizations** | Option<**Vec<String>**> | Organization names as extracted that funded the study. Never merged with the assay provider. | [optional]
**evidence_type** | Option<**String**> | Always literature_assay_validation, which tells these rows from the other evidence types. | [optional]
**evidence_quote** | Option<**Vec<String>**> | The sentences of the paper the facts were read from. Null when the extraction gave none. | [optional]
**confidence** | Option<**String**> | The extraction model's self-reported confidence. Not a reliability signal: it does not track whether the extraction is correct; never use it as a quality gate. | [optional]
**title** | Option<**String**> | The paper's title. | [optional]
**published_at** | Option<**String**> | The paper's publication timestamp, in UTC. Null for about one paper in seven, whose source gives a bare year. | [optional]
**extraction_model** | Option<**String**> | The model that extracted the facts. | [optional]
**extracted_at** | Option<**String**> | When the facts were extracted, in UTC. | [optional]
**source_url** | Option<**String**> | The address of the paper the facts were extracted from. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


