# PharmaObservatory::ApiV1TrialsHistoryAndPredictionGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **nct_id** | **String** |  | [optional] |
| **found** | **Boolean** |  | [optional] |
| **record_history** | [**ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory**](ApiV1TrialsHistoryAndPredictionGet200ResponseRecordHistory.md) |  | [optional] |
| **predictions** | [**Array&lt;ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner&gt;**](ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner.md) |  | [optional] |
| **predictions_status** | [**ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsStatus**](ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsStatus.md) |  | [optional] |
| **build_state** | [**ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState**](ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState.md) |  |  |
| **sections** | [**ApiV1TrialsHistoryAndPredictionGet200ResponseSections**](ApiV1TrialsHistoryAndPredictionGet200ResponseSections.md) |  |  |
| **error** | **String** |  | [optional] |
| **unbuilt_relations** | **Array&lt;String&gt;** |  | [optional] |
| **message** | **String** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1TrialsHistoryAndPredictionGet200Response.new(
  nct_id: null,
  found: null,
  record_history: null,
  predictions: null,
  predictions_status: null,
  build_state: null,
  sections: null,
  error: null,
  unbuilt_relations: null,
  message: null
)
```

