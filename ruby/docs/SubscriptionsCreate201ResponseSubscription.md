# PharmaObservatory::SubscriptionsCreate201ResponseSubscription

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **label** | **String** |  |  |
| **target** | [**SubscriptionsCreate201ResponseSubscriptionTarget**](SubscriptionsCreate201ResponseSubscriptionTarget.md) |  |  |
| **event_types** | **Array&lt;String&gt;** |  |  |
| **paused** | **Boolean** |  |  |
| **created_at** | **String** |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SubscriptionsCreate201ResponseSubscription.new(
  id: null,
  label: null,
  target: null,
  event_types: null,
  paused: null,
  created_at: null
)
```

