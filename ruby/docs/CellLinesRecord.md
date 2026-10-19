# PharmaObservatory::CellLinesRecord

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **object_type** | **String** |  |  |
| **id** | **String** |  |  |
| **record** | [**CellLinesRow**](CellLinesRow.md) |  |  |
| **attributes** | **Object** |  | [optional] |
| **datasets** | [**Array&lt;DatasetListing&gt;**](DatasetListing.md) |  |  |
| **attribution** | [**Array&lt;DatasetAttribution&gt;**](DatasetAttribution.md) |  | [optional] |
| **source_credit** | **String** |  | [optional] |
| **freshness** | [**BiomarkerCdxAssertionsResponseFreshness**](BiomarkerCdxAssertionsResponseFreshness.md) |  | [optional] |
| **build** | [**BiomarkerCdxAssertionsResponseBuild**](BiomarkerCdxAssertionsResponseBuild.md) |  | [optional] |
| **build_state** | [**DatasetBuildState**](DatasetBuildState.md) |  | [optional] |
| **restriction** | **Hash&lt;String, Object&gt;** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CellLinesRecord.new(
  object_type: null,
  id: null,
  record: null,
  attributes: null,
  datasets: null,
  attribution: null,
  source_credit: null,
  freshness: null,
  build: null,
  build_state: null,
  restriction: null
)
```

