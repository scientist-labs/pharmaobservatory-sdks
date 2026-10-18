# PharmaObservatory::OrgsOrgRecruitment200ResponseContext

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **state** | **String** |  |  |
| **realized_state** | **String** |  |  |
| **kind** | **String** |  |  |
| **cutoff** | **String** |  |  |
| **min_n** | **Integer** |  |  |
| **cap** | **Integer** |  |  |
| **model_version** | **String** |  |  |
| **method** | **String** |  |  |
| **n_trials** | **Integer** |  |  |
| **n_evaluated** | **Integer** |  |  |
| **n_not_evaluated** | **Integer** |  |  |
| **statuses** | [**OrgsOrgRecruitment200ResponseContextStatuses**](OrgsOrgRecruitment200ResponseContextStatuses.md) |  |  |
| **n_superseded** | **Integer** |  |  |
| **n_open** | **Integer** |  |  |
| **n_at_cap** | **Integer** |  |  |
| **n_overdue** | **Integer** |  |  |
| **n_error_measured** | **Integer** |  |  |
| **predicted** | [**OrgsOrgRecruitment200ResponseContextPredicted**](OrgsOrgRecruitment200ResponseContextPredicted.md) |  |  |
| **n_realized** | **Integer** |  |  |
| **n_realized_over_cap** | **Integer** |  |  |
| **realized** | [**OrgsOrgRecruitment200ResponseContextRealized**](OrgsOrgRecruitment200ResponseContextRealized.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsOrgRecruitment200ResponseContext.new(
  state: null,
  realized_state: null,
  kind: null,
  cutoff: null,
  min_n: null,
  cap: null,
  model_version: null,
  method: null,
  n_trials: null,
  n_evaluated: null,
  n_not_evaluated: null,
  statuses: null,
  n_superseded: null,
  n_open: null,
  n_at_cap: null,
  n_overdue: null,
  n_error_measured: null,
  predicted: null,
  n_realized: null,
  n_realized_over_cap: null,
  realized: null
)
```

