# RwdSourceConditionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data_source_id** | Option<**String**> | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study's data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional]
**data_source_name_key** | Option<**String**> | The normalised name key the study register is joined to the data source on (the study register carries no data source id). | [optional]
**condition_id** | Option<**String**> | The condition's MeSH descriptor id (D followed by six to nine digits), the id of the conditions object type. | [optional]
**condition_label** | Option<**String**> | The condition's MeSH heading. | [optional]
**mesh_tree_category** | Option<**String**> | The MeSH tree category the condition sits under. | [optional]
**mesh_tree_prefix** | Option<**String**> | The MeSH tree prefix of that category. | [optional]
**is_disease_descriptor** | Option<**bool**> | Whether the MeSH descriptor is a disease descriptor. | [optional]
**ema_condition_terms** | Option<**Vec<String>**> | The register's own condition terms that resolved to this MeSH descriptor. | [optional]
**ema_condition_term_count** | Option<**i32**> | The number of register condition terms behind the row. | [optional]
**study_count** | Option<**i32**> | How many registered studies named this source. A registered study NAMED the source; it is never evidence that the source holds data on a condition. Read it with study_register_total and studies_naming_a_registered_source, because a zero is not a finding that the source is unused. | [optional]
**linked_study_ids** | Option<**Vec<String>**> | The ids of the registered studies behind the row (studies of this condition that named this source). | [optional]
**source_study_count** | Option<**i32**> | How many registered studies named this source, whatever their condition: the denominator under study_count on this row. | [optional]
**source_condition_study_count** | Option<**i32**> | How many distinct registered studies of this source carry any MeSH-resolved condition, whatever the condition. | [optional]
**source_condition_count** | Option<**i32**> | How many distinct conditions have a row for this source. | [optional]
**register_source_count** | Option<**i32**> | The number of data sources in the whole data source register at the snapshot: the denominator every source count is over. Not a count of real-world data sources in the world. | [optional]
**sources_with_any_study_link** | Option<**i32**> | How many registered sources are named by at least one registered study. | [optional]
**sources_with_any_condition_link** | Option<**i32**> | How many registered sources have at least one condition row (named by a study with a resolved condition). | [optional]
**study_register_total** | Option<**i32**> | The number of studies in the whole study register at the snapshot. | [optional]
**studies_naming_a_registered_source** | Option<**i32**> | How many registered studies name at least one data source that is itself registered; the denominator under study_count. | [optional]
**condition_term_total** | Option<**i32**> | How many distinct condition terms the study register holds. | [optional]
**condition_term_resolved** | Option<**i32**> | How many of those terms resolved to a MeSH descriptor. | [optional]
**linked_condition_count** | Option<**i32**> | How many distinct MeSH conditions appear on any row of the dataset. | [optional]
**snapshot_label** | Option<**String**> | The label of the catalogue snapshot the row was read from (a date). | [optional]
**accessed_month** | Option<**String**> | The month and year the catalogue was accessed (YYYY-MM), which EMA's site asks to be cited with its URL. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


