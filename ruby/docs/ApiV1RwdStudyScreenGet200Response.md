# PharmaObservatory::ApiV1RwdStudyScreenGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **filters** | [**ApiV1RwdStudyScreenGet200ResponseFilters**](ApiV1RwdStudyScreenGet200ResponseFilters.md) |  |  |
| **matched_studies** | **Integer** |  |  |
| **register_total** | **Integer** |  |  |
| **delivery_distribution** | [**ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution**](ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution.md) |  |  |
| **coverage_caveat** | **String** |  |  |
| **tier_provisional_note** | **String** |  |  |
| **build_state** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  |  |
| **restriction** | [**ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction**](ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction.md) |  |  |
| **condition_filter_built** | **Boolean** | false means a condition was named and the condition mart is not built in this warehouse, so the filter could not run and matched_studies, register_total and delivery_distribution are unknown (null), not zero. true means the filter ran, or no condition was named. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1RwdStudyScreenGet200Response.new(
  filters: null,
  matched_studies: null,
  register_total: null,
  delivery_distribution: null,
  coverage_caveat: null,
  tier_provisional_note: null,
  build_state: null,
  restriction: null,
  condition_filter_built: null
)
```

