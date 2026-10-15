# ApiV1BiomarkerAssayWhitespaceGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**mode** | **String** |  | 
**gene** | Option<**String**> |  | [optional]
**gene_recognized** | Option<**bool**> |  | [optional]
**in_population** | Option<**bool**> |  | [optional]
**build_state** | [**models::ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](_api_v1_biomarker_assay_screen_get_200_response_build_state.md) |  | 
**filters** | Option<[**models::ApiV1BiomarkerAssayWhitespaceGet200ResponseFilters**](_api_v1_biomarker_assay_whitespace_get_200_response_filters.md)> |  | [optional]
**bucket_counts** | Option<[**models::ApiV1BiomarkerAssayWhitespaceGet200ResponseBucketCounts**](_api_v1_biomarker_assay_whitespace_get_200_response_bucket_counts.md)> |  | [optional]
**n_genes_matching** | Option<**i32**> |  | [optional]
**n_genes_returned** | Option<**i32**> |  | [optional]
**truncated** | Option<**bool**> |  | [optional]
**genes** | Option<[**Vec<models::ApiV1BiomarkerAssayWhitespaceGet200ResponseGenesInner>**](_api_v1_biomarker_assay_whitespace_get_200_response_genes_inner.md)> |  | [optional]
**profile_link** | Option<[**models::ApiV1BiomarkerAssayWhitespaceGet200ResponseProfileLink**](_api_v1_biomarker_assay_whitespace_get_200_response_profile_link.md)> |  | [optional]
**gene_symbol** | Option<**String**> |  | [optional]
**biomarker_display_name** | Option<**String**> |  | [optional]
**whitespace_bucket** | Option<**String**> |  | [optional]
**whitespace_bucket_label** | Option<**String**> |  | [optional]
**whitespace_bucket_rule_id** | Option<**String**> |  | [optional]
**gene_evidence_tier** | Option<**String**> |  | [optional]
**gene_evidence_tier_label** | Option<**String**> |  | [optional]
**gene_supply_status** | Option<**String**> |  | [optional]
**gene_supply_status_label** | Option<**String**> |  | [optional]
**pipeline_program_count** | Option<[**serde_json::Value**](.md)> |  | [optional]
**pipeline_max_phase** | Option<[**serde_json::Value**](.md)> |  | [optional]
**validated_evidence** | Option<[**Vec<models::ApiV1BiomarkerAssayWhitespaceGet200ResponseValidatedEvidenceInner>**](_api_v1_biomarker_assay_whitespace_get_200_response_validated_evidence_inner.md)> |  | [optional]
**literature_evidence** | Option<[**Vec<models::ApiV1BiomarkerAssayWhitespaceGet200ResponseLiteratureEvidenceInner>**](_api_v1_biomarker_assay_whitespace_get_200_response_literature_evidence_inner.md)> |  | [optional]
**confidence_caveat** | Option<**String**> |  | [optional]
**supplier_declarations** | Option<[**Vec<models::ApiV1BiomarkerAssayWhitespaceGet200ResponseSupplierDeclarationsInner>**](_api_v1_biomarker_assay_whitespace_get_200_response_supplier_declarations_inner.md)> |  | [optional]
**validated_evidence_built** | Option<**bool**> |  | [optional]
**literature_evidence_built** | Option<**bool**> |  | [optional]
**supplier_declarations_built** | Option<**bool**> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


