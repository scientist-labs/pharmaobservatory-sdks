# PharmaObservatory::QueryError

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **error_code** | **String** |  |  |
| **message** | **String** |  |  |
| **hint** | **String** |  | [optional] |
| **docs** | **String** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::QueryError.new(
  error_code: null,
  message: null,
  hint: null,
  docs: null
)
```

