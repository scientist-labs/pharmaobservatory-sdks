# PharmaObservatory::ApiV1OrgPipelineStagesGet200ResponseSections

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **in_development** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  |  |
| **abandoned_dormant** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  |  |
| **approved_marketed** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  |  |
| **trial_start_forecast_series** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | [optional] |
| **trial_start_forecast_years** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | [optional] |
| **recruitment_rollup_series** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | [optional] |
| **recruitment_rollup_phase** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1OrgPipelineStagesGet200ResponseSections.new(
  in_development: null,
  abandoned_dormant: null,
  approved_marketed: null,
  trial_start_forecast_series: null,
  trial_start_forecast_years: null,
  recruitment_rollup_series: null,
  recruitment_rollup_phase: null
)
```

