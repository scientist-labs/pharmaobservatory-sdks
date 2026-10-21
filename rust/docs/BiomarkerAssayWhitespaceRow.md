# BiomarkerAssayWhitespaceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**biomarker_id** | Option<**String**> | The gene's biomarker key. | [optional]
**route_key** | Option<**String**> | The gene's route key, the entry key for the biomarkers object type; the gene symbol. | [optional]
**gene_symbol** | Option<**String**> | The gene symbol. | [optional]
**biomarker_display_name** | Option<**String**> | The gene's display name. | [optional]
**biomarker_identity_status** | Option<**String**> | Whether the gene's identity is confirmed. | [optional]
**target_id** | Option<**String**> | The target's Ensembl gene id; null when the gene does not resolve one, in which case the pipeline columns are null. | [optional]
**whitespace_bucket** | Option<**String**> | The classification, one of nine values crossing gene_evidence_tier with gene_supply_status. Labels: no_evidence_no_supply 'No evidence, explicit no-assay'; no_evidence_supply_ambiguous 'No evidence, ambiguous supply'; supply_no_evidence 'Supply declared, no evidence'; literature_evidence_no_supply 'Literature evidence, no confirmed supply'; literature_evidence_supply_ambiguous 'Literature evidence, ambiguous supply'; literature_evidence_served 'Literature evidence, supply declared'; validated_evidence_no_supply 'Validated-ladder evidence, no confirmed supply'; validated_evidence_supply_ambiguous 'Validated-ladder evidence, ambiguous supply'; served 'Validated-ladder evidence, supply declared'. | [optional]
**whitespace_bucket_rule_id** | Option<**String**> | The rule that assigned the bucket, one rule per bucket. | [optional]
**gene_evidence_tier** | Option<**String**> | Evidence axis, a precedence and not a score. validated: at least one row of the validation ladder (biomarker_validation) for the gene. literature_only: no validated row but at least one extracted PubMed assay fact. none: neither. | [optional]
**gene_supply_status** | Option<**String**> | Supply axis, a precedence. declared_positive: a supplier declares an assay (validated or not). declared_ambiguous: a response conflicts or gives detail with no option chosen. declared_no_assay: a supplier answered that it has no assay. not_on_survey: no answer at all. | [optional]
**validated_evidence_row_count** | Option<**i32**> | Rows of the validation ladder for the gene (biomarker_validation). A count, not a score, and not combined with the literature count. | [optional]
**literature_evidence_row_count** | Option<**i32**> | Extracted PubMed assay facts resolved to the gene (biomarker_assay_evidence). A count, not a score. | [optional]
**evidence_types_present** | Option<**Vec<String>**> | The sorted set of evidence types behind the gene's evidence, including literature_assay_validation when literature evidence exists. A set, never a rank. | [optional]
**evidence_tier_rule_ids_present** | Option<**Vec<String>**> | The sorted set of validation ladder rule ids behind the validated evidence. For display; do not derive a rank from it. | [optional]
**evidence_validation_ids** | Option<**Vec<String>**> | biomarker_validation_id values of the validated evidence. Empty when gene_evidence_tier is not validated. | [optional]
**literature_evidence_ids** | Option<**Vec<String>**> | biomarker_assay_evidence_id values of the extracted PubMed facts. Populated whichever evidence tier governs the bucket. | [optional]
**literature_evidence_document_ids** | Option<**Vec<String>**> | The source document ids behind literature_evidence_ids. | [optional]
**supplier_declaration_ids** | Option<**Vec<String>**> | Ids of the supplier survey declarations behind gene_supply_status (the winning precedence tier). Empty when not_on_survey. The declarations are not a published dataset. | [optional]
**supplier_positive_response_count** | Option<**i32**> | Supplier declarations that an assay exists, validated or unvalidated, for the gene. | [optional]
**supplier_conflicting_response_count** | Option<**i32**> | Declarations whose own cells disagree. | [optional]
**supplier_detail_only_response_count** | Option<**i32**> | Declarations that give free text and no selected option. | [optional]
**supplier_no_assay_response_count** | Option<**i32**> | Declarations that the supplier has no assay for the gene. | [optional]
**pipeline_program_count** | Option<**i32**> | Drug programs against the gene's target across every phase band. Null means the gene has no target row, never zero activity. Heat only; not part of the bucket. | [optional]
**pipeline_active_program_count** | Option<**i32**> | Active drug programs against the gene's target. Same nullability as pipeline_program_count. | [optional]
**pipeline_max_phase** | Option<**i32**> | The highest phase band (0 to 4) any program against the target reached; the app maps it to a phase label. Null as pipeline_program_count. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


