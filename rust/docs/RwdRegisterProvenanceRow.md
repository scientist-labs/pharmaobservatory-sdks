# RwdRegisterProvenanceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**register_key** | Option<**String**> | Which of the five exports the row concerns: study, data_source_human, data_source_veterinary, institution or network. The entry key of the rwd_registers object type. | [optional]
**register_label** | Option<**String**> | The register's plain-words label. | [optional]
**register_ordinal** | Option<**i32**> | The register's position in the lane's order. | [optional]
**provenance_source** | Option<**String**> | The landed source name of the export. | [optional]
**export_is_empty_at_source** | Option<**bool**> | True when the export held no rows at the source. | [optional]
**landed_rows** | Option<**i64**> | How many rows the export landed. | [optional]
**build_status** | Option<**String**> | The status of the latest landing run. | [optional]
**build_snapshot_id** | Option<**String**> | The snapshot id of the latest landing run. | [optional]
**build_row_count** | Option<**i64**> | The row count of the latest landing run. | [optional]
**build_started_at** | Option<**String**> | When the latest landing run started (UTC). | [optional]
**build_finished_at** | Option<**String**> | When the latest landing run finished (UTC). | [optional]
**last_completed_snapshot_id** | Option<**String**> | The snapshot id of the last completed landing. | [optional]
**last_completed_row_count** | Option<**i64**> | The row count of the last completed landing. | [optional]
**last_completed_at** | Option<**String**> | When the last completed landing finished (UTC). | [optional]
**registrant_clock_asof** | Option<[**String**](string.md)> | The cutoff of the registrant clock: the snapshot date less twelve months; a record last updated before it counts as older than twelve months. | [optional]
**register_publishes_update_date** | Option<**bool**> | Whether the export carries an update date column. | [optional]
**records_with_an_update_date** | Option<**i64**> | How many records carry an update date. | [optional]
**records_updated_over_twelve_months** | Option<**i64**> | How many records were last updated more than twelve months before the snapshot. | [optional]
**share_updated_over_twelve_months** | Option<**f64**> | The share of records with an update date that are older than twelve months. | [optional]
**register_publishes_outdated_flag** | Option<**bool**> | Whether the export carries EMA's Outdated flag. | [optional]
**records_ema_flags_outdated** | Option<**i64**> | How many records EMA flags Outdated. | [optional]
**records_outdated_flag_blank** | Option<**i64**> | How many records have no Outdated flag value. | [optional]
**columns_landed** | Option<**i64**> | How many columns of the export this lane lands. | [optional]
**columns_withheld** | Option<**i64**> | How many columns of the export the lane withholds (contact details of named individuals). | [optional]
**snapshot_label** | Option<**String**> | The label of the catalogue snapshot the row was read from (a date). | [optional]
**snapshot_date** | Option<[**String**](string.md)> | The date of the catalogue snapshot. | [optional]
**accessed_month** | Option<**String**> | The month and year the catalogue was accessed (YYYY-MM), which EMA's site asks to be cited with its URL. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


