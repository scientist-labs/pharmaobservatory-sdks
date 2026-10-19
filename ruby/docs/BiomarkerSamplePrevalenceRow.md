# PharmaObservatory::BiomarkerSamplePrevalenceRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **route_key** | **String** | The route key of the gene the variants belong to, the entry key for the biomarkers object type. | [optional] |
| **gene_symbol** | **String** | The gene symbol. | [optional] |
| **variant_biomarker_id** | **String** | The variant&#39;s own biomarker key. | [optional] |
| **variant_route_key** | **String** | The variant&#39;s own route key (the id its web page uses). | [optional] |
| **amino_acid_change** | **String** | The protein change, for example p.V600E. | [optional] |
| **gdc_ssm_id** | **String** | The identifier of the somatic mutation in the Genomic Data Commons. | [optional] |
| **gdc_project_id** | **String** | The Genomic Data Commons research project (cohort) the counts come from. | [optional] |
| **gdc_project_name** | **String** | The name of the research project. | [optional] |
| **condition_id** | **String** | The condition key (a MeSH, Medical Subject Headings, descriptor id) the project&#39;s disease maps to, where one is mapped. | [optional] |
| **condition_name** | **String** | The mapped condition&#39;s MeSH heading. | [optional] |
| **source_disease_types** | **Array&lt;String&gt;** | The disease types the source states for the project. | [optional] |
| **source_primary_sites** | **Array&lt;String&gt;** | The primary sites the source states for the project. | [optional] |
| **affected_case_count** | **Integer** | Cases in the cohort with the exact variant. | [optional] |
| **ssm_profiled_case_count** | **Integer** | Cases in the cohort profiled for somatic mutations, the denominator of the observed frequency. | [optional] |
| **all_case_count** | **Integer** | All cases in the project. | [optional] |
| **observed_frequency** | **Float** | Affected cases over mutation-profiled cases in this cohort. | [optional] |
| **frequency_ci95_low** | **Float** | Lower bound of the 95 percent confidence interval of the observed frequency. | [optional] |
| **frequency_ci95_high** | **Float** | Upper bound of the 95 percent confidence interval of the observed frequency. | [optional] |
| **gdc_data_release** | **String** | The Genomic Data Commons data release the counts come from. | [optional] |
| **claim_type** | **String** | What the figure states; an observed frequency in a named cohort. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerSamplePrevalenceRow.new(
  route_key: null,
  gene_symbol: null,
  variant_biomarker_id: null,
  variant_route_key: null,
  amino_acid_change: null,
  gdc_ssm_id: null,
  gdc_project_id: null,
  gdc_project_name: null,
  condition_id: null,
  condition_name: null,
  source_disease_types: null,
  source_primary_sites: null,
  affected_case_count: null,
  ssm_profiled_case_count: null,
  all_case_count: null,
  observed_frequency: null,
  frequency_ci95_low: null,
  frequency_ci95_high: null,
  gdc_data_release: null,
  claim_type: null
)
```

