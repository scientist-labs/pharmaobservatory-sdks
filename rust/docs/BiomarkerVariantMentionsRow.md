# BiomarkerVariantMentionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**route_key** | Option<**String**> | The route key of the gene the variants belong to, the entry key for the biomarkers object type. | [optional]
**gene_symbol** | Option<**String**> | The gene symbol. | [optional]
**variant_biomarker_id** | Option<**String**> | The variant's own biomarker key. | [optional]
**variant_route_key** | Option<**String**> | The variant's own route key (the id its web page uses). | [optional]
**display_name** | Option<**String**> | The variant's display name. | [optional]
**amino_acid_change** | Option<**String**> | The protein change, for example V600E. | [optional]
**biomarker_kind** | Option<**String**> | The kind of biomarker the row describes (a protein variant). | [optional]
**evidence_type** | Option<**String**> | The evidence type (a literature mention, never an actionability claim). | [optional]
**source** | Option<**String**> | The source of the evidence. | [optional]
**n_papers** | Option<**i64**> | Distinct papers that name the variant. | [optional]
**first_publication_year** | Option<**i32**> | Earliest publication year. | [optional]
**last_publication_year** | Option<**i32**> | Latest publication year. | [optional]
**sample_pmids** | Option<**Vec<String>**> | Up to five sample PubMed ids, newest first. | [optional]
**sample_citations** | Option<**Vec<String>**> | Up to five sample citations, newest first. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


