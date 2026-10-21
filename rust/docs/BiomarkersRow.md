# BiomarkersRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**biomarker_id** | Option<**String**> | The biomarker key. | [optional]
**route_key** | Option<**String**> | The biomarker's route key, the id the web page uses and the one entry key every biomarker dataset takes (the gene symbol for a gene biomarker). | [optional]
**biomarker_kind** | Option<**String**> | gene, protein_variant or another feature kind. | [optional]
**identity_status** | Option<**String**> | Whether the biomarker's identity is confirmed. | [optional]
**display_name** | Option<**String**> | The biomarker's display name. | [optional]
**parent_biomarker_id** | Option<**String**> | The parent biomarker (a variant's gene). | [optional]
**gene_symbol** | Option<**String**> | The gene symbol. | [optional]
**amino_acid_change** | Option<**String**> | The amino-acid change, for a variant. | [optional]
**target_id** | Option<**String**> | The target's Ensembl gene id. | [optional]
**hgnc_id** | Option<**String**> | The HUGO Gene Nomenclature Committee (HGNC) gene id. | [optional]
**uniprot_accession** | Option<**String**> | The UniProt accession. | [optional]
**target_protein_name** | Option<**String**> | The recommended name of the protein the biomarker's gene encodes: the UniProt protein of the accession, else the first UniProt protein of the gene symbol, else the Open Targets target name. Null when none names it. | [optional]
**feature_code** | Option<**String**> | Feature code for a non-gene feature. | [optional]
**feature_type** | Option<**String**> | Feature type for a non-gene feature. | [optional]
**evidence_tier** | Option<**i32**> | Strongest evidence tier: 1 companion diagnostic of the US Food and Drug Administration (FDA) through 6 no evidence. | [optional]
**evidence_tier_label** | Option<**String**> | The tier's name. | [optional]
**n_fda_cdx_assertions** | Option<**i32**> | FDA companion diagnostic assertions. | [optional]
**n_fda_bqp_assertions** | Option<**i32**> | FDA biomarker qualification project assertions. | [optional]
**n_cancer_actionable_cgi** | Option<**i32**> | Cancer Genome Interpreter actionable rows. | [optional]
**n_ot_association** | Option<**i32**> | Open Targets target-disease associations. | [optional]
**n_trial_mention** | Option<**i32**> | Trials that name the biomarker. | [optional]
**n_publication_variant_mentions** | Option<**i32**> | Publication variant mentions. | [optional]
**has_clinical_validation_evidence** | Option<**bool**> | Whether clinical validation evidence exists. | [optional]
**has_analytical_validation_evidence** | Option<**bool**> | Whether analytical validation evidence exists. | [optional]
**n_assay_provider_papers** | Option<**i32**> | Papers naming an assay provider for the biomarker. | [optional]
**n_assay_provider_suppliers** | Option<**i32**> | Assay-provider suppliers named. | [optional]
**n_diagnostic_provider_suppliers** | Option<**i32**> | Of those, diagnostic providers. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


