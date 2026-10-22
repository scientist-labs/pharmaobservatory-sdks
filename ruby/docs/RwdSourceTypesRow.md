# PharmaObservatory::RwdSourceTypesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **data_source_id** | **String** | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study&#39;s data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional] |
| **item_ordinal** | **Integer** | The position of the item in the register&#39;s own list, from 1. | [optional] |
| **type_value** | **String** | The declared data source type, as typed after the register&#39;s list is split. | [optional] |
| **type_key** | **String** | A stable lowercase key for the data source type; the filter value. | [optional] |
| **type_label** | **String** | The register&#39;s most common spelling of the type. | [optional] |
| **is_other_sentinel** | **Boolean** | True when the item is the register&#39;s Other sentinel; the holder&#39;s free text for it is on the source row. | [optional] |
| **source_value** | **String** | The whole register answer the item was split from. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdSourceTypesRow.new(
  data_source_id: null,
  item_ordinal: null,
  type_value: null,
  type_key: null,
  type_label: null,
  is_other_sentinel: null,
  source_value: null,
  restriction_tier: null,
  attribution: null
)
```

