# PharmaObservatory::CellLinesGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **page** | **String** |  |  |
| **id** | **String** |  |  |
| **identity** | [**CellLinesGet200ResponseIdentity**](CellLinesGet200ResponseIdentity.md) |  |  |
| **sections** | [**Array&lt;BiomarkersGet200ResponseSectionsInner&gt;**](BiomarkersGet200ResponseSectionsInner.md) |  |  |
| **attribution** | [**Array&lt;CellLinesList200ResponseAttributionInner&gt;**](CellLinesList200ResponseAttributionInner.md) |  |  |
| **web_url** | **String** |  |  |
| **build_state** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  |  |
| **resolution_confidence** | [**BiomarkersGet200ResponseResolutionConfidence**](BiomarkersGet200ResponseResolutionConfidence.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CellLinesGet200Response.new(
  page: null,
  id: null,
  identity: null,
  sections: null,
  attribution: null,
  web_url: null,
  build_state: null,
  resolution_confidence: null
)
```

