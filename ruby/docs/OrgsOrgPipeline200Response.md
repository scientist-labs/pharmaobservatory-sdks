# PharmaObservatory::OrgsOrgPipeline200Response

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
| **rows** | [**Array&lt;OrgsOrgPipeline200ResponseRowsInner&gt;**](OrgsOrgPipeline200ResponseRowsInner.md) |  |  |
| **attribution** | [**Array&lt;CellLinesList200ResponseAttributionInner&gt;**](CellLinesList200ResponseAttributionInner.md) |  |  |
| **context** | [**OrgsOrgPipeline200ResponseContext**](OrgsOrgPipeline200ResponseContext.md) |  |  |
| **build_state** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  |  |
| **restriction** | [**OrgsBenzingaCatalysts200ResponseRestriction**](OrgsBenzingaCatalysts200ResponseRestriction.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsOrgPipeline200Response.new(
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
  build_state: null,
  restriction: null
)
```

