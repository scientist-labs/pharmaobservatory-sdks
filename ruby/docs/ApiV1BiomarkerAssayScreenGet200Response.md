# PharmaObservatory::ApiV1BiomarkerAssayScreenGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **filters** | [**ApiV1BiomarkerAssayScreenGet200ResponseFilters**](ApiV1BiomarkerAssayScreenGet200ResponseFilters.md) |  |  |
| **matched_count** | **Integer** |  |  |
| **confidence_caveat** | **String** |  |  |
| **evidence** | [**Array&lt;ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner&gt;**](ApiV1BiomarkerAssayScreenGet200ResponseEvidenceInner.md) |  |  |
| **build_state** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1BiomarkerAssayScreenGet200Response.new(
  filters: null,
  matched_count: null,
  confidence_caveat: null,
  evidence: null,
  build_state: null
)
```

