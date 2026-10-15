# PharmaObservatory::ApiV1TrialsEligibilityGeneralizabilityScanGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **mode** | **String** |  | [optional] |
| **filters** | [**ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseFilters**](ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseFilters.md) |  | [optional] |
| **matched_count** | **Integer** |  | [optional] |
| **returned_count** | **Integer** |  | [optional] |
| **trials** | [**Array&lt;ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner&gt;**](ApiV1TrialsEligibilityGeneralizabilityScanGet200ResponseTrialsInner.md) |  | [optional] |
| **build_state** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  |  |
| **error** | **String** |  | [optional] |
| **message** | **String** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1TrialsEligibilityGeneralizabilityScanGet200Response.new(
  mode: null,
  filters: null,
  matched_count: null,
  returned_count: null,
  trials: null,
  build_state: null,
  error: null,
  message: null
)
```

