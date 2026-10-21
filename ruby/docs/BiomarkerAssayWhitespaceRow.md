# PharmaObservatory::BiomarkerAssayWhitespaceRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **biomarker_id** | **String** | The gene&#39;s biomarker key. | [optional] |
| **route_key** | **String** | The gene&#39;s route key, the entry key for the biomarkers object type; the gene symbol. | [optional] |
| **gene_symbol** | **String** | The gene symbol. | [optional] |
| **biomarker_display_name** | **String** | The gene&#39;s display name. | [optional] |
| **biomarker_identity_status** | **String** | Whether the gene&#39;s identity is confirmed. | [optional] |
| **target_id** | **String** | The target&#39;s Ensembl gene id; null when the gene does not resolve one, in which case the pipeline columns are null. | [optional] |
| **whitespace_bucket** | **String** | The classification, one of nine values crossing gene_evidence_tier with gene_supply_status. Labels: no_evidence_no_supply &#39;No evidence, explicit no-assay&#39;; no_evidence_supply_ambiguous &#39;No evidence, ambiguous supply&#39;; supply_no_evidence &#39;Supply declared, no evidence&#39;; literature_evidence_no_supply &#39;Literature evidence, no confirmed supply&#39;; literature_evidence_supply_ambiguous &#39;Literature evidence, ambiguous supply&#39;; literature_evidence_served &#39;Literature evidence, supply declared&#39;; validated_evidence_no_supply &#39;Validated-ladder evidence, no confirmed supply&#39;; validated_evidence_supply_ambiguous &#39;Validated-ladder evidence, ambiguous supply&#39;; served &#39;Validated-ladder evidence, supply declared&#39;. | [optional] |
| **whitespace_bucket_rule_id** | **String** | The rule that assigned the bucket, one rule per bucket. | [optional] |
| **gene_evidence_tier** | **String** | Evidence axis, a precedence and not a score. validated: at least one row of the validation ladder (biomarker_validation) for the gene. literature_only: no validated row but at least one extracted PubMed assay fact. none: neither. | [optional] |
| **gene_supply_status** | **String** | Supply axis, a precedence. declared_positive: a supplier declares an assay (validated or not). declared_ambiguous: a response conflicts or gives detail with no option chosen. declared_no_assay: a supplier answered that it has no assay. not_on_survey: no answer at all. | [optional] |
| **validated_evidence_row_count** | **Integer** | Rows of the validation ladder for the gene (biomarker_validation). A count, not a score, and not combined with the literature count. | [optional] |
| **literature_evidence_row_count** | **Integer** | Extracted PubMed assay facts resolved to the gene (biomarker_assay_evidence). A count, not a score. | [optional] |
| **evidence_types_present** | **Array&lt;String&gt;** | The sorted set of evidence types behind the gene&#39;s evidence, including literature_assay_validation when literature evidence exists. A set, never a rank. | [optional] |
| **evidence_tier_rule_ids_present** | **Array&lt;String&gt;** | The sorted set of validation ladder rule ids behind the validated evidence. For display; do not derive a rank from it. | [optional] |
| **evidence_validation_ids** | **Array&lt;String&gt;** | biomarker_validation_id values of the validated evidence. Empty when gene_evidence_tier is not validated. | [optional] |
| **literature_evidence_ids** | **Array&lt;String&gt;** | biomarker_assay_evidence_id values of the extracted PubMed facts. Populated whichever evidence tier governs the bucket. | [optional] |
| **literature_evidence_document_ids** | **Array&lt;String&gt;** | The source document ids behind literature_evidence_ids. | [optional] |
| **supplier_declaration_ids** | **Array&lt;String&gt;** | Ids of the supplier survey declarations behind gene_supply_status (the winning precedence tier). Empty when not_on_survey. The declarations are not a published dataset. | [optional] |
| **supplier_positive_response_count** | **Integer** | Supplier declarations that an assay exists, validated or unvalidated, for the gene. | [optional] |
| **supplier_conflicting_response_count** | **Integer** | Declarations whose own cells disagree. | [optional] |
| **supplier_detail_only_response_count** | **Integer** | Declarations that give free text and no selected option. | [optional] |
| **supplier_no_assay_response_count** | **Integer** | Declarations that the supplier has no assay for the gene. | [optional] |
| **pipeline_program_count** | **Integer** | Drug programs against the gene&#39;s target across every phase band. Null means the gene has no target row, never zero activity. Heat only; not part of the bucket. | [optional] |
| **pipeline_active_program_count** | **Integer** | Active drug programs against the gene&#39;s target. Same nullability as pipeline_program_count. | [optional] |
| **pipeline_max_phase** | **Integer** | The highest phase band (0 to 4) any program against the target reached; the app maps it to a phase label. Null as pipeline_program_count. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerAssayWhitespaceRow.new(
  biomarker_id: null,
  route_key: null,
  gene_symbol: null,
  biomarker_display_name: null,
  biomarker_identity_status: null,
  target_id: null,
  whitespace_bucket: null,
  whitespace_bucket_rule_id: null,
  gene_evidence_tier: null,
  gene_supply_status: null,
  validated_evidence_row_count: null,
  literature_evidence_row_count: null,
  evidence_types_present: null,
  evidence_tier_rule_ids_present: null,
  evidence_validation_ids: null,
  literature_evidence_ids: null,
  literature_evidence_document_ids: null,
  supplier_declaration_ids: null,
  supplier_positive_response_count: null,
  supplier_conflicting_response_count: null,
  supplier_detail_only_response_count: null,
  supplier_no_assay_response_count: null,
  pipeline_program_count: null,
  pipeline_active_program_count: null,
  pipeline_max_phase: null
)
```

