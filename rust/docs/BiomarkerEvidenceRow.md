# BiomarkerEvidenceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**biomarker_evidence_id** | Option<**String**> | The evidence assertion id. | [optional]
**route_key** | Option<**String**> | The biomarker route key of the gene this evidence is about (the gene symbol); the entry key for the biomarkers object type. | [optional]
**gene_symbol** | Option<**String**> | The gene symbol the evidence is about. | [optional]
**target_id** | Option<**String**> | The target's Ensembl gene id. | [optional]
**uniprot_accession** | Option<**String**> | The UniProt accession. | [optional]
**biomarker** | Option<**String**> | The biomarker as the source names it (the gene or the variant). | [optional]
**evidence_type** | Option<**String**> | cancer_actionable_cgi, ot_association or trial_mention. | [optional]
**source** | Option<**String**> | The source the assertion comes from. | [optional]
**condition_id** | Option<**String**> | The condition key. | [optional]
**condition_name** | Option<**String**> | The condition as the source names it. | [optional]
**therapeutic_area** | Option<**String**> | The therapeutic area. | [optional]
**drug_name** | Option<**String**> | The drug as filed. | [optional]
**drug_id** | Option<**String**> | The drug key. | [optional]
**molecule_id** | Option<**String**> | The molecule key. Null today: the evidence mart carries no molecule column. | [optional]
**response** | Option<**String**> | The reported response (for example Responsive, Resistant). | [optional]
**score** | Option<**f64**> | The association score, for an Open Targets row. | [optional]
**variant_count** | Option<**i32**> | Variants the assertion covers. | [optional]
**trial_id** | Option<**String**> | The trial registry id, for a trial-mention row. | [optional]
**sample_type** | Option<**String**> | The sample type, for a trial-mention row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


