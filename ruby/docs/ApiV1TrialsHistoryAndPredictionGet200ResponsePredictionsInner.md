# PharmaObservatory::ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **kind** | **String** |  |  |
| **registry_value** | **Boolean** |  |  |
| **marker** | **String** |  |  |
| **estimate_form** | **String** |  |  |
| **state** | **String** |  |  |
| **method** | **String** |  |  |
| **model_name** | **String** |  |  |
| **model_version** | **String** |  |  |
| **cutoff** | **String** |  |  |
| **reason** | [**ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerReason**](ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerReason.md) |  |  |
| **statement** | **String** |  |  |
| **figure** | [**ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerFigure**](ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerFigure.md) |  |  |
| **baseline** | [**ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerBaseline**](ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerBaseline.md) |  |  |
| **registered_primary_completion** | [**ApiV1TrialsHistoryPredictionsInnerRegisteredPrimaryCompletion**](ApiV1TrialsHistoryPredictionsInnerRegisteredPrimaryCompletion.md) |  |  |
| **measured_error** | [**ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredError**](ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredError.md) |  |  |
| **inputs_as_registered** | [**ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerInputsAsRegistered**](ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerInputsAsRegistered.md) |  |  |
| **cutoff_statement** | **String** |  |  |
| **history_runs_past_cutoff** | **Boolean** |  |  |
| **newest_history_as_of** | **String** |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner.new(
  kind: null,
  registry_value: null,
  marker: null,
  estimate_form: null,
  state: null,
  method: null,
  model_name: null,
  model_version: null,
  cutoff: null,
  reason: null,
  statement: null,
  figure: null,
  baseline: null,
  registered_primary_completion: null,
  measured_error: null,
  inputs_as_registered: null,
  cutoff_statement: null,
  history_runs_past_cutoff: null,
  newest_history_as_of: null
)
```

