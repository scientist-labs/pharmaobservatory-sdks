# PharmaObservatory::ApiV1TrialsRecordsGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **requested** | **Array&lt;String&gt;** |  |  |
| **trials** | [**Array&lt;ApiV1TrialsRecordsGet200ResponseTrialsInner&gt;**](ApiV1TrialsRecordsGet200ResponseTrialsInner.md) |  |  |
| **not_found** | **Array&lt;String&gt;** |  |  |
| **extraction_coverage** | [**ApiV1TrialsRecordsGet200ResponseExtractionCoverage**](ApiV1TrialsRecordsGet200ResponseExtractionCoverage.md) |  | [optional] |
| **build_state** | [**ApiV1TrialsRecordsGet200ResponseBuildState**](ApiV1TrialsRecordsGet200ResponseBuildState.md) |  |  |
| **sections** | [**ApiV1TrialsRecordsGet200ResponseSections**](ApiV1TrialsRecordsGet200ResponseSections.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1TrialsRecordsGet200Response.new(
  requested: null,
  trials: null,
  not_found: null,
  extraction_coverage: null,
  build_state: null,
  sections: null
)
```

