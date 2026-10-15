# PharmaObservatory::ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **available** | **Boolean** |  |  |
| **sites_available** | **Boolean** |  |  |
| **coverage_available** | **Boolean** |  |  |
| **snapshot_count** | **Integer** |  |  |
| **first_as_of** | **String** |  |  |
| **last_as_of** | **String** |  |  |
| **window** | **String** |  |  |
| **fully_captured** | **Boolean** |  |  |
| **observations** | [**Array&lt;ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner&gt;**](ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner.md) |  |  |
| **changes** | [**Array&lt;ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryChangesInner&gt;**](ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryChangesInner.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory.new(
  available: null,
  sites_available: null,
  coverage_available: null,
  snapshot_count: null,
  first_as_of: null,
  last_as_of: null,
  window: null,
  fully_captured: null,
  observations: null,
  changes: null
)
```

