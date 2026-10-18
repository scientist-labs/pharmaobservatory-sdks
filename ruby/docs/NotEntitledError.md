# PharmaObservatory::NotEntitledError

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **error** | **String** |  |  |
| **visible** | **Boolean** |  | [optional] |
| **reason** | **String** |  | [optional] |
| **tier** | **String** |  |  |
| **attribution** | **String** |  |  |
| **page** | **String** |  | [optional] |
| **id** | **String** |  | [optional] |
| **section** | **String** |  | [optional] |
| **heading** | **String** |  | [optional] |
| **lens** | **String** |  | [optional] |
| **web_url** | **String** |  | [optional] |
| **build_state** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  | [optional] |
| **license** | **String** |  | [optional] |
| **restriction** | [**OrgsBenzingaCatalysts200ResponseRestriction**](OrgsBenzingaCatalysts200ResponseRestriction.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::NotEntitledError.new(
  error: null,
  visible: null,
  reason: null,
  tier: null,
  attribution: null,
  page: null,
  id: null,
  section: null,
  heading: null,
  lens: null,
  web_url: null,
  build_state: null,
  license: null,
  restriction: null
)
```

