# PharmaObservatory::ChangesList200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **items** | [**Array&lt;ChangesList200ResponseItemsInner&gt;**](ChangesList200ResponseItemsInner.md) |  |  |
| **count** | **Integer** |  |  |
| **next_cursor** | **String** |  |  |
| **has_more** | **Boolean** |  |  |
| **build_state** | [**ChangesList200ResponseBuildState**](ChangesList200ResponseBuildState.md) |  |  |
| **oldest_available** | **String** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ChangesList200Response.new(
  items: null,
  count: null,
  next_cursor: null,
  has_more: null,
  build_state: null,
  oldest_available: null
)
```

