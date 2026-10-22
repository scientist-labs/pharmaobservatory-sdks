# RwdSourceTypesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data_source_id** | Option<**String**> | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study's data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional]
**item_ordinal** | Option<**i32**> | The position of the item in the register's own list, from 1. | [optional]
**type_value** | Option<**String**> | The declared data source type, as typed after the register's list is split. | [optional]
**type_key** | Option<**String**> | A stable lowercase key for the data source type; the filter value. | [optional]
**type_label** | Option<**String**> | The register's most common spelling of the type. | [optional]
**is_other_sentinel** | Option<**bool**> | True when the item is the register's Other sentinel; the holder's free text for it is on the source row. | [optional]
**source_value** | Option<**String**> | The whole register answer the item was split from. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


