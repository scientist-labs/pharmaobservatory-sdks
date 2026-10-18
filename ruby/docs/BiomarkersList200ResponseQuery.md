# PharmaObservatory::BiomarkersList200ResponseQuery

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **q** | **String** |  |  |
| **filters** | [**BiomarkersList200ResponseQueryFilters**](BiomarkersList200ResponseQueryFilters.md) |  |  |
| **sort** | **String** |  |  |
| **dir** | **String** |  |  |
| **evidence_floor** | **Boolean** |  |  |
| **validated_floor** | **Boolean** |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkersList200ResponseQuery.new(
  q: null,
  filters: null,
  sort: null,
  dir: null,
  evidence_floor: null,
  validated_floor: null
)
```

