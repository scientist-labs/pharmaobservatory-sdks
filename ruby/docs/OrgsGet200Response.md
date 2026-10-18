# PharmaObservatory::OrgsGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **page** | **String** |  |  |
| **id** | **String** |  |  |
| **identity** | [**OrgsGet200ResponseIdentity**](OrgsGet200ResponseIdentity.md) |  |  |
| **sections** | [**Array&lt;OrgsGet200ResponseSectionsInner&gt;**](OrgsGet200ResponseSectionsInner.md) |  |  |
| **attribution** | [**Array&lt;BiomarkersList200ResponseAttributionInner&gt;**](BiomarkersList200ResponseAttributionInner.md) |  |  |
| **web_url** | **String** |  |  |
| **build_state** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  |  |
| **resolution_confidence** | [**BiomarkersGet200ResponseResolutionConfidence**](BiomarkersGet200ResponseResolutionConfidence.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsGet200Response.new(
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

