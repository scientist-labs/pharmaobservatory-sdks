# PharmaObservatory::BiomarkersSampleScreening200ResponseRowsInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **amino_acid_change** | **String** |  |  |
| **gdc_project_id** | **String** |  |  |
| **source_disease_types** | **Array&lt;String&gt;** |  |  |
| **source_primary_sites** | **Array&lt;String&gt;** |  |  |
| **affected_case_count** | **Integer** |  |  |
| **ssm_profiled_case_count** | **Integer** |  |  |
| **observed_frequency** | **Float** |  |  |
| **frequency_ci95_low** | **Float** |  |  |
| **frequency_ci95_high** | **Float** |  |  |
| **gdc_data_release** | **String** |  |  |
| **reference_screens** | [**BiomarkersSampleScreening200ResponseRowsInnerReferenceScreens**](BiomarkersSampleScreening200ResponseRowsInnerReferenceScreens.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkersSampleScreening200ResponseRowsInner.new(
  amino_acid_change: null,
  gdc_project_id: null,
  source_disease_types: null,
  source_primary_sites: null,
  affected_case_count: null,
  ssm_profiled_case_count: null,
  observed_frequency: null,
  frequency_ci95_low: null,
  frequency_ci95_high: null,
  gdc_data_release: null,
  reference_screens: null
)
```

