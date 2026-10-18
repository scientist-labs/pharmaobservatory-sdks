# PharmaObservatory::CellLinesList200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **page** | **String** |  |  |
| **id** | **Object** |  |  |
| **section** | **String** |  |  |
| **heading** | **String** |  |  |
| **lens** | **Object** |  |  |
| **web_url** | **String** |  |  |
| **shape** | [**BiomarkersList200ResponseShape**](BiomarkersList200ResponseShape.md) |  |  |
| **columns** | [**Array&lt;BiomarkersList200ResponseColumnsInner&gt;**](BiomarkersList200ResponseColumnsInner.md) |  |  |
| **rows** | [**Array&lt;CellLinesList200ResponseRowsInner&gt;**](CellLinesList200ResponseRowsInner.md) |  |  |
| **attribution** | [**Array&lt;CellLinesList200ResponseAttributionInner&gt;**](CellLinesList200ResponseAttributionInner.md) |  |  |
| **pagination** | [**BiomarkersList200ResponsePagination**](BiomarkersList200ResponsePagination.md) |  |  |
| **query** | [**CellLinesList200ResponseQuery**](CellLinesList200ResponseQuery.md) |  |  |
| **build_state** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  |  |
| **facets** | [**Array&lt;BiomarkersList200ResponseFacetsInner&gt;**](BiomarkersList200ResponseFacetsInner.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CellLinesList200Response.new(
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
  pagination: null,
  query: null,
  build_state: null,
  facets: null
)
```

