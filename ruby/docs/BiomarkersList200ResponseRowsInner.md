# PharmaObservatory::BiomarkersList200ResponseRowsInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **route_key** | **String** |  |  |
| **display_name** | **String** |  |  |
| **biomarker_kind** | **String** |  |  |
| **identity_status** | **String** |  |  |
| **gene_symbol** | **String** |  |  |
| **amino_acid_change** | **Object** |  |  |
| **n_fda_cdx_assertions** | **Integer** |  |  |
| **n_fda_bqp_assertions** | **Integer** |  |  |
| **n_cancer_actionable_cgi** | **Integer** |  |  |
| **n_ot_association** | **Integer** |  |  |
| **n_trial_mention** | **Integer** |  |  |
| **n_publication_variant_mentions** | **Integer** |  |  |
| **kind_label** | **String** |  |  |
| **evidence_tier** | **Integer** |  |  |
| **evidence_tier_label** | **String** |  |  |
| **has_clinical_validation_evidence** | **Boolean** |  |  |
| **has_analytical_validation_evidence** | **Boolean** |  |  |
| **n_assay_provider_papers** | **Integer** |  |  |
| **n_assay_provider_suppliers** | **Integer** |  |  |
| **n_diagnostic_provider_suppliers** | **Integer** |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkersList200ResponseRowsInner.new(
  route_key: null,
  display_name: null,
  biomarker_kind: null,
  identity_status: null,
  gene_symbol: null,
  amino_acid_change: null,
  n_fda_cdx_assertions: null,
  n_fda_bqp_assertions: null,
  n_cancer_actionable_cgi: null,
  n_ot_association: null,
  n_trial_mention: null,
  n_publication_variant_mentions: null,
  kind_label: null,
  evidence_tier: null,
  evidence_tier_label: null,
  has_clinical_validation_evidence: null,
  has_analytical_validation_evidence: null,
  n_assay_provider_papers: null,
  n_assay_provider_suppliers: null,
  n_diagnostic_provider_suppliers: null
)
```

