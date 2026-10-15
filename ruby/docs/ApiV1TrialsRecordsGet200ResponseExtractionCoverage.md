# PharmaObservatory::ApiV1TrialsRecordsGet200ResponseExtractionCoverage

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **available** | **Boolean** |  |  |
| **trials_swept** | **Integer** |  | [optional] |
| **trials_with_result_posted** | **Integer** |  | [optional] |
| **trials_with_clean_read** | **Integer** |  | [optional] |
| **result_posted_rate** | **Float** |  | [optional] |
| **clean_read_rate_of_posted** | **Float** |  | [optional] |
| **caveat** | **String** |  | [optional] |
| **reason** | **String** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1TrialsRecordsGet200ResponseExtractionCoverage.new(
  available: null,
  trials_swept: null,
  trials_with_result_posted: null,
  trials_with_clean_read: null,
  result_posted_rate: null,
  clean_read_rate_of_posted: null,
  caveat: null,
  reason: null
)
```

