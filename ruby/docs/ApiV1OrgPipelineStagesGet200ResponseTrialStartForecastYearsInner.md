# PharmaObservatory::ApiV1OrgPipelineStagesGet200ResponseTrialStartForecastYearsInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **year** | **Integer** |  |  |
| **offset** | **Integer** |  |  |
| **months_observed** | **Integer** |  |  |
| **observed** | **Integer** |  |  |
| **planned_remaining** | **Integer** |  |  |
| **point** | **Float** |  |  |
| **about** | **Integer** |  |  |
| **horizon** | **String** |  |  |
| **smape_pct** | **String** |  |  |
| **rel_bias** | **String** |  |  |
| **relmae** | **String** |  |  |
| **n_series** | **Integer** |  |  |
| **backtest_kind** | **String** |  |  |
| **method** | **String** |  |  |
| **planned_k** | **Object** |  |  |
| **stricter_agrees** | **Object** |  |  |
| **stricter_bias** | **Object** |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1OrgPipelineStagesGet200ResponseTrialStartForecastYearsInner.new(
  year: null,
  offset: null,
  months_observed: null,
  observed: null,
  planned_remaining: null,
  point: null,
  about: null,
  horizon: null,
  smape_pct: null,
  rel_bias: null,
  relmae: null,
  n_series: null,
  backtest_kind: null,
  method: null,
  planned_k: null,
  stricter_agrees: null,
  stricter_bias: null
)
```

