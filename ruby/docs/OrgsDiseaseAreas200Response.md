# PharmaObservatory::OrgsDiseaseAreas200Response

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
| **rows** | [**Array&lt;OrgsDiseaseAreas200ResponseRowsInner&gt;**](OrgsDiseaseAreas200ResponseRowsInner.md) |  |  |
| **attribution** | [**Array&lt;BiomarkersCgi200ResponseAttributionInner&gt;**](BiomarkersCgi200ResponseAttributionInner.md) |  |  |
| **pagination** | [**BiomarkersAssayProviderLedger200ResponsePagination**](BiomarkersAssayProviderLedger200ResponsePagination.md) |  |  |
| **context** | [**OrgsApprovedAssets200ResponseContext**](OrgsApprovedAssets200ResponseContext.md) |  |  |
| **build_state** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsDiseaseAreas200Response.new(
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
  context: null,
  build_state: null
)
```

