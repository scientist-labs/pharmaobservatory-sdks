# PharmaObservatory::RwdRegisterProvenanceRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **register_key** | **String** | Which of the five exports the row concerns: study, data_source_human, data_source_veterinary, institution or network. The entry key of the rwd_registers object type. | [optional] |
| **register_label** | **String** | The register&#39;s plain-words label. | [optional] |
| **register_ordinal** | **Integer** | The register&#39;s position in the lane&#39;s order. | [optional] |
| **provenance_source** | **String** | The landed source name of the export. | [optional] |
| **export_is_empty_at_source** | **Boolean** | True when the export held no rows at the source. | [optional] |
| **landed_rows** | **Integer** | How many rows the export landed. | [optional] |
| **build_status** | **String** | The status of the latest landing run. | [optional] |
| **build_snapshot_id** | **String** | The snapshot id of the latest landing run. | [optional] |
| **build_row_count** | **Integer** | The row count of the latest landing run. | [optional] |
| **build_started_at** | **Time** | When the latest landing run started (UTC). | [optional] |
| **build_finished_at** | **Time** | When the latest landing run finished (UTC). | [optional] |
| **last_completed_snapshot_id** | **String** | The snapshot id of the last completed landing. | [optional] |
| **last_completed_row_count** | **Integer** | The row count of the last completed landing. | [optional] |
| **last_completed_at** | **Time** | When the last completed landing finished (UTC). | [optional] |
| **registrant_clock_asof** | **Date** | The cutoff of the registrant clock: the snapshot date less twelve months; a record last updated before it counts as older than twelve months. | [optional] |
| **register_publishes_update_date** | **Boolean** | Whether the export carries an update date column. | [optional] |
| **records_with_an_update_date** | **Integer** | How many records carry an update date. | [optional] |
| **records_updated_over_twelve_months** | **Integer** | How many records were last updated more than twelve months before the snapshot. | [optional] |
| **share_updated_over_twelve_months** | **Float** | The share of records with an update date that are older than twelve months. | [optional] |
| **register_publishes_outdated_flag** | **Boolean** | Whether the export carries EMA&#39;s Outdated flag. | [optional] |
| **records_ema_flags_outdated** | **Integer** | How many records EMA flags Outdated. | [optional] |
| **records_outdated_flag_blank** | **Integer** | How many records have no Outdated flag value. | [optional] |
| **columns_landed** | **Integer** | How many columns of the export this lane lands. | [optional] |
| **columns_withheld** | **Integer** | How many columns of the export the lane withholds (contact details of named individuals). | [optional] |
| **snapshot_label** | **String** | The label of the catalogue snapshot the row was read from (a date). | [optional] |
| **snapshot_date** | **Date** | The date of the catalogue snapshot. | [optional] |
| **accessed_month** | **String** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdRegisterProvenanceRow.new(
  register_key: null,
  register_label: null,
  register_ordinal: null,
  provenance_source: null,
  export_is_empty_at_source: null,
  landed_rows: null,
  build_status: null,
  build_snapshot_id: null,
  build_row_count: null,
  build_started_at: null,
  build_finished_at: null,
  last_completed_snapshot_id: null,
  last_completed_row_count: null,
  last_completed_at: null,
  registrant_clock_asof: null,
  register_publishes_update_date: null,
  records_with_an_update_date: null,
  records_updated_over_twelve_months: null,
  share_updated_over_twelve_months: null,
  register_publishes_outdated_flag: null,
  records_ema_flags_outdated: null,
  records_outdated_flag_blank: null,
  columns_landed: null,
  columns_withheld: null,
  snapshot_label: null,
  snapshot_date: null,
  accessed_month: null,
  restriction_tier: null,
  attribution: null
)
```

