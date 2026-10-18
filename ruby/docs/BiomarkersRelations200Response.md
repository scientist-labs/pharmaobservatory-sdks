# PharmaObservatory::BiomarkersRelations200Response

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
| **rows** | [**Array&lt;BiomarkersRelations200ResponseRowsInner&gt;**](BiomarkersRelations200ResponseRowsInner.md) |  |  |
| **attribution** | [**Array&lt;BiomarkersAssayProviderLedger200ResponseAttributionInner&gt;**](BiomarkersAssayProviderLedger200ResponseAttributionInner.md) |  |  |
| **pagination** | [**BiomarkersAssayProviderLedger200ResponsePagination**](BiomarkersAssayProviderLedger200ResponsePagination.md) |  |  |
| **query** | [**BiomarkersRelations200ResponseQuery**](BiomarkersRelations200ResponseQuery.md) |  |  |
| **build_state** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  |  |
| **context** | [**BiomarkersRelations200ResponseContext**](BiomarkersRelations200ResponseContext.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkersRelations200Response.new(
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
  context: null
)
```

