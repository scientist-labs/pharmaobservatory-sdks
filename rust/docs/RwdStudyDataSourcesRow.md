# RwdStudyDataSourcesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**study_id** | Option<**String**> | The register's own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional]
**item_ordinal** | Option<**i32**> | The position of the item in the register's own list, from 1. | [optional]
**data_source_name** | Option<**String**> | The data source name as the study named it, from the register's Data source(s) list. | [optional]
**data_source_name_key** | Option<**String**> | The normalised name key the study register is joined to the data source on (the study register carries no data source id). | [optional]
**matched_data_source_register** | Option<**bool**> | Whether the name matches a registered data source. | [optional]
**data_source_id** | Option<**String**> | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study's data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional]
**data_source_ids** | Option<**Vec<String>**> | Every registered data source id the name matches (more than one when two entries share a name). | [optional]
**register_row_count** | Option<**i32**> | How many register entries share the name. | [optional]
**study_status** | Option<**String**> | The study's status as the register words it (Planned, Ongoing, Finalised, and so on). | [optional]
**mandate_stratum** | Option<**String**> | mandated when a regulator required the study (a post-authorisation obligation), voluntary when the registrant says it did not, unknown when unanswered. Delivery timeliness is only comparable within a stratum. | [optional]
**required_by_regulator** | Option<**String**> | The registrant's answer to whether a regulatory body required the study (Yes, No or Unknown). | [optional]
**study_start_actual_on** | Option<[**String**](string.md)> | Actual study start date. | [optional]
**started_in_last_5_years** | Option<**bool**> | Whether the study started within five years of the snapshot. | [optional]
**condition_count** | Option<**i32**> | How many distinct condition terms the study names. | [optional]
**conducting_institution_count** | Option<**i32**> | How many conducting institutions the study names. | [optional]
**snapshot_label** | Option<**String**> | The label of the catalogue snapshot the row was read from (a date). | [optional]
**accessed_month** | Option<**String**> | The month and year the catalogue was accessed (YYYY-MM), which EMA's site asks to be cited with its URL. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


