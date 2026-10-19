# PharmaObservatory::TrialOutcomeMeasuresRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **outcome_id** | **Integer** | The outcome measure&#39;s registry id. | [optional] |
| **outcome_type** | **String** | PRIMARY, SECONDARY or OTHER. | [optional] |
| **title** | **String** | The outcome measure&#39;s title. | [optional] |
| **time_frame** | **String** | The time frame the measure covers. | [optional] |
| **population** | **String** | The population the measure is reported for. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialOutcomeMeasuresRow.new(
  trial_id: null,
  outcome_id: null,
  outcome_type: null,
  title: null,
  time_frame: null,
  population: null
)
```

