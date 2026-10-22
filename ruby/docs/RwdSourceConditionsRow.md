# PharmaObservatory::RwdSourceConditionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **data_source_id** | **String** | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study&#39;s data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional] |
| **data_source_name_key** | **String** | The normalised name key the study register is joined to the data source on (the study register carries no data source id). | [optional] |
| **condition_id** | **String** | The condition&#39;s MeSH descriptor id (D followed by six to nine digits), the id of the conditions object type. | [optional] |
| **condition_label** | **String** | The condition&#39;s MeSH heading. | [optional] |
| **mesh_tree_category** | **String** | The MeSH tree category the condition sits under. | [optional] |
| **mesh_tree_prefix** | **String** | The MeSH tree prefix of that category. | [optional] |
| **is_disease_descriptor** | **Boolean** | Whether the MeSH descriptor is a disease descriptor. | [optional] |
| **ema_condition_terms** | **Array&lt;String&gt;** | The register&#39;s own condition terms that resolved to this MeSH descriptor. | [optional] |
| **ema_condition_term_count** | **Integer** | The number of register condition terms behind the row. | [optional] |
| **study_count** | **Integer** | How many registered studies named this source. A registered study NAMED the source; it is never evidence that the source holds data on a condition. Read it with study_register_total and studies_naming_a_registered_source, because a zero is not a finding that the source is unused. | [optional] |
| **linked_study_ids** | **Array&lt;String&gt;** | The ids of the registered studies behind the row (studies of this condition that named this source). | [optional] |
| **source_study_count** | **Integer** | How many registered studies named this source, whatever their condition: the denominator under study_count on this row. | [optional] |
| **source_condition_study_count** | **Integer** | How many distinct registered studies of this source carry any MeSH-resolved condition, whatever the condition. | [optional] |
| **source_condition_count** | **Integer** | How many distinct conditions have a row for this source. | [optional] |
| **register_source_count** | **Integer** | The number of data sources in the whole data source register at the snapshot: the denominator every source count is over. Not a count of real-world data sources in the world. | [optional] |
| **sources_with_any_study_link** | **Integer** | How many registered sources are named by at least one registered study. | [optional] |
| **sources_with_any_condition_link** | **Integer** | How many registered sources have at least one condition row (named by a study with a resolved condition). | [optional] |
| **study_register_total** | **Integer** | The number of studies in the whole study register at the snapshot. | [optional] |
| **studies_naming_a_registered_source** | **Integer** | How many registered studies name at least one data source that is itself registered; the denominator under study_count. | [optional] |
| **condition_term_total** | **Integer** | How many distinct condition terms the study register holds. | [optional] |
| **condition_term_resolved** | **Integer** | How many of those terms resolved to a MeSH descriptor. | [optional] |
| **linked_condition_count** | **Integer** | How many distinct MeSH conditions appear on any row of the dataset. | [optional] |
| **snapshot_label** | **String** | The label of the catalogue snapshot the row was read from (a date). | [optional] |
| **accessed_month** | **String** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdSourceConditionsRow.new(
  data_source_id: null,
  data_source_name_key: null,
  condition_id: null,
  condition_label: null,
  mesh_tree_category: null,
  mesh_tree_prefix: null,
  is_disease_descriptor: null,
  ema_condition_terms: null,
  ema_condition_term_count: null,
  study_count: null,
  linked_study_ids: null,
  source_study_count: null,
  source_condition_study_count: null,
  source_condition_count: null,
  register_source_count: null,
  sources_with_any_study_link: null,
  sources_with_any_condition_link: null,
  study_register_total: null,
  studies_naming_a_registered_source: null,
  condition_term_total: null,
  condition_term_resolved: null,
  linked_condition_count: null,
  snapshot_label: null,
  accessed_month: null,
  restriction_tier: null,
  attribution: null
)
```

