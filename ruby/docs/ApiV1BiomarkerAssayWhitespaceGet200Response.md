# PharmaObservatory::ApiV1BiomarkerAssayWhitespaceGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **mode** | **String** |  |  |
| **gene** | **String** |  | [optional] |
| **gene_recognized** | **Boolean** |  | [optional] |
| **in_population** | **Boolean** |  | [optional] |
| **build_state** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  |  |
| **filters** | [**ApiV1BiomarkerAssayWhitespaceGet200ResponseFilters**](ApiV1BiomarkerAssayWhitespaceGet200ResponseFilters.md) |  | [optional] |
| **bucket_counts** | [**ApiV1BiomarkerAssayWhitespaceGet200ResponseBucketCounts**](ApiV1BiomarkerAssayWhitespaceGet200ResponseBucketCounts.md) |  | [optional] |
| **n_genes_matching** | **Integer** |  | [optional] |
| **n_genes_returned** | **Integer** |  | [optional] |
| **truncated** | **Boolean** |  | [optional] |
| **genes** | [**Array&lt;ApiV1BiomarkerAssayWhitespaceGet200ResponseGenesInner&gt;**](ApiV1BiomarkerAssayWhitespaceGet200ResponseGenesInner.md) |  | [optional] |
| **profile_link** | [**ApiV1BiomarkerAssayWhitespaceGet200ResponseProfileLink**](ApiV1BiomarkerAssayWhitespaceGet200ResponseProfileLink.md) |  | [optional] |
| **gene_symbol** | **String** |  | [optional] |
| **biomarker_display_name** | **String** |  | [optional] |
| **whitespace_bucket** | **String** |  | [optional] |
| **whitespace_bucket_label** | **String** |  | [optional] |
| **whitespace_bucket_rule_id** | **String** |  | [optional] |
| **gene_evidence_tier** | **String** |  | [optional] |
| **gene_evidence_tier_label** | **String** |  | [optional] |
| **gene_supply_status** | **String** |  | [optional] |
| **gene_supply_status_label** | **String** |  | [optional] |
| **pipeline_program_count** | **Object** |  | [optional] |
| **pipeline_max_phase** | **Object** |  | [optional] |
| **validated_evidence** | [**Array&lt;ApiV1BiomarkerAssayWhitespaceGet200ResponseValidatedEvidenceInner&gt;**](ApiV1BiomarkerAssayWhitespaceGet200ResponseValidatedEvidenceInner.md) |  | [optional] |
| **literature_evidence** | [**Array&lt;ApiV1BiomarkerAssayWhitespaceGet200ResponseLiteratureEvidenceInner&gt;**](ApiV1BiomarkerAssayWhitespaceGet200ResponseLiteratureEvidenceInner.md) |  | [optional] |
| **confidence_caveat** | **String** |  | [optional] |
| **supplier_declarations** | [**Array&lt;ApiV1BiomarkerAssayWhitespaceGet200ResponseSupplierDeclarationsInner&gt;**](ApiV1BiomarkerAssayWhitespaceGet200ResponseSupplierDeclarationsInner.md) |  | [optional] |
| **validated_evidence_built** | **Boolean** |  | [optional] |
| **literature_evidence_built** | **Boolean** |  | [optional] |
| **supplier_declarations_built** | **Boolean** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1BiomarkerAssayWhitespaceGet200Response.new(
  mode: null,
  gene: null,
  gene_recognized: null,
  in_population: null,
  build_state: null,
  filters: null,
  bucket_counts: null,
  n_genes_matching: null,
  n_genes_returned: null,
  truncated: null,
  genes: null,
  profile_link: null,
  gene_symbol: null,
  biomarker_display_name: null,
  whitespace_bucket: null,
  whitespace_bucket_label: null,
  whitespace_bucket_rule_id: null,
  gene_evidence_tier: null,
  gene_evidence_tier_label: null,
  gene_supply_status: null,
  gene_supply_status_label: null,
  pipeline_program_count: null,
  pipeline_max_phase: null,
  validated_evidence: null,
  literature_evidence: null,
  confidence_caveat: null,
  supplier_declarations: null,
  validated_evidence_built: null,
  literature_evidence_built: null,
  supplier_declarations_built: null
)
```

