# PharmaObservatory::ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **kind** | **String** |  |  |
| **as_of** | **String** |  |  |
| **enrollment** | [**ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerEnrollment**](ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerEnrollment.md) |  |  |
| **start_date** | [**ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate**](ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate.md) |  |  |
| **primary_completion_date** | [**ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate**](ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate.md) |  |  |
| **completion_date** | [**ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate**](ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerStartDate.md) |  |  |
| **overall_status** | [**ApiV1TrialsHistoryRecordHistoryObservationsInnerOverallStatus**](ApiV1TrialsHistoryRecordHistoryObservationsInnerOverallStatus.md) |  |  |
| **sites** | [**ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerSites**](ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInnerSites.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistoryObservationsInner.new(
  kind: null,
  as_of: null,
  enrollment: null,
  start_date: null,
  primary_completion_date: null,
  completion_date: null,
  overall_status: null,
  sites: null
)
```

