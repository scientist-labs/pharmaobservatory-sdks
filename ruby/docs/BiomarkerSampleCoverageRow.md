# PharmaObservatory::BiomarkerSampleCoverageRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **route_key** | **String** | The route key of the gene the variants belong to, the entry key for the biomarkers object type. | [optional] |
| **gene_symbol** | **String** | The gene symbol. | [optional] |
| **variant_biomarker_id** | **String** | The variant&#39;s own biomarker key. | [optional] |
| **variant_route_key** | **String** | The variant&#39;s own route key (the id its web page uses). | [optional] |
| **amino_acid_change** | **String** | The protein change, for example p.V600E. | [optional] |
| **gdc_data_release** | **String** | The Genomic Data Commons data release the match was made against. | [optional] |
| **normalization_status** | **String** | How the variant matched: resolved_single_ssm, no_exact_ssm, multiple_exact_ssms or source_error. | [optional] |
| **match_count** | **Integer** | Somatic mutations in the catalogue that encode the same protein change. | [optional] |
| **resolution_reason** | **String** | The reason for the status, in words. | [optional] |
| **measured_project_count** | **Integer** | Research projects with an observed frequency for the variant in biomarker_sample_prevalence. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerSampleCoverageRow.new(
  route_key: null,
  gene_symbol: null,
  variant_biomarker_id: null,
  variant_route_key: null,
  amino_acid_change: null,
  gdc_data_release: null,
  normalization_status: null,
  match_count: null,
  resolution_reason: null,
  measured_project_count: null
)
```

