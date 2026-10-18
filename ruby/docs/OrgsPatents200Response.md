# PharmaObservatory::OrgsPatents200Response

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
| **rows** | [**Array&lt;OrgsPatents200ResponseRowsInner&gt;**](OrgsPatents200ResponseRowsInner.md) |  |  |
| **attribution** | [**Array&lt;BiomarkersAssayProviderLedger200ResponseAttributionInner&gt;**](BiomarkersAssayProviderLedger200ResponseAttributionInner.md) |  |  |
| **context** | [**OrgsPatents200ResponseContext**](OrgsPatents200ResponseContext.md) |  |  |
| **build_state** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsPatents200Response.new(
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

