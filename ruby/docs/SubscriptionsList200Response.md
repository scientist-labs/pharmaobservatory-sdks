# PharmaObservatory::SubscriptionsList200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **subscriptions** | [**Array&lt;SubscriptionsList200ResponseSubscriptionsInner&gt;**](SubscriptionsList200ResponseSubscriptionsInner.md) |  |  |
| **count** | **Integer** |  |  |
| **build_state** | [**ChangesList200ResponseBuildState**](ChangesList200ResponseBuildState.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SubscriptionsList200Response.new(
  subscriptions: null,
  count: null,
  build_state: null
)
```

