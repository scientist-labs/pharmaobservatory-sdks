# PharmaObservatory::RwdSourceLinkagesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **data_source_id** | **String** | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study&#39;s data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional] |
| **linkage_ordinal** | **Integer** | The position of the linkage among the source&#39;s declared linkages, from 1. | [optional] |
| **declared_linkage_count** | **Integer** | How many linkages the holder declared; 0 when none. | [optional] |
| **is_pre_linked** | **Boolean** | Whether the holder answered Yes to Pre linked for this linkage. | [optional] |
| **pre_linked_answer** | **String** | The holder&#39;s Yes or No answer to Pre linked for this linkage. | [optional] |
| **linked_data_source_name** | **String** | The linked data source named for this linkage, from the register&#39;s list. | [optional] |
| **linked_data_source_name_aligned** | **Boolean** | False when the register list for this column splits into a different number of pieces than the number of linkages declared; the row value is then null and the whole answer is in linked_data_source_name_source. | [optional] |
| **linked_data_source_name_source** | **String** | The whole comma-joined answer the linked data source name was split from, kept so an unaligned row can still be read. | [optional] |
| **names_other_sentinel** | **Boolean** | True when the linked data source name is the register&#39;s Other data source sentinel. | [optional] |
| **names_registered_source** | **Boolean** | True when the linked data source name is itself a registered data source. | [optional] |
| **linked_data_source_other** | **String** | The linked data source named in free text when the holder chose Other data source. | [optional] |
| **linked_data_source_other_aligned** | **Boolean** | False when the register list for this column splits into a different number of pieces than the number of linkages declared; the row value is then null and the whole answer is in linked_data_source_other_source. | [optional] |
| **linked_data_source_other_source** | **String** | The whole answer the other-data-source text was split from. | [optional] |
| **linkage_completeness** | **String** | The linkage completeness declared for this linkage, as typed. | [optional] |
| **linkage_completeness_aligned** | **Boolean** | False when the register list for this column splits into a different number of pieces than the number of linkages declared; the row value is then null and the whole answer is in linkage_completeness_source. | [optional] |
| **linkage_completeness_source** | **String** | The whole answer linkage completeness was split from. | [optional] |
| **linkage_variable** | **String** | The linkage variable declared for this linkage, as typed. | [optional] |
| **linkage_variable_aligned** | **Boolean** | False when the register list for this column splits into a different number of pieces than the number of linkages declared; the row value is then null and the whole answer is in linkage_variable_source. | [optional] |
| **linkage_variable_source** | **String** | The whole answer the linkage variable was split from. | [optional] |
| **linkage_strategy** | **String** | The linkage strategy declared for this linkage, as typed. | [optional] |
| **linkage_strategy_aligned** | **Boolean** | False when the register list for this column splits into a different number of pieces than the number of linkages declared; the row value is then null and the whole answer is in linkage_strategy_source. | [optional] |
| **linkage_strategy_source** | **String** | The whole answer the linkage strategy was split from. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdSourceLinkagesRow.new(
  data_source_id: null,
  linkage_ordinal: null,
  declared_linkage_count: null,
  is_pre_linked: null,
  pre_linked_answer: null,
  linked_data_source_name: null,
  linked_data_source_name_aligned: null,
  linked_data_source_name_source: null,
  names_other_sentinel: null,
  names_registered_source: null,
  linked_data_source_other: null,
  linked_data_source_other_aligned: null,
  linked_data_source_other_source: null,
  linkage_completeness: null,
  linkage_completeness_aligned: null,
  linkage_completeness_source: null,
  linkage_variable: null,
  linkage_variable_aligned: null,
  linkage_variable_source: null,
  linkage_strategy: null,
  linkage_strategy_aligned: null,
  linkage_strategy_source: null,
  restriction_tier: null,
  attribution: null
)
```

