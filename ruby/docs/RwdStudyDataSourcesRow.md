# PharmaObservatory::RwdStudyDataSourcesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **study_id** | **String** | The register&#39;s own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional] |
| **item_ordinal** | **Integer** | The position of the item in the register&#39;s own list, from 1. | [optional] |
| **data_source_name** | **String** | The data source name as the study named it, from the register&#39;s Data source(s) list. | [optional] |
| **data_source_name_key** | **String** | The normalised name key the study register is joined to the data source on (the study register carries no data source id). | [optional] |
| **matched_data_source_register** | **Boolean** | Whether the name matches a registered data source. | [optional] |
| **data_source_id** | **String** | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study&#39;s data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional] |
| **data_source_ids** | **Array&lt;String&gt;** | Every registered data source id the name matches (more than one when two entries share a name). | [optional] |
| **register_row_count** | **Integer** | How many register entries share the name. | [optional] |
| **study_status** | **String** | The study&#39;s status as the register words it (Planned, Ongoing, Finalised, and so on). | [optional] |
| **mandate_stratum** | **String** | mandated when a regulator required the study (a post-authorisation obligation), voluntary when the registrant says it did not, unknown when unanswered. Delivery timeliness is only comparable within a stratum. | [optional] |
| **required_by_regulator** | **String** | The registrant&#39;s answer to whether a regulatory body required the study (Yes, No or Unknown). | [optional] |
| **study_start_actual_on** | **Date** | Actual study start date. | [optional] |
| **started_in_last_5_years** | **Boolean** | Whether the study started within five years of the snapshot. | [optional] |
| **condition_count** | **Integer** | How many distinct condition terms the study names. | [optional] |
| **conducting_institution_count** | **Integer** | How many conducting institutions the study names. | [optional] |
| **snapshot_label** | **String** | The label of the catalogue snapshot the row was read from (a date). | [optional] |
| **accessed_month** | **String** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdStudyDataSourcesRow.new(
  study_id: null,
  item_ordinal: null,
  data_source_name: null,
  data_source_name_key: null,
  matched_data_source_register: null,
  data_source_id: null,
  data_source_ids: null,
  register_row_count: null,
  study_status: null,
  mandate_stratum: null,
  required_by_regulator: null,
  study_start_actual_on: null,
  started_in_last_5_years: null,
  condition_count: null,
  conducting_institution_count: null,
  snapshot_label: null,
  accessed_month: null,
  restriction_tier: null,
  attribution: null
)
```

