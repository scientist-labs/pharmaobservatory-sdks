# PharmaObservatory::ConditionBurdenRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **condition_name** | **String** | The condition&#39;s MeSH heading. | [optional] |
| **cause_name** | **String** | The Global Burden of Disease cause matched to the condition. | [optional] |
| **dalys_millions** | **Float** | Global disability-adjusted life years for the matched cause, in millions. | [optional] |
| **deaths_millions** | **Float** | Global deaths for the matched cause, in millions. | [optional] |
| **druggability** | **Integer** | Druggability class of the matched cause: 1 highly druggable, 2 druggable, 3 challenging, 4 very challenging. | [optional] |
| **n_trials** | **Integer** | Trials for the condition, counted under each condition they are tagged with and each condition above it in the MeSH tree. | [optional] |
| **trials_per_million_dalys** | **Float** | Trials per million DALYs: n_trials divided by dalys_millions. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ConditionBurdenRow.new(
  condition_id: null,
  condition_name: null,
  cause_name: null,
  dalys_millions: null,
  deaths_millions: null,
  druggability: null,
  n_trials: null,
  trials_per_million_dalys: null
)
```

