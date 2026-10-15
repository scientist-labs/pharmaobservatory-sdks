# PharmaObservatory::ApiV1TrialsReportingCompletenessGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **filters** | [**ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseFilters**](ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseFilters.md) |  | [optional] |
| **matched_count** | **Integer** |  | [optional] |
| **reported_count** | **Integer** |  | [optional] |
| **reported_share** | **Float** |  | [optional] |
| **demographic_field_population_rate** | [**ApiV1TrialsReportingCompletenessGet200ResponseDemographicFieldPopulationRate**](ApiV1TrialsReportingCompletenessGet200ResponseDemographicFieldPopulationRate.md) |  | [optional] |
| **build_state** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  |  |
| **error** | **String** |  | [optional] |
| **message** | **String** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1TrialsReportingCompletenessGet200Response.new(
  filters: null,
  matched_count: null,
  reported_count: null,
  reported_share: null,
  demographic_field_population_rate: null,
  build_state: null,
  error: null,
  message: null
)
```

