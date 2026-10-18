# PharmaObservatory::OrgsFunding200ResponseContext

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **order_note** | **String** |  |  |
| **remainder_noun** | **String** |  |  |
| **undisclosed_events** | **Integer** |  | [optional] |
| **summary** | [**OrgsFunding200ResponseContextSummary**](OrgsFunding200ResponseContextSummary.md) |  | [optional] |
| **downward_corrections** | **Integer** |  | [optional] |
| **downward_corrections_usd** | **Integer** |  | [optional] |
| **unresolved_recognition_records** | **Integer** |  | [optional] |
| **unresolved_usd** | **Integer** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsFunding200ResponseContext.new(
  order_note: null,
  remainder_noun: null,
  undisclosed_events: null,
  summary: null,
  downward_corrections: null,
  downward_corrections_usd: null,
  unresolved_recognition_records: null,
  unresolved_usd: null
)
```

