# PharmaObservatory::ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** |  |  |
| **name** | **String** |  |  |
| **found** | **Boolean** |  |  |
| **dalys_m** | **Float** |  |  |
| **trials** | **Integer** |  |  |
| **trials_per_m_daly** | **String** |  |  |
| **rank** | **Integer** |  |  |
| **total_ranked** | **Integer** |  |  |
| **build_state** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  |  |
| **restriction** | [**ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction**](ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response.new(
  condition_id: null,
  name: null,
  found: null,
  dalys_m: null,
  trials: null,
  trials_per_m_daly: null,
  rank: null,
  total_ranked: null,
  build_state: null,
  restriction: null
)
```

