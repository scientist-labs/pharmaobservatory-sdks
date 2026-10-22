# RwdSourceEtlMappingsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data_source_id** | Option<**String**> | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study's data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional]
**mapping_ordinal** | Option<**i32**> | The position of the mapping among the source's declared CDM mappings, from 1. | [optional]
**declared_mapping_count** | Option<**i32**> | How many CDM mappings the holder declared, after register join artefacts are removed; 0 when the question was never answered. | [optional]
**etl_status** | Option<**String**> | The extract-transform-load (ETL) status of the CDM mapping as typed; one comma-joined answer per declared mapping (rwd_source_etl_mappings has one row per mapping). | [optional]
**etl_status_aligned** | Option<**bool**> | False when the register list for this column splits into a different number of pieces than the number of mappings declared; the row value is then null and the whole answer is in etl_status_source. | [optional]
**etl_status_source** | Option<**String**> | The whole answer the ETL status was split from. | [optional]
**etl_cdm_version** | Option<**String**> | The CDM version(s) declared for the ETL, comma-joined as typed. | [optional]
**etl_cdm_version_aligned** | Option<**bool**> | False when the register list for this column splits into a different number of pieces than the number of mappings declared; the row value is then null and the whole answer is in etl_cdm_version_source. | [optional]
**etl_cdm_version_source** | Option<**String**> | The whole answer the CDM version was split from. | [optional]
**etl_frequency** | Option<**String**> | The ETL frequency declared, comma-joined as typed. | [optional]
**etl_frequency_aligned** | Option<**bool**> | False when the register list for this column splits into a different number of pieces than the number of mappings declared; the row value is then null and the whole answer is in etl_frequency_source. | [optional]
**etl_frequency_source** | Option<**String**> | The whole answer the ETL frequency was split from. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


