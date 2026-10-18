# PharmaObservatory::OrgsPeers200Response

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
| **rows** | [**Array&lt;OrgsPeers200ResponseRowsInner&gt;**](OrgsPeers200ResponseRowsInner.md) |  |  |
| **attribution** | **Array&lt;Object&gt;** |  |  |
| **build_state** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsPeers200Response.new(
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
  build_state: null
)
```

