# PharmaObservatory::ConditionAdverseEventsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **adverse_event_term** | **String** | The adverse event term as the trials report it. | [optional] |
| **subjects_affected** | **Float** | Subjects affected, summed across the condition&#39;s trials. | [optional] |
| **n_trials** | **Integer** | Trials for this condition that report the term. | [optional] |
| **n_terms_total** | **Integer** | Distinct serious adverse event terms across the condition&#39;s trials. The rows are only the top 10 of these, so a row count is not the term count; use this. | [optional] |
| **n_trials_total** | **Integer** | Results-posting trials of the condition that report at least one serious adverse event term. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ConditionAdverseEventsRow.new(
  condition_id: null,
  adverse_event_term: null,
  subjects_affected: null,
  n_trials: null,
  n_terms_total: null,
  n_trials_total: null
)
```

