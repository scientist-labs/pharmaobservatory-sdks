# PharmaObservatory::OrgsOrgOwnership200ResponseContext

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **view** | **String** |  |  |
| **limit** | **Integer** |  |  |
| **totals** | [**OrgsCurrentOwnership200ResponseContextTotals**](OrgsCurrentOwnership200ResponseContextTotals.md) |  |  |
| **truncated** | [**OrgsOrgOwnership200ResponseContextTruncated**](OrgsOrgOwnership200ResponseContextTruncated.md) |  |  |
| **personal_data_owners** | **Boolean** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsOrgOwnership200ResponseContext.new(
  view: null,
  limit: null,
  totals: null,
  truncated: null,
  personal_data_owners: null
)
```

