# PharmaObservatory::ConditionsRecordAttributes

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_burden** | [**ConditionBurdenRow**](ConditionBurdenRow.md) |  | [optional] |
| **condition_delivery** | [**ConditionDeliveryRow**](ConditionDeliveryRow.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ConditionsRecordAttributes.new(
  condition_burden: null,
  condition_delivery: null
)
```

