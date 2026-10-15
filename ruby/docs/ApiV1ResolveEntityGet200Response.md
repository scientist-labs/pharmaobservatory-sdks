# PharmaObservatory::ApiV1ResolveEntityGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **spine** | **String** |  |  |
| **mode** | **String** |  |  |
| **query** | **String** |  |  |
| **source_type** | **String** |  |  |
| **raw_key** | **String** |  |  |
| **candidates** | [**Array&lt;ApiV1ResolveEntityGet200ResponseCandidatesInner&gt;**](ApiV1ResolveEntityGet200ResponseCandidatesInner.md) |  |  |
| **persistence_stable** | **Boolean** |  |  |
| **message** | **String** |  | [optional] |
| **resolution_confidence** | [**ApiV1ResolveEntityGet200ResponseResolutionConfidence**](ApiV1ResolveEntityGet200ResponseResolutionConfidence.md) |  |  |
| **build_state** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1ResolveEntityGet200Response.new(
  spine: null,
  mode: null,
  query: null,
  source_type: null,
  raw_key: null,
  candidates: null,
  persistence_stable: null,
  message: null,
  resolution_confidence: null,
  build_state: null
)
```

