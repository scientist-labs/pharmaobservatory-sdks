# PharmaObservatory::ApiV1TrialsHistoryAndPredictionGet200ResponseSections

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  |  |
| **record_history** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | [optional] |
| **site_history** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | [optional] |
| **history_coverage** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | [optional] |
| **recruitment_prediction** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1TrialsHistoryAndPredictionGet200ResponseSections.new(
  trial: null,
  record_history: null,
  site_history: null,
  history_coverage: null,
  recruitment_prediction: null
)
```

