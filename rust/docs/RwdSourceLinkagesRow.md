# RwdSourceLinkagesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data_source_id** | Option<**String**> | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study's data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional]
**linkage_ordinal** | Option<**i32**> | The position of the linkage among the source's declared linkages, from 1. | [optional]
**declared_linkage_count** | Option<**i32**> | How many linkages the holder declared; 0 when none. | [optional]
**is_pre_linked** | Option<**bool**> | Whether the holder answered Yes to Pre linked for this linkage. | [optional]
**pre_linked_answer** | Option<**String**> | The holder's Yes or No answer to Pre linked for this linkage. | [optional]
**linked_data_source_name** | Option<**String**> | The linked data source named for this linkage, from the register's list. | [optional]
**linked_data_source_name_aligned** | Option<**bool**> | False when the register list for this column splits into a different number of pieces than the number of linkages declared; the row value is then null and the whole answer is in linked_data_source_name_source. | [optional]
**linked_data_source_name_source** | Option<**String**> | The whole comma-joined answer the linked data source name was split from, kept so an unaligned row can still be read. | [optional]
**names_other_sentinel** | Option<**bool**> | True when the linked data source name is the register's Other data source sentinel. | [optional]
**names_registered_source** | Option<**bool**> | True when the linked data source name is itself a registered data source. | [optional]
**linked_data_source_other** | Option<**String**> | The linked data source named in free text when the holder chose Other data source. | [optional]
**linked_data_source_other_aligned** | Option<**bool**> | False when the register list for this column splits into a different number of pieces than the number of linkages declared; the row value is then null and the whole answer is in linked_data_source_other_source. | [optional]
**linked_data_source_other_source** | Option<**String**> | The whole answer the other-data-source text was split from. | [optional]
**linkage_completeness** | Option<**String**> | The linkage completeness declared for this linkage, as typed. | [optional]
**linkage_completeness_aligned** | Option<**bool**> | False when the register list for this column splits into a different number of pieces than the number of linkages declared; the row value is then null and the whole answer is in linkage_completeness_source. | [optional]
**linkage_completeness_source** | Option<**String**> | The whole answer linkage completeness was split from. | [optional]
**linkage_variable** | Option<**String**> | The linkage variable declared for this linkage, as typed. | [optional]
**linkage_variable_aligned** | Option<**bool**> | False when the register list for this column splits into a different number of pieces than the number of linkages declared; the row value is then null and the whole answer is in linkage_variable_source. | [optional]
**linkage_variable_source** | Option<**String**> | The whole answer the linkage variable was split from. | [optional]
**linkage_strategy** | Option<**String**> | The linkage strategy declared for this linkage, as typed. | [optional]
**linkage_strategy_aligned** | Option<**bool**> | False when the register list for this column splits into a different number of pieces than the number of linkages declared; the row value is then null and the whole answer is in linkage_strategy_source. | [optional]
**linkage_strategy_source** | Option<**String**> | The whole answer the linkage strategy was split from. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


