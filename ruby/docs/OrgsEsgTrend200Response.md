# PharmaObservatory::OrgsEsgTrend200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **page** | **String** |  |  |
| **id** | **String** |  |  |
| **section** | **String** |  |  |
| **heading** | **String** |  |  |
| **lens** | **String** |  |  |
| **web_url** | **String** |  |  |
| **shape** | [**BiomarkersList200ResponseShape**](BiomarkersList200ResponseShape.md) |  |  |
| **columns** | [**Array&lt;BiomarkersList200ResponseColumnsInner&gt;**](BiomarkersList200ResponseColumnsInner.md) |  |  |
| **rows** | [**Array&lt;OrgsEsgTrend200ResponseRowsInner&gt;**](OrgsEsgTrend200ResponseRowsInner.md) |  |  |
| **attribution** | [**Array&lt;OrgsEsgTrend200ResponseAttributionInner&gt;**](OrgsEsgTrend200ResponseAttributionInner.md) |  |  |
| **context** | [**OrgsEsgTrend200ResponseContext**](OrgsEsgTrend200ResponseContext.md) |  |  |
| **build_state** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsEsgTrend200Response.new(
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
  context: null,
  build_state: null
)
```

