# PharmaObservatory::CellLinesIdentity200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **page** | **String** |  |  |
| **id** | **String** |  |  |
| **section** | **String** |  |  |
| **heading** | **String** |  |  |
| **lens** | **Object** |  |  |
| **web_url** | **String** |  |  |
| **shape** | [**BiomarkersList200ResponseShape**](BiomarkersList200ResponseShape.md) |  |  |
| **columns** | [**Array&lt;BiomarkersList200ResponseColumnsInner&gt;**](BiomarkersList200ResponseColumnsInner.md) |  |  |
| **rows** | [**Array&lt;CellLinesGet200ResponseIdentityRowsInner&gt;**](CellLinesGet200ResponseIdentityRowsInner.md) |  |  |
| **attribution** | [**Array&lt;CellLinesList200ResponseAttributionInner&gt;**](CellLinesList200ResponseAttributionInner.md) |  |  |
| **build_state** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  |  |
| **resolution_confidence** | [**BiomarkersGet200ResponseResolutionConfidence**](BiomarkersGet200ResponseResolutionConfidence.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CellLinesIdentity200Response.new(
  page: null,
  id: null,
  section: null,
  heading: null,
  lens: null,
  web_url: null,
  shape: null,
  columns: null,
  rows: null,
  attribution: null,
  build_state: null,
  resolution_confidence: null
)
```

