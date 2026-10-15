# PharmaObservatory::ApiV1EnrollmentBenchmarkGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **error** | **String** |  | [optional] |
| **message** | **String** |  | [optional] |
| **build_state** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  |  |
| **filters** | [**ApiV1EnrollmentBenchmarkGet200ResponseFilters**](ApiV1EnrollmentBenchmarkGet200ResponseFilters.md) |  | [optional] |
| **matched** | **Boolean** |  | [optional] |
| **benchmark** | [**ApiV1EnrollmentBenchmarkGet200ResponseBenchmark**](ApiV1EnrollmentBenchmarkGet200ResponseBenchmark.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1EnrollmentBenchmarkGet200Response.new(
  error: null,
  message: null,
  build_state: null,
  filters: null,
  matched: null,
  benchmark: null
)
```

