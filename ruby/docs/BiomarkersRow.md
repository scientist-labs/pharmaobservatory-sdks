# PharmaObservatory::BiomarkersRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **biomarker_id** | **String** | The biomarker key. | [optional] |
| **route_key** | **String** | The biomarker&#39;s route key, the id the web page uses and the one entry key every biomarker dataset takes (the gene symbol for a gene biomarker). | [optional] |
| **biomarker_kind** | **String** | gene, protein_variant or another feature kind. | [optional] |
| **identity_status** | **String** | Whether the biomarker&#39;s identity is confirmed. | [optional] |
| **display_name** | **String** | The biomarker&#39;s display name. | [optional] |
| **parent_biomarker_id** | **String** | The parent biomarker (a variant&#39;s gene). | [optional] |
| **gene_symbol** | **String** | The gene symbol. | [optional] |
| **amino_acid_change** | **String** | The amino-acid change, for a variant. | [optional] |
| **target_id** | **String** | The target&#39;s Ensembl gene id. | [optional] |
| **hgnc_id** | **String** | The HGNC gene id. | [optional] |
| **uniprot_accession** | **String** | The UniProt accession. | [optional] |
| **feature_code** | **String** | Feature code for a non-gene feature. | [optional] |
| **feature_type** | **String** | Feature type for a non-gene feature. | [optional] |
| **evidence_tier** | **Integer** | Strongest evidence tier: 1 FDA companion diagnostic through 6 no evidence. | [optional] |
| **evidence_tier_label** | **String** | The tier&#39;s name. | [optional] |
| **n_fda_cdx_assertions** | **Integer** | FDA companion diagnostic assertions. | [optional] |
| **n_fda_bqp_assertions** | **Integer** | FDA biomarker qualification project assertions. | [optional] |
| **n_cancer_actionable_cgi** | **Integer** | Cancer Genome Interpreter actionable rows. | [optional] |
| **n_ot_association** | **Integer** | Open Targets target-disease associations. | [optional] |
| **n_trial_mention** | **Integer** | Trials that name the biomarker. | [optional] |
| **n_publication_variant_mentions** | **Integer** | Publication variant mentions. | [optional] |
| **has_clinical_validation_evidence** | **Boolean** | Whether clinical validation evidence exists. | [optional] |
| **has_analytical_validation_evidence** | **Boolean** | Whether analytical validation evidence exists. | [optional] |
| **n_assay_provider_papers** | **Integer** | Papers naming an assay provider for the biomarker. | [optional] |
| **n_assay_provider_suppliers** | **Integer** | Assay-provider suppliers named. | [optional] |
| **n_diagnostic_provider_suppliers** | **Integer** | Of those, diagnostic providers. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkersRow.new(
  biomarker_id: null,
  route_key: null,
  biomarker_kind: null,
  identity_status: null,
  display_name: null,
  parent_biomarker_id: null,
  gene_symbol: null,
  amino_acid_change: null,
  target_id: null,
  hgnc_id: null,
  uniprot_accession: null,
  feature_code: null,
  feature_type: null,
  evidence_tier: null,
  evidence_tier_label: null,
  n_fda_cdx_assertions: null,
  n_fda_bqp_assertions: null,
  n_cancer_actionable_cgi: null,
  n_ot_association: null,
  n_trial_mention: null,
  n_publication_variant_mentions: null,
  has_clinical_validation_evidence: null,
  has_analytical_validation_evidence: null,
  n_assay_provider_papers: null,
  n_assay_provider_suppliers: null,
  n_diagnostic_provider_suppliers: null
)
```

