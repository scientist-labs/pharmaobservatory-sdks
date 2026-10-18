# PharmaObservatory::OrgsFunding200ResponseRowsInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **event_date** | **String** |  |  |
| **round_type** | **String** |  |  |
| **amount_usd** | **Float** |  |  |
| **amount_is_disclosed** | **Boolean** |  |  |
| **source** | **String** |  |  |
| **funding_event_id** | **String** |  | [optional] |
| **round_subtype** | **String** |  | [optional] |
| **is_dilutive** | **Boolean** |  | [optional] |
| **currency** | **String** |  | [optional] |
| **capital_treatment** | **String** |  | [optional] |
| **cumulative_raised** | **Float** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsFunding200ResponseRowsInner.new(
  event_date: null,
  round_type: null,
  amount_usd: null,
  amount_is_disclosed: null,
  source: null,
  funding_event_id: null,
  round_subtype: null,
  is_dilutive: null,
  currency: null,
  capital_treatment: null,
  cumulative_raised: null
)
```

