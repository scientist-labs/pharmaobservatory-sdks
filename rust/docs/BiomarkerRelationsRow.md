# BiomarkerRelationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**route_key** | Option<**String**> | The biomarker's own route key, the entry key for the biomarkers object type (the gene symbol for a gene, the variant key for a protein variant). | [optional]
**biomarker_id** | Option<**String**> | The biomarker's key. | [optional]
**biomarker_kind** | Option<**String**> | Whether the row describes a gene or a protein variant. | [optional]
**gene_symbol** | Option<**String**> | The gene symbol. | [optional]
**amino_acid_change** | Option<**String**> | The protein change, for a variant. | [optional]
**related_entity_kind** | Option<**String**> | The kind of the related entity (gene, drug or condition). | [optional]
**related_entity_id** | Option<**String**> | The related entity's key in its own vocabulary (a gene symbol, a drug id such as unii:..., or a condition key). | [optional]
**related_display_name** | Option<**String**> | The related entity's name. | [optional]
**relation_type** | Option<**String**> | The source's relation label, as the source states it. | [optional]
**pmid** | Option<**String**> | The PubMed id of the paper the relation was extracted from. | [optional]
**title** | Option<**String**> | The paper's title. | [optional]
**journal_nlm_id** | Option<**String**> | The journal's National Library of Medicine (NLM) id. | [optional]
**journal_title** | Option<**String**> | The journal's title. | [optional]
**pub_year** | Option<**i32**> | The paper's publication year. | [optional]
**n_source_assertions** | Option<**i64**> | Source assertions of this relation in the paper. | [optional]
**n_relation_papers** | Option<**i64**> | Papers that carry this biomarker, related entity and relation label. | [optional]
**evidence_type** | Option<**String**> | The evidence type (a text-mined relation, never a clinical claim). | [optional]
**source** | Option<**String**> | The source of the evidence. | [optional]
**display_order** | Option<**i64**> | The order the web page lists the relation in within its kind. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


