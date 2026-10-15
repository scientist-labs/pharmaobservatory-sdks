# PharmaObservatory::ApiV1CandidateSitesGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **filters** | [**ApiV1CandidateSitesGet200ResponseFilters**](ApiV1CandidateSitesGet200ResponseFilters.md) |  | [optional] |
| **matched_count** | **Integer** |  | [optional] |
| **saturated_count** | **Integer** |  | [optional] |
| **returned_count** | **Integer** |  | [optional] |
| **sites** | [**Array&lt;ApiV1CandidateSitesGet200ResponseSitesInner&gt;**](ApiV1CandidateSitesGet200ResponseSitesInner.md) |  | [optional] |
| **build_state** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  |  |
| **error** | **String** |  | [optional] |
| **message** | **String** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1CandidateSitesGet200Response.new(
  filters: null,
  matched_count: null,
  saturated_count: null,
  returned_count: null,
  sites: null,
  build_state: null,
  error: null,
  message: null
)
```

