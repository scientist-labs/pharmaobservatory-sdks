# PharmaObservatory::ApiV1TrialsDiscontinuationReasonsGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **error** | **String** |  | [optional] |
| **message** | **String** |  | [optional] |
| **build_state** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  |  |
| **filters** | [**ApiV1TrialsDiscontinuationReasonsGet200ResponseFilters**](ApiV1TrialsDiscontinuationReasonsGet200ResponseFilters.md) |  | [optional] |
| **matched_count** | **Integer** |  | [optional] |
| **trials_with_why_stopped** | **Integer** |  | [optional] |
| **why_stopped_coverage_rate** | **Float** |  | [optional] |
| **coverage_caveat** | **String** |  | [optional] |
| **trials** | [**Array&lt;ApiV1TrialsDiscontinuationReasonsGet200ResponseTrialsInner&gt;**](ApiV1TrialsDiscontinuationReasonsGet200ResponseTrialsInner.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1TrialsDiscontinuationReasonsGet200Response.new(
  error: null,
  message: null,
  build_state: null,
  filters: null,
  matched_count: null,
  trials_with_why_stopped: null,
  why_stopped_coverage_rate: null,
  coverage_caveat: null,
  trials: null
)
```

