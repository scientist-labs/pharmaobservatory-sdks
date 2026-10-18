# PharmaObservatory::OrgsPdufaReconciliation200Response

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
| **rows** | [**Array&lt;OrgsPdufaReconciliation200ResponseRowsInner&gt;**](OrgsPdufaReconciliation200ResponseRowsInner.md) |  |  |
| **attribution** | [**Array&lt;OrgsBenzingaCatalysts200ResponseAttributionInner&gt;**](OrgsBenzingaCatalysts200ResponseAttributionInner.md) |  |  |
| **context** | [**OrgsPdufaReconciliation200ResponseContext**](OrgsPdufaReconciliation200ResponseContext.md) |  |  |
| **build_state** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  |  |
| **restriction** | [**OrgsBenzingaCatalysts200ResponseRestriction**](OrgsBenzingaCatalysts200ResponseRestriction.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsPdufaReconciliation200Response.new(
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

