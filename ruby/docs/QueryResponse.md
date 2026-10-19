# PharmaObservatory::QueryResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **columns** | [**Array&lt;QueryResponseColumnsInner&gt;**](QueryResponseColumnsInner.md) |  |  |
| **rows** | **Array&lt;Array&lt;Object&gt;&gt;** |  |  |
| **total_rows** | **Integer** |  | [optional] |
| **truncated** | **Boolean** |  | [optional] |
| **result_id** | **String** |  |  |
| **next_cursor** | **String** |  | [optional] |
| **attribution** | [**Array&lt;DatasetAttribution&gt;**](DatasetAttribution.md) |  | [optional] |
| **source_credit** | **String** |  | [optional] |
| **withheld_tiers** | **Array&lt;String&gt;** |  | [optional] |
| **withheld_columns** | **Array&lt;String&gt;** |  | [optional] |
| **build_state** | [**DatasetBuildState**](DatasetBuildState.md) |  | [optional] |
| **dataset** | **String** |  |  |
| **notice** | **String** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::QueryResponse.new(
  columns: null,
  rows: null,
  total_rows: null,
  truncated: null,
  result_id: null,
  next_cursor: null,
  attribution: null,
  source_credit: null,
  withheld_tiers: null,
  withheld_columns: null,
  build_state: null,
  dataset: null,
  notice: null
)
```

