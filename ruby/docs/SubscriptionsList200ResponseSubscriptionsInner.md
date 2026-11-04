# PharmaObservatory::SubscriptionsList200ResponseSubscriptionsInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **label** | **String** |  |  |
| **target** | [**SubscriptionsList200ResponseSubscriptionsInnerTarget**](SubscriptionsList200ResponseSubscriptionsInnerTarget.md) |  |  |
| **event_types** | **Array&lt;String&gt;** |  |  |
| **paused** | **Boolean** |  |  |
| **created_at** | **String** |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SubscriptionsList200ResponseSubscriptionsInner.new(
  id: null,
  label: null,
  target: null,
  event_types: null,
  paused: null,
  created_at: null
)
```

