# RwdConductorFootprintRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**conductor_name_key** | Option<**String**> | The normalised institution name key the conductor rollup is keyed on; every register institution whose name folds to it is one conductor. | [optional]
**conductor_id** | Option<**String**> | The conductor slug: the register's normalised institution name key with spaces as hyphens, the id in the web page address (/rwd-conductors/<id>). The entry key of the rwd_conductors object type; injective because no name key carries a hyphen. | [optional]
**conductor_name** | Option<**String**> | The conductor's display name: one register spelling of the key. | [optional]
**dimension** | Option<**String**> | What the footprint value is: delivered_country (a country of a conducting study) or delivered_data_source (a registered data source a conducting study names). | [optional]
**value** | Option<**String**> | The country or data source name the row counts. | [optional]
**study_count** | Option<**i32**> | Distinct conducting studies of the conductor for this country or data source. | [optional]
**conductor_study_count** | Option<**i32**> | Distinct conducting studies of the conductor, the denominator beside study_count. | [optional]
**data_source_ids** | Option<**Vec<String>**> | The register Data source IDs of the registered sources with this name (delivered_data_source rows only; null for a country). | [optional]
**snapshot_label** | Option<**String**> | The label of the catalogue snapshot the row was read from (a date). | [optional]
**accessed_month** | Option<**String**> | The month and year the catalogue was accessed (YYYY-MM), which EMA's site asks to be cited with its URL. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


