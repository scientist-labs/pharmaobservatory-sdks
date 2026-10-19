# PharmaObservatory::BiomarkerEvidenceRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **biomarker_evidence_id** | **String** | The evidence assertion id. | [optional] |
| **route_key** | **String** | The biomarker route key of the gene this evidence is about (the gene symbol); the entry key for the biomarkers object type. | [optional] |
| **gene_symbol** | **String** | The gene symbol the evidence is about. | [optional] |
| **target_id** | **String** | The target&#39;s Ensembl gene id. | [optional] |
| **uniprot_accession** | **String** | The UniProt accession. | [optional] |
| **biomarker** | **String** | The biomarker as the source names it (the gene or the variant). | [optional] |
| **evidence_type** | **String** | cancer_actionable_cgi, ot_association or trial_mention. | [optional] |
| **source** | **String** | The source the assertion comes from. | [optional] |
| **condition_id** | **String** | The condition key. | [optional] |
| **condition_name** | **String** | The condition as the source names it. | [optional] |
| **therapeutic_area** | **String** | The therapeutic area. | [optional] |
| **drug_name** | **String** | The drug as filed. | [optional] |
| **drug_id** | **String** | The drug key. | [optional] |
| **molecule_id** | **String** | The molecule key. Null on every row: no evidence row carries a molecule. | [optional] |
| **response** | **String** | The reported response (for example Responsive, Resistant). | [optional] |
| **score** | **Float** | The association score, for an Open Targets row. | [optional] |
| **variant_count** | **Integer** | Variants the assertion covers. | [optional] |
| **trial_id** | **String** | The trial registry id, for a trial-mention row. | [optional] |
| **sample_type** | **String** | The sample type, for a trial-mention row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerEvidenceRow.new(
  biomarker_evidence_id: null,
  route_key: null,
  gene_symbol: null,
  target_id: null,
  uniprot_accession: null,
  biomarker: null,
  evidence_type: null,
  source: null,
  condition_id: null,
  condition_name: null,
  therapeutic_area: null,
  drug_name: null,
  drug_id: null,
  molecule_id: null,
  response: null,
  score: null,
  variant_count: null,
  trial_id: null,
  sample_type: null
)
```

