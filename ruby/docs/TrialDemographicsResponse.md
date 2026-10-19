# PharmaObservatory::TrialDemographicsResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **dataset** | **String** |  |  |
| **grain** | **String** |  | [optional] |
| **object_type** | **String** |  | [optional] |
| **id** | **String** |  | [optional] |
| **scope** | [**BiomarkerCdxAssertionsResponseScope**](BiomarkerCdxAssertionsResponseScope.md) |  | [optional] |
| **columns** | [**Array&lt;DatasetColumn&gt;**](DatasetColumn.md) |  |  |
| **rows** | [**Array&lt;TrialDemographicsRow&gt;**](TrialDemographicsRow.md) |  |  |
| **total** | **Integer** | The exact count of the filtered rows on an ungrouped read, carried by the cursor; null when grouped. | [optional] |
| **next_cursor** | **String** |  | [optional] |
| **sort** | **Array&lt;String&gt;** |  | [optional] |
| **filters** | **Array&lt;Hash&lt;String, Object&gt;&gt;** |  | [optional] |
| **group_by** | **Array&lt;String&gt;** |  | [optional] |
| **measures** | **Array&lt;String&gt;** |  | [optional] |
| **attribution** | [**Array&lt;DatasetAttribution&gt;**](DatasetAttribution.md) |  | [optional] |
| **source_credit** | **String** |  | [optional] |
| **withheld_tiers** | **Array&lt;String&gt;** |  | [optional] |
| **withheld_columns** | **Array&lt;String&gt;** |  | [optional] |
| **freshness** | [**BiomarkerCdxAssertionsResponseFreshness**](BiomarkerCdxAssertionsResponseFreshness.md) |  | [optional] |
| **build** | [**BiomarkerCdxAssertionsResponseBuild**](BiomarkerCdxAssertionsResponseBuild.md) |  | [optional] |
| **build_state** | [**DatasetBuildState**](DatasetBuildState.md) |  | [optional] |
| **restriction** | **Hash&lt;String, Object&gt;** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialDemographicsResponse.new(
  dataset: null,
  grain: null,
  object_type: null,
  id: null,
  scope: null,
  columns: null,
  rows: null,
  total: null,
  next_cursor: null,
  sort: null,
  filters: null,
  group_by: null,
  measures: null,
  attribution: null,
  source_credit: null,
  withheld_tiers: null,
  withheld_columns: null,
  freshness: null,
  build: null,
  build_state: null,
  restriction: null
)
```

