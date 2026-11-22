# PharmaObservatory::ConditionSiteCoverageRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **start_year** | **Integer** | Year the trials started. Null for a trial with no start date. | [optional] |
| **n_trials** | **Integer** | Trials of the condition that started that year. | [optional] |
| **n_trials_located** | **Integer** | Of those trials, how many list at least one facility with a coordinate. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ConditionSiteCoverageRow.new(
  condition_id: null,
  start_year: null,
  n_trials: null,
  n_trials_located: null
)
```

