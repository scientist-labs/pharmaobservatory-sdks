# PharmaObservatory::ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** |  |  |
| **molecule_name** | **String** |  |  |
| **registered_as** | **Array&lt;String&gt;** |  |  |
| **matched_from_name_part** | **Boolean** |  |  |
| **brand_count** | **Integer** |  |  |
| **brands** | [**Array&lt;ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInnerBrandsInner&gt;**](ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInnerBrandsInner.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner.new(
  molecule_id: null,
  molecule_name: null,
  registered_as: null,
  matched_from_name_part: null,
  brand_count: null,
  brands: null
)
```

