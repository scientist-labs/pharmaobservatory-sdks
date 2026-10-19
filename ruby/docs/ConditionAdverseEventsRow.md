# PharmaObservatory::ConditionAdverseEventsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **adverse_event_term** | **String** | The adverse event term as the trials report it. | [optional] |
| **subjects_affected** | **Float** | Subjects affected, summed across the condition&#39;s trials. | [optional] |
| **n_trials** | **Integer** | Trials for this condition that report the term. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ConditionAdverseEventsRow.new(
  condition_id: null,
  adverse_event_term: null,
  subjects_affected: null,
  n_trials: null
)
```

