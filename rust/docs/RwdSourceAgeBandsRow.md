# RwdSourceAgeBandsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data_source_id** | Option<**String**> | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study's data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional]
**band_ordinal** | Option<**i32**> | The position of the age band in the source's declared list, from 1. | [optional]
**age_band** | Option<**String**> | The age band label as the holder declared it. Bands can overlap (see is_nested_scheme); never sum them. | [optional]
**population_size** | Option<**i64**> | The population size the holder declared for the source (all registered persons). | [optional]
**active_population_size** | Option<**i64**> | The active population size the holder declared for the source. | [optional]
**has_population_breakdown** | Option<**bool**> | Whether the holder published a population size breakdown by age band at all. | [optional]
**has_active_population_breakdown** | Option<**bool**> | Whether the holder published an active population breakdown by age band at all. | [optional]
**is_nested_scheme** | Option<**bool**> | True when the source's bands overlap, so a sum over bands double counts. | [optional]
**snapshot_label** | Option<**String**> | The label of the catalogue snapshot the row was read from (a date). | [optional]
**accessed_month** | Option<**String**> | The month and year the catalogue was accessed (YYYY-MM), which EMA's site asks to be cited with its URL. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


