# PharmaObservatory::RwdSourceVocabulariesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **data_source_id** | **String** | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study&#39;s data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional] |
| **capability_key** | **String** | Which capability the coding scheme answers: exposure (prescriptions, dispensing, medicinal product, indication) or outcomes (diagnosis, cause of death, procedures, and so on). | [optional] |
| **capability_label** | **String** | The capability&#39;s label as the register words it. | [optional] |
| **source_column** | **String** | The register column the vocabulary item came from. | [optional] |
| **item_ordinal** | **Integer** | The position of the item in the register&#39;s own list, from 1. | [optional] |
| **vocabulary_value** | **String** | The declared coding scheme as typed. | [optional] |
| **vocabulary_key** | **String** | A stable lowercase key for the coding scheme; the filter value. | [optional] |
| **vocabulary_label** | **String** | The register&#39;s most common spelling of the coding scheme. | [optional] |
| **is_free_text** | **Boolean** | True when the item is a free-text answer to the register&#39;s other-vocabulary question, not a coding scheme from its list. | [optional] |
| **is_other_sentinel** | **Boolean** | True when the item is the register&#39;s Other sentinel; the holder&#39;s free text for it is on the source row. | [optional] |
| **source_value** | **String** | The whole register answer the item was split from. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdSourceVocabulariesRow.new(
  data_source_id: null,
  capability_key: null,
  capability_label: null,
  source_column: null,
  item_ordinal: null,
  vocabulary_value: null,
  vocabulary_key: null,
  vocabulary_label: null,
  is_free_text: null,
  is_other_sentinel: null,
  source_value: null,
  restriction_tier: null,
  attribution: null
)
```

