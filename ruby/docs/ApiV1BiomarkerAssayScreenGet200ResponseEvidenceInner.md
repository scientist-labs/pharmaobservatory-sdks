# PharmaObservatory::ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **biomarker_assay_evidence_id** | **String** |  |  |
| **pmid** | **String** |  |  |
| **gene_symbol_raw** | **String** |  |  |
| **biomarker_id** | **String** |  |  |
| **gene_resolution_status** | **String** |  |  |
| **gene_resolution_method** | **String** |  |  |
| **assay_name** | **String** |  |  |
| **specimen_type** | **String** |  |  |
| **comparator_method** | **String** |  |  |
| **reported_performance** | **String** |  |  |
| **validation_type_mentioned** | **String** |  |  |
| **assay_provider_company** | **Array&lt;String&gt;** |  |  |
| **funding_organizations** | **Array&lt;String&gt;** |  |  |
| **evidence_quote** | **Array&lt;String&gt;** |  |  |
| **confidence** | **String** |  |  |
| **source_ref** | [**ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInnerSourceRef**](ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInnerSourceRef.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner.new(
  biomarker_assay_evidence_id: null,
  pmid: null,
  gene_symbol_raw: null,
  biomarker_id: null,
  gene_resolution_status: null,
  gene_resolution_method: null,
  assay_name: null,
  specimen_type: null,
  comparator_method: null,
  reported_performance: null,
  validation_type_mentioned: null,
  assay_provider_company: null,
  funding_organizations: null,
  evidence_quote: null,
  confidence: null,
  source_ref: null
)
```

